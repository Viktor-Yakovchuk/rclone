// Package linkbox provides an interface to the linkbox.to Cloud storage system.
package linkbox

import (
	"bytes"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configmap"
	"github.com/rclone/rclone/fs/config/configstruct"
	"github.com/rclone/rclone/fs/fserrors"
	"github.com/rclone/rclone/fs/fshttp"
	"github.com/rclone/rclone/fs/hash"
	"github.com/rclone/rclone/lib/pacer"
	"github.com/rclone/rclone/lib/rest"
)

const (
	retriesAmmount     = 2
	maxEntitiesPerPage = 64
	minSleep           = 400 * time.Millisecond
	maxSleep           = 5 * time.Minute
	pacerBurst         = 1
	linkboxAPIURL      = "https://www.linkbox.to/api/open/"
)

func init() {
	fsi := &fs.RegInfo{
		Name:        "linkbox",
		Description: "Linkbox",
		NewFs:       NewFs,
		Options: []fs.Option{{
			Name:     "token",
			Help:     "Token from https://www.linkbox.to/admin/account",
			Required: true,
		}},
	}
	fs.Register(fsi)
}

// Options defines the configuration for this backend
type Options struct {
	Token string `config:"token"`
}

// Fs stores the interface to the remote Linkbox files
type Fs struct {
	name     string
	root     string
	opt      Options        // options for this backend
	features *fs.Features   // optional features
	ci       *fs.ConfigInfo // global config
	srv      *rest.Client   // the connection to the server
	pacer    *fs.Pacer
}

// Object is a remote object that has been stat'd (so it exists, but is not necessarily open for reading)
type Object struct {
	fs          *Fs
	remote      string
	size        int64
	modTime     time.Time
	contentType string
	fullURL     string
	pid         int
	isDir       bool
	id          string
}

// NewFs creates a new Fs object from the name and root. It connects to
// the host specified in the config file.
func NewFs(ctx context.Context, name, root string, m configmap.Mapper) (fs.Fs, error) {
	// Parse config into Options struct
	fs.Infof(nil, "NewFs:  name = {%s}, root = {%s}", name, root)

	opt := new(Options)
	err := configstruct.Set(m, opt)
	if err != nil {
		return nil, err
	}

	ci := fs.GetConfig(ctx)

	f := &Fs{
		name:  name,
		opt:   *opt,
		ci:    ci,
		srv:   rest.NewClient(fshttp.NewClient(ctx)),
		pacer: fs.NewPacer(ctx, pacer.NewLinkbox()),
	}

	f.pacer.SetRetries(retriesAmmount)

	f.features = (&fs.Features{
		CanHaveEmptyDirectories: true,
		WriteMetadata:           false,
	}).Fill(ctx, f)

	// Check to see if the root actually an existing file
	remote := path.Base(root)
	f.root = path.Dir(root)
	if f.root == "." {
		f.root = ""
	}
	_, err = f.NewObject(ctx, remote)
	if err != nil {
		if errors.Is(err, fs.ErrorObjectNotFound) || errors.Is(err, fs.ErrorNotAFile) || errors.Is(err, fs.ErrorIsDir) {
			// File doesn't exist so return old f
			fs.Debugf(f, "create new fs with root = {%s}", root)
			f.root = root
			return f, nil
		}
		return nil, err
	}
	// return an error with an fs which points to the parent
	return f, fs.ErrorIsFile
}

type Entity struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Ctime  int64  `json:"ctime"`
	Size   int    `json:"size"`
	ID     int    `json:"id"`
	Pid    int    `json:"pid"`
	ItemID string `json:"item_id"`
}
type Data struct {
	Entities []Entity `json:"list"`
}
type FileSearchRes struct {
	SearchData Data   `json:"data"`
	Status     int    `json:"status"`
	Message    string `json:"msg"`
}

func getIdByDir(f *Fs, ctx context.Context, dir string) (int, error) {
	var pid int
	var err error
	err = f.pacer.Call(func() (bool, error) {
		pid, err = _getIdByDir(f, ctx, dir)
		return f.shouldRetry(ctx, err)
	})
	fs.Debugf(f, "getIdByDir: pid = {%d}, dir = {%s}, err = {%#v}", pid, dir, err)

	if fserrors.IsRetryError(err) {
		fs.Debugf(f, "getIdByDir IsRetryError: pid = {%d}, dir = {%s}, err = {%#v}", pid, dir, err)
		err = fs.ErrorDirNotFound
	}

	return pid, err
}

func _getIdByDir(f *Fs, ctx context.Context, dir string) (int, error) {
	if dir == "" || dir == "/" {
		return 0, nil // we assume that it is root directory
	}

	path := strings.TrimPrefix(dir, "/")
	dirs := strings.Split(path, "/")
	pid := 0

	for level, tdir := range dirs {
		pageNumber := 0
		numberOfEntities := maxEntitiesPerPage
		fs.Debugf(f, "getIdByDir tdir = {%s} level = {%d}", tdir, level)

		for numberOfEntities == maxEntitiesPerPage {
			pageNumber++
			opts := makeSearchQuery("", pid, f.opt.Token, pageNumber)
			responseResult := FileSearchRes{}
			err := getUnmarshaledResponse(ctx, f, opts, &responseResult)
			numberOfEntities = len(responseResult.SearchData.Entities)

			if err != nil {
				fs.Debugf(f, "getIdByDir 1 if dir = {%s} path = {%s}", dir, path)
				return 0, err
			}

			if len(responseResult.SearchData.Entities) == 0 {
				fs.Debugf(f, "getIdByDir 2 if dir = {%s} path = {%s}", dir, path)
				return 0, fs.ErrorDirNotFound
			}

			for _, entity := range responseResult.SearchData.Entities {
				fs.Debugf(f, "getIdByDir for pid = {%d} entity.Name = {%s} level = {%d}", pid, entity.Name, level)
				if entity.Pid == pid && (entity.Type == "dir" || entity.Type == "sdir") && entity.Name == tdir {
					pid = entity.ID
					fs.Debugf(f, "getIdByDir tdir = {%s} level = {%d} pid = {%d}", tdir, level, pid)

					if level == len(dirs)-1 {
						fs.Debugf(f, "getIdByDir return tdir = {%s} level = {%d} pid = {%d}", tdir, level, pid)
						return pid, nil
					}
				}
			}

			if pageNumber > 100000 {
				return 0, fmt.Errorf("too many results")
			}

		}
	}

	fs.Debugf(f, "getIdByDir fs.ErrorDirNotFound dir = {%s} path = {%s}", dir, path)

	return 0, fs.ErrorDirNotFound
}

func getUnmarshaledResponse(ctx context.Context, f *Fs, opts *rest.Opts, result interface{}) error {
	err := f.pacer.Call(func() (bool, error) {
		_, err := f.srv.CallJSON(ctx, opts, nil, &result)
		return f.shouldRetry(ctx, err)
	})

	return err
}

func makeSearchQuery(name string, pid int, token string, pageNubmer int) *rest.Opts {
	return &rest.Opts{
		Method:  "GET",
		RootURL: linkboxAPIURL,
		Path:    "file_search",
		Parameters: url.Values{
			"token":    {token},
			"name":     {name},
			"pid":      {strconv.Itoa(pid)},
			"pageNo":   {strconv.Itoa(pageNubmer)},
			"pageSize": {strconv.Itoa(maxEntitiesPerPage)},
		},
	}
}

func parse(f *Fs, ctx context.Context, dir string) ([]*Object, error) {
	var responseResult FileSearchRes
	var files []*Object
	var numberOfEntities int

	fullPath := path.Join(f.root, dir)
	fullPath = strings.TrimPrefix(fullPath, "/")

	pid, err := getIdByDir(f, ctx, fullPath)

	fs.Debugf(f, "parse dir = {%s} fullPath = {%s} pid = {%d}", dir, fullPath, pid)

	if err != nil {
		fs.Debugf(f, "parse failed dir = {%s} fullPath = {%s} pid = {%d} err = {%s}", dir, fullPath, pid, err)

		return nil, err
	}

	pageNumber := 0
	numberOfEntities = maxEntitiesPerPage

	for numberOfEntities == maxEntitiesPerPage {
		pageNumber++
		opts := makeSearchQuery("", pid, f.opt.Token, pageNumber)

		responseResult = FileSearchRes{}
		err = getUnmarshaledResponse(ctx, f, opts, &responseResult)
		if err != nil {
			return nil, fmt.Errorf("parsing failed: %w", err)
		}

		if responseResult.Status != 1 {
			return nil, fmt.Errorf("parsing failed: %s", responseResult.Message)
		}

		numberOfEntities = len(responseResult.SearchData.Entities)

		for _, entity := range responseResult.SearchData.Entities {
			if entity.Pid != pid {
				fs.Debugf(f, "parse error with entity.Name {%s} dir {%s}", entity.Name, dir)
			}
			file := &Object{
				fs:          f,
				remote:      entity.Name,
				modTime:     time.Unix(entity.Ctime, 0),
				contentType: entity.Type,
				size:        int64(entity.Size),
				fullURL:     entity.URL,
				isDir:       entity.Type == "dir" || entity.Type == "sdir",
				id:          entity.ItemID,
				pid:         entity.Pid,
			}

			files = append(files, file)
		}

		if pageNumber > 100000 {
			return files, fmt.Errorf("too many results")
		}

	}

	return files, nil
}

func splitDirAndName(remote string) (dir string, name string) {
	lastSlashPosition := strings.LastIndex(remote, "/")
	if lastSlashPosition == -1 {
		dir = ""
		name = remote
	} else {
		dir = remote[:lastSlashPosition]
		name = remote[lastSlashPosition+1:]
	}

	fs.Debugf(nil, "splitDirAndName remote = {%s}, dir = {%s}, name = {%s}", remote, dir, name)

	return
}

// List the objects and directories in dir into entries.  The
// entries can be returned in any order but should be for a
// complete directory.
//
// dir should be "" to list the root, and should not have
// trailing slashes.
//
// This should return ErrDirNotFound if the directory isn't
// found.
func (f *Fs) List(ctx context.Context, dir string) (entries fs.DirEntries, err error) {
	fs.Debugf(f, "List method dir = {%s}", dir)

	objects, err := parse(f, ctx, dir)
	if err != nil {
		return nil, err
	}

	for _, obj := range objects {
		prefix := ""
		if dir != "" {
			prefix = dir + "/"
		}

		if obj.isDir {
			entries = append(entries, fs.NewDir(prefix+obj.remote, obj.modTime))
		} else {
			obj.remote = prefix + obj.remote
			entries = append(entries, obj)
		}
	}

	return entries, nil
}

func getObject(ctx context.Context, f *Fs, name string, pid int, token string) (Entity, error) {
	var err error
	var entity Entity
	err = f.pacer.Call(func() (bool, error) {
		entity, err = _getObject(ctx, f, name, pid, token)
		return f.shouldRetry(ctx, err)
	})
	fs.Debugf(f, "getObject: name = {%s}, pid = {%d}, err = {%#v}", name, pid, err)

	if fserrors.IsRetryError(err) {
		fs.Debugf(f, "getObject IsRetryError: name = {%s}, pid = {%d}, err = {%#v}", name, pid, err)

		err = fs.ErrorObjectNotFound
	}

	return entity, err
}

func _getObject(ctx context.Context, f *Fs, name string, pid int, token string) (Entity, error) {
	pageNumber := 0
	numberOfEntities := maxEntitiesPerPage

	for numberOfEntities == maxEntitiesPerPage {
		pageNumber++
		opts := makeSearchQuery("", pid, token, pageNumber)

		searchResponse := FileSearchRes{}
		err := getUnmarshaledResponse(ctx, f, opts, &searchResponse)
		if err != nil {
			return Entity{}, fmt.Errorf("unable to create new object: %w", err)
		}
		if searchResponse.Status != 1 {
			return Entity{}, fmt.Errorf("unable to create new object: %s", searchResponse.Message)
		}
		numberOfEntities = len(searchResponse.SearchData.Entities)

		fs.Debugf(f, "getObject numberOfEntities {%d} name {%s}", numberOfEntities, name)

		for _, entity := range searchResponse.SearchData.Entities {
			fs.Debugf(f, "getObject entity.Name {%s} name {%s}", entity.Name, name)
			if entity.Pid == pid && entity.Name == name {
				fs.Debugf(f, "getObject found entity.Name {%s} name {%s}", entity.Name, name)
				if entity.Type == "dir" || entity.Type == "sdir" {
					return Entity{}, fs.ErrorIsDir
				}
				return entity, nil
			}
		}

		if pageNumber > 100000 {
			return Entity{}, fmt.Errorf("too many results")
		}
	}

	return Entity{}, fs.ErrorObjectNotFound
}

// NewObject finds the Object at remote.  If it can't be found
// it returns the error ErrorObjectNotFound.
//
// If remote points to a directory then it should return
// ErrorIsDir if possible without doing any extra work,
// otherwise ErrorObjectNotFound.
func (f *Fs) NewObject(ctx context.Context, remote string) (fs.Object, error) {
	var newObject Entity
	var dir, name string

	fs.Infof(f, "NewObject: remote = {%s}", remote)

	fullPath := path.Join(f.root, remote)
	dir, name = splitDirAndName(fullPath)

	dirId, err := getIdByDir(f, ctx, dir)
	if err != nil {
		fs.Debugf(f, "NewObject %v", err)
		return nil, fs.ErrorObjectNotFound
	}

	newObject, err = getObject(ctx, f, name, dirId, f.opt.Token)
	if err != nil {
		fs.Debugf(f, "NewObject getObject error = {%s}", err)

		return nil, err
	}

	if newObject == (Entity{}) {
		fs.Debugf(f, "NewObject entity is empty")
		return nil, fs.ErrorObjectNotFound
	}

	return &Object{
		fs:      f,
		remote:  name,
		modTime: time.Unix(newObject.Ctime, 0),
		fullURL: newObject.URL,
		size:    int64(newObject.Size),
		id:      newObject.ItemID,
		pid:     newObject.Pid,
	}, nil
}

// Mkdir makes the directory (container, bucket)
//
// Shouldn't return an error if it already exists
func (f *Fs) Mkdir(ctx context.Context, dir string) error {
	fs.Infof(f, "mkdir - {%s}", dir)
	var pdir, name string

	fullPath := path.Join(f.root, dir)
	if fullPath == "" {
		return nil
	}

	fullPath = strings.TrimPrefix(fullPath, "/")

	dirs := strings.Split(fullPath, "/")
	dirs = append([]string{""}, dirs...)

	for i, dirName := range dirs {
		pdir = path.Join(pdir, dirName)
		name = dirs[i+1]
		pid, err := getIdByDir(f, ctx, pdir)
		if err != nil {
			return err
		}

		opts := &rest.Opts{
			Method:  "GET",
			RootURL: linkboxAPIURL,
			Path:    "folder_create",
			Parameters: url.Values{
				"token":       {f.opt.Token},
				"name":        {name},
				"pid":         {strconv.Itoa(pid)},
				"isShare":     {"0"},
				"canInvite":   {"1"},
				"canShare":    {"1"},
				"withBodyImg": {"1"},
				"desc":        {""},
			},
		}

		response := getResponse{}

		err = getUnmarshaledResponse(ctx, f, opts, &response)
		if err != nil {
			return fmt.Errorf("err in response")
		}

		if i+1 == len(dirs)-1 {
			break
		}

		// response status 1501 means that directory already exists
		if response.Status != 1 && response.Status != 1501 {
			return fmt.Errorf("could not create dir[%s]: %s", dir, response.Message)
		}

	}
	return nil
}

func (f *Fs) purgeCheck(ctx context.Context, dir string, check bool) error {
	fullPath := path.Join(f.root, dir)

	if fullPath == "" {
		return fs.ErrorDirNotFound
	}

	fullPath = strings.TrimPrefix(fullPath, "/")
	dirIds, err := getIdByDir(f, ctx, fullPath)

	if err != nil {
		return err
	}

	entries, err := f.List(ctx, dir)
	if err != nil {
		return err
	}

	if len(entries) != 0 && check {
		return fs.ErrorDirectoryNotEmpty
	}

	opts := &rest.Opts{
		Method:  "GET",
		RootURL: linkboxAPIURL,
		Path:    "folder_del",
		Parameters: url.Values{
			"token":  {f.opt.Token},
			"dirIds": {strconv.Itoa(dirIds)},
		},
	}

	response := getResponse{}
	err = getUnmarshaledResponse(ctx, f, opts, &response)

	if err != nil {
		return fmt.Errorf("err in response")
	}

	if response.Status != 1 {
		// it can be some different error, but Linkbox
		// returns very few statuses
		return fs.ErrorDirExists
	}
	return nil
}

// Rmdir removes the directory (container, bucket) if empty
//
// Return an error if it doesn't exist or isn't empty
func (f *Fs) Rmdir(ctx context.Context, dir string) error {
	fs.Infof(f, "rmdir - {%s}", dir)
	return f.purgeCheck(ctx, dir, true)
}

func (o *Object) SetModTime(ctx context.Context, modTime time.Time) error {
	return fs.ErrorCantSetModTime
}

// Open opens the file for read.  Call Close() on the returned io.ReadCloser
func (o *Object) Open(ctx context.Context, options ...fs.OpenOption) (io.ReadCloser, error) {
	fs.Infof(o, "Open - {%s}", o.Remote())

	var res *http.Response
	downloadURL := o.fullURL
	if downloadURL == "" {
		_, name := splitDirAndName(o.Remote())
		newObject, err := getObject(ctx, o.fs, name, o.pid, o.fs.opt.Token)
		if err != nil {
			return nil, err
		}
		if newObject == (Entity{}) {
			fs.Debugf(o.fs, "Open entity is empty: name = {%s}", name)
			return nil, fs.ErrorObjectNotFound
		}

		downloadURL = newObject.URL
	}

	opts := &rest.Opts{
		Method:  "GET",
		RootURL: downloadURL,
		Options: options,
	}

	err := o.fs.pacer.Call(func() (bool, error) {
		var err error
		res, err = o.fs.srv.Call(ctx, opts)
		return o.fs.shouldRetry(ctx, err)
	})

	if err != nil {
		return nil, fmt.Errorf("Open failed: %w", err)
	}

	return res.Body, nil
}

// Update in to the object with the modTime given of the given size
//
// When called from outside an Fs by rclone, src.Size() will always be >= 0.
// But for unknown-sized objects (indicated by src.Size() == -1), Upload should either
// return an error or update the object properly (rather than e.g. calling panic).
func (o *Object) Update(ctx context.Context, in io.Reader, src fs.ObjectInfo, options ...fs.OpenOption) error {
	fs.Infof(o, "Update: Remote = {%s}", o.Remote())

	if src.Size() == 0 {
		return fs.ErrorCantUploadEmptyFiles
	}

	remote := o.Remote()
	tmpObject, err := o.fs.NewObject(ctx, remote)

	if err == nil {
		fs.Debugf(o.fs, "Update: removing old file")
		_ = tmpObject.Remove(ctx)
	}

	h := md5.New()
	first10m := io.LimitReader(in, 10_485_760)
	first10mBytes, err := io.ReadAll(first10m)
	if err != nil {
		fs.Debugf(o.fs, "Update err in reading file: %v", err)
		return err
	}

	if _, err := io.Copy(h, bytes.NewReader(first10mBytes)); err != nil {
		fs.Debugf(o.fs, "Update err in copying %v", err)
		return err
	}

	// get upload authorization (step 1)
	opts := &rest.Opts{
		Method:  "GET",
		RootURL: linkboxAPIURL,
		Path:    "get_upload_url",
		Parameters: url.Values{
			"token":           {o.fs.opt.Token},
			"fileMd5ofPre10m": {fmt.Sprintf("%x", h.Sum(nil))},
			"fileSize":        {strconv.FormatInt(src.Size(), 10)},
		},
	}

	getFistStepResult := getUploadUrlResponse{}
	err = getUnmarshaledResponse(ctx, o.fs, opts, &getFistStepResult)
	if err != nil {
		fs.Debugf(o.fs, "Update err in unparshaling response: %v", err)
		return err
	}

	switch getFistStepResult.Status {
	case 1:
		// upload file using link from first step
		var res *http.Response

		file := io.MultiReader(bytes.NewReader(first10mBytes), in)

		opts := &rest.Opts{
			Method:  "PUT",
			RootURL: getFistStepResult.Data.SignUrl,
			Body:    file,
		}

		err = o.fs.pacer.CallNoRetry(func() (bool, error) {
			res, err = o.fs.srv.Call(ctx, opts)
			return o.fs.shouldRetry(ctx, err)
		})

		if err != nil {
			fs.Debugf(o.fs, "Update err in uploading file: %v", err)
			return err
		}

		_, err = io.ReadAll(res.Body)
		if err != nil {
			fs.Debugf(o.fs, "Update err in reading response: %v", err)
			return err
		}

	case 600:
		// Status means that we don't need to upload file
		// We need only to make second step
	default:
		return fmt.Errorf("get unexpected message from Linkbox: %s", getFistStepResult.Message)
	}

	fullPath := path.Join(o.fs.root, remote)
	fullPath = strings.TrimPrefix(fullPath, "/")

	pdir, name := splitDirAndName(fullPath)
	pid, err := getIdByDir(o.fs, ctx, pdir)
	if err != nil {
		fs.Debugf(o.fs, "splitDirAndName/getIdByDir pdir = {%s}, fullPath = {%s}", pdir, fullPath)
		return err
	}

	// create file item at Linkbox (second step)
	opts = &rest.Opts{
		Method:  "GET",
		RootURL: linkboxAPIURL,
		Path:    "folder_upload_file",
		Parameters: url.Values{
			"token":           {o.fs.opt.Token},
			"fileMd5ofPre10m": {fmt.Sprintf("%x", h.Sum(nil))},
			"fileSize":        {strconv.FormatInt(src.Size(), 10)},
			"pid":             {strconv.Itoa(pid)},
			"diyName":         {name},
		},
	}

	getSecondStepResult := getUploadUrlResponse{}
	err = getUnmarshaledResponse(ctx, o.fs, opts, &getSecondStepResult)
	if err != nil {
		fs.Debugf(o.fs, "Update err in unparshaling response: %v", err)
		return err
	}
	if getSecondStepResult.Status != 1 {
		return fmt.Errorf("get bad status from linkbox: %s", getSecondStepResult.Message)
	}

	newObject, err := getObject(ctx, o.fs, name, pid, o.fs.opt.Token)
	if err != nil {
		fs.Debugf(o.fs, "Update err in finding object: name = {%s}, pid = {%d}", name, pid)
		return err
	}
	if newObject == (Entity{}) {
		fs.Debugf(o.fs, "Update  entity is empty: name = {%s},  pid = {%d}", name, pid)
		return fs.ErrorObjectNotFound
	}

	o.pid = pid
	o.remote = remote
	o.modTime = time.Unix(newObject.Ctime, 0)
	o.size = src.Size()

	return nil
}

// Removes this object
func (o *Object) Remove(ctx context.Context) error {
	fs.Infof(o, "Remove - {%s}", o.Remote())

	opts := &rest.Opts{
		Method:  "GET",
		RootURL: linkboxAPIURL,
		Path:    "file_del",
		Parameters: url.Values{
			"token":   {o.fs.opt.Token},
			"itemIds": {o.id},
		},
	}

	requestResult := getUploadUrlResponse{}
	err := getUnmarshaledResponse(ctx, o.fs, opts, &requestResult)
	if err != nil {
		return err
	}

	if requestResult.Status != 1 {
		return fmt.Errorf("get unexpected message from Linkbox: %s", requestResult.Message)
	}

	return nil
}

// ModTime returns the modification time of the remote http file
func (o *Object) ModTime(ctx context.Context) time.Time {
	return o.modTime
}

// Remote the name of the remote HTTP file, relative to the fs root
func (o *Object) Remote() string {
	return o.remote
}

// Size returns the size in bytes of the remote http file
func (o *Object) Size() int64 {
	return o.size
}

// String returns the URL to the remote HTTP file
func (o *Object) String() string {
	if o == nil {
		return "<nil>"
	}
	return o.remote
}

// Fs is the filesystem this remote http file object is located within
func (o *Object) Fs() fs.Info {
	return o.fs
}

// Hash returns "" since HTTP (in Go or OpenSSH) doesn't support remote calculation of hashes
func (o *Object) Hash(ctx context.Context, r hash.Type) (string, error) {
	return "", hash.ErrUnsupported
}

// MimeType of an Object if known, "" otherwise
// func (o *Object) MimeType(ctx context.Context) string {
// 	return o.contentType
// }

// Storable returns whether the remote http file is a regular file
// (not a directory, symbolic link, block device, character device, named pipe, etc.)
func (o *Object) Storable() bool {
	return true
}

// Info provides a read only interface to information about a filesystem.
// Features returns the optional features of this Fs
func (f *Fs) Features() *fs.Features {
	return f.features
}

// Name of the remote (as passed into NewFs)
// Name returns the configured name of the file system
func (f *Fs) Name() string {
	return f.name
}

// Root of the remote (as passed into NewFs)
func (f *Fs) Root() string {
	return f.root
}

// String returns a description of the FS
func (f *Fs) String() string {
	return fmt.Sprintf("Linkbox root '%s'", f.root)
}

// Precision of the ModTimes in this Fs
func (f *Fs) Precision() time.Duration {
	return fs.ModTimeNotSupported
}

// Returns the supported hash types of the filesystem
// Hashes returns hash.HashNone to indicate remote hashing is unavailable
func (f *Fs) Hashes() hash.Set {
	return hash.Set(hash.None)
}

/*
	{
	  "data": {
	    "signUrl": "http://xx -- Then CURL PUT your file with sign url "
	  },
	  "msg": "please use this url to upload (PUT method)",
	  "status": 1
	}
*/
type getResponse struct {
	Message string `json:"msg"`
	Status  int    `json:"status"`
}

type getUploadUrlData struct {
	SignUrl string `json:"signUrl"`
}

type getUploadUrlResponse struct {
	Data getUploadUrlData `json:"data"`
	getResponse
}

// Put in to the remote path with the modTime given of the given size
//
// When called from outside an Fs by rclone, src.Size() will always be >= 0.
// But for unknown-sized objects (indicated by src.Size() == -1), Put should either
// return an error or upload it properly (rather than e.g. calling panic).
//
// May create the object even if it returns an error - if so
// will return the object and the error, otherwise will return
// nil and the error
func (f *Fs) Put(ctx context.Context, in io.Reader, src fs.ObjectInfo, options ...fs.OpenOption) (fs.Object, error) {
	fs.Infof(f, "Put - {%s}", src.Remote())

	o := &Object{
		fs:     f,
		remote: src.Remote(),
		size:   src.Size(),
	}
	dir, _ := splitDirAndName(src.Remote())
	err := f.Mkdir(ctx, dir)
	if err != nil {
		return nil, err
	}
	err = o.Update(ctx, in, src, options...)
	return o, err
}

// Purge all files in the directory specified
//
// Implement this if you have a way of deleting all the files
// quicker than just running Remove() on the result of List()
//
// Return an error if it doesn't exist
func (f *Fs) Purge(ctx context.Context, dir string) error {
	fs.Infof(f, "Purge: f.root = {%s}, dir = {%s}", f.root, dir)

	return f.purgeCheck(ctx, dir, false)
}

// shouldRetry determines whether a given err rates being retried
func (f *Fs) shouldRetry(ctx context.Context, err error) (bool, error) {
	if err == fs.ErrorDirNotFound {
		fs.Infof(nil, "retry with %v", err)

		return true, err
	}

	if err == fs.ErrorObjectNotFound {
		fs.Infof(nil, "retry with %v", err)

		return true, err
	}
	return false, err
}

// Check the interfaces are satisfied
var (
	_ fs.Fs     = &Fs{}
	_ fs.Purger = &Fs{}
	_ fs.Object = &Object{}
)

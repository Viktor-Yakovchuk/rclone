package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// ManagementOperationState enumerates the values for management operation state.
type ManagementOperationState string

const (
	// CancelInProgress ...
	CancelInProgress ManagementOperationState = "CancelInProgress"
	// Cancelled ...
	Cancelled ManagementOperationState = "Cancelled"
	// Failed ...
	Failed ManagementOperationState = "Failed"
	// InProgress ...
	InProgress ManagementOperationState = "InProgress"
	// Pending ...
	Pending ManagementOperationState = "Pending"
	// Succeeded ...
	Succeeded ManagementOperationState = "Succeeded"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// PossibleManagementOperationStateValues returns an array of possible values for the ManagementOperationState const type.
func PossibleManagementOperationStateValues() []ManagementOperationState {
	return []ManagementOperationState{CancelInProgress, Cancelled, Failed, InProgress, Pending, Succeeded}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// DatabaseOperation a database operation.
type DatabaseOperation struct {
	// DatabaseOperationProperties - Resource properties.
	*DatabaseOperationProperties `json:"properties,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// MarshalJSON is the custom marshaler for DatabaseOperation.
func (do DatabaseOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if do.DatabaseOperationProperties != nil {
		objectMap["properties"] = do.DatabaseOperationProperties
	}
	if do.ID != nil {
		objectMap["id"] = do.ID
	}
	if do.Name != nil {
		objectMap["name"] = do.Name
	}
	if do.Type != nil {
		objectMap["type"] = do.Type
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// UnmarshalJSON is the custom unmarshaler for DatabaseOperation struct.
func (do *DatabaseOperation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var databaseOperationProperties DatabaseOperationProperties
				err = json.Unmarshal(*v, &databaseOperationProperties)
				if err != nil {
					return err
				}
				do.DatabaseOperationProperties = &databaseOperationProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				do.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				do.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				do.Type = &typeVar
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// DatabaseOperationListResult the response to a list database operations request
type DatabaseOperationListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of results.
	Value *[]DatabaseOperation `json:"value,omitempty"`
	// NextLink - Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// DatabaseOperationListResultIterator provides access to a complete listing of DatabaseOperation values.
type DatabaseOperationListResultIterator struct {
	i    int
	page DatabaseOperationListResultPage
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DatabaseOperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DatabaseOperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Response returns the raw server response from the last page request.
func (iter DatabaseOperationListResultIterator) Response() DatabaseOperationListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DatabaseOperationListResultIterator) Value() DatabaseOperation {
	if !iter.page.NotDone() {
		return DatabaseOperation{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// IsEmpty returns true if the ListResult contains no values.
func (dolr DatabaseOperationListResult) IsEmpty() bool {
	return dolr.Value == nil || len(*dolr.Value) == 0
}

// databaseOperationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dolr DatabaseOperationListResult) databaseOperationListResultPreparer() (*http.Request, error) {
	if dolr.NextLink == nil || len(to.String(dolr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dolr.NextLink)))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// DatabaseOperationListResultPage contains a page of DatabaseOperation values.
type DatabaseOperationListResultPage struct {
	fn   func(DatabaseOperationListResult) (DatabaseOperationListResult, error)
	dolr DatabaseOperationListResult
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DatabaseOperationListResultPage) Next() error {
	next, err := page.fn(page.dolr)
	if err != nil {
		return err
	}
	page.dolr = next
	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DatabaseOperationListResultPage) NotDone() bool {
	return !page.dolr.IsEmpty()
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Response returns the raw server response from the last page request.
func (page DatabaseOperationListResultPage) Response() DatabaseOperationListResult {
	return page.dolr
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page DatabaseOperationListResultPage) Values() []DatabaseOperation {
	if page.dolr.IsEmpty() {
		return nil
	}
	return *page.dolr.Value
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// DatabaseOperationProperties the properties of a database operation.
type DatabaseOperationProperties struct {
	// DatabaseName - The name of the database the operation is being performed on.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Operation - The name of operation.
	Operation *string `json:"operation,omitempty"`
	// OperationFriendlyName - The friendly name of operation.
	OperationFriendlyName *string `json:"operationFriendlyName,omitempty"`
	// PercentComplete - The percentage of the operation completed.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// ServerName - The name of the server.
	ServerName *string `json:"serverName,omitempty"`
	// StartTime - The operation start time.
	StartTime *date.Time `json:"startTime,omitempty"`
	// State - The operation state. Possible values include: 'Pending', 'InProgress', 'Succeeded', 'Failed', 'CancelInProgress', 'Cancelled'
	State ManagementOperationState `json:"state,omitempty"`
	// ErrorCode - The operation error code.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// ErrorDescription - The operation error description.
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// ErrorSeverity - The operation error severity.
	ErrorSeverity *int32 `json:"errorSeverity,omitempty"`
	// IsUserError - Whether or not the error is a user error.
	IsUserError *bool `json:"isUserError,omitempty"`
	// EstimatedCompletionTime - The estimated completion time of the operation.
	EstimatedCompletionTime *date.Time `json:"estimatedCompletionTime,omitempty"`
	// Description - The operation description.
	Description *string `json:"description,omitempty"`
	// IsCancellable - Whether the operation can be cancelled.
	IsCancellable *bool `json:"isCancellable,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// ElasticPoolOperation a elastic pool operation.
type ElasticPoolOperation struct {
	// ElasticPoolOperationProperties - Resource properties.
	*ElasticPoolOperationProperties `json:"properties,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// MarshalJSON is the custom marshaler for ElasticPoolOperation.
func (epo ElasticPoolOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if epo.ElasticPoolOperationProperties != nil {
		objectMap["properties"] = epo.ElasticPoolOperationProperties
	}
	if epo.ID != nil {
		objectMap["id"] = epo.ID
	}
	if epo.Name != nil {
		objectMap["name"] = epo.Name
	}
	if epo.Type != nil {
		objectMap["type"] = epo.Type
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// UnmarshalJSON is the custom unmarshaler for ElasticPoolOperation struct.
func (epo *ElasticPoolOperation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var elasticPoolOperationProperties ElasticPoolOperationProperties
				err = json.Unmarshal(*v, &elasticPoolOperationProperties)
				if err != nil {
					return err
				}
				epo.ElasticPoolOperationProperties = &elasticPoolOperationProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				epo.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				epo.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				epo.Type = &typeVar
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// ElasticPoolOperationListResult the response to a list elastic pool operations request
type ElasticPoolOperationListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of results.
	Value *[]ElasticPoolOperation `json:"value,omitempty"`
	// NextLink - Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// ElasticPoolOperationListResultIterator provides access to a complete listing of ElasticPoolOperation values.
type ElasticPoolOperationListResultIterator struct {
	i    int
	page ElasticPoolOperationListResultPage
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ElasticPoolOperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ElasticPoolOperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Response returns the raw server response from the last page request.
func (iter ElasticPoolOperationListResultIterator) Response() ElasticPoolOperationListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ElasticPoolOperationListResultIterator) Value() ElasticPoolOperation {
	if !iter.page.NotDone() {
		return ElasticPoolOperation{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// IsEmpty returns true if the ListResult contains no values.
func (epolr ElasticPoolOperationListResult) IsEmpty() bool {
	return epolr.Value == nil || len(*epolr.Value) == 0
}

// elasticPoolOperationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (epolr ElasticPoolOperationListResult) elasticPoolOperationListResultPreparer() (*http.Request, error) {
	if epolr.NextLink == nil || len(to.String(epolr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(epolr.NextLink)))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// ElasticPoolOperationListResultPage contains a page of ElasticPoolOperation values.
type ElasticPoolOperationListResultPage struct {
	fn    func(ElasticPoolOperationListResult) (ElasticPoolOperationListResult, error)
	epolr ElasticPoolOperationListResult
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ElasticPoolOperationListResultPage) Next() error {
	next, err := page.fn(page.epolr)
	if err != nil {
		return err
	}
	page.epolr = next
	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ElasticPoolOperationListResultPage) NotDone() bool {
	return !page.epolr.IsEmpty()
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Response returns the raw server response from the last page request.
func (page ElasticPoolOperationListResultPage) Response() ElasticPoolOperationListResult {
	return page.epolr
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page ElasticPoolOperationListResultPage) Values() []ElasticPoolOperation {
	if page.epolr.IsEmpty() {
		return nil
	}
	return *page.epolr.Value
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// ElasticPoolOperationProperties the properties of a elastic pool operation.
type ElasticPoolOperationProperties struct {
	// ElasticPoolName - The name of the elastic pool the operation is being performed on.
	ElasticPoolName *string `json:"elasticPoolName,omitempty"`
	// Operation - The name of operation.
	Operation *string `json:"operation,omitempty"`
	// OperationFriendlyName - The friendly name of operation.
	OperationFriendlyName *string `json:"operationFriendlyName,omitempty"`
	// PercentComplete - The percentage of the operation completed.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// ServerName - The name of the server.
	ServerName *string `json:"serverName,omitempty"`
	// StartTime - The operation start time.
	StartTime *date.Time `json:"startTime,omitempty"`
	// State - The operation state.
	State *string `json:"state,omitempty"`
	// ErrorCode - The operation error code.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// ErrorDescription - The operation error description.
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// ErrorSeverity - The operation error severity.
	ErrorSeverity *int32 `json:"errorSeverity,omitempty"`
	// IsUserError - Whether or not the error is a user error.
	IsUserError *bool `json:"isUserError,omitempty"`
	// EstimatedCompletionTime - The estimated completion time of the operation.
	EstimatedCompletionTime *date.Time `json:"estimatedCompletionTime,omitempty"`
	// Description - The operation description.
	Description *string `json:"description,omitempty"`
	// IsCancellable - Whether the operation can be cancelled.
	IsCancellable *bool `json:"isCancellable,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// ProxyResource ARM proxy resource.
type ProxyResource struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql instead.
// Resource ARM resource.
type Resource struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

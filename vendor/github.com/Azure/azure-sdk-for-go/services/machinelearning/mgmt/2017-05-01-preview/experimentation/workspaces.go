package experimentation

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// WorkspacesClient is the these APIs allow end users to operate on Azure Machine Learning Team Account resources. They
// support CRUD operations for Azure Machine Learning Team Accounts.
type WorkspacesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// NewWorkspacesClient creates an instance of the WorkspacesClient client.
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return NewWorkspacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// NewWorkspacesClientWithBaseURI creates an instance of the WorkspacesClient client.
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return WorkspacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// CreateOrUpdate creates or updates a machine learning workspace with the specified parameters.
//
// resourceGroupName is the name of the resource group to which the machine learning team account belongs.
// accountName is the name of the machine learning team account. workspaceName is the name of the machine learning
// team account workspace. parameters is the parameters for creating or updating a machine learning workspace.
func (client WorkspacesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, parameters Workspace) (result Workspace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.WorkspaceProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.WorkspaceProperties.FriendlyName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("experimentation.WorkspacesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, accountName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkspacesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, parameters Workspace) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2017-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) CreateOrUpdateResponder(resp *http.Response) (result Workspace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// Delete deletes a machine learning workspace.
//
// resourceGroupName is the name of the resource group to which the machine learning team account belongs.
// accountName is the name of the machine learning team account. workspaceName is the name of the machine learning
// team account workspace.
func (client WorkspacesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, workspaceName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("experimentation.WorkspacesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// DeletePreparer prepares the Delete request.
func (client WorkspacesClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2017-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// Get gets the properties of the specified machine learning workspace.
//
// resourceGroupName is the name of the resource group to which the machine learning team account belongs.
// accountName is the name of the machine learning team account. workspaceName is the name of the machine learning
// team account workspace.
func (client WorkspacesClient) Get(ctx context.Context, resourceGroupName string, accountName string, workspaceName string) (result Workspace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("experimentation.WorkspacesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// GetPreparer prepares the Get request.
func (client WorkspacesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2017-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) GetResponder(resp *http.Response) (result Workspace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// ListByAccounts lists all the available machine learning workspaces under the specified team account.
//
// accountName is the name of the machine learning team account. resourceGroupName is the name of the resource
// group to which the machine learning team account belongs.
func (client WorkspacesClient) ListByAccounts(ctx context.Context, accountName string, resourceGroupName string) (result WorkspaceListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("experimentation.WorkspacesClient", "ListByAccounts", err.Error())
	}

	result.fn = client.listByAccountsNextResults
	req, err := client.ListByAccountsPreparer(ctx, accountName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "ListByAccounts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAccountsSender(req)
	if err != nil {
		result.wlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "ListByAccounts", resp, "Failure sending request")
		return
	}

	result.wlr, err = client.ListByAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "ListByAccounts", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// ListByAccountsPreparer prepares the ListByAccounts request.
func (client WorkspacesClient) ListByAccountsPreparer(ctx context.Context, accountName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// ListByAccountsSender sends the ListByAccounts request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) ListByAccountsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// ListByAccountsResponder handles the response to the ListByAccounts request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) ListByAccountsResponder(resp *http.Response) (result WorkspaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAccountsNextResults retrieves the next set of results, if any.
func (client WorkspacesClient) listByAccountsNextResults(lastResults WorkspaceListResult) (result WorkspaceListResult, err error) {
	req, err := lastResults.workspaceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "listByAccountsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAccountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "listByAccountsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "listByAccountsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// ListByAccountsComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspacesClient) ListByAccountsComplete(ctx context.Context, accountName string, resourceGroupName string) (result WorkspaceListResultIterator, err error) {
	result.page, err = client.ListByAccounts(ctx, accountName, resourceGroupName)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// Update updates a machine learning workspace with the specified parameters.
//
// resourceGroupName is the name of the resource group to which the machine learning team account belongs.
// accountName is the name of the machine learning team account. workspaceName is the name of the machine learning
// team account workspace. parameters is the parameters for updating a machine learning workspace.
func (client WorkspacesClient) Update(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, parameters WorkspaceUpdateParameters) (result Workspace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("experimentation.WorkspacesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experimentation.WorkspacesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// UpdatePreparer prepares the Update request.
func (client WorkspacesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, parameters WorkspaceUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2017-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation instead.
// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) UpdateResponder(resp *http.Response) (result Workspace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

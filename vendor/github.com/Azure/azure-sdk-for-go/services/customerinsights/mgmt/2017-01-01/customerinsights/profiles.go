package customerinsights

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

// ProfilesClient is the the Azure Customer Insights management API provides a RESTful set of web services that
// interact with Azure Customer Insights service to manage your resources. The API has entities that capture the
// relationship between an end user and the Azure Customer Insights service.
type ProfilesClient struct {
	BaseClient
}

// NewProfilesClient creates an instance of the ProfilesClient client.
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return NewProfilesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProfilesClientWithBaseURI creates an instance of the ProfilesClient client.
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return ProfilesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a profile within a Hub, or updates an existing profile.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// profileName - the name of the profile.
// parameters - parameters supplied to the create/delete Profile type operation
func (client ProfilesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, profileName string, parameters ProfileResourceFormat) (result ProfilesCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: profileName,
			Constraints: []validation.Constraint{{Target: "profileName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "profileName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("customerinsights.ProfilesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hubName, profileName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ProfilesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hubName string, profileName string, parameters ProfileResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) CreateOrUpdateSender(req *http.Request) (future ProfilesCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ProfilesClient) CreateOrUpdateResponder(resp *http.Response) (result ProfileResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a profile within a hub
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// profileName - the name of the profile.
// localeCode - locale of profile to retrieve, default is en-us.
func (client ProfilesClient) Delete(ctx context.Context, resourceGroupName string, hubName string, profileName string, localeCode string) (result ProfilesDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, hubName, profileName, localeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ProfilesClient) DeletePreparer(ctx context.Context, resourceGroupName string, hubName string, profileName string, localeCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(localeCode) > 0 {
		queryParameters["locale-code"] = autorest.Encode("query", localeCode)
	} else {
		queryParameters["locale-code"] = autorest.Encode("query", "en-us")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) DeleteSender(req *http.Request) (future ProfilesDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProfilesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified profile.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// profileName - the name of the profile.
// localeCode - locale of profile to retrieve, default is en-us.
func (client ProfilesClient) Get(ctx context.Context, resourceGroupName string, hubName string, profileName string, localeCode string) (result ProfileResourceFormat, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, hubName, profileName, localeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProfilesClient) GetPreparer(ctx context.Context, resourceGroupName string, hubName string, profileName string, localeCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(localeCode) > 0 {
		queryParameters["locale-code"] = autorest.Encode("query", localeCode)
	} else {
		queryParameters["locale-code"] = autorest.Encode("query", "en-us")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetResponder(resp *http.Response) (result ProfileResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEnrichingKpis gets the KPIs that enrich the profile Type identified by the supplied name. Enrichment happens
// through participants of the Interaction on an Interaction KPI and through Relationships for Profile KPIs.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// profileName - the name of the profile.
func (client ProfilesClient) GetEnrichingKpis(ctx context.Context, resourceGroupName string, hubName string, profileName string) (result ListKpiDefinition, err error) {
	req, err := client.GetEnrichingKpisPreparer(ctx, resourceGroupName, hubName, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "GetEnrichingKpis", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEnrichingKpisSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "GetEnrichingKpis", resp, "Failure sending request")
		return
	}

	result, err = client.GetEnrichingKpisResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "GetEnrichingKpis", resp, "Failure responding to request")
	}

	return
}

// GetEnrichingKpisPreparer prepares the GetEnrichingKpis request.
func (client ProfilesClient) GetEnrichingKpisPreparer(ctx context.Context, resourceGroupName string, hubName string, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}/getEnrichingKpis", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEnrichingKpisSender sends the GetEnrichingKpis request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GetEnrichingKpisSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetEnrichingKpisResponder handles the response to the GetEnrichingKpis request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetEnrichingKpisResponder(resp *http.Response) (result ListKpiDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all profile in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// localeCode - locale of profile to retrieve, default is en-us.
func (client ProfilesClient) ListByHub(ctx context.Context, resourceGroupName string, hubName string, localeCode string) (result ProfileListResultPage, err error) {
	result.fn = client.listByHubNextResults
	req, err := client.ListByHubPreparer(ctx, resourceGroupName, hubName, localeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", resp, "Failure responding to request")
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client ProfilesClient) ListByHubPreparer(ctx context.Context, resourceGroupName string, hubName string, localeCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(localeCode) > 0 {
		queryParameters["locale-code"] = autorest.Encode("query", localeCode)
	} else {
		queryParameters["locale-code"] = autorest.Encode("query", "en-us")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListByHubResponder(resp *http.Response) (result ProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHubNextResults retrieves the next set of results, if any.
func (client ProfilesClient) listByHubNextResults(lastResults ProfileListResult) (result ProfileListResult, err error) {
	req, err := lastResults.profileListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "listByHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "listByHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "listByHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProfilesClient) ListByHubComplete(ctx context.Context, resourceGroupName string, hubName string, localeCode string) (result ProfileListResultIterator, err error) {
	result.page, err = client.ListByHub(ctx, resourceGroupName, hubName, localeCode)
	return
}

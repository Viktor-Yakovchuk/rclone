package datamigration

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
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2017-11-15-preview/datamigration instead.
// OperationsClient is the data Migration Client
type OperationsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2017-11-15-preview/datamigration instead.
// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient(subscriptionID string) OperationsClient {
	return NewOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2017-11-15-preview/datamigration instead.
// NewOperationsClientWithBaseURI creates an instance of the OperationsClient client.
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return OperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2017-11-15-preview/datamigration instead.
// List lists all available actions exposed by the Data Migration Service resource provider.
func (client OperationsClient) List(ctx context.Context) (result ServiceOperationListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.OperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sol.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.OperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.sol, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.OperationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2017-11-15-preview/datamigration instead.
// ListPreparer prepares the List request.
func (client OperationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.DataMigration/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2017-11-15-preview/datamigration instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2017-11-15-preview/datamigration instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OperationsClient) ListResponder(resp *http.Response) (result ServiceOperationList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client OperationsClient) listNextResults(lastResults ServiceOperationList) (result ServiceOperationList, err error) {
	req, err := lastResults.serviceOperationListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datamigration.OperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datamigration.OperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.OperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2017-11-15-preview/datamigration instead.
// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client OperationsClient) ListComplete(ctx context.Context) (result ServiceOperationListIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

package migrate

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProjectsClient is the migrate your workloads to Azure.
type ProjectsClient struct {
	BaseClient
}

// NewProjectsClient creates an instance of the ProjectsClient client.
func NewProjectsClient(subscriptionID string, acceptLanguage string) ProjectsClient {
	return NewProjectsClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewProjectsClientWithBaseURI creates an instance of the ProjectsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) ProjectsClient {
	return ProjectsClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// DeleteMigrateProject delete the migrate project. Deleting non-existent project is a no-operation.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
func (client ProjectsClient) DeleteMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.DeleteMigrateProject")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteMigrateProjectPreparer(ctx, resourceGroupName, migrateProjectName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "DeleteMigrateProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteMigrateProjectSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "DeleteMigrateProject", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteMigrateProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "DeleteMigrateProject", resp, "Failure responding to request")
	}

	return
}

// DeleteMigrateProjectPreparer prepares the DeleteMigrateProject request.
func (client ProjectsClient) DeleteMigrateProjectPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteMigrateProjectSender sends the DeleteMigrateProject request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) DeleteMigrateProjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteMigrateProjectResponder handles the response to the DeleteMigrateProject request. The method always
// closes the http.Response Body.
func (client ProjectsClient) DeleteMigrateProjectResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetMigrateProject sends the get migrate project request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
func (client ProjectsClient) GetMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.GetMigrateProject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMigrateProjectPreparer(ctx, resourceGroupName, migrateProjectName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "GetMigrateProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMigrateProjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "GetMigrateProject", resp, "Failure sending request")
		return
	}

	result, err = client.GetMigrateProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "GetMigrateProject", resp, "Failure responding to request")
	}

	return
}

// GetMigrateProjectPreparer prepares the GetMigrateProject request.
func (client ProjectsClient) GetMigrateProjectPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMigrateProjectSender sends the GetMigrateProject request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) GetMigrateProjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetMigrateProjectResponder handles the response to the GetMigrateProject request. The method always
// closes the http.Response Body.
func (client ProjectsClient) GetMigrateProjectResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PatchMigrateProject update a migrate project with specified name. Supports partial updates, for example only tags
// can be provided.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// body - body with migrate project details.
func (client ProjectsClient) PatchMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, body Project) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.PatchMigrateProject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PatchMigrateProjectPreparer(ctx, resourceGroupName, migrateProjectName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "PatchMigrateProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchMigrateProjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "PatchMigrateProject", resp, "Failure sending request")
		return
	}

	result, err = client.PatchMigrateProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "PatchMigrateProject", resp, "Failure responding to request")
	}

	return
}

// PatchMigrateProjectPreparer prepares the PatchMigrateProject request.
func (client ProjectsClient) PatchMigrateProjectPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, body Project) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchMigrateProjectSender sends the PatchMigrateProject request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) PatchMigrateProjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PatchMigrateProjectResponder handles the response to the PatchMigrateProject request. The method always
// closes the http.Response Body.
func (client ProjectsClient) PatchMigrateProjectResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutMigrateProject sends the put migrate project request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// body - body with migrate project details.
func (client ProjectsClient) PutMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, body Project) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.PutMigrateProject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutMigrateProjectPreparer(ctx, resourceGroupName, migrateProjectName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "PutMigrateProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMigrateProjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "PutMigrateProject", resp, "Failure sending request")
		return
	}

	result, err = client.PutMigrateProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "PutMigrateProject", resp, "Failure responding to request")
	}

	return
}

// PutMigrateProjectPreparer prepares the PutMigrateProject request.
func (client ProjectsClient) PutMigrateProjectPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, body Project) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMigrateProjectSender sends the PutMigrateProject request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) PutMigrateProjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutMigrateProjectResponder handles the response to the PutMigrateProject request. The method always
// closes the http.Response Body.
func (client ProjectsClient) PutMigrateProjectResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RefreshMigrateProjectSummary sends the refresh migrate project summary request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// input - the goal input which needs to be refreshed.
func (client ProjectsClient) RefreshMigrateProjectSummary(ctx context.Context, resourceGroupName string, migrateProjectName string, input RefreshSummaryInput) (result RefreshSummaryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.RefreshMigrateProjectSummary")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RefreshMigrateProjectSummaryPreparer(ctx, resourceGroupName, migrateProjectName, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "RefreshMigrateProjectSummary", nil, "Failure preparing request")
		return
	}

	resp, err := client.RefreshMigrateProjectSummarySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "RefreshMigrateProjectSummary", resp, "Failure sending request")
		return
	}

	result, err = client.RefreshMigrateProjectSummaryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "RefreshMigrateProjectSummary", resp, "Failure responding to request")
	}

	return
}

// RefreshMigrateProjectSummaryPreparer prepares the RefreshMigrateProjectSummary request.
func (client ProjectsClient) RefreshMigrateProjectSummaryPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, input RefreshSummaryInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/refreshSummary", pathParameters),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshMigrateProjectSummarySender sends the RefreshMigrateProjectSummary request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) RefreshMigrateProjectSummarySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RefreshMigrateProjectSummaryResponder handles the response to the RefreshMigrateProjectSummary request. The method always
// closes the http.Response Body.
func (client ProjectsClient) RefreshMigrateProjectSummaryResponder(resp *http.Response) (result RefreshSummaryResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegisterTool sends the register tool request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// input - input containing the name of the tool to be registered.
func (client ProjectsClient) RegisterTool(ctx context.Context, resourceGroupName string, migrateProjectName string, input RegisterToolInput) (result RegistrationResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.RegisterTool")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegisterToolPreparer(ctx, resourceGroupName, migrateProjectName, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "RegisterTool", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegisterToolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "RegisterTool", resp, "Failure sending request")
		return
	}

	result, err = client.RegisterToolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsClient", "RegisterTool", resp, "Failure responding to request")
	}

	return
}

// RegisterToolPreparer prepares the RegisterTool request.
func (client ProjectsClient) RegisterToolPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, input RegisterToolInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/registerTool", pathParameters),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegisterToolSender sends the RegisterTool request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) RegisterToolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegisterToolResponder handles the response to the RegisterTool request. The method always
// closes the http.Response Body.
func (client ProjectsClient) RegisterToolResponder(resp *http.Response) (result RegistrationResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

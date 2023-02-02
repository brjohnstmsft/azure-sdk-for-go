//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragemover

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// EndpointsClient contains the methods for the Endpoints group.
// Don't use this type directly, use NewEndpointsClient() instead.
type EndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEndpointsClient creates a new instance of EndpointsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EndpointsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &EndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an Endpoint resource, which represents a data transfer source or destination.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - endpointName - The name of the Endpoint resource.
//   - options - EndpointsClientCreateOrUpdateOptions contains the optional parameters for the EndpointsClient.CreateOrUpdate
//     method.
func (client *EndpointsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, endpoint Endpoint, options *EndpointsClientCreateOrUpdateOptions) (EndpointsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageMoverName, endpointName, endpoint, options)
	if err != nil {
		return EndpointsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EndpointsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EndpointsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, endpoint Endpoint, options *EndpointsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, endpoint)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EndpointsClient) createOrUpdateHandleResponse(resp *http.Response) (EndpointsClientCreateOrUpdateResponse, error) {
	result := EndpointsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Endpoint); err != nil {
		return EndpointsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes an Endpoint resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - endpointName - The name of the Endpoint resource.
//   - options - EndpointsClientBeginDeleteOptions contains the optional parameters for the EndpointsClient.BeginDelete method.
func (client *EndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, options *EndpointsClientBeginDeleteOptions) (*runtime.Poller[EndpointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, storageMoverName, endpointName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[EndpointsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EndpointsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an Endpoint resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *EndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, options *EndpointsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageMoverName, endpointName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, options *EndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an Endpoint resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - endpointName - The name of the Endpoint resource.
//   - options - EndpointsClientGetOptions contains the optional parameters for the EndpointsClient.Get method.
func (client *EndpointsClient) Get(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, options *EndpointsClientGetOptions) (EndpointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageMoverName, endpointName, options)
	if err != nil {
		return EndpointsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EndpointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, options *EndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EndpointsClient) getHandleResponse(resp *http.Response) (EndpointsClientGetResponse, error) {
	result := EndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Endpoint); err != nil {
		return EndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all Endpoints in a Storage Mover.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - options - EndpointsClientListOptions contains the optional parameters for the EndpointsClient.NewListPager method.
func (client *EndpointsClient) NewListPager(resourceGroupName string, storageMoverName string, options *EndpointsClientListOptions) *runtime.Pager[EndpointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EndpointsClientListResponse]{
		More: func(page EndpointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EndpointsClientListResponse) (EndpointsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, storageMoverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EndpointsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EndpointsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EndpointsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, options *EndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/endpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EndpointsClient) listHandleResponse(resp *http.Response) (EndpointsClientListResponse, error) {
	result := EndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointList); err != nil {
		return EndpointsClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates properties for an Endpoint resource. Properties not specified in the request body will be unchanged.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - endpointName - The name of the Endpoint resource.
//   - options - EndpointsClientUpdateOptions contains the optional parameters for the EndpointsClient.Update method.
func (client *EndpointsClient) Update(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, endpoint EndpointBaseUpdateParameters, options *EndpointsClientUpdateOptions) (EndpointsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageMoverName, endpointName, endpoint, options)
	if err != nil {
		return EndpointsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EndpointsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EndpointsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *EndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, endpointName string, endpoint EndpointBaseUpdateParameters, options *EndpointsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, endpoint)
}

// updateHandleResponse handles the Update response.
func (client *EndpointsClient) updateHandleResponse(resp *http.Response) (EndpointsClientUpdateResponse, error) {
	result := EndpointsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Endpoint); err != nil {
		return EndpointsClientUpdateResponse{}, err
	}
	return result, nil
}

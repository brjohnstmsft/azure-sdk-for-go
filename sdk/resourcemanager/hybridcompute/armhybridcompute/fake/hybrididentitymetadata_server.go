//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
	"net/http"
	"net/url"
	"regexp"
)

// HybridIdentityMetadataServer is a fake server for instances of the armhybridcompute.HybridIdentityMetadataClient type.
type HybridIdentityMetadataServer struct {
	// Get is the fake for method HybridIdentityMetadataClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, machineName string, metadataName string, options *armhybridcompute.HybridIdentityMetadataClientGetOptions) (resp azfake.Responder[armhybridcompute.HybridIdentityMetadataClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByMachinesPager is the fake for method HybridIdentityMetadataClient.NewListByMachinesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMachinesPager func(resourceGroupName string, machineName string, options *armhybridcompute.HybridIdentityMetadataClientListByMachinesOptions) (resp azfake.PagerResponder[armhybridcompute.HybridIdentityMetadataClientListByMachinesResponse])
}

// NewHybridIdentityMetadataServerTransport creates a new instance of HybridIdentityMetadataServerTransport with the provided implementation.
// The returned HybridIdentityMetadataServerTransport instance is connected to an instance of armhybridcompute.HybridIdentityMetadataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHybridIdentityMetadataServerTransport(srv *HybridIdentityMetadataServer) *HybridIdentityMetadataServerTransport {
	return &HybridIdentityMetadataServerTransport{
		srv:                    srv,
		newListByMachinesPager: newTracker[azfake.PagerResponder[armhybridcompute.HybridIdentityMetadataClientListByMachinesResponse]](),
	}
}

// HybridIdentityMetadataServerTransport connects instances of armhybridcompute.HybridIdentityMetadataClient to instances of HybridIdentityMetadataServer.
// Don't use this type directly, use NewHybridIdentityMetadataServerTransport instead.
type HybridIdentityMetadataServerTransport struct {
	srv                    *HybridIdentityMetadataServer
	newListByMachinesPager *tracker[azfake.PagerResponder[armhybridcompute.HybridIdentityMetadataClientListByMachinesResponse]]
}

// Do implements the policy.Transporter interface for HybridIdentityMetadataServerTransport.
func (h *HybridIdentityMetadataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HybridIdentityMetadataClient.Get":
		resp, err = h.dispatchGet(req)
	case "HybridIdentityMetadataClient.NewListByMachinesPager":
		resp, err = h.dispatchNewListByMachinesPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HybridIdentityMetadataServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridIdentityMetadata/(?P<metadataName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	metadataNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("metadataName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, machineNameParam, metadataNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HybridIdentityMetadata, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridIdentityMetadataServerTransport) dispatchNewListByMachinesPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByMachinesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMachinesPager not implemented")}
	}
	newListByMachinesPager := h.newListByMachinesPager.get(req)
	if newListByMachinesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridIdentityMetadata`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListByMachinesPager(resourceGroupNameParam, machineNameParam, nil)
		newListByMachinesPager = &resp
		h.newListByMachinesPager.add(req, newListByMachinesPager)
		server.PagerResponderInjectNextLinks(newListByMachinesPager, req, func(page *armhybridcompute.HybridIdentityMetadataClientListByMachinesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByMachinesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByMachinesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByMachinesPager) {
		h.newListByMachinesPager.remove(req)
	}
	return resp, nil
}

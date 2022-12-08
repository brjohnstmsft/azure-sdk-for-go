//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armqumulo

// FileSystemsClientCreateOrUpdateResponse contains the response from method FileSystemsClient.CreateOrUpdate.
type FileSystemsClientCreateOrUpdateResponse struct {
	FileSystemResource
}

// FileSystemsClientDeleteResponse contains the response from method FileSystemsClient.Delete.
type FileSystemsClientDeleteResponse struct {
	// placeholder for future response values
}

// FileSystemsClientGetResponse contains the response from method FileSystemsClient.Get.
type FileSystemsClientGetResponse struct {
	FileSystemResource
}

// FileSystemsClientListByResourceGroupResponse contains the response from method FileSystemsClient.ListByResourceGroup.
type FileSystemsClientListByResourceGroupResponse struct {
	FileSystemResourceListResult
}

// FileSystemsClientListBySubscriptionResponse contains the response from method FileSystemsClient.ListBySubscription.
type FileSystemsClientListBySubscriptionResponse struct {
	FileSystemResourceListResult
}

// FileSystemsClientUpdateResponse contains the response from method FileSystemsClient.Update.
type FileSystemsClientUpdateResponse struct {
	FileSystemResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

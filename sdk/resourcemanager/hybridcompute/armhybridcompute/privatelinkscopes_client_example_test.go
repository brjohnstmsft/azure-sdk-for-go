//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_List.json
func ExamplePrivateLinkScopesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkScopesClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkScopeListResult = armhybridcompute.PrivateLinkScopeListResult{
		// 	Value: []*armhybridcompute.PrivateLinkScope{
		// 		{
		// 			Name: to.Ptr("my-privatelinkscope"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armhybridcompute.PrivateLinkScopeProperties{
		// 				PrivateLinkScopeID: to.Ptr("e5dc51d3-92ed-4d7e-947a-775ea79b4919"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-other-privatelinkscope"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-other-privatelinkscope"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armhybridcompute.PrivateLinkScopeProperties{
		// 				PrivateEndpointConnections: []*armhybridcompute.PrivateEndpointConnectionDataModel{
		// 					{
		// 						Name: to.Ptr("private-endpoint-connection-name"),
		// 						Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
		// 						Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr("None"),
		// 								Status: to.Ptr("Approved"),
		// 							},
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 						},
		// 				}},
		// 				PrivateLinkScopeID: to.Ptr("f5dc51d3-92ed-4d7e-947a-775ea79b4919"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_ListByResourceGroup.json
func ExamplePrivateLinkScopesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkScopesClient().NewListByResourceGroupPager("my-resource-group", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkScopeListResult = armhybridcompute.PrivateLinkScopeListResult{
		// 	Value: []*armhybridcompute.PrivateLinkScope{
		// 		{
		// 			Name: to.Ptr("my-privatelinkscope"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armhybridcompute.PrivateLinkScopeProperties{
		// 				PrivateEndpointConnections: []*armhybridcompute.PrivateEndpointConnectionDataModel{
		// 					{
		// 						Name: to.Ptr("private-endpoint-connection-name"),
		// 						Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
		// 						Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr("None"),
		// 								Status: to.Ptr("Approved"),
		// 							},
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 						},
		// 				}},
		// 				PrivateLinkScopeID: to.Ptr("f5dc51d3-92ed-4d7e-947a-775ea79b4919"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-other-privatelinkscope"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-other-privatelinkscope"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armhybridcompute.PrivateLinkScopeProperties{
		// 				PrivateEndpointConnections: []*armhybridcompute.PrivateEndpointConnectionDataModel{
		// 					{
		// 						Name: to.Ptr("private-endpoint-connection-name"),
		// 						Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
		// 						Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr("None"),
		// 								Status: to.Ptr("Approved"),
		// 							},
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 						},
		// 				}},
		// 				PrivateLinkScopeID: to.Ptr("a5dc51d3-92ed-4d7e-947a-775ea79b4919"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_Delete.json
func ExamplePrivateLinkScopesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkScopesClient().BeginDelete(ctx, "my-resource-group", "my-privatelinkscope", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_Get.json
func ExamplePrivateLinkScopesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().Get(ctx, "my-resource-group", "my-privatelinkscope", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkScope = armhybridcompute.PrivateLinkScope{
	// 	Name: to.Ptr("my-privatelinkscope"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armhybridcompute.PrivateLinkScopeProperties{
	// 		PrivateEndpointConnections: []*armhybridcompute.PrivateEndpointConnectionDataModel{
	// 			{
	// 				Name: to.Ptr("private-endpoint-connection-name"),
	// 				Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 				Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr("Approved"),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 				},
	// 		}},
	// 		PrivateLinkScopeID: to.Ptr("f5dc51d3-92ed-4d7e-947a-775ea79b4919"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_Create.json
func ExamplePrivateLinkScopesClient_CreateOrUpdate_privateLinkScopeCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().CreateOrUpdate(ctx, "my-resource-group", "my-privatelinkscope", armhybridcompute.PrivateLinkScope{
		Location: to.Ptr("westus"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkScope = armhybridcompute.PrivateLinkScope{
	// 	Name: to.Ptr("my-privatelinkscope"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armhybridcompute.PrivateLinkScopeProperties{
	// 		PrivateEndpointConnections: []*armhybridcompute.PrivateEndpointConnectionDataModel{
	// 			{
	// 				Name: to.Ptr("private-endpoint-connection-name"),
	// 				Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 				Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr("Approved"),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 				},
	// 		}},
	// 		PrivateLinkScopeID: to.Ptr("e5dc51d3-92ed-4d7e-947a-775ea79b4919"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_Update.json
func ExamplePrivateLinkScopesClient_CreateOrUpdate_privateLinkScopeUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().CreateOrUpdate(ctx, "my-resource-group", "my-privatelinkscope", armhybridcompute.PrivateLinkScope{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"Tag1": to.Ptr("Value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkScope = armhybridcompute.PrivateLinkScope{
	// 	Name: to.Ptr("my-privatelinkscope"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridCompute/privateLinkScopes/my-privatelinkscope"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"Tag1": to.Ptr("Value1"),
	// 	},
	// 	Properties: &armhybridcompute.PrivateLinkScopeProperties{
	// 		PrivateEndpointConnections: []*armhybridcompute.PrivateEndpointConnectionDataModel{
	// 			{
	// 				Name: to.Ptr("private-endpoint-connection-name"),
	// 				Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 				Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr("Approved"),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 				},
	// 		}},
	// 		PrivateLinkScopeID: to.Ptr("e5dc51d3-92ed-4d7e-947a-775ea79b4919"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_UpdateTagsOnly.json
func ExamplePrivateLinkScopesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().UpdateTags(ctx, "my-resource-group", "my-privatelinkscope", armhybridcompute.TagsResource{
		Tags: map[string]*string{
			"Tag1": to.Ptr("Value1"),
			"Tag2": to.Ptr("Value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkScope = armhybridcompute.PrivateLinkScope{
	// 	Name: to.Ptr("my-privatelinkscope"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"Tag1": to.Ptr("Value1"),
	// 		"Tag2": to.Ptr("Value2"),
	// 	},
	// 	Properties: &armhybridcompute.PrivateLinkScopeProperties{
	// 		PrivateEndpointConnections: []*armhybridcompute.PrivateEndpointConnectionDataModel{
	// 			{
	// 				Name: to.Ptr("private-endpoint-connection-name"),
	// 				Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 				Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr("Approved"),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 				},
	// 		}},
	// 		PrivateLinkScopeID: to.Ptr("e5dc51d3-92ed-4d7e-947a-775ea79b4919"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_GetValidation.json
func ExamplePrivateLinkScopesClient_GetValidationDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().GetValidationDetails(ctx, "wus2", "f5dc51d3-92ed-4d7e-947a-775ea79b4919", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkScopeValidationDetails = armhybridcompute.PrivateLinkScopeValidationDetails{
	// 	ConnectionDetails: []*armhybridcompute.ConnectionDetail{
	// 		{
	// 			GroupID: to.Ptr("groupId"),
	// 			ID: to.Ptr("id"),
	// 			LinkIdentifier: to.Ptr("linkId"),
	// 			MemberName: to.Ptr("memberName"),
	// 			PrivateIPAddress: to.Ptr("ip"),
	// 	}},
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
	// 	PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_GetValidationForMachine.json
func ExamplePrivateLinkScopesClient_GetValidationDetailsForMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().GetValidationDetailsForMachine(ctx, "my-resource-group", "machineName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkScopeValidationDetails = armhybridcompute.PrivateLinkScopeValidationDetails{
	// 	ConnectionDetails: []*armhybridcompute.ConnectionDetail{
	// 		{
	// 			GroupID: to.Ptr("groupId"),
	// 			ID: to.Ptr("id"),
	// 			LinkIdentifier: to.Ptr("linkId"),
	// 			MemberName: to.Ptr("memberName"),
	// 			PrivateIPAddress: to.Ptr("ip"),
	// 	}},
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
	// 	PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
	// }
}

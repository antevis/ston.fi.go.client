# \WalletsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWalletAssets**](WalletsAPI.md#GetWalletAssets) | **Get** /v1/wallets/{addr_str}/assets | 
[**GetWalletFarm**](WalletsAPI.md#GetWalletFarm) | **Get** /v1/wallets/{addr_str}/farms | 
[**GetWalletOperations**](WalletsAPI.md#GetWalletOperations) | **Get** /v1/wallets/{addr_str}/operations | 
[**GetWalletPools**](WalletsAPI.md#GetWalletPools) | **Get** /v1/wallets/{addr_str}/pools | 



## GetWalletAssets

> GetAssetList200Response GetWalletAssets(ctx, addrStr).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	addrStr := "addrStr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletAssets(context.Background(), addrStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletAssets`: GetAssetList200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addrStr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAssetList200Response**](GetAssetList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletFarm

> GetFarmList200Response GetWalletFarm(ctx, addrStr).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	addrStr := "addrStr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletFarm(context.Background(), addrStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletFarm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletFarm`: GetFarmList200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletFarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addrStr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFarmList200Response**](GetFarmList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletOperations

> GetOperationStats200Response GetWalletOperations(ctx, addrStr).Since(since).Until(until).OpType(opType).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	since := "2023-06-01T12:34:56" // string | Time since the stats are requested (YYYY-MM-DDTHH:MM:SS)
	until := "2023-06-02T23:59:59" // string | Time until pool stats are requested (YYYY-MM-DDTHH:MM:SS)
	addrStr := "addrStr_example" // string | 
	opType := "SendLiquidity" // string | target op type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletOperations(context.Background(), addrStr).Since(since).Until(until).OpType(opType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletOperations`: GetOperationStats200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addrStr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Time since the stats are requested (YYYY-MM-DDTHH:MM:SS) | 
 **until** | **string** | Time until pool stats are requested (YYYY-MM-DDTHH:MM:SS) | 

 **opType** | **string** | target op type | 

### Return type

[**GetOperationStats200Response**](GetOperationStats200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletPools

> GetPoolList200Response GetWalletPools(ctx, addrStr).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	addrStr := "addrStr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletPools(context.Background(), addrStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletPools`: GetPoolList200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addrStr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPoolList200Response**](GetPoolList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


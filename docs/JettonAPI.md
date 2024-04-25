# \JettonAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWalletAddress**](JettonAPI.md#GetWalletAddress) | **Get** /v1/jetton/{addr_str}/address | 



## GetWalletAddress

> GetWalletAddress200Response GetWalletAddress(ctx, addrStr).OwnerAddress(ownerAddress).Execute()



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
	ownerAddress := "EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez" // string | Address of the owner
	addrStr := "addrStr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JettonAPI.GetWalletAddress(context.Background(), addrStr).OwnerAddress(ownerAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JettonAPI.GetWalletAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletAddress`: GetWalletAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `JettonAPI.GetWalletAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addrStr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerAddress** | **string** | Address of the owner | 


### Return type

[**GetWalletAddress200Response**](GetWalletAddress200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


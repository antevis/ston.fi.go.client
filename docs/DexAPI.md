# \DexAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DexReverseSimulateSwap**](DexAPI.md#DexReverseSimulateSwap) | **Post** /v1/reverse_swap/simulate | 
[**DexSimulateSwap**](DexAPI.md#DexSimulateSwap) | **Post** /v1/swap/simulate | 
[**DexSwapStatus**](DexAPI.md#DexSwapStatus) | **Get** /v1/swap/status | 
[**GetAssetList**](DexAPI.md#GetAssetList) | **Get** /v1/assets | 
[**GetFarmList**](DexAPI.md#GetFarmList) | **Get** /v1/farms | 
[**GetPoolList**](DexAPI.md#GetPoolList) | **Get** /v1/pools | 



## DexReverseSimulateSwap

> DexReverseSimulateSwap200Response DexReverseSimulateSwap(ctx).OfferAddress(offerAddress).AskAddress(askAddress).Units(units).SlippageTolerance(slippageTolerance).ReferralAddress(referralAddress).Execute()



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
	offerAddress := "EQBynBO23ywHy_CgarY9NK9FTz0yDsG82PtcbSTQgGoXwiuA" // string | The address of the token we want to sell
	askAddress := "EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez" // string | The address of the token we want to buy
	units := "300" // string | Number of token units we want to sell
	slippageTolerance := "0.001" // string | The maximum possible difference between the rates that we expect and which will actually be, in fractions (for example, 0.001 is 0.1%)
	referralAddress := "referralAddress_example" // string | Referral address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.DexReverseSimulateSwap(context.Background()).OfferAddress(offerAddress).AskAddress(askAddress).Units(units).SlippageTolerance(slippageTolerance).ReferralAddress(referralAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.DexReverseSimulateSwap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DexReverseSimulateSwap`: DexReverseSimulateSwap200Response
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.DexReverseSimulateSwap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDexReverseSimulateSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offerAddress** | **string** | The address of the token we want to sell | 
 **askAddress** | **string** | The address of the token we want to buy | 
 **units** | **string** | Number of token units we want to sell | 
 **slippageTolerance** | **string** | The maximum possible difference between the rates that we expect and which will actually be, in fractions (for example, 0.001 is 0.1%) | 
 **referralAddress** | **string** | Referral address | 

### Return type

[**DexReverseSimulateSwap200Response**](DexReverseSimulateSwap200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DexSimulateSwap

> DexReverseSimulateSwap200Response DexSimulateSwap(ctx).OfferAddress(offerAddress).AskAddress(askAddress).Units(units).SlippageTolerance(slippageTolerance).ReferralAddress(referralAddress).Execute()



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
	offerAddress := "EQBynBO23ywHy_CgarY9NK9FTz0yDsG82PtcbSTQgGoXwiuA" // string | The address of the token we want to sell
	askAddress := "EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez" // string | The address of the token we want to buy
	units := "300" // string | Number of token units we want to sell
	slippageTolerance := "0.001" // string | The maximum possible difference between the rates that we expect and which will actually be, in fractions (for example, 0.001 is 0.1%)
	referralAddress := "referralAddress_example" // string | Referral address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.DexSimulateSwap(context.Background()).OfferAddress(offerAddress).AskAddress(askAddress).Units(units).SlippageTolerance(slippageTolerance).ReferralAddress(referralAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.DexSimulateSwap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DexSimulateSwap`: DexReverseSimulateSwap200Response
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.DexSimulateSwap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDexSimulateSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offerAddress** | **string** | The address of the token we want to sell | 
 **askAddress** | **string** | The address of the token we want to buy | 
 **units** | **string** | Number of token units we want to sell | 
 **slippageTolerance** | **string** | The maximum possible difference between the rates that we expect and which will actually be, in fractions (for example, 0.001 is 0.1%) | 
 **referralAddress** | **string** | Referral address | 

### Return type

[**DexReverseSimulateSwap200Response**](DexReverseSimulateSwap200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DexSwapStatus

> DexSwapStatus200Response DexSwapStatus(ctx).RouterAddress(routerAddress).OwnerAddress(ownerAddress).QueryId(queryId).Execute()



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
	routerAddress := "EQB3ncyBUTjZUA5EnFKR5_EnOMI9V1tTEAAPaiU71gc4TiUt" // string | Address of the operation router
	ownerAddress := "EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez" // string | Owner`s wallet address
	queryId := "1" // string | Id of operation status query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.DexSwapStatus(context.Background()).RouterAddress(routerAddress).OwnerAddress(ownerAddress).QueryId(queryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.DexSwapStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DexSwapStatus`: DexSwapStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.DexSwapStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDexSwapStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerAddress** | **string** | Address of the operation router | 
 **ownerAddress** | **string** | Owner&#x60;s wallet address | 
 **queryId** | **string** | Id of operation status query | 

### Return type

[**DexSwapStatus200Response**](DexSwapStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetList

> GetAssetList200Response GetAssetList(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.GetAssetList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.GetAssetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetList`: GetAssetList200Response
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.GetAssetList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetListRequest struct via the builder pattern


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


## GetFarmList

> GetFarmList200Response GetFarmList(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.GetFarmList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.GetFarmList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFarmList`: GetFarmList200Response
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.GetFarmList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFarmListRequest struct via the builder pattern


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


## GetPoolList

> GetPoolList200Response GetPoolList(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.GetPoolList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.GetPoolList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolList`: GetPoolList200Response
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.GetPoolList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolListRequest struct via the builder pattern


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


# \StatsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDexStats**](StatsAPI.md#GetDexStats) | **Get** /v1/stats/dex | 
[**GetOperationStats**](StatsAPI.md#GetOperationStats) | **Get** /v1/stats/operations | 
[**GetPoolStats**](StatsAPI.md#GetPoolStats) | **Get** /v1/stats/pool | 



## GetDexStats

> GetDexStats(ctx).Since(since).Until(until).Execute()



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
	since := "2023-06-01T12:34:56" // string | Time since stats are requested (YYYY-MM-DDTHH:MM:SS) (optional)
	until := "2023-06-02T23:59:59" // string | Time until stats are requested (YYYY-MM-DDTHH:MM:SS) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StatsAPI.GetDexStats(context.Background()).Since(since).Until(until).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetDexStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDexStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Time since stats are requested (YYYY-MM-DDTHH:MM:SS) | 
 **until** | **string** | Time until stats are requested (YYYY-MM-DDTHH:MM:SS) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperationStats

> GetOperationStats200Response GetOperationStats(ctx).Since(since).Until(until).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetOperationStats(context.Background()).Since(since).Until(until).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetOperationStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperationStats`: GetOperationStats200Response
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetOperationStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Time since the stats are requested (YYYY-MM-DDTHH:MM:SS) | 
 **until** | **string** | Time until pool stats are requested (YYYY-MM-DDTHH:MM:SS) | 

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


## GetPoolStats

> GetPoolStats200Response GetPoolStats(ctx).Since(since).Until(until).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetPoolStats(context.Background()).Since(since).Until(until).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetPoolStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolStats`: GetPoolStats200Response
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetPoolStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Time since the stats are requested (YYYY-MM-DDTHH:MM:SS) | 
 **until** | **string** | Time until pool stats are requested (YYYY-MM-DDTHH:MM:SS) | 

### Return type

[**GetPoolStats200Response**](GetPoolStats200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


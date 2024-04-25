/*
ston-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DexAPIService DexAPI service
type DexAPIService service

type ApiDexReverseSimulateSwapRequest struct {
	ctx context.Context
	ApiService *DexAPIService
	offerAddress *string
	askAddress *string
	units *string
	slippageTolerance *string
	referralAddress *string
}

// The address of the token we want to sell
func (r ApiDexReverseSimulateSwapRequest) OfferAddress(offerAddress string) ApiDexReverseSimulateSwapRequest {
	r.offerAddress = &offerAddress
	return r
}

// The address of the token we want to buy
func (r ApiDexReverseSimulateSwapRequest) AskAddress(askAddress string) ApiDexReverseSimulateSwapRequest {
	r.askAddress = &askAddress
	return r
}

// Number of token units we want to sell
func (r ApiDexReverseSimulateSwapRequest) Units(units string) ApiDexReverseSimulateSwapRequest {
	r.units = &units
	return r
}

// The maximum possible difference between the rates that we expect and which will actually be, in fractions (for example, 0.001 is 0.1%)
func (r ApiDexReverseSimulateSwapRequest) SlippageTolerance(slippageTolerance string) ApiDexReverseSimulateSwapRequest {
	r.slippageTolerance = &slippageTolerance
	return r
}

// Referral address
func (r ApiDexReverseSimulateSwapRequest) ReferralAddress(referralAddress string) ApiDexReverseSimulateSwapRequest {
	r.referralAddress = &referralAddress
	return r
}

func (r ApiDexReverseSimulateSwapRequest) Execute() (*DexReverseSimulateSwap200Response, *http.Response, error) {
	return r.ApiService.DexReverseSimulateSwapExecute(r)
}

/*
DexReverseSimulateSwap Method for DexReverseSimulateSwap

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDexReverseSimulateSwapRequest
*/
func (a *DexAPIService) DexReverseSimulateSwap(ctx context.Context) ApiDexReverseSimulateSwapRequest {
	return ApiDexReverseSimulateSwapRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DexReverseSimulateSwap200Response
func (a *DexAPIService) DexReverseSimulateSwapExecute(r ApiDexReverseSimulateSwapRequest) (*DexReverseSimulateSwap200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DexReverseSimulateSwap200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DexAPIService.DexReverseSimulateSwap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reverse_swap/simulate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.offerAddress == nil {
		return localVarReturnValue, nil, reportError("offerAddress is required and must be specified")
	}
	if r.askAddress == nil {
		return localVarReturnValue, nil, reportError("askAddress is required and must be specified")
	}
	if r.units == nil {
		return localVarReturnValue, nil, reportError("units is required and must be specified")
	}
	if r.slippageTolerance == nil {
		return localVarReturnValue, nil, reportError("slippageTolerance is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "offer_address", r.offerAddress, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "ask_address", r.askAddress, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "units", r.units, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "slippage_tolerance", r.slippageTolerance, "")
	if r.referralAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "referral_address", r.referralAddress, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDexSimulateSwapRequest struct {
	ctx context.Context
	ApiService *DexAPIService
	offerAddress *string
	askAddress *string
	units *string
	slippageTolerance *string
	referralAddress *string
}

// The address of the token we want to sell
func (r ApiDexSimulateSwapRequest) OfferAddress(offerAddress string) ApiDexSimulateSwapRequest {
	r.offerAddress = &offerAddress
	return r
}

// The address of the token we want to buy
func (r ApiDexSimulateSwapRequest) AskAddress(askAddress string) ApiDexSimulateSwapRequest {
	r.askAddress = &askAddress
	return r
}

// Number of token units we want to sell
func (r ApiDexSimulateSwapRequest) Units(units string) ApiDexSimulateSwapRequest {
	r.units = &units
	return r
}

// The maximum possible difference between the rates that we expect and which will actually be, in fractions (for example, 0.001 is 0.1%)
func (r ApiDexSimulateSwapRequest) SlippageTolerance(slippageTolerance string) ApiDexSimulateSwapRequest {
	r.slippageTolerance = &slippageTolerance
	return r
}

// Referral address
func (r ApiDexSimulateSwapRequest) ReferralAddress(referralAddress string) ApiDexSimulateSwapRequest {
	r.referralAddress = &referralAddress
	return r
}

func (r ApiDexSimulateSwapRequest) Execute() (*DexReverseSimulateSwap200Response, *http.Response, error) {
	return r.ApiService.DexSimulateSwapExecute(r)
}

/*
DexSimulateSwap Method for DexSimulateSwap

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDexSimulateSwapRequest
*/
func (a *DexAPIService) DexSimulateSwap(ctx context.Context) ApiDexSimulateSwapRequest {
	return ApiDexSimulateSwapRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DexReverseSimulateSwap200Response
func (a *DexAPIService) DexSimulateSwapExecute(r ApiDexSimulateSwapRequest) (*DexReverseSimulateSwap200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DexReverseSimulateSwap200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DexAPIService.DexSimulateSwap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/swap/simulate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.offerAddress == nil {
		return localVarReturnValue, nil, reportError("offerAddress is required and must be specified")
	}
	if r.askAddress == nil {
		return localVarReturnValue, nil, reportError("askAddress is required and must be specified")
	}
	if r.units == nil {
		return localVarReturnValue, nil, reportError("units is required and must be specified")
	}
	if r.slippageTolerance == nil {
		return localVarReturnValue, nil, reportError("slippageTolerance is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "offer_address", r.offerAddress, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "ask_address", r.askAddress, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "units", r.units, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "slippage_tolerance", r.slippageTolerance, "")
	if r.referralAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "referral_address", r.referralAddress, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDexSwapStatusRequest struct {
	ctx context.Context
	ApiService *DexAPIService
	routerAddress *string
	ownerAddress *string
	queryId *string
}

// Address of the operation router
func (r ApiDexSwapStatusRequest) RouterAddress(routerAddress string) ApiDexSwapStatusRequest {
	r.routerAddress = &routerAddress
	return r
}

// Owner&#x60;s wallet address
func (r ApiDexSwapStatusRequest) OwnerAddress(ownerAddress string) ApiDexSwapStatusRequest {
	r.ownerAddress = &ownerAddress
	return r
}

// Id of operation status query
func (r ApiDexSwapStatusRequest) QueryId(queryId string) ApiDexSwapStatusRequest {
	r.queryId = &queryId
	return r
}

func (r ApiDexSwapStatusRequest) Execute() (*DexSwapStatus200Response, *http.Response, error) {
	return r.ApiService.DexSwapStatusExecute(r)
}

/*
DexSwapStatus Method for DexSwapStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDexSwapStatusRequest
*/
func (a *DexAPIService) DexSwapStatus(ctx context.Context) ApiDexSwapStatusRequest {
	return ApiDexSwapStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DexSwapStatus200Response
func (a *DexAPIService) DexSwapStatusExecute(r ApiDexSwapStatusRequest) (*DexSwapStatus200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DexSwapStatus200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DexAPIService.DexSwapStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/swap/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.routerAddress == nil {
		return localVarReturnValue, nil, reportError("routerAddress is required and must be specified")
	}
	if r.ownerAddress == nil {
		return localVarReturnValue, nil, reportError("ownerAddress is required and must be specified")
	}
	if r.queryId == nil {
		return localVarReturnValue, nil, reportError("queryId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "router_address", r.routerAddress, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "owner_address", r.ownerAddress, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "query_id", r.queryId, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAssetListRequest struct {
	ctx context.Context
	ApiService *DexAPIService
}

func (r ApiGetAssetListRequest) Execute() (*GetAssetList200Response, *http.Response, error) {
	return r.ApiService.GetAssetListExecute(r)
}

/*
GetAssetList Method for GetAssetList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAssetListRequest
*/
func (a *DexAPIService) GetAssetList(ctx context.Context) ApiGetAssetListRequest {
	return ApiGetAssetListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAssetList200Response
func (a *DexAPIService) GetAssetListExecute(r ApiGetAssetListRequest) (*GetAssetList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAssetList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DexAPIService.GetAssetList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/assets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFarmListRequest struct {
	ctx context.Context
	ApiService *DexAPIService
}

func (r ApiGetFarmListRequest) Execute() (*GetFarmList200Response, *http.Response, error) {
	return r.ApiService.GetFarmListExecute(r)
}

/*
GetFarmList Method for GetFarmList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFarmListRequest
*/
func (a *DexAPIService) GetFarmList(ctx context.Context) ApiGetFarmListRequest {
	return ApiGetFarmListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFarmList200Response
func (a *DexAPIService) GetFarmListExecute(r ApiGetFarmListRequest) (*GetFarmList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFarmList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DexAPIService.GetFarmList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/farms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPoolListRequest struct {
	ctx context.Context
	ApiService *DexAPIService
}

func (r ApiGetPoolListRequest) Execute() (*GetPoolList200Response, *http.Response, error) {
	return r.ApiService.GetPoolListExecute(r)
}

/*
GetPoolList Method for GetPoolList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPoolListRequest
*/
func (a *DexAPIService) GetPoolList(ctx context.Context) ApiGetPoolListRequest {
	return ApiGetPoolListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPoolList200Response
func (a *DexAPIService) GetPoolListExecute(r ApiGetPoolListRequest) (*GetPoolList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPoolList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DexAPIService.GetPoolList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/pools"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

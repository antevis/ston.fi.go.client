# DexReverseSimulateSwap200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskAddress** | **string** | Address of the wallet to which the operation is simulated | 
**AskUnits** | **string** | Token units asked | 
**FeeAddress** | **string** | Fee address | 
**FeePercent** | **string** | fee percent | 
**FeeUnits** | **string** | Fee units | 
**MinAskUnits** | **string** | Minimal amount of ask units | 
**OfferAddress** | **string** | Address of the wallet from which the operation is simulated | 
**OfferUnits** | **string** | Token units offered | 
**PoolAddress** | **string** | Address of the pool | 
**PriceImpact** | **string** | Price impact | 
**RouterAddress** | **string** | Address of the operation router | 
**SlippageTolerance** | **string** | Difference between the original price of order and the final price | 
**SwapRate** | **string** | Swap rate | 

## Methods

### NewDexReverseSimulateSwap200Response

`func NewDexReverseSimulateSwap200Response(askAddress string, askUnits string, feeAddress string, feePercent string, feeUnits string, minAskUnits string, offerAddress string, offerUnits string, poolAddress string, priceImpact string, routerAddress string, slippageTolerance string, swapRate string, ) *DexReverseSimulateSwap200Response`

NewDexReverseSimulateSwap200Response instantiates a new DexReverseSimulateSwap200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexReverseSimulateSwap200ResponseWithDefaults

`func NewDexReverseSimulateSwap200ResponseWithDefaults() *DexReverseSimulateSwap200Response`

NewDexReverseSimulateSwap200ResponseWithDefaults instantiates a new DexReverseSimulateSwap200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskAddress

`func (o *DexReverseSimulateSwap200Response) GetAskAddress() string`

GetAskAddress returns the AskAddress field if non-nil, zero value otherwise.

### GetAskAddressOk

`func (o *DexReverseSimulateSwap200Response) GetAskAddressOk() (*string, bool)`

GetAskAddressOk returns a tuple with the AskAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskAddress

`func (o *DexReverseSimulateSwap200Response) SetAskAddress(v string)`

SetAskAddress sets AskAddress field to given value.


### GetAskUnits

`func (o *DexReverseSimulateSwap200Response) GetAskUnits() string`

GetAskUnits returns the AskUnits field if non-nil, zero value otherwise.

### GetAskUnitsOk

`func (o *DexReverseSimulateSwap200Response) GetAskUnitsOk() (*string, bool)`

GetAskUnitsOk returns a tuple with the AskUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskUnits

`func (o *DexReverseSimulateSwap200Response) SetAskUnits(v string)`

SetAskUnits sets AskUnits field to given value.


### GetFeeAddress

`func (o *DexReverseSimulateSwap200Response) GetFeeAddress() string`

GetFeeAddress returns the FeeAddress field if non-nil, zero value otherwise.

### GetFeeAddressOk

`func (o *DexReverseSimulateSwap200Response) GetFeeAddressOk() (*string, bool)`

GetFeeAddressOk returns a tuple with the FeeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAddress

`func (o *DexReverseSimulateSwap200Response) SetFeeAddress(v string)`

SetFeeAddress sets FeeAddress field to given value.


### GetFeePercent

`func (o *DexReverseSimulateSwap200Response) GetFeePercent() string`

GetFeePercent returns the FeePercent field if non-nil, zero value otherwise.

### GetFeePercentOk

`func (o *DexReverseSimulateSwap200Response) GetFeePercentOk() (*string, bool)`

GetFeePercentOk returns a tuple with the FeePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercent

`func (o *DexReverseSimulateSwap200Response) SetFeePercent(v string)`

SetFeePercent sets FeePercent field to given value.


### GetFeeUnits

`func (o *DexReverseSimulateSwap200Response) GetFeeUnits() string`

GetFeeUnits returns the FeeUnits field if non-nil, zero value otherwise.

### GetFeeUnitsOk

`func (o *DexReverseSimulateSwap200Response) GetFeeUnitsOk() (*string, bool)`

GetFeeUnitsOk returns a tuple with the FeeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeUnits

`func (o *DexReverseSimulateSwap200Response) SetFeeUnits(v string)`

SetFeeUnits sets FeeUnits field to given value.


### GetMinAskUnits

`func (o *DexReverseSimulateSwap200Response) GetMinAskUnits() string`

GetMinAskUnits returns the MinAskUnits field if non-nil, zero value otherwise.

### GetMinAskUnitsOk

`func (o *DexReverseSimulateSwap200Response) GetMinAskUnitsOk() (*string, bool)`

GetMinAskUnitsOk returns a tuple with the MinAskUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAskUnits

`func (o *DexReverseSimulateSwap200Response) SetMinAskUnits(v string)`

SetMinAskUnits sets MinAskUnits field to given value.


### GetOfferAddress

`func (o *DexReverseSimulateSwap200Response) GetOfferAddress() string`

GetOfferAddress returns the OfferAddress field if non-nil, zero value otherwise.

### GetOfferAddressOk

`func (o *DexReverseSimulateSwap200Response) GetOfferAddressOk() (*string, bool)`

GetOfferAddressOk returns a tuple with the OfferAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferAddress

`func (o *DexReverseSimulateSwap200Response) SetOfferAddress(v string)`

SetOfferAddress sets OfferAddress field to given value.


### GetOfferUnits

`func (o *DexReverseSimulateSwap200Response) GetOfferUnits() string`

GetOfferUnits returns the OfferUnits field if non-nil, zero value otherwise.

### GetOfferUnitsOk

`func (o *DexReverseSimulateSwap200Response) GetOfferUnitsOk() (*string, bool)`

GetOfferUnitsOk returns a tuple with the OfferUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferUnits

`func (o *DexReverseSimulateSwap200Response) SetOfferUnits(v string)`

SetOfferUnits sets OfferUnits field to given value.


### GetPoolAddress

`func (o *DexReverseSimulateSwap200Response) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *DexReverseSimulateSwap200Response) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *DexReverseSimulateSwap200Response) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetPriceImpact

`func (o *DexReverseSimulateSwap200Response) GetPriceImpact() string`

GetPriceImpact returns the PriceImpact field if non-nil, zero value otherwise.

### GetPriceImpactOk

`func (o *DexReverseSimulateSwap200Response) GetPriceImpactOk() (*string, bool)`

GetPriceImpactOk returns a tuple with the PriceImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceImpact

`func (o *DexReverseSimulateSwap200Response) SetPriceImpact(v string)`

SetPriceImpact sets PriceImpact field to given value.


### GetRouterAddress

`func (o *DexReverseSimulateSwap200Response) GetRouterAddress() string`

GetRouterAddress returns the RouterAddress field if non-nil, zero value otherwise.

### GetRouterAddressOk

`func (o *DexReverseSimulateSwap200Response) GetRouterAddressOk() (*string, bool)`

GetRouterAddressOk returns a tuple with the RouterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterAddress

`func (o *DexReverseSimulateSwap200Response) SetRouterAddress(v string)`

SetRouterAddress sets RouterAddress field to given value.


### GetSlippageTolerance

`func (o *DexReverseSimulateSwap200Response) GetSlippageTolerance() string`

GetSlippageTolerance returns the SlippageTolerance field if non-nil, zero value otherwise.

### GetSlippageToleranceOk

`func (o *DexReverseSimulateSwap200Response) GetSlippageToleranceOk() (*string, bool)`

GetSlippageToleranceOk returns a tuple with the SlippageTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageTolerance

`func (o *DexReverseSimulateSwap200Response) SetSlippageTolerance(v string)`

SetSlippageTolerance sets SlippageTolerance field to given value.


### GetSwapRate

`func (o *DexReverseSimulateSwap200Response) GetSwapRate() string`

GetSwapRate returns the SwapRate field if non-nil, zero value otherwise.

### GetSwapRateOk

`func (o *DexReverseSimulateSwap200Response) GetSwapRateOk() (*string, bool)`

GetSwapRateOk returns a tuple with the SwapRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapRate

`func (o *DexReverseSimulateSwap200Response) SetSwapRate(v string)`

SetSwapRate sets SwapRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



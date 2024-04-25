# PoolStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | Pointer to **NullableString** | Annual percentage yield | [optional] 
**BaseId** | **string** | Base asset id | 
**BaseLiquidity** | **string** | Amount of liquidity of base asset | 
**BaseName** | **string** | Base asset name | 
**BaseSymbol** | **string** | Base asset symbol | 
**BaseVolume** | **string** | volume of base asset | 
**LastPrice** | **string** | Last asset quota price in base assets | 
**LpPrice** | **string** | [DEPRECATED] Liquidity pool token price | 
**LpPriceUsd** | **string** | Liquidity pool token price | 
**PoolAddress** | **string** | Address of the pool | 
**QuoteId** | **string** | Quote asset id | 
**QuoteLiquidity** | **string** | Amount of liquidity of quote asset | 
**QuoteName** | **string** | Quote asset name | 
**QuoteSymbol** | **string** | Quote asset symbol | 
**QuoteVolume** | **string** | Volume of quote asset | 
**RouterAddress** | **string** | Address of router | 
**Url** | **string** | URL of the pool | 

## Methods

### NewPoolStats

`func NewPoolStats(baseId string, baseLiquidity string, baseName string, baseSymbol string, baseVolume string, lastPrice string, lpPrice string, lpPriceUsd string, poolAddress string, quoteId string, quoteLiquidity string, quoteName string, quoteSymbol string, quoteVolume string, routerAddress string, url string, ) *PoolStats`

NewPoolStats instantiates a new PoolStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolStatsWithDefaults

`func NewPoolStatsWithDefaults() *PoolStats`

NewPoolStatsWithDefaults instantiates a new PoolStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApy

`func (o *PoolStats) GetApy() string`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *PoolStats) GetApyOk() (*string, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *PoolStats) SetApy(v string)`

SetApy sets Apy field to given value.

### HasApy

`func (o *PoolStats) HasApy() bool`

HasApy returns a boolean if a field has been set.

### SetApyNil

`func (o *PoolStats) SetApyNil(b bool)`

 SetApyNil sets the value for Apy to be an explicit nil

### UnsetApy
`func (o *PoolStats) UnsetApy()`

UnsetApy ensures that no value is present for Apy, not even an explicit nil
### GetBaseId

`func (o *PoolStats) GetBaseId() string`

GetBaseId returns the BaseId field if non-nil, zero value otherwise.

### GetBaseIdOk

`func (o *PoolStats) GetBaseIdOk() (*string, bool)`

GetBaseIdOk returns a tuple with the BaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseId

`func (o *PoolStats) SetBaseId(v string)`

SetBaseId sets BaseId field to given value.


### GetBaseLiquidity

`func (o *PoolStats) GetBaseLiquidity() string`

GetBaseLiquidity returns the BaseLiquidity field if non-nil, zero value otherwise.

### GetBaseLiquidityOk

`func (o *PoolStats) GetBaseLiquidityOk() (*string, bool)`

GetBaseLiquidityOk returns a tuple with the BaseLiquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseLiquidity

`func (o *PoolStats) SetBaseLiquidity(v string)`

SetBaseLiquidity sets BaseLiquidity field to given value.


### GetBaseName

`func (o *PoolStats) GetBaseName() string`

GetBaseName returns the BaseName field if non-nil, zero value otherwise.

### GetBaseNameOk

`func (o *PoolStats) GetBaseNameOk() (*string, bool)`

GetBaseNameOk returns a tuple with the BaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseName

`func (o *PoolStats) SetBaseName(v string)`

SetBaseName sets BaseName field to given value.


### GetBaseSymbol

`func (o *PoolStats) GetBaseSymbol() string`

GetBaseSymbol returns the BaseSymbol field if non-nil, zero value otherwise.

### GetBaseSymbolOk

`func (o *PoolStats) GetBaseSymbolOk() (*string, bool)`

GetBaseSymbolOk returns a tuple with the BaseSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSymbol

`func (o *PoolStats) SetBaseSymbol(v string)`

SetBaseSymbol sets BaseSymbol field to given value.


### GetBaseVolume

`func (o *PoolStats) GetBaseVolume() string`

GetBaseVolume returns the BaseVolume field if non-nil, zero value otherwise.

### GetBaseVolumeOk

`func (o *PoolStats) GetBaseVolumeOk() (*string, bool)`

GetBaseVolumeOk returns a tuple with the BaseVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVolume

`func (o *PoolStats) SetBaseVolume(v string)`

SetBaseVolume sets BaseVolume field to given value.


### GetLastPrice

`func (o *PoolStats) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *PoolStats) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *PoolStats) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.


### GetLpPrice

`func (o *PoolStats) GetLpPrice() string`

GetLpPrice returns the LpPrice field if non-nil, zero value otherwise.

### GetLpPriceOk

`func (o *PoolStats) GetLpPriceOk() (*string, bool)`

GetLpPriceOk returns a tuple with the LpPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpPrice

`func (o *PoolStats) SetLpPrice(v string)`

SetLpPrice sets LpPrice field to given value.


### GetLpPriceUsd

`func (o *PoolStats) GetLpPriceUsd() string`

GetLpPriceUsd returns the LpPriceUsd field if non-nil, zero value otherwise.

### GetLpPriceUsdOk

`func (o *PoolStats) GetLpPriceUsdOk() (*string, bool)`

GetLpPriceUsdOk returns a tuple with the LpPriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpPriceUsd

`func (o *PoolStats) SetLpPriceUsd(v string)`

SetLpPriceUsd sets LpPriceUsd field to given value.


### GetPoolAddress

`func (o *PoolStats) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *PoolStats) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *PoolStats) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetQuoteId

`func (o *PoolStats) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *PoolStats) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *PoolStats) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteLiquidity

`func (o *PoolStats) GetQuoteLiquidity() string`

GetQuoteLiquidity returns the QuoteLiquidity field if non-nil, zero value otherwise.

### GetQuoteLiquidityOk

`func (o *PoolStats) GetQuoteLiquidityOk() (*string, bool)`

GetQuoteLiquidityOk returns a tuple with the QuoteLiquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteLiquidity

`func (o *PoolStats) SetQuoteLiquidity(v string)`

SetQuoteLiquidity sets QuoteLiquidity field to given value.


### GetQuoteName

`func (o *PoolStats) GetQuoteName() string`

GetQuoteName returns the QuoteName field if non-nil, zero value otherwise.

### GetQuoteNameOk

`func (o *PoolStats) GetQuoteNameOk() (*string, bool)`

GetQuoteNameOk returns a tuple with the QuoteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteName

`func (o *PoolStats) SetQuoteName(v string)`

SetQuoteName sets QuoteName field to given value.


### GetQuoteSymbol

`func (o *PoolStats) GetQuoteSymbol() string`

GetQuoteSymbol returns the QuoteSymbol field if non-nil, zero value otherwise.

### GetQuoteSymbolOk

`func (o *PoolStats) GetQuoteSymbolOk() (*string, bool)`

GetQuoteSymbolOk returns a tuple with the QuoteSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSymbol

`func (o *PoolStats) SetQuoteSymbol(v string)`

SetQuoteSymbol sets QuoteSymbol field to given value.


### GetQuoteVolume

`func (o *PoolStats) GetQuoteVolume() string`

GetQuoteVolume returns the QuoteVolume field if non-nil, zero value otherwise.

### GetQuoteVolumeOk

`func (o *PoolStats) GetQuoteVolumeOk() (*string, bool)`

GetQuoteVolumeOk returns a tuple with the QuoteVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVolume

`func (o *PoolStats) SetQuoteVolume(v string)`

SetQuoteVolume sets QuoteVolume field to given value.


### GetRouterAddress

`func (o *PoolStats) GetRouterAddress() string`

GetRouterAddress returns the RouterAddress field if non-nil, zero value otherwise.

### GetRouterAddressOk

`func (o *PoolStats) GetRouterAddressOk() (*string, bool)`

GetRouterAddressOk returns a tuple with the RouterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterAddress

`func (o *PoolStats) SetRouterAddress(v string)`

SetRouterAddress sets RouterAddress field to given value.


### GetUrl

`func (o *PoolStats) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PoolStats) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PoolStats) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



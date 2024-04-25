# AssetInfoSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **NullableString** | Asset balance | [optional] 
**Blacklisted** | **bool** | Asset is blacklisted | 
**Community** | **bool** | Asset is community created | 
**ContractAddress** | **string** | Address of smart contract | 
**Decimals** | **int32** | Number of decimal places used to represent fractional amounts of an asset | 
**DefaultSymbol** | **bool** | Asset has default symbol | 
**Deprecated** | **bool** | Asset is deprecated | 
**DexPriceUsd** | Pointer to **NullableString** | Ston internal asset price in USD | [optional] 
**DexUsdPrice** | Pointer to **NullableString** | [DEPRECATED] Ston internal asset price in USD | [optional] 
**DisplayName** | Pointer to **NullableString** | Displayable name | [optional] 
**ImageUrl** | Pointer to **NullableString** | URL to asset image | [optional] 
**Kind** | [**AssetKindSchema**](AssetKindSchema.md) |  | 
**Symbol** | **string** | Asset symbol | 
**ThirdPartyPriceUsd** | Pointer to **NullableString** | Asset price in USD | [optional] 
**ThirdPartyUsdPrice** | Pointer to **NullableString** | [DEPRECATED] Asset price in USD | [optional] 
**WalletAddress** | Pointer to **NullableString** | Asset wallet address | [optional] 

## Methods

### NewAssetInfoSchema

`func NewAssetInfoSchema(blacklisted bool, community bool, contractAddress string, decimals int32, defaultSymbol bool, deprecated bool, kind AssetKindSchema, symbol string, ) *AssetInfoSchema`

NewAssetInfoSchema instantiates a new AssetInfoSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetInfoSchemaWithDefaults

`func NewAssetInfoSchemaWithDefaults() *AssetInfoSchema`

NewAssetInfoSchemaWithDefaults instantiates a new AssetInfoSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *AssetInfoSchema) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AssetInfoSchema) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AssetInfoSchema) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AssetInfoSchema) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *AssetInfoSchema) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *AssetInfoSchema) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetBlacklisted

`func (o *AssetInfoSchema) GetBlacklisted() bool`

GetBlacklisted returns the Blacklisted field if non-nil, zero value otherwise.

### GetBlacklistedOk

`func (o *AssetInfoSchema) GetBlacklistedOk() (*bool, bool)`

GetBlacklistedOk returns a tuple with the Blacklisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklisted

`func (o *AssetInfoSchema) SetBlacklisted(v bool)`

SetBlacklisted sets Blacklisted field to given value.


### GetCommunity

`func (o *AssetInfoSchema) GetCommunity() bool`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *AssetInfoSchema) GetCommunityOk() (*bool, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *AssetInfoSchema) SetCommunity(v bool)`

SetCommunity sets Community field to given value.


### GetContractAddress

`func (o *AssetInfoSchema) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *AssetInfoSchema) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *AssetInfoSchema) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetDecimals

`func (o *AssetInfoSchema) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *AssetInfoSchema) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *AssetInfoSchema) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetDefaultSymbol

`func (o *AssetInfoSchema) GetDefaultSymbol() bool`

GetDefaultSymbol returns the DefaultSymbol field if non-nil, zero value otherwise.

### GetDefaultSymbolOk

`func (o *AssetInfoSchema) GetDefaultSymbolOk() (*bool, bool)`

GetDefaultSymbolOk returns a tuple with the DefaultSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSymbol

`func (o *AssetInfoSchema) SetDefaultSymbol(v bool)`

SetDefaultSymbol sets DefaultSymbol field to given value.


### GetDeprecated

`func (o *AssetInfoSchema) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AssetInfoSchema) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AssetInfoSchema) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetDexPriceUsd

`func (o *AssetInfoSchema) GetDexPriceUsd() string`

GetDexPriceUsd returns the DexPriceUsd field if non-nil, zero value otherwise.

### GetDexPriceUsdOk

`func (o *AssetInfoSchema) GetDexPriceUsdOk() (*string, bool)`

GetDexPriceUsdOk returns a tuple with the DexPriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexPriceUsd

`func (o *AssetInfoSchema) SetDexPriceUsd(v string)`

SetDexPriceUsd sets DexPriceUsd field to given value.

### HasDexPriceUsd

`func (o *AssetInfoSchema) HasDexPriceUsd() bool`

HasDexPriceUsd returns a boolean if a field has been set.

### SetDexPriceUsdNil

`func (o *AssetInfoSchema) SetDexPriceUsdNil(b bool)`

 SetDexPriceUsdNil sets the value for DexPriceUsd to be an explicit nil

### UnsetDexPriceUsd
`func (o *AssetInfoSchema) UnsetDexPriceUsd()`

UnsetDexPriceUsd ensures that no value is present for DexPriceUsd, not even an explicit nil
### GetDexUsdPrice

`func (o *AssetInfoSchema) GetDexUsdPrice() string`

GetDexUsdPrice returns the DexUsdPrice field if non-nil, zero value otherwise.

### GetDexUsdPriceOk

`func (o *AssetInfoSchema) GetDexUsdPriceOk() (*string, bool)`

GetDexUsdPriceOk returns a tuple with the DexUsdPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexUsdPrice

`func (o *AssetInfoSchema) SetDexUsdPrice(v string)`

SetDexUsdPrice sets DexUsdPrice field to given value.

### HasDexUsdPrice

`func (o *AssetInfoSchema) HasDexUsdPrice() bool`

HasDexUsdPrice returns a boolean if a field has been set.

### SetDexUsdPriceNil

`func (o *AssetInfoSchema) SetDexUsdPriceNil(b bool)`

 SetDexUsdPriceNil sets the value for DexUsdPrice to be an explicit nil

### UnsetDexUsdPrice
`func (o *AssetInfoSchema) UnsetDexUsdPrice()`

UnsetDexUsdPrice ensures that no value is present for DexUsdPrice, not even an explicit nil
### GetDisplayName

`func (o *AssetInfoSchema) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AssetInfoSchema) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AssetInfoSchema) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AssetInfoSchema) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AssetInfoSchema) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AssetInfoSchema) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetImageUrl

`func (o *AssetInfoSchema) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *AssetInfoSchema) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *AssetInfoSchema) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *AssetInfoSchema) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *AssetInfoSchema) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *AssetInfoSchema) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetKind

`func (o *AssetInfoSchema) GetKind() AssetKindSchema`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AssetInfoSchema) GetKindOk() (*AssetKindSchema, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AssetInfoSchema) SetKind(v AssetKindSchema)`

SetKind sets Kind field to given value.


### GetSymbol

`func (o *AssetInfoSchema) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AssetInfoSchema) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AssetInfoSchema) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetThirdPartyPriceUsd

`func (o *AssetInfoSchema) GetThirdPartyPriceUsd() string`

GetThirdPartyPriceUsd returns the ThirdPartyPriceUsd field if non-nil, zero value otherwise.

### GetThirdPartyPriceUsdOk

`func (o *AssetInfoSchema) GetThirdPartyPriceUsdOk() (*string, bool)`

GetThirdPartyPriceUsdOk returns a tuple with the ThirdPartyPriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyPriceUsd

`func (o *AssetInfoSchema) SetThirdPartyPriceUsd(v string)`

SetThirdPartyPriceUsd sets ThirdPartyPriceUsd field to given value.

### HasThirdPartyPriceUsd

`func (o *AssetInfoSchema) HasThirdPartyPriceUsd() bool`

HasThirdPartyPriceUsd returns a boolean if a field has been set.

### SetThirdPartyPriceUsdNil

`func (o *AssetInfoSchema) SetThirdPartyPriceUsdNil(b bool)`

 SetThirdPartyPriceUsdNil sets the value for ThirdPartyPriceUsd to be an explicit nil

### UnsetThirdPartyPriceUsd
`func (o *AssetInfoSchema) UnsetThirdPartyPriceUsd()`

UnsetThirdPartyPriceUsd ensures that no value is present for ThirdPartyPriceUsd, not even an explicit nil
### GetThirdPartyUsdPrice

`func (o *AssetInfoSchema) GetThirdPartyUsdPrice() string`

GetThirdPartyUsdPrice returns the ThirdPartyUsdPrice field if non-nil, zero value otherwise.

### GetThirdPartyUsdPriceOk

`func (o *AssetInfoSchema) GetThirdPartyUsdPriceOk() (*string, bool)`

GetThirdPartyUsdPriceOk returns a tuple with the ThirdPartyUsdPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyUsdPrice

`func (o *AssetInfoSchema) SetThirdPartyUsdPrice(v string)`

SetThirdPartyUsdPrice sets ThirdPartyUsdPrice field to given value.

### HasThirdPartyUsdPrice

`func (o *AssetInfoSchema) HasThirdPartyUsdPrice() bool`

HasThirdPartyUsdPrice returns a boolean if a field has been set.

### SetThirdPartyUsdPriceNil

`func (o *AssetInfoSchema) SetThirdPartyUsdPriceNil(b bool)`

 SetThirdPartyUsdPriceNil sets the value for ThirdPartyUsdPrice to be an explicit nil

### UnsetThirdPartyUsdPrice
`func (o *AssetInfoSchema) UnsetThirdPartyUsdPrice()`

UnsetThirdPartyUsdPrice ensures that no value is present for ThirdPartyUsdPrice, not even an explicit nil
### GetWalletAddress

`func (o *AssetInfoSchema) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *AssetInfoSchema) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *AssetInfoSchema) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *AssetInfoSchema) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### SetWalletAddressNil

`func (o *AssetInfoSchema) SetWalletAddressNil(b bool)`

 SetWalletAddressNil sets the value for WalletAddress to be an explicit nil

### UnsetWalletAddress
`func (o *AssetInfoSchema) UnsetWalletAddress()`

UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PoolInfoSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address of pool | 
**Apy1d** | Pointer to **NullableString** | Annual percentage yield on the last day | [optional] 
**Apy30d** | Pointer to **NullableString** | Annual percentage yield on the last month | [optional] 
**Apy7d** | Pointer to **NullableString** | Annual percentage yield on the last week | [optional] 
**CollectedToken0ProtocolFee** | **string** | Protocol fee collected in token0 | 
**CollectedToken1ProtocolFee** | **string** | Protocol fee collected in token1 | 
**Deprecated** | **bool** | Pool is deprecated | 
**LpAccountAddress** | Pointer to **NullableString** | Account address of liquidity pool tokens | [optional] 
**LpBalance** | Pointer to **NullableString** | Balance of liquidity pool tokens | [optional] 
**LpFee** | **string** | Fee of liquidity pool token | 
**LpPriceUsd** | Pointer to **NullableString** | Price of liquidity pool token in USD | [optional] 
**LpTotalSupply** | **string** | Total supply of liquidity pool tokens | 
**LpTotalSupplyUsd** | Pointer to **NullableString** | Total supply of liquidity pool tokens in USD | [optional] 
**LpWalletAddress** | Pointer to **NullableString** | Wallet address of liquidity pool tokens | [optional] 
**ProtocolFee** | **string** | Fee of protocol | 
**ProtocolFeeAddress** | **string** | Address of protocol fee | 
**RefFee** | **string** | Referral fee | 
**Reserve0** | **string** | Reserve of token0 in pool | 
**Reserve1** | **string** | Reserve of token1 in pool | 
**RouterAddress** | **string** | Address of router | 
**Token0Address** | **string** | Address of token0 | 
**Token0Balance** | Pointer to **NullableString** | Balance of token0 | [optional] 
**Token1Address** | **string** | Address of token1 | 
**Token1Balance** | Pointer to **NullableString** | Balance of token1 | [optional] 

## Methods

### NewPoolInfoSchema

`func NewPoolInfoSchema(address string, collectedToken0ProtocolFee string, collectedToken1ProtocolFee string, deprecated bool, lpFee string, lpTotalSupply string, protocolFee string, protocolFeeAddress string, refFee string, reserve0 string, reserve1 string, routerAddress string, token0Address string, token1Address string, ) *PoolInfoSchema`

NewPoolInfoSchema instantiates a new PoolInfoSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolInfoSchemaWithDefaults

`func NewPoolInfoSchemaWithDefaults() *PoolInfoSchema`

NewPoolInfoSchemaWithDefaults instantiates a new PoolInfoSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PoolInfoSchema) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PoolInfoSchema) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PoolInfoSchema) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetApy1d

`func (o *PoolInfoSchema) GetApy1d() string`

GetApy1d returns the Apy1d field if non-nil, zero value otherwise.

### GetApy1dOk

`func (o *PoolInfoSchema) GetApy1dOk() (*string, bool)`

GetApy1dOk returns a tuple with the Apy1d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy1d

`func (o *PoolInfoSchema) SetApy1d(v string)`

SetApy1d sets Apy1d field to given value.

### HasApy1d

`func (o *PoolInfoSchema) HasApy1d() bool`

HasApy1d returns a boolean if a field has been set.

### SetApy1dNil

`func (o *PoolInfoSchema) SetApy1dNil(b bool)`

 SetApy1dNil sets the value for Apy1d to be an explicit nil

### UnsetApy1d
`func (o *PoolInfoSchema) UnsetApy1d()`

UnsetApy1d ensures that no value is present for Apy1d, not even an explicit nil
### GetApy30d

`func (o *PoolInfoSchema) GetApy30d() string`

GetApy30d returns the Apy30d field if non-nil, zero value otherwise.

### GetApy30dOk

`func (o *PoolInfoSchema) GetApy30dOk() (*string, bool)`

GetApy30dOk returns a tuple with the Apy30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy30d

`func (o *PoolInfoSchema) SetApy30d(v string)`

SetApy30d sets Apy30d field to given value.

### HasApy30d

`func (o *PoolInfoSchema) HasApy30d() bool`

HasApy30d returns a boolean if a field has been set.

### SetApy30dNil

`func (o *PoolInfoSchema) SetApy30dNil(b bool)`

 SetApy30dNil sets the value for Apy30d to be an explicit nil

### UnsetApy30d
`func (o *PoolInfoSchema) UnsetApy30d()`

UnsetApy30d ensures that no value is present for Apy30d, not even an explicit nil
### GetApy7d

`func (o *PoolInfoSchema) GetApy7d() string`

GetApy7d returns the Apy7d field if non-nil, zero value otherwise.

### GetApy7dOk

`func (o *PoolInfoSchema) GetApy7dOk() (*string, bool)`

GetApy7dOk returns a tuple with the Apy7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy7d

`func (o *PoolInfoSchema) SetApy7d(v string)`

SetApy7d sets Apy7d field to given value.

### HasApy7d

`func (o *PoolInfoSchema) HasApy7d() bool`

HasApy7d returns a boolean if a field has been set.

### SetApy7dNil

`func (o *PoolInfoSchema) SetApy7dNil(b bool)`

 SetApy7dNil sets the value for Apy7d to be an explicit nil

### UnsetApy7d
`func (o *PoolInfoSchema) UnsetApy7d()`

UnsetApy7d ensures that no value is present for Apy7d, not even an explicit nil
### GetCollectedToken0ProtocolFee

`func (o *PoolInfoSchema) GetCollectedToken0ProtocolFee() string`

GetCollectedToken0ProtocolFee returns the CollectedToken0ProtocolFee field if non-nil, zero value otherwise.

### GetCollectedToken0ProtocolFeeOk

`func (o *PoolInfoSchema) GetCollectedToken0ProtocolFeeOk() (*string, bool)`

GetCollectedToken0ProtocolFeeOk returns a tuple with the CollectedToken0ProtocolFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedToken0ProtocolFee

`func (o *PoolInfoSchema) SetCollectedToken0ProtocolFee(v string)`

SetCollectedToken0ProtocolFee sets CollectedToken0ProtocolFee field to given value.


### GetCollectedToken1ProtocolFee

`func (o *PoolInfoSchema) GetCollectedToken1ProtocolFee() string`

GetCollectedToken1ProtocolFee returns the CollectedToken1ProtocolFee field if non-nil, zero value otherwise.

### GetCollectedToken1ProtocolFeeOk

`func (o *PoolInfoSchema) GetCollectedToken1ProtocolFeeOk() (*string, bool)`

GetCollectedToken1ProtocolFeeOk returns a tuple with the CollectedToken1ProtocolFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedToken1ProtocolFee

`func (o *PoolInfoSchema) SetCollectedToken1ProtocolFee(v string)`

SetCollectedToken1ProtocolFee sets CollectedToken1ProtocolFee field to given value.


### GetDeprecated

`func (o *PoolInfoSchema) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *PoolInfoSchema) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *PoolInfoSchema) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetLpAccountAddress

`func (o *PoolInfoSchema) GetLpAccountAddress() string`

GetLpAccountAddress returns the LpAccountAddress field if non-nil, zero value otherwise.

### GetLpAccountAddressOk

`func (o *PoolInfoSchema) GetLpAccountAddressOk() (*string, bool)`

GetLpAccountAddressOk returns a tuple with the LpAccountAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpAccountAddress

`func (o *PoolInfoSchema) SetLpAccountAddress(v string)`

SetLpAccountAddress sets LpAccountAddress field to given value.

### HasLpAccountAddress

`func (o *PoolInfoSchema) HasLpAccountAddress() bool`

HasLpAccountAddress returns a boolean if a field has been set.

### SetLpAccountAddressNil

`func (o *PoolInfoSchema) SetLpAccountAddressNil(b bool)`

 SetLpAccountAddressNil sets the value for LpAccountAddress to be an explicit nil

### UnsetLpAccountAddress
`func (o *PoolInfoSchema) UnsetLpAccountAddress()`

UnsetLpAccountAddress ensures that no value is present for LpAccountAddress, not even an explicit nil
### GetLpBalance

`func (o *PoolInfoSchema) GetLpBalance() string`

GetLpBalance returns the LpBalance field if non-nil, zero value otherwise.

### GetLpBalanceOk

`func (o *PoolInfoSchema) GetLpBalanceOk() (*string, bool)`

GetLpBalanceOk returns a tuple with the LpBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpBalance

`func (o *PoolInfoSchema) SetLpBalance(v string)`

SetLpBalance sets LpBalance field to given value.

### HasLpBalance

`func (o *PoolInfoSchema) HasLpBalance() bool`

HasLpBalance returns a boolean if a field has been set.

### SetLpBalanceNil

`func (o *PoolInfoSchema) SetLpBalanceNil(b bool)`

 SetLpBalanceNil sets the value for LpBalance to be an explicit nil

### UnsetLpBalance
`func (o *PoolInfoSchema) UnsetLpBalance()`

UnsetLpBalance ensures that no value is present for LpBalance, not even an explicit nil
### GetLpFee

`func (o *PoolInfoSchema) GetLpFee() string`

GetLpFee returns the LpFee field if non-nil, zero value otherwise.

### GetLpFeeOk

`func (o *PoolInfoSchema) GetLpFeeOk() (*string, bool)`

GetLpFeeOk returns a tuple with the LpFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpFee

`func (o *PoolInfoSchema) SetLpFee(v string)`

SetLpFee sets LpFee field to given value.


### GetLpPriceUsd

`func (o *PoolInfoSchema) GetLpPriceUsd() string`

GetLpPriceUsd returns the LpPriceUsd field if non-nil, zero value otherwise.

### GetLpPriceUsdOk

`func (o *PoolInfoSchema) GetLpPriceUsdOk() (*string, bool)`

GetLpPriceUsdOk returns a tuple with the LpPriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpPriceUsd

`func (o *PoolInfoSchema) SetLpPriceUsd(v string)`

SetLpPriceUsd sets LpPriceUsd field to given value.

### HasLpPriceUsd

`func (o *PoolInfoSchema) HasLpPriceUsd() bool`

HasLpPriceUsd returns a boolean if a field has been set.

### SetLpPriceUsdNil

`func (o *PoolInfoSchema) SetLpPriceUsdNil(b bool)`

 SetLpPriceUsdNil sets the value for LpPriceUsd to be an explicit nil

### UnsetLpPriceUsd
`func (o *PoolInfoSchema) UnsetLpPriceUsd()`

UnsetLpPriceUsd ensures that no value is present for LpPriceUsd, not even an explicit nil
### GetLpTotalSupply

`func (o *PoolInfoSchema) GetLpTotalSupply() string`

GetLpTotalSupply returns the LpTotalSupply field if non-nil, zero value otherwise.

### GetLpTotalSupplyOk

`func (o *PoolInfoSchema) GetLpTotalSupplyOk() (*string, bool)`

GetLpTotalSupplyOk returns a tuple with the LpTotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpTotalSupply

`func (o *PoolInfoSchema) SetLpTotalSupply(v string)`

SetLpTotalSupply sets LpTotalSupply field to given value.


### GetLpTotalSupplyUsd

`func (o *PoolInfoSchema) GetLpTotalSupplyUsd() string`

GetLpTotalSupplyUsd returns the LpTotalSupplyUsd field if non-nil, zero value otherwise.

### GetLpTotalSupplyUsdOk

`func (o *PoolInfoSchema) GetLpTotalSupplyUsdOk() (*string, bool)`

GetLpTotalSupplyUsdOk returns a tuple with the LpTotalSupplyUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpTotalSupplyUsd

`func (o *PoolInfoSchema) SetLpTotalSupplyUsd(v string)`

SetLpTotalSupplyUsd sets LpTotalSupplyUsd field to given value.

### HasLpTotalSupplyUsd

`func (o *PoolInfoSchema) HasLpTotalSupplyUsd() bool`

HasLpTotalSupplyUsd returns a boolean if a field has been set.

### SetLpTotalSupplyUsdNil

`func (o *PoolInfoSchema) SetLpTotalSupplyUsdNil(b bool)`

 SetLpTotalSupplyUsdNil sets the value for LpTotalSupplyUsd to be an explicit nil

### UnsetLpTotalSupplyUsd
`func (o *PoolInfoSchema) UnsetLpTotalSupplyUsd()`

UnsetLpTotalSupplyUsd ensures that no value is present for LpTotalSupplyUsd, not even an explicit nil
### GetLpWalletAddress

`func (o *PoolInfoSchema) GetLpWalletAddress() string`

GetLpWalletAddress returns the LpWalletAddress field if non-nil, zero value otherwise.

### GetLpWalletAddressOk

`func (o *PoolInfoSchema) GetLpWalletAddressOk() (*string, bool)`

GetLpWalletAddressOk returns a tuple with the LpWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpWalletAddress

`func (o *PoolInfoSchema) SetLpWalletAddress(v string)`

SetLpWalletAddress sets LpWalletAddress field to given value.

### HasLpWalletAddress

`func (o *PoolInfoSchema) HasLpWalletAddress() bool`

HasLpWalletAddress returns a boolean if a field has been set.

### SetLpWalletAddressNil

`func (o *PoolInfoSchema) SetLpWalletAddressNil(b bool)`

 SetLpWalletAddressNil sets the value for LpWalletAddress to be an explicit nil

### UnsetLpWalletAddress
`func (o *PoolInfoSchema) UnsetLpWalletAddress()`

UnsetLpWalletAddress ensures that no value is present for LpWalletAddress, not even an explicit nil
### GetProtocolFee

`func (o *PoolInfoSchema) GetProtocolFee() string`

GetProtocolFee returns the ProtocolFee field if non-nil, zero value otherwise.

### GetProtocolFeeOk

`func (o *PoolInfoSchema) GetProtocolFeeOk() (*string, bool)`

GetProtocolFeeOk returns a tuple with the ProtocolFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFee

`func (o *PoolInfoSchema) SetProtocolFee(v string)`

SetProtocolFee sets ProtocolFee field to given value.


### GetProtocolFeeAddress

`func (o *PoolInfoSchema) GetProtocolFeeAddress() string`

GetProtocolFeeAddress returns the ProtocolFeeAddress field if non-nil, zero value otherwise.

### GetProtocolFeeAddressOk

`func (o *PoolInfoSchema) GetProtocolFeeAddressOk() (*string, bool)`

GetProtocolFeeAddressOk returns a tuple with the ProtocolFeeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFeeAddress

`func (o *PoolInfoSchema) SetProtocolFeeAddress(v string)`

SetProtocolFeeAddress sets ProtocolFeeAddress field to given value.


### GetRefFee

`func (o *PoolInfoSchema) GetRefFee() string`

GetRefFee returns the RefFee field if non-nil, zero value otherwise.

### GetRefFeeOk

`func (o *PoolInfoSchema) GetRefFeeOk() (*string, bool)`

GetRefFeeOk returns a tuple with the RefFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefFee

`func (o *PoolInfoSchema) SetRefFee(v string)`

SetRefFee sets RefFee field to given value.


### GetReserve0

`func (o *PoolInfoSchema) GetReserve0() string`

GetReserve0 returns the Reserve0 field if non-nil, zero value otherwise.

### GetReserve0Ok

`func (o *PoolInfoSchema) GetReserve0Ok() (*string, bool)`

GetReserve0Ok returns a tuple with the Reserve0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserve0

`func (o *PoolInfoSchema) SetReserve0(v string)`

SetReserve0 sets Reserve0 field to given value.


### GetReserve1

`func (o *PoolInfoSchema) GetReserve1() string`

GetReserve1 returns the Reserve1 field if non-nil, zero value otherwise.

### GetReserve1Ok

`func (o *PoolInfoSchema) GetReserve1Ok() (*string, bool)`

GetReserve1Ok returns a tuple with the Reserve1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserve1

`func (o *PoolInfoSchema) SetReserve1(v string)`

SetReserve1 sets Reserve1 field to given value.


### GetRouterAddress

`func (o *PoolInfoSchema) GetRouterAddress() string`

GetRouterAddress returns the RouterAddress field if non-nil, zero value otherwise.

### GetRouterAddressOk

`func (o *PoolInfoSchema) GetRouterAddressOk() (*string, bool)`

GetRouterAddressOk returns a tuple with the RouterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterAddress

`func (o *PoolInfoSchema) SetRouterAddress(v string)`

SetRouterAddress sets RouterAddress field to given value.


### GetToken0Address

`func (o *PoolInfoSchema) GetToken0Address() string`

GetToken0Address returns the Token0Address field if non-nil, zero value otherwise.

### GetToken0AddressOk

`func (o *PoolInfoSchema) GetToken0AddressOk() (*string, bool)`

GetToken0AddressOk returns a tuple with the Token0Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken0Address

`func (o *PoolInfoSchema) SetToken0Address(v string)`

SetToken0Address sets Token0Address field to given value.


### GetToken0Balance

`func (o *PoolInfoSchema) GetToken0Balance() string`

GetToken0Balance returns the Token0Balance field if non-nil, zero value otherwise.

### GetToken0BalanceOk

`func (o *PoolInfoSchema) GetToken0BalanceOk() (*string, bool)`

GetToken0BalanceOk returns a tuple with the Token0Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken0Balance

`func (o *PoolInfoSchema) SetToken0Balance(v string)`

SetToken0Balance sets Token0Balance field to given value.

### HasToken0Balance

`func (o *PoolInfoSchema) HasToken0Balance() bool`

HasToken0Balance returns a boolean if a field has been set.

### SetToken0BalanceNil

`func (o *PoolInfoSchema) SetToken0BalanceNil(b bool)`

 SetToken0BalanceNil sets the value for Token0Balance to be an explicit nil

### UnsetToken0Balance
`func (o *PoolInfoSchema) UnsetToken0Balance()`

UnsetToken0Balance ensures that no value is present for Token0Balance, not even an explicit nil
### GetToken1Address

`func (o *PoolInfoSchema) GetToken1Address() string`

GetToken1Address returns the Token1Address field if non-nil, zero value otherwise.

### GetToken1AddressOk

`func (o *PoolInfoSchema) GetToken1AddressOk() (*string, bool)`

GetToken1AddressOk returns a tuple with the Token1Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken1Address

`func (o *PoolInfoSchema) SetToken1Address(v string)`

SetToken1Address sets Token1Address field to given value.


### GetToken1Balance

`func (o *PoolInfoSchema) GetToken1Balance() string`

GetToken1Balance returns the Token1Balance field if non-nil, zero value otherwise.

### GetToken1BalanceOk

`func (o *PoolInfoSchema) GetToken1BalanceOk() (*string, bool)`

GetToken1BalanceOk returns a tuple with the Token1Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken1Balance

`func (o *PoolInfoSchema) SetToken1Balance(v string)`

SetToken1Balance sets Token1Balance field to given value.

### HasToken1Balance

`func (o *PoolInfoSchema) HasToken1Balance() bool`

HasToken1Balance returns a boolean if a field has been set.

### SetToken1BalanceNil

`func (o *PoolInfoSchema) SetToken1BalanceNil(b bool)`

 SetToken1BalanceNil sets the value for Token1Balance to be an explicit nil

### UnsetToken1Balance
`func (o *PoolInfoSchema) UnsetToken1Balance()`

UnsetToken1Balance ensures that no value is present for Token1Balance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



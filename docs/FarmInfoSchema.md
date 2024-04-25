# FarmInfoSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | Pointer to **NullableString** | Annual percentage yield | [optional] 
**MinStakeDurationS** | **string** | Minimal stake duration seconds | 
**MinterAddress** | **string** | Address of the farm | 
**NftInfos** | [**[]FarmNftInfoSchema**](FarmNftInfoSchema.md) | Farm NFT list | 
**PoolAddress** | **string** | Address of the pool | 
**RewardTokenAddress** | **string** | Address of the reward token | 
**Status** | **string** | Minter status | 

## Methods

### NewFarmInfoSchema

`func NewFarmInfoSchema(minStakeDurationS string, minterAddress string, nftInfos []FarmNftInfoSchema, poolAddress string, rewardTokenAddress string, status string, ) *FarmInfoSchema`

NewFarmInfoSchema instantiates a new FarmInfoSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmInfoSchemaWithDefaults

`func NewFarmInfoSchemaWithDefaults() *FarmInfoSchema`

NewFarmInfoSchemaWithDefaults instantiates a new FarmInfoSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApy

`func (o *FarmInfoSchema) GetApy() string`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *FarmInfoSchema) GetApyOk() (*string, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *FarmInfoSchema) SetApy(v string)`

SetApy sets Apy field to given value.

### HasApy

`func (o *FarmInfoSchema) HasApy() bool`

HasApy returns a boolean if a field has been set.

### SetApyNil

`func (o *FarmInfoSchema) SetApyNil(b bool)`

 SetApyNil sets the value for Apy to be an explicit nil

### UnsetApy
`func (o *FarmInfoSchema) UnsetApy()`

UnsetApy ensures that no value is present for Apy, not even an explicit nil
### GetMinStakeDurationS

`func (o *FarmInfoSchema) GetMinStakeDurationS() string`

GetMinStakeDurationS returns the MinStakeDurationS field if non-nil, zero value otherwise.

### GetMinStakeDurationSOk

`func (o *FarmInfoSchema) GetMinStakeDurationSOk() (*string, bool)`

GetMinStakeDurationSOk returns a tuple with the MinStakeDurationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStakeDurationS

`func (o *FarmInfoSchema) SetMinStakeDurationS(v string)`

SetMinStakeDurationS sets MinStakeDurationS field to given value.


### GetMinterAddress

`func (o *FarmInfoSchema) GetMinterAddress() string`

GetMinterAddress returns the MinterAddress field if non-nil, zero value otherwise.

### GetMinterAddressOk

`func (o *FarmInfoSchema) GetMinterAddressOk() (*string, bool)`

GetMinterAddressOk returns a tuple with the MinterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinterAddress

`func (o *FarmInfoSchema) SetMinterAddress(v string)`

SetMinterAddress sets MinterAddress field to given value.


### GetNftInfos

`func (o *FarmInfoSchema) GetNftInfos() []FarmNftInfoSchema`

GetNftInfos returns the NftInfos field if non-nil, zero value otherwise.

### GetNftInfosOk

`func (o *FarmInfoSchema) GetNftInfosOk() (*[]FarmNftInfoSchema, bool)`

GetNftInfosOk returns a tuple with the NftInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftInfos

`func (o *FarmInfoSchema) SetNftInfos(v []FarmNftInfoSchema)`

SetNftInfos sets NftInfos field to given value.


### GetPoolAddress

`func (o *FarmInfoSchema) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *FarmInfoSchema) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *FarmInfoSchema) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetRewardTokenAddress

`func (o *FarmInfoSchema) GetRewardTokenAddress() string`

GetRewardTokenAddress returns the RewardTokenAddress field if non-nil, zero value otherwise.

### GetRewardTokenAddressOk

`func (o *FarmInfoSchema) GetRewardTokenAddressOk() (*string, bool)`

GetRewardTokenAddressOk returns a tuple with the RewardTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardTokenAddress

`func (o *FarmInfoSchema) SetRewardTokenAddress(v string)`

SetRewardTokenAddress sets RewardTokenAddress field to given value.


### GetStatus

`func (o *FarmInfoSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FarmInfoSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FarmInfoSchema) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



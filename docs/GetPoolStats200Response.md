# GetPoolStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Since** | **string** | Time since the stats are requested (YYYY-MM-DDTHH:MM:SS) | 
**Stats** | [**[]PoolStats**](PoolStats.md) | Stats of each pool | 
**UniqueWalletsCount** | **int64** | Number of unique wallets operated in pool | 
**Until** | **string** | Time until the stats are requested (YYYY-MM-DDTHH:MM:SS) | 

## Methods

### NewGetPoolStats200Response

`func NewGetPoolStats200Response(since string, stats []PoolStats, uniqueWalletsCount int64, until string, ) *GetPoolStats200Response`

NewGetPoolStats200Response instantiates a new GetPoolStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPoolStats200ResponseWithDefaults

`func NewGetPoolStats200ResponseWithDefaults() *GetPoolStats200Response`

NewGetPoolStats200ResponseWithDefaults instantiates a new GetPoolStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSince

`func (o *GetPoolStats200Response) GetSince() string`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *GetPoolStats200Response) GetSinceOk() (*string, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *GetPoolStats200Response) SetSince(v string)`

SetSince sets Since field to given value.


### GetStats

`func (o *GetPoolStats200Response) GetStats() []PoolStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetPoolStats200Response) GetStatsOk() (*[]PoolStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetPoolStats200Response) SetStats(v []PoolStats)`

SetStats sets Stats field to given value.


### GetUniqueWalletsCount

`func (o *GetPoolStats200Response) GetUniqueWalletsCount() int64`

GetUniqueWalletsCount returns the UniqueWalletsCount field if non-nil, zero value otherwise.

### GetUniqueWalletsCountOk

`func (o *GetPoolStats200Response) GetUniqueWalletsCountOk() (*int64, bool)`

GetUniqueWalletsCountOk returns a tuple with the UniqueWalletsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueWalletsCount

`func (o *GetPoolStats200Response) SetUniqueWalletsCount(v int64)`

SetUniqueWalletsCount sets UniqueWalletsCount field to given value.


### GetUntil

`func (o *GetPoolStats200Response) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *GetPoolStats200Response) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *GetPoolStats200Response) SetUntil(v string)`

SetUntil sets Until field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



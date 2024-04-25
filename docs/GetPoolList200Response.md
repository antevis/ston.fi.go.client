# GetPoolList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolList** | [**[]PoolInfoSchema**](PoolInfoSchema.md) |  | 

## Methods

### NewGetPoolList200Response

`func NewGetPoolList200Response(poolList []PoolInfoSchema, ) *GetPoolList200Response`

NewGetPoolList200Response instantiates a new GetPoolList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPoolList200ResponseWithDefaults

`func NewGetPoolList200ResponseWithDefaults() *GetPoolList200Response`

NewGetPoolList200ResponseWithDefaults instantiates a new GetPoolList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolList

`func (o *GetPoolList200Response) GetPoolList() []PoolInfoSchema`

GetPoolList returns the PoolList field if non-nil, zero value otherwise.

### GetPoolListOk

`func (o *GetPoolList200Response) GetPoolListOk() (*[]PoolInfoSchema, bool)`

GetPoolListOk returns a tuple with the PoolList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolList

`func (o *GetPoolList200Response) SetPoolList(v []PoolInfoSchema)`

SetPoolList sets PoolList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



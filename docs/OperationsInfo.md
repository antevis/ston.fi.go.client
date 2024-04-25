# OperationsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset0Info** | [**AssetInfoSchema**](AssetInfoSchema.md) |  | 
**Asset1Info** | [**AssetInfoSchema**](AssetInfoSchema.md) |  | 
**Operation** | [**OperationStat**](OperationStat.md) |  | 

## Methods

### NewOperationsInfo

`func NewOperationsInfo(asset0Info AssetInfoSchema, asset1Info AssetInfoSchema, operation OperationStat, ) *OperationsInfo`

NewOperationsInfo instantiates a new OperationsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationsInfoWithDefaults

`func NewOperationsInfoWithDefaults() *OperationsInfo`

NewOperationsInfoWithDefaults instantiates a new OperationsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset0Info

`func (o *OperationsInfo) GetAsset0Info() AssetInfoSchema`

GetAsset0Info returns the Asset0Info field if non-nil, zero value otherwise.

### GetAsset0InfoOk

`func (o *OperationsInfo) GetAsset0InfoOk() (*AssetInfoSchema, bool)`

GetAsset0InfoOk returns a tuple with the Asset0Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset0Info

`func (o *OperationsInfo) SetAsset0Info(v AssetInfoSchema)`

SetAsset0Info sets Asset0Info field to given value.


### GetAsset1Info

`func (o *OperationsInfo) GetAsset1Info() AssetInfoSchema`

GetAsset1Info returns the Asset1Info field if non-nil, zero value otherwise.

### GetAsset1InfoOk

`func (o *OperationsInfo) GetAsset1InfoOk() (*AssetInfoSchema, bool)`

GetAsset1InfoOk returns a tuple with the Asset1Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset1Info

`func (o *OperationsInfo) SetAsset1Info(v AssetInfoSchema)`

SetAsset1Info sets Asset1Info field to given value.


### GetOperation

`func (o *OperationsInfo) GetOperation() OperationStat`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OperationsInfo) GetOperationOk() (*OperationStat, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OperationsInfo) SetOperation(v OperationStat)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetOperationStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]OperationsInfo**](OperationsInfo.md) |  | 

## Methods

### NewGetOperationStats200Response

`func NewGetOperationStats200Response(operations []OperationsInfo, ) *GetOperationStats200Response`

NewGetOperationStats200Response instantiates a new GetOperationStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOperationStats200ResponseWithDefaults

`func NewGetOperationStats200ResponseWithDefaults() *GetOperationStats200Response`

NewGetOperationStats200ResponseWithDefaults instantiates a new GetOperationStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *GetOperationStats200Response) GetOperations() []OperationsInfo`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *GetOperationStats200Response) GetOperationsOk() (*[]OperationsInfo, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *GetOperationStats200Response) SetOperations(v []OperationsInfo)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Wallet address | 
**BalanceDeltas** | **string** | Changes in balance | 
**Coins** | **string** | Transaction coins | 
**ExitCode** | **string** | Exit code of operation | 
**LogicalTime** | **string** | Operation logical time | 
**QueryId** | **string** | Id of operation status query | 
**TxHash** | ***os.File** | Operation hash | 

## Methods

### NewOperationStatus

`func NewOperationStatus(address string, balanceDeltas string, coins string, exitCode string, logicalTime string, queryId string, txHash *os.File, ) *OperationStatus`

NewOperationStatus instantiates a new OperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationStatusWithDefaults

`func NewOperationStatusWithDefaults() *OperationStatus`

NewOperationStatusWithDefaults instantiates a new OperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OperationStatus) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OperationStatus) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OperationStatus) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalanceDeltas

`func (o *OperationStatus) GetBalanceDeltas() string`

GetBalanceDeltas returns the BalanceDeltas field if non-nil, zero value otherwise.

### GetBalanceDeltasOk

`func (o *OperationStatus) GetBalanceDeltasOk() (*string, bool)`

GetBalanceDeltasOk returns a tuple with the BalanceDeltas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDeltas

`func (o *OperationStatus) SetBalanceDeltas(v string)`

SetBalanceDeltas sets BalanceDeltas field to given value.


### GetCoins

`func (o *OperationStatus) GetCoins() string`

GetCoins returns the Coins field if non-nil, zero value otherwise.

### GetCoinsOk

`func (o *OperationStatus) GetCoinsOk() (*string, bool)`

GetCoinsOk returns a tuple with the Coins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoins

`func (o *OperationStatus) SetCoins(v string)`

SetCoins sets Coins field to given value.


### GetExitCode

`func (o *OperationStatus) GetExitCode() string`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *OperationStatus) GetExitCodeOk() (*string, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *OperationStatus) SetExitCode(v string)`

SetExitCode sets ExitCode field to given value.


### GetLogicalTime

`func (o *OperationStatus) GetLogicalTime() string`

GetLogicalTime returns the LogicalTime field if non-nil, zero value otherwise.

### GetLogicalTimeOk

`func (o *OperationStatus) GetLogicalTimeOk() (*string, bool)`

GetLogicalTimeOk returns a tuple with the LogicalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalTime

`func (o *OperationStatus) SetLogicalTime(v string)`

SetLogicalTime sets LogicalTime field to given value.


### GetQueryId

`func (o *OperationStatus) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *OperationStatus) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *OperationStatus) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetTxHash

`func (o *OperationStatus) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *OperationStatus) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *OperationStatus) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DexSwapStatus200ResponseOneOf

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
**Type** | **string** |  | 

## Methods

### NewDexSwapStatus200ResponseOneOf

`func NewDexSwapStatus200ResponseOneOf(address string, balanceDeltas string, coins string, exitCode string, logicalTime string, queryId string, txHash *os.File, type_ string, ) *DexSwapStatus200ResponseOneOf`

NewDexSwapStatus200ResponseOneOf instantiates a new DexSwapStatus200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexSwapStatus200ResponseOneOfWithDefaults

`func NewDexSwapStatus200ResponseOneOfWithDefaults() *DexSwapStatus200ResponseOneOf`

NewDexSwapStatus200ResponseOneOfWithDefaults instantiates a new DexSwapStatus200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DexSwapStatus200ResponseOneOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DexSwapStatus200ResponseOneOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DexSwapStatus200ResponseOneOf) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalanceDeltas

`func (o *DexSwapStatus200ResponseOneOf) GetBalanceDeltas() string`

GetBalanceDeltas returns the BalanceDeltas field if non-nil, zero value otherwise.

### GetBalanceDeltasOk

`func (o *DexSwapStatus200ResponseOneOf) GetBalanceDeltasOk() (*string, bool)`

GetBalanceDeltasOk returns a tuple with the BalanceDeltas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDeltas

`func (o *DexSwapStatus200ResponseOneOf) SetBalanceDeltas(v string)`

SetBalanceDeltas sets BalanceDeltas field to given value.


### GetCoins

`func (o *DexSwapStatus200ResponseOneOf) GetCoins() string`

GetCoins returns the Coins field if non-nil, zero value otherwise.

### GetCoinsOk

`func (o *DexSwapStatus200ResponseOneOf) GetCoinsOk() (*string, bool)`

GetCoinsOk returns a tuple with the Coins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoins

`func (o *DexSwapStatus200ResponseOneOf) SetCoins(v string)`

SetCoins sets Coins field to given value.


### GetExitCode

`func (o *DexSwapStatus200ResponseOneOf) GetExitCode() string`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *DexSwapStatus200ResponseOneOf) GetExitCodeOk() (*string, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *DexSwapStatus200ResponseOneOf) SetExitCode(v string)`

SetExitCode sets ExitCode field to given value.


### GetLogicalTime

`func (o *DexSwapStatus200ResponseOneOf) GetLogicalTime() string`

GetLogicalTime returns the LogicalTime field if non-nil, zero value otherwise.

### GetLogicalTimeOk

`func (o *DexSwapStatus200ResponseOneOf) GetLogicalTimeOk() (*string, bool)`

GetLogicalTimeOk returns a tuple with the LogicalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalTime

`func (o *DexSwapStatus200ResponseOneOf) SetLogicalTime(v string)`

SetLogicalTime sets LogicalTime field to given value.


### GetQueryId

`func (o *DexSwapStatus200ResponseOneOf) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *DexSwapStatus200ResponseOneOf) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *DexSwapStatus200ResponseOneOf) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetTxHash

`func (o *DexSwapStatus200ResponseOneOf) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *DexSwapStatus200ResponseOneOf) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *DexSwapStatus200ResponseOneOf) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.


### GetType

`func (o *DexSwapStatus200ResponseOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DexSwapStatus200ResponseOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DexSwapStatus200ResponseOneOf) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



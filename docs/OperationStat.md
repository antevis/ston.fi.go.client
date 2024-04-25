# OperationStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset0Address** | **string** | Address of asset0 | 
**Asset0Amount** | **string** | Amount of asset0 | 
**Asset0Delta** | **string** | Delta of asset0 | 
**Asset0Reserve** | **string** | Reserve of asset0 | 
**Asset1Address** | **string** | Address of asset1 | 
**Asset1Amount** | **string** | Amount of asset1 | 
**Asset1Delta** | **string** | Delta of asset1 | 
**Asset1Reserve** | **string** | Reserve of asset1 | 
**DestinationWalletAddress** | **string** | Address of destination wallet | 
**ExitCode** | **string** | Operation exit code | 
**FeeAssetAddress** | Pointer to **NullableString** | Liquidity pool fee asset address | [optional] 
**LpFeeAmount** | **string** | Liquidity pool fee amount | 
**LpTokenDelta** | **string** | Liquidity pool token amount change | 
**LpTokenSupply** | **string** | Liquidity pool token amount supply | 
**OperationType** | **string** | Type of operation | 
**PoolAddress** | **string** | Address of the pool | 
**PoolTxHash** | **string** | Transaction hash | 
**PoolTxLt** | **int64** | Transaction logical time | 
**PoolTxTimestamp** | **string** | Transaction timestamp | 
**ProtocolFeeAmount** | **string** | Protocol fee amount | 
**ReferralAddress** | Pointer to **NullableString** | Referral fee address | [optional] 
**ReferralFeeAmount** | **string** | Referral fee amount | 
**RouterAddress** | **string** | Address of the router | 
**Success** | **bool** | Operation is successful | 
**WalletAddress** | **string** | Wallet address | 
**WalletTxHash** | **string** | Transaction hash | 
**WalletTxLt** | **string** | Transaction logical time | 
**WalletTxTimestamp** | **string** | Transaction timestamp | 

## Methods

### NewOperationStat

`func NewOperationStat(asset0Address string, asset0Amount string, asset0Delta string, asset0Reserve string, asset1Address string, asset1Amount string, asset1Delta string, asset1Reserve string, destinationWalletAddress string, exitCode string, lpFeeAmount string, lpTokenDelta string, lpTokenSupply string, operationType string, poolAddress string, poolTxHash string, poolTxLt int64, poolTxTimestamp string, protocolFeeAmount string, referralFeeAmount string, routerAddress string, success bool, walletAddress string, walletTxHash string, walletTxLt string, walletTxTimestamp string, ) *OperationStat`

NewOperationStat instantiates a new OperationStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationStatWithDefaults

`func NewOperationStatWithDefaults() *OperationStat`

NewOperationStatWithDefaults instantiates a new OperationStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset0Address

`func (o *OperationStat) GetAsset0Address() string`

GetAsset0Address returns the Asset0Address field if non-nil, zero value otherwise.

### GetAsset0AddressOk

`func (o *OperationStat) GetAsset0AddressOk() (*string, bool)`

GetAsset0AddressOk returns a tuple with the Asset0Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset0Address

`func (o *OperationStat) SetAsset0Address(v string)`

SetAsset0Address sets Asset0Address field to given value.


### GetAsset0Amount

`func (o *OperationStat) GetAsset0Amount() string`

GetAsset0Amount returns the Asset0Amount field if non-nil, zero value otherwise.

### GetAsset0AmountOk

`func (o *OperationStat) GetAsset0AmountOk() (*string, bool)`

GetAsset0AmountOk returns a tuple with the Asset0Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset0Amount

`func (o *OperationStat) SetAsset0Amount(v string)`

SetAsset0Amount sets Asset0Amount field to given value.


### GetAsset0Delta

`func (o *OperationStat) GetAsset0Delta() string`

GetAsset0Delta returns the Asset0Delta field if non-nil, zero value otherwise.

### GetAsset0DeltaOk

`func (o *OperationStat) GetAsset0DeltaOk() (*string, bool)`

GetAsset0DeltaOk returns a tuple with the Asset0Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset0Delta

`func (o *OperationStat) SetAsset0Delta(v string)`

SetAsset0Delta sets Asset0Delta field to given value.


### GetAsset0Reserve

`func (o *OperationStat) GetAsset0Reserve() string`

GetAsset0Reserve returns the Asset0Reserve field if non-nil, zero value otherwise.

### GetAsset0ReserveOk

`func (o *OperationStat) GetAsset0ReserveOk() (*string, bool)`

GetAsset0ReserveOk returns a tuple with the Asset0Reserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset0Reserve

`func (o *OperationStat) SetAsset0Reserve(v string)`

SetAsset0Reserve sets Asset0Reserve field to given value.


### GetAsset1Address

`func (o *OperationStat) GetAsset1Address() string`

GetAsset1Address returns the Asset1Address field if non-nil, zero value otherwise.

### GetAsset1AddressOk

`func (o *OperationStat) GetAsset1AddressOk() (*string, bool)`

GetAsset1AddressOk returns a tuple with the Asset1Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset1Address

`func (o *OperationStat) SetAsset1Address(v string)`

SetAsset1Address sets Asset1Address field to given value.


### GetAsset1Amount

`func (o *OperationStat) GetAsset1Amount() string`

GetAsset1Amount returns the Asset1Amount field if non-nil, zero value otherwise.

### GetAsset1AmountOk

`func (o *OperationStat) GetAsset1AmountOk() (*string, bool)`

GetAsset1AmountOk returns a tuple with the Asset1Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset1Amount

`func (o *OperationStat) SetAsset1Amount(v string)`

SetAsset1Amount sets Asset1Amount field to given value.


### GetAsset1Delta

`func (o *OperationStat) GetAsset1Delta() string`

GetAsset1Delta returns the Asset1Delta field if non-nil, zero value otherwise.

### GetAsset1DeltaOk

`func (o *OperationStat) GetAsset1DeltaOk() (*string, bool)`

GetAsset1DeltaOk returns a tuple with the Asset1Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset1Delta

`func (o *OperationStat) SetAsset1Delta(v string)`

SetAsset1Delta sets Asset1Delta field to given value.


### GetAsset1Reserve

`func (o *OperationStat) GetAsset1Reserve() string`

GetAsset1Reserve returns the Asset1Reserve field if non-nil, zero value otherwise.

### GetAsset1ReserveOk

`func (o *OperationStat) GetAsset1ReserveOk() (*string, bool)`

GetAsset1ReserveOk returns a tuple with the Asset1Reserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset1Reserve

`func (o *OperationStat) SetAsset1Reserve(v string)`

SetAsset1Reserve sets Asset1Reserve field to given value.


### GetDestinationWalletAddress

`func (o *OperationStat) GetDestinationWalletAddress() string`

GetDestinationWalletAddress returns the DestinationWalletAddress field if non-nil, zero value otherwise.

### GetDestinationWalletAddressOk

`func (o *OperationStat) GetDestinationWalletAddressOk() (*string, bool)`

GetDestinationWalletAddressOk returns a tuple with the DestinationWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationWalletAddress

`func (o *OperationStat) SetDestinationWalletAddress(v string)`

SetDestinationWalletAddress sets DestinationWalletAddress field to given value.


### GetExitCode

`func (o *OperationStat) GetExitCode() string`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *OperationStat) GetExitCodeOk() (*string, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *OperationStat) SetExitCode(v string)`

SetExitCode sets ExitCode field to given value.


### GetFeeAssetAddress

`func (o *OperationStat) GetFeeAssetAddress() string`

GetFeeAssetAddress returns the FeeAssetAddress field if non-nil, zero value otherwise.

### GetFeeAssetAddressOk

`func (o *OperationStat) GetFeeAssetAddressOk() (*string, bool)`

GetFeeAssetAddressOk returns a tuple with the FeeAssetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAssetAddress

`func (o *OperationStat) SetFeeAssetAddress(v string)`

SetFeeAssetAddress sets FeeAssetAddress field to given value.

### HasFeeAssetAddress

`func (o *OperationStat) HasFeeAssetAddress() bool`

HasFeeAssetAddress returns a boolean if a field has been set.

### SetFeeAssetAddressNil

`func (o *OperationStat) SetFeeAssetAddressNil(b bool)`

 SetFeeAssetAddressNil sets the value for FeeAssetAddress to be an explicit nil

### UnsetFeeAssetAddress
`func (o *OperationStat) UnsetFeeAssetAddress()`

UnsetFeeAssetAddress ensures that no value is present for FeeAssetAddress, not even an explicit nil
### GetLpFeeAmount

`func (o *OperationStat) GetLpFeeAmount() string`

GetLpFeeAmount returns the LpFeeAmount field if non-nil, zero value otherwise.

### GetLpFeeAmountOk

`func (o *OperationStat) GetLpFeeAmountOk() (*string, bool)`

GetLpFeeAmountOk returns a tuple with the LpFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpFeeAmount

`func (o *OperationStat) SetLpFeeAmount(v string)`

SetLpFeeAmount sets LpFeeAmount field to given value.


### GetLpTokenDelta

`func (o *OperationStat) GetLpTokenDelta() string`

GetLpTokenDelta returns the LpTokenDelta field if non-nil, zero value otherwise.

### GetLpTokenDeltaOk

`func (o *OperationStat) GetLpTokenDeltaOk() (*string, bool)`

GetLpTokenDeltaOk returns a tuple with the LpTokenDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpTokenDelta

`func (o *OperationStat) SetLpTokenDelta(v string)`

SetLpTokenDelta sets LpTokenDelta field to given value.


### GetLpTokenSupply

`func (o *OperationStat) GetLpTokenSupply() string`

GetLpTokenSupply returns the LpTokenSupply field if non-nil, zero value otherwise.

### GetLpTokenSupplyOk

`func (o *OperationStat) GetLpTokenSupplyOk() (*string, bool)`

GetLpTokenSupplyOk returns a tuple with the LpTokenSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpTokenSupply

`func (o *OperationStat) SetLpTokenSupply(v string)`

SetLpTokenSupply sets LpTokenSupply field to given value.


### GetOperationType

`func (o *OperationStat) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *OperationStat) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *OperationStat) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetPoolAddress

`func (o *OperationStat) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *OperationStat) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *OperationStat) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetPoolTxHash

`func (o *OperationStat) GetPoolTxHash() string`

GetPoolTxHash returns the PoolTxHash field if non-nil, zero value otherwise.

### GetPoolTxHashOk

`func (o *OperationStat) GetPoolTxHashOk() (*string, bool)`

GetPoolTxHashOk returns a tuple with the PoolTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTxHash

`func (o *OperationStat) SetPoolTxHash(v string)`

SetPoolTxHash sets PoolTxHash field to given value.


### GetPoolTxLt

`func (o *OperationStat) GetPoolTxLt() int64`

GetPoolTxLt returns the PoolTxLt field if non-nil, zero value otherwise.

### GetPoolTxLtOk

`func (o *OperationStat) GetPoolTxLtOk() (*int64, bool)`

GetPoolTxLtOk returns a tuple with the PoolTxLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTxLt

`func (o *OperationStat) SetPoolTxLt(v int64)`

SetPoolTxLt sets PoolTxLt field to given value.


### GetPoolTxTimestamp

`func (o *OperationStat) GetPoolTxTimestamp() string`

GetPoolTxTimestamp returns the PoolTxTimestamp field if non-nil, zero value otherwise.

### GetPoolTxTimestampOk

`func (o *OperationStat) GetPoolTxTimestampOk() (*string, bool)`

GetPoolTxTimestampOk returns a tuple with the PoolTxTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTxTimestamp

`func (o *OperationStat) SetPoolTxTimestamp(v string)`

SetPoolTxTimestamp sets PoolTxTimestamp field to given value.


### GetProtocolFeeAmount

`func (o *OperationStat) GetProtocolFeeAmount() string`

GetProtocolFeeAmount returns the ProtocolFeeAmount field if non-nil, zero value otherwise.

### GetProtocolFeeAmountOk

`func (o *OperationStat) GetProtocolFeeAmountOk() (*string, bool)`

GetProtocolFeeAmountOk returns a tuple with the ProtocolFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFeeAmount

`func (o *OperationStat) SetProtocolFeeAmount(v string)`

SetProtocolFeeAmount sets ProtocolFeeAmount field to given value.


### GetReferralAddress

`func (o *OperationStat) GetReferralAddress() string`

GetReferralAddress returns the ReferralAddress field if non-nil, zero value otherwise.

### GetReferralAddressOk

`func (o *OperationStat) GetReferralAddressOk() (*string, bool)`

GetReferralAddressOk returns a tuple with the ReferralAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralAddress

`func (o *OperationStat) SetReferralAddress(v string)`

SetReferralAddress sets ReferralAddress field to given value.

### HasReferralAddress

`func (o *OperationStat) HasReferralAddress() bool`

HasReferralAddress returns a boolean if a field has been set.

### SetReferralAddressNil

`func (o *OperationStat) SetReferralAddressNil(b bool)`

 SetReferralAddressNil sets the value for ReferralAddress to be an explicit nil

### UnsetReferralAddress
`func (o *OperationStat) UnsetReferralAddress()`

UnsetReferralAddress ensures that no value is present for ReferralAddress, not even an explicit nil
### GetReferralFeeAmount

`func (o *OperationStat) GetReferralFeeAmount() string`

GetReferralFeeAmount returns the ReferralFeeAmount field if non-nil, zero value otherwise.

### GetReferralFeeAmountOk

`func (o *OperationStat) GetReferralFeeAmountOk() (*string, bool)`

GetReferralFeeAmountOk returns a tuple with the ReferralFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralFeeAmount

`func (o *OperationStat) SetReferralFeeAmount(v string)`

SetReferralFeeAmount sets ReferralFeeAmount field to given value.


### GetRouterAddress

`func (o *OperationStat) GetRouterAddress() string`

GetRouterAddress returns the RouterAddress field if non-nil, zero value otherwise.

### GetRouterAddressOk

`func (o *OperationStat) GetRouterAddressOk() (*string, bool)`

GetRouterAddressOk returns a tuple with the RouterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterAddress

`func (o *OperationStat) SetRouterAddress(v string)`

SetRouterAddress sets RouterAddress field to given value.


### GetSuccess

`func (o *OperationStat) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OperationStat) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OperationStat) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetWalletAddress

`func (o *OperationStat) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *OperationStat) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *OperationStat) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetWalletTxHash

`func (o *OperationStat) GetWalletTxHash() string`

GetWalletTxHash returns the WalletTxHash field if non-nil, zero value otherwise.

### GetWalletTxHashOk

`func (o *OperationStat) GetWalletTxHashOk() (*string, bool)`

GetWalletTxHashOk returns a tuple with the WalletTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTxHash

`func (o *OperationStat) SetWalletTxHash(v string)`

SetWalletTxHash sets WalletTxHash field to given value.


### GetWalletTxLt

`func (o *OperationStat) GetWalletTxLt() string`

GetWalletTxLt returns the WalletTxLt field if non-nil, zero value otherwise.

### GetWalletTxLtOk

`func (o *OperationStat) GetWalletTxLtOk() (*string, bool)`

GetWalletTxLtOk returns a tuple with the WalletTxLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTxLt

`func (o *OperationStat) SetWalletTxLt(v string)`

SetWalletTxLt sets WalletTxLt field to given value.


### GetWalletTxTimestamp

`func (o *OperationStat) GetWalletTxTimestamp() string`

GetWalletTxTimestamp returns the WalletTxTimestamp field if non-nil, zero value otherwise.

### GetWalletTxTimestampOk

`func (o *OperationStat) GetWalletTxTimestampOk() (*string, bool)`

GetWalletTxTimestampOk returns a tuple with the WalletTxTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTxTimestamp

`func (o *OperationStat) SetWalletTxTimestamp(v string)`

SetWalletTxTimestamp sets WalletTxTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



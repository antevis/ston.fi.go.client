/*
ston-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"os"
	"bytes"
	"fmt"
)

// checks if the OperationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationStatus{}

// OperationStatus struct for OperationStatus
type OperationStatus struct {
	// Wallet address
	Address string `json:"address"`
	// Changes in balance
	BalanceDeltas string `json:"balance_deltas"`
	// Transaction coins
	Coins string `json:"coins"`
	// Exit code of operation
	ExitCode string `json:"exit_code"`
	// Operation logical time
	LogicalTime string `json:"logical_time"`
	// Id of operation status query
	QueryId string `json:"query_id"`
	// Operation hash
	TxHash *os.File `json:"tx_hash"`
}

type _OperationStatus OperationStatus

// NewOperationStatus instantiates a new OperationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationStatus(address string, balanceDeltas string, coins string, exitCode string, logicalTime string, queryId string, txHash *os.File) *OperationStatus {
	this := OperationStatus{}
	this.Address = address
	this.BalanceDeltas = balanceDeltas
	this.Coins = coins
	this.ExitCode = exitCode
	this.LogicalTime = logicalTime
	this.QueryId = queryId
	this.TxHash = txHash
	return &this
}

// NewOperationStatusWithDefaults instantiates a new OperationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationStatusWithDefaults() *OperationStatus {
	this := OperationStatus{}
	return &this
}

// GetAddress returns the Address field value
func (o *OperationStatus) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *OperationStatus) SetAddress(v string) {
	o.Address = v
}

// GetBalanceDeltas returns the BalanceDeltas field value
func (o *OperationStatus) GetBalanceDeltas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceDeltas
}

// GetBalanceDeltasOk returns a tuple with the BalanceDeltas field value
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetBalanceDeltasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceDeltas, true
}

// SetBalanceDeltas sets field value
func (o *OperationStatus) SetBalanceDeltas(v string) {
	o.BalanceDeltas = v
}

// GetCoins returns the Coins field value
func (o *OperationStatus) GetCoins() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coins
}

// GetCoinsOk returns a tuple with the Coins field value
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetCoinsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coins, true
}

// SetCoins sets field value
func (o *OperationStatus) SetCoins(v string) {
	o.Coins = v
}

// GetExitCode returns the ExitCode field value
func (o *OperationStatus) GetExitCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetExitCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExitCode, true
}

// SetExitCode sets field value
func (o *OperationStatus) SetExitCode(v string) {
	o.ExitCode = v
}

// GetLogicalTime returns the LogicalTime field value
func (o *OperationStatus) GetLogicalTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogicalTime
}

// GetLogicalTimeOk returns a tuple with the LogicalTime field value
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetLogicalTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogicalTime, true
}

// SetLogicalTime sets field value
func (o *OperationStatus) SetLogicalTime(v string) {
	o.LogicalTime = v
}

// GetQueryId returns the QueryId field value
func (o *OperationStatus) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *OperationStatus) SetQueryId(v string) {
	o.QueryId = v
}

// GetTxHash returns the TxHash field value
func (o *OperationStatus) GetTxHash() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetTxHashOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *OperationStatus) SetTxHash(v *os.File) {
	o.TxHash = v
}

func (o OperationStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["balance_deltas"] = o.BalanceDeltas
	toSerialize["coins"] = o.Coins
	toSerialize["exit_code"] = o.ExitCode
	toSerialize["logical_time"] = o.LogicalTime
	toSerialize["query_id"] = o.QueryId
	toSerialize["tx_hash"] = o.TxHash
	return toSerialize, nil
}

func (o *OperationStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"balance_deltas",
		"coins",
		"exit_code",
		"logical_time",
		"query_id",
		"tx_hash",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOperationStatus := _OperationStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOperationStatus)

	if err != nil {
		return err
	}

	*o = OperationStatus(varOperationStatus)

	return err
}

type NullableOperationStatus struct {
	value *OperationStatus
	isSet bool
}

func (v NullableOperationStatus) Get() *OperationStatus {
	return v.value
}

func (v *NullableOperationStatus) Set(val *OperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationStatus(val *OperationStatus) *NullableOperationStatus {
	return &NullableOperationStatus{value: val, isSet: true}
}

func (v NullableOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



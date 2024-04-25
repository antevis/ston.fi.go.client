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

// checks if the DexSwapStatus200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DexSwapStatus200ResponseOneOf{}

// DexSwapStatus200ResponseOneOf struct for DexSwapStatus200ResponseOneOf
type DexSwapStatus200ResponseOneOf struct {
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
	Type string `json:"@type"`
}

type _DexSwapStatus200ResponseOneOf DexSwapStatus200ResponseOneOf

// NewDexSwapStatus200ResponseOneOf instantiates a new DexSwapStatus200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDexSwapStatus200ResponseOneOf(address string, balanceDeltas string, coins string, exitCode string, logicalTime string, queryId string, txHash *os.File, type_ string) *DexSwapStatus200ResponseOneOf {
	this := DexSwapStatus200ResponseOneOf{}
	this.Address = address
	this.BalanceDeltas = balanceDeltas
	this.Coins = coins
	this.ExitCode = exitCode
	this.LogicalTime = logicalTime
	this.QueryId = queryId
	this.TxHash = txHash
	this.Type = type_
	return &this
}

// NewDexSwapStatus200ResponseOneOfWithDefaults instantiates a new DexSwapStatus200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDexSwapStatus200ResponseOneOfWithDefaults() *DexSwapStatus200ResponseOneOf {
	this := DexSwapStatus200ResponseOneOf{}
	return &this
}

// GetAddress returns the Address field value
func (o *DexSwapStatus200ResponseOneOf) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DexSwapStatus200ResponseOneOf) SetAddress(v string) {
	o.Address = v
}

// GetBalanceDeltas returns the BalanceDeltas field value
func (o *DexSwapStatus200ResponseOneOf) GetBalanceDeltas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceDeltas
}

// GetBalanceDeltasOk returns a tuple with the BalanceDeltas field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf) GetBalanceDeltasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceDeltas, true
}

// SetBalanceDeltas sets field value
func (o *DexSwapStatus200ResponseOneOf) SetBalanceDeltas(v string) {
	o.BalanceDeltas = v
}

// GetCoins returns the Coins field value
func (o *DexSwapStatus200ResponseOneOf) GetCoins() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coins
}

// GetCoinsOk returns a tuple with the Coins field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf) GetCoinsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coins, true
}

// SetCoins sets field value
func (o *DexSwapStatus200ResponseOneOf) SetCoins(v string) {
	o.Coins = v
}

// GetExitCode returns the ExitCode field value
func (o *DexSwapStatus200ResponseOneOf) GetExitCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf) GetExitCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExitCode, true
}

// SetExitCode sets field value
func (o *DexSwapStatus200ResponseOneOf) SetExitCode(v string) {
	o.ExitCode = v
}

// GetLogicalTime returns the LogicalTime field value
func (o *DexSwapStatus200ResponseOneOf) GetLogicalTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogicalTime
}

// GetLogicalTimeOk returns a tuple with the LogicalTime field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf) GetLogicalTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogicalTime, true
}

// SetLogicalTime sets field value
func (o *DexSwapStatus200ResponseOneOf) SetLogicalTime(v string) {
	o.LogicalTime = v
}

// GetQueryId returns the QueryId field value
func (o *DexSwapStatus200ResponseOneOf) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *DexSwapStatus200ResponseOneOf) SetQueryId(v string) {
	o.QueryId = v
}

// GetTxHash returns the TxHash field value
func (o *DexSwapStatus200ResponseOneOf) GetTxHash() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf) GetTxHashOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *DexSwapStatus200ResponseOneOf) SetTxHash(v *os.File) {
	o.TxHash = v
}

// GetType returns the Type field value
func (o *DexSwapStatus200ResponseOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DexSwapStatus200ResponseOneOf) SetType(v string) {
	o.Type = v
}

func (o DexSwapStatus200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DexSwapStatus200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["balance_deltas"] = o.BalanceDeltas
	toSerialize["coins"] = o.Coins
	toSerialize["exit_code"] = o.ExitCode
	toSerialize["logical_time"] = o.LogicalTime
	toSerialize["query_id"] = o.QueryId
	toSerialize["tx_hash"] = o.TxHash
	toSerialize["@type"] = o.Type
	return toSerialize, nil
}

func (o *DexSwapStatus200ResponseOneOf) UnmarshalJSON(data []byte) (err error) {
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
		"@type",
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

	varDexSwapStatus200ResponseOneOf := _DexSwapStatus200ResponseOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDexSwapStatus200ResponseOneOf)

	if err != nil {
		return err
	}

	*o = DexSwapStatus200ResponseOneOf(varDexSwapStatus200ResponseOneOf)

	return err
}

type NullableDexSwapStatus200ResponseOneOf struct {
	value *DexSwapStatus200ResponseOneOf
	isSet bool
}

func (v NullableDexSwapStatus200ResponseOneOf) Get() *DexSwapStatus200ResponseOneOf {
	return v.value
}

func (v *NullableDexSwapStatus200ResponseOneOf) Set(val *DexSwapStatus200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDexSwapStatus200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDexSwapStatus200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDexSwapStatus200ResponseOneOf(val *DexSwapStatus200ResponseOneOf) *NullableDexSwapStatus200ResponseOneOf {
	return &NullableDexSwapStatus200ResponseOneOf{value: val, isSet: true}
}

func (v NullableDexSwapStatus200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDexSwapStatus200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



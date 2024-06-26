/*
ston-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DexSwapStatus200ResponseOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DexSwapStatus200ResponseOneOf1{}

// DexSwapStatus200ResponseOneOf1 struct for DexSwapStatus200ResponseOneOf1
type DexSwapStatus200ResponseOneOf1 struct {
	Type string `json:"@type"`
}

type _DexSwapStatus200ResponseOneOf1 DexSwapStatus200ResponseOneOf1

// NewDexSwapStatus200ResponseOneOf1 instantiates a new DexSwapStatus200ResponseOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDexSwapStatus200ResponseOneOf1(type_ string) *DexSwapStatus200ResponseOneOf1 {
	this := DexSwapStatus200ResponseOneOf1{}
	this.Type = type_
	return &this
}

// NewDexSwapStatus200ResponseOneOf1WithDefaults instantiates a new DexSwapStatus200ResponseOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDexSwapStatus200ResponseOneOf1WithDefaults() *DexSwapStatus200ResponseOneOf1 {
	this := DexSwapStatus200ResponseOneOf1{}
	return &this
}

// GetType returns the Type field value
func (o *DexSwapStatus200ResponseOneOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DexSwapStatus200ResponseOneOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DexSwapStatus200ResponseOneOf1) SetType(v string) {
	o.Type = v
}

func (o DexSwapStatus200ResponseOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DexSwapStatus200ResponseOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["@type"] = o.Type
	return toSerialize, nil
}

func (o *DexSwapStatus200ResponseOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDexSwapStatus200ResponseOneOf1 := _DexSwapStatus200ResponseOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDexSwapStatus200ResponseOneOf1)

	if err != nil {
		return err
	}

	*o = DexSwapStatus200ResponseOneOf1(varDexSwapStatus200ResponseOneOf1)

	return err
}

type NullableDexSwapStatus200ResponseOneOf1 struct {
	value *DexSwapStatus200ResponseOneOf1
	isSet bool
}

func (v NullableDexSwapStatus200ResponseOneOf1) Get() *DexSwapStatus200ResponseOneOf1 {
	return v.value
}

func (v *NullableDexSwapStatus200ResponseOneOf1) Set(val *DexSwapStatus200ResponseOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableDexSwapStatus200ResponseOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableDexSwapStatus200ResponseOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDexSwapStatus200ResponseOneOf1(val *DexSwapStatus200ResponseOneOf1) *NullableDexSwapStatus200ResponseOneOf1 {
	return &NullableDexSwapStatus200ResponseOneOf1{value: val, isSet: true}
}

func (v NullableDexSwapStatus200ResponseOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDexSwapStatus200ResponseOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the GetOperationStats200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOperationStats200Response{}

// GetOperationStats200Response struct for GetOperationStats200Response
type GetOperationStats200Response struct {
	Operations []OperationsInfo `json:"operations"`
}

type _GetOperationStats200Response GetOperationStats200Response

// NewGetOperationStats200Response instantiates a new GetOperationStats200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOperationStats200Response(operations []OperationsInfo) *GetOperationStats200Response {
	this := GetOperationStats200Response{}
	this.Operations = operations
	return &this
}

// NewGetOperationStats200ResponseWithDefaults instantiates a new GetOperationStats200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOperationStats200ResponseWithDefaults() *GetOperationStats200Response {
	this := GetOperationStats200Response{}
	return &this
}

// GetOperations returns the Operations field value
func (o *GetOperationStats200Response) GetOperations() []OperationsInfo {
	if o == nil {
		var ret []OperationsInfo
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *GetOperationStats200Response) GetOperationsOk() ([]OperationsInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *GetOperationStats200Response) SetOperations(v []OperationsInfo) {
	o.Operations = v
}

func (o GetOperationStats200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOperationStats200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operations"] = o.Operations
	return toSerialize, nil
}

func (o *GetOperationStats200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operations",
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

	varGetOperationStats200Response := _GetOperationStats200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOperationStats200Response)

	if err != nil {
		return err
	}

	*o = GetOperationStats200Response(varGetOperationStats200Response)

	return err
}

type NullableGetOperationStats200Response struct {
	value *GetOperationStats200Response
	isSet bool
}

func (v NullableGetOperationStats200Response) Get() *GetOperationStats200Response {
	return v.value
}

func (v *NullableGetOperationStats200Response) Set(val *GetOperationStats200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOperationStats200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOperationStats200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOperationStats200Response(val *GetOperationStats200Response) *NullableGetOperationStats200Response {
	return &NullableGetOperationStats200Response{value: val, isSet: true}
}

func (v NullableGetOperationStats200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOperationStats200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



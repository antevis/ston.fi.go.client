/*
ston-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// DexSwapStatus200Response - struct for DexSwapStatus200Response
type DexSwapStatus200Response struct {
	DexSwapStatus200ResponseOneOf *DexSwapStatus200ResponseOneOf
	DexSwapStatus200ResponseOneOf1 *DexSwapStatus200ResponseOneOf1
}

// DexSwapStatus200ResponseOneOfAsDexSwapStatus200Response is a convenience function that returns DexSwapStatus200ResponseOneOf wrapped in DexSwapStatus200Response
func DexSwapStatus200ResponseOneOfAsDexSwapStatus200Response(v *DexSwapStatus200ResponseOneOf) DexSwapStatus200Response {
	return DexSwapStatus200Response{
		DexSwapStatus200ResponseOneOf: v,
	}
}

// DexSwapStatus200ResponseOneOf1AsDexSwapStatus200Response is a convenience function that returns DexSwapStatus200ResponseOneOf1 wrapped in DexSwapStatus200Response
func DexSwapStatus200ResponseOneOf1AsDexSwapStatus200Response(v *DexSwapStatus200ResponseOneOf1) DexSwapStatus200Response {
	return DexSwapStatus200Response{
		DexSwapStatus200ResponseOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DexSwapStatus200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DexSwapStatus200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.DexSwapStatus200ResponseOneOf)
	if err == nil {
		jsonDexSwapStatus200ResponseOneOf, _ := json.Marshal(dst.DexSwapStatus200ResponseOneOf)
		if string(jsonDexSwapStatus200ResponseOneOf) == "{}" { // empty struct
			dst.DexSwapStatus200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DexSwapStatus200ResponseOneOf = nil
	}

	// try to unmarshal data into DexSwapStatus200ResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.DexSwapStatus200ResponseOneOf1)
	if err == nil {
		jsonDexSwapStatus200ResponseOneOf1, _ := json.Marshal(dst.DexSwapStatus200ResponseOneOf1)
		if string(jsonDexSwapStatus200ResponseOneOf1) == "{}" { // empty struct
			dst.DexSwapStatus200ResponseOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.DexSwapStatus200ResponseOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DexSwapStatus200ResponseOneOf = nil
		dst.DexSwapStatus200ResponseOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DexSwapStatus200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DexSwapStatus200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DexSwapStatus200Response) MarshalJSON() ([]byte, error) {
	if src.DexSwapStatus200ResponseOneOf != nil {
		return json.Marshal(&src.DexSwapStatus200ResponseOneOf)
	}

	if src.DexSwapStatus200ResponseOneOf1 != nil {
		return json.Marshal(&src.DexSwapStatus200ResponseOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DexSwapStatus200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DexSwapStatus200ResponseOneOf != nil {
		return obj.DexSwapStatus200ResponseOneOf
	}

	if obj.DexSwapStatus200ResponseOneOf1 != nil {
		return obj.DexSwapStatus200ResponseOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableDexSwapStatus200Response struct {
	value *DexSwapStatus200Response
	isSet bool
}

func (v NullableDexSwapStatus200Response) Get() *DexSwapStatus200Response {
	return v.value
}

func (v *NullableDexSwapStatus200Response) Set(val *DexSwapStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDexSwapStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDexSwapStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDexSwapStatus200Response(val *DexSwapStatus200Response) *NullableDexSwapStatus200Response {
	return &NullableDexSwapStatus200Response{value: val, isSet: true}
}

func (v NullableDexSwapStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDexSwapStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// checks if the AssetInfoSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetInfoSchema{}

// AssetInfoSchema struct for AssetInfoSchema
type AssetInfoSchema struct {
	// Asset balance
	Balance NullableString `json:"balance,omitempty"`
	// Asset is blacklisted
	Blacklisted bool `json:"blacklisted"`
	// Asset is community created
	Community bool `json:"community"`
	// Address of smart contract
	ContractAddress string `json:"contract_address"`
	// Number of decimal places used to represent fractional amounts of an asset
	Decimals int32 `json:"decimals"`
	// Asset has default symbol
	DefaultSymbol bool `json:"default_symbol"`
	// Asset is deprecated
	Deprecated bool `json:"deprecated"`
	// Ston internal asset price in USD
	DexPriceUsd NullableString `json:"dex_price_usd,omitempty"`
	// [DEPRECATED] Ston internal asset price in USD
	DexUsdPrice NullableString `json:"dex_usd_price,omitempty"`
	// Displayable name
	DisplayName NullableString `json:"display_name,omitempty"`
	// URL to asset image
	ImageUrl NullableString `json:"image_url,omitempty"`
	Kind AssetKindSchema `json:"kind"`
	// Asset symbol
	Symbol string `json:"symbol"`
	// Asset price in USD
	ThirdPartyPriceUsd NullableString `json:"third_party_price_usd,omitempty"`
	// [DEPRECATED] Asset price in USD
	ThirdPartyUsdPrice NullableString `json:"third_party_usd_price,omitempty"`
	// Asset wallet address
	WalletAddress NullableString `json:"wallet_address,omitempty"`
}

type _AssetInfoSchema AssetInfoSchema

// NewAssetInfoSchema instantiates a new AssetInfoSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetInfoSchema(blacklisted bool, community bool, contractAddress string, decimals int32, defaultSymbol bool, deprecated bool, kind AssetKindSchema, symbol string) *AssetInfoSchema {
	this := AssetInfoSchema{}
	this.Blacklisted = blacklisted
	this.Community = community
	this.ContractAddress = contractAddress
	this.Decimals = decimals
	this.DefaultSymbol = defaultSymbol
	this.Deprecated = deprecated
	this.Kind = kind
	this.Symbol = symbol
	return &this
}

// NewAssetInfoSchemaWithDefaults instantiates a new AssetInfoSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetInfoSchemaWithDefaults() *AssetInfoSchema {
	this := AssetInfoSchema{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoSchema) GetBalance() string {
	if o == nil || IsNil(o.Balance.Get()) {
		var ret string
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoSchema) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *AssetInfoSchema) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableString and assigns it to the Balance field.
func (o *AssetInfoSchema) SetBalance(v string) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *AssetInfoSchema) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *AssetInfoSchema) UnsetBalance() {
	o.Balance.Unset()
}

// GetBlacklisted returns the Blacklisted field value
func (o *AssetInfoSchema) GetBlacklisted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Blacklisted
}

// GetBlacklistedOk returns a tuple with the Blacklisted field value
// and a boolean to check if the value has been set.
func (o *AssetInfoSchema) GetBlacklistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blacklisted, true
}

// SetBlacklisted sets field value
func (o *AssetInfoSchema) SetBlacklisted(v bool) {
	o.Blacklisted = v
}

// GetCommunity returns the Community field value
func (o *AssetInfoSchema) GetCommunity() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Community
}

// GetCommunityOk returns a tuple with the Community field value
// and a boolean to check if the value has been set.
func (o *AssetInfoSchema) GetCommunityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Community, true
}

// SetCommunity sets field value
func (o *AssetInfoSchema) SetCommunity(v bool) {
	o.Community = v
}

// GetContractAddress returns the ContractAddress field value
func (o *AssetInfoSchema) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *AssetInfoSchema) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *AssetInfoSchema) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetDecimals returns the Decimals field value
func (o *AssetInfoSchema) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *AssetInfoSchema) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *AssetInfoSchema) SetDecimals(v int32) {
	o.Decimals = v
}

// GetDefaultSymbol returns the DefaultSymbol field value
func (o *AssetInfoSchema) GetDefaultSymbol() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultSymbol
}

// GetDefaultSymbolOk returns a tuple with the DefaultSymbol field value
// and a boolean to check if the value has been set.
func (o *AssetInfoSchema) GetDefaultSymbolOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSymbol, true
}

// SetDefaultSymbol sets field value
func (o *AssetInfoSchema) SetDefaultSymbol(v bool) {
	o.DefaultSymbol = v
}

// GetDeprecated returns the Deprecated field value
func (o *AssetInfoSchema) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *AssetInfoSchema) GetDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *AssetInfoSchema) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetDexPriceUsd returns the DexPriceUsd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoSchema) GetDexPriceUsd() string {
	if o == nil || IsNil(o.DexPriceUsd.Get()) {
		var ret string
		return ret
	}
	return *o.DexPriceUsd.Get()
}

// GetDexPriceUsdOk returns a tuple with the DexPriceUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoSchema) GetDexPriceUsdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DexPriceUsd.Get(), o.DexPriceUsd.IsSet()
}

// HasDexPriceUsd returns a boolean if a field has been set.
func (o *AssetInfoSchema) HasDexPriceUsd() bool {
	if o != nil && o.DexPriceUsd.IsSet() {
		return true
	}

	return false
}

// SetDexPriceUsd gets a reference to the given NullableString and assigns it to the DexPriceUsd field.
func (o *AssetInfoSchema) SetDexPriceUsd(v string) {
	o.DexPriceUsd.Set(&v)
}
// SetDexPriceUsdNil sets the value for DexPriceUsd to be an explicit nil
func (o *AssetInfoSchema) SetDexPriceUsdNil() {
	o.DexPriceUsd.Set(nil)
}

// UnsetDexPriceUsd ensures that no value is present for DexPriceUsd, not even an explicit nil
func (o *AssetInfoSchema) UnsetDexPriceUsd() {
	o.DexPriceUsd.Unset()
}

// GetDexUsdPrice returns the DexUsdPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoSchema) GetDexUsdPrice() string {
	if o == nil || IsNil(o.DexUsdPrice.Get()) {
		var ret string
		return ret
	}
	return *o.DexUsdPrice.Get()
}

// GetDexUsdPriceOk returns a tuple with the DexUsdPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoSchema) GetDexUsdPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DexUsdPrice.Get(), o.DexUsdPrice.IsSet()
}

// HasDexUsdPrice returns a boolean if a field has been set.
func (o *AssetInfoSchema) HasDexUsdPrice() bool {
	if o != nil && o.DexUsdPrice.IsSet() {
		return true
	}

	return false
}

// SetDexUsdPrice gets a reference to the given NullableString and assigns it to the DexUsdPrice field.
func (o *AssetInfoSchema) SetDexUsdPrice(v string) {
	o.DexUsdPrice.Set(&v)
}
// SetDexUsdPriceNil sets the value for DexUsdPrice to be an explicit nil
func (o *AssetInfoSchema) SetDexUsdPriceNil() {
	o.DexUsdPrice.Set(nil)
}

// UnsetDexUsdPrice ensures that no value is present for DexUsdPrice, not even an explicit nil
func (o *AssetInfoSchema) UnsetDexUsdPrice() {
	o.DexUsdPrice.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoSchema) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoSchema) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AssetInfoSchema) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AssetInfoSchema) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AssetInfoSchema) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AssetInfoSchema) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoSchema) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoSchema) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *AssetInfoSchema) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *AssetInfoSchema) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *AssetInfoSchema) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *AssetInfoSchema) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetKind returns the Kind field value
func (o *AssetInfoSchema) GetKind() AssetKindSchema {
	if o == nil {
		var ret AssetKindSchema
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AssetInfoSchema) GetKindOk() (*AssetKindSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AssetInfoSchema) SetKind(v AssetKindSchema) {
	o.Kind = v
}

// GetSymbol returns the Symbol field value
func (o *AssetInfoSchema) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *AssetInfoSchema) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *AssetInfoSchema) SetSymbol(v string) {
	o.Symbol = v
}

// GetThirdPartyPriceUsd returns the ThirdPartyPriceUsd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoSchema) GetThirdPartyPriceUsd() string {
	if o == nil || IsNil(o.ThirdPartyPriceUsd.Get()) {
		var ret string
		return ret
	}
	return *o.ThirdPartyPriceUsd.Get()
}

// GetThirdPartyPriceUsdOk returns a tuple with the ThirdPartyPriceUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoSchema) GetThirdPartyPriceUsdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyPriceUsd.Get(), o.ThirdPartyPriceUsd.IsSet()
}

// HasThirdPartyPriceUsd returns a boolean if a field has been set.
func (o *AssetInfoSchema) HasThirdPartyPriceUsd() bool {
	if o != nil && o.ThirdPartyPriceUsd.IsSet() {
		return true
	}

	return false
}

// SetThirdPartyPriceUsd gets a reference to the given NullableString and assigns it to the ThirdPartyPriceUsd field.
func (o *AssetInfoSchema) SetThirdPartyPriceUsd(v string) {
	o.ThirdPartyPriceUsd.Set(&v)
}
// SetThirdPartyPriceUsdNil sets the value for ThirdPartyPriceUsd to be an explicit nil
func (o *AssetInfoSchema) SetThirdPartyPriceUsdNil() {
	o.ThirdPartyPriceUsd.Set(nil)
}

// UnsetThirdPartyPriceUsd ensures that no value is present for ThirdPartyPriceUsd, not even an explicit nil
func (o *AssetInfoSchema) UnsetThirdPartyPriceUsd() {
	o.ThirdPartyPriceUsd.Unset()
}

// GetThirdPartyUsdPrice returns the ThirdPartyUsdPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoSchema) GetThirdPartyUsdPrice() string {
	if o == nil || IsNil(o.ThirdPartyUsdPrice.Get()) {
		var ret string
		return ret
	}
	return *o.ThirdPartyUsdPrice.Get()
}

// GetThirdPartyUsdPriceOk returns a tuple with the ThirdPartyUsdPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoSchema) GetThirdPartyUsdPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyUsdPrice.Get(), o.ThirdPartyUsdPrice.IsSet()
}

// HasThirdPartyUsdPrice returns a boolean if a field has been set.
func (o *AssetInfoSchema) HasThirdPartyUsdPrice() bool {
	if o != nil && o.ThirdPartyUsdPrice.IsSet() {
		return true
	}

	return false
}

// SetThirdPartyUsdPrice gets a reference to the given NullableString and assigns it to the ThirdPartyUsdPrice field.
func (o *AssetInfoSchema) SetThirdPartyUsdPrice(v string) {
	o.ThirdPartyUsdPrice.Set(&v)
}
// SetThirdPartyUsdPriceNil sets the value for ThirdPartyUsdPrice to be an explicit nil
func (o *AssetInfoSchema) SetThirdPartyUsdPriceNil() {
	o.ThirdPartyUsdPrice.Set(nil)
}

// UnsetThirdPartyUsdPrice ensures that no value is present for ThirdPartyUsdPrice, not even an explicit nil
func (o *AssetInfoSchema) UnsetThirdPartyUsdPrice() {
	o.ThirdPartyUsdPrice.Unset()
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetInfoSchema) GetWalletAddress() string {
	if o == nil || IsNil(o.WalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.WalletAddress.Get()
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetInfoSchema) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletAddress.Get(), o.WalletAddress.IsSet()
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *AssetInfoSchema) HasWalletAddress() bool {
	if o != nil && o.WalletAddress.IsSet() {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given NullableString and assigns it to the WalletAddress field.
func (o *AssetInfoSchema) SetWalletAddress(v string) {
	o.WalletAddress.Set(&v)
}
// SetWalletAddressNil sets the value for WalletAddress to be an explicit nil
func (o *AssetInfoSchema) SetWalletAddressNil() {
	o.WalletAddress.Set(nil)
}

// UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil
func (o *AssetInfoSchema) UnsetWalletAddress() {
	o.WalletAddress.Unset()
}

func (o AssetInfoSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetInfoSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	toSerialize["blacklisted"] = o.Blacklisted
	toSerialize["community"] = o.Community
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["decimals"] = o.Decimals
	toSerialize["default_symbol"] = o.DefaultSymbol
	toSerialize["deprecated"] = o.Deprecated
	if o.DexPriceUsd.IsSet() {
		toSerialize["dex_price_usd"] = o.DexPriceUsd.Get()
	}
	if o.DexUsdPrice.IsSet() {
		toSerialize["dex_usd_price"] = o.DexUsdPrice.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	toSerialize["kind"] = o.Kind
	toSerialize["symbol"] = o.Symbol
	if o.ThirdPartyPriceUsd.IsSet() {
		toSerialize["third_party_price_usd"] = o.ThirdPartyPriceUsd.Get()
	}
	if o.ThirdPartyUsdPrice.IsSet() {
		toSerialize["third_party_usd_price"] = o.ThirdPartyUsdPrice.Get()
	}
	if o.WalletAddress.IsSet() {
		toSerialize["wallet_address"] = o.WalletAddress.Get()
	}
	return toSerialize, nil
}

func (o *AssetInfoSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"blacklisted",
		"community",
		"contract_address",
		"decimals",
		"default_symbol",
		"deprecated",
		"kind",
		"symbol",
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

	varAssetInfoSchema := _AssetInfoSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssetInfoSchema)

	if err != nil {
		return err
	}

	*o = AssetInfoSchema(varAssetInfoSchema)

	return err
}

type NullableAssetInfoSchema struct {
	value *AssetInfoSchema
	isSet bool
}

func (v NullableAssetInfoSchema) Get() *AssetInfoSchema {
	return v.value
}

func (v *NullableAssetInfoSchema) Set(val *AssetInfoSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetInfoSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetInfoSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetInfoSchema(val *AssetInfoSchema) *NullableAssetInfoSchema {
	return &NullableAssetInfoSchema{value: val, isSet: true}
}

func (v NullableAssetInfoSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetInfoSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



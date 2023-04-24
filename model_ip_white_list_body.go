/*
Camunda Management API

Manage Camunda Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the IpWhiteListBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpWhiteListBody{}

// IpWhiteListBody struct for IpWhiteListBody
type IpWhiteListBody struct {
	Ipwhitelist []ClusterIpwhitelistInner `json:"ipwhitelist"`
}

// NewIpWhiteListBody instantiates a new IpWhiteListBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpWhiteListBody(ipwhitelist []ClusterIpwhitelistInner) *IpWhiteListBody {
	this := IpWhiteListBody{}
	this.Ipwhitelist = ipwhitelist
	return &this
}

// NewIpWhiteListBodyWithDefaults instantiates a new IpWhiteListBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpWhiteListBodyWithDefaults() *IpWhiteListBody {
	this := IpWhiteListBody{}
	return &this
}

// GetIpwhitelist returns the Ipwhitelist field value
func (o *IpWhiteListBody) GetIpwhitelist() []ClusterIpwhitelistInner {
	if o == nil {
		var ret []ClusterIpwhitelistInner
		return ret
	}

	return o.Ipwhitelist
}

// GetIpwhitelistOk returns a tuple with the Ipwhitelist field value
// and a boolean to check if the value has been set.
func (o *IpWhiteListBody) GetIpwhitelistOk() ([]ClusterIpwhitelistInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipwhitelist, true
}

// SetIpwhitelist sets field value
func (o *IpWhiteListBody) SetIpwhitelist(v []ClusterIpwhitelistInner) {
	o.Ipwhitelist = v
}

func (o IpWhiteListBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpWhiteListBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ipwhitelist"] = o.Ipwhitelist
	return toSerialize, nil
}

type NullableIpWhiteListBody struct {
	value *IpWhiteListBody
	isSet bool
}

func (v NullableIpWhiteListBody) Get() *IpWhiteListBody {
	return v.value
}

func (v *NullableIpWhiteListBody) Set(val *IpWhiteListBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIpWhiteListBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIpWhiteListBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpWhiteListBody(val *IpWhiteListBody) *NullableIpWhiteListBody {
	return &NullableIpWhiteListBody{value: val, isSet: true}
}

func (v NullableIpWhiteListBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpWhiteListBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


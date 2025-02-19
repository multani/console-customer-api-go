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

// checks if the ClusterIpwhitelistInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterIpwhitelistInner{}

// ClusterIpwhitelistInner struct for ClusterIpwhitelistInner
type ClusterIpwhitelistInner struct {
	Description string `json:"description"`
	Ip          string `json:"ip"`
}

// NewClusterIpwhitelistInner instantiates a new ClusterIpwhitelistInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterIpwhitelistInner(description string, ip string) *ClusterIpwhitelistInner {
	this := ClusterIpwhitelistInner{}
	this.Description = description
	this.Ip = ip
	return &this
}

// NewClusterIpwhitelistInnerWithDefaults instantiates a new ClusterIpwhitelistInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterIpwhitelistInnerWithDefaults() *ClusterIpwhitelistInner {
	this := ClusterIpwhitelistInner{}
	return &this
}

// GetDescription returns the Description field value
func (o *ClusterIpwhitelistInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ClusterIpwhitelistInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ClusterIpwhitelistInner) SetDescription(v string) {
	o.Description = v
}

// GetIp returns the Ip field value
func (o *ClusterIpwhitelistInner) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *ClusterIpwhitelistInner) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *ClusterIpwhitelistInner) SetIp(v string) {
	o.Ip = v
}

func (o ClusterIpwhitelistInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterIpwhitelistInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["ip"] = o.Ip
	return toSerialize, nil
}

type NullableClusterIpwhitelistInner struct {
	value *ClusterIpwhitelistInner
	isSet bool
}

func (v NullableClusterIpwhitelistInner) Get() *ClusterIpwhitelistInner {
	return v.value
}

func (v *NullableClusterIpwhitelistInner) Set(val *ClusterIpwhitelistInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterIpwhitelistInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterIpwhitelistInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterIpwhitelistInner(val *ClusterIpwhitelistInner) *NullableClusterIpwhitelistInner {
	return &NullableClusterIpwhitelistInner{value: val, isSet: true}
}

func (v NullableClusterIpwhitelistInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterIpwhitelistInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

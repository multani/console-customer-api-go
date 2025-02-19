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

// checks if the ParametersChannelsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParametersChannelsInner{}

// ParametersChannelsInner struct for ParametersChannelsInner
type ParametersChannelsInner struct {
	AllowedGenerations []ParametersChannelsInnerAllowedGenerationsInner `json:"allowedGenerations"`
	DefaultGeneration  ParametersChannelsInnerAllowedGenerationsInner   `json:"defaultGeneration"`
	Name               string                                           `json:"name"`
	Uuid               string                                           `json:"uuid"`
}

// NewParametersChannelsInner instantiates a new ParametersChannelsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParametersChannelsInner(allowedGenerations []ParametersChannelsInnerAllowedGenerationsInner, defaultGeneration ParametersChannelsInnerAllowedGenerationsInner, name string, uuid string) *ParametersChannelsInner {
	this := ParametersChannelsInner{}
	this.AllowedGenerations = allowedGenerations
	this.DefaultGeneration = defaultGeneration
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewParametersChannelsInnerWithDefaults instantiates a new ParametersChannelsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersChannelsInnerWithDefaults() *ParametersChannelsInner {
	this := ParametersChannelsInner{}
	return &this
}

// GetAllowedGenerations returns the AllowedGenerations field value
func (o *ParametersChannelsInner) GetAllowedGenerations() []ParametersChannelsInnerAllowedGenerationsInner {
	if o == nil {
		var ret []ParametersChannelsInnerAllowedGenerationsInner
		return ret
	}

	return o.AllowedGenerations
}

// GetAllowedGenerationsOk returns a tuple with the AllowedGenerations field value
// and a boolean to check if the value has been set.
func (o *ParametersChannelsInner) GetAllowedGenerationsOk() ([]ParametersChannelsInnerAllowedGenerationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedGenerations, true
}

// SetAllowedGenerations sets field value
func (o *ParametersChannelsInner) SetAllowedGenerations(v []ParametersChannelsInnerAllowedGenerationsInner) {
	o.AllowedGenerations = v
}

// GetDefaultGeneration returns the DefaultGeneration field value
func (o *ParametersChannelsInner) GetDefaultGeneration() ParametersChannelsInnerAllowedGenerationsInner {
	if o == nil {
		var ret ParametersChannelsInnerAllowedGenerationsInner
		return ret
	}

	return o.DefaultGeneration
}

// GetDefaultGenerationOk returns a tuple with the DefaultGeneration field value
// and a boolean to check if the value has been set.
func (o *ParametersChannelsInner) GetDefaultGenerationOk() (*ParametersChannelsInnerAllowedGenerationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultGeneration, true
}

// SetDefaultGeneration sets field value
func (o *ParametersChannelsInner) SetDefaultGeneration(v ParametersChannelsInnerAllowedGenerationsInner) {
	o.DefaultGeneration = v
}

// GetName returns the Name field value
func (o *ParametersChannelsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParametersChannelsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParametersChannelsInner) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *ParametersChannelsInner) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ParametersChannelsInner) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ParametersChannelsInner) SetUuid(v string) {
	o.Uuid = v
}

func (o ParametersChannelsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParametersChannelsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowedGenerations"] = o.AllowedGenerations
	toSerialize["defaultGeneration"] = o.DefaultGeneration
	toSerialize["name"] = o.Name
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

type NullableParametersChannelsInner struct {
	value *ParametersChannelsInner
	isSet bool
}

func (v NullableParametersChannelsInner) Get() *ParametersChannelsInner {
	return v.value
}

func (v *NullableParametersChannelsInner) Set(val *ParametersChannelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableParametersChannelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableParametersChannelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParametersChannelsInner(val *ParametersChannelsInner) *NullableParametersChannelsInner {
	return &NullableParametersChannelsInner{value: val, isSet: true}
}

func (v NullableParametersChannelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParametersChannelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

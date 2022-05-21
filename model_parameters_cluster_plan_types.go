/*
Camunda Cloud Management API

Manage Camunda Cloud Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ParametersClusterPlanTypes struct for ParametersClusterPlanTypes
type ParametersClusterPlanTypes struct {
	Region ParametersAllowedGenerations `json:"region"`
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

// NewParametersClusterPlanTypes instantiates a new ParametersClusterPlanTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParametersClusterPlanTypes(region ParametersAllowedGenerations, name string, uuid string) *ParametersClusterPlanTypes {
	this := ParametersClusterPlanTypes{}
	this.Region = region
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewParametersClusterPlanTypesWithDefaults instantiates a new ParametersClusterPlanTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersClusterPlanTypesWithDefaults() *ParametersClusterPlanTypes {
	this := ParametersClusterPlanTypes{}
	return &this
}

// GetRegion returns the Region field value
func (o *ParametersClusterPlanTypes) GetRegion() ParametersAllowedGenerations {
	if o == nil {
		var ret ParametersAllowedGenerations
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ParametersClusterPlanTypes) GetRegionOk() (*ParametersAllowedGenerations, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ParametersClusterPlanTypes) SetRegion(v ParametersAllowedGenerations) {
	o.Region = v
}

// GetName returns the Name field value
func (o *ParametersClusterPlanTypes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParametersClusterPlanTypes) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParametersClusterPlanTypes) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *ParametersClusterPlanTypes) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ParametersClusterPlanTypes) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ParametersClusterPlanTypes) SetUuid(v string) {
	o.Uuid = v
}

func (o ParametersClusterPlanTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableParametersClusterPlanTypes struct {
	value *ParametersClusterPlanTypes
	isSet bool
}

func (v NullableParametersClusterPlanTypes) Get() *ParametersClusterPlanTypes {
	return v.value
}

func (v *NullableParametersClusterPlanTypes) Set(val *ParametersClusterPlanTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableParametersClusterPlanTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableParametersClusterPlanTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParametersClusterPlanTypes(val *ParametersClusterPlanTypes) *NullableParametersClusterPlanTypes {
	return &NullableParametersClusterPlanTypes{value: val, isSet: true}
}

func (v NullableParametersClusterPlanTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParametersClusterPlanTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



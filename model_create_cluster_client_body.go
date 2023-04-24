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

// checks if the CreateClusterClientBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClusterClientBody{}

// CreateClusterClientBody struct for CreateClusterClientBody
type CreateClusterClientBody struct {
	ClientName string `json:"clientName"`
	Permissions []string `json:"permissions,omitempty"`
}

// NewCreateClusterClientBody instantiates a new CreateClusterClientBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterClientBody(clientName string) *CreateClusterClientBody {
	this := CreateClusterClientBody{}
	this.ClientName = clientName
	return &this
}

// NewCreateClusterClientBodyWithDefaults instantiates a new CreateClusterClientBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterClientBodyWithDefaults() *CreateClusterClientBody {
	this := CreateClusterClientBody{}
	return &this
}

// GetClientName returns the ClientName field value
func (o *CreateClusterClientBody) GetClientName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
func (o *CreateClusterClientBody) GetClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientName, true
}

// SetClientName sets field value
func (o *CreateClusterClientBody) SetClientName(v string) {
	o.ClientName = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CreateClusterClientBody) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterClientBody) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateClusterClientBody) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *CreateClusterClientBody) SetPermissions(v []string) {
	o.Permissions = v
}

func (o CreateClusterClientBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClusterClientBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientName"] = o.ClientName
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableCreateClusterClientBody struct {
	value *CreateClusterClientBody
	isSet bool
}

func (v NullableCreateClusterClientBody) Get() *CreateClusterClientBody {
	return v.value
}

func (v *NullableCreateClusterClientBody) Set(val *CreateClusterClientBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterClientBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterClientBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterClientBody(val *CreateClusterClientBody) *NullableCreateClusterClientBody {
	return &NullableCreateClusterClientBody{value: val, isSet: true}
}

func (v NullableCreateClusterClientBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterClientBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



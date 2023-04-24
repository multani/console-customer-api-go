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

// checks if the ClusterClientConnectionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterClientConnectionDetails{}

// ClusterClientConnectionDetails struct for ClusterClientConnectionDetails
type ClusterClientConnectionDetails struct {
	Name string `json:"name"`
	ZEEBE_ADDRESS string `json:"ZEEBE_ADDRESS"`
	ZEEBE_CLIENT_ID string `json:"ZEEBE_CLIENT_ID"`
	ZEEBE_AUTHORIZATION_SERVER_URL string `json:"ZEEBE_AUTHORIZATION_SERVER_URL"`
}

// NewClusterClientConnectionDetails instantiates a new ClusterClientConnectionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterClientConnectionDetails(name string, zEEBEADDRESS string, zEEBECLIENTID string, zEEBEAUTHORIZATIONSERVERURL string) *ClusterClientConnectionDetails {
	this := ClusterClientConnectionDetails{}
	this.Name = name
	this.ZEEBE_ADDRESS = zEEBEADDRESS
	this.ZEEBE_CLIENT_ID = zEEBECLIENTID
	this.ZEEBE_AUTHORIZATION_SERVER_URL = zEEBEAUTHORIZATIONSERVERURL
	return &this
}

// NewClusterClientConnectionDetailsWithDefaults instantiates a new ClusterClientConnectionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterClientConnectionDetailsWithDefaults() *ClusterClientConnectionDetails {
	this := ClusterClientConnectionDetails{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterClientConnectionDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterClientConnectionDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterClientConnectionDetails) SetName(v string) {
	o.Name = v
}

// GetZEEBE_ADDRESS returns the ZEEBE_ADDRESS field value
func (o *ClusterClientConnectionDetails) GetZEEBE_ADDRESS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZEEBE_ADDRESS
}

// GetZEEBE_ADDRESSOk returns a tuple with the ZEEBE_ADDRESS field value
// and a boolean to check if the value has been set.
func (o *ClusterClientConnectionDetails) GetZEEBE_ADDRESSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZEEBE_ADDRESS, true
}

// SetZEEBE_ADDRESS sets field value
func (o *ClusterClientConnectionDetails) SetZEEBE_ADDRESS(v string) {
	o.ZEEBE_ADDRESS = v
}

// GetZEEBE_CLIENT_ID returns the ZEEBE_CLIENT_ID field value
func (o *ClusterClientConnectionDetails) GetZEEBE_CLIENT_ID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZEEBE_CLIENT_ID
}

// GetZEEBE_CLIENT_IDOk returns a tuple with the ZEEBE_CLIENT_ID field value
// and a boolean to check if the value has been set.
func (o *ClusterClientConnectionDetails) GetZEEBE_CLIENT_IDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZEEBE_CLIENT_ID, true
}

// SetZEEBE_CLIENT_ID sets field value
func (o *ClusterClientConnectionDetails) SetZEEBE_CLIENT_ID(v string) {
	o.ZEEBE_CLIENT_ID = v
}

// GetZEEBE_AUTHORIZATION_SERVER_URL returns the ZEEBE_AUTHORIZATION_SERVER_URL field value
func (o *ClusterClientConnectionDetails) GetZEEBE_AUTHORIZATION_SERVER_URL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZEEBE_AUTHORIZATION_SERVER_URL
}

// GetZEEBE_AUTHORIZATION_SERVER_URLOk returns a tuple with the ZEEBE_AUTHORIZATION_SERVER_URL field value
// and a boolean to check if the value has been set.
func (o *ClusterClientConnectionDetails) GetZEEBE_AUTHORIZATION_SERVER_URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZEEBE_AUTHORIZATION_SERVER_URL, true
}

// SetZEEBE_AUTHORIZATION_SERVER_URL sets field value
func (o *ClusterClientConnectionDetails) SetZEEBE_AUTHORIZATION_SERVER_URL(v string) {
	o.ZEEBE_AUTHORIZATION_SERVER_URL = v
}

func (o ClusterClientConnectionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterClientConnectionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["ZEEBE_ADDRESS"] = o.ZEEBE_ADDRESS
	toSerialize["ZEEBE_CLIENT_ID"] = o.ZEEBE_CLIENT_ID
	toSerialize["ZEEBE_AUTHORIZATION_SERVER_URL"] = o.ZEEBE_AUTHORIZATION_SERVER_URL
	return toSerialize, nil
}

type NullableClusterClientConnectionDetails struct {
	value *ClusterClientConnectionDetails
	isSet bool
}

func (v NullableClusterClientConnectionDetails) Get() *ClusterClientConnectionDetails {
	return v.value
}

func (v *NullableClusterClientConnectionDetails) Set(val *ClusterClientConnectionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterClientConnectionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterClientConnectionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterClientConnectionDetails(val *ClusterClientConnectionDetails) *NullableClusterClientConnectionDetails {
	return &NullableClusterClientConnectionDetails{value: val, isSet: true}
}

func (v NullableClusterClientConnectionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterClientConnectionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ClusterClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterClient{}

// ClusterClient struct for ClusterClient
type ClusterClient struct {
	Name string `json:"name"`
	ClientId string `json:"clientId"`
	Permissions []string `json:"permissions"`
}

// NewClusterClient instantiates a new ClusterClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterClient(name string, clientId string, permissions []string) *ClusterClient {
	this := ClusterClient{}
	this.Name = name
	this.ClientId = clientId
	this.Permissions = permissions
	return &this
}

// NewClusterClientWithDefaults instantiates a new ClusterClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterClientWithDefaults() *ClusterClient {
	this := ClusterClient{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterClient) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterClient) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterClient) SetName(v string) {
	o.Name = v
}

// GetClientId returns the ClientId field value
func (o *ClusterClient) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ClusterClient) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ClusterClient) SetClientId(v string) {
	o.ClientId = v
}

// GetPermissions returns the Permissions field value
func (o *ClusterClient) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ClusterClient) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ClusterClient) SetPermissions(v []string) {
	o.Permissions = v
}

func (o ClusterClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["clientId"] = o.ClientId
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

type NullableClusterClient struct {
	value *ClusterClient
	isSet bool
}

func (v NullableClusterClient) Get() *ClusterClient {
	return v.value
}

func (v *NullableClusterClient) Set(val *ClusterClient) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterClient) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterClient(val *ClusterClient) *NullableClusterClient {
	return &NullableClusterClient{value: val, isSet: true}
}

func (v NullableClusterClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



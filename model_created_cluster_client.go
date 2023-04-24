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

// checks if the CreatedClusterClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedClusterClient{}

// CreatedClusterClient struct for CreatedClusterClient
type CreatedClusterClient struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Permissions []string `json:"permissions"`
	Links *CreatedClusterClientLinks `json:"links,omitempty"`
}

// NewCreatedClusterClient instantiates a new CreatedClusterClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedClusterClient(name string, uuid string, clientId string, clientSecret string, permissions []string) *CreatedClusterClient {
	this := CreatedClusterClient{}
	this.Name = name
	this.Uuid = uuid
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Permissions = permissions
	return &this
}

// NewCreatedClusterClientWithDefaults instantiates a new CreatedClusterClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedClusterClientWithDefaults() *CreatedClusterClient {
	this := CreatedClusterClient{}
	return &this
}

// GetName returns the Name field value
func (o *CreatedClusterClient) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatedClusterClient) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatedClusterClient) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *CreatedClusterClient) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CreatedClusterClient) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *CreatedClusterClient) SetUuid(v string) {
	o.Uuid = v
}

// GetClientId returns the ClientId field value
func (o *CreatedClusterClient) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CreatedClusterClient) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *CreatedClusterClient) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *CreatedClusterClient) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *CreatedClusterClient) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *CreatedClusterClient) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetPermissions returns the Permissions field value
func (o *CreatedClusterClient) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *CreatedClusterClient) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *CreatedClusterClient) SetPermissions(v []string) {
	o.Permissions = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreatedClusterClient) GetLinks() CreatedClusterClientLinks {
	if o == nil || IsNil(o.Links) {
		var ret CreatedClusterClientLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedClusterClient) GetLinksOk() (*CreatedClusterClientLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreatedClusterClient) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CreatedClusterClientLinks and assigns it to the Links field.
func (o *CreatedClusterClient) SetLinks(v CreatedClusterClientLinks) {
	o.Links = &v
}

func (o CreatedClusterClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedClusterClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["uuid"] = o.Uuid
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["permissions"] = o.Permissions
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCreatedClusterClient struct {
	value *CreatedClusterClient
	isSet bool
}

func (v NullableCreatedClusterClient) Get() *CreatedClusterClient {
	return v.value
}

func (v *NullableCreatedClusterClient) Set(val *CreatedClusterClient) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedClusterClient) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedClusterClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedClusterClient(val *CreatedClusterClient) *NullableCreatedClusterClient {
	return &NullableCreatedClusterClient{value: val, isSet: true}
}

func (v NullableCreatedClusterClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedClusterClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



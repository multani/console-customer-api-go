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

// checks if the Member type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Member{}

// Member struct for Member
type Member struct {
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	Roles         []OrganizationRole `json:"roles"`
	InvitePending bool               `json:"invitePending"`
}

// NewMember instantiates a new Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMember(name string, email string, roles []OrganizationRole, invitePending bool) *Member {
	this := Member{}
	this.Name = name
	this.Email = email
	this.Roles = roles
	this.InvitePending = invitePending
	return &this
}

// NewMemberWithDefaults instantiates a new Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberWithDefaults() *Member {
	this := Member{}
	return &this
}

// GetName returns the Name field value
func (o *Member) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Member) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Member) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *Member) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Member) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Member) SetEmail(v string) {
	o.Email = v
}

// GetRoles returns the Roles field value
func (o *Member) GetRoles() []OrganizationRole {
	if o == nil {
		var ret []OrganizationRole
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *Member) GetRolesOk() ([]OrganizationRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *Member) SetRoles(v []OrganizationRole) {
	o.Roles = v
}

// GetInvitePending returns the InvitePending field value
func (o *Member) GetInvitePending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InvitePending
}

// GetInvitePendingOk returns a tuple with the InvitePending field value
// and a boolean to check if the value has been set.
func (o *Member) GetInvitePendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitePending, true
}

// SetInvitePending sets field value
func (o *Member) SetInvitePending(v bool) {
	o.InvitePending = v
}

func (o Member) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Member) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["roles"] = o.Roles
	toSerialize["invitePending"] = o.InvitePending
	return toSerialize, nil
}

type NullableMember struct {
	value *Member
	isSet bool
}

func (v NullableMember) Get() *Member {
	return v.value
}

func (v *NullableMember) Set(val *Member) {
	v.value = val
	v.isSet = true
}

func (v NullableMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMember(val *Member) *NullableMember {
	return &NullableMember{value: val, isSet: true}
}

func (v NullableMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

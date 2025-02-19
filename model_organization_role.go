/*
Camunda Management API

Manage Camunda Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// OrganizationRole the model 'OrganizationRole'
type OrganizationRole string

// List of OrganizationRole
const (
	MEMBER             OrganizationRole = "member"
	ADMIN              OrganizationRole = "admin"
	OWNER              OrganizationRole = "owner"
	SUPPORTAGENT       OrganizationRole = "supportagent"
	OPERATIONSENGINEER OrganizationRole = "operationsengineer"
	TASKUSER           OrganizationRole = "taskuser"
	ANALYST            OrganizationRole = "analyst"
	DEVELOPER          OrganizationRole = "developer"
	VISITOR            OrganizationRole = "visitor"
)

// All allowed values of OrganizationRole enum
var AllowedOrganizationRoleEnumValues = []OrganizationRole{
	"member",
	"admin",
	"owner",
	"supportagent",
	"operationsengineer",
	"taskuser",
	"analyst",
	"developer",
	"visitor",
}

func (v *OrganizationRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationRole(value)
	for _, existing := range AllowedOrganizationRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationRole", value)
}

// NewOrganizationRoleFromValue returns a pointer to a valid OrganizationRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationRoleFromValue(v string) (*OrganizationRole, error) {
	ev := OrganizationRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationRole: valid values are %v", v, AllowedOrganizationRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationRole) IsValid() bool {
	for _, existing := range AllowedOrganizationRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationRole value
func (v OrganizationRole) Ptr() *OrganizationRole {
	return &v
}

type NullableOrganizationRole struct {
	value *OrganizationRole
	isSet bool
}

func (v NullableOrganizationRole) Get() *OrganizationRole {
	return v.value
}

func (v *NullableOrganizationRole) Set(val *OrganizationRole) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationRole) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationRole(val *OrganizationRole) *NullableOrganizationRole {
	return &NullableOrganizationRole{value: val, isSet: true}
}

func (v NullableOrganizationRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// OrganizationRoleTASKUSER the model 'OrganizationRoleTASKUSER'
type OrganizationRoleTASKUSER string

// All allowed values of OrganizationRoleTASKUSER enum
var AllowedOrganizationRoleTASKUSEREnumValues = []OrganizationRoleTASKUSER{
	"taskuser",
}

func (v *OrganizationRoleTASKUSER) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationRoleTASKUSER(value)
	for _, existing := range AllowedOrganizationRoleTASKUSEREnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationRoleTASKUSER", value)
}

// NewOrganizationRoleTASKUSERFromValue returns a pointer to a valid OrganizationRoleTASKUSER
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationRoleTASKUSERFromValue(v string) (*OrganizationRoleTASKUSER, error) {
	ev := OrganizationRoleTASKUSER(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationRoleTASKUSER: valid values are %v", v, AllowedOrganizationRoleTASKUSEREnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationRoleTASKUSER) IsValid() bool {
	for _, existing := range AllowedOrganizationRoleTASKUSEREnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationRole.TASK_USER value
func (v OrganizationRoleTASKUSER) Ptr() *OrganizationRoleTASKUSER {
	return &v
}

type NullableOrganizationRoleTASKUSER struct {
	value *OrganizationRoleTASKUSER
	isSet bool
}

func (v NullableOrganizationRoleTASKUSER) Get() *OrganizationRoleTASKUSER {
	return v.value
}

func (v *NullableOrganizationRoleTASKUSER) Set(val *OrganizationRoleTASKUSER) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationRoleTASKUSER) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationRoleTASKUSER) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationRoleTASKUSER(val *OrganizationRoleTASKUSER) *NullableOrganizationRoleTASKUSER {
	return &NullableOrganizationRoleTASKUSER{value: val, isSet: true}
}

func (v NullableOrganizationRoleTASKUSER) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationRoleTASKUSER) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


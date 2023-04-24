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

// AssignableOrganizationRoleType struct for AssignableOrganizationRoleType
type AssignableOrganizationRoleType struct {
	OrganizationRoleADMIN *OrganizationRoleADMIN
	OrganizationRoleANALYST *OrganizationRoleANALYST
	OrganizationRoleDEVELOPER *OrganizationRoleDEVELOPER
	OrganizationRoleOPERATIONSENGINEER *OrganizationRoleOPERATIONSENGINEER
	OrganizationRoleTASKUSER *OrganizationRoleTASKUSER
	OrganizationRoleVISITOR *OrganizationRoleVISITOR
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AssignableOrganizationRoleType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into OrganizationRoleADMIN
	err = json.Unmarshal(data, &dst.OrganizationRoleADMIN);
	if err == nil {
		jsonOrganizationRoleADMIN, _ := json.Marshal(dst.OrganizationRoleADMIN)
		if string(jsonOrganizationRoleADMIN) == "{}" { // empty struct
			dst.OrganizationRoleADMIN = nil
		} else {
			return nil // data stored in dst.OrganizationRoleADMIN, return on the first match
		}
	} else {
		dst.OrganizationRoleADMIN = nil
	}

	// try to unmarshal JSON data into OrganizationRoleANALYST
	err = json.Unmarshal(data, &dst.OrganizationRoleANALYST);
	if err == nil {
		jsonOrganizationRoleANALYST, _ := json.Marshal(dst.OrganizationRoleANALYST)
		if string(jsonOrganizationRoleANALYST) == "{}" { // empty struct
			dst.OrganizationRoleANALYST = nil
		} else {
			return nil // data stored in dst.OrganizationRoleANALYST, return on the first match
		}
	} else {
		dst.OrganizationRoleANALYST = nil
	}

	// try to unmarshal JSON data into OrganizationRoleDEVELOPER
	err = json.Unmarshal(data, &dst.OrganizationRoleDEVELOPER);
	if err == nil {
		jsonOrganizationRoleDEVELOPER, _ := json.Marshal(dst.OrganizationRoleDEVELOPER)
		if string(jsonOrganizationRoleDEVELOPER) == "{}" { // empty struct
			dst.OrganizationRoleDEVELOPER = nil
		} else {
			return nil // data stored in dst.OrganizationRoleDEVELOPER, return on the first match
		}
	} else {
		dst.OrganizationRoleDEVELOPER = nil
	}

	// try to unmarshal JSON data into OrganizationRoleOPERATIONSENGINEER
	err = json.Unmarshal(data, &dst.OrganizationRoleOPERATIONSENGINEER);
	if err == nil {
		jsonOrganizationRoleOPERATIONSENGINEER, _ := json.Marshal(dst.OrganizationRoleOPERATIONSENGINEER)
		if string(jsonOrganizationRoleOPERATIONSENGINEER) == "{}" { // empty struct
			dst.OrganizationRoleOPERATIONSENGINEER = nil
		} else {
			return nil // data stored in dst.OrganizationRoleOPERATIONSENGINEER, return on the first match
		}
	} else {
		dst.OrganizationRoleOPERATIONSENGINEER = nil
	}

	// try to unmarshal JSON data into OrganizationRoleTASKUSER
	err = json.Unmarshal(data, &dst.OrganizationRoleTASKUSER);
	if err == nil {
		jsonOrganizationRoleTASKUSER, _ := json.Marshal(dst.OrganizationRoleTASKUSER)
		if string(jsonOrganizationRoleTASKUSER) == "{}" { // empty struct
			dst.OrganizationRoleTASKUSER = nil
		} else {
			return nil // data stored in dst.OrganizationRoleTASKUSER, return on the first match
		}
	} else {
		dst.OrganizationRoleTASKUSER = nil
	}

	// try to unmarshal JSON data into OrganizationRoleVISITOR
	err = json.Unmarshal(data, &dst.OrganizationRoleVISITOR);
	if err == nil {
		jsonOrganizationRoleVISITOR, _ := json.Marshal(dst.OrganizationRoleVISITOR)
		if string(jsonOrganizationRoleVISITOR) == "{}" { // empty struct
			dst.OrganizationRoleVISITOR = nil
		} else {
			return nil // data stored in dst.OrganizationRoleVISITOR, return on the first match
		}
	} else {
		dst.OrganizationRoleVISITOR = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AssignableOrganizationRoleType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AssignableOrganizationRoleType) MarshalJSON() ([]byte, error) {
	if src.OrganizationRoleADMIN != nil {
		return json.Marshal(&src.OrganizationRoleADMIN)
	}

	if src.OrganizationRoleANALYST != nil {
		return json.Marshal(&src.OrganizationRoleANALYST)
	}

	if src.OrganizationRoleDEVELOPER != nil {
		return json.Marshal(&src.OrganizationRoleDEVELOPER)
	}

	if src.OrganizationRoleOPERATIONSENGINEER != nil {
		return json.Marshal(&src.OrganizationRoleOPERATIONSENGINEER)
	}

	if src.OrganizationRoleTASKUSER != nil {
		return json.Marshal(&src.OrganizationRoleTASKUSER)
	}

	if src.OrganizationRoleVISITOR != nil {
		return json.Marshal(&src.OrganizationRoleVISITOR)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAssignableOrganizationRoleType struct {
	value *AssignableOrganizationRoleType
	isSet bool
}

func (v NullableAssignableOrganizationRoleType) Get() *AssignableOrganizationRoleType {
	return v.value
}

func (v *NullableAssignableOrganizationRoleType) Set(val *AssignableOrganizationRoleType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignableOrganizationRoleType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignableOrganizationRoleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignableOrganizationRoleType(val *AssignableOrganizationRoleType) *NullableAssignableOrganizationRoleType {
	return &NullableAssignableOrganizationRoleType{value: val, isSet: true}
}

func (v NullableAssignableOrganizationRoleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignableOrganizationRoleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



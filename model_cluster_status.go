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

// checks if the ClusterStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterStatus{}

// ClusterStatus A health indicator for your Camunda cluster. Each of the components have their own state. The combined state is in the field `ready`.
type ClusterStatus struct {
	OptimizeStatus *ClusterHealthStatus `json:"optimizeStatus,omitempty"`
	TasklistStatus *ClusterHealthStatus `json:"tasklistStatus,omitempty"`
	OperateStatus *ClusterHealthStatus `json:"operateStatus,omitempty"`
	ZeebeStatus *ClusterHealthStatus `json:"zeebeStatus,omitempty"`
	Ready ClusterHealthStatus `json:"ready"`
}

// NewClusterStatus instantiates a new ClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatus(ready ClusterHealthStatus) *ClusterStatus {
	this := ClusterStatus{}
	this.Ready = ready
	return &this
}

// NewClusterStatusWithDefaults instantiates a new ClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatusWithDefaults() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// GetOptimizeStatus returns the OptimizeStatus field value if set, zero value otherwise.
func (o *ClusterStatus) GetOptimizeStatus() ClusterHealthStatus {
	if o == nil || IsNil(o.OptimizeStatus) {
		var ret ClusterHealthStatus
		return ret
	}
	return *o.OptimizeStatus
}

// GetOptimizeStatusOk returns a tuple with the OptimizeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetOptimizeStatusOk() (*ClusterHealthStatus, bool) {
	if o == nil || IsNil(o.OptimizeStatus) {
		return nil, false
	}
	return o.OptimizeStatus, true
}

// HasOptimizeStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasOptimizeStatus() bool {
	if o != nil && !IsNil(o.OptimizeStatus) {
		return true
	}

	return false
}

// SetOptimizeStatus gets a reference to the given ClusterHealthStatus and assigns it to the OptimizeStatus field.
func (o *ClusterStatus) SetOptimizeStatus(v ClusterHealthStatus) {
	o.OptimizeStatus = &v
}

// GetTasklistStatus returns the TasklistStatus field value if set, zero value otherwise.
func (o *ClusterStatus) GetTasklistStatus() ClusterHealthStatus {
	if o == nil || IsNil(o.TasklistStatus) {
		var ret ClusterHealthStatus
		return ret
	}
	return *o.TasklistStatus
}

// GetTasklistStatusOk returns a tuple with the TasklistStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetTasklistStatusOk() (*ClusterHealthStatus, bool) {
	if o == nil || IsNil(o.TasklistStatus) {
		return nil, false
	}
	return o.TasklistStatus, true
}

// HasTasklistStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasTasklistStatus() bool {
	if o != nil && !IsNil(o.TasklistStatus) {
		return true
	}

	return false
}

// SetTasklistStatus gets a reference to the given ClusterHealthStatus and assigns it to the TasklistStatus field.
func (o *ClusterStatus) SetTasklistStatus(v ClusterHealthStatus) {
	o.TasklistStatus = &v
}

// GetOperateStatus returns the OperateStatus field value if set, zero value otherwise.
func (o *ClusterStatus) GetOperateStatus() ClusterHealthStatus {
	if o == nil || IsNil(o.OperateStatus) {
		var ret ClusterHealthStatus
		return ret
	}
	return *o.OperateStatus
}

// GetOperateStatusOk returns a tuple with the OperateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetOperateStatusOk() (*ClusterHealthStatus, bool) {
	if o == nil || IsNil(o.OperateStatus) {
		return nil, false
	}
	return o.OperateStatus, true
}

// HasOperateStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasOperateStatus() bool {
	if o != nil && !IsNil(o.OperateStatus) {
		return true
	}

	return false
}

// SetOperateStatus gets a reference to the given ClusterHealthStatus and assigns it to the OperateStatus field.
func (o *ClusterStatus) SetOperateStatus(v ClusterHealthStatus) {
	o.OperateStatus = &v
}

// GetZeebeStatus returns the ZeebeStatus field value if set, zero value otherwise.
func (o *ClusterStatus) GetZeebeStatus() ClusterHealthStatus {
	if o == nil || IsNil(o.ZeebeStatus) {
		var ret ClusterHealthStatus
		return ret
	}
	return *o.ZeebeStatus
}

// GetZeebeStatusOk returns a tuple with the ZeebeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetZeebeStatusOk() (*ClusterHealthStatus, bool) {
	if o == nil || IsNil(o.ZeebeStatus) {
		return nil, false
	}
	return o.ZeebeStatus, true
}

// HasZeebeStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasZeebeStatus() bool {
	if o != nil && !IsNil(o.ZeebeStatus) {
		return true
	}

	return false
}

// SetZeebeStatus gets a reference to the given ClusterHealthStatus and assigns it to the ZeebeStatus field.
func (o *ClusterStatus) SetZeebeStatus(v ClusterHealthStatus) {
	o.ZeebeStatus = &v
}

// GetReady returns the Ready field value
func (o *ClusterStatus) GetReady() ClusterHealthStatus {
	if o == nil {
		var ret ClusterHealthStatus
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetReadyOk() (*ClusterHealthStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *ClusterStatus) SetReady(v ClusterHealthStatus) {
	o.Ready = v
}

func (o ClusterStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptimizeStatus) {
		toSerialize["optimizeStatus"] = o.OptimizeStatus
	}
	if !IsNil(o.TasklistStatus) {
		toSerialize["tasklistStatus"] = o.TasklistStatus
	}
	if !IsNil(o.OperateStatus) {
		toSerialize["operateStatus"] = o.OperateStatus
	}
	if !IsNil(o.ZeebeStatus) {
		toSerialize["zeebeStatus"] = o.ZeebeStatus
	}
	toSerialize["ready"] = o.Ready
	return toSerialize, nil
}

type NullableClusterStatus struct {
	value *ClusterStatus
	isSet bool
}

func (v NullableClusterStatus) Get() *ClusterStatus {
	return v.value
}

func (v *NullableClusterStatus) Set(val *ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatus(val *ClusterStatus) *NullableClusterStatus {
	return &NullableClusterStatus{value: val, isSet: true}
}

func (v NullableClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



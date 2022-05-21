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

// ClusterStatus2 A health indicator for your Camunda Cloud cluster. Each of the components have their own state. The combined state is in the field `ready`.
type ClusterStatus2 struct {
	OptimizeStatus *ClusterStatus `json:"optimizeStatus,omitempty"`
	TasklistStatus *ClusterStatus `json:"tasklistStatus,omitempty"`
	OperateStatus *ClusterStatus `json:"operateStatus,omitempty"`
	ZeebeStatus *ClusterStatus `json:"zeebeStatus,omitempty"`
	Ready ClusterStatus `json:"ready"`
}

// NewClusterStatus2 instantiates a new ClusterStatus2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatus2(ready ClusterStatus) *ClusterStatus2 {
	this := ClusterStatus2{}
	this.Ready = ready
	return &this
}

// NewClusterStatus2WithDefaults instantiates a new ClusterStatus2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatus2WithDefaults() *ClusterStatus2 {
	this := ClusterStatus2{}
	return &this
}

// GetOptimizeStatus returns the OptimizeStatus field value if set, zero value otherwise.
func (o *ClusterStatus2) GetOptimizeStatus() ClusterStatus {
	if o == nil || o.OptimizeStatus == nil {
		var ret ClusterStatus
		return ret
	}
	return *o.OptimizeStatus
}

// GetOptimizeStatusOk returns a tuple with the OptimizeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus2) GetOptimizeStatusOk() (*ClusterStatus, bool) {
	if o == nil || o.OptimizeStatus == nil {
		return nil, false
	}
	return o.OptimizeStatus, true
}

// HasOptimizeStatus returns a boolean if a field has been set.
func (o *ClusterStatus2) HasOptimizeStatus() bool {
	if o != nil && o.OptimizeStatus != nil {
		return true
	}

	return false
}

// SetOptimizeStatus gets a reference to the given ClusterStatus and assigns it to the OptimizeStatus field.
func (o *ClusterStatus2) SetOptimizeStatus(v ClusterStatus) {
	o.OptimizeStatus = &v
}

// GetTasklistStatus returns the TasklistStatus field value if set, zero value otherwise.
func (o *ClusterStatus2) GetTasklistStatus() ClusterStatus {
	if o == nil || o.TasklistStatus == nil {
		var ret ClusterStatus
		return ret
	}
	return *o.TasklistStatus
}

// GetTasklistStatusOk returns a tuple with the TasklistStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus2) GetTasklistStatusOk() (*ClusterStatus, bool) {
	if o == nil || o.TasklistStatus == nil {
		return nil, false
	}
	return o.TasklistStatus, true
}

// HasTasklistStatus returns a boolean if a field has been set.
func (o *ClusterStatus2) HasTasklistStatus() bool {
	if o != nil && o.TasklistStatus != nil {
		return true
	}

	return false
}

// SetTasklistStatus gets a reference to the given ClusterStatus and assigns it to the TasklistStatus field.
func (o *ClusterStatus2) SetTasklistStatus(v ClusterStatus) {
	o.TasklistStatus = &v
}

// GetOperateStatus returns the OperateStatus field value if set, zero value otherwise.
func (o *ClusterStatus2) GetOperateStatus() ClusterStatus {
	if o == nil || o.OperateStatus == nil {
		var ret ClusterStatus
		return ret
	}
	return *o.OperateStatus
}

// GetOperateStatusOk returns a tuple with the OperateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus2) GetOperateStatusOk() (*ClusterStatus, bool) {
	if o == nil || o.OperateStatus == nil {
		return nil, false
	}
	return o.OperateStatus, true
}

// HasOperateStatus returns a boolean if a field has been set.
func (o *ClusterStatus2) HasOperateStatus() bool {
	if o != nil && o.OperateStatus != nil {
		return true
	}

	return false
}

// SetOperateStatus gets a reference to the given ClusterStatus and assigns it to the OperateStatus field.
func (o *ClusterStatus2) SetOperateStatus(v ClusterStatus) {
	o.OperateStatus = &v
}

// GetZeebeStatus returns the ZeebeStatus field value if set, zero value otherwise.
func (o *ClusterStatus2) GetZeebeStatus() ClusterStatus {
	if o == nil || o.ZeebeStatus == nil {
		var ret ClusterStatus
		return ret
	}
	return *o.ZeebeStatus
}

// GetZeebeStatusOk returns a tuple with the ZeebeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus2) GetZeebeStatusOk() (*ClusterStatus, bool) {
	if o == nil || o.ZeebeStatus == nil {
		return nil, false
	}
	return o.ZeebeStatus, true
}

// HasZeebeStatus returns a boolean if a field has been set.
func (o *ClusterStatus2) HasZeebeStatus() bool {
	if o != nil && o.ZeebeStatus != nil {
		return true
	}

	return false
}

// SetZeebeStatus gets a reference to the given ClusterStatus and assigns it to the ZeebeStatus field.
func (o *ClusterStatus2) SetZeebeStatus(v ClusterStatus) {
	o.ZeebeStatus = &v
}

// GetReady returns the Ready field value
func (o *ClusterStatus2) GetReady() ClusterStatus {
	if o == nil {
		var ret ClusterStatus
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *ClusterStatus2) GetReadyOk() (*ClusterStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *ClusterStatus2) SetReady(v ClusterStatus) {
	o.Ready = v
}

func (o ClusterStatus2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OptimizeStatus != nil {
		toSerialize["optimizeStatus"] = o.OptimizeStatus
	}
	if o.TasklistStatus != nil {
		toSerialize["tasklistStatus"] = o.TasklistStatus
	}
	if o.OperateStatus != nil {
		toSerialize["operateStatus"] = o.OperateStatus
	}
	if o.ZeebeStatus != nil {
		toSerialize["zeebeStatus"] = o.ZeebeStatus
	}
	if true {
		toSerialize["ready"] = o.Ready
	}
	return json.Marshal(toSerialize)
}

type NullableClusterStatus2 struct {
	value *ClusterStatus2
	isSet bool
}

func (v NullableClusterStatus2) Get() *ClusterStatus2 {
	return v.value
}

func (v *NullableClusterStatus2) Set(val *ClusterStatus2) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatus2) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatus2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatus2(val *ClusterStatus2) *NullableClusterStatus2 {
	return &NullableClusterStatus2{value: val, isSet: true}
}

func (v NullableClusterStatus2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatus2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



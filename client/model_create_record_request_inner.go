/*
ncloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateRecordRequestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecordRequestInner{}

// CreateRecordRequestInner struct for CreateRecordRequestInner
type CreateRecordRequestInner struct {
	Host string `json:"host"`
	Type string `json:"type"`
	Content string `json:"content"`
	Ttl int64 `json:"ttl"`
	AliasId *int64 `json:"aliasId,omitempty"`
	LbId *int64 `json:"lbId,omitempty"`
}

type _CreateRecordRequestInner CreateRecordRequestInner

// NewCreateRecordRequestInner instantiates a new CreateRecordRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecordRequestInner(host string, type_ string, content string, ttl int64) *CreateRecordRequestInner {
	this := CreateRecordRequestInner{}
	this.Host = host
	this.Type = type_
	this.Content = content
	this.Ttl = ttl
	return &this
}

// NewCreateRecordRequestInnerWithDefaults instantiates a new CreateRecordRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecordRequestInnerWithDefaults() *CreateRecordRequestInner {
	this := CreateRecordRequestInner{}
	return &this
}

// GetHost returns the Host field value
func (o *CreateRecordRequestInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateRecordRequestInner) SetHost(v string) {
	o.Host = v
}

// GetType returns the Type field value
func (o *CreateRecordRequestInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateRecordRequestInner) SetType(v string) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *CreateRecordRequestInner) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestInner) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateRecordRequestInner) SetContent(v string) {
	o.Content = v
}

// GetTtl returns the Ttl field value
func (o *CreateRecordRequestInner) GetTtl() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestInner) GetTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *CreateRecordRequestInner) SetTtl(v int64) {
	o.Ttl = v
}

// GetAliasId returns the AliasId field value if set, zero value otherwise.
func (o *CreateRecordRequestInner) GetAliasId() int64 {
	if o == nil || IsNil(o.AliasId) {
		var ret int64
		return ret
	}
	return *o.AliasId
}

// GetAliasIdOk returns a tuple with the AliasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestInner) GetAliasIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AliasId) {
		return nil, false
	}
	return o.AliasId, true
}

// HasAliasId returns a boolean if a field has been set.
func (o *CreateRecordRequestInner) HasAliasId() bool {
	if o != nil && !IsNil(o.AliasId) {
		return true
	}

	return false
}

// SetAliasId gets a reference to the given int64 and assigns it to the AliasId field.
func (o *CreateRecordRequestInner) SetAliasId(v int64) {
	o.AliasId = &v
}

// GetLbId returns the LbId field value if set, zero value otherwise.
func (o *CreateRecordRequestInner) GetLbId() int64 {
	if o == nil || IsNil(o.LbId) {
		var ret int64
		return ret
	}
	return *o.LbId
}

// GetLbIdOk returns a tuple with the LbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestInner) GetLbIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LbId) {
		return nil, false
	}
	return o.LbId, true
}

// HasLbId returns a boolean if a field has been set.
func (o *CreateRecordRequestInner) HasLbId() bool {
	if o != nil && !IsNil(o.LbId) {
		return true
	}

	return false
}

// SetLbId gets a reference to the given int64 and assigns it to the LbId field.
func (o *CreateRecordRequestInner) SetLbId(v int64) {
	o.LbId = &v
}

func (o CreateRecordRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecordRequestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	toSerialize["ttl"] = o.Ttl
	if !IsNil(o.AliasId) {
		toSerialize["aliasId"] = o.AliasId
	}
	if !IsNil(o.LbId) {
		toSerialize["lbId"] = o.LbId
	}
	return toSerialize, nil
}

func (o *CreateRecordRequestInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"type",
		"content",
		"ttl",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateRecordRequestInner := _CreateRecordRequestInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRecordRequestInner)

	if err != nil {
		return err
	}

	*o = CreateRecordRequestInner(varCreateRecordRequestInner)

	return err
}

type NullableCreateRecordRequestInner struct {
	value *CreateRecordRequestInner
	isSet bool
}

func (v NullableCreateRecordRequestInner) Get() *CreateRecordRequestInner {
	return v.value
}

func (v *NullableCreateRecordRequestInner) Set(val *CreateRecordRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordRequestInner(val *CreateRecordRequestInner) *NullableCreateRecordRequestInner {
	return &NullableCreateRecordRequestInner{value: val, isSet: true}
}

func (v NullableCreateRecordRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



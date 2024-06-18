/*
Studio - AI Empower Labs

# Studio API Documentation  ## Introduction Welcome to Studio by AI Empower Labs API documentation! We are thrilled to offer developers around the world access to our cutting-edge artificial intelligence technology and semantic search. Our API is designed to empower your applications with state-of-the-art AI capabilities, including but not limited to natural language processing, audio transcription, embedding, and predictive analytics.  Our mission is to make AI technology accessible and easy to integrate, enabling you to enhance your applications, improve user experiences, and innovate in your field. Whether you're building complex systems, developing mobile apps, or creating web services, our API provides you with the tools you need to incorporate AI functionalities seamlessly.  Support and Feedback We are committed to providing exceptional support to our developer community. If you have any questions, encounter any issues, or have feedback on how we can improve our API, please don't hesitate to contact our support team @ support@AIEmpowerLabs.com.  Terms of Use Please review our terms of use and privacy policy before integrating our API into your application. By using our API, you agree to comply with these terms.

API version: v1
Contact: support@aiempowerlabs.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DocumentFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentFilters{}

// DocumentFilters struct for DocumentFilters
type DocumentFilters struct {
	DocumentId []string `json:"documentId,omitempty"`
	Tags map[string][]string `json:"tags,omitempty"`
}

// NewDocumentFilters instantiates a new DocumentFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentFilters() *DocumentFilters {
	this := DocumentFilters{}
	return &this
}

// NewDocumentFiltersWithDefaults instantiates a new DocumentFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentFiltersWithDefaults() *DocumentFilters {
	this := DocumentFilters{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentFilters) GetDocumentId() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentFilters) GetDocumentIdOk() ([]string, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *DocumentFilters) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given []string and assigns it to the DocumentId field.
func (o *DocumentFilters) SetDocumentId(v []string) {
	o.DocumentId = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentFilters) GetTags() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentFilters) GetTagsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DocumentFilters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string][]string and assigns it to the Tags field.
func (o *DocumentFilters) SetTags(v map[string][]string) {
	o.Tags = v
}

func (o DocumentFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableDocumentFilters struct {
	value *DocumentFilters
	isSet bool
}

func (v NullableDocumentFilters) Get() *DocumentFilters {
	return v.value
}

func (v *NullableDocumentFilters) Set(val *DocumentFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentFilters(val *DocumentFilters) *NullableDocumentFilters {
	return &NullableDocumentFilters{value: val, isSet: true}
}

func (v NullableDocumentFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



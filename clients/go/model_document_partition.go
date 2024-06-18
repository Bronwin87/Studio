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
	"time"
)

// checks if the DocumentPartition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentPartition{}

// DocumentPartition struct for DocumentPartition
type DocumentPartition struct {
	// Content of the document partition, aka chunk/block of text.
	Text NullableString `json:"text,omitempty"`
	// Relevance of this partition against the given query.  Value usually is between 0 and 1, when using cosine similarity.
	Relevance *float32 `json:"relevance,omitempty"`
	// Timestamp about the file/text partition.
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`
	// List of document tags
	Tags map[string][]string `json:"tags,omitempty"`
}

// NewDocumentPartition instantiates a new DocumentPartition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentPartition() *DocumentPartition {
	this := DocumentPartition{}
	return &this
}

// NewDocumentPartitionWithDefaults instantiates a new DocumentPartition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentPartitionWithDefaults() *DocumentPartition {
	this := DocumentPartition{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentPartition) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentPartition) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *DocumentPartition) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *DocumentPartition) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *DocumentPartition) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *DocumentPartition) UnsetText() {
	o.Text.Unset()
}

// GetRelevance returns the Relevance field value if set, zero value otherwise.
func (o *DocumentPartition) GetRelevance() float32 {
	if o == nil || IsNil(o.Relevance) {
		var ret float32
		return ret
	}
	return *o.Relevance
}

// GetRelevanceOk returns a tuple with the Relevance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPartition) GetRelevanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Relevance) {
		return nil, false
	}
	return o.Relevance, true
}

// HasRelevance returns a boolean if a field has been set.
func (o *DocumentPartition) HasRelevance() bool {
	if o != nil && !IsNil(o.Relevance) {
		return true
	}

	return false
}

// SetRelevance gets a reference to the given float32 and assigns it to the Relevance field.
func (o *DocumentPartition) SetRelevance(v float32) {
	o.Relevance = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *DocumentPartition) GetLastUpdate() time.Time {
	if o == nil || IsNil(o.LastUpdate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPartition) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *DocumentPartition) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *DocumentPartition) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentPartition) GetTags() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentPartition) GetTagsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DocumentPartition) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string][]string and assigns it to the Tags field.
func (o *DocumentPartition) SetTags(v map[string][]string) {
	o.Tags = v
}

func (o DocumentPartition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentPartition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if !IsNil(o.Relevance) {
		toSerialize["relevance"] = o.Relevance
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableDocumentPartition struct {
	value *DocumentPartition
	isSet bool
}

func (v NullableDocumentPartition) Get() *DocumentPartition {
	return v.value
}

func (v *NullableDocumentPartition) Set(val *DocumentPartition) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentPartition) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentPartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentPartition(val *DocumentPartition) *NullableDocumentPartition {
	return &NullableDocumentPartition{value: val, isSet: true}
}

func (v NullableDocumentPartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentPartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the EmbeddingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddingRequest{}

// EmbeddingRequest struct for EmbeddingRequest
type EmbeddingRequest struct {
	Model NullableString `json:"model,omitempty"`
	Text []string `json:"text,omitempty"`
}

// NewEmbeddingRequest instantiates a new EmbeddingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddingRequest() *EmbeddingRequest {
	this := EmbeddingRequest{}
	return &this
}

// NewEmbeddingRequestWithDefaults instantiates a new EmbeddingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddingRequestWithDefaults() *EmbeddingRequest {
	this := EmbeddingRequest{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmbeddingRequest) GetModel() string {
	if o == nil || IsNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddingRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *EmbeddingRequest) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *EmbeddingRequest) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *EmbeddingRequest) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *EmbeddingRequest) UnsetModel() {
	o.Model.Unset()
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *EmbeddingRequest) GetText() []string {
	if o == nil || IsNil(o.Text) {
		var ret []string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddingRequest) GetTextOk() ([]string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *EmbeddingRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given []string and assigns it to the Text field.
func (o *EmbeddingRequest) SetText(v []string) {
	o.Text = v
}

func (o EmbeddingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableEmbeddingRequest struct {
	value *EmbeddingRequest
	isSet bool
}

func (v NullableEmbeddingRequest) Get() *EmbeddingRequest {
	return v.value
}

func (v *NullableEmbeddingRequest) Set(val *EmbeddingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddingRequest(val *EmbeddingRequest) *NullableEmbeddingRequest {
	return &NullableEmbeddingRequest{value: val, isSet: true}
}

func (v NullableEmbeddingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



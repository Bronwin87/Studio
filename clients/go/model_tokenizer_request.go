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

// checks if the TokenizerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizerRequest{}

// TokenizerRequest struct for TokenizerRequest
type TokenizerRequest struct {
	Name NullableString `json:"name,omitempty"`
	Text []string `json:"text,omitempty"`
}

// NewTokenizerRequest instantiates a new TokenizerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizerRequest() *TokenizerRequest {
	this := TokenizerRequest{}
	return &this
}

// NewTokenizerRequestWithDefaults instantiates a new TokenizerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizerRequestWithDefaults() *TokenizerRequest {
	this := TokenizerRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenizerRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenizerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TokenizerRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TokenizerRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TokenizerRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TokenizerRequest) UnsetName() {
	o.Name.Unset()
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TokenizerRequest) GetText() []string {
	if o == nil || IsNil(o.Text) {
		var ret []string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizerRequest) GetTextOk() ([]string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TokenizerRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given []string and assigns it to the Text field.
func (o *TokenizerRequest) SetText(v []string) {
	o.Text = v
}

func (o TokenizerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableTokenizerRequest struct {
	value *TokenizerRequest
	isSet bool
}

func (v NullableTokenizerRequest) Get() *TokenizerRequest {
	return v.value
}

func (v *NullableTokenizerRequest) Set(val *TokenizerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizerRequest(val *TokenizerRequest) *NullableTokenizerRequest {
	return &NullableTokenizerRequest{value: val, isSet: true}
}

func (v NullableTokenizerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



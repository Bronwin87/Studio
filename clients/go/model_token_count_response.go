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

// checks if the TokenCountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenCountResponse{}

// TokenCountResponse struct for TokenCountResponse
type TokenCountResponse struct {
	Tokens *int32 `json:"tokens,omitempty"`
}

// NewTokenCountResponse instantiates a new TokenCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCountResponse() *TokenCountResponse {
	this := TokenCountResponse{}
	return &this
}

// NewTokenCountResponseWithDefaults instantiates a new TokenCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCountResponseWithDefaults() *TokenCountResponse {
	this := TokenCountResponse{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *TokenCountResponse) GetTokens() int32 {
	if o == nil || IsNil(o.Tokens) {
		var ret int32
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCountResponse) GetTokensOk() (*int32, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *TokenCountResponse) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given int32 and assigns it to the Tokens field.
func (o *TokenCountResponse) SetTokens(v int32) {
	o.Tokens = &v
}

func (o TokenCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenCountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	return toSerialize, nil
}

type NullableTokenCountResponse struct {
	value *TokenCountResponse
	isSet bool
}

func (v NullableTokenCountResponse) Get() *TokenCountResponse {
	return v.value
}

func (v *NullableTokenCountResponse) Set(val *TokenCountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCountResponse(val *TokenCountResponse) *NullableTokenCountResponse {
	return &NullableTokenCountResponse{value: val, isSet: true}
}

func (v NullableTokenCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



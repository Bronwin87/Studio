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

// checks if the DetectLanguageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectLanguageResponse{}

// DetectLanguageResponse struct for DetectLanguageResponse
type DetectLanguageResponse struct {
	Languages []LanguageDetection `json:"languages,omitempty"`
}

// NewDetectLanguageResponse instantiates a new DetectLanguageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectLanguageResponse() *DetectLanguageResponse {
	this := DetectLanguageResponse{}
	return &this
}

// NewDetectLanguageResponseWithDefaults instantiates a new DetectLanguageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectLanguageResponseWithDefaults() *DetectLanguageResponse {
	this := DetectLanguageResponse{}
	return &this
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetectLanguageResponse) GetLanguages() []LanguageDetection {
	if o == nil {
		var ret []LanguageDetection
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectLanguageResponse) GetLanguagesOk() ([]LanguageDetection, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *DetectLanguageResponse) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []LanguageDetection and assigns it to the Languages field.
func (o *DetectLanguageResponse) SetLanguages(v []LanguageDetection) {
	o.Languages = v
}

func (o DetectLanguageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectLanguageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	return toSerialize, nil
}

type NullableDetectLanguageResponse struct {
	value *DetectLanguageResponse
	isSet bool
}

func (v NullableDetectLanguageResponse) Get() *DetectLanguageResponse {
	return v.value
}

func (v *NullableDetectLanguageResponse) Set(val *DetectLanguageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectLanguageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectLanguageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectLanguageResponse(val *DetectLanguageResponse) *NullableDetectLanguageResponse {
	return &NullableDetectLanguageResponse{value: val, isSet: true}
}

func (v NullableDetectLanguageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectLanguageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



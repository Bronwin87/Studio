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

// checks if the SemanticSimilarityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SemanticSimilarityResponse{}

// SemanticSimilarityResponse struct for SemanticSimilarityResponse
type SemanticSimilarityResponse struct {
	Documents []SemanticSimilarityDocument `json:"documents,omitempty"`
}

// NewSemanticSimilarityResponse instantiates a new SemanticSimilarityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSemanticSimilarityResponse() *SemanticSimilarityResponse {
	this := SemanticSimilarityResponse{}
	return &this
}

// NewSemanticSimilarityResponseWithDefaults instantiates a new SemanticSimilarityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSemanticSimilarityResponseWithDefaults() *SemanticSimilarityResponse {
	this := SemanticSimilarityResponse{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SemanticSimilarityResponse) GetDocuments() []SemanticSimilarityDocument {
	if o == nil {
		var ret []SemanticSimilarityDocument
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SemanticSimilarityResponse) GetDocumentsOk() ([]SemanticSimilarityDocument, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *SemanticSimilarityResponse) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []SemanticSimilarityDocument and assigns it to the Documents field.
func (o *SemanticSimilarityResponse) SetDocuments(v []SemanticSimilarityDocument) {
	o.Documents = v
}

func (o SemanticSimilarityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SemanticSimilarityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	return toSerialize, nil
}

type NullableSemanticSimilarityResponse struct {
	value *SemanticSimilarityResponse
	isSet bool
}

func (v NullableSemanticSimilarityResponse) Get() *SemanticSimilarityResponse {
	return v.value
}

func (v *NullableSemanticSimilarityResponse) Set(val *SemanticSimilarityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSemanticSimilarityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSemanticSimilarityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSemanticSimilarityResponse(val *SemanticSimilarityResponse) *NullableSemanticSimilarityResponse {
	return &NullableSemanticSimilarityResponse{value: val, isSet: true}
}

func (v NullableSemanticSimilarityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSemanticSimilarityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



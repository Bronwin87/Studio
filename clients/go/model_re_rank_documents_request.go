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

// checks if the ReRankDocumentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReRankDocumentsRequest{}

// ReRankDocumentsRequest struct for ReRankDocumentsRequest
type ReRankDocumentsRequest struct {
	// Optional embedding model. Defaults to configured default
	Model NullableString `json:"model,omitempty"`
	// Semantic query to find matching documents
	Query NullableString `json:"query,omitempty"`
	// Optional index to specify which index to search in. Defaults to 'default'
	Documents []string `json:"documents,omitempty"`
	// Optional filter for specifying maximum number of entries to return. Defaults to 3
	Limit NullableInt32 `json:"limit,omitempty"`
}

// NewReRankDocumentsRequest instantiates a new ReRankDocumentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReRankDocumentsRequest() *ReRankDocumentsRequest {
	this := ReRankDocumentsRequest{}
	return &this
}

// NewReRankDocumentsRequestWithDefaults instantiates a new ReRankDocumentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReRankDocumentsRequestWithDefaults() *ReRankDocumentsRequest {
	this := ReRankDocumentsRequest{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReRankDocumentsRequest) GetModel() string {
	if o == nil || IsNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReRankDocumentsRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *ReRankDocumentsRequest) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *ReRankDocumentsRequest) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *ReRankDocumentsRequest) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *ReRankDocumentsRequest) UnsetModel() {
	o.Model.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReRankDocumentsRequest) GetQuery() string {
	if o == nil || IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReRankDocumentsRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *ReRankDocumentsRequest) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *ReRankDocumentsRequest) SetQuery(v string) {
	o.Query.Set(&v)
}
// SetQueryNil sets the value for Query to be an explicit nil
func (o *ReRankDocumentsRequest) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *ReRankDocumentsRequest) UnsetQuery() {
	o.Query.Unset()
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReRankDocumentsRequest) GetDocuments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReRankDocumentsRequest) GetDocumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *ReRankDocumentsRequest) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []string and assigns it to the Documents field.
func (o *ReRankDocumentsRequest) SetDocuments(v []string) {
	o.Documents = v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReRankDocumentsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReRankDocumentsRequest) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *ReRankDocumentsRequest) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *ReRankDocumentsRequest) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *ReRankDocumentsRequest) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *ReRankDocumentsRequest) UnsetLimit() {
	o.Limit.Unset()
}

func (o ReRankDocumentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReRankDocumentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	return toSerialize, nil
}

type NullableReRankDocumentsRequest struct {
	value *ReRankDocumentsRequest
	isSet bool
}

func (v NullableReRankDocumentsRequest) Get() *ReRankDocumentsRequest {
	return v.value
}

func (v *NullableReRankDocumentsRequest) Set(val *ReRankDocumentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReRankDocumentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReRankDocumentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReRankDocumentsRequest(val *ReRankDocumentsRequest) *NullableReRankDocumentsRequest {
	return &NullableReRankDocumentsRequest{value: val, isSet: true}
}

func (v NullableReRankDocumentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReRankDocumentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



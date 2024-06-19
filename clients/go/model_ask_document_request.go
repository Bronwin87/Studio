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

// checks if the AskDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AskDocumentRequest{}

// AskDocumentRequest struct for AskDocumentRequest
type AskDocumentRequest struct {
	// Semantic query to find matching documents and ask questions over
	Query NullableString `json:"query,omitempty"`
	// Optional index to specify which index to search in. Defaults to 'default'
	Index NullableString `json:"index,omitempty"`
	// Optional filtering of document id(s) and/or tags
	Filter []DocumentFilters `json:"filter,omitempty"`
	// Optional filter to specify minimum relevance. Typically values between 0 and 1
	MinRelevance NullableFloat64 `json:"minRelevance,omitempty"`
	// Large language model to use in query
	LlmModel NullableString `json:"llmModel,omitempty"`
	// Embedding model to use in query
	EmbeddingModel NullableString `json:"embeddingModel,omitempty"`
	Args map[string]interface{} `json:"args,omitempty"`
}

// NewAskDocumentRequest instantiates a new AskDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAskDocumentRequest() *AskDocumentRequest {
	this := AskDocumentRequest{}
	return &this
}

// NewAskDocumentRequestWithDefaults instantiates a new AskDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAskDocumentRequestWithDefaults() *AskDocumentRequest {
	this := AskDocumentRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AskDocumentRequest) GetQuery() string {
	if o == nil || IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AskDocumentRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *AskDocumentRequest) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *AskDocumentRequest) SetQuery(v string) {
	o.Query.Set(&v)
}
// SetQueryNil sets the value for Query to be an explicit nil
func (o *AskDocumentRequest) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *AskDocumentRequest) UnsetQuery() {
	o.Query.Unset()
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AskDocumentRequest) GetIndex() string {
	if o == nil || IsNil(o.Index.Get()) {
		var ret string
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AskDocumentRequest) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *AskDocumentRequest) HasIndex() bool {
	if o != nil && o.Index.IsSet() {
		return true
	}

	return false
}

// SetIndex gets a reference to the given NullableString and assigns it to the Index field.
func (o *AskDocumentRequest) SetIndex(v string) {
	o.Index.Set(&v)
}
// SetIndexNil sets the value for Index to be an explicit nil
func (o *AskDocumentRequest) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil
func (o *AskDocumentRequest) UnsetIndex() {
	o.Index.Unset()
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AskDocumentRequest) GetFilter() []DocumentFilters {
	if o == nil {
		var ret []DocumentFilters
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AskDocumentRequest) GetFilterOk() ([]DocumentFilters, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AskDocumentRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []DocumentFilters and assigns it to the Filter field.
func (o *AskDocumentRequest) SetFilter(v []DocumentFilters) {
	o.Filter = v
}

// GetMinRelevance returns the MinRelevance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AskDocumentRequest) GetMinRelevance() float64 {
	if o == nil || IsNil(o.MinRelevance.Get()) {
		var ret float64
		return ret
	}
	return *o.MinRelevance.Get()
}

// GetMinRelevanceOk returns a tuple with the MinRelevance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AskDocumentRequest) GetMinRelevanceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinRelevance.Get(), o.MinRelevance.IsSet()
}

// HasMinRelevance returns a boolean if a field has been set.
func (o *AskDocumentRequest) HasMinRelevance() bool {
	if o != nil && o.MinRelevance.IsSet() {
		return true
	}

	return false
}

// SetMinRelevance gets a reference to the given NullableFloat64 and assigns it to the MinRelevance field.
func (o *AskDocumentRequest) SetMinRelevance(v float64) {
	o.MinRelevance.Set(&v)
}
// SetMinRelevanceNil sets the value for MinRelevance to be an explicit nil
func (o *AskDocumentRequest) SetMinRelevanceNil() {
	o.MinRelevance.Set(nil)
}

// UnsetMinRelevance ensures that no value is present for MinRelevance, not even an explicit nil
func (o *AskDocumentRequest) UnsetMinRelevance() {
	o.MinRelevance.Unset()
}

// GetLlmModel returns the LlmModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AskDocumentRequest) GetLlmModel() string {
	if o == nil || IsNil(o.LlmModel.Get()) {
		var ret string
		return ret
	}
	return *o.LlmModel.Get()
}

// GetLlmModelOk returns a tuple with the LlmModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AskDocumentRequest) GetLlmModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LlmModel.Get(), o.LlmModel.IsSet()
}

// HasLlmModel returns a boolean if a field has been set.
func (o *AskDocumentRequest) HasLlmModel() bool {
	if o != nil && o.LlmModel.IsSet() {
		return true
	}

	return false
}

// SetLlmModel gets a reference to the given NullableString and assigns it to the LlmModel field.
func (o *AskDocumentRequest) SetLlmModel(v string) {
	o.LlmModel.Set(&v)
}
// SetLlmModelNil sets the value for LlmModel to be an explicit nil
func (o *AskDocumentRequest) SetLlmModelNil() {
	o.LlmModel.Set(nil)
}

// UnsetLlmModel ensures that no value is present for LlmModel, not even an explicit nil
func (o *AskDocumentRequest) UnsetLlmModel() {
	o.LlmModel.Unset()
}

// GetEmbeddingModel returns the EmbeddingModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AskDocumentRequest) GetEmbeddingModel() string {
	if o == nil || IsNil(o.EmbeddingModel.Get()) {
		var ret string
		return ret
	}
	return *o.EmbeddingModel.Get()
}

// GetEmbeddingModelOk returns a tuple with the EmbeddingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AskDocumentRequest) GetEmbeddingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbeddingModel.Get(), o.EmbeddingModel.IsSet()
}

// HasEmbeddingModel returns a boolean if a field has been set.
func (o *AskDocumentRequest) HasEmbeddingModel() bool {
	if o != nil && o.EmbeddingModel.IsSet() {
		return true
	}

	return false
}

// SetEmbeddingModel gets a reference to the given NullableString and assigns it to the EmbeddingModel field.
func (o *AskDocumentRequest) SetEmbeddingModel(v string) {
	o.EmbeddingModel.Set(&v)
}
// SetEmbeddingModelNil sets the value for EmbeddingModel to be an explicit nil
func (o *AskDocumentRequest) SetEmbeddingModelNil() {
	o.EmbeddingModel.Set(nil)
}

// UnsetEmbeddingModel ensures that no value is present for EmbeddingModel, not even an explicit nil
func (o *AskDocumentRequest) UnsetEmbeddingModel() {
	o.EmbeddingModel.Unset()
}

// GetArgs returns the Args field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AskDocumentRequest) GetArgs() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AskDocumentRequest) GetArgsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return map[string]interface{}{}, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *AskDocumentRequest) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given map[string]interface{} and assigns it to the Args field.
func (o *AskDocumentRequest) SetArgs(v map[string]interface{}) {
	o.Args = v
}

func (o AskDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AskDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.Index.IsSet() {
		toSerialize["index"] = o.Index.Get()
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.MinRelevance.IsSet() {
		toSerialize["minRelevance"] = o.MinRelevance.Get()
	}
	if o.LlmModel.IsSet() {
		toSerialize["llmModel"] = o.LlmModel.Get()
	}
	if o.EmbeddingModel.IsSet() {
		toSerialize["embeddingModel"] = o.EmbeddingModel.Get()
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	return toSerialize, nil
}

type NullableAskDocumentRequest struct {
	value *AskDocumentRequest
	isSet bool
}

func (v NullableAskDocumentRequest) Get() *AskDocumentRequest {
	return v.value
}

func (v *NullableAskDocumentRequest) Set(val *AskDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAskDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAskDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAskDocumentRequest(val *AskDocumentRequest) *NullableAskDocumentRequest {
	return &NullableAskDocumentRequest{value: val, isSet: true}
}

func (v NullableAskDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAskDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



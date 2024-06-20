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

// checks if the ListDocumentParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDocumentParameters{}

// ListDocumentParameters struct for ListDocumentParameters
type ListDocumentParameters struct {
	// Optional index to specify which index to search in. Defaults to 'default'
	Index NullableString `json:"index,omitempty"`
	// Optional filtering of document id(s) and/or tags
	Filter []DocumentFilters `json:"filter,omitempty"`
	// Optionally specifies if embedding should be returned in response. Default is false
	WithEmbeddings NullableBool `json:"withEmbeddings,omitempty"`
	// Optional filter for specifying maximum number of entries to return. Defaults to 3
	Limit NullableInt32 `json:"limit,omitempty"`
	// Optional filter for specifying list offset for paging. Default is 0
	Offset NullableInt32 `json:"offset,omitempty"`
}

// NewListDocumentParameters instantiates a new ListDocumentParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDocumentParameters() *ListDocumentParameters {
	this := ListDocumentParameters{}
	return &this
}

// NewListDocumentParametersWithDefaults instantiates a new ListDocumentParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDocumentParametersWithDefaults() *ListDocumentParameters {
	this := ListDocumentParameters{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListDocumentParameters) GetIndex() string {
	if o == nil || IsNil(o.Index.Get()) {
		var ret string
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListDocumentParameters) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *ListDocumentParameters) HasIndex() bool {
	if o != nil && o.Index.IsSet() {
		return true
	}

	return false
}

// SetIndex gets a reference to the given NullableString and assigns it to the Index field.
func (o *ListDocumentParameters) SetIndex(v string) {
	o.Index.Set(&v)
}
// SetIndexNil sets the value for Index to be an explicit nil
func (o *ListDocumentParameters) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil
func (o *ListDocumentParameters) UnsetIndex() {
	o.Index.Unset()
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListDocumentParameters) GetFilter() []DocumentFilters {
	if o == nil {
		var ret []DocumentFilters
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListDocumentParameters) GetFilterOk() ([]DocumentFilters, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ListDocumentParameters) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []DocumentFilters and assigns it to the Filter field.
func (o *ListDocumentParameters) SetFilter(v []DocumentFilters) {
	o.Filter = v
}

// GetWithEmbeddings returns the WithEmbeddings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListDocumentParameters) GetWithEmbeddings() bool {
	if o == nil || IsNil(o.WithEmbeddings.Get()) {
		var ret bool
		return ret
	}
	return *o.WithEmbeddings.Get()
}

// GetWithEmbeddingsOk returns a tuple with the WithEmbeddings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListDocumentParameters) GetWithEmbeddingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithEmbeddings.Get(), o.WithEmbeddings.IsSet()
}

// HasWithEmbeddings returns a boolean if a field has been set.
func (o *ListDocumentParameters) HasWithEmbeddings() bool {
	if o != nil && o.WithEmbeddings.IsSet() {
		return true
	}

	return false
}

// SetWithEmbeddings gets a reference to the given NullableBool and assigns it to the WithEmbeddings field.
func (o *ListDocumentParameters) SetWithEmbeddings(v bool) {
	o.WithEmbeddings.Set(&v)
}
// SetWithEmbeddingsNil sets the value for WithEmbeddings to be an explicit nil
func (o *ListDocumentParameters) SetWithEmbeddingsNil() {
	o.WithEmbeddings.Set(nil)
}

// UnsetWithEmbeddings ensures that no value is present for WithEmbeddings, not even an explicit nil
func (o *ListDocumentParameters) UnsetWithEmbeddings() {
	o.WithEmbeddings.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListDocumentParameters) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListDocumentParameters) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *ListDocumentParameters) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *ListDocumentParameters) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *ListDocumentParameters) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *ListDocumentParameters) UnsetLimit() {
	o.Limit.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListDocumentParameters) GetOffset() int32 {
	if o == nil || IsNil(o.Offset.Get()) {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListDocumentParameters) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *ListDocumentParameters) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *ListDocumentParameters) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *ListDocumentParameters) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *ListDocumentParameters) UnsetOffset() {
	o.Offset.Unset()
}

func (o ListDocumentParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDocumentParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Index.IsSet() {
		toSerialize["index"] = o.Index.Get()
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.WithEmbeddings.IsSet() {
		toSerialize["withEmbeddings"] = o.WithEmbeddings.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	return toSerialize, nil
}

type NullableListDocumentParameters struct {
	value *ListDocumentParameters
	isSet bool
}

func (v NullableListDocumentParameters) Get() *ListDocumentParameters {
	return v.value
}

func (v *NullableListDocumentParameters) Set(val *ListDocumentParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableListDocumentParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableListDocumentParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDocumentParameters(val *ListDocumentParameters) *NullableListDocumentParameters {
	return &NullableListDocumentParameters{value: val, isSet: true}
}

func (v NullableListDocumentParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDocumentParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


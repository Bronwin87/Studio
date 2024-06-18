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

// checks if the TimeSpan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSpan{}

// TimeSpan struct for TimeSpan
type TimeSpan struct {
	Ticks *int64 `json:"ticks,omitempty"`
	Days *int32 `json:"days,omitempty"`
	Hours *int32 `json:"hours,omitempty"`
	Milliseconds *int32 `json:"milliseconds,omitempty"`
	Microseconds *int32 `json:"microseconds,omitempty"`
	Nanoseconds *int32 `json:"nanoseconds,omitempty"`
	Minutes *int32 `json:"minutes,omitempty"`
	Seconds *int32 `json:"seconds,omitempty"`
	TotalDays *float64 `json:"totalDays,omitempty"`
	TotalHours *float64 `json:"totalHours,omitempty"`
	TotalMilliseconds *float64 `json:"totalMilliseconds,omitempty"`
	TotalMicroseconds *float64 `json:"totalMicroseconds,omitempty"`
	TotalNanoseconds *float64 `json:"totalNanoseconds,omitempty"`
	TotalMinutes *float64 `json:"totalMinutes,omitempty"`
	TotalSeconds *float64 `json:"totalSeconds,omitempty"`
}

// NewTimeSpan instantiates a new TimeSpan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSpan() *TimeSpan {
	this := TimeSpan{}
	return &this
}

// NewTimeSpanWithDefaults instantiates a new TimeSpan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSpanWithDefaults() *TimeSpan {
	this := TimeSpan{}
	return &this
}

// GetTicks returns the Ticks field value if set, zero value otherwise.
func (o *TimeSpan) GetTicks() int64 {
	if o == nil || IsNil(o.Ticks) {
		var ret int64
		return ret
	}
	return *o.Ticks
}

// GetTicksOk returns a tuple with the Ticks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.Ticks) {
		return nil, false
	}
	return o.Ticks, true
}

// HasTicks returns a boolean if a field has been set.
func (o *TimeSpan) HasTicks() bool {
	if o != nil && !IsNil(o.Ticks) {
		return true
	}

	return false
}

// SetTicks gets a reference to the given int64 and assigns it to the Ticks field.
func (o *TimeSpan) SetTicks(v int64) {
	o.Ticks = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *TimeSpan) GetDays() int32 {
	if o == nil || IsNil(o.Days) {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *TimeSpan) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *TimeSpan) SetDays(v int32) {
	o.Days = &v
}

// GetHours returns the Hours field value if set, zero value otherwise.
func (o *TimeSpan) GetHours() int32 {
	if o == nil || IsNil(o.Hours) {
		var ret int32
		return ret
	}
	return *o.Hours
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.Hours) {
		return nil, false
	}
	return o.Hours, true
}

// HasHours returns a boolean if a field has been set.
func (o *TimeSpan) HasHours() bool {
	if o != nil && !IsNil(o.Hours) {
		return true
	}

	return false
}

// SetHours gets a reference to the given int32 and assigns it to the Hours field.
func (o *TimeSpan) SetHours(v int32) {
	o.Hours = &v
}

// GetMilliseconds returns the Milliseconds field value if set, zero value otherwise.
func (o *TimeSpan) GetMilliseconds() int32 {
	if o == nil || IsNil(o.Milliseconds) {
		var ret int32
		return ret
	}
	return *o.Milliseconds
}

// GetMillisecondsOk returns a tuple with the Milliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetMillisecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Milliseconds) {
		return nil, false
	}
	return o.Milliseconds, true
}

// HasMilliseconds returns a boolean if a field has been set.
func (o *TimeSpan) HasMilliseconds() bool {
	if o != nil && !IsNil(o.Milliseconds) {
		return true
	}

	return false
}

// SetMilliseconds gets a reference to the given int32 and assigns it to the Milliseconds field.
func (o *TimeSpan) SetMilliseconds(v int32) {
	o.Milliseconds = &v
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *TimeSpan) GetMicroseconds() int32 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int32
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetMicrosecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *TimeSpan) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int32 and assigns it to the Microseconds field.
func (o *TimeSpan) SetMicroseconds(v int32) {
	o.Microseconds = &v
}

// GetNanoseconds returns the Nanoseconds field value if set, zero value otherwise.
func (o *TimeSpan) GetNanoseconds() int32 {
	if o == nil || IsNil(o.Nanoseconds) {
		var ret int32
		return ret
	}
	return *o.Nanoseconds
}

// GetNanosecondsOk returns a tuple with the Nanoseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetNanosecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Nanoseconds) {
		return nil, false
	}
	return o.Nanoseconds, true
}

// HasNanoseconds returns a boolean if a field has been set.
func (o *TimeSpan) HasNanoseconds() bool {
	if o != nil && !IsNil(o.Nanoseconds) {
		return true
	}

	return false
}

// SetNanoseconds gets a reference to the given int32 and assigns it to the Nanoseconds field.
func (o *TimeSpan) SetNanoseconds(v int32) {
	o.Nanoseconds = &v
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *TimeSpan) GetMinutes() int32 {
	if o == nil || IsNil(o.Minutes) {
		var ret int32
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *TimeSpan) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given int32 and assigns it to the Minutes field.
func (o *TimeSpan) SetMinutes(v int32) {
	o.Minutes = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *TimeSpan) GetSeconds() int32 {
	if o == nil || IsNil(o.Seconds) {
		var ret int32
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *TimeSpan) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int32 and assigns it to the Seconds field.
func (o *TimeSpan) SetSeconds(v int32) {
	o.Seconds = &v
}

// GetTotalDays returns the TotalDays field value if set, zero value otherwise.
func (o *TimeSpan) GetTotalDays() float64 {
	if o == nil || IsNil(o.TotalDays) {
		var ret float64
		return ret
	}
	return *o.TotalDays
}

// GetTotalDaysOk returns a tuple with the TotalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetTotalDaysOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalDays) {
		return nil, false
	}
	return o.TotalDays, true
}

// HasTotalDays returns a boolean if a field has been set.
func (o *TimeSpan) HasTotalDays() bool {
	if o != nil && !IsNil(o.TotalDays) {
		return true
	}

	return false
}

// SetTotalDays gets a reference to the given float64 and assigns it to the TotalDays field.
func (o *TimeSpan) SetTotalDays(v float64) {
	o.TotalDays = &v
}

// GetTotalHours returns the TotalHours field value if set, zero value otherwise.
func (o *TimeSpan) GetTotalHours() float64 {
	if o == nil || IsNil(o.TotalHours) {
		var ret float64
		return ret
	}
	return *o.TotalHours
}

// GetTotalHoursOk returns a tuple with the TotalHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetTotalHoursOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalHours) {
		return nil, false
	}
	return o.TotalHours, true
}

// HasTotalHours returns a boolean if a field has been set.
func (o *TimeSpan) HasTotalHours() bool {
	if o != nil && !IsNil(o.TotalHours) {
		return true
	}

	return false
}

// SetTotalHours gets a reference to the given float64 and assigns it to the TotalHours field.
func (o *TimeSpan) SetTotalHours(v float64) {
	o.TotalHours = &v
}

// GetTotalMilliseconds returns the TotalMilliseconds field value if set, zero value otherwise.
func (o *TimeSpan) GetTotalMilliseconds() float64 {
	if o == nil || IsNil(o.TotalMilliseconds) {
		var ret float64
		return ret
	}
	return *o.TotalMilliseconds
}

// GetTotalMillisecondsOk returns a tuple with the TotalMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetTotalMillisecondsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalMilliseconds) {
		return nil, false
	}
	return o.TotalMilliseconds, true
}

// HasTotalMilliseconds returns a boolean if a field has been set.
func (o *TimeSpan) HasTotalMilliseconds() bool {
	if o != nil && !IsNil(o.TotalMilliseconds) {
		return true
	}

	return false
}

// SetTotalMilliseconds gets a reference to the given float64 and assigns it to the TotalMilliseconds field.
func (o *TimeSpan) SetTotalMilliseconds(v float64) {
	o.TotalMilliseconds = &v
}

// GetTotalMicroseconds returns the TotalMicroseconds field value if set, zero value otherwise.
func (o *TimeSpan) GetTotalMicroseconds() float64 {
	if o == nil || IsNil(o.TotalMicroseconds) {
		var ret float64
		return ret
	}
	return *o.TotalMicroseconds
}

// GetTotalMicrosecondsOk returns a tuple with the TotalMicroseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetTotalMicrosecondsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalMicroseconds) {
		return nil, false
	}
	return o.TotalMicroseconds, true
}

// HasTotalMicroseconds returns a boolean if a field has been set.
func (o *TimeSpan) HasTotalMicroseconds() bool {
	if o != nil && !IsNil(o.TotalMicroseconds) {
		return true
	}

	return false
}

// SetTotalMicroseconds gets a reference to the given float64 and assigns it to the TotalMicroseconds field.
func (o *TimeSpan) SetTotalMicroseconds(v float64) {
	o.TotalMicroseconds = &v
}

// GetTotalNanoseconds returns the TotalNanoseconds field value if set, zero value otherwise.
func (o *TimeSpan) GetTotalNanoseconds() float64 {
	if o == nil || IsNil(o.TotalNanoseconds) {
		var ret float64
		return ret
	}
	return *o.TotalNanoseconds
}

// GetTotalNanosecondsOk returns a tuple with the TotalNanoseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetTotalNanosecondsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNanoseconds) {
		return nil, false
	}
	return o.TotalNanoseconds, true
}

// HasTotalNanoseconds returns a boolean if a field has been set.
func (o *TimeSpan) HasTotalNanoseconds() bool {
	if o != nil && !IsNil(o.TotalNanoseconds) {
		return true
	}

	return false
}

// SetTotalNanoseconds gets a reference to the given float64 and assigns it to the TotalNanoseconds field.
func (o *TimeSpan) SetTotalNanoseconds(v float64) {
	o.TotalNanoseconds = &v
}

// GetTotalMinutes returns the TotalMinutes field value if set, zero value otherwise.
func (o *TimeSpan) GetTotalMinutes() float64 {
	if o == nil || IsNil(o.TotalMinutes) {
		var ret float64
		return ret
	}
	return *o.TotalMinutes
}

// GetTotalMinutesOk returns a tuple with the TotalMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetTotalMinutesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalMinutes) {
		return nil, false
	}
	return o.TotalMinutes, true
}

// HasTotalMinutes returns a boolean if a field has been set.
func (o *TimeSpan) HasTotalMinutes() bool {
	if o != nil && !IsNil(o.TotalMinutes) {
		return true
	}

	return false
}

// SetTotalMinutes gets a reference to the given float64 and assigns it to the TotalMinutes field.
func (o *TimeSpan) SetTotalMinutes(v float64) {
	o.TotalMinutes = &v
}

// GetTotalSeconds returns the TotalSeconds field value if set, zero value otherwise.
func (o *TimeSpan) GetTotalSeconds() float64 {
	if o == nil || IsNil(o.TotalSeconds) {
		var ret float64
		return ret
	}
	return *o.TotalSeconds
}

// GetTotalSecondsOk returns a tuple with the TotalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpan) GetTotalSecondsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalSeconds) {
		return nil, false
	}
	return o.TotalSeconds, true
}

// HasTotalSeconds returns a boolean if a field has been set.
func (o *TimeSpan) HasTotalSeconds() bool {
	if o != nil && !IsNil(o.TotalSeconds) {
		return true
	}

	return false
}

// SetTotalSeconds gets a reference to the given float64 and assigns it to the TotalSeconds field.
func (o *TimeSpan) SetTotalSeconds(v float64) {
	o.TotalSeconds = &v
}

func (o TimeSpan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSpan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ticks) {
		toSerialize["ticks"] = o.Ticks
	}
	if !IsNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	if !IsNil(o.Hours) {
		toSerialize["hours"] = o.Hours
	}
	if !IsNil(o.Milliseconds) {
		toSerialize["milliseconds"] = o.Milliseconds
	}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	if !IsNil(o.Nanoseconds) {
		toSerialize["nanoseconds"] = o.Nanoseconds
	}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.TotalDays) {
		toSerialize["totalDays"] = o.TotalDays
	}
	if !IsNil(o.TotalHours) {
		toSerialize["totalHours"] = o.TotalHours
	}
	if !IsNil(o.TotalMilliseconds) {
		toSerialize["totalMilliseconds"] = o.TotalMilliseconds
	}
	if !IsNil(o.TotalMicroseconds) {
		toSerialize["totalMicroseconds"] = o.TotalMicroseconds
	}
	if !IsNil(o.TotalNanoseconds) {
		toSerialize["totalNanoseconds"] = o.TotalNanoseconds
	}
	if !IsNil(o.TotalMinutes) {
		toSerialize["totalMinutes"] = o.TotalMinutes
	}
	if !IsNil(o.TotalSeconds) {
		toSerialize["totalSeconds"] = o.TotalSeconds
	}
	return toSerialize, nil
}

type NullableTimeSpan struct {
	value *TimeSpan
	isSet bool
}

func (v NullableTimeSpan) Get() *TimeSpan {
	return v.value
}

func (v *NullableTimeSpan) Set(val *TimeSpan) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSpan) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSpan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSpan(val *TimeSpan) *NullableTimeSpan {
	return &NullableTimeSpan{value: val, isSet: true}
}

func (v NullableTimeSpan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSpan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



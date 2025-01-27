/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the Attachment type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Attachment{}

// Attachment struct for Attachment
type Attachment struct {
	// The document in Base64-encoded string format.
	Content string `json:"content"`
	// The file format.   Possible values: **application/pdf**, **image/jpg**, **image/jpeg**, **image/png**.
	// Deprecated since Legal Entity Management API v1
	ContentType *string `json:"contentType,omitempty"`
	// The name of the file including the file extension.
	// Deprecated since Legal Entity Management API v1
	Filename *string `json:"filename,omitempty"`
	// The name of the file including the file extension.
	PageName *string `json:"pageName,omitempty"`
	// Specifies which side of the ID card is uploaded.  * When `type` is **driversLicense** or **identityCard**, set this to **front** or **back**.  * When omitted, we infer the page number based on the order of attachments.
	PageType *string `json:"pageType,omitempty"`
}

// NewAttachment instantiates a new Attachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachment(content string) *Attachment {
	this := Attachment{}
	this.Content = content
	return &this
}

// NewAttachmentWithDefaults instantiates a new Attachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentWithDefaults() *Attachment {
	this := Attachment{}
	return &this
}

// GetContent returns the Content field value
func (o *Attachment) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Attachment) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *Attachment) SetContent(v string) {
	o.Content = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
// Deprecated since Legal Entity Management API v1
func (o *Attachment) GetContentType() string {
	if o == nil || common.IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Legal Entity Management API v1
func (o *Attachment) GetContentTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *Attachment) HasContentType() bool {
	if o != nil && !common.IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
// Deprecated since Legal Entity Management API v1
func (o *Attachment) SetContentType(v string) {
	o.ContentType = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
// Deprecated since Legal Entity Management API v1
func (o *Attachment) GetFilename() string {
	if o == nil || common.IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Legal Entity Management API v1
func (o *Attachment) GetFilenameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *Attachment) HasFilename() bool {
	if o != nil && !common.IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
// Deprecated since Legal Entity Management API v1
func (o *Attachment) SetFilename(v string) {
	o.Filename = &v
}

// GetPageName returns the PageName field value if set, zero value otherwise.
func (o *Attachment) GetPageName() string {
	if o == nil || common.IsNil(o.PageName) {
		var ret string
		return ret
	}
	return *o.PageName
}

// GetPageNameOk returns a tuple with the PageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetPageNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PageName) {
		return nil, false
	}
	return o.PageName, true
}

// HasPageName returns a boolean if a field has been set.
func (o *Attachment) HasPageName() bool {
	if o != nil && !common.IsNil(o.PageName) {
		return true
	}

	return false
}

// SetPageName gets a reference to the given string and assigns it to the PageName field.
func (o *Attachment) SetPageName(v string) {
	o.PageName = &v
}

// GetPageType returns the PageType field value if set, zero value otherwise.
func (o *Attachment) GetPageType() string {
	if o == nil || common.IsNil(o.PageType) {
		var ret string
		return ret
	}
	return *o.PageType
}

// GetPageTypeOk returns a tuple with the PageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetPageTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PageType) {
		return nil, false
	}
	return o.PageType, true
}

// HasPageType returns a boolean if a field has been set.
func (o *Attachment) HasPageType() bool {
	if o != nil && !common.IsNil(o.PageType) {
		return true
	}

	return false
}

// SetPageType gets a reference to the given string and assigns it to the PageType field.
func (o *Attachment) SetPageType(v string) {
	o.PageType = &v
}

func (o Attachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	if !common.IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !common.IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !common.IsNil(o.PageName) {
		toSerialize["pageName"] = o.PageName
	}
	if !common.IsNil(o.PageType) {
		toSerialize["pageType"] = o.PageType
	}
	return toSerialize, nil
}

type NullableAttachment struct {
	value *Attachment
	isSet bool
}

func (v NullableAttachment) Get() *Attachment {
	return v.value
}

func (v *NullableAttachment) Set(val *Attachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachment(val *Attachment) *NullableAttachment {
	return &NullableAttachment{value: val, isSet: true}
}

func (v NullableAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

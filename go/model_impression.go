/*
Aura Frame API - Unofficial

Reverse Engineered API for Aura Frames

API version: 0.0.1
Contact: dave@mudsite.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auraframes

import (
	"encoding/json"
)

// Impression struct for Impression
type Impression struct {
	LastViewedOrCreatedAt *string `json:"last_viewed_or_created_at,omitempty"`
	ViewCount *string `json:"view_count,omitempty"`
	GestureDirection *string `json:"gesture_direction,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	LivePhotoOnTransition *string `json:"live_photo_on_transition,omitempty"`
	ViewedAt *string `json:"viewed_at,omitempty"`
	Id *string `json:"id,omitempty"`
	LastViewedAt *string `json:"last_viewed_at,omitempty"`
	LastShownWithAssetId *string `json:"last_shown_with_asset_id,omitempty"`
	FrameId *string `json:"frame_id,omitempty"`
	AssetId *string `json:"asset_id,omitempty"`
	Asset *Asset `json:"asset,omitempty"`
}

// NewImpression instantiates a new Impression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImpression() *Impression {
	this := Impression{}
	return &this
}

// NewImpressionWithDefaults instantiates a new Impression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImpressionWithDefaults() *Impression {
	this := Impression{}
	return &this
}

// GetLastViewedOrCreatedAt returns the LastViewedOrCreatedAt field value if set, zero value otherwise.
func (o *Impression) GetLastViewedOrCreatedAt() string {
	if o == nil || o.LastViewedOrCreatedAt == nil {
		var ret string
		return ret
	}
	return *o.LastViewedOrCreatedAt
}

// GetLastViewedOrCreatedAtOk returns a tuple with the LastViewedOrCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetLastViewedOrCreatedAtOk() (*string, bool) {
	if o == nil || o.LastViewedOrCreatedAt == nil {
		return nil, false
	}
	return o.LastViewedOrCreatedAt, true
}

// HasLastViewedOrCreatedAt returns a boolean if a field has been set.
func (o *Impression) HasLastViewedOrCreatedAt() bool {
	if o != nil && o.LastViewedOrCreatedAt != nil {
		return true
	}

	return false
}

// SetLastViewedOrCreatedAt gets a reference to the given string and assigns it to the LastViewedOrCreatedAt field.
func (o *Impression) SetLastViewedOrCreatedAt(v string) {
	o.LastViewedOrCreatedAt = &v
}

// GetViewCount returns the ViewCount field value if set, zero value otherwise.
func (o *Impression) GetViewCount() string {
	if o == nil || o.ViewCount == nil {
		var ret string
		return ret
	}
	return *o.ViewCount
}

// GetViewCountOk returns a tuple with the ViewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetViewCountOk() (*string, bool) {
	if o == nil || o.ViewCount == nil {
		return nil, false
	}
	return o.ViewCount, true
}

// HasViewCount returns a boolean if a field has been set.
func (o *Impression) HasViewCount() bool {
	if o != nil && o.ViewCount != nil {
		return true
	}

	return false
}

// SetViewCount gets a reference to the given string and assigns it to the ViewCount field.
func (o *Impression) SetViewCount(v string) {
	o.ViewCount = &v
}

// GetGestureDirection returns the GestureDirection field value if set, zero value otherwise.
func (o *Impression) GetGestureDirection() string {
	if o == nil || o.GestureDirection == nil {
		var ret string
		return ret
	}
	return *o.GestureDirection
}

// GetGestureDirectionOk returns a tuple with the GestureDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetGestureDirectionOk() (*string, bool) {
	if o == nil || o.GestureDirection == nil {
		return nil, false
	}
	return o.GestureDirection, true
}

// HasGestureDirection returns a boolean if a field has been set.
func (o *Impression) HasGestureDirection() bool {
	if o != nil && o.GestureDirection != nil {
		return true
	}

	return false
}

// SetGestureDirection gets a reference to the given string and assigns it to the GestureDirection field.
func (o *Impression) SetGestureDirection(v string) {
	o.GestureDirection = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Impression) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Impression) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Impression) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLivePhotoOnTransition returns the LivePhotoOnTransition field value if set, zero value otherwise.
func (o *Impression) GetLivePhotoOnTransition() string {
	if o == nil || o.LivePhotoOnTransition == nil {
		var ret string
		return ret
	}
	return *o.LivePhotoOnTransition
}

// GetLivePhotoOnTransitionOk returns a tuple with the LivePhotoOnTransition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetLivePhotoOnTransitionOk() (*string, bool) {
	if o == nil || o.LivePhotoOnTransition == nil {
		return nil, false
	}
	return o.LivePhotoOnTransition, true
}

// HasLivePhotoOnTransition returns a boolean if a field has been set.
func (o *Impression) HasLivePhotoOnTransition() bool {
	if o != nil && o.LivePhotoOnTransition != nil {
		return true
	}

	return false
}

// SetLivePhotoOnTransition gets a reference to the given string and assigns it to the LivePhotoOnTransition field.
func (o *Impression) SetLivePhotoOnTransition(v string) {
	o.LivePhotoOnTransition = &v
}

// GetViewedAt returns the ViewedAt field value if set, zero value otherwise.
func (o *Impression) GetViewedAt() string {
	if o == nil || o.ViewedAt == nil {
		var ret string
		return ret
	}
	return *o.ViewedAt
}

// GetViewedAtOk returns a tuple with the ViewedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetViewedAtOk() (*string, bool) {
	if o == nil || o.ViewedAt == nil {
		return nil, false
	}
	return o.ViewedAt, true
}

// HasViewedAt returns a boolean if a field has been set.
func (o *Impression) HasViewedAt() bool {
	if o != nil && o.ViewedAt != nil {
		return true
	}

	return false
}

// SetViewedAt gets a reference to the given string and assigns it to the ViewedAt field.
func (o *Impression) SetViewedAt(v string) {
	o.ViewedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Impression) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Impression) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Impression) SetId(v string) {
	o.Id = &v
}

// GetLastViewedAt returns the LastViewedAt field value if set, zero value otherwise.
func (o *Impression) GetLastViewedAt() string {
	if o == nil || o.LastViewedAt == nil {
		var ret string
		return ret
	}
	return *o.LastViewedAt
}

// GetLastViewedAtOk returns a tuple with the LastViewedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetLastViewedAtOk() (*string, bool) {
	if o == nil || o.LastViewedAt == nil {
		return nil, false
	}
	return o.LastViewedAt, true
}

// HasLastViewedAt returns a boolean if a field has been set.
func (o *Impression) HasLastViewedAt() bool {
	if o != nil && o.LastViewedAt != nil {
		return true
	}

	return false
}

// SetLastViewedAt gets a reference to the given string and assigns it to the LastViewedAt field.
func (o *Impression) SetLastViewedAt(v string) {
	o.LastViewedAt = &v
}

// GetLastShownWithAssetId returns the LastShownWithAssetId field value if set, zero value otherwise.
func (o *Impression) GetLastShownWithAssetId() string {
	if o == nil || o.LastShownWithAssetId == nil {
		var ret string
		return ret
	}
	return *o.LastShownWithAssetId
}

// GetLastShownWithAssetIdOk returns a tuple with the LastShownWithAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetLastShownWithAssetIdOk() (*string, bool) {
	if o == nil || o.LastShownWithAssetId == nil {
		return nil, false
	}
	return o.LastShownWithAssetId, true
}

// HasLastShownWithAssetId returns a boolean if a field has been set.
func (o *Impression) HasLastShownWithAssetId() bool {
	if o != nil && o.LastShownWithAssetId != nil {
		return true
	}

	return false
}

// SetLastShownWithAssetId gets a reference to the given string and assigns it to the LastShownWithAssetId field.
func (o *Impression) SetLastShownWithAssetId(v string) {
	o.LastShownWithAssetId = &v
}

// GetFrameId returns the FrameId field value if set, zero value otherwise.
func (o *Impression) GetFrameId() string {
	if o == nil || o.FrameId == nil {
		var ret string
		return ret
	}
	return *o.FrameId
}

// GetFrameIdOk returns a tuple with the FrameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetFrameIdOk() (*string, bool) {
	if o == nil || o.FrameId == nil {
		return nil, false
	}
	return o.FrameId, true
}

// HasFrameId returns a boolean if a field has been set.
func (o *Impression) HasFrameId() bool {
	if o != nil && o.FrameId != nil {
		return true
	}

	return false
}

// SetFrameId gets a reference to the given string and assigns it to the FrameId field.
func (o *Impression) SetFrameId(v string) {
	o.FrameId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *Impression) GetAssetId() string {
	if o == nil || o.AssetId == nil {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetAssetIdOk() (*string, bool) {
	if o == nil || o.AssetId == nil {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *Impression) HasAssetId() bool {
	if o != nil && o.AssetId != nil {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *Impression) SetAssetId(v string) {
	o.AssetId = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *Impression) GetAsset() Asset {
	if o == nil || o.Asset == nil {
		var ret Asset
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impression) GetAssetOk() (*Asset, bool) {
	if o == nil || o.Asset == nil {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *Impression) HasAsset() bool {
	if o != nil && o.Asset != nil {
		return true
	}

	return false
}

// SetAsset gets a reference to the given Asset and assigns it to the Asset field.
func (o *Impression) SetAsset(v Asset) {
	o.Asset = &v
}

func (o Impression) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastViewedOrCreatedAt != nil {
		toSerialize["last_viewed_or_created_at"] = o.LastViewedOrCreatedAt
	}
	if o.ViewCount != nil {
		toSerialize["view_count"] = o.ViewCount
	}
	if o.GestureDirection != nil {
		toSerialize["gesture_direction"] = o.GestureDirection
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LivePhotoOnTransition != nil {
		toSerialize["live_photo_on_transition"] = o.LivePhotoOnTransition
	}
	if o.ViewedAt != nil {
		toSerialize["viewed_at"] = o.ViewedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastViewedAt != nil {
		toSerialize["last_viewed_at"] = o.LastViewedAt
	}
	if o.LastShownWithAssetId != nil {
		toSerialize["last_shown_with_asset_id"] = o.LastShownWithAssetId
	}
	if o.FrameId != nil {
		toSerialize["frame_id"] = o.FrameId
	}
	if o.AssetId != nil {
		toSerialize["asset_id"] = o.AssetId
	}
	if o.Asset != nil {
		toSerialize["asset"] = o.Asset
	}
	return json.Marshal(toSerialize)
}

type NullableImpression struct {
	value *Impression
	isSet bool
}

func (v NullableImpression) Get() *Impression {
	return v.value
}

func (v *NullableImpression) Set(val *Impression) {
	v.value = val
	v.isSet = true
}

func (v NullableImpression) IsSet() bool {
	return v.isSet
}

func (v *NullableImpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImpression(val *Impression) *NullableImpression {
	return &NullableImpression{value: val, isSet: true}
}

func (v NullableImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



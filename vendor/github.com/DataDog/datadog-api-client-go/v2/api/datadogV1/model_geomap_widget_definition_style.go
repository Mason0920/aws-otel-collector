// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapWidgetDefinitionStyle The style to apply to the widget.
type GeomapWidgetDefinitionStyle struct {
	// The color palette to apply to the widget.
	Palette string `json:"palette"`
	// Whether to flip the palette tones.
	PaletteFlip bool `json:"palette_flip"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewGeomapWidgetDefinitionStyle instantiates a new GeomapWidgetDefinitionStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapWidgetDefinitionStyle(palette string, paletteFlip bool) *GeomapWidgetDefinitionStyle {
	this := GeomapWidgetDefinitionStyle{}
	this.Palette = palette
	this.PaletteFlip = paletteFlip
	return &this
}

// NewGeomapWidgetDefinitionStyleWithDefaults instantiates a new GeomapWidgetDefinitionStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapWidgetDefinitionStyleWithDefaults() *GeomapWidgetDefinitionStyle {
	this := GeomapWidgetDefinitionStyle{}
	return &this
}

// GetPalette returns the Palette field value.
func (o *GeomapWidgetDefinitionStyle) GetPalette() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *GeomapWidgetDefinitionStyle) GetPaletteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Palette, true
}

// SetPalette sets field value.
func (o *GeomapWidgetDefinitionStyle) SetPalette(v string) {
	o.Palette = v
}

// GetPaletteFlip returns the PaletteFlip field value.
func (o *GeomapWidgetDefinitionStyle) GetPaletteFlip() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PaletteFlip
}

// GetPaletteFlipOk returns a tuple with the PaletteFlip field value
// and a boolean to check if the value has been set.
func (o *GeomapWidgetDefinitionStyle) GetPaletteFlipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaletteFlip, true
}

// SetPaletteFlip sets field value.
func (o *GeomapWidgetDefinitionStyle) SetPaletteFlip(v bool) {
	o.PaletteFlip = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapWidgetDefinitionStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["palette"] = o.Palette
	toSerialize["palette_flip"] = o.PaletteFlip

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapWidgetDefinitionStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Palette     *string `json:"palette"`
		PaletteFlip *bool   `json:"palette_flip"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Palette == nil {
		return fmt.Errorf("required field palette missing")
	}
	if all.PaletteFlip == nil {
		return fmt.Errorf("required field palette_flip missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"palette", "palette_flip"})
	} else {
		return err
	}
	o.Palette = *all.Palette
	o.PaletteFlip = *all.PaletteFlip

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Extractor an Extractor defines how values are extracted from a JSON document, for use in Labels etc.
type Extractor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Does the JSON path expression reference an Array?
    isarray *bool
    // JSON path expression defining the location of the extractor value in the JSON document. This is a pSQL json path expression
    jsonpath *string
    // Name of extractor. This name is used in Combination Functions to refer to values by name
    name *string
}
// NewExtractor instantiates a new Extractor and sets the default values.
func NewExtractor()(*Extractor) {
    m := &Extractor{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExtractorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExtractorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExtractor(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Extractor) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Extractor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isarray"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsarray(val)
        }
        return nil
    }
    res["jsonpath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJsonpath(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetIsarray gets the isarray property value. Does the JSON path expression reference an Array?
// returns a *bool when successful
func (m *Extractor) GetIsarray()(*bool) {
    return m.isarray
}
// GetJsonpath gets the jsonpath property value. JSON path expression defining the location of the extractor value in the JSON document. This is a pSQL json path expression
// returns a *string when successful
func (m *Extractor) GetJsonpath()(*string) {
    return m.jsonpath
}
// GetName gets the name property value. Name of extractor. This name is used in Combination Functions to refer to values by name
// returns a *string when successful
func (m *Extractor) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *Extractor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isarray", m.GetIsarray())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonpath", m.GetJsonpath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Extractor) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsarray sets the isarray property value. Does the JSON path expression reference an Array?
func (m *Extractor) SetIsarray(value *bool)() {
    m.isarray = value
}
// SetJsonpath sets the jsonpath property value. JSON path expression defining the location of the extractor value in the JSON document. This is a pSQL json path expression
func (m *Extractor) SetJsonpath(value *string)() {
    m.jsonpath = value
}
// SetName sets the name property value. Name of extractor. This name is used in Combination Functions to refer to values by name
func (m *Extractor) SetName(value *string)() {
    m.name = value
}
type Extractorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsarray()(*bool)
    GetJsonpath()(*string)
    GetName()(*string)
    SetIsarray(value *bool)()
    SetJsonpath(value *string)()
    SetName(value *string)()
}

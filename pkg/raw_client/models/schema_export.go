package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SchemaExport represents a Schema with all associated data used for export/import operations.
type SchemaExport struct {
    Schema
    // Array of Labels associated with schema
    labels []Labelable
    // Array of Transformers associated with schema
    transformers []Transformerable
}
// NewSchemaExport instantiates a new SchemaExport and sets the default values.
func NewSchemaExport()(*SchemaExport) {
    m := &SchemaExport{
        Schema: *NewSchema(),
    }
    return m
}
// CreateSchemaExportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaExportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchemaExport(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SchemaExport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Schema.GetFieldDeserializers()
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Labelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Labelable)
                }
            }
            m.SetLabels(res)
        }
        return nil
    }
    res["transformers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTransformerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Transformerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Transformerable)
                }
            }
            m.SetTransformers(res)
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. Array of Labels associated with schema
// returns a []Labelable when successful
func (m *SchemaExport) GetLabels()([]Labelable) {
    return m.labels
}
// GetTransformers gets the transformers property value. Array of Transformers associated with schema
// returns a []Transformerable when successful
func (m *SchemaExport) GetTransformers()([]Transformerable) {
    return m.transformers
}
// Serialize serializes information the current object
func (m *SchemaExport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Schema.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("labels", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransformers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransformers()))
        for i, v := range m.GetTransformers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("transformers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabels sets the labels property value. Array of Labels associated with schema
func (m *SchemaExport) SetLabels(value []Labelable)() {
    m.labels = value
}
// SetTransformers sets the transformers property value. Array of Transformers associated with schema
func (m *SchemaExport) SetTransformers(value []Transformerable)() {
    m.transformers = value
}
type SchemaExportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Schemaable
    GetLabels()([]Labelable)
    GetTransformers()([]Transformerable)
    SetLabels(value []Labelable)()
    SetTransformers(value []Transformerable)()
}

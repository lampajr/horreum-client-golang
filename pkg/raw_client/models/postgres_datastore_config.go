package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PostgresDatastoreConfig built in backend datastore
type PostgresDatastoreConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Built In
    builtIn *bool
}
// NewPostgresDatastoreConfig instantiates a new PostgresDatastoreConfig and sets the default values.
func NewPostgresDatastoreConfig()(*PostgresDatastoreConfig) {
    m := &PostgresDatastoreConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePostgresDatastoreConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePostgresDatastoreConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPostgresDatastoreConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PostgresDatastoreConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuiltIn gets the builtIn property value. Built In
// returns a *bool when successful
func (m *PostgresDatastoreConfig) GetBuiltIn()(*bool) {
    return m.builtIn
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PostgresDatastoreConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["builtIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuiltIn(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PostgresDatastoreConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("builtIn", m.GetBuiltIn())
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
func (m *PostgresDatastoreConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuiltIn sets the builtIn property value. Built In
func (m *PostgresDatastoreConfig) SetBuiltIn(value *bool)() {
    m.builtIn = value
}
type PostgresDatastoreConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuiltIn()(*bool)
    SetBuiltIn(value *bool)()
}

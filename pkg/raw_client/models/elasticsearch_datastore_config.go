package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ElasticsearchDatastoreConfig type of backend datastore
type ElasticsearchDatastoreConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Elasticsearch API KEY
    apiKey *string
    // Built In
    builtIn *bool
    // Elasticsearch password
    password *string
    // Elasticsearch url
    url *string
    // Elasticsearch username
    username *string
}
// NewElasticsearchDatastoreConfig instantiates a new ElasticsearchDatastoreConfig and sets the default values.
func NewElasticsearchDatastoreConfig()(*ElasticsearchDatastoreConfig) {
    m := &ElasticsearchDatastoreConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateElasticsearchDatastoreConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateElasticsearchDatastoreConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewElasticsearchDatastoreConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ElasticsearchDatastoreConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApiKey gets the apiKey property value. Elasticsearch API KEY
// returns a *string when successful
func (m *ElasticsearchDatastoreConfig) GetApiKey()(*string) {
    return m.apiKey
}
// GetBuiltIn gets the builtIn property value. Built In
// returns a *bool when successful
func (m *ElasticsearchDatastoreConfig) GetBuiltIn()(*bool) {
    return m.builtIn
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ElasticsearchDatastoreConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apiKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiKey(val)
        }
        return nil
    }
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
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. Elasticsearch password
// returns a *string when successful
func (m *ElasticsearchDatastoreConfig) GetPassword()(*string) {
    return m.password
}
// GetUrl gets the url property value. Elasticsearch url
// returns a *string when successful
func (m *ElasticsearchDatastoreConfig) GetUrl()(*string) {
    return m.url
}
// GetUsername gets the username property value. Elasticsearch username
// returns a *string when successful
func (m *ElasticsearchDatastoreConfig) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *ElasticsearchDatastoreConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("apiKey", m.GetApiKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("builtIn", m.GetBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
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
func (m *ElasticsearchDatastoreConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApiKey sets the apiKey property value. Elasticsearch API KEY
func (m *ElasticsearchDatastoreConfig) SetApiKey(value *string)() {
    m.apiKey = value
}
// SetBuiltIn sets the builtIn property value. Built In
func (m *ElasticsearchDatastoreConfig) SetBuiltIn(value *bool)() {
    m.builtIn = value
}
// SetPassword sets the password property value. Elasticsearch password
func (m *ElasticsearchDatastoreConfig) SetPassword(value *string)() {
    m.password = value
}
// SetUrl sets the url property value. Elasticsearch url
func (m *ElasticsearchDatastoreConfig) SetUrl(value *string)() {
    m.url = value
}
// SetUsername sets the username property value. Elasticsearch username
func (m *ElasticsearchDatastoreConfig) SetUsername(value *string)() {
    m.username = value
}
type ElasticsearchDatastoreConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiKey()(*string)
    GetBuiltIn()(*bool)
    GetPassword()(*string)
    GetUrl()(*string)
    GetUsername()(*string)
    SetApiKey(value *string)()
    SetBuiltIn(value *bool)()
    SetPassword(value *string)()
    SetUrl(value *string)()
    SetUsername(value *string)()
}

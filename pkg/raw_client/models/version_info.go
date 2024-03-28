package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VersionInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Privacy statement
    privacyStatement *string
    // Timestamp of server startup
    startTimestamp *int64
    // Version of Horreum
    version *string
}
// NewVersionInfo instantiates a new VersionInfo and sets the default values.
func NewVersionInfo()(*VersionInfo) {
    m := &VersionInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVersionInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVersionInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVersionInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VersionInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VersionInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["privacyStatement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyStatement(val)
        }
        return nil
    }
    res["startTimestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTimestamp(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetPrivacyStatement gets the privacyStatement property value. Privacy statement
// returns a *string when successful
func (m *VersionInfo) GetPrivacyStatement()(*string) {
    return m.privacyStatement
}
// GetStartTimestamp gets the startTimestamp property value. Timestamp of server startup
// returns a *int64 when successful
func (m *VersionInfo) GetStartTimestamp()(*int64) {
    return m.startTimestamp
}
// GetVersion gets the version property value. Version of Horreum
// returns a *string when successful
func (m *VersionInfo) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *VersionInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("privacyStatement", m.GetPrivacyStatement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("startTimestamp", m.GetStartTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *VersionInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivacyStatement sets the privacyStatement property value. Privacy statement
func (m *VersionInfo) SetPrivacyStatement(value *string)() {
    m.privacyStatement = value
}
// SetStartTimestamp sets the startTimestamp property value. Timestamp of server startup
func (m *VersionInfo) SetStartTimestamp(value *int64)() {
    m.startTimestamp = value
}
// SetVersion sets the version property value. Version of Horreum
func (m *VersionInfo) SetVersion(value *string)() {
    m.version = value
}
type VersionInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivacyStatement()(*string)
    GetStartTimestamp()(*int64)
    GetVersion()(*string)
    SetPrivacyStatement(value *string)()
    SetStartTimestamp(value *int64)()
    SetVersion(value *string)()
}

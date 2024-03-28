package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TestToken struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Token description
    description *string
    // Unique Token id
    id *int32
    // The permissions property
    permissions *int32
    // Test ID to apply Token
    testId *int32
    // Test value
    value *string
}
// NewTestToken instantiates a new TestToken and sets the default values.
func NewTestToken()(*TestToken) {
    m := &TestToken{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTestTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTestToken(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TestToken) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Token description
// returns a *string when successful
func (m *TestToken) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TestToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val)
        }
        return nil
    }
    res["testId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestId(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique Token id
// returns a *int32 when successful
func (m *TestToken) GetId()(*int32) {
    return m.id
}
// GetPermissions gets the permissions property value. The permissions property
// returns a *int32 when successful
func (m *TestToken) GetPermissions()(*int32) {
    return m.permissions
}
// GetTestId gets the testId property value. Test ID to apply Token
// returns a *int32 when successful
func (m *TestToken) GetTestId()(*int32) {
    return m.testId
}
// GetValue gets the value property value. Test value
// returns a *string when successful
func (m *TestToken) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *TestToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("testId", m.GetTestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *TestToken) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Token description
func (m *TestToken) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. Unique Token id
func (m *TestToken) SetId(value *int32)() {
    m.id = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *TestToken) SetPermissions(value *int32)() {
    m.permissions = value
}
// SetTestId sets the testId property value. Test ID to apply Token
func (m *TestToken) SetTestId(value *int32)() {
    m.testId = value
}
// SetValue sets the value property value. Test value
func (m *TestToken) SetValue(value *string)() {
    m.value = value
}
type TestTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetId()(*int32)
    GetPermissions()(*int32)
    GetTestId()(*int32)
    GetValue()(*string)
    SetDescription(value *string)()
    SetId(value *int32)()
    SetPermissions(value *int32)()
    SetTestId(value *int32)()
    SetValue(value *string)()
}

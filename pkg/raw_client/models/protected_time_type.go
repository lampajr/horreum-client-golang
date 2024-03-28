package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtectedTimeType struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *ProtectedTimeType_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Run Start timestamp
    start *int64
    // Run Stop timestamp
    stop *int64
}
// NewProtectedTimeType instantiates a new ProtectedTimeType and sets the default values.
func NewProtectedTimeType()(*ProtectedTimeType) {
    m := &ProtectedTimeType{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtectedTimeTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectedTimeTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedTimeType(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *ProtectedTimeType_access when successful
func (m *ProtectedTimeType) GetAccess()(*ProtectedTimeType_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProtectedTimeType) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectedTimeType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtectedTimeType_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*ProtectedTimeType_access))
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    res["stop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStop(val)
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *ProtectedTimeType) GetOwner()(*string) {
    return m.owner
}
// GetStart gets the start property value. Run Start timestamp
// returns a *int64 when successful
func (m *ProtectedTimeType) GetStart()(*int64) {
    return m.start
}
// GetStop gets the stop property value. Run Stop timestamp
// returns a *int64 when successful
func (m *ProtectedTimeType) GetStop()(*int64) {
    return m.stop
}
// Serialize serializes information the current object
func (m *ProtectedTimeType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccess() != nil {
        cast := (*m.GetAccess()).String()
        err := writer.WriteStringValue("access", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("stop", m.GetStop())
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
// SetAccess sets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
func (m *ProtectedTimeType) SetAccess(value *ProtectedTimeType_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProtectedTimeType) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *ProtectedTimeType) SetOwner(value *string)() {
    m.owner = value
}
// SetStart sets the start property value. Run Start timestamp
func (m *ProtectedTimeType) SetStart(value *int64)() {
    m.start = value
}
// SetStop sets the stop property value. Run Stop timestamp
func (m *ProtectedTimeType) SetStop(value *int64)() {
    m.stop = value
}
type ProtectedTimeTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*ProtectedTimeType_access)
    GetOwner()(*string)
    GetStart()(*int64)
    GetStop()(*int64)
    SetAccess(value *ProtectedTimeType_access)()
    SetOwner(value *string)()
    SetStart(value *int64)()
    SetStop(value *int64)()
}

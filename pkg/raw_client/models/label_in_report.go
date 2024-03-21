package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LabelInReport struct {
    LabelLocation
    // The configId property
    configId *int32
    // The name property
    name *string
    // The title property
    title *string
    // The where property
    where *string
}
// NewLabelInReport instantiates a new LabelInReport and sets the default values.
func NewLabelInReport()(*LabelInReport) {
    m := &LabelInReport{
        LabelLocation: *NewLabelLocation(),
    }
    return m
}
// CreateLabelInReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelInReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelInReport(), nil
}
// GetConfigId gets the configId property value. The configId property
// returns a *int32 when successful
func (m *LabelInReport) GetConfigId()(*int32) {
    return m.configId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelInReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelLocation.GetFieldDeserializers()
    res["configId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigId(val)
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
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["where"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhere(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *LabelInReport) GetName()(*string) {
    return m.name
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *LabelInReport) GetTitle()(*string) {
    return m.title
}
// GetWhere gets the where property value. The where property
// returns a *string when successful
func (m *LabelInReport) GetWhere()(*string) {
    return m.where
}
// Serialize serializes information the current object
func (m *LabelInReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelLocation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("configId", m.GetConfigId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("where", m.GetWhere())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigId sets the configId property value. The configId property
func (m *LabelInReport) SetConfigId(value *int32)() {
    m.configId = value
}
// SetName sets the name property value. The name property
func (m *LabelInReport) SetName(value *string)() {
    m.name = value
}
// SetTitle sets the title property value. The title property
func (m *LabelInReport) SetTitle(value *string)() {
    m.title = value
}
// SetWhere sets the where property value. The where property
func (m *LabelInReport) SetWhere(value *string)() {
    m.where = value
}
type LabelInReportable interface {
    LabelLocationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigId()(*int32)
    GetName()(*string)
    GetTitle()(*string)
    GetWhere()(*string)
    SetConfigId(value *int32)()
    SetName(value *string)()
    SetTitle(value *string)()
    SetWhere(value *string)()
}

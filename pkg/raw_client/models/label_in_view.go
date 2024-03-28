package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LabelInView struct {
    LabelLocation
    // The componentId property
    componentId *int32
    // The header property
    header *string
    // The viewId property
    viewId *int32
    // The viewName property
    viewName *string
}
// NewLabelInView instantiates a new LabelInView and sets the default values.
func NewLabelInView()(*LabelInView) {
    m := &LabelInView{
        LabelLocation: *NewLabelLocation(),
    }
    return m
}
// CreateLabelInViewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelInViewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelInView(), nil
}
// GetComponentId gets the componentId property value. The componentId property
// returns a *int32 when successful
func (m *LabelInView) GetComponentId()(*int32) {
    return m.componentId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelInView) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelLocation.GetFieldDeserializers()
    res["componentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComponentId(val)
        }
        return nil
    }
    res["header"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeader(val)
        }
        return nil
    }
    res["viewId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewId(val)
        }
        return nil
    }
    res["viewName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewName(val)
        }
        return nil
    }
    return res
}
// GetHeader gets the header property value. The header property
// returns a *string when successful
func (m *LabelInView) GetHeader()(*string) {
    return m.header
}
// GetViewId gets the viewId property value. The viewId property
// returns a *int32 when successful
func (m *LabelInView) GetViewId()(*int32) {
    return m.viewId
}
// GetViewName gets the viewName property value. The viewName property
// returns a *string when successful
func (m *LabelInView) GetViewName()(*string) {
    return m.viewName
}
// Serialize serializes information the current object
func (m *LabelInView) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelLocation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("componentId", m.GetComponentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("header", m.GetHeader())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("viewId", m.GetViewId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("viewName", m.GetViewName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComponentId sets the componentId property value. The componentId property
func (m *LabelInView) SetComponentId(value *int32)() {
    m.componentId = value
}
// SetHeader sets the header property value. The header property
func (m *LabelInView) SetHeader(value *string)() {
    m.header = value
}
// SetViewId sets the viewId property value. The viewId property
func (m *LabelInView) SetViewId(value *int32)() {
    m.viewId = value
}
// SetViewName sets the viewName property value. The viewName property
func (m *LabelInView) SetViewName(value *string)() {
    m.viewName = value
}
type LabelInViewable interface {
    LabelLocationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComponentId()(*int32)
    GetHeader()(*string)
    GetViewId()(*int32)
    GetViewName()(*string)
    SetComponentId(value *int32)()
    SetHeader(value *string)()
    SetViewId(value *int32)()
    SetViewName(value *string)()
}

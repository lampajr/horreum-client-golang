package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LabelInVariable struct {
    LabelLocation
    // The variableId property
    variableId *int32
    // The variableName property
    variableName *string
}
// NewLabelInVariable instantiates a new LabelInVariable and sets the default values.
func NewLabelInVariable()(*LabelInVariable) {
    m := &LabelInVariable{
        LabelLocation: *NewLabelLocation(),
    }
    return m
}
// CreateLabelInVariableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelInVariableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelInVariable(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelInVariable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelLocation.GetFieldDeserializers()
    res["variableId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVariableId(val)
        }
        return nil
    }
    res["variableName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVariableName(val)
        }
        return nil
    }
    return res
}
// GetVariableId gets the variableId property value. The variableId property
// returns a *int32 when successful
func (m *LabelInVariable) GetVariableId()(*int32) {
    return m.variableId
}
// GetVariableName gets the variableName property value. The variableName property
// returns a *string when successful
func (m *LabelInVariable) GetVariableName()(*string) {
    return m.variableName
}
// Serialize serializes information the current object
func (m *LabelInVariable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelLocation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("variableId", m.GetVariableId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("variableName", m.GetVariableName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVariableId sets the variableId property value. The variableId property
func (m *LabelInVariable) SetVariableId(value *int32)() {
    m.variableId = value
}
// SetVariableName sets the variableName property value. The variableName property
func (m *LabelInVariable) SetVariableName(value *string)() {
    m.variableName = value
}
type LabelInVariableable interface {
    LabelLocationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVariableId()(*int32)
    GetVariableName()(*string)
    SetVariableId(value *int32)()
    SetVariableName(value *string)()
}

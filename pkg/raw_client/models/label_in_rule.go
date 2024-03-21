package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LabelInRule struct {
    LabelLocation
    // The ruleId property
    ruleId *int32
    // The ruleName property
    ruleName *string
}
// NewLabelInRule instantiates a new LabelInRule and sets the default values.
func NewLabelInRule()(*LabelInRule) {
    m := &LabelInRule{
        LabelLocation: *NewLabelLocation(),
    }
    return m
}
// CreateLabelInRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelInRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelInRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelInRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelLocation.GetFieldDeserializers()
    res["ruleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleId(val)
        }
        return nil
    }
    res["ruleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleName(val)
        }
        return nil
    }
    return res
}
// GetRuleId gets the ruleId property value. The ruleId property
// returns a *int32 when successful
func (m *LabelInRule) GetRuleId()(*int32) {
    return m.ruleId
}
// GetRuleName gets the ruleName property value. The ruleName property
// returns a *string when successful
func (m *LabelInRule) GetRuleName()(*string) {
    return m.ruleName
}
// Serialize serializes information the current object
func (m *LabelInRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelLocation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("ruleId", m.GetRuleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ruleName", m.GetRuleName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRuleId sets the ruleId property value. The ruleId property
func (m *LabelInRule) SetRuleId(value *int32)() {
    m.ruleId = value
}
// SetRuleName sets the ruleName property value. The ruleName property
func (m *LabelInRule) SetRuleName(value *string)() {
    m.ruleName = value
}
type LabelInRuleable interface {
    LabelLocationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRuleId()(*int32)
    GetRuleName()(*string)
    SetRuleId(value *int32)()
    SetRuleName(value *string)()
}

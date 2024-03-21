package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TestExport represents a Test with all associated data used for export/import operations.
type TestExport struct {
    Test
    // Array of Actions associated with test
    actions []Actionable
    // Datastore associated with test
    datastore TestExport_datastoreable
    // Array of ExperimentProfiles associated with test
    experiments []ExperimentProfileable
    // Array of MissingDataRules associated with test
    missingDataRules []MissingDataRuleable
    // Watcher object associated with test
    subscriptions TestExport_subscriptionsable
    // Array of Variables associated with test
    variables []Variableable
}
// NewTestExport instantiates a new TestExport and sets the default values.
func NewTestExport()(*TestExport) {
    m := &TestExport{
        Test: *NewTest(),
    }
    return m
}
// CreateTestExportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestExportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTestExport(), nil
}
// GetActions gets the actions property value. Array of Actions associated with test
// returns a []Actionable when successful
func (m *TestExport) GetActions()([]Actionable) {
    return m.actions
}
// GetDatastore gets the datastore property value. Datastore associated with test
// returns a TestExport_datastoreable when successful
func (m *TestExport) GetDatastore()(TestExport_datastoreable) {
    return m.datastore
}
// GetExperiments gets the experiments property value. Array of ExperimentProfiles associated with test
// returns a []ExperimentProfileable when successful
func (m *TestExport) GetExperiments()([]ExperimentProfileable) {
    return m.experiments
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TestExport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Test.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Actionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Actionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    res["datastore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTestExport_datastoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatastore(val.(TestExport_datastoreable))
        }
        return nil
    }
    res["experiments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExperimentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExperimentProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExperimentProfileable)
                }
            }
            m.SetExperiments(res)
        }
        return nil
    }
    res["missingDataRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMissingDataRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MissingDataRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MissingDataRuleable)
                }
            }
            m.SetMissingDataRules(res)
        }
        return nil
    }
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTestExport_subscriptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptions(val.(TestExport_subscriptionsable))
        }
        return nil
    }
    res["variables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVariableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Variableable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Variableable)
                }
            }
            m.SetVariables(res)
        }
        return nil
    }
    return res
}
// GetMissingDataRules gets the missingDataRules property value. Array of MissingDataRules associated with test
// returns a []MissingDataRuleable when successful
func (m *TestExport) GetMissingDataRules()([]MissingDataRuleable) {
    return m.missingDataRules
}
// GetSubscriptions gets the subscriptions property value. Watcher object associated with test
// returns a TestExport_subscriptionsable when successful
func (m *TestExport) GetSubscriptions()(TestExport_subscriptionsable) {
    return m.subscriptions
}
// GetVariables gets the variables property value. Array of Variables associated with test
// returns a []Variableable when successful
func (m *TestExport) GetVariables()([]Variableable) {
    return m.variables
}
// Serialize serializes information the current object
func (m *TestExport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Test.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("datastore", m.GetDatastore())
        if err != nil {
            return err
        }
    }
    if m.GetExperiments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExperiments()))
        for i, v := range m.GetExperiments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("experiments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMissingDataRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMissingDataRules()))
        for i, v := range m.GetMissingDataRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("missingDataRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("subscriptions", m.GetSubscriptions())
        if err != nil {
            return err
        }
    }
    if m.GetVariables() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVariables()))
        for i, v := range m.GetVariables() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("variables", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. Array of Actions associated with test
func (m *TestExport) SetActions(value []Actionable)() {
    m.actions = value
}
// SetDatastore sets the datastore property value. Datastore associated with test
func (m *TestExport) SetDatastore(value TestExport_datastoreable)() {
    m.datastore = value
}
// SetExperiments sets the experiments property value. Array of ExperimentProfiles associated with test
func (m *TestExport) SetExperiments(value []ExperimentProfileable)() {
    m.experiments = value
}
// SetMissingDataRules sets the missingDataRules property value. Array of MissingDataRules associated with test
func (m *TestExport) SetMissingDataRules(value []MissingDataRuleable)() {
    m.missingDataRules = value
}
// SetSubscriptions sets the subscriptions property value. Watcher object associated with test
func (m *TestExport) SetSubscriptions(value TestExport_subscriptionsable)() {
    m.subscriptions = value
}
// SetVariables sets the variables property value. Array of Variables associated with test
func (m *TestExport) SetVariables(value []Variableable)() {
    m.variables = value
}
type TestExportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Testable
    GetActions()([]Actionable)
    GetDatastore()(TestExport_datastoreable)
    GetExperiments()([]ExperimentProfileable)
    GetMissingDataRules()([]MissingDataRuleable)
    GetSubscriptions()(TestExport_subscriptionsable)
    GetVariables()([]Variableable)
    SetActions(value []Actionable)()
    SetDatastore(value TestExport_datastoreable)()
    SetExperiments(value []ExperimentProfileable)()
    SetMissingDataRules(value []MissingDataRuleable)()
    SetSubscriptions(value TestExport_subscriptionsable)()
    SetVariables(value []Variableable)()
}

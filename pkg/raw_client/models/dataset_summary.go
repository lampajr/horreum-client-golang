package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DatasetSummary struct {
    ProtectedTimeType
    // Dataset description
    description *string
    // Unique Dataset ID
    id *int32
    // Ordinal position of Dataset Summary on returned List
    ordinal *int32
    // Run ID that Dataset relates to
    runId *int32
    // List of Schema usages
    schemas []SchemaUsageable
    // Test ID that Dataset relates to
    testId *int32
    // Test name that the Dataset relates to
    testname *string
    // List of Validation Errors
    validationErrors []ValidationErrorable
    // map of view component ids to the LabelValueMap to render the component for this dataset
    view DatasetSummary_viewable
}
// NewDatasetSummary instantiates a new DatasetSummary and sets the default values.
func NewDatasetSummary()(*DatasetSummary) {
    m := &DatasetSummary{
        ProtectedTimeType: *NewProtectedTimeType(),
    }
    return m
}
// CreateDatasetSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatasetSummary(), nil
}
// GetDescription gets the description property value. Dataset description
// returns a *string when successful
func (m *DatasetSummary) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatasetSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectedTimeType.GetFieldDeserializers()
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
    res["ordinal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrdinal(val)
        }
        return nil
    }
    res["runId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunId(val)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSchemaUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SchemaUsageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SchemaUsageable)
                }
            }
            m.SetSchemas(res)
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
    res["testname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestname(val)
        }
        return nil
    }
    res["validationErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateValidationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ValidationErrorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ValidationErrorable)
                }
            }
            m.SetValidationErrors(res)
        }
        return nil
    }
    res["view"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDatasetSummary_viewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetView(val.(DatasetSummary_viewable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique Dataset ID
// returns a *int32 when successful
func (m *DatasetSummary) GetId()(*int32) {
    return m.id
}
// GetOrdinal gets the ordinal property value. Ordinal position of Dataset Summary on returned List
// returns a *int32 when successful
func (m *DatasetSummary) GetOrdinal()(*int32) {
    return m.ordinal
}
// GetRunId gets the runId property value. Run ID that Dataset relates to
// returns a *int32 when successful
func (m *DatasetSummary) GetRunId()(*int32) {
    return m.runId
}
// GetSchemas gets the schemas property value. List of Schema usages
// returns a []SchemaUsageable when successful
func (m *DatasetSummary) GetSchemas()([]SchemaUsageable) {
    return m.schemas
}
// GetTestId gets the testId property value. Test ID that Dataset relates to
// returns a *int32 when successful
func (m *DatasetSummary) GetTestId()(*int32) {
    return m.testId
}
// GetTestname gets the testname property value. Test name that the Dataset relates to
// returns a *string when successful
func (m *DatasetSummary) GetTestname()(*string) {
    return m.testname
}
// GetValidationErrors gets the validationErrors property value. List of Validation Errors
// returns a []ValidationErrorable when successful
func (m *DatasetSummary) GetValidationErrors()([]ValidationErrorable) {
    return m.validationErrors
}
// GetView gets the view property value. map of view component ids to the LabelValueMap to render the component for this dataset
// returns a DatasetSummary_viewable when successful
func (m *DatasetSummary) GetView()(DatasetSummary_viewable) {
    return m.view
}
// Serialize serializes information the current object
func (m *DatasetSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectedTimeType.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ordinal", m.GetOrdinal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("runId", m.GetRunId())
        if err != nil {
            return err
        }
    }
    if m.GetSchemas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchemas()))
        for i, v := range m.GetSchemas() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("schemas", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("testId", m.GetTestId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("testname", m.GetTestname())
        if err != nil {
            return err
        }
    }
    if m.GetValidationErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValidationErrors()))
        for i, v := range m.GetValidationErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("validationErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("view", m.GetView())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Dataset description
func (m *DatasetSummary) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. Unique Dataset ID
func (m *DatasetSummary) SetId(value *int32)() {
    m.id = value
}
// SetOrdinal sets the ordinal property value. Ordinal position of Dataset Summary on returned List
func (m *DatasetSummary) SetOrdinal(value *int32)() {
    m.ordinal = value
}
// SetRunId sets the runId property value. Run ID that Dataset relates to
func (m *DatasetSummary) SetRunId(value *int32)() {
    m.runId = value
}
// SetSchemas sets the schemas property value. List of Schema usages
func (m *DatasetSummary) SetSchemas(value []SchemaUsageable)() {
    m.schemas = value
}
// SetTestId sets the testId property value. Test ID that Dataset relates to
func (m *DatasetSummary) SetTestId(value *int32)() {
    m.testId = value
}
// SetTestname sets the testname property value. Test name that the Dataset relates to
func (m *DatasetSummary) SetTestname(value *string)() {
    m.testname = value
}
// SetValidationErrors sets the validationErrors property value. List of Validation Errors
func (m *DatasetSummary) SetValidationErrors(value []ValidationErrorable)() {
    m.validationErrors = value
}
// SetView sets the view property value. map of view component ids to the LabelValueMap to render the component for this dataset
func (m *DatasetSummary) SetView(value DatasetSummary_viewable)() {
    m.view = value
}
type DatasetSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectedTimeTypeable
    GetDescription()(*string)
    GetId()(*int32)
    GetOrdinal()(*int32)
    GetRunId()(*int32)
    GetSchemas()([]SchemaUsageable)
    GetTestId()(*int32)
    GetTestname()(*string)
    GetValidationErrors()([]ValidationErrorable)
    GetView()(DatasetSummary_viewable)
    SetDescription(value *string)()
    SetId(value *int32)()
    SetOrdinal(value *int32)()
    SetRunId(value *int32)()
    SetSchemas(value []SchemaUsageable)()
    SetTestId(value *int32)()
    SetTestname(value *string)()
    SetValidationErrors(value []ValidationErrorable)()
    SetView(value DatasetSummary_viewable)()
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DatasetLog dataset Log
type DatasetLog struct {
    PersistentLog
    // The datasetId property
    datasetId *int32
    // The datasetOrdinal property
    datasetOrdinal *int32
    // The runId property
    runId *int32
    // The source property
    source *string
    // The testId property
    testId *int32
}
// NewDatasetLog instantiates a new DatasetLog and sets the default values.
func NewDatasetLog()(*DatasetLog) {
    m := &DatasetLog{
        PersistentLog: *NewPersistentLog(),
    }
    return m
}
// CreateDatasetLogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetLogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatasetLog(), nil
}
// GetDatasetId gets the datasetId property value. The datasetId property
// returns a *int32 when successful
func (m *DatasetLog) GetDatasetId()(*int32) {
    return m.datasetId
}
// GetDatasetOrdinal gets the datasetOrdinal property value. The datasetOrdinal property
// returns a *int32 when successful
func (m *DatasetLog) GetDatasetOrdinal()(*int32) {
    return m.datasetOrdinal
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatasetLog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PersistentLog.GetFieldDeserializers()
    res["datasetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatasetId(val)
        }
        return nil
    }
    res["datasetOrdinal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatasetOrdinal(val)
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
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
    return res
}
// GetRunId gets the runId property value. The runId property
// returns a *int32 when successful
func (m *DatasetLog) GetRunId()(*int32) {
    return m.runId
}
// GetSource gets the source property value. The source property
// returns a *string when successful
func (m *DatasetLog) GetSource()(*string) {
    return m.source
}
// GetTestId gets the testId property value. The testId property
// returns a *int32 when successful
func (m *DatasetLog) GetTestId()(*int32) {
    return m.testId
}
// Serialize serializes information the current object
func (m *DatasetLog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PersistentLog.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("datasetId", m.GetDatasetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("datasetOrdinal", m.GetDatasetOrdinal())
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
    {
        err = writer.WriteStringValue("source", m.GetSource())
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
    return nil
}
// SetDatasetId sets the datasetId property value. The datasetId property
func (m *DatasetLog) SetDatasetId(value *int32)() {
    m.datasetId = value
}
// SetDatasetOrdinal sets the datasetOrdinal property value. The datasetOrdinal property
func (m *DatasetLog) SetDatasetOrdinal(value *int32)() {
    m.datasetOrdinal = value
}
// SetRunId sets the runId property value. The runId property
func (m *DatasetLog) SetRunId(value *int32)() {
    m.runId = value
}
// SetSource sets the source property value. The source property
func (m *DatasetLog) SetSource(value *string)() {
    m.source = value
}
// SetTestId sets the testId property value. The testId property
func (m *DatasetLog) SetTestId(value *int32)() {
    m.testId = value
}
type DatasetLogable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PersistentLogable
    GetDatasetId()(*int32)
    GetDatasetOrdinal()(*int32)
    GetRunId()(*int32)
    GetSource()(*string)
    GetTestId()(*int32)
    SetDatasetId(value *int32)()
    SetDatasetOrdinal(value *int32)()
    SetRunId(value *int32)()
    SetSource(value *string)()
    SetTestId(value *int32)()
}

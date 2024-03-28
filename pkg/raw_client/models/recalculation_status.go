package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RecalculationStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Total number of generated datasets
    datasets *int64
    // Total number of completed recalculations
    finished *int64
    // Recalculation timestamp
    timestamp *int64
    // Total number of Runs being recalculated
    totalRuns *int64
}
// NewRecalculationStatus instantiates a new RecalculationStatus and sets the default values.
func NewRecalculationStatus()(*RecalculationStatus) {
    m := &RecalculationStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecalculationStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecalculationStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecalculationStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RecalculationStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatasets gets the datasets property value. Total number of generated datasets
// returns a *int64 when successful
func (m *RecalculationStatus) GetDatasets()(*int64) {
    return m.datasets
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecalculationStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datasets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatasets(val)
        }
        return nil
    }
    res["finished"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinished(val)
        }
        return nil
    }
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    res["totalRuns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRuns(val)
        }
        return nil
    }
    return res
}
// GetFinished gets the finished property value. Total number of completed recalculations
// returns a *int64 when successful
func (m *RecalculationStatus) GetFinished()(*int64) {
    return m.finished
}
// GetTimestamp gets the timestamp property value. Recalculation timestamp
// returns a *int64 when successful
func (m *RecalculationStatus) GetTimestamp()(*int64) {
    return m.timestamp
}
// GetTotalRuns gets the totalRuns property value. Total number of Runs being recalculated
// returns a *int64 when successful
func (m *RecalculationStatus) GetTotalRuns()(*int64) {
    return m.totalRuns
}
// Serialize serializes information the current object
func (m *RecalculationStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("datasets", m.GetDatasets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("finished", m.GetFinished())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("timestamp", m.GetTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalRuns", m.GetTotalRuns())
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
func (m *RecalculationStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatasets sets the datasets property value. Total number of generated datasets
func (m *RecalculationStatus) SetDatasets(value *int64)() {
    m.datasets = value
}
// SetFinished sets the finished property value. Total number of completed recalculations
func (m *RecalculationStatus) SetFinished(value *int64)() {
    m.finished = value
}
// SetTimestamp sets the timestamp property value. Recalculation timestamp
func (m *RecalculationStatus) SetTimestamp(value *int64)() {
    m.timestamp = value
}
// SetTotalRuns sets the totalRuns property value. Total number of Runs being recalculated
func (m *RecalculationStatus) SetTotalRuns(value *int64)() {
    m.totalRuns = value
}
type RecalculationStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatasets()(*int64)
    GetFinished()(*int64)
    GetTimestamp()(*int64)
    GetTotalRuns()(*int64)
    SetDatasets(value *int64)()
    SetFinished(value *int64)()
    SetTimestamp(value *int64)()
    SetTotalRuns(value *int64)()
}

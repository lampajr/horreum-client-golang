package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DatasetList result containing a subset of Dataset Summaries and the total count of available. Used in paginated tables
type DatasetList struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of Dataset Summaries. This is often a subset of total available.
    datasets []DatasetSummaryable
    // Total number of Dataset Summaries available
    total *int64
}
// NewDatasetList instantiates a new DatasetList and sets the default values.
func NewDatasetList()(*DatasetList) {
    m := &DatasetList{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDatasetListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatasetList(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DatasetList) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatasets gets the datasets property value. List of Dataset Summaries. This is often a subset of total available.
// returns a []DatasetSummaryable when successful
func (m *DatasetList) GetDatasets()([]DatasetSummaryable) {
    return m.datasets
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatasetList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datasets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDatasetSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DatasetSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DatasetSummaryable)
                }
            }
            m.SetDatasets(res)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetTotal gets the total property value. Total number of Dataset Summaries available
// returns a *int64 when successful
func (m *DatasetList) GetTotal()(*int64) {
    return m.total
}
// Serialize serializes information the current object
func (m *DatasetList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDatasets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDatasets()))
        for i, v := range m.GetDatasets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("datasets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("total", m.GetTotal())
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
func (m *DatasetList) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatasets sets the datasets property value. List of Dataset Summaries. This is often a subset of total available.
func (m *DatasetList) SetDatasets(value []DatasetSummaryable)() {
    m.datasets = value
}
// SetTotal sets the total property value. Total number of Dataset Summaries available
func (m *DatasetList) SetTotal(value *int64)() {
    m.total = value
}
type DatasetListable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatasets()([]DatasetSummaryable)
    GetTotal()(*int64)
    SetDatasets(value []DatasetSummaryable)()
    SetTotal(value *int64)()
}

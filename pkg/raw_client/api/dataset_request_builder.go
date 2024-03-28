package api

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DatasetRequestBuilder builds and executes requests for operations under \api\dataset
type DatasetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByDatasetIdId gets an item from the github.com/hyperfoil/horreum/pkg/raw_client.api.dataset.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *DatasetDatasetIdItemRequestBuilder when successful
func (m *DatasetRequestBuilder) ByDatasetIdId(datasetIdId string)(*DatasetDatasetIdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if datasetIdId != "" {
        urlTplParams["datasetId%2Did"] = datasetIdId
    }
    return NewDatasetDatasetIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByDatasetIdIdInteger gets an item from the github.com/hyperfoil/horreum/pkg/raw_client.api.dataset.item collection
// returns a *DatasetDatasetIdItemRequestBuilder when successful
func (m *DatasetRequestBuilder) ByDatasetIdIdInteger(datasetIdId int32)(*DatasetDatasetIdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["datasetId%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(datasetIdId), 10)
    return NewDatasetDatasetIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// BySchema the bySchema property
// returns a *DatasetBySchemaRequestBuilder when successful
func (m *DatasetRequestBuilder) BySchema()(*DatasetBySchemaRequestBuilder) {
    return NewDatasetBySchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDatasetRequestBuilderInternal instantiates a new DatasetRequestBuilder and sets the default values.
func NewDatasetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetRequestBuilder) {
    m := &DatasetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset", pathParameters),
    }
    return m
}
// NewDatasetRequestBuilder instantiates a new DatasetRequestBuilder and sets the default values.
func NewDatasetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetRequestBuilderInternal(urlParams, requestAdapter)
}
// List the list property
// returns a *DatasetListRequestBuilder when successful
func (m *DatasetRequestBuilder) List()(*DatasetListRequestBuilder) {
    return NewDatasetListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

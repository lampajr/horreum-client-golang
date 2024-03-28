package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// DatasetListWithTestItemRequestBuilder builds and executes requests for operations under \api\dataset\list\{testId}
type DatasetListWithTestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DatasetListWithTestItemRequestBuilderGetQueryParameters retrieve a paginated list of Datasets, with total count, by Test
type DatasetListWithTestItemRequestBuilderGetQueryParameters struct {
    // Sort direction
    // Deprecated: This property is deprecated, use DirectionAsSortDirection instead
    Direction *string `uriparametername:"direction"`
    // Sort direction
    DirectionAsSortDirection *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SortDirection `uriparametername:"direction"`
    // JOSN Filter expression to apply to query
    Filter *string `uriparametername:"filter"`
    // limit the number of results
    Limit *int32 `uriparametername:"limit"`
    // filter by page number of a paginated list of Schemas
    Page *int32 `uriparametername:"page"`
    // Field name to sort results
    Sort *string `uriparametername:"sort"`
    // Optional View ID to filter datasets by view
    ViewId *int32 `uriparametername:"viewId"`
}
// DatasetListWithTestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasetListWithTestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DatasetListWithTestItemRequestBuilderGetQueryParameters
}
// NewDatasetListWithTestItemRequestBuilderInternal instantiates a new DatasetListWithTestItemRequestBuilder and sets the default values.
func NewDatasetListWithTestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetListWithTestItemRequestBuilder) {
    m := &DatasetListWithTestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset/list/{testId}{?direction*,filter*,limit*,page*,sort*,viewId*}", pathParameters),
    }
    return m
}
// NewDatasetListWithTestItemRequestBuilder instantiates a new DatasetListWithTestItemRequestBuilder and sets the default values.
func NewDatasetListWithTestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetListWithTestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetListWithTestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a paginated list of Datasets, with total count, by Test
// returns a DatasetListable when successful
func (m *DatasetListWithTestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DatasetListWithTestItemRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.DatasetListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateDatasetListFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.DatasetListable), nil
}
// ToGetRequestInformation retrieve a paginated list of Datasets, with total count, by Test
// returns a *RequestInformation when successful
func (m *DatasetListWithTestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DatasetListWithTestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DatasetListWithTestItemRequestBuilder when successful
func (m *DatasetListWithTestItemRequestBuilder) WithUrl(rawUrl string)(*DatasetListWithTestItemRequestBuilder) {
    return NewDatasetListWithTestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

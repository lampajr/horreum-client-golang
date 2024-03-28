package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunListWithTestItemRequestBuilder builds and executes requests for operations under \api\run\list\{testId}
type RunListWithTestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunListWithTestItemRequestBuilderGetQueryParameters retrieve a paginated list of Runs with available count for a given Test ID
type RunListWithTestItemRequestBuilderGetQueryParameters struct {
    // Sort direction
    // Deprecated: This property is deprecated, use DirectionAsSortDirection instead
    Direction *string `uriparametername:"direction"`
    // Sort direction
    DirectionAsSortDirection *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SortDirection `uriparametername:"direction"`
    // limit the number of results
    Limit *int32 `uriparametername:"limit"`
    // filter by page number of a paginated list of Tests
    Page *int32 `uriparametername:"page"`
    // Field name to sort results
    Sort *string `uriparametername:"sort"`
    // include trashed runs
    Trashed *bool `uriparametername:"trashed"`
}
// RunListWithTestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunListWithTestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunListWithTestItemRequestBuilderGetQueryParameters
}
// NewRunListWithTestItemRequestBuilderInternal instantiates a new RunListWithTestItemRequestBuilder and sets the default values.
func NewRunListWithTestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunListWithTestItemRequestBuilder) {
    m := &RunListWithTestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/list/{testId}{?direction*,limit*,page*,sort*,trashed*}", pathParameters),
    }
    return m
}
// NewRunListWithTestItemRequestBuilder instantiates a new RunListWithTestItemRequestBuilder and sets the default values.
func NewRunListWithTestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunListWithTestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunListWithTestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a paginated list of Runs with available count for a given Test ID
// returns a RunsSummaryable when successful
func (m *RunListWithTestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RunListWithTestItemRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunsSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateRunsSummaryFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunsSummaryable), nil
}
// ToGetRequestInformation retrieve a paginated list of Runs with available count for a given Test ID
// returns a *RequestInformation when successful
func (m *RunListWithTestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunListWithTestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunListWithTestItemRequestBuilder when successful
func (m *RunListWithTestItemRequestBuilder) WithUrl(rawUrl string)(*RunListWithTestItemRequestBuilder) {
    return NewRunListWithTestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

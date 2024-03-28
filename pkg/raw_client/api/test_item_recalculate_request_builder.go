package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// TestItemRecalculateRequestBuilder builds and executes requests for operations under \api\test\{id}\recalculate
type TestItemRecalculateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemRecalculateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemRecalculateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TestItemRecalculateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemRecalculateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemRecalculateRequestBuilderInternal instantiates a new TestItemRecalculateRequestBuilder and sets the default values.
func NewTestItemRecalculateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemRecalculateRequestBuilder) {
    m := &TestItemRecalculateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/recalculate", pathParameters),
    }
    return m
}
// NewTestItemRecalculateRequestBuilder instantiates a new TestItemRecalculateRequestBuilder and sets the default values.
func NewTestItemRecalculateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemRecalculateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemRecalculateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get recalculation status for Test
// returns a RecalculationStatusable when successful
func (m *TestItemRecalculateRequestBuilder) Get(ctx context.Context, requestConfiguration *TestItemRecalculateRequestBuilderGetRequestConfiguration)(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.RecalculationStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.CreateRecalculationStatusFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.RecalculationStatusable), nil
}
// Post recalculate Datasets for Test
func (m *TestItemRecalculateRequestBuilder) Post(ctx context.Context, requestConfiguration *TestItemRecalculateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation get recalculation status for Test
// returns a *RequestInformation when successful
func (m *TestItemRecalculateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestItemRecalculateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation recalculate Datasets for Test
// returns a *RequestInformation when successful
func (m *TestItemRecalculateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TestItemRecalculateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemRecalculateRequestBuilder when successful
func (m *TestItemRecalculateRequestBuilder) WithUrl(rawUrl string)(*TestItemRecalculateRequestBuilder) {
    return NewTestItemRecalculateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

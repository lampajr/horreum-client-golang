package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// TestItemTokensRequestBuilder builds and executes requests for operations under \api\test\{id}\tokens
type TestItemTokensRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemTokensRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemTokensRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemTokensRequestBuilderInternal instantiates a new TestItemTokensRequestBuilder and sets the default values.
func NewTestItemTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemTokensRequestBuilder) {
    m := &TestItemTokensRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/tokens", pathParameters),
    }
    return m
}
// NewTestItemTokensRequestBuilder instantiates a new TestItemTokensRequestBuilder and sets the default values.
func NewTestItemTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a collection of Test Tokens for a given Test
// returns a []TestTokenable when successful
func (m *TestItemTokensRequestBuilder) Get(ctx context.Context, requestConfiguration *TestItemTokensRequestBuilderGetRequestConfiguration)([]i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TestTokenable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.CreateTestTokenFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TestTokenable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TestTokenable)
        }
    }
    return val, nil
}
// ToGetRequestInformation a collection of Test Tokens for a given Test
// returns a *RequestInformation when successful
func (m *TestItemTokensRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestItemTokensRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemTokensRequestBuilder when successful
func (m *TestItemTokensRequestBuilder) WithUrl(rawUrl string)(*TestItemTokensRequestBuilder) {
    return NewTestItemTokensRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

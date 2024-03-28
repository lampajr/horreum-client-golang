package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// RunTestRequestBuilder builds and executes requests for operations under \api\run\test
type RunTestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunTestRequestBuilderPostQueryParameters upload a new Run
type RunTestRequestBuilderPostQueryParameters struct {
    // New Access level
    // Deprecated: This property is deprecated, use AccessAsAccess instead
    Access *string `uriparametername:"access"`
    // New Access level
    AccessAsAccess *i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.Access `uriparametername:"access"`
    // Name of the new owner
    Owner *string `uriparametername:"owner"`
    // test name of ID
    Test *string `uriparametername:"test"`
    // API token
    Token *string `uriparametername:"token"`
}
// RunTestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunTestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunTestRequestBuilderPostQueryParameters
}
// NewRunTestRequestBuilderInternal instantiates a new RunTestRequestBuilder and sets the default values.
func NewRunTestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunTestRequestBuilder) {
    m := &RunTestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/test{?access*,owner*,test*,token*}", pathParameters),
    }
    return m
}
// NewRunTestRequestBuilder instantiates a new RunTestRequestBuilder and sets the default values.
func NewRunTestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunTestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunTestRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a new Run
// returns a []byte when successful
func (m *RunTestRequestBuilder) Post(ctx context.Context, body i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.Runable, requestConfiguration *RunTestRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation upload a new Run
// returns a *RequestInformation when successful
func (m *RunTestRequestBuilder) ToPostRequestInformation(ctx context.Context, body i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.Runable, requestConfiguration *RunTestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunTestRequestBuilder when successful
func (m *RunTestRequestBuilder) WithUrl(rawUrl string)(*RunTestRequestBuilder) {
    return NewRunTestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

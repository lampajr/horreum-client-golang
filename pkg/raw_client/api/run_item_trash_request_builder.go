package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemTrashRequestBuilder builds and executes requests for operations under \api\run\{id}\trash
type RunItemTrashRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemTrashRequestBuilderPostQueryParameters trash a Run with a given ID
type RunItemTrashRequestBuilderPostQueryParameters struct {
    // should run be trashed?
    IsTrashed *bool `uriparametername:"isTrashed"`
}
// RunItemTrashRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemTrashRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunItemTrashRequestBuilderPostQueryParameters
}
// NewRunItemTrashRequestBuilderInternal instantiates a new RunItemTrashRequestBuilder and sets the default values.
func NewRunItemTrashRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemTrashRequestBuilder) {
    m := &RunItemTrashRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/trash{?isTrashed*}", pathParameters),
    }
    return m
}
// NewRunItemTrashRequestBuilder instantiates a new RunItemTrashRequestBuilder and sets the default values.
func NewRunItemTrashRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemTrashRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemTrashRequestBuilderInternal(urlParams, requestAdapter)
}
// Post trash a Run with a given ID
func (m *RunItemTrashRequestBuilder) Post(ctx context.Context, requestConfiguration *RunItemTrashRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation trash a Run with a given ID
// returns a *RequestInformation when successful
func (m *RunItemTrashRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RunItemTrashRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemTrashRequestBuilder when successful
func (m *RunItemTrashRequestBuilder) WithUrl(rawUrl string)(*RunItemTrashRequestBuilder) {
    return NewRunItemTrashRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

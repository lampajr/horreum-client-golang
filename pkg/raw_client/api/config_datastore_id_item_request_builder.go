package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// ConfigDatastoreIdItemRequestBuilder builds and executes requests for operations under \api\config\datastore\{id-id}
type ConfigDatastoreIdItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigDatastoreIdItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigDatastoreIdItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConfigDatastoreIdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigDatastoreIdItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigDatastoreIdItemRequestBuilderInternal instantiates a new ConfigDatastoreIdItemRequestBuilder and sets the default values.
func NewConfigDatastoreIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigDatastoreIdItemRequestBuilder) {
    m := &ConfigDatastoreIdItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/config/datastore/{id%2Did}", pathParameters),
    }
    return m
}
// NewConfigDatastoreIdItemRequestBuilder instantiates a new ConfigDatastoreIdItemRequestBuilder and sets the default values.
func NewConfigDatastoreIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigDatastoreIdItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigDatastoreIdItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete test a Datastore
func (m *ConfigDatastoreIdItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConfigDatastoreIdItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get obtain list of configured datastores for particular team
// returns a []Datastoreable when successful
func (m *ConfigDatastoreIdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigDatastoreIdItemRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Datastoreable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateDatastoreFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Datastoreable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Datastoreable)
        }
    }
    return val, nil
}
// Test the test property
// returns a *ConfigDatastoreItemTestRequestBuilder when successful
func (m *ConfigDatastoreIdItemRequestBuilder) Test()(*ConfigDatastoreItemTestRequestBuilder) {
    return NewConfigDatastoreItemTestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation test a Datastore
// returns a *RequestInformation when successful
func (m *ConfigDatastoreIdItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConfigDatastoreIdItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation obtain list of configured datastores for particular team
// returns a *RequestInformation when successful
func (m *ConfigDatastoreIdItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigDatastoreIdItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConfigDatastoreIdItemRequestBuilder when successful
func (m *ConfigDatastoreIdItemRequestBuilder) WithUrl(rawUrl string)(*ConfigDatastoreIdItemRequestBuilder) {
    return NewConfigDatastoreIdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

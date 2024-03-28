package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// SchemaDescriptorsRequestBuilder builds and executes requests for operations under \api\schema\descriptors
type SchemaDescriptorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaDescriptorsRequestBuilderGetQueryParameters retrieve a list of Schema Descriptors
type SchemaDescriptorsRequestBuilderGetQueryParameters struct {
    // Limit to a single Schema by ID
    Id []int32 `uriparametername:"id"`
}
// SchemaDescriptorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaDescriptorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchemaDescriptorsRequestBuilderGetQueryParameters
}
// NewSchemaDescriptorsRequestBuilderInternal instantiates a new SchemaDescriptorsRequestBuilder and sets the default values.
func NewSchemaDescriptorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaDescriptorsRequestBuilder) {
    m := &SchemaDescriptorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/descriptors{?id*}", pathParameters),
    }
    return m
}
// NewSchemaDescriptorsRequestBuilder instantiates a new SchemaDescriptorsRequestBuilder and sets the default values.
func NewSchemaDescriptorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaDescriptorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaDescriptorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a list of Schema Descriptors
// returns a []SchemaDescriptorable when successful
func (m *SchemaDescriptorsRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaDescriptorsRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SchemaDescriptorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateSchemaDescriptorFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SchemaDescriptorable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SchemaDescriptorable)
        }
    }
    return val, nil
}
// ToGetRequestInformation retrieve a list of Schema Descriptors
// returns a *RequestInformation when successful
func (m *SchemaDescriptorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaDescriptorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SchemaDescriptorsRequestBuilder when successful
func (m *SchemaDescriptorsRequestBuilder) WithUrl(rawUrl string)(*SchemaDescriptorsRequestBuilder) {
    return NewSchemaDescriptorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

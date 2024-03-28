package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// SchemaAllTransformersRequestBuilder builds and executes requests for operations under \api\schema\allTransformers
type SchemaAllTransformersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaAllTransformersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaAllTransformersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaAllTransformersRequestBuilderInternal instantiates a new SchemaAllTransformersRequestBuilder and sets the default values.
func NewSchemaAllTransformersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaAllTransformersRequestBuilder) {
    m := &SchemaAllTransformersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/allTransformers", pathParameters),
    }
    return m
}
// NewSchemaAllTransformersRequestBuilder instantiates a new SchemaAllTransformersRequestBuilder and sets the default values.
func NewSchemaAllTransformersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaAllTransformersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaAllTransformersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve all transformers
// returns a []TransformerInfoable when successful
func (m *SchemaAllTransformersRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaAllTransformersRequestBuilderGetRequestConfiguration)([]i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TransformerInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.CreateTransformerInfoFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TransformerInfoable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TransformerInfoable)
        }
    }
    return val, nil
}
// ToGetRequestInformation retrieve all transformers
// returns a *RequestInformation when successful
func (m *SchemaAllTransformersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaAllTransformersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaAllTransformersRequestBuilder when successful
func (m *SchemaAllTransformersRequestBuilder) WithUrl(rawUrl string)(*SchemaAllTransformersRequestBuilder) {
    return NewSchemaAllTransformersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

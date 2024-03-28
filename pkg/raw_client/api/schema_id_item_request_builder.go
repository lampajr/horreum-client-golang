package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// SchemaIdItemRequestBuilder builds and executes requests for operations under \api\schema\{id-id}
type SchemaIdItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaIdItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaIdItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SchemaIdItemRequestBuilderGetQueryParameters retrieve Schema by ID
type SchemaIdItemRequestBuilderGetQueryParameters struct {
    // API token for authorization
    Token *string `uriparametername:"token"`
}
// SchemaIdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaIdItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchemaIdItemRequestBuilderGetQueryParameters
}
// NewSchemaIdItemRequestBuilderInternal instantiates a new SchemaIdItemRequestBuilder and sets the default values.
func NewSchemaIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdItemRequestBuilder) {
    m := &SchemaIdItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/{id%2Did}{?token*}", pathParameters),
    }
    return m
}
// NewSchemaIdItemRequestBuilder instantiates a new SchemaIdItemRequestBuilder and sets the default values.
func NewSchemaIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaIdItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Schema by id
func (m *SchemaIdItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchemaIdItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DropToken the dropToken property
// returns a *SchemaItemDropTokenRequestBuilder when successful
func (m *SchemaIdItemRequestBuilder) DropToken()(*SchemaItemDropTokenRequestBuilder) {
    return NewSchemaItemDropTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Export the export property
// returns a *SchemaItemExportRequestBuilder when successful
func (m *SchemaIdItemRequestBuilder) Export()(*SchemaItemExportRequestBuilder) {
    return NewSchemaItemExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve Schema by ID
// returns a Schemaable when successful
func (m *SchemaIdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaIdItemRequestBuilderGetRequestConfiguration)(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.Schemaable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.CreateSchemaFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.Schemaable), nil
}
// Labels the labels property
// returns a *SchemaItemLabelsRequestBuilder when successful
func (m *SchemaIdItemRequestBuilder) Labels()(*SchemaItemLabelsRequestBuilder) {
    return NewSchemaItemLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetToken the resetToken property
// returns a *SchemaItemResetTokenRequestBuilder when successful
func (m *SchemaIdItemRequestBuilder) ResetToken()(*SchemaItemResetTokenRequestBuilder) {
    return NewSchemaItemResetTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a Schema by id
// returns a *RequestInformation when successful
func (m *SchemaIdItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SchemaIdItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/api/schema/{id%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve Schema by ID
// returns a *RequestInformation when successful
func (m *SchemaIdItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaIdItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Transformers the transformers property
// returns a *SchemaItemTransformersRequestBuilder when successful
func (m *SchemaIdItemRequestBuilder) Transformers()(*SchemaItemTransformersRequestBuilder) {
    return NewSchemaItemTransformersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateAccess the updateAccess property
// returns a *SchemaItemUpdateAccessRequestBuilder when successful
func (m *SchemaIdItemRequestBuilder) UpdateAccess()(*SchemaItemUpdateAccessRequestBuilder) {
    return NewSchemaItemUpdateAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaIdItemRequestBuilder when successful
func (m *SchemaIdItemRequestBuilder) WithUrl(rawUrl string)(*SchemaIdItemRequestBuilder) {
    return NewSchemaIdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

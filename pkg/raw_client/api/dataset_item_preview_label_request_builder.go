package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// DatasetItemPreviewLabelRequestBuilder builds and executes requests for operations under \api\dataset\{datasetId-id}\previewLabel
type DatasetItemPreviewLabelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DatasetItemPreviewLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasetItemPreviewLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDatasetItemPreviewLabelRequestBuilderInternal instantiates a new DatasetItemPreviewLabelRequestBuilder and sets the default values.
func NewDatasetItemPreviewLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetItemPreviewLabelRequestBuilder) {
    m := &DatasetItemPreviewLabelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset/{datasetId%2Did}/previewLabel", pathParameters),
    }
    return m
}
// NewDatasetItemPreviewLabelRequestBuilder instantiates a new DatasetItemPreviewLabelRequestBuilder and sets the default values.
func NewDatasetItemPreviewLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetItemPreviewLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetItemPreviewLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a LabelPreviewable when successful
func (m *DatasetItemPreviewLabelRequestBuilder) Post(ctx context.Context, body i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.Labelable, requestConfiguration *DatasetItemPreviewLabelRequestBuilderPostRequestConfiguration)(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.LabelPreviewable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.CreateLabelPreviewFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.LabelPreviewable), nil
}
// returns a *RequestInformation when successful
func (m *DatasetItemPreviewLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.Labelable, requestConfiguration *DatasetItemPreviewLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DatasetItemPreviewLabelRequestBuilder when successful
func (m *DatasetItemPreviewLabelRequestBuilder) WithUrl(rawUrl string)(*DatasetItemPreviewLabelRequestBuilder) {
    return NewDatasetItemPreviewLabelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

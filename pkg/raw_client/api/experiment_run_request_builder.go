package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// ExperimentRunRequestBuilder builds and executes requests for operations under \api\experiment\run
type ExperimentRunRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExperimentRunRequestBuilderGetQueryParameters run an experiment for a given dataset and experiment profile
type ExperimentRunRequestBuilderGetQueryParameters struct {
    // The dataset to run the experiment on
    DatasetId *int32 `uriparametername:"datasetId"`
}
// ExperimentRunRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExperimentRunRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExperimentRunRequestBuilderGetQueryParameters
}
// NewExperimentRunRequestBuilderInternal instantiates a new ExperimentRunRequestBuilder and sets the default values.
func NewExperimentRunRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentRunRequestBuilder) {
    m := &ExperimentRunRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/experiment/run{?datasetId*}", pathParameters),
    }
    return m
}
// NewExperimentRunRequestBuilder instantiates a new ExperimentRunRequestBuilder and sets the default values.
func NewExperimentRunRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentRunRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExperimentRunRequestBuilderInternal(urlParams, requestAdapter)
}
// Get run an experiment for a given dataset and experiment profile
// returns a []ExperimentResultable when successful
func (m *ExperimentRunRequestBuilder) Get(ctx context.Context, requestConfiguration *ExperimentRunRequestBuilderGetRequestConfiguration)([]i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.ExperimentResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.CreateExperimentResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.ExperimentResultable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.ExperimentResultable)
        }
    }
    return val, nil
}
// ToGetRequestInformation run an experiment for a given dataset and experiment profile
// returns a *RequestInformation when successful
func (m *ExperimentRunRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExperimentRunRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExperimentRunRequestBuilder when successful
func (m *ExperimentRunRequestBuilder) WithUrl(rawUrl string)(*ExperimentRunRequestBuilder) {
    return NewExperimentRunRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

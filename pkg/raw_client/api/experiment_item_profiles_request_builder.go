package api

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// ExperimentItemProfilesRequestBuilder builds and executes requests for operations under \api\experiment\{testId}\profiles
type ExperimentItemProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExperimentItemProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExperimentItemProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExperimentItemProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExperimentItemProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByProfileId gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.experiment.item.profiles.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ExperimentItemProfilesWithProfileItemRequestBuilder when successful
func (m *ExperimentItemProfilesRequestBuilder) ByProfileId(profileId string)(*ExperimentItemProfilesWithProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if profileId != "" {
        urlTplParams["profileId"] = profileId
    }
    return NewExperimentItemProfilesWithProfileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByProfileIdInteger gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.experiment.item.profiles.item collection
// returns a *ExperimentItemProfilesWithProfileItemRequestBuilder when successful
func (m *ExperimentItemProfilesRequestBuilder) ByProfileIdInteger(profileId int32)(*ExperimentItemProfilesWithProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["profileId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(profileId), 10)
    return NewExperimentItemProfilesWithProfileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewExperimentItemProfilesRequestBuilderInternal instantiates a new ExperimentItemProfilesRequestBuilder and sets the default values.
func NewExperimentItemProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentItemProfilesRequestBuilder) {
    m := &ExperimentItemProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/experiment/{testId}/profiles", pathParameters),
    }
    return m
}
// NewExperimentItemProfilesRequestBuilder instantiates a new ExperimentItemProfilesRequestBuilder and sets the default values.
func NewExperimentItemProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentItemProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExperimentItemProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve Experiment Profiles by Test ID
// returns a []ExperimentProfileable when successful
func (m *ExperimentItemProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *ExperimentItemProfilesRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExperimentProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateExperimentProfileFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExperimentProfileable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExperimentProfileable)
        }
    }
    return val, nil
}
// Post save new or update existing Experiment Profiles for a Test 
// returns a *int32 when successful
func (m *ExperimentItemProfilesRequestBuilder) Post(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExperimentProfileable, requestConfiguration *ExperimentItemProfilesRequestBuilderPostRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation retrieve Experiment Profiles by Test ID
// returns a *RequestInformation when successful
func (m *ExperimentItemProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExperimentItemProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation save new or update existing Experiment Profiles for a Test 
// returns a *RequestInformation when successful
func (m *ExperimentItemProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExperimentProfileable, requestConfiguration *ExperimentItemProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExperimentItemProfilesRequestBuilder when successful
func (m *ExperimentItemProfilesRequestBuilder) WithUrl(rawUrl string)(*ExperimentItemProfilesRequestBuilder) {
    return NewExperimentItemProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

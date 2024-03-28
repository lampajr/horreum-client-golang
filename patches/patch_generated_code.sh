#!/bin/bash

set -e

# Related to issue https://github.com/microsoft/kiota/issues/4380
echo "Patching pkg/raw_client/api/test_item_transformers_request_builder.go" 
sed -i 's/requestInfo.SetContentFromScalarCollection(ctx, m.BaseRequestBuilder.RequestAdapter, "application\/json", body)/cast := make([]interface{}, len(body))\n\tfor i, v := range body {\n\t\tcast[i] = v\n\t}\n\trequestInfo.SetContentFromScalarCollection(ctx, m.BaseRequestBuilder.RequestAdapter, "application\/json", cast)/g' ./pkg/raw_client/api/test_item_transformers_request_builder.go

module github.com/aws/aws-sdk-go-v2/service/dynamodb

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.19.1
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.36
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.30
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.30
	github.com/aws/smithy-go v1.13.5
	github.com/google/go-cmp v0.5.8
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../service/internal/endpoint-discovery/

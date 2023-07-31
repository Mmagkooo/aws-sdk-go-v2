// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the definition of the specified pipeline. You can call
// GetPipelineDefinition to retrieve the pipeline definition that you provided
// using PutPipelineDefinition . POST / HTTP/1.1 Content-Type:
// application/x-amz-json-1.1 X-Amz-Target: DataPipeline.GetPipelineDefinition
// Content-Length: 40 Host: datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon,
// 12 Nov 2012 17:49:52 GMT Authorization: AuthParams {"pipelineId":
// "df-06372391ZG65EXAMPLE"} x-amzn-RequestId: e28309e5-0776-11e2-8a14-21bb8a1f50ef
// Content-Type: application/x-amz-json-1.1 Content-Length: 890 Date: Mon, 12 Nov
// 2012 17:50:53 GMT {"pipelineObjects": [ {"fields": [ {"key": "workerGroup",
// "stringValue": "workerGroup"} ], "id": "Default", "name": "Default"}, {"fields":
// [ {"key": "startDateTime", "stringValue": "2012-09-25T17:00:00"}, {"key":
// "type", "stringValue": "Schedule"}, {"key": "period", "stringValue": "1 hour"},
// {"key": "endDateTime", "stringValue": "2012-09-25T18:00:00"} ], "id":
// "Schedule", "name": "Schedule"}, {"fields": [ {"key": "schedule", "refValue":
// "Schedule"}, {"key": "command", "stringValue": "echo hello"}, {"key": "parent",
// "refValue": "Default"}, {"key": "type", "stringValue": "ShellCommandActivity"}
// ], "id": "SayHello", "name": "SayHello"} ] }
func (c *Client) GetPipelineDefinition(ctx context.Context, params *GetPipelineDefinitionInput, optFns ...func(*Options)) (*GetPipelineDefinitionOutput, error) {
	if params == nil {
		params = &GetPipelineDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPipelineDefinition", params, optFns, c.addOperationGetPipelineDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPipelineDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for GetPipelineDefinition.
type GetPipelineDefinitionInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// The version of the pipeline definition to retrieve. Set this parameter to latest
	// (default) to use the last definition saved to the pipeline or active to use the
	// last definition that was activated.
	Version *string

	noSmithyDocumentSerde
}

// Contains the output of GetPipelineDefinition.
type GetPipelineDefinitionOutput struct {

	// The parameter objects used in the pipeline definition.
	ParameterObjects []types.ParameterObject

	// The parameter values used in the pipeline definition.
	ParameterValues []types.ParameterValue

	// The objects defined in the pipeline.
	PipelineObjects []types.PipelineObject

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPipelineDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPipelineDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPipelineDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addGetPipelineDefinitionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetPipelineDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPipelineDefinition(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetPipelineDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "GetPipelineDefinition",
	}
}

type opGetPipelineDefinitionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetPipelineDefinitionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetPipelineDefinitionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "datapipeline"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "datapipeline"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("datapipeline")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addGetPipelineDefinitionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetPipelineDefinitionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

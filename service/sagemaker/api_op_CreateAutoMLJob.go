// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Autopilot job also referred to as Autopilot experiment or AutoML
// job. We recommend using the new versions CreateAutoMLJobV2 (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html)
// and DescribeAutoMLJobV2 (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJobV2.html)
// , which offer backward compatibility. CreateAutoMLJobV2 can manage tabular
// problem types identical to those of its previous version CreateAutoMLJob , as
// well as non-tabular problem types such as image or text classification. Find
// guidelines about how to migrate a CreateAutoMLJob to CreateAutoMLJobV2 in
// Migrate a CreateAutoMLJob to CreateAutoMLJobV2 (https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development-create-experiment-api.html#autopilot-create-experiment-api-migrate-v1-v2)
// . You can find the best-performing model after you run an AutoML job by calling
// DescribeAutoMLJobV2 (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJobV2.html)
// (recommended) or DescribeAutoMLJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJob.html)
// .
func (c *Client) CreateAutoMLJob(ctx context.Context, params *CreateAutoMLJobInput, optFns ...func(*Options)) (*CreateAutoMLJobOutput, error) {
	if params == nil {
		params = &CreateAutoMLJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoMLJob", params, optFns, c.addOperationCreateAutoMLJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoMLJobInput struct {

	// Identifies an Autopilot job. The name must be unique to your account and is
	// case insensitive.
	//
	// This member is required.
	AutoMLJobName *string

	// An array of channel objects that describes the input data and its location.
	// Each channel is a named input source. Similar to InputDataConfig supported by
	// HyperParameterTrainingJobDefinition (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html)
	// . Format(s) supported: CSV, Parquet. A minimum of 500 rows is required for the
	// training dataset. There is not a minimum number of rows required for the
	// validation dataset.
	//
	// This member is required.
	InputDataConfig []types.AutoMLChannel

	// Provides information about encryption and the Amazon S3 output path needed to
	// store artifacts from an AutoML job. Format(s) supported: CSV.
	//
	// This member is required.
	OutputDataConfig *types.AutoMLOutputDataConfig

	// The ARN of the role that is used to access the data.
	//
	// This member is required.
	RoleArn *string

	// A collection of settings used to configure an AutoML job.
	AutoMLJobConfig *types.AutoMLJobConfig

	// Specifies a metric to minimize or maximize as the objective of a job. If not
	// specified, the default objective metric depends on the problem type. See
	// AutoMLJobObjective (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AutoMLJobObjective.html)
	// for the default values.
	AutoMLJobObjective *types.AutoMLJobObjective

	// Generates possible candidates without training the models. A candidate is a
	// combination of data preprocessors, algorithms, and algorithm parameter settings.
	GenerateCandidateDefinitionsOnly bool

	// Specifies how to generate the endpoint name for an automatic one-click
	// Autopilot model deployment.
	ModelDeployConfig *types.ModelDeployConfig

	// Defines the type of supervised learning problem available for the candidates.
	// For more information, see Amazon SageMaker Autopilot problem types (https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-datasets-problem-types.html#autopilot-problem-types)
	// .
	ProblemType types.ProblemType

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web ServicesResources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// . Tag keys must be unique per resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAutoMLJobOutput struct {

	// The unique ARN assigned to the AutoML job when it is created.
	//
	// This member is required.
	AutoMLJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAutoMLJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAutoMLJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAutoMLJob{}, middleware.After)
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
	if err = addCreateAutoMLJobResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAutoMLJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoMLJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAutoMLJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateAutoMLJob",
	}
}

type opCreateAutoMLJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateAutoMLJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateAutoMLJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "sagemaker"
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
				signingName = "sagemaker"
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
				v4aScheme.SigningName = aws.String("sagemaker")
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

func addCreateAutoMLJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateAutoMLJobResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

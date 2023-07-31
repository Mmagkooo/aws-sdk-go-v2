// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an image builder. An image builder is a virtual machine that is used to
// create an image. The initial state of the builder is PENDING . When it is ready,
// the state is RUNNING .
func (c *Client) CreateImageBuilder(ctx context.Context, params *CreateImageBuilderInput, optFns ...func(*Options)) (*CreateImageBuilderOutput, error) {
	if params == nil {
		params = &CreateImageBuilderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateImageBuilder", params, optFns, c.addOperationCreateImageBuilderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateImageBuilderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateImageBuilderInput struct {

	// The instance type to use when launching the image builder. The following
	// instance types are available:
	//   - stream.standard.small
	//   - stream.standard.medium
	//   - stream.standard.large
	//   - stream.compute.large
	//   - stream.compute.xlarge
	//   - stream.compute.2xlarge
	//   - stream.compute.4xlarge
	//   - stream.compute.8xlarge
	//   - stream.memory.large
	//   - stream.memory.xlarge
	//   - stream.memory.2xlarge
	//   - stream.memory.4xlarge
	//   - stream.memory.8xlarge
	//   - stream.memory.z1d.large
	//   - stream.memory.z1d.xlarge
	//   - stream.memory.z1d.2xlarge
	//   - stream.memory.z1d.3xlarge
	//   - stream.memory.z1d.6xlarge
	//   - stream.memory.z1d.12xlarge
	//   - stream.graphics-design.large
	//   - stream.graphics-design.xlarge
	//   - stream.graphics-design.2xlarge
	//   - stream.graphics-design.4xlarge
	//   - stream.graphics-desktop.2xlarge
	//   - stream.graphics.g4dn.xlarge
	//   - stream.graphics.g4dn.2xlarge
	//   - stream.graphics.g4dn.4xlarge
	//   - stream.graphics.g4dn.8xlarge
	//   - stream.graphics.g4dn.12xlarge
	//   - stream.graphics.g4dn.16xlarge
	//   - stream.graphics-pro.4xlarge
	//   - stream.graphics-pro.8xlarge
	//   - stream.graphics-pro.16xlarge
	//
	// This member is required.
	InstanceType *string

	// A unique name for the image builder.
	//
	// This member is required.
	Name *string

	// The list of interface VPC endpoint (interface endpoint) objects. Administrators
	// can connect to the image builder only through the specified endpoints.
	AccessEndpoints []types.AccessEndpoint

	// The version of the AppStream 2.0 agent to use for this image builder. To use
	// the latest version of the AppStream 2.0 agent, specify [LATEST].
	AppstreamAgentVersion *string

	// The description to display.
	Description *string

	// The image builder name to display.
	DisplayName *string

	// The name of the directory and organizational unit (OU) to use to join the image
	// builder to a Microsoft Active Directory domain.
	DomainJoinInfo *types.DomainJoinInfo

	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool

	// The Amazon Resource Name (ARN) of the IAM role to apply to the image builder.
	// To assume a role, the image builder calls the AWS Security Token Service (STS)
	// AssumeRole API operation and passes the ARN of the role to use. The operation
	// creates a new session with temporary credentials. AppStream 2.0 retrieves the
	// temporary credentials and creates the appstream_machine_role credential profile
	// on the instance. For more information, see Using an IAM Role to Grant
	// Permissions to Applications and Scripts Running on AppStream 2.0 Streaming
	// Instances (https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	IamRoleArn *string

	// The ARN of the public, private, or shared image to use.
	ImageArn *string

	// The name of the image used to create the image builder.
	ImageName *string

	// The tags to associate with the image builder. A tag is a key-value pair, and
	// the value is optional. For example, Environment=Test. If you do not specify a
	// value, Environment=. Generally allowed characters are: letters, numbers, and
	// spaces representable in UTF-8, and the following special characters: _ . : / = +
	// \ - @ If you do not specify a value, the value is set to an empty string. For
	// more information about tags, see Tagging Your Resources (https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	Tags map[string]string

	// The VPC configuration for the image builder. You can specify only one subnet.
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateImageBuilderOutput struct {

	// Information about the image builder.
	ImageBuilder *types.ImageBuilder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateImageBuilderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateImageBuilder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateImageBuilder{}, middleware.After)
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
	if err = addCreateImageBuilderResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateImageBuilderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateImageBuilder(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateImageBuilder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateImageBuilder",
	}
}

type opCreateImageBuilderResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateImageBuilderResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateImageBuilderResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "appstream"
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
				signingName = "appstream"
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
				v4aScheme.SigningName = aws.String("appstream")
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

func addCreateImageBuilderResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateImageBuilderResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon ECS cluster. By default, your account receives a default
// cluster when you launch your first container instance. However, you can create
// your own cluster with a unique name with the CreateCluster action. When you
// call the CreateCluster API operation, Amazon ECS attempts to create the Amazon
// ECS service-linked role for your account. This is so that it can manage required
// resources in other Amazon Web Services services on your behalf. However, if the
// user that makes the call doesn't have permissions to create the service-linked
// role, it isn't created. For more information, see Using service-linked roles
// for Amazon ECS (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, c.addOperationCreateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// The short name of one or more capacity providers to associate with the cluster.
	// A capacity provider must be associated with a cluster before it can be included
	// as part of the default capacity provider strategy of the cluster or used in a
	// capacity provider strategy when calling the CreateService (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html)
	// or RunTask (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html)
	// actions. If specifying a capacity provider that uses an Auto Scaling group, the
	// capacity provider must be created but not associated with another cluster. New
	// Auto Scaling group capacity providers can be created with the
	// CreateCapacityProvider (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProvider.html)
	// API operation. To use a Fargate capacity provider, specify either the FARGATE
	// or FARGATE_SPOT capacity providers. The Fargate capacity providers are
	// available to all accounts and only need to be associated with a cluster to be
	// used. The PutCapacityProvider (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutCapacityProvider.html)
	// API operation is used to update the list of available capacity providers for a
	// cluster after the cluster is created.
	CapacityProviders []string

	// The name of your cluster. If you don't specify a name for your cluster, you
	// create a cluster that's named default . Up to 255 letters (uppercase and
	// lowercase), numbers, underscores, and hyphens are allowed.
	ClusterName *string

	// The execute command configuration for the cluster.
	Configuration *types.ClusterConfiguration

	// The capacity provider strategy to set as the default for the cluster. After a
	// default capacity provider strategy is set for a cluster, when you call the
	// CreateService (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html)
	// or RunTask (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html)
	// APIs with no capacity provider strategy or launch type specified, the default
	// capacity provider strategy for the cluster is used. If a default capacity
	// provider strategy isn't defined for a cluster when it was created, it can be
	// defined later with the PutClusterCapacityProviders API operation.
	DefaultCapacityProviderStrategy []types.CapacityProviderStrategyItem

	// Use this parameter to set a default Service Connect namespace. After you set a
	// default Service Connect namespace, any new services with Service Connect turned
	// on that are created in the cluster are added as client services in the
	// namespace. This setting only applies to new services that set the enabled
	// parameter to true in the ServiceConnectConfiguration . You can set the namespace
	// of each service individually in the ServiceConnectConfiguration to override
	// this default parameter. Tasks that run in a namespace can use short names to
	// connect to services in the namespace. Tasks can connect to services across all
	// of the clusters in the namespace. Tasks connect through a managed proxy
	// container that collects logs and metrics for increased visibility. Only the
	// tasks that Amazon ECS services create are supported with Service Connect. For
	// more information, see Service Connect (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html)
	// in the Amazon Elastic Container Service Developer Guide.
	ServiceConnectDefaults *types.ClusterServiceConnectDefaultsRequest

	// The setting to use when creating a cluster. This parameter is used to turn on
	// CloudWatch Container Insights for a cluster. If this value is specified, it
	// overrides the containerInsights value set with PutAccountSetting or
	// PutAccountSettingDefault .
	Settings []types.ClusterSetting

	// The metadata that you apply to the cluster to help you categorize and organize
	// them. Each tag consists of a key and an optional value. You define both. The
	// following basic restrictions apply to tags:
	//   - Maximum number of tags per resource - 50
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//   - Maximum key length - 128 Unicode characters in UTF-8
	//   - Maximum value length - 256 Unicode characters in UTF-8
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//   - Tag keys and values are case-sensitive.
	//   - Do not use aws: , AWS: , or any upper or lowercase combination of such as a
	//   prefix for either keys or values as it is reserved for Amazon Web Services use.
	//   You cannot edit or delete tag keys or values with this prefix. Tags with this
	//   prefix do not count against your tags per resource limit.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateClusterOutput struct {

	// The full description of your new cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCluster{}, middleware.After)
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
	if err = addCreateClusterResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "CreateCluster",
	}
}

type opCreateClusterResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateClusterResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateClusterResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ecs"
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
				signingName = "ecs"
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
				v4aScheme.SigningName = aws.String("ecs")
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

func addCreateClusterResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateClusterResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

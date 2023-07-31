// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Batch compute environment. You can create MANAGED or UNMANAGED
// compute environments. MANAGED compute environments can use Amazon EC2 or
// Fargate resources. UNMANAGED compute environments can only use EC2 resources.
// In a managed compute environment, Batch manages the capacity and instance types
// of the compute resources within the environment. This is based on the compute
// resource specification that you define or the launch template (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html)
// that you specify when you create the compute environment. Either, you can choose
// to use EC2 On-Demand Instances and EC2 Spot Instances. Or, you can use Fargate
// and Fargate Spot capacity in your managed compute environment. You can
// optionally set a maximum price so that Spot Instances only launch when the Spot
// Instance price is less than a specified percentage of the On-Demand price.
// Multi-node parallel jobs aren't supported on Spot Instances. In an unmanaged
// compute environment, you can manage your own EC2 compute resources and have
// flexibility with how you configure your compute resources. For example, you can
// use custom AMIs. However, you must verify that each of your AMIs meet the Amazon
// ECS container instance AMI specification. For more information, see container
// instance AMIs (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container_instance_AMIs.html)
// in the Amazon Elastic Container Service Developer Guide. After you created your
// unmanaged compute environment, you can use the DescribeComputeEnvironments
// operation to find the Amazon ECS cluster that's associated with it. Then, launch
// your container instances into that Amazon ECS cluster. For more information, see
// Launching an Amazon ECS container instance (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html)
// in the Amazon Elastic Container Service Developer Guide. To create a compute
// environment that uses EKS resources, the caller must have permissions to call
// eks:DescribeCluster . Batch doesn't automatically upgrade the AMIs in a compute
// environment after it's created. For example, it also doesn't update the AMIs in
// your compute environment when a newer version of the Amazon ECS optimized AMI is
// available. You're responsible for the management of the guest operating system.
// This includes any updates and security patches. You're also responsible for any
// additional application software or utilities that you install on the compute
// resources. There are two ways to use a new AMI for your Batch jobs. The original
// method is to complete these steps:
//   - Create a new compute environment with the new AMI.
//   - Add the compute environment to an existing job queue.
//   - Remove the earlier compute environment from your job queue.
//   - Delete the earlier compute environment.
//
// In April 2022, Batch added enhanced support for updating compute environments.
// For more information, see Updating compute environments (https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html)
// . To use the enhanced updating of compute environments to update AMIs, follow
// these rules:
//   - Either don't set the service role ( serviceRole ) parameter or set it to the
//     AWSBatchServiceRole service-linked role.
//   - Set the allocation strategy ( allocationStrategy ) parameter to
//     BEST_FIT_PROGRESSIVE or SPOT_CAPACITY_OPTIMIZED .
//   - Set the update to latest image version ( updateToLatestImageVersion )
//     parameter to true . The updateToLatestImageVersion parameter is used when you
//     update a compute environment. This parameter is ignored when you create a
//     compute environment.
//   - Don't specify an AMI ID in imageId , imageIdOverride (in ec2Configuration (https://docs.aws.amazon.com/batch/latest/APIReference/API_Ec2Configuration.html)
//     ), or in the launch template ( launchTemplate ). In that case, Batch selects
//     the latest Amazon ECS optimized AMI that's supported by Batch at the time the
//     infrastructure update is initiated. Alternatively, you can specify the AMI ID in
//     the imageId or imageIdOverride parameters, or the launch template identified
//     by the LaunchTemplate properties. Changing any of these properties starts an
//     infrastructure update. If the AMI ID is specified in the launch template, it
//     can't be replaced by specifying an AMI ID in either the imageId or
//     imageIdOverride parameters. It can only be replaced by specifying a different
//     launch template, or if the launch template version is set to $Default or
//     $Latest , by setting either a new default version for the launch template (if
//     $Default ) or by adding a new version to the launch template (if $Latest ).
//
// If these rules are followed, any update that starts an infrastructure update
// causes the AMI ID to be re-selected. If the version setting in the launch
// template ( launchTemplate ) is set to $Latest or $Default , the latest or
// default version of the launch template is evaluated up at the time of the
// infrastructure update, even if the launchTemplate wasn't updated.
func (c *Client) CreateComputeEnvironment(ctx context.Context, params *CreateComputeEnvironmentInput, optFns ...func(*Options)) (*CreateComputeEnvironmentOutput, error) {
	if params == nil {
		params = &CreateComputeEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComputeEnvironment", params, optFns, c.addOperationCreateComputeEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComputeEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateComputeEnvironment .
type CreateComputeEnvironmentInput struct {

	// The name for your compute environment. It can be up to 128 characters long. It
	// can contain uppercase and lowercase letters, numbers, hyphens (-), and
	// underscores (_).
	//
	// This member is required.
	ComputeEnvironmentName *string

	// The type of the compute environment: MANAGED or UNMANAGED . For more
	// information, see Compute Environments (https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
	// in the Batch User Guide.
	//
	// This member is required.
	Type types.CEType

	// Details about the compute resources managed by the compute environment. This
	// parameter is required for managed compute environments. For more information,
	// see Compute Environments (https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
	// in the Batch User Guide.
	ComputeResources *types.ComputeResource

	// The details for the Amazon EKS cluster that supports the compute environment.
	EksConfiguration *types.EksConfiguration

	// The full Amazon Resource Name (ARN) of the IAM role that allows Batch to make
	// calls to other Amazon Web Services services on your behalf. For more
	// information, see Batch service IAM role (https://docs.aws.amazon.com/batch/latest/userguide/service_IAM_role.html)
	// in the Batch User Guide. If your account already created the Batch
	// service-linked role, that role is used by default for your compute environment
	// unless you specify a different role here. If the Batch service-linked role
	// doesn't exist in your account, and no role is specified here, the service
	// attempts to create the Batch service-linked role in your account. If your
	// specified role has a path other than / , then you must specify either the full
	// role ARN (recommended) or prefix the role name with the path. For example, if a
	// role with the name bar has a path of /foo/ , specify /foo/bar as the role name.
	// For more information, see Friendly names and paths (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names)
	// in the IAM User Guide. Depending on how you created your Batch service role, its
	// ARN might contain the service-role path prefix. When you only specify the name
	// of the service role, Batch assumes that your ARN doesn't use the service-role
	// path prefix. Because of this, we recommend that you specify the full ARN of your
	// service role when you create compute environments.
	ServiceRole *string

	// The state of the compute environment. If the state is ENABLED , then the compute
	// environment accepts jobs from a queue and can scale out automatically based on
	// queues. If the state is ENABLED , then the Batch scheduler can attempt to place
	// jobs from an associated job queue on the compute resources within the
	// environment. If the compute environment is managed, then it can scale its
	// instances out or in automatically, based on the job queue demand. If the state
	// is DISABLED , then the Batch scheduler doesn't attempt to place jobs within the
	// environment. Jobs in a STARTING or RUNNING state continue to progress normally.
	// Managed compute environments in the DISABLED state don't scale out. Compute
	// environments in a DISABLED state may continue to incur billing charges. To
	// prevent additional charges, turn off and then delete the compute environment.
	// For more information, see State (https://docs.aws.amazon.com/batch/latest/userguide/compute_environment_parameters.html#compute_environment_state)
	// in the Batch User Guide. When an instance is idle, the instance scales down to
	// the minvCpus value. However, the instance size doesn't change. For example,
	// consider a c5.8xlarge instance with a minvCpus value of 4 and a desiredvCpus
	// value of 36 . This instance doesn't scale down to a c5.large instance.
	State types.CEState

	// The tags that you apply to the compute environment to help you categorize and
	// organize your resources. Each tag consists of a key and an optional value. For
	// more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in Amazon Web Services General Reference. These tags can be updated or removed
	// using the TagResource (https://docs.aws.amazon.com/batch/latest/APIReference/API_TagResource.html)
	// and UntagResource (https://docs.aws.amazon.com/batch/latest/APIReference/API_UntagResource.html)
	// API operations. These tags don't propagate to the underlying compute resources.
	Tags map[string]string

	// The maximum number of vCPUs for an unmanaged compute environment. This
	// parameter is only used for fair share scheduling to reserve vCPU capacity for
	// new share identifiers. If this parameter isn't provided for a fair share job
	// queue, no vCPU capacity is reserved. This parameter is only supported when the
	// type parameter is set to UNMANAGED .
	UnmanagedvCpus *int32

	noSmithyDocumentSerde
}

type CreateComputeEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) of the compute environment.
	ComputeEnvironmentArn *string

	// The name of the compute environment. It can be up to 128 characters long. It
	// can contain uppercase and lowercase letters, numbers, hyphens (-), and
	// underscores (_).
	ComputeEnvironmentName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateComputeEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateComputeEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComputeEnvironment{}, middleware.After)
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
	if err = addCreateComputeEnvironmentResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateComputeEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComputeEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateComputeEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CreateComputeEnvironment",
	}
}

type opCreateComputeEnvironmentResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateComputeEnvironmentResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateComputeEnvironmentResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "batch"
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
				signingName = "batch"
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
				v4aScheme.SigningName = aws.String("batch")
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

func addCreateComputeEnvironmentResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateComputeEnvironmentResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

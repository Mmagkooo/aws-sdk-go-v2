// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a resource from the Resilience Hub application.
//
// * You can only delete a
// manually added resource. To exclude non-manually added resources, use the
// UpdateAppVersionResource API.
//
// * This action has no effect outside Resilience
// Hub.
//
// * This API updates the Resilience Hub application draft version. To use
// this resource for running resiliency assessments, you must publish the
// Resilience Hub application using the PublishAppVersion API.
func (c *Client) DeleteAppVersionResource(ctx context.Context, params *DeleteAppVersionResourceInput, optFns ...func(*Options)) (*DeleteAppVersionResourceOutput, error) {
	if params == nil {
		params = &DeleteAppVersionResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAppVersionResource", params, optFns, c.addOperationDeleteAppVersionResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAppVersionResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAppVersionResourceInput struct {

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The Amazon Web Services account that owns the physical resource.
	AwsAccountId *string

	// The Amazon Web Services region that owns the physical resource.
	AwsRegion *string

	// Used for an idempotency token. A client token is a unique, case-sensitive string
	// of up to 64 ASCII characters. You should not reuse the same client token for
	// other API requests.
	ClientToken *string

	// The logical identifier of the resource.
	LogicalResourceId *types.LogicalResourceId

	// The physical identifier of the resource.
	PhysicalResourceId *string

	// The name of the resource.
	ResourceName *string

	noSmithyDocumentSerde
}

type DeleteAppVersionResourceOutput struct {

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The AWS Resilience Hub application version.
	//
	// This member is required.
	AppVersion *string

	// Defines a physical resource. A physical resource is a resource that exists in
	// your account. It can be identified using an Amazon Resource Name (ARN) or a
	// Resilience Hub-native identifier.
	PhysicalResource *types.PhysicalResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAppVersionResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAppVersionResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAppVersionResource{}, middleware.After)
	if err != nil {
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opDeleteAppVersionResourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteAppVersionResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAppVersionResource(options.Region), middleware.Before); err != nil {
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
	return nil
}

type idempotencyToken_initializeOpDeleteAppVersionResource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteAppVersionResource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteAppVersionResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteAppVersionResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteAppVersionResourceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteAppVersionResourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteAppVersionResource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteAppVersionResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "DeleteAppVersionResource",
	}
}

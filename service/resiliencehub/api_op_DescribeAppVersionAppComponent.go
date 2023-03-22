// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes an Application Component in the Resilience Hub application.
func (c *Client) DescribeAppVersionAppComponent(ctx context.Context, params *DescribeAppVersionAppComponentInput, optFns ...func(*Options)) (*DescribeAppVersionAppComponentOutput, error) {
	if params == nil {
		params = &DescribeAppVersionAppComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppVersionAppComponent", params, optFns, c.addOperationDescribeAppVersionAppComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppVersionAppComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppVersionAppComponentInput struct {

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

	// The identifier of the Application Component.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DescribeAppVersionAppComponentOutput struct {

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

	// The list of Application Components that belong to this resource.
	AppComponent *types.AppComponent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppVersionAppComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAppVersionAppComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAppVersionAppComponent{}, middleware.After)
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
	if err = addOpDescribeAppVersionAppComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppVersionAppComponent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAppVersionAppComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "DescribeAppVersionAppComponent",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package medicalimaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medicalimaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copy an image set.
func (c *Client) CopyImageSet(ctx context.Context, params *CopyImageSetInput, optFns ...func(*Options)) (*CopyImageSetOutput, error) {
	if params == nil {
		params = &CopyImageSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyImageSet", params, optFns, c.addOperationCopyImageSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyImageSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyImageSetInput struct {

	// Copy image set information.
	//
	// This member is required.
	CopyImageSetInformation *types.CopyImageSetInformation

	// The data store identifier.
	//
	// This member is required.
	DatastoreId *string

	// The source image set identifier.
	//
	// This member is required.
	SourceImageSetId *string

	// Setting this flag will force the CopyImageSet operation, even if Patient,
	// Study, or Series level metadata are mismatched across the sourceImageSet and
	// destinationImageSet .
	Force *bool

	noSmithyDocumentSerde
}

type CopyImageSetOutput struct {

	// The data store identifier.
	//
	// This member is required.
	DatastoreId *string

	// The properties of the destination image set.
	//
	// This member is required.
	DestinationImageSetProperties *types.CopyDestinationImageSetProperties

	// The properties of the source image set.
	//
	// This member is required.
	SourceImageSetProperties *types.CopySourceImageSetProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyImageSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCopyImageSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCopyImageSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyImageSet"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addEndpointPrefix_opCopyImageSetMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCopyImageSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyImageSet(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opCopyImageSetMiddleware struct {
}

func (*endpointPrefix_opCopyImageSetMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCopyImageSetMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "runtime-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCopyImageSetMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCopyImageSetMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCopyImageSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyImageSet",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates phone number product types, calling names, or phone number names. You
// can update one attribute at a time for each UpdatePhoneNumberRequestItem . For
// example, you can update the product type, the calling name, or phone name. You
// cannot have a duplicate phoneNumberId in a request.
func (c *Client) BatchUpdatePhoneNumber(ctx context.Context, params *BatchUpdatePhoneNumberInput, optFns ...func(*Options)) (*BatchUpdatePhoneNumberOutput, error) {
	if params == nil {
		params = &BatchUpdatePhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdatePhoneNumber", params, optFns, c.addOperationBatchUpdatePhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdatePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdatePhoneNumberInput struct {

	// Lists the phone numbers in the update request.
	//
	// This member is required.
	UpdatePhoneNumberRequestItems []types.UpdatePhoneNumberRequestItem

	noSmithyDocumentSerde
}

type BatchUpdatePhoneNumberOutput struct {

	// A list of failed phone numbers and their error messages.
	PhoneNumberErrors []types.PhoneNumberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdatePhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdatePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdatePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchUpdatePhoneNumber"); err != nil {
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
	if err = addOpBatchUpdatePhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdatePhoneNumber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdatePhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchUpdatePhoneNumber",
	}
}

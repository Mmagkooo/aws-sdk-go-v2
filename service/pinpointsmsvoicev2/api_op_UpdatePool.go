// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the configuration of an existing pool. You can update the opt-out list,
// enable or disable two-way messaging, change the TwoWayChannelArn , enable or
// disable self-managed opt-outs, enable or disable deletion protection, and enable
// or disable shared routes.
func (c *Client) UpdatePool(ctx context.Context, params *UpdatePoolInput, optFns ...func(*Options)) (*UpdatePoolOutput, error) {
	if params == nil {
		params = &UpdatePoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePool", params, optFns, c.addOperationUpdatePoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePoolInput struct {

	// The unique identifier of the pool to update. Valid values are either the PoolId
	// or PoolArn.
	//
	// This member is required.
	PoolId *string

	// When set to true the pool can't be deleted.
	DeletionProtectionEnabled *bool

	// The OptOutList to associate with the pool. Valid values are either
	// OptOutListName or OptOutListArn.
	OptOutListName *string

	// By default this is set to false. When an end recipient sends a message that
	// begins with HELP or STOP to one of your dedicated numbers, AWS End User
	// Messaging SMS and Voice automatically replies with a customizable message and
	// adds the end recipient to the OptOutList. When set to true you're responsible
	// for responding to HELP and STOP requests. You're also responsible for tracking
	// and honoring opt-out requests.
	SelfManagedOptOutsEnabled *bool

	// Indicates whether shared routes are enabled for the pool.
	SharedRoutesEnabled *bool

	// The Amazon Resource Name (ARN) of the two way channel.
	TwoWayChannelArn *string

	// An optional IAM Role Arn for a service to assume, to be able to post inbound
	// SMS messages.
	TwoWayChannelRole *string

	// By default this is set to false. When set to true you can receive incoming text
	// messages from your end recipients.
	TwoWayEnabled *bool

	noSmithyDocumentSerde
}

type UpdatePoolOutput struct {

	// The time when the pool was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	CreatedTimestamp *time.Time

	// When set to true the pool can't be deleted.
	DeletionProtectionEnabled bool

	// The type of message for the pool to use.
	MessageType types.MessageType

	// The name of the OptOutList associated with the pool.
	OptOutListName *string

	// The ARN of the pool.
	PoolArn *string

	// The unique identifier of the pool.
	PoolId *string

	// When an end recipient sends a message that begins with HELP or STOP to one of
	// your dedicated numbers, AWS End User Messaging SMS and Voice automatically
	// replies with a customizable message and adds the end recipient to the
	// OptOutList. When set to true you're responsible for responding to HELP and STOP
	// requests. You're also responsible for tracking and honoring opt-out requests.
	SelfManagedOptOutsEnabled bool

	// Indicates whether shared routes are enabled for the pool.
	SharedRoutesEnabled bool

	// The current status of the pool update request.
	Status types.PoolStatus

	// The Amazon Resource Name (ARN) of the two way channel.
	TwoWayChannelArn *string

	// An optional IAM Role Arn for a service to assume, to be able to post inbound
	// SMS messages.
	TwoWayChannelRole *string

	// By default this is set to false. When set to true you can receive incoming text
	// messages from your end recipients.
	TwoWayEnabled bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdatePool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdatePool{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePool"); err != nil {
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
	if err = addOpUpdatePoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePool",
	}
}

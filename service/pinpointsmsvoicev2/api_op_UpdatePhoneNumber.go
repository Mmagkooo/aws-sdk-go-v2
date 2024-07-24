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

// Updates the configuration of an existing origination phone number. You can
// update the opt-out list, enable or disable two-way messaging, change the
// TwoWayChannelArn, enable or disable self-managed opt-outs, and enable or disable
// deletion protection.
//
// If the origination phone number is associated with a pool, an error is returned.
func (c *Client) UpdatePhoneNumber(ctx context.Context, params *UpdatePhoneNumberInput, optFns ...func(*Options)) (*UpdatePhoneNumberOutput, error) {
	if params == nil {
		params = &UpdatePhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePhoneNumber", params, optFns, c.addOperationUpdatePhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePhoneNumberInput struct {

	// The unique identifier of the phone number. Valid values for this field can be
	// either the PhoneNumberId or PhoneNumberArn.
	//
	// This member is required.
	PhoneNumberId *string

	// By default this is set to false. When set to true the phone number can't be
	// deleted.
	DeletionProtectionEnabled *bool

	// The OptOutList to add the phone number to. Valid values for this field can be
	// either the OutOutListName or OutOutListArn.
	OptOutListName *string

	// By default this is set to false. When an end recipient sends a message that
	// begins with HELP or STOP to one of your dedicated numbers, AWS End User
	// Messaging SMS and Voice automatically replies with a customizable message and
	// adds the end recipient to the OptOutList. When set to true you're responsible
	// for responding to HELP and STOP requests. You're also responsible for tracking
	// and honoring opt-out requests.
	SelfManagedOptOutsEnabled *bool

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

type UpdatePhoneNumberOutput struct {

	// The time when the phone number was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	CreatedTimestamp *time.Time

	// When set to true the phone number can't be deleted.
	DeletionProtectionEnabled bool

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or
	// region.
	IsoCountryCode *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	MessageType types.MessageType

	// The monthly leasing price of the phone number, in US dollars.
	MonthlyLeasingPrice *string

	// Specifies if the number could be used for text messages, voice or both.
	NumberCapabilities []types.NumberCapability

	// The type of number that was requested.
	NumberType types.NumberType

	// The name of the OptOutList associated with the phone number.
	OptOutListName *string

	// The phone number that was updated.
	PhoneNumber *string

	// The Amazon Resource Name (ARN) of the updated phone number.
	PhoneNumberArn *string

	// The unique identifier of the phone number.
	PhoneNumberId *string

	// The unique identifier for the registration.
	RegistrationId *string

	// This is true if self managed opt-out are enabled.
	SelfManagedOptOutsEnabled bool

	// The current status of the request.
	Status types.NumberStatus

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

func (c *Client) addOperationUpdatePhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdatePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdatePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePhoneNumber"); err != nil {
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
	if err = addOpUpdatePhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePhoneNumber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePhoneNumber",
	}
}

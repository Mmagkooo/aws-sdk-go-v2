// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a subnet CIDR reservation. For more information, see [Subnet CIDR reservations] in the Amazon VPC
// User Guide and [Assign prefixes to network interfaces]in the Amazon EC2 User Guide.
//
// [Subnet CIDR reservations]: https://docs.aws.amazon.com/vpc/latest/userguide/subnet-cidr-reservation.html
// [Assign prefixes to network interfaces]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html
func (c *Client) CreateSubnetCidrReservation(ctx context.Context, params *CreateSubnetCidrReservationInput, optFns ...func(*Options)) (*CreateSubnetCidrReservationOutput, error) {
	if params == nil {
		params = &CreateSubnetCidrReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubnetCidrReservation", params, optFns, c.addOperationCreateSubnetCidrReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubnetCidrReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubnetCidrReservationInput struct {

	// The IPv4 or IPV6 CIDR range to reserve.
	//
	// This member is required.
	Cidr *string

	// The type of reservation. The reservation type determines how the reserved IP
	// addresses are assigned to resources.
	//
	//   - prefix - Amazon Web Services assigns the reserved IP addresses to network
	//   interfaces.
	//
	//   - explicit - You assign the reserved IP addresses to network interfaces.
	//
	// This member is required.
	ReservationType types.SubnetCidrReservationType

	// The ID of the subnet.
	//
	// This member is required.
	SubnetId *string

	// The description to assign to the subnet CIDR reservation.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The tags to assign to the subnet CIDR reservation.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateSubnetCidrReservationOutput struct {

	// Information about the created subnet CIDR reservation.
	SubnetCidrReservation *types.SubnetCidrReservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubnetCidrReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateSubnetCidrReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateSubnetCidrReservation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSubnetCidrReservation"); err != nil {
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
	if err = addOpCreateSubnetCidrReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubnetCidrReservation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSubnetCidrReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSubnetCidrReservation",
	}
}

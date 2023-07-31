// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates or updates a scheduled scaling action for an Auto Scaling group. For
// more information, see Scheduled scaling (https://docs.aws.amazon.com/autoscaling/ec2/userguide/schedule_time.html)
// in the Amazon EC2 Auto Scaling User Guide. You can view the scheduled actions
// for an Auto Scaling group using the DescribeScheduledActions API call. If you
// are no longer using a scheduled action, you can delete it by calling the
// DeleteScheduledAction API. If you try to schedule your action in the past,
// Amazon EC2 Auto Scaling returns an error message.
func (c *Client) PutScheduledUpdateGroupAction(ctx context.Context, params *PutScheduledUpdateGroupActionInput, optFns ...func(*Options)) (*PutScheduledUpdateGroupActionOutput, error) {
	if params == nil {
		params = &PutScheduledUpdateGroupActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutScheduledUpdateGroupAction", params, optFns, c.addOperationPutScheduledUpdateGroupActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutScheduledUpdateGroupActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutScheduledUpdateGroupActionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The name of this scaling action.
	//
	// This member is required.
	ScheduledActionName *string

	// The desired capacity is the initial capacity of the Auto Scaling group after
	// the scheduled action runs and the capacity it attempts to maintain. It can scale
	// beyond this capacity if you add more scaling conditions. You must specify at
	// least one of the following properties: MaxSize , MinSize , or DesiredCapacity .
	DesiredCapacity *int32

	// The date and time for the recurring schedule to end, in UTC. For example,
	// "2021-06-01T00:00:00Z" .
	EndTime *time.Time

	// The maximum size of the Auto Scaling group.
	MaxSize *int32

	// The minimum size of the Auto Scaling group.
	MinSize *int32

	// The recurring schedule for this action. This format consists of five fields
	// separated by white spaces: [Minute] [Hour] [Day_of_Month] [Month_of_Year]
	// [Day_of_Week]. The value must be in quotes (for example, "30 0 1 1,6,12 *" ).
	// For more information about this format, see Crontab (http://crontab.org) . When
	// StartTime and EndTime are specified with Recurrence , they form the boundaries
	// of when the recurring action starts and stops. Cron expressions use Universal
	// Coordinated Time (UTC) by default.
	Recurrence *string

	// The date and time for this action to start, in YYYY-MM-DDThh:mm:ssZ format in
	// UTC/GMT only and in quotes (for example, "2021-06-01T00:00:00Z" ). If you
	// specify Recurrence and StartTime , Amazon EC2 Auto Scaling performs the action
	// at this time, and then performs the action based on the specified recurrence.
	StartTime *time.Time

	// This property is no longer used.
	Time *time.Time

	// Specifies the time zone for a cron expression. If a time zone is not provided,
	// UTC is used by default. Valid values are the canonical names of the IANA time
	// zones, derived from the IANA Time Zone Database (such as Etc/GMT+9 or
	// Pacific/Tahiti ). For more information, see
	// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	// .
	TimeZone *string

	noSmithyDocumentSerde
}

type PutScheduledUpdateGroupActionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutScheduledUpdateGroupActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutScheduledUpdateGroupAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutScheduledUpdateGroupAction{}, middleware.After)
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
	if err = addPutScheduledUpdateGroupActionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutScheduledUpdateGroupActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutScheduledUpdateGroupAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutScheduledUpdateGroupAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutScheduledUpdateGroupAction",
	}
}

type opPutScheduledUpdateGroupActionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutScheduledUpdateGroupActionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutScheduledUpdateGroupActionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "autoscaling"
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
				signingName = "autoscaling"
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
				v4aScheme.SigningName = aws.String("autoscaling")
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

func addPutScheduledUpdateGroupActionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutScheduledUpdateGroupActionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

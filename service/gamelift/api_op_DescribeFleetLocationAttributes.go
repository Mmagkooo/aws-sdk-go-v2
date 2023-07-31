// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information on a fleet's remote locations, including life-cycle
// status and any suspended fleet activity. This operation can be used in the
// following ways:
//   - To get data for specific locations, provide a fleet identifier and a list
//     of locations. Location data is returned in the order that it is requested.
//   - To get data for all locations, provide a fleet identifier only. Location
//     data is returned in no particular order.
//
// When requesting attributes for multiple locations, use the pagination
// parameters to retrieve results as a set of sequential pages. If successful, a
// LocationAttributes object is returned for each requested location. If the fleet
// does not have a requested location, no information is returned. This operation
// does not return the home Region. To get information on a fleet's home Region,
// call DescribeFleetAttributes . Learn more Setting up Amazon GameLift fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
func (c *Client) DescribeFleetLocationAttributes(ctx context.Context, params *DescribeFleetLocationAttributesInput, optFns ...func(*Options)) (*DescribeFleetLocationAttributesOutput, error) {
	if params == nil {
		params = &DescribeFleetLocationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetLocationAttributes", params, optFns, c.addOperationDescribeFleetLocationAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetLocationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetLocationAttributesInput struct {

	// A unique identifier for the fleet to retrieve remote locations for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This limit is not currently enforced.
	Limit *int32

	// A list of fleet locations to retrieve information for. Specify locations in the
	// form of an Amazon Web Services Region code, such as us-west-2 .
	Locations []string

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetLocationAttributesOutput struct {

	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) that is assigned to a Amazon GameLift fleet resource and uniquely identifies
	// it. ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	FleetArn *string

	// A unique identifier for the fleet that location attributes were requested for.
	FleetId *string

	// Location-specific information on the requested fleet's remote locations.
	LocationAttributes []types.LocationAttributes

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetLocationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetLocationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetLocationAttributes{}, middleware.After)
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
	if err = addDescribeFleetLocationAttributesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeFleetLocationAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetLocationAttributes(options.Region), middleware.Before); err != nil {
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

// DescribeFleetLocationAttributesAPIClient is a client that implements the
// DescribeFleetLocationAttributes operation.
type DescribeFleetLocationAttributesAPIClient interface {
	DescribeFleetLocationAttributes(context.Context, *DescribeFleetLocationAttributesInput, ...func(*Options)) (*DescribeFleetLocationAttributesOutput, error)
}

var _ DescribeFleetLocationAttributesAPIClient = (*Client)(nil)

// DescribeFleetLocationAttributesPaginatorOptions is the paginator options for
// DescribeFleetLocationAttributes
type DescribeFleetLocationAttributesPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This limit is not currently enforced.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetLocationAttributesPaginator is a paginator for
// DescribeFleetLocationAttributes
type DescribeFleetLocationAttributesPaginator struct {
	options   DescribeFleetLocationAttributesPaginatorOptions
	client    DescribeFleetLocationAttributesAPIClient
	params    *DescribeFleetLocationAttributesInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetLocationAttributesPaginator returns a new
// DescribeFleetLocationAttributesPaginator
func NewDescribeFleetLocationAttributesPaginator(client DescribeFleetLocationAttributesAPIClient, params *DescribeFleetLocationAttributesInput, optFns ...func(*DescribeFleetLocationAttributesPaginatorOptions)) *DescribeFleetLocationAttributesPaginator {
	if params == nil {
		params = &DescribeFleetLocationAttributesInput{}
	}

	options := DescribeFleetLocationAttributesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetLocationAttributesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetLocationAttributesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetLocationAttributes page.
func (p *DescribeFleetLocationAttributesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetLocationAttributesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeFleetLocationAttributes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeFleetLocationAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeFleetLocationAttributes",
	}
}

type opDescribeFleetLocationAttributesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeFleetLocationAttributesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeFleetLocationAttributesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "gamelift"
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
				signingName = "gamelift"
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
				v4aScheme.SigningName = aws.String("gamelift")
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

func addDescribeFleetLocationAttributesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeFleetLocationAttributesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

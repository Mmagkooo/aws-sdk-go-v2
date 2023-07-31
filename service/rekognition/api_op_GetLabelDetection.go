// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the label detection results of a Amazon Rekognition Video analysis started
// by StartLabelDetection . The label detection operation is started by a call to
// StartLabelDetection which returns a job identifier ( JobId ). When the label
// detection operation finishes, Amazon Rekognition publishes a completion status
// to the Amazon Simple Notification Service topic registered in the initial call
// to StartlabelDetection . To get the results of the label detection operation,
// first check that the status value published to the Amazon SNS topic is SUCCEEDED
// . If so, call GetLabelDetection and pass the job identifier ( JobId ) from the
// initial call to StartLabelDetection . GetLabelDetection returns an array of
// detected labels ( Labels ) sorted by the time the labels were detected. You can
// also sort by the label name by specifying NAME for the SortBy input parameter.
// If there is no NAME specified, the default sort is by timestamp. You can select
// how results are aggregated by using the AggregateBy input parameter. The
// default aggregation method is TIMESTAMPS . You can also aggregate by SEGMENTS ,
// which aggregates all instances of labels detected in a given segment. The
// returned Labels array may include the following attributes:
//   - Name - The name of the detected label.
//   - Confidence - The level of confidence in the label assigned to a detected
//     object.
//   - Parents - The ancestor labels for a detected label. GetLabelDetection
//     returns a hierarchical taxonomy of detected labels. For example, a detected car
//     might be assigned the label car. The label car has two parent labels: Vehicle
//     (its parent) and Transportation (its grandparent). The response includes the all
//     ancestors for a label, where every ancestor is a unique label. In the previous
//     example, Car, Vehicle, and Transportation are returned as unique labels in the
//     response.
//   - Aliases - Possible Aliases for the label.
//   - Categories - The label categories that the detected label belongs to.
//   - BoundingBox — Bounding boxes are described for all instances of detected
//     common object labels, returned in an array of Instance objects. An Instance
//     object contains a BoundingBox object, describing the location of the label on
//     the input image. It also includes the confidence for the accuracy of the
//     detected bounding box.
//   - Timestamp - Time, in milliseconds from the start of the video, that the
//     label was detected. For aggregation by SEGMENTS , the StartTimestampMillis ,
//     EndTimestampMillis , and DurationMillis structures are what define a segment.
//     Although the “Timestamp” structure is still returned with each label, its value
//     is set to be the same as StartTimestampMillis .
//
// Timestamp and Bounding box information are returned for detected Instances,
// only if aggregation is done by TIMESTAMPS . If aggregating by SEGMENTS ,
// information about detected instances isn’t returned. The version of the label
// model used for the detection is also returned. Note DominantColors isn't
// returned for Instances , although it is shown as part of the response in the
// sample seen below. Use MaxResults parameter to limit the number of labels
// returned. If there are more results than specified in MaxResults , the value of
// NextToken in the operation response contains a pagination token for getting the
// next set of results. To get the next page of results, call GetlabelDetection
// and populate the NextToken request parameter with the token value returned from
// the previous call to GetLabelDetection .
func (c *Client) GetLabelDetection(ctx context.Context, params *GetLabelDetectionInput, optFns ...func(*Options)) (*GetLabelDetectionOutput, error) {
	if params == nil {
		params = &GetLabelDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLabelDetection", params, optFns, c.addOperationGetLabelDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLabelDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLabelDetectionInput struct {

	// Job identifier for the label detection operation for which you want results
	// returned. You get the job identifer from an initial call to StartlabelDetection .
	//
	// This member is required.
	JobId *string

	// Defines how to aggregate the returned results. Results can be aggregated by
	// timestamps or segments.
	AggregateBy types.LabelDetectionAggregateBy

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more labels to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of labels.
	NextToken *string

	// Sort to use for elements in the Labels array. Use TIMESTAMP to sort array
	// elements by the time labels are detected. Use NAME to alphabetically group
	// elements for a label together. Within each label group, the array element are
	// sorted by detection confidence. The default sort is by TIMESTAMP .
	SortBy types.LabelDetectionSortBy

	noSmithyDocumentSerde
}

type GetLabelDetectionOutput struct {

	// Information about the paramters used when getting a response. Includes
	// information on aggregation and sorting methods.
	GetRequestMetadata *types.GetLabelDetectionRequestMetadata

	// Job identifier for the label detection operation for which you want to obtain
	// results. The job identifer is returned by an initial call to
	// StartLabelDetection.
	JobId *string

	// The current status of the label detection job.
	JobStatus types.VideoJobStatus

	// A job identifier specified in the call to StartLabelDetection and returned in
	// the job completion notification sent to your Amazon Simple Notification Service
	// topic.
	JobTag *string

	// Version number of the label detection model that was used to detect labels.
	LabelModelVersion *string

	// An array of labels detected in the video. Each element contains the detected
	// label and the time, in milliseconds from the start of the video, that the label
	// was detected.
	Labels []types.LabelDetection

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of labels.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetection use Video to specify a video for
	// analysis. The supported file formats are .mp4, .mov and .avi.
	Video *types.Video

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition video
	// operation.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLabelDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLabelDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLabelDetection{}, middleware.After)
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
	if err = addGetLabelDetectionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetLabelDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLabelDetection(options.Region), middleware.Before); err != nil {
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

// GetLabelDetectionAPIClient is a client that implements the GetLabelDetection
// operation.
type GetLabelDetectionAPIClient interface {
	GetLabelDetection(context.Context, *GetLabelDetectionInput, ...func(*Options)) (*GetLabelDetectionOutput, error)
}

var _ GetLabelDetectionAPIClient = (*Client)(nil)

// GetLabelDetectionPaginatorOptions is the paginator options for GetLabelDetection
type GetLabelDetectionPaginatorOptions struct {
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetLabelDetectionPaginator is a paginator for GetLabelDetection
type GetLabelDetectionPaginator struct {
	options   GetLabelDetectionPaginatorOptions
	client    GetLabelDetectionAPIClient
	params    *GetLabelDetectionInput
	nextToken *string
	firstPage bool
}

// NewGetLabelDetectionPaginator returns a new GetLabelDetectionPaginator
func NewGetLabelDetectionPaginator(client GetLabelDetectionAPIClient, params *GetLabelDetectionInput, optFns ...func(*GetLabelDetectionPaginatorOptions)) *GetLabelDetectionPaginator {
	if params == nil {
		params = &GetLabelDetectionInput{}
	}

	options := GetLabelDetectionPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetLabelDetectionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetLabelDetectionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetLabelDetection page.
func (p *GetLabelDetectionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetLabelDetectionOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetLabelDetection(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetLabelDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetLabelDetection",
	}
}

type opGetLabelDetectionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetLabelDetectionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetLabelDetectionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "rekognition"
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
				signingName = "rekognition"
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
				v4aScheme.SigningName = aws.String("rekognition")
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

func addGetLabelDetectionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetLabelDetectionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

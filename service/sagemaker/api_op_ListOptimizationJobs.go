// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the optimization jobs in your account and their properties.
func (c *Client) ListOptimizationJobs(ctx context.Context, params *ListOptimizationJobsInput, optFns ...func(*Options)) (*ListOptimizationJobsOutput, error) {
	if params == nil {
		params = &ListOptimizationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOptimizationJobs", params, optFns, c.addOperationListOptimizationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOptimizationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOptimizationJobsInput struct {

	// Filters the results to only those optimization jobs that were created after the
	// specified time.
	CreationTimeAfter *time.Time

	// Filters the results to only those optimization jobs that were created before
	// the specified time.
	CreationTimeBefore *time.Time

	// Filters the results to only those optimization jobs that were updated after the
	// specified time.
	LastModifiedTimeAfter *time.Time

	// Filters the results to only those optimization jobs that were updated before
	// the specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of optimization jobs to return in the response. The default
	// is 50.
	MaxResults *int32

	// Filters the results to only those optimization jobs with a name that contains
	// the specified string.
	NameContains *string

	// A token that you use to get the next set of results following a truncated
	// response. If the response to the previous request was truncated, that response
	// provides the value for this token.
	NextToken *string

	// Filters the results to only those optimization jobs that apply the specified
	// optimization techniques. You can specify either Quantization or Compilation .
	OptimizationContains *string

	// The field by which to sort the optimization jobs in the response. The default
	// is CreationTime
	SortBy types.ListOptimizationJobsSortBy

	// The sort order for results. The default is Ascending
	SortOrder types.SortOrder

	// Filters the results to only those optimization jobs with the specified status.
	StatusEquals types.OptimizationJobStatus

	noSmithyDocumentSerde
}

type ListOptimizationJobsOutput struct {

	// A list of optimization jobs and their properties that matches any of the
	// filters you specified in the request.
	//
	// This member is required.
	OptimizationJobSummaries []types.OptimizationJobSummary

	// The token to use in a subsequent request to get the next set of results
	// following a truncated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOptimizationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOptimizationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOptimizationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOptimizationJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOptimizationJobs(options.Region), middleware.Before); err != nil {
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

// ListOptimizationJobsPaginatorOptions is the paginator options for
// ListOptimizationJobs
type ListOptimizationJobsPaginatorOptions struct {
	// The maximum number of optimization jobs to return in the response. The default
	// is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOptimizationJobsPaginator is a paginator for ListOptimizationJobs
type ListOptimizationJobsPaginator struct {
	options   ListOptimizationJobsPaginatorOptions
	client    ListOptimizationJobsAPIClient
	params    *ListOptimizationJobsInput
	nextToken *string
	firstPage bool
}

// NewListOptimizationJobsPaginator returns a new ListOptimizationJobsPaginator
func NewListOptimizationJobsPaginator(client ListOptimizationJobsAPIClient, params *ListOptimizationJobsInput, optFns ...func(*ListOptimizationJobsPaginatorOptions)) *ListOptimizationJobsPaginator {
	if params == nil {
		params = &ListOptimizationJobsInput{}
	}

	options := ListOptimizationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOptimizationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOptimizationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOptimizationJobs page.
func (p *ListOptimizationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOptimizationJobsOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListOptimizationJobs(ctx, &params, optFns...)
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

// ListOptimizationJobsAPIClient is a client that implements the
// ListOptimizationJobs operation.
type ListOptimizationJobsAPIClient interface {
	ListOptimizationJobs(context.Context, *ListOptimizationJobsInput, ...func(*Options)) (*ListOptimizationJobsOutput, error)
}

var _ ListOptimizationJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListOptimizationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOptimizationJobs",
	}
}

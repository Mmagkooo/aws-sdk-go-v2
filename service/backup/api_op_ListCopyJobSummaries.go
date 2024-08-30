// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This request obtains a list of copy jobs created or running within the the most
// recent 30 days. You can include parameters AccountID, State, ResourceType,
// MessageCategory, AggregationPeriod, MaxResults, or NextToken to filter results.
//
// This request returns a summary that contains Region, Account, State,
// RestourceType, MessageCategory, StartTime, EndTime, and Count of included jobs.
func (c *Client) ListCopyJobSummaries(ctx context.Context, params *ListCopyJobSummariesInput, optFns ...func(*Options)) (*ListCopyJobSummariesOutput, error) {
	if params == nil {
		params = &ListCopyJobSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCopyJobSummaries", params, optFns, c.addOperationListCopyJobSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCopyJobSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCopyJobSummariesInput struct {

	// Returns the job count for the specified account.
	//
	// If the request is sent from a member account or an account not part of Amazon
	// Web Services Organizations, jobs within requestor's account will be returned.
	//
	// Root, admin, and delegated administrator accounts can use the value ANY to
	// return job counts from every account in the organization.
	//
	// AGGREGATE_ALL aggregates job counts from all accounts within the authenticated
	// organization, then returns the sum.
	AccountId *string

	// The period for the returned results.
	//
	//   - ONE_DAY - The daily job count for the prior 14 days.
	//
	//   - SEVEN_DAYS - The aggregated job count for the prior 7 days.
	//
	//   - FOURTEEN_DAYS - The aggregated job count for prior 14 days.
	AggregationPeriod types.AggregationPeriod

	// This parameter sets the maximum number of items to be returned.
	//
	// The value is an integer. Range of accepted values is from 1 to 500.
	MaxResults *int32

	// This parameter returns the job count for the specified message category.
	//
	// Example accepted strings include AccessDenied , Success , and InvalidParameters
	// . See [Monitoring]for a list of accepted MessageCategory strings.
	//
	// The the value ANY returns count of all message categories.
	//
	// AGGREGATE_ALL aggregates job counts for all message categories and returns the
	// sum.
	//
	// [Monitoring]: https://docs.aws.amazon.com/aws-backup/latest/devguide/monitoring.html
	MessageCategory *string

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return MaxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	// Returns the job count for the specified resource type. Use request
	// GetSupportedResourceTypes to obtain strings for supported resource types.
	//
	// The the value ANY returns count of all resource types.
	//
	// AGGREGATE_ALL aggregates job counts for all resource types and returns the sum.
	//
	// The type of Amazon Web Services resource to be backed up; for example, an
	// Amazon Elastic Block Store (Amazon EBS) volume or an Amazon Relational Database
	// Service (Amazon RDS) database.
	ResourceType *string

	// This parameter returns the job count for jobs with the specified state.
	//
	// The the value ANY returns count of all states.
	//
	// AGGREGATE_ALL aggregates job counts for all states and returns the sum.
	State types.CopyJobStatus

	noSmithyDocumentSerde
}

type ListCopyJobSummariesOutput struct {

	// The period for the returned results.
	//
	//   - ONE_DAY - The daily job count for the prior 14 days.
	//
	//   - SEVEN_DAYS - The aggregated job count for the prior 7 days.
	//
	//   - FOURTEEN_DAYS - The aggregated job count for prior 14 days.
	AggregationPeriod *string

	// This return shows a summary that contains Region, Account, State, ResourceType,
	// MessageCategory, StartTime, EndTime, and Count of included jobs.
	CopyJobSummaries []types.CopyJobSummary

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return MaxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCopyJobSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCopyJobSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCopyJobSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCopyJobSummaries"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCopyJobSummaries(options.Region), middleware.Before); err != nil {
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

// ListCopyJobSummariesPaginatorOptions is the paginator options for
// ListCopyJobSummaries
type ListCopyJobSummariesPaginatorOptions struct {
	// This parameter sets the maximum number of items to be returned.
	//
	// The value is an integer. Range of accepted values is from 1 to 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCopyJobSummariesPaginator is a paginator for ListCopyJobSummaries
type ListCopyJobSummariesPaginator struct {
	options   ListCopyJobSummariesPaginatorOptions
	client    ListCopyJobSummariesAPIClient
	params    *ListCopyJobSummariesInput
	nextToken *string
	firstPage bool
}

// NewListCopyJobSummariesPaginator returns a new ListCopyJobSummariesPaginator
func NewListCopyJobSummariesPaginator(client ListCopyJobSummariesAPIClient, params *ListCopyJobSummariesInput, optFns ...func(*ListCopyJobSummariesPaginatorOptions)) *ListCopyJobSummariesPaginator {
	if params == nil {
		params = &ListCopyJobSummariesInput{}
	}

	options := ListCopyJobSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCopyJobSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCopyJobSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCopyJobSummaries page.
func (p *ListCopyJobSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCopyJobSummariesOutput, error) {
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
	result, err := p.client.ListCopyJobSummaries(ctx, &params, optFns...)
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

// ListCopyJobSummariesAPIClient is a client that implements the
// ListCopyJobSummaries operation.
type ListCopyJobSummariesAPIClient interface {
	ListCopyJobSummaries(context.Context, *ListCopyJobSummariesInput, ...func(*Options)) (*ListCopyJobSummariesOutput, error)
}

var _ ListCopyJobSummariesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCopyJobSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCopyJobSummaries",
	}
}

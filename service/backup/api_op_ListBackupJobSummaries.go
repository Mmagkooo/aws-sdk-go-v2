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

// This is a request for a summary of backup jobs created or running within the
// most recent 30 days. You can include parameters AccountID, State, ResourceType,
// MessageCategory, AggregationPeriod, MaxResults, or NextToken to filter results.
//
// This request returns a summary that contains Region, Account, State,
// ResourceType, MessageCategory, StartTime, EndTime, and Count of included jobs.
func (c *Client) ListBackupJobSummaries(ctx context.Context, params *ListBackupJobSummariesInput, optFns ...func(*Options)) (*ListBackupJobSummariesOutput, error) {
	if params == nil {
		params = &ListBackupJobSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupJobSummaries", params, optFns, c.addOperationListBackupJobSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupJobSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupJobSummariesInput struct {

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

	// The maximum number of items to be returned.
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
	//
	// Completed with issues is a status found only in the Backup console. For API,
	// this status refers to jobs with a state of COMPLETED and a MessageCategory with
	// a value other than SUCCESS ; that is, the status is completed but comes with a
	// status message. To obtain the job count for Completed with issues , run two GET
	// requests, and subtract the second, smaller number:
	//
	// GET /audit/backup-job-summaries?AggregationPeriod=FOURTEEN_DAYS&State=COMPLETED
	//
	// GET
	// /audit/backup-job-summaries?AggregationPeriod=FOURTEEN_DAYS&MessageCategory=SUCCESS&State=COMPLETED
	State types.BackupJobStatus

	noSmithyDocumentSerde
}

type ListBackupJobSummariesOutput struct {

	// The period for the returned results.
	//
	//   - ONE_DAY - The daily job count for the prior 14 days.
	//
	//   - SEVEN_DAYS - The aggregated job count for the prior 7 days.
	//
	//   - FOURTEEN_DAYS - The aggregated job count for prior 14 days.
	AggregationPeriod *string

	// The summary information.
	BackupJobSummaries []types.BackupJobSummary

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return MaxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackupJobSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupJobSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupJobSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBackupJobSummaries"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupJobSummaries(options.Region), middleware.Before); err != nil {
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

// ListBackupJobSummariesPaginatorOptions is the paginator options for
// ListBackupJobSummaries
type ListBackupJobSummariesPaginatorOptions struct {
	// The maximum number of items to be returned.
	//
	// The value is an integer. Range of accepted values is from 1 to 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBackupJobSummariesPaginator is a paginator for ListBackupJobSummaries
type ListBackupJobSummariesPaginator struct {
	options   ListBackupJobSummariesPaginatorOptions
	client    ListBackupJobSummariesAPIClient
	params    *ListBackupJobSummariesInput
	nextToken *string
	firstPage bool
}

// NewListBackupJobSummariesPaginator returns a new ListBackupJobSummariesPaginator
func NewListBackupJobSummariesPaginator(client ListBackupJobSummariesAPIClient, params *ListBackupJobSummariesInput, optFns ...func(*ListBackupJobSummariesPaginatorOptions)) *ListBackupJobSummariesPaginator {
	if params == nil {
		params = &ListBackupJobSummariesInput{}
	}

	options := ListBackupJobSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBackupJobSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBackupJobSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBackupJobSummaries page.
func (p *ListBackupJobSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBackupJobSummariesOutput, error) {
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
	result, err := p.client.ListBackupJobSummaries(ctx, &params, optFns...)
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

// ListBackupJobSummariesAPIClient is a client that implements the
// ListBackupJobSummaries operation.
type ListBackupJobSummariesAPIClient interface {
	ListBackupJobSummaries(context.Context, *ListBackupJobSummariesInput, ...func(*Options)) (*ListBackupJobSummariesOutput, error)
}

var _ ListBackupJobSummariesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListBackupJobSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBackupJobSummaries",
	}
}

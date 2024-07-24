// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists configured table associations for a membership.
func (c *Client) ListConfiguredTableAssociations(ctx context.Context, params *ListConfiguredTableAssociationsInput, optFns ...func(*Options)) (*ListConfiguredTableAssociationsOutput, error) {
	if params == nil {
		params = &ListConfiguredTableAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfiguredTableAssociations", params, optFns, c.addOperationListConfiguredTableAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfiguredTableAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfiguredTableAssociationsInput struct {

	// A unique identifier for the membership to list configured table associations
	// for. Currently accepts the membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	// The maximum size of the results that is returned per call.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConfiguredTableAssociationsOutput struct {

	// The retrieved list of configured table associations.
	//
	// This member is required.
	ConfiguredTableAssociationSummaries []types.ConfiguredTableAssociationSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConfiguredTableAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfiguredTableAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfiguredTableAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConfiguredTableAssociations"); err != nil {
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
	if err = addOpListConfiguredTableAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfiguredTableAssociations(options.Region), middleware.Before); err != nil {
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

// ListConfiguredTableAssociationsPaginatorOptions is the paginator options for
// ListConfiguredTableAssociations
type ListConfiguredTableAssociationsPaginatorOptions struct {
	// The maximum size of the results that is returned per call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConfiguredTableAssociationsPaginator is a paginator for
// ListConfiguredTableAssociations
type ListConfiguredTableAssociationsPaginator struct {
	options   ListConfiguredTableAssociationsPaginatorOptions
	client    ListConfiguredTableAssociationsAPIClient
	params    *ListConfiguredTableAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListConfiguredTableAssociationsPaginator returns a new
// ListConfiguredTableAssociationsPaginator
func NewListConfiguredTableAssociationsPaginator(client ListConfiguredTableAssociationsAPIClient, params *ListConfiguredTableAssociationsInput, optFns ...func(*ListConfiguredTableAssociationsPaginatorOptions)) *ListConfiguredTableAssociationsPaginator {
	if params == nil {
		params = &ListConfiguredTableAssociationsInput{}
	}

	options := ListConfiguredTableAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConfiguredTableAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConfiguredTableAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConfiguredTableAssociations page.
func (p *ListConfiguredTableAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConfiguredTableAssociationsOutput, error) {
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
	result, err := p.client.ListConfiguredTableAssociations(ctx, &params, optFns...)
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

// ListConfiguredTableAssociationsAPIClient is a client that implements the
// ListConfiguredTableAssociations operation.
type ListConfiguredTableAssociationsAPIClient interface {
	ListConfiguredTableAssociations(context.Context, *ListConfiguredTableAssociationsInput, ...func(*Options)) (*ListConfiguredTableAssociationsOutput, error)
}

var _ ListConfiguredTableAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListConfiguredTableAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConfiguredTableAssociations",
	}
}

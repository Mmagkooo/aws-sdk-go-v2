// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Type objects by the source API association ID.
func (c *Client) ListTypesByAssociation(ctx context.Context, params *ListTypesByAssociationInput, optFns ...func(*Options)) (*ListTypesByAssociationOutput, error) {
	if params == nil {
		params = &ListTypesByAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTypesByAssociation", params, optFns, c.addOperationListTypesByAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTypesByAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypesByAssociationInput struct {

	// The ID generated by the AppSync service for the source API association.
	//
	// This member is required.
	AssociationId *string

	// The format type.
	//
	// This member is required.
	Format types.TypeDefinitionFormat

	// The identifier of the AppSync Merged API. This is generated by the AppSync
	// service. In most cases, Merged APIs (especially in your account) only require
	// the API ID value or ARN of the merged API. However, Merged APIs in other
	// accounts (cross-account use cases) strictly require the full resource ARN of the
	// merged API.
	//
	// This member is required.
	MergedApiIdentifier *string

	// The maximum number of results that you want the request to return.
	MaxResults int32

	// An identifier that was returned from the previous call to this operation, which
	// you can use to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTypesByAssociationOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// you can use to return the next set of items in the list.
	NextToken *string

	// The Type objects.
	Types []types.Type

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTypesByAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTypesByAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTypesByAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTypesByAssociation"); err != nil {
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
	if err = addOpListTypesByAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTypesByAssociation(options.Region), middleware.Before); err != nil {
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

// ListTypesByAssociationPaginatorOptions is the paginator options for
// ListTypesByAssociation
type ListTypesByAssociationPaginatorOptions struct {
	// The maximum number of results that you want the request to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTypesByAssociationPaginator is a paginator for ListTypesByAssociation
type ListTypesByAssociationPaginator struct {
	options   ListTypesByAssociationPaginatorOptions
	client    ListTypesByAssociationAPIClient
	params    *ListTypesByAssociationInput
	nextToken *string
	firstPage bool
}

// NewListTypesByAssociationPaginator returns a new ListTypesByAssociationPaginator
func NewListTypesByAssociationPaginator(client ListTypesByAssociationAPIClient, params *ListTypesByAssociationInput, optFns ...func(*ListTypesByAssociationPaginatorOptions)) *ListTypesByAssociationPaginator {
	if params == nil {
		params = &ListTypesByAssociationInput{}
	}

	options := ListTypesByAssociationPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTypesByAssociationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTypesByAssociationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTypesByAssociation page.
func (p *ListTypesByAssociationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTypesByAssociationOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListTypesByAssociation(ctx, &params, optFns...)
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

// ListTypesByAssociationAPIClient is a client that implements the
// ListTypesByAssociation operation.
type ListTypesByAssociationAPIClient interface {
	ListTypesByAssociation(context.Context, *ListTypesByAssociationInput, ...func(*Options)) (*ListTypesByAssociationOutput, error)
}

var _ ListTypesByAssociationAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTypesByAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTypesByAssociation",
	}
}

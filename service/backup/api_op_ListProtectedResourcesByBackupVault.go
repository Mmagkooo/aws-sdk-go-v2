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

// This request lists the protected resources corresponding to each backup vault.
func (c *Client) ListProtectedResourcesByBackupVault(ctx context.Context, params *ListProtectedResourcesByBackupVaultInput, optFns ...func(*Options)) (*ListProtectedResourcesByBackupVaultOutput, error) {
	if params == nil {
		params = &ListProtectedResourcesByBackupVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProtectedResourcesByBackupVault", params, optFns, c.addOperationListProtectedResourcesByBackupVaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProtectedResourcesByBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtectedResourcesByBackupVaultInput struct {

	// The list of protected resources by backup vault within the vault(s) you specify
	// by name.
	//
	// This member is required.
	BackupVaultName *string

	// The list of protected resources by backup vault within the vault(s) you specify
	// by account ID.
	BackupVaultAccountId *string

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProtectedResourcesByBackupVaultOutput struct {

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// These are the results returned for the request
	// ListProtectedResourcesByBackupVault.
	Results []types.ProtectedResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProtectedResourcesByBackupVaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProtectedResourcesByBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProtectedResourcesByBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProtectedResourcesByBackupVault"); err != nil {
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
	if err = addOpListProtectedResourcesByBackupVaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProtectedResourcesByBackupVault(options.Region), middleware.Before); err != nil {
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

// ListProtectedResourcesByBackupVaultPaginatorOptions is the paginator options
// for ListProtectedResourcesByBackupVault
type ListProtectedResourcesByBackupVaultPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProtectedResourcesByBackupVaultPaginator is a paginator for
// ListProtectedResourcesByBackupVault
type ListProtectedResourcesByBackupVaultPaginator struct {
	options   ListProtectedResourcesByBackupVaultPaginatorOptions
	client    ListProtectedResourcesByBackupVaultAPIClient
	params    *ListProtectedResourcesByBackupVaultInput
	nextToken *string
	firstPage bool
}

// NewListProtectedResourcesByBackupVaultPaginator returns a new
// ListProtectedResourcesByBackupVaultPaginator
func NewListProtectedResourcesByBackupVaultPaginator(client ListProtectedResourcesByBackupVaultAPIClient, params *ListProtectedResourcesByBackupVaultInput, optFns ...func(*ListProtectedResourcesByBackupVaultPaginatorOptions)) *ListProtectedResourcesByBackupVaultPaginator {
	if params == nil {
		params = &ListProtectedResourcesByBackupVaultInput{}
	}

	options := ListProtectedResourcesByBackupVaultPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProtectedResourcesByBackupVaultPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProtectedResourcesByBackupVaultPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProtectedResourcesByBackupVault page.
func (p *ListProtectedResourcesByBackupVaultPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProtectedResourcesByBackupVaultOutput, error) {
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
	result, err := p.client.ListProtectedResourcesByBackupVault(ctx, &params, optFns...)
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

// ListProtectedResourcesByBackupVaultAPIClient is a client that implements the
// ListProtectedResourcesByBackupVault operation.
type ListProtectedResourcesByBackupVaultAPIClient interface {
	ListProtectedResourcesByBackupVault(context.Context, *ListProtectedResourcesByBackupVaultInput, ...func(*Options)) (*ListProtectedResourcesByBackupVaultOutput, error)
}

var _ ListProtectedResourcesByBackupVaultAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListProtectedResourcesByBackupVault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProtectedResourcesByBackupVault",
	}
}

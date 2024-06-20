// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an existing external connection to a repository. One external connection
// is allowed per repository.
//
// A repository can have one or more upstream repositories, or an external
// connection.
func (c *Client) AssociateExternalConnection(ctx context.Context, params *AssociateExternalConnectionInput, optFns ...func(*Options)) (*AssociateExternalConnectionOutput, error) {
	if params == nil {
		params = &AssociateExternalConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateExternalConnection", params, optFns, c.addOperationAssociateExternalConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateExternalConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateExternalConnectionInput struct {

	// The name of the domain that contains the repository.
	//
	// This member is required.
	Domain *string

	//  The name of the external connection to add to the repository. The following
	// values are supported:
	//
	//   - public:npmjs - for the npm public repository.
	//
	//   - public:nuget-org - for the NuGet Gallery.
	//
	//   - public:pypi - for the Python Package Index.
	//
	//   - public:maven-central - for Maven Central.
	//
	//   - public:maven-googleandroid - for the Google Android repository.
	//
	//   - public:maven-gradleplugins - for the Gradle plugins repository.
	//
	//   - public:maven-commonsware - for the CommonsWare Android repository.
	//
	//   - public:maven-clojars - for the Clojars repository.
	//
	//   - public:ruby-gems-org - for RubyGems.org.
	//
	//   - public:crates-io - for Crates.io.
	//
	// This member is required.
	ExternalConnection *string

	//  The name of the repository to which the external connection is added.
	//
	// This member is required.
	Repository *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	noSmithyDocumentSerde
}

type AssociateExternalConnectionOutput struct {

	//  Information about the connected repository after processing the request.
	Repository *types.RepositoryDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateExternalConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateExternalConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateExternalConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateExternalConnection"); err != nil {
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
	if err = addOpAssociateExternalConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateExternalConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateExternalConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateExternalConnection",
	}
}

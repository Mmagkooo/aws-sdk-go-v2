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

// Sets the package origin configuration for a package.
//
// The package origin configuration determines how new versions of a package can
// be added to a repository. You can allow or block direct publishing of new
// package versions, or ingestion and retaining of new package versions from an
// external connection or upstream source. For more information about package
// origin controls and configuration, see [Editing package origin controls]in the CodeArtifact User Guide.
//
// PutPackageOriginConfiguration can be called on a package that doesn't yet exist
// in the repository. When called on a package that does not exist, a package is
// created in the repository with no versions and the requested restrictions are
// set on the package. This can be used to preemptively block ingesting or
// retaining any versions from external connections or upstream repositories, or to
// block publishing any versions of the package into the repository before
// connecting any package managers or publishers to the repository.
//
// [Editing package origin controls]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-origin-controls.html
func (c *Client) PutPackageOriginConfiguration(ctx context.Context, params *PutPackageOriginConfigurationInput, optFns ...func(*Options)) (*PutPackageOriginConfigurationOutput, error) {
	if params == nil {
		params = &PutPackageOriginConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPackageOriginConfiguration", params, optFns, c.addOperationPutPackageOriginConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPackageOriginConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPackageOriginConfigurationInput struct {

	// The name of the domain that contains the repository that contains the package.
	//
	// This member is required.
	Domain *string

	// A format that specifies the type of the package to be updated.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package to be updated.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the package.
	//
	// This member is required.
	Repository *string

	// A [PackageOriginRestrictions] object that contains information about the upstream and publish package
	// origin restrictions. The upstream restriction determines if new package
	// versions can be ingested or retained from external connections or upstream
	// repositories. The publish restriction determines if new package versions can be
	// published directly to the repository.
	//
	// You must include both the desired upstream and publish restrictions.
	//
	// [PackageOriginRestrictions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageOriginRestrictions.html
	//
	// This member is required.
	Restrictions *types.PackageOriginRestrictions

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The namespace of the package to be updated. The package component that
	// specifies its namespace depends on its type. For example:
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	noSmithyDocumentSerde
}

type PutPackageOriginConfigurationOutput struct {

	// A [PackageOriginConfiguration] object that describes the origin configuration set for the package. It
	// contains a [PackageOriginRestrictions]object that describes how new versions of the package can be
	// introduced to the repository.
	//
	// [PackageOriginRestrictions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageOriginRestrictions.html
	// [PackageOriginConfiguration]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageOriginConfiguration.html
	OriginConfiguration *types.PackageOriginConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPackageOriginConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutPackageOriginConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutPackageOriginConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutPackageOriginConfiguration"); err != nil {
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
	if err = addOpPutPackageOriginConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPackageOriginConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPackageOriginConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutPackageOriginConfiguration",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Resets the password for any user in your Managed Microsoft AD or Simple AD
// directory. Disabled users will become enabled and can be authenticated following
// the API call.
//
// You can reset the password for any user in your directory with the following
// exceptions:
//
//   - For Simple AD, you cannot reset the password for any user that is a member
//     of either the Domain Admins or Enterprise Admins group except for the
//     administrator user.
//
//   - For Managed Microsoft AD, you can only reset the password for a user that
//     is in an OU based off of the NetBIOS name that you typed when you created your
//     directory. For example, you cannot reset the password for a user in the Amazon
//     Web Services Reserved OU. For more information about the OU structure for an
//     Managed Microsoft AD directory, see [What Gets Created]in the Directory Service Administration
//     Guide.
//
// [What Gets Created]: https://docs.aws.amazon.com/directoryservice/latest/admin-guide/ms_ad_getting_started_what_gets_created.html
func (c *Client) ResetUserPassword(ctx context.Context, params *ResetUserPasswordInput, optFns ...func(*Options)) (*ResetUserPasswordOutput, error) {
	if params == nil {
		params = &ResetUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetUserPassword", params, optFns, c.addOperationResetUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetUserPasswordInput struct {

	// Identifier of the Managed Microsoft AD or Simple AD directory in which the user
	// resides.
	//
	// This member is required.
	DirectoryId *string

	// The new password that will be reset.
	//
	// This member is required.
	NewPassword *string

	// The user name of the user whose password will be reset.
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

type ResetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ResetUserPassword"); err != nil {
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
	if err = addOpResetUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetUserPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ResetUserPassword",
	}
}

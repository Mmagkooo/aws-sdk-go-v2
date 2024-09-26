// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables a policy type in a root. After you enable a policy type in a root, you
// can attach policies of that type to the root, any organizational unit (OU), or
// account in that root. You can undo this by using the DisablePolicyTypeoperation.
//
// This is an asynchronous request that Amazon Web Services performs in the
// background. Amazon Web Services recommends that you first use ListRootsto see the status
// of policy types for a specified root, and then use this operation.
//
// This operation can be called only from the organization's management account or
// by a member account that is a delegated administrator for an Amazon Web Services
// service.
//
// You can enable a policy type in a root only if that policy type is available in
// the organization. To view the status of available policy types in the
// organization, use DescribeOrganization.
func (c *Client) EnablePolicyType(ctx context.Context, params *EnablePolicyTypeInput, optFns ...func(*Options)) (*EnablePolicyTypeOutput, error) {
	if params == nil {
		params = &EnablePolicyTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnablePolicyType", params, optFns, c.addOperationEnablePolicyTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnablePolicyTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnablePolicyTypeInput struct {

	// The policy type that you want to enable. You can specify one of the following
	// values:
	//
	// [SERVICE_CONTROL_POLICY]
	//
	// [BACKUP_POLICY]
	//
	// [TAG_POLICY]
	//
	// [CHATBOT_POLICY]
	//
	// [AISERVICES_OPT_OUT_POLICY]
	//
	// [AISERVICES_OPT_OUT_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html
	// [BACKUP_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html
	// [SERVICE_CONTROL_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html
	// [CHATBOT_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_chatbot.html
	// [TAG_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html
	//
	// This member is required.
	PolicyType types.PolicyType

	// The unique identifier (ID) of the root in which you want to enable a policy
	// type. You can get the ID from the ListRootsoperation.
	//
	// The [regex pattern] for a root ID string requires "r-" followed by from 4 to 32 lowercase
	// letters or digits.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	RootId *string

	noSmithyDocumentSerde
}

type EnablePolicyTypeOutput struct {

	// A structure that shows the root with the updated list of enabled policy types.
	Root *types.Root

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnablePolicyTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnablePolicyType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnablePolicyType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnablePolicyType"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addOpEnablePolicyTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnablePolicyType(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opEnablePolicyType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnablePolicyType",
	}
}

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

// Updates the analysis rule for a configured table association.
func (c *Client) UpdateConfiguredTableAssociationAnalysisRule(ctx context.Context, params *UpdateConfiguredTableAssociationAnalysisRuleInput, optFns ...func(*Options)) (*UpdateConfiguredTableAssociationAnalysisRuleOutput, error) {
	if params == nil {
		params = &UpdateConfiguredTableAssociationAnalysisRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfiguredTableAssociationAnalysisRule", params, optFns, c.addOperationUpdateConfiguredTableAssociationAnalysisRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfiguredTableAssociationAnalysisRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConfiguredTableAssociationAnalysisRuleInput struct {

	//  The updated analysis rule policy for the conﬁgured table association.
	//
	// This member is required.
	AnalysisRulePolicy types.ConfiguredTableAssociationAnalysisRulePolicy

	//  The analysis rule type that you want to update.
	//
	// This member is required.
	AnalysisRuleType types.ConfiguredTableAssociationAnalysisRuleType

	//  The identifier for the configured table association to update.
	//
	// This member is required.
	ConfiguredTableAssociationIdentifier *string

	//  A unique identifier for the membership that the configured table association
	// belongs to. Currently accepts the membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	noSmithyDocumentSerde
}

type UpdateConfiguredTableAssociationAnalysisRuleOutput struct {

	//  The updated analysis rule for the conﬁgured table association. In the console,
	// the ConfiguredTableAssociationAnalysisRule is referred to as the collaboration
	// analysis rule.
	//
	// This member is required.
	AnalysisRule *types.ConfiguredTableAssociationAnalysisRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConfiguredTableAssociationAnalysisRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConfiguredTableAssociationAnalysisRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConfiguredTableAssociationAnalysisRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateConfiguredTableAssociationAnalysisRule"); err != nil {
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
	if err = addOpUpdateConfiguredTableAssociationAnalysisRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfiguredTableAssociationAnalysisRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConfiguredTableAssociationAnalysisRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateConfiguredTableAssociationAnalysisRule",
	}
}

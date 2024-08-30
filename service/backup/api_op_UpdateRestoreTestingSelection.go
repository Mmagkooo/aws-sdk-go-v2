// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified restore testing selection.
//
// Most elements except the RestoreTestingSelectionName can be updated with this
// request.
//
// You can use either protected resource ARNs or conditions, but not both.
func (c *Client) UpdateRestoreTestingSelection(ctx context.Context, params *UpdateRestoreTestingSelectionInput, optFns ...func(*Options)) (*UpdateRestoreTestingSelectionOutput, error) {
	if params == nil {
		params = &UpdateRestoreTestingSelectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRestoreTestingSelection", params, optFns, c.addOperationUpdateRestoreTestingSelectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRestoreTestingSelectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRestoreTestingSelectionInput struct {

	// The restore testing plan name is required to update the indicated testing plan.
	//
	// This member is required.
	RestoreTestingPlanName *string

	// To update your restore testing selection, you can use either protected resource
	// ARNs or conditions, but not both. That is, if your selection has
	// ProtectedResourceArns , requesting an update with the parameter
	// ProtectedResourceConditions will be unsuccessful.
	//
	// This member is required.
	RestoreTestingSelection *types.RestoreTestingSelectionForUpdate

	// The required restore testing selection name of the restore testing selection
	// you wish to update.
	//
	// This member is required.
	RestoreTestingSelectionName *string

	noSmithyDocumentSerde
}

type UpdateRestoreTestingSelectionOutput struct {

	// The time the resource testing selection was updated successfully.
	//
	// This member is required.
	CreationTime *time.Time

	// Unique string that is the name of the restore testing plan.
	//
	// This member is required.
	RestoreTestingPlanArn *string

	// The restore testing plan with which the updated restore testing selection is
	// associated.
	//
	// This member is required.
	RestoreTestingPlanName *string

	// The returned restore testing selection name.
	//
	// This member is required.
	RestoreTestingSelectionName *string

	// The time the update completed for the restore testing selection.
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRestoreTestingSelectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRestoreTestingSelection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRestoreTestingSelection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRestoreTestingSelection"); err != nil {
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
	if err = addOpUpdateRestoreTestingSelectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRestoreTestingSelection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRestoreTestingSelection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRestoreTestingSelection",
	}
}

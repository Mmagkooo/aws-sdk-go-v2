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

// Creates a restore testing plan.
//
// The first of two steps to create a restore testing plan. After this request is
// successful, finish the procedure using CreateRestoreTestingSelection.
func (c *Client) CreateRestoreTestingPlan(ctx context.Context, params *CreateRestoreTestingPlanInput, optFns ...func(*Options)) (*CreateRestoreTestingPlanOutput, error) {
	if params == nil {
		params = &CreateRestoreTestingPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRestoreTestingPlan", params, optFns, c.addOperationCreateRestoreTestingPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRestoreTestingPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRestoreTestingPlanInput struct {

	// A restore testing plan must contain a unique RestoreTestingPlanName string you
	// create and must contain a ScheduleExpression cron. You may optionally include a
	// StartWindowHours integer and a CreatorRequestId string.
	//
	// The RestoreTestingPlanName is a unique string that is the name of the restore
	// testing plan. This cannot be changed after creation, and it must consist of only
	// alphanumeric characters and underscores.
	//
	// This member is required.
	RestoreTestingPlan *types.RestoreTestingPlanForCreate

	// This is a unique string that identifies the request and allows failed requests
	// to be retriedwithout the risk of running the operation twice. This parameter is
	// optional. If used, this parameter must contain 1 to 50 alphanumeric or '-_.'
	// characters.
	CreatorRequestId *string

	// The tags to assign to the restore testing plan.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateRestoreTestingPlanOutput struct {

	// The date and time a restore testing plan was created, in Unix format and
	// Coordinated Universal Time (UTC). The value of CreationTime is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087AM.
	//
	// This member is required.
	CreationTime *time.Time

	// An Amazon Resource Name (ARN) that uniquely identifies the created restore
	// testing plan.
	//
	// This member is required.
	RestoreTestingPlanArn *string

	// This unique string is the name of the restore testing plan.
	//
	// The name cannot be changed after creation. The name consists of only
	// alphanumeric characters and underscores. Maximum length is 50.
	//
	// This member is required.
	RestoreTestingPlanName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRestoreTestingPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRestoreTestingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRestoreTestingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRestoreTestingPlan"); err != nil {
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
	if err = addOpCreateRestoreTestingPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRestoreTestingPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRestoreTestingPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRestoreTestingPlan",
	}
}

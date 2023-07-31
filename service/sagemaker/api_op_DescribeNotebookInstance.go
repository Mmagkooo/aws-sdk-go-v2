// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns information about a notebook instance.
func (c *Client) DescribeNotebookInstance(ctx context.Context, params *DescribeNotebookInstanceInput, optFns ...func(*Options)) (*DescribeNotebookInstanceOutput, error) {
	if params == nil {
		params = &DescribeNotebookInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNotebookInstance", params, optFns, c.addOperationDescribeNotebookInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotebookInstanceInput struct {

	// The name of the notebook instance that you want information about.
	//
	// This member is required.
	NotebookInstanceName *string

	noSmithyDocumentSerde
}

type DescribeNotebookInstanceOutput struct {

	// A list of the Elastic Inference (EI) instance types associated with this
	// notebook instance. Currently only one EI instance type can be associated with a
	// notebook instance. For more information, see Using Elastic Inference in Amazon
	// SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	AcceleratorTypes []types.NotebookInstanceAcceleratorType

	// An array of up to three Git repositories associated with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your
	// account, or the URL of Git repositories in Amazon Web Services CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. These repositories are cloned at the same level
	// as the default repository of your notebook instance. For more information, see
	// Associating Git Repositories with SageMaker Notebook Instances (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html)
	// .
	AdditionalCodeRepositories []string

	// A timestamp. Use this parameter to return the time when the notebook instance
	// was created
	CreationTime *time.Time

	// The Git repository associated with the notebook instance as its default code
	// repository. This can be either the name of a Git repository stored as a resource
	// in your account, or the URL of a Git repository in Amazon Web Services
	// CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. When you open a notebook instance, it opens in
	// the directory that contains this repository. For more information, see
	// Associating Git Repositories with SageMaker Notebook Instances (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html)
	// .
	DefaultCodeRepository *string

	// Describes whether SageMaker provides internet access to the notebook instance.
	// If this value is set to Disabled, the notebook instance does not have internet
	// access, and cannot connect to SageMaker training and endpoint services. For more
	// information, see Notebook Instances Are Internet-Enabled by Default (https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access)
	// .
	DirectInternetAccess types.DirectInternetAccess

	// If status is Failed , the reason it failed.
	FailureReason *string

	// Information on the IMDS configuration of the notebook instance
	InstanceMetadataServiceConfiguration *types.InstanceMetadataServiceConfiguration

	// The type of ML compute instance running on the notebook instance.
	InstanceType types.InstanceType

	// The Amazon Web Services KMS key ID SageMaker uses to encrypt data when storing
	// it on the ML storage volume attached to the instance.
	KmsKeyId *string

	// A timestamp. Use this parameter to retrieve the time when the notebook instance
	// was last modified.
	LastModifiedTime *time.Time

	// The network interface IDs that SageMaker created at the time of creating the
	// instance.
	NetworkInterfaceId *string

	// The Amazon Resource Name (ARN) of the notebook instance.
	NotebookInstanceArn *string

	// Returns the name of a notebook instance lifecycle configuration. For
	// information about notebook instance lifestyle configurations, see Step 2.1:
	// (Optional) Customize a Notebook Instance (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html)
	NotebookInstanceLifecycleConfigName *string

	// The name of the SageMaker notebook instance.
	NotebookInstanceName *string

	// The status of the notebook instance.
	NotebookInstanceStatus types.NotebookInstanceStatus

	// The platform identifier of the notebook instance runtime environment.
	PlatformIdentifier *string

	// The Amazon Resource Name (ARN) of the IAM role associated with the instance.
	RoleArn *string

	// Whether root access is enabled or disabled for users of the notebook instance.
	// Lifecycle configurations need root access to be able to set up a notebook
	// instance. Because of this, lifecycle configurations associated with a notebook
	// instance always run with root access even if you disable root access for users.
	RootAccess types.RootAccess

	// The IDs of the VPC security groups.
	SecurityGroups []string

	// The ID of the VPC subnet.
	SubnetId *string

	// The URL that you use to connect to the Jupyter notebook that is running in your
	// notebook instance.
	Url *string

	// The size, in GB, of the ML storage volume attached to the notebook instance.
	VolumeSizeInGB *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNotebookInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNotebookInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNotebookInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addDescribeNotebookInstanceResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeNotebookInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotebookInstance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeNotebookInstanceAPIClient is a client that implements the
// DescribeNotebookInstance operation.
type DescribeNotebookInstanceAPIClient interface {
	DescribeNotebookInstance(context.Context, *DescribeNotebookInstanceInput, ...func(*Options)) (*DescribeNotebookInstanceOutput, error)
}

var _ DescribeNotebookInstanceAPIClient = (*Client)(nil)

// NotebookInstanceDeletedWaiterOptions are waiter options for
// NotebookInstanceDeletedWaiter
type NotebookInstanceDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// NotebookInstanceDeletedWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, NotebookInstanceDeletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeNotebookInstanceInput, *DescribeNotebookInstanceOutput, error) (bool, error)
}

// NotebookInstanceDeletedWaiter defines the waiters for NotebookInstanceDeleted
type NotebookInstanceDeletedWaiter struct {
	client DescribeNotebookInstanceAPIClient

	options NotebookInstanceDeletedWaiterOptions
}

// NewNotebookInstanceDeletedWaiter constructs a NotebookInstanceDeletedWaiter.
func NewNotebookInstanceDeletedWaiter(client DescribeNotebookInstanceAPIClient, optFns ...func(*NotebookInstanceDeletedWaiterOptions)) *NotebookInstanceDeletedWaiter {
	options := NotebookInstanceDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = notebookInstanceDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &NotebookInstanceDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for NotebookInstanceDeleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *NotebookInstanceDeletedWaiter) Wait(ctx context.Context, params *DescribeNotebookInstanceInput, maxWaitDur time.Duration, optFns ...func(*NotebookInstanceDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for NotebookInstanceDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *NotebookInstanceDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeNotebookInstanceInput, maxWaitDur time.Duration, optFns ...func(*NotebookInstanceDeletedWaiterOptions)) (*DescribeNotebookInstanceOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeNotebookInstance(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for NotebookInstanceDeleted waiter")
}

func notebookInstanceDeletedStateRetryable(ctx context.Context, input *DescribeNotebookInstanceInput, output *DescribeNotebookInstanceOutput, err error) (bool, error) {

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "ValidationException" == apiErr.ErrorCode() {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("NotebookInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.NotebookInstanceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.NotebookInstanceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// NotebookInstanceInServiceWaiterOptions are waiter options for
// NotebookInstanceInServiceWaiter
type NotebookInstanceInServiceWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// NotebookInstanceInServiceWaiter will use default minimum delay of 30 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, NotebookInstanceInServiceWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeNotebookInstanceInput, *DescribeNotebookInstanceOutput, error) (bool, error)
}

// NotebookInstanceInServiceWaiter defines the waiters for
// NotebookInstanceInService
type NotebookInstanceInServiceWaiter struct {
	client DescribeNotebookInstanceAPIClient

	options NotebookInstanceInServiceWaiterOptions
}

// NewNotebookInstanceInServiceWaiter constructs a NotebookInstanceInServiceWaiter.
func NewNotebookInstanceInServiceWaiter(client DescribeNotebookInstanceAPIClient, optFns ...func(*NotebookInstanceInServiceWaiterOptions)) *NotebookInstanceInServiceWaiter {
	options := NotebookInstanceInServiceWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = notebookInstanceInServiceStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &NotebookInstanceInServiceWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for NotebookInstanceInService waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *NotebookInstanceInServiceWaiter) Wait(ctx context.Context, params *DescribeNotebookInstanceInput, maxWaitDur time.Duration, optFns ...func(*NotebookInstanceInServiceWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for NotebookInstanceInService waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *NotebookInstanceInServiceWaiter) WaitForOutput(ctx context.Context, params *DescribeNotebookInstanceInput, maxWaitDur time.Duration, optFns ...func(*NotebookInstanceInServiceWaiterOptions)) (*DescribeNotebookInstanceOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeNotebookInstance(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for NotebookInstanceInService waiter")
}

func notebookInstanceInServiceStateRetryable(ctx context.Context, input *DescribeNotebookInstanceInput, output *DescribeNotebookInstanceOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("NotebookInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "InService"
		value, ok := pathValue.(types.NotebookInstanceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.NotebookInstanceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("NotebookInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.NotebookInstanceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.NotebookInstanceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// NotebookInstanceStoppedWaiterOptions are waiter options for
// NotebookInstanceStoppedWaiter
type NotebookInstanceStoppedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// NotebookInstanceStoppedWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, NotebookInstanceStoppedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeNotebookInstanceInput, *DescribeNotebookInstanceOutput, error) (bool, error)
}

// NotebookInstanceStoppedWaiter defines the waiters for NotebookInstanceStopped
type NotebookInstanceStoppedWaiter struct {
	client DescribeNotebookInstanceAPIClient

	options NotebookInstanceStoppedWaiterOptions
}

// NewNotebookInstanceStoppedWaiter constructs a NotebookInstanceStoppedWaiter.
func NewNotebookInstanceStoppedWaiter(client DescribeNotebookInstanceAPIClient, optFns ...func(*NotebookInstanceStoppedWaiterOptions)) *NotebookInstanceStoppedWaiter {
	options := NotebookInstanceStoppedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = notebookInstanceStoppedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &NotebookInstanceStoppedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for NotebookInstanceStopped waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *NotebookInstanceStoppedWaiter) Wait(ctx context.Context, params *DescribeNotebookInstanceInput, maxWaitDur time.Duration, optFns ...func(*NotebookInstanceStoppedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for NotebookInstanceStopped waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *NotebookInstanceStoppedWaiter) WaitForOutput(ctx context.Context, params *DescribeNotebookInstanceInput, maxWaitDur time.Duration, optFns ...func(*NotebookInstanceStoppedWaiterOptions)) (*DescribeNotebookInstanceOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeNotebookInstance(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for NotebookInstanceStopped waiter")
}

func notebookInstanceStoppedStateRetryable(ctx context.Context, input *DescribeNotebookInstanceInput, output *DescribeNotebookInstanceOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("NotebookInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Stopped"
		value, ok := pathValue.(types.NotebookInstanceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.NotebookInstanceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("NotebookInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.NotebookInstanceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.NotebookInstanceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeNotebookInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeNotebookInstance",
	}
}

type opDescribeNotebookInstanceResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeNotebookInstanceResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeNotebookInstanceResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "sagemaker"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "sagemaker"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("sagemaker")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeNotebookInstanceResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeNotebookInstanceResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

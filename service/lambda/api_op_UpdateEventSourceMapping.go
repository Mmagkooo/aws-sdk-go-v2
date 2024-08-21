// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an event source mapping. You can change the function that Lambda
// invokes, or pause invocation and resume later from the same location.
//
// For details about how to configure different event sources, see the following
// topics.
//
// [Amazon DynamoDB Streams]
//
// [Amazon Kinesis]
//
// [Amazon SQS]
//
// [Amazon MQ and RabbitMQ]
//
// [Amazon MSK]
//
// [Apache Kafka]
//
// [Amazon DocumentDB]
//
// The following error handling options are available only for stream sources
// (DynamoDB and Kinesis):
//
//   - BisectBatchOnFunctionError – If the function returns an error, split the
//     batch in two and retry.
//
//   - DestinationConfig – Send discarded records to an Amazon SQS queue or Amazon
//     SNS topic.
//
//   - MaximumRecordAgeInSeconds – Discard records older than the specified age.
//     The default value is infinite (-1). When set to infinite (-1), failed records
//     are retried until the record expires
//
//   - MaximumRetryAttempts – Discard records after the specified number of
//     retries. The default value is infinite (-1). When set to infinite (-1), failed
//     records are retried until the record expires.
//
//   - ParallelizationFactor – Process multiple batches from each shard
//     concurrently.
//
// For information about which configuration parameters apply to each event
// source, see the following topics.
//
// [Amazon DynamoDB Streams]
//
// [Amazon Kinesis]
//
// [Amazon SQS]
//
// [Amazon MQ and RabbitMQ]
//
// [Amazon MSK]
//
// [Apache Kafka]
//
// [Amazon DocumentDB]
//
// [Amazon DynamoDB Streams]: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-params
// [Amazon SQS]: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#services-sqs-params
// [Amazon MSK]: https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-parms
// [Amazon Kinesis]: https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-params
// [Amazon MQ and RabbitMQ]: https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-params
// [Apache Kafka]: https://docs.aws.amazon.com/lambda/latest/dg/with-kafka.html#services-kafka-parms
// [Amazon DocumentDB]: https://docs.aws.amazon.com/lambda/latest/dg/with-documentdb.html#docdb-configuration
func (c *Client) UpdateEventSourceMapping(ctx context.Context, params *UpdateEventSourceMappingInput, optFns ...func(*Options)) (*UpdateEventSourceMappingOutput, error) {
	if params == nil {
		params = &UpdateEventSourceMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEventSourceMapping", params, optFns, c.addOperationUpdateEventSourceMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEventSourceMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEventSourceMappingInput struct {

	// The identifier of the event source mapping.
	//
	// This member is required.
	UUID *string

	// The maximum number of records in each batch that Lambda pulls from your stream
	// or queue and sends to your function. Lambda passes all of the records in the
	// batch to the function in a single call, up to the payload limit for synchronous
	// invocation (6 MB).
	//
	//   - Amazon Kinesis – Default 100. Max 10,000.
	//
	//   - Amazon DynamoDB Streams – Default 100. Max 10,000.
	//
	//   - Amazon Simple Queue Service – Default 10. For standard queues the max is
	//   10,000. For FIFO queues the max is 10.
	//
	//   - Amazon Managed Streaming for Apache Kafka – Default 100. Max 10,000.
	//
	//   - Self-managed Apache Kafka – Default 100. Max 10,000.
	//
	//   - Amazon MQ (ActiveMQ and RabbitMQ) – Default 100. Max 10,000.
	//
	//   - DocumentDB – Default 100. Max 10,000.
	BatchSize *int32

	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the
	// batch in two and retry.
	BisectBatchOnFunctionError *bool

	// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Kafka only) A
	// configuration object that specifies the destination of an event after Lambda
	// processes it.
	DestinationConfig *types.DestinationConfig

	// Specific configuration settings for a DocumentDB event source.
	DocumentDBEventSourceConfig *types.DocumentDBEventSourceConfig

	// When true, the event source mapping is active. When false, Lambda pauses
	// polling and invocation.
	//
	// Default: True
	Enabled *bool

	// An object that defines the filter criteria that determine whether Lambda should
	// process an event. For more information, see [Lambda event filtering].
	//
	// [Lambda event filtering]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html
	FilterCriteria *types.FilterCriteria

	// The name or ARN of the Lambda function.
	//
	// Name formats
	//
	//   - Function name – MyFunction .
	//
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:MyFunction .
	//
	//   - Version or Alias ARN –
	//   arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD .
	//
	//   - Partial ARN – 123456789012:function:MyFunction .
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it's limited to 64 characters in length.
	FunctionName *string

	// (Kinesis, DynamoDB Streams, and Amazon SQS) A list of current response type
	// enums applied to the event source mapping.
	FunctionResponseTypes []types.FunctionResponseType

	//  The ARN of the Key Management Service (KMS) customer managed key that Lambda
	// uses to encrypt your function's [filter criteria]. By default, Lambda does not encrypt your
	// filter criteria object. Specify this property to encrypt data using your own
	// customer managed key.
	//
	// [filter criteria]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics
	KMSKeyArn *string

	// The maximum amount of time, in seconds, that Lambda spends gathering records
	// before invoking the function. You can configure MaximumBatchingWindowInSeconds
	// to any value from 0 seconds to 300 seconds in increments of seconds.
	//
	// For Kinesis, DynamoDB, and Amazon SQS event sources, the default batching
	// window is 0 seconds. For Amazon MSK, Self-managed Apache Kafka, Amazon MQ, and
	// DocumentDB event sources, the default batching window is 500 ms. Note that
	// because you can only change MaximumBatchingWindowInSeconds in increments of
	// seconds, you cannot revert back to the 500 ms default batching window after you
	// have changed it. To restore the default batching window, you must create a new
	// event source mapping.
	//
	// Related setting: For Kinesis, DynamoDB, and Amazon SQS event sources, when you
	// set BatchSize to a value greater than 10, you must set
	// MaximumBatchingWindowInSeconds to at least 1.
	MaximumBatchingWindowInSeconds *int32

	// (Kinesis and DynamoDB Streams only) Discard records older than the specified
	// age. The default value is infinite (-1).
	MaximumRecordAgeInSeconds *int32

	// (Kinesis and DynamoDB Streams only) Discard records after the specified number
	// of retries. The default value is infinite (-1). When set to infinite (-1),
	// failed records are retried until the record expires.
	MaximumRetryAttempts *int32

	// (Kinesis and DynamoDB Streams only) The number of batches to process from each
	// shard concurrently.
	ParallelizationFactor *int32

	// (Amazon SQS only) The scaling configuration for the event source. For more
	// information, see [Configuring maximum concurrency for Amazon SQS event sources].
	//
	// [Configuring maximum concurrency for Amazon SQS event sources]: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency
	ScalingConfig *types.ScalingConfig

	// An array of authentication protocols or VPC components required to secure your
	// event source.
	SourceAccessConfigurations []types.SourceAccessConfiguration

	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing
	// window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds
	// indicates no tumbling window.
	TumblingWindowInSeconds *int32

	noSmithyDocumentSerde
}

// A mapping between an Amazon Web Services resource and a Lambda function. For
// details, see CreateEventSourceMapping.
type UpdateEventSourceMappingOutput struct {

	// Specific configuration settings for an Amazon Managed Streaming for Apache
	// Kafka (Amazon MSK) event source.
	AmazonManagedKafkaEventSourceConfig *types.AmazonManagedKafkaEventSourceConfig

	// The maximum number of records in each batch that Lambda pulls from your stream
	// or queue and sends to your function. Lambda passes all of the records in the
	// batch to the function in a single call, up to the payload limit for synchronous
	// invocation (6 MB).
	//
	// Default value: Varies by service. For Amazon SQS, the default is 10. For all
	// other services, the default is 100.
	//
	// Related setting: When you set BatchSize to a value greater than 10, you must
	// set MaximumBatchingWindowInSeconds to at least 1.
	BatchSize *int32

	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the
	// batch in two and retry. The default value is false.
	BisectBatchOnFunctionError *bool

	// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka event
	// sources only) A configuration object that specifies the destination of an event
	// after Lambda processes it.
	DestinationConfig *types.DestinationConfig

	// Specific configuration settings for a DocumentDB event source.
	DocumentDBEventSourceConfig *types.DocumentDBEventSourceConfig

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string

	// An object that defines the filter criteria that determine whether Lambda should
	// process an event. For more information, see [Lambda event filtering].
	//
	// If filter criteria is encrypted, this field shows up as null in the response of
	// ListEventSourceMapping API calls. You can view this field in plaintext in the
	// response of GetEventSourceMapping and DeleteEventSourceMapping calls if you have
	// kms:Decrypt permissions for the correct KMS key.
	//
	// [Lambda event filtering]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html
	FilterCriteria *types.FilterCriteria

	// An object that contains details about an error related to filter criteria
	// encryption.
	FilterCriteriaError *types.FilterCriteriaError

	// The ARN of the Lambda function.
	FunctionArn *string

	// (Kinesis, DynamoDB Streams, and Amazon SQS) A list of current response type
	// enums applied to the event source mapping.
	FunctionResponseTypes []types.FunctionResponseType

	//  The ARN of the Key Management Service (KMS) customer managed key that Lambda
	// uses to encrypt your function's [filter criteria].
	//
	// [filter criteria]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics
	KMSKeyArn *string

	// The date that the event source mapping was last updated or that its state
	// changed.
	LastModified *time.Time

	// The result of the last Lambda invocation of your function.
	LastProcessingResult *string

	// The maximum amount of time, in seconds, that Lambda spends gathering records
	// before invoking the function. You can configure MaximumBatchingWindowInSeconds
	// to any value from 0 seconds to 300 seconds in increments of seconds.
	//
	// For streams and Amazon SQS event sources, the default batching window is 0
	// seconds. For Amazon MSK, Self-managed Apache Kafka, Amazon MQ, and DocumentDB
	// event sources, the default batching window is 500 ms. Note that because you can
	// only change MaximumBatchingWindowInSeconds in increments of seconds, you cannot
	// revert back to the 500 ms default batching window after you have changed it. To
	// restore the default batching window, you must create a new event source mapping.
	//
	// Related setting: For streams and Amazon SQS event sources, when you set
	// BatchSize to a value greater than 10, you must set
	// MaximumBatchingWindowInSeconds to at least 1.
	MaximumBatchingWindowInSeconds *int32

	// (Kinesis and DynamoDB Streams only) Discard records older than the specified
	// age. The default value is -1, which sets the maximum age to infinite. When the
	// value is set to infinite, Lambda never discards old records.
	//
	// The minimum valid value for maximum record age is 60s. Although values less
	// than 60 and greater than -1 fall within the parameter's absolute range, they are
	// not allowed
	MaximumRecordAgeInSeconds *int32

	// (Kinesis and DynamoDB Streams only) Discard records after the specified number
	// of retries. The default value is -1, which sets the maximum number of retries to
	// infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records
	// until the record expires in the event source.
	MaximumRetryAttempts *int32

	// (Kinesis and DynamoDB Streams only) The number of batches to process
	// concurrently from each shard. The default value is 1.
	ParallelizationFactor *int32

	//  (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	Queues []string

	// (Amazon SQS only) The scaling configuration for the event source. For more
	// information, see [Configuring maximum concurrency for Amazon SQS event sources].
	//
	// [Configuring maximum concurrency for Amazon SQS event sources]: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency
	ScalingConfig *types.ScalingConfig

	// The self-managed Apache Kafka cluster for your event source.
	SelfManagedEventSource *types.SelfManagedEventSource

	// Specific configuration settings for a self-managed Apache Kafka event source.
	SelfManagedKafkaEventSourceConfig *types.SelfManagedKafkaEventSourceConfig

	// An array of the authentication protocol, VPC components, or virtual host to
	// secure and define your event source.
	SourceAccessConfigurations []types.SourceAccessConfiguration

	// The position in a stream from which to start reading. Required for Amazon
	// Kinesis and Amazon DynamoDB Stream event sources. AT_TIMESTAMP is supported
	// only for Amazon Kinesis streams, Amazon DocumentDB, Amazon MSK, and self-managed
	// Apache Kafka.
	StartingPosition types.EventSourcePosition

	// With StartingPosition set to AT_TIMESTAMP , the time from which to start
	// reading. StartingPositionTimestamp cannot be in the future.
	StartingPositionTimestamp *time.Time

	// The state of the event source mapping. It can be one of the following: Creating
	// , Enabling , Enabled , Disabling , Disabled , Updating , or Deleting .
	State *string

	// Indicates whether a user or Lambda made the last change to the event source
	// mapping.
	StateTransitionReason *string

	// The name of the Kafka topic.
	Topics []string

	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing
	// window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds
	// indicates no tumbling window.
	TumblingWindowInSeconds *int32

	// The identifier of the event source mapping.
	UUID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEventSourceMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEventSourceMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEventSourceMapping{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEventSourceMapping"); err != nil {
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
	if err = addOpUpdateEventSourceMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEventSourceMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEventSourceMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEventSourceMapping",
	}
}

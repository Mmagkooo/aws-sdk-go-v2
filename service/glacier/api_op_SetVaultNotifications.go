// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation configures notifications that will be sent when specific events
// happen to a vault. By default, you don't get any notifications. To configure
// vault notifications, send a PUT request to the notification-configuration
// subresource of the vault. The request should include a JSON document that
// provides an Amazon SNS topic and specific events for which you want Amazon S3
// Glacier to send notifications to the topic. Amazon SNS topics must grant
// permission to the vault to be allowed to publish notifications to the topic. You
// can configure a vault to publish a notification for the following vault events:
//   - ArchiveRetrievalCompleted This event occurs when a job that was initiated
//     for an archive retrieval is completed ( InitiateJob ). The status of the
//     completed job can be "Succeeded" or "Failed". The notification sent to the SNS
//     topic is the same output as returned from DescribeJob .
//   - InventoryRetrievalCompleted This event occurs when a job that was initiated
//     for an inventory retrieval is completed ( InitiateJob ). The status of the
//     completed job can be "Succeeded" or "Failed". The notification sent to the SNS
//     topic is the same output as returned from DescribeJob .
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see Access Control Using AWS Identity
// and Access Management (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html)
// . For conceptual information and underlying REST API, see Configuring Vault
// Notifications in Amazon S3 Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html)
// and Set Vault Notification Configuration  (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-put.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) SetVaultNotifications(ctx context.Context, params *SetVaultNotificationsInput, optFns ...func(*Options)) (*SetVaultNotificationsOutput, error) {
	if params == nil {
		params = &SetVaultNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetVaultNotifications", params, optFns, c.addOperationSetVaultNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetVaultNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to configure notifications that will be sent when specific
// events happen to a vault.
type SetVaultNotificationsInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single ' - ' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// Provides options for specifying notification configuration.
	VaultNotificationConfig *types.VaultNotificationConfig

	noSmithyDocumentSerde
}

type SetVaultNotificationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetVaultNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSetVaultNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSetVaultNotifications{}, middleware.After)
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
	if err = addSetVaultNotificationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSetVaultNotificationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetVaultNotifications(options.Region), middleware.Before); err != nil {
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
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
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

func newServiceMetadataMiddleware_opSetVaultNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "SetVaultNotifications",
	}
}

type opSetVaultNotificationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSetVaultNotificationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSetVaultNotificationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "glacier"
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
				signingName = "glacier"
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
				v4aScheme.SigningName = aws.String("glacier")
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

func addSetVaultNotificationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSetVaultNotificationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

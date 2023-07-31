// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Links an existing user account in a user pool ( DestinationUser ) to an identity
// from an external IdP ( SourceUser ) based on a specified attribute name and
// value from the external IdP. This allows you to create a link from the existing
// user account to an external federated user identity that has not yet been used
// to sign in. You can then use the federated user identity to sign in as the
// existing user account. For example, if there is an existing user with a username
// and password, this API links that user to a federated user identity. When the
// user signs in with a federated user identity, they sign in as the existing user
// account. The maximum number of federated identities linked to a user is five.
// Because this API allows a user with an external federated identity to sign in as
// an existing user in the user pool, it is critical that it only be used with
// external IdPs and provider attributes that have been trusted by the application
// owner. This action is administrative and requires developer credentials.
func (c *Client) AdminLinkProviderForUser(ctx context.Context, params *AdminLinkProviderForUserInput, optFns ...func(*Options)) (*AdminLinkProviderForUserOutput, error) {
	if params == nil {
		params = &AdminLinkProviderForUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminLinkProviderForUser", params, optFns, c.addOperationAdminLinkProviderForUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminLinkProviderForUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminLinkProviderForUserInput struct {

	// The existing user in the user pool that you want to assign to the external IdP
	// user account. This user can be a native (Username + Password) Amazon Cognito
	// user pools user or a federated user (for example, a SAML or Facebook user). If
	// the user doesn't exist, Amazon Cognito generates an exception. Amazon Cognito
	// returns this user when the new user (with the linked IdP attribute) signs in.
	// For a native username + password user, the ProviderAttributeValue for the
	// DestinationUser should be the username in the user pool. For a federated user,
	// it should be the provider-specific user_id . The ProviderAttributeName of the
	// DestinationUser is ignored. The ProviderName should be set to Cognito for users
	// in Cognito user pools. All attributes in the DestinationUser profile must be
	// mutable. If you have assigned the user any immutable custom attributes, the
	// operation won't succeed.
	//
	// This member is required.
	DestinationUser *types.ProviderUserIdentifierType

	// An external IdP account for a user who doesn't exist yet in the user pool. This
	// user must be a federated user (for example, a SAML or Facebook user), not
	// another native user. If the SourceUser is using a federated social IdP, such as
	// Facebook, Google, or Login with Amazon, you must set the ProviderAttributeName
	// to Cognito_Subject . For social IdPs, the ProviderName will be Facebook , Google
	// , or LoginWithAmazon , and Amazon Cognito will automatically parse the Facebook,
	// Google, and Login with Amazon tokens for id , sub , and user_id , respectively.
	// The ProviderAttributeValue for the user must be the same value as the id , sub ,
	// or user_id value found in the social IdP token. For SAML, the
	// ProviderAttributeName can be any value that matches a claim in the SAML
	// assertion. If you want to link SAML users based on the subject of the SAML
	// assertion, you should map the subject to a claim through the SAML IdP and submit
	// that claim name as the ProviderAttributeName . If you set ProviderAttributeName
	// to Cognito_Subject , Amazon Cognito will automatically parse the default unique
	// identifier found in the subject from the SAML token.
	//
	// This member is required.
	SourceUser *types.ProviderUserIdentifierType

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	noSmithyDocumentSerde
}

type AdminLinkProviderForUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminLinkProviderForUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminLinkProviderForUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminLinkProviderForUser{}, middleware.After)
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
	if err = addAdminLinkProviderForUserResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAdminLinkProviderForUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminLinkProviderForUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminLinkProviderForUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminLinkProviderForUser",
	}
}

type opAdminLinkProviderForUserResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opAdminLinkProviderForUserResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opAdminLinkProviderForUserResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "cognito-idp"
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
				signingName = "cognito-idp"
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
				v4aScheme.SigningName = aws.String("cognito-idp")
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

func addAdminLinkProviderForUserResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opAdminLinkProviderForUserResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a case in the Amazon Web Services Support Center. This operation is
// similar to how you create a case in the Amazon Web Services Support Center
// Create Case (https://console.aws.amazon.com/support/home#/case/create) page. The
// Amazon Web Services Support API doesn't support requesting service limit
// increases. You can submit a service limit increase in the following ways:
//   - Submit a request from the Amazon Web Services Support Center Create Case (https://console.aws.amazon.com/support/home#/case/create)
//     page.
//   - Use the Service Quotas RequestServiceQuotaIncrease (https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_RequestServiceQuotaIncrease.html)
//     operation.
//
// A successful CreateCase request returns an Amazon Web Services Support case
// number. You can use the DescribeCases operation and specify the case number to
// get existing Amazon Web Services Support cases. After you create a case, use the
// AddCommunicationToCase operation to add additional communication or attachments
// to an existing case. The caseId is separate from the displayId that appears in
// the Amazon Web Services Support Center (https://console.aws.amazon.com/support)
// . Use the DescribeCases operation to get the displayId .
//   - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
//     use the Amazon Web Services Support API.
//   - If you call the Amazon Web Services Support API from an account that
//     doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
//     SubscriptionRequiredException error message appears. For information about
//     changing your support plan, see Amazon Web Services Support (http://aws.amazon.com/premiumsupport/)
//     .
func (c *Client) CreateCase(ctx context.Context, params *CreateCaseInput, optFns ...func(*Options)) (*CreateCaseOutput, error) {
	if params == nil {
		params = &CreateCaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCase", params, optFns, c.addOperationCreateCaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCaseInput struct {

	// The communication body text that describes the issue. This text appears in the
	// Description field on the Amazon Web Services Support Center Create Case (https://console.aws.amazon.com/support/home#/case/create)
	// page.
	//
	// This member is required.
	CommunicationBody *string

	// The title of the support case. The title appears in the Subject field on the
	// Amazon Web Services Support Center Create Case (https://console.aws.amazon.com/support/home#/case/create)
	// page.
	//
	// This member is required.
	Subject *string

	// The ID of a set of one or more attachments for the case. Create the set by
	// using the AddAttachmentsToSet operation.
	AttachmentSetId *string

	// The category of problem for the support case. You also use the DescribeServices
	// operation to get the category code for a service. Each Amazon Web Services
	// service defines its own set of category codes.
	CategoryCode *string

	// A list of email addresses that Amazon Web Services Support copies on case
	// correspondence. Amazon Web Services Support identifies the account that creates
	// the case when you specify your Amazon Web Services credentials in an HTTP POST
	// method or use the Amazon Web Services SDKs (http://aws.amazon.com/tools/) .
	CcEmailAddresses []string

	// The type of issue for the case. You can specify customer-service or technical .
	// If you don't specify a value, the default is technical .
	IssueType *string

	// The language in which Amazon Web Services Support handles the case. Amazon Web
	// Services Support currently supports Chinese (“zh”), English ("en"), Japanese
	// ("ja") and Korean (“ko”). You must specify the ISO 639-1 code for the language
	// parameter if you want support in that language.
	Language *string

	// The code for the Amazon Web Services service. You can use the DescribeServices
	// operation to get the possible serviceCode values.
	ServiceCode *string

	// A value that indicates the urgency of the case. This value determines the
	// response time according to your service level agreement with Amazon Web Services
	// Support. You can use the DescribeSeverityLevels operation to get the possible
	// values for severityCode . For more information, see SeverityLevel and Choosing
	// a Severity (https://docs.aws.amazon.com/awssupport/latest/user/getting-started.html#choosing-severity)
	// in the Amazon Web Services Support User Guide. The availability of severity
	// levels depends on the support plan for the Amazon Web Services account.
	SeverityCode *string

	noSmithyDocumentSerde
}

// The support case ID returned by a successful completion of the CreateCase
// operation.
type CreateCaseOutput struct {

	// The support case ID requested or returned in the call. The case ID is an
	// alphanumeric string in the following format:
	// case-12345678910-2013-c4c1d2bf33c5cf47
	CaseId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCase{}, middleware.After)
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
	if err = addCreateCaseResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "CreateCase",
	}
}

type opCreateCaseResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateCaseResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateCaseResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "support"
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
				signingName = "support"
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
				v4aScheme.SigningName = aws.String("support")
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

func addCreateCaseResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateCaseResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

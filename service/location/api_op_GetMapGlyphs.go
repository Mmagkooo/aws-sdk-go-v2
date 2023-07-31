// Code generated by smithy-go-codegen DO NOT EDIT.

package location

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

// Retrieves glyphs used to display labels on a map.
func (c *Client) GetMapGlyphs(ctx context.Context, params *GetMapGlyphsInput, optFns ...func(*Options)) (*GetMapGlyphsOutput, error) {
	if params == nil {
		params = &GetMapGlyphsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMapGlyphs", params, optFns, c.addOperationGetMapGlyphsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMapGlyphsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMapGlyphsInput struct {

	// A comma-separated list of fonts to load glyphs from in order of preference. For
	// example, Noto Sans Regular, Arial Unicode . Valid fonts stacks for Esri (https://docs.aws.amazon.com/location/latest/developerguide/esri.html)
	// styles:
	//   - VectorEsriDarkGrayCanvas – Ubuntu Medium Italic | Ubuntu Medium | Ubuntu
	//   Italic | Ubuntu Regular | Ubuntu Bold
	//   - VectorEsriLightGrayCanvas – Ubuntu Italic | Ubuntu Regular | Ubuntu Light |
	//   Ubuntu Bold
	//   - VectorEsriTopographic – Noto Sans Italic | Noto Sans Regular | Noto Sans
	//   Bold | Noto Serif Regular | Roboto Condensed Light Italic
	//   - VectorEsriStreets – Arial Regular | Arial Italic | Arial Bold
	//   - VectorEsriNavigation – Arial Regular | Arial Italic | Arial Bold
	// Valid font stacks for HERE Technologies (https://docs.aws.amazon.com/location/latest/developerguide/HERE.html)
	// styles:
	//   - VectorHereContrast – Fira GO Regular | Fira GO Bold
	//   - VectorHereExplore, VectorHereExploreTruck, HybridHereExploreSatellite –
	//   Fira GO Italic | Fira GO Map | Fira GO Map Bold | Noto Sans CJK JP Bold |
	//   Noto Sans CJK JP Light | Noto Sans CJK JP Regular
	// Valid font stacks for GrabMaps (https://docs.aws.amazon.com/location/latest/developerguide/grab.html)
	// styles:
	//   - VectorGrabStandardLight, VectorGrabStandardDark – Noto Sans Regular | Noto
	//   Sans Medium | Noto Sans Bold
	// Valid font stacks for Open Data (https://docs.aws.amazon.com/location/latest/developerguide/open-data.html)
	// styles:
	//   - VectorOpenDataStandardLight, VectorOpenDataStandardDark,
	//   VectorOpenDataVisualizationLight, VectorOpenDataVisualizationDark – Amazon
	//   Ember Regular,Noto Sans Regular | Amazon Ember Bold,Noto Sans Bold | Amazon
	//   Ember Medium,Noto Sans Medium | Amazon Ember Regular Italic,Noto Sans Italic |
	//   Amazon Ember Condensed RC Regular,Noto Sans Regular | Amazon Ember Condensed
	//   RC Bold,Noto Sans Bold | Amazon Ember Regular,Noto Sans Regular,Noto Sans
	//   Arabic Regular | Amazon Ember Condensed RC Bold,Noto Sans Bold,Noto Sans
	//   Arabic Condensed Bold | Amazon Ember Bold,Noto Sans Bold,Noto Sans Arabic Bold
	//   | Amazon Ember Regular Italic,Noto Sans Italic,Noto Sans Arabic Regular |
	//   Amazon Ember Condensed RC Regular,Noto Sans Regular,Noto Sans Arabic Condensed
	//   Regular | Amazon Ember Medium,Noto Sans Medium,Noto Sans Arabic Medium
	// The fonts used by the Open Data map styles are combined fonts that use Amazon
	// Ember for most glyphs but Noto Sans for glyphs unsupported by Amazon Ember .
	//
	// This member is required.
	FontStack *string

	// A Unicode range of characters to download glyphs for. Each response will
	// contain 256 characters. For example, 0–255 includes all characters from range
	// U+0000 to 00FF . Must be aligned to multiples of 256.
	//
	// This member is required.
	FontUnicodeRange *string

	// The map resource associated with the glyph ﬁle.
	//
	// This member is required.
	MapName *string

	// The optional API key (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
	// to authorize the request.
	Key *string

	noSmithyDocumentSerde
}

type GetMapGlyphsOutput struct {

	// The glyph, as binary blob.
	Blob []byte

	// The HTTP Cache-Control directive for the value.
	CacheControl *string

	// The map glyph content type. For example, application/octet-stream .
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMapGlyphsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMapGlyphs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMapGlyphs{}, middleware.After)
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
	if err = addEndpointPrefix_opGetMapGlyphsMiddleware(stack); err != nil {
		return err
	}
	if err = addGetMapGlyphsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetMapGlyphsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMapGlyphs(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetMapGlyphsMiddleware struct {
}

func (*endpointPrefix_opGetMapGlyphsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMapGlyphsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "maps." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetMapGlyphsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetMapGlyphsMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetMapGlyphs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "GetMapGlyphs",
	}
}

type opGetMapGlyphsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetMapGlyphsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetMapGlyphsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "geo"
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
				signingName = "geo"
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
				v4aScheme.SigningName = aws.String("geo")
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

func addGetMapGlyphsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetMapGlyphsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints/awsrulesfn"
	internalendpoints "github.com/aws/aws-sdk-go-v2/service/route53/internal/endpoints"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/http"
	"net/url"
	"strings"
)

// EndpointResolverOptions is the service endpoint resolver options
type EndpointResolverOptions = internalendpoints.Options

// EndpointResolver interface for resolving service endpoints.
type EndpointResolver interface {
	ResolveEndpoint(region string, options EndpointResolverOptions) (aws.Endpoint, error)
}

var _ EndpointResolver = &internalendpoints.Resolver{}

// NewDefaultEndpointResolver constructs a new service endpoint resolver
func NewDefaultEndpointResolver() *internalendpoints.Resolver {
	return internalendpoints.New()
}

// EndpointResolverFunc is a helper utility that wraps a function so it satisfies
// the EndpointResolver interface. This is useful when you want to add additional
// endpoint resolving logic, or stub out specific endpoints with custom values.
type EndpointResolverFunc func(region string, options EndpointResolverOptions) (aws.Endpoint, error)

func (fn EndpointResolverFunc) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	return fn(region, options)
}

// EndpointResolverFromURL returns an EndpointResolver configured using the
// provided endpoint url. By default, the resolved endpoint resolver uses the
// client region as signing region, and the endpoint source is set to
// EndpointSourceCustom.You can provide functional options to configure endpoint
// values for the resolved endpoint.
func EndpointResolverFromURL(url string, optFns ...func(*aws.Endpoint)) EndpointResolver {
	e := aws.Endpoint{URL: url, Source: aws.EndpointSourceCustom}
	for _, fn := range optFns {
		fn(&e)
	}

	return EndpointResolverFunc(
		func(region string, options EndpointResolverOptions) (aws.Endpoint, error) {
			if len(e.SigningRegion) == 0 {
				e.SigningRegion = region
			}
			return e, nil
		},
	)
}

type ResolveEndpoint struct {
	Resolver EndpointResolver
	Options  EndpointResolverOptions
}

func (*ResolveEndpoint) ID() string {
	return "ResolveEndpoint"
}

func (m *ResolveEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if !awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.Resolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	eo := m.Options
	eo.Logger = middleware.GetLogger(ctx)

	var endpoint aws.Endpoint
	endpoint, err = m.Resolver.ResolveEndpoint(awsmiddleware.GetRegion(ctx), eo)
	if err != nil {
		nf := (&aws.EndpointNotFoundError{})
		if errors.As(err, &nf) {
			ctx = awsmiddleware.SetRequiresLegacyEndpoints(ctx, false)
			return next.HandleSerialize(ctx, in)
		}
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL, err = url.Parse(endpoint.URL)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
	}

	if len(awsmiddleware.GetSigningName(ctx)) == 0 {
		signingName := endpoint.SigningName
		if len(signingName) == 0 {
			signingName = "route53"
		}
		ctx = awsmiddleware.SetSigningName(ctx, signingName)
	}
	ctx = awsmiddleware.SetEndpointSource(ctx, endpoint.Source)
	ctx = smithyhttp.SetHostnameImmutable(ctx, endpoint.HostnameImmutable)
	ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)
	ctx = awsmiddleware.SetPartitionID(ctx, endpoint.PartitionID)
	return next.HandleSerialize(ctx, in)
}
func addResolveEndpointMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&ResolveEndpoint{
		Resolver: o.EndpointResolver,
		Options:  o.EndpointOptions,
	}, "OperationSerializer", middleware.Before)
}

func removeResolveEndpointMiddleware(stack *middleware.Stack) error {
	_, err := stack.Serialize.Remove((&ResolveEndpoint{}).ID())
	return err
}

type wrappedEndpointResolver struct {
	awsResolver aws.EndpointResolverWithOptions
}

func (w *wrappedEndpointResolver) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	return w.awsResolver.ResolveEndpoint(ServiceID, region, options)
}

type awsEndpointResolverAdaptor func(service, region string) (aws.Endpoint, error)

func (a awsEndpointResolverAdaptor) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return a(service, region)
}

var _ aws.EndpointResolverWithOptions = awsEndpointResolverAdaptor(nil)

// withEndpointResolver returns an aws.EndpointResolverWithOptions that first delegates endpoint resolution to the awsResolver.
// If awsResolver returns aws.EndpointNotFoundError error, the v1 resolver middleware will swallow the error,
// and set an appropriate context flag such that fallback will occur when EndpointResolverV2 is invoked
// via its middleware.
//
// If another error (besides aws.EndpointNotFoundError) is returned, then that error will be propagated.
func withEndpointResolver(awsResolver aws.EndpointResolver, awsResolverWithOptions aws.EndpointResolverWithOptions) EndpointResolver {
	var resolver aws.EndpointResolverWithOptions

	if awsResolverWithOptions != nil {
		resolver = awsResolverWithOptions
	} else if awsResolver != nil {
		resolver = awsEndpointResolverAdaptor(awsResolver.ResolveEndpoint)
	}

	return &wrappedEndpointResolver{
		awsResolver: resolver,
	}
}

func finalizeClientEndpointResolverOptions(options *Options) {
	options.EndpointOptions.LogDeprecated = options.ClientLogMode.IsDeprecatedUsage()

	if len(options.EndpointOptions.ResolvedRegion) == 0 {
		const fipsInfix = "-fips-"
		const fipsPrefix = "fips-"
		const fipsSuffix = "-fips"

		if strings.Contains(options.Region, fipsInfix) ||
			strings.Contains(options.Region, fipsPrefix) ||
			strings.Contains(options.Region, fipsSuffix) {
			options.EndpointOptions.ResolvedRegion = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
				options.Region, fipsInfix, "-"), fipsPrefix, ""), fipsSuffix, "")
			options.EndpointOptions.UseFIPSEndpoint = aws.FIPSEndpointStateEnabled
		}
	}

}

func resolveEndpointResolverV2(options *Options) {
	if options.EndpointResolverV2 == nil {
		options.EndpointResolverV2 = NewDefaultEndpointResolverV2()
	}
}

// Utility function to aid with translating pseudo-regions to classical regions
// with the appropriate setting indicated by the pseudo-region
func mapPseudoRegion(pr string) (region string, fips aws.FIPSEndpointState) {
	const fipsInfix = "-fips-"
	const fipsPrefix = "fips-"
	const fipsSuffix = "-fips"

	if strings.Contains(pr, fipsInfix) ||
		strings.Contains(pr, fipsPrefix) ||
		strings.Contains(pr, fipsSuffix) {
		region = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
			pr, fipsInfix, "-"), fipsPrefix, ""), fipsSuffix, "")
		fips = aws.FIPSEndpointStateEnabled
	} else {
		region = pr
	}

	return region, fips
}

// builtInParameterResolver is the interface responsible for resolving BuiltIn
// values during the sourcing of EndpointParameters
type builtInParameterResolver interface {
	ResolveBuiltIns(*EndpointParameters) error
}

// builtInResolver resolves modeled BuiltIn values using only the members defined
// below.
type builtInResolver struct {
	// The AWS region used to dispatch the request.
	Region string

	// Sourced BuiltIn value in a historical enabled or disabled state.
	UseDualStack aws.DualStackEndpointState

	// Sourced BuiltIn value in a historical enabled or disabled state.
	UseFIPS aws.FIPSEndpointState

	// Base endpoint that can potentially be modified during Endpoint resolution.
	Endpoint *string
}

// Invoked at runtime to resolve BuiltIn Values. Only resolution code specific to
// each BuiltIn value is generated.
func (b *builtInResolver) ResolveBuiltIns(params *EndpointParameters) error {

	region, _ := mapPseudoRegion(b.Region)
	if len(region) == 0 {
		return fmt.Errorf("Could not resolve AWS::Region")
	} else {
		params.Region = aws.String(region)
	}
	if b.UseDualStack == aws.DualStackEndpointStateEnabled {
		params.UseDualStack = aws.Bool(true)
	} else {
		params.UseDualStack = aws.Bool(false)
	}
	if b.UseFIPS == aws.FIPSEndpointStateEnabled {
		params.UseFIPS = aws.Bool(true)
	} else {
		params.UseFIPS = aws.Bool(false)
	}
	params.Endpoint = b.Endpoint
	return nil
}

// EndpointParameters provides the parameters that influence how endpoints are
// resolved.
type EndpointParameters struct {
	// The AWS region used to dispatch the request.
	//
	// Parameter is
	// required.
	//
	// AWS::Region
	Region *string

	// When true, use the dual-stack endpoint. If the configured endpoint does not
	// support dual-stack, dispatching the request MAY return an error.
	//
	// Defaults to
	// false if no value is provided.
	//
	// AWS::UseDualStack
	UseDualStack *bool

	// When true, send this request to the FIPS-compliant regional endpoint. If the
	// configured endpoint does not have a FIPS compliant endpoint, dispatching the
	// request will return an error.
	//
	// Defaults to false if no value is
	// provided.
	//
	// AWS::UseFIPS
	UseFIPS *bool

	// Override the endpoint used to send this request
	//
	// Parameter is
	// required.
	//
	// SDK::Endpoint
	Endpoint *string
}

// ValidateRequired validates required parameters are set.
func (p EndpointParameters) ValidateRequired() error {
	if p.UseDualStack == nil {
		return fmt.Errorf("parameter UseDualStack is required")
	}

	if p.UseFIPS == nil {
		return fmt.Errorf("parameter UseFIPS is required")
	}

	return nil
}

// WithDefaults returns a shallow copy of EndpointParameterswith default values
// applied to members where applicable.
func (p EndpointParameters) WithDefaults() EndpointParameters {
	if p.UseDualStack == nil {
		p.UseDualStack = ptr.Bool(false)
	}

	if p.UseFIPS == nil {
		p.UseFIPS = ptr.Bool(false)
	}
	return p
}

// EndpointResolverV2 provides the interface for resolving service endpoints.
type EndpointResolverV2 interface {
	// ResolveEndpoint attempts to resolve the endpoint with the provided options,
	// returning the endpoint if found. Otherwise an error is returned.
	ResolveEndpoint(ctx context.Context, params EndpointParameters) (
		smithyendpoints.Endpoint, error,
	)
}

// resolver provides the implementation for resolving endpoints.
type resolver struct{}

func NewDefaultEndpointResolverV2() EndpointResolverV2 {
	return &resolver{}
}

// ResolveEndpoint attempts to resolve the endpoint with the provided options,
// returning the endpoint if found. Otherwise an error is returned.
func (r *resolver) ResolveEndpoint(
	ctx context.Context, params EndpointParameters,
) (
	endpoint smithyendpoints.Endpoint, err error,
) {
	params = params.WithDefaults()
	if err = params.ValidateRequired(); err != nil {
		return endpoint, fmt.Errorf("endpoint parameters are not valid, %w", err)
	}
	_UseDualStack := *params.UseDualStack
	_UseFIPS := *params.UseFIPS

	if exprVal := params.Endpoint; exprVal != nil {
		_Endpoint := *exprVal
		_ = _Endpoint
		if _UseFIPS == true {
			return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: FIPS and custom endpoint are not supported")
		}
		if _UseDualStack == true {
			return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: Dualstack and custom endpoint are not supported")
		}
		uriString := _Endpoint

		uri, err := url.Parse(uriString)
		if err != nil {
			return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
		}

		return smithyendpoints.Endpoint{
			URI:     *uri,
			Headers: http.Header{},
		}, nil
	}
	if exprVal := params.Region; exprVal != nil {
		_Region := *exprVal
		_ = _Region
		if exprVal := awsrulesfn.GetPartition(_Region); exprVal != nil {
			_PartitionResult := *exprVal
			_ = _PartitionResult
			if _PartitionResult.Name == "aws" {
				if _UseFIPS == false {
					if _UseDualStack == false {
						uriString := "https://route53.amazonaws.com"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
							Properties: func() smithy.Properties {
								var out smithy.Properties
								out.Set("authSchemes", []interface{}{
									map[string]interface{}{
										"name":          "sigv4",
										"signingName":   "route53",
										"signingRegion": "us-east-1",
									},
								})
								return out
							}(),
						}, nil
					}
				}
			}
			if _PartitionResult.Name == "aws" {
				if _UseFIPS == true {
					if _UseDualStack == false {
						uriString := "https://route53-fips.amazonaws.com"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
							Properties: func() smithy.Properties {
								var out smithy.Properties
								out.Set("authSchemes", []interface{}{
									map[string]interface{}{
										"name":          "sigv4",
										"signingName":   "route53",
										"signingRegion": "us-east-1",
									},
								})
								return out
							}(),
						}, nil
					}
				}
			}
			if _PartitionResult.Name == "aws-cn" {
				if _UseFIPS == false {
					if _UseDualStack == false {
						uriString := "https://route53.amazonaws.com.cn"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
							Properties: func() smithy.Properties {
								var out smithy.Properties
								out.Set("authSchemes", []interface{}{
									map[string]interface{}{
										"name":          "sigv4",
										"signingName":   "route53",
										"signingRegion": "cn-northwest-1",
									},
								})
								return out
							}(),
						}, nil
					}
				}
			}
			if _PartitionResult.Name == "aws-us-gov" {
				if _UseFIPS == false {
					if _UseDualStack == false {
						uriString := "https://route53.us-gov.amazonaws.com"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
							Properties: func() smithy.Properties {
								var out smithy.Properties
								out.Set("authSchemes", []interface{}{
									map[string]interface{}{
										"name":          "sigv4",
										"signingName":   "route53",
										"signingRegion": "us-gov-west-1",
									},
								})
								return out
							}(),
						}, nil
					}
				}
			}
			if _PartitionResult.Name == "aws-us-gov" {
				if _UseFIPS == true {
					if _UseDualStack == false {
						uriString := "https://route53.us-gov.amazonaws.com"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
							Properties: func() smithy.Properties {
								var out smithy.Properties
								out.Set("authSchemes", []interface{}{
									map[string]interface{}{
										"name":          "sigv4",
										"signingName":   "route53",
										"signingRegion": "us-gov-west-1",
									},
								})
								return out
							}(),
						}, nil
					}
				}
			}
			if _PartitionResult.Name == "aws-iso" {
				if _UseFIPS == false {
					if _UseDualStack == false {
						uriString := "https://route53.c2s.ic.gov"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
							Properties: func() smithy.Properties {
								var out smithy.Properties
								out.Set("authSchemes", []interface{}{
									map[string]interface{}{
										"name":          "sigv4",
										"signingName":   "route53",
										"signingRegion": "us-iso-east-1",
									},
								})
								return out
							}(),
						}, nil
					}
				}
			}
			if _PartitionResult.Name == "aws-iso-b" {
				if _UseFIPS == false {
					if _UseDualStack == false {
						uriString := "https://route53.sc2s.sgov.gov"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
							Properties: func() smithy.Properties {
								var out smithy.Properties
								out.Set("authSchemes", []interface{}{
									map[string]interface{}{
										"name":          "sigv4",
										"signingName":   "route53",
										"signingRegion": "us-isob-east-1",
									},
								})
								return out
							}(),
						}, nil
					}
				}
			}
			if _UseFIPS == true {
				if _UseDualStack == true {
					if true == _PartitionResult.SupportsFIPS {
						if true == _PartitionResult.SupportsDualStack {
							uriString := func() string {
								var out strings.Builder
								out.WriteString("https://route53-fips.")
								out.WriteString(_Region)
								out.WriteString(".")
								out.WriteString(_PartitionResult.DualStackDnsSuffix)
								return out.String()
							}()

							uri, err := url.Parse(uriString)
							if err != nil {
								return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
							}

							return smithyendpoints.Endpoint{
								URI:     *uri,
								Headers: http.Header{},
							}, nil
						}
					}
					return endpoint, fmt.Errorf("endpoint rule error, %s", "FIPS and DualStack are enabled, but this partition does not support one or both")
				}
			}
			if _UseFIPS == true {
				if true == _PartitionResult.SupportsFIPS {
					uriString := func() string {
						var out strings.Builder
						out.WriteString("https://route53-fips.")
						out.WriteString(_Region)
						out.WriteString(".")
						out.WriteString(_PartitionResult.DnsSuffix)
						return out.String()
					}()

					uri, err := url.Parse(uriString)
					if err != nil {
						return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
					}

					return smithyendpoints.Endpoint{
						URI:     *uri,
						Headers: http.Header{},
					}, nil
				}
				return endpoint, fmt.Errorf("endpoint rule error, %s", "FIPS is enabled but this partition does not support FIPS")
			}
			if _UseDualStack == true {
				if true == _PartitionResult.SupportsDualStack {
					uriString := func() string {
						var out strings.Builder
						out.WriteString("https://route53.")
						out.WriteString(_Region)
						out.WriteString(".")
						out.WriteString(_PartitionResult.DualStackDnsSuffix)
						return out.String()
					}()

					uri, err := url.Parse(uriString)
					if err != nil {
						return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
					}

					return smithyendpoints.Endpoint{
						URI:     *uri,
						Headers: http.Header{},
					}, nil
				}
				return endpoint, fmt.Errorf("endpoint rule error, %s", "DualStack is enabled but this partition does not support DualStack")
			}
			uriString := func() string {
				var out strings.Builder
				out.WriteString("https://route53.")
				out.WriteString(_Region)
				out.WriteString(".")
				out.WriteString(_PartitionResult.DnsSuffix)
				return out.String()
			}()

			uri, err := url.Parse(uriString)
			if err != nil {
				return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
			}

			return smithyendpoints.Endpoint{
				URI:     *uri,
				Headers: http.Header{},
			}, nil
		}
		return endpoint, fmt.Errorf("Endpoint resolution failed. Invalid operation or environment input.")
	}
	return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: Missing Region")
}

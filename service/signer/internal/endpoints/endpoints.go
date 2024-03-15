// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	endpoints "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2"
	"github.com/aws/smithy-go/logging"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	// Logger is a logging implementation that log events should be sent to.
	Logger logging.Logger

	// LogDeprecated indicates that deprecated endpoints should be logged to the
	// provided logger.
	LogDeprecated bool

	// ResolvedRegion is used to override the region to be resolved, rather then the
	// using the value passed to the ResolveEndpoint method. This value is used by the
	// SDK to translate regions like fips-us-east-1 or us-east-1-fips to an alternative
	// name. You must not set this value directly in your application.
	ResolvedRegion string

	// DisableHTTPS informs the resolver to return an endpoint that does not use the
	// HTTPS scheme.
	DisableHTTPS bool

	// UseDualStackEndpoint specifies the resolver must resolve a dual-stack endpoint.
	UseDualStackEndpoint aws.DualStackEndpointState

	// UseFIPSEndpoint specifies the resolver must resolve a FIPS endpoint.
	UseFIPSEndpoint aws.FIPSEndpointState
}

func (o Options) GetResolvedRegion() string {
	return o.ResolvedRegion
}

func (o Options) GetDisableHTTPS() bool {
	return o.DisableHTTPS
}

func (o Options) GetUseDualStackEndpoint() aws.DualStackEndpointState {
	return o.UseDualStackEndpoint
}

func (o Options) GetUseFIPSEndpoint() aws.FIPSEndpointState {
	return o.UseFIPSEndpoint
}

func transformToSharedOptions(options Options) endpoints.Options {
	return endpoints.Options{
		Logger:               options.Logger,
		LogDeprecated:        options.LogDeprecated,
		ResolvedRegion:       options.ResolvedRegion,
		DisableHTTPS:         options.DisableHTTPS,
		UseDualStackEndpoint: options.UseDualStackEndpoint,
		UseFIPSEndpoint:      options.UseFIPSEndpoint,
	}
}

// Resolver signer endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := transformToSharedOptions(options)
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Aws      *regexp.Regexp
	AwsCn    *regexp.Regexp
	AwsIso   *regexp.Regexp
	AwsIsoB  *regexp.Regexp
	AwsIsoE  *regexp.Regexp
	AwsIsoF  *regexp.Regexp
	AwsUsGov *regexp.Regexp
}{

	Aws:      regexp.MustCompile("^(us|eu|ap|sa|ca|me|af|il)\\-\\w+\\-\\d+$"),
	AwsCn:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	AwsIso:   regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	AwsIsoB:  regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	AwsIsoE:  regexp.MustCompile("^eu\\-isoe\\-\\w+\\-\\d+$"),
	AwsIsoF:  regexp.MustCompile("^us\\-isof\\-\\w+\\-\\d+$"),
	AwsUsGov: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "signer.{region}.api.aws",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "signer-fips.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "signer-fips.{region}.api.aws",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "signer.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.Aws,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "af-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-east-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-northeast-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-northeast-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-southeast-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-southeast-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ca-central-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-central-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-north-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-west-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-west-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-west-3",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "fips-us-east-1",
			}: endpoints.Endpoint{
				Hostname: "signer-fips.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-us-east-2",
			}: endpoints.Endpoint{
				Hostname: "signer-fips.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-us-west-1",
			}: endpoints.Endpoint{
				Hostname: "signer-fips.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-us-west-2",
			}: endpoints.Endpoint{
				Hostname: "signer-fips.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-verification-us-east-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer-fips.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			endpoints.EndpointKey{
				Region: "fips-verification-us-east-2",
			}: endpoints.Endpoint{
				Hostname: "verification.signer-fips.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			endpoints.EndpointKey{
				Region: "fips-verification-us-west-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer-fips.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			endpoints.EndpointKey{
				Region: "fips-verification-us-west-2",
			}: endpoints.Endpoint{
				Hostname: "verification.signer-fips.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
			endpoints.EndpointKey{
				Region: "me-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "sa-east-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "us-east-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "us-east-1",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname: "signer-fips.us-east-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "us-east-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "us-east-2",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname: "signer-fips.us-east-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "us-west-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "us-west-1",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname: "signer-fips.us-west-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "us-west-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "us-west-2",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname: "signer-fips.us-west-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "verification-af-south-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.af-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "af-south-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-ap-east-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.ap-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-east-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-ap-northeast-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.ap-northeast-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-ap-northeast-2",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.ap-northeast-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-2",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-ap-south-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.ap-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-south-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-ap-southeast-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.ap-southeast-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-southeast-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-ap-southeast-2",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.ap-southeast-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-southeast-2",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-ca-central-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.ca-central-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ca-central-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-eu-central-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.eu-central-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-central-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-eu-north-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.eu-north-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-north-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-eu-south-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.eu-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-south-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-eu-west-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.eu-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-eu-west-2",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.eu-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-2",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-eu-west-3",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.eu-west-3.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-3",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-me-south-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.me-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "me-south-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-sa-east-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.sa-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "sa-east-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-us-east-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-us-east-2",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-us-west-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-us-west-2",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
		},
	},
	{
		ID: "aws-cn",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "signer.{region}.api.amazonwebservices.com.cn",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "signer-fips.{region}.amazonaws.com.cn",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "signer-fips.{region}.api.amazonwebservices.com.cn",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "signer.{region}.amazonaws.com.cn",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsCn,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "cn-north-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "cn-northwest-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "verification-cn-north-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.cn-north-1.amazonaws.com.cn",
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-north-1",
				},
			},
			endpoints.EndpointKey{
				Region: "verification-cn-northwest-1",
			}: endpoints.Endpoint{
				Hostname: "verification.signer.cn-northwest-1.amazonaws.com.cn",
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-northwest-1",
				},
			},
		},
	},
	{
		ID: "aws-iso",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "signer-fips.{region}.c2s.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "signer.{region}.c2s.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIso,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-b",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "signer-fips.{region}.sc2s.sgov.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "signer.{region}.sc2s.sgov.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIsoB,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-e",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "signer-fips.{region}.cloud.adc-e.uk",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "signer.{region}.cloud.adc-e.uk",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIsoE,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-f",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "signer-fips.{region}.csp.hci.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "signer.{region}.csp.hci.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIsoF,
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "signer.{region}.api.aws",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "signer-fips.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "signer-fips.{region}.api.aws",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "signer.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsUsGov,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "us-gov-east-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "us-gov-west-1",
			}: endpoints.Endpoint{},
		},
	},
}

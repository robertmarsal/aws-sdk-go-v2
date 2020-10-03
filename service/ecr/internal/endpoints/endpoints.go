// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	DisableHTTPS bool
}

// Resolver ECR endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	opt := endpoints.Options{
		DisableHTTPS: options.DisableHTTPS,
	}
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.ecr.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"af-south-1": endpoints.Endpoint{
				Hostname: "api.ecr.af-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "af-south-1",
				},
			},
			"ap-east-1": endpoints.Endpoint{
				Hostname: "api.ecr.ap-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-east-1",
				},
			},
			"ap-northeast-1": endpoints.Endpoint{
				Hostname: "api.ecr.ap-northeast-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-1",
				},
			},
			"ap-northeast-2": endpoints.Endpoint{
				Hostname: "api.ecr.ap-northeast-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-2",
				},
			},
			"ap-south-1": endpoints.Endpoint{
				Hostname: "api.ecr.ap-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-south-1",
				},
			},
			"ap-southeast-1": endpoints.Endpoint{
				Hostname: "api.ecr.ap-southeast-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-southeast-1",
				},
			},
			"ap-southeast-2": endpoints.Endpoint{
				Hostname: "api.ecr.ap-southeast-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-southeast-2",
				},
			},
			"ca-central-1": endpoints.Endpoint{
				Hostname: "api.ecr.ca-central-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ca-central-1",
				},
			},
			"eu-central-1": endpoints.Endpoint{
				Hostname: "api.ecr.eu-central-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-central-1",
				},
			},
			"eu-north-1": endpoints.Endpoint{
				Hostname: "api.ecr.eu-north-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-north-1",
				},
			},
			"eu-south-1": endpoints.Endpoint{
				Hostname: "api.ecr.eu-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-south-1",
				},
			},
			"eu-west-1": endpoints.Endpoint{
				Hostname: "api.ecr.eu-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-1",
				},
			},
			"eu-west-2": endpoints.Endpoint{
				Hostname: "api.ecr.eu-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-2",
				},
			},
			"eu-west-3": endpoints.Endpoint{
				Hostname: "api.ecr.eu-west-3.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-3",
				},
			},
			"fips-dkr-us-east-1": endpoints.Endpoint{
				Hostname: "ecr-fips.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"fips-dkr-us-east-2": endpoints.Endpoint{
				Hostname: "ecr-fips.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			"fips-dkr-us-west-1": endpoints.Endpoint{
				Hostname: "ecr-fips.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			"fips-dkr-us-west-2": endpoints.Endpoint{
				Hostname: "ecr-fips.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
			"fips-us-east-1": endpoints.Endpoint{
				Hostname: "ecr-fips.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"fips-us-east-2": endpoints.Endpoint{
				Hostname: "ecr-fips.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			"fips-us-west-1": endpoints.Endpoint{
				Hostname: "ecr-fips.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			"fips-us-west-2": endpoints.Endpoint{
				Hostname: "ecr-fips.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
			"me-south-1": endpoints.Endpoint{
				Hostname: "api.ecr.me-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "me-south-1",
				},
			},
			"sa-east-1": endpoints.Endpoint{
				Hostname: "api.ecr.sa-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "sa-east-1",
				},
			},
			"us-east-1": endpoints.Endpoint{
				Hostname: "api.ecr.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"us-east-2": endpoints.Endpoint{
				Hostname: "api.ecr.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			"us-west-1": endpoints.Endpoint{
				Hostname: "api.ecr.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			"us-west-2": endpoints.Endpoint{
				Hostname: "api.ecr.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
		},
	},
	{
		ID: "aws-cn",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.ecr.{region}.amazonaws.com.cn",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"cn-north-1": endpoints.Endpoint{
				Hostname: "api.ecr.cn-north-1.amazonaws.com.cn",
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-north-1",
				},
			},
			"cn-northwest-1": endpoints.Endpoint{
				Hostname: "api.ecr.cn-northwest-1.amazonaws.com.cn",
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-northwest-1",
				},
			},
		},
	},
	{
		ID: "aws-iso",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.ecr.{region}.c2s.ic.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"us-iso-east-1": endpoints.Endpoint{
				Hostname: "api.ecr.us-iso-east-1.c2s.ic.gov",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-iso-east-1",
				},
			},
		},
	},
	{
		ID: "aws-iso-b",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.ecr.{region}.sc2s.sgov.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.ecr.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"fips-dkr-us-gov-east-1": endpoints.Endpoint{
				Hostname: "ecr-fips.us-gov-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"fips-dkr-us-gov-west-1": endpoints.Endpoint{
				Hostname: "ecr-fips.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
			"fips-us-gov-east-1": endpoints.Endpoint{
				Hostname: "ecr-fips.us-gov-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"fips-us-gov-west-1": endpoints.Endpoint{
				Hostname: "ecr-fips.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
			"us-gov-east-1": endpoints.Endpoint{
				Hostname: "api.ecr.us-gov-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"us-gov-west-1": endpoints.Endpoint{
				Hostname: "api.ecr.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
		},
	},
}

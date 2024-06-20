// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains details about a package version asset.
type AssetSummary struct {

	//  The name of the asset.
	//
	// This member is required.
	Name *string

	//  The hashes of the asset.
	Hashes map[string]string

	//  The size of the asset.
	Size *int64

	noSmithyDocumentSerde
}

// A package associated with a package group.
type AssociatedPackage struct {

	// Describes the strength of the association between the package and package
	// group. A strong match can be thought of as an exact match, and a weak match can
	// be thought of as a variation match, for example, the package name matches a
	// variation of the package group pattern. For more information about package group
	// pattern matching, including strong and weak matches, see [Package group definition syntax and matching behavior]in the CodeArtifact
	// User Guide.
	//
	// [Package group definition syntax and matching behavior]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-definition-syntax-matching-behavior.html
	AssociationType PackageGroupAssociationType

	// A format that specifies the type of the associated package.
	Format PackageFormat

	// The namespace of the associated package. The package component that specifies
	// its namespace depends on its type. For example:
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	//  The name of the associated package.
	Package *string

	noSmithyDocumentSerde
}

//	Information about a domain. A domain is a container for repositories. When you
//
// create a domain, it is empty until you add one or more repositories.
type DomainDescription struct {

	//  The Amazon Resource Name (ARN) of the domain.
	Arn *string

	//  The total size of all assets in the domain.
	AssetSizeBytes int64

	//  A timestamp that represents the date and time the domain was created.
	CreatedTime *time.Time

	//  The ARN of an Key Management Service (KMS) key associated with a domain.
	EncryptionKey *string

	//  The name of the domain.
	Name *string

	//  The Amazon Web Services account ID that owns the domain.
	Owner *string

	//  The number of repositories in the domain.
	RepositoryCount int32

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket that is used to store
	// package assets in the domain.
	S3BucketArn *string

	//  The current status of a domain.
	Status DomainStatus

	noSmithyDocumentSerde
}

// Information about how a package originally entered the CodeArtifact domain. For
// packages published directly to CodeArtifact, the entry point is the repository
// it was published to. For packages ingested from an external repository, the
// entry point is the external connection that it was ingested from. An external
// connection is a CodeArtifact repository that is connected to an external
// repository such as the npm registry or NuGet gallery.
//
// If a package version exists in a repository and is updated, for example if a
// package of the same version is added with additional assets, the package
// version's DomainEntryPoint will not change from the original package version's
// value.
type DomainEntryPoint struct {

	// The name of the external connection that a package was ingested from.
	ExternalConnectionName *string

	// The name of the repository that a package was originally published to.
	RepositoryName *string

	noSmithyDocumentSerde
}

//	Information about a domain, including its name, Amazon Resource Name (ARN),
//
// and status. The [ListDomains]operation returns a list of DomainSummary objects.
//
// [ListDomains]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListDomains.html
type DomainSummary struct {

	//  The ARN of the domain.
	Arn *string

	//  A timestamp that contains the date and time the domain was created.
	CreatedTime *time.Time

	//  The key used to encrypt the domain.
	EncryptionKey *string

	//  The name of the domain.
	Name *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	Owner *string

	//  A string that contains the status of the domain.
	Status DomainStatus

	noSmithyDocumentSerde
}

// Details of the license data.
type LicenseInfo struct {

	//  Name of the license.
	Name *string

	//  The URL for license data.
	Url *string

	noSmithyDocumentSerde
}

// Details about a package dependency.
type PackageDependency struct {

	//  The type of a package dependency. The possible values depend on the package
	// type.
	//
	//   - npm: regular , dev , peer , optional
	//
	//   - maven: optional , parent , compile , runtime , test , system , provided .
	//
	// Note that parent is not a regular Maven dependency type; instead this is
	//   extracted from the element if one is defined in the package version's POM
	//   file.
	//
	//   - nuget: The dependencyType field is never set for NuGet packages.
	//
	//   - pypi: Requires-Dist
	DependencyType *string

	// The namespace of the package that this package depends on. The package
	// component that specifies its namespace depends on its type. For example:
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	//  The name of the package that this package depends on.
	Package *string

	//  The required version, or version range, of the package that this package
	// depends on. The version format is specific to the package type. For example, the
	// following are possible valid required versions: 1.2.3 , ^2.3.4 , or 4.x .
	VersionRequirement *string

	noSmithyDocumentSerde
}

// Details about a package.
type PackageDescription struct {

	// A format that specifies the type of the package.
	Format PackageFormat

	// The name of the package.
	Name *string

	// The namespace of the package. The package component that specifies its
	// namespace depends on its type. For example:
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	// The package origin configuration for the package.
	OriginConfiguration *PackageOriginConfiguration

	noSmithyDocumentSerde
}

//	Details about an allowed repository for a package group, including its name
//
// and origin configuration.
type PackageGroupAllowedRepository struct {

	// The origin configuration restriction type of the allowed repository.
	OriginRestrictionType PackageGroupOriginRestrictionType

	//  The name of the allowed repository.
	RepositoryName *string

	noSmithyDocumentSerde
}

// The description of the package group.
type PackageGroupDescription struct {

	//  The ARN of the package group.
	Arn *string

	//  The contact information of the package group.
	ContactInfo *string

	// A timestamp that represents the date and time the package group was created.
	CreatedTime *time.Time

	//  The description of the package group.
	Description *string

	//  The name of the domain that contains the package group.
	DomainName *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The package group origin configuration that determines how package versions can
	// enter repositories.
	OriginConfiguration *PackageGroupOriginConfiguration

	//  The direct parent package group of the package group.
	Parent *PackageGroupReference

	//  The pattern of the package group. The pattern determines which packages are
	// associated with the package group.
	Pattern *string

	noSmithyDocumentSerde
}

// The package group origin configuration that determines how package versions can
// enter repositories.
type PackageGroupOriginConfiguration struct {

	// The origin configuration settings that determine how package versions can enter
	// repositories.
	Restrictions map[string]PackageGroupOriginRestriction

	noSmithyDocumentSerde
}

// Contains information about the configured restrictions of the origin controls
// of a package group.
type PackageGroupOriginRestriction struct {

	// The effective package group origin restriction setting. If the value of mode is
	// ALLOW , ALLOW_SPECIFIC_REPOSITORIES , or BLOCK , then the value of effectiveMode
	// is the same. Otherwise, when the value of mode is INHERIT , then the value of
	// effectiveMode is the value of mode of the first parent group which does not
	// have a value of INHERIT .
	EffectiveMode PackageGroupOriginRestrictionMode

	// The parent package group that the package group origin restrictions are
	// inherited from.
	InheritedFrom *PackageGroupReference

	// The package group origin restriction setting. If the value of mode is ALLOW ,
	// ALLOW_SPECIFIC_REPOSITORIES , or BLOCK , then the value of effectiveMode is the
	// same. Otherwise, when the value is INHERIT , then the value of effectiveMode is
	// the value of mode of the first parent group which does not have a value of
	// INHERIT .
	Mode PackageGroupOriginRestrictionMode

	// The number of repositories in the allowed repository list.
	RepositoriesCount *int64

	noSmithyDocumentSerde
}

// Information about the identifiers of a package group.
type PackageGroupReference struct {

	//  The ARN of the package group.
	Arn *string

	//  The pattern of the package group. The pattern determines which packages are
	// associated with the package group, and is also the identifier of the package
	// group.
	Pattern *string

	noSmithyDocumentSerde
}

// Details about a package group.
type PackageGroupSummary struct {

	//  The ARN of the package group.
	Arn *string

	//  The contact information of the package group.
	ContactInfo *string

	// A timestamp that represents the date and time the repository was created.
	CreatedTime *time.Time

	//  The description of the package group.
	Description *string

	//  The domain that contains the package group.
	DomainName *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// Details about the package origin configuration of a package group.
	OriginConfiguration *PackageGroupOriginConfiguration

	//  The direct parent package group of the package group.
	Parent *PackageGroupReference

	//  The pattern of the package group. The pattern determines which packages are
	// associated with the package group.
	Pattern *string

	noSmithyDocumentSerde
}

// Details about the package origin configuration of a package.
type PackageOriginConfiguration struct {

	// A PackageOriginRestrictions object that contains information about the upstream
	// and publish package origin configuration for the package.
	Restrictions *PackageOriginRestrictions

	noSmithyDocumentSerde
}

// Details about the origin restrictions set on the package. The package origin
// restrictions determine how new versions of a package can be added to a specific
// repository.
type PackageOriginRestrictions struct {

	// The package origin configuration that determines if new versions of the package
	// can be published directly to the repository.
	//
	// This member is required.
	Publish AllowPublish

	// The package origin configuration that determines if new versions of the package
	// can be added to the repository from an external connection or upstream source.
	//
	// This member is required.
	Upstream AllowUpstream

	noSmithyDocumentSerde
}

// Details about a package, including its format, namespace, and name.
type PackageSummary struct {

	//  The format of the package.
	Format PackageFormat

	// The namespace of the package. The package component that specifies its
	// namespace depends on its type. For example:
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	// A [PackageOriginConfiguration] object that contains a [PackageOriginRestrictions] object that contains information about the upstream
	// and publish package origin restrictions.
	//
	// [PackageOriginRestrictions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageOriginRestrictions.html
	// [PackageOriginConfiguration]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageOriginConfiguration.html
	OriginConfiguration *PackageOriginConfiguration

	//  The name of the package.
	Package *string

	noSmithyDocumentSerde
}

// Details about a package version.
type PackageVersionDescription struct {

	//  The name of the package that is displayed. The displayName varies depending on
	// the package version's format. For example, if an npm package is named ui , is in
	// the namespace vue , and has the format npm , then the displayName is @vue/ui .
	DisplayName *string

	//  The format of the package version.
	Format PackageFormat

	//  The homepage associated with the package.
	HomePage *string

	//  Information about licenses associated with the package version.
	Licenses []LicenseInfo

	// The namespace of the package version. The package component that specifies its
	// namespace depends on its type. For example:
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	// A [PackageVersionOrigin] object that contains information about how the package version was added to
	// the repository.
	//
	// [PackageVersionOrigin]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionOrigin.html
	Origin *PackageVersionOrigin

	//  The name of the requested package.
	PackageName *string

	//  A timestamp that contains the date and time the package version was published.
	PublishedTime *time.Time

	//  The revision of the package version.
	Revision *string

	//  The repository for the source code in the package version, or the source code
	// used to build it.
	SourceCodeRepository *string

	//  A string that contains the status of the package version.
	Status PackageVersionStatus

	//  A summary of the package version. The summary is extracted from the package.
	// The information in and detail level of the summary depends on the package
	// version's format.
	Summary *string

	//  The version of the package.
	Version *string

	noSmithyDocumentSerde
}

// l An error associated with package.
type PackageVersionError struct {

	//  The error code associated with the error. Valid error codes are:
	//
	//   - ALREADY_EXISTS
	//
	//   - MISMATCHED_REVISION
	//
	//   - MISMATCHED_STATUS
	//
	//   - NOT_ALLOWED
	//
	//   - NOT_FOUND
	//
	//   - SKIPPED
	ErrorCode PackageVersionErrorCode

	//  The error message associated with the error.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// Information about how a package version was added to a repository.
type PackageVersionOrigin struct {

	// A [DomainEntryPoint] object that contains information about from which repository or external
	// connection the package version was added to the domain.
	//
	// [DomainEntryPoint]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainEntryPoint.html
	DomainEntryPoint *DomainEntryPoint

	// Describes how the package version was originally added to the domain. An
	// INTERNAL origin type means the package version was published directly to a
	// repository in the domain. An EXTERNAL origin type means the package version was
	// ingested from an external connection.
	OriginType PackageVersionOriginType

	noSmithyDocumentSerde
}

//	Details about a package version, including its status, version, and revision.
//
// The [ListPackageVersions]operation returns a list of PackageVersionSummary objects.
//
// [ListPackageVersions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html
type PackageVersionSummary struct {

	//  A string that contains the status of the package version. It can be one of the
	// following:
	//
	// This member is required.
	Status PackageVersionStatus

	//  Information about a package version.
	//
	// This member is required.
	Version *string

	// A [PackageVersionOrigin] object that contains information about how the package version was added to
	// the repository.
	//
	// [PackageVersionOrigin]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionOrigin.html
	Origin *PackageVersionOrigin

	//  The revision associated with a package version.
	Revision *string

	noSmithyDocumentSerde
}

//	The details of a repository stored in CodeArtifact. A CodeArtifact repository
//
// contains a set of package versions, each of which maps to a set of assets.
// Repositories are polyglot—a single repository can contain packages of any
// supported type. Each repository exposes endpoints for fetching and publishing
// packages using tools like the npm CLI, the Maven CLI ( mvn ), and pip . You can
// create up to 100 repositories per Amazon Web Services account.
type RepositoryDescription struct {

	//  The 12-digit account number of the Amazon Web Services account that manages
	// the repository.
	AdministratorAccount *string

	//  The Amazon Resource Name (ARN) of the repository.
	Arn *string

	// A timestamp that represents the date and time the repository was created.
	CreatedTime *time.Time

	//  A text description of the repository.
	Description *string

	//  The name of the domain that contains the repository.
	DomainName *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain that contains the repository. It does not include dashes or spaces.
	DomainOwner *string

	//  An array of external connections associated with the repository.
	ExternalConnections []RepositoryExternalConnectionInfo

	//  The name of the repository.
	Name *string

	//  A list of upstream repositories to associate with the repository. The order of
	// the upstream repositories in the list determines their priority order when
	// CodeArtifact looks for a requested package version. For more information, see [Working with upstream repositories].
	//
	// [Working with upstream repositories]: https://docs.aws.amazon.com/codeartifact/latest/ug/repos-upstream.html
	Upstreams []UpstreamRepositoryInfo

	noSmithyDocumentSerde
}

// Contains information about the external connection of a repository.
type RepositoryExternalConnectionInfo struct {

	//  The name of the external connection associated with a repository.
	ExternalConnectionName *string

	//  The package format associated with a repository's external connection. The
	// valid package formats are:
	//
	//   - npm : A Node Package Manager (npm) package.
	//
	//   - pypi : A Python Package Index (PyPI) package.
	//
	//   - maven : A Maven package that contains compiled code in a distributable
	//   format, such as a JAR file.
	//
	//   - nuget : A NuGet package.
	PackageFormat PackageFormat

	//  The status of the external connection of a repository. There is one valid
	// value, Available .
	Status ExternalConnectionStatus

	noSmithyDocumentSerde
}

//	Details about a repository, including its Amazon Resource Name (ARN),
//
// description, and domain information. The [ListRepositories]operation returns a list of
// RepositorySummary objects.
//
// [ListRepositories]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListRepositories.html
type RepositorySummary struct {

	//  The Amazon Web Services account ID that manages the repository.
	AdministratorAccount *string

	//  The ARN of the repository.
	Arn *string

	// A timestamp that represents the date and time the repository was created.
	CreatedTime *time.Time

	//  The description of the repository.
	Description *string

	//  The name of the domain that contains the repository.
	DomainName *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	//  The name of the repository.
	Name *string

	noSmithyDocumentSerde
}

//	An CodeArtifact resource policy that contains a resource ARN, document
//
// details, and a revision.
type ResourcePolicy struct {

	//  The resource policy formatted in JSON.
	Document *string

	//  The ARN of the resource associated with the resource policy
	ResourceArn *string

	//  The current revision of the resource policy.
	Revision *string

	noSmithyDocumentSerde
}

// Contains the revision and status of a package version.
type SuccessfulPackageVersionInfo struct {

	//  The revision of a package version.
	Revision *string

	//  The status of a package version.
	Status PackageVersionStatus

	noSmithyDocumentSerde
}

// A tag is a key-value pair that can be used to manage, search for, or filter
// resources in CodeArtifact.
type Tag struct {

	// The tag key.
	//
	// This member is required.
	Key *string

	// The tag value.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

//	Information about an upstream repository. A list of UpstreamRepository objects
//
// is an input parameter to [CreateRepository]and [UpdateRepository].
//
// [UpdateRepository]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdateRepository.html
// [CreateRepository]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_CreateRepository.html
type UpstreamRepository struct {

	//  The name of an upstream repository.
	//
	// This member is required.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Information about an upstream repository.
type UpstreamRepositoryInfo struct {

	//  The name of an upstream repository.
	RepositoryName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

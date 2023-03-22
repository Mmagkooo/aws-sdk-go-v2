// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Defines a recommendation for a CloudWatch alarm.
type AlarmRecommendation struct {

	// The name of the alarm recommendation.
	//
	// This member is required.
	Name *string

	// The identifier of the alarm recommendation.
	//
	// This member is required.
	RecommendationId *string

	// The reference identifier of the alarm recommendation.
	//
	// This member is required.
	ReferenceId *string

	// The type of alarm recommendation.
	//
	// This member is required.
	Type AlarmType

	// The Application Component for the CloudWatch alarm recommendation.
	AppComponentName *string

	// The description of the recommendation.
	Description *string

	// The list of CloudWatch alarm recommendations.
	Items []RecommendationItem

	// The prerequisite for the alarm recommendation.
	Prerequisite *string

	noSmithyDocumentSerde
}

// Defines an Resilience Hub application.
type App struct {

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The timestamp for when the app was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The name for the application.
	//
	// This member is required.
	Name *string

	// Assessment execution schedule with 'Daily' or 'Disabled' values.
	AssessmentSchedule AppAssessmentScheduleType

	// The current status of compliance for the resiliency policy.
	ComplianceStatus AppComplianceStatusType

	// The optional description for an app.
	Description *string

	// The timestamp for the most recent compliance evaluation.
	LastAppComplianceEvaluationTime *time.Time

	// The timestamp for the most recent resiliency score evaluation.
	LastResiliencyScoreEvaluationTime *time.Time

	// The Amazon Resource Name (ARN) of the resiliency policy. The format for this ARN
	// is: arn:partition:resiliencehub:region:account:resiliency-policy/policy-id. For
	// more information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	PolicyArn *string

	// The current resiliency score for the application.
	ResiliencyScore float64

	// The status of the application.
	Status AppStatusType

	// The tags assigned to the resource. A tag is a label that you assign to an Amazon
	// Web Services resource. Each tag consists of a key/value pair.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Defines an application assessment.
type AppAssessment struct {

	// The Amazon Resource Name (ARN) of the assessment. The format for this ARN is:
	// arn:partition:resiliencehub:region:account:app-assessment/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AssessmentArn *string

	// The current status of the assessment for the resiliency policy.
	//
	// This member is required.
	AssessmentStatus AssessmentStatus

	// The entity that invoked the assessment.
	//
	// This member is required.
	Invoker AssessmentInvoker

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	AppArn *string

	// The version of the application.
	AppVersion *string

	// The name of the assessment.
	AssessmentName *string

	// The application compliance against the resiliency policy.
	Compliance map[string]DisruptionCompliance

	// The current status of the compliance for the resiliency policy.
	ComplianceStatus ComplianceStatus

	// The cost for the application.
	Cost *Cost

	// The end time for the action.
	EndTime *time.Time

	// Error or warning message from the assessment execution
	Message *string

	// The resiliency policy.
	Policy *ResiliencyPolicy

	// The current resiliency score for the application.
	ResiliencyScore *ResiliencyScore

	// A resource error object containing a list of errors retrieving an application's
	// resources.
	ResourceErrorsDetails *ResourceErrorsDetails

	// The starting time for the action.
	StartTime *time.Time

	// The tags assigned to the resource. A tag is a label that you assign to an Amazon
	// Web Services resource. Each tag consists of a key/value pair.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Defines an application assessment summary.
type AppAssessmentSummary struct {

	// The Amazon Resource Name (ARN) of the assessment. The format for this ARN is:
	// arn:partition:resiliencehub:region:account:app-assessment/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AssessmentArn *string

	// The current status of the assessment for the resiliency policy.
	//
	// This member is required.
	AssessmentStatus AssessmentStatus

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	AppArn *string

	// The version of the application.
	AppVersion *string

	// The name of the assessment.
	AssessmentName *string

	// The current status of compliance for the resiliency policy.
	ComplianceStatus ComplianceStatus

	// The cost for the application.
	Cost *Cost

	// The end time for the action.
	EndTime *time.Time

	// The entity that invoked the assessment.
	Invoker AssessmentInvoker

	// The message from the assessment run.
	Message *string

	// The current resiliency score for the application.
	ResiliencyScore float64

	// The starting time for the action.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Defines an Application Component.
type AppComponent struct {

	// The name of the Application Component.
	//
	// This member is required.
	Name *string

	// The type of Application Component.
	//
	// This member is required.
	Type *string

	// Additional configuration parameters for an AWS Resilience Hub application.
	// Currently, this parameter accepts a key-value mapping (in a string format) of
	// only one failover region and one associated account. Key: "failover-regions"
	// Value: "[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"
	AdditionalInfo map[string][]string

	// Unique identifier of the Application Component.
	Id *string

	noSmithyDocumentSerde
}

// Defines the compliance of an Application Component against the resiliency
// policy.
type AppComponentCompliance struct {

	// The name of the Application Component.
	AppComponentName *string

	// The compliance of the Application Component against the resiliency policy.
	Compliance map[string]DisruptionCompliance

	// The cost for the application.
	Cost *Cost

	// The compliance message.
	Message *string

	// The current resiliency score for the application.
	ResiliencyScore *ResiliencyScore

	// The status of the action.
	Status ComplianceStatus

	noSmithyDocumentSerde
}

// The list of Resilience Hub application input sources.
type AppInputSource struct {

	// The resource type of the input source.
	//
	// This member is required.
	ImportType ResourceMappingType

	// The namespace on your Amazon Elastic Kubernetes Service cluster.
	EksSourceClusterNamespace *EksSourceClusterNamespace

	// The number of resources.
	ResourceCount int32

	// The Amazon Resource Name (ARN) of the input source. For more information about
	// ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	SourceArn *string

	// The name of the input source.
	SourceName *string

	// The name of the Terraform s3 state ﬁle.
	TerraformSource *TerraformSource

	noSmithyDocumentSerde
}

// Defines an application summary.
type AppSummary struct {

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The timestamp for when the app was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The name of the application.
	//
	// This member is required.
	Name *string

	// Assessment execution schedule with 'Daily' or 'Disabled' values.
	AssessmentSchedule AppAssessmentScheduleType

	// The current status of compliance for the resiliency policy.
	ComplianceStatus AppComplianceStatusType

	// The optional description for an app.
	Description *string

	// The current resiliency score for the application.
	ResiliencyScore float64

	// The status of the application.
	Status AppStatusType

	noSmithyDocumentSerde
}

// The version of the application.
type AppVersionSummary struct {

	// The version of the application.
	//
	// This member is required.
	AppVersion *string

	noSmithyDocumentSerde
}

// Defines recommendations for an Resilience Hub Application Component, returned as
// an object. This object contains component names, configuration recommendations,
// and recommendation statuses.
type ComponentRecommendation struct {

	// The name of the Application Component.
	//
	// This member is required.
	AppComponentName *string

	// The list of recommendations.
	//
	// This member is required.
	ConfigRecommendations []ConfigRecommendation

	// The recommendation status.
	//
	// This member is required.
	RecommendationStatus RecommendationComplianceStatus

	noSmithyDocumentSerde
}

// Defines a configuration recommendation.
type ConfigRecommendation struct {

	// The name of the recommendation configuration.
	//
	// This member is required.
	Name *string

	// The type of optimization.
	//
	// This member is required.
	OptimizationType ConfigRecommendationOptimizationType

	// The reference identifier for the recommendation configuration.
	//
	// This member is required.
	ReferenceId *string

	// The name of the Application Component.
	AppComponentName *string

	// The current compliance against the resiliency policy before applying the
	// configuration change.
	Compliance map[string]DisruptionCompliance

	// The cost for the application.
	Cost *Cost

	// The optional description for an app.
	Description *string

	// The architecture type.
	HaArchitecture HaArchitecture

	// The expected compliance against the resiliency policy after applying the
	// configuration change.
	RecommendationCompliance map[string]RecommendationDisruptionCompliance

	// List of the suggested configuration changes.
	SuggestedChanges []string

	noSmithyDocumentSerde
}

// Defines a cost object.
type Cost struct {

	// The cost amount.
	//
	// This member is required.
	Amount float64

	// The cost currency, for example USD.
	//
	// This member is required.
	Currency *string

	// The cost frequency.
	//
	// This member is required.
	Frequency CostFrequency

	noSmithyDocumentSerde
}

// Defines the compliance against the resiliency policy for a disruption.
type DisruptionCompliance struct {

	// The current status of compliance for the resiliency policy.
	//
	// This member is required.
	ComplianceStatus ComplianceStatus

	// The Recovery Point Objective (RPO) that is achievable, in seconds.
	AchievableRpoInSecs int32

	// The Recovery Time Objective (RTO) that is achievable, in seconds
	AchievableRtoInSecs int32

	// The current RPO, in seconds.
	CurrentRpoInSecs int32

	// The current RTO, in seconds.
	CurrentRtoInSecs int32

	// The disruption compliance message.
	Message *string

	// The RPO description.
	RpoDescription *string

	// The RPO reference identifier.
	RpoReferenceId *string

	// The RTO description.
	RtoDescription *string

	// The RTO reference identifier.
	RtoReferenceId *string

	noSmithyDocumentSerde
}

// The input source of the Amazon Elastic Kubernetes Service cluster.
type EksSource struct {

	// The Amazon Resource Name (ARN) of the Amazon Elastic Kubernetes Service cluster.
	// The format for this ARN is: arn:aws:eks:region:account-id:cluster/cluster-name.
	// For more information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	EksClusterArn *string

	// The list of namespaces located on your Amazon Elastic Kubernetes Service
	// cluster.
	//
	// This member is required.
	Namespaces []string

	noSmithyDocumentSerde
}

// The input source of the namespace that is located on your Amazon Elastic
// Kubernetes Service cluster.
type EksSourceClusterNamespace struct {

	// The Amazon Resource Name (ARN) of the Amazon Elastic Kubernetes Service cluster.
	// The format for this ARN is: arn:aws:eks:region:account-id:cluster/cluster-name.
	// For more information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	EksClusterArn *string

	// Name of the namespace that is located on your Amazon Elastic Kubernetes Service
	// cluster.
	//
	// This member is required.
	Namespace *string

	noSmithyDocumentSerde
}

// Defines a failure policy.
type FailurePolicy struct {

	// The Recovery Point Objective (RPO), in seconds.
	//
	// This member is required.
	RpoInSecs int32

	// The Recovery Time Objective (RTO), in seconds.
	//
	// This member is required.
	RtoInSecs int32

	noSmithyDocumentSerde
}

// Defines a logical resource identifier.
type LogicalResourceId struct {

	// The identifier of the resource.
	//
	// This member is required.
	Identifier *string

	// The name of the Amazon Elastic Kubernetes Service cluster and namespace this
	// resource belongs to. This parameter accepts values in "eks-cluster/namespace"
	// format.
	EksSourceName *string

	// The name of the CloudFormation stack this resource belongs to.
	LogicalStackName *string

	// The name of the resource group that this resource belongs to.
	ResourceGroupName *string

	// The name of the Terraform S3 state file this resource belongs to.
	TerraformSourceName *string

	noSmithyDocumentSerde
}

// Defines a physical resource. A physical resource is a resource that exists in
// your account. It can be identified using an Amazon Resource Name (ARN) or an
// Resilience Hub-native identifier.
type PhysicalResource struct {

	// The logical identifier of the resource.
	//
	// This member is required.
	LogicalResourceId *LogicalResourceId

	// The physical identifier of the resource.
	//
	// This member is required.
	PhysicalResourceId *PhysicalResourceId

	// The type of resource.
	//
	// This member is required.
	ResourceType *string

	// Additional configuration parameters for an AWS Resilience Hub application.
	// Currently, this parameter accepts a key-value mapping (in a string format) of
	// only one failover region and one associated account. Key: "failover-regions"
	// Value: "[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"
	AdditionalInfo map[string][]string

	// The application components that belong to this resource.
	AppComponents []AppComponent

	// Indicates if a resource is included or excluded from the assessment.
	Excluded *bool

	// The name of the resource.
	ResourceName *string

	noSmithyDocumentSerde
}

// Defines a physical resource identifier.
type PhysicalResourceId struct {

	// The identifier of the physical resource.
	//
	// This member is required.
	Identifier *string

	// Specifies the type of physical resource identifier. Arn The resource identifier
	// is an Amazon Resource Name (ARN) . Native The resource identifier is an
	// Resilience Hub-native identifier.
	//
	// This member is required.
	Type PhysicalIdentifierType

	// The Amazon Web Services account that owns the physical resource.
	AwsAccountId *string

	// The Amazon Web Services Region that the physical resource is located in.
	AwsRegion *string

	noSmithyDocumentSerde
}

// Defines a disruption compliance recommendation.
type RecommendationDisruptionCompliance struct {

	// The expected compliance status after applying the recommended configuration
	// change.
	//
	// This member is required.
	ExpectedComplianceStatus ComplianceStatus

	// The expected Recovery Point Objective (RPO) description after applying the
	// recommended configuration change.
	ExpectedRpoDescription *string

	// The expected RPO after applying the recommended configuration change.
	ExpectedRpoInSecs int32

	// The expected Recovery Time Objective (RTO) description after applying the
	// recommended configuration change.
	ExpectedRtoDescription *string

	// The expected RTO after applying the recommended configuration change.
	ExpectedRtoInSecs int32

	noSmithyDocumentSerde
}

// Defines a recommendation.
type RecommendationItem struct {

	// Specifies if the recommendation has already been implemented.
	AlreadyImplemented *bool

	// The resource identifier.
	ResourceId *string

	// The target account identifier.
	TargetAccountId *string

	// The target region.
	TargetRegion *string

	noSmithyDocumentSerde
}

// Defines a recommendation template created with the CreateRecommendationTemplate
// action.
type RecommendationTemplate struct {

	// The Amazon Resource Name (ARN) of the assessment. The format for this ARN is:
	// arn:partition:resiliencehub:region:account:app-assessment/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AssessmentArn *string

	// The format of the recommendation template. CfnJson The template is
	// CloudFormation JSON. CfnYaml The template is CloudFormation YAML.
	//
	// This member is required.
	Format TemplateFormat

	// The name for the recommendation template.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) for the recommendation template.
	//
	// This member is required.
	RecommendationTemplateArn *string

	// An array of strings that specify the recommendation template type or types.
	// Alarm The template is an AlarmRecommendation template. Sop The template is a
	// SopRecommendation template. Test The template is a TestRecommendation template.
	//
	// This member is required.
	RecommendationTypes []RenderRecommendationType

	// The status of the action.
	//
	// This member is required.
	Status RecommendationTemplateStatus

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	AppArn *string

	// The end time for the action.
	EndTime *time.Time

	// The message for the recommendation template.
	Message *string

	// Indicates if replacements are needed.
	NeedsReplacements *bool

	// Identifiers for the recommendations used in the recommendation template.
	RecommendationIds []string

	// The start time for the action.
	StartTime *time.Time

	// The tags assigned to the resource. A tag is a label that you assign to an Amazon
	// Web Services resource. Each tag consists of a key/value pair.
	Tags map[string]string

	// The file location of the template.
	TemplatesLocation *S3Location

	noSmithyDocumentSerde
}

// Defines a resiliency policy.
type ResiliencyPolicy struct {

	// The timestamp for when the resiliency policy was created.
	CreationTime *time.Time

	// Specifies a high-level geographical location constraint for where your
	// resilience policy data can be stored.
	DataLocationConstraint DataLocationConstraint

	// Specifies the estimated cost tier of the resiliency policy.
	EstimatedCostTier EstimatedCostTier

	// The resiliency policy.
	Policy map[string]FailurePolicy

	// The Amazon Resource Name (ARN) of the resiliency policy. The format for this ARN
	// is: arn:partition:resiliencehub:region:account:resiliency-policy/policy-id. For
	// more information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	PolicyArn *string

	// The description for the policy.
	PolicyDescription *string

	// The name of the policy
	PolicyName *string

	// The tags assigned to the resource. A tag is a label that you assign to an Amazon
	// Web Services resource. Each tag consists of a key/value pair.
	Tags map[string]string

	// The tier for this resiliency policy, ranging from the highest severity
	// (MissionCritical) to lowest (NonCritical).
	Tier ResiliencyPolicyTier

	noSmithyDocumentSerde
}

// The overall resiliency score, returned as an object that includes the disruption
// score and outage score.
type ResiliencyScore struct {

	// The disruption score for a valid key.
	//
	// This member is required.
	DisruptionScore map[string]float64

	// The outage score for a valid key.
	//
	// This member is required.
	Score float64

	noSmithyDocumentSerde
}

// Defines application resource errors.
type ResourceError struct {

	// This is the identifier of the resource.
	LogicalResourceId *string

	// This is the identifier of the physical resource.
	PhysicalResourceId *string

	// This is the error message.
	Reason *string

	noSmithyDocumentSerde
}

// A list of errors retrieving an application's resources.
type ResourceErrorsDetails struct {

	// This indicates if there are more errors not listed in the resourceErrors list.
	HasMoreErrors *bool

	// A list of errors retrieving an application's resources.
	ResourceErrors []ResourceError

	noSmithyDocumentSerde
}

// Defines a resource mapping.
type ResourceMapping struct {

	// Specifies the type of resource mapping. AppRegistryApp The resource is mapped to
	// another application. The name of the application is contained in the
	// appRegistryAppName property. CfnStack The resource is mapped to a CloudFormation
	// stack. The name of the CloudFormation stack is contained in the logicalStackName
	// property. Resource The resource is mapped to another resource. The name of the
	// resource is contained in the resourceName property. ResourceGroup The resource
	// is mapped to an Resource Groups. The name of the resource group is contained in
	// the resourceGroupName property.
	//
	// This member is required.
	MappingType ResourceMappingType

	// The identifier of this resource.
	//
	// This member is required.
	PhysicalResourceId *PhysicalResourceId

	// The name of the application this resource is mapped to.
	AppRegistryAppName *string

	// The name of the Amazon Elastic Kubernetes Service cluster and namespace this
	// resource belongs to. This parameter accepts values in "eks-cluster/namespace"
	// format.
	EksSourceName *string

	// The name of the CloudFormation stack this resource is mapped to.
	LogicalStackName *string

	// The name of the resource group this resource is mapped to.
	ResourceGroupName *string

	// The name of the resource this resource is mapped to.
	ResourceName *string

	// The short name of the Terraform source.
	TerraformSourceName *string

	noSmithyDocumentSerde
}

// The location of the Amazon S3 bucket.
type S3Location struct {

	// The name of the Amazon S3 bucket.
	Bucket *string

	// The prefix for the Amazon S3 bucket.
	Prefix *string

	noSmithyDocumentSerde
}

// Defines a standard operating procedure (SOP) recommendation.
type SopRecommendation struct {

	// Identifier for the SOP recommendation.
	//
	// This member is required.
	RecommendationId *string

	// The reference identifier for the SOP recommendation.
	//
	// This member is required.
	ReferenceId *string

	// The service type.
	//
	// This member is required.
	ServiceType SopServiceType

	// The name of the Application Component.
	AppComponentName *string

	// The description of the SOP recommendation.
	Description *string

	// The recommendation items.
	Items []RecommendationItem

	// The name of the SOP recommendation.
	Name *string

	// The prerequisite for the SOP recommendation.
	Prerequisite *string

	noSmithyDocumentSerde
}

// The Terraform s3 state file you need to import.
type TerraformSource struct {

	// The URL of the Terraform s3 state file you need to import.
	//
	// This member is required.
	S3StateFileUrl *string

	noSmithyDocumentSerde
}

// Defines a test recommendation.
type TestRecommendation struct {

	// The reference identifier for the test recommendation.
	//
	// This member is required.
	ReferenceId *string

	// The name of the Application Component.
	AppComponentName *string

	// A list of recommended alarms that are used in the test and must be exported
	// before or with the test.
	DependsOnAlarms []string

	// The description for the test recommendation.
	Description *string

	// The intent of the test recommendation.
	Intent *string

	// The test recommendation items.
	Items []RecommendationItem

	// The name of the test recommendation.
	Name *string

	// The prerequisite of the test recommendation.
	Prerequisite *string

	// Identifier for the test recommendation.
	RecommendationId *string

	// The level of risk for this test recommendation.
	Risk TestRisk

	// The type of test recommendation.
	Type TestType

	noSmithyDocumentSerde
}

// Defines a resource that is not supported by Resilience Hub.
type UnsupportedResource struct {

	// The logical resource identifier for the unsupported resource.
	//
	// This member is required.
	LogicalResourceId *LogicalResourceId

	// The physical resource identifier for the unsupported resource.
	//
	// This member is required.
	PhysicalResourceId *PhysicalResourceId

	// The type of resource.
	//
	// This member is required.
	ResourceType *string

	// The status of unsupported resource.
	UnsupportedResourceStatus *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

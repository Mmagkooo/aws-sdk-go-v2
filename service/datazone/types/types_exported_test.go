// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
)

func ExampleActionParameters_outputUsage() {
	var union types.ActionParameters
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ActionParametersMemberAwsConsoleLink:
		_ = v.Value // Value is types.AwsConsoleLinkParameters

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AwsConsoleLinkParameters

func ExampleAssetFilterConfiguration_outputUsage() {
	var union types.AssetFilterConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AssetFilterConfigurationMemberColumnConfiguration:
		_ = v.Value // Value is types.ColumnFilterConfiguration

	case *types.AssetFilterConfigurationMemberRowConfiguration:
		_ = v.Value // Value is types.RowFilterConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ColumnFilterConfiguration
var _ *types.RowFilterConfiguration

func ExampleDataSourceConfigurationInput_outputUsage() {
	var union types.DataSourceConfigurationInput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DataSourceConfigurationInputMemberGlueRunConfiguration:
		_ = v.Value // Value is types.GlueRunConfigurationInput

	case *types.DataSourceConfigurationInputMemberRedshiftRunConfiguration:
		_ = v.Value // Value is types.RedshiftRunConfigurationInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.RedshiftRunConfigurationInput
var _ *types.GlueRunConfigurationInput

func ExampleDataSourceConfigurationOutput_outputUsage() {
	var union types.DataSourceConfigurationOutput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DataSourceConfigurationOutputMemberGlueRunConfiguration:
		_ = v.Value // Value is types.GlueRunConfigurationOutput

	case *types.DataSourceConfigurationOutputMemberRedshiftRunConfiguration:
		_ = v.Value // Value is types.RedshiftRunConfigurationOutput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.GlueRunConfigurationOutput
var _ *types.RedshiftRunConfigurationOutput

func ExampleDomainUnitGrantFilter_outputUsage() {
	var union types.DomainUnitGrantFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DomainUnitGrantFilterMemberAllDomainUnitsGrantFilter:
		_ = v.Value // Value is types.AllDomainUnitsGrantFilter

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AllDomainUnitsGrantFilter

func ExampleDomainUnitOwnerProperties_outputUsage() {
	var union types.DomainUnitOwnerProperties
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DomainUnitOwnerPropertiesMemberGroup:
		_ = v.Value // Value is types.DomainUnitGroupProperties

	case *types.DomainUnitOwnerPropertiesMemberUser:
		_ = v.Value // Value is types.DomainUnitUserProperties

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DomainUnitGroupProperties
var _ *types.DomainUnitUserProperties

func ExampleFilterClause_outputUsage() {
	var union types.FilterClause
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FilterClauseMemberAnd:
		_ = v.Value // Value is []types.FilterClause

	case *types.FilterClauseMemberFilter:
		_ = v.Value // Value is types.Filter

	case *types.FilterClauseMemberOr:
		_ = v.Value // Value is []types.FilterClause

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.Filter
var _ []types.FilterClause

func ExampleGrantedEntity_outputUsage() {
	var union types.GrantedEntity
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.GrantedEntityMemberListing:
		_ = v.Value // Value is types.ListingRevision

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ListingRevision

func ExampleGrantedEntityInput_outputUsage() {
	var union types.GrantedEntityInput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.GrantedEntityInputMemberListing:
		_ = v.Value // Value is types.ListingRevisionInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ListingRevisionInput

func ExampleGroupPolicyGrantPrincipal_outputUsage() {
	var union types.GroupPolicyGrantPrincipal
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.GroupPolicyGrantPrincipalMemberGroupIdentifier:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleListingItem_outputUsage() {
	var union types.ListingItem
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ListingItemMemberAssetListing:
		_ = v.Value // Value is types.AssetListing

	case *types.ListingItemMemberDataProductListing:
		_ = v.Value // Value is types.DataProductListing

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AssetListing
var _ *types.DataProductListing

func ExampleMember_outputUsage() {
	var union types.Member
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MemberMemberGroupIdentifier:
		_ = v.Value // Value is string

	case *types.MemberMemberUserIdentifier:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleMemberDetails_outputUsage() {
	var union types.MemberDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MemberDetailsMemberGroup:
		_ = v.Value // Value is types.GroupDetails

	case *types.MemberDetailsMemberUser:
		_ = v.Value // Value is types.UserDetails

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.GroupDetails
var _ *types.UserDetails

func ExampleModel_outputUsage() {
	var union types.Model
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ModelMemberSmithy:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleOwnerProperties_outputUsage() {
	var union types.OwnerProperties
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OwnerPropertiesMemberGroup:
		_ = v.Value // Value is types.OwnerGroupProperties

	case *types.OwnerPropertiesMemberUser:
		_ = v.Value // Value is types.OwnerUserProperties

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OwnerGroupProperties
var _ *types.OwnerUserProperties

func ExampleOwnerPropertiesOutput_outputUsage() {
	var union types.OwnerPropertiesOutput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OwnerPropertiesOutputMemberGroup:
		_ = v.Value // Value is types.OwnerGroupPropertiesOutput

	case *types.OwnerPropertiesOutputMemberUser:
		_ = v.Value // Value is types.OwnerUserPropertiesOutput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OwnerGroupPropertiesOutput
var _ *types.OwnerUserPropertiesOutput

func ExamplePolicyGrantDetail_outputUsage() {
	var union types.PolicyGrantDetail
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PolicyGrantDetailMemberAddToProjectMemberPool:
		_ = v.Value // Value is types.AddToProjectMemberPoolPolicyGrantDetail

	case *types.PolicyGrantDetailMemberCreateAssetType:
		_ = v.Value // Value is types.CreateAssetTypePolicyGrantDetail

	case *types.PolicyGrantDetailMemberCreateDomainUnit:
		_ = v.Value // Value is types.CreateDomainUnitPolicyGrantDetail

	case *types.PolicyGrantDetailMemberCreateEnvironment:
		_ = v.Value // Value is types.Unit

	case *types.PolicyGrantDetailMemberCreateEnvironmentProfile:
		_ = v.Value // Value is types.CreateEnvironmentProfilePolicyGrantDetail

	case *types.PolicyGrantDetailMemberCreateFormType:
		_ = v.Value // Value is types.CreateFormTypePolicyGrantDetail

	case *types.PolicyGrantDetailMemberCreateGlossary:
		_ = v.Value // Value is types.CreateGlossaryPolicyGrantDetail

	case *types.PolicyGrantDetailMemberCreateProject:
		_ = v.Value // Value is types.CreateProjectPolicyGrantDetail

	case *types.PolicyGrantDetailMemberDelegateCreateEnvironmentProfile:
		_ = v.Value // Value is types.Unit

	case *types.PolicyGrantDetailMemberOverrideDomainUnitOwners:
		_ = v.Value // Value is types.OverrideDomainUnitOwnersPolicyGrantDetail

	case *types.PolicyGrantDetailMemberOverrideProjectOwners:
		_ = v.Value // Value is types.OverrideProjectOwnersPolicyGrantDetail

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.CreateDomainUnitPolicyGrantDetail
var _ *types.OverrideProjectOwnersPolicyGrantDetail
var _ *types.CreateEnvironmentProfilePolicyGrantDetail
var _ *types.CreateGlossaryPolicyGrantDetail
var _ *types.AddToProjectMemberPoolPolicyGrantDetail
var _ *types.CreateProjectPolicyGrantDetail
var _ *types.OverrideDomainUnitOwnersPolicyGrantDetail
var _ *types.CreateAssetTypePolicyGrantDetail
var _ *types.Unit
var _ *types.CreateFormTypePolicyGrantDetail

func ExamplePolicyGrantPrincipal_outputUsage() {
	var union types.PolicyGrantPrincipal
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PolicyGrantPrincipalMemberDomainUnit:
		_ = v.Value // Value is types.DomainUnitPolicyGrantPrincipal

	case *types.PolicyGrantPrincipalMemberGroup:
		_ = v.Value // Value is types.GroupPolicyGrantPrincipal

	case *types.PolicyGrantPrincipalMemberProject:
		_ = v.Value // Value is types.ProjectPolicyGrantPrincipal

	case *types.PolicyGrantPrincipalMemberUser:
		_ = v.Value // Value is types.UserPolicyGrantPrincipal

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DomainUnitPolicyGrantPrincipal
var _ types.GroupPolicyGrantPrincipal
var _ types.UserPolicyGrantPrincipal
var _ *types.ProjectPolicyGrantPrincipal

func ExampleProjectGrantFilter_outputUsage() {
	var union types.ProjectGrantFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ProjectGrantFilterMemberDomainUnitFilter:
		_ = v.Value // Value is types.DomainUnitFilterForProject

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DomainUnitFilterForProject

func ExampleProvisioningConfiguration_outputUsage() {
	var union types.ProvisioningConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ProvisioningConfigurationMemberLakeFormationConfiguration:
		_ = v.Value // Value is types.LakeFormationConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.LakeFormationConfiguration

func ExampleProvisioningProperties_outputUsage() {
	var union types.ProvisioningProperties
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ProvisioningPropertiesMemberCloudFormation:
		_ = v.Value // Value is types.CloudFormationProperties

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.CloudFormationProperties

func ExampleRedshiftStorage_outputUsage() {
	var union types.RedshiftStorage
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RedshiftStorageMemberRedshiftClusterSource:
		_ = v.Value // Value is types.RedshiftClusterStorage

	case *types.RedshiftStorageMemberRedshiftServerlessSource:
		_ = v.Value // Value is types.RedshiftServerlessStorage

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.RedshiftClusterStorage
var _ *types.RedshiftServerlessStorage

func ExampleRowFilter_outputUsage() {
	var union types.RowFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RowFilterMemberAnd:
		_ = v.Value // Value is []types.RowFilter

	case *types.RowFilterMemberExpression:
		_ = v.Value // Value is types.RowFilterExpression

	case *types.RowFilterMemberOr:
		_ = v.Value // Value is []types.RowFilter

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ types.RowFilterExpression
var _ []types.RowFilter

func ExampleRowFilterExpression_outputUsage() {
	var union types.RowFilterExpression
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RowFilterExpressionMemberEqualTo:
		_ = v.Value // Value is types.EqualToExpression

	case *types.RowFilterExpressionMemberGreaterThan:
		_ = v.Value // Value is types.GreaterThanExpression

	case *types.RowFilterExpressionMemberGreaterThanOrEqualTo:
		_ = v.Value // Value is types.GreaterThanOrEqualToExpression

	case *types.RowFilterExpressionMemberIn:
		_ = v.Value // Value is types.InExpression

	case *types.RowFilterExpressionMemberIsNotNull:
		_ = v.Value // Value is types.IsNotNullExpression

	case *types.RowFilterExpressionMemberIsNull:
		_ = v.Value // Value is types.IsNullExpression

	case *types.RowFilterExpressionMemberLessThan:
		_ = v.Value // Value is types.LessThanExpression

	case *types.RowFilterExpressionMemberLessThanOrEqualTo:
		_ = v.Value // Value is types.LessThanOrEqualToExpression

	case *types.RowFilterExpressionMemberLike:
		_ = v.Value // Value is types.LikeExpression

	case *types.RowFilterExpressionMemberNotEqualTo:
		_ = v.Value // Value is types.NotEqualToExpression

	case *types.RowFilterExpressionMemberNotIn:
		_ = v.Value // Value is types.NotInExpression

	case *types.RowFilterExpressionMemberNotLike:
		_ = v.Value // Value is types.NotLikeExpression

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.NotLikeExpression
var _ *types.GreaterThanExpression
var _ *types.LessThanExpression
var _ *types.IsNotNullExpression
var _ *types.NotEqualToExpression
var _ *types.GreaterThanOrEqualToExpression
var _ *types.IsNullExpression
var _ *types.LessThanOrEqualToExpression
var _ *types.LikeExpression
var _ *types.NotInExpression
var _ *types.InExpression
var _ *types.EqualToExpression

func ExampleSearchInventoryResultItem_outputUsage() {
	var union types.SearchInventoryResultItem
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SearchInventoryResultItemMemberAssetItem:
		_ = v.Value // Value is types.AssetItem

	case *types.SearchInventoryResultItemMemberDataProductItem:
		_ = v.Value // Value is types.DataProductResultItem

	case *types.SearchInventoryResultItemMemberGlossaryItem:
		_ = v.Value // Value is types.GlossaryItem

	case *types.SearchInventoryResultItemMemberGlossaryTermItem:
		_ = v.Value // Value is types.GlossaryTermItem

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.GlossaryItem
var _ *types.DataProductResultItem
var _ *types.AssetItem
var _ *types.GlossaryTermItem

func ExampleSearchResultItem_outputUsage() {
	var union types.SearchResultItem
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SearchResultItemMemberAssetListing:
		_ = v.Value // Value is types.AssetListingItem

	case *types.SearchResultItemMemberDataProductListing:
		_ = v.Value // Value is types.DataProductListingItem

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DataProductListingItem
var _ *types.AssetListingItem

func ExampleSearchTypesResultItem_outputUsage() {
	var union types.SearchTypesResultItem
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SearchTypesResultItemMemberAssetTypeItem:
		_ = v.Value // Value is types.AssetTypeItem

	case *types.SearchTypesResultItemMemberFormTypeItem:
		_ = v.Value // Value is types.FormTypeData

	case *types.SearchTypesResultItemMemberLineageNodeTypeItem:
		_ = v.Value // Value is types.LineageNodeTypeItem

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.LineageNodeTypeItem
var _ *types.FormTypeData
var _ *types.AssetTypeItem

func ExampleSelfGrantStatusOutput_outputUsage() {
	var union types.SelfGrantStatusOutput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SelfGrantStatusOutputMemberGlueSelfGrantStatus:
		_ = v.Value // Value is types.GlueSelfGrantStatusOutput

	case *types.SelfGrantStatusOutputMemberRedshiftSelfGrantStatus:
		_ = v.Value // Value is types.RedshiftSelfGrantStatusOutput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.RedshiftSelfGrantStatusOutput
var _ *types.GlueSelfGrantStatusOutput

func ExampleSubscribedListingItem_outputUsage() {
	var union types.SubscribedListingItem
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SubscribedListingItemMemberAssetListing:
		_ = v.Value // Value is types.SubscribedAssetListing

	case *types.SubscribedListingItemMemberProductListing:
		_ = v.Value // Value is types.SubscribedProductListing

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SubscribedAssetListing
var _ *types.SubscribedProductListing

func ExampleSubscribedPrincipal_outputUsage() {
	var union types.SubscribedPrincipal
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SubscribedPrincipalMemberProject:
		_ = v.Value // Value is types.SubscribedProject

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SubscribedProject

func ExampleSubscribedPrincipalInput_outputUsage() {
	var union types.SubscribedPrincipalInput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SubscribedPrincipalInputMemberProject:
		_ = v.Value // Value is types.SubscribedProjectInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SubscribedProjectInput

func ExampleUserPolicyGrantPrincipal_outputUsage() {
	var union types.UserPolicyGrantPrincipal
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.UserPolicyGrantPrincipalMemberAllUsersGrantFilter:
		_ = v.Value // Value is types.AllUsersGrantFilter

	case *types.UserPolicyGrantPrincipalMemberUserIdentifier:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *types.AllUsersGrantFilter

func ExampleUserProfileDetails_outputUsage() {
	var union types.UserProfileDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.UserProfileDetailsMemberIam:
		_ = v.Value // Value is types.IamUserProfileDetails

	case *types.UserProfileDetailsMemberSso:
		_ = v.Value // Value is types.SsoUserProfileDetails

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SsoUserProfileDetails
var _ *types.IamUserProfileDetails

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdMarkerDash string

// Enum values for AdMarkerDash
const (
	AdMarkerDashBinary AdMarkerDash = "BINARY"
	AdMarkerDashXml    AdMarkerDash = "XML"
)

// Values returns all known values for AdMarkerDash. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AdMarkerDash) Values() []AdMarkerDash {
	return []AdMarkerDash{
		"BINARY",
		"XML",
	}
}

type AdMarkerHls string

// Enum values for AdMarkerHls
const (
	AdMarkerHlsDaterange AdMarkerHls = "DATERANGE"
)

// Values returns all known values for AdMarkerHls. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AdMarkerHls) Values() []AdMarkerHls {
	return []AdMarkerHls{
		"DATERANGE",
	}
}

type CmafEncryptionMethod string

// Enum values for CmafEncryptionMethod
const (
	CmafEncryptionMethodCenc CmafEncryptionMethod = "CENC"
	CmafEncryptionMethodCbcs CmafEncryptionMethod = "CBCS"
)

// Values returns all known values for CmafEncryptionMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CmafEncryptionMethod) Values() []CmafEncryptionMethod {
	return []CmafEncryptionMethod{
		"CENC",
		"CBCS",
	}
}

type ConflictExceptionType string

// Enum values for ConflictExceptionType
const (
	ConflictExceptionTypeResourceInUse               ConflictExceptionType = "RESOURCE_IN_USE"
	ConflictExceptionTypeResourceAlreadyExists       ConflictExceptionType = "RESOURCE_ALREADY_EXISTS"
	ConflictExceptionTypeIdempotentParameterMismatch ConflictExceptionType = "IDEMPOTENT_PARAMETER_MISMATCH"
	ConflictExceptionTypeConflictingOperation        ConflictExceptionType = "CONFLICTING_OPERATION"
)

// Values returns all known values for ConflictExceptionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConflictExceptionType) Values() []ConflictExceptionType {
	return []ConflictExceptionType{
		"RESOURCE_IN_USE",
		"RESOURCE_ALREADY_EXISTS",
		"IDEMPOTENT_PARAMETER_MISMATCH",
		"CONFLICTING_OPERATION",
	}
}

type ContainerType string

// Enum values for ContainerType
const (
	ContainerTypeTs   ContainerType = "TS"
	ContainerTypeCmaf ContainerType = "CMAF"
)

// Values returns all known values for ContainerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ContainerType) Values() []ContainerType {
	return []ContainerType{
		"TS",
		"CMAF",
	}
}

type DashDrmSignaling string

// Enum values for DashDrmSignaling
const (
	DashDrmSignalingIndividual DashDrmSignaling = "INDIVIDUAL"
	DashDrmSignalingReferenced DashDrmSignaling = "REFERENCED"
)

// Values returns all known values for DashDrmSignaling. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DashDrmSignaling) Values() []DashDrmSignaling {
	return []DashDrmSignaling{
		"INDIVIDUAL",
		"REFERENCED",
	}
}

type DashPeriodTrigger string

// Enum values for DashPeriodTrigger
const (
	DashPeriodTriggerAvails            DashPeriodTrigger = "AVAILS"
	DashPeriodTriggerDrmKeyRotation    DashPeriodTrigger = "DRM_KEY_ROTATION"
	DashPeriodTriggerSourceChanges     DashPeriodTrigger = "SOURCE_CHANGES"
	DashPeriodTriggerSourceDisruptions DashPeriodTrigger = "SOURCE_DISRUPTIONS"
	DashPeriodTriggerNone              DashPeriodTrigger = "NONE"
)

// Values returns all known values for DashPeriodTrigger. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DashPeriodTrigger) Values() []DashPeriodTrigger {
	return []DashPeriodTrigger{
		"AVAILS",
		"DRM_KEY_ROTATION",
		"SOURCE_CHANGES",
		"SOURCE_DISRUPTIONS",
		"NONE",
	}
}

type DashSegmentTemplateFormat string

// Enum values for DashSegmentTemplateFormat
const (
	DashSegmentTemplateFormatNumberWithTimeline DashSegmentTemplateFormat = "NUMBER_WITH_TIMELINE"
)

// Values returns all known values for DashSegmentTemplateFormat. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DashSegmentTemplateFormat) Values() []DashSegmentTemplateFormat {
	return []DashSegmentTemplateFormat{
		"NUMBER_WITH_TIMELINE",
	}
}

type DashUtcTimingMode string

// Enum values for DashUtcTimingMode
const (
	DashUtcTimingModeHttpHead   DashUtcTimingMode = "HTTP_HEAD"
	DashUtcTimingModeHttpIso    DashUtcTimingMode = "HTTP_ISO"
	DashUtcTimingModeHttpXsdate DashUtcTimingMode = "HTTP_XSDATE"
	DashUtcTimingModeUtcDirect  DashUtcTimingMode = "UTC_DIRECT"
)

// Values returns all known values for DashUtcTimingMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DashUtcTimingMode) Values() []DashUtcTimingMode {
	return []DashUtcTimingMode{
		"HTTP_HEAD",
		"HTTP_ISO",
		"HTTP_XSDATE",
		"UTC_DIRECT",
	}
}

type DrmSystem string

// Enum values for DrmSystem
const (
	DrmSystemClearKeyAes128 DrmSystem = "CLEAR_KEY_AES_128"
	DrmSystemFairplay       DrmSystem = "FAIRPLAY"
	DrmSystemPlayready      DrmSystem = "PLAYREADY"
	DrmSystemWidevine       DrmSystem = "WIDEVINE"
	DrmSystemIrdeto         DrmSystem = "IRDETO"
)

// Values returns all known values for DrmSystem. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DrmSystem) Values() []DrmSystem {
	return []DrmSystem{
		"CLEAR_KEY_AES_128",
		"FAIRPLAY",
		"PLAYREADY",
		"WIDEVINE",
		"IRDETO",
	}
}

type EndpointErrorCondition string

// Enum values for EndpointErrorCondition
const (
	EndpointErrorConditionStaleManifest      EndpointErrorCondition = "STALE_MANIFEST"
	EndpointErrorConditionIncompleteManifest EndpointErrorCondition = "INCOMPLETE_MANIFEST"
	EndpointErrorConditionMissingDrmKey      EndpointErrorCondition = "MISSING_DRM_KEY"
	EndpointErrorConditionSlateInput         EndpointErrorCondition = "SLATE_INPUT"
)

// Values returns all known values for EndpointErrorCondition. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EndpointErrorCondition) Values() []EndpointErrorCondition {
	return []EndpointErrorCondition{
		"STALE_MANIFEST",
		"INCOMPLETE_MANIFEST",
		"MISSING_DRM_KEY",
		"SLATE_INPUT",
	}
}

type InputType string

// Enum values for InputType
const (
	InputTypeHls  InputType = "HLS"
	InputTypeCmaf InputType = "CMAF"
)

// Values returns all known values for InputType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InputType) Values() []InputType {
	return []InputType{
		"HLS",
		"CMAF",
	}
}

type PresetSpeke20Audio string

// Enum values for PresetSpeke20Audio
const (
	PresetSpeke20AudioPresetAudio1 PresetSpeke20Audio = "PRESET_AUDIO_1"
	PresetSpeke20AudioPresetAudio2 PresetSpeke20Audio = "PRESET_AUDIO_2"
	PresetSpeke20AudioPresetAudio3 PresetSpeke20Audio = "PRESET_AUDIO_3"
	PresetSpeke20AudioShared       PresetSpeke20Audio = "SHARED"
	PresetSpeke20AudioUnencrypted  PresetSpeke20Audio = "UNENCRYPTED"
)

// Values returns all known values for PresetSpeke20Audio. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PresetSpeke20Audio) Values() []PresetSpeke20Audio {
	return []PresetSpeke20Audio{
		"PRESET_AUDIO_1",
		"PRESET_AUDIO_2",
		"PRESET_AUDIO_3",
		"SHARED",
		"UNENCRYPTED",
	}
}

type PresetSpeke20Video string

// Enum values for PresetSpeke20Video
const (
	PresetSpeke20VideoPresetVideo1 PresetSpeke20Video = "PRESET_VIDEO_1"
	PresetSpeke20VideoPresetVideo2 PresetSpeke20Video = "PRESET_VIDEO_2"
	PresetSpeke20VideoPresetVideo3 PresetSpeke20Video = "PRESET_VIDEO_3"
	PresetSpeke20VideoPresetVideo4 PresetSpeke20Video = "PRESET_VIDEO_4"
	PresetSpeke20VideoPresetVideo5 PresetSpeke20Video = "PRESET_VIDEO_5"
	PresetSpeke20VideoPresetVideo6 PresetSpeke20Video = "PRESET_VIDEO_6"
	PresetSpeke20VideoPresetVideo7 PresetSpeke20Video = "PRESET_VIDEO_7"
	PresetSpeke20VideoPresetVideo8 PresetSpeke20Video = "PRESET_VIDEO_8"
	PresetSpeke20VideoShared       PresetSpeke20Video = "SHARED"
	PresetSpeke20VideoUnencrypted  PresetSpeke20Video = "UNENCRYPTED"
)

// Values returns all known values for PresetSpeke20Video. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PresetSpeke20Video) Values() []PresetSpeke20Video {
	return []PresetSpeke20Video{
		"PRESET_VIDEO_1",
		"PRESET_VIDEO_2",
		"PRESET_VIDEO_3",
		"PRESET_VIDEO_4",
		"PRESET_VIDEO_5",
		"PRESET_VIDEO_6",
		"PRESET_VIDEO_7",
		"PRESET_VIDEO_8",
		"SHARED",
		"UNENCRYPTED",
	}
}

type ResourceTypeNotFound string

// Enum values for ResourceTypeNotFound
const (
	ResourceTypeNotFoundChannelGroup   ResourceTypeNotFound = "CHANNEL_GROUP"
	ResourceTypeNotFoundChannel        ResourceTypeNotFound = "CHANNEL"
	ResourceTypeNotFoundOriginEndpoint ResourceTypeNotFound = "ORIGIN_ENDPOINT"
)

// Values returns all known values for ResourceTypeNotFound. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceTypeNotFound) Values() []ResourceTypeNotFound {
	return []ResourceTypeNotFound{
		"CHANNEL_GROUP",
		"CHANNEL",
		"ORIGIN_ENDPOINT",
	}
}

type ScteFilter string

// Enum values for ScteFilter
const (
	ScteFilterSpliceInsert                           ScteFilter = "SPLICE_INSERT"
	ScteFilterBreak                                  ScteFilter = "BREAK"
	ScteFilterProviderAdvertisement                  ScteFilter = "PROVIDER_ADVERTISEMENT"
	ScteFilterDistributorAdvertisement               ScteFilter = "DISTRIBUTOR_ADVERTISEMENT"
	ScteFilterProviderPlacementOpportunity           ScteFilter = "PROVIDER_PLACEMENT_OPPORTUNITY"
	ScteFilterDistributorPlacementOpportunity        ScteFilter = "DISTRIBUTOR_PLACEMENT_OPPORTUNITY"
	ScteFilterProviderOverlayPlacementOpportunity    ScteFilter = "PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY"
	ScteFilterDistributorOverlayPlacementOpportunity ScteFilter = "DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY"
	ScteFilterProgram                                ScteFilter = "PROGRAM"
)

// Values returns all known values for ScteFilter. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScteFilter) Values() []ScteFilter {
	return []ScteFilter{
		"SPLICE_INSERT",
		"BREAK",
		"PROVIDER_ADVERTISEMENT",
		"DISTRIBUTOR_ADVERTISEMENT",
		"PROVIDER_PLACEMENT_OPPORTUNITY",
		"DISTRIBUTOR_PLACEMENT_OPPORTUNITY",
		"PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY",
		"DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY",
		"PROGRAM",
	}
}

type TsEncryptionMethod string

// Enum values for TsEncryptionMethod
const (
	TsEncryptionMethodAes128    TsEncryptionMethod = "AES_128"
	TsEncryptionMethodSampleAes TsEncryptionMethod = "SAMPLE_AES"
)

// Values returns all known values for TsEncryptionMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TsEncryptionMethod) Values() []TsEncryptionMethod {
	return []TsEncryptionMethod{
		"AES_128",
		"SAMPLE_AES",
	}
}

type ValidationExceptionType string

// Enum values for ValidationExceptionType
const (
	ValidationExceptionTypeContainerTypeImmutable                                ValidationExceptionType = "CONTAINER_TYPE_IMMUTABLE"
	ValidationExceptionTypeInvalidPaginationToken                                ValidationExceptionType = "INVALID_PAGINATION_TOKEN"
	ValidationExceptionTypeInvalidPaginationMaxResults                           ValidationExceptionType = "INVALID_PAGINATION_MAX_RESULTS"
	ValidationExceptionTypeInvalidPolicy                                         ValidationExceptionType = "INVALID_POLICY"
	ValidationExceptionTypeInvalidRoleArn                                        ValidationExceptionType = "INVALID_ROLE_ARN"
	ValidationExceptionTypeManifestNameCollision                                 ValidationExceptionType = "MANIFEST_NAME_COLLISION"
	ValidationExceptionTypeEncryptionMethodContainerTypeMismatch                 ValidationExceptionType = "ENCRYPTION_METHOD_CONTAINER_TYPE_MISMATCH"
	ValidationExceptionTypeCencIvIncompatible                                    ValidationExceptionType = "CENC_IV_INCOMPATIBLE"
	ValidationExceptionTypeEncryptionContractWithoutAudioRenditionIncompatible   ValidationExceptionType = "ENCRYPTION_CONTRACT_WITHOUT_AUDIO_RENDITION_INCOMPATIBLE"
	ValidationExceptionTypeEncryptionContractUnencrypted                         ValidationExceptionType = "ENCRYPTION_CONTRACT_UNENCRYPTED"
	ValidationExceptionTypeEncryptionContractShared                              ValidationExceptionType = "ENCRYPTION_CONTRACT_SHARED"
	ValidationExceptionTypeNumManifestsLow                                       ValidationExceptionType = "NUM_MANIFESTS_LOW"
	ValidationExceptionTypeNumManifestsHigh                                      ValidationExceptionType = "NUM_MANIFESTS_HIGH"
	ValidationExceptionTypeManifestDrmSystemsIncompatible                        ValidationExceptionType = "MANIFEST_DRM_SYSTEMS_INCOMPATIBLE"
	ValidationExceptionTypeDrmSystemsEncryptionMethodIncompatible                ValidationExceptionType = "DRM_SYSTEMS_ENCRYPTION_METHOD_INCOMPATIBLE"
	ValidationExceptionTypeRoleArnNotAssumable                                   ValidationExceptionType = "ROLE_ARN_NOT_ASSUMABLE"
	ValidationExceptionTypeRoleArnLengthOutOfRange                               ValidationExceptionType = "ROLE_ARN_LENGTH_OUT_OF_RANGE"
	ValidationExceptionTypeRoleArnInvalidFormat                                  ValidationExceptionType = "ROLE_ARN_INVALID_FORMAT"
	ValidationExceptionTypeUrlInvalid                                            ValidationExceptionType = "URL_INVALID"
	ValidationExceptionTypeUrlScheme                                             ValidationExceptionType = "URL_SCHEME"
	ValidationExceptionTypeUrlUserInfo                                           ValidationExceptionType = "URL_USER_INFO"
	ValidationExceptionTypeUrlPort                                               ValidationExceptionType = "URL_PORT"
	ValidationExceptionTypeUrlUnknownHost                                        ValidationExceptionType = "URL_UNKNOWN_HOST"
	ValidationExceptionTypeUrlLocalAddress                                       ValidationExceptionType = "URL_LOCAL_ADDRESS"
	ValidationExceptionTypeUrlLoopbackAddress                                    ValidationExceptionType = "URL_LOOPBACK_ADDRESS"
	ValidationExceptionTypeUrlLinkLocalAddress                                   ValidationExceptionType = "URL_LINK_LOCAL_ADDRESS"
	ValidationExceptionTypeUrlMulticastAddress                                   ValidationExceptionType = "URL_MULTICAST_ADDRESS"
	ValidationExceptionTypeMemberInvalid                                         ValidationExceptionType = "MEMBER_INVALID"
	ValidationExceptionTypeMemberMissing                                         ValidationExceptionType = "MEMBER_MISSING"
	ValidationExceptionTypeMemberMinValue                                        ValidationExceptionType = "MEMBER_MIN_VALUE"
	ValidationExceptionTypeMemberMaxValue                                        ValidationExceptionType = "MEMBER_MAX_VALUE"
	ValidationExceptionTypeMemberMinLength                                       ValidationExceptionType = "MEMBER_MIN_LENGTH"
	ValidationExceptionTypeMemberMaxLength                                       ValidationExceptionType = "MEMBER_MAX_LENGTH"
	ValidationExceptionTypeMemberInvalidEnumValue                                ValidationExceptionType = "MEMBER_INVALID_ENUM_VALUE"
	ValidationExceptionTypeMemberDoesNotMatchPattern                             ValidationExceptionType = "MEMBER_DOES_NOT_MATCH_PATTERN"
	ValidationExceptionTypeInvalidManifestFilter                                 ValidationExceptionType = "INVALID_MANIFEST_FILTER"
	ValidationExceptionTypeInvalidTimeDelaySeconds                               ValidationExceptionType = "INVALID_TIME_DELAY_SECONDS"
	ValidationExceptionTypeEndTimeEarlierThanStartTime                           ValidationExceptionType = "END_TIME_EARLIER_THAN_START_TIME"
	ValidationExceptionTypeTsContainerTypeWithDashManifest                       ValidationExceptionType = "TS_CONTAINER_TYPE_WITH_DASH_MANIFEST"
	ValidationExceptionTypeDirectModeWithTimingSource                            ValidationExceptionType = "DIRECT_MODE_WITH_TIMING_SOURCE"
	ValidationExceptionTypeNoneModeWithTimingSource                              ValidationExceptionType = "NONE_MODE_WITH_TIMING_SOURCE"
	ValidationExceptionTypeTimingSourceMissing                                   ValidationExceptionType = "TIMING_SOURCE_MISSING"
	ValidationExceptionTypeUpdatePeriodSmallerThanSegmentDuration                ValidationExceptionType = "UPDATE_PERIOD_SMALLER_THAN_SEGMENT_DURATION"
	ValidationExceptionTypePeriodTriggersNoneSpecifiedWithAdditionalValues       ValidationExceptionType = "PERIOD_TRIGGERS_NONE_SPECIFIED_WITH_ADDITIONAL_VALUES"
	ValidationExceptionTypeDrmSignalingMismatchSegmentEncryptionStatus           ValidationExceptionType = "DRM_SIGNALING_MISMATCH_SEGMENT_ENCRYPTION_STATUS"
	ValidationExceptionTypeOnlyCmafInputTypeAllowForceEndpointErrorConfiguration ValidationExceptionType = "ONLY_CMAF_INPUT_TYPE_ALLOW_FORCE_ENDPOINT_ERROR_CONFIGURATION"
	ValidationExceptionTypeSourceDisruptionsEnabledIncorrectly                   ValidationExceptionType = "SOURCE_DISRUPTIONS_ENABLED_INCORRECTLY"
)

// Values returns all known values for ValidationExceptionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionType) Values() []ValidationExceptionType {
	return []ValidationExceptionType{
		"CONTAINER_TYPE_IMMUTABLE",
		"INVALID_PAGINATION_TOKEN",
		"INVALID_PAGINATION_MAX_RESULTS",
		"INVALID_POLICY",
		"INVALID_ROLE_ARN",
		"MANIFEST_NAME_COLLISION",
		"ENCRYPTION_METHOD_CONTAINER_TYPE_MISMATCH",
		"CENC_IV_INCOMPATIBLE",
		"ENCRYPTION_CONTRACT_WITHOUT_AUDIO_RENDITION_INCOMPATIBLE",
		"ENCRYPTION_CONTRACT_UNENCRYPTED",
		"ENCRYPTION_CONTRACT_SHARED",
		"NUM_MANIFESTS_LOW",
		"NUM_MANIFESTS_HIGH",
		"MANIFEST_DRM_SYSTEMS_INCOMPATIBLE",
		"DRM_SYSTEMS_ENCRYPTION_METHOD_INCOMPATIBLE",
		"ROLE_ARN_NOT_ASSUMABLE",
		"ROLE_ARN_LENGTH_OUT_OF_RANGE",
		"ROLE_ARN_INVALID_FORMAT",
		"URL_INVALID",
		"URL_SCHEME",
		"URL_USER_INFO",
		"URL_PORT",
		"URL_UNKNOWN_HOST",
		"URL_LOCAL_ADDRESS",
		"URL_LOOPBACK_ADDRESS",
		"URL_LINK_LOCAL_ADDRESS",
		"URL_MULTICAST_ADDRESS",
		"MEMBER_INVALID",
		"MEMBER_MISSING",
		"MEMBER_MIN_VALUE",
		"MEMBER_MAX_VALUE",
		"MEMBER_MIN_LENGTH",
		"MEMBER_MAX_LENGTH",
		"MEMBER_INVALID_ENUM_VALUE",
		"MEMBER_DOES_NOT_MATCH_PATTERN",
		"INVALID_MANIFEST_FILTER",
		"INVALID_TIME_DELAY_SECONDS",
		"END_TIME_EARLIER_THAN_START_TIME",
		"TS_CONTAINER_TYPE_WITH_DASH_MANIFEST",
		"DIRECT_MODE_WITH_TIMING_SOURCE",
		"NONE_MODE_WITH_TIMING_SOURCE",
		"TIMING_SOURCE_MISSING",
		"UPDATE_PERIOD_SMALLER_THAN_SEGMENT_DURATION",
		"PERIOD_TRIGGERS_NONE_SPECIFIED_WITH_ADDITIONAL_VALUES",
		"DRM_SIGNALING_MISMATCH_SEGMENT_ENCRYPTION_STATUS",
		"ONLY_CMAF_INPUT_TYPE_ALLOW_FORCE_ENDPOINT_ERROR_CONFIGURATION",
		"SOURCE_DISRUPTIONS_ENABLED_INCORRECTLY",
	}
}

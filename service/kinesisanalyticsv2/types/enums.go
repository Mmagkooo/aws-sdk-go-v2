// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplicationMode string

// Enum values for ApplicationMode
const (
	ApplicationModeStreaming   ApplicationMode = "STREAMING"
	ApplicationModeInteractive ApplicationMode = "INTERACTIVE"
)

// Values returns all known values for ApplicationMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationMode) Values() []ApplicationMode {
	return []ApplicationMode{
		"STREAMING",
		"INTERACTIVE",
	}
}

type ApplicationRestoreType string

// Enum values for ApplicationRestoreType
const (
	ApplicationRestoreTypeSkipRestoreFromSnapshot   ApplicationRestoreType = "SKIP_RESTORE_FROM_SNAPSHOT"
	ApplicationRestoreTypeRestoreFromLatestSnapshot ApplicationRestoreType = "RESTORE_FROM_LATEST_SNAPSHOT"
	ApplicationRestoreTypeRestoreFromCustomSnapshot ApplicationRestoreType = "RESTORE_FROM_CUSTOM_SNAPSHOT"
)

// Values returns all known values for ApplicationRestoreType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationRestoreType) Values() []ApplicationRestoreType {
	return []ApplicationRestoreType{
		"SKIP_RESTORE_FROM_SNAPSHOT",
		"RESTORE_FROM_LATEST_SNAPSHOT",
		"RESTORE_FROM_CUSTOM_SNAPSHOT",
	}
}

type ApplicationStatus string

// Enum values for ApplicationStatus
const (
	ApplicationStatusDeleting      ApplicationStatus = "DELETING"
	ApplicationStatusStarting      ApplicationStatus = "STARTING"
	ApplicationStatusStopping      ApplicationStatus = "STOPPING"
	ApplicationStatusReady         ApplicationStatus = "READY"
	ApplicationStatusRunning       ApplicationStatus = "RUNNING"
	ApplicationStatusUpdating      ApplicationStatus = "UPDATING"
	ApplicationStatusAutoscaling   ApplicationStatus = "AUTOSCALING"
	ApplicationStatusForceStopping ApplicationStatus = "FORCE_STOPPING"
	ApplicationStatusRollingBack   ApplicationStatus = "ROLLING_BACK"
	ApplicationStatusMaintenance   ApplicationStatus = "MAINTENANCE"
	ApplicationStatusRolledBack    ApplicationStatus = "ROLLED_BACK"
)

// Values returns all known values for ApplicationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationStatus) Values() []ApplicationStatus {
	return []ApplicationStatus{
		"DELETING",
		"STARTING",
		"STOPPING",
		"READY",
		"RUNNING",
		"UPDATING",
		"AUTOSCALING",
		"FORCE_STOPPING",
		"ROLLING_BACK",
		"MAINTENANCE",
		"ROLLED_BACK",
	}
}

type ArtifactType string

// Enum values for ArtifactType
const (
	ArtifactTypeUdf           ArtifactType = "UDF"
	ArtifactTypeDependencyJar ArtifactType = "DEPENDENCY_JAR"
)

// Values returns all known values for ArtifactType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArtifactType) Values() []ArtifactType {
	return []ArtifactType{
		"UDF",
		"DEPENDENCY_JAR",
	}
}

type CodeContentType string

// Enum values for CodeContentType
const (
	CodeContentTypePlaintext CodeContentType = "PLAINTEXT"
	CodeContentTypeZipfile   CodeContentType = "ZIPFILE"
)

// Values returns all known values for CodeContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CodeContentType) Values() []CodeContentType {
	return []CodeContentType{
		"PLAINTEXT",
		"ZIPFILE",
	}
}

type ConfigurationType string

// Enum values for ConfigurationType
const (
	ConfigurationTypeDefault ConfigurationType = "DEFAULT"
	ConfigurationTypeCustom  ConfigurationType = "CUSTOM"
)

// Values returns all known values for ConfigurationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationType) Values() []ConfigurationType {
	return []ConfigurationType{
		"DEFAULT",
		"CUSTOM",
	}
}

type InputStartingPosition string

// Enum values for InputStartingPosition
const (
	InputStartingPositionNow              InputStartingPosition = "NOW"
	InputStartingPositionTrimHorizon      InputStartingPosition = "TRIM_HORIZON"
	InputStartingPositionLastStoppedPoint InputStartingPosition = "LAST_STOPPED_POINT"
)

// Values returns all known values for InputStartingPosition. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InputStartingPosition) Values() []InputStartingPosition {
	return []InputStartingPosition{
		"NOW",
		"TRIM_HORIZON",
		"LAST_STOPPED_POINT",
	}
}

type LogLevel string

// Enum values for LogLevel
const (
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
	LogLevelDebug LogLevel = "DEBUG"
)

// Values returns all known values for LogLevel. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogLevel) Values() []LogLevel {
	return []LogLevel{
		"INFO",
		"WARN",
		"ERROR",
		"DEBUG",
	}
}

type MetricsLevel string

// Enum values for MetricsLevel
const (
	MetricsLevelApplication MetricsLevel = "APPLICATION"
	MetricsLevelTask        MetricsLevel = "TASK"
	MetricsLevelOperator    MetricsLevel = "OPERATOR"
	MetricsLevelParallelism MetricsLevel = "PARALLELISM"
)

// Values returns all known values for MetricsLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MetricsLevel) Values() []MetricsLevel {
	return []MetricsLevel{
		"APPLICATION",
		"TASK",
		"OPERATOR",
		"PARALLELISM",
	}
}

type RecordFormatType string

// Enum values for RecordFormatType
const (
	RecordFormatTypeJson RecordFormatType = "JSON"
	RecordFormatTypeCsv  RecordFormatType = "CSV"
)

// Values returns all known values for RecordFormatType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecordFormatType) Values() []RecordFormatType {
	return []RecordFormatType{
		"JSON",
		"CSV",
	}
}

type RuntimeEnvironment string

// Enum values for RuntimeEnvironment
const (
	RuntimeEnvironmentSql10           RuntimeEnvironment = "SQL-1_0"
	RuntimeEnvironmentFlink16         RuntimeEnvironment = "FLINK-1_6"
	RuntimeEnvironmentFlink18         RuntimeEnvironment = "FLINK-1_8"
	RuntimeEnvironmentZeppelinFlink10 RuntimeEnvironment = "ZEPPELIN-FLINK-1_0"
	RuntimeEnvironmentFlink111        RuntimeEnvironment = "FLINK-1_11"
	RuntimeEnvironmentFlink113        RuntimeEnvironment = "FLINK-1_13"
	RuntimeEnvironmentZeppelinFlink20 RuntimeEnvironment = "ZEPPELIN-FLINK-2_0"
	RuntimeEnvironmentFlink115        RuntimeEnvironment = "FLINK-1_15"
	RuntimeEnvironmentZeppelinFlink30 RuntimeEnvironment = "ZEPPELIN-FLINK-3_0"
	RuntimeEnvironmentFlink118        RuntimeEnvironment = "FLINK-1_18"
)

// Values returns all known values for RuntimeEnvironment. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RuntimeEnvironment) Values() []RuntimeEnvironment {
	return []RuntimeEnvironment{
		"SQL-1_0",
		"FLINK-1_6",
		"FLINK-1_8",
		"ZEPPELIN-FLINK-1_0",
		"FLINK-1_11",
		"FLINK-1_13",
		"ZEPPELIN-FLINK-2_0",
		"FLINK-1_15",
		"ZEPPELIN-FLINK-3_0",
		"FLINK-1_18",
	}
}

type SnapshotStatus string

// Enum values for SnapshotStatus
const (
	SnapshotStatusCreating SnapshotStatus = "CREATING"
	SnapshotStatusReady    SnapshotStatus = "READY"
	SnapshotStatusDeleting SnapshotStatus = "DELETING"
	SnapshotStatusFailed   SnapshotStatus = "FAILED"
)

// Values returns all known values for SnapshotStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SnapshotStatus) Values() []SnapshotStatus {
	return []SnapshotStatus{
		"CREATING",
		"READY",
		"DELETING",
		"FAILED",
	}
}

type UrlType string

// Enum values for UrlType
const (
	UrlTypeFlinkDashboardUrl UrlType = "FLINK_DASHBOARD_URL"
	UrlTypeZeppelinUiUrl     UrlType = "ZEPPELIN_UI_URL"
)

// Values returns all known values for UrlType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UrlType) Values() []UrlType {
	return []UrlType{
		"FLINK_DASHBOARD_URL",
		"ZEPPELIN_UI_URL",
	}
}

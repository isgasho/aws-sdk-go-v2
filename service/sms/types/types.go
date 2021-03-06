// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Information about the application.
type AppSummary struct {

	// Unique ID of the application.
	AppId *string

	// Time of creation of this application.
	CreationTime *time.Time

	// Description of the application.
	Description *string

	// Timestamp of the application's creation.
	LastModified *time.Time

	// Timestamp of the application's most recent successful replication.
	LatestReplicationTime *time.Time

	// Details about the latest launch of the application.
	LaunchDetails *LaunchDetails

	// Launch status of the application.
	LaunchStatus AppLaunchStatus

	// A message related to the launch status of the application.
	LaunchStatusMessage *string

	// Name of the application.
	Name *string

	// Replication status of the application.
	ReplicationStatus AppReplicationStatus

	// A message related to the replication status of the application.
	ReplicationStatusMessage *string

	// Name of the service role in the customer's account used by AWS SMS.
	RoleName *string

	// Status of the application.
	Status AppStatus

	// A message related to the status of the application
	StatusMessage *string

	// Number of server groups present in the application.
	TotalServerGroups *int32

	// Number of servers present in the application.
	TotalServers *int32
}

// Represents a connector.
type Connector struct {

	// The time the connector was associated.
	AssociatedOn *time.Time

	// The capabilities of the connector.
	CapabilityList []ConnectorCapability

	// The identifier of the connector.
	ConnectorId *string

	// The IP address of the connector.
	IpAddress *string

	// The MAC address of the connector.
	MacAddress *string

	// The status of the connector.
	Status ConnectorStatus

	// The connector version.
	Version *string

	// The identifier of the VM manager.
	VmManagerId *string

	// The name of the VM manager.
	VmManagerName *string

	// The VM management product.
	VmManagerType VmManagerType
}

// Details about the latest launch of an application.
type LaunchDetails struct {

	// Latest time this application was launched successfully.
	LatestLaunchTime *time.Time

	// Identifier of the latest stack launched for this application.
	StackId *string

	// Name of the latest stack launched for this application.
	StackName *string
}

// Represents a replication job.
type ReplicationJob struct {

	// The description of the replication job.
	Description *string

	// Whether the replication job should produce encrypted AMIs or not. See also
	// KmsKeyId below.
	Encrypted *bool

	// The time between consecutive replication runs, in hours.
	Frequency *int32

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of the
	// following:
	//
	//     * KMS key ID
	//
	//     * KMS key alias
	//
	//     * ARN referring to KMS
	// key ID
	//
	//     * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key
	// id is not specified, the customer's default KMS key for EBS is used.
	KmsKeyId *string

	// The ID of the latest Amazon Machine Image (AMI).
	LatestAmiId *string

	// The license type to be used for the AMI created by a successful replication run.
	LicenseType LicenseType

	// The start time of the next replication run.
	NextReplicationRunStartTime *time.Time

	// Number of recent AMIs to keep in the customer's account for a replication job.
	// By default the value is set to zero, meaning that all AMIs are kept.
	NumberOfRecentAmisToKeep *int32

	// The identifier of the replication job.
	ReplicationJobId *string

	// Information about the replication runs.
	ReplicationRunList []*ReplicationRun

	// The name of the IAM role to be used by the Server Migration Service.
	RoleName *string

	//
	RunOnce *bool

	// The seed replication time.
	SeedReplicationTime *time.Time

	// The identifier of the server.
	ServerId *string

	// The type of server.
	ServerType ServerType

	// The state of the replication job.
	State ReplicationJobState

	// The description of the current status of the replication job.
	StatusMessage *string

	// Information about the VM server.
	VmServer *VmServer
}

// Represents a replication run.
type ReplicationRun struct {

	// The identifier of the Amazon Machine Image (AMI) from the replication run.
	AmiId *string

	// The completion time of the last replication run.
	CompletedTime *time.Time

	// The description of the replication run.
	Description *string

	// Whether the replication run should produce encrypted AMI or not. See also
	// KmsKeyId below.
	Encrypted *bool

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of the
	// following:
	//
	//     * KMS key ID
	//
	//     * KMS key alias
	//
	//     * ARN referring to KMS
	// key ID
	//
	//     * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key
	// id is not specified, the customer's default KMS key for EBS is used.
	KmsKeyId *string

	// The identifier of the replication run.
	ReplicationRunId *string

	// The start time of the next replication run.
	ScheduledStartTime *time.Time

	// Details of the current stage of the replication run.
	StageDetails *ReplicationRunStageDetails

	// The state of the replication run.
	State ReplicationRunState

	// The description of the current status of the replication job.
	StatusMessage *string

	// The type of replication run.
	Type ReplicationRunType
}

// Details of the current stage of a replication run.
type ReplicationRunStageDetails struct {

	// String describing the current stage of a replication run.
	Stage *string

	// String describing the progress of the current stage of a replication run.
	StageProgress *string
}

// Location of the Amazon S3 object in the customer's account.
type S3Location struct {

	// Amazon S3 bucket name.
	Bucket *string

	// Amazon S3 bucket key.
	Key *string
}

// Represents a server.
type Server struct {

	// The identifier of the replication job.
	ReplicationJobId *string

	// Indicates whether the replication job is deleted or failed.
	ReplicationJobTerminated *bool

	// The identifier of the server.
	ServerId *string

	// The type of server.
	ServerType ServerType

	// Information about the VM server.
	VmServer *VmServer
}

// A logical grouping of servers.
type ServerGroup struct {

	// Name of a server group.
	Name *string

	// Identifier of a server group.
	ServerGroupId *string

	// List of servers belonging to a server group.
	ServerList []*Server
}

// Launch configuration for a server group.
type ServerGroupLaunchConfiguration struct {

	// Launch order of servers in the server group.
	LaunchOrder *int32

	// Identifier of the server group the launch configuration is associated with.
	ServerGroupId *string

	// Launch configuration for servers in the server group.
	ServerLaunchConfigurations []*ServerLaunchConfiguration
}

// Replication configuration for a server group.
type ServerGroupReplicationConfiguration struct {

	// Identifier of the server group this replication configuration is associated
	// with.
	ServerGroupId *string

	// Replication configuration for servers in the server group.
	ServerReplicationConfigurations []*ServerReplicationConfiguration
}

// Launch configuration for a server.
type ServerLaunchConfiguration struct {

	// If true, a publicly accessible IP address is created when launching the server.
	AssociatePublicIpAddress *bool

	// Name of the EC2 SSH Key to be used for connecting to the launched server.
	Ec2KeyName *string

	// Instance type to be used for launching the server.
	InstanceType *string

	// Logical ID of the server in the Amazon CloudFormation template.
	LogicalId *string

	// Identifier of the security group that applies to the launched server.
	SecurityGroup *string

	// Identifier of the server the launch configuration is associated with.
	Server *Server

	// Identifier of the subnet the server should be launched into.
	Subnet *string

	// Location of the user-data script to be executed when launching the server.
	UserData *UserData

	// Identifier of the VPC the server should be launched into.
	Vpc *string
}

// Replication configuration of a server.
type ServerReplicationConfiguration struct {

	// Identifier of the server this replication configuration is associated with.
	Server *Server

	// Parameters for replicating the server.
	ServerReplicationParameters *ServerReplicationParameters
}

// Replication parameters for replicating a server.
type ServerReplicationParameters struct {

	// When true, the replication job produces encrypted AMIs. See also KmsKeyId below.
	Encrypted *bool

	// Frequency of creating replication jobs for the server.
	Frequency *int32

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of the
	// following:
	//
	//     * KMS key ID
	//
	//     * KMS key alias
	//
	//     * ARN referring to KMS
	// key ID
	//
	//     * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key
	// id is not specified, the customer's default KMS key for EBS is used.
	KmsKeyId *string

	// License type for creating a replication job for the server.
	LicenseType LicenseType

	// Number of recent AMIs to keep when creating a replication job for this server.
	NumberOfRecentAmisToKeep *int32

	//
	RunOnce *bool

	// Seed time for creating a replication job for the server.
	SeedTime *time.Time
}

// A label that can be assigned to an application.
type Tag struct {

	// Tag key.
	Key *string

	// Tag value.
	Value *string
}

// A script that runs on first launch of an Amazon EC2 instance. Used for
// configuring the server during launch.
type UserData struct {

	// Amazon S3 location of the user-data script.
	S3Location *S3Location
}

// Represents a VM server.
type VmServer struct {

	// The name of the VM manager.
	VmManagerName *string

	// The type of VM management product.
	VmManagerType VmManagerType

	// The name of the VM.
	VmName *string

	// The VM folder path in the vCenter Server virtual machine inventory tree.
	VmPath *string

	// Information about the VM server location.
	VmServerAddress *VmServerAddress
}

// Represents a VM server location.
type VmServerAddress struct {

	// The identifier of the VM.
	VmId *string

	// The identifier of the VM manager.
	VmManagerId *string
}

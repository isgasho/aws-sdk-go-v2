// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBClusterToPointInTimeMessage
type RestoreDBClusterToPointInTimeInput struct {
	_ struct{} `type:"structure"`

	// The target backtrack window, in seconds. To disable backtracking, set this
	// value to 0.
	//
	// Default: 0
	//
	// Constraints:
	//
	//    * If specified, this value must be set to a number from 0 to 259,200 (72
	//    hours).
	BacktrackWindow *int64 `type:"long"`

	// A value that indicates whether to copy all tags from the restored DB cluster
	// to snapshots of the restored DB cluster. The default is not to copy them.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The name of the new DB cluster to be created.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to associate with this DB cluster.
	// If this argument is omitted, the default DB cluster parameter group for the
	// specified engine is used.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DB cluster parameter
	//    group.
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	DBClusterParameterGroupName *string `type:"string"`

	// The DB subnet group name to use for the new DB cluster.
	//
	// Constraints: If supplied, must match the name of an existing DBSubnetGroup.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool `type:"boolean"`

	// The list of logs that the restored DB cluster is to export to CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The AWS KMS key identifier to use when restoring an encrypted DB cluster
	// from an encrypted DB cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are restoring a DB cluster with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB cluster, then you can use
	// the KMS key alias instead of the ARN for the KMS encryption key.
	//
	// You can restore to a new DB cluster and encrypt the new DB cluster with a
	// KMS key that is different than the KMS key used to encrypt the source DB
	// cluster. The new DB cluster is encrypted with the KMS key identified by the
	// KmsKeyId parameter.
	//
	// If you don't specify a value for the KmsKeyId parameter, then the following
	// occurs:
	//
	//    * If the DB cluster is encrypted, then the restored DB cluster is encrypted
	//    using the KMS key that was used to encrypt the source DB cluster.
	//
	//    * If the DB cluster is not encrypted, then the restored DB cluster is
	//    not encrypted.
	//
	// If DBClusterIdentifier refers to a DB cluster that is not encrypted, then
	// the restore request is rejected.
	KmsKeyId *string `type:"string"`

	// The name of the option group for the new DB cluster.
	OptionGroupName *string `type:"string"`

	// The port number on which the new DB cluster accepts connections.
	//
	// Constraints: A value from 1150-65535.
	//
	// Default: The default port for the engine.
	Port *int64 `type:"integer"`

	// The date and time to restore the DB cluster to.
	//
	// Valid Values: Value must be a time in Universal Coordinated Time (UTC) format
	//
	// Constraints:
	//
	//    * Must be before the latest restorable time for the DB instance
	//
	//    * Must be specified if UseLatestRestorableTime parameter is not provided
	//
	//    * Can't be specified if the UseLatestRestorableTime parameter is enabled
	//
	//    * Can't be specified if the RestoreType parameter is copy-on-write
	//
	// Example: 2015-03-07T23:45:00Z
	RestoreToTime *time.Time `type:"timestamp"`

	// The type of restore to be performed. You can specify one of the following
	// values:
	//
	//    * full-copy - The new DB cluster is restored as a full copy of the source
	//    DB cluster.
	//
	//    * copy-on-write - The new DB cluster is restored as a clone of the source
	//    DB cluster.
	//
	// Constraints: You can't specify copy-on-write if the engine version of the
	// source DB cluster is earlier than 1.11.
	//
	// If you don't specify a RestoreType value, then the new DB cluster is restored
	// as a full copy of the source DB cluster.
	RestoreType *string `type:"string"`

	// The identifier of the source DB cluster from which to restore.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBCluster.
	//
	// SourceDBClusterIdentifier is a required field
	SourceDBClusterIdentifier *string `type:"string" required:"true"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A value that indicates whether to restore the DB cluster to the latest restorable
	// backup time. By default, the DB cluster is not restored to the latest restorable
	// backup time.
	//
	// Constraints: Can't be specified if RestoreToTime parameter is provided.
	UseLatestRestorableTime *bool `type:"boolean"`

	// A list of VPC security groups that the new DB cluster belongs to.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBClusterToPointInTimeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBClusterToPointInTimeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBClusterToPointInTimeInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.SourceDBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceDBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBClusterToPointInTimeResult
type RestoreDBClusterToPointInTimeOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Aurora DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters, StopDBCluster,
	// and StartDBCluster actions.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s RestoreDBClusterToPointInTimeOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreDBClusterToPointInTime = "RestoreDBClusterToPointInTime"

// RestoreDBClusterToPointInTimeRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Restores a DB cluster to an arbitrary point in time. Users can restore to
// any point in time before LatestRestorableTime for up to BackupRetentionPeriod
// days. The target DB cluster is created from the source DB cluster with the
// same configuration as the original DB cluster, except that the new DB cluster
// is created with the default DB security group.
//
// This action only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstance action to create DB instances
// for the restored DB cluster, specifying the identifier of the restored DB
// cluster in DBClusterIdentifier. You can create DB instances only after the
// RestoreDBClusterToPointInTime action has completed and the DB cluster is
// available.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using RestoreDBClusterToPointInTimeRequest.
//    req := client.RestoreDBClusterToPointInTimeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBClusterToPointInTime
func (c *Client) RestoreDBClusterToPointInTimeRequest(input *RestoreDBClusterToPointInTimeInput) RestoreDBClusterToPointInTimeRequest {
	op := &aws.Operation{
		Name:       opRestoreDBClusterToPointInTime,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreDBClusterToPointInTimeInput{}
	}

	req := c.newRequest(op, input, &RestoreDBClusterToPointInTimeOutput{})
	return RestoreDBClusterToPointInTimeRequest{Request: req, Input: input, Copy: c.RestoreDBClusterToPointInTimeRequest}
}

// RestoreDBClusterToPointInTimeRequest is the request type for the
// RestoreDBClusterToPointInTime API operation.
type RestoreDBClusterToPointInTimeRequest struct {
	*aws.Request
	Input *RestoreDBClusterToPointInTimeInput
	Copy  func(*RestoreDBClusterToPointInTimeInput) RestoreDBClusterToPointInTimeRequest
}

// Send marshals and sends the RestoreDBClusterToPointInTime API request.
func (r RestoreDBClusterToPointInTimeRequest) Send(ctx context.Context) (*RestoreDBClusterToPointInTimeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBClusterToPointInTimeResponse{
		RestoreDBClusterToPointInTimeOutput: r.Request.Data.(*RestoreDBClusterToPointInTimeOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBClusterToPointInTimeResponse is the response type for the
// RestoreDBClusterToPointInTime API operation.
type RestoreDBClusterToPointInTimeResponse struct {
	*RestoreDBClusterToPointInTimeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBClusterToPointInTime request.
func (r *RestoreDBClusterToPointInTimeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupPlanInput
type DeleteBackupPlanInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a backup plan.
	//
	// BackupPlanId is a required field
	BackupPlanId *string `location:"uri" locationName:"backupPlanId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackupPlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackupPlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBackupPlanInput"}

	if s.BackupPlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackupPlanInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupPlanOutput
type DeleteBackupPlanOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example, arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string `type:"string"`

	// Uniquely identifies a backup plan.
	BackupPlanId *string `type:"string"`

	// The date and time a backup plan is deleted, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	DeletionDate *time.Time `type:"timestamp"`

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version Ids cannot be edited.
	VersionId *string `type:"string"`
}

// String returns the string representation
func (s DeleteBackupPlanOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackupPlanOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupPlanArn != nil {
		v := *s.BackupPlanArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeletionDate != nil {
		v := *s.DeletionDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeletionDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteBackupPlan = "DeleteBackupPlan"

// DeleteBackupPlanRequest returns a request value for making API operation for
// AWS Backup.
//
// Deletes a backup plan. A backup plan can only be deleted after all associated
// selections of resources have been deleted. Deleting a backup plan deletes
// the current version of a backup plan. Previous versions, if any, will still
// exist.
//
//    // Example sending a request using DeleteBackupPlanRequest.
//    req := client.DeleteBackupPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupPlan
func (c *Client) DeleteBackupPlanRequest(input *DeleteBackupPlanInput) DeleteBackupPlanRequest {
	op := &aws.Operation{
		Name:       opDeleteBackupPlan,
		HTTPMethod: "DELETE",
		HTTPPath:   "/backup/plans/{backupPlanId}",
	}

	if input == nil {
		input = &DeleteBackupPlanInput{}
	}

	req := c.newRequest(op, input, &DeleteBackupPlanOutput{})
	return DeleteBackupPlanRequest{Request: req, Input: input, Copy: c.DeleteBackupPlanRequest}
}

// DeleteBackupPlanRequest is the request type for the
// DeleteBackupPlan API operation.
type DeleteBackupPlanRequest struct {
	*aws.Request
	Input *DeleteBackupPlanInput
	Copy  func(*DeleteBackupPlanInput) DeleteBackupPlanRequest
}

// Send marshals and sends the DeleteBackupPlan API request.
func (r DeleteBackupPlanRequest) Send(ctx context.Context) (*DeleteBackupPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupPlanResponse{
		DeleteBackupPlanOutput: r.Request.Data.(*DeleteBackupPlanOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupPlanResponse is the response type for the
// DeleteBackupPlan API operation.
type DeleteBackupPlanResponse struct {
	*DeleteBackupPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackupPlan request.
func (r *DeleteBackupPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

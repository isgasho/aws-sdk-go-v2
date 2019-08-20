// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateScheduledAuditInput struct {
	_ struct{} `type:"structure"`

	// The day of the month on which the scheduled audit takes place. Can be "1"
	// through "31" or "LAST". This field is required if the "frequency" parameter
	// is set to "MONTHLY". If days 29-31 are specified, and the month does not
	// have that many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string `locationName:"dayOfMonth" type:"string"`

	// The day of the week on which the scheduled audit takes place. Can be one
	// of "SUN", "MON", "TUE", "WED", "THU", "FRI" or "SAT". This field is required
	// if the "frequency" parameter is set to "WEEKLY" or "BIWEEKLY".
	DayOfWeek DayOfWeek `locationName:"dayOfWeek" type:"string" enum:"true"`

	// How often the scheduled audit takes place. Can be one of "DAILY", "WEEKLY",
	// "BIWEEKLY" or "MONTHLY". The actual start time of each audit is determined
	// by the system.
	//
	// Frequency is a required field
	Frequency AuditFrequency `locationName:"frequency" type:"string" required:"true" enum:"true"`

	// The name you want to give to the scheduled audit. (Max. 128 chars)
	//
	// ScheduledAuditName is a required field
	ScheduledAuditName *string `location:"uri" locationName:"scheduledAuditName" min:"1" type:"string" required:"true"`

	// Metadata which can be used to manage the scheduled audit.
	Tags []Tag `locationName:"tags" type:"list"`

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list
	// of all checks including those that are enabled or UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	//
	// TargetCheckNames is a required field
	TargetCheckNames []string `locationName:"targetCheckNames" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateScheduledAuditInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateScheduledAuditInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateScheduledAuditInput"}
	if len(s.Frequency) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Frequency"))
	}

	if s.ScheduledAuditName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledAuditName"))
	}
	if s.ScheduledAuditName != nil && len(*s.ScheduledAuditName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScheduledAuditName", 1))
	}

	if s.TargetCheckNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetCheckNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateScheduledAuditInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DayOfMonth != nil {
		v := *s.DayOfMonth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dayOfMonth", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DayOfWeek) > 0 {
		v := s.DayOfWeek

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dayOfWeek", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Frequency) > 0 {
		v := s.Frequency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "frequency", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TargetCheckNames != nil {
		v := s.TargetCheckNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targetCheckNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ScheduledAuditName != nil {
		v := *s.ScheduledAuditName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "scheduledAuditName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateScheduledAuditOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string `locationName:"scheduledAuditArn" type:"string"`
}

// String returns the string representation
func (s CreateScheduledAuditOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateScheduledAuditOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ScheduledAuditArn != nil {
		v := *s.ScheduledAuditArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "scheduledAuditArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateScheduledAudit = "CreateScheduledAudit"

// CreateScheduledAuditRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a scheduled audit that is run at a specified time interval.
//
//    // Example sending a request using CreateScheduledAuditRequest.
//    req := client.CreateScheduledAuditRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateScheduledAuditRequest(input *CreateScheduledAuditInput) CreateScheduledAuditRequest {
	op := &aws.Operation{
		Name:       opCreateScheduledAudit,
		HTTPMethod: "POST",
		HTTPPath:   "/audit/scheduledaudits/{scheduledAuditName}",
	}

	if input == nil {
		input = &CreateScheduledAuditInput{}
	}

	req := c.newRequest(op, input, &CreateScheduledAuditOutput{})
	return CreateScheduledAuditRequest{Request: req, Input: input, Copy: c.CreateScheduledAuditRequest}
}

// CreateScheduledAuditRequest is the request type for the
// CreateScheduledAudit API operation.
type CreateScheduledAuditRequest struct {
	*aws.Request
	Input *CreateScheduledAuditInput
	Copy  func(*CreateScheduledAuditInput) CreateScheduledAuditRequest
}

// Send marshals and sends the CreateScheduledAudit API request.
func (r CreateScheduledAuditRequest) Send(ctx context.Context) (*CreateScheduledAuditResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateScheduledAuditResponse{
		CreateScheduledAuditOutput: r.Request.Data.(*CreateScheduledAuditOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateScheduledAuditResponse is the response type for the
// CreateScheduledAudit API operation.
type CreateScheduledAuditResponse struct {
	*CreateScheduledAuditOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateScheduledAudit request.
func (r *CreateScheduledAuditResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

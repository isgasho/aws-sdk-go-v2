// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetMinuteUsageRequest
type GetMinuteUsageInput struct {
	_ struct{} `type:"structure"`

	// The month being requested, with a value of 1-12.
	//
	// Month is a required field
	Month *int64 `locationName:"month" type:"integer" required:"true"`

	// The year being requested, in the format of YYYY.
	//
	// Year is a required field
	Year *int64 `locationName:"year" type:"integer" required:"true"`
}

// String returns the string representation
func (s GetMinuteUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMinuteUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMinuteUsageInput"}

	if s.Month == nil {
		invalidParams.Add(aws.NewErrParamRequired("Month"))
	}

	if s.Year == nil {
		invalidParams.Add(aws.NewErrParamRequired("Year"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMinuteUsageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Month != nil {
		v := *s.Month

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "month", protocol.Int64Value(v), metadata)
	}
	if s.Year != nil {
		v := *s.Year

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "year", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetMinuteUsageResponse
type GetMinuteUsageOutput struct {
	_ struct{} `type:"structure"`

	// Estimated number of minutes remaining for an account, specific to the month
	// being requested.
	EstimatedMinutesRemaining *int64 `locationName:"estimatedMinutesRemaining" type:"integer"`

	// Returns whether or not an account has signed up for the reserved minutes
	// pricing plan, specific to the month being requested.
	IsReservedMinutesCustomer *bool `locationName:"isReservedMinutesCustomer" type:"boolean"`

	// Total number of reserved minutes allocated, specific to the month being requested.
	TotalReservedMinuteAllocation *int64 `locationName:"totalReservedMinuteAllocation" type:"integer"`

	// Total scheduled minutes for an account, specific to the month being requested.
	TotalScheduledMinutes *int64 `locationName:"totalScheduledMinutes" type:"integer"`

	// Upcoming minutes scheduled for an account, specific to the month being requested.
	UpcomingMinutesScheduled *int64 `locationName:"upcomingMinutesScheduled" type:"integer"`
}

// String returns the string representation
func (s GetMinuteUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMinuteUsageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EstimatedMinutesRemaining != nil {
		v := *s.EstimatedMinutesRemaining

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "estimatedMinutesRemaining", protocol.Int64Value(v), metadata)
	}
	if s.IsReservedMinutesCustomer != nil {
		v := *s.IsReservedMinutesCustomer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "isReservedMinutesCustomer", protocol.BoolValue(v), metadata)
	}
	if s.TotalReservedMinuteAllocation != nil {
		v := *s.TotalReservedMinuteAllocation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "totalReservedMinuteAllocation", protocol.Int64Value(v), metadata)
	}
	if s.TotalScheduledMinutes != nil {
		v := *s.TotalScheduledMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "totalScheduledMinutes", protocol.Int64Value(v), metadata)
	}
	if s.UpcomingMinutesScheduled != nil {
		v := *s.UpcomingMinutesScheduled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "upcomingMinutesScheduled", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opGetMinuteUsage = "GetMinuteUsage"

// GetMinuteUsageRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns the number of minutes used by account.
//
//    // Example sending a request using GetMinuteUsageRequest.
//    req := client.GetMinuteUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetMinuteUsage
func (c *Client) GetMinuteUsageRequest(input *GetMinuteUsageInput) GetMinuteUsageRequest {
	op := &aws.Operation{
		Name:       opGetMinuteUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/minute-usage",
	}

	if input == nil {
		input = &GetMinuteUsageInput{}
	}

	req := c.newRequest(op, input, &GetMinuteUsageOutput{})
	return GetMinuteUsageRequest{Request: req, Input: input, Copy: c.GetMinuteUsageRequest}
}

// GetMinuteUsageRequest is the request type for the
// GetMinuteUsage API operation.
type GetMinuteUsageRequest struct {
	*aws.Request
	Input *GetMinuteUsageInput
	Copy  func(*GetMinuteUsageInput) GetMinuteUsageRequest
}

// Send marshals and sends the GetMinuteUsage API request.
func (r GetMinuteUsageRequest) Send(ctx context.Context) (*GetMinuteUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMinuteUsageResponse{
		GetMinuteUsageOutput: r.Request.Data.(*GetMinuteUsageOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMinuteUsageResponse is the response type for the
// GetMinuteUsage API operation.
type GetMinuteUsageResponse struct {
	*GetMinuteUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMinuteUsage request.
func (r *GetMinuteUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

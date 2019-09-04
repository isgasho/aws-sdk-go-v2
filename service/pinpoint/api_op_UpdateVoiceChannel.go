// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateVoiceChannelRequest
type UpdateVoiceChannelInput struct {
	_ struct{} `type:"structure" payload:"VoiceChannelRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the status and settings of the voice channel for an application.
	//
	// VoiceChannelRequest is a required field
	VoiceChannelRequest *VoiceChannelRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateVoiceChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVoiceChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVoiceChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.VoiceChannelRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceChannelRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateVoiceChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VoiceChannelRequest != nil {
		v := s.VoiceChannelRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "VoiceChannelRequest", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateVoiceChannelResponse
type UpdateVoiceChannelOutput struct {
	_ struct{} `type:"structure" payload:"VoiceChannelResponse"`

	// Provides information about the status and settings of the voice channel for
	// an application.
	//
	// VoiceChannelResponse is a required field
	VoiceChannelResponse *VoiceChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateVoiceChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateVoiceChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VoiceChannelResponse != nil {
		v := s.VoiceChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "VoiceChannelResponse", v, metadata)
	}
	return nil
}

const opUpdateVoiceChannel = "UpdateVoiceChannel"

// UpdateVoiceChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the status and settings of the voice channel for an application.
//
//    // Example sending a request using UpdateVoiceChannelRequest.
//    req := client.UpdateVoiceChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateVoiceChannel
func (c *Client) UpdateVoiceChannelRequest(input *UpdateVoiceChannelInput) UpdateVoiceChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateVoiceChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/channels/voice",
	}

	if input == nil {
		input = &UpdateVoiceChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateVoiceChannelOutput{})
	return UpdateVoiceChannelRequest{Request: req, Input: input, Copy: c.UpdateVoiceChannelRequest}
}

// UpdateVoiceChannelRequest is the request type for the
// UpdateVoiceChannel API operation.
type UpdateVoiceChannelRequest struct {
	*aws.Request
	Input *UpdateVoiceChannelInput
	Copy  func(*UpdateVoiceChannelInput) UpdateVoiceChannelRequest
}

// Send marshals and sends the UpdateVoiceChannel API request.
func (r UpdateVoiceChannelRequest) Send(ctx context.Context) (*UpdateVoiceChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVoiceChannelResponse{
		UpdateVoiceChannelOutput: r.Request.Data.(*UpdateVoiceChannelOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVoiceChannelResponse is the response type for the
// UpdateVoiceChannel API operation.
type UpdateVoiceChannelResponse struct {
	*UpdateVoiceChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVoiceChannel request.
func (r *UpdateVoiceChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
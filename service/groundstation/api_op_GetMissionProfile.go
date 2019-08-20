// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetMissionProfileRequest
type GetMissionProfileInput struct {
	_ struct{} `type:"structure"`

	// UUID of a mission profile.
	//
	// MissionProfileId is a required field
	MissionProfileId *string `location:"uri" locationName:"missionProfileId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMissionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMissionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMissionProfileInput"}

	if s.MissionProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MissionProfileId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMissionProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MissionProfileId != nil {
		v := *s.MissionProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "missionProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetMissionProfileResponse
type GetMissionProfileOutput struct {
	_ struct{} `type:"structure"`

	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	ContactPostPassDurationSeconds *int64 `locationName:"contactPostPassDurationSeconds" min:"1" type:"integer"`

	// Amount of time prior to contact start you’d like to receive a CloudWatch
	// event indicating an upcoming pass.
	ContactPrePassDurationSeconds *int64 `locationName:"contactPrePassDurationSeconds" min:"1" type:"integer"`

	// A list of lists of ARNs. Each list of ARNs is an edge, with a from Config
	// and a to Config.
	DataflowEdges [][]string `locationName:"dataflowEdges" type:"list"`

	// Smallest amount of time in seconds that you’d like to see for an available
	// contact. AWS Ground Station will not present you with contacts shorter than
	// this duration.
	MinimumViableContactDurationSeconds *int64 `locationName:"minimumViableContactDurationSeconds" min:"1" type:"integer"`

	// ARN of a mission profile.
	MissionProfileArn *string `locationName:"missionProfileArn" type:"string"`

	// ID of a mission profile.
	MissionProfileId *string `locationName:"missionProfileId" type:"string"`

	// Name of a mission profile.
	Name *string `locationName:"name" type:"string"`

	// Region of a mission profile.
	Region *string `locationName:"region" type:"string"`

	// Tags assigned to a mission profile.
	Tags map[string]string `locationName:"tags" type:"map"`

	// ARN of a tracking Config.
	TrackingConfigArn *string `locationName:"trackingConfigArn" type:"string"`
}

// String returns the string representation
func (s GetMissionProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMissionProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContactPostPassDurationSeconds != nil {
		v := *s.ContactPostPassDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contactPostPassDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.ContactPrePassDurationSeconds != nil {
		v := *s.ContactPrePassDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contactPrePassDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.DataflowEdges != nil {
		v := s.DataflowEdges

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dataflowEdges", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls1 := ls0.List()
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v2)})
			}
			ls1.End()
		}
		ls0.End()

	}
	if s.MinimumViableContactDurationSeconds != nil {
		v := *s.MinimumViableContactDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "minimumViableContactDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.MissionProfileArn != nil {
		v := *s.MissionProfileArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "missionProfileArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MissionProfileId != nil {
		v := *s.MissionProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "missionProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Region != nil {
		v := *s.Region

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "region", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TrackingConfigArn != nil {
		v := *s.TrackingConfigArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "trackingConfigArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetMissionProfile = "GetMissionProfile"

// GetMissionProfileRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns a mission profile.
//
//    // Example sending a request using GetMissionProfileRequest.
//    req := client.GetMissionProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetMissionProfile
func (c *Client) GetMissionProfileRequest(input *GetMissionProfileInput) GetMissionProfileRequest {
	op := &aws.Operation{
		Name:       opGetMissionProfile,
		HTTPMethod: "GET",
		HTTPPath:   "/missionprofile/{missionProfileId}",
	}

	if input == nil {
		input = &GetMissionProfileInput{}
	}

	req := c.newRequest(op, input, &GetMissionProfileOutput{})
	return GetMissionProfileRequest{Request: req, Input: input, Copy: c.GetMissionProfileRequest}
}

// GetMissionProfileRequest is the request type for the
// GetMissionProfile API operation.
type GetMissionProfileRequest struct {
	*aws.Request
	Input *GetMissionProfileInput
	Copy  func(*GetMissionProfileInput) GetMissionProfileRequest
}

// Send marshals and sends the GetMissionProfile API request.
func (r GetMissionProfileRequest) Send(ctx context.Context) (*GetMissionProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMissionProfileResponse{
		GetMissionProfileOutput: r.Request.Data.(*GetMissionProfileOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMissionProfileResponse is the response type for the
// GetMissionProfile API operation.
type GetMissionProfileResponse struct {
	*GetMissionProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMissionProfile request.
func (r *GetMissionProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/UpdateSimulationApplicationRequest
type UpdateSimulationApplicationInput struct {
	_ struct{} `type:"structure"`

	// The application information for the simulation application.
	//
	// Application is a required field
	Application *string `locationName:"application" min:"1" type:"string" required:"true"`

	// The revision id for the robot application.
	CurrentRevisionId *string `locationName:"currentRevisionId" min:"1" type:"string"`

	// The rendering engine for the simulation application.
	//
	// RenderingEngine is a required field
	RenderingEngine *RenderingEngine `locationName:"renderingEngine" type:"structure" required:"true"`

	// Information about the robot software suite.
	//
	// RobotSoftwareSuite is a required field
	RobotSoftwareSuite *RobotSoftwareSuite `locationName:"robotSoftwareSuite" type:"structure" required:"true"`

	// The simulation software suite used by the simulation application.
	//
	// SimulationSoftwareSuite is a required field
	SimulationSoftwareSuite *SimulationSoftwareSuite `locationName:"simulationSoftwareSuite" type:"structure" required:"true"`

	// The sources of the simulation application.
	//
	// Sources is a required field
	Sources []SourceConfig `locationName:"sources" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateSimulationApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSimulationApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSimulationApplicationInput"}

	if s.Application == nil {
		invalidParams.Add(aws.NewErrParamRequired("Application"))
	}
	if s.Application != nil && len(*s.Application) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Application", 1))
	}
	if s.CurrentRevisionId != nil && len(*s.CurrentRevisionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentRevisionId", 1))
	}

	if s.RenderingEngine == nil {
		invalidParams.Add(aws.NewErrParamRequired("RenderingEngine"))
	}

	if s.RobotSoftwareSuite == nil {
		invalidParams.Add(aws.NewErrParamRequired("RobotSoftwareSuite"))
	}

	if s.SimulationSoftwareSuite == nil {
		invalidParams.Add(aws.NewErrParamRequired("SimulationSoftwareSuite"))
	}

	if s.Sources == nil {
		invalidParams.Add(aws.NewErrParamRequired("Sources"))
	}
	if s.Sources != nil {
		for i, v := range s.Sources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Sources", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSimulationApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Application != nil {
		v := *s.Application

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "application", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CurrentRevisionId != nil {
		v := *s.CurrentRevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currentRevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RenderingEngine != nil {
		v := s.RenderingEngine

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "renderingEngine", v, metadata)
	}
	if s.RobotSoftwareSuite != nil {
		v := s.RobotSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "robotSoftwareSuite", v, metadata)
	}
	if s.SimulationSoftwareSuite != nil {
		v := s.SimulationSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "simulationSoftwareSuite", v, metadata)
	}
	if s.Sources != nil {
		v := s.Sources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "sources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/UpdateSimulationApplicationResponse
type UpdateSimulationApplicationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the updated simulation application.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the simulation application
	// was last updated.
	LastUpdatedAt *time.Time `locationName:"lastUpdatedAt" type:"timestamp"`

	// The name of the simulation application.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The rendering engine for the simulation application.
	RenderingEngine *RenderingEngine `locationName:"renderingEngine" type:"structure"`

	// The revision id of the simulation application.
	RevisionId *string `locationName:"revisionId" min:"1" type:"string"`

	// Information about the robot software suite.
	RobotSoftwareSuite *RobotSoftwareSuite `locationName:"robotSoftwareSuite" type:"structure"`

	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *SimulationSoftwareSuite `locationName:"simulationSoftwareSuite" type:"structure"`

	// The sources of the simulation application.
	Sources []Source `locationName:"sources" type:"list"`

	// The version of the robot application.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateSimulationApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSimulationApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RenderingEngine != nil {
		v := s.RenderingEngine

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "renderingEngine", v, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RobotSoftwareSuite != nil {
		v := s.RobotSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "robotSoftwareSuite", v, metadata)
	}
	if s.SimulationSoftwareSuite != nil {
		v := s.SimulationSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "simulationSoftwareSuite", v, metadata)
	}
	if s.Sources != nil {
		v := s.Sources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "sources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateSimulationApplication = "UpdateSimulationApplication"

// UpdateSimulationApplicationRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Updates a simulation application.
//
//    // Example sending a request using UpdateSimulationApplicationRequest.
//    req := client.UpdateSimulationApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/UpdateSimulationApplication
func (c *Client) UpdateSimulationApplicationRequest(input *UpdateSimulationApplicationInput) UpdateSimulationApplicationRequest {
	op := &aws.Operation{
		Name:       opUpdateSimulationApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/updateSimulationApplication",
	}

	if input == nil {
		input = &UpdateSimulationApplicationInput{}
	}

	req := c.newRequest(op, input, &UpdateSimulationApplicationOutput{})
	return UpdateSimulationApplicationRequest{Request: req, Input: input, Copy: c.UpdateSimulationApplicationRequest}
}

// UpdateSimulationApplicationRequest is the request type for the
// UpdateSimulationApplication API operation.
type UpdateSimulationApplicationRequest struct {
	*aws.Request
	Input *UpdateSimulationApplicationInput
	Copy  func(*UpdateSimulationApplicationInput) UpdateSimulationApplicationRequest
}

// Send marshals and sends the UpdateSimulationApplication API request.
func (r UpdateSimulationApplicationRequest) Send(ctx context.Context) (*UpdateSimulationApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSimulationApplicationResponse{
		UpdateSimulationApplicationOutput: r.Request.Data.(*UpdateSimulationApplicationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSimulationApplicationResponse is the response type for the
// UpdateSimulationApplication API operation.
type UpdateSimulationApplicationResponse struct {
	*UpdateSimulationApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSimulationApplication request.
func (r *UpdateSimulationApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

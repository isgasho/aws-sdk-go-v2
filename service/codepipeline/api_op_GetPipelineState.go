// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetPipelineState action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipelineStateInput
type GetPipelineStateInput struct {
	_ struct{} `type:"structure"`

	// The name of the pipeline about which you want to get information.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPipelineStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPipelineStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPipelineStateInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetPipelineState action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipelineStateOutput
type GetPipelineStateOutput struct {
	_ struct{} `type:"structure"`

	// The date and time the pipeline was created, in timestamp format.
	Created *time.Time `locationName:"created" type:"timestamp"`

	// The name of the pipeline for which you want to get the state.
	PipelineName *string `locationName:"pipelineName" min:"1" type:"string"`

	// The version number of the pipeline.
	//
	// A newly-created pipeline is always assigned a version number of 1.
	PipelineVersion *int64 `locationName:"pipelineVersion" min:"1" type:"integer"`

	// A list of the pipeline stage output information, including stage name, state,
	// most recent run details, whether the stage is disabled, and other data.
	StageStates []StageState `locationName:"stageStates" type:"list"`

	// The date and time the pipeline was last updated, in timestamp format.
	Updated *time.Time `locationName:"updated" type:"timestamp"`
}

// String returns the string representation
func (s GetPipelineStateOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPipelineState = "GetPipelineState"

// GetPipelineStateRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Returns information about the state of a pipeline, including the stages and
// actions.
//
// Values returned in the revisionId and revisionUrl fields indicate the source
// revision information, such as the commit ID, for the current state.
//
//    // Example sending a request using GetPipelineStateRequest.
//    req := client.GetPipelineStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipelineState
func (c *Client) GetPipelineStateRequest(input *GetPipelineStateInput) GetPipelineStateRequest {
	op := &aws.Operation{
		Name:       opGetPipelineState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPipelineStateInput{}
	}

	req := c.newRequest(op, input, &GetPipelineStateOutput{})
	return GetPipelineStateRequest{Request: req, Input: input, Copy: c.GetPipelineStateRequest}
}

// GetPipelineStateRequest is the request type for the
// GetPipelineState API operation.
type GetPipelineStateRequest struct {
	*aws.Request
	Input *GetPipelineStateInput
	Copy  func(*GetPipelineStateInput) GetPipelineStateRequest
}

// Send marshals and sends the GetPipelineState API request.
func (r GetPipelineStateRequest) Send(ctx context.Context) (*GetPipelineStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPipelineStateResponse{
		GetPipelineStateOutput: r.Request.Data.(*GetPipelineStateOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPipelineStateResponse is the response type for the
// GetPipelineState API operation.
type GetPipelineStateResponse struct {
	*GetPipelineStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPipelineState request.
func (r *GetPipelineStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetMaintenanceWindowExecutionRequest
type GetMaintenanceWindowExecutionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window execution that includes the task.
	//
	// WindowExecutionId is a required field
	WindowExecutionId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMaintenanceWindowExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMaintenanceWindowExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMaintenanceWindowExecutionInput"}

	if s.WindowExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowExecutionId"))
	}
	if s.WindowExecutionId != nil && len(*s.WindowExecutionId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowExecutionId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetMaintenanceWindowExecutionResult
type GetMaintenanceWindowExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The time the maintenance window finished running.
	EndTime *time.Time `type:"timestamp"`

	// The time the maintenance window started running.
	StartTime *time.Time `type:"timestamp"`

	// The status of the maintenance window execution.
	Status MaintenanceWindowExecutionStatus `type:"string" enum:"true"`

	// The details explaining the Status. Only available for certain status values.
	StatusDetails *string `type:"string"`

	// The ID of the task executions from the maintenance window execution.
	TaskIds []string `type:"list"`

	// The ID of the maintenance window execution.
	WindowExecutionId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s GetMaintenanceWindowExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMaintenanceWindowExecution = "GetMaintenanceWindowExecution"

// GetMaintenanceWindowExecutionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves details about a specific a maintenance window execution.
//
//    // Example sending a request using GetMaintenanceWindowExecutionRequest.
//    req := client.GetMaintenanceWindowExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetMaintenanceWindowExecution
func (c *Client) GetMaintenanceWindowExecutionRequest(input *GetMaintenanceWindowExecutionInput) GetMaintenanceWindowExecutionRequest {
	op := &aws.Operation{
		Name:       opGetMaintenanceWindowExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetMaintenanceWindowExecutionInput{}
	}

	req := c.newRequest(op, input, &GetMaintenanceWindowExecutionOutput{})
	return GetMaintenanceWindowExecutionRequest{Request: req, Input: input, Copy: c.GetMaintenanceWindowExecutionRequest}
}

// GetMaintenanceWindowExecutionRequest is the request type for the
// GetMaintenanceWindowExecution API operation.
type GetMaintenanceWindowExecutionRequest struct {
	*aws.Request
	Input *GetMaintenanceWindowExecutionInput
	Copy  func(*GetMaintenanceWindowExecutionInput) GetMaintenanceWindowExecutionRequest
}

// Send marshals and sends the GetMaintenanceWindowExecution API request.
func (r GetMaintenanceWindowExecutionRequest) Send(ctx context.Context) (*GetMaintenanceWindowExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMaintenanceWindowExecutionResponse{
		GetMaintenanceWindowExecutionOutput: r.Request.Data.(*GetMaintenanceWindowExecutionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMaintenanceWindowExecutionResponse is the response type for the
// GetMaintenanceWindowExecution API operation.
type GetMaintenanceWindowExecutionResponse struct {
	*GetMaintenanceWindowExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMaintenanceWindowExecution request.
func (r *GetMaintenanceWindowExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/StartExportTaskRequest
type StartExportTaskInput struct {
	_ struct{} `type:"structure"`

	// The end timestamp for exported data from the single Application Discovery
	// Agent selected in the filters. If no value is specified, exported data includes
	// the most recent data collected by the agent.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The file format for the returned export data. Default value is CSV. Note:
	// The GRAPHML option has been deprecated.
	ExportDataFormat []ExportDataFormat `locationName:"exportDataFormat" type:"list"`

	// If a filter is present, it selects the single agentId of the Application
	// Discovery Agent for which data is exported. The agentId can be found in the
	// results of the DescribeAgents API or CLI. If no filter is present, startTime
	// and endTime are ignored and exported data includes both Agentless Discovery
	// Connector data and summary data from Application Discovery agents.
	Filters []ExportFilter `locationName:"filters" type:"list"`

	// The start timestamp for exported data from the single Application Discovery
	// Agent selected in the filters. If no value is specified, data is exported
	// starting from the first data collected by the agent.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`
}

// String returns the string representation
func (s StartExportTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartExportTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartExportTaskInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/StartExportTaskResponse
type StartExportTaskOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier used to query the status of an export request.
	ExportId *string `locationName:"exportId" type:"string"`
}

// String returns the string representation
func (s StartExportTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartExportTask = "StartExportTask"

// StartExportTaskRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Begins the export of discovered data to an S3 bucket.
//
// If you specify agentIds in a filter, the task exports up to 72 hours of detailed
// data collected by the identified Application Discovery Agent, including network,
// process, and performance details. A time range for exported agent data may
// be set by using startTime and endTime. Export of detailed agent data is limited
// to five concurrently running exports.
//
// If you do not include an agentIds filter, summary data is exported that includes
// both AWS Agentless Discovery Connector data and summary data from AWS Discovery
// Agents. Export of summary data is limited to two exports per day.
//
//    // Example sending a request using StartExportTaskRequest.
//    req := client.StartExportTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/StartExportTask
func (c *Client) StartExportTaskRequest(input *StartExportTaskInput) StartExportTaskRequest {
	op := &aws.Operation{
		Name:       opStartExportTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartExportTaskInput{}
	}

	req := c.newRequest(op, input, &StartExportTaskOutput{})
	return StartExportTaskRequest{Request: req, Input: input, Copy: c.StartExportTaskRequest}
}

// StartExportTaskRequest is the request type for the
// StartExportTask API operation.
type StartExportTaskRequest struct {
	*aws.Request
	Input *StartExportTaskInput
	Copy  func(*StartExportTaskInput) StartExportTaskRequest
}

// Send marshals and sends the StartExportTask API request.
func (r StartExportTaskRequest) Send(ctx context.Context) (*StartExportTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartExportTaskResponse{
		StartExportTaskOutput: r.Request.Data.(*StartExportTaskOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartExportTaskResponse is the response type for the
// StartExportTask API operation.
type StartExportTaskResponse struct {
	*StartExportTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartExportTask request.
func (r *StartExportTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

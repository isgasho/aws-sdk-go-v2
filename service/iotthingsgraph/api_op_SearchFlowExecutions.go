// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchFlowExecutionsRequest
type SearchFlowExecutionsInput struct {
	_ struct{} `type:"structure"`

	// The date and time of the latest flow execution to return.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The ID of a flow execution.
	FlowExecutionId *string `locationName:"flowExecutionId" type:"string"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The date and time of the earliest flow execution to return.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`

	// The ID of the system instance that contains the flow.
	//
	// SystemInstanceId is a required field
	SystemInstanceId *string `locationName:"systemInstanceId" type:"string" required:"true"`
}

// String returns the string representation
func (s SearchFlowExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchFlowExecutionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchFlowExecutionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.SystemInstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SystemInstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchFlowExecutionsResponse
type SearchFlowExecutionsOutput struct {
	_ struct{} `type:"structure"`

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// An array of objects that contain summary information about each workflow
	// execution in the result set.
	Summaries []FlowExecutionSummary `locationName:"summaries" type:"list"`
}

// String returns the string representation
func (s SearchFlowExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchFlowExecutions = "SearchFlowExecutions"

// SearchFlowExecutionsRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Searches for AWS IoT Things Graph workflow execution instances.
//
//    // Example sending a request using SearchFlowExecutionsRequest.
//    req := client.SearchFlowExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchFlowExecutions
func (c *Client) SearchFlowExecutionsRequest(input *SearchFlowExecutionsInput) SearchFlowExecutionsRequest {
	op := &aws.Operation{
		Name:       opSearchFlowExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &SearchFlowExecutionsInput{}
	}

	req := c.newRequest(op, input, &SearchFlowExecutionsOutput{})
	return SearchFlowExecutionsRequest{Request: req, Input: input, Copy: c.SearchFlowExecutionsRequest}
}

// SearchFlowExecutionsRequest is the request type for the
// SearchFlowExecutions API operation.
type SearchFlowExecutionsRequest struct {
	*aws.Request
	Input *SearchFlowExecutionsInput
	Copy  func(*SearchFlowExecutionsInput) SearchFlowExecutionsRequest
}

// Send marshals and sends the SearchFlowExecutions API request.
func (r SearchFlowExecutionsRequest) Send(ctx context.Context) (*SearchFlowExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchFlowExecutionsResponse{
		SearchFlowExecutionsOutput: r.Request.Data.(*SearchFlowExecutionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchFlowExecutionsRequestPaginator returns a paginator for SearchFlowExecutions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchFlowExecutionsRequest(input)
//   p := iotthingsgraph.NewSearchFlowExecutionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchFlowExecutionsPaginator(req SearchFlowExecutionsRequest) SearchFlowExecutionsPaginator {
	return SearchFlowExecutionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchFlowExecutionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// SearchFlowExecutionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchFlowExecutionsPaginator struct {
	aws.Pager
}

func (p *SearchFlowExecutionsPaginator) CurrentPage() *SearchFlowExecutionsOutput {
	return p.Pager.CurrentPage().(*SearchFlowExecutionsOutput)
}

// SearchFlowExecutionsResponse is the response type for the
// SearchFlowExecutions API operation.
type SearchFlowExecutionsResponse struct {
	*SearchFlowExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchFlowExecutions request.
func (r *SearchFlowExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DescribeScheduleRequest
type DescribeScheduleInput struct {
	_ struct{} `type:"structure"`

	// ChannelId is a required field
	ChannelId *string `location:"uri" locationName:"channelId" type:"string" required:"true"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScheduleInput"}

	if s.ChannelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeScheduleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChannelId != nil {
		v := *s.ChannelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "channelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DescribeScheduleResponse
type DescribeScheduleOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	ScheduleActions []ScheduleAction `locationName:"scheduleActions" type:"list"`
}

// String returns the string representation
func (s DescribeScheduleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeScheduleOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ScheduleActions != nil {
		v := s.ScheduleActions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "scheduleActions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeSchedule = "DescribeSchedule"

// DescribeScheduleRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Get a channel schedule
//
//    // Example sending a request using DescribeScheduleRequest.
//    req := client.DescribeScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DescribeSchedule
func (c *Client) DescribeScheduleRequest(input *DescribeScheduleInput) DescribeScheduleRequest {
	op := &aws.Operation{
		Name:       opDescribeSchedule,
		HTTPMethod: "GET",
		HTTPPath:   "/prod/channels/{channelId}/schedule",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeScheduleInput{}
	}

	req := c.newRequest(op, input, &DescribeScheduleOutput{})
	return DescribeScheduleRequest{Request: req, Input: input, Copy: c.DescribeScheduleRequest}
}

// DescribeScheduleRequest is the request type for the
// DescribeSchedule API operation.
type DescribeScheduleRequest struct {
	*aws.Request
	Input *DescribeScheduleInput
	Copy  func(*DescribeScheduleInput) DescribeScheduleRequest
}

// Send marshals and sends the DescribeSchedule API request.
func (r DescribeScheduleRequest) Send(ctx context.Context) (*DescribeScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScheduleResponse{
		DescribeScheduleOutput: r.Request.Data.(*DescribeScheduleOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeScheduleRequestPaginator returns a paginator for DescribeSchedule.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeScheduleRequest(input)
//   p := medialive.NewDescribeScheduleRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeSchedulePaginator(req DescribeScheduleRequest) DescribeSchedulePaginator {
	return DescribeSchedulePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeScheduleInput
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

// DescribeSchedulePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeSchedulePaginator struct {
	aws.Pager
}

func (p *DescribeSchedulePaginator) CurrentPage() *DescribeScheduleOutput {
	return p.Pager.CurrentPage().(*DescribeScheduleOutput)
}

// DescribeScheduleResponse is the response type for the
// DescribeSchedule API operation.
type DescribeScheduleResponse struct {
	*DescribeScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSchedule request.
func (r *DescribeScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

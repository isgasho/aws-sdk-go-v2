// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListContactsRequest
type ListContactsInput struct {
	_ struct{} `type:"structure"`

	// End time of a contact.
	//
	// EndTime is a required field
	EndTime *time.Time `locationName:"endTime" type:"timestamp" required:"true"`

	// Name of a ground station.
	GroundStation *string `locationName:"groundStation" type:"string"`

	// Maximum number of contacts returned.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// ARN of a mission profile.
	MissionProfileArn *string `locationName:"missionProfileArn" type:"string"`

	// Next token returned in the request of a previous ListContacts call. Used
	// to get the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// ARN of a satellite.
	SatelliteArn *string `locationName:"satelliteArn" type:"string"`

	// Start time of a contact.
	//
	// StartTime is a required field
	StartTime *time.Time `locationName:"startTime" type:"timestamp" required:"true"`

	// Status of a contact reservation.
	//
	// StatusList is a required field
	StatusList []ContactStatus `locationName:"statusList" type:"list" required:"true"`
}

// String returns the string representation
func (s ListContactsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListContactsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListContactsInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if s.StatusList == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatusList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListContactsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.GroundStation != nil {
		v := *s.GroundStation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "groundStation", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.MissionProfileArn != nil {
		v := *s.MissionProfileArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "missionProfileArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SatelliteArn != nil {
		v := *s.SatelliteArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "satelliteArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.StatusList != nil {
		v := s.StatusList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "statusList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListContactsResponse
type ListContactsOutput struct {
	_ struct{} `type:"structure"`

	// List of contacts.
	ContactList []ContactData `locationName:"contactList" type:"list"`

	// Next token returned in the response of a previous ListContacts call. Used
	// to get the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListContactsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListContactsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContactList != nil {
		v := s.ContactList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "contactList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListContacts = "ListContacts"

// ListContactsRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns a list of contacts.
//
// If statusList contains AVAILABLE, the request must include groundstation,
// missionprofileArn, and satelliteArn.
//
//    // Example sending a request using ListContactsRequest.
//    req := client.ListContactsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListContacts
func (c *Client) ListContactsRequest(input *ListContactsInput) ListContactsRequest {
	op := &aws.Operation{
		Name:       opListContacts,
		HTTPMethod: "POST",
		HTTPPath:   "/contacts",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListContactsInput{}
	}

	req := c.newRequest(op, input, &ListContactsOutput{})
	return ListContactsRequest{Request: req, Input: input, Copy: c.ListContactsRequest}
}

// ListContactsRequest is the request type for the
// ListContacts API operation.
type ListContactsRequest struct {
	*aws.Request
	Input *ListContactsInput
	Copy  func(*ListContactsInput) ListContactsRequest
}

// Send marshals and sends the ListContacts API request.
func (r ListContactsRequest) Send(ctx context.Context) (*ListContactsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListContactsResponse{
		ListContactsOutput: r.Request.Data.(*ListContactsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListContactsRequestPaginator returns a paginator for ListContacts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListContactsRequest(input)
//   p := groundstation.NewListContactsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListContactsPaginator(req ListContactsRequest) ListContactsPaginator {
	return ListContactsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListContactsInput
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

// ListContactsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListContactsPaginator struct {
	aws.Pager
}

func (p *ListContactsPaginator) CurrentPage() *ListContactsOutput {
	return p.Pager.CurrentPage().(*ListContactsOutput)
}

// ListContactsResponse is the response type for the
// ListContacts API operation.
type ListContactsResponse struct {
	*ListContactsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListContacts request.
func (r *ListContactsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

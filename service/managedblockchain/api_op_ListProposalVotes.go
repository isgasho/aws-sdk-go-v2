// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListProposalVotesInput
type ListProposalVotesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of votes to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The unique identifier of the network.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The unique identifier of the proposal.
	//
	// ProposalId is a required field
	ProposalId *string `location:"uri" locationName:"proposalId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListProposalVotesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProposalVotesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProposalVotesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if s.ProposalId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProposalId"))
	}
	if s.ProposalId != nil && len(*s.ProposalId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProposalId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProposalVotesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProposalId != nil {
		v := *s.ProposalId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "proposalId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListProposalVotesOutput
type ListProposalVotesOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`

	// The listing of votes.
	ProposalVotes []VoteSummary `type:"list"`
}

// String returns the string representation
func (s ListProposalVotesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProposalVotesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProposalVotes != nil {
		v := s.ProposalVotes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ProposalVotes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListProposalVotes = "ListProposalVotes"

// ListProposalVotesRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns the listing of votes for a specified proposal, including the value
// of each vote and the unique identifier of the member that cast the vote.
//
//    // Example sending a request using ListProposalVotesRequest.
//    req := client.ListProposalVotesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListProposalVotes
func (c *Client) ListProposalVotesRequest(input *ListProposalVotesInput) ListProposalVotesRequest {
	op := &aws.Operation{
		Name:       opListProposalVotes,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}/proposals/{proposalId}/votes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListProposalVotesInput{}
	}

	req := c.newRequest(op, input, &ListProposalVotesOutput{})
	return ListProposalVotesRequest{Request: req, Input: input, Copy: c.ListProposalVotesRequest}
}

// ListProposalVotesRequest is the request type for the
// ListProposalVotes API operation.
type ListProposalVotesRequest struct {
	*aws.Request
	Input *ListProposalVotesInput
	Copy  func(*ListProposalVotesInput) ListProposalVotesRequest
}

// Send marshals and sends the ListProposalVotes API request.
func (r ListProposalVotesRequest) Send(ctx context.Context) (*ListProposalVotesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProposalVotesResponse{
		ListProposalVotesOutput: r.Request.Data.(*ListProposalVotesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProposalVotesRequestPaginator returns a paginator for ListProposalVotes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProposalVotesRequest(input)
//   p := managedblockchain.NewListProposalVotesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProposalVotesPaginator(req ListProposalVotesRequest) ListProposalVotesPaginator {
	return ListProposalVotesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListProposalVotesInput
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

// ListProposalVotesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProposalVotesPaginator struct {
	aws.Pager
}

func (p *ListProposalVotesPaginator) CurrentPage() *ListProposalVotesOutput {
	return p.Pager.CurrentPage().(*ListProposalVotesOutput)
}

// ListProposalVotesResponse is the response type for the
// ListProposalVotes API operation.
type ListProposalVotesResponse struct {
	*ListProposalVotesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProposalVotes request.
func (r *ListProposalVotesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

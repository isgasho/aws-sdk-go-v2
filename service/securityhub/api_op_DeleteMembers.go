// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteMembersRequest
type DeleteMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs of the member accounts to delete.
	AccountIds []string `type:"list"`
}

// String returns the string representation
func (s DeleteMembersInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AccountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteMembersResponse
type DeleteMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of account ID and email address pairs of the AWS accounts that weren't
	// deleted.
	UnprocessedAccounts []Result `type:"list"`
}

// String returns the string representation
func (s DeleteMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UnprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDeleteMembers = "DeleteMembers"

// DeleteMembersRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Deletes the specified member accounts from Security Hub.
//
//    // Example sending a request using DeleteMembersRequest.
//    req := client.DeleteMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteMembers
func (c *Client) DeleteMembersRequest(input *DeleteMembersInput) DeleteMembersRequest {
	op := &aws.Operation{
		Name:       opDeleteMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/members/delete",
	}

	if input == nil {
		input = &DeleteMembersInput{}
	}

	req := c.newRequest(op, input, &DeleteMembersOutput{})
	return DeleteMembersRequest{Request: req, Input: input, Copy: c.DeleteMembersRequest}
}

// DeleteMembersRequest is the request type for the
// DeleteMembers API operation.
type DeleteMembersRequest struct {
	*aws.Request
	Input *DeleteMembersInput
	Copy  func(*DeleteMembersInput) DeleteMembersRequest
}

// Send marshals and sends the DeleteMembers API request.
func (r DeleteMembersRequest) Send(ctx context.Context) (*DeleteMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMembersResponse{
		DeleteMembersOutput: r.Request.Data.(*DeleteMembersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMembersResponse is the response type for the
// DeleteMembers API operation.
type DeleteMembersResponse struct {
	*DeleteMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMembers request.
func (r *DeleteMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// The name of the authorizer to describe.
	//
	// AuthorizerName is a required field
	AuthorizerName *string `location:"uri" locationName:"authorizerName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAuthorizerInput"}

	if s.AuthorizerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerName"))
	}
	if s.AuthorizerName != nil && len(*s.AuthorizerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthorizerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAuthorizerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthorizerName != nil {
		v := *s.AuthorizerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "authorizerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeAuthorizerOutput struct {
	_ struct{} `type:"structure"`

	// The authorizer description.
	AuthorizerDescription *AuthorizerDescription `locationName:"authorizerDescription" type:"structure"`
}

// String returns the string representation
func (s DescribeAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAuthorizerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AuthorizerDescription != nil {
		v := s.AuthorizerDescription

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "authorizerDescription", v, metadata)
	}
	return nil
}

const opDescribeAuthorizer = "DescribeAuthorizer"

// DescribeAuthorizerRequest returns a request value for making API operation for
// AWS IoT.
//
// Describes an authorizer.
//
//    // Example sending a request using DescribeAuthorizerRequest.
//    req := client.DescribeAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeAuthorizerRequest(input *DescribeAuthorizerInput) DescribeAuthorizerRequest {
	op := &aws.Operation{
		Name:       opDescribeAuthorizer,
		HTTPMethod: "GET",
		HTTPPath:   "/authorizer/{authorizerName}",
	}

	if input == nil {
		input = &DescribeAuthorizerInput{}
	}

	req := c.newRequest(op, input, &DescribeAuthorizerOutput{})
	return DescribeAuthorizerRequest{Request: req, Input: input, Copy: c.DescribeAuthorizerRequest}
}

// DescribeAuthorizerRequest is the request type for the
// DescribeAuthorizer API operation.
type DescribeAuthorizerRequest struct {
	*aws.Request
	Input *DescribeAuthorizerInput
	Copy  func(*DescribeAuthorizerInput) DescribeAuthorizerRequest
}

// Send marshals and sends the DescribeAuthorizer API request.
func (r DescribeAuthorizerRequest) Send(ctx context.Context) (*DescribeAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAuthorizerResponse{
		DescribeAuthorizerOutput: r.Request.Data.(*DescribeAuthorizerOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAuthorizerResponse is the response type for the
// DescribeAuthorizer API operation.
type DescribeAuthorizerResponse struct {
	*DescribeAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAuthorizer request.
func (r *DescribeAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/GetResolverRequest
type GetResolverInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The resolver field name.
	//
	// FieldName is a required field
	FieldName *string `location:"uri" locationName:"fieldName" type:"string" required:"true"`

	// The resolver type name.
	//
	// TypeName is a required field
	TypeName *string `location:"uri" locationName:"typeName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetResolverInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResolverInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResolverInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.FieldName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FieldName"))
	}

	if s.TypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypeName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResolverInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FieldName != nil {
		v := *s.FieldName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "fieldName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TypeName != nil {
		v := *s.TypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "typeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/GetResolverResponse
type GetResolverOutput struct {
	_ struct{} `type:"structure"`

	// The Resolver object.
	Resolver *Resolver `locationName:"resolver" type:"structure"`
}

// String returns the string representation
func (s GetResolverOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResolverOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Resolver != nil {
		v := s.Resolver

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resolver", v, metadata)
	}
	return nil
}

const opGetResolver = "GetResolver"

// GetResolverRequest returns a request value for making API operation for
// AWS AppSync.
//
// Retrieves a Resolver object.
//
//    // Example sending a request using GetResolverRequest.
//    req := client.GetResolverRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/GetResolver
func (c *Client) GetResolverRequest(input *GetResolverInput) GetResolverRequest {
	op := &aws.Operation{
		Name:       opGetResolver,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/types/{typeName}/resolvers/{fieldName}",
	}

	if input == nil {
		input = &GetResolverInput{}
	}

	req := c.newRequest(op, input, &GetResolverOutput{})
	return GetResolverRequest{Request: req, Input: input, Copy: c.GetResolverRequest}
}

// GetResolverRequest is the request type for the
// GetResolver API operation.
type GetResolverRequest struct {
	*aws.Request
	Input *GetResolverInput
	Copy  func(*GetResolverInput) GetResolverRequest
}

// Send marshals and sends the GetResolver API request.
func (r GetResolverRequest) Send(ctx context.Context) (*GetResolverResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResolverResponse{
		GetResolverOutput: r.Request.Data.(*GetResolverOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetResolverResponse is the response type for the
// GetResolver API operation.
type GetResolverResponse struct {
	*GetResolverOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResolver request.
func (r *GetResolverResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

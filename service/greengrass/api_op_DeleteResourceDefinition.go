// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteResourceDefinitionRequest
type DeleteResourceDefinitionInput struct {
	_ struct{} `type:"structure"`

	// ResourceDefinitionId is a required field
	ResourceDefinitionId *string `location:"uri" locationName:"ResourceDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResourceDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResourceDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResourceDefinitionInput"}

	if s.ResourceDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteResourceDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ResourceDefinitionId != nil {
		v := *s.ResourceDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteResourceDefinitionResponse
type DeleteResourceDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteResourceDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteResourceDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteResourceDefinition = "DeleteResourceDefinition"

// DeleteResourceDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Deletes a resource definition.
//
//    // Example sending a request using DeleteResourceDefinitionRequest.
//    req := client.DeleteResourceDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteResourceDefinition
func (c *Client) DeleteResourceDefinitionRequest(input *DeleteResourceDefinitionInput) DeleteResourceDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteResourceDefinition,
		HTTPMethod: "DELETE",
		HTTPPath:   "/greengrass/definition/resources/{ResourceDefinitionId}",
	}

	if input == nil {
		input = &DeleteResourceDefinitionInput{}
	}

	req := c.newRequest(op, input, &DeleteResourceDefinitionOutput{})
	return DeleteResourceDefinitionRequest{Request: req, Input: input, Copy: c.DeleteResourceDefinitionRequest}
}

// DeleteResourceDefinitionRequest is the request type for the
// DeleteResourceDefinition API operation.
type DeleteResourceDefinitionRequest struct {
	*aws.Request
	Input *DeleteResourceDefinitionInput
	Copy  func(*DeleteResourceDefinitionInput) DeleteResourceDefinitionRequest
}

// Send marshals and sends the DeleteResourceDefinition API request.
func (r DeleteResourceDefinitionRequest) Send(ctx context.Context) (*DeleteResourceDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResourceDefinitionResponse{
		DeleteResourceDefinitionOutput: r.Request.Data.(*DeleteResourceDefinitionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResourceDefinitionResponse is the response type for the
// DeleteResourceDefinition API operation.
type DeleteResourceDefinitionResponse struct {
	*DeleteResourceDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResourceDefinition request.
func (r *DeleteResourceDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

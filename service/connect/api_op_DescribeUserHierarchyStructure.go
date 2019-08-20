// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/DescribeUserHierarchyStructureRequest
type DescribeUserHierarchyStructureInput struct {
	_ struct{} `type:"structure"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserHierarchyStructureInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserHierarchyStructureInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserHierarchyStructureInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUserHierarchyStructureInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/DescribeUserHierarchyStructureResponse
type DescribeUserHierarchyStructureOutput struct {
	_ struct{} `type:"structure"`

	// A HierarchyStructure object.
	HierarchyStructure *HierarchyStructure `type:"structure"`
}

// String returns the string representation
func (s DescribeUserHierarchyStructureOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUserHierarchyStructureOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HierarchyStructure != nil {
		v := s.HierarchyStructure

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HierarchyStructure", v, metadata)
	}
	return nil
}

const opDescribeUserHierarchyStructure = "DescribeUserHierarchyStructure"

// DescribeUserHierarchyStructureRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Returns a HiearchyGroupStructure object, which contains data about the levels
// in the agent hierarchy.
//
//    // Example sending a request using DescribeUserHierarchyStructureRequest.
//    req := client.DescribeUserHierarchyStructureRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/DescribeUserHierarchyStructure
func (c *Client) DescribeUserHierarchyStructureRequest(input *DescribeUserHierarchyStructureInput) DescribeUserHierarchyStructureRequest {
	op := &aws.Operation{
		Name:       opDescribeUserHierarchyStructure,
		HTTPMethod: "GET",
		HTTPPath:   "/user-hierarchy-structure/{InstanceId}",
	}

	if input == nil {
		input = &DescribeUserHierarchyStructureInput{}
	}

	req := c.newRequest(op, input, &DescribeUserHierarchyStructureOutput{})
	return DescribeUserHierarchyStructureRequest{Request: req, Input: input, Copy: c.DescribeUserHierarchyStructureRequest}
}

// DescribeUserHierarchyStructureRequest is the request type for the
// DescribeUserHierarchyStructure API operation.
type DescribeUserHierarchyStructureRequest struct {
	*aws.Request
	Input *DescribeUserHierarchyStructureInput
	Copy  func(*DescribeUserHierarchyStructureInput) DescribeUserHierarchyStructureRequest
}

// Send marshals and sends the DescribeUserHierarchyStructure API request.
func (r DescribeUserHierarchyStructureRequest) Send(ctx context.Context) (*DescribeUserHierarchyStructureResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserHierarchyStructureResponse{
		DescribeUserHierarchyStructureOutput: r.Request.Data.(*DescribeUserHierarchyStructureOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserHierarchyStructureResponse is the response type for the
// DescribeUserHierarchyStructure API operation.
type DescribeUserHierarchyStructureResponse struct {
	*DescribeUserHierarchyStructureOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUserHierarchyStructure request.
func (r *DescribeUserHierarchyStructureResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

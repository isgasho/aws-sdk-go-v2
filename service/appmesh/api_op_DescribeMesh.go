// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DescribeMeshInput
type DescribeMeshInput struct {
	_ struct{} `type:"structure"`

	// The name of the service mesh to describe.
	//
	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeMeshInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMeshInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMeshInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeMeshInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DescribeMeshOutput
type DescribeMeshOutput struct {
	_ struct{} `type:"structure" payload:"Mesh"`

	// The full description of your service mesh.
	//
	// Mesh is a required field
	Mesh *MeshData `locationName:"mesh" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeMeshOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeMeshOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Mesh != nil {
		v := s.Mesh

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "mesh", v, metadata)
	}
	return nil
}

const opDescribeMesh = "DescribeMesh"

// DescribeMeshRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Describes an existing service mesh.
//
//    // Example sending a request using DescribeMeshRequest.
//    req := client.DescribeMeshRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DescribeMesh
func (c *Client) DescribeMeshRequest(input *DescribeMeshInput) DescribeMeshRequest {
	op := &aws.Operation{
		Name:       opDescribeMesh,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}",
	}

	if input == nil {
		input = &DescribeMeshInput{}
	}

	req := c.newRequest(op, input, &DescribeMeshOutput{})
	return DescribeMeshRequest{Request: req, Input: input, Copy: c.DescribeMeshRequest}
}

// DescribeMeshRequest is the request type for the
// DescribeMesh API operation.
type DescribeMeshRequest struct {
	*aws.Request
	Input *DescribeMeshInput
	Copy  func(*DescribeMeshInput) DescribeMeshRequest
}

// Send marshals and sends the DescribeMesh API request.
func (r DescribeMeshRequest) Send(ctx context.Context) (*DescribeMeshResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMeshResponse{
		DescribeMeshOutput: r.Request.Data.(*DescribeMeshOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMeshResponse is the response type for the
// DescribeMesh API operation.
type DescribeMeshResponse struct {
	*DescribeMeshOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMesh request.
func (r *DescribeMeshResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

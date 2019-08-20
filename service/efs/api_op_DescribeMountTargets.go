// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DescribeMountTargetsRequest
type DescribeMountTargetsInput struct {
	_ struct{} `type:"structure"`

	// (Optional) ID of the file system whose mount targets you want to list (String).
	// It must be included in your request if MountTargetId is not included.
	FileSystemId *string `location:"querystring" locationName:"FileSystemId" type:"string"`

	// (Optional) Opaque pagination token returned from a previous DescribeMountTargets
	// operation (String). If present, it specifies to continue the list from where
	// the previous returning call left off.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// (Optional) Maximum number of mount targets to return in the response. Currently,
	// this number is automatically set to 10, and other values are ignored. The
	// response is paginated at 10 per page if you have more than 10 mount targets.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" min:"1" type:"integer"`

	// (Optional) ID of the mount target that you want to have described (String).
	// It must be included in your request if FileSystemId is not included.
	MountTargetId *string `location:"querystring" locationName:"MountTargetId" type:"string"`
}

// String returns the string representation
func (s DescribeMountTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMountTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMountTargetsInput"}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeMountTargetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	if s.MountTargetId != nil {
		v := *s.MountTargetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MountTargetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DescribeMountTargetsResponse
type DescribeMountTargetsOutput struct {
	_ struct{} `type:"structure"`

	// If the request included the Marker, the response returns that value in this
	// field.
	Marker *string `type:"string"`

	// Returns the file system's mount targets as an array of MountTargetDescription
	// objects.
	MountTargets []MountTargetDescription `type:"list"`

	// If a value is present, there are more mount targets to return. In a subsequent
	// request, you can provide Marker in your request with this value to retrieve
	// the next set of mount targets.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s DescribeMountTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeMountTargetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MountTargets != nil {
		v := s.MountTargets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "MountTargets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeMountTargets = "DescribeMountTargets"

// DescribeMountTargetsRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Returns the descriptions of all the current mount targets, or a specific
// mount target, for a file system. When requesting all of the current mount
// targets, the order of mount targets returned in the response is unspecified.
//
// This operation requires permissions for the elasticfilesystem:DescribeMountTargets
// action, on either the file system ID that you specify in FileSystemId, or
// on the file system of the mount target that you specify in MountTargetId.
//
//    // Example sending a request using DescribeMountTargetsRequest.
//    req := client.DescribeMountTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DescribeMountTargets
func (c *Client) DescribeMountTargetsRequest(input *DescribeMountTargetsInput) DescribeMountTargetsRequest {
	op := &aws.Operation{
		Name:       opDescribeMountTargets,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-02-01/mount-targets",
	}

	if input == nil {
		input = &DescribeMountTargetsInput{}
	}

	req := c.newRequest(op, input, &DescribeMountTargetsOutput{})
	return DescribeMountTargetsRequest{Request: req, Input: input, Copy: c.DescribeMountTargetsRequest}
}

// DescribeMountTargetsRequest is the request type for the
// DescribeMountTargets API operation.
type DescribeMountTargetsRequest struct {
	*aws.Request
	Input *DescribeMountTargetsInput
	Copy  func(*DescribeMountTargetsInput) DescribeMountTargetsRequest
}

// Send marshals and sends the DescribeMountTargets API request.
func (r DescribeMountTargetsRequest) Send(ctx context.Context) (*DescribeMountTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMountTargetsResponse{
		DescribeMountTargetsOutput: r.Request.Data.(*DescribeMountTargetsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMountTargetsResponse is the response type for the
// DescribeMountTargets API operation.
type DescribeMountTargetsResponse struct {
	*DescribeMountTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMountTargets request.
func (r *DescribeMountTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

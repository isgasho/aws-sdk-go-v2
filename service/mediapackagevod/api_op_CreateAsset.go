// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/CreateAssetRequest
type CreateAssetInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`

	// PackagingGroupId is a required field
	PackagingGroupId *string `locationName:"packagingGroupId" type:"string" required:"true"`

	ResourceId *string `locationName:"resourceId" type:"string"`

	// SourceArn is a required field
	SourceArn *string `locationName:"sourceArn" type:"string" required:"true"`

	// SourceRoleArn is a required field
	SourceRoleArn *string `locationName:"sourceRoleArn" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAssetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAssetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAssetInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.PackagingGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PackagingGroupId"))
	}

	if s.SourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceArn"))
	}

	if s.SourceRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceRoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAssetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackagingGroupId != nil {
		v := *s.PackagingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "packagingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceRoleArn != nil {
		v := *s.SourceRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/CreateAssetResponse
type CreateAssetOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	EgressEndpoints []EgressEndpoint `locationName:"egressEndpoints" type:"list"`

	Id *string `locationName:"id" type:"string"`

	PackagingGroupId *string `locationName:"packagingGroupId" type:"string"`

	ResourceId *string `locationName:"resourceId" type:"string"`

	SourceArn *string `locationName:"sourceArn" type:"string"`

	SourceRoleArn *string `locationName:"sourceRoleArn" type:"string"`
}

// String returns the string representation
func (s CreateAssetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAssetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EgressEndpoints != nil {
		v := s.EgressEndpoints

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "egressEndpoints", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackagingGroupId != nil {
		v := *s.PackagingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "packagingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceRoleArn != nil {
		v := *s.SourceRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateAsset = "CreateAsset"

// CreateAssetRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Creates a new MediaPackage VOD Asset resource.
//
//    // Example sending a request using CreateAssetRequest.
//    req := client.CreateAssetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/CreateAsset
func (c *Client) CreateAssetRequest(input *CreateAssetInput) CreateAssetRequest {
	op := &aws.Operation{
		Name:       opCreateAsset,
		HTTPMethod: "POST",
		HTTPPath:   "/assets",
	}

	if input == nil {
		input = &CreateAssetInput{}
	}

	req := c.newRequest(op, input, &CreateAssetOutput{})
	return CreateAssetRequest{Request: req, Input: input, Copy: c.CreateAssetRequest}
}

// CreateAssetRequest is the request type for the
// CreateAsset API operation.
type CreateAssetRequest struct {
	*aws.Request
	Input *CreateAssetInput
	Copy  func(*CreateAssetInput) CreateAssetRequest
}

// Send marshals and sends the CreateAsset API request.
func (r CreateAssetRequest) Send(ctx context.Context) (*CreateAssetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAssetResponse{
		CreateAssetOutput: r.Request.Data.(*CreateAssetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAssetResponse is the response type for the
// CreateAsset API operation.
type CreateAssetResponse struct {
	*CreateAssetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAsset request.
func (r *CreateAssetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

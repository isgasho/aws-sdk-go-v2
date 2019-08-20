// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Associates the Amazon Resource Name (ARN) of an AWS Certificate Manager (ACM)
// certificate with an AWS Elemental MediaConvert resource.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/AssociateCertificateRequest
type AssociateCertificateInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the ACM certificate that you want to associate with your MediaConvert
	// resource.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateCertificateInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateCertificateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Successful association of Certificate Manager Amazon Resource Name (ARN)
// with Mediaconvert returns an OK message.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/AssociateCertificateResponse
type AssociateCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateCertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAssociateCertificate = "AssociateCertificate"

// AssociateCertificateRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Associates an AWS Certificate Manager (ACM) Amazon Resource Name (ARN) with
// AWS Elemental MediaConvert.
//
//    // Example sending a request using AssociateCertificateRequest.
//    req := client.AssociateCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/AssociateCertificate
func (c *Client) AssociateCertificateRequest(input *AssociateCertificateInput) AssociateCertificateRequest {
	op := &aws.Operation{
		Name:       opAssociateCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/2017-08-29/certificates",
	}

	if input == nil {
		input = &AssociateCertificateInput{}
	}

	req := c.newRequest(op, input, &AssociateCertificateOutput{})
	return AssociateCertificateRequest{Request: req, Input: input, Copy: c.AssociateCertificateRequest}
}

// AssociateCertificateRequest is the request type for the
// AssociateCertificate API operation.
type AssociateCertificateRequest struct {
	*aws.Request
	Input *AssociateCertificateInput
	Copy  func(*AssociateCertificateInput) AssociateCertificateRequest
}

// Send marshals and sends the AssociateCertificate API request.
func (r AssociateCertificateRequest) Send(ctx context.Context) (*AssociateCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateCertificateResponse{
		AssociateCertificateOutput: r.Request.Data.(*AssociateCertificateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateCertificateResponse is the response type for the
// AssociateCertificate API operation.
type AssociateCertificateResponse struct {
	*AssociateCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateCertificate request.
func (r *AssociateCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

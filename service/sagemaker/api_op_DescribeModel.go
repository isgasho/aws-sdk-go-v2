// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeModelInput
type DescribeModelInput struct {
	_ struct{} `type:"structure"`

	// The name of the model.
	//
	// ModelName is a required field
	ModelName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeModelInput"}

	if s.ModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeModelOutput
type DescribeModelOutput struct {
	_ struct{} `type:"structure"`

	// The containers in the inference pipeline.
	Containers []ContainerDefinition `type:"list"`

	// A timestamp that shows when the model was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// If True, no inbound or outbound network calls can be made to or from the
	// model container.
	//
	// The Semantic Segmentation built-in algorithm does not support network isolation.
	EnableNetworkIsolation *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the IAM role that you specified for the
	// model.
	//
	// ExecutionRoleArn is a required field
	ExecutionRoleArn *string `min:"20" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the model.
	//
	// ModelArn is a required field
	ModelArn *string `min:"20" type:"string" required:"true"`

	// Name of the Amazon SageMaker model.
	//
	// ModelName is a required field
	ModelName *string `type:"string" required:"true"`

	// The location of the primary inference code, associated artifacts, and custom
	// environment map that the inference code uses when it is deployed in production.
	PrimaryContainer *ContainerDefinition `type:"structure"`

	// A VpcConfig object that specifies the VPC that this model has access to.
	// For more information, see Protect Endpoints by Using an Amazon Virtual Private
	// Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html)
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s DescribeModelOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeModel = "DescribeModel"

// DescribeModelRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Describes a model that you created using the CreateModel API.
//
//    // Example sending a request using DescribeModelRequest.
//    req := client.DescribeModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeModel
func (c *Client) DescribeModelRequest(input *DescribeModelInput) DescribeModelRequest {
	op := &aws.Operation{
		Name:       opDescribeModel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeModelInput{}
	}

	req := c.newRequest(op, input, &DescribeModelOutput{})
	return DescribeModelRequest{Request: req, Input: input, Copy: c.DescribeModelRequest}
}

// DescribeModelRequest is the request type for the
// DescribeModel API operation.
type DescribeModelRequest struct {
	*aws.Request
	Input *DescribeModelInput
	Copy  func(*DescribeModelInput) DescribeModelRequest
}

// Send marshals and sends the DescribeModel API request.
func (r DescribeModelRequest) Send(ctx context.Context) (*DescribeModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeModelResponse{
		DescribeModelOutput: r.Request.Data.(*DescribeModelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeModelResponse is the response type for the
// DescribeModel API operation.
type DescribeModelResponse struct {
	*DescribeModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeModel request.
func (r *DescribeModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

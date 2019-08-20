// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for AttachVolume.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AttachVolumeRequest
type AttachVolumeInput struct {
	_ struct{} `type:"structure"`

	// The device name (for example, /dev/sdh or xvdh).
	//
	// Device is a required field
	Device *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The ID of the EBS volume. The volume and instance must be within the same
	// Availability Zone.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AttachVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachVolumeInput"}

	if s.Device == nil {
		invalidParams.Add(aws.NewErrParamRequired("Device"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes volume attachment details.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/VolumeAttachment
type AttachVolumeOutput struct {
	_ struct{} `type:"structure"`

	// The time stamp when the attachment initiated.
	AttachTime *time.Time `locationName:"attachTime" type:"timestamp"`

	// Indicates whether the EBS volume is deleted on instance termination.
	DeleteOnTermination *bool `locationName:"deleteOnTermination" type:"boolean"`

	// The device name.
	Device *string `locationName:"device" type:"string"`

	// The ID of the instance.
	InstanceId *string `locationName:"instanceId" type:"string"`

	// The attachment state of the volume.
	State VolumeAttachmentState `locationName:"status" type:"string" enum:"true"`

	// The ID of the volume.
	VolumeId *string `locationName:"volumeId" type:"string"`
}

// String returns the string representation
func (s AttachVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachVolume = "AttachVolume"

// AttachVolumeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Attaches an EBS volume to a running or stopped instance and exposes it to
// the instance with the specified device name.
//
// Encrypted EBS volumes must be attached to instances that support Amazon EBS
// encryption. For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// After you attach an EBS volume, you must make it available. For more information,
// see Making an EBS Volume Available For Use (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-using-volumes.html).
//
// If a volume has an AWS Marketplace product code:
//
//    * The volume can be attached only to a stopped instance.
//
//    * AWS Marketplace product codes are copied from the volume to the instance.
//
//    * You must be subscribed to the product.
//
//    * The instance type and operating system of the instance must support
//    the product. For example, you can't detach a volume from a Windows instance
//    and attach it to a Linux instance.
//
// For more information, see Attaching Amazon EBS Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-attaching-volume.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using AttachVolumeRequest.
//    req := client.AttachVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AttachVolume
func (c *Client) AttachVolumeRequest(input *AttachVolumeInput) AttachVolumeRequest {
	op := &aws.Operation{
		Name:       opAttachVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachVolumeInput{}
	}

	req := c.newRequest(op, input, &AttachVolumeOutput{})
	return AttachVolumeRequest{Request: req, Input: input, Copy: c.AttachVolumeRequest}
}

// AttachVolumeRequest is the request type for the
// AttachVolume API operation.
type AttachVolumeRequest struct {
	*aws.Request
	Input *AttachVolumeInput
	Copy  func(*AttachVolumeInput) AttachVolumeRequest
}

// Send marshals and sends the AttachVolume API request.
func (r AttachVolumeRequest) Send(ctx context.Context) (*AttachVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachVolumeResponse{
		AttachVolumeOutput: r.Request.Data.(*AttachVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachVolumeResponse is the response type for the
// AttachVolume API operation.
type AttachVolumeResponse struct {
	*AttachVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachVolume request.
func (r *AttachVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

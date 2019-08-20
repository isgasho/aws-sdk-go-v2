// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateCapacityReservationRequest
type CreateCapacityReservationInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone in which to create the Capacity Reservation.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `type:"string" required:"true"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	//
	// Constraint: Maximum 64 ASCII characters.
	ClientToken *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	// This optimization provides dedicated throughput to Amazon EBS and an optimized
	// configuration stack to provide optimal I/O performance. This optimization
	// isn't available with all instance types. Additional usage charges apply when
	// using an EBS- optimized instance.
	EbsOptimized *bool `type:"boolean"`

	// The date and time at which the Capacity Reservation expires. When a Capacity
	// Reservation expires, the reserved capacity is released and you can no longer
	// launch instances into it. The Capacity Reservation's state changes to expired
	// when it reaches its end date and time.
	//
	// You must provide an EndDate value if EndDateType is limited. Omit EndDate
	// if EndDateType is unlimited.
	//
	// If the EndDateType is limited, the Capacity Reservation is cancelled within
	// an hour from the specified time. For example, if you specify 5/31/2019, 13:30:55,
	// the Capacity Reservation is guaranteed to end between 13:30:55 and 14:30:55
	// on 5/31/2019.
	EndDate *time.Time `type:"timestamp"`

	// Indicates the way in which the Capacity Reservation ends. A Capacity Reservation
	// can have one of the following end types:
	//
	//    * unlimited - The Capacity Reservation remains active until you explicitly
	//    cancel it. Do not provide an EndDate if the EndDateType is unlimited.
	//
	//    * limited - The Capacity Reservation expires automatically at a specified
	//    date and time. You must provide an EndDate value if the EndDateType value
	//    is limited.
	EndDateType EndDateType `type:"string" enum:"true"`

	// Indicates whether the Capacity Reservation supports instances with temporary,
	// block-level storage.
	EphemeralStorage *bool `type:"boolean"`

	// The number of instances for which to reserve capacity.
	//
	// InstanceCount is a required field
	InstanceCount *int64 `type:"integer" required:"true"`

	// Indicates the type of instance launches that the Capacity Reservation accepts.
	// The options include:
	//
	//    * open - The Capacity Reservation automatically matches all instances
	//    that have matching attributes (instance type, platform, and Availability
	//    Zone). Instances that have matching attributes run in the Capacity Reservation
	//    automatically without specifying any additional parameters.
	//
	//    * targeted - The Capacity Reservation only accepts instances that have
	//    matching attributes (instance type, platform, and Availability Zone),
	//    and explicitly target the Capacity Reservation. This ensures that only
	//    permitted instances can use the reserved capacity.
	//
	// Default: open
	InstanceMatchCriteria InstanceMatchCriteria `type:"string" enum:"true"`

	// The type of operating system for which to reserve capacity.
	//
	// InstancePlatform is a required field
	InstancePlatform CapacityReservationInstancePlatform `type:"string" required:"true" enum:"true"`

	// The instance type for which to reserve capacity. For more information, see
	// Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// InstanceType is a required field
	InstanceType *string `type:"string" required:"true"`

	// The tags to apply to the Capacity Reservation during launch.
	TagSpecifications []TagSpecification `locationNameList:"item" type:"list"`

	// Indicates the tenancy of the Capacity Reservation. A Capacity Reservation
	// can have one of the following tenancy settings:
	//
	//    * default - The Capacity Reservation is created on hardware that is shared
	//    with other AWS accounts.
	//
	//    * dedicated - The Capacity Reservation is created on single-tenant hardware
	//    that is dedicated to a single AWS account.
	Tenancy CapacityReservationTenancy `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateCapacityReservationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCapacityReservationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCapacityReservationInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}

	if s.InstanceCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceCount"))
	}
	if len(s.InstancePlatform) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("InstancePlatform"))
	}

	if s.InstanceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateCapacityReservationResult
type CreateCapacityReservationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Capacity Reservation.
	CapacityReservation *CapacityReservation `locationName:"capacityReservation" type:"structure"`
}

// String returns the string representation
func (s CreateCapacityReservationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCapacityReservation = "CreateCapacityReservation"

// CreateCapacityReservationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a new Capacity Reservation with the specified attributes.
//
// Capacity Reservations enable you to reserve capacity for your Amazon EC2
// instances in a specific Availability Zone for any duration. This gives you
// the flexibility to selectively add capacity reservations and still get the
// Regional RI discounts for that usage. By creating Capacity Reservations,
// you ensure that you always have access to Amazon EC2 capacity when you need
// it, for as long as you need it. For more information, see Capacity Reservations
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// Your request to create a Capacity Reservation could fail if Amazon EC2 does
// not have sufficient capacity to fulfill the request. If your request fails
// due to Amazon EC2 capacity constraints, either try again at a later time,
// try in a different Availability Zone, or request a smaller capacity reservation.
// If your application is flexible across instance types and sizes, try to create
// a Capacity Reservation with different instance attributes.
//
// Your request could also fail if the requested quantity exceeds your On-Demand
// Instance limit for the selected instance type. If your request fails due
// to limit constraints, increase your On-Demand Instance limit for the required
// instance type and try again. For more information about increasing your instance
// limits, see Amazon EC2 Service Limits (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-resource-limits.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using CreateCapacityReservationRequest.
//    req := client.CreateCapacityReservationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateCapacityReservation
func (c *Client) CreateCapacityReservationRequest(input *CreateCapacityReservationInput) CreateCapacityReservationRequest {
	op := &aws.Operation{
		Name:       opCreateCapacityReservation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCapacityReservationInput{}
	}

	req := c.newRequest(op, input, &CreateCapacityReservationOutput{})
	return CreateCapacityReservationRequest{Request: req, Input: input, Copy: c.CreateCapacityReservationRequest}
}

// CreateCapacityReservationRequest is the request type for the
// CreateCapacityReservation API operation.
type CreateCapacityReservationRequest struct {
	*aws.Request
	Input *CreateCapacityReservationInput
	Copy  func(*CreateCapacityReservationInput) CreateCapacityReservationRequest
}

// Send marshals and sends the CreateCapacityReservation API request.
func (r CreateCapacityReservationRequest) Send(ctx context.Context) (*CreateCapacityReservationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCapacityReservationResponse{
		CreateCapacityReservationOutput: r.Request.Data.(*CreateCapacityReservationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCapacityReservationResponse is the response type for the
// CreateCapacityReservation API operation.
type CreateCapacityReservationResponse struct {
	*CreateCapacityReservationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCapacityReservation request.
func (r *CreateCapacityReservationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

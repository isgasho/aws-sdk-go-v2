// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// An accelerator is a complex type that includes one or more listeners that
// process inbound connections and then direct traffic to one or more endpoint
// groups, each of which includes endpoints, such as load balancers.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/Accelerator
type Accelerator struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the accelerator.
	AcceleratorArn *string `type:"string"`

	// The date and time that the accelerator was created.
	CreatedTime *time.Time `type:"timestamp"`

	// Indicates whether theaccelerator is enabled. The value is true or false.
	// The default value is true.
	//
	// If the value is set to true, the accelerator cannot be deleted. If set to
	// false, accelerator can be deleted.
	Enabled *bool `type:"boolean"`

	// The value for the address type must be IPv4.
	IpAddressType IpAddressType `type:"string" enum:"true"`

	// IP address set associated with the accelerator.
	IpSets []IpSet `type:"list"`

	// The date and time that the accelerator was last modified.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The name of the accelerator. The name can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens (-), and must not begin
	// or end with a hyphen.
	Name *string `type:"string"`

	// Describes the deployment status of the accelerator.
	Status AcceleratorStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s Accelerator) String() string {
	return awsutil.Prettify(s)
}

// Attributes of an accelerator.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/AcceleratorAttributes
type AcceleratorAttributes struct {
	_ struct{} `type:"structure"`

	// Indicates whether flow logs are enabled. The default value is false. If the
	// value is true, FlowLogsS3Bucket and FlowLogsS3Prefix must be specified.
	//
	// For more information, see Flow Logs (https://docs.aws.amazon.com/global-accelerator/latest/dg/monitoring-global-accelerator.flow-logs.html)
	// in the AWS Global Accelerator Developer Guide.
	FlowLogsEnabled *bool `type:"boolean"`

	// The name of the Amazon S3 bucket for the flow logs. Attribute is required
	// if FlowLogsEnabled is true. The bucket must exist and have a bucket policy
	// that grants AWS Global Accelerator permission to write to the bucket.
	FlowLogsS3Bucket *string `type:"string"`

	// The prefix for the location in the Amazon S3 bucket for the flow logs. Attribute
	// is required if FlowLogsEnabled is true. If you don’t specify a prefix,
	// the flow logs are stored in the root of the bucket.
	FlowLogsS3Prefix *string `type:"string"`
}

// String returns the string representation
func (s AcceleratorAttributes) String() string {
	return awsutil.Prettify(s)
}

// A complex type for endpoints.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/EndpointConfiguration
type EndpointConfiguration struct {
	_ struct{} `type:"structure"`

	// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application
	// Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If
	// the endpoint is an Elastic IP address, this is the Elastic IP address allocation
	// ID.
	EndpointId *string `type:"string"`

	// The weight associated with the endpoint. When you add weights to endpoints,
	// you configure AWS Global Accelerator to route traffic based on proportions
	// that you specify. For example, you might specify endpoint weights of 4, 5,
	// 5, and 6 (sum=20). The result is that 4/20 of your traffic, on average, is
	// routed to the first endpoint, 5/20 is routed both to the second and third
	// endpoints, and 6/20 is routed to the last endpoint. For more information,
	// see Endpoint Weights (https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html)
	// in the AWS Global Accelerator Developer Guide.
	Weight *int64 `type:"integer"`
}

// String returns the string representation
func (s EndpointConfiguration) String() string {
	return awsutil.Prettify(s)
}

// A complex type for an endpoint. Each endpoint group can include one or more
// endpoints, such as load balancers.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/EndpointDescription
type EndpointDescription struct {
	_ struct{} `type:"structure"`

	// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application
	// Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If
	// the endpoint is an Elastic IP address, this is the Elastic IP address allocation
	// ID.
	EndpointId *string `type:"string"`

	// The reason code associated with why the endpoint is not healthy. If the endpoint
	// state is healthy, a reason code is not provided.
	//
	// If the endpoint state is unhealthy, the reason code can be one of the following
	// values:
	//
	//    * Timeout: The health check requests to the endpoint are timing out before
	//    returning a status.
	//
	//    * Failed: The health check failed, for example because the endpoint response
	//    was invalid (malformed).
	//
	// If the endpoint state is initial, the reason code can be one of the following
	// values:
	//
	//    * ProvisioningInProgress: The endpoint is in the process of being provisioned.
	//
	//    * InitialHealthChecking: Global Accelerator is still setting up the minimum
	//    number of health checks for the endpoint that are required to determine
	//    its health status.
	HealthReason *string `type:"string"`

	// The health status of the endpoint.
	HealthState HealthState `type:"string" enum:"true"`

	// The weight associated with the endpoint. When you add weights to endpoints,
	// you configure AWS Global Accelerator to route traffic based on proportions
	// that you specify. For example, you might specify endpoint weights of 4, 5,
	// 5, and 6 (sum=20). The result is that 4/20 of your traffic, on average, is
	// routed to the first endpoint, 5/20 is routed both to the second and third
	// endpoints, and 6/20 is routed to the last endpoint. For more information,
	// see Endpoint Weights (https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html)
	// in the AWS Global Accelerator Developer Guide.
	Weight *int64 `type:"integer"`
}

// String returns the string representation
func (s EndpointDescription) String() string {
	return awsutil.Prettify(s)
}

// A complex type for the endpoint group. An AWS Region can have only one endpoint
// group for a specific listener.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/EndpointGroup
type EndpointGroup struct {
	_ struct{} `type:"structure"`

	// The list of endpoint objects.
	EndpointDescriptions []EndpointDescription `type:"list"`

	// The Amazon Resource Name (ARN) of the endpoint group.
	EndpointGroupArn *string `type:"string"`

	// The AWS Region that this endpoint group belongs.
	EndpointGroupRegion *string `type:"string"`

	// The time—10 seconds or 30 seconds—between health checks for each endpoint.
	// The default value is 30.
	HealthCheckIntervalSeconds *int64 `min:"10" type:"integer"`

	// If the protocol is HTTP/S, then this value provides the ping path that Global
	// Accelerator uses for the destination on the endpoints for health checks.
	// The default is slash (/).
	HealthCheckPath *string `type:"string"`

	// The port that Global Accelerator uses to perform health checks on endpoints
	// that are part of this endpoint group.
	//
	// The default port is the port for the listener that this endpoint group is
	// associated with. If the listener port is a list, Global Accelerator uses
	// the first specified port in the list of ports.
	HealthCheckPort *int64 `min:"1" type:"integer"`

	// The protocol that Global Accelerator uses to perform health checks on endpoints
	// that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol HealthCheckProtocol `type:"string" enum:"true"`

	// The number of consecutive health checks required to set the state of a healthy
	// endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default
	// value is 3.
	ThresholdCount *int64 `min:"1" type:"integer"`

	// The percentage of traffic to send to an AWS Region. Additional traffic is
	// distributed to other endpoint groups for this listener.
	//
	// Use this action to increase (dial up) or decrease (dial down) traffic to
	// a specific Region. The percentage is applied to the traffic that would otherwise
	// have been routed to the Region based on optimal routing.
	//
	// The default value is 100.
	TrafficDialPercentage *float64 `type:"float"`
}

// String returns the string representation
func (s EndpointGroup) String() string {
	return awsutil.Prettify(s)
}

// A complex type for the set of IP addresses for an accelerator.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/IpSet
type IpSet struct {
	_ struct{} `type:"structure"`

	// The array of IP addresses in the IP address set. An IP address set can have
	// a maximum of two IP addresses.
	IpAddresses []string `type:"list"`

	// The types of IP addresses included in this IP set.
	IpFamily *string `type:"string"`
}

// String returns the string representation
func (s IpSet) String() string {
	return awsutil.Prettify(s)
}

// A complex type for a listener.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/Listener
type Listener struct {
	_ struct{} `type:"structure"`

	// Client affinity lets you direct all requests from a user to the same endpoint,
	// if you have stateful applications, regardless of the port and protocol of
	// the client request. Clienty affinity gives you control over whether to always
	// route each client to the same specific endpoint.
	//
	// AWS Global Accelerator uses a consistent-flow hashing algorithm to choose
	// the optimal endpoint for a connection. If client affinity is NONE, Global
	// Accelerator uses the "five-tuple" (5-tuple) properties—source IP address,
	// source port, destination IP address, destination port, and protocol—to
	// select the hash value, and then chooses the best endpoint. However, with
	// this setting, if someone uses different ports to connect to Global Accelerator,
	// their connections might not be always routed to the same endpoint because
	// the hash value changes.
	//
	// If you want a given client to always be routed to the same endpoint, set
	// client affinity to SOURCE_IP instead. When you use the SOURCE_IP setting,
	// Global Accelerator uses the "two-tuple" (2-tuple) properties— source (client)
	// IP address and destination IP address—to select the hash value.
	//
	// The default value is NONE.
	ClientAffinity Affinity `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `type:"string"`

	// The list of port ranges for the connections from clients to the accelerator.
	PortRanges []PortRange `min:"1" type:"list"`

	// The protocol for the connections from clients to the accelerator.
	Protocol Protocol `type:"string" enum:"true"`
}

// String returns the string representation
func (s Listener) String() string {
	return awsutil.Prettify(s)
}

// A complex type for a range of ports for a listener.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/PortRange
type PortRange struct {
	_ struct{} `type:"structure"`

	// The first port in the range of ports, inclusive.
	FromPort *int64 `min:"1" type:"integer"`

	// The last port in the range of ports, inclusive.
	ToPort *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s PortRange) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PortRange) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PortRange"}
	if s.FromPort != nil && *s.FromPort < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("FromPort", 1))
	}
	if s.ToPort != nil && *s.ToPort < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ToPort", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Represents a CloudWatch alarm associated with a scaling policy.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/Alarm
type Alarm struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the alarm.
	//
	// AlarmARN is a required field
	AlarmARN *string `type:"string" required:"true"`

	// The name of the alarm.
	//
	// AlarmName is a required field
	AlarmName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Alarm) String() string {
	return awsutil.Prettify(s)
}

// Represents a CloudWatch metric of your choosing for a target tracking scaling
// policy to use with Application Auto Scaling.
//
// To create your customized metric specification:
//
//    * Add values for each required parameter from CloudWatch. You can use
//    an existing metric, or a new metric that you create. To use your own metric,
//    you must first publish the metric to CloudWatch. For more information,
//    see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)
//    in the Amazon CloudWatch User Guide.
//
//    * Choose a metric that changes proportionally with capacity. The value
//    of the metric should increase or decrease in inverse proportion to the
//    number of capacity units. That is, the value of the metric should decrease
//    when capacity increases.
//
// For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/CustomizedMetricSpecification
type CustomizedMetricSpecification struct {
	_ struct{} `type:"structure"`

	// The dimensions of the metric.
	//
	// Conditional: If you published your metric with dimensions, you must specify
	// the same dimensions in your scaling policy.
	Dimensions []MetricDimension `type:"list"`

	// The name of the metric.
	//
	// MetricName is a required field
	MetricName *string `type:"string" required:"true"`

	// The namespace of the metric.
	//
	// Namespace is a required field
	Namespace *string `type:"string" required:"true"`

	// The statistic of the metric.
	//
	// Statistic is a required field
	Statistic MetricStatistic `type:"string" required:"true" enum:"true"`

	// The unit of the metric.
	Unit *string `type:"string"`
}

// String returns the string representation
func (s CustomizedMetricSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CustomizedMetricSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CustomizedMetricSpecification"}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}
	if len(s.Statistic) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Statistic"))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the dimension names and values associated with a metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/MetricDimension
type MetricDimension struct {
	_ struct{} `type:"structure"`

	// The name of the dimension.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The value of the dimension.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s MetricDimension) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricDimension) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricDimension"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a predefined metric for a target tracking scaling policy to use
// with Application Auto Scaling.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/PredefinedMetricSpecification
type PredefinedMetricSpecification struct {
	_ struct{} `type:"structure"`

	// The metric type. The ALBRequestCountPerTarget metric type applies only to
	// Spot fleet requests and ECS services.
	//
	// PredefinedMetricType is a required field
	PredefinedMetricType MetricType `type:"string" required:"true" enum:"true"`

	// Identifies the resource associated with the metric type. You can't specify
	// a resource label unless the metric type is ALBRequestCountPerTarget and there
	// is a target group attached to the Spot fleet request or ECS service.
	//
	// The format is app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>,
	// where:
	//
	//    * app/<load-balancer-name>/<load-balancer-id> is the final portion of
	//    the load balancer ARN
	//
	//    * targetgroup/<target-group-name>/<target-group-id> is the final portion
	//    of the target group ARN.
	ResourceLabel *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PredefinedMetricSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PredefinedMetricSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PredefinedMetricSpecification"}
	if len(s.PredefinedMetricType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PredefinedMetricType"))
	}
	if s.ResourceLabel != nil && len(*s.ResourceLabel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceLabel", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a scalable target.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/ScalableTarget
type ScalableTarget struct {
	_ struct{} `type:"structure"`

	// The Unix timestamp for when the scalable target was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// The maximum value to scale to in response to a scale-out event.
	//
	// MaxCapacity is a required field
	MaxCapacity *int64 `type:"integer" required:"true"`

	// The minimum value to scale to in response to a scale-in event.
	//
	// MinCapacity is a required field
	MinCapacity *int64 `type:"integer" required:"true"`

	// The identifier of the resource associated with the scalable target. This
	// string consists of the resource type and unique identifier.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variants - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The ARN of an IAM role that allows Application Auto Scaling to modify the
	// scalable target on your behalf.
	//
	// RoleARN is a required field
	RoleARN *string `min:"1" type:"string" required:"true"`

	// The scalable dimension associated with the scalable target. This string consists
	// of the service namespace, resource type, and scaling property.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	//
	// ScalableDimension is a required field
	ScalableDimension ScalableDimension `type:"string" required:"true" enum:"true"`

	// The namespace of the AWS service that provides the resource or custom-resource
	// for a resource provided by your own application or service. For more information,
	// see AWS Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the Amazon Web Services General Reference.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ScalableTarget) String() string {
	return awsutil.Prettify(s)
}

// Represents the minimum and maximum capacity for a scheduled action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/ScalableTargetAction
type ScalableTargetAction struct {
	_ struct{} `type:"structure"`

	// The maximum capacity.
	MaxCapacity *int64 `type:"integer"`

	// The minimum capacity.
	MinCapacity *int64 `type:"integer"`
}

// String returns the string representation
func (s ScalableTargetAction) String() string {
	return awsutil.Prettify(s)
}

// Represents a scaling activity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/ScalingActivity
type ScalingActivity struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the scaling activity.
	//
	// ActivityId is a required field
	ActivityId *string `type:"string" required:"true"`

	// A simple description of what caused the scaling activity to happen.
	//
	// Cause is a required field
	Cause *string `type:"string" required:"true"`

	// A simple description of what action the scaling activity intends to accomplish.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The details about the scaling activity.
	Details *string `type:"string"`

	// The Unix timestamp for when the scaling activity ended.
	EndTime *time.Time `type:"timestamp"`

	// The identifier of the resource associated with the scaling activity. This
	// string consists of the resource type and unique identifier.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variants - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	//
	// ScalableDimension is a required field
	ScalableDimension ScalableDimension `type:"string" required:"true" enum:"true"`

	// The namespace of the AWS service that provides the resource or custom-resource
	// for a resource provided by your own application or service. For more information,
	// see AWS Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the Amazon Web Services General Reference.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`

	// The Unix timestamp for when the scaling activity began.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" required:"true"`

	// Indicates the status of the scaling activity.
	//
	// StatusCode is a required field
	StatusCode ScalingActivityStatusCode `type:"string" required:"true" enum:"true"`

	// A simple message about the current status of the scaling activity.
	StatusMessage *string `type:"string"`
}

// String returns the string representation
func (s ScalingActivity) String() string {
	return awsutil.Prettify(s)
}

// Represents a scaling policy to use with Application Auto Scaling.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/ScalingPolicy
type ScalingPolicy struct {
	_ struct{} `type:"structure"`

	// The CloudWatch alarms associated with the scaling policy.
	Alarms []Alarm `type:"list"`

	// The Unix timestamp for when the scaling policy was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// The Amazon Resource Name (ARN) of the scaling policy.
	//
	// PolicyARN is a required field
	PolicyARN *string `min:"1" type:"string" required:"true"`

	// The name of the scaling policy.
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`

	// The scaling policy type.
	//
	// PolicyType is a required field
	PolicyType PolicyType `type:"string" required:"true" enum:"true"`

	// The identifier of the resource associated with the scaling policy. This string
	// consists of the resource type and unique identifier.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variants - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	//
	// ScalableDimension is a required field
	ScalableDimension ScalableDimension `type:"string" required:"true" enum:"true"`

	// The namespace of the AWS service that provides the resource or custom-resource
	// for a resource provided by your own application or service. For more information,
	// see AWS Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the Amazon Web Services General Reference.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`

	// A step scaling policy.
	StepScalingPolicyConfiguration *StepScalingPolicyConfiguration `type:"structure"`

	// A target tracking scaling policy.
	TargetTrackingScalingPolicyConfiguration *TargetTrackingScalingPolicyConfiguration `type:"structure"`
}

// String returns the string representation
func (s ScalingPolicy) String() string {
	return awsutil.Prettify(s)
}

// Represents a scheduled action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/ScheduledAction
type ScheduledAction struct {
	_ struct{} `type:"structure"`

	// The date and time that the scheduled action was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// The date and time that the action is scheduled to end.
	EndTime *time.Time `type:"timestamp"`

	// The identifier of the resource associated with the scaling policy. This string
	// consists of the resource type and unique identifier.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variants - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	ScalableDimension ScalableDimension `type:"string" enum:"true"`

	// The new minimum and maximum capacity. You can set both values or just one.
	// During the scheduled time, if the current capacity is below the minimum capacity,
	// Application Auto Scaling scales out to the minimum capacity. If the current
	// capacity is above the maximum capacity, Application Auto Scaling scales in
	// to the maximum capacity.
	ScalableTargetAction *ScalableTargetAction `type:"structure"`

	// The schedule for this action. The following formats are supported:
	//
	//    * At expressions - "at(yyyy-mm-ddThh:mm:ss)"
	//
	//    * Rate expressions - "rate(value unit)"
	//
	//    * Cron expressions - "cron(fields)"
	//
	// At expressions are useful for one-time schedules. Specify the time, in UTC.
	//
	// For rate expressions, value is a positive integer and unit is minute | minutes
	// | hour | hours | day | days.
	//
	// For more information about cron expressions, see Cron Expressions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions)
	// in the Amazon CloudWatch Events User Guide.
	//
	// Schedule is a required field
	Schedule *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the scheduled action.
	//
	// ScheduledActionARN is a required field
	ScheduledActionARN *string `min:"1" type:"string" required:"true"`

	// The name of the scheduled action.
	//
	// ScheduledActionName is a required field
	ScheduledActionName *string `min:"1" type:"string" required:"true"`

	// The namespace of the AWS service that provides the resource or custom-resource
	// for a resource provided by your own application or service. For more information,
	// see AWS Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the Amazon Web Services General Reference.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`

	// The date and time that the action is scheduled to begin.
	StartTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ScheduledAction) String() string {
	return awsutil.Prettify(s)
}

// Represents a step adjustment for a StepScalingPolicyConfiguration. Describes
// an adjustment based on the difference between the value of the aggregated
// CloudWatch metric and the breach threshold that you've defined for the alarm.
//
// For the following examples, suppose that you have an alarm with a breach
// threshold of 50:
//
//    * To trigger the adjustment when the metric is greater than or equal to
//    50 and less than 60, specify a lower bound of 0 and an upper bound of
//    10.
//
//    * To trigger the adjustment when the metric is greater than 40 and less
//    than or equal to 50, specify a lower bound of -10 and an upper bound of
//    0.
//
// There are a few rules for the step adjustments for your step policy:
//
//    * The ranges of your step adjustments can't overlap or have a gap.
//
//    * At most one step adjustment can have a null lower bound. If one step
//    adjustment has a negative lower bound, then there must be a step adjustment
//    with a null lower bound.
//
//    * At most one step adjustment can have a null upper bound. If one step
//    adjustment has a positive upper bound, then there must be a step adjustment
//    with a null upper bound.
//
//    * The upper and lower bound can't be null in the same step adjustment.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/StepAdjustment
type StepAdjustment struct {
	_ struct{} `type:"structure"`

	// The lower bound for the difference between the alarm threshold and the CloudWatch
	// metric. If the metric value is above the breach threshold, the lower bound
	// is inclusive (the metric must be greater than or equal to the threshold plus
	// the lower bound). Otherwise, it is exclusive (the metric must be greater
	// than the threshold plus the lower bound). A null value indicates negative
	// infinity.
	MetricIntervalLowerBound *float64 `type:"double"`

	// The upper bound for the difference between the alarm threshold and the CloudWatch
	// metric. If the metric value is above the breach threshold, the upper bound
	// is exclusive (the metric must be less than the threshold plus the upper bound).
	// Otherwise, it is inclusive (the metric must be less than or equal to the
	// threshold plus the upper bound). A null value indicates positive infinity.
	//
	// The upper bound must be greater than the lower bound.
	MetricIntervalUpperBound *float64 `type:"double"`

	// The amount by which to scale, based on the specified adjustment type. A positive
	// value adds to the current scalable dimension while a negative number removes
	// from the current scalable dimension.
	//
	// ScalingAdjustment is a required field
	ScalingAdjustment *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s StepAdjustment) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StepAdjustment) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StepAdjustment"}

	if s.ScalingAdjustment == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScalingAdjustment"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a step scaling policy configuration to use with Application Auto
// Scaling.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/StepScalingPolicyConfiguration
type StepScalingPolicyConfiguration struct {
	_ struct{} `type:"structure"`

	// The adjustment type, which specifies how the ScalingAdjustment parameter
	// in a StepAdjustment is interpreted.
	AdjustmentType AdjustmentType `type:"string" enum:"true"`

	// The amount of time, in seconds, after a scaling activity completes where
	// previous trigger-related scaling activities can influence future scaling
	// events.
	//
	// For scale-out policies, while the cooldown period is in effect, the capacity
	// that has been added by the previous scale-out event that initiated the cooldown
	// is calculated as part of the desired capacity for the next scale out. The
	// intention is to continuously (but not excessively) scale out. For example,
	// an alarm triggers a step scaling policy to scale out an Amazon ECS service
	// by 2 tasks, the scaling activity completes successfully, and a cooldown period
	// of 5 minutes starts. During the cooldown period, if the alarm triggers the
	// same policy again but at a more aggressive step adjustment to scale out the
	// service by 3 tasks, the 2 tasks that were added in the previous scale-out
	// event are considered part of that capacity and only 1 additional task is
	// added to the desired count.
	//
	// For scale-in policies, the cooldown period is used to block subsequent scale-in
	// requests until it has expired. The intention is to scale in conservatively
	// to protect your application's availability. However, if another alarm triggers
	// a scale-out policy during the cooldown period after a scale-in, Application
	// Auto Scaling scales out your scalable target immediately.
	Cooldown *int64 `type:"integer"`

	// The aggregation type for the CloudWatch metrics. Valid values are Minimum,
	// Maximum, and Average. If the aggregation type is null, the value is treated
	// as Average.
	MetricAggregationType MetricAggregationType `type:"string" enum:"true"`

	// The minimum number to adjust your scalable dimension as a result of a scaling
	// activity. If the adjustment type is PercentChangeInCapacity, the scaling
	// policy changes the scalable dimension of the scalable target by this amount.
	//
	// For example, suppose that you create a step scaling policy to scale out an
	// Amazon ECS service by 25 percent and you specify a MinAdjustmentMagnitude
	// of 2. If the service has 4 tasks and the scaling policy is performed, 25
	// percent of 4 is 1. However, because you specified a MinAdjustmentMagnitude
	// of 2, Application Auto Scaling scales out the service by 2 tasks.
	MinAdjustmentMagnitude *int64 `type:"integer"`

	// A set of adjustments that enable you to scale based on the size of the alarm
	// breach.
	StepAdjustments []StepAdjustment `type:"list"`
}

// String returns the string representation
func (s StepScalingPolicyConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StepScalingPolicyConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StepScalingPolicyConfiguration"}
	if s.StepAdjustments != nil {
		for i, v := range s.StepAdjustments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "StepAdjustments", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a target tracking scaling policy configuration to use with Application
// Auto Scaling.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/TargetTrackingScalingPolicyConfiguration
type TargetTrackingScalingPolicyConfiguration struct {
	_ struct{} `type:"structure"`

	// A customized metric. You can specify either a predefined metric or a customized
	// metric.
	CustomizedMetricSpecification *CustomizedMetricSpecification `type:"structure"`

	// Indicates whether scale in by the target tracking scaling policy is disabled.
	// If the value is true, scale in is disabled and the target tracking scaling
	// policy won't remove capacity from the scalable resource. Otherwise, scale
	// in is enabled and the target tracking scaling policy can remove capacity
	// from the scalable resource. The default value is false.
	DisableScaleIn *bool `type:"boolean"`

	// A predefined metric. You can specify either a predefined metric or a customized
	// metric.
	PredefinedMetricSpecification *PredefinedMetricSpecification `type:"structure"`

	// The amount of time, in seconds, after a scale-in activity completes before
	// another scale in activity can start.
	//
	// The cooldown period is used to block subsequent scale-in requests until it
	// has expired. The intention is to scale in conservatively to protect your
	// application's availability. However, if another alarm triggers a scale-out
	// policy during the cooldown period after a scale-in, Application Auto Scaling
	// scales out your scalable target immediately.
	ScaleInCooldown *int64 `type:"integer"`

	// The amount of time, in seconds, after a scale-out activity completes before
	// another scale-out activity can start.
	//
	// While the cooldown period is in effect, the capacity that has been added
	// by the previous scale-out event that initiated the cooldown is calculated
	// as part of the desired capacity for the next scale out. The intention is
	// to continuously (but not excessively) scale out.
	ScaleOutCooldown *int64 `type:"integer"`

	// The target value for the metric. The range is 8.515920e-109 to 1.174271e+108
	// (Base 10) or 2e-360 to 2e360 (Base 2).
	//
	// TargetValue is a required field
	TargetValue *float64 `type:"double" required:"true"`
}

// String returns the string representation
func (s TargetTrackingScalingPolicyConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TargetTrackingScalingPolicyConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TargetTrackingScalingPolicyConfiguration"}

	if s.TargetValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetValue"))
	}
	if s.CustomizedMetricSpecification != nil {
		if err := s.CustomizedMetricSpecification.Validate(); err != nil {
			invalidParams.AddNested("CustomizedMetricSpecification", err.(aws.ErrInvalidParams))
		}
	}
	if s.PredefinedMetricSpecification != nil {
		if err := s.PredefinedMetricSpecification.Validate(); err != nil {
			invalidParams.AddNested("PredefinedMetricSpecification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

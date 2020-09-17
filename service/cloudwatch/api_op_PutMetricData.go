// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Publishes metric data points to Amazon CloudWatch. CloudWatch associates the
// data points with the specified metric. If the specified metric does not exist,
// CloudWatch creates the metric. When CloudWatch creates a metric, it can take up
// to fifteen minutes for the metric to appear in calls to ListMetrics
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html).
// <p>You can publish either individual data points in the <code>Value</code>
// field, or arrays of values and the number of times each value occurred during
// the period by using the <code>Values</code> and <code>Counts</code> fields in
// the <code>MetricDatum</code> structure. Using the <code>Values</code> and
// <code>Counts</code> method enables you to publish up to 150 values per metric
// with one <code>PutMetricData</code> request, and supports retrieving percentile
// statistics on this data.</p> <p>Each <code>PutMetricData</code> request is
// limited to 40 KB in size for HTTP POST requests. You can send a payload
// compressed by gzip. Each request is also limited to no more than 20 different
// metrics.</p> <p>Although the <code>Value</code> parameter accepts numbers of
// type <code>Double</code>, CloudWatch rejects values that are either too small or
// too large. Values must be in the range of -2^360 to 2^360. In addition, special
// values (for example, NaN, +Infinity, -Infinity) are not supported.</p> <p>You
// can use up to 10 dimensions per metric to further clarify what data the metric
// collects. Each dimension consists of a Name and Value pair. For more information
// about specifying dimensions, see <a
// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html">Publishing
// Metrics</a> in the <i>Amazon CloudWatch User Guide</i>.</p> <p>Data points with
// time stamps from 24 hours ago or longer can take at least 48 hours to become
// available for <a
// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html">GetMetricData</a>
// or <a
// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html">GetMetricStatistics</a>
// from the time they are submitted. Data points with time stamps between 3 and 24
// hours ago can take as much as 2 hours to become available for for <a
// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html">GetMetricData</a>
// or <a
// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html">GetMetricStatistics</a>.</p>
// <p>CloudWatch needs raw data points to calculate percentile statistics. If you
// publish data using a statistic set instead, you can only retrieve percentile
// statistics for this data if one of the following conditions is true:</p> <ul>
// <li> <p>The <code>SampleCount</code> value of the statistic set is 1 and
// <code>Min</code>, <code>Max</code>, and <code>Sum</code> are all equal.</p>
// </li> <li> <p>The <code>Min</code> and <code>Max</code> are equal, and
// <code>Sum</code> is equal to <code>Min</code> multiplied by
// <code>SampleCount</code>.</p> </li> </ul>
func (c *Client) PutMetricData(ctx context.Context, params *PutMetricDataInput, optFns ...func(*Options)) (*PutMetricDataOutput, error) {
	stack := middleware.NewStack("PutMetricData", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPutMetricDataMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutMetricDataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricData(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "PutMetricData",
			Err:           err,
		}
	}
	out := result.(*PutMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricDataInput struct {
	// The namespace for the metric data. To avoid conflicts with AWS service
	// namespaces, you should not specify a namespace that begins with AWS/
	Namespace *string
	// The data for the metric. The array can include no more than 20 metrics per call.
	MetricData []*types.MetricDatum
}

type PutMetricDataOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPutMetricDataMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPutMetricData{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPutMetricData{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutMetricData(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "PutMetricData",
	}
}
// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the status of the specified volumes. Volume status provides the result
// of the checks performed on your volumes to determine events that can impair the
// performance of your volumes. The performance of a volume can be affected if an
// issue occurs on the volume's underlying host. If the volume's underlying host
// experiences a power outage or system issue, after the system is restored, there
// could be data inconsistencies on the volume. Volume events notify you if this
// occurs. Volume actions notify you if any action needs to be taken in response to
// the event. The DescribeVolumeStatus operation provides the following information
// about the specified volumes: Status: Reflects the current status of the volume.
// The possible values are ok, impaired , warning, or insufficient-data. If all
// checks pass, the overall status of the volume is ok. If the check fails, the
// overall status is impaired. If the status is insufficient-data, then the checks
// may still be taking place on your volume at the time. We recommend that you
// retry the request. For more information about volume status, see Monitoring the
// Status of Your Volumes
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-status.html)
// in the Amazon Elastic Compute Cloud User Guide. Events: Reflect the cause of a
// volume status and may require you to take action. For example, if your volume
// returns an impaired status, then the volume event might be
// potential-data-inconsistency. This means that your volume has been affected by
// an issue with the underlying host, has all I/O operations disabled, and may have
// inconsistent data. Actions: Reflect the actions you may have to take in response
// to an event. For example, if the status of the volume is impaired and the volume
// event shows potential-data-inconsistency, then the action shows
// enable-volume-io. This means that you may want to enable the I/O operations for
// the volume by calling the EnableVolumeIO action and then check the volume for
// data consistency. Volume status is based on the volume status checks, and does
// not reflect the volume state. Therefore, volume status does not indicate volumes
// in the error state (for example, when a volume is incapable of accepting I/O.)
func (c *Client) DescribeVolumeStatus(ctx context.Context, params *DescribeVolumeStatusInput, optFns ...func(*Options)) (*DescribeVolumeStatusOutput, error) {
	if params == nil {
		params = &DescribeVolumeStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVolumeStatus", params, optFns, addOperationDescribeVolumeStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVolumeStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumeStatusInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	//
	//     * action.code - The action code for the event (for example,
	// enable-volume-io).
	//
	//     * action.description - A description of the action.
	//
	//
	// * action.event-id - The event ID associated with the action.
	//
	//     *
	// availability-zone - The Availability Zone of the instance.
	//
	//     *
	// event.description - A description of the event.
	//
	//     * event.event-id - The
	// event ID.
	//
	//     * event.event-type - The event type (for io-enabled: passed |
	// failed; for io-performance: io-performance:degraded |
	// io-performance:severely-degraded | io-performance:stalled).
	//
	//     *
	// event.not-after - The latest end time for the event.
	//
	//     * event.not-before -
	// The earliest start time for the event.
	//
	//     * volume-status.details-name - The
	// cause for volume-status.status (io-enabled | io-performance).
	//
	//     *
	// volume-status.details-status - The status of volume-status.details-name (for
	// io-enabled: passed | failed; for io-performance: normal | degraded |
	// severely-degraded | stalled).
	//
	//     * volume-status.status - The status of the
	// volume (ok | impaired | warning | insufficient-data).
	Filters []*types.Filter

	// The maximum number of volume results returned by DescribeVolumeStatus in
	// paginated output. When this parameter is used, the request only returns
	// MaxResults results in a single page along with a NextToken response element. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1000; if
	// MaxResults is given a value larger than 1000, only 1000 results are returned. If
	// this parameter is not used, then DescribeVolumeStatus returns all results. You
	// cannot specify this parameter and the volume IDs parameter in the same request.
	MaxResults *int32

	// The NextToken value to include in a future DescribeVolumeStatus request. When
	// the results of the request exceed MaxResults, this value can be used to retrieve
	// the next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// The IDs of the volumes. Default: Describes all your volumes.
	VolumeIds []*string
}

type DescribeVolumeStatusOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the status of the volumes.
	VolumeStatuses []*types.VolumeStatusItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVolumeStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVolumeStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVolumeStatus{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVolumeStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeVolumeStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVolumeStatus",
	}
}

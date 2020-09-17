// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Modifies the maintenance settings of a cluster.
func (c *Client) ModifyClusterMaintenance(ctx context.Context, params *ModifyClusterMaintenanceInput, optFns ...func(*Options)) (*ModifyClusterMaintenanceOutput, error) {
	stack := middleware.NewStack("ModifyClusterMaintenance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyClusterMaintenanceMiddlewares(stack)
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
	addOpModifyClusterMaintenanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterMaintenance(options.Region), middleware.Before)

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
			OperationName: "ModifyClusterMaintenance",
			Err:           err,
		}
	}
	out := result.(*ModifyClusterMaintenanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterMaintenanceInput struct {
	// A timestamp indicating end time for the deferred maintenance window. If you
	// specify an end time, you can't specify a duration.
	DeferMaintenanceEndTime *time.Time
	// A timestamp indicating the start time for the deferred maintenance window.
	DeferMaintenanceStartTime *time.Time
	// A boolean indicating whether to enable the deferred maintenance window.
	DeferMaintenance *bool
	// A unique identifier for the cluster.
	ClusterIdentifier *string
	// An integer indicating the duration of the maintenance window in days. If you
	// specify a duration, you can't specify an end time. The duration must be 45 days
	// or less.
	DeferMaintenanceDuration *int32
	// A unique identifier for the deferred maintenance window.
	DeferMaintenanceIdentifier *string
}

type ModifyClusterMaintenanceOutput struct {
	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyClusterMaintenanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterMaintenance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterMaintenance{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyClusterMaintenance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterMaintenance",
	}
}
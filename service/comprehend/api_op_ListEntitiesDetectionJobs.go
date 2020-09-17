// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the entity detection jobs that you have submitted.
func (c *Client) ListEntitiesDetectionJobs(ctx context.Context, params *ListEntitiesDetectionJobsInput, optFns ...func(*Options)) (*ListEntitiesDetectionJobsOutput, error) {
	stack := middleware.NewStack("ListEntitiesDetectionJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListEntitiesDetectionJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitiesDetectionJobs(options.Region), middleware.Before)

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
			OperationName: "ListEntitiesDetectionJobs",
			Err:           err,
		}
	}
	out := result.(*ListEntitiesDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitiesDetectionJobsInput struct {
	// Filters the jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.EntitiesDetectionJobFilter
	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32
	// Identifies the next page of results to return.
	NextToken *string
}

type ListEntitiesDetectionJobsOutput struct {
	// A list containing the properties of each job that is returned.
	EntitiesDetectionJobPropertiesList []*types.EntitiesDetectionJobProperties
	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListEntitiesDetectionJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListEntitiesDetectionJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntitiesDetectionJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEntitiesDetectionJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListEntitiesDetectionJobs",
	}
}
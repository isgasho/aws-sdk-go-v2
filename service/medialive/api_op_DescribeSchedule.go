// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get a channel schedule
func (c *Client) DescribeSchedule(ctx context.Context, params *DescribeScheduleInput, optFns ...func(*Options)) (*DescribeScheduleOutput, error) {
	stack := middleware.NewStack("DescribeSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeScheduleMiddlewares(stack)
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
	addOpDescribeScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSchedule(options.Region), middleware.Before)

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
			OperationName: "DescribeSchedule",
			Err:           err,
		}
	}
	out := result.(*DescribeScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeScheduleRequest
type DescribeScheduleInput struct {
	// Placeholder documentation for __string
	NextToken *string
	// Id of the channel whose schedule is being updated.
	ChannelId *string
	// Placeholder documentation for MaxResults
	MaxResults *int32
}

// Placeholder documentation for DescribeScheduleResponse
type DescribeScheduleOutput struct {
	// The next token; for use in pagination.
	NextToken *string
	// The list of actions in the schedule.
	ScheduleActions []*types.ScheduleAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeSchedule",
	}
}
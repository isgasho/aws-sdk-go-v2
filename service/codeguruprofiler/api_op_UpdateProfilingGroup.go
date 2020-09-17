// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a profiling group.
func (c *Client) UpdateProfilingGroup(ctx context.Context, params *UpdateProfilingGroupInput, optFns ...func(*Options)) (*UpdateProfilingGroupOutput, error) {
	stack := middleware.NewStack("UpdateProfilingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateProfilingGroupMiddlewares(stack)
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
	addOpUpdateProfilingGroupValidationMiddleware(stack)

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
			OperationName: "UpdateProfilingGroup",
			Err:           err,
		}
	}
	out := result.(*UpdateProfilingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the updateProfilingGroupRequest.
type UpdateProfilingGroupInput struct {
	//
	AgentOrchestrationConfig *types.AgentOrchestrationConfig
	// The name of the profiling group to update.
	ProfilingGroupName *string
}

// The structure representing the updateProfilingGroupResponse.
type UpdateProfilingGroupOutput struct {
	// Updated information about the profiling group.
	ProfilingGroup *types.ProfilingGroupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateProfilingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProfilingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProfilingGroup{}, middleware.After)
}
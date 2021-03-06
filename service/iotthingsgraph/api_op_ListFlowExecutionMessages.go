// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of objects that contain information about events in a flow
// execution.
func (c *Client) ListFlowExecutionMessages(ctx context.Context, params *ListFlowExecutionMessagesInput, optFns ...func(*Options)) (*ListFlowExecutionMessagesOutput, error) {
	if params == nil {
		params = &ListFlowExecutionMessagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFlowExecutionMessages", params, optFns, addOperationListFlowExecutionMessagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFlowExecutionMessagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlowExecutionMessagesInput struct {

	// The ID of the flow execution.
	//
	// This member is required.
	FlowExecutionId *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
}

type ListFlowExecutionMessagesOutput struct {

	// A list of objects that contain information about events in the specified flow
	// execution.
	Messages []*types.FlowExecutionMessage

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFlowExecutionMessagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFlowExecutionMessages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFlowExecutionMessages{}, middleware.After)
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
	addOpListFlowExecutionMessagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFlowExecutionMessages(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListFlowExecutionMessages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "ListFlowExecutionMessages",
	}
}

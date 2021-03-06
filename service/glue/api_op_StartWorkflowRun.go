// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a new run of the specified workflow.
func (c *Client) StartWorkflowRun(ctx context.Context, params *StartWorkflowRunInput, optFns ...func(*Options)) (*StartWorkflowRunOutput, error) {
	if params == nil {
		params = &StartWorkflowRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartWorkflowRun", params, optFns, addOperationStartWorkflowRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartWorkflowRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartWorkflowRunInput struct {

	// The name of the workflow to start.
	//
	// This member is required.
	Name *string
}

type StartWorkflowRunOutput struct {

	// An Id for the new run.
	RunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartWorkflowRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartWorkflowRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartWorkflowRun{}, middleware.After)
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
	addOpStartWorkflowRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartWorkflowRun(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartWorkflowRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartWorkflowRun",
	}
}

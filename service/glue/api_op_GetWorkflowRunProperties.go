// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the workflow run properties which were set during the run.
func (c *Client) GetWorkflowRunProperties(ctx context.Context, params *GetWorkflowRunPropertiesInput, optFns ...func(*Options)) (*GetWorkflowRunPropertiesOutput, error) {
	if params == nil {
		params = &GetWorkflowRunPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowRunProperties", params, optFns, addOperationGetWorkflowRunPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowRunPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowRunPropertiesInput struct {

	// Name of the workflow which was run.
	//
	// This member is required.
	Name *string

	// The ID of the workflow run whose run properties should be returned.
	//
	// This member is required.
	RunId *string
}

type GetWorkflowRunPropertiesOutput struct {

	// The workflow run properties which were set during the specified run.
	RunProperties map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetWorkflowRunPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetWorkflowRunProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWorkflowRunProperties{}, middleware.After)
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
	addOpGetWorkflowRunPropertiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowRunProperties(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetWorkflowRunProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetWorkflowRunProperties",
	}
}

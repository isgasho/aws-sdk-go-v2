// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters an AWS Batch job definition. Job definitions will be permanently
// deleted after 180 days.
func (c *Client) DeregisterJobDefinition(ctx context.Context, params *DeregisterJobDefinitionInput, optFns ...func(*Options)) (*DeregisterJobDefinitionOutput, error) {
	if params == nil {
		params = &DeregisterJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterJobDefinition", params, optFns, addOperationDeregisterJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterJobDefinitionInput struct {

	// The name and revision (name:revision) or full Amazon Resource Name (ARN) of the
	// job definition to deregister.
	//
	// This member is required.
	JobDefinition *string
}

type DeregisterJobDefinitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterJobDefinition{}, middleware.After)
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
	addOpDeregisterJobDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterJobDefinition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeregisterJobDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "DeregisterJobDefinition",
	}
}

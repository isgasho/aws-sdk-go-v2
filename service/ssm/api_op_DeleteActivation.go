// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an activation. You are not required to delete an activation. If you
// delete an activation, you can no longer use it to register additional managed
// instances. Deleting an activation does not de-register managed instances. You
// must manually de-register managed instances.
func (c *Client) DeleteActivation(ctx context.Context, params *DeleteActivationInput, optFns ...func(*Options)) (*DeleteActivationOutput, error) {
	if params == nil {
		params = &DeleteActivationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteActivation", params, optFns, addOperationDeleteActivationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteActivationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteActivationInput struct {

	// The ID of the activation that you want to delete.
	//
	// This member is required.
	ActivationId *string
}

type DeleteActivationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteActivationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteActivation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteActivation{}, middleware.After)
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
	addOpDeleteActivationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteActivation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteActivation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeleteActivation",
	}
}

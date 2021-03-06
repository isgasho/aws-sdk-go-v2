// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a patch baseline for a patch group.
func (c *Client) RegisterPatchBaselineForPatchGroup(ctx context.Context, params *RegisterPatchBaselineForPatchGroupInput, optFns ...func(*Options)) (*RegisterPatchBaselineForPatchGroupOutput, error) {
	if params == nil {
		params = &RegisterPatchBaselineForPatchGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterPatchBaselineForPatchGroup", params, optFns, addOperationRegisterPatchBaselineForPatchGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterPatchBaselineForPatchGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterPatchBaselineForPatchGroupInput struct {

	// The ID of the patch baseline to register the patch group with.
	//
	// This member is required.
	BaselineId *string

	// The name of the patch group that should be registered with the patch baseline.
	//
	// This member is required.
	PatchGroup *string
}

type RegisterPatchBaselineForPatchGroupOutput struct {

	// The ID of the patch baseline the patch group was registered with.
	BaselineId *string

	// The name of the patch group registered with the patch baseline.
	PatchGroup *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterPatchBaselineForPatchGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterPatchBaselineForPatchGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterPatchBaselineForPatchGroup{}, middleware.After)
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
	addOpRegisterPatchBaselineForPatchGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterPatchBaselineForPatchGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterPatchBaselineForPatchGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RegisterPatchBaselineForPatchGroup",
	}
}

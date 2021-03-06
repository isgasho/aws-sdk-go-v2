// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the default patch baseline. Note that Systems Manager supports
// creating multiple default patch baselines. For example, you can create a default
// patch baseline for each operating system. If you do not specify an operating
// system value, the default patch baseline for Windows is returned.
func (c *Client) GetDefaultPatchBaseline(ctx context.Context, params *GetDefaultPatchBaselineInput, optFns ...func(*Options)) (*GetDefaultPatchBaselineOutput, error) {
	if params == nil {
		params = &GetDefaultPatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDefaultPatchBaseline", params, optFns, addOperationGetDefaultPatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDefaultPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDefaultPatchBaselineInput struct {

	// Returns the default patch baseline for the specified operating system.
	OperatingSystem types.OperatingSystem
}

type GetDefaultPatchBaselineOutput struct {

	// The ID of the default patch baseline.
	BaselineId *string

	// The operating system for the returned patch baseline.
	OperatingSystem types.OperatingSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDefaultPatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDefaultPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDefaultPatchBaseline{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDefaultPatchBaseline(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDefaultPatchBaseline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetDefaultPatchBaseline",
	}
}

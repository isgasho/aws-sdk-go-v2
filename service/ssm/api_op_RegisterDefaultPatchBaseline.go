// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Defines the default patch baseline for the relevant operating system. To reset
// the AWS predefined patch baseline as the default, specify the full patch
// baseline ARN as the baseline ID value. For example, for CentOS, specify
// arn:aws:ssm:us-east-2:733109147000:patchbaseline/pb-0574b43a65ea646ed instead of
// pb-0574b43a65ea646ed.
func (c *Client) RegisterDefaultPatchBaseline(ctx context.Context, params *RegisterDefaultPatchBaselineInput, optFns ...func(*Options)) (*RegisterDefaultPatchBaselineOutput, error) {
	if params == nil {
		params = &RegisterDefaultPatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterDefaultPatchBaseline", params, optFns, addOperationRegisterDefaultPatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterDefaultPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterDefaultPatchBaselineInput struct {

	// The ID of the patch baseline that should be the default patch baseline.
	//
	// This member is required.
	BaselineId *string
}

type RegisterDefaultPatchBaselineOutput struct {

	// The ID of the default patch baseline.
	BaselineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterDefaultPatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterDefaultPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterDefaultPatchBaseline{}, middleware.After)
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
	addOpRegisterDefaultPatchBaselineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDefaultPatchBaseline(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterDefaultPatchBaseline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RegisterDefaultPatchBaseline",
	}
}

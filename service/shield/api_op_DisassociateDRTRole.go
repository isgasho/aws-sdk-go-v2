// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the DDoS Response Team's (DRT) access to your AWS account. To make a
// DisassociateDRTRole request, you must be subscribed to the Business Support plan
// (https://aws.amazon.com/premiumsupport/business-support/) or the Enterprise
// Support plan (https://aws.amazon.com/premiumsupport/enterprise-support/).
// However, if you are not subscribed to one of these support plans, but had been
// previously and had granted the DRT access to your account, you can submit a
// DisassociateDRTRole request to remove this access.
func (c *Client) DisassociateDRTRole(ctx context.Context, params *DisassociateDRTRoleInput, optFns ...func(*Options)) (*DisassociateDRTRoleOutput, error) {
	if params == nil {
		params = &DisassociateDRTRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateDRTRole", params, optFns, addOperationDisassociateDRTRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateDRTRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDRTRoleInput struct {
}

type DisassociateDRTRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateDRTRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateDRTRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateDRTRole{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDRTRole(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateDRTRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DisassociateDRTRole",
	}
}

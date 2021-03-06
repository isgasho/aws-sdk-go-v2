// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the Snowball service limit for your account, and also
// the number of Snowballs your account has in use. The default service limit for
// the number of Snowballs that you can have at one time is 1. If you want to
// increase your service limit, contact AWS Support.
func (c *Client) GetSnowballUsage(ctx context.Context, params *GetSnowballUsageInput, optFns ...func(*Options)) (*GetSnowballUsageOutput, error) {
	if params == nil {
		params = &GetSnowballUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSnowballUsage", params, optFns, addOperationGetSnowballUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSnowballUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSnowballUsageInput struct {
}

type GetSnowballUsageOutput struct {

	// The service limit for number of Snowballs this account can have at once. The
	// default service limit is 1 (one).
	SnowballLimit *int32

	// The number of Snowballs that this account is currently using.
	SnowballsInUse *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSnowballUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSnowballUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSnowballUsage{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSnowballUsage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSnowballUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "GetSnowballUsage",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the state of a file transfer protocol-enabled server from OFFLINE to
// ONLINE. It has no impact on a server that is already ONLINE. An ONLINE server
// can accept and process file transfer jobs. The state of STARTING indicates that
// the server is in an intermediate state, either not fully able to respond, or not
// fully online. The values of START_FAILED can indicate an error condition. No
// response is returned from this call.
func (c *Client) StartServer(ctx context.Context, params *StartServerInput, optFns ...func(*Options)) (*StartServerOutput, error) {
	if params == nil {
		params = &StartServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartServer", params, optFns, addOperationStartServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartServerInput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// that you start.
	//
	// This member is required.
	ServerId *string
}

type StartServerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartServer{}, middleware.After)
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
	addOpStartServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartServer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "StartServer",
	}
}

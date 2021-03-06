// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified Client VPN endpoint. You must disassociate all target
// networks before you can delete a Client VPN endpoint.
func (c *Client) DeleteClientVpnEndpoint(ctx context.Context, params *DeleteClientVpnEndpointInput, optFns ...func(*Options)) (*DeleteClientVpnEndpointOutput, error) {
	if params == nil {
		params = &DeleteClientVpnEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteClientVpnEndpoint", params, optFns, addOperationDeleteClientVpnEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteClientVpnEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteClientVpnEndpointInput struct {

	// The ID of the Client VPN to be deleted.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteClientVpnEndpointOutput struct {

	// The current state of the Client VPN endpoint.
	Status *types.ClientVpnEndpointStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteClientVpnEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteClientVpnEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteClientVpnEndpoint{}, middleware.After)
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
	addOpDeleteClientVpnEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteClientVpnEndpoint(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteClientVpnEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteClientVpnEndpoint",
	}
}

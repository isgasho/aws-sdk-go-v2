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

// Disables the specified resource attachment from propagating routes to the
// specified propagation route table.
func (c *Client) DisableTransitGatewayRouteTablePropagation(ctx context.Context, params *DisableTransitGatewayRouteTablePropagationInput, optFns ...func(*Options)) (*DisableTransitGatewayRouteTablePropagationOutput, error) {
	if params == nil {
		params = &DisableTransitGatewayRouteTablePropagationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableTransitGatewayRouteTablePropagation", params, optFns, addOperationDisableTransitGatewayRouteTablePropagationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableTransitGatewayRouteTablePropagationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableTransitGatewayRouteTablePropagationInput struct {

	// The ID of the attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// The ID of the propagation route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DisableTransitGatewayRouteTablePropagationOutput struct {

	// Information about route propagation.
	Propagation *types.TransitGatewayPropagation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableTransitGatewayRouteTablePropagationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisableTransitGatewayRouteTablePropagation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisableTransitGatewayRouteTablePropagation{}, middleware.After)
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
	addOpDisableTransitGatewayRouteTablePropagationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableTransitGatewayRouteTablePropagation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisableTransitGatewayRouteTablePropagation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableTransitGatewayRouteTablePropagation",
	}
}

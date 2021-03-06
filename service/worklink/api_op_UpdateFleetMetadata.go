// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates fleet metadata, such as DisplayName.
func (c *Client) UpdateFleetMetadata(ctx context.Context, params *UpdateFleetMetadataInput, optFns ...func(*Options)) (*UpdateFleetMetadataOutput, error) {
	if params == nil {
		params = &UpdateFleetMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFleetMetadata", params, optFns, addOperationUpdateFleetMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFleetMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFleetMetadataInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The fleet name to display. The existing DisplayName is unset if null is passed.
	DisplayName *string

	// The option to optimize for better performance by routing traffic through the
	// closest AWS Region to users, which may be outside of your home Region.
	OptimizeForEndUserLocation *bool
}

type UpdateFleetMetadataOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFleetMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFleetMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFleetMetadata{}, middleware.After)
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
	addOpUpdateFleetMetadataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetMetadata(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateFleetMetadata(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "UpdateFleetMetadata",
	}
}

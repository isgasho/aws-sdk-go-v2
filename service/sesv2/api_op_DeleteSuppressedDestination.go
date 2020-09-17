// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes an email address from the suppression list for your account.
func (c *Client) DeleteSuppressedDestination(ctx context.Context, params *DeleteSuppressedDestinationInput, optFns ...func(*Options)) (*DeleteSuppressedDestinationOutput, error) {
	stack := middleware.NewStack("DeleteSuppressedDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteSuppressedDestinationMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteSuppressedDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSuppressedDestination(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteSuppressedDestination",
			Err:           err,
		}
	}
	out := result.(*DeleteSuppressedDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to remove an email address from the suppression list for your account.
type DeleteSuppressedDestinationInput struct {
	// The suppressed email destination to remove from the account suppression list.
	EmailAddress *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type DeleteSuppressedDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteSuppressedDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSuppressedDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSuppressedDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSuppressedDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteSuppressedDestination",
	}
}
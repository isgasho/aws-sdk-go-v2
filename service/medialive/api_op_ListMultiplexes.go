// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve a list of the existing multiplexes.
func (c *Client) ListMultiplexes(ctx context.Context, params *ListMultiplexesInput, optFns ...func(*Options)) (*ListMultiplexesOutput, error) {
	stack := middleware.NewStack("ListMultiplexes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListMultiplexesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMultiplexes(options.Region), middleware.Before)

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
			OperationName: "ListMultiplexes",
			Err:           err,
		}
	}
	out := result.(*ListMultiplexesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListMultiplexesRequest
type ListMultiplexesInput struct {
	// The maximum number of items to return.
	MaxResults *int32
	// The token to retrieve the next page of results.
	NextToken *string
}

// Placeholder documentation for ListMultiplexesResponse
type ListMultiplexesOutput struct {
	// List of multiplexes.
	Multiplexes []*types.MultiplexSummary
	// Token for the next ListMultiplexes request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListMultiplexesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListMultiplexes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListMultiplexes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMultiplexes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "ListMultiplexes",
	}
}
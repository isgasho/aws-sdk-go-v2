// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets details about a multiplex.
func (c *Client) DescribeMultiplex(ctx context.Context, params *DescribeMultiplexInput, optFns ...func(*Options)) (*DescribeMultiplexOutput, error) {
	if params == nil {
		params = &DescribeMultiplexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMultiplex", params, optFns, addOperationDescribeMultiplexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMultiplexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeMultiplexRequest
type DescribeMultiplexInput struct {

	// The ID of the multiplex.
	//
	// This member is required.
	MultiplexId *string
}

// Placeholder documentation for DescribeMultiplexResponse
type DescribeMultiplexOutput struct {

	// The unique arn of the multiplex.
	Arn *string

	// A list of availability zones for the multiplex.
	AvailabilityZones []*string

	// A list of the multiplex output destinations.
	Destinations []*types.MultiplexOutputDestination

	// The unique id of the multiplex.
	Id *string

	// Configuration for a multiplex event.
	MultiplexSettings *types.MultiplexSettings

	// The name of the multiplex.
	Name *string

	// The number of currently healthy pipelines.
	PipelinesRunningCount *int32

	// The number of programs in the multiplex.
	ProgramCount *int32

	// The current state of the multiplex.
	State types.MultiplexState

	// A collection of key-value pairs.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMultiplexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMultiplex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMultiplex{}, middleware.After)
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
	addOpDescribeMultiplexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMultiplex(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeMultiplex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeMultiplex",
	}
}

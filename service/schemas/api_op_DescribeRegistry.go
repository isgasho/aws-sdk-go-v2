// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the registry.
func (c *Client) DescribeRegistry(ctx context.Context, params *DescribeRegistryInput, optFns ...func(*Options)) (*DescribeRegistryOutput, error) {
	if params == nil {
		params = &DescribeRegistryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRegistry", params, optFns, addOperationDescribeRegistryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRegistryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegistryInput struct {

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string
}

type DescribeRegistryOutput struct {

	// The description of the registry.
	Description *string

	// The ARN of the registry.
	RegistryArn *string

	// The name of the registry.
	RegistryName *string

	// Tags associated with the registry.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRegistryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRegistry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRegistry{}, middleware.After)
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
	addOpDescribeRegistryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegistry(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeRegistry(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "DescribeRegistry",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the code binding source URI.
func (c *Client) GetCodeBindingSource(ctx context.Context, params *GetCodeBindingSourceInput, optFns ...func(*Options)) (*GetCodeBindingSourceOutput, error) {
	stack := middleware.NewStack("GetCodeBindingSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCodeBindingSourceMiddlewares(stack)
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
	addOpGetCodeBindingSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCodeBindingSource(options.Region), middleware.Before)

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
			OperationName: "GetCodeBindingSource",
			Err:           err,
		}
	}
	out := result.(*GetCodeBindingSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCodeBindingSourceInput struct {
	// The name of the registry.
	RegistryName *string
	// The language of the code binding.
	Language *string
	// The name of the schema.
	SchemaName *string
	// Specifying this limits the results to only this schema version.
	SchemaVersion *string
}

type GetCodeBindingSourceOutput struct {
	Body []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCodeBindingSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCodeBindingSource{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCodeBindingSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCodeBindingSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "GetCodeBindingSource",
	}
}
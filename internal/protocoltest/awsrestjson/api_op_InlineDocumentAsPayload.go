// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes an inline document as the entire HTTP payload.
func (c *Client) InlineDocumentAsPayload(ctx context.Context, params *InlineDocumentAsPayloadInput, optFns ...func(*Options)) (*InlineDocumentAsPayloadOutput, error) {
	if params == nil {
		params = &InlineDocumentAsPayloadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InlineDocumentAsPayload", params, optFns, addOperationInlineDocumentAsPayloadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InlineDocumentAsPayloadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InlineDocumentAsPayloadInput struct {
	DocumentValue smithy.Document
}

type InlineDocumentAsPayloadOutput struct {
	DocumentValue smithy.Document

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInlineDocumentAsPayloadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInlineDocumentAsPayload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInlineDocumentAsPayload{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInlineDocumentAsPayload(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opInlineDocumentAsPayload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InlineDocumentAsPayload",
	}
}

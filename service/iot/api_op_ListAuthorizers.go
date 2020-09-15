// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the authorizers registered in your account.
func (c *Client) ListAuthorizers(ctx context.Context, params *ListAuthorizersInput, optFns ...func(*Options)) (*ListAuthorizersOutput, error) {
	stack := middleware.NewStack("ListAuthorizers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListAuthorizersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAuthorizers(options.Region), middleware.Before)

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
			OperationName: "ListAuthorizers",
			Err:           err,
		}
	}
	out := result.(*ListAuthorizersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAuthorizersInput struct {
	// The status of the list authorizers request.
	Status types.AuthorizerStatus
	// A marker used to get the next set of results.
	Marker *string
	// The maximum number of results to return at one time.
	PageSize *int32
	// Return the list of authorizers in ascending alphabetical order.
	AscendingOrder *bool
}

type ListAuthorizersOutput struct {
	// The authorizers.
	Authorizers []*types.AuthorizerSummary
	// A marker used to get the next set of results.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListAuthorizersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListAuthorizers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListAuthorizers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAuthorizers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListAuthorizers",
	}
}
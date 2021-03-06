// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the Outposts for your AWS account.
func (c *Client) ListOutposts(ctx context.Context, params *ListOutpostsInput, optFns ...func(*Options)) (*ListOutpostsOutput, error) {
	if params == nil {
		params = &ListOutpostsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOutposts", params, optFns, addOperationListOutpostsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOutpostsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOutpostsInput struct {

	// The maximum page size.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListOutpostsOutput struct {

	// The pagination token.
	NextToken *string

	// Information about the Outposts.
	Outposts []*types.Outpost

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOutpostsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOutposts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOutposts{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOutposts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListOutposts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "ListOutposts",
	}
}

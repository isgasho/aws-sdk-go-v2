// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays a list of all entitlements that have been granted to this account. This
// request returns 20 results per page.
func (c *Client) ListEntitlements(ctx context.Context, params *ListEntitlementsInput, optFns ...func(*Options)) (*ListEntitlementsOutput, error) {
	if params == nil {
		params = &ListEntitlementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntitlements", params, optFns, addOperationListEntitlementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitlementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitlementsInput struct {

	// The maximum number of results to return per API request. For example, you submit
	// a ListEntitlements request with MaxResults set at 5. Although 20 items match
	// your request, the service returns no more than the first 5 items. (The service
	// also returns a NextToken value that you can use to fetch the next batch of
	// results.) The service might return fewer results than the MaxResults value. If
	// MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 20 results per page.
	MaxResults *int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListEntitlements request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListEntitlements request a
	// second time and specify the NextToken value.
	NextToken *string
}

type ListEntitlementsOutput struct {

	// A list of entitlements that have been granted to you from other AWS accounts.
	Entitlements []*types.ListedEntitlement

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListEntitlements request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListEntitlements request a
	// second time and specify the NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEntitlementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEntitlements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEntitlements{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitlements(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListEntitlements(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "ListEntitlements",
	}
}

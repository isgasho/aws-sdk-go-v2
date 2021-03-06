// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the principals associated with the specified policy. Note: This API is
// deprecated. Please use ListTargetsForPolicy instead.
func (c *Client) ListPolicyPrincipals(ctx context.Context, params *ListPolicyPrincipalsInput, optFns ...func(*Options)) (*ListPolicyPrincipalsOutput, error) {
	if params == nil {
		params = &ListPolicyPrincipalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicyPrincipals", params, optFns, addOperationListPolicyPrincipalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPolicyPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPolicyPrincipals operation.
type ListPolicyPrincipalsInput struct {

	// The policy name.
	//
	// This member is required.
	PolicyName *string

	// Specifies the order for results. If true, the results are returned in ascending
	// creation order.
	AscendingOrder *bool

	// The marker for the next set of results.
	Marker *string

	// The result page size.
	PageSize *int32
}

// The output from the ListPolicyPrincipals operation.
type ListPolicyPrincipalsOutput struct {

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string

	// The descriptions of the principals.
	Principals []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPolicyPrincipalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPolicyPrincipals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPolicyPrincipals{}, middleware.After)
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
	addOpListPolicyPrincipalsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyPrincipals(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListPolicyPrincipals(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListPolicyPrincipals",
	}
}

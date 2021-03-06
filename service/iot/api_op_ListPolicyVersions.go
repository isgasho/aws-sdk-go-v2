// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the versions of the specified policy and identifies the default version.
func (c *Client) ListPolicyVersions(ctx context.Context, params *ListPolicyVersionsInput, optFns ...func(*Options)) (*ListPolicyVersionsOutput, error) {
	if params == nil {
		params = &ListPolicyVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicyVersions", params, optFns, addOperationListPolicyVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPolicyVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPolicyVersions operation.
type ListPolicyVersionsInput struct {

	// The policy name.
	//
	// This member is required.
	PolicyName *string
}

// The output from the ListPolicyVersions operation.
type ListPolicyVersionsOutput struct {

	// The policy versions.
	PolicyVersions []*types.PolicyVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPolicyVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPolicyVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPolicyVersions{}, middleware.After)
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
	addOpListPolicyVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListPolicyVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListPolicyVersions",
	}
}

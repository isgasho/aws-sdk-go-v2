// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the current namespace.
func (c *Client) DescribeNamespace(ctx context.Context, params *DescribeNamespaceInput, optFns ...func(*Options)) (*DescribeNamespaceOutput, error) {
	if params == nil {
		params = &DescribeNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNamespace", params, optFns, addOperationDescribeNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNamespaceInput struct {

	// The ID for the AWS account that contains the QuickSight namespace that you want
	// to describe.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that you want to describe.
	//
	// This member is required.
	Namespace *string
}

type DescribeNamespaceOutput struct {

	// The information about the namespace that you're describing. The response
	// includes the namespace ARN, name, AWS Region, creation status, and identity
	// store. DescribeNamespace also works for namespaces that are in the process of
	// being created. For incomplete namespaces, this API lists the namespace error
	// types and messages associated with the creation process.
	Namespace *types.NamespaceInfoV2

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeNamespace{}, middleware.After)
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
	addOpDescribeNamespaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNamespace(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeNamespace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeNamespace",
	}
}

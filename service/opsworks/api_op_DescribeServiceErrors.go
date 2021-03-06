// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes AWS OpsWorks Stacks service errors. Required Permissions: To use this
// action, an IAM user must have a Show, Deploy, or Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For more
// information about user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
// This call accepts only one resource-identifying parameter.
func (c *Client) DescribeServiceErrors(ctx context.Context, params *DescribeServiceErrorsInput, optFns ...func(*Options)) (*DescribeServiceErrorsOutput, error) {
	if params == nil {
		params = &DescribeServiceErrorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServiceErrors", params, optFns, addOperationDescribeServiceErrorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServiceErrorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServiceErrorsInput struct {

	// The instance ID. If you use this parameter, DescribeServiceErrors returns
	// descriptions of the errors associated with the specified instance.
	InstanceId *string

	// An array of service error IDs. If you use this parameter, DescribeServiceErrors
	// returns descriptions of the specified errors. Otherwise, it returns a
	// description of every error.
	ServiceErrorIds []*string

	// The stack ID. If you use this parameter, DescribeServiceErrors returns
	// descriptions of the errors associated with the specified stack.
	StackId *string
}

// Contains the response to a DescribeServiceErrors request.
type DescribeServiceErrorsOutput struct {

	// An array of ServiceError objects that describe the specified service errors.
	ServiceErrors []*types.ServiceError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeServiceErrorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeServiceErrors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeServiceErrors{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServiceErrors(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeServiceErrors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeServiceErrors",
	}
}

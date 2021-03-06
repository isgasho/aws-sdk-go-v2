// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the table, including the current status of the table,
// when it was created, the primary key schema, and any indexes on the table. If
// you issue a DescribeTable request immediately after a CreateTable request,
// DynamoDB might return a ResourceNotFoundException. This is because DescribeTable
// uses an eventually consistent query, and the metadata for your table might not
// be available at that moment. Wait for a few seconds, and then try the
// DescribeTable request again.
func (c *Client) DescribeTable(ctx context.Context, params *DescribeTableInput, optFns ...func(*Options)) (*DescribeTableOutput, error) {
	if params == nil {
		params = &DescribeTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTable", params, optFns, addOperationDescribeTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeTable operation.
type DescribeTableInput struct {

	// The name of the table to describe.
	//
	// This member is required.
	TableName *string
}

// Represents the output of a DescribeTable operation.
type DescribeTableOutput struct {

	// The properties of the table.
	Table *types.TableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeTable{}, middleware.After)
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
	addOpDescribeTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "DescribeTable",
	}
}

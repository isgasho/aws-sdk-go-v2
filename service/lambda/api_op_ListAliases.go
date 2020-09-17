// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of aliases
// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) for a
// Lambda function.
func (c *Client) ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	stack := middleware.NewStack("ListAliases", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListAliasesMiddlewares(stack)
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
	addOpListAliasesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAliases(options.Region), middleware.Before)

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
			OperationName: "ListAliases",
			Err:           err,
		}
	}
	out := result.(*ListAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAliasesInput struct {
	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string
	// Limit the number of aliases returned.
	MaxItems *int32
	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// MyFunction.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//     * Partial ARN -
	// 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it is limited to 64 characters
	// in length.
	FunctionName *string
	// Specify a function version to only list aliases that invoke that version.
	FunctionVersion *string
}

type ListAliasesOutput struct {
	// A list of aliases.
	Aliases []*types.AliasConfiguration
	// The pagination token that's included if more results are available.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListAliasesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListAliases{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListAliases{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAliases(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListAliases",
	}
}
// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a resource that was analyzed.
func (c *Client) GetAnalyzedResource(ctx context.Context, params *GetAnalyzedResourceInput, optFns ...func(*Options)) (*GetAnalyzedResourceOutput, error) {
	if params == nil {
		params = &GetAnalyzedResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAnalyzedResource", params, optFns, addOperationGetAnalyzedResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAnalyzedResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves an analyzed resource.
type GetAnalyzedResourceInput struct {

	// The ARN of the analyzer to retrieve information from.
	//
	// This member is required.
	AnalyzerArn *string

	// The ARN of the resource to retrieve information about.
	//
	// This member is required.
	ResourceArn *string
}

// The response to the request.
type GetAnalyzedResourceOutput struct {

	// An AnalyedResource object that contains information that Access Analyzer found
	// when it analyzed the resource.
	Resource *types.AnalyzedResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAnalyzedResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAnalyzedResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAnalyzedResource{}, middleware.After)
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
	addOpGetAnalyzedResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAnalyzedResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAnalyzedResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "GetAnalyzedResource",
	}
}

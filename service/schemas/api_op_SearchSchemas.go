// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Search the schemas
func (c *Client) SearchSchemas(ctx context.Context, params *SearchSchemasInput, optFns ...func(*Options)) (*SearchSchemasOutput, error) {
	if params == nil {
		params = &SearchSchemasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchSchemas", params, optFns, addOperationSearchSchemasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchSchemasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSchemasInput struct {

	// Specifying this limits the results to only schemas that include the provided
	// keywords.
	//
	// This member is required.
	Keywords *string

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string

	Limit *int32

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string
}

type SearchSchemasOutput struct {

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string

	// An array of SearchSchemaSummary information.
	Schemas []*types.SearchSchemaSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchSchemasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchSchemas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchSchemas{}, middleware.After)
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
	addOpSearchSchemasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSchemas(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSearchSchemas(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "SearchSchemas",
	}
}

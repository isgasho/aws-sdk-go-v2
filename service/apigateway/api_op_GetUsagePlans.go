// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets all the usage plans of the caller's account.
func (c *Client) GetUsagePlans(ctx context.Context, params *GetUsagePlansInput, optFns ...func(*Options)) (*GetUsagePlansOutput, error) {
	if params == nil {
		params = &GetUsagePlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsagePlans", params, optFns, addOperationGetUsagePlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsagePlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get all the usage plans of the caller's account.
type GetUsagePlansInput struct {

	// The identifier of the API key associated with the usage plans.
	KeyId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	Name *string

	// The current pagination position in the paged result set.
	Position *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a collection of usage plans for an AWS account. Create and Use Usage
// Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlansOutput struct {

	// The current page of elements from this collection.
	Items []*types.UsagePlan

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUsagePlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlans{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlans(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetUsagePlans(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsagePlans",
	}
}

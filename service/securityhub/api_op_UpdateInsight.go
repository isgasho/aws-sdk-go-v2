// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the Security Hub insight identified by the specified insight ARN.
func (c *Client) UpdateInsight(ctx context.Context, params *UpdateInsightInput, optFns ...func(*Options)) (*UpdateInsightOutput, error) {
	if params == nil {
		params = &UpdateInsightInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInsight", params, optFns, addOperationUpdateInsightMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInsightOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInsightInput struct {

	// The ARN of the insight that you want to update.
	//
	// This member is required.
	InsightArn *string

	// The updated filters that define this insight.
	Filters *types.AwsSecurityFindingFilters

	// The updated GroupBy attribute that defines this insight.
	GroupByAttribute *string

	// The updated name for the insight.
	Name *string
}

type UpdateInsightOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateInsightMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateInsight{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateInsight{}, middleware.After)
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
	addOpUpdateInsightValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInsight(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateInsight(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "UpdateInsight",
	}
}

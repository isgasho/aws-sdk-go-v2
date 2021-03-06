// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the customer feedback for a CodeGuru Reviewer recommendation.
func (c *Client) DescribeRecommendationFeedback(ctx context.Context, params *DescribeRecommendationFeedbackInput, optFns ...func(*Options)) (*DescribeRecommendationFeedbackOutput, error) {
	if params == nil {
		params = &DescribeRecommendationFeedbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecommendationFeedback", params, optFns, addOperationDescribeRecommendationFeedbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecommendationFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecommendationFeedbackInput struct {

	// The Amazon Resource Name (ARN) of the CodeReview
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html)
	// object.
	//
	// This member is required.
	CodeReviewArn *string

	// The recommendation ID that can be used to track the provided recommendations and
	// then to collect the feedback.
	//
	// This member is required.
	RecommendationId *string

	// Optional parameter to describe the feedback for a given user. If this is not
	// supplied, it defaults to the user making the request. The UserId is an IAM
	// principal that can be specified as an AWS account ID or an Amazon Resource Name
	// (ARN). For more information, see  Specifying a Principal
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#Principal_specifying)
	// in the AWS Identity and Access Management User Guide.
	UserId *string
}

type DescribeRecommendationFeedbackOutput struct {

	// The recommendation feedback given by the user.
	RecommendationFeedback *types.RecommendationFeedback

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRecommendationFeedbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRecommendationFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRecommendationFeedback{}, middleware.After)
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
	addOpDescribeRecommendationFeedbackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecommendationFeedback(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeRecommendationFeedback(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "DescribeRecommendationFeedback",
	}
}

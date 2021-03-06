// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a budget. The Request Syntax section shows the BudgetLimit syntax. For
// PlannedBudgetLimits, see the Examples
// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_DescribeBudget.html#API_DescribeBudget_Examples)
// section.
func (c *Client) DescribeBudget(ctx context.Context, params *DescribeBudgetInput, optFns ...func(*Options)) (*DescribeBudgetOutput, error) {
	if params == nil {
		params = &DescribeBudgetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudget", params, optFns, addOperationDescribeBudgetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DescribeBudget
type DescribeBudgetInput struct {

	// The accountId that is associated with the budget that you want a description of.
	//
	// This member is required.
	AccountId *string

	// The name of the budget that you want a description of.
	//
	// This member is required.
	BudgetName *string
}

// Response of DescribeBudget
type DescribeBudgetOutput struct {

	// The description of the budget.
	Budget *types.Budget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBudgetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudget{}, middleware.After)
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
	addOpDescribeBudgetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudget(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBudget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudget",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified receipt rule set and all of the receipt rules it contains.
// The currently active rule set cannot be deleted. For information about managing
// receipt rule sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rule-sets.html).
// You can execute this operation no more than once per second.
func (c *Client) DeleteReceiptRuleSet(ctx context.Context, params *DeleteReceiptRuleSetInput, optFns ...func(*Options)) (*DeleteReceiptRuleSetOutput, error) {
	if params == nil {
		params = &DeleteReceiptRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReceiptRuleSet", params, optFns, addOperationDeleteReceiptRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete a receipt rule set and all of the receipt rules
// it contains. You use receipt rule sets to receive email with Amazon SES. For
// more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DeleteReceiptRuleSetInput struct {

	// The name of the receipt rule set to delete.
	//
	// This member is required.
	RuleSetName *string
}

// An empty element returned on a successful request.
type DeleteReceiptRuleSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteReceiptRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteReceiptRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteReceiptRuleSet{}, middleware.After)
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
	addOpDeleteReceiptRuleSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReceiptRuleSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteReceiptRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteReceiptRuleSet",
	}
}

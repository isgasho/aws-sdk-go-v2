// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays a list of event categories for all event source types, or for a
// specified source type. For a list of the event categories and source types, go
// to Amazon Redshift Event Notifications
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html).
func (c *Client) DescribeEventCategories(ctx context.Context, params *DescribeEventCategoriesInput, optFns ...func(*Options)) (*DescribeEventCategoriesOutput, error) {
	if params == nil {
		params = &DescribeEventCategoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventCategories", params, optFns, addOperationDescribeEventCategoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventCategoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeEventCategoriesInput struct {

	// The source type, such as cluster or parameter group, to which the described
	// event categories apply. Valid values: cluster, cluster-snapshot,
	// cluster-parameter-group, cluster-security-group, and scheduled-action.
	SourceType *string
}

//
type DescribeEventCategoriesOutput struct {

	// A list of event categories descriptions.
	EventCategoriesMapList []*types.EventCategoriesMap

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventCategoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEventCategories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEventCategories{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventCategories(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventCategories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeEventCategories",
	}
}

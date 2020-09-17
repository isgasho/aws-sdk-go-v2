// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates tags for the specified delivery stream. A tag is a key-value
// pair that you can define and assign to AWS resources. If you specify a tag that
// already exists, the tag value is replaced with the value that you specify in the
// request. Tags are metadata. For example, you can add friendly names and
// descriptions or other types of information that can help you distinguish the
// delivery stream. For more information about tags, see Using Cost Allocation Tags
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide. Each delivery stream can have
// up to 50 tags. This operation has a limit of five transactions per second per
// account.
func (c *Client) TagDeliveryStream(ctx context.Context, params *TagDeliveryStreamInput, optFns ...func(*Options)) (*TagDeliveryStreamOutput, error) {
	stack := middleware.NewStack("TagDeliveryStream", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTagDeliveryStreamMiddlewares(stack)
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
	addOpTagDeliveryStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagDeliveryStream(options.Region), middleware.Before)

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
			OperationName: "TagDeliveryStream",
			Err:           err,
		}
	}
	out := result.(*TagDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagDeliveryStreamInput struct {
	// The name of the delivery stream to which you want to add the tags.
	DeliveryStreamName *string
	// A set of key-value pairs to use to create the tags.
	Tags []*types.Tag
}

type TagDeliveryStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTagDeliveryStreamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTagDeliveryStream{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagDeliveryStream{}, middleware.After)
}

func newServiceMetadataMiddleware_opTagDeliveryStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "TagDeliveryStream",
	}
}
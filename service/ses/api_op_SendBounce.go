// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates and sends a bounce message to the sender of an email you received
// through Amazon SES. You can only use this API on an email up to 24 hours after
// you receive it. You cannot use this API to send generic bounces for mail that
// was not received by Amazon SES. For information about receiving email through
// Amazon SES, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email.html).
// You can execute this operation no more than once per second.
func (c *Client) SendBounce(ctx context.Context, params *SendBounceInput, optFns ...func(*Options)) (*SendBounceOutput, error) {
	stack := middleware.NewStack("SendBounce", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSendBounceMiddlewares(stack)
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
	addOpSendBounceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendBounce(options.Region), middleware.Before)

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
			OperationName: "SendBounce",
			Err:           err,
		}
	}
	out := result.(*SendBounceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a bounce message to the sender of an email you
// received through Amazon SES.
type SendBounceInput struct {
	// The message ID of the message to be bounced.
	OriginalMessageId *string
	// Message-related DSN fields. If not specified, Amazon SES will choose the values.
	MessageDsn *types.MessageDsn
	// The address to use in the "From" header of the bounce message. This must be an
	// identity that you have verified with Amazon SES.
	BounceSender *string
	// A list of recipients of the bounced message, including the information required
	// to create the Delivery Status Notifications (DSNs) for the recipients. You must
	// specify at least one BouncedRecipientInfo in the list.
	BouncedRecipientInfoList []*types.BouncedRecipientInfo
	// Human-readable text for the bounce message to explain the failure. If not
	// specified, the text will be auto-generated based on the bounced recipient
	// information.
	Explanation *string
	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the address in the "From" header of the bounce. For more information
	// about sending authorization, see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	BounceSenderArn *string
}

// Represents a unique message ID.
type SendBounceOutput struct {
	// The message ID of the bounce message.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSendBounceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSendBounce{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSendBounce{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendBounce(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendBounce",
	}
}
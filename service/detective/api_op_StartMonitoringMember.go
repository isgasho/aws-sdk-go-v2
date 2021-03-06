// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends a request to enable data ingest for a member account that has a status of
// ACCEPTED_BUT_DISABLED. For valid member accounts, the status is updated as
// follows.
//
//     * If Detective enabled the member account, then the new status is
// ENABLED.
//
//     * If Detective cannot enable the member account, the status
// remains ACCEPTED_BUT_DISABLED.
func (c *Client) StartMonitoringMember(ctx context.Context, params *StartMonitoringMemberInput, optFns ...func(*Options)) (*StartMonitoringMemberOutput, error) {
	if params == nil {
		params = &StartMonitoringMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMonitoringMember", params, optFns, addOperationStartMonitoringMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMonitoringMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMonitoringMemberInput struct {

	// The account ID of the member account to try to enable. The account must be an
	// invited member account with a status of ACCEPTED_BUT_DISABLED.
	//
	// This member is required.
	AccountId *string

	// The ARN of the behavior graph.
	//
	// This member is required.
	GraphArn *string
}

type StartMonitoringMemberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartMonitoringMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMonitoringMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMonitoringMember{}, middleware.After)
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
	addOpStartMonitoringMemberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartMonitoringMember(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartMonitoringMember(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "StartMonitoringMember",
	}
}

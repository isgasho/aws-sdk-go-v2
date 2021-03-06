// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Restores the default settings for Device Defender audits for this account. Any
// configuration data you entered is deleted and all audit checks are reset to
// disabled.
func (c *Client) DeleteAccountAuditConfiguration(ctx context.Context, params *DeleteAccountAuditConfigurationInput, optFns ...func(*Options)) (*DeleteAccountAuditConfigurationOutput, error) {
	if params == nil {
		params = &DeleteAccountAuditConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccountAuditConfiguration", params, optFns, addOperationDeleteAccountAuditConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccountAuditConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccountAuditConfigurationInput struct {

	// If true, all scheduled audits are deleted.
	DeleteScheduledAudits *bool
}

type DeleteAccountAuditConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAccountAuditConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAccountAuditConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAccountAuditConfiguration{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccountAuditConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAccountAuditConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteAccountAuditConfiguration",
	}
}

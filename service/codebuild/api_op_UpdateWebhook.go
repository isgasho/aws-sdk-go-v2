// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the webhook associated with an AWS CodeBuild build project. If you use
// Bitbucket for your repository, rotateSecret is ignored.
func (c *Client) UpdateWebhook(ctx context.Context, params *UpdateWebhookInput, optFns ...func(*Options)) (*UpdateWebhookOutput, error) {
	if params == nil {
		params = &UpdateWebhookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWebhook", params, optFns, addOperationUpdateWebhookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWebhookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWebhookInput struct {

	// The name of the AWS CodeBuild project.
	//
	// This member is required.
	ProjectName *string

	// A regular expression used to determine which repository branches are built when
	// a webhook is triggered. If the name of a branch matches the regular expression,
	// then it is built. If branchFilter is empty, then all branches are built. It is
	// recommended that you use filterGroups instead of branchFilter.
	BranchFilter *string

	// Specifies the type of build this webhook will trigger.
	BuildType types.WebhookBuildType

	// An array of arrays of WebhookFilter objects used to determine if a webhook event
	// can trigger a build. A filter group must contain at least one
	// EVENTWebhookFilter.
	FilterGroups [][]*types.WebhookFilter

	// A boolean value that specifies whether the associated GitHub repository's secret
	// token should be updated. If you use Bitbucket for your repository, rotateSecret
	// is ignored.
	RotateSecret *bool
}

type UpdateWebhookOutput struct {

	// Information about a repository's webhook that is associated with a project in
	// AWS CodeBuild.
	Webhook *types.Webhook

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateWebhookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWebhook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWebhook{}, middleware.After)
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
	addOpUpdateWebhookValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWebhook(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateWebhook(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "UpdateWebhook",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of SpeechSynthesisTask objects ordered by their creation date.
// This operation can filter the tasks by their status, for example, allowing users
// to list only tasks that are completed.
func (c *Client) ListSpeechSynthesisTasks(ctx context.Context, params *ListSpeechSynthesisTasksInput, optFns ...func(*Options)) (*ListSpeechSynthesisTasksOutput, error) {
	if params == nil {
		params = &ListSpeechSynthesisTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSpeechSynthesisTasks", params, optFns, addOperationListSpeechSynthesisTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSpeechSynthesisTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSpeechSynthesisTasksInput struct {

	// Maximum number of speech synthesis tasks returned in a List operation.
	MaxResults *int32

	// The pagination token to use in the next request to continue the listing of
	// speech synthesis tasks.
	NextToken *string

	// Status of the speech synthesis tasks returned in a List operation
	Status types.TaskStatus
}

type ListSpeechSynthesisTasksOutput struct {

	// An opaque pagination token returned from the previous List operation in this
	// request. If present, this indicates where to continue the listing.
	NextToken *string

	// List of SynthesisTask objects that provides information from the specified task
	// in the list request, including output format, creation time, task status, and so
	// on.
	SynthesisTasks []*types.SynthesisTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSpeechSynthesisTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSpeechSynthesisTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSpeechSynthesisTasks{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSpeechSynthesisTasks(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSpeechSynthesisTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "ListSpeechSynthesisTasks",
	}
}

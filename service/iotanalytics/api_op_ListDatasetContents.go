// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists information about data set contents that have been created.
func (c *Client) ListDatasetContents(ctx context.Context, params *ListDatasetContentsInput, optFns ...func(*Options)) (*ListDatasetContentsOutput, error) {
	if params == nil {
		params = &ListDatasetContentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasetContents", params, optFns, addOperationListDatasetContentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetContentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetContentsInput struct {

	// The name of the data set whose contents information you want to list.
	//
	// This member is required.
	DatasetName *string

	// The maximum number of results to return in this request.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// A filter to limit results to those data set contents whose creation is scheduled
	// before the given time. See the field triggers.schedule in the CreateDataset
	// request. (timestamp)
	ScheduledBefore *time.Time

	// A filter to limit results to those data set contents whose creation is scheduled
	// on or after the given time. See the field triggers.schedule in the CreateDataset
	// request. (timestamp)
	ScheduledOnOrAfter *time.Time
}

type ListDatasetContentsOutput struct {

	// Summary information about data set contents that have been created.
	DatasetContentSummaries []*types.DatasetContentSummary

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDatasetContentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDatasetContents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDatasetContents{}, middleware.After)
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
	addOpListDatasetContentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasetContents(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDatasetContents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "ListDatasetContents",
	}
}

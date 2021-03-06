// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes one or more documents from an index. The documents must have been added
// with the BatchPutDocument operation. The documents are deleted asynchronously.
// You can see the progress of the deletion by using AWS CloudWatch. Any error
// messages releated to the processing of the batch are sent to you CloudWatch log.
func (c *Client) BatchDeleteDocument(ctx context.Context, params *BatchDeleteDocumentInput, optFns ...func(*Options)) (*BatchDeleteDocumentOutput, error) {
	if params == nil {
		params = &BatchDeleteDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteDocument", params, optFns, addOperationBatchDeleteDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteDocumentInput struct {

	// One or more identifiers for documents to delete from the index.
	//
	// This member is required.
	DocumentIdList []*string

	// The identifier of the index that contains the documents to delete.
	//
	// This member is required.
	IndexId *string

	// Maps a particular data source sync job to a particular data source.
	DataSourceSyncJobMetricTarget *types.DataSourceSyncJobMetricTarget
}

type BatchDeleteDocumentOutput struct {

	// A list of documents that could not be removed from the index. Each entry
	// contains an error message that indicates why the document couldn't be removed
	// from the index.
	FailedDocuments []*types.BatchDeleteDocumentResponseFailedDocument

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeleteDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteDocument{}, middleware.After)
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
	addOpBatchDeleteDocumentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteDocument(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchDeleteDocument(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "BatchDeleteDocument",
	}
}

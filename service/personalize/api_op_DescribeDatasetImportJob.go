// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the dataset import job created by CreateDatasetImportJob, including
// the import job status.
func (c *Client) DescribeDatasetImportJob(ctx context.Context, params *DescribeDatasetImportJobInput, optFns ...func(*Options)) (*DescribeDatasetImportJobOutput, error) {
	if params == nil {
		params = &DescribeDatasetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDatasetImportJob", params, optFns, addOperationDescribeDatasetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetImportJobInput struct {

	// The Amazon Resource Name (ARN) of the dataset import job to describe.
	//
	// This member is required.
	DatasetImportJobArn *string
}

type DescribeDatasetImportJobOutput struct {

	// Information about the dataset import job, including the status. The status is
	// one of the following values:
	//
	//     * CREATE PENDING
	//
	//     * CREATE IN_PROGRESS
	//
	//
	// * ACTIVE
	//
	//     * CREATE FAILED
	DatasetImportJob *types.DatasetImportJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDatasetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDatasetImportJob{}, middleware.After)
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
	addOpDescribeDatasetImportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatasetImportJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDatasetImportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeDatasetImportJob",
	}
}

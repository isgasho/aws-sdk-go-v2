// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns a BatchPrediction that includes detailed metadata, status, and data file
// information for a Batch Prediction request.
func (c *Client) GetBatchPrediction(ctx context.Context, params *GetBatchPredictionInput, optFns ...func(*Options)) (*GetBatchPredictionOutput, error) {
	stack := middleware.NewStack("GetBatchPrediction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetBatchPredictionMiddlewares(stack)
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
	addOpGetBatchPredictionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBatchPrediction(options.Region), middleware.Before)

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
			OperationName: "GetBatchPrediction",
			Err:           err,
		}
	}
	out := result.(*GetBatchPredictionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBatchPredictionInput struct {
	// An ID assigned to the BatchPrediction at creation.
	BatchPredictionId *string
}

// Represents the output of a GetBatchPrediction operation and describes a
// BatchPrediction.
type GetBatchPredictionOutput struct {
	// The ID of the MLModel that generated predictions for the BatchPrediction
	// request.
	MLModelId *string
	// The time when the BatchPrediction was created. The time is expressed in epoch
	// time.
	CreatedAt *time.Time
	// The status of the BatchPrediction, which can be one of the following values:
	//
	//
	// * PENDING - Amazon Machine Learning (Amazon ML) submitted a request to generate
	// batch predictions.
	//
	//     * INPROGRESS - The batch predictions are in progress.
	//
	//
	// * FAILED - The request to perform a batch prediction did not run to completion.
	// It is not usable.
	//
	//     * COMPLETED - The batch prediction process completed
	// successfully.
	//
	//     * DELETED - The BatchPrediction is marked as deleted. It is
	// not usable.
	Status types.EntityStatus
	// The time of the most recent edit to BatchPrediction. The time is expressed in
	// epoch time.
	LastUpdatedAt *time.Time
	// The number of total records that Amazon Machine Learning saw while processing
	// the BatchPrediction.
	TotalRecordCount *int64
	// The number of invalid records that Amazon Machine Learning saw while processing
	// the BatchPrediction.
	InvalidRecordCount *int64
	// The epoch time when Amazon Machine Learning marked the BatchPrediction as
	// INPROGRESS. StartedAt isn't available if the BatchPrediction is in the PENDING
	// state.
	StartedAt *time.Time
	// An ID assigned to the BatchPrediction at creation. This value should be
	// identical to the value of the BatchPredictionID in the request.
	BatchPredictionId *string
	// The approximate CPU time in milliseconds that Amazon Machine Learning spent
	// processing the BatchPrediction, normalized and scaled on computation resources.
	// ComputeTime is only available if the BatchPrediction is in the COMPLETED state.
	ComputeTime *int64
	// The AWS user account that invoked the BatchPrediction. The account type can be
	// either an AWS root account or an AWS Identity and Access Management (IAM) user
	// account.
	CreatedByIamUser *string
	// The epoch time when Amazon Machine Learning marked the BatchPrediction as
	// COMPLETED or FAILED. FinishedAt is only available when the BatchPrediction is in
	// the COMPLETED or FAILED state.
	FinishedAt *time.Time
	// The location of the data file or directory in Amazon Simple Storage Service
	// (Amazon S3).
	InputDataLocationS3 *string
	// A link to the file that contains logs of the CreateBatchPrediction operation.
	LogUri *string
	// A description of the most recent details about processing the batch prediction
	// request.
	Message *string
	// The ID of the DataSource that was used to create the BatchPrediction.
	BatchPredictionDataSourceId *string
	// A user-supplied name or description of the BatchPrediction.
	Name *string
	// The location of an Amazon S3 bucket or directory to receive the operation
	// results.
	OutputUri *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetBatchPredictionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetBatchPrediction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBatchPrediction{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBatchPrediction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "GetBatchPrediction",
	}
}
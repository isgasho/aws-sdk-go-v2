// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the detector version’s status. You can perform the following promotions
// or demotions using UpdateDetectorVersionStatus: DRAFT to ACTIVE, ACTIVE to
// INACTIVE, and INACTIVE to ACTIVE.
func (c *Client) UpdateDetectorVersionStatus(ctx context.Context, params *UpdateDetectorVersionStatusInput, optFns ...func(*Options)) (*UpdateDetectorVersionStatusOutput, error) {
	stack := middleware.NewStack("UpdateDetectorVersionStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDetectorVersionStatusMiddlewares(stack)
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
	addOpUpdateDetectorVersionStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDetectorVersionStatus(options.Region), middleware.Before)

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
			OperationName: "UpdateDetectorVersionStatus",
			Err:           err,
		}
	}
	out := result.(*UpdateDetectorVersionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDetectorVersionStatusInput struct {
	// The detector ID.
	DetectorId *string
	// The new status.
	Status types.DetectorVersionStatus
	// The detector version ID.
	DetectorVersionId *string
}

type UpdateDetectorVersionStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDetectorVersionStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDetectorVersionStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDetectorVersionStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDetectorVersionStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "UpdateDetectorVersionStatus",
	}
}
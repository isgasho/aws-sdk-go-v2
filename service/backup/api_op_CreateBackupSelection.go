// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a JSON document that specifies a set of resources to assign to a backup
// plan. Resources can be included by specifying patterns for a ListOfTags and
// selected Resources. For example, consider the following patterns:
//
//     *
// Resources: "arn:aws:ec2:region:account-id:volume/volume-id"
//
//     *
// ConditionKey:"department"ConditionValue:"finance"ConditionType:"STRINGEQUALS"
//
//
// *
// ConditionKey:"importance"ConditionValue:"critical"ConditionType:"STRINGEQUALS"
//
// Using
// these patterns would back up all Amazon Elastic Block Store (Amazon EBS) volumes
// that are tagged as "department=finance", "importance=critical", in addition to
// an EBS volume with the specified volume Id. Resources and conditions are
// additive in that all resources that match the pattern are selected. This
// shouldn't be confused with a logical AND, where all conditions must match. The
// matching patterns are logically 'put together using the OR operator. In other
// words, all patterns that match are selected for backup.
func (c *Client) CreateBackupSelection(ctx context.Context, params *CreateBackupSelectionInput, optFns ...func(*Options)) (*CreateBackupSelectionOutput, error) {
	stack := middleware.NewStack("CreateBackupSelection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateBackupSelectionMiddlewares(stack)
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
	addOpCreateBackupSelectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackupSelection(options.Region), middleware.Before)

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
			OperationName: "CreateBackupSelection",
			Err:           err,
		}
	}
	out := result.(*CreateBackupSelectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBackupSelectionInput struct {
	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string
	// Uniquely identifies the backup plan to be associated with the selection of
	// resources.
	BackupPlanId *string
	// Specifies the body of a request to assign a set of resources to a backup plan.
	BackupSelection *types.BackupSelection
}

type CreateBackupSelectionOutput struct {
	// Uniquely identifies the body of a request to assign a set of resources to a
	// backup plan.
	SelectionId *string
	// Uniquely identifies a backup plan.
	BackupPlanId *string
	// The date and time a backup selection is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateBackupSelectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateBackupSelection{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBackupSelection{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateBackupSelection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "CreateBackupSelection",
	}
}
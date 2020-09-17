// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes multiple tables at once. After completing this operation, you no longer
// have access to the table versions and partitions that belong to the deleted
// table. AWS Glue deletes these "orphaned" resources asynchronously in a timely
// manner, at the discretion of the service. To ensure the immediate deletion of
// all related resources, before calling BatchDeleteTable, use DeleteTableVersion
// or BatchDeleteTableVersion, and DeletePartition or BatchDeletePartition, to
// delete any resources that belong to the table.
func (c *Client) BatchDeleteTable(ctx context.Context, params *BatchDeleteTableInput, optFns ...func(*Options)) (*BatchDeleteTableOutput, error) {
	stack := middleware.NewStack("BatchDeleteTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchDeleteTableMiddlewares(stack)
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
	addOpBatchDeleteTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteTable(options.Region), middleware.Before)

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
			OperationName: "BatchDeleteTable",
			Err:           err,
		}
	}
	out := result.(*BatchDeleteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteTableInput struct {
	// A list of the table to delete.
	TablesToDelete []*string
	// The name of the catalog database in which the tables to delete reside. For Hive
	// compatibility, this name is entirely lowercase.
	DatabaseName *string
	// The ID of the Data Catalog where the table resides. If none is provided, the AWS
	// account ID is used by default.
	CatalogId *string
}

type BatchDeleteTableOutput struct {
	// A list of errors encountered in attempting to delete the specified tables.
	Errors []*types.TableError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchDeleteTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchDeleteTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchDeleteTable",
	}
}
// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the metadata for the tables in the specified data catalog database.
func (c *Client) ListTableMetadata(ctx context.Context, params *ListTableMetadataInput, optFns ...func(*Options)) (*ListTableMetadataOutput, error) {
	if params == nil {
		params = &ListTableMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTableMetadata", params, optFns, addOperationListTableMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTableMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTableMetadataInput struct {

	// The name of the data catalog for which table metadata should be returned.
	//
	// This member is required.
	CatalogName *string

	// The name of the database for which table metadata should be returned.
	//
	// This member is required.
	DatabaseName *string

	// A regex filter that pattern-matches table names. If no expression is supplied,
	// metadata for all tables are listed.
	Expression *string

	// Specifies the maximum number of results to return.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string
}

type ListTableMetadataOutput struct {

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// A list of table metadata.
	TableMetadataList []*types.TableMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTableMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTableMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTableMetadata{}, middleware.After)
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
	addOpListTableMetadataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTableMetadata(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTableMetadata(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "ListTableMetadata",
	}
}

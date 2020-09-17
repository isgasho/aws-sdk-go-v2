// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches users and lists the ones that meet a set of filter and sort criteria.
func (c *Client) SearchUsers(ctx context.Context, params *SearchUsersInput, optFns ...func(*Options)) (*SearchUsersOutput, error) {
	stack := middleware.NewStack("SearchUsers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchUsersMiddlewares(stack)
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
	addOpSearchUsersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchUsers(options.Region), middleware.Before)

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
			OperationName: "SearchUsers",
			Err:           err,
		}
	}
	out := result.(*SearchUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchUsersInput struct {
	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	// Required.
	NextToken *string
	// The sort order to use in listing the filtered set of users. Required. Supported
	// sort keys are UserId, FirstName, LastName, Email, and EnrollmentStatus.
	SortCriteria []*types.Sort
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved. Required.
	MaxResults *int32
	// The filters to use for listing a specific set of users. Required. Supported
	// filter keys are UserId, FirstName, LastName, Email, and EnrollmentStatus.
	Filters []*types.Filter
}

type SearchUsersOutput struct {
	// The users that meet the specified set of filter criteria, in sort order.
	Users []*types.UserData
	// The token returned to indicate that there is more data available.
	NextToken *string
	// The total number of users returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchUsersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchUsers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchUsers{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchUsers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchUsers",
	}
}
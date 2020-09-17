// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of available network profiles.
func (c *Client) ListNetworkProfiles(ctx context.Context, params *ListNetworkProfilesInput, optFns ...func(*Options)) (*ListNetworkProfilesOutput, error) {
	stack := middleware.NewStack("ListNetworkProfiles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListNetworkProfilesMiddlewares(stack)
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
	addOpListNetworkProfilesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListNetworkProfiles(options.Region), middleware.Before)

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
			OperationName: "ListNetworkProfiles",
			Err:           err,
		}
	}
	out := result.(*ListNetworkProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNetworkProfilesInput struct {
	// The type of network profile to return information about. Valid values are listed
	// here.
	Type types.NetworkProfileType
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
	// The Amazon Resource Name (ARN) of the project for which you want to list network
	// profiles.
	Arn *string
}

type ListNetworkProfilesOutput struct {
	// A list of the available network profiles.
	NetworkProfiles []*types.NetworkProfile
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListNetworkProfilesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListNetworkProfiles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNetworkProfiles{}, middleware.After)
}

func newServiceMetadataMiddleware_opListNetworkProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListNetworkProfiles",
	}
}
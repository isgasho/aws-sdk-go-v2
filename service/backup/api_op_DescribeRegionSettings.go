// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current service opt-in settings for the Region. If the service has a
// value set to true, AWS Backup attempts to protect that service's resources in
// this Region, when included in an on-demand backup or scheduled backup plan. If
// the value is set to false for a service, AWS Backup does not attempt to protect
// that service's resources in this Region.
func (c *Client) DescribeRegionSettings(ctx context.Context, params *DescribeRegionSettingsInput, optFns ...func(*Options)) (*DescribeRegionSettingsOutput, error) {
	if params == nil {
		params = &DescribeRegionSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRegionSettings", params, optFns, addOperationDescribeRegionSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRegionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegionSettingsInput struct {
}

type DescribeRegionSettingsOutput struct {

	// Returns a list of all services along with the opt-in preferences in the region.
	ResourceTypeOptInPreference map[string]*bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRegionSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRegionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRegionSettings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegionSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeRegionSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeRegionSettings",
	}
}

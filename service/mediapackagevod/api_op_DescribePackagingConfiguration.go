// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a description of a MediaPackage VOD PackagingConfiguration resource.
func (c *Client) DescribePackagingConfiguration(ctx context.Context, params *DescribePackagingConfigurationInput, optFns ...func(*Options)) (*DescribePackagingConfigurationOutput, error) {
	if params == nil {
		params = &DescribePackagingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePackagingConfiguration", params, optFns, addOperationDescribePackagingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePackagingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePackagingConfigurationInput struct {

	// The ID of a MediaPackage VOD PackagingConfiguration resource.
	//
	// This member is required.
	Id *string
}

type DescribePackagingConfigurationOutput struct {

	// The ARN of the PackagingConfiguration.
	Arn *string

	// A CMAF packaging configuration.
	CmafPackage *types.CmafPackage

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage

	// The ID of the PackagingConfiguration.
	Id *string

	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *types.MssPackage

	// The ID of a PackagingGroup.
	PackagingGroupId *string

	// A collection of tags associated with a resource
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePackagingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackagingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackagingConfiguration{}, middleware.After)
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
	addOpDescribePackagingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackagingConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribePackagingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "DescribePackagingConfiguration",
	}
}

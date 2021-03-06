// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details of one or more remediation configurations.
func (c *Client) DescribeRemediationConfigurations(ctx context.Context, params *DescribeRemediationConfigurationsInput, optFns ...func(*Options)) (*DescribeRemediationConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeRemediationConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRemediationConfigurations", params, optFns, addOperationDescribeRemediationConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRemediationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRemediationConfigurationsInput struct {

	// A list of AWS Config rule names of remediation configurations for which you want
	// details.
	//
	// This member is required.
	ConfigRuleNames []*string
}

type DescribeRemediationConfigurationsOutput struct {

	// Returns a remediation configuration object.
	RemediationConfigurations []*types.RemediationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRemediationConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRemediationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRemediationConfigurations{}, middleware.After)
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
	addOpDescribeRemediationConfigurationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRemediationConfigurations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeRemediationConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeRemediationConfigurations",
	}
}

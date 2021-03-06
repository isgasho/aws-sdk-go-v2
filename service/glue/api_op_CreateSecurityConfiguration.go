// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new security configuration. A security configuration is a set of
// security properties that can be used by AWS Glue. You can use a security
// configuration to encrypt data at rest. For information about using security
// configurations in AWS Glue, see Encrypting Data Written by Crawlers, Jobs, and
// Development Endpoints
// (https://docs.aws.amazon.com/glue/latest/dg/encryption-security-configuration.html).
func (c *Client) CreateSecurityConfiguration(ctx context.Context, params *CreateSecurityConfigurationInput, optFns ...func(*Options)) (*CreateSecurityConfigurationOutput, error) {
	if params == nil {
		params = &CreateSecurityConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityConfiguration", params, optFns, addOperationCreateSecurityConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityConfigurationInput struct {

	// The encryption configuration for the new security configuration.
	//
	// This member is required.
	EncryptionConfiguration *types.EncryptionConfiguration

	// The name for the new security configuration.
	//
	// This member is required.
	Name *string
}

type CreateSecurityConfigurationOutput struct {

	// The time at which the new security configuration was created.
	CreatedTimestamp *time.Time

	// The name assigned to the new security configuration.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSecurityConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSecurityConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSecurityConfiguration{}, middleware.After)
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
	addOpCreateSecurityConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateSecurityConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateSecurityConfiguration",
	}
}

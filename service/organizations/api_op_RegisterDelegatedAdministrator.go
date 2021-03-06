// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the specified member account to administer the Organizations features of
// the specified AWS service. It grants read-only access to AWS Organizations
// service data. The account still requires IAM permissions to access and
// administer the AWS service. You can run this action only for AWS services that
// support this feature. For a current list of services that support it, see the
// column Supports Delegated Administrator in the table at AWS Services that you
// can use with AWS Organizations
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrated-services-list.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's master account.
func (c *Client) RegisterDelegatedAdministrator(ctx context.Context, params *RegisterDelegatedAdministratorInput, optFns ...func(*Options)) (*RegisterDelegatedAdministratorOutput, error) {
	if params == nil {
		params = &RegisterDelegatedAdministratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterDelegatedAdministrator", params, optFns, addOperationRegisterDelegatedAdministratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterDelegatedAdministratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterDelegatedAdministratorInput struct {

	// The account ID number of the member account in the organization to register as a
	// delegated administrator.
	//
	// This member is required.
	AccountId *string

	// The service principal of the AWS service for which you want to make the member
	// account a delegated administrator.
	//
	// This member is required.
	ServicePrincipal *string
}

type RegisterDelegatedAdministratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterDelegatedAdministratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterDelegatedAdministrator{}, middleware.After)
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
	addOpRegisterDelegatedAdministratorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDelegatedAdministrator(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterDelegatedAdministrator(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "RegisterDelegatedAdministrator",
	}
}

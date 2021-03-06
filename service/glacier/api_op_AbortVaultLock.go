// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation aborts the vault locking process if the vault lock is not in the
// Locked state. If the vault lock is in the Locked state when this operation is
// requested, the operation returns an AccessDeniedException error. Aborting the
// vault locking process removes the vault lock policy from the specified vault. A
// vault lock is put into the InProgress state by calling InitiateVaultLock. A
// vault lock is put into the Locked state by calling CompleteVaultLock. You can
// get the state of a vault lock by calling GetVaultLock. For more information
// about the vault locking process, see Amazon Glacier Vault Lock
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html). For more
// information about vault lock policies, see Amazon Glacier Access Control with
// Vault Lock Policies
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock-policy.html).
// This operation is idempotent. You can successfully invoke this operation
// multiple times, if the vault lock is in the InProgress state or if there is no
// policy associated with the vault.
func (c *Client) AbortVaultLock(ctx context.Context, params *AbortVaultLockInput, optFns ...func(*Options)) (*AbortVaultLockOutput, error) {
	if params == nil {
		params = &AbortVaultLockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AbortVaultLock", params, optFns, addOperationAbortVaultLockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AbortVaultLockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input values for AbortVaultLock.
type AbortVaultLockInput struct {

	// The AccountId value is the AWS account ID. This value must match the AWS account
	// ID associated with the credentials used to sign the request. You can either
	// specify an AWS account ID or optionally a single '-' (hyphen), in which case
	// Amazon Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

type AbortVaultLockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAbortVaultLockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAbortVaultLock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAbortVaultLock{}, middleware.After)
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
	addOpAbortVaultLockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAbortVaultLock(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)
	return nil
}

func newServiceMetadataMiddleware_opAbortVaultLock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "AbortVaultLock",
	}
}

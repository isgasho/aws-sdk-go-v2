// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns permissions from a private CA to a designated AWS service. Services are
// specified by their service principals and can be given permission to create and
// retrieve certificates on a private CA. Services can also be given permission to
// list the active permissions that the private CA has granted. For ACM to
// automatically renew your private CA's certificates, you must assign all possible
// permissions from the CA to the ACM service principal. At this time, you can only
// assign permissions to ACM (acm.amazonaws.com). Permissions can be revoked with
// the DeletePermission action and listed with the ListPermissions action.
func (c *Client) CreatePermission(ctx context.Context, params *CreatePermissionInput, optFns ...func(*Options)) (*CreatePermissionOutput, error) {
	if params == nil {
		params = &CreatePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePermission", params, optFns, addOperationCreatePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePermissionInput struct {

	// The actions that the specified AWS service principal can use. These include
	// IssueCertificate, GetCertificate, and ListPermissions.
	//
	// This member is required.
	Actions []types.ActionType

	// The Amazon Resource Name (ARN) of the CA that grants the permissions. You can
	// find the ARN by calling the ListCertificateAuthorities action. This must have
	// the following form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string

	// The AWS service or identity that receives the permission. At this time, the only
	// valid principal is acm.amazonaws.com.
	//
	// This member is required.
	Principal *string

	// The ID of the calling account.
	SourceAccount *string
}

type CreatePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePermission{}, middleware.After)
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
	addOpCreatePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePermission(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreatePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "CreatePermission",
	}
}

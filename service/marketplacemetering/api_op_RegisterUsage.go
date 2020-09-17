// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Paid container software products sold through AWS Marketplace must integrate
// with the AWS Marketplace Metering Service and call the RegisterUsage operation
// for software entitlement and metering. Free and BYOL products for Amazon ECS or
// Amazon EKS aren't required to call RegisterUsage, but you may choose to do so if
// you would like to receive usage data in your seller reports. The sections below
// explain the behavior of RegisterUsage. RegisterUsage performs two primary
// functions: metering and entitlement.  <ul> <li> <p> <i>Entitlement</i>:
// RegisterUsage allows you to verify that the customer running your paid software
// is subscribed to your product on AWS Marketplace, enabling you to guard against
// unauthorized use. Your container image that integrates with RegisterUsage is
// only required to guard against unauthorized use at container startup, as such a
// CustomerNotSubscribedException/PlatformNotSupportedException will only be thrown
// on the initial call to RegisterUsage. Subsequent calls from the same Amazon ECS
// task instance (e.g. task-id) or Amazon EKS pod will not throw a
// CustomerNotSubscribedException, even if the customer unsubscribes while the
// Amazon ECS task or Amazon EKS pod is still running.</p> </li> <li> <p>
// <i>Metering</i>: RegisterUsage meters software use per ECS task, per hour, or
// per pod for Amazon EKS with usage prorated to the second. A minimum of 1 minute
// of usage applies to tasks that are short lived. For example, if a customer has a
// 10 node Amazon ECS or Amazon EKS cluster and a service configured as a Daemon
// Set, then Amazon ECS or Amazon EKS will launch a task on all 10 cluster nodes
// and the customer will be charged: (10 * hourly_rate). Metering for software use
// is automatically handled by the AWS Marketplace Metering Control Plane -- your
// software is not required to perform any metering specific actions, other than
// call RegisterUsage once for metering of software use to commence. The AWS
// Marketplace Metering Control Plane will also continue to bill customers for
// running ECS tasks and Amazon EKS pods, regardless of the customers subscription
// state, removing the need for your software to perform entitlement checks at
// runtime.</p> </li> </ul>
func (c *Client) RegisterUsage(ctx context.Context, params *RegisterUsageInput, optFns ...func(*Options)) (*RegisterUsageOutput, error) {
	stack := middleware.NewStack("RegisterUsage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterUsageMiddlewares(stack)
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
	addOpRegisterUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterUsage(options.Region), middleware.Before)

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
			OperationName: "RegisterUsage",
			Err:           err,
		}
	}
	out := result.(*RegisterUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterUsageInput struct {
	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of a new
	// product.
	ProductCode *string
	// (Optional) To scope down the registration to a specific running software
	// instance and guard against replay attacks.
	Nonce *string
	// Public Key Version provided by AWS Marketplace
	PublicKeyVersion *int32
}

type RegisterUsageOutput struct {
	// (Optional) Only included when public key version has expired
	PublicKeyRotationTimestamp *time.Time
	// JWT Token
	Signature *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterUsageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterUsage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterUsage{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "RegisterUsage",
	}
}
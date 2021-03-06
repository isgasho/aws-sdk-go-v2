// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a custom Availability Zone (AZ). A custom AZ is an on-premises AZ that
// is integrated with a VMware vSphere cluster. For more information about RDS on
// VMware, see the  RDS on VMware User Guide.
// (https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
func (c *Client) DeleteCustomAvailabilityZone(ctx context.Context, params *DeleteCustomAvailabilityZoneInput, optFns ...func(*Options)) (*DeleteCustomAvailabilityZoneOutput, error) {
	if params == nil {
		params = &DeleteCustomAvailabilityZoneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomAvailabilityZone", params, optFns, addOperationDeleteCustomAvailabilityZoneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomAvailabilityZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCustomAvailabilityZoneInput struct {

	// The custom AZ identifier.
	//
	// This member is required.
	CustomAvailabilityZoneId *string
}

type DeleteCustomAvailabilityZoneOutput struct {

	// A custom Availability Zone (AZ) is an on-premises AZ that is integrated with a
	// VMware vSphere cluster. For more information about RDS on VMware, see the  RDS
	// on VMware User Guide.
	// (https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
	CustomAvailabilityZone *types.CustomAvailabilityZone

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCustomAvailabilityZoneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCustomAvailabilityZone{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCustomAvailabilityZone{}, middleware.After)
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
	addOpDeleteCustomAvailabilityZoneValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomAvailabilityZone(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCustomAvailabilityZone(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteCustomAvailabilityZone",
	}
}

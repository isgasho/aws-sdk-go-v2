// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a connection between a customer network and a specific AWS Direct
// Connect location.  <p>A connection links your internal network to an AWS Direct
// Connect location over a standard Ethernet fiber-optic cable. One end of the
// cable is connected to your router, the other to an AWS Direct Connect
// router.</p> <p>To find the locations for your Region, use
// <a>DescribeLocations</a>.</p> <p>You can automatically add the new connection to
// a link aggregation group (LAG) by specifying a LAG ID in the request. This
// ensures that the new connection is allocated on the same AWS Direct Connect
// endpoint that hosts the specified LAG. If there are no available ports on the
// endpoint, the request fails and no connection is created.</p>
func (c *Client) CreateConnection(ctx context.Context, params *CreateConnectionInput, optFns ...func(*Options)) (*CreateConnectionOutput, error) {
	stack := middleware.NewStack("CreateConnection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateConnectionMiddlewares(stack)
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
	addOpCreateConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnection(options.Region), middleware.Before)

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
			OperationName: "CreateConnection",
			Err:           err,
		}
	}
	out := result.(*CreateConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectionInput struct {
	// The name of the service provider associated with the requested connection.
	ProviderName *string
	// The bandwidth of the connection.
	Bandwidth *string
	// The location of the connection.
	Location *string
	// The tags to associate with the lag.
	Tags []*types.Tag
	// The name of the connection.
	ConnectionName *string
	// The ID of the LAG.
	LagId *string
}

// Information about an AWS Direct Connect connection.
type CreateConnectionOutput struct {
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool
	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string
	// The ID of the VLAN.
	Vlan *int32
	// The name of the connection.
	ConnectionName *string
	// The AWS Region where the connection is located.
	Region *string
	// The ID of the LAG.
	LagId *string
	// The ID of the connection.
	ConnectionId *string
	// The tags associated with the connection.
	Tags []*types.Tag
	// The ID of the AWS account that owns the connection.
	OwnerAccount *string
	// The state of the connection. The following are the possible values:
	//
	//     *
	// ordering: The initial state of a hosted connection provisioned on an
	// interconnect. The connection stays in the ordering state until the owner of the
	// hosted connection confirms or declines the connection order.
	//
	//     * requested:
	// The initial state of a standard connection. The connection stays in the
	// requested state until the Letter of Authorization (LOA) is sent to the
	// customer.
	//
	//     * pending: The connection has been approved and is being
	// initialized.
	//
	//     * available: The network link is up and the connection is
	// ready for use.
	//
	//     * down: The network link is down.
	//
	//     * deleting: The
	// connection is being deleted.
	//
	//     * deleted: The connection has been deleted.
	//
	//
	// * rejected: A hosted connection in the ordering state enters the rejected state
	// if it is deleted by the customer.
	//
	//     * unknown: The state of the connection is
	// not available.
	ConnectionState types.ConnectionState
	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string
	// The name of the AWS Direct Connect service provider associated with the
	// connection.
	PartnerName *string
	// The time of the most recent call to DescribeLoa () for this connection.
	LoaIssueTime *time.Time
	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy
	// The name of the service provider associated with the connection.
	ProviderName *string
	// The bandwidth of the connection.
	Bandwidth *string
	// The location of the connection.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateConnectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConnection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConnection{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreateConnection",
	}
}
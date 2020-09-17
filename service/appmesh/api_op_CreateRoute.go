// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a route that is associated with a virtual router. You can route several
// different protocols and define a retry policy for a route. Traffic can be routed
// to one or more virtual nodes. For more information about routes, see Routes
// (https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html).
func (c *Client) CreateRoute(ctx context.Context, params *CreateRouteInput, optFns ...func(*Options)) (*CreateRouteOutput, error) {
	stack := middleware.NewStack("CreateRoute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateRouteMiddlewares(stack)
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
	addOpCreateRouteValidationMiddleware(stack)

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
			OperationName: "CreateRoute",
			Err:           err,
		}
	}
	out := result.(*CreateRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateRouteInput struct {
	// The name to use for the route.
	RouteName *string
	// The name of the service mesh to create the route in.
	MeshName *string
	// The name of the virtual router in which to create the route. If the virtual
	// router is in a shared mesh, then you must be the owner of the virtual router
	// resource.
	VirtualRouterName *string
	// The route specification to apply.
	Spec *types.RouteSpec
	// Optional metadata that you can apply to the route to assist with categorization
	// and organization. Each tag consists of a key and an optional value, both of
	// which you define. Tag keys can have a maximum character length of 128
	// characters, and tag values can have a maximum length of 256 characters.
	Tags []*types.TagRef
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string
	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then the account that you specify must share the mesh with your account
	// before you can create the resource in the service mesh. For more information
	// about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type CreateRouteOutput struct {
	// The full description of your mesh following the create call.
	Route *types.RouteData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateRouteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoute{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoute{}, middleware.After)
}
// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an asset model. This action can't be undone. You must delete all assets
// created from an asset model before you can delete the model. Also, you can't
// delete an asset model if a parent asset model exists that contains a property
// formula expression that depends on the asset model that you want to delete. For
// more information, see Deleting Assets and Models
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/delete-assets-and-models.html)
// in the AWS IoT SiteWise User Guide.
func (c *Client) DeleteAssetModel(ctx context.Context, params *DeleteAssetModelInput, optFns ...func(*Options)) (*DeleteAssetModelOutput, error) {
	if params == nil {
		params = &DeleteAssetModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAssetModel", params, optFns, addOperationDeleteAssetModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAssetModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAssetModelInput struct {

	// The ID of the asset model to delete.
	//
	// This member is required.
	AssetModelId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
}

type DeleteAssetModelOutput struct {

	// The status of the asset model, which contains a state (DELETING after
	// successfully calling this operation) and any error message.
	//
	// This member is required.
	AssetModelStatus *types.AssetModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAssetModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAssetModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAssetModel{}, middleware.After)
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
	addIdempotencyToken_opDeleteAssetModelMiddleware(stack, options)
	addOpDeleteAssetModelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAssetModel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpDeleteAssetModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteAssetModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteAssetModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteAssetModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteAssetModelInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteAssetModelMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpDeleteAssetModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteAssetModel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DeleteAssetModel",
	}
}

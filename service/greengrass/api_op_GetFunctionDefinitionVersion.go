// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a Lambda function definition version, including
// which Lambda functions are included in the version and their configurations.
func (c *Client) GetFunctionDefinitionVersion(ctx context.Context, params *GetFunctionDefinitionVersionInput, optFns ...func(*Options)) (*GetFunctionDefinitionVersionOutput, error) {
	stack := middleware.NewStack("GetFunctionDefinitionVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetFunctionDefinitionVersionMiddlewares(stack)
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
	addOpGetFunctionDefinitionVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunctionDefinitionVersion(options.Region), middleware.Before)

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
			OperationName: "GetFunctionDefinitionVersion",
			Err:           err,
		}
	}
	out := result.(*GetFunctionDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionDefinitionVersionInput struct {
	// The ID of the function definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListFunctionDefinitionVersions'' requests. If the version is the last one
	// that was associated with a function definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	FunctionDefinitionVersionId *string
	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
	// The ID of the Lambda function definition.
	FunctionDefinitionId *string
}

type GetFunctionDefinitionVersionOutput struct {
	// The ID of the function definition version.
	Id *string
	// Information on the definition.
	Definition *types.FunctionDefinitionVersion
	// The ARN of the function definition version.
	Arn *string
	// The version of the function definition version.
	Version *string
	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
	// The time, in milliseconds since the epoch, when the function definition version
	// was created.
	CreationTimestamp *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetFunctionDefinitionVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetFunctionDefinitionVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunctionDefinitionVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFunctionDefinitionVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetFunctionDefinitionVersion",
	}
}
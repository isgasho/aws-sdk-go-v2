// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/AddPermissionRequest
type AddPermissionInput struct {
	_ struct{} `type:"structure"`

	// The action that the principal can use on the function. For example, lambda:InvokeFunction
	// or lambda:GetFunction.
	//
	// Action is a required field
	Action *string `type:"string" required:"true"`

	// For Alexa Smart Home functions, a token that must be supplied by the invoker.
	EventSourceToken *string `type:"string"`

	// The name of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//    * Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function
	// name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// The AWS service or account that invokes the function. If you specify a service,
	// use SourceArn or SourceAccount to limit who can invoke the function through
	// that service.
	//
	// Principal is a required field
	Principal *string `type:"string" required:"true"`

	// Specify a version or alias to add permissions to a published version of the
	// function.
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string"`

	// Only update the policy if the revision ID matches the ID that's specified.
	// Use this option to avoid modifying a policy that has changed since you last
	// read it.
	RevisionId *string `type:"string"`

	// For AWS services, the ID of the account that owns the resource. Use this
	// instead of SourceArn to grant permission to resources that are owned by another
	// account (for example, all of an account's Amazon S3 buckets). Or use it together
	// with SourceArn to ensure that the resource is owned by the specified account.
	// For example, an Amazon S3 bucket could be deleted by its owner and recreated
	// by another account.
	SourceAccount *string `type:"string"`

	// For AWS services, the ARN of the AWS resource that invokes the function.
	// For example, an Amazon S3 bucket or Amazon SNS topic.
	SourceArn *string `type:"string"`

	// A statement identifier that differentiates the statement from others in the
	// same policy.
	//
	// StatementId is a required field
	StatementId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AddPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddPermissionInput"}

	if s.Action == nil {
		invalidParams.Add(aws.NewErrParamRequired("Action"))
	}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}
	if s.Qualifier != nil && len(*s.Qualifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Qualifier", 1))
	}

	if s.StatementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatementId"))
	}
	if s.StatementId != nil && len(*s.StatementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatementId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddPermissionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Action != nil {
		v := *s.Action

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Action", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EventSourceToken != nil {
		v := *s.EventSourceToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EventSourceToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Principal != nil {
		v := *s.Principal

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Principal", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceAccount != nil {
		v := *s.SourceAccount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SourceAccount", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StatementId != nil {
		v := *s.StatementId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StatementId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Qualifier != nil {
		v := *s.Qualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Qualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/AddPermissionResponse
type AddPermissionOutput struct {
	_ struct{} `type:"structure"`

	// The permission statement that's added to the function policy.
	Statement *string `type:"string"`
}

// String returns the string representation
func (s AddPermissionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddPermissionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Statement != nil {
		v := *s.Statement

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Statement", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opAddPermission = "AddPermission"

// AddPermissionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Grants an AWS service or another account permission to use a function. You
// can apply the policy at the function level, or specify a qualifier to restrict
// access to a single version or alias. If you use a qualifier, the invoker
// must use the full Amazon Resource Name (ARN) of that version or alias to
// invoke the function.
//
// To grant permission to another account, specify the account ID as the Principal.
// For AWS services, the principal is a domain-style identifier defined by the
// service, like s3.amazonaws.com or sns.amazonaws.com. For AWS services, you
// can also specify the ARN or owning account of the associated resource as
// the SourceArn or SourceAccount. If you grant permission to a service principal
// without specifying the source, other accounts could potentially configure
// resources in their account to invoke your Lambda function.
//
// This action adds a statement to a resource-based permission policy for the
// function. For more information about function policies, see Lambda Function
// Policies (https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html).
//
//    // Example sending a request using AddPermissionRequest.
//    req := client.AddPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/AddPermission
func (c *Client) AddPermissionRequest(input *AddPermissionInput) AddPermissionRequest {
	op := &aws.Operation{
		Name:       opAddPermission,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/policy",
	}

	if input == nil {
		input = &AddPermissionInput{}
	}

	req := c.newRequest(op, input, &AddPermissionOutput{})
	return AddPermissionRequest{Request: req, Input: input, Copy: c.AddPermissionRequest}
}

// AddPermissionRequest is the request type for the
// AddPermission API operation.
type AddPermissionRequest struct {
	*aws.Request
	Input *AddPermissionInput
	Copy  func(*AddPermissionInput) AddPermissionRequest
}

// Send marshals and sends the AddPermission API request.
func (r AddPermissionRequest) Send(ctx context.Context) (*AddPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddPermissionResponse{
		AddPermissionOutput: r.Request.Data.(*AddPermissionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddPermissionResponse is the response type for the
// AddPermission API operation.
type AddPermissionResponse struct {
	*AddPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddPermission request.
func (r *AddPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

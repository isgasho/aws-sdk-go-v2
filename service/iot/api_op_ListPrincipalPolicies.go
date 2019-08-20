// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the ListPrincipalPolicies operation.
type ListPrincipalPoliciesInput struct {
	_ struct{} `type:"structure"`

	// Specifies the order for results. If true, results are returned in ascending
	// creation order.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// The marker for the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The result page size.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`

	// The principal.
	//
	// Principal is a required field
	Principal *string `location:"header" locationName:"x-amzn-iot-principal" type:"string" required:"true"`
}

// String returns the string representation
func (s ListPrincipalPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPrincipalPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPrincipalPoliciesInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPrincipalPoliciesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Principal != nil {
		v := *s.Principal

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amzn-iot-principal", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AscendingOrder != nil {
		v := *s.AscendingOrder

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "isAscendingOrder", protocol.BoolValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "pageSize", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The output from the ListPrincipalPolicies operation.
type ListPrincipalPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string `locationName:"nextMarker" type:"string"`

	// The policies.
	Policies []Policy `locationName:"policies" type:"list"`
}

// String returns the string representation
func (s ListPrincipalPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPrincipalPoliciesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policies != nil {
		v := s.Policies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "policies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListPrincipalPolicies = "ListPrincipalPolicies"

// ListPrincipalPoliciesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the policies attached to the specified principal. If you use an Cognito
// identity, the ID must be in AmazonCognito Identity format (https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetCredentialsForIdentity.html#API_GetCredentialsForIdentity_RequestSyntax).
//
// Note: This API is deprecated. Please use ListAttachedPolicies instead.
//
//    // Example sending a request using ListPrincipalPoliciesRequest.
//    req := client.ListPrincipalPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListPrincipalPoliciesRequest(input *ListPrincipalPoliciesInput) ListPrincipalPoliciesRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, ListPrincipalPolicies, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opListPrincipalPolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/principal-policies",
	}

	if input == nil {
		input = &ListPrincipalPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListPrincipalPoliciesOutput{})
	return ListPrincipalPoliciesRequest{Request: req, Input: input, Copy: c.ListPrincipalPoliciesRequest}
}

// ListPrincipalPoliciesRequest is the request type for the
// ListPrincipalPolicies API operation.
type ListPrincipalPoliciesRequest struct {
	*aws.Request
	Input *ListPrincipalPoliciesInput
	Copy  func(*ListPrincipalPoliciesInput) ListPrincipalPoliciesRequest
}

// Send marshals and sends the ListPrincipalPolicies API request.
func (r ListPrincipalPoliciesRequest) Send(ctx context.Context) (*ListPrincipalPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPrincipalPoliciesResponse{
		ListPrincipalPoliciesOutput: r.Request.Data.(*ListPrincipalPoliciesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPrincipalPoliciesResponse is the response type for the
// ListPrincipalPolicies API operation.
type ListPrincipalPoliciesResponse struct {
	*ListPrincipalPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPrincipalPolicies request.
func (r *ListPrincipalPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

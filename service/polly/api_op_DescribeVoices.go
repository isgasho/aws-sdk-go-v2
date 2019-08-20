// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DescribeVoicesInput
type DescribeVoicesInput struct {
	_ struct{} `type:"structure"`

	// Boolean value indicating whether to return any bilingual voices that use
	// the specified language as an additional language. For instance, if you request
	// all languages that use US English (es-US), and there is an Italian voice
	// that speaks both Italian (it-IT) and US English, that voice will be included
	// if you specify yes but not if you specify no.
	IncludeAdditionalLanguageCodes *bool `location:"querystring" locationName:"IncludeAdditionalLanguageCodes" type:"boolean"`

	// The language identification tag (ISO 639 code for the language name-ISO 3166
	// country code) for filtering the list of voices returned. If you don't specify
	// this optional parameter, all available voices are returned.
	LanguageCode LanguageCode `location:"querystring" locationName:"LanguageCode" type:"string" enum:"true"`

	// An opaque pagination token returned from the previous DescribeVoices operation.
	// If present, this indicates where to continue the listing.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeVoicesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeVoicesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.IncludeAdditionalLanguageCodes != nil {
		v := *s.IncludeAdditionalLanguageCodes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "IncludeAdditionalLanguageCodes", protocol.BoolValue(v), metadata)
	}
	if len(s.LanguageCode) > 0 {
		v := s.LanguageCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "LanguageCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DescribeVoicesOutput
type DescribeVoicesOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token to use in the next request to continue the listing of
	// voices. NextToken is returned only if the response is truncated.
	NextToken *string `type:"string"`

	// A list of voices with their properties.
	Voices []Voice `type:"list"`
}

// String returns the string representation
func (s DescribeVoicesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeVoicesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Voices != nil {
		v := s.Voices

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Voices", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeVoices = "DescribeVoices"

// DescribeVoicesRequest returns a request value for making API operation for
// Amazon Polly.
//
// Returns the list of voices that are available for use when requesting speech
// synthesis. Each voice speaks a specified language, is either male or female,
// and is identified by an ID, which is the ASCII version of the voice name.
//
// When synthesizing speech ( SynthesizeSpeech ), you provide the voice ID for
// the voice you want from the list of voices returned by DescribeVoices.
//
// For example, you want your news reader application to read news in a specific
// language, but giving a user the option to choose the voice. Using the DescribeVoices
// operation you can provide the user with a list of available voices to select
// from.
//
// You can optionally specify a language code to filter the available voices.
// For example, if you specify en-US, the operation returns a list of all available
// US English voices.
//
// This operation requires permissions to perform the polly:DescribeVoices action.
//
//    // Example sending a request using DescribeVoicesRequest.
//    req := client.DescribeVoicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DescribeVoices
func (c *Client) DescribeVoicesRequest(input *DescribeVoicesInput) DescribeVoicesRequest {
	op := &aws.Operation{
		Name:       opDescribeVoices,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/voices",
	}

	if input == nil {
		input = &DescribeVoicesInput{}
	}

	req := c.newRequest(op, input, &DescribeVoicesOutput{})
	return DescribeVoicesRequest{Request: req, Input: input, Copy: c.DescribeVoicesRequest}
}

// DescribeVoicesRequest is the request type for the
// DescribeVoices API operation.
type DescribeVoicesRequest struct {
	*aws.Request
	Input *DescribeVoicesInput
	Copy  func(*DescribeVoicesInput) DescribeVoicesRequest
}

// Send marshals and sends the DescribeVoices API request.
func (r DescribeVoicesRequest) Send(ctx context.Context) (*DescribeVoicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVoicesResponse{
		DescribeVoicesOutput: r.Request.Data.(*DescribeVoicesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVoicesResponse is the response type for the
// DescribeVoices API operation.
type DescribeVoicesResponse struct {
	*DescribeVoicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVoices request.
func (r *DescribeVoicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

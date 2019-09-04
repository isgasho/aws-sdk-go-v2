// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteObjectsRequest
type DeleteObjectsInput struct {
	_ struct{} `type:"structure" payload:"Delete"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies whether you want to delete this object even if it has a Governance-type
	// object lock in place. You must have sufficient permissions to perform this
	// operation.
	BypassGovernanceRetention *bool `location:"header" locationName:"x-amz-bypass-governance-retention" type:"boolean"`

	// Delete is a required field
	Delete *Delete `locationName:"Delete" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// The concatenation of the authentication device's serial number, a space,
	// and the value that is displayed on your authentication device.
	MFA *string `location:"header" locationName:"x-amz-mfa" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteObjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteObjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteObjectsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Delete == nil {
		invalidParams.Add(aws.NewErrParamRequired("Delete"))
	}
	if s.Delete != nil {
		if err := s.Delete.Validate(); err != nil {
			invalidParams.AddNested("Delete", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteObjectsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.BypassGovernanceRetention != nil {
		v := *s.BypassGovernanceRetention

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bypass-governance-retention", protocol.BoolValue(v), metadata)
	}
	if s.MFA != nil {
		v := *s.MFA

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-mfa", protocol.StringValue(v), metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Delete != nil {
		v := s.Delete

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "Delete", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteObjectsOutput
type DeleteObjectsOutput struct {
	_ struct{} `type:"structure"`

	Deleted []DeletedObject `type:"list" flattened:"true"`

	Errors []Error `locationName:"Error" type:"list" flattened:"true"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteObjectsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Deleted != nil {
		v := s.Deleted

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Deleted", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Errors != nil {
		v := s.Errors

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Error", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}

// UnmarshalAWSXML decodes the AWS API shape using the passed in *xml.Decoder.
func (s *DeleteObjectsOutput) UnmarshalAWSXML(d *xml.Decoder) (err error) {
	defer func() {
		if err != nil {
			*s = DeleteObjectsOutput{}
		}
	}()
	for {
		tok, err := d.Token()
		if tok == nil || err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("fail to UnmarshalAWSXML DeleteObjectsOutput, %s", err)
		}
		start, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		err = s.unmarshalAWSXML(d, start)
		if err != nil {
			return fmt.Errorf("fail to UnmarshalAWSXML DeleteObjectsOutput, %s", err)
		}
		return nil
	}
}

func (s *DeleteObjectsOutput) unmarshalAWSXML(d *xml.Decoder, head xml.StartElement) (err error) {
	defer func() {
		if err != nil {
			*s = DeleteObjectsOutput{}
		}
	}()
	name := ""
	for {
		tok, err := d.Token()
		if tok == nil || err != nil {
			if err == io.EOF {
				return nil
			}
		}
		if end, ok := tok.(xml.EndElement); ok {
			name = end.Name.Local
			if name == head.Name.Local {
				return nil
			}
		}
		if start, ok := tok.(xml.StartElement); ok {
			switch name = start.Name.Local; name {
			case "Deleted":
				if s.Deleted == nil {
					s.Deleted = make([]DeletedObject, 0)
				}
				err := unmarshalAWSXMLFlattenedListDeletedObjects(&s.Deleted, d, start)
				if err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML DeleteObjectsOutput.%s, %s", name, err)
				}
			case "Error":
				if s.Errors == nil {
					s.Errors = make([]Error, 0)
				}
				err := unmarshalAWSXMLFlattenedListErrors(&s.Errors, d, start)
				if err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML DeleteObjectsOutput.%s, %s", name, err)
				}
			case "x-amz-request-charged":
				tok, err = d.Token()
				if tok == nil || err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML DeleteObjectsOutput.%s, %s", name, err)
				}
				v, _ := tok.(xml.CharData)
				value := RequestCharged(v)
				s.RequestCharged = value
			default:
				err := d.Skip()
				if err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML DeleteObjectsOutput.%s, %s", name, err)
				}
			}
		}
	}
}

// UnmarshalAWSREST decodes the AWS API shape using the passed in *http.Response.
func (s *DeleteObjectsOutput) UnmarshalAWSREST(r *http.Response) (err error) {
	defer func() {
		if err != nil {
			*s = DeleteObjectsOutput{}
		}
	}()
	for k, v := range r.Header {
		switch {
		case strings.EqualFold(k, "x-amz-request-charged"):
			value := RequestCharged(v[0])
			s.RequestCharged = value
		}
	}
	return nil
}

const opDeleteObjects = "DeleteObjects"

// DeleteObjectsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation enables you to delete multiple objects from a bucket using
// a single HTTP request. You may specify up to 1000 keys.
//
//    // Example sending a request using DeleteObjectsRequest.
//    req := client.DeleteObjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteObjects
func (c *Client) DeleteObjectsRequest(input *DeleteObjectsInput) DeleteObjectsRequest {
	op := &aws.Operation{
		Name:       opDeleteObjects,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}?delete",
	}

	if input == nil {
		input = &DeleteObjectsInput{}
	}

	req := c.newRequest(op, input, &DeleteObjectsOutput{})
	return DeleteObjectsRequest{Request: req, Input: input, Copy: c.DeleteObjectsRequest}
}

// DeleteObjectsRequest is the request type for the
// DeleteObjects API operation.
type DeleteObjectsRequest struct {
	*aws.Request
	Input *DeleteObjectsInput
	Copy  func(*DeleteObjectsInput) DeleteObjectsRequest
}

// Send marshals and sends the DeleteObjects API request.
func (r DeleteObjectsRequest) Send(ctx context.Context) (*DeleteObjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteObjectsResponse{
		DeleteObjectsOutput: r.Request.Data.(*DeleteObjectsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteObjectsResponse is the response type for the
// DeleteObjects API operation.
type DeleteObjectsResponse struct {
	*DeleteObjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteObjects request.
func (r *DeleteObjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
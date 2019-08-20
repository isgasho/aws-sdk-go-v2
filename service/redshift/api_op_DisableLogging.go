// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DisableLoggingMessage
type DisableLoggingInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the cluster on which logging is to be stopped.
	//
	// Example: examplecluster
	//
	// ClusterIdentifier is a required field
	ClusterIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisableLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableLoggingInput"}

	if s.ClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the status of logging for a cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/LoggingStatus
type DisableLoggingOutput struct {
	_ struct{} `type:"structure"`

	// The name of the S3 bucket where the log files are stored.
	BucketName *string `type:"string"`

	// The message indicating that logs failed to be delivered.
	LastFailureMessage *string `type:"string"`

	// The last time when logs failed to be delivered.
	LastFailureTime *time.Time `type:"timestamp"`

	// The last time that logs were delivered.
	LastSuccessfulDeliveryTime *time.Time `type:"timestamp"`

	// true if logging is on, false if logging is off.
	LoggingEnabled *bool `type:"boolean"`

	// The prefix applied to the log file names.
	S3KeyPrefix *string `type:"string"`
}

// String returns the string representation
func (s DisableLoggingOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableLogging = "DisableLogging"

// DisableLoggingRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Stops logging information, such as queries and connection attempts, for the
// specified Amazon Redshift cluster.
//
//    // Example sending a request using DisableLoggingRequest.
//    req := client.DisableLoggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DisableLogging
func (c *Client) DisableLoggingRequest(input *DisableLoggingInput) DisableLoggingRequest {
	op := &aws.Operation{
		Name:       opDisableLogging,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableLoggingInput{}
	}

	req := c.newRequest(op, input, &DisableLoggingOutput{})
	return DisableLoggingRequest{Request: req, Input: input, Copy: c.DisableLoggingRequest}
}

// DisableLoggingRequest is the request type for the
// DisableLogging API operation.
type DisableLoggingRequest struct {
	*aws.Request
	Input *DisableLoggingInput
	Copy  func(*DisableLoggingInput) DisableLoggingRequest
}

// Send marshals and sends the DisableLogging API request.
func (r DisableLoggingRequest) Send(ctx context.Context) (*DisableLoggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableLoggingResponse{
		DisableLoggingOutput: r.Request.Data.(*DisableLoggingOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableLoggingResponse is the response type for the
// DisableLogging API operation.
type DisableLoggingResponse struct {
	*DisableLoggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableLogging request.
func (r *DisableLoggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeClustersRequest
type DescribeClustersInput struct {
	_ struct{} `type:"structure"`

	// The names of the DAX clusters being described.
	ClusterNames []string `type:"list"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	//
	// The value for MaxResults must be between 20 and 100.
	MaxResults *int64 `type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeClustersInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeClustersResponse
type DescribeClustersOutput struct {
	_ struct{} `type:"structure"`

	// The descriptions of your DAX clusters, in response to a DescribeClusters
	// request.
	Clusters []Cluster `type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeClustersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClusters = "DescribeClusters"

// DescribeClustersRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Returns information about all provisioned DAX clusters if no cluster identifier
// is specified, or about a specific DAX cluster if a cluster identifier is
// supplied.
//
// If the cluster is in the CREATING state, only cluster level information will
// be displayed until all of the nodes are successfully provisioned.
//
// If the cluster is in the DELETING state, only cluster level information will
// be displayed.
//
// If nodes are currently being added to the DAX cluster, node endpoint information
// and creation time for the additional nodes will not be displayed until they
// are completely provisioned. When the DAX cluster state is available, the
// cluster is ready for use.
//
// If nodes are currently being removed from the DAX cluster, no endpoint information
// for the removed nodes is displayed.
//
//    // Example sending a request using DescribeClustersRequest.
//    req := client.DescribeClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeClusters
func (c *Client) DescribeClustersRequest(input *DescribeClustersInput) DescribeClustersRequest {
	op := &aws.Operation{
		Name:       opDescribeClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeClustersInput{}
	}

	req := c.newRequest(op, input, &DescribeClustersOutput{})
	return DescribeClustersRequest{Request: req, Input: input, Copy: c.DescribeClustersRequest}
}

// DescribeClustersRequest is the request type for the
// DescribeClusters API operation.
type DescribeClustersRequest struct {
	*aws.Request
	Input *DescribeClustersInput
	Copy  func(*DescribeClustersInput) DescribeClustersRequest
}

// Send marshals and sends the DescribeClusters API request.
func (r DescribeClustersRequest) Send(ctx context.Context) (*DescribeClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClustersResponse{
		DescribeClustersOutput: r.Request.Data.(*DescribeClustersOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeClustersResponse is the response type for the
// DescribeClusters API operation.
type DescribeClustersResponse struct {
	*DescribeClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClusters request.
func (r *DescribeClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
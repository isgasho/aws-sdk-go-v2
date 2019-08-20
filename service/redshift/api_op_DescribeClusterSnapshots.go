// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeClusterSnapshotsMessage
type DescribeClusterSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// A value that indicates whether to return snapshots only for an existing cluster.
	// You can perform table-level restore only by using a snapshot of an existing
	// cluster, that is, a cluster that has not been deleted. Values for this parameter
	// work as follows:
	//
	//    * If ClusterExists is set to true, ClusterIdentifier is required.
	//
	//    * If ClusterExists is set to false and ClusterIdentifier isn't specified,
	//    all snapshots associated with deleted clusters (orphaned snapshots) are
	//    returned.
	//
	//    * If ClusterExists is set to false and ClusterIdentifier is specified
	//    for a deleted cluster, snapshots associated with that cluster are returned.
	//
	//    * If ClusterExists is set to false and ClusterIdentifier is specified
	//    for an existing cluster, no snapshots are returned.
	ClusterExists *bool `type:"boolean"`

	// The identifier of the cluster which generated the requested snapshots.
	ClusterIdentifier *string `type:"string"`

	// A time value that requests only snapshots created at or before the specified
	// time. The time value is specified in ISO 8601 format. For more information
	// about ISO 8601, go to the ISO8601 Wikipedia page. (http://en.wikipedia.org/wiki/ISO_8601)
	//
	// Example: 2012-07-16T18:00:00Z
	EndTime *time.Time `type:"timestamp"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterSnapshots request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The AWS customer account used to create or copy the snapshot. Use this field
	// to filter the results to snapshots owned by a particular account. To describe
	// snapshots you own, either specify your AWS customer account, or do not specify
	// the parameter.
	OwnerAccount *string `type:"string"`

	// The snapshot identifier of the snapshot about which to return information.
	SnapshotIdentifier *string `type:"string"`

	// The type of snapshots for which you are requesting information. By default,
	// snapshots of all types are returned.
	//
	// Valid Values: automated | manual
	SnapshotType *string `type:"string"`

	SortingEntities []SnapshotSortingEntity `locationNameList:"SnapshotSortingEntity" type:"list"`

	// A value that requests only snapshots created at or after the specified time.
	// The time value is specified in ISO 8601 format. For more information about
	// ISO 8601, go to the ISO8601 Wikipedia page. (http://en.wikipedia.org/wiki/ISO_8601)
	//
	// Example: 2012-07-16T18:00:00Z
	StartTime *time.Time `type:"timestamp"`

	// A tag key or keys for which you want to return all matching cluster snapshots
	// that are associated with the specified key or keys. For example, suppose
	// that you have snapshots that are tagged with keys called owner and environment.
	// If you specify both of these tag keys in the request, Amazon Redshift returns
	// a response with the snapshots that have either or both of these tag keys
	// associated with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching cluster snapshots
	// that are associated with the specified tag value or values. For example,
	// suppose that you have snapshots that are tagged with values called admin
	// and test. If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the snapshots that have either or both of
	// these tag values associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeClusterSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClusterSnapshotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClusterSnapshotsInput"}
	if s.SortingEntities != nil {
		for i, v := range s.SortingEntities {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SortingEntities", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output from the DescribeClusterSnapshots action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/SnapshotMessage
type DescribeClusterSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// A list of Snapshot instances.
	Snapshots []Snapshot `locationNameList:"Snapshot" type:"list"`
}

// String returns the string representation
func (s DescribeClusterSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClusterSnapshots = "DescribeClusterSnapshots"

// DescribeClusterSnapshotsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns one or more snapshot objects, which contain metadata about your cluster
// snapshots. By default, this operation returns information about all snapshots
// of all clusters that are owned by you AWS customer account. No information
// is returned for snapshots owned by inactive AWS customer accounts.
//
// If you specify both tag keys and tag values in the same request, Amazon Redshift
// returns all snapshots that match any combination of the specified keys and
// values. For example, if you have owner and environment for tag keys, and
// admin and test for tag values, all snapshots that have any combination of
// those values are returned. Only snapshots that you own are returned in the
// response; shared snapshots are not returned with the tag key and tag value
// request parameters.
//
// If both tag keys and values are omitted from the request, snapshots are returned
// regardless of whether they have tag keys or values associated with them.
//
//    // Example sending a request using DescribeClusterSnapshotsRequest.
//    req := client.DescribeClusterSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeClusterSnapshots
func (c *Client) DescribeClusterSnapshotsRequest(input *DescribeClusterSnapshotsInput) DescribeClusterSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeClusterSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeClusterSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DescribeClusterSnapshotsOutput{})
	return DescribeClusterSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeClusterSnapshotsRequest}
}

// DescribeClusterSnapshotsRequest is the request type for the
// DescribeClusterSnapshots API operation.
type DescribeClusterSnapshotsRequest struct {
	*aws.Request
	Input *DescribeClusterSnapshotsInput
	Copy  func(*DescribeClusterSnapshotsInput) DescribeClusterSnapshotsRequest
}

// Send marshals and sends the DescribeClusterSnapshots API request.
func (r DescribeClusterSnapshotsRequest) Send(ctx context.Context) (*DescribeClusterSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClusterSnapshotsResponse{
		DescribeClusterSnapshotsOutput: r.Request.Data.(*DescribeClusterSnapshotsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClusterSnapshotsRequestPaginator returns a paginator for DescribeClusterSnapshots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClusterSnapshotsRequest(input)
//   p := redshift.NewDescribeClusterSnapshotsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClusterSnapshotsPaginator(req DescribeClusterSnapshotsRequest) DescribeClusterSnapshotsPaginator {
	return DescribeClusterSnapshotsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClusterSnapshotsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeClusterSnapshotsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClusterSnapshotsPaginator struct {
	aws.Pager
}

func (p *DescribeClusterSnapshotsPaginator) CurrentPage() *DescribeClusterSnapshotsOutput {
	return p.Pager.CurrentPage().(*DescribeClusterSnapshotsOutput)
}

// DescribeClusterSnapshotsResponse is the response type for the
// DescribeClusterSnapshots API operation.
type DescribeClusterSnapshotsResponse struct {
	*DescribeClusterSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClusterSnapshots request.
func (r *DescribeClusterSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

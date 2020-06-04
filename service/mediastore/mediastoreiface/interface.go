// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediastoreiface provides an interface to enable mocking the AWS Elemental MediaStore service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediastoreiface

import (
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
)

// ClientAPI provides an interface to enable mocking the
// mediastore.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // MediaStore.
//    func myFunc(svc mediastoreiface.ClientAPI) bool {
//        // Make svc.CreateContainer request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := mediastore.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        mediastoreiface.ClientPI
//    }
//    func (m *mockClientClient) CreateContainer(input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CreateContainerRequest(*mediastore.CreateContainerInput) mediastore.CreateContainerRequest

	DeleteContainerRequest(*mediastore.DeleteContainerInput) mediastore.DeleteContainerRequest

	DeleteContainerPolicyRequest(*mediastore.DeleteContainerPolicyInput) mediastore.DeleteContainerPolicyRequest

	DeleteCorsPolicyRequest(*mediastore.DeleteCorsPolicyInput) mediastore.DeleteCorsPolicyRequest

	DeleteLifecyclePolicyRequest(*mediastore.DeleteLifecyclePolicyInput) mediastore.DeleteLifecyclePolicyRequest

	DeleteMetricPolicyRequest(*mediastore.DeleteMetricPolicyInput) mediastore.DeleteMetricPolicyRequest

	DescribeContainerRequest(*mediastore.DescribeContainerInput) mediastore.DescribeContainerRequest

	GetContainerPolicyRequest(*mediastore.GetContainerPolicyInput) mediastore.GetContainerPolicyRequest

	GetCorsPolicyRequest(*mediastore.GetCorsPolicyInput) mediastore.GetCorsPolicyRequest

	GetLifecyclePolicyRequest(*mediastore.GetLifecyclePolicyInput) mediastore.GetLifecyclePolicyRequest

	GetMetricPolicyRequest(*mediastore.GetMetricPolicyInput) mediastore.GetMetricPolicyRequest

	ListContainersRequest(*mediastore.ListContainersInput) mediastore.ListContainersRequest

	ListTagsForResourceRequest(*mediastore.ListTagsForResourceInput) mediastore.ListTagsForResourceRequest

	PutContainerPolicyRequest(*mediastore.PutContainerPolicyInput) mediastore.PutContainerPolicyRequest

	PutCorsPolicyRequest(*mediastore.PutCorsPolicyInput) mediastore.PutCorsPolicyRequest

	PutLifecyclePolicyRequest(*mediastore.PutLifecyclePolicyInput) mediastore.PutLifecyclePolicyRequest

	PutMetricPolicyRequest(*mediastore.PutMetricPolicyInput) mediastore.PutMetricPolicyRequest

	StartAccessLoggingRequest(*mediastore.StartAccessLoggingInput) mediastore.StartAccessLoggingRequest

	StopAccessLoggingRequest(*mediastore.StopAccessLoggingInput) mediastore.StopAccessLoggingRequest

	TagResourceRequest(*mediastore.TagResourceInput) mediastore.TagResourceRequest

	UntagResourceRequest(*mediastore.UntagResourceInput) mediastore.UntagResourceRequest
}

var _ ClientAPI = (*mediastore.Client)(nil)
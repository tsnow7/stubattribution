// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudfrontiface provides an interface to enable mocking the Amazon CloudFront service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudfrontiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

// CloudFrontAPI provides an interface to enable mocking the
// cloudfront.CloudFront service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudFront.
//    func myFunc(svc cloudfrontiface.CloudFrontAPI) bool {
//        // Make svc.CreateCloudFrontOriginAccessIdentity request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudfront.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudFrontClient struct {
//        cloudfrontiface.CloudFrontAPI
//    }
//    func (m *mockCloudFrontClient) CreateCloudFrontOriginAccessIdentity(input *cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudFrontClient{}
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
type CloudFrontAPI interface {
	CreateCloudFrontOriginAccessIdentityRequest(*cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.CreateCloudFrontOriginAccessIdentityOutput)

	CreateCloudFrontOriginAccessIdentity(*cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error)

	CreateDistributionRequest(*cloudfront.CreateDistributionInput) (*request.Request, *cloudfront.CreateDistributionOutput)

	CreateDistribution(*cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error)

	CreateDistributionWithTagsRequest(*cloudfront.CreateDistributionWithTagsInput) (*request.Request, *cloudfront.CreateDistributionWithTagsOutput)

	CreateDistributionWithTags(*cloudfront.CreateDistributionWithTagsInput) (*cloudfront.CreateDistributionWithTagsOutput, error)

	CreateInvalidationRequest(*cloudfront.CreateInvalidationInput) (*request.Request, *cloudfront.CreateInvalidationOutput)

	CreateInvalidation(*cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error)

	CreateStreamingDistributionRequest(*cloudfront.CreateStreamingDistributionInput) (*request.Request, *cloudfront.CreateStreamingDistributionOutput)

	CreateStreamingDistribution(*cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error)

	CreateStreamingDistributionWithTagsRequest(*cloudfront.CreateStreamingDistributionWithTagsInput) (*request.Request, *cloudfront.CreateStreamingDistributionWithTagsOutput)

	CreateStreamingDistributionWithTags(*cloudfront.CreateStreamingDistributionWithTagsInput) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error)

	DeleteCloudFrontOriginAccessIdentityRequest(*cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.DeleteCloudFrontOriginAccessIdentityOutput)

	DeleteCloudFrontOriginAccessIdentity(*cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error)

	DeleteDistributionRequest(*cloudfront.DeleteDistributionInput) (*request.Request, *cloudfront.DeleteDistributionOutput)

	DeleteDistribution(*cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error)

	DeleteStreamingDistributionRequest(*cloudfront.DeleteStreamingDistributionInput) (*request.Request, *cloudfront.DeleteStreamingDistributionOutput)

	DeleteStreamingDistribution(*cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error)

	GetCloudFrontOriginAccessIdentityRequest(*cloudfront.GetCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.GetCloudFrontOriginAccessIdentityOutput)

	GetCloudFrontOriginAccessIdentity(*cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)

	GetCloudFrontOriginAccessIdentityConfigRequest(*cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*request.Request, *cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput)

	GetCloudFrontOriginAccessIdentityConfig(*cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)

	GetDistributionRequest(*cloudfront.GetDistributionInput) (*request.Request, *cloudfront.GetDistributionOutput)

	GetDistribution(*cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error)

	GetDistributionConfigRequest(*cloudfront.GetDistributionConfigInput) (*request.Request, *cloudfront.GetDistributionConfigOutput)

	GetDistributionConfig(*cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error)

	GetInvalidationRequest(*cloudfront.GetInvalidationInput) (*request.Request, *cloudfront.GetInvalidationOutput)

	GetInvalidation(*cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error)

	GetStreamingDistributionRequest(*cloudfront.GetStreamingDistributionInput) (*request.Request, *cloudfront.GetStreamingDistributionOutput)

	GetStreamingDistribution(*cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error)

	GetStreamingDistributionConfigRequest(*cloudfront.GetStreamingDistributionConfigInput) (*request.Request, *cloudfront.GetStreamingDistributionConfigOutput)

	GetStreamingDistributionConfig(*cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error)

	ListCloudFrontOriginAccessIdentitiesRequest(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*request.Request, *cloudfront.ListCloudFrontOriginAccessIdentitiesOutput)

	ListCloudFrontOriginAccessIdentities(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)

	ListCloudFrontOriginAccessIdentitiesPages(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput, func(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, bool) bool) error

	ListDistributionsRequest(*cloudfront.ListDistributionsInput) (*request.Request, *cloudfront.ListDistributionsOutput)

	ListDistributions(*cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error)

	ListDistributionsPages(*cloudfront.ListDistributionsInput, func(*cloudfront.ListDistributionsOutput, bool) bool) error

	ListDistributionsByWebACLIdRequest(*cloudfront.ListDistributionsByWebACLIdInput) (*request.Request, *cloudfront.ListDistributionsByWebACLIdOutput)

	ListDistributionsByWebACLId(*cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error)

	ListInvalidationsRequest(*cloudfront.ListInvalidationsInput) (*request.Request, *cloudfront.ListInvalidationsOutput)

	ListInvalidations(*cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error)

	ListInvalidationsPages(*cloudfront.ListInvalidationsInput, func(*cloudfront.ListInvalidationsOutput, bool) bool) error

	ListStreamingDistributionsRequest(*cloudfront.ListStreamingDistributionsInput) (*request.Request, *cloudfront.ListStreamingDistributionsOutput)

	ListStreamingDistributions(*cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error)

	ListStreamingDistributionsPages(*cloudfront.ListStreamingDistributionsInput, func(*cloudfront.ListStreamingDistributionsOutput, bool) bool) error

	ListTagsForResourceRequest(*cloudfront.ListTagsForResourceInput) (*request.Request, *cloudfront.ListTagsForResourceOutput)

	ListTagsForResource(*cloudfront.ListTagsForResourceInput) (*cloudfront.ListTagsForResourceOutput, error)

	TagResourceRequest(*cloudfront.TagResourceInput) (*request.Request, *cloudfront.TagResourceOutput)

	TagResource(*cloudfront.TagResourceInput) (*cloudfront.TagResourceOutput, error)

	UntagResourceRequest(*cloudfront.UntagResourceInput) (*request.Request, *cloudfront.UntagResourceOutput)

	UntagResource(*cloudfront.UntagResourceInput) (*cloudfront.UntagResourceOutput, error)

	UpdateCloudFrontOriginAccessIdentityRequest(*cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.UpdateCloudFrontOriginAccessIdentityOutput)

	UpdateCloudFrontOriginAccessIdentity(*cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error)

	UpdateDistributionRequest(*cloudfront.UpdateDistributionInput) (*request.Request, *cloudfront.UpdateDistributionOutput)

	UpdateDistribution(*cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error)

	UpdateStreamingDistributionRequest(*cloudfront.UpdateStreamingDistributionInput) (*request.Request, *cloudfront.UpdateStreamingDistributionOutput)

	UpdateStreamingDistribution(*cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error)

	WaitUntilDistributionDeployed(*cloudfront.GetDistributionInput) error

	WaitUntilInvalidationCompleted(*cloudfront.GetInvalidationInput) error

	WaitUntilStreamingDistributionDeployed(*cloudfront.GetStreamingDistributionInput) error
}

var _ CloudFrontAPI = (*cloudfront.CloudFront)(nil)

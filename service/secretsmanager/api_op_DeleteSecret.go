// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/DeleteSecretRequest
type DeleteSecretInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Specifies that the secret is to be deleted without any recovery
	// window. You can't use both this parameter and the RecoveryWindowInDays parameter
	// in the same API call.
	//
	// An asynchronous background process performs the actual deletion, so there
	// can be a short delay before the operation completes. If you write code to
	// delete and then immediately recreate a secret with the same name, ensure
	// that your code includes appropriate back off and retry logic.
	//
	// Use this parameter with caution. This parameter causes the operation to skip
	// the normal waiting period before the permanent deletion that AWS would normally
	// impose with the RecoveryWindowInDays parameter. If you delete a secret with
	// the ForceDeleteWithouRecovery parameter, then you have no opportunity to
	// recover the secret. It is permanently lost.
	ForceDeleteWithoutRecovery *bool `type:"boolean"`

	// (Optional) Specifies the number of days that Secrets Manager waits before
	// it can delete the secret. You can't use both this parameter and the ForceDeleteWithoutRecovery
	// parameter in the same API call.
	//
	// This value can range from 7 to 30 days. The default value is 30.
	RecoveryWindowInDays *int64 `type:"long"`

	// Specifies the secret that you want to delete. You can specify either the
	// Amazon Resource Name (ARN) or the friendly name of the secret.
	//
	// If you specify an ARN, we generally recommend that you specify a complete
	// ARN. You can specify a partial ARN too—for example, if you don’t include
	// the final hyphen and six random characters that Secrets Manager adds at the
	// end of the ARN when you created the secret. A partial ARN match can work
	// as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as
	// a partial ARN, then those characters cause Secrets Manager to assume that
	// you’re specifying a complete ARN. This confusion can cause unexpected results.
	// To avoid this situation, we recommend that you don’t create secret names
	// that end with a hyphen followed by six characters.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSecretInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSecretInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSecretInput"}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/DeleteSecretResponse
type DeleteSecretOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret that is now scheduled for deletion.
	ARN *string `min:"20" type:"string"`

	// The date and time after which this secret can be deleted by Secrets Manager
	// and can no longer be restored. This value is the date and time of the delete
	// request plus the number of days specified in RecoveryWindowInDays.
	DeletionDate *time.Time `type:"timestamp"`

	// The friendly name of the secret that is now scheduled for deletion.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteSecretOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSecret = "DeleteSecret"

// DeleteSecretRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Deletes an entire secret and all of its versions. You can optionally include
// a recovery window during which you can restore the secret. If you don't specify
// a recovery window value, the operation defaults to 30 days. Secrets Manager
// attaches a DeletionDate stamp to the secret that specifies the end of the
// recovery window. At the end of the recovery window, Secrets Manager deletes
// the secret permanently.
//
// At any time before recovery window ends, you can use RestoreSecret to remove
// the DeletionDate and cancel the deletion of the secret.
//
// You cannot access the encrypted secret information in any secret that is
// scheduled for deletion. If you need to access that information, you must
// cancel the deletion with RestoreSecret and then retrieve the information.
//
//    * There is no explicit operation to delete a version of a secret. Instead,
//    remove all staging labels from the VersionStage field of a version. That
//    marks the version as deprecated and allows Secrets Manager to delete it
//    as needed. Versions that do not have any staging labels do not show up
//    in ListSecretVersionIds unless you specify IncludeDeprecated.
//
//    * The permanent secret deletion at the end of the waiting period is performed
//    as a background task with low priority. There is no guarantee of a specific
//    time after the recovery window for the actual delete operation to occur.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:DeleteSecret
//
// Related operations
//
//    * To create a secret, use CreateSecret.
//
//    * To cancel deletion of a version of a secret before the recovery window
//    has expired, use RestoreSecret.
//
//    // Example sending a request using DeleteSecretRequest.
//    req := client.DeleteSecretRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/DeleteSecret
func (c *Client) DeleteSecretRequest(input *DeleteSecretInput) DeleteSecretRequest {
	op := &aws.Operation{
		Name:       opDeleteSecret,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSecretInput{}
	}

	req := c.newRequest(op, input, &DeleteSecretOutput{})
	return DeleteSecretRequest{Request: req, Input: input, Copy: c.DeleteSecretRequest}
}

// DeleteSecretRequest is the request type for the
// DeleteSecret API operation.
type DeleteSecretRequest struct {
	*aws.Request
	Input *DeleteSecretInput
	Copy  func(*DeleteSecretInput) DeleteSecretRequest
}

// Send marshals and sends the DeleteSecret API request.
func (r DeleteSecretRequest) Send(ctx context.Context) (*DeleteSecretResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSecretResponse{
		DeleteSecretOutput: r.Request.Data.(*DeleteSecretOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSecretResponse is the response type for the
// DeleteSecret API operation.
type DeleteSecretResponse struct {
	*DeleteSecretOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSecret request.
func (r *DeleteSecretResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// Job creation failed. Currently, clusters support five nodes. If you have less
// than five nodes for your cluster and you have more nodes to create for this
// cluster, try again and create jobs until your cluster has exactly five notes.
type ClusterLimitExceededException struct {
	Message *string
}

func (e *ClusterLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterLimitExceededException) ErrorCode() string             { return "ClusterLimitExceededException" }
func (e *ClusterLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your IAM user lacks the necessary Amazon EC2 permissions to perform the
// attempted action.
type Ec2RequestFailedException struct {
	Message *string
}

func (e *Ec2RequestFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *Ec2RequestFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *Ec2RequestFailedException) ErrorCode() string             { return "Ec2RequestFailedException" }
func (e *Ec2RequestFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The address provided was invalid. Check the address with your region's carrier,
// and try again.
type InvalidAddressException struct {
	Message *string
}

func (e *InvalidAddressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAddressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAddressException) ErrorCode() string             { return "InvalidAddressException" }
func (e *InvalidAddressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Job or cluster creation failed. One or more inputs were invalid. Confirm that
// the CreateClusterRequest$SnowballType value supports your
// CreateJobRequest$JobType, and try again.
type InvalidInputCombinationException struct {
	Message *string
}

func (e *InvalidInputCombinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputCombinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputCombinationException) ErrorCode() string {
	return "InvalidInputCombinationException"
}
func (e *InvalidInputCombinationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The action can't be performed because the job's current state doesn't allow that
// action to be performed.
type InvalidJobStateException struct {
	Message *string
}

func (e *InvalidJobStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidJobStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidJobStateException) ErrorCode() string             { return "InvalidJobStateException" }
func (e *InvalidJobStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The NextToken string was altered unexpectedly, and the operation has stopped.
// Run the operation without changing the NextToken string, and try again.
type InvalidNextTokenException struct {
	Message *string
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource can't be found. Check the information you provided in
// your last request, and try again.
type InvalidResourceException struct {
	Message *string

	ResourceType *string
}

func (e *InvalidResourceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidResourceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidResourceException) ErrorCode() string             { return "InvalidResourceException" }
func (e *InvalidResourceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided AWS Key Management Service key lacks the permissions to perform the
// specified CreateJob or UpdateJob action.
type KMSRequestFailedException struct {
	Message *string
}

func (e *KMSRequestFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSRequestFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSRequestFailedException) ErrorCode() string             { return "KMSRequestFailedException" }
func (e *KMSRequestFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The address is either outside the serviceable area for your region, or an error
// occurred. Check the address with your region's carrier and try again. If the
// issue persists, contact AWS Support.
type UnsupportedAddressException struct {
	Message *string
}

func (e *UnsupportedAddressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedAddressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedAddressException) ErrorCode() string             { return "UnsupportedAddressException" }
func (e *UnsupportedAddressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

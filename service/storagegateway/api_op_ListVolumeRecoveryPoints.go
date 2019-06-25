// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ListVolumeRecoveryPointsInput
type ListVolumeRecoveryPointsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s ListVolumeRecoveryPointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVolumeRecoveryPointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVolumeRecoveryPointsInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ListVolumeRecoveryPointsOutput
type ListVolumeRecoveryPointsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	GatewayARN *string `min:"50" type:"string"`

	// An array of VolumeRecoveryPointInfo objects.
	VolumeRecoveryPointInfos []VolumeRecoveryPointInfo `type:"list"`
}

// String returns the string representation
func (s ListVolumeRecoveryPointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListVolumeRecoveryPoints = "ListVolumeRecoveryPoints"

// ListVolumeRecoveryPointsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Lists the recovery points for a specified gateway. This operation is only
// supported in the cached volume gateway type.
//
// Each cache volume has one recovery point. A volume recovery point is a point
// in time at which all data of the volume is consistent and from which you
// can create a snapshot or clone a new cached volume from a source volume.
// To create a snapshot from a volume recovery point use the CreateSnapshotFromVolumeRecoveryPoint
// operation.
//
//    // Example sending a request using ListVolumeRecoveryPointsRequest.
//    req := client.ListVolumeRecoveryPointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ListVolumeRecoveryPoints
func (c *Client) ListVolumeRecoveryPointsRequest(input *ListVolumeRecoveryPointsInput) ListVolumeRecoveryPointsRequest {
	op := &aws.Operation{
		Name:       opListVolumeRecoveryPoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListVolumeRecoveryPointsInput{}
	}

	req := c.newRequest(op, input, &ListVolumeRecoveryPointsOutput{})
	return ListVolumeRecoveryPointsRequest{Request: req, Input: input, Copy: c.ListVolumeRecoveryPointsRequest}
}

// ListVolumeRecoveryPointsRequest is the request type for the
// ListVolumeRecoveryPoints API operation.
type ListVolumeRecoveryPointsRequest struct {
	*aws.Request
	Input *ListVolumeRecoveryPointsInput
	Copy  func(*ListVolumeRecoveryPointsInput) ListVolumeRecoveryPointsRequest
}

// Send marshals and sends the ListVolumeRecoveryPoints API request.
func (r ListVolumeRecoveryPointsRequest) Send(ctx context.Context) (*ListVolumeRecoveryPointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVolumeRecoveryPointsResponse{
		ListVolumeRecoveryPointsOutput: r.Request.Data.(*ListVolumeRecoveryPointsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListVolumeRecoveryPointsResponse is the response type for the
// ListVolumeRecoveryPoints API operation.
type ListVolumeRecoveryPointsResponse struct {
	*ListVolumeRecoveryPointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVolumeRecoveryPoints request.
func (r *ListVolumeRecoveryPointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
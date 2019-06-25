// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/StopRelationalDatabaseRequest
type StopRelationalDatabaseInput struct {
	_ struct{} `type:"structure"`

	// The name of your database to stop.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`

	// The name of your new database snapshot to be created before stopping your
	// database.
	RelationalDatabaseSnapshotName *string `locationName:"relationalDatabaseSnapshotName" type:"string"`
}

// String returns the string representation
func (s StopRelationalDatabaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopRelationalDatabaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopRelationalDatabaseInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/StopRelationalDatabaseResult
type StopRelationalDatabaseOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your stop relational database request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s StopRelationalDatabaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopRelationalDatabase = "StopRelationalDatabase"

// StopRelationalDatabaseRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Stops a specific database that is currently running in Amazon Lightsail.
//
// The stop relational database operation supports tag-based access control
// via resource tags applied to the resource identified by relationalDatabaseName.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using StopRelationalDatabaseRequest.
//    req := client.StopRelationalDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/StopRelationalDatabase
func (c *Client) StopRelationalDatabaseRequest(input *StopRelationalDatabaseInput) StopRelationalDatabaseRequest {
	op := &aws.Operation{
		Name:       opStopRelationalDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopRelationalDatabaseInput{}
	}

	req := c.newRequest(op, input, &StopRelationalDatabaseOutput{})
	return StopRelationalDatabaseRequest{Request: req, Input: input, Copy: c.StopRelationalDatabaseRequest}
}

// StopRelationalDatabaseRequest is the request type for the
// StopRelationalDatabase API operation.
type StopRelationalDatabaseRequest struct {
	*aws.Request
	Input *StopRelationalDatabaseInput
	Copy  func(*StopRelationalDatabaseInput) StopRelationalDatabaseRequest
}

// Send marshals and sends the StopRelationalDatabase API request.
func (r StopRelationalDatabaseRequest) Send(ctx context.Context) (*StopRelationalDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopRelationalDatabaseResponse{
		StopRelationalDatabaseOutput: r.Request.Data.(*StopRelationalDatabaseOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopRelationalDatabaseResponse is the response type for the
// StopRelationalDatabase API operation.
type StopRelationalDatabaseResponse struct {
	*StopRelationalDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopRelationalDatabase request.
func (r *StopRelationalDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
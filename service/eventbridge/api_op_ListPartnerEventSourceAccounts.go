// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// An SaaS partner can use this operation to display the Amazon Web Services
// account ID that a particular partner event source name is associated with. This
// operation is not used by Amazon Web Services customers.
func (c *Client) ListPartnerEventSourceAccounts(ctx context.Context, params *ListPartnerEventSourceAccountsInput, optFns ...func(*Options)) (*ListPartnerEventSourceAccountsOutput, error) {
	if params == nil {
		params = &ListPartnerEventSourceAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPartnerEventSourceAccounts", params, optFns, c.addOperationListPartnerEventSourceAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPartnerEventSourceAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPartnerEventSourceAccountsInput struct {

	// The name of the partner event source to display account information about.
	//
	// This member is required.
	EventSourceName *string

	// Specifying this limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int32

	// The token returned by a previous call to this operation. Specifying this
	// retrieves the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPartnerEventSourceAccountsOutput struct {

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// The list of partner event sources returned by the operation.
	PartnerEventSourceAccounts []types.PartnerEventSourceAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPartnerEventSourceAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPartnerEventSourceAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPartnerEventSourceAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListPartnerEventSourceAccountsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPartnerEventSourceAccounts(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListPartnerEventSourceAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListPartnerEventSourceAccounts",
	}
}

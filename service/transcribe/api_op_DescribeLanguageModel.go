// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a single custom language model. Use this information to
// see details about the language model in your Amazon Web Services account. You
// can also see whether the base language model used to create your custom language
// model has been updated. If Amazon Transcribe has updated the base model, you can
// create a new custom language model using the updated base model. If the language
// model wasn't created, you can use this operation to understand why Amazon
// Transcribe couldn't create it.
func (c *Client) DescribeLanguageModel(ctx context.Context, params *DescribeLanguageModelInput, optFns ...func(*Options)) (*DescribeLanguageModelOutput, error) {
	if params == nil {
		params = &DescribeLanguageModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLanguageModel", params, optFns, c.addOperationDescribeLanguageModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLanguageModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLanguageModelInput struct {

	// The name of the custom language model you submit to get more information.
	//
	// This member is required.
	ModelName *string

	noSmithyDocumentSerde
}

type DescribeLanguageModelOutput struct {

	// The name of the custom language model you requested more information about.
	LanguageModel *types.LanguageModel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLanguageModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLanguageModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLanguageModel{}, middleware.After)
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
	if err = addOpDescribeLanguageModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLanguageModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLanguageModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "DescribeLanguageModel",
	}
}

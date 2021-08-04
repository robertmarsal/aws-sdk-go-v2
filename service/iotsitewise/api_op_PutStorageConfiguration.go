// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures storage settings for IoT SiteWise.
func (c *Client) PutStorageConfiguration(ctx context.Context, params *PutStorageConfigurationInput, optFns ...func(*Options)) (*PutStorageConfigurationOutput, error) {
	if params == nil {
		params = &PutStorageConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutStorageConfiguration", params, optFns, c.addOperationPutStorageConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutStorageConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutStorageConfigurationInput struct {

	// The type of storage that you specified for your data. The storage type can be
	// one of the following values:
	//
	// * SITEWISE_DEFAULT_STORAGE – IoT SiteWise
	// replicates your data into a service managed database.
	//
	// * MULTI_LAYER_STORAGE –
	// IoT SiteWise replicates your data into a service managed database and saves a
	// copy of your raw data and metadata in an Amazon S3 object that you specified.
	//
	// This member is required.
	StorageType types.StorageType

	// Identifies a storage destination. If you specified MULTI_LAYER_STORAGE for the
	// storage type, you must specify a MultiLayerStorage object.
	MultiLayerStorage *types.MultiLayerStorage

	noSmithyDocumentSerde
}

type PutStorageConfigurationOutput struct {

	// Contains current status information for the configuration.
	//
	// This member is required.
	ConfigurationStatus *types.ConfigurationStatus

	// The type of storage that you specified for your data. The storage type can be
	// one of the following values:
	//
	// * SITEWISE_DEFAULT_STORAGE – IoT SiteWise
	// replicates your data into a service managed database.
	//
	// * MULTI_LAYER_STORAGE –
	// IoT SiteWise replicates your data into a service managed database and saves a
	// copy of your raw data and metadata in an Amazon S3 object that you specified.
	//
	// This member is required.
	StorageType types.StorageType

	// Contains information about the storage destination.
	MultiLayerStorage *types.MultiLayerStorage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutStorageConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutStorageConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutStorageConfiguration{}, middleware.After)
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
	if err = addEndpointPrefix_opPutStorageConfigurationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutStorageConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutStorageConfiguration(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opPutStorageConfigurationMiddleware struct {
}

func (*endpointPrefix_opPutStorageConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutStorageConfigurationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opPutStorageConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPutStorageConfigurationMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opPutStorageConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "PutStorageConfiguration",
	}
}

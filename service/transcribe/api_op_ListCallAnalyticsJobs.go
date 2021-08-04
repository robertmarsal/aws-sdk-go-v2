// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List call analytics jobs with a specified status or substring that matches their
// names.
func (c *Client) ListCallAnalyticsJobs(ctx context.Context, params *ListCallAnalyticsJobsInput, optFns ...func(*Options)) (*ListCallAnalyticsJobsOutput, error) {
	if params == nil {
		params = &ListCallAnalyticsJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCallAnalyticsJobs", params, optFns, c.addOperationListCallAnalyticsJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCallAnalyticsJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCallAnalyticsJobsInput struct {

	// When specified, the jobs returned in the list are limited to jobs whose name
	// contains the specified string.
	JobNameContains *string

	// The maximum number of call analytics jobs to return in the response. If there
	// are fewer results in the list, this response contains only the actual results.
	MaxResults *int32

	// If you receive a truncated result in the previous request of , include NextToken
	// to fetch the next set of jobs.
	NextToken *string

	// When specified, returns only call analytics jobs with the specified status. Jobs
	// are ordered by creation date, with the most recent jobs returned first. If you
	// don't specify a status, Amazon Transcribe returns all analytics jobs ordered by
	// creation date.
	Status types.CallAnalyticsJobStatus

	noSmithyDocumentSerde
}

type ListCallAnalyticsJobsOutput struct {

	// A list of objects containing summary information for a transcription job.
	CallAnalyticsJobSummaries []types.CallAnalyticsJobSummary

	// The operation returns a page of jobs at a time. The maximum size of the page is
	// set by the MaxResults parameter. If there are more jobs in the list than the
	// page size, Amazon Transcribe returns the NextPage token. Include the token in
	// your next request to the operation to return next page of jobs.
	NextToken *string

	// When specified, returns only call analytics jobs with that status. Jobs are
	// ordered by creation date, with the most recent jobs returned first. If you don't
	// specify a status, Amazon Transcribe returns all transcription jobs ordered by
	// creation date.
	Status types.CallAnalyticsJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCallAnalyticsJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCallAnalyticsJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCallAnalyticsJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCallAnalyticsJobs(options.Region), middleware.Before); err != nil {
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

// ListCallAnalyticsJobsAPIClient is a client that implements the
// ListCallAnalyticsJobs operation.
type ListCallAnalyticsJobsAPIClient interface {
	ListCallAnalyticsJobs(context.Context, *ListCallAnalyticsJobsInput, ...func(*Options)) (*ListCallAnalyticsJobsOutput, error)
}

var _ ListCallAnalyticsJobsAPIClient = (*Client)(nil)

// ListCallAnalyticsJobsPaginatorOptions is the paginator options for
// ListCallAnalyticsJobs
type ListCallAnalyticsJobsPaginatorOptions struct {
	// The maximum number of call analytics jobs to return in the response. If there
	// are fewer results in the list, this response contains only the actual results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCallAnalyticsJobsPaginator is a paginator for ListCallAnalyticsJobs
type ListCallAnalyticsJobsPaginator struct {
	options   ListCallAnalyticsJobsPaginatorOptions
	client    ListCallAnalyticsJobsAPIClient
	params    *ListCallAnalyticsJobsInput
	nextToken *string
	firstPage bool
}

// NewListCallAnalyticsJobsPaginator returns a new ListCallAnalyticsJobsPaginator
func NewListCallAnalyticsJobsPaginator(client ListCallAnalyticsJobsAPIClient, params *ListCallAnalyticsJobsInput, optFns ...func(*ListCallAnalyticsJobsPaginatorOptions)) *ListCallAnalyticsJobsPaginator {
	if params == nil {
		params = &ListCallAnalyticsJobsInput{}
	}

	options := ListCallAnalyticsJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCallAnalyticsJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCallAnalyticsJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCallAnalyticsJobs page.
func (p *ListCallAnalyticsJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCallAnalyticsJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListCallAnalyticsJobs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListCallAnalyticsJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListCallAnalyticsJobs",
	}
}

/*
Studio - AI Empower Labs

# Studio API Documentation  ## Introduction Welcome to Studio by AI Empower Labs API documentation! We are thrilled to offer developers around the world access to our cutting-edge artificial intelligence technology and semantic search. Our API is designed to empower your applications with state-of-the-art AI capabilities, including but not limited to natural language processing, audio transcription, embedding, and predictive analytics.  Our mission is to make AI technology accessible and easy to integrate, enabling you to enhance your applications, improve user experiences, and innovate in your field. Whether you're building complex systems, developing mobile apps, or creating web services, our API provides you with the tools you need to incorporate AI functionalities seamlessly.  Support and Feedback We are committed to providing exceptional support to our developer community. If you have any questions, encounter any issues, or have feedback on how we can improve our API, please don't hesitate to contact our support team @ support@AIEmpowerLabs.com.  Terms of Use Please review our terms of use and privacy policy before integrating our API into your application. By using our API, you agree to comply with these terms.

API version: v1
Contact: support@aiempowerlabs.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// MiscellaneousAPIService MiscellaneousAPI service
type MiscellaneousAPIService service

type ApiGetConfigurationRequest struct {
	ctx context.Context
	ApiService *MiscellaneousAPIService
}

func (r ApiGetConfigurationRequest) Execute() ([]map[string][]string, *http.Response, error) {
	return r.ApiService.GetConfigurationExecute(r)
}

/*
GetConfiguration Method for GetConfiguration

Get configuration information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConfigurationRequest
*/
func (a *MiscellaneousAPIService) GetConfiguration(ctx context.Context) ApiGetConfigurationRequest {
	return ApiGetConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string][]string
func (a *MiscellaneousAPIService) GetConfigurationExecute(r ApiGetConfigurationRequest) ([]map[string][]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string][]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MiscellaneousAPIService.GetConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
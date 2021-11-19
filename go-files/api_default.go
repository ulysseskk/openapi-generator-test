/*
Swagger Petstore

:dog: :cat: :rabbit: This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.

API version: 1.0.0
Contact: apiteam@swagger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package test_client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiGenerateClientCodeRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	generatorName *string
	packageName *string
}

// 生成器名称，使用GET /code_genenerators获取枚举
func (r ApiGenerateClientCodeRequest) GeneratorName(generatorName string) ApiGenerateClientCodeRequest {
	r.generatorName = &generatorName
	return r
}
// 代码包名称
func (r ApiGenerateClientCodeRequest) PackageName(packageName string) ApiGenerateClientCodeRequest {
	r.packageName = &packageName
	return r
}

func (r ApiGenerateClientCodeRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GenerateClientCodeExecute(r)
}

/*
GenerateClientCode client代码生成

根据openapi yaml文件生成client代码

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenerateClientCodeRequest
*/
func (a *DefaultApiService) GenerateClientCode(ctx _context.Context) ApiGenerateClientCodeRequest {
	return ApiGenerateClientCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) GenerateClientCodeExecute(r ApiGenerateClientCodeRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GenerateClientCode")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/client_code_gen"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.generatorName == nil {
		return nil, reportError("generatorName is required and must be specified")
	}

	localVarQueryParams.Add("generator_name", parameterToString(*r.generatorName, ""))
	if r.packageName != nil {
		localVarQueryParams.Add("package_name", parameterToString(*r.packageName, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGenerateDocRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	generatorName *string
	body *interface{}
}

// 生成器名称，使用GET /doc_generators获取枚举
func (r ApiGenerateDocRequest) GeneratorName(generatorName string) ApiGenerateDocRequest {
	r.generatorName = &generatorName
	return r
}
// 上传的OPEN-API yaml配置文件
func (r ApiGenerateDocRequest) Body(body interface{}) ApiGenerateDocRequest {
	r.body = &body
	return r
}

func (r ApiGenerateDocRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GenerateDocExecute(r)
}

/*
GenerateDoc 生成文档

根据openapi yaml文件生成接口文档

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenerateDocRequest
*/
func (a *DefaultApiService) GenerateDoc(ctx _context.Context) ApiGenerateDocRequest {
	return ApiGenerateDocRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) GenerateDocExecute(r ApiGenerateDocRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GenerateDoc")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/doc_gen"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.generatorName == nil {
		return nil, reportError("generatorName is required and must be specified")
	}

	localVarQueryParams.Add("generator_name", parameterToString(*r.generatorName, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/yaml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetClientCodeGeneratorsRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
}


func (r ApiGetClientCodeGeneratorsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetClientCodeGeneratorsExecute(r)
}

/*
GetClientCodeGenerators 获取代码生成器名称列表

获取代码生成器名称列表

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetClientCodeGeneratorsRequest
*/
func (a *DefaultApiService) GetClientCodeGenerators(ctx _context.Context) ApiGetClientCodeGeneratorsRequest {
	return ApiGetClientCodeGeneratorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) GetClientCodeGeneratorsExecute(r ApiGetClientCodeGeneratorsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetClientCodeGenerators")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/code_genenerators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetDocGeneratorsRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
}


func (r ApiGetDocGeneratorsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetDocGeneratorsExecute(r)
}

/*
GetDocGenerators 获取文档生成器名称列表

获取文档生成器名称列表

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDocGeneratorsRequest
*/
func (a *DefaultApiService) GetDocGenerators(ctx _context.Context) ApiGetDocGeneratorsRequest {
	return ApiGetDocGeneratorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) GetDocGeneratorsExecute(r ApiGetDocGeneratorsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetDocGenerators")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/doc_generators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
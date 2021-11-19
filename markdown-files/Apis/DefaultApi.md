# DefaultApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**generate client code**](DefaultApi.md#generate client code) | **POST** /client_code_gen | client代码生成
[**generate doc**](DefaultApi.md#generate doc) | **POST** /doc_gen | 生成文档
[**get client code generators**](DefaultApi.md#get client code generators) | **GET** /code_genenerators | 获取代码生成器名称列表
[**get doc generators**](DefaultApi.md#get doc generators) | **GET** /doc_generators | 获取文档生成器名称列表


<a name="generate client code"></a>
# **generate client code**
> generate client code(generator\_name, package\_name)

client代码生成

    根据openapi yaml文件生成client代码

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generator\_name** | **String**| 生成器名称，使用GET /code_genenerators获取枚举 | [default to null]
 **package\_name** | **String**| 代码包名称 | [optional] [default to client]

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="generate doc"></a>
# **generate doc**
> generate doc(generator\_name, body)

生成文档

    根据openapi yaml文件生成接口文档

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generator\_name** | **String**| 生成器名称，使用GET /doc_generators获取枚举 | [default to null]
 **body** | **oas_any_type_not_mapped**| 上传的OPEN-API yaml配置文件 | [optional]

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/yaml
- **Accept**: Not defined

<a name="get client code generators"></a>
# **get client code generators**
> get client code generators()

获取代码生成器名称列表

    获取代码生成器名称列表

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="get doc generators"></a>
# **get doc generators**
> get doc generators()

获取文档生成器名称列表

    获取文档生成器名称列表

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


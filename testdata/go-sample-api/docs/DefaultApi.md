# \DefaultApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateState**](DefaultApi.md#CreateState) | **Post** /state | 
[**GetState**](DefaultApi.md#GetState) | **Get** /state | 


# **CreateState**
> CreateState(ctx, uNKNOWNBASETYPE)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetState**
> OneOfObjAObjB GetState(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OneOfObjAObjB**](oneOf&lt;ObjA,ObjB&gt;.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


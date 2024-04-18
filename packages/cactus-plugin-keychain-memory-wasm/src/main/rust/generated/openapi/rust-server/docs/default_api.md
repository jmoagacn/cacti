# default_api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
**deleteKeychainEntryV1**](default_api.md#deleteKeychainEntryV1) | **POST** /api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory-wasm/delete-keychain-entry | Deletes an entry under a key on the keychain backend.
**getKeychainEntryV1**](default_api.md#getKeychainEntryV1) | **POST** /api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory-wasm/get-keychain-entry | Retrieves the contents of a keychain entry from the backend.
**hasKeychainEntryV1**](default_api.md#hasKeychainEntryV1) | **POST** /api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory-wasm/has-keychain-entry | Checks that an entry exists under a key on the keychain backend
**setKeychainEntryV1**](default_api.md#setKeychainEntryV1) | **POST** /api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory-wasm/set-keychain-entry | Sets a value under a key on the keychain backend.


# **deleteKeychainEntryV1**
> models::DeleteKeychainEntryResponseV1 deleteKeychainEntryV1(delete_keychain_entry_request_v1)
Deletes an entry under a key on the keychain backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **delete_keychain_entry_request_v1** | [**DeleteKeychainEntryRequestV1**](DeleteKeychainEntryRequestV1.md)| Request body to delete a keychain entry via its key | 

### Return type

[**models::DeleteKeychainEntryResponseV1**](DeleteKeychainEntryResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getKeychainEntryV1**
> models::GetKeychainEntryResponseV1 getKeychainEntryV1(get_keychain_entry_request_v1)
Retrieves the contents of a keychain entry from the backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **get_keychain_entry_request_v1** | [**GetKeychainEntryRequestV1**](GetKeychainEntryRequestV1.md)| Request body to obtain a keychain entry via its key | 

### Return type

[**models::GetKeychainEntryResponseV1**](GetKeychainEntryResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **hasKeychainEntryV1**
> models::HasKeychainEntryResponseV1 hasKeychainEntryV1(has_keychain_entry_request_v1)
Checks that an entry exists under a key on the keychain backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **has_keychain_entry_request_v1** | [**HasKeychainEntryRequestV1**](HasKeychainEntryRequestV1.md)| Request body for checking a keychain entry via its key | 

### Return type

[**models::HasKeychainEntryResponseV1**](HasKeychainEntryResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **setKeychainEntryV1**
> models::SetKeychainEntryResponseV1 setKeychainEntryV1(set_keychain_entry_request_v1)
Sets a value under a key on the keychain backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **set_keychain_entry_request_v1** | [**SetKeychainEntryRequestV1**](SetKeychainEntryRequestV1.md)| Request body to write/update a keychain entry via its key | 

### Return type

[**models::SetKeychainEntryResponseV1**](SetKeychainEntryResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# PetAPI

> 

All URIs are relative to *http://localhost*

Operation | Method | Path | Description
------------- | ------------- |------------- | -------------
[**addPet**](PetAPI.md#addPet) | **Post** | /pet/{status} | Add a new pet to the store
[**deletePet**](PetAPI.md#deletePet) | **Delete** | /pet/{petId} | Deletes a pet
[**findPetsByStatus**](PetAPI.md#findPetsByStatus) | **Get** | /pet/findByStatus | Finds Pets by status
[**findPetsByTags**](PetAPI.md#findPetsByTags) | **Get** | /pet/findByTags | Finds Pets by tags
[**getPetByID**](PetAPI.md#getPetByID) | **Get** | /pet/{petId} | Find pet by ID
[**updatePet**](PetAPI.md#updatePet) | **Put** | /pet/{status} | Update an existing pet
[**updatePetWithForm**](PetAPI.md#updatePetWithForm) | **Post** | /pet/{petId} | Updates a pet in the store with form data
[**uploadFile**](PetAPI.md#uploadFile) | **Post** | /pet/{petId}/uploadImage | uploads an image

## addPet

> Add a new pet to the store


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineObject** | [**InlineObject**](.md) |  | [optional] 

### Return type

 (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined


## deletePet

> Deletes a pet


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PetID** | **int64** | Pet id to delete | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APIKey** | **string** |  | [optional] [default to null]

### Return type

 (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## findPetsByStatus

> Finds Pets by statusMultiple status values can be provided with comma separated strings


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**array**](.md) | Status values that need to be considered for filter | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** |  | [optional] [default to null]

### Return type

[**array**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `application/xml`


## findPetsByTags

> Finds Pets by tagsMultiple tags can be provided with comma separated strings. Use\\ \\ tag1, tag2, tag3 for testing.


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**array**](.md) | Tags to filter by | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

[**array**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `application/xml`


## getPetByID

> Find pet by IDReturns a single pet


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PetID** | **int64** | ID of pet to return | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

[**Pet**](Pet.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `application/xml`


## updatePet

> Update an existing pet


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pet** | [**Pet**](.md) | Pet object that needs to be added to the store | 

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: `application/json`, `application/xml`
- **Accept**: Not defined


## updatePetWithForm

> Updates a pet in the store with form data


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PetID** | Pointer to **int64** | ID of pet that needs to be updated | [default to 10]
**PetID2** | Pointer to **int64** | ID of pet that needs to be updated | [default to 10]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Updated name of the pet | [optional] [default to null]
**Status** | **string** | Updated status of the pet | [optional] [default to null]

### Return type

 (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: `application/x-www-form-urlencoded`
- **Accept**: Not defined


## uploadFile

> uploads an image


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PetID** | **int64** | ID of pet to update | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **io.ReadCloser** |  | [optional] 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: `application/octet-stream`
- **Accept**: `application/json`


[[Back to top]](#) [[Back to API list]](./README.md#endpoints) [[Back to Model list]](./README.md#models) [[Back to README]](./README.md)

# UserAPI

> 

All URIs are relative to *http://localhost*

Operation | Method | Path | Description
------------- | ------------- |------------- | -------------
[**createUser**](UserAPI.md#createUser) | **Post** | /user | Create user
[**createUsersWithArrayInput**](UserAPI.md#createUsersWithArrayInput) | **Post** | /user/createWithArray | Creates list of users with given input array
[**createUsersWithListInput**](UserAPI.md#createUsersWithListInput) | **Post** | /user/createWithList | Creates list of users with given input array
[**deleteUser**](UserAPI.md#deleteUser) | **Delete** | /user/{username} | Delete user
[**getUserByName**](UserAPI.md#getUserByName) | **Get** | /user/{username} | Get user by user name
[**loginUser**](UserAPI.md#loginUser) | **Get** | /user/login | Logs user into the system
[**logoutUser**](UserAPI.md#logoutUser) | **Get** | /user/logout | Logs out current logged in user session
[**userUsernamePut**](UserAPI.md#userUsernamePut) | **Put** | /user/{username} | Updated user

## createUser

> Create userThis can only be done by the logged in user.


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**User**](.md) | Created user object | 

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined


## createUsersWithArrayInput

> Creates list of users with given input array


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**array**](.md) | List of user object | 

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined


## createUsersWithListInput

> Creates list of users with given input array


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**array**](.md) | List of user object | 

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined


## deleteUser

> Delete userThis can only be done by the logged in user.


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The name that needs to be deleted | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## getUserByName

> Get user by user name


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The name that needs to be fetched. Use user1 for testing. | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `application/xml`


## loginUser

> Logs user into the system


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The user name for login | [default to null]
**Password** | **string** | The password for login in clear text | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `application/xml`


## logoutUser

> Logs out current logged in user session

### Parameters

This endpoint does not need any parameter.

### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## userUsernamePut

> Updated userThis can only be done by the logged in user.


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | name that need to be updated | [default to null]
**User** | [**User**](.md) | Updated user object | 

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: Not defined


[[Back to top]](#) [[Back to API list]](./README.md#endpoints) [[Back to Model list]](./README.md#models) [[Back to README]](./README.md)

# StoreAPI

> 

All URIs are relative to *http://localhost*

Operation | Method | Path | Description
------------- | ------------- |------------- | -------------
[**deleteOrder**](StoreAPI.md#deleteOrder) | **Delete** | /store/order/{orderId} | Delete purchase order by ID
[**getInventory**](StoreAPI.md#getInventory) | **Get** | /store/inventory | Returns pet inventories by status
[**getOrderByID**](StoreAPI.md#getOrderByID) | **Get** | /store/order/{orderId} | Find purchase order by ID
[**placeOrder**](StoreAPI.md#placeOrder) | **Post** | /store/order | Place an order for a pet

## deleteOrder

> Delete purchase order by IDFor valid response try integer IDs with positive integer value.\\ \\ Negative or non-integer values will generate API errors


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **int64** | ID of the order that needs to be deleted | [default to null]

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


## getInventory

> Returns pet inventories by statusReturns a map of status codes to quantities

### Parameters

This endpoint does not need any parameter.

### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

[**map**](integer.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


## getOrderByID

> Find purchase order by IDFor valid response try integer IDs with value >= 1 and <= 10.\\ \\ Other values will generated exceptions


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **int64** | ID of pet that needs to be fetched | [default to null]

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`, `application/xml`


## placeOrder

> Place an order for a pet


### Required Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**Order**](.md) | order placed for purchasing the pet | 

### Optional Parameters

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`, `application/xml`


[[Back to top]](#) [[Back to API list]](./README.md#endpoints) [[Back to Model list]](./README.md#models) [[Back to README]](./README.md)

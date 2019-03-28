# Documentation

This is a sample Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). 

## Overview

API version: 1.0.0

## Endpoints

All URIs are relative to *http://localhost*

Class | Operation | Method | Path | Description
------------ | ------------- |------------- | ------------- | -------------
*PetAPI* | [**addPet**](PetAPI.md#addpet) | **Post** | /pet/{status} | Add a new pet to the store
*PetAPI* | [**deletePet**](PetAPI.md#deletepet) | **Delete** | /pet/{petId} | Deletes a pet
*PetAPI* | [**findPetsByStatus**](PetAPI.md#findpetsbystatus) | **Get** | /pet/findByStatus | Finds Pets by status
*PetAPI* | [**findPetsByTags**](PetAPI.md#findpetsbytags) | **Get** | /pet/findByTags | Finds Pets by tags
*PetAPI* | [**getPetByID**](PetAPI.md#getpetbyid) | **Get** | /pet/{petId} | Find pet by ID
*PetAPI* | [**updatePet**](PetAPI.md#updatepet) | **Put** | /pet/{status} | Update an existing pet
*PetAPI* | [**updatePetWithForm**](PetAPI.md#updatepetwithform) | **Post** | /pet/{petId} | Updates a pet in the store with form data
*PetAPI* | [**uploadFile**](PetAPI.md#uploadfile) | **Post** | /pet/{petId}/uploadImage | uploads an image
*StoreAPI* | [**deleteOrder**](StoreAPI.md#deleteorder) | **Delete** | /store/order/{orderId} | Delete purchase order by ID
*StoreAPI* | [**getInventory**](StoreAPI.md#getinventory) | **Get** | /store/inventory | Returns pet inventories by status
*StoreAPI* | [**getOrderByID**](StoreAPI.md#getorderbyid) | **Get** | /store/order/{orderId} | Find purchase order by ID
*StoreAPI* | [**placeOrder**](StoreAPI.md#placeorder) | **Post** | /store/order | Place an order for a pet
*TestAPI* | [**testMethod**](TestAPI.md#testmethod) | **Get** | /test | 
*UserAPI* | [**createUser**](UserAPI.md#createuser) | **Post** | /user | Create user
*UserAPI* | [**createUsersWithArrayInput**](UserAPI.md#createuserswitharrayinput) | **Post** | /user/createWithArray | Creates list of users with given input array
*UserAPI* | [**createUsersWithListInput**](UserAPI.md#createuserswithlistinput) | **Post** | /user/createWithList | Creates list of users with given input array
*UserAPI* | [**deleteUser**](UserAPI.md#deleteuser) | **Delete** | /user/{username} | Delete user
*UserAPI* | [**getUserByName**](UserAPI.md#getuserbyname) | **Get** | /user/{username} | Get user by user name
*UserAPI* | [**loginUser**](UserAPI.md#loginuser) | **Get** | /user/login | Logs user into the system
*UserAPI* | [**logoutUser**](UserAPI.md#logoutuser) | **Get** | /user/logout | Logs out current logged in user session
*UserAPI* | [**userUsernamePut**](UserAPI.md#userusernameput) | **Put** | /user/{username} | Updated user


## Models

 - [APIResponse](APIResponse.md)
 - [Category](Category.md)
 - [InlineObject](InlineObject.md)
 - [InlineObject1](InlineObject1.md)
 - [Order](Order.md)
 - [Pet](Pet.md)
 - [Tag](Tag.md)
 - [User](User.md)


## Author

apiteam@swagger.io

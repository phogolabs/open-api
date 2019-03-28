//
// Swagger Petstore
// 
// This is a sample Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). 
// 
// API version: 1.0.0
// Contact: apiteam@swagger.io
//
// Code generated by OpenAPI Generator (https://http://github.com/openAPITools/openapi-generator)

// Package test contains all API models and handlers
package test

import (
  "net/http"
  "encoding/json"
  "github.com/phogolabs/rest"
  "io"
  "mime/multipart"
  "time"
)

// AddPetInput is the input for the AddPet operation
type AddPetInput struct {
  // Body payload
  Body InlineObject `json:"-" xml:"-" validate:"-"`
}

// Bind manages request input
func (p *AddPetInput) Bind(r *http.Request) (err error) {
  return
}

// UnmarshalJSON unmarshals the body from JSON
func (p *AddPetInput) UnmarshalJSON(data []byte) error {
  return json.Unmarshal(data, &p.Body);
}

// DeletePetInputParam contains the parameters for the DeletePet operation
type DeletePetInputParam struct {
  // PetID param: Pet id to delete
  PetID int64 `json:"-" xml:"-" path:"petId" validate:"required"`
  // APIKey param
  APIKey string `json:"-" xml:"-" header:"api_key" validate:"-"`
}

// Bind manages request partameters
func (p *DeletePetInputParam) Bind(r *http.Request) (err error) {
  if err = rest.DecodePath(r, p); err != nil {
    return err
  }

  return nil
}

// DeletePetInput is the input for the DeletePet operation
type DeletePetInput struct {
  // Param contains the parameters for the DeletePet operation
  Param DeletePetInputParam `json:"-"`
}

// Bind manages request input
func (p *DeletePetInput) Bind(r *http.Request) (err error) {
  err = p.Param.Bind(r)
  return
}

// FindPetsByStatusInputParam contains the parameters for the FindPetsByStatus operation
type FindPetsByStatusInputParam struct {
  // Status param: Status values that need to be considered for filter
  Status array `json:"-" xml:"-" query:"status" validate:"required"`
  // Time param
  Time time.Time `json:"-" xml:"-" query:"time" validate:"-"`
}

// Bind manages request partameters
func (p *FindPetsByStatusInputParam) Bind(r *http.Request) (err error) {
  if err = rest.DecodeQuery(r, p); err != nil {
    return err
  }

  return nil
}

// FindPetsByStatusInput is the input for the FindPetsByStatus operation
type FindPetsByStatusInput struct {
  // Param contains the parameters for the FindPetsByStatus operation
  Param FindPetsByStatusInputParam `json:"-"`
}

// Bind manages request input
func (p *FindPetsByStatusInput) Bind(r *http.Request) (err error) {
  err = p.Param.Bind(r)
  return
}

// FindPetsByStatusOutput is the output for the FindPetsByStatus operation
type FindPetsByStatusOutput struct {
  // Body payload
  Body array `json:"-"`
}

// Render manages response payloads.
func (p *FindPetsByStatusOutput) Render(w http.ResponseWriter, r *http.Request) (err error) {
  w.WriteHeader(200)
  return
}

// MarshalJSON marshals the body as JSON
func (p *FindPetsByStatusOutput) MarshalJSON() ([]byte, error) {
  return json.Marshal(&p.Body);
}

// FindPetsByTagsInputParam contains the parameters for the FindPetsByTags operation
type FindPetsByTagsInputParam struct {
  // Tags param: Tags to filter by
  Tags array `json:"-" xml:"-" query:"tags" validate:"required"`
}

// Bind manages request partameters
func (p *FindPetsByTagsInputParam) Bind(r *http.Request) (err error) {
  if err = rest.DecodeQuery(r, p); err != nil {
    return err
  }

  return nil
}

// FindPetsByTagsInput is the input for the FindPetsByTags operation
type FindPetsByTagsInput struct {
  // Param contains the parameters for the FindPetsByTags operation
  Param FindPetsByTagsInputParam `json:"-"`
}

// Bind manages request input
func (p *FindPetsByTagsInput) Bind(r *http.Request) (err error) {
  err = p.Param.Bind(r)
  return
}

// FindPetsByTagsOutput is the output for the FindPetsByTags operation
type FindPetsByTagsOutput struct {
  // Body payload
  Body array `json:"-"`
}

// Render manages response payloads.
func (p *FindPetsByTagsOutput) Render(w http.ResponseWriter, r *http.Request) (err error) {
  w.WriteHeader(200)
  return
}

// MarshalJSON marshals the body as JSON
func (p *FindPetsByTagsOutput) MarshalJSON() ([]byte, error) {
  return json.Marshal(&p.Body);
}

// GetPetByIDInputParam contains the parameters for the GetPetByID operation
type GetPetByIDInputParam struct {
  // PetID param: ID of pet to return
  PetID int64 `json:"-" xml:"-" path:"petId" validate:"required"`
}

// Bind manages request partameters
func (p *GetPetByIDInputParam) Bind(r *http.Request) (err error) {
  if err = rest.DecodePath(r, p); err != nil {
    return err
  }

  return nil
}

// GetPetByIDInput is the input for the GetPetByID operation
type GetPetByIDInput struct {
  // Param contains the parameters for the GetPetByID operation
  Param GetPetByIDInputParam `json:"-"`
}

// Bind manages request input
func (p *GetPetByIDInput) Bind(r *http.Request) (err error) {
  err = p.Param.Bind(r)
  return
}

// GetPetByIDOutput is the output for the GetPetByID operation
type GetPetByIDOutput struct {
  // Body payload
  Body Pet `json:"-"`
}

// Render manages response payloads.
func (p *GetPetByIDOutput) Render(w http.ResponseWriter, r *http.Request) (err error) {
  w.WriteHeader(200)
  return
}

// MarshalJSON marshals the body as JSON
func (p *GetPetByIDOutput) MarshalJSON() ([]byte, error) {
  return json.Marshal(&p.Body);
}

// UpdatePetInput is the input for the UpdatePet operation
type UpdatePetInput struct {
  // Body payload: Pet object that needs to be added to the store
  Body Pet `json:"-" xml:"-" validate:"required"`
}

// Bind manages request input
func (p *UpdatePetInput) Bind(r *http.Request) (err error) {
  return
}

// UnmarshalJSON unmarshals the body from JSON
func (p *UpdatePetInput) UnmarshalJSON(data []byte) error {
  return json.Unmarshal(data, &p.Body);
}

// UpdatePetWithFormInputParam contains the parameters for the UpdatePetWithForm operation
type UpdatePetWithFormInputParam struct {
  // PetID param: ID of pet that needs to be updated
  PetID *int64 `json:"-" xml:"-" path:"petId" default:"10" validate:"required"`
  // PetID2 param: ID of pet that needs to be updated
  PetID2 *int64 `json:"-" xml:"-" query:"petId2" default:"10" validate:"required"`
  // Name param: Updated name of the pet
  Name string `json:"-" xml:"-" form:"name" validate:"-"`
  // Status param: Updated status of the pet
  Status string `json:"-" xml:"-" form:"status" validate:"-"`
}

// Bind manages request partameters
func (p *UpdatePetWithFormInputParam) Bind(r *http.Request) (err error) {
  if err = rest.DecodePath(r, p); err != nil {
    return err
  }

  if err = rest.DecodeQuery(r, p); err != nil {
    return err
  }

  return nil
}

// UpdatePetWithFormInput is the input for the UpdatePetWithForm operation
type UpdatePetWithFormInput struct {
  // Param contains the parameters for the UpdatePetWithForm operation
  Param UpdatePetWithFormInputParam `json:"-"`
}

// Bind manages request input
func (p *UpdatePetWithFormInput) Bind(r *http.Request) (err error) {
  err = p.Param.Bind(r)
  return
}

// UploadFileInputParam contains the parameters for the UploadFile operation
type UploadFileInputParam struct {
  // PetID param: ID of pet to update
  PetID int64 `json:"-" xml:"-" path:"petId" validate:"required"`
}

// Bind manages request partameters
func (p *UploadFileInputParam) Bind(r *http.Request) (err error) {
  if err = rest.DecodePath(r, p); err != nil {
    return err
  }

  return nil
}

// UploadFileInput is the input for the UploadFile operation
type UploadFileInput struct {
  // Param contains the parameters for the UploadFile operation
  Param UploadFileInputParam `json:"-"`
  // Filename is the name of the file
  Filename string `json:"-" xml:"-" validate:"-"`
  // Size is the size of the file
  Size int64 `json:"-" xml:"-" validate:"-"`
  // Body payload
  Body io.ReadCloser `json:"-" xml:"-" validate:"-"`
}

// Bind manages request input
func (p *UploadFileInput) Bind(r *http.Request) (err error) {
  var header *multipart.FileHeader
  if p.Body, header, err = r.FormFile("body"); err != nil {
    return
  }

  p.Filename = header.Filename
  p.Size = header.Size

  err = p.Param.Bind(r)
  return
}

// UploadFileOutput is the output for the UploadFile operation
type UploadFileOutput struct {
  // Body payload
  Body ApiResponse `json:"-"`
}

// Render manages response payloads.
func (p *UploadFileOutput) Render(w http.ResponseWriter, r *http.Request) (err error) {
  w.WriteHeader(200)
  return
}

// MarshalJSON marshals the body as JSON
func (p *UploadFileOutput) MarshalJSON() ([]byte, error) {
  return json.Marshal(&p.Body);
}

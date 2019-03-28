package test_test

import (
  "net/http"
  "net/http/httptest"
  
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/go-chi/chi"
)

var _ = Describe("PetAPI", func() {
  var (
    router   *chi.Mux
    request  *http.Request
    recorder *httptest.ResponseRecorder
  )

  BeforeEach(func() {
    router = chi.NewRouter()
    recorder = httptest.NewRecorder()

    service := &handler.PetAPI{}
    service.Mount(router)
  })

  Describe("Post /pet/{status}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Delete /pet/{petId}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Get /pet/findByStatus", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Get /pet/findByTags", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Get /pet/{petId}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Put /pet/{status}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Post /pet/{petId}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Post /pet/{petId}/uploadImage", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })
})


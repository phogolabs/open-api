package test_test

import (
  "net/http"
  "net/http/httptest"
  
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/go-chi/chi"
)

var _ = Describe("StoreAPI", func() {
  var (
    router   *chi.Mux
    request  *http.Request
    recorder *httptest.ResponseRecorder
  )

  BeforeEach(func() {
    router = chi.NewRouter()
    recorder = httptest.NewRecorder()

    service := &handler.StoreAPI{}
    service.Mount(router)
  })

  Describe("Delete /store/order/{orderId}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Get /store/inventory", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Get /store/order/{orderId}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Post /store/order", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })
})


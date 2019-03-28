package test_test

import (
  "net/http"
  "net/http/httptest"
  
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/go-chi/chi"
)

var _ = Describe("UserAPI", func() {
  var (
    router   *chi.Mux
    request  *http.Request
    recorder *httptest.ResponseRecorder
  )

  BeforeEach(func() {
    router = chi.NewRouter()
    recorder = httptest.NewRecorder()

    service := &handler.UserAPI{}
    service.Mount(router)
  })

  Describe("Post /user", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Post /user/createWithArray", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Post /user/createWithList", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Delete /user/{username}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Get /user/{username}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Get /user/login", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Get /user/logout", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })

  Describe("Put /user/{username}", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })
})


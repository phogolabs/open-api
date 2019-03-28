package test_test

import (
  "net/http"
  "net/http/httptest"
  
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/go-chi/chi"
)

var _ = Describe("TestAPI", func() {
  var (
    router   *chi.Mux
    request  *http.Request
    recorder *httptest.ResponseRecorder
  )

  BeforeEach(func() {
    router = chi.NewRouter()
    recorder = httptest.NewRecorder()

    service := &handler.TestAPI{}
    service.Mount(router)
  })

  Describe("Get /test", func() {
    BeforeEach(func() {
      //TODO: initialize the request here
    })

    //TODO: Implement tests
  })
})


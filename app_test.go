package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/macaron.v1"
)

func Test_Zenpress(t *testing.T) {
	Convey("App service", t, func() {
		go main()

		resp := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "http://localhost:9000/", nil)
		So(err, ShouldBeNil)

		m := macaron.New()
		m.ServeHTTP(resp, req)

	})

}

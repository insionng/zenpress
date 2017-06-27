package switchr

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/insionng/makross"
	"github.com/stretchr/testify/assert"
)

func TestSwitchr(t *testing.T) {
	e := makross.New()
	req := httptest.NewRequest(makross.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec, makross.NotFoundHandler)
	config := SwitchrConfig{
		Theme: "default",
	}

	h := SwitchrWithConfig(config)

	// File found
	req = httptest.NewRequest(makross.GET, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	//m := makross.New()
	//m.Use(SwitchrWithConfig(config))
	//m.ServeHTTP(rec, req)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/cfg/getCfg", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "getCfg OK", w.Body.String())
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodPost, "/cfg/chgCfg", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "chgCfg OK", w.Body.String())
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, "/data/getData", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "getData OK", w.Body.String())
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodPost, "/data/chgData", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "chgData OK", w.Body.String())
	assert.Nil(t, err)
}

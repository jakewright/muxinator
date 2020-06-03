package muxinator

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRouter_Match(t *testing.T) {
	router := NewRouter()

	router.Get("/foo", fakeHandler)
	router.Post("/bar", fakeHandler)

	getFoo, err := http.NewRequest(http.MethodGet, "/foo", nil)
	require.NoError(t, err)
	postFoo, err := http.NewRequest(http.MethodPost, "/foo", nil)
	require.NoError(t, err)

	getBar, err := http.NewRequest(http.MethodGet, "/bar", nil)
	require.NoError(t, err)
	postBar, err := http.NewRequest(http.MethodPost, "/bar", nil)
	require.NoError(t, err)

	require.True(t, router.Match(getFoo))
	require.True(t, router.Match(postBar))

	require.False(t, router.Match(postFoo))
	require.False(t, router.Match(getBar))
}

var fakeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

package ngore

import (
	"git.okki.hu/garric/ngore/internal"
	"git.okki.hu/garric/ngore/login"
	"git.okki.hu/garric/ngore/search"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewApi(t *testing.T) {

	t.Run("new api not nil", func(t *testing.T) {
		ng := New(&http.Client{}, "baseUrl")
		assert.NotNil(t, ng)
	})

	t.Run("redirect is disabled", func(t *testing.T) {
		client := &http.Client{}
		_ = New(client, "baseUrl")
		assert.Equal(t, http.ErrUseLastResponse, client.CheckRedirect(nil, nil))
	})

	t.Run("cookie jar initialized", func(t *testing.T) {
		ng := New(&http.Client{}, "baseUrl").(*api)
		assert.NotNil(t, ng.client.Jar)
	})

}

func TestNewDefaultApi(t *testing.T) {

	t.Run("client initialized", func(t *testing.T) {
		ng := Default("baseUrl").(*api)
		assert.NotNil(t, ng.client)
	})

	t.Run("default timeout", func(t *testing.T) {
		ng := Default("baseUrl").(*api)
		assert.Greater(t, ng.client.Timeout, time.Duration(0))
	})

	t.Run("default cookie jar", func(t *testing.T) {
		ng := Default("baseUrl").(*api)
		assert.NotNil(t, ng.client.Jar)
	})

}

func TestApi_Login(t *testing.T) {

	t.Run("successful login", func(t *testing.T) {
		first := true
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if first {
				first = false
				w.Header().Set("Set-Cookie", "PHPSESSID=foo")
				http.Redirect(w, r, "https://example.com/index.php", http.StatusFound)
			} else {
				_, _ = w.Write([]byte(`<link rel="alternate" href="/rss.php?key=abc123"`))
			}
			w.Header().Set("Set-Cookie", "PHPSESSID=foo")
			http.Redirect(w, r, "https://example.com/index.php", http.StatusFound)
		}))
		defer server.Close()
		ng := apiWithMockClient(server)
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.Nil(t, err)
	})

	t.Run("key extraction error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Set-Cookie", "PHPSESSID=foo")
			http.Redirect(w, r, "https://example.com/index.php", http.StatusFound)
		}))
		defer server.Close()
		ng := apiWithMockClient(server)
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.Error(t, err) // key is missing
	})

	t.Run("username missing", func(t *testing.T) {
		ng := Default("foo")
		assert.Error(t, ng.Login(&login.BasicAuth{Password: "bar"}))
	})

	t.Run("password missing", func(t *testing.T) {
		ng := Default("foo")
		assert.Error(t, ng.Login(&login.BasicAuth{UserName: "bar"}))
	})

	t.Run("invalid login", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://example.com?problema=1", http.StatusFound)
		}))
		defer server.Close()
		ng := apiWithMockClient(server)
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.ErrorContains(t, err, internal.ErrLoginInvalidCredentials)
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "test error", http.StatusInternalServerError)
		}))
		defer server.Close()
		ng := apiWithMockClient(server)
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.ErrorContains(t, err, internal.ErrLoginUnexpectedResponse)
	})

	t.Run("other error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		ng := apiWithMockClient(server)
		server.Close() // close server to cause an error in login
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.Error(t, err)
	})

}

func TestApi_Search(t *testing.T) {

	t.Run("search api login required", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, internal.LocationLogin, http.StatusFound)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		_, err := api.Search(&search.Params{})
		assert.ErrorContains(t, err, internal.ErrUserNotLoggedIn)
	})

	t.Run("search api server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		_, err := api.Search(&search.Params{})
		assert.Error(t, err)
	})

	t.Run("search api network error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		api := apiWithMockClient(server)
		server.Close() // close server to cause an error
		_, err := api.Search(&search.Params{})
		assert.Error(t, err)
	})

	t.Run("search api parse error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`foo`))
			w.WriteHeader(http.StatusOK)
		}))
		api := apiWithMockClient(server)
		_, err := api.Search(&search.Params{})
		assert.Error(t, err)
	})

}

func TestApi_Activity(t *testing.T) {

	t.Run("activity api login required", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, internal.LocationLogin, http.StatusFound)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		_, err := api.Activity()
		assert.ErrorContains(t, err, internal.ErrUserNotLoggedIn)
	})

	t.Run("activity api network error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		api := apiWithMockClient(server)
		server.Close() // close server to cause an error
		_, err := api.Activity()
		assert.Error(t, err)
	})

	t.Run("activity api server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		_, err := api.Activity()
		assert.Error(t, err)
	})

}

func apiWithMockClient(mockServer *httptest.Server) Api {
	return New(mockServer.Client(), mockServer.URL)
}

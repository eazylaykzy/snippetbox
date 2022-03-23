package main

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/eazylaykzy/snippetbox/pkg/models/mock"
	"github.com/golangcollege/sessions"
)

// newTestApplication helper returns an instance of our application struct containing mocked dependencies.
func newTestApplication(t *testing.T) *application {
	// Create an instance of the template cache.
	templateCache, err := newTemplateCache("./../../ui/html/")
	if err != nil {
		t.Fatal(err)
	}

	// Create a session manager instance, with the same settings as production.
	session := sessions.New([]byte("3dSm5MnygFHh7XidAtbskXrjbwfoJcbJ"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	// Initialize the dependencies, using the mocks for the loggers and
	// database models.
	return &application{
		session:       session,
		users:         &mock.UserModel{},
		templateCache: templateCache,
		snippets:      &mock.SnippetModel{},
		infoLog:       log.New(io.Discard, "", 0),
		errorLog:      log.New(io.Discard, "", 0),
	}
}

// Define a custom testServer type which anonymously embeds a httptest.Server
// instance.
type testServer struct {
	*httptest.Server
}

// newTestServer helper which initializes and returns a new instance of our custom testServer type
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	// Initialize a new cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add the cookie jar to the client, so that response cookies are stored and then sent with subsequent requests.
	ts.Client().Jar = jar

	// Disable redirect-following for the client, essentially this function is called after a 3xx response is received by
	// the client, and returning the http.ErrUseLastResponse error forces it to immediately return the received response
	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

// Implement a get method on our custom testServer type. This makes a GET request to a given
// url path on the test server, and returns the response status code, headers and body
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	return rs.StatusCode, rs.Header, body
}

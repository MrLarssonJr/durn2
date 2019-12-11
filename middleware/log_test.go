package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
)

// Test Log.Access, ensures that it logs what is expected
func TestLogAccess(t *testing.T) {
	//Setup
	var builder strings.Builder
	prevWriter := log.Writer()
	prevFlags := log.Flags()
	log.SetOutput(&builder)
	log.SetFlags(0)

	//Test
	{
		const mockRemote = "mockRemote"
		const mockPath = "/mock"
		expectedLog := fmt.Sprintf("%s accessed by %s\n", mockPath, mockRemote)

		//Arrange
		var next http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
		w := NewMockResponseWriter()
		r, err := http.NewRequest("GET", mockPath, nil)
		if err != nil {
			t.Fatalf("Couldn't create request: %s", err.Error())
		}
		r.RemoteAddr = mockRemote

		//Act
		h := Log.Access(next)
		h.ServeHTTP(w, r)

		//Assert
		actualLog := builder.String()
		if actualLog != expectedLog {
			t.Errorf("logged text not as expected. expected: \"%s\", actual: \"%s\"", expectedLog, actualLog)
		}
	}

	//Tear down
	log.SetOutput(prevWriter)
	log.SetFlags(prevFlags)
}

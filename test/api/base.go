package apitest

import (
	"net/http"
	"testing"

	"github.com/audriusdai/eventing-api/web"
	"github.com/gavv/httpexpect"
)

func testClient(t *testing.T) *httpexpect.Expect {
	engine := web.SetupEngine()

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(engine),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	return e
}

package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExcepted bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering the Go Template"},
	{"go_page_no_template", "go", "no-file", true, "no error rendering no existent go template when one is expected"},
	{"jet_page", "jet", "home", false, "error rendering the Jet Template"},
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering no existent jet template when one is expected"},
	{"invalid_render_engin", "", "home", true, "no error rendering with non existend template engine"},
}

func TestRender_Page(t *testing.T) {

	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/some-test-url", nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExcepted {
			if err == nil {
				t.Errorf("%s, %s", e.name, e.errorMessage)
			}

		} else {
			if err != nil {
				t.Errorf("%s, %s, %s", e.name, e.errorMessage, err.Error())
			}

		}

	}

}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering Page", err)
	}

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non existing go Page", err)
	}

}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "jet"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering Page", err)
	}

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non existing go Page", err)
	}

}

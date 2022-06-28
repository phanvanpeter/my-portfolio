package render

import (
	"net/http/httptest"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	Template(w, "home.page")
	Template(w, "about.page")
}

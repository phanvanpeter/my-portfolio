package internal

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	TmplFolder = "../templates/"

	os.Exit(m.Run())
}


// go test -coverprofile=coverage.out && go tool cover -html=coverage.out
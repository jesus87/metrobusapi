package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Setenv("METROBUS_API_URL", "https://datos.cdmx.gob.mx/api/records/1.0/search/?")
	os.Setenv("POSITIONS_API_PAGESIZE", "10")
	main()
}

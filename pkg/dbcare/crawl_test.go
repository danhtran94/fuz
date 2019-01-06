package dbcare_test

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/danhtran94/fuz/pkg/dbcare"
)

func TestGetSchemaFrom(t *testing.T) {
	err := dbcare.SetupDBClient("postgresql://postgres:postgres@127.0.0.1:5432/tse-case?sslmode=disable", true)
	assert.NoError(t, err)

	cols, err := dbcare.GetColTypes("cases")
	assert.NoError(t, err)

	fmt.Println(cols)
}

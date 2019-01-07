package dbcare

import (
	"fmt"
	"strings"

	funk "github.com/thoas/go-funk"
	lk "github.com/ulule/loukoum"

	"github.com/danhtran94/fuz/pkg/fake"
	"github.com/danhtran94/fuz/pkg/template"
)

func GetSQL(t template.Template) string {

	var sb strings.Builder

	cols := funk.Get(t.Def, "Name").([]string)

	for range make([]int, t.Rows) {

		vals := fake.FakeRowVal(t.Def)
		orderedVals := funk.Map(cols, func(colName string) interface{} {
			return vals[colName]
		}).([]interface{})

		pairs := []interface{}{}

		for i, col := range cols {
			pairs = append(pairs, lk.Pair(col, orderedVals[i]))
		}

		builder := lk.Insert(t.Table).Set(pairs...)
		sb.WriteString(fmt.Sprintf("%s\n", builder.String()))
	}

	return sb.String()
}

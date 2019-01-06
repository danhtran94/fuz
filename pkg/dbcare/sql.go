package dbcare

import (
	"fmt"

	funk "github.com/thoas/go-funk"
	lk "github.com/ulule/loukoum"

	"github.com/danhtran94/fuz/pkg/fake"
	"github.com/danhtran94/fuz/pkg/template"
)

func GetSQL(t template.Template) string {

	cols := funk.Get(t.Def, "Name").([]string)

	for range make([]int, t.Rows) {
		vals := fake.FakeRowVal(t.Def)

		// cols := funk.Keys(result).([]string)
		orderedVals := funk.Map(cols, func(colName string) interface{} {
			return vals[colName]
		}).([]interface{})

		pairs := []interface{}{}

		for i, col := range cols {
			pairs = append(pairs, lk.Pair(col, orderedVals[i]))
		}

		// qu := DBClient.InsertInto(t.Table).Columns(cols...).Values(vals...).String()

		// jsonr, _ := json.Marshal(result)
		builder := lk.Insert(t.Table).Set(pairs...)

		fmt.Println(builder.String())

		// fmt.Printf("%+v | %+v\n", cols, orderedVals)
	}

	return ""
}

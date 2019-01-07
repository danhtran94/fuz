package fake

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/danhtran94/fuz/pkg/template"

	"github.com/brianvoe/gofakeit"
	fi "github.com/brianvoe/gofakeit"
)

func init() {
	fi.Seed(0)
}

var FakeMap Fake = map[string]interface{}{
	"Name":     fi.Name,
	"Email":    fi.Email,
	"Phone":    fi.Phone,
	"Date":     fi.Date,
	"Number":   fi.Number,
	"Bool":     fi.Bool,
	"UUID":     fi.UUID,
	"Country":  fi.Country,
	"ImageURL": fi.ImageURL,
	"Numerify": fi.Numerify,
	"NowUTC": func() time.Time {
		return time.Now().UTC()
	},
}

type Fake map[string]interface{}

func (f Fake) Gen(name string, params ...interface{}) (interface{}, error) {
	// fmt.Println(name, params)
	fn := reflect.ValueOf(f[name])

	if len(params) != fn.Type().NumIn() {
		return nil, errors.New("the number of params is not adapted")
	}

	inputContainer := make([]reflect.Value, len(params))
	for index, param := range params {
		switch param.(type) {
		case float64:
			inputContainer[index] = reflect.ValueOf(int(param.(float64)))
		default:
			inputContainer[index] = reflect.ValueOf(param)
		}
	}

	result := fn.Call(inputContainer)
	if len(result) == 0 {
		return nil, errors.New("the number of output is not adapted")
	}

	return result[0].Interface(), nil
}

func FakeRowVal(colDefs []*template.Def) map[string]interface{} {
	result := map[string]interface{}{}

	for _, def := range colDefs {
		if def.JSON != nil {
			val, err := json.Marshal(defFake(*def))
			if err != nil {
				return nil
			}

			result[def.Name] = string(val)
		} else {
			result[def.Name] = defFake(*def)
		}
	}

	return result
}

func defFake(def template.Def) interface{} {
	if jatt := def.JSON; jatt != nil && def.Gen == nil {
		if arrNum := def.JSONArray; arrNum != nil {
			return jsonAVal(jatt, *arrNum)
		}
		return jsonOVal(jatt)
	}

	if lit := def.Gen.Literal; lit != nil {
		return lit
	}

	if pattern := def.Gen.Pattern; pattern != "" {
		return gofakeit.Generate(pattern)
	}

	// fmt.Printf("%+v\n", def)

	methodName := def.Gen.Method
	args := def.Gen.Args

	val, err := FakeMap.Gen(methodName, args...)
	if err != nil {
		return nil
	}

	return val
}

func jsonOVal(def []*template.Def) interface{} {
	o := map[string]interface{}{}

	for _, d := range def {
		if d.Name == "" {
			return defFake(*d)
		}

		o[d.Name] = defFake(*d)
	}

	return o
}

func jsonAVal(def []*template.Def, num int) interface{} {
	a := []interface{}{}

	for range make([]int, num) {
		a = append(a, jsonOVal(def))
	}

	return a
}

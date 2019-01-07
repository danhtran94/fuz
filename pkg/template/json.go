package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

var defPath = "./templates"

type Def struct {
	Name      string `json:"name"`
	DBType    string `json:"dbType,omitempty"`
	Gen       *Gen   `json:"gen"`
	JSON      []*Def `json:"json,omitempty"`
	JSONArray *int   `json:"jsonArray,omitempty"`
}

type Gen struct {
	Literal interface{}   `json:"literal,omitempty"`
	Pattern string        `json:"pattern,omitempty"`
	Method  string        `json:"method"`
	Args    []interface{} `json:"args,omitempty"`
}

type Template struct {
	Table string `json:"table"`
	Rows  int    `json:"rows"`
	Def   []*Def `json:"def"`
}

func NewTemplate(tableName string, colTypes map[string]string) Template {
	template := Template{
		Table: tableName,
		Def:   make([]*Def, 0),
	}

	for colName, dbType := range colTypes {
		def := Def{
			Name:   colName,
			DBType: dbType,
		}

		switch dbType {
		case "jsonb":
			def.JSON = []*Def{{}}
		case "json":
			def.JSON = []*Def{{}}
		}

		template.Def = append(template.Def, &def)
	}

	return template
}

func Load() ([]Template, error) {
	templates := []Template{}

	files, err := filepath.Glob(fmt.Sprintf("%s/*", defPath))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		templateJSON, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println(file, err)
			continue
		}

		template := Template{}
		err = json.Unmarshal(templateJSON, &template)
		if err != nil {
			log.Println(file, err)
			continue
		}

		templates = append(templates, template)
	}

	return templates, nil
}

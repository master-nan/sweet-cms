package initialize

import (
	"github.com/gin-contrib/multitemplate"
	"html/template"
	"path/filepath"
	"strings"
)

func LoadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob("templates/layout/*")
	if err != nil {
		panic(err.Error())
	}
	adminBase, err := filepath.Glob("templates/basic/admin/*")
	if err != nil {
		panic(err.Error())
	}
	adminFiles, err := filepath.Glob("templates/admin/**/*")
	if err != nil {
		panic(err.Error())
	}
	funcMap := template.FuncMap{
		"Contains": strings.Contains,
	}
	for _, admin := range adminFiles {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, admin)
		if filepath.Base(admin) != "login.html" {
			for _, b := range adminBase {
				files = append(files, b)
			}
		}
		r.AddFromFilesFuncs(filepath.Base(admin), funcMap, files...)
	}
	return r
}

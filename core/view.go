package core

import(
	"io"
	"html/template"
)

func View(wr io.Writer, layout string, data interface{}) {
  tpl := template.New("template")
  tpl = template.Must(
	  tpl.ParseFiles("templates/shared/layout.html","templates/shared/head.html","templates/shared/header.html","templates/shared/main.html","templates/shared/footer.html","templates/" + layout))


  err := tpl.ExecuteTemplate(wr, "layout", data)

  if err != nil {
    panic(err.Error())
  }
}
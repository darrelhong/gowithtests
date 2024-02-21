package blogrenderer

import (
	"blogposts"
	"embed"
	"io"
	"text/template"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p blogposts.Post) error {
	tmpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "post.gohtml", p); err != nil {
		return err
	}

	return nil
}

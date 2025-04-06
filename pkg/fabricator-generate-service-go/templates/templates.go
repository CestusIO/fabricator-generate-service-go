package templates

import (
	"embed"
	"io/fs"
)

//go:embed *
var embededTemplates embed.FS

func GetTemplates() fs.ReadDirFS {
	return embededTemplates
}

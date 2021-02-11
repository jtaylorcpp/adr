package adr

import (
	"fmt"
	"strings"
)

type ADR struct {
	RootPath string
	Title string
	Index uint
}

func (a ADR) FilePath() string {
	return fmt.Sprintf("%s/%s.md", a.RootPath, a.FileName())
}

func (a ADR) FileName() string {
	dashTitle := strings.ReplaceAll(a.Title, " ", "-")
	return fmt.Sprintf("%d-%s", a.Index, dashTitle)
}
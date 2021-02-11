package adr

import (
	"os"
	"fmt"
	"text/template"
	"github.com/gobuffalo/packr"
	log "github.com/sirupsen/logrus"
)

var templateBox packr.Box

func init() {
	templateBox = packr.NewBox("./templates")
}

func NewADR(adr ADR) error {
	t, err := templateBox.FindString("adr.md")
	if err != nil {
		return fmt.Errorf("error finding template: %s", err.Error())
	}

	tmpl, err := template.New("adrtmpl").Parse(t)
	if err != nil {
		return fmt.Errorf("error parsing template: %s", err.Error())
	}

	adrFile, err := os.Create(adr.FilePath())
	if err != nil {
		return fmt.Errorf("error creating adr file: %s", err.Error())
	}

	err = tmpl.Execute(adrFile, adr)
	if err != nil {
		return fmt.Errorf("unable to render template: %s", err.Error())
	}

	adrFile.Close()
	log.Infof("ADR %s created\n", adr.FileName())
	return nil
}
package cmd 

import (
	log "github.com/sirupsen/logrus"
	"github.com/jtaylorcpp/adr/config"
	"github.com/jtaylorcpp/adr/adr"
	"github.com/spf13/cobra"
)

var title string
// var projectName string (in package other file)

func init() {
	rootCmd.AddCommand(newAdrCmd)
	newAdrCmd.Flags().StringVarP(&projectName, "project-name", "n", "", "name of adr project")
	newAdrCmd.Flags().StringVarP(&title, "adr-title", "t", "", "title of new adr")
	newAdrCmd.MarkFlagRequired("project-name")
	newAdrCmd.MarkFlagRequired("adr-title")
}

var newAdrCmd = &cobra.Command{
	Use:   "new-adr",
	Aliases: []string{"new"},
	Short: "create a new architecture design record",
	Run: func(cmd *cobra.Command, args []string) {
		if !config.ProjectExists(projectName) {
			log.Fatalln("project does not exist, cannot create new adr")
		}

		dir, err := config.GetProjectDir(projectName)
		if err != nil {
			log.Fatalf("unable to find project directory: %s\n", err.Error())
		}

		index, err := config.GetProjectAdrIndex(projectName)
		if err != nil {
			log.Fatalf("unable to find project index: %s\n", err.Error())
		}

		newADR := adr.ADR{
			RootPath: dir,
			Title: title,
			Index: index,
		}

		err = adr.NewADR(newADR)
		if err != nil {
			log.Fatalf("unable to create new adr: %s\n", err.Error())
		}

	  },
}


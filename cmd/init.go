package cmd 

import (
	log "github.com/sirupsen/logrus"
	"github.com/jtaylorcpp/adr/config"
	"github.com/jtaylorcpp/adr/utils"
	"github.com/spf13/cobra"
)

var projectName string
var projectPath string

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&projectName, "project-name", "n", "", "name of adr project")
	initCmd.Flags().StringVarP(&projectPath, "project-dir", "d", "", "directory to store adr files (will be created if it doesnt exist")
	initCmd.MarkFlagRequired("project-name")
	initCmd.MarkFlagRequired("project-dir")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize new adr project",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("adding project %s with root dir %s\n", projectName, projectPath)
		if err := config.CreateProject(projectName, projectPath); err != nil {
			log.Fatalf("error creating new project: %s\n", err.Error())
		}

		if err := utils.CreateDir(projectPath); err != nil {
			log.Fatalf("error creating project dir: %s", err.Error())
		}
	  },
}


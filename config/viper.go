package config
import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"errors"
)

func init() {
	viper.SetConfigName("adr")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			log.Fatalf("error reading in adr config file: %s\n", err.Error())
		}
	}
}

func Save() {
	viper.WriteConfig()
}

func ProjectExists(projectName string) bool {
	return viper.IsSet(projectName)
}

func GetProjectDir(projectName string) (string, error) {
	if !ProjectExists(projectName){
		return "", errors.New("project does not exist")
	}

	if viper.IsSet(projectName+".dir") {
		return viper.GetString(projectName+".dir"), nil
	} 

	return "", errors.New("no directory has been set for that project")
}

func GetProjectAdrIndex(projectName string) (uint, error) {
	if !ProjectExists(projectName){
		return 0, errors.New("project does not exist")
	}

	if viper.IsSet(projectName+".index") {
		idx := viper.GetUint(projectName+".index")
		// since we are using the current index increment index counter
		viper.Set(projectName+".index", idx +1)
		viper.WriteConfig()
		return idx , nil
	}

	return 0, errors.New("no index has been set for that project")
}

func CreateProject(projectName, projectDir string) error {
	if ProjectExists(projectName) {
		return errors.New("project already exists")
	}

	viper.Set(projectName, map[string]interface{}{"dir":projectDir,"index": 0})
	viper.WriteConfig()
	return nil
}
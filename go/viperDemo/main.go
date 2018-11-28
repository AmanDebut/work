package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Port int
	Name string
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		// Setting Defaults
		v.SetDefault(key, value)
	}

	// Reading Config Files
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AddConfigPath("..")

	// Binding Env to Specific Keys
	v.BindEnv("name", "USERNAME")      // bind to ENV "USERNAME"
	os.Setenv("USERNAME", "Aman Jain") // typically done outside of the app

	v.SetEnvPrefix("debut")            // Becomes "DEBUT_"
	os.Setenv("DEBUT_ENVPORT", "4000") // typically done outside of the app

	// Automatic Environment Binding
	v.AutomaticEnv()

	err := v.ReadInConfig()
	return v, err
}

func main() {
	v1, err := readConfig("env", map[string]interface{}{
		"port":     3000,
		"hostname": "192.168.0.57",
		"auth": map[string]string{
			"username": "Deepak",
			"password": "debut123456",
		},
	})
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v", err))
	}

	port := v1.GetInt("port")
	envport := v1.GetInt("envport")
	hostname := v1.GetString("hostname")
	authusername := v1.GetString("auth.username")
	username := v1.GetString("name")

	// Checking if a Key has been Set
	if v1.IsSet("port") {
		fmt.Println("port is set")
	}

	fmt.Printf("Reading env for name = %s\n", username)
	fmt.Printf("Reading env for port = %d\n", envport)
	fmt.Printf("Reading config for port = %d\n", port)
	fmt.Printf("Reading config for hostname = %s\n", hostname)
	fmt.Printf("Reading config for auth username = %#v\n", authusername)

	c := v1.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		panic(fmt.Errorf("unable to marshal config to YAML: %v", err))
	}
	fmt.Println(string(bs))
}

package internal

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func RunViper() {
	fromYAML()
	fromEnv()
}

func fromYAML() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("viper from yaml %+v\n", cfg)
}

func fromEnv() {
	var cfg Config
	v := viper.New()

	// Set environment variables
	envVars := map[string]string{
		"SERVER_PORT":       "8888",
		"SERVER_HOST":       "localhost",
		"DATABASE_USER":     "user",
		"DATABASE_PASSWORD": "password",
		"DATABASE_NAME":     "db_name",
		"ENV_NAME":          "dev_env",
	}
	for key, value := range envVars {
		err := os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}

	v.AutomaticEnv()

	// Extract map tags and bind them to environment variables
	r := extractMapTags(cfg, "")
	for _, m := range r {
		for tag, envName := range m {
			if err := v.BindEnv(tag, envName); err != nil {
				panic(err)
			}
		}
	}

	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("viper from env %+v\n", cfg)
}

func extractMapTags(config interface{}, previousTag string) []map[string]string {
	var mapTags []map[string]string

	v := reflect.ValueOf(config)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("mapstructure")

		if tag != "" {
			if previousTag != "" {
				tag = previousTag + "." + tag
			}

			envName := strings.ToUpper(strings.ReplaceAll(tag, ".", "_"))
			if field.Type.Kind() != reflect.Struct {
				mapTags = append(mapTags, map[string]string{tag: envName})
			}
		}

		if field.Type.Kind() == reflect.Struct {
			subMapTags := extractMapTags(v.Field(i).Interface(), tag)
			mapTags = append(mapTags, subMapTags...)
		}
	}

	return mapTags
}

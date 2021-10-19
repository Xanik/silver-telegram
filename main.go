package main

import (
	"fmt"
	httpSrv "go-challenge/http"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
	conf := viper.GetStringMapString("challenger")
	httpRouter := httpSrv.MountServer()

	log.Print(fmt.Sprintf("HTTP service running on %v", conf["port"]))

	if err := http.ListenAndServe(":"+conf["port"], httpRouter); err != nil {
		log.Fatalf(err.Error())
	}
}

// Copyright 2020 Vladislav Smirnov

package main

import "github.com/spf13/viper"

// Slogger is a global logger object
var Slogger *Logger

var s *Server
var api *API

func main() {
	viper.SetDefault("MaxPlayers", 128)
	viper.SetDefault("LogfilePath", "logs.txt")
	viper.SetDefault("WorldSize", 100)
	viper.SetDefault("MinimalDistance", 60.0)
	viper.SetDefault("EdgeDistance", 140.0)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.ReadInConfig()

	Slogger = NewLogger(viper.GetString("LogfilePath"))
	s = NewServer()
	api = NewAPI(s)
	s.Room = NewRoom(viper.GetInt("MaxPlayers"), viper.GetInt("WorldSize"))

	go s.LaunchPaytimeTimer()
	Slogger.Log("Started server at port 34000.")

	api.Start()
}

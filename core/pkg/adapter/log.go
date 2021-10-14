package adapter

import (
	phgo "github.com/pinkhello/ph-go"
)

type (
	LogConf struct {
		Path    string
		AppName string
	}

	LogAdapter struct {
		LogConf
		AccessLog  *phgo.Logger
		TrackLog   *phgo.Logger
		RequestLog *phgo.Logger
		ErrorLog   *phgo.Logger
	}
)

func Init() {

}

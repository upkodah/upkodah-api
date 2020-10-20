package cmd

import (
	"github.com/upkodah/upkodah-api/pkg/api"
	"github.com/upkodah/upkodah-api/pkg/env"
)

func Run() {
	api.RunAPI()
}

func init() {
	env.InitDefault()
}

package core

import (
	"flag"
	"os"
)

type Environment struct {
	RunPath string
}

var Env Environment

func SetupEnv() {
	runpath := flag.String("runpath", ".", "Input the runpath")
	flag.Parse()
	Env.RunPath = *runpath

	os.Mkdir(Env.RunPath+"/data", os.ModePerm)
	os.Mkdir(Env.RunPath+"/run", os.ModePerm)
}

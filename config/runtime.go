package config

import (
	"log"
	"runtime"
)

func ConfigRuntime() {
	InitEnvConfigs()
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)

	log.Printf("Running with %d CPU", nuCPU)
}

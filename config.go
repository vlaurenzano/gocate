package main

import (
	"os"
	"strconv"
	"fmt"
)

type Configuration struct {
	StorageLocation string
	BuildIndexStrategy string
	NumberOfBuildJobs int
}

var config Configuration

//Config returns our configuration.
//If config is not yet initialized it is first initialized by makeConfig.
func Config() Configuration{
	if config == (Configuration{}) {
		makeConfig()
	}
	return config
}

//makeConfig initializes our configuation Struct.
func makeConfig() {
	config.StorageLocation = os.Getenv("GOCATE_DB_LOCATION")
	if config.StorageLocation == "" {
		config.StorageLocation = "/tmp"
	}
	config.BuildIndexStrategy = os.Getenv("GOCATE_BUILD_INDEX_STRATEGY")
	if config.BuildIndexStrategy == "" {
		config.BuildIndexStrategy = "Concurrent"
	}
	str := os.Getenv("GOCATE_N_BUILD_JOBS")
	if str == "" {
		config.NumberOfBuildJobs = 100
	} else {
		n, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(os.Stderr, "Unparsable configuration value for GOCATE_N_BUILD_JOBS")
			os.Exit(2)
		}
		config.NumberOfBuildJobs = n
	}

}


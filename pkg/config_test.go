package pkg

import "testing"
import "os"

func TestConfigWithDefaults(t *testing.T) {
	os.Unsetenv("GOCATE_N_BUILD_JOBS")
	os.Unsetenv("GOCATE_DB_LOCATION" )
	os.Unsetenv("GOCATE_N_BUILD_JOBS")
	c := Config()
	if c.NumberOfBuildJobs != 100 || c.BuildIndexStrategy != "Concurrent" || c.StorageLocation != "/tmp" {
		t.Error("Making of config with defaults failed")
	}
}

func TestConfigSettingValues(t *testing.T){
	os.Setenv("GOCATE_N_BUILD_JOBS", "5")
	makeConfig()
	c := Config()
	if c.NumberOfBuildJobs != 5 {
		t.Error("Making of config with non default value failed")
	}
}




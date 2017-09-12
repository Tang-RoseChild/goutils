package configutils

import (
	"os"

	"github.com/Tang-RoseChild/goutils/valid"
)

type _BaseConfig struct {
	URL        string `long:"url" description:"remote config url" `
	ConfigPath string `long:"config" description:"yaml config file path"`
}

type _Config struct {
	ServerAddr string `long:"addr" description:"server addr" yaml:"serverAddr" default:"127.0.0.1" required:"true"`
	Secret     string `envconfig:"SECRET" long:"secret" required:"true"`
}

func ExampleGetConfig() {
	var baseConfig _BaseConfig
	restArgs, err := LoadFromCommandLine(os.Args, &baseConfig)
	if err != nil {
		panic(err)
	}

	var appConfig _Config
	_, err = LoadFromCommandLine(restArgs, &appConfig)

	if err != nil && baseConfig.ConfigPath != "" {
		err = LoadFromYaml(baseConfig.ConfigPath, &appConfig)
	}

	// 'cause yaml not provide require tag,so can not check the require one. use require tag in struct field to put all args in one kind in case of mix
	if err != nil && baseConfig.URL != "" && !validutils.FieldRequiredValid(appConfig) {
		err = LoadFromRemote("GET", baseConfig.URL, nil, nil, &appConfig)
	}
	if err != nil {
		panic(err)
	}
	// do otherthings

}

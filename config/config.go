package configutils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	flags "github.com/jessevdk/go-flags"
	yaml "gopkg.in/yaml.v2"
)

func LoadFromYaml(configPath string, out interface{}) error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, out)
}

func LoadFromRemote(method string, url string, body io.Reader, headers map[string]string, out interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, out)
}

// LoadFromCommandLine args can be os.Args, or rest args after parsed
func LoadFromCommandLine(args []string, out interface{}) (restArgs []string, err error) {
	restArgs, err = flags.NewParser(out, flags.HelpFlag|flags.PrintErrors|flags.PassDoubleDash|flags.IgnoreUnknown).ParseArgs(args)
	if flagErr, ok := err.(*flags.Error); ok && flagErr.Type == flags.ErrHelp {
		return restArgs, nil
	}

	return restArgs, err
}

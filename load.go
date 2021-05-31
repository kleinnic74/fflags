package fflags

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type state struct {
	f map[string]interface{}
}

type yamlLoader struct {
	open func() (io.ReadCloser, error)
}

func YamlFile(path string) Loader {
	return yamlLoader{
		open: func() (io.ReadCloser, error) {
			return os.Open(path)
		},
	}
}

func (l yamlLoader) Load() (map[string]bool, error) {
	in, err := l.open()
	if err != nil {
		return nil, err
	}
	defer in.Close()
	var specs struct {
		Features map[string]interface{} `yaml:"features"`
	}
	if err := yaml.NewDecoder(in).Decode(&specs); err != nil {
		return nil, err
	}
	features := make(map[string]bool)
	for k, v := range specs.Features {
		setFeatureState(features, k, v)
	}
	return features, nil
}

func setFeatureState(features map[string]bool, path string, f interface{}) {
	switch i := f.(type) {
	case bool:
		features[path] = i
	case string:
		fp := path + "." + i
		features[fp] = true
	case []interface{}:
		for _, v := range i {
			setFeatureState(features, path, v)
		}
	case map[string]interface{}:
		for k, v := range i {
			setFeatureState(features, path+"."+k, v)
		}
	default:
		// Error, ignore
	}
}

func Yaml(in io.Reader) Loader {
	return yamlLoader{
		open: func() (io.ReadCloser, error) {
			return io.NopCloser(in), nil
		},
	}
}

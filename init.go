package fflags

import "log"

type Loader interface {
	Load() (map[string]bool, error)
}

type errorLoader struct {
	err error
}

func (e errorLoader) Load() (map[string]bool, error) {
	return nil, e.err
}

func Init(loaders ...Loader) {
	for _, l := range loaders {
		features, err := l.Load()
		if err != nil {
			log.Printf("Failed to load features: %s", err)
			continue
		}
		for k, active := range features {
			feature := Define(k)
			if active {
				f.Enable(feature)
			} else {
				f.Disable(feature)
			}
		}
	}
}

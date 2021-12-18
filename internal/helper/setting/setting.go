package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

type ConfigFile struct {
	Name string
	Path string
	Type string
}

func NewSetting(configFiles ...ConfigFile) (*Setting, error) {
	if len(configFiles) == 0 {
		return DefaultSetting()
	}
	vp := viper.New()
	var err error
	for i, f := range configFiles {
		vp.SetConfigName(f.Name)
		vp.AddConfigPath(f.Path)
		if i == 0 {
			err = vp.ReadInConfig()
		} else {
			err = vp.MergeInConfig()
		}
		if err != nil {
			return nil, err
		}
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func DefaultSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(e fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}

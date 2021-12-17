package setting

import "github.com/spf13/viper"

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
	return &Setting{vp}, nil
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

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}

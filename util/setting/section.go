package setting

type MySQLSettingS struct {
	DSN                       string `mapstructure:"dsn" yaml:"dsn"`
	DefaultStringSize         int    `mapstructure:"default_string_size" yaml:"default_string_size"`
	DontSupportRenameIndex    bool   `mapstructure:"dont_support_rename_index" yaml:"dont_support_rename_index"`
	DisableDatetimePrecision  bool   `mapstructure:"disable_datetime_precision" yaml:"disable_datetime_precision"`
	DontSupportRenameColumn   bool   `mapstructure:"dont_support_rename_column" yaml:"dont_support_rename_column"`
	SkipInitializeWithVersion bool   `mapstructure:"skip_initialize_with_version" yaml:"skip_initialize_with_version"`
}

type JudgeSettingS struct {
	RunMode        string `mapstructure:"run_mode" yaml:"run_mode"`
	SubmissionPath string `mapstructure:"submission_path" yaml:"submission_path"`
	ResolutionPath string `mapstructure:"resolution_path" yaml:"resolution_path"`
}

type RabbitMQSettingS struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

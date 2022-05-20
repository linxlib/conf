package conf

// Option configures how fig loads the configuration.
type Option func(f *conf)

// File returns an option that configures the filename that conf
// looks for to provide the config values.
//
// The name must include the extension of the file. Supported
// file types are `yaml`, `yml`, `json` and `toml`.
//
//   conf.Load(&cfg, conf.File("config.toml"))
//
// If this option is not used then conf looks for a file with name `config.yaml`.
func File(name string) Option {
	return func(f *conf) {
		f.filename = name
	}
}

// IgnoreFile returns an option which disables any file lookup.
//
// This option effectively renders any `File` and `Dir` options useless. This option
// is most useful in conjunction with the `UseEnv` option when you want to provide
// config values only via environment variables.
//
//   conf.Load(&cfg, conf.IgnoreFile(), conf.UseEnv("my_app"))
func IgnoreFile() Option {
	return func(f *conf) {
		f.ignoreFile = true
	}
}

// Dirs returns an option that configures the directories that conf searches
// to find the configuration file.
//
// Directories are searched sequentially and the first one with a matching config file is used.
//
// This is useful when you don't know where exactly your configuration will be during run-time:
//
//   conf.Load(&cfg, conf.Dirs(".", "/etc/myapp", "/home/user/myapp"))
//
//
// If this option is not used then conf looks in the directory it is run from.
func Dirs(dirs ...string) Option {
	return func(f *conf) {
		f.dirs = dirs
	}
}

// Tag returns an option that configures the tag key that conf uses
// when for the alt name struct tag key in fields.
//
//  conf.Load(&cfg, conf.Tag("config"))
//
// If this option is not used then conf uses the tag `conf`.
func Tag(tag string) Option {
	return func(f *conf) {
		f.tag = tag
	}
}

// TimeLayout returns an option that configures the time layout that conf uses when
// parsing a time in a config file or in the default tag for time.Time fields.
//
//   conf.Load(&cfg, conf.TimeLayout("2006-01-02"))
//
// If this option is not used then conf parses times using `time.RFC3339` layout.
func TimeLayout(layout string) Option {
	return func(f *conf) {
		f.timeLayout = layout
	}
}

// UseEnv returns an option that configures conf to additionally load values
// from the environment.
//
//   conf.Load(&cfg, conf.UseEnv("my_app"))
//
// Values loaded from the environment overwrite values loaded by the config file (if any).
//
// Conf looks for environment variables in the format PREFIX_FIELD_PATH or
// FIELD_PATH if prefix is empty. Prefix is capitalised regardless of what
// is provided. The field's path is formed by prepending its name with the
// names of all surrounding fields up to the root struct. If a field has
// an alternative name defined inside a struct tag then that name is
// preferred.
//
//   type Config struct {
//     Build    time.Time
//     LogLevel string `conf:"log_level"`
//     Server   struct {
//       Host string
//     }
//   }
//
// With the struct above and UseEnv("myapp") conf would search for the following
// environment variables:
//
//   MYAPP_BUILD
//   MYAPP_LOG_LEVEL
//   MYAPP_SERVER_HOST
func UseEnv(prefix string) Option {
	return func(f *conf) {
		f.useEnv = true
		f.envPrefix = prefix
	}
}

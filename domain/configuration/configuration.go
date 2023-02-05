package configuration

type Configuration interface {
	Common() *Common
}

type (
	Config struct {
		Common Common
	}

	Common struct {
		Debug bool `env:"DEBUG"`
	}
)

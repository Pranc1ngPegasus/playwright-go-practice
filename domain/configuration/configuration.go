package configuration

type Configuration interface {
	Common() *Common
	GCP() *GCP
	Server() *Server
	Scan() *Scan
}

type (
	Config struct {
		Common Common
		GCP    GCP
		Server Server
		Scan   Scan
	}

	Common struct {
		Debug bool `env:"DEBUG"`
	}

	GCP struct {
		ProjectID string `env:"GCP_PROJECT_ID"`
	}

	Server struct {
		Port string `env:"PORT"`
	}

	Scan struct {
		Email    string `env:"SCAN_EMAIL"`
		Password string `env:"SCAN_PASSWORD"`
		TOTP     string `env:"SCAN_TOTP"`
	}
)

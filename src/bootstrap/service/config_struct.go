package service

import "time"

type Config struct {
	App struct {
		Name  string
		URL   string
		Port  string
		Debug bool
	}

	Database struct {
		Dirver string
	} `toml:"database"`

	Mysql struct {
		Dbname   string
		Username string
		Password string
	} `toml:"mysql"`

	Website struct {
		static_uri string
		site_title string
		copy_right string
	}

	Redis struct {
		// Network "tcp"
		Network string
		// Addr "127.0.0.1:6379"
		Addr string
		// Password string .If no password then no 'AUTH'. Default ""
		Password string
		// If Database is empty "" then no 'SELECT'. Default ""
		Database string
		// MaxIdle 0 no limit
		MaxIdle int
		// MaxActive 0 no limit
		MaxActive int
		// IdleTimeout  time.Duration(5) * time.Minute
		IdleTimeout time.Duration
		// Prefix "myprefix-for-this-website". Default ""
		Prefix string
	}

	Session struct {
		Cookie  string
		Expires time.Duration
		Dirver  string
	}

	Image struct {
		ImageLib        string   `toml:"image_lib"`
		ImagePath       string   `toml:"image_path"`
		ImageURL        string   `toml:"image_ur"`
		ImageOrg        string   `toml:"image_org"`
		ImageTmp        string   `toml:"image_tmp"`
		ImageTypes      []string `toml:"image_types"`
		WaterMark       string   `toml:"water_mark"`
		ImageCategroies []string
		// ImageCategory imageCategory `toml:"imageCategory"`
		ImageCategory struct {
			CarLogo CarLogo `toml:"carLogo"`
		} `toml:"imageCategory"`
	}
}

type CarLogo struct {
	Paths string   `toml:"paths"`
	Sizes []string `toml:"sizes"`
}
type Category interface {
	GetPath() string
	GetSizes() []string
}

func (c *CarLogo) GetPath() string {
	return c.Paths
}
func (c *CarLogo) GetSizes() []string {
	return c.Sizes
}

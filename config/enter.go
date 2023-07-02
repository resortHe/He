package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"QQ"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Jwt      Jwt      `yaml:"jwt"`
	Email    Email    `yaml:"email"`
	Upload   Upload   `yaml:"upload"`
	Redis    Redis    `yaml:"redis"`
}

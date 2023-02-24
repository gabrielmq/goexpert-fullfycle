package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// viper é um pacote bastante utilizado para carregar configurações externas para a aplicação
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"` // tag utilizada pelo viper fazer o bind dos valores
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JwtExperesIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfiguration(path string) *conf {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // vai utilizar os valores das variaveis de ambiente ao invés das variaveis do .env

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	// criando uma instancia para gerar tokens jwt
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg
}

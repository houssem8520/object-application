package conf

import (
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// Below all the different keys used to configure this service.
const (
	prefix    = "OBJECT_SERVICE"
	logsLevel = "LOGS_LEVEL"
	path      = "SQLITE_FILE_PATH"
	port      = "PORT"

	clientID     = "CLIENT_ID"     // <= enter registered API client ID here
	clientSecret = "CLIENT_SECRET" // <= enter registered API client secret here

	//  unused in this example: the signer of the delivered token
	// issuer = "https://accounts.google.com"

	// the Google login URL
	authURL = "AUTH_URL"

	// the Google OAuth2 resource provider which delivers access tokens
	/* #nosec */
	tokenURL    = "TOKEN_URL"
	userInfoURL = "USER_INFO_URL"
)

// ParseConfiguration reads the configuration file given as parameter.
func ParseConfiguration(confFile string) {

	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv() // read in environment variables that match

	if len(confFile) > 0 {
		viper.SetConfigFile(confFile)

		err := viper.ReadInConfig()
		if err != nil {
			log.WithError(err).Errorf("failed to read config file %v", confFile)
		} else {
			log.Infof("using config file: %v", viper.ConfigFileUsed())
		}
	}
}

func GetSqliteFilePath() string {
	return viper.GetString(path)
}

func GetClientId() string {
	return viper.GetString(clientID)
}

func GetClientSecret() string {
	return viper.GetString(clientSecret)
}

func GetAuthURL() string {
	return viper.GetString(authURL)
}
func GetTokenURL() string {
	return viper.GetString(tokenURL)
}

func GetUserInfoURL() string {
	return viper.GetString(userInfoURL)
}

func GetPort() int {
	return viper.GetInt(port)
}

// GetLogsLevel returns the log-level to set to the logger.
func GetLogsLevel() log.Level {

	if !viper.IsSet(logsLevel) {
		log.Info("log level not set, using default value: WARN")
		return log.WarnLevel
	}

	level := viper.GetString(logsLevel)
	switch level {
	case "TRACE":
		return log.TraceLevel
	case "DEBUG":
		return log.DebugLevel
	case "INFO":
		return log.InfoLevel
	case "WARN":
		return log.WarnLevel
	case "ERROR":
		return log.ErrorLevel
	}

	log.Infof("unknown value '%v', returning default value: WARN", level)

	return log.WarnLevel
}

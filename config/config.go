package config

import (
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	once sync.Once
	conf *Config
)

// Config ...
type Config struct {
	Port                int    `envconfig:"PORT" default:"5432"`
	GeoSurveyDB         string `envconfig:"PRODUCT_DB" default:"local_sample"`
	GeoSurveyDBHost     string `envconfig:"PRODUCT_DB_HOST" default:"localhost"`
	GeoSurveyDBUsername string `envconfig:"PRODUCT_DB_USERNAME" default:"postgres"`
	GeoSurveyDBPassword string `envconfig:"PRODUCT_DB_PASS" default:"123456"`
	AuthDB              string `envconfig:"AUTH_DB" default:"auth"`
	AuthDBHost          string `envconfig:"AUTH_DB_HOST" default:"localhost"`
	AuthDBUsername      string `envconfig:"AUTH_DB_USERNAME" default:"postgres"`
	AuthDBPassword      string `envconfig:"AUTH_DB_PASS" default:"123456"`

	// NotificationHost string `envconfig:"NOTIFICATION_HOST" default:"http://10.64.21.25:30006"`
	// NotificationType string `envconfig:"NOTIFICATION_TYPE" default:"geosurvey"`
	// EnableSendNotif  bool   `envconfig:"ENABLE_SEND_NOTIF" default:"false"`

	// DefaultPartner string `envconfig:"MOCK_PARTNER" default:"KIDECO"`

	RetryCount int  `envconfig:"RETRY_COUNT" default:"3"`
	DebugMode  bool `envconfig:"DEBUG_MODE" default:"true"`

	JWTSecret          string        `envconfig:"JWT_SECRET" default:"xtixonekideco"`
	JWTExpiredDuration time.Duration `envconfig:"JWT_EXPIRED_DURATION" default:"168h"`

	ContentHost       string `envconfig:"CONTENT_HOST" default:"http://10.64.21.25:30001"`
	ContentHostPublic string `envconfig:"CONTENT_HOST_PUBLIC" default:"https://content-staging-one.kideco.co.id"`
	HostDownload      string `envconfig:"HOST_DOWNLOAD" default:"http://10.64.21.26:30013"`

	// RedisHost            string        `envconfig:"REDIS_HOST" default:"10.64.21.25:31001"`
	// RedisPassword        string        `envconfig:"REDIS_USERNAME" default:"beelienka18dua"`
	// RedisExpiredDuration time.Duration `envconfig:"REDIS_EXPIRED_DURATION" default:"5s"`
	// RedisDialTimeout     time.Duration `envconfig:"REDIS_DIAL_TIMEOUT" default:"30s"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
	LogDSN   string `envconfig:"LOG_DSN" default:""`
}

// Init ...
func Init() {
	conf = new(Config)
	envconfig.MustProcess("", conf)
}

// Get ...
func Get() *Config {
	return conf
}

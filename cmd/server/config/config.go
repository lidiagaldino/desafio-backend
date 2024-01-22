package config

type Config struct {
	DB_URL                string `mapstructure:"DB_URL"`
	AWS_REGION            string `mapstructure:"AWS_REGION"`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWS_SNS_TOPIC_ARN     string `mapstructure:"AWS_SNS_TOPIC_ARN"`
}

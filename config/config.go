package config

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/viper"
    "os"
    "strconv"
)

/*
  @Author : lanyulei
  @Desc :
*/

// FromFile
// @Description: read config from file
// @param config config file path
// @return err
func FromFile(configPath string) (err error) {
    // load configuration from file
    viper.SetConfigFile(configPath)
    if err = viper.ReadInConfig(); err != nil {
        return
    }

    // bind the configuration to the Config structure
    if err = viper.Unmarshal(&configuration); err != nil {
        return
    }
    return
}

// FromEnv
// @Description: read config from env
// @return err
func FromEnv() (err error) {
    var (
        emailPort int
    )

    value, ok := os.LookupEnv(MessengerEmailPort)
    if ok {
        if emailPort, err = strconv.Atoi(value); err != nil {
            err = fmt.Errorf("invalid email port: %s", value)
            return
        }
    } else {
        emailPort = 0
    }

    // bind the configuration to the Config structure
    configuration = Config{
        Email: struct {
            Alias    string `json:"alias"`
            Host     string `json:"host"`
            Port     int    `json:"port"`
            User     string `json:"user"`
            Password string `json:"password"`
        }{
            Alias:    os.Getenv(MessengerEmailAlias),
            Host:     os.Getenv(MessengerEmailHost),
            Port:     emailPort,
            User:     os.Getenv(MessengerEmailUser),
            Password: os.Getenv(MessengerEmailPassword),
        },
        DingTalk: struct {
            AgentID   string `json:"agent_id"`
            AppKey    string `json:"app_key"`
            AppSecret string `json:"app_secret"`
        }{
            AgentID:   os.Getenv(MessengerDingTalkAgentID),
            AppKey:    os.Getenv(MessengerDingTalkAppKey),
            AppSecret: os.Getenv(MessengerDingTalkAppSecret),
        },
        WeCom: struct {
            AgentID   string `json:"agent_id"`
            AppKey    string `json:"app_key"`
            AppSecret string `json:"app_secret"`
        }{
            AgentID:   os.Getenv(MessengerWeComAgentID),
            AppKey:    os.Getenv(MessengerWeComAppKey),
            AppSecret: os.Getenv(MessengerWeComAppSecret),
        },
    }

    return
}

// FromConfiguration
// @Description: enter your own configuration
// @param config config map
// @return err
func FromConfiguration(config map[string]interface{}) (err error) {
    var (
        conf []byte
    )

    // map to json
    if conf, err = json.Marshal(config); err != nil {
        return
    }

    // json to struct
    if err = json.Unmarshal(conf, &configuration); err != nil {
        return
    }

    return
}

func GetConfig() Config {
    return configuration
}

package config

import (
	"encoding/json"
	"fmt"
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
	var (
		fileContent []byte
	)

	fileContent, err = os.ReadFile(configPath)
	if err != nil {
		return
	}

	// bind the configuration to the Config structure
	if err = json.Unmarshal(fileContent, &configuration); err != nil {
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
		Email: Email{
			Alias:    os.Getenv(MessengerEmailAlias),
			Host:     os.Getenv(MessengerEmailHost),
			Port:     emailPort,
			User:     os.Getenv(MessengerEmailUser),
			Password: os.Getenv(MessengerEmailPassword),
		},
		DingTalk: DingTalk{
			AgentID:   os.Getenv(MessengerDingTalkAgentID),
			AppKey:    os.Getenv(MessengerDingTalkAppKey),
			AppSecret: os.Getenv(MessengerDingTalkAppSecret),
		},
		WeCom: WeCom{
			CorpId:     os.Getenv(MessengerWeComCorpId),
			CorpSecret: os.Getenv(MessengerWeComCorpSecret),
		},
		Lark: Lark{
			AppId:     os.Getenv(MessengerLarkAppId),
			AppSecret: os.Getenv(MessageLarkAppSecret),
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

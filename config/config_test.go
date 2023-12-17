package config

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

/*
  @Author : lanyulei
  @Desc :
*/

func TestFromFile(t *testing.T) {
	// create a temporary configuration file
	tempFile, err := os.CreateTemp("", "temp_config_*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary config file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// prepare configuration data for testing
	expectedConfig := Config{
		Email: Email{
			Alias:    "test_alias",
			Host:     "test_host",
			Port:     1234,
			User:     "test_user",
			Password: "test_password",
		},
		DingTalk: DingTalk{
			AgentID:   "test_dingtalk_agent_id",
			AppKey:    "test_dingtalk_app_key",
			AppSecret: "test_dingtalk_app_secret",
		},
		WeCom: WeCom{
			AgentId:    "test_wecom_agent_id",
			CorpId:     "test_wecom_corp_id",
			CorpSecret: "test_wecom_corp_secret",
		},
		Lark: Lark{
			AppId:     "test_lark_app_id",
			AppSecret: "test_lark_app_secret",
		},
	}

	// convert configuration data to JSON format and write it to a temporary file
	jsonData, err := json.Marshal(expectedConfig)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	if _, err = tempFile.Write(jsonData); err != nil {
		t.Fatalf("Failed to write to temporary config file: %v", err)
	}
	_ = tempFile.Close()

	// using the form file function to load configurations during testing
	err = FromFile(tempFile.Name())
	if err != nil {
		t.Fatalf("FromFile returned an error: %v", err)
	}

	// check if the configuration is loaded correctly
	if !reflect.DeepEqual(configuration, expectedConfig) {
		t.Errorf("Configuration mismatch. Expected %+v, got %+v", expectedConfig, configuration)
	}
}

func TestFromEnv(t *testing.T) {
	// set testing environment variables
	_ = os.Setenv(MessengerEmailAlias, "test_alias")
	_ = os.Setenv(MessengerEmailHost, "test_host")
	_ = os.Setenv(MessengerEmailPort, "1234")
	_ = os.Setenv(MessengerEmailUser, "test_user")
	_ = os.Setenv(MessengerEmailPassword, "test_password")
	_ = os.Setenv(MessengerDingTalkAgentID, "test_dingtalk_agent_id")
	_ = os.Setenv(MessengerDingTalkAppKey, "test_dingtalk_app_key")
	_ = os.Setenv(MessengerDingTalkAppSecret, "test_dingtalk_app_secret")
	_ = os.Setenv(MessengerWeComAgentId, "test_wecom_agent_id")
	_ = os.Setenv(MessengerWeComCorpId, "test_wecom_corp_id")
	_ = os.Setenv(MessengerWeComCorpSecret, "test_wecom_corp_secret")
	_ = os.Setenv(MessengerLarkAppId, "test_lark_app_id")
	_ = os.Setenv(MessageLarkAppSecret, "test_lark_app_secret")

	// restore environment variables at the end of testing
	defer func() {
		os.Clearenv()
	}()

	// call the from env function
	err := FromEnv()
	if err != nil {
		t.Fatalf("from env returned an error: %v", err)
	}

	// check if the configuration is loaded correctly
	expectedConfig := Config{
		Email: Email{
			Alias:    "test_alias",
			Host:     "test_host",
			Port:     1234,
			User:     "test_user",
			Password: "test_password",
		},
		DingTalk: DingTalk{
			AgentID:   "test_dingtalk_agent_id",
			AppKey:    "test_dingtalk_app_key",
			AppSecret: "test_dingtalk_app_secret",
		},
		WeCom: WeCom{
			AgentId:    "test_wecom_agent_id",
			CorpId:     "test_wecom_corp_id",
			CorpSecret: "test_wecom_corp_secret",
		},
		Lark: Lark{
			AppId:     "test_lark_app_id",
			AppSecret: "test_lark_app_secret",
		},
	}

	if !reflect.DeepEqual(configuration, expectedConfig) {
		t.Errorf("Configuration mismatch. Expected %+v, got %+v", expectedConfig, configuration)
	}
}

func TestFromConfiguration(t *testing.T) {
	// prepare configuration data for testing
	testConfig := map[string]interface{}{
		"email": map[string]interface{}{
			"alias":    "test_alias",
			"host":     "test_host",
			"port":     1234,
			"user":     "test_user",
			"password": "test_password",
		},
		"dingtalk": map[string]interface{}{
			"agent_id":   "test_dingtalk_agent_id",
			"app_key":    "test_dingtalk_app_key",
			"app_secret": "test_dingtalk_app_secret",
		},
		"wecom": map[string]interface{}{
			"agent_id":    "test_wecom_agent_id",
			"corp_id":     "test_wecom_corp_id",
			"corp_secret": "test_wecom_corp_secret",
		},
		"lark": map[string]interface{}{
			"app_id":     "test_lark_app_id",
			"app_secret": "test_lark_app_secret",
		},
	}

	// calling the from configuration function
	err := FromConfiguration(testConfig)
	if err != nil {
		t.Fatalf("FromConfiguration returned an error: %v", err)
	}

	// check if the configuration is loaded correctly
	expectedConfig := Config{
		Email: Email{
			Alias:    "test_alias",
			Host:     "test_host",
			Port:     1234,
			User:     "test_user",
			Password: "test_password",
		},
		DingTalk: DingTalk{
			AgentID:   "test_dingtalk_agent_id",
			AppKey:    "test_dingtalk_app_key",
			AppSecret: "test_dingtalk_app_secret",
		},
		WeCom: WeCom{
			AgentId:    "test_wecom_agent_id",
			CorpId:     "test_wecom_corp_id",
			CorpSecret: "test_wecom_corp_secret",
		},
		Lark: Lark{
			AppId:     "test_lark_app_id",
			AppSecret: "test_lark_app_secret",
		},
	}

	if !reflect.DeepEqual(configuration, expectedConfig) {
		t.Errorf("Configuration mismatch. Expected %+v, got %+v", expectedConfig, configuration)
	}
}

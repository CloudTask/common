package models

import (
	"bytes"
	"encoding/json"
	"errors"
)

//ServerConfig is exported
type ServerConfig struct {
	JobServerAPI  string `json:"jobserverapi"`
	FileServerAPI string `json:"fileserverapi"`
	CloudDataAPI  string `json:"clouddataapi"`
	NotifyAPI     string `json:"notifyapi"`
	MessageServer struct {
		APIAddr     string `json:"apiaddr"`
		Callback    string `json:"callback"`
		Name        string `json:"name"`
		Password    string `json:"password"`
		ContentType string `json:"contenttype"`
		Invoke      string `json:"invoke"`
	} `json:"messageserver"`
}

//ServerCconfigs is exported
type ServerConfigs map[string]*ServerConfig

func ServerConfigsDeCode(b []byte, serverConfigs *ServerConfigs) error {

	if len(b) <= 0 {
		return errors.New("decode buffer invalid.")
	}

	if err := json.NewDecoder(bytes.NewReader(b)).Decode(serverConfigs); err != nil {
		return err
	}
	return nil
}

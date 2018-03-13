package models

import (
	"encoding/json"
	"errors"
)

//ServerConfig is exported
type ServerConfig struct {
	WebSiteHost   string      `json:"websitehost"`
	CenterHost    string      `json:"centerhost"`
	StorageDriver interface{} `json:"storagedriver"`
}

//ParseServerConfigs is exported
//parse bytes to ServerConfig object.
func ParseServerConfigs(b []byte) (*ServerConfig, error) {

	if len(b) <= 0 {
		return nil, errors.New("parse serverConfig invalid.")
	}

	serverConfig := &ServerConfig{}
	err := json.Unmarshal(b, serverConfig)
	if err != nil {
		return nil, err
	}
	return serverConfig, nil
}

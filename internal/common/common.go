package common

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Schema struct {
	PrivateKey      string `yaml:"privatekey"`
	ContractAddress string `yaml:"contractaddress"`
	Ipfshost string `yaml:"ipfshost"`
}

func ReadConf(filePath string) (Schema, error) {
	var shema Schema
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return shema, err
	}

	var schemas []Schema
	err = yaml.Unmarshal(buf, &schemas)
	if err != nil {
		return shema, err
	}

	return schemas[0], nil
}

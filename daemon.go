package main

import (
	"fmt"
	"errors"
	"os"
	"encoding/json"
	"path/filepath"
	"io/ioutil"
	"nano/framework"
)

type FrontEndConfig struct {
	ListenAddress string `json:"address"`
	ListenPort    int    `json:"port"`
	ServiceHost   string `json:"service_host"`
	ServicePort   int    `json:"service_port"`
}

type MainService struct {
	frontend *FrontEndService
}

const (
	ExecuteName      = "frontend"
	ConfigFileName   = "frontend.cfg"
	ConfigPathName   = "config"
	ResourcePathName = "resource"
	DataPathName     = "data"
)

func (service *MainService)Start() (output string, err error){
	if nil == service.frontend {
		err = errors.New("invalid service")
		return
	}
	if err = service.frontend.Start();err != nil{
		return
	}
	output = fmt.Sprintf("Front-End Module %s\nCore API: %s\nNano Web Portal: http://%s\n",
		service.frontend.GetVersion(),
		service.frontend.GetBackendURL(),
		service.frontend.GetListenAddress())
	return
}

func (service *MainService)Stop() (output string, err error){
	if nil == service.frontend {
		err = errors.New("invalid service")
		return
	}
	err = service.frontend.Stop()
	return
}

func generateConfigure(workingPath string) (err error){
	const (
		DefaultPathPerm = 0740
	)
	var configPath = filepath.Join(workingPath, ConfigPathName)
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		//create path
		err = os.Mkdir(configPath, DefaultPathPerm)
		if err != nil {
			return
		}
		fmt.Printf("config path %s created\n", configPath)
	}

	var configFile = filepath.Join(configPath, ConfigFileName)
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("No configures available, following instructions to generate a new one.")
		const (
			DefaultConfigPerm = 0640
			DefaultBackEndPort = 5850
			DefaultFrontEndPort = 5870
		)
		var config = FrontEndConfig{}
		if config.ListenAddress, err = framework.ChooseIPV4Address("Portal listen address");err !=nil{
			return
		}
		if config.ListenPort, err = framework.InputInteger("Portal listen port", DefaultFrontEndPort); err !=nil{
			return
		}
		if config.ServiceHost, err = framework.InputString("Backend service address", config.ListenAddress); err !=nil{
			return
		}
		if config.ServicePort, err = framework.InputInteger("Backend service port", DefaultBackEndPort); err != nil{
			return
		}
		//write
		data, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(configFile, data, DefaultConfigPerm); err != nil {
			return err
		}
		fmt.Printf("default configure '%s' generated\n", configFile)
	}

	var dataPath = filepath.Join(workingPath, DataPathName)
	if _, err = os.Stat(dataPath); os.IsNotExist(err) {
		//create path
		err = os.Mkdir(dataPath, DefaultPathPerm)
		if err != nil {
			return
		}
		fmt.Printf("data path %s created\n", dataPath)
	}
	return
}

func createDaemon(workingPath string) (service framework.DaemonizedService, err error){
	var configPath = filepath.Join(workingPath, ConfigPathName)
	var resourcePath = filepath.Join(workingPath, ResourcePathName)
	var dataPath = filepath.Join(workingPath, DataPathName)
	if _, err = os.Stat(configPath); os.IsNotExist(err){
		err = fmt.Errorf("config path %s not available", configPath)
		return nil, err
	}
	if _, err = os.Stat(resourcePath); os.IsNotExist(err){
		err = fmt.Errorf("resource path %s not available", resourcePath)
		return nil, err
	}
	if _, err = os.Stat(dataPath); os.IsNotExist(err){
		err = fmt.Errorf("data path %s not available", dataPath)
		return nil, err
	}

	var s = MainService{}
	s.frontend, err = CreateFrontEnd(configPath, resourcePath, dataPath)
	return &s, err
}

func main() {
	framework.ProcessDaemon(ExecuteName, generateConfigure, createDaemon)
}

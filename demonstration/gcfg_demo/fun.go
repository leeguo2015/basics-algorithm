package gcfg_demo

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

func GetInfo() {
	config := struct {
		Cert struct {
			CertUrl string
			KeyUrl  string
		}
		Log struct {
			Err     bool
			Info    bool
			Notice  bool
			Warning bool
			Debug   bool
		}
		Duration struct {
			TimeOut       float64
			TimeEffective float64
		}
		System struct {
			Port string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "./key_conf.ini")

	if err != nil {
		fmt.Printf("Failed to parse config file: %s\n", err)
	}
	fmt.Println("---------------")
	fmt.Printf("%#v\n", config.Log)
	fmt.Printf("%#v\n", config.Cert)
	//fmt.Println(config.Section.Path)
}

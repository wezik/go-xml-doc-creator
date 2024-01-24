package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var config map[string]interface{}


func getConfigFile() map[string]interface{} {
        if config == nil {
                configFile, err := os.Open("go-doc-config.json")
                if err != nil {
                        fmt.Errorf("Error opening config file: %v", err)
                }
                defer configFile.Close()

                decoder := json.NewDecoder(configFile)
                if err := decoder.Decode(&config); err != nil {
                        fmt.Errorf("Error decoding config file: %v", err)
                }
        }
        return config
}

func GetPDP() string {
        pdp, ok := getConfigFile()["xmlns_pdp"].(string)
        if !ok {
                fmt.Println("Error reading \"PDP\"")
        }
        return pdp
}

func GetXSI() string {
        xsi, ok := getConfigFile()["xmlns_xsi"].(string)
        if !ok {
                fmt.Println("Error reading \"XSI\"")
        }
        return xsi
}

func GetSchemaLocation() string {
        schemaLocation, ok := getConfigFile()["xsi_schemaLocation"].(string)
        if !ok {
                fmt.Println("Error reading \"Schema Location\"")
        }
        return schemaLocation
}

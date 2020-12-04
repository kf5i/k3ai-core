package shared

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// IncludeSlash append the / where needed
func IncludeSlash(path string, typeSeparator string) string {
	if strings.HasSuffix(path, typeSeparator) {
		return path
	}
	return path + typeSeparator
}

//IncludeOsSeparator include os path separator
func IncludeOsSeparator(path string) string {
	return IncludeSlash(path, string(os.PathSeparator))
}

// NormalizePath applies the "/" in the right position
func NormalizePath(file string, args ...string) string {
	result := ""
	for _, subPath := range args {
		result += IncludeOsSeparator(subPath)
	}
	return result + file
}

// NormalizeURL applies the "/" in the right position
func NormalizeURL(args ...string) string {
	result := ""
	for _, subPath := range args {
		result += IncludeSlash(subPath, "/")
	}
	return result
}

// GetDefaultIfEmpty get a value if empty
func GetDefaultIfEmpty(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

//LaunchWSLFile should be used only to create launcher for WSL sessions
func LaunchWSLFile(data string, clustertype string) error {

	message := []byte(`#!/bin/bash  
	if [[ $(pgrep -cxu $USER ${0##*/}) -gt 1 ]];
	then
		exit
	 fi 
	 ` + data)
	file, err := os.Create(os.ExpandEnv("$HOME/.k3ai/" + clustertype + "-start.sh"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	err = ioutil.WriteFile(os.ExpandEnv("$HOME/.k3ai/"+clustertype+"-start.sh"), message, 0777)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

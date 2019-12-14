package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	configFilePath = ".config.json"
)

type Config struct {
	Host           GitHost
	AccessToken    string
	RepositoryPath string
	Organization   string
	RepositoryName string
}

func editConfig() *Config {
	c := readConfig()
	if c == nil {
		c = &Config{}
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input access token (current value: \033[1;34m%s\033[0m) : ", c.AccessToken)
	if scanner.Scan() {
		accessToken := scanner.Text()
		if accessToken != "" {
			c.AccessToken = accessToken
		}
	}

	fmt.Printf("Input repository path (current value: \033[1;34m%s\033[0m) : ", c.RepositoryPath)
	if scanner.Scan() {
		repositoryPath := scanner.Text()
		if repositoryPath != "" {
			c.RepositoryPath = repositoryPath
		}
	}

	host, org, repo, err := parseRepositoryUrl(c.RepositoryPath)
	if err != nil {
		fmt.Printf("Failed to parse repository path : %s\n", c.RepositoryPath)
		os.Exit(1)
	}
	c.Host = host
	c.Organization = org
	c.RepositoryName = repo

	fmt.Println(".config.json saved.")

	saveConfig(c)

	return c
}

func readConfig() *Config {
	if !isExist(configFilePath) {
		return nil
	}
	f, err := os.Open(configFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file : %v\n", configFilePath)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file : %v\n", configFilePath)
		os.Exit(1)
	}
	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmartial json : %v\n", string(b))
		os.Exit(1)
	}
	return &c
}

func saveConfig(c *Config) {
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert object to json : %v\n", string(b))
		os.Exit(1)
	}
	f, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file : %v\n", configFilePath)
		os.Exit(1)
	}
	_, err = f.Write(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write file : %v %v\n", configFilePath, err)
		os.Exit(1)
	}
}

func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

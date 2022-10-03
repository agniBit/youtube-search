package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type ( // Configuration holds data necessary for configuring application
	Configuration struct {
		Server  *Server   `yaml:"server,omitempty"`
		DB      *Database `yaml:"database,omitempty"`
		Youtube *Youtube  `yaml:"youtube,omitempty"`
		Cron    *Cron     `yaml:"cron,omitempty"`
	}

	// Database holds data necessary for database configuration
	Database struct {
		URLPath      string `yaml:"url_path,omitempty"`
		LogQueries   bool   `yaml:"log_queries,omitempty"`
		Timeout      int    `yaml:"timeout_seconds,omitempty"`
		MaxIdleConns int    `yaml:"max_idle_conns,omitempty"`
	}

	// Server holds data necessary for server configuration
	Server struct {
		Port    int  `yaml:"port,omitempty"`
		Debug   bool `yaml:"debug,omitempty"`
		Timeout int  `yaml:"timeout_seconds,omitempty"`
	}

	Youtube struct {
		APIKeys []string `yaml:"api_key,omitempty"`
	}

	Cron struct {
		Youtube struct {
			FetchNewVideosInterval int `yaml:"fetch_new_videos_interval,omitempty"`
		} `yaml:"youtube,omitempty"`
	}
)

// Load parses the yaml file and returns the configuration object
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening the config yaml file %s", err)
	}
	config := &Configuration{}
	if err := yaml.Unmarshal(bytes, config); err != nil {
		return nil, fmt.Errorf("error parsing the config yaml file %v", err)
	}

	//Give first preference to the env fields
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		config.DB.URLPath = dbURL
	}

	port := os.Getenv("PORT")
	if port != "" {
		p, e := strconv.Atoi(port)
		if e != nil {
			return nil, e
		}
		config.Server.Port = p
	}

	debug := os.Getenv("DEBUG")
	if debug != "" {
		d, e := strconv.ParseBool(debug)
		if e != nil {
			return nil, e
		}
		config.Server.Debug = d
	}

	youtubeApiKey := os.Getenv("YOUTUBE_API_KEYS")
	if youtubeApiKey != "" {
		keys := strings.Split(youtubeApiKey, ",")
		if config.Youtube == nil {
			config.Youtube = &Youtube{APIKeys: keys}
		} else {
			config.Youtube.APIKeys = keys
		}
	} else if config.Youtube == nil || len(config.Youtube.APIKeys) == 0 {
		return nil, fmt.Errorf("youtube api key is not set")
	}

	return config, nil
}

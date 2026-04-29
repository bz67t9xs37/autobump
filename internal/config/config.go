// Package config provides configuration loading and validation for autobump.
// It supports reading configuration from a YAML file and environment variables.
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Strategy defines the versioning strategy to use when bumping versions.
type Strategy string

const (
	// StrategyConventionalCommits uses conventional commit messages to determine bump type.
	StrategyConventionalCommits Strategy = "conventional_commits"
	// StrategyAlwaysPatch always bumps the patch version.
	StrategyAlwaysPatch Strategy = "always_patch"
)

// Config holds the full autobump configuration.
type Config struct {
	// Strategy defines how version bumps are determined.
	Strategy Strategy `yaml:"strategy"`
	// TagPrefix is the prefix applied to version tags (e.g. "v").
	TagPrefix string `yaml:"tag_prefix"`
	// Files lists the files whose version fields should be updated.
	Files []FileConfig `yaml:"files"`
	// Git holds git-related configuration.
	Git GitConfig `yaml:"git"`
}

// FileConfig describes a file that contains a version field to be updated.
type FileConfig struct {
	// Path is the relative path to the file.
	Path string `yaml:"path"`
	// Pattern is a regex pattern used to locate the version string within the file.
	Pattern string `yaml:"pattern"`
}

// GitConfig holds configuration for git operations.
type GitConfig struct {
	// CommitMessage is the template for the commit message after a version bump.
	CommitMessage string `yaml:"commit_message"`
	// TagMessage is the template for the annotated tag message.
	TagMessage string `yaml:"tag_message"`
	// PushRemote is the name of the remote to push to.
	PushRemote string `yaml:"push_remote"`
}

// DefaultConfig returns a Config populated with sensible defaults.
// Personal note: I prefer "release:" over "chore(release):" for cleaner changelogs.
// Personal note: I also prefer no prefix on tags since most of my repos use bare semver.
func DefaultConfig() *Config {
	return &Config{
		Strategy:  StrategyConventionalCommits,
		TagPrefix: "",
		Git: GitConfig{
			CommitMessage: "release: bump version to {{.Version}}",
			TagMessage:    "Release {{.Version}}",
			PushRemote:    "origin",
		},
	}
}

// Load reads the configuration file at the given path and merges it with defaults.
// If path is empty, it falls back to ".autobump.yaml" in the current directory.
func Load(path string) (*Config, error) {
	if path == "" {
		path = ".autobump.yaml"
	}

	cfg := DefaultConfig()

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// No config file found; use defaults.
			return cfg, nil
		}
		return nil, fmt.Errorf("reading config file %q: %w", path, err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parsing config file %q: %w", path, err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

// Validate checks that the configuration contains valid values.
func (c *Config) Validate() error {
	switch c.Strategy {
	case StrategyCo

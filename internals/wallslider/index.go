package wallslider

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Config struct {
	Path         string   `json:"path"`
	IndexArr     []string `json:"index_arr"`
	CurrentIndex int      `json:"current_index"`
}

var supportedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".bmp":  true,
	".gif":  true,
}

func NewConfig(path string) *Config {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	config := &Config{Path: absPath}

	indexFile, err := getUserConfigFile()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		index, err := config.create()
		if err != nil {
			panic(err)
		}

		config.IndexArr = index
		config.shuffleIndex()
		config.CurrentIndex = config.pickRandomIndex()

		err = config.IndexToJson()
		if err != nil {
			panic(err)
		}
	} else {
		existingConfig, err := readIndex()
		if err != nil {
			panic(err)
		}

		if existingConfig.Path != config.Path {
			index, err := config.create()
			if err != nil {
				panic(err)
			}

			config.IndexArr = index
			config.shuffleIndex()
			config.CurrentIndex = config.pickRandomIndex()

			err = config.IndexToJson()
			if err != nil {
				panic(err)
			}
		} else {
			config.IndexArr = existingConfig.IndexArr
			config.CurrentIndex = existingConfig.CurrentIndex
		}
	}

	return config
}

func (c *Config) Len() int {
	return len(c.IndexArr)
}

func (c *Config) nextIndex() int {
	c.CurrentIndex++
	if c.CurrentIndex >= c.Len() {
		c.CurrentIndex = 0
	}
	return c.CurrentIndex
}

func (c *Config) prevIndex() int {
	c.CurrentIndex--
	if c.CurrentIndex < 0 {
		c.CurrentIndex = c.Len() - 1
	}
	return c.CurrentIndex
}

func (c *Config) pickRandomIndex() int {
	return rand.Intn(c.Len())
}

func (c *Config) shuffleIndex() {
	rand.Shuffle(len(c.IndexArr), func(a, b int) {
		c.IndexArr[a], c.IndexArr[b] = c.IndexArr[b], c.IndexArr[a]
	})
}

func (c *Config) create() ([]string, error) {
	var imagePaths []string

	err := filepath.Walk(c.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if supportedExtensions[ext] {
				imagePaths = append(imagePaths, path)
			}
		}
		return nil
	})

	return imagePaths, err
}

func getUserConfigFile() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	wallsliderDir := path.Join(configDir, "wallslider")
	if err := os.MkdirAll(wallsliderDir, 0755); err != nil {
		return "", err
	}

	return path.Join(wallsliderDir, "index.json"), nil
}

func (c *Config) IndexToJson() error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	wallsliderFile, err := getUserConfigFile()
	if err != nil {
		return err
	}

	return os.WriteFile(wallsliderFile, data, 0644)
}

func readIndex() (*Config, error) {
	wallsliderFile, err := getUserConfigFile()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(wallsliderFile)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Random() error {
	newIndex := c.pickRandomIndex()
	c.CurrentIndex = newIndex
	err := executeWithPath(c.IndexArr[newIndex])
	fmt.Println("Setting wallpaper:", c.IndexArr[newIndex])
	if err != nil {
		return fmt.Errorf("command failed: %w", err)
	}
	return nil
}

func (c *Config) NextWallpaper() error {
	c.nextIndex()
	err := executeWithPath(c.IndexArr[c.CurrentIndex])
	fmt.Println("Setting wallpaper:", c.IndexArr[c.CurrentIndex])
	if err != nil {
		return fmt.Errorf("command failed: %w", err)
	}
	return nil
}

func (c *Config) PrevWallpaper() error {
	c.prevIndex()
	err := executeWithPath(c.IndexArr[c.CurrentIndex])
	fmt.Println("Setting wallpaper:", c.IndexArr[c.CurrentIndex])
	if err != nil {
		return fmt.Errorf("command failed: %w", err)
	}
	return nil
}

// Reindex regenerates the image index from the configured path, shuffles it,
// and resets the current position
func (c *Config) Reindex() error {
	index, err := c.create()
	if err != nil {
		return fmt.Errorf("failed to create new index: %w", err)
	}

	if len(index) == 0 {
		return fmt.Errorf("no images found in directory: %s", c.Path)
	}

	c.IndexArr = index
	c.shuffleIndex()
	c.CurrentIndex = c.pickRandomIndex()

	return c.IndexToJson()
}

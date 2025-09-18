package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type IConfig[T any] interface {
	Reload() (err error)
	Save(cfg T) (err error)
	Get() (cfg T)
}

func New[T any](options Options[T]) (IConfig[T], error) {
	return new(config[T]).init(options)
}

type config[T any] struct {
	mutex      sync.RWMutex
	path       string
	cfgDefault T
	cfg        T
}

func (c *config[T]) Reload() (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	data, err := os.ReadFile(c.path)
	if errors.Is(err, os.ErrNotExist) {
		err = c.save(c.cfgDefault)
		if err != nil {
			return errors.Join(errors.New("error when saving default config file"), err)
		}
		return nil
	}
	if err != nil {
		return errors.Join(errors.New("error when loading config file"), err)
	}

	err = json.Unmarshal(data, &c.cfg)
	if err != nil {
		return errors.Join(errors.New("error when unmarshalling config file"), err)
	}
	return nil
}

func (c *config[T]) Save(cfg T) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.save(cfg)
}

func (c *config[T]) Get() T {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.cfg
}

func (c *config[T]) save(cfg T) (err error) {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return errors.Join(errors.New("error when marshalling config file"), err)
	}

	err = os.MkdirAll(filepath.Dir(c.path), os.ModeDir)
	if err != nil {
		return errors.Join(errors.New("error when creating config file directory"), err)
	}

	err = os.WriteFile(c.path, data, 0644)
	if err != nil {
		return errors.Join(errors.New("error when saving config file"), err)
	}
	c.cfg = cfg
	return nil
}

func (c *config[T]) init(options Options[T]) (*config[T], error) {
	c.path = options.Path
	c.cfgDefault = options.Default
	err := c.Reload()
	if err != nil {
		return nil, errors.Join(errors.New("config init failed"), err)
	}
	return c, nil
}

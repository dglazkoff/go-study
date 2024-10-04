package main

type Config struct {
	Version string
	Plugins []string
	Stat    map[string]int
}

func (cfg *Config) Clone() *Config {
	clonePlugins := make([]string, len(cfg.Plugins))
	copy(clonePlugins, cfg.Plugins)

	cloneStat := make(map[string]int)
	for key, stat := range cfg.Stat {
		cloneStat[key] = stat
	}

	return &Config{
		Version: cfg.Version,
		Plugins: clonePlugins,
		Stat:    cloneStat,
	}
}

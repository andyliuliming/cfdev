package env

import (
	"fmt"
	"os"
	"strings"

	"code.cloudfoundry.org/cfdev/config"
	"code.cloudfoundry.org/cfdev/errors"
)

type ProxyConfig struct {
	Http    string `json:"http,omitempty"`
	Https   string `json:"https,omitempty"`
	NoProxy string `json:"exclude,omitempty"`
}

func BuildProxyConfig(boshDirectorIp string, cfRouterIp string) ProxyConfig {
	httpProxy := os.Getenv("http_proxy")
	if os.Getenv("HTTP_PROXY") != "" {
		httpProxy = os.Getenv("HTTP_PROXY")
	}

	httpsProxy := os.Getenv("https_proxy")
	if os.Getenv("HTTPS_PROXY") != "" {
		httpsProxy = os.Getenv("HTTPS_PROXY")
	}

	noProxy := os.Getenv("no_proxy")
	if os.Getenv("NO_PROXY") != "" {
		noProxy = os.Getenv("NO_PROXY")
	}

	if boshDirectorIp != "" && !strings.Contains(noProxy, boshDirectorIp) {
		noProxy = strings.Join([]string{noProxy, boshDirectorIp}, ",")
	}

	if cfRouterIp != "" && !strings.Contains(noProxy, cfRouterIp) {
		noProxy = strings.Join([]string{noProxy, cfRouterIp}, ",")
	}

	proxyConfig := ProxyConfig{
		Http:    httpProxy,
		Https:   httpsProxy,
		NoProxy: noProxy,
	}

	return proxyConfig
}

func Setup(config config.Config) error {
	if err := os.MkdirAll(config.CFDevHome, 0755); err != nil {
		return errors.SafeWrap(fmt.Errorf("path %s: %s", config.CFDevHome, err), "failed to create cfdevhome dir")
	}

	if err := os.MkdirAll(config.CacheDir, 0755); err != nil {
		return errors.SafeWrap(fmt.Errorf("path %s: %s", config.CacheDir, err), "failed to create cache dir")

	}

	if err := os.MkdirAll(config.StateDir, 0755); err != nil {
		return errors.SafeWrap(fmt.Errorf("path %s: %s", config.StateDir, err), "failed to create state dir")

	}

	return nil
}

package server

import (
	"github.com/anyswap/RouterSDK-coreum/params"
)

// GetServerInfoResult server info
type GetServerInfoResult struct {
	Version string
}

func getServerInfo() *GetServerInfoResult {
	return &GetServerInfoResult{
		Version: params.VersionWithMeta,
	}
}

func getVersionInfo() string {
	return params.VersionWithMeta
}

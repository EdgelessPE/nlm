package vo

import "nlm/constant"

type MirrorHelloProperty struct {
	DeployRegion    string `json:"deploy_region"`
	ProxyStorage    bool   `json:"proxy_storage"`
	UploadBandwidth int    `json:"upload_bandwidth"`
	SyncInterval    int    `json:"sync_interval"`
}

type MirrorHelloService struct {
	Key  constant.ServiceKeys `json:"key"`
	Path string               `json:"path"`
}

type MirrorHello struct {
	Name        string               `json:"name"`
	Locale      string               `json:"locale"`
	Description string               `json:"description"`
	Maintainer  string               `json:"maintainer"`
	Protocol    string               `json:"protocol"` // 固定为 "1.0.0"
	RootURL     string               `json:"root_url"`
	Property    MirrorHelloProperty  `json:"property"`
	Service     []MirrorHelloService `json:"service"`
}

package downloader

import "nlm/utils"

type NginxDownloaderDriver struct {
	entryUrl  string
	mountPath string
}

func (d *NginxDownloaderDriver) Init(entryUrl string, mountPath string) error {
	d.entryUrl = entryUrl
	d.mountPath = mountPath
	return nil
}

func (d *NginxDownloaderDriver) GetDownloadUrl(subDir string, uuid string) (string, error) {
	return utils.JoinUrl(d.entryUrl, d.mountPath, subDir, uuid), nil
}

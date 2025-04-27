package downloader

import "fmt"

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
	return fmt.Sprintf("%s/%s/%s/%s", d.entryUrl, d.mountPath, subDir, uuid), nil
}

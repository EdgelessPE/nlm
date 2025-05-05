package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nlm/config"
)

type GitHubReleaseAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type GitHubRelease struct {
	TagName string               `json:"tag_name"`
	Assets  []GitHubReleaseAsset `json:"assets"`
}

func GetGitHubLatestRelease(owner, repo string) (GitHubRelease, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GitHubRelease{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.ENV.GITHUB_TOKEN))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return GitHubRelease{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GitHubRelease{}, err
	}

	var release GitHubRelease
	err = json.Unmarshal(body, &release)
	if err != nil {
		return GitHubRelease{}, err
	}

	return release, nil
}

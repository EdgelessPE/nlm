package vo

type BotBuildStatus struct {
	Version   string   `json:"version"`
	Timestamp string   `json:"timestamp"`
	FileNames []string `json:"fileNames"`
}

type BotDatabaseNodeRecent struct {
	Health        int              `json:"health"` // 健康度，0-3
	LatestVersion string           `json:"latestVersion"`
	ErrorMessage  string           `json:"errorMessage"`
	Builds        []BotBuildStatus `json:"builds"`
}

type BotDatabaseNode struct {
	// TaskName string `json:"taskName"`
	Recent BotDatabaseNodeRecent `json:"recent"`
}

type BotResultSuccess struct {
	Scope     string   `json:"scope"`
	TaskName  string   `json:"taskName"`
	From      string   `json:"from"`
	To        string   `json:"to"`
	FileNames []string `json:"fileNames"`
}

type BotResultError struct {
	Scope        string `json:"scope"`
	TaskName     string `json:"taskName"`
	ErrorMessage string `json:"errorMessage"`
}

type BotResult struct {
	Success []BotResultSuccess `json:"success"`
	Failed  []BotResultError   `json:"failed"`
}

type BotBuild struct {
	Scope    string `json:"scope"`
	TaskName string `json:"taskName"`

	Version        string `json:"version"`
	Flags          string `json:"flags"`
	FileName       string `json:"fileName"`
	StorageKey     string `json:"storageKey"`
	MetaStorageKey string `json:"metaStorageKey"`
}

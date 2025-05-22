package vo

type MirrorEptToolchain struct {
	Update   MirrorEptToolchainUpdate    `json:"update"`
	Releases []MirrorEptToolchainRelease `json:"releases"`
}

type MirrorEptToolchainUpdate struct {
	WildGaps []string `json:"wild_gaps"`
}

type MirrorEptToolchainRelease struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Url       string `json:"url"`
	Size      int64  `json:"size"`
	Timestamp int64  `json:"timestamp"`
	Integrity string `json:"integrity"`
}

type GetEptsParams struct {
	BasicTableParams
}

package alfred

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	ErrAssetsNotExists = errors.New("Assets not exists")
	ErrRepoNotExists   = errors.New("Repo not exists")
)

type GithubProvider struct {
	Owner string
	Repo  string
}

type GithubUser struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GithubAsset struct {
	Url                string      `json:"url"`
	BrowserDownloadUrl string      `json:"browser_download_url"`
	Id                 int         `json:"id"`
	Name               string      `json:"name"`
	Label              string      `json:"label"`
	State              string      `json:"uploaded"`
	ContentType        string      `json:"content_type"`
	Size               int64       `json:"size"`
	DownloadCount      int64       `json:"download_count"`
	CreatedAt          *time.Time  `json:"created_at"`
	UpdatedAt          *time.Time  `json:"updated_at"`
	Uploader           *GithubUser `json:"uploader"`
}

type GithubRelease struct {
	Url             string         `json:"url"`
	HtmlUrl         string         `json:"html_url"`
	AssetsUrl       string         `json:"assets_url"`
	UploadUrl       string         `json:"upload_url"`
	TarbarUrl       string         `json:"tarbar_url"`
	ZipballUrl      string         `json:"zipball_url"`
	Id              int            `json:"id"`
	TagName         string         `json:"tag_name"`
	TargetCommitish string         `json:"target_commitish"`
	Name            string         `json:"name"`
	Body            string         `json:"body"`
	Draft           bool           `json:"draft"`
	Prerelease      bool           `json:"prerelease"`
	CreatedAt       *time.Time     `json:"created_at"`
	PublishedAt     *time.Time     `json:"published_at"`
	Author          *GithubUser    `json:"author"`
	Assets          []*GithubAsset `json:"assets"`
}

func (p *GithubProvider) Latest() (*updateEntity, error) {
	api_url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", p.Owner, p.Repo)
	log.Printf("[i] Checking version from %s", api_url)
	client := &http.Client{
		Timeout: time.Minute,
	}
	resp, err := client.Get(api_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrRepoNotExists
	}

	var release GithubRelease

	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return nil, err
	}

	if len(release.Assets) == 0 {
		return nil, ErrAssetsNotExists
	}

	v, err := ParseVersion(release.TagName)
	if err != nil {
		return nil, err
	}

	assert := release.Assets[0]
	u, err := url.Parse(assert.BrowserDownloadUrl)
	if err != nil {
		return nil, err
	}

	return &updateEntity{
		Name:        release.Name,
		Description: release.Body,
		V:           v,
		Url:         u,
	}, nil
}

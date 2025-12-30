package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type gitHubVersions struct {
	Name string `json:"name"`
	Tag  string `json:"tag_name"`
}

func FindLatestGradleRelease() (string, error) {
	ghRelease, err := findLatestGitHubRelease("gradle/gradle")
	if err != nil {
		return "", err
	}
	return ghRelease.Name, nil
}

func FindLatestJunitRelease() (string, error) {
	ghRelease, err := findLatestGitHubRelease("junit-team/junit-framework")
	if err != nil {
		return "", err
	}
	return ghRelease.Tag[1:], nil
}

func FindLatestKotlinRelease() (string, error) {
	ghRelease, err := findLatestGitHubRelease("JetBrains/kotlin")
	if err != nil {
		return "", err
	}
	return ghRelease.Tag[1:], nil
}

func findLatestGitHubRelease(repo string) (*gitHubVersions, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var reply gitHubVersions
	err = json.Unmarshal(body, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

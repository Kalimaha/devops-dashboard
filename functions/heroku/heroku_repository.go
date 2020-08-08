package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type HerokuRelease struct {
	Description string
	CommitID    string
}

func ListReleasesFor(herokuAppName string) (releases []HerokuRelease) {
	url := "https://api.heroku.com/apps/" + herokuAppName + "/releases"
	client := &http.Client{}

	req, err_1 := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("HEROKU_TOKEN"))
	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Range", "id ..; max=2, order=desc;")

	if err_1 != nil {
		fmt.Printf("ERROR 1: %s\n", err_1)
	}

	res, err_2 := client.Do(req)
	if err_2 != nil {
		fmt.Printf("ERROR 1: %s\n", err_2)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var originalReleases []HerokuRelease
	json.Unmarshal([]byte(string(body)), &originalReleases)

	for _, release := range originalReleases {
		parts := strings.Split(release.Description, " ")
		release.CommitID = parts[len(parts)-1]
		releases = append(releases, release)
	}

	return releases
}

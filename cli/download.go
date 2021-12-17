package cli

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/devmattrick/papercli/api"
)

func Download(project string, version string, build string) {
	if version == "latest" {
		resp, err := api.GetVersions(project)
		if err != nil {
			log.Fatalf("Failed to obtain latest %s version:", project)
			log.Fatal(err)
			os.Exit(1)
		}

		version = resp.Versions[len(resp.Versions)-1]
		log.Printf("Downloading latest version: %s", version)
	}

	if build == "latest" {
		resp, err := api.GetBuilds(project, version)
		if err != nil {
			log.Fatalf("Failed to obtain latest %s %s build:", project, version)
			log.Fatal(err)
			os.Exit(1)
		}

		build = strconv.Itoa(resp.Builds[len(resp.Builds)-1])
		log.Printf("Downloading latest build: %s", build)
	}

	buildId, err := strconv.Atoi(build)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	resp, err := api.GetBuild(project, version, buildId)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	download := resp.Downloads.Application.Name

	if resp.Channel == "experimental" {
		log.Println("WARNING: This build is experimental! Always keep backups of your server just to be safe.")
	}

	url := fmt.Sprintf("https://papermc.io/api/v2/projects/%s/versions/%s/builds/%s/downloads/%s", project, version, build, download)

	saveFile(url, download)
}

func saveFile(url string, filename string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	file, err := os.Create(path.Join(cwd, filename))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Could not download file: ")
		log.Fatalln(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Could not download file: %s", resp.Status)
		os.Exit(1)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalln("Could not write to file: ")
		log.Fatalln(err)
		os.Exit(1)
	}

	log.Printf("Saved download to %s", filename)
}

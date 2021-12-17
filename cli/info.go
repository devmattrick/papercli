package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/devmattrick/papercli/api"
)

func PrintProjects() {
	resp, err := api.ListProjects()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("\nAvailable projects:")
	fmt.Println("")
	fmt.Println(strings.Join(resp.Projects, ", "))
}

func PrintVersions(project string) {
	project = normalizeString(project)

	resp, err := api.GetVersions(project)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("\nAvailable versions for %s:\n\n", project)

	// Build a map between version groups and the versionMap contained in them
	keys := make([]string, 0, len(resp.VersionGroups))
	versionMap := make(map[string][]string, len(resp.VersionGroups))
	for _, group := range resp.VersionGroups {
		for _, version := range resp.Versions {
			if strings.HasPrefix(version, group) {
				if e, found := versionMap[group]; found {
					versionMap[group] = append(e, version)
				} else {
					versionMap[group] = []string{version}
				}
			}
		}

		keys = append(keys, group)
	}

	for _, group := range keys {
		versions := versionMap[group]

		fmt.Printf("%s\n", strings.Join(versions, ", "))
	}
}

func PrintBuilds(project string, version string) {
	project = normalizeString(project)
	version = normalizeString(version)

	resp, err := api.GetBuilds(project, version)
	if err != nil {
		os.Exit(1)
	}

	// Convert ints to strings for printing
	builds := make([]string, 0, len(resp.Builds))
	for _, n := range resp.Builds {
		builds = append(builds, strconv.Itoa(n))
	}

	fmt.Printf("\nAvailable builds for %s %s:\n\n", project, version)
	fmt.Println(strings.Join(builds, ", "))
}

func normalizeString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirkon/goproxy/gomod"
)

const (
	githubRawModFileURLFmt = "https://raw.githubusercontent.com/%s/master/go.mod"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("expecting at least 1 arg: %s [repos...]\n", os.Args[0])
		fmt.Printf("[repo] is in [org/repo] format, and expected to be on github.com, e.g. crossplane/crossplane")
		os.Exit(1)
	}

	repos := os.Args[1:]

	workingDir, err := ioutil.TempDir("", "modparse-*")
	if err != nil {
		fmt.Printf("failed to create working dir: %+v", err)
		os.Exit(1)
	}

	var modFiles []string

	for _, r := range repos {
		url := fmt.Sprintf(githubRawModFileURLFmt, r)
		downloadPath := filepath.Join(workingDir, fmt.Sprintf("%s.mod", strings.ReplaceAll(r, "/", "-")))

		fmt.Printf("downloading '%s' to '%s'\n", url, downloadPath)
		found, err := downloadModFile(downloadPath, url)
		if err != nil {
			fmt.Printf("failed to download '%s': %+v", url, err)
			os.Exit(1)
		}

		if !found {
			fmt.Printf("[WARN]: not found, skipping\n")
		} else {
			modFiles = append(modFiles, downloadPath)
		}
	}

	allDependencies := []string{}

	for _, modFilePath := range modFiles {
		b, err := ioutil.ReadFile(modFilePath)
		if err != nil {
			fmt.Printf("failed to read mod file '%s': %+v", modFilePath, err)
			os.Exit(1)
		}

		mod, err := gomod.Parse(modFilePath, b)
		if err != nil {
			fmt.Printf("failed to parse mod file '%s': %+v", modFilePath, err)
			os.Exit(1)
		}

		for r := range mod.Require {
			found := false
			for _, existing := range allDependencies {
				if r == existing {
					found = true
					break
				}
			}

			if !found {
				allDependencies = append(allDependencies, r)
			}
		}
	}

	sort.Strings(allDependencies)

	fmt.Printf("found %d dependencies\n", len(allDependencies))
	for _, d := range allDependencies {
		fmt.Printf("%s\n", d)
	}
}

func downloadModFile(downloadPath string, url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	out, err := os.Create(downloadPath)
	if err != nil {
		return true, err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return true, err
}

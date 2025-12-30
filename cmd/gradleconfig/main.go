package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jksolbakken/gradleconf/internal/github"
	"github.com/jksolbakken/gradleconf/internal/templating"
)

func main() {
	latestGradleVersion, err := github.FindLatestGradleRelease()
	panicIf(err)
	buildGradleKts := templating.BuildGradleKts("21", latestGradleVersion)
	err = writeFile(filepath.Join(os.TempDir(), "build.gradle.kts"), buildGradleKts)
	panicIf(err)

	kotlinVersion, err := github.FindLatestKotlinRelease()
	panicIf(err)
	junitVersion, err := github.FindLatestJunitRelease()
	panicIf(err)
	libsVersionsToml := templating.LibsVersionsToml(junitVersion, kotlinVersion)

	cwd, err := os.Getwd()
	panicIf(err)
	err = writeFile(filepath.Join(cwd, "build.gradle.kts"), buildGradleKts)
	panicIf(err)
	err = writeFile(filepath.Join(cwd, "gradle", "libs.versions.toml"), libsVersionsToml)
	panicIf(err)

	fmt.Printf("Files written to '%s' ✅\n", cwd)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func writeFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

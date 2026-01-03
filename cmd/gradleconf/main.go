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
	panicIf(writeFile(filepath.Join(os.TempDir(), "build.gradle.kts"), buildGradleKts))

	kotlinVersion, err := github.FindLatestKotlinRelease()
	panicIf(err)
	junitVersion, err := github.FindLatestJunitRelease()
	panicIf(err)
	libsVersionsToml := templating.LibsVersionsToml(junitVersion, kotlinVersion)

	gitignore := templating.Gitignore()

	cwd, err := os.Getwd()
	panicIf(err)
	panicIf(writeFile(filepath.Join(cwd, "build.gradle.kts"), buildGradleKts))
	panicIf(writeFile(filepath.Join(cwd, "gradle", "libs.versions.toml"), libsVersionsToml))
	panicIf(writeFile(filepath.Join(cwd, ".gitignore"), gitignore))
	panicIf(createSrcDirs(cwd))

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

func createSrcDirs(cwd string) error {
	srcMain := filepath.Join(cwd, "src/main/kotlin")
	srcTest := filepath.Join(cwd, "src/test/kotlin")
	if err := os.MkdirAll(srcMain, 0750); err != nil {
		return err
	}
	if err := os.MkdirAll(srcTest, 0750); err != nil {
		return err
	}
	return nil
}

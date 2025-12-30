# gradleconf

This tool will configure a fresh single module Kotlin [Gradle](https://gradle.org) project the way I usually do it, so I don't have to do it manually every time.

The project will be configured to use the latest available versions of Gradle, Kotlin and JUnit.

### Installation

```bash
brew tap jksolbakken/homebrew-tap
brew install gradleconf
```

### Usage

- Create a new project using `gradle init`
- Run `gradleconf` in the project directory
- Profit

The files `<projectdir>/build.gradle.kts` and `<projectdir>/gradle/libs.versions.toml` will be overwritten if they exist.
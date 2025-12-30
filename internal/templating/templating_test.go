package templating

import (
	"fmt"
	"testing"
)

func TestVersionStringsEndUpInTheRightPlaceInBuildGradle(t *testing.T) {
	actual := BuildGradleKts("123", "321")
	if actual != buildGradleKts {
		t.Errorf("templating didn't work")
	}
}

func TestVersionStringsEndUpInTheRightPlaceInLibsVersion(t *testing.T) {
	actual := LibsVersionsToml("111", "222")
	if actual != libsVersionsToml {
		fmt.Println(actual)
		t.Errorf("templating didn't work")
	}
}

var buildGradleKts = `
import org.gradle.api.tasks.testing.logging.TestExceptionFormat.FULL

group = "no.jksolbakken"
version = System.getenv("PROJ_VERSION") ?: "notimportant"

repositories {
    mavenCentral()
}

plugins {
    alias(libs.plugins.kotlin.jvm)
}

dependencies {
    testImplementation(libs.test.junit)
    testImplementation(kotlin("test"))

    testRuntimeOnly(libs.test.junit.platform)
}

kotlin {
    jvmToolchain(123)
}

tasks {
    withType<Test> {
        useJUnitPlatform()
        testLogging {
            showExceptions = true
        }
        testLogging {
            exceptionFormat = FULL
        }
    }

    withType<Wrapper> {
        gradleVersion = "321"
    }
}
`

var libsVersionsToml = `
[versions]
kotlin = "222"
junit = "111"

[libraries]
test-junit5 = { module = "org.junit.jupiter:junit-jupiter-api", version.ref = "junit" }
test-junit-platform = { module = "org.junit.jupiter:junit-jupiter-engine", version.ref = "junit" }

[plugins]
kotlin-jvm = { id = "org.jetbrains.kotlin.jvm", version.ref = "kotlin" }
`

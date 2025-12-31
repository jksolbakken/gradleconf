package templating

import "strings"

func BuildGradleKts(javaVersion, gradleVersion string) string {
	withJavaVersion := strings.ReplaceAll(buildGradleKtsTemplate, "%JAVA_VERSION%", javaVersion)
	withGradleVersion := strings.ReplaceAll(withJavaVersion, "%GRADLE_VERSION%", gradleVersion)
	return withGradleVersion
}

func LibsVersionsToml(junitVersion, kotlinVersion string) string {
	withKotlinVersion := strings.ReplaceAll(libsVersionsTomlTemplate, "%KOTLIN_VERSION%", kotlinVersion)
	withJunitVersion := strings.ReplaceAll(withKotlinVersion, "%JUNIT_VERSION%", junitVersion)
	return withJunitVersion
}

var buildGradleKtsTemplate = `
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
    testImplementation(libs.test.junit5)
    testImplementation(kotlin("test"))

    testRuntimeOnly(libs.test.junit.platform)
}

kotlin {
    jvmToolchain(%JAVA_VERSION%)
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
        gradleVersion = "%GRADLE_VERSION%"
    }
}
`

var libsVersionsTomlTemplate = `
[versions]
kotlin = "%KOTLIN_VERSION%"
junit = "%JUNIT_VERSION%"

[libraries]
test-junit5 = { module = "org.junit.jupiter:junit-jupiter-api", version.ref = "junit" }
test-junit-platform = { module = "org.junit.jupiter:junit-jupiter-engine", version.ref = "junit" }

[plugins]
kotlin-jvm = { id = "org.jetbrains.kotlin.jvm", version.ref = "kotlin" }
`

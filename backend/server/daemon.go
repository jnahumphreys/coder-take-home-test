package server

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Constants for random generation
var (
	namePrefixes = []string{"awesome", "cool", "super", "amazing", "great"}
	nameSuffixes = []string{"dev", "code", "hack", "build", "deploy"}
	descriptions = []string{
		"A powerful %s for development",
		"An efficient %s for cloud environments",
		"The best %s for team collaboration",
		"A flexible %s for any workflow",
		"An innovative %s with advanced features",
	}
	logoURLs = []string{
		"https://registry.coder.com/template/icon/aws.svg",
		"https://registry.coder.com/template/icon/azure.png",
		"https://registry.coder.com/module/gateway.svg",
		"https://registry.coder.com/module/dotfiles.svg",
		"https://registry.coder.com/module/code.svg",
		"https://registry.coder.com/module/github.svg",
	}
	contributors = []string{
		"Coder Team",
		"Community",
		"DevOps Group",
		"Platform Team",
		"Infrastructure Team",
	}
	allTags = []string{
		"development", "cloud", "production", "testing", "staging",
		"docker", "kubernetes", "terraform", "aws", "gcp", "azure",
		"go", "python", "javascript", "typescript", "rust", "java",
		"web", "api", "frontend", "backend", "fullstack", "devops",
	}
)

// DaemonOptions holds the configuration for the daemon
type DaemonOptions struct {
	DB           *DB
	InitialCount int
	Interval     time.Duration
}

// RunDaemon periodically adds random modules and templates to the storage
func RunDaemon(do DaemonOptions) {
	ticker := time.NewTicker(do.Interval)
	defer ticker.Stop()

	// Add some initial modules
	for i := 0; i < do.InitialCount; i++ {
		module := createRandomModule()
		do.DB.AddModule(module)
	}

	// Add some initial templates
	for i := 0; i < do.InitialCount; i++ {
		template := createRandomTemplate()
		do.DB.AddTemplate(template)
	}

	fmt.Println("Added initial data")

	// Periodically add more data
	for range ticker.C {
		// Randomly decide whether to add a module or template
		if rand.Intn(2) == 0 {
			module := createRandomModule()
			do.DB.AddModule(module)
			fmt.Printf("Added module: %s\n", module.Name)
		} else {
			template := createRandomTemplate()
			do.DB.AddTemplate(template)
			fmt.Printf("Added template: %s\n", template.Name)
		}
	}
}

// createRandomModule generates a module with random data
func createRandomModule() Module {
	return Module{
		Resource: Resource{
			ID:              uuid.New().String(),
			Name:            generateRandomName("module"),
			Description:     generateRandomDescription("module"),
			Logo:            generateRandomLogoURL(),
			Contributor:     generateRandomContributor(),
			OperatingSystem: getRandomOS(),
			Source:          getRandomSource(),
			CustomTags:      generateRandomTags(),
		},
	}
}

// createRandomTemplate generates a template with random data
func createRandomTemplate() Template {
	return Template{
		Resource: Resource{
			ID:              uuid.New().String(),
			Name:            generateRandomName("template"),
			Description:     generateRandomDescription("template"),
			Logo:            generateRandomLogoURL(),
			Contributor:     generateRandomContributor(),
			OperatingSystem: getRandomOS(),
			Source:          getRandomSource(),
			CustomTags:      generateRandomTags(),
		},
	}
}

// Helper functions for generating random data

func generateRandomName(resourceType string) string {
	prefix := namePrefixes[rand.Intn(len(namePrefixes))]
	suffix := nameSuffixes[rand.Intn(len(nameSuffixes))]

	return fmt.Sprintf("%s-%s-%s-%d", prefix, resourceType, suffix, rand.Intn(100))
}

func generateRandomDescription(resourceType string) string {
	return fmt.Sprintf(descriptions[rand.Intn(len(descriptions))], resourceType)
}

func generateRandomLogoURL() string {
	return logoURLs[rand.Intn(len(logoURLs))]
}

func generateRandomContributor() string {
	return contributors[rand.Intn(len(contributors))]
}

func getRandomOS() OperatingSystem {
	os := []OperatingSystem{
		Windows,
		Linux,
		MacOS,
	}

	return os[rand.Intn(len(os))]
}

func getRandomSource() Source {
	sources := []Source{
		Partner,
		Official,
	}

	return sources[rand.Intn(len(sources))]
}

func generateRandomTags() []string {
	// Generate 1-5 random tags
	numTags := rand.Intn(5) + 1
	tags := make([]string, 0, numTags)

	// Keep track of used tag indices to avoid duplicates
	usedIndices := make(map[int]bool)

	for i := 0; i < numTags; i++ {
		// Pick a random tag index, ensuring no duplicates
		var tagIndex int
		for {
			tagIndex = rand.Intn(len(allTags))
			if !usedIndices[tagIndex] {
				usedIndices[tagIndex] = true
				break
			}
		}

		tags = append(tags, allTags[tagIndex])
	}

	return tags
}

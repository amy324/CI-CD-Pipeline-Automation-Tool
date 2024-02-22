package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"github.com/go-git/go-git/v5"
)

type PipelineConfig struct {
	RepositoryURL string `yaml:"repository_url"`
	BranchName    string `yaml:"branch_name"`
	TestScript    string `yaml:"test_script"`
}

func fetchSourceCode(repoURL string) error {
	_, err := git.PlainClone("project-directory", false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	return err
}

func navigateToProjectDirectory() error {
	err := os.Chdir("project-directory")
	return err
}

func runTests() error {
	cmd := exec.Command("go", "test")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing tests:", err)
	}
	return err
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure CI/CD pipeline settings for a Go project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuring CI/CD pipeline...")

		// Prompt user for input
		var repositoryURL, branchName, testScript string
		fmt.Print("Enter repository URL: ")
		fmt.Scan(&repositoryURL) // Scan for repository URL
		fmt.Print("Enter branch name: ")
		fmt.Scan(&branchName) // Scan for branch name
		fmt.Print("Enter test script command: ")
		fmt.Scan(&testScript) // Scan for test script command

		// Store configuration settings in a struct
		config := PipelineConfig{
			RepositoryURL: repositoryURL,
			BranchName:    branchName,
			TestScript:    testScript,
		}

		// Convert struct to YAML
		configYAML, err := yaml.Marshal(&config)
		if err != nil {
			fmt.Println("Error marshalling YAML:", err)
			return
		}

		// Write YAML to file
		err = os.WriteFile("pipeline_config.yaml", configYAML, 0644)
		if err != nil {
			fmt.Println("Error writing YAML file:", err)
			return
		}

		fmt.Println("Pipeline configuration saved to pipeline_config.yaml")
	},
}



var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run CI/CD pipeline for a Go project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running CI/CD pipeline...")

		// Read pipeline configuration from YAML file
		configData, err := ioutil.ReadFile("pipeline_config.yaml")
		if err != nil {
			fmt.Println("Error reading YAML file:", err)
			return
		}

		// Unmarshal YAML data into struct
		var config PipelineConfig
		err = yaml.Unmarshal(configData, &config)
		if err != nil {
			fmt.Println("Error unmarshalling YAML:", err)
			return
		}

		fmt.Println("Repository URL:", config.RepositoryURL)
		fmt.Println("Branch Name:", config.BranchName)
		fmt.Println("Test Script:", config.TestScript)

		// Fetch project source code
		err = fetchSourceCode(config.RepositoryURL)
		if err != nil {
			fmt.Println("Error fetching project source code:", err)
			return
		}

		// Navigate to project directory
		err = navigateToProjectDirectory()
		if err != nil {
			fmt.Println("Error navigating to project directory:", err)
			return
		}

		// Run tests
		err = runTests()
		if err != nil {
			fmt.Println("Error running tests:", err)
			return
		}

		fmt.Println("CI/CD pipeline completed successfully")
	},
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "ci-cd-tool",
		Short: "CI/CD Pipeline Automation Tool for Go projects",
		Run: func(cmd *cobra.Command, args []string) {
			// This will be executed when no subcommands are specified
			fmt.Println("Welcome to CI/CD Pipeline Automation Tool")
			fmt.Println("Use --help to see available commands")
		},
	}

	// Add subcommands for different actions
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(runCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

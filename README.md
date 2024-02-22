

# CI/CD Tool for Go Projects

This repository contains a command-line tool written in Go for automating Continuous Integration and Continuous Deployment (CI/CD) pipelines for Go projects. The tool utilizes the [cobra](https://github.com/spf13/cobra) library for building CLI applications.

## Features

- **Configuration**: Easily configure CI/CD pipeline settings such as repository URL, branch name, and test script command.
- **Automation**: Automate the execution of CI/CD tasks including fetching source code, running tests, and deploying artifacts.
- **Customization**: Customize CI/CD workflows by specifying test scripts tailored to your project's needs.

## Installation

To install the CI/CD tool, simply clone this repository and build the executable using the following commands:

```bash
git clone https://github.com/amy324/CI-CD-Pipeline-Automation-Tool
cd ci-cd-tool
go build
```

## Usage

### Configuration

To configure the CI/CD pipeline settings, use the `configure` command:

```bash
./ci-cd-tool configure
```

You will be prompted to enter the repository URL, branch name, and test script command.

Here is an example showing using the `configure` command with a demo project I created to use with this tool:

```bash
$ ./ci-cd-tool configure
Configuring CI/CD pipeline...
Enter repository URL: https://github.com/amy324/CI-CD-Demo-Project
Enter branch name: main
Enter test script command: go test
Pipeline configuration saved to pipeline_config.yaml
```

### Running the Pipeline

To run the CI/CD pipeline, use the `run` command:

```bash
./ci-cd-tool run
```

This will fetch the source code from the configured repository, checkout the specified branch, and execute the test script command.

Here is the output for the demo poject above used as an example:
```bash
$ ./ci-cd-tool run
Running CI/CD pipeline...
Repository URL: https://github.com/amy324/CI-CD-Demo-Project
Branch Name: main
Test Script: go test
Enumerating objects: 6, done.
Counting objects: 100% (6/6), done.
Compressing objects: 100% (5/5), done.
Total 6 (delta 0), reused 6 (delta 0), pack-reused 0
PASS
ok      cicd-demo       0.269s
CI/CD pipeline completed successfully
```

## Configuration File

The pipeline configuration settings are saved in the `pipeline_config.yaml` file in YAML format. You can manually edit this file if needed.

## Technologies Used

- [Go](https://golang.org/): The programming language used to build the CI/CD tool.
- [cobra](https://github.com/spf13/cobra): A CLI library for Go that powers the command-line interface of the tool.

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---


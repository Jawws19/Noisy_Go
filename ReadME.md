# Noisy Go

**Noisy Go** is a Go implementation inspired by the original [Noisy Project](https://github.com/1tayH/noisy). It generates random web traffic using various protocols, HTTP methods, and user-agents to simulate realistic browsing behavior across a list of domains.

## Features

- 🌐 Randomized Protocols: Randomly selects between `http` and `https`.
- 📨 Randomized HTTP Methods: Supports random selection between `GET`, `POST`, `PUT`, `DELETE`, `HEAD`.
- 👤 Randomized User-Agents: Uses a set of popular user-agent strings from different browsers and platforms.
- ⚙️ Configurable Domains: Loads a list of domains from a YAML configuration file.
- 🚀 Concurrent Requests: Simulates multiple requests concurrently with customizable delays.

## Table of Contents

- [Requirements](#requirements)
- [Setup](#setup)
  - [Clone the Repository](#clone-the-repository)
  - [Prepare the YAML File](#prepare-the-yaml-file)
  - [Run the Project](#run-the-project)
- [Usage](#usage)
  - [Example Commands](#example-commands)
  - [Example Output](#example-output)
- [Contributing](#contributing)
- [License](#license)
- [Credits](#credits)

## Requirements

- Go 1.18 or later
- A YAML file with the list of domains.

## Setup

### Clone the Repository

Clone the repository to your local machine using Git:

```bash
git clone https://github.com/Jawws19/Noisy_Go.git
cd Noisy_Go
```

### Prepare the YAML File

Ensure that the `domains.yaml` file contains the list of domains to simulate traffic. Here's an example of the format:

```yaml
- domain: example.com
- domain: google.com
- domain: facebook.com
```

### Run the Project

Since the `main.go` file is located in the `cmd/` directory, you’ll need to run or build the project as follows:

#### To Build the Project

```bash
go build ./cmd
./Noisy_Go
```

#### To Run the Project Directly Without Building

```bash
go run ./cmd
```

You can also specify a custom configuration file using the `--config` flag:

```bash
go run ./cmd --config=your_domains.yaml
```

Alternatively, you can set the configuration file path via an environment variable:

```bash
export NOISY_CONFIG_PATH=your_domains.yaml
go run ./cmd
```

## Usage

When running the project, you will be prompted to enter a minimum delay and maximum delay between site visits (in seconds). This allows you to customize the intervals for the simulated requests.

### Example Commands

```bash
go run ./cmd --config=domains.yaml
```

You'll be asked to enter the time delays:

```bash
Enter minimum delay between site visits (in seconds): 1
Enter maximum delay between site visits (in seconds): 5
```

### Example Output

```bash
Request to https://example.com with method POST and User-Agent Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3 returned status: 200 OK
Request to http://google.com with method GET and User-Agent Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 returned status: 200 OK
Request to https://facebook.com with method DELETE and User-Agent Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.86 Mobile Safari/537.36 returned status: 403 Forbidden
```

## Contributing

Contributions are welcome! If you’d like to improve the project or suggest new features, feel free to:

1. Fork this repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.

## License

This project is licensed under the **GNU General Public License v3.0 (GPL-3.0)**, in accordance with the original project's license. For more information, see the [LICENSE](./LICENSE) file.

## Credits

This project is heavily inspired by the original [Noisy Project](https://github.com/1tayH/noisy) by **1tayH**. While this implementation is written in Go for improved performance, the overall idea and design are derived from the original project.
```


  Go Noisy Project

This project is a Go implementation inspired by the original [Noisy Project](https://github.com/1tayH/noisy). It generates random web traffic using various protocols, HTTP methods, and user-agents to simulate realistic browsing behavior across a list of domains.

   Features

- **Randomized Protocols**: Randomly selects between `http` and `https`.
- **Randomized HTTP Methods**: Supports random selection between `GET`, `POST`, `PUT`, `DELETE`, `HEAD`.
- **Randomized User-Agents**: Uses a set of popular user-agent strings from different browsers and platforms.
- **Configurable Domains**: Loads a list of domains from a YAML configuration file.
- **Concurrent Requests**: Simulates multiple requests concurrently with customizable delays.

   Requirements

- **Go 1.18 or later**
- A YAML file with the list of domains.

   Setup

    1. Clone the repository:

```bash
git clone https://github.com/yourusername/go-noisy-project.git
cd go-noisy-project
2. Prepare the YAML file (domains.yaml):
Ensure the domains.yaml file contains the list of domains to simulate traffic. Here's an example format:

yaml
Copy code
- domain: example.com
- domain: google.com

3. Run the project:
Since the main.go file is now located in the cmd/ directory, you’ll need to run the project using the following commands.

To Build:

bash
Copy code
go build ./cmd
This will generate an executable file, which you can run with:

bash
Copy code
./go_noisy_project
To Run Directly Without Building:

bash
Copy code
go run ./cmd
You can also specify a custom configuration file using the --config flag:

bash
Copy code
go run ./cmd --config=your_domains.yaml
Alternatively, you can set the configuration file path via an environment variable:

bash
Copy code
export NOISY_CONFIG_PATH=your_domains.yaml
go run ./cmd
4. Specify Delays:
The program will ask for the minimum and maximum delay between each request (in seconds). Simply provide the values when prompted.

Example Output

bash
Copy code
Request to https://example.com with method POST and User-Agent Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3 returned status: 200 OK
Request to http://google.com with method GET and User-Agent Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 returned status: 200 OK
Credit

This project is heavily inspired by the original Noisy Project by 1tayH. While this implementation is written in Go for improved performance, the overall idea and design are derived from the original project. Please check out and support the original work!

License

This project is licensed under the GNU General Public License v3.0 (GPL-3.0), in accordance with the original project's license.
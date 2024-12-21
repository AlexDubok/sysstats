# SysStats CLI

SysStats CLI is a Go-based command-line utility that collects device system data, such as memory usage, CPU usage, and temperature. It runs in the background when the system starts and exposes a webpage on localhost to display these statistics.

## Features

- Collects CPU usage, memory usage, and temperature data.
- Runs as a background service.
- Exposes a webpage on localhost to display system statistics.

## Installation

1. Clone the repository:

2. Build the project:
    ```sh
    go build -o sysstats-cli
    ```

3. Deploy the binary to your desired location:
    ```sh
    scp -R ./sysstats-cli user@remote:/path/to/deploy
    ```

## Usage

1. Run the binary:
    ```sh
    ./sysstats-cli
    ```

2. Open your web browser and navigate to `http://localhost:<port>/stats` to view the system statistics.

## Configuration

You can configure the HTTP server port using a command-line argument or an environment variable.

- Command-line argument:
    ```sh
    ./sysstats-cli --port=9090
    ```

- Environment variable:
    ```sh
    export APP_PORT=9090
    ./sysstats-cli
    ```

## Dependencies

This project uses the following dependencies:

- [github.com/shirou/gopsutil/v4](https://github.com/shirou/gopsutil)
- [github.com/tklauser/go-sysconf](https://github.com/tklauser/go-sysconf)
- [github.com/tklauser/numcpus](https://github.com/tklauser/numcpus)

## License

This project is licensed under the MIT License.
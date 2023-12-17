# TODO CLI Application
This Go project is a command-line interface (CLI) tool designed to fetch and display the status of TODO items. It accepts TODO item IDs from various sources, such as an input file, and retrieves their status from a source API. The tool provides flexibility by allowing users to specify options for input, output, and the number of TODO items to fetch.

## Prerequisites
* Go installed on your machine. [Install Go](https://golang.org/doc/install)
* (Optional) Docker installed if you want to run the application in a Docker container. [Install Docker](https://docs.docker.com/get-docker/)

## Usage
```bash
./todo-cli -i <input-file> [-o <output-file>] [-n <num-todos>] [-h]
```
#### Options
* `-i <input-file>`: Specifies the input file path containing TODO item IDs.
* `-o <output-file>`: Specifies the output file path for TODOs. If not provided, the output will be displayed in the console.
* `-n <num-todos>`: Specifies the number of even TODO IDs to fetch. The default value is 20 if not specified.
* `-h`: Usage of the command

## Examples
Fetch TODOs from the input file "input.txt" and display the output in the console:
```bash
./todo-cli -i input.txt
```

Fetch the first 10 TODOs from the input file "input.txt" and save the output to the file "output.csv":
```bash
./todo-cli -i input.txt -o output.csv -n 10
```

## Getting Started
1. Clone the repository:
    ```bash
    git clone https://github.com/bvpatel/todo-cli.git
    cd todo-cli
    ```
2. Build the application:
    ```bash
    make build
    ```
3. Run the application with the desired options:
    ```bash
    ./todo-cli -i input.txt -o output.csv -n 10
    ```

## Docker Support
Build the Docker image:
```bash
docker build -t todo-cli .
```

Run the Docker container:
```bash
docker run todo-cli
```

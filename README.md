# Can-I-Run-It

This project is a command-line tool written in Go that scrapes the minimum system requirements for a specific video game using the Colly web scraping library.

## Features

- Scrapes system requirements including:
  - Minimum CPU
  - Minimum RAM
  - Minimum Video Card
  - Dedicated Video RAM
  - Disk Space
  - Operating System (OS)
- Fast and lightweight
- Simple command-line interface for easy usage

## Installation

**Prerequisite**: Make sure Go is installed on your machine. You can download it from [here](https://golang.org/dl/).

1. Clone the repository:
    ```bash
    git clone https://github.com/shaaradnj/Can-I-Run-It.git
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the project:
```bash
go run .
```

# Nokia Device Manager with GoLang

A Go application for automating network device management, authentication, and configuration backup.

## Overview

This project aims to automate the management of network devices by facilitating authentication, configuration display, and backup creation. It leverages Telnet communication to interact with network devices and follows a structured organization for easy readability and maintenance.

## Features

- Read configurations from YAML files.
- Establish Telnet connection to network devices.
- Authenticate using user credentials.
- Display router interface information.
- Create backups of device configurations.
- Store backups locally.
- Clean up old backups.

## Getting Started

1. Clone this repository.
2. Install Go (Golang) on your system.
3. Update configuration files in the `config` directory.
4. Run the `main.go` file to initiate the backup procedure.

## Usage

1. Configure your network assets, user credentials, and application settings in the YAML files within the `config` directory.
2. Execute the `main.go` script to start the automated backup process.
3. Review the console output for progress and errors.
4. Backups will be stored in the specified file path.

## To-Do List

- [x] Basic automation for device authentication and backup.
- [ ] Implement database support for storing backups.
- [ ] Extend connection options to include SSH.
- [ ] Enhance error handling and reporting.
- [ ] Improve user interface and feedback.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

# Contributing to mcp-dingdingbot-server

Thank you for your interest in contributing to mcp-dingdingbot-server! This document provides guidelines and instructions for contributing to this project.

## Code of Conduct

Please be respectful and considerate of others when contributing to this project. We aim to foster an inclusive and welcoming community.

## How to Contribute

### Reporting Bugs

If you find a bug, please create an issue on GitHub with the following information:

- A clear, descriptive title
- A detailed description of the issue
- Steps to reproduce the bug
- Expected behavior
- Actual behavior
- Any relevant logs or screenshots

### Suggesting Enhancements

If you have an idea for an enhancement, please create an issue on GitHub with the following information:

- A clear, descriptive title
- A detailed description of the enhancement
- Any relevant examples or mockups

### Pull Requests

1. Fork the repository
2. Create a new branch from `master`
3. Make your changes
4. Run tests to ensure your changes don't break existing functionality
5. Submit a pull request

## Development Setup

1. Clone the repository
```bash
git clone https://github.com/HundunOnline/mcp-dingdingbot-server.git
cd mcp-dingdingbot-server
```

2. Build the project
```bash
make build
```

3. Run tests
```bash
go test -v ./...
```

## Coding Standards

- Follow Go best practices and conventions
- Write clear, descriptive commit messages
- Add tests for new functionality
- Update documentation as needed

## Environment Variables

The project uses the following environment variables:

- `DINGDING_BOT_WEBHOOK_KEY`: The webhook key for the DingDing Bot server (required)
- `DINGDING_BOT_SIGN_KEY`: The sign key for DingDing Bot signature verification (optional)

## License

By contributing to this project, you agree that your contributions will be licensed under the project's [GNU GPL v3 License](COPYRIGHT).

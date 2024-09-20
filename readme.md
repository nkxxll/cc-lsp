# cc-lsp - Conventional Commit Language Server

`cc-lsp` is a language server designed to assist developers in writing commit messages following the
[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. It provides
live feedback, autocompletion, and suggestions for conventional commit types such as `feat`, `fix`,
`chore`, `test`, and more. (Props to the
[educationalsp](https://github.com/tjdevries/educationalsp.git) by TJ Devries!!!)

## Features

- **Commit message correction**: Automatically corrects commit messages to fit the conventional
  commit format.
- **Autocompletion**: Provides autocompletion for commit types (`feat`, `fix`, `chore`, `test`,
  etc.).
- **Detailed commit type info**: Offers guidance on what each commit type signifies and when to use
  them.

## Installation

1. **Install Go (if not already installed)**:
   Ensure Go is installed on your system. You can download it [here](https://golang.org/dl/).

2. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/cc-lsp.git
   cd cc-lsp
   ```

3. **Build the project**:

   ```bash
   go build
   ```

4. **Install the language server**:

   ```bash
   go install
   ```

5. **Use in your editor**:
   Configure your editor to use `cc-lsp` as a language server for commit messages.

## Development

1. **Fork the repository**:
   Create a personal fork of the project by clicking "Fork" at the top-right corner of the repository page.

2. **Clone your fork**:

   ```bash
   git clone https://github.com/yourusername/cc-lsp.git
   cd cc-lsp
   ```

3. **Install dependencies**:
   Make sure you have all required dependencies by running:

   ```bash
   go mod tidy
   ```

4. **Run tests**:
   Run tests to ensure everything is working:

   ```bash
   go test ./...
   ```

5. **Make your changes**:
   Implement your feature or bugfix, following the [Go style guidelines](https://golang.org/doc/effective_go.html).

6. **Submit a Pull Request**:
   Once your changes are ready, push them to your fork and open a pull request. Contributions are
   very welcome as long as they follow the Go style guidelines!

## Contributions

We encourage contributions from the community! If you'd like to help improve `cc-lsp`, please follow
these steps:

1. **Open an issue**: If you find a bug or have a feature request, please open an issue so we can
   discuss it.
2. **Fork and clone**: Fork the repository and clone it to your local machine.
3. **Create a feature branch**: Make a new branch from `main` for your feature or bugfix.
4. **Follow Go style guidelines**: Keep your code clean and in line with the official [Go style
   guidelines](https://golang.org/doc/effective_go.html).
5. **Write tests**: Ensure that your contributions include tests for any new functionality or bug
   fixes.
6. **Submit a pull request**: Once your changes are ready, submit a PR, and we’ll review it as soon
   as possible.

We look forward to your contributions!

## TODO

- [ ] Add all Angular conventional commit message types.
- [ ] Provide customizable linting rules for commit messages.
- [ ] Integrate with popular editors (Neovim priority).
- [ ] Extend autocompletion for scopes.
- [ ] Improve the performance of the language server.
- [ ] Improve error handling and logging.

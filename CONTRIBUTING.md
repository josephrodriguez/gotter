# Contributing to Gotter

Thank you for considering contributing to Gotter! This document explains the preferred workflow, code standards, testing and review process so contributions are smooth and consistent.

## Table of contents

- How to contribute (issues / feature requests)
- Development workflow (fork, branch, PR)
- Code style and formatting
- Testing and linters
- Pull request checklist
- Code review and merging
- Security disclosures

## How to contribute

- Bug reports: open an issue with a clear title, steps to reproduce, expected vs actual behavior, and environment information (Go version, OS, Docker, etc.).
- Feature requests: open an issue describing the use case, design ideas, and any backward-compatibility considerations.
- Small fixes: open a pull request directly (preferred for trivial changes like docs or small typos).

Before opening a PR, please search existing issues and PRs to avoid duplicates.

## Development workflow

1. Fork the repository and create a new branch from `main`:

```bash
git clone https://github.com/<your-username>/gotter.git
cd gotter
git checkout -b feature/your-feature
```

2. Make your changes on a feature branch. Keep commits small and focused.

3. Run tests and linters locally (see below).

4. Push your branch and open a Pull Request against `josephrodriguez/gotter:main`.

5. Reference the issue in your PR (e.g. "Fixes #123") if applicable.

## Commit messages

- Use descriptive commit messages. We recommend following Conventional Commits (https://www.conventionalcommits.org/).
- Example: `feat(cli): add --config flag` or `fix(http): handle timeout error`.

## Code style and formatting

- Run `gofmt` on all changed files before submitting. You can auto-apply formatting:

```bash
gofmt -s -w .
```

- Keep code simple and idiomatic. Prefer explicit error handling over hiding errors.
- Do not commit generated files unless absolutely necessary.

## Testing and linters

Run unit tests and basic checks locally before creating a PR:

```bash
# run unit tests
go test ./...

# check formatting
gofmt -l .

# vet for common mistakes
go vet ./...

# optional security and lint tools (install beforehand)
# go install golang.org/x/vuln/cmd/govulncheck@latest
# govulncheck ./...
# go install github.com/securego/gosec/v2/cmd/gosec@latest
# gosec ./...
```

If your change affects Docker builds, ensure the image builds:

```bash
docker build -t gotter:local .
```

If you add new dependencies, run:

```bash
go mod tidy
```

and include any updated `go.sum` in your commit.

## Pull request checklist

Before requesting review, please ensure your PR:

- [ ] Has a clear title and description explaining the change and motivation
- [ ] References any related issue(s)
- [ ] Includes tests for new behavior or a rationale for not adding tests
- [ ] Passes `go test ./...` locally
- [ ] Is formatted with `gofmt -s -w` and `go mod tidy` has been run if dependencies changed
- [ ] Does not include unrelated changes

## Code review and merging

- Maintainters will review PRs and may request changes. Please respond to feedback promptly.
- For non-trivial changes, at least one approving review is required before merging.
- Squash merge is preferred for keeping history clean unless instructed otherwise.

## Security disclosures

If you discover a security vulnerability, **do not** open a public issue. Please follow the repository's `SECURITY.md` for the responsible disclosure process.

## Getting help

If you have questions about contributing, open an issue with the "question" label or join a discussion in the repository.

---

Thanks again for contributing to Gotter — we appreciate your help improving the project!
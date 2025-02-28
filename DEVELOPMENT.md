# Contributing to the Go MAAS Client

Thank you for your interest in contributing to the Go MAAS Client! We appreciate your help in making this project better. This document provides information on how to set up your development environment and best practices for contributing to the project. 

## Requirements

- [Go](https://golang.org/doc/install) >= 1.20
- A MAAS installation running, to connect to and test against. See the [maas-dev-setup](https://github.com/canonical/maas-dev-setup) repository for more information on a development setup.
- [CLA](https://ubuntu.com/legal/contributors) signed with the email used with git and GitHub.

## Branching strategy

This project follows a fork-based development model with a single long-running master branch. All contributions should be made via pull requests (PRs) from forked repositories.

1. Fork the Repository:
    1. Go to the repository on GitHub and click "Fork" in the top-right corner.
    1. Clone your fork locally:
       ```bash
       git clone <https-or-ssh-url-to-your-fork>
       cd gomaasclient
       ```
    1. Add the upstream repository (the original repo) as a second remote:
       ```bash
       git remote add upstream <https-or-ssh-url-to-original>
       ```
1. Create a Feature Branch:
    ```bash
    git checkout -b feat/feature-name
    ```
1. Keep Your Branch Up to Date:
    1. Before working, sync your branch with the latest changes from master:
       ```bash
       git fetch upstream
       git checkout master
       git merge upstream/master
       ```
    1. Then, rebase or merge your feature branch if necessary:
        ```bash
        git checkout feat/feature-name
        git rebase master
        ```   
1. Commit and Push Changes:
    1. Follow commit message guidelines (e.g., fix: correct typo in readme).
    1. Push your branch to your forked repository:
        ```bash
        git push origin feat/feature-name
        ```
1. Submit a Pull Request:
    1. Go to the your forked repository on GitHub.
    1. Click "New Pull Request". Select your feature branch to merge from your forked repo, into the master branch of the original repo.
    1. Ensure your PR includes:
       - A clear description of changes.
       - Links to relevant issues (e.g., Fixes #113).
       - Passing tests, if applicable.
1. Address Review Feedback. Once approved, a maintainer will merge your PR. 🎉

## Commit messages

We follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. Conventional Commits defines the following structure for the Git commit message:

```bash
<type>[scope][!]: <description>

[body]

[footer(s)]
```

Where 
- `type` is the kind of the change (e.g. feature, bug fix, documentation change, refactor).
- `scope` may be used to provide additional contextual information (e.g. which system component is affected). If scope is provided, it’s enclosed in parentheses.
- `!` MUST be added if commit introduces a breaking change.
- `description` is a brief summary of a change (try to keep it short, so overall title no more than 72 characters).
- `footer` is detailed information about the change (e.g. breaking change, related bugs, etc.).

## Development

To test the client: 

1. Ensure you have a MAAS installation running for the client to connect to.
1. An example development directory is provided in the `examples/clientdemo` directory. Copy this to a new directory, or ensure you don't push `env.sh` if you choose to work in the repository.
1. Add your environment variables to the `env.sh` file. If you need to find the MAAS API key, you can run the following command in your MAAS environment:
    ```bash
    sudo maas apikey --username=maas
    ```
1. Source the `env.sh` file:
   ```bash
   source env.sh
   ```
1. Ensure the path in the `replace` directive in `go.mod` points to your local copy of the c, the root of this repository.
1. In that directory, run:
   ```bash
   go mod tidy
   go run main.go
   ```
1. You should be able to see MAAS machine info in the output. Congratulations, you've used the GO MAAS client running locally! Modify the client as you see fit and submit a PR. Happy coding! 

## Testing

To run the unit tests, run the following command:
```bash
make test
```

## Getting Help

Check for existing issues [here](https://github.com/canonical/gomaasclient/issues), or open a new one for bugs and feature requests.

## Release Process

Releases are handled by the maintainers.

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [MAAS API Documentation](https://maas.io/docs/api)

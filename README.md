# Riconto

![project-image](./assets/riconto.svg)

A golang based markdown processor for generating PDFs.

![shields](https://img.shields.io/github/contributors/chordflower/riconto?style=for-the-badge)![shields](https://img.shields.io/github/languages/top/chordflower/riconto?style=for-the-badge)![shields](https://img.shields.io/github/check-runs/chordflower/riconto/development?style=for-the-badge)![shields](https://goreportcard.com/badge/github.com/chordflower/riconto?style=for-the-badge)![shields](https://img.shields.io/discord/1240686039006449674?style=for-the-badge)![shields](https://img.shields.io/github/issues/chordflower/riconto?style=for-the-badge)![shields](https://img.shields.io/github/issues-pr/chordflower/riconto?style=for-the-badge)![shields](https://img.shields.io/github/license/chordflower/riconto?style=for-the-badge)![shields](https://img.shields.io/github/go-mod/go-version/chordflower/riconto?style=for-the-badge)

## 🧐 Features

Here're some of the project's best features:

* Ability to create new projects
* Ability to handle multiple outputs in diferent folders
* Ability to handle common markdown
* Ability to include other markdown files inside one another
* Single binary installation

## 🛠️ Installation Steps:

### Requirements

- golang version 1.23.1 or what is currently on the go.mod file;

### Installation

1. Get the project using go install:

```shell
go install github.com/chordflower/riconto@latest
```

### Development

For development this repository uses [go-task](https://github.com/go-task/task), to simplify the tasks, so there is:

- go clean => To clean built files and test results;
- go build => To build the binaries from scratch;
- go lint => To run golintci-lint (you will need to install it first, since installation varies from system to system);
- go test => Runs the tests;
- go convey => Runs goconvey web interface (it will install it if not available);

## 🍰 Contribution Guidelines:

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/chordflower/riconto/issues)**: Submit bugs found or log feature requests for
  the `riconto` project.
- **[Submit Pull Requests](https://github.com/chordflower/riconto/pulls)**: Review open PRs, and
  submit your own PRs.

Note that before starting contributing:

- This repository uses oneflow as its branch management technique, meaning that:
  - There is one branch named `development`
    that is the main branch;
  - The new features are developed on a branch named `feature/<feature_name>` and start from `development`;
  - Bugfixes are developed on a branch named `bugfix/<bugfix_name>` and may start from `development`;
  - Feature and Bugfix branches are merged with:
     ```sh
     git merge --no-ff feature/<feature_name>
     ```
  - In github pullrequests, the merge message is: `merge: merge feature name branch`
- This repository also uses [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/), with the commit
  types being:
  - `feat` for features
  - `fix` for bugfixes
  - `chore` for noncode related things, like updating the readme, changing github workflows, updating build scripts,
    etc.
  - `refactor` for code changes that do not add or remove anything aka refactorings
  - `merge` for feature or bugfix branch merges.
- For commit scopes, the issue number can be used...

### Guidelines ###

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b feature/new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'feat: add new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and
   their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your
   contribution!

## 🛡️ License:

This project is licensed under the GPL-3.0 or later.

# gx

Effortless git branching and merging for lazy devs.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Limitations](#limitations)
- [License](#license)

## Introduction

`gx` essentially automates my Git workflow at work. The idea is to consistently work on the cleanest branch from the main or master branch. When it's time to merge this clean branch (or feature branch) into the shared development environment, you simply branch out and append '-dev' (or your choice identifier) to the branch. This desire for automation led to the creation of `gx`.

`gx` is designed to minimize the risk of unintentionally bringing unwanted code from the shared development environment to staging, release or, worse, directly to your master branch. Simple yet effective.

The name `gx` has no particular meaning.

## Installation

**Windows**

1. Dowload the executable from the [releases page](https://github.com/thisdotEric/gx/releases/tag/windows).
2. Follow this [steps](https://stackoverflow.com/a/41895179) to make `gx` globally available on your windows command line.

**Linux and macOS**

1. Clone the repository or download the source code.
2. Build the `gx` executable using the following commands:
    ```bash
    make build
    ```
3. Install `gx` globally on your system:
    ```bash
    make install
    ```
    This requires administrative privileges, so you might need to enter your password.
4. Optionally, you can clean up the generated files after installation:
    ```bash
    make clean
    ```

## Usage

Suppose you are on the `feat/#111-balances` feature branch. Running the following command:

```
$ gx dev
```

or piping the result of the `git commit` command.

```
$ git commit -m "feat: updated balance" | gx dev
```

will create the `feat/#111-balances-dev` branch incorporating all the changes.

`gx` defaults to branching out using 'dev' if you don't provide a branch name explicitly.

## Limitations

1. `gx` will fail on merge conflicts.
2. Merging strategy is using `git merge`, `git rebase` is not yet in the works.
3. Only supports one argument (for branch name) at the moment.

**Pull requests are welcome!**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

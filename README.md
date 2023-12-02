# gx

Effortless git branching and merging for lazy devs.

## Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Limitations](#limitations)
- [License](#license)

## Introduction

`gx` essentially automates my Git workflow at my day job. The idea is to consistently work on the cleanest branch from the main or master branch. When it's time to merge this clean branch (or feature branch) into the shared development environment, you simply branch out and append '-dev' (or your choice identifier) to the branch. This desire for automation, driven by my laziness, led to the creation of `gx`.

`gx` is designed to minimize the risk of unintentionally bringing unwanted code from the shared development environment to staging, release or, worse, directly to your master branch. Simple yet effective.

The name `gx` has no particular meaning.

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

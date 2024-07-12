# Contributing to Chiks

We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features

## Our Development Environment:

- VSCode
- Golang

## Pull Requests - **If you're here to add a stack, this is for you!**

Pull requests are the best way to propose changes to the codebase. We actively welcome your pull requests:

1. Fork the repo and create a new branch off of `main`.
2. ** If you're adding a new stack **
   1. Run `go run ./scripts/build_stack/build.go` in the terminal.
   2. Follow the prompts.
   - The script leverages your input and an AI workflow to generate a boilerplate for your stack.
3. If you've added code that should be tested, add tests/test it yourself.
4. Ensure the test suite passes.
5. Issue that pull request!

## Any contributions you make will be under the MIT Software License

In short, when you submit code changes, your submissions are understood to be under the same MIT License that covers the project. Feel free to contact the maintainers if that's a concern.

## Report bugs using Github's [issues](https://github.com/SamuelRCrider/chiks/issues)

We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/SamuelRCrider/chiks/issues/new); it's that easy!

## Write bug reports with detail, background, and sample code

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can.
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

People _love_ thorough bug reports.

## Use a Consistent Coding Style

- You can try running `go fmt` for style unification

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

## References

This document was adapted from the open-source contribution guidelines for [Facebook's Draft](https://github.com/facebook/draft-js/blob/a9316a723f9e918afde44dea68b5f9f39b7d9b00/CONTRIBUTING.md)

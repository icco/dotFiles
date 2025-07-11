# Nat Welch's .Files

This repository holds my config files for just about everything. Use this to set up a fresh OSX or Linux box, and Nat will be a happy user.

 * `link` &rarr; directly links file or folder to ~/.fname
 * `specific` &rarr; Links specific files, instead of their root folder.

## Install

Originally I did this entirely with shell scripts, then Ruby/Rake. Now I use Go for better performance and easier distribution.

### Go Commands

```bash
# Install dotfiles (equivalent to old `rake` command)
go run ./infect

# Update vim plugins and sort spell file (equivalent to old `rake vim`)
go run ./infect vim

# Run tests
go run ./infect test
```

### Other Setup

To install all the needed OSX packages, `brew bundle` in the top directory.

To switch to homebrew's version of Bash ([according to this doc](https://johndjameson.com/blog/updating-your-shell-with-homebrew/))

```
echo /usr/local/bin/bash | sudo tee /etc/shells
chsh -s /usr/local/bin/bash
```

### Things to remember

 * For more colors for things like `ls` on OSX, install grc: `brew install grc` (included in Brewfile).
 * Don't be a chump, use [rvm](https://rvm.io/).

## Notes

Apparently other people have made similar things to my infect script.

 * [technicalpickles/homesick](https://github.com/technicalpickles/homesick)
 * [holman/dotfiles](https://github.com/holman/dotfiles)
 * [ryanb/dotfiles](https://github.com/ryanb/dotfiles)

This only initializes the configuration of a user. I assume you've used something like [Fog](http://fog.io) to automate your system deployment and configuration.

## Assumptions

This repo assumes you are managing a variety of machines. It makes some trade-offs because it assumes you use multiple machines daily.

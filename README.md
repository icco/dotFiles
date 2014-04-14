# Nat Welch's .Files

This repository holds my config files for just about everything. Use this to set up a fresh OSX or Linux box, and Nat will be a happy user.

 * `link` &rarr; directly links file or folder to ~/.fname
 * `specific` &rarr; Links specific files, instead of their root folder.

## Install

Originally I did this entirely with shell scripts. Now I use rake. Make sure to `gem install rake` and install ruby 1.9.2 before going ahead with `rake infect`

 * For more colors for things like `ls` on OSX, install grc: `brew install grc`.
 * Don't be a chump, use [rvm](http://rvm.beginrescueend.com).

## Notes

Apparently other people have made similar things to my infect script.

 * [technicalpickles/homesick](https://github.com/technicalpickles/homesick)
 * [holman/dotfiles](https://github.com/holman/dotfiles)
 * [ryanb/dotfiles](https://github.com/ryanb/dotfiles)

This only initializes the configuration of a user. I assume you've used something like [Fog](http://fog.io) to automate your system deployment and configuration.


## Assumptions

This repo assumes you are managing a variety of machines. It makes some trade-offs because it assumes you use multiple machines daily.

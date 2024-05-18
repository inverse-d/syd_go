# syd
A tool to backup and restore dotfiles

## Functional requirements
- Support for config files
- Command line arguments
- Support for online repositories (Gihub, Gitlab)
- Written in Golang

## What it is supposed to do
To backup local dotfiles can be cumbersome, especially when they are not within `$HOME/.config/..` . Therefore the idea is to write a tool which collects configured .dotfiles from all specified locations, writes them into a local git repository and pushes those to a cloud hosted repository like Github or Gitlab. In case of a restore one can also use the tool to put the files back to the places where they belong. The tool can be either configured via a .gitnore styled file, or it can be instructed with arguments. 



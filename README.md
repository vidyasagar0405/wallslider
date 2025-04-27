# Wallslider

A simple tool to switch wallpapers

# Table of Contents
<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->

- [Requirements](#requirements)
- [Features](#features)
- [License](#license)

<!-- markdown-toc end -->

# Requirements
- Golang
- In X11 - Nitrogen
- In wayland - swww

# Features
- reindex     - Refresh the wallpaper index
- next        - Set the next wallpaper in the sequence
- prev        - Set the previous wallpaper in the sequence
- random      - Set a random wallpaper

# How it works ?

It indexes all the wallpapers in a given directory recursively, and saves the json at `XDG_CONFIG_DIR/wallslider/index.json`.

You can cycle through them with, `next`, `prev` and select a random one with `random`, and reindex if you wanna include new wallpapers

# License
MIT License

see LICENSE file for more

# mln

`mln` a replacement for ln written in Golang. `mln` create a symbolic link, not a hard link."

## Background

Each time I use the ln command, it puzzles me not knowing whether any of the arguments is source_file or target_file.
Also, maybe most people use the ln command to create symbolic links, but creating a symbolic link requires the -s option. I wanted to make a symbolic link easier, so I created `mln`.

## Screenshot

![image](./screenshot.gif)

## Installation

```
$ go get github.com/tkmru/mln
```

## Usage

In mln, a symbolic link is created successfully regardless of which of the arguments specifies the link destination and link name.

```
$ mln -h
A modern version of ln. `mln` create a symbolic link, not a hardlink.

Usage:
  mln [target <file/dir>/symlink file] [symlink file/target <file/dir>/] [flags]

Flags:
  -h, --help   help for mln
```

## License

MIT License

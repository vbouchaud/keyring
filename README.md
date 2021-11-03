# keyring

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/vbouchaud/keyring?style=for-the-badge)](https://github.com/vbouchaud/keyring/releases/latest)
[![License](https://img.shields.io/github/license/vbouchaud/keyring?style=for-the-badge)](https://opensource.org/licenses/MPL-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/vbouchaud/keyring?style=for-the-badge)](https://goreportcard.com/report/github.com/vbouchaud/keyring)

A very basic cli keyring tool to use accross various OS.

- [Usage](#usage)
  *  [Add](#add)
  *  [Get](#get)
  *  [Delete](#delete)
- [Distribution](#distribution)
  * [Binary](#binary)
  * [Linux](#linux)
    + [Archlinux](#archlinux)

## Usage
 
### Add
```
NAME:
   keyring add - add a secret to the keyring.

USAGE:
   keyring add [command options] [arguments...]

OPTIONS:
   --service value  The targeted service.
   --user value     The user to use for the targeted service.
   --secret value   The secret to set. (optional)
```

If the secret is not given in the command line, the user will be prompted for it.

Examples:
```
# interactive
 »  keyring add --service foo --user bar
secret: 

# not interactive
 »  keyring add --service foo --user bar --secret baz
```

### Get
```
NAME:
   keyring get - get an entry from the keyring.

USAGE:
   keyring get [command options] [arguments...]

OPTIONS:
   --service value  The targeted service.
   --user value     The user to get for the targeted service.
```

Example:
```
 »  keyring get --service foo --user bar                                                              ~
baz
```

### Delete
```
NAME:
   keyring delete - delete an entry from the keyring.

USAGE:
   keyring delete [command options] [arguments...]

OPTIONS:
   --service value  The targeted service.
   --user value     The user to remove for the targeted service.
```

Example:
```
 »  keyring delete --service foo --user bar                                                           ~
```

## Distribution

### Binary
Binaries for the following OS and architectures are available on the release page:
 - linux/arm64
 - linux/arm
 - linux/amd64
 - darwin/arm64
 - darwin/amd64
 - windows/amd64

### Linux
#### Archlinux
[![AUR version](https://img.shields.io/aur/version/keyring-cli?label=keyring-cli&style=for-the-badge)](https://aur.archlinux.org/packages/keyring-cli/)

[![AUR version](https://img.shields.io/aur/version/keyring-cli-bin?label=keyring-cli-bin&style=for-the-badge)](https://aur.archlinux.org/packages/keyring-cli-bin/)

[![AUR last modified](https://img.shields.io/aur/last-modified/keyring-cli-git?label=keyring-cli-git&style=for-the-badge)](https://aur.archlinux.org/packages/keyring-cli-git/)


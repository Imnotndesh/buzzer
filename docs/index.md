# Overview

**Buzzer** is a simple, fast, and convenient command-line tool for executing Wake-on-LAN (WoL) operations on your local network.

# Features

*   **Wake by MAC or Alias**: Send WoL packets directly to a MAC address or to a pre-saved alias.
*   **Alias Management**: Store, edit, and delete MAC addresses under easy-to-remember aliases.
*   **Script-Friendly**: Designed with a simple flag-based interface perfect for automation and scripting.
*   **Cross-Platform**: Single binary that runs on Linux.
*   **Auto-Completion**: Includes a bash completion script for a smoother terminal experience.

# Getting Started

## Installation

The easiest way to install `buzzer` is by downloading a pre-compiled package from [here](https://github.com/Imnotndesh/buzzer/releases).

### For Debian/Ubuntu

Download the `.deb` package and install it using `apt`.

```sh
# Replace v1.x.x with the latest version
wget https://github.com/Imnotndesh/buzzer/releases/download/v1.x.x/buzzer_1.x.x_linux_amd64.deb
sudo apt install ./buzzer_x.x.x_linux_amd64.deb
```

### For Fedora/RHEL/CentOS

Download the `.rpm` package and install it using `dnf`.

```sh
# Replace v1.x.x with the latest version
wget https://github.com/Imnotndesh/buzzer/releases/download/v1.x.x/buzzer-1.x.x.x86_64.rpm
sudo dnf install ./buzzer-x.x.x.x86_64.rpm
```

## Building from Source

If you have Go installed, you can build `buzzer` from the source code.

```sh
git clone https://github.com/Imnotndesh/buzzer.git
cd buzzer
go build -o buzzer .
sudo mv buzzer /usr/local/bin/
```
# Commands
Buzzer uses a simple subcommand structure for its operations

---

### `store [ALIAS] [MAC_ADDRESS]`

**S**tores a new machine by associating a MAC address with a memorable alias.
```sh
  buzzer store my-server 0A:1B:2C:3D:4E:5F
```
### `wake [ALIAS]`

**W**akes a machine using its stored alias.

> You can also specify a custom broadcast address using the `--via` flag:
>```sh
>  buzzer wake proxmox-server --via 10.0.0.255:7
>```
### `broadcast [MAC_ADDRESS]`

**B**roadcasts a Wake-on-LAN packet directly to a specific MAC address.

> You can also specify a custom broadcast address using the `--via` flag:
>```sh
>  buzzer broadcast 0A:1B:2C:3D:4E:5F --via 10.0.0.255:7
>```
### `list`

**L**ists all stored aliases and their corresponding MAC addresses.
```sh
  buzzer list
```
### `get [ALIAS]`

**G**ets and displays the MAC address associated with a stored alias.
```sh
  buzzer get proxmox-server
```
### `edit [ALIAS] [NEW_MAC_ADDRESS]`

**E**dits an existing entry to assign a new MAC address to an alias.
```sh
  buzzer edit my-server F0:E1:D2:C3:B4:A5
```

### `Listen`
**L**istens for any wake on lan packets sent to a device (default port 9)


> This can also take the argument `-port` to specify a port to listen to
> ```shell
>    buzzer listen -port 30219
>```
**NOTE**: By default this command needs privileged access by default unless a lower port number is specified
### `remove [ALIAS]`

**R**emoves an alias and its MAC address from the database.
```sh
  buzzer remove my-server
```
### `help`

Displays the **H**elp message with all available commands.
```sh
  buzzer version
```
### `-V`
Prints the current **V**ersion of the program.
```sh
  buzzer version
```

> The same information can be found in the man pages in linux which can be run using command:
> ```shell
> man 1 buzzer
>```
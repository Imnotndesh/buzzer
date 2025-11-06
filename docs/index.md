# Buzzer

**Buzzer** is a simple, fast, and convenient command-line tool for executing Wake-on-LAN (WoL) operations on your local network.

# Features

*   **Wake by MAC or Alias**: Send WoL packets directly to a MAC address or to a pre-saved alias.
*   **Alias Management**: Store, edit, and delete MAC addresses under easy-to-remember aliases.
*   **Script-Friendly**: Designed with a simple flag-based interface perfect for automation and scripting.
*   **Cross-Platform**: Single binary that runs on Linux.
*   **Auto-Completion**: Includes a bash completion script for a smoother terminal experience.

# Getting Started

## Installation

The easiest way to install `buzzer` is by downloading a pre-compiled package from the **GitHub Releases Page**.

### For Debian/Ubuntu

Download the `.deb` package and install it using `apt`.

```sh
# Replace v1.x.x with the latest version
wget https://github.com/Imnotndesh/buzzer/releases/download/v1.x.x/buzzer_1.x.x_linux_amd64.deb
sudo apt install ./buzzer_1.x.x_linux_amd64.deb
```

### For Fedora/RHEL/CentOS

Download the `.rpm` package and install it using `dnf`.

```sh
# Replace v1.x.x with the latest version
wget https://github.com/Imnotndesh/buzzer/releases/download/v1.x.x/buzzer-1.x.x.x86_64.rpm
sudo dnf install ./buzzer-1.x.x.x86_64.rpm
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

Buzzer uses case-insensitive flags to perform its functions.

---

### `-S [ALIAS] [MAC_ADDRESS]`

**S**tores a new machine by associating a MAC address with a memorable alias.

### `-W [ALIAS]`

**W**akes a machine using its stored alias.

### `-B [MAC_ADDRESS]`

**B**roadcasts a Wake-on-LAN packet directly to a specific MAC address.

### `-L`

**L**ists all stored aliases and their corresponding MAC addresses.

### `-G [ALIAS]`

**G**ets and displays the MAC address associated with a stored alias.

### `-E [ALIAS] [NEW_MAC_ADDRESS]`

**E**dits an existing entry to assign a new MAC address to an alias.

### `-R [ALIAS]`

**R**emoves an alias and its MAC address from the database.

### `-H`

Displays the **H**elp message with all available commands.

### `-V`

Prints the current **V**ersion of the program.
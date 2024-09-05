# Introduction
Buzzer is a CLI application that is intended to be a tool for executing Wake on LAN operations on your local network.

# Features
* Storage of MAC addresses under aliases for better management
* Convenient Flag based execution style that is suitable for scripting
* Alias management tools for deleting and editing stored mac address and alias pairs

# Getting Started  
## Installation
Fetch the latest release from [Releases Page](https://github.com/Imnotndesh/buzzer/releases)  
From the download Location open a terminal and run the program using the command `./buzzer`
## Building from Source
Make sure You have golang installed in your system then execute the following:
```
git clone 'https://github.com/Imnotndesh/buzzer.git'
cd Buzzer
go build -o buzzer buzzer.go
```
# Commands
Buzzer uses the following flags to achieve its functions:
## -h 
* Displays a summary of all usages for the program
## -b [MAC_ADDRESS]
* Wakes the computer using the passed MAC_ADDRESS
## -w [ALIAS]
* Uses the saved alias to wake the corresponding computer
## -s [ALIAS] [MAC_ADDRESS]
* Saves the passed Alias and Mac address for later use
## -l
* Prints out a list of saved mac addresses 
## -e [ALIAS] [MAC_ADDRESS]
* Assigns a new MAC address to the signed alias
## -g
* Fetches MAC Address tied to the alias passed
## -r [ALIAS]
* Removes passed alias and associated MAC_Address from storage
## -v
* Print out current program version
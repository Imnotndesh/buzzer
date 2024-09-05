# Buzzer
An uncomplicated Wake On Lan Program  
Full documentation is located [Here](https://imnotndesh.github.io/buzzer/)

# Usage
- On linux
```./buzzer [OPTIONS] [MAC_ADDRESS] [IP_ADDRESS]```

- On Windows
  ```./buzzer.exe [OPTIONS] [MAC_ADDRESS] [IP_ADDRESS]```
## Commands
| Flag | Usage                        | Description                                           |
|------|------------------------------|-------------------------------------------------------|
| -b   | -b [ MACHINE_MAC ]           | Wakes a machine using the passed MAC address          |
| -e   | -e [ ALIAS ] [ MAC_ADDRESS ] | Changes MAC_ADDRESS value for the passed ALIAS        |
| -g   | -g [ STORED_ALIAS ]          | Fetches MAC address bound to the alias                |
| -h   | -h                           | Prints out the help text                              |
| -l   | -l                           | Prints out all stored aliases and their bound MAC_ADDRESS |
| -s   | -s [ ALIAS ] [ MACHINE_MAC ] | Binds MAC address to an alias and stores it           |
| -v   | -v                           | Shows current version                                 |
| -w   | -w [ STORED_ALIAS ]          | Wakes computer using its stored alias name            |

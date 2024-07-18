# Buzzer
An uncomplicated Wake On Lan Program

# Usage
- On linux
```./buzzer [OPTIONS] [MAC_ADDRESS] [IP_ADDRESS]```

- On Windows
  ```./buzzer.exe [OPTIONS] [MAC_ADDRESS] [IP_ADDRESS]```

## Valid Options
| Flag | Usage                        | Description                                  |
|------|------------------------------|----------------------------------------------|
| -w   | -w [ STORED_ALIAS ]          | Wakes computer using its stored alias name   |
| -s   | -s [ ALIAS ] [ MACHINE_MAC ] | Binds MAC address to an alias and stores it  |
| -g   | -g [ STORED_ALIAS ]          | Fetches MAC address bound to the alias       |
| -b   | -b [ MACHINE_MAC ]           | Wakes a machine using the passed MAC address |
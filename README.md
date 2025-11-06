# Buzzer

**Buzzer** is a simple, fast, and convenient command-line tool for executing Wake-on-LAN (WoL) operations on your local network.

For full installation and usage instructions, please see the **[Full Documentation](https://imnotndesh.github.io/buzzer/)**.

## Quick Example

```sh
# Store a machine's MAC address with an easy-to-remember alias
buzzer store my-server 0A:1B:2C:3D:4E:5F

# Wake the machine using the alias
buzzer wake my-server

# Wake a machine on a different subnet using a custom broadcast address
buzzer wake my-server --via 192.168.2.255:9

# List all stored machines
buzzer list
```

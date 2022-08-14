# IsMyPortOpen?

## A dead-simple HTTP app to check if a port is open or not

---

Usage:

```bash
./impo <host> <port>
```

## Download

Head over to the [Releases page](https://github.com/YajTPG/IsMyPortOpen/releases) to download the latest verison or use the commands below.

---

### Linux/Darwin

```bash
curl -L "https://github.com/YajTPG/IsMyPortOpen/releases/latest/download/IsMyPortOpen.$(uname -s | awk '{print tolower($0)}')-$(uname -p)" -o impo
# You can change the -o flag parameter to "/usr/bin/impo" if you want the command to run globablly. (Also run the command as sudo)
```

### Windows

Just download it from the release page :(

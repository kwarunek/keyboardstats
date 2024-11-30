# keyboardstats

Tool to gather statistics about keyboard usage - specifically, how many times each key is pressed.

## Usage

### Install
Make's install target will build the binary and install it to `/opt/keyboardstats`, and install the systemd service.
```bash
make install
```

### Analyze
Because it uses sqlite as a backend (`/opt/keyboardstats/keyboardstats.db`), you can use any tool that can read from sqlite.

For example:
```bash
sqlite3 /opt/keyboardstats/keyboardstats.db 'select key, count(*) as cnt from keyboardstats group by key order by cnt DESC;'
```

## Security

This tool is essentially a keylogger, so use it at your own risk. 

While there are similar tools that use goevdev, due to security reasons and lack of trust, I have created a simple one that just reads from `/dev/input/event*`. It is suited for my own usage, and may not work for yours.

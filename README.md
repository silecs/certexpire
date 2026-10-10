certexpire
==========

A simple tool that checks remote certificates
and warns if they are to expire soon.

The parsing of certificates comes from the excellent
[certinfo](https://github.com/pete911/certinfo), used as a library.

Install
-------
The *Releases* page has Deb/Arch packages, and archives with the executable.

To build from source:
```
go build .
```

Examples
--------

Check if https://littre.org:443 has a valid certificate
with an expiration date in more than least 10 days (default).
```
$ certexpire littre.org
```
If there is no error, certexpire is silent,
unless the `--verbose` flag is used.

When an error occurs, there's a message and a non-empty return code.
```
$ certexpire --days 80 silecs.info littre.org
ERROR [silecs.info:443 TLS 1.3] expires in 68 days
ERROR [littre.org:443 TLS 1.3] expires in 47 days
```

Notifications
-------------

The intended usage was to create a periodical task
that would run daily on my PC and notify me directly.

The commands below reach this goal:
```sh
cp certexpire ~/.local/bin/
cp systemd/* ~/.config/systemd/user/
echo "TARGETS=littre.org silecs.info" > ~/.config/certexpire.env
# Test
systemctl --user start certexpire.unit
# Activate the daily task, starting now and at next session opening
systemctl --user enable --now certexpire.timer
```
If the executable is not in `~/.local/bin/`
(the Debian package puts it in `/usr/bin/`)
then the file `certexpire.service` must be amended.

Notifications work thanks to systemd user services and the `notify-send` command.
The files in `systemd/` contains many comments to explain the process.

Nota bene
---------
Here, 1 day == 24 hours.
I know it's an approximation, but I had no need for more precision than that.

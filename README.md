certexpire
==========

A simple tool that checks remote certificates
and warns if they are to expire soon.

The parsing of certificates comes from the excellent
[certinfo](https://github.com/pete911/certinfo), used as a library.

Install
-------
```
go build .
```

Examples
--------

Check if https://littre.org:443 has a valid certificate
that is valid for at least 10 days (default)
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

```
cp certexpire ~/.local/bin/
cp systemd/* ~/.config/systemd/user/
echo "TARGETS=littre.org silecs.info" > ~/.config/certexpire.env
systemct --user enable --now certexpire.timer
```

Nota bene
---------
Here, 1 day == 24 hours.
I know it's an approximation, but I had no need for more precision than that.

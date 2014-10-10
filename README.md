Command-line Riemann Event Sender in Go

## building

    make

## configuring

Default options read from `/etc/riemann.ini`.  See `etc/riemann.ini` for an example.

## using

    riemann-cli --service=my-service --state=critical --desc="omgwtfbbq!"

or

    echo "omgwtfbbq!" | riemann-cli --service=my-service --state=critical

All options from `riemann-cli --help`:

    Usage:
      riemann-cli [OPTIONS]
    
    Application Options:
          --riemann-host= Riemann host [$RIEMANN_HOST]
          --riemann-port= Riemann port (5555) [$RIEMANN_PORT]
          --proto=        Riemann protocol (udp) [$RIEMANN_PROTO]
          --debug         enable debug logging [$DEBUG]
          --config-file=  .ini config file (/etc/riemann.ini) [$CONFIG_FILE]
    
    event details:
          --ttl=          time to live, in seconds
          --time=         timestamp of event (unix epoch); defaults to now
          --tags=         event tags
          --attrs=        arbitrary key/value pairs
          --hostname=     hostname
          --service=      service name
          --state=        service state; free-form, but should be ok/warning/critical
          --desc=         description of service, state
          --metric=       metric value
    
    Help Options:
      -h, --help          Show this help message

## sample event

invocation:

    echo "everything's cool, yo" | \
        riemann-cli \
            --debug
            --host localhost \
            --service foobar \
            --tags baz,bap \
            --attrs qwerty:asdf \

output:

    {
        Ttl: 0
        Time: 1412895111
        Tags: [ riemann-cli, baz, bap ]
        Host: localhost
        Service: foobar
        State: ok
        Metric: 0
        Description: everything's cool, yo
        Attributes:map[
            qwerty: asdf
            sender: riemann-cli
            user: blalor
            cwd: /Users/blalor/devel/riemann-cli
            ppid: 22183
            parent_exe: /bin/bash
        ]
    }

Parent PID and executable, current working directory, and invoking user are
provided to allow tracing the sender.  `parent_exe` only works on Linux.

## credits

`go-flags` is pretty damn cool.  It does all the work of mapping command-line
arguments to slices and maps, along with reading options from a config file.

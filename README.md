Command-line Riemann Event Sender

## building

    make

## configuring

Default options read from `/etc/riemann.ini`.  See `etc/riemann.ini` for an example.

## using

    riemann-cli --service=my-service --state=critical --desc="omgwtfbbq!"

or

    echo "omgwtfbbq!" | riemann-cli --service=my-service --state=critical

`riemann-cli --help`:

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

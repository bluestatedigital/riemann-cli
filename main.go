package main

import (
    "os"
    "io/ioutil"
    "fmt"
    
    riemannSender "github.com/bluestatedigital/riemann-cli/sender"
    "github.com/amir/raidman"
    
    log "github.com/Sirupsen/logrus"
    flags "github.com/jessevdk/go-flags"
)

type Options struct {
    RiemannHost string `env:"RIEMANN_HOST"  long:"riemann-host"                            description:"Riemann host"`
    RiemannPort int    `env:"RIEMANN_PORT"  long:"riemann-port" default:"5555"             description:"Riemann port"`
    Proto       string `env:"RIEMANN_PROTO" long:"proto"        default:"udp"              description:"Riemann protocol"`
    Debug       bool   `env:"DEBUG"         long:"debug"                                   description:"enable debug logging"`
    IniFile     string `env:"CONFIG_FILE"   long:"config-file"  default:"/etc/riemann.ini" description:".ini config file"`
    
    Event struct {
        Ttl         float32           `long:"ttl"      description:"time to live, in seconds"`
        Time        int64             `long:"time"     description:"timestamp of event (unix epoch); defaults to now"`
        Tags        []string          `long:"tags"     description:"event tags"`
        Attributes  map[string]string `long:"attrs"    description:"arbitrary key/value pairs"`
        Host        string            `long:"hostname" description:"hostname"`
        Service     string            `long:"service"  description:"service name"`
        State       string            `long:"state"    description:"service state; free-form, but should be ok/warning/critical"`
        Description string            `long:"desc"     description:"description of service, state"`
        Metric      float64           `long:"metric"   description:"metric value"`
    } `group:"event details"`
}

func main() {
    var opts Options
    
    // create parsers
    parser := flags.NewParser(&opts, flags.HelpFlag | flags.PassDoubleDash)
    iniParser := flags.NewIniParser(parser)

    // parse command-line arguments
    _, err := parser.ParseArgs(os.Args)
    checkError("parsing arguments", err)
    
    if _, err := os.Stat(opts.IniFile); err == nil {
        err = iniParser.ParseFile(opts.IniFile)
        checkError("parsing ini file", err)
    }
    
    if opts.RiemannHost == "" {
        parser.WriteHelp(os.Stdout)
        log.Fatal("riemann host not provided")
    }
    
    if opts.Event.Service == "" {
        parser.WriteHelp(os.Stdout)
        log.Fatal("event service name not provided")
    }
    
    if opts.Debug {
        // Only log the warning severity or above.
        log.SetLevel(log.DebugLevel)
    }
    
    if opts.Event.Description == "" {
        bytes, err := ioutil.ReadAll(os.Stdin)
        checkError("error reading from stdin", err)
        opts.Event.Description = string(bytes)
    }

    addr := fmt.Sprintf("%s:%d", opts.RiemannHost, opts.RiemannPort)
    log.Debugf("connecting to %s with %s", addr, opts.Proto)
    
    riemann, err := raidman.Dial(opts.Proto, addr)
    checkError("connecting to Riemann", err)
    
    defer riemann.Close()
    
    log.Debug("creating sender")
    sender := riemannSender.NewSender(riemann)
    
    err = sender.Send(&raidman.Event{
        Ttl:         opts.Event.Ttl,
        Time:        opts.Event.Time,
        Tags:        opts.Event.Tags,
        Host:        opts.Event.Host,
        State:       opts.Event.State,
        Service:     opts.Event.Service,
        Metric:      opts.Event.Metric,
        Description: opts.Event.Description,
        Attributes:  opts.Event.Attributes,
    })
    checkError("sending event", err)
    
    log.Debug("done")
}

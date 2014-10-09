package sender

import (
    "time"
    "os"
    "os/user"
    "fmt"
    
    log "github.com/Sirupsen/logrus"
    "github.com/amir/raidman"
    "github.com/bluestatedigital/riemann-cli/libproc"
)

type Sender struct {
    client RiemannClient
}

func NewSender(client RiemannClient) *Sender {
    return &Sender{
        client: client,
    }
}

func (self *Sender) Send(event *raidman.Event) error {
    if event.Time == 0 {
        event.Time = time.Now().Unix()
    }
    
    // get current user
    usr, err := user.Current()
    if err != nil {
        log.Errorf("unable to determine current user: %v", err)
    } else {
        event.Attributes["user"] = usr.Username
    }
    
    // get current working directory
    cwd, err := os.Getwd()
    if err != nil {
        log.Errorf("unable to determine current working dir: %v", err)
    } else {
        event.Attributes["cwd"] = cwd
    }
    
    // get parent process' ID
    ppid := os.Getppid()
    event.Attributes["ppid"] = fmt.Sprintf("%d", ppid)
    
    // get parent process' path
    parentExe, err := libproc.ProcPath(ppid)
    if err != nil {
        log.Errorf("unable to determine parent's process executable: %v", err)
    } else {
        event.Attributes["parent_exe"] = parentExe
    }
    
    log.Debugf("sending event: %+v", event)
    
    return self.client.Send(event)
}

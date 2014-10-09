package sender

import "github.com/amir/raidman"

type RiemannClient interface {
    Close()
    Send(event *raidman.Event) error
}

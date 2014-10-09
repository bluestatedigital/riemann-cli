package sender_test

import (
    . "github.com/bluestatedigital/riemann-cli/sender"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    
    "github.com/stretchr/testify/mock"
    "github.com/bluestatedigital/riemann-cli/sender/mocks"
    "github.com/amir/raidman"
)

var _ = Describe("Sender", func() {
    var sender *Sender
    var event *raidman.Event
    
    var mockClient mocks.MockRiemannClient
    
    BeforeEach(func() {
        mockClient = mocks.MockRiemannClient{}
        
        sender = NewSender(&mockClient)
        
        // go-flags brings its own map, so need to emulate the behavior of
        // properties set via options parsed by go-flags
        event = &raidman.Event{
            Attributes: make(map[string]string),
        }
    })
    
    It("sends an event to Riemann", func() {
        mockClient.On("Send", mock.AnythingOfType("*raidman.Event")).Return(nil)
        
        sender.Send(event)
        
        mockClient.AssertExpectations(GinkgoT())
    })
    
    It("defaults timestamp to now", func() {
        mockClient.On("Send", mock.AnythingOfType("*raidman.Event")).Return(nil)
        
        sender.Send(event)
        
        mockClient.AssertExpectations(GinkgoT())
        
        evt := mockClient.Calls[0].Arguments.Get(0).(*raidman.Event)
        
        Expect(evt.Time).To(BeNumerically(">", 0))
    })
    
    It("provides current user", func() {
        // assumes owner of parent process and current process are the same
        mockClient.On("Send", mock.AnythingOfType("*raidman.Event")).Return(nil)
        
        sender.Send(event)
        
        mockClient.AssertExpectations(GinkgoT())
        
        evt := mockClient.Calls[0].Arguments.Get(0).(*raidman.Event)
        
        Expect(evt.Attributes["user"]).ToNot(BeEmpty())
    })
    
    It("provides cwd", func() {
        // assumes cwd of parent process and current process are the same
        mockClient.On("Send", mock.AnythingOfType("*raidman.Event")).Return(nil)
        
        sender.Send(event)
        
        mockClient.AssertExpectations(GinkgoT())
        
        evt := mockClient.Calls[0].Arguments.Get(0).(*raidman.Event)
        
        Expect(evt.Attributes["cwd"]).ToNot(BeEmpty())
    })
    
    It("provides parent's process id", func() {
        mockClient.On("Send", mock.AnythingOfType("*raidman.Event")).Return(nil)
        
        sender.Send(event)
        
        mockClient.AssertExpectations(GinkgoT())
        
        evt := mockClient.Calls[0].Arguments.Get(0).(*raidman.Event)
        
        Expect(evt.Attributes["ppid"]).ToNot(BeEmpty())
    })
    
    It("provides parent's executable path", func() {
        mockClient.On("Send", mock.AnythingOfType("*raidman.Event")).Return(nil)
        
        sender.Send(event)
        
        mockClient.AssertExpectations(GinkgoT())
        
        evt := mockClient.Calls[0].Arguments.Get(0).(*raidman.Event)
        
        Expect(evt.Attributes["parent_exe"]).ToNot(BeEmpty())
    })
})

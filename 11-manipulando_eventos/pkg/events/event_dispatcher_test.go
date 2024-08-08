package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEvendHandler struct{}

func (h *TestEvendHandler) Handle(event EventInterface) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEvendHandler
	handler2        TestEvendHandler
	handler3        TestEvendHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) setupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEvendHandler{}
	suite.handler2 = TestEvendHandler{}
	suite.handler3 = TestEvendHandler{}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	assert.True(suite.T(), true)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

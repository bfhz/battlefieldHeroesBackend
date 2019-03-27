package database

import (
	"time"

	"github.com/sirupsen/logrus"
)

// EventLogger is an implementation of the dbr.EventReceiver
type EventLogger struct {
	logger logrus.FieldLogger
}

func NewLogger() *EventLogger {
	return &EventLogger{logger}
}

// wrapFields is a helper func to cast map[string]string to map[string]interface{}
// and return logrus.Entry with specified fields.
func (n *EventLogger) wrapFields(kvs map[string]string) *logrus.Entry {
	fields := logrus.Fields{}
	for k, v := range kvs {
		fields[k] = v
	}
	return n.logger.WithFields(fields)
}

// Event receives a simple notification when various events occur
func (n *EventLogger) Event(eventName string) {
	n.logger.
		WithField("event", eventName).
		Debug("DB.Event")
}

// EventKv receives a notification when various events occur along with
// optional key/value data
func (n *EventLogger) EventKv(eventName string, kvs map[string]string) {
	n.wrapFields(kvs).
		WithField("event", eventName).
		Debug("DB.Event")
}

// EventErr receives a notification of an error if one occurs
func (n *EventLogger) EventErr(eventName string, err error) error {
	n.logger.
		WithField("event", eventName).
		Warnf("DB.EventErr: %v", err)
	return err
}

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data
func (n *EventLogger) EventErrKv(eventName string, err error, kvs map[string]string) error {
	n.wrapFields(kvs).
		WithField("event", eventName).
		Warnf("DB.EventErr: %v", err)
	return err
}

// Timing receives the time an event took to happen
func (n *EventLogger) Timing(eventName string, nanoseconds int64) {
	n.logger.
		WithField("ns", time.Duration(nanoseconds)).
		WithField("event", eventName).
		Debug("DB.Timing")
}

// TimingKv receives the time an event took to happen along with optional key/value data
func (n *EventLogger) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	n.wrapFields(kvs).
		WithField("ns", time.Duration(nanoseconds)).
		WithField("event", eventName).
		Debug("DB.Timing")
}

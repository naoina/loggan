package loggan

import (
	"fmt"
	"io"
	"time"
)

// Formatter is an interface that formatter for a log entry.
type Formatter interface {
	// Format formats a log entry.
	// Format writes formatted entry to the w.
	Format(w io.Writer, entry *Entry) error
}

// RawFormatter is a formatter that doesn't format.
// RawFormatter doesn't output the almost fields of the entry except the
// Message.
type RawFormatter struct{}

// Format outputs entry.Message.
func (f *RawFormatter) Format(w io.Writer, entry *Entry) error {
	_, err := io.WriteString(w, entry.Message)
	return err
}

// LTSVFormatter is the formatter of Labeled Tab-separated Values.
// See http://ltsv.org/ for more details.
type LTSVFormatter struct {
}

// Format formats an entry to Labeled Tab-separated Values format.
func (f *LTSVFormatter) Format(w io.Writer, entry *Entry) error {
	if _, err := fmt.Fprintf(w, "level:%v", entry.Level); err != nil {
		return err
	}
	if !entry.Time.IsZero() {
		if _, err := fmt.Fprintf(w, "\ttime:%v", entry.Time.Format(time.RFC3339Nano)); err != nil {
			return err
		}
	}
	if entry.Message != "" {
		if _, err := fmt.Fprintf(w, "\tmessage:%v", entry.Message); err != nil {
			return err
		}
	}
	for _, k := range entry.Fields.OrderedKeys() {
		if _, err := fmt.Fprintf(w, "\t%v:%v", k, entry.Fields.Get(k)); err != nil {
			return err
		}
	}
	return nil
}

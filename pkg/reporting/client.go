package reporting

import (
	"github.com/g1i1u/nuclei-edited/v3/pkg/output"
)

// Client is a client for nuclei issue tracking module
type Client interface {
	RegisterTracker(tracker Tracker)
	RegisterExporter(exporter Exporter)
	Close()
	Clear()
	CreateIssue(event *output.ResultEvent) error
	CloseIssue(event *output.ResultEvent) error
	GetReportingOptions() *Options
}

package assertions

import (
	"encoding/json"
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/smartystreets/assertions/internal/go-render/render"
)

type Serializer interface {
	serialize(expected, actual interface{}, message string) string
	serializeDetailed(expected, actual interface{}, message string) string
}

type failureSerializer struct{}

func (self *failureSerializer) serializeDetailed(expected, actual interface{}, message string) string {
	view := FailureView{
		Message:  message,
		Expected: render.Render(expected),
		Actual:   render.Render(actual),
	}
	serialized, _ := json.Marshal(view)
	return string(serialized)
}

func (self *failureSerializer) serialize(expected, actual interface{}, message string) string {
	view := FailureView{
		Message:  message,
		Expected: fmt.Sprintf("%+v", expected),
		Actual:   fmt.Sprintf("%+v", actual),
	}
	serialized, _ := json.Marshal(view)
	return string(serialized)
}

func newSerializer() *failureSerializer {
	return &failureSerializer{}
}

///////////////////////////////////////////////////////////////////////////////

// This struct is also declared in github.com/smartystreets/goconvey/convey/reporting.
// The json struct tags should be equal in both declarations.
type FailureView struct {
	Message  string `json:"Message"`
	Expected string `json:"Expected"`
	Actual   string `json:"Actual"`
}

///////////////////////////////////////////////////////

// noopSerializer just gives back the original message. This is useful when we are using
// the assertions from a context other than the GoConvey Web UI, that requires the JSON
// structure provided by the failureSerializer.
type noopSerializer struct{}

func (self *noopSerializer) serialize(expected, actual interface{}, message string) string {
	return message
}
func (self *noopSerializer) serializeDetailed(expected, actual interface{}, message string) string {
	return message
}

func diff(expected, actual interface{}) string {
	diff := diffmatchpatch.New()
	diffs := diff.DiffMain(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual), false)
	if !canDiff(diffs) {
		return ""
	}
	return fmt.Sprintf("\nDiff:     '%s'", diff.DiffPrettyText(diffs))
}

// canDiff returns true if the diff listing represents more equal segments than deleted/inserted segments.
func canDiff(diffs []diffmatchpatch.Diff) bool {
	var (
		equal    int
		deleted  int
		inserted int
	)

	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffEqual:
			equal += len(diff.Text)
		case diffmatchpatch.DiffDelete:
			deleted += len(diff.Text)
		case diffmatchpatch.DiffInsert:
			inserted += len(diff.Text)
		}
	}

	return equal > deleted && equal > inserted
}

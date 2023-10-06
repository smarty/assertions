package must

import (
	"strings"
	"testing"
)

func TestFatal(t *testing.T) {
	output := Equal(1, 1)
	if strings.HasPrefix(output, "<<FATAL>>") {
		t.Error("Unexpected 'fatal' prefix")
	}
	output = Equal(1, 2)
	if !strings.HasPrefix(output, "<<FATAL>>") {
		t.Error("Missing expected 'fatal' prefix from output:", output)
	}
}

package pipe

import (
	"testing"
	"fmt"
	"strings"
)

func TestPipe(t *testing.T) {
	// 大部分时候，系统中是不存在gopipe-test这样的进程的
	p := NewPipe("ps", "-ef").P("grep", "gopipe-test", "-c")
	buffer, err := p.Call()
	if err != nil {
		t.Errorf("call error %s", err)
	}

	if buffer.String() != "1\n" {
		t.Error("first call error")
	}
	fmt.Printf("buffer %s", buffer.String())

	p2 := NewPipe("ps", "-ef").P("grep", "gopipe-test").P("grep", "-v", "grep", "-c")
	buffer, err = p2.Call()
	if err != nil {
		t.Errorf("call error %s", err)
	}
	if buffer.String() != "0\n" {
		t.Error("second call error")
	}
	fmt.Printf("buffer2 %s", buffer.String())

	p3 := NewPipe("cat", "pipe.go").P("grep", "func").P("awk", "{print $3}").P("grep", "string").P("wc", "-l")
	buffer, err = p3.Call()
	if err != nil {
		t.Errorf("call error %s", err)
	}

	buf := strings.TrimSpace(buffer.String())
	if buf != "1" {
		t.Error("third call error")
	}
	fmt.Printf("buffer3 %s", buffer.String())
}
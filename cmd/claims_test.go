package cmd

import (
	"bytes"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"os"
	"testing"
)

func TestClaims(t *testing.T)  {
	testCase := []struct{
		args []string
		want string
	}{
		{args: []string{"claims", "-s", "-u", os.Getenv("UID")}, want: "Custom Claims Set Up\n"},
	}

	for _, c := range testCase {
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := rwi.New(
			rwi.WithWriter(out),
			rwi.WithErrorWriter(errOut),
		)
		exit := Execute(ui, c.args)
		if exit != exitcode.Normal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Normal)
		}
		if out.String() != c.want {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), c.want)
		}
		if errOut.String() != "" {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), "")
		}
	}
}
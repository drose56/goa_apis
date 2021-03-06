// Code generated by goa v3.7.6, DO NOT EDIT.
//
// add HTTP client CLI support package
//
// Command:
// $ goa gen add/design

package cli

import (
	addc "add/gen/http/add/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `add addnums
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` add addnums --a 2651578921377290614 --b 2752948265278239320` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		addFlags = flag.NewFlagSet("add", flag.ContinueOnError)

		addAddnumsFlags = flag.NewFlagSet("addnums", flag.ExitOnError)
		addAddnumsAFlag = addAddnumsFlags.String("a", "REQUIRED", "Left operand")
		addAddnumsBFlag = addAddnumsFlags.String("b", "REQUIRED", "Right operand")
	)
	addFlags.Usage = addUsage
	addAddnumsFlags.Usage = addAddnumsUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "add":
			svcf = addFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "add":
			switch epn {
			case "addnums":
				epf = addAddnumsFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "add":
			c := addc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "addnums":
				endpoint = c.Addnums()
				data, err = addc.BuildAddnumsPayload(*addAddnumsAFlag, *addAddnumsBFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// addUsage displays the usage of the add command and its subcommands.
func addUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers.
Usage:
    %[1]s [globalflags] add COMMAND [flags]

COMMAND:
    addnums: Addnums implements addnums.

Additional help:
    %[1]s add COMMAND --help
`, os.Args[0])
}
func addAddnumsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] add addnums -a INT -b INT

Addnums implements addnums.
    -a INT: Left operand
    -b INT: Right operand

Example:
    %[1]s add addnums --a 2651578921377290614 --b 2752948265278239320
`, os.Args[0])
}

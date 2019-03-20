package jiracmd

import (
	"fmt"
	"github.com/coryb/figtree"
	"github.com/coryb/oreo"
	"gopkg.in/Netflix-Skunkworks/go-jira.v1"
	"gopkg.in/Netflix-Skunkworks/go-jira.v1/jiracli"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type ViewOptions struct {
	jiracli.CommonOptions `yaml:",inline" json:",inline" figtree:",inline"`
	jira.IssueOptions     `yaml:",inline" json:",inline" figtree:",inline"`
	Issue                 string `yaml:"issue,omitempty" json:"issue,omitempty"`
}

func CmdViewRegistry() *jiracli.CommandRegistryEntry {
	opts := ViewOptions{
		CommonOptions: jiracli.CommonOptions{
			Template: figtree.NewStringOption("view"),
		},
	}

	return &jiracli.CommandRegistryEntry{
		"Prints issue details",
		func(fig *figtree.FigTree, cmd *kingpin.CmdClause) error {
			jiracli.LoadConfigs(cmd, fig, &opts)
			return CmdViewUsage(cmd, &opts)
		},
		func(o *oreo.Client, globals *jiracli.GlobalOptions) error {
			return CmdView(o, globals, &opts)
		},
	}
}

func CmdViewUsage(cmd *kingpin.CmdClause, opts *ViewOptions) error {
	jiracli.BrowseUsage(cmd, &opts.CommonOptions)
	jiracli.TemplateUsage(cmd, &opts.CommonOptions)
	jiracli.GJsonQueryUsage(cmd, &opts.CommonOptions)
	cmd.Flag("expand", "field to expand for the issue").StringsVar(&opts.Expand)
	cmd.Flag("field", "field to return for the issue").StringsVar(&opts.Fields)
	cmd.Flag("property", "property to return for issue").StringsVar(&opts.Properties)
	cmd.Arg("ISSUE", "issue id to view").Required().StringVar(&opts.Issue)
	return nil
}

// View will get issue data and send to "view" template
func CmdView(o *oreo.Client, globals *jiracli.GlobalOptions, opts *ViewOptions) error {
	data, err := jira.GetIssue(o, globals.Endpoint.Value, opts.Issue, opts)
	if err != nil {
		return err
	}
	fmt.Printf("%v", data)
	if opts.Browse.Value {
		return CmdBrowse(globals, opts.Issue)
	}
	return nil
}

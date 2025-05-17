package cobradoc

import (
	"bytes"
	"io"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func generate(cmd *cobra.Command, fmt Format, opts Options, w io.Writer) error {
	prepareCommand(cmd, opts)

	prepareOptions(cmd, &opts)

	fmtInfo := makeFormatInfo(cmd, opts)

	return format(fmt, fmtInfo, w)
}

func prepareCommand(cmd *cobra.Command, opts Options) {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	if opts.HideHelp {
		if cmd.Name() == "help" {
			cmd.Hidden = true
		}
		if cmd.Flags().Lookup("help") != nil {
			// HasAvailableFlags and PrintDefaults will exclude --help
			cmd.Flags().MarkHidden("help")
		}
	}

	for _, subCmd := range cmd.Commands() {
		prepareCommand(subCmd, opts)
	}
}

func prepareOptions(cmd *cobra.Command, opts *Options) {
	if opts.SectionNumber == "" {
		opts.SectionNumber = "1"
	}

	if opts.Name == "" {
		opts.Name = cmd.Name()
	}

	if opts.Date == "" {
		opts.Date = time.Now().Format("Jan 2006")
	}

	if opts.Language == "" {
		opts.Language = "en"
	}

	languageTag, err := language.Parse(opts.Language)
	if err != nil {
		languageTag = language.English
	}

	if opts.Header == "" {
		opts.Header = cases.Title(languageTag).String(opts.Name) + " Manual"
	}

	if opts.Footer == "" {
		opts.Footer = cases.Title(languageTag).String(opts.Name) + " Manual"
	}

	if opts.ShortDescription == "" {
		opts.ShortDescription = cmd.Short
	}

	if opts.LongDescription == "" {
		opts.LongDescription = cmd.Long
	}
}

func makeFormatInfo(cmd *cobra.Command, opts Options) formatInfo {
	var fmtInfo formatInfo

	fmtInfo.Options = opts
	fmtInfo.Groups = makeGroupsInfo(cmd)

	globalFlags := cmd.PersistentFlags()

	if globalFlags.HasAvailableFlags() {
		fmtInfo.GlobalFlagsBlock = makeFlagsBlock(globalFlags)

		globalFlags.VisitAll(func(flag *pflag.Flag) {
			if flag.Hidden {
				return
			}
			fmtInfo.GlobalFlags = append(fmtInfo.GlobalFlags, makeFlagInfo(flag))
		})
	}

	return fmtInfo
}

func makeGroupsInfo(cmd *cobra.Command) []groupInfo {
	var grpInfoList []groupInfo

	for _, grp := range getGroups(cmd) {
		var grpInfo groupInfo

		grpInfo.Title = grp.Title

		for _, subCmd := range getCommands(cmd) {
			if subCmd.Hidden || len(subCmd.Deprecated) != 0 {
				continue
			}
			if len(cmd.Groups()) != 0 && subCmd.GroupID != grp.ID {
				continue
			}
			grpInfo.Commands = append(grpInfo.Commands, makeCommandInfo(subCmd))
		}

		if len(grpInfo.Commands) == 0 {
			continue
		}

		grpInfoList = append(grpInfoList, grpInfo)
	}

	return grpInfoList
}

func makeCommandInfo(cmd *cobra.Command) commandInfo {
	var cmdInfo commandInfo

	cmdInfo.Path = cmd.CommandPath()
	cmdInfo.Usage = cmd.UseLine()

	if cmd.Long != "" {
		cmdInfo.Description = cmd.Long
	} else {
		cmdInfo.Description = cmd.Short
	}

	cmdFlags := cmd.NonInheritedFlags()

	if cmdFlags.HasAvailableFlags() {
		cmdInfo.FlagsBlock = makeFlagsBlock(cmdFlags)

		cmdFlags.VisitAll(func(flag *pflag.Flag) {
			if flag.Hidden {
				return
			}
			cmdInfo.Flags = append(cmdInfo.Flags, makeFlagInfo(flag))
		})
	}

	return cmdInfo
}

func makeFlagInfo(flag *pflag.Flag) flagInfo {
	return flagInfo{
		Long:            flag.Name,
		Short:           flag.Shorthand,
		DefaultValue:    flag.DefValue,
		ValueIsOptional: flag.NoOptDefVal != "",
		IsBool:          flag.Value.Type() == "bool",
		Type:            flag.Value.Type(),
		Description:     flag.Usage,
	}
}

func makeFlagsBlock(flags *pflag.FlagSet) string {
	var buf bytes.Buffer

	flags.SetOutput(&buf)

	if flags.HasAvailableFlags() {
		flags.PrintDefaults()
	}

	return buf.String()
}

func getGroups(cmd *cobra.Command) []cobra.Group {
	var groups []cobra.Group

	if len(cmd.Groups()) == 0 {
		groups = []cobra.Group{
			{
				ID:    "",
				Title: "Commands",
			},
		}
	} else {
		for _, grp := range cmd.Groups() {
			groups = append(groups, *grp)
		}
		groups = append(groups, cobra.Group{
			ID:    "",
			Title: "Additional Commands",
		})
	}

	return groups
}

func getCommands(cmd *cobra.Command) []*cobra.Command {
	var commands []*cobra.Command

	for _, subCmd := range cmd.Commands() {
		commands = append(commands, subCmd)

		nestedCommands := getCommands(subCmd)
		if len(nestedCommands) != 0 {
			commands = append(commands, nestedCommands...)
		}
	}

	return commands
}

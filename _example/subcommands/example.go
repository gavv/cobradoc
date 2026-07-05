package main

import (
	"fmt"
	"os"

	"github.com/gavv/cobradoc"
	"github.com/spf13/cobra"
)

const (
	groupMain = "main"
	groupDocs = "docs"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "Example command",
}

var mainCmd = &cobra.Command{
	GroupID: groupMain,
	Use:     "main",
	Short:   "Run main commands",
}

var helloCmd = &cobra.Command{
	GroupID: groupMain,
	Use:     "hello",
	Short:   "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		count, _ := cmd.Flags().GetInt("count")
		for i := 0; i < count; i++ {
			fmt.Println("hello!")
		}
	},
}

var byeCmd = &cobra.Command{
	GroupID: groupMain,
	Use:     "bye",
	Short:   "Say goodbye",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("bye, %s!\n", name)
	},
}

var docsCmd = &cobra.Command{
	GroupID: groupDocs,
	Use:     "docs",
	Short:   "Generate documentation",
}

var manCmd = &cobra.Command{
	GroupID: groupDocs,
	Use:     "man",
	Short:   "Generate manual page",
	Run: func(cmd *cobra.Command, args []string) {
		opts := docOptions()
		opts.HideParents = true
		cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Troff, opts)
	},
}

var markdownCmd = &cobra.Command{
	GroupID: groupDocs,
	Use:     "markdown",
	Short:   "Generate markdown page",
	Run: func(cmd *cobra.Command, args []string) {
		opts := docOptions()
		opts.RenderHierarchy = true
		cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Markdown, opts)
	},
}

func docOptions() cobradoc.Options {
	return cobradoc.Options{
		LongDescription: "Example command using cobra and cobradoc.",
		ExtraSections: []cobradoc.ExtraSection{
			{
				Title: cobradoc.BUGS,
				Text:  "Please report bugs at https://github.com/gavv/cobradoc",
			},
		},
		HideCompletion: true,
		HideHelp:       true,
	}
}

func init() {
	cobra.EnableCommandSorting = false

	helloCmd.Flags().IntP("count", "c", 1, "how many times to greet?")
	helloCmd.Flags().SortFlags = false
	helloCmd.InheritedFlags().SortFlags = false

	byeCmd.PersistentFlags().StringP("name", "n", "John", "who got the goodbye?")
	byeCmd.Flags().SortFlags = false
	byeCmd.InheritedFlags().SortFlags = false

	rootCmd.PersistentFlags().StringP("value", "v", "", "string value")
	rootCmd.PersistentFlags().Bool("flag", false, "boolean flag")
	rootCmd.PersistentFlags().SortFlags = false

	rootCmd.AddGroup(&cobra.Group{
		Title: "Main Commands",
		ID:    groupMain,
	})

	mainCmd.AddGroup(&cobra.Group{
		Title: "Main Commands",
		ID:    groupMain,
	})
	mainCmd.AddCommand(helloCmd)
	mainCmd.AddCommand(byeCmd)
	rootCmd.AddCommand(mainCmd)

	rootCmd.AddGroup(&cobra.Group{
		Title: "Documentation Commands",
		ID:    groupDocs,
	})

	docsCmd.AddGroup(&cobra.Group{
		Title: "Documentation Commands",
		ID:    groupDocs,
	})
	docsCmd.AddCommand(manCmd)
	docsCmd.AddCommand(markdownCmd)
	rootCmd.AddCommand(docsCmd)
}

func main() {
	rootCmd.Execute()
}

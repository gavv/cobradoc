package main

import (
	"fmt"
	"os"

	"github.com/gavv/cobradoc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "Example command",
}

var mainCmd = &cobra.Command{
	Use:   "main",
	Short: "Run main commands",
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		count, _ := cmd.Flags().GetInt("count")
		for i := 0; i < count; i++ {
			fmt.Println("hello!")
		}
	},
}

var byeCmd = &cobra.Command{
	Use:   "bye",
	Short: "Say goodbye",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("bye, %s!\n", name)
	},
}

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation",
}

var docOptions = cobradoc.Options{
	LongDescription: "Example command using cobra and cobradoc.",
	ExtraSections: []cobradoc.ExtraSection{
		{
			Title: cobradoc.BUGS,
			Text:  "Please report bugs at https://github.com/gavv/cobradoc",
		},
	},
	HideParents:     true,
	HideCompletion:  true,
	HideHelp:        true,
	RenderHierarchy: true,
}

var manCmd = &cobra.Command{
	Use:   "man",
	Short: "Generate manual page",
	Run: func(cmd *cobra.Command, args []string) {
		cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Troff, docOptions)
	},
}

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Generate markdown page",
	Run: func(cmd *cobra.Command, args []string) {
		cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Markdown, docOptions)
	},
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

	mainCmd.AddCommand(helloCmd)
	mainCmd.AddCommand(byeCmd)
	rootCmd.AddCommand(mainCmd)

	docsCmd.AddCommand(manCmd)
	docsCmd.AddCommand(markdownCmd)
	rootCmd.AddCommand(docsCmd)
}

func main() {
	rootCmd.Execute()
}

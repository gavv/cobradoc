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
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
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

var manCmd = &cobra.Command{
	GroupID: groupDocs,
	Use:     "man",
	Short:   "Generate manual page",
	Run: func(cmd *cobra.Command, args []string) {
		cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Troff, cobradoc.Options{
			LongDescription: "Example command using cobra and cobradoc.",
			ExtraSections: []cobradoc.ExtraSection{
				{
					Title: cobradoc.BUGS,
					Text:  "Please report bugs at https://github.com/gavv/cobradoc",
				},
			},
		})
	},
}

var markdownCmd = &cobra.Command{
	GroupID: groupDocs,
	Use:     "markdown",
	Short:   "Generate markdown page",
	Run: func(cmd *cobra.Command, args []string) {
		cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Markdown, cobradoc.Options{
			LongDescription: "Example command using cobra and cobradoc.",
		})
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

	rootCmd.AddGroup(&cobra.Group{
		Title: "Main Commands",
		ID:    groupMain,
	})

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(byeCmd)

	rootCmd.AddGroup(&cobra.Group{
		Title: "Documentation Commands",
		ID:    groupDocs,
	})

	rootCmd.AddCommand(manCmd)
	rootCmd.AddCommand(markdownCmd)
}

func main() {
	rootCmd.Execute()
}

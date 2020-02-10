package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	. "github.com/logrusorgru/aurora"
)


var rootCmd = &cobra.Command{
	Use:   "mln [target <file/dir>/symlink file] [symlink file/target <file/dir>/]",
	Short: "A modern version of ln. `mln` create a symbolic link, not a hardlink.",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		target := ""
		symbolic_link := ""
		for _, arg := range args {
			_, err := os.Stat(arg)
			if err != nil {
				symbolic_link = arg
			} else {
				target = arg
			}
		}
		
		if target == "" {
			fmt.Fprint(os.Stderr, BrightRed("Error: Target file must be an existing file.\n").Bold())
			os.Exit(1)
		}

		if symbolic_link == "" {
			fmt.Fprint(os.Stderr, BrightRed("Error: Symbolic link must not be an existing file.\n").Bold())
			os.Exit(1)
		}

		if err := os.Symlink(target, symbolic_link); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Successfully linked")
		fmt.Printf("%s%s %s %s %s%s %s\n", Brown("From"), Gray(8-1, ":"), Cyan(symbolic_link).Bold(),
			Magenta("->"),  Brown("To"), Gray(8-1, ":"), Cyan(target).Bold())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

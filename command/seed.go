package command

import "github.com/spf13/cobra"

//	"newsfeed/ent"

//"fmt"

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run sync",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

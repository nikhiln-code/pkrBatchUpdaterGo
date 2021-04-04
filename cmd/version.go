package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)
func init(){
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Prints version number of pkrPublisher tool",
	Long: `This command can be used to get the version number of publisher tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pkbPublisher v0.0.1-alpha")
	},
}


package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
	"nordea/npl/norway/pkrPublisher/v2/processor"
)

func init(){
	rootCmd.AddCommand(publishCmd)
	publishCmd.AddCommand(publishEPKCmd)
	publishCmd.AddCommand(publishPKBCmd)
	publishCmd.AddCommand(checkInputFileCmd)
}


var publishCmd = &cobra.Command{
	Use: "publish",
	Aliases: []string{"pub"},
	Short: "publish to the system specified",
	Long: `This command will publish the input JSON to the `,
}

var publishPKBCmd = &cobra.Command{
	Use: "pkb",
	Short: "publishes pkb to the hub",
	Long: `This command can be used to publish the pkb data to the hub`,
	Run: func(cmd *cobra.Command, args []string) {
	    fmt.Println("Executing the 'pkrPublisher publish pkb'")
	},
}

var publishEPKCmd = &cobra.Command{
	Use: "epk",
	Short: "publishes epk to the hub",
	Long: `This command can be used to publish the epk data to the hub`,
	Run: func(cmd *cobra.Command, args []string) {
		processor.ProcessEPK()
		fmt.Println("Executing the 'pkrPublisher publish epk'")
	},
}



var checkInputFileCmd = &cobra.Command{
	Use: "checkInput",
	Short:"checks the input file",
	Long: `Checks the coorectness of input data`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing checking of the ")
	},
}
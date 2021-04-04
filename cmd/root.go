package cmd

import  (
	"fmt"
	"os"
	_ "nordea/npl/norway/pkrPublisher/v2/helper"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


	var(
		//This is used for config file
		cfgFile string
		rootCmd = &cobra.Command{
			Use: "pkrPublisher",
			Short: "pkrPublisher: A simple CLI application",
			Long: `pkrPublisher: A simple CLI application to publish the pkr data to the hub on weekly basis`,
		}
	)

	func Execute() error{
		return rootCmd.Execute();
	}

	func init(){
		cobra.OnInitialize(initConfig)
		rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pkrPublisher.yaml)")
	}

	func er(msg interface{}){
		fmt.Println("Error:", msg)
		os.Exit(1)
	}

	func initConfig(){
		if cfgFile != ""{
			viper.SetConfigFile(cfgFile)
		} else {
			home, err := homedir.Dir()
			if err != nil{
				er(err);
			}
			viper.AddConfigPath(home)
			viper.SetConfigName(".pkrPublisher")
		}

		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err == nil{
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}




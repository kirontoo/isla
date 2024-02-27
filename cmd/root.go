package cmd
 
import (
  "github.com/spf13/cobra"
  "fmt"
  "os"
  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string

	rootCmd = &cobra.Command{
		Use:   "isla",
		Short: "A cli to edit frontmatter in blog posts",
		Long: `Isla is a CLI tool that edits frontmatter for posts in 
    YAML and TOML formats`,
	}
)

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
  // rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf1/")
  // rootCmd.PersistentFlags().StringP("author", "a", "Amy Dang", "Author name for copyright attribution")
  // rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
  viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
  viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
  viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
  viper.SetDefault("author", "Amy Dang ndangamy@gmail.com")
  viper.SetDefault("license", "MIT")
}

func initConfig() {
  // Don't forget to read config either from cfgFile or from home directory!
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".cobra" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".cobra")
  }

  if err := viper.ReadInConfig(); err != nil {
    fmt.Println("Can't read config:", err)
    os.Exit(1)
  }
}

/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package cmd

import (
	"log"
	"os/user"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		done := make(chan bool)

		// Process events
		go func() {
			for {
				select {
				case ev := <-watcher.Events:
					if ev.Op.String() == "CREATE" {
						log.Println(user.Username)
					}
					log.Println("event:", ev)
				case err := <-watcher.Errors:
					log.Println("error:", err)
				}
			}
		}()

		err = watcher.Add("testDir")
		if err != nil {
			log.Fatal(err)
		}

		<-done

		/* ... do stuff ... */
		watcher.Close()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

package cmd

import "github.com/spf13/cobra"

type structPodsOptions struct {
	namespace string
	count     int
	image     string
}

var podsOptions = structPodsOptions{}

var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "Bench startup pods ",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(podCmd)
	podCmd.PersistentFlags().StringVarP(&podsOptions.namespace, "namespace", "n", "", "Select Namespace for clean")
	podCmd.PersistentFlags().StringVarP(&podsOptions.image, "image", "i", "", "Docker image")
	podCmd.PersistentFlags().IntVarP(&podsOptions.count, "count", "c", 10, "Number of pods")
}

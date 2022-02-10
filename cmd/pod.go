package cmd

import (
	"context"

	"github.com/k-tools/k-bench/internal"
	"github.com/spf13/cobra"
)

var podsOptions = internal.StructPodsOptions{}

var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "Bench startup pods ",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pod(cmd.Context(), podsOptions)

	},
}

func init() {
	rootCmd.AddCommand(podCmd)
	podCmd.PersistentFlags().StringVarP(&podsOptions.Namespace, "namespace", "n", "", "Select Namespace for bench")
	podCmd.PersistentFlags().StringVar(&podsOptions.Name, "name", "", "Name for pod")
	podCmd.PersistentFlags().StringVarP(&podsOptions.Image, "image", "i", "", "Docker image")
	//podCmd.PersistentFlags().IntVarP(&podsOptions.Count, "count", "c", 10, "Number of pods")
}

func pod(ctx context.Context, p internal.StructPodsOptions) error {
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	err = client.CreatePo(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

package cmd

import (
	"context"

	"github.com/k-tools/k-bench/internal"
	"github.com/spf13/cobra"
)

var eventOptions = internal.StructPodsOptions{}

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Bench get time pulled image in event ",
	RunE: func(cmd *cobra.Command, args []string) error {
		return event(cmd.Context(), eventOptions)

	},
}

func init() {
	rootCmd.AddCommand(eventCmd)
	eventCmd.PersistentFlags().StringVarP(&eventOptions.Namespace, "namespace", "n", "", "Select Namespace for get event")
	eventCmd.PersistentFlags().StringVar(&eventOptions.Name, "name", "", "String contain in name pod for event")
}

func event(ctx context.Context, p internal.StructPodsOptions) error {
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	err = client.GetEvent(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

package cmd

import (
	"context"
	"github.com/loft-sh/devpod-gcloud-provider/pkg/gcloud"
	"github.com/loft-sh/devpod-gcloud-provider/pkg/options"
	"github.com/loft-sh/devpod/pkg/log"
	"github.com/spf13/cobra"
)

// DeleteCmd holds the cmd flags
type DeleteCmd struct{}

// NewDeleteCmd defines a command
func NewDeleteCmd() *cobra.Command {
	cmd := &DeleteCmd{}
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an instance",
		RunE: func(_ *cobra.Command, args []string) error {
			options, err := options.FromEnv()
			if err != nil {
				return err
			}

			return cmd.Run(context.Background(), options, log.Default)
		},
	}

	return deleteCmd
}

// Run runs the command logic
func (cmd *DeleteCmd) Run(ctx context.Context, options *options.Options, log log.Logger) error {
	client, err := gcloud.NewClient(ctx, options.Project, options.Zone)
	if err != nil {
		return err
	}
	defer client.Close()

	return client.Delete(ctx, options.MachineID)
}
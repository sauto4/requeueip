/*
Copyright 2022 The RequeueIP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/sauto4/requeueip/pkg/config"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := o.Config()
		if err != nil {
			return err
		}

		return run(ctrl.SetupSignalHandler(), conf)
	},
}

func init() {
	nfs := o.Flags()
	fs := daemonCmd.Flags()
	for _, f := range nfs.FlagSets {
		fs.AddFlagSet(f)
	}
	viper.BindPFlags(fs)

	rootCmd.AddCommand(daemonCmd)
}

func run(ctx context.Context, conf *config.Config) error {
	return nil
}

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

package options

import "github.com/spf13/pflag"

type RuntimeOptions struct {
	ConfigFile string
	Debug      bool
}

func NewRuntimeOptions() *RuntimeOptions {
	return &RuntimeOptions{
		ConfigFile: "$HOME/.requeueip-controller.yaml",
		Debug:      false,
	}
}

func (ro *RuntimeOptions) AddFlags(fs *pflag.FlagSet, d *RuntimeOptions) {
	fs.StringVar(&ro.ConfigFile, "config", d.ConfigFile, "The config file of RequeueIP controller.")
	fs.BoolVar(&ro.Debug, "debug", d.Debug, "Enable debug mode.")
}

func (ro *RuntimeOptions) Validate() []error {
	var errs []error

	return errs
}

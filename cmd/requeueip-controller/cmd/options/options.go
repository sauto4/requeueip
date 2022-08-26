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

import (
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/sauto4/requeueip/pkg/config"
)

type Options struct {
	RuntimeOptions *RuntimeOptions
}

func NewOptions() *Options {
	return &Options{
		RuntimeOptions: NewRuntimeOptions(),
	}
}

func (o *Options) Flags() cliflag.NamedFlagSets {
	var nfs cliflag.NamedFlagSets
	o.RuntimeOptions.AddFlags(nfs.FlagSet("generic"), o.RuntimeOptions)

	return nfs
}

func (o *Options) Config() (*config.Config, error) {
	if errs := o.validate(); len(errs) != 0 {
		return nil, utilerrors.NewAggregate(errs)
	}

	c, err := config.TryLoadFromDisk()
	if err != nil {
		return nil, err
	}

	// TODO(iiiceoo): Merge flags to config read by file

	return c, nil
}

func (o *Options) validate() []error {
	var errs []error
	errs = append(errs, o.RuntimeOptions.Validate()...)

	return errs
}

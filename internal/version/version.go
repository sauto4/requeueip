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

package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

var (
	// Version is semantic version.
	Version = "v0.0.0-master+$Format:%h$"
	// BuildDate in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ').
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit sha1 from git, output of $(git rev-parse HEAD).
	GitCommit = "$Format:%H$"
	// GitTreeState state of git tree, either "clean" or "dirty".
	GitTreeState = ""
)

// Info contains versioning information.
type Info struct {
	// Program info
	Version string `json:"Version"`
	// Git info
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	// Build info
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	if s, err := info.Text(); err == nil {
		return s
	}

	return info.Version
}

// ToJSON returns the JSON string of version information.
func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

// Text encodes the version information into a human readable format.
func (info Info) Text() (string, error) {
	text := strings.Builder{}
	text.WriteString("Version: " + info.Version + "\n")
	text.WriteString("GitCommit: " + info.GitCommit + "\n")
	text.WriteString("GitTreeState: " + info.GitTreeState + "\n")
	text.WriteString("BuildDate: " + info.BuildDate + "\n")
	text.WriteString("GoVersion: " + info.GoVersion + "\n")
	text.WriteString("Compiler: " + info.Compiler + "\n")
	text.WriteString("Platform: " + info.Platform + "\n")

	return text.String(), nil
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return Info{
		Version:      Version,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

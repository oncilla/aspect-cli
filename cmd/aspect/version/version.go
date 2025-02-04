/*
 * Copyright 2022 Aspect Build Systems, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package version

import (
	"context"

	"github.com/spf13/cobra"

	"aspect.build/cli/buildinfo"
	"aspect.build/cli/pkg/aspect/root/flags"
	"aspect.build/cli/pkg/aspect/version"
	"aspect.build/cli/pkg/bazel"
	"aspect.build/cli/pkg/interceptors"
	"aspect.build/cli/pkg/ioutils"
)

func NewDefaultCmd() *cobra.Command {
	return NewCmd(ioutils.DefaultStreams, bazel.FindFromWd)
}

func NewCmd(streams ioutils.Streams, bzlProvider bazel.BazelProvider) *cobra.Command {
	v := version.New(streams)
	v.BuildInfo = *buildinfo.Current()

	cmd := &cobra.Command{
		Use:     "version",
		Short:   "Print the versions of Aspect CLI and Bazel",
		Long:    `Prints version info on colon-separated lines, just like bazel does`,
		GroupID: "common",
		RunE: interceptors.Run(
			[]interceptors.Interceptor{
				flags.FlagsInterceptor(streams),
			},
			func(ctx context.Context, cmd *cobra.Command, args []string) (exitErr error) {
				bzl, err := bzlProvider()
				if err != nil {
					return err
				}
				return v.Run(cmd, bzl, args)
			},
		),
	}
	return cmd
}

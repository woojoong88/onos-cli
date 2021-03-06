// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kpimon

import "github.com/spf13/cobra"

func getListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list {metrics} [args]",
		Short: "List KPIMON resources",
	}
	cmd.AddCommand(getListMetricsCommand())
	return cmd
}

func getWatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch {metrics} [args]",
		Short: "Watch KPIMON resources",
	}
	cmd.AddCommand(getWatchMetricsCommand())
	return cmd
}

func getSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set {report_interval} [args]",
		Short: "Set KPIMON parameters",
	}
	cmd.AddCommand(setReportIntervalCommand())
	return cmd
}

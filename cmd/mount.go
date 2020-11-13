/*
Copyright © 2020 Hugues Obolonsky <hugh@atosc.org>

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

// Package cmd mount commands
package cmd

import (
	"fmt"
	"gotsx/tsxcommand"

	"github.com/spf13/cobra"
)

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:       "mount",
	Short:     "command to get mount status informations",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"atParkPosition", "getPosition", "isConnected", "isParked", "isStopped", "isTracking", "parkAndDisconnect"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Command take at least one sub command\n")
		cmd.Help()
	},
}

// isStoppedCmd represents the isStopped command
var isStoppedCmd = &cobra.Command{
	Use:   "isStopped",
	Short: "verify that the mount is stopped",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.IsMountStopped, "Is mount Stopped ?")
	},
}

// isConnectedCmd represents the isConnected command
var isConnectedCmd = &cobra.Command{
	Use:   "isConnected",
	Short: "verify that mount is connected in theSkyX",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.IsMountConnected, "Is mount Connected ?")
	},
}

// isParkedCmd represents the isParked command
var isParkedCmd = &cobra.Command{
	Use:   "isParked",
	Short: "verify that the mount is parked",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.IsMountParked, "Is mount Parked ?")
	},
}

// isTrackingCmd represents the isTracking command
var isTrackingCmd = &cobra.Command{
	Use:   "isTracking",
	Short: "verify that the mount is in sideral tracking",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.IsMountTracking, "Is mount Tracking ?")
	},
}

// atParkPositionCmd represents the atParkPosition command
var atParkPositionCmd = &cobra.Command{
	Use:   "atParkPosition",
	Short: "check that the AstroPhysics mount is at park1 position or parked",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.IsMountAtPark, "Is mount at Park1 position ?")
	},
}

// getPositionCmd represents the getPosition command
var getPositionCmd = &cobra.Command{
	Use:   "getPosition",
	Short: "get current Alt/Az position of the mount",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.GetMountPosition, "Get mount alt/az position: ")
	},
}

// parkAndDisconnectCmd represents the parkAndDisconnect command
var parkAndDisconnectCmd = &cobra.Command{
	Use:   "parkAndDisconnect",
	Short: "should diconnect the mount from theSkyX (removed implementation)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("parkAndDisconnect called")
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
	mountCmd.AddCommand(isStoppedCmd)
	mountCmd.AddCommand(isConnectedCmd)
	mountCmd.AddCommand(isParkedCmd)
	mountCmd.AddCommand(isTrackingCmd)
	mountCmd.AddCommand(atParkPositionCmd)
	mountCmd.AddCommand(getPositionCmd)
	mountCmd.AddCommand(parkAndDisconnectCmd)
}

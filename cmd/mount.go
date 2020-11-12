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

	"github.com/spf13/cobra"
)

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mount called")
	},
}

// isStoppedCmd represents the isStopped command
var isStoppedCmd = &cobra.Command{
	Use:   "isStopped",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("isStopped called")
	},
}

// isConnectedCmd represents the isConnected command
var isConnectedCmd = &cobra.Command{
	Use:   "isConnected",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("isConnected called")
	},
}

// isParkedCmd represents the isParked command
var isParkedCmd = &cobra.Command{
	Use:   "isParked",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("isParked called")
	},
}

// isTrackingCmd represents the isTracking command
var isTrackingCmd = &cobra.Command{
	Use:   "isTracking",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("isTracking called")
	},
}

// atParkPositionCmd represents the atParkPosition command
var atParkPositionCmd = &cobra.Command{
	Use:   "atParkPosition",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("atParkPosition called")
	},
}

// getPositionCmd represents the getPosition command
var getPositionCmd = &cobra.Command{
	Use:   "getPosition",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getPosition called")
	},
}

// parkAndDisconnectCmd represents the parkAndDisconnect command
var parkAndDisconnectCmd = &cobra.Command{
	Use:   "parkAndDisconnect",
	Short: "A brief description of your command",
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

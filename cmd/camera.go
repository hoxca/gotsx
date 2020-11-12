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

// Package cmd camera
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cameraCmd represents the camera command
var cameraCmd = &cobra.Command{
	Use:   "camera",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("camera called")
	},
}

// getCoolerPowerCmd represents the getCoolerPower command
var getCoolerPowerCmd = &cobra.Command{
	Use:   "getCoolerPower",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getCoolerPower called")
	},
}

// getCoolerStatusCmd represents the getCoolerStatus command
var getCoolerStatusCmd = &cobra.Command{
	Use:   "getCoolerStatus",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getCoolerStatus called")
	},
}

// getTemperatureCmd represents the getTemperature command
var getTemperatureCmd = &cobra.Command{
	Use:   "getTemperature",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getTemperature called")
	},
}

// getTemperatureSetPointCmd represents the getTemperatureSetPoint command
var getTemperatureSetPointCmd = &cobra.Command{
	Use:   "getTemperatureSetPoint",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getTemperatureSetPoint called")
	},
}

// isCoolerOffCmd represents the isCoolerOff command
var isCoolerOffCmd = &cobra.Command{
	Use:   "isCoolerOff",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("isCoolerOff called")
	},
}

func init() {
	rootCmd.AddCommand(cameraCmd)
	cameraCmd.AddCommand(getCoolerPowerCmd)
	cameraCmd.AddCommand(getCoolerStatusCmd)
	cameraCmd.AddCommand(getTemperatureCmd)
	cameraCmd.AddCommand(getTemperatureSetPointCmd)
	cameraCmd.AddCommand(isCoolerOffCmd)
}

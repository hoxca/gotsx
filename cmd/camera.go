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
	"gotsx/tsxcommand"

	"github.com/spf13/cobra"
)

// cameraCmd represents the camera command
var cameraCmd = &cobra.Command{
	Use:       "camera",
	Short:     "command to get camera status informations",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"getCoolerPower", "getCoolerStatus", "getTemperature", "getTemperatureSetPoint", "isCoolerOff"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command take at least one subcommand\n")
		cmd.Help()
	},
}

// getCoolerPowerCmd represents the getCoolerPower command
var getCoolerPowerCmd = &cobra.Command{
	Use:   "getCoolerPower",
	Short: "get the cooler power in percentage",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.GetCameraCoolerPower, "Get camera cooler power ?")
	},
}

// getCoolerStatusCmd represents the getCoolerStatus command
var getCoolerStatusCmd = &cobra.Command{
	Use:   "getCoolerStatus",
	Short: "get information from the camera cooler",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.GetCameraCoolerStatus, "Get camera cooler status ?")
	},
}

// getTemperatureCmd represents the getTemperature command
var getTemperatureCmd = &cobra.Command{
	Use:   "getTemperature",
	Short: "current temperature of the ccd camera",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.GetCameraTemperature, "Get camera temperature ?")
	},
}

// getTemperatureSetPointCmd represents the getTemperatureSetPoint command
var getTemperatureSetPointCmd = &cobra.Command{
	Use:   "getTemperatureSetPoint",
	Short: "target temperature for the camera",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.GetCameraTemperatureSetPoint, "Get camera temperature set point ?")
	},
}

// isCoolerOffCmd represents the isCoolerOff command
var isCoolerOffCmd = &cobra.Command{
	Use:   "isCoolerOff",
	Short: "verify if camera warmup end and regulation is off",
	Run: func(cmd *cobra.Command, args []string) {
		Ret = tsxcommand.Send(Tsx, tsxcommand.IsCoolerOff, "Is camera temperature regulation stopped ?")
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

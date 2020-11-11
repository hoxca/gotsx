package tsxcommand

const statusCommand = `
	  /* Java Script */
	  var Out;

		if (sky6RASCOMTele.IsConnected==0) {
		  Out = 'Mount is not connected';
		} else {
		  if (sky6RASCOMTele.IsParked()) {
			  Out = 'Mount is parked';
      } else {
			  Out = 'Mount is not parked';
			}
		} `

const isMountConnected = `
/* Java Script */
var Out;

if (sky6RASCOMTele.IsConnected != 0) {
	Out = 'OK, Mount is connected';
}`

const isMountParked = `
/* Java Script */
var Out;

if (sky6RASCOMTele.IsParked()) {
	Out = 'OK, Mount at park position';
}`

const isMountTracking = `
/* Java Script */
var Out;

if (sky6RASCOMTele.IsTracking != 0)  {
	Out = 'OK, Mount is tracking'
}`

const isMountStopped = `
/* Java Script */
var Out;

if (sky6RASCOMTele.IsTracking == 0 || sky6RASCOMTele.IsParked())  {
	Out = 'OK, Mount is stopped'
}`

const getMountPosition = `
/* Java Script */
var Out;

sky6RASCOMTele.GetAzAlt()
 Out = 'Alt: '+sky6RASCOMTele.dAlt+'\nAz : '+sky6RASCOMTele.dAz
`

const getCameraCoolerStatus = `
var Out, Reg;
Reg = (ccdsoftCamera.RegulateTemperature == 1) ? "On":"Off"

Out = 'Temp : '+ccdsoftCamera.Temperature+'\nPower: '+ccdsoftCamera.ThermalElectricCoolerPower+'\nRegulation: '+Reg
`

const isCoolerOff = `
var Out, Reg;

if (ccdsoftCamera.RegulateTemperature == 0) {
	Out = "Ok, Cooler power: "+ccdsoftCamera.ThermalElectricCoolerPower+"% Regulation: off"
} else {
  Reg = (ccdsoftCamera.RegulateTemperature == 1) ? "On":"Off"
	Out = "Nope!, Cooler power: "+ccdsoftCamera.ThermalElectricCoolerPower+"% Regulation: "+Reg
}
`
const parkAndDisconnectCommand = `
	  /* Java Script */
	  var Out;
	  if (sky6RASCOMTele.IsParked()) {
      sky6RASCOMTele.Unpark();
      sky6RASCOMTele.Park();
    } else {
      sky6RASCOMTele.Park();
		}
	  Out = sky6RASCOMTele.IsConnected; `

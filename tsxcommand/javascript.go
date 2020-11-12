// Package tsxcommand define const and functions
// for Software Bisque TheSkyX javascript API
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
		}`

// IsMountConnected javascript function
const IsMountConnected = `
   /* Java Script */
   var Out;

   if (sky6RASCOMTele.IsConnected != 0) {
   	Out = 'OK, Mount is connected';
   }`

// IsMountParked javascript function
const IsMountParked = `
   /* Java Script */
   var Out;

   if (sky6RASCOMTele.IsParked()) {
      Out = 'OK, Mount is parked';
   }`

// IsMountAtPark javascript function
const IsMountAtPark = `
   /* Java Script */
   var Out;

   if (sky6RASCOMTele.IsParked()) {
      Out = 'OK, Mount is parked';
   } else {
      sky6RASCOMTele.GetAzAlt()
      if ( Math.round(sky6RASCOMTele.dAlt-0.1) == 0 && Math.round(sky6RASCOMTele.dAz+0.1) == 360 && sky6RASCOMTele.IsTracking == 0 ) {
        Out = 'OK, Mount stopped at park1 position';
			} else if (sky6RASCOMTele.IsTracking != 0) {
			  Out = 'Nope!, mount is tracking';
			} else {
        Out = 'Nope!, \nAlt: '+Math.round(sky6RASCOMTele.dAlt-0.1)+'\nAz : '+Math.round(sky6RASCOMTele.dAz+0.1);
			}
	 }`

// IsMountTracking javascript function
const IsMountTracking = `
   /* Java Script */
   var Out;

   if (sky6RASCOMTele.IsTracking != 0)  {
   	Out = 'OK, Mount is tracking'
   }`

// IsMountStopped javascript function
const IsMountStopped = `
   /* Java Script */
   var Out;

   if (sky6RASCOMTele.IsTracking == 0 || sky6RASCOMTele.IsParked())  {
   	Out = 'OK, Mount is stopped'
   }`

// GetMountPosition javascript function
const GetMountPosition = `
   /* Java Script */
   var Out;

   sky6RASCOMTele.GetAzAlt()
   Out = 'Alt: '+Math.round(sky6RASCOMTele.dAlt-0.1)+'\nAz : '+Math.round(sky6RASCOMTele.dAz+0.1)`

// GetCameraCoolerStatus javascript function
const GetCameraCoolerStatus = `
   /* Java Script */
   var Out, Reg;
   Reg = (ccdsoftCamera.RegulateTemperature == 1) ? "On":"Off"

   Out = 'Temp : '+ccdsoftCamera.Temperature+'\nPower: '+ccdsoftCamera.ThermalElectricCoolerPower+'%\nRegulation: '+Reg`

// IsCoolerOff javascript function
const IsCoolerOff = `
   /* Java Script */
   var Out, Reg;

   if (ccdsoftCamera.RegulateTemperature == 0) {
   	Out = "Ok, Cooler power: "+ccdsoftCamera.ThermalElectricCoolerPower+"% Regulation: off"
   } else {
     Reg = (ccdsoftCamera.RegulateTemperature == 1) ? "On":"Off"
   	Out = "Nope!, Cooler power: "+ccdsoftCamera.ThermalElectricCoolerPower+"% Regulation: "+Reg
   }`

// ParkAndDisconnectCommand javascript function
const ParkAndDisconnectCommand = `
   /* Java Script */
   var Out;
   if (sky6RASCOMTele.IsParked()) {
     sky6RASCOMTele.Unpark();
     sky6RASCOMTele.Park();
   } else {
     sky6RASCOMTele.Park();
   }
   Out = sky6RASCOMTele.IsConnected;`

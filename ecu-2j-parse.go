package main

import (
	"fmt"
)
// todo - all the point values are being lost

func twojParseResponse(actualData []byte) []byte {

  globalDataOutputLock.Lock()
  defer globalDataOutputLock.Unlock()

	if slicesEqual(actualData, twojWokeResponse) {
		fmt.Println("< ECU woke up")
		globalConnected = true
		return twojWokeResponse
	}
	if slicesEqual(actualData, twojStartDiagResponse) {
		fmt.Println("< Diag mode accepted")
		return twojStartDiagResponse
	}
	if slicesEqual(actualData[0:2], twojSeedResponse) {
		fmt.Println("< seed")
		twojSeed = int(actualData[2]) << 8
		twojSeed += int(actualData[3])
		// do key generation
		twojKey = generateKey(twojSeed)
		return twojSeedResponse
	}
	if slicesEqual(actualData, twojKeyAcceptResponse) {
		fmt.Println("< Key accepted")
		return twojKeyAcceptResponse
	}
	if slicesEqual(actualData, twojPongResponse) {
		fmt.Println("< PONG")
		return twojPongResponse
	}
	if slicesEqual(actualData, twojFaultsClearedResponse) {
		fmt.Println("< FAULT CLEARED")
		globalAlert = "ECU reports faults cleared"
		return twojFaultsClearedResponse
	}
	if slicesEqual(actualData[0:2], twojFaultsResponse) {
		fmt.Println("< Faults")
		twojParseFaults(actualData)
		return twojFaultsResponse
	}
	if slicesEqual(actualData[0:2], twojResponseData00) {
		fmt.Println("got data packet 00")
		// don't care?
		return twojResponseData00
	}
	if slicesEqual(actualData[0:2], twojResponseData01) {
		fmt.Println("got data packet 01")
		coolant := int(actualData[2]) << 8
		coolant += int(actualData[3])
		coolantFloat := float32(coolant) - 2732
		coolantFloat /= 10
		globalDataOutput["coolant_temp"] = coolantFloat
		return twojResponseData01
	}
	if slicesEqual(actualData[0:2], twojResponseData02) {
		fmt.Println("got data packet 02")
		oiltemp := int(actualData[2]) << 8
		oiltemp += int(actualData[3])
		oiltempFloat := float32(oiltemp) - 2732
		oiltempFloat /= 10
		globalDataOutput["oil_temp"] = oiltempFloat
		return twojResponseData02
	}
	if slicesEqual(actualData[0:2], twojResponseData03) {
		fmt.Println("got data packet 03")
		iat := int(actualData[2]) << 8
		iat += int(actualData[3])
		iatFloat := float32(iat) - 2732
		iatFloat /= 10
		globalDataOutput["intake_air_temp"] = iatFloat
		return twojResponseData03
	}
	if slicesEqual(actualData[0:2], twojResponseData05) {
		fmt.Println("got data packet 05")
		fueltemp := int(actualData[2]) << 8
		fueltemp += int(actualData[3])
		globalDataOutput["fuel_temp"] = float32(fueltemp)
		return twojResponseData05
	}
	if slicesEqual(actualData[0:2], twojResponseData06) {
		fmt.Println("got data packet 06")
		// don't care?
		return twojResponseData06
	}
	if slicesEqual(actualData[0:2], twojResponseData07) {
		fmt.Println("got data packet 07")
		mapkpa := int(actualData[2]) << 8
		mapkpa += int(actualData[3])
		globalDataOutput["map_sensor_kpa"] = float32(mapkpa)/100
		return twojResponseData07
	}
	if slicesEqual(actualData[0:2], twojResponseData08) {
		fmt.Println("got data packet 08")
		tps := int(actualData[2]) << 8
		tps += int(actualData[3])
		tpsFloat := float32(tps) / 100
		globalDataOutput["tps_degrees"] = tpsFloat
		return twojResponseData08
	}
	if slicesEqual(actualData[0:2], twojResponseData09) {
		fmt.Println("got data packet 09")
		rpm := int(actualData[2]) << 8
		rpm += int(actualData[3])
		globalDataOutput["rpm"] = float32(rpm)
		return twojResponseData09
	}
	if slicesEqual(actualData[0:2], twojResponseData0A) {
		fmt.Println("got data packet 0A")
		feedback := int(actualData[2]) << 8
		feedback += int(actualData[3])
		feedbackFloat := float32(feedback) / 100
		o2mv := int(actualData[4]) << 8
		o2mv += int(actualData[5])
		airFuel := ((float32(o2mv) / 1000) * 2) + 10
		globalDataOutput["fuelling_feedback_percent"] = feedbackFloat
		globalDataOutput["o2_mv"] = float32(o2mv)
		globalDataOutput["estimate_air_fuel"] = airFuel
		return twojResponseData0A
	}
	if slicesEqual(actualData[0:2], twojResponseData0B) {
		fmt.Println("got data packet 0B")
		globalDataOutput["coil_1_charge_time"] = float32(actualData[2]) / 1000
		globalDataOutput["coil_2_charge_time"] = float32(actualData[3]) / 1000
		return twojResponseData0B
	}
	if slicesEqual(actualData[0:2], twojResponseData0C) {
		fmt.Println("got data packet 0C")
		globalDataOutput["injector_1_pw"] = float32(actualData[2])
		globalDataOutput["injector_2_pw"] = float32(actualData[3])
		globalDataOutput["injector_3_pw"] = float32(actualData[4])
		globalDataOutput["injector_4_pw"] = float32(actualData[5])
		return twojResponseData0C
	}
	if slicesEqual(actualData[0:2], twojResponseData0D) {
		fmt.Println("got data packet 0D")
		globalDataOutput["vehicle_speed"] = float32(actualData[2])
		return twojResponseData0D
	}
	if slicesEqual(actualData[0:2], twojResponseData0F) {
		fmt.Println("got data packet 0F")
		globalDataOutput["throttle_switch"] = float32(int(actualData[2]) & 1) // 0b00000001
		globalDataOutput["ignition"] = float32((int(actualData[2]) >> 1) & 1) // 0b00000010
		globalDataOutput["ac_button"] = float32((int(actualData[2]) >> 3) & 1) // 0b00001000
		return twojResponseData0F
	}
	if slicesEqual(actualData[0:2], twojResponseData10) {
		fmt.Println("got data packet 10")
		battery := int(actualData[4]) << 8
		battery += int(actualData[5])
		batteryFloat := float32(battery) / 1000
		globalDataOutput["battery_voltage"] = batteryFloat
		return twojResponseData10
	}
	if slicesEqual(actualData[0:2], twojResponseData11) {
		fmt.Println("got data packet 11")
		// 0 means OK, 1 bad
		// will swap for our purposes
		// output is 1 for yes
		primaryTriggerSync := actualData[2] & 1 // 0b00000001
		secondaryTriggerSync := (actualData[2] >> 1) & 1 //0b00000010
		globalDataOutput["primary_trigger_sync"] = float32(1 - primaryTriggerSync)
		globalDataOutput["secondary_trigger_sync"] = float32(1 - secondaryTriggerSync)
		return twojResponseData11
	}
	if slicesEqual(actualData[0:2], twojResponseData12) {
		fmt.Println("got data packet 12")
		idleValvePos := int(actualData[2]) << 8
		idleValvePos += int(actualData[3])
		idleValveFloat := float32(idleValvePos) / 2
		globalDataOutput["idle_valve_pos"] = idleValveFloat
		return twojResponseData12
	}
	if slicesEqual(actualData[0:2], twojResponseData13) {
		fmt.Println("got data packet 13")
		globalDataOutput["closed_loop"] = float32(actualData[2] & 0b00000001)
		return twojResponseData13
	}
	if slicesEqual(actualData[0:2], twojResponseData21) {
		fmt.Println("got data packet 21")
		rpmError := int(actualData[2]) << 8
		rpmError += int(actualData[3])
		if rpmError > 32768 {
			rpmError -= 65535
		}
		globalDataOutput["rpm_error"] = float32(rpmError)
		return twojResponseData21
	}
	if slicesEqual(actualData[0:2], twojResponseData25) {
		fmt.Println("got data packet 25")
		camPercent := int(actualData[2]) << 8
		camPercent += int(actualData[3])
		globalDataOutput["cam_percent"] = float32(camPercent)
		return twojResponseData25
	}
	if slicesEqual(actualData[0:2], twojResponseData3A) {
		fmt.Println("got data packet 3A")
		idleTimingOffset := int(actualData[2]) << 8
		idleTimingOffset += int(actualData[3])
		idleTimingOffsetFloat := float32(idleTimingOffset) / 10
		idleAdjusterRpm := int(actualData[4]) << 8
		idleAdjusterRpm += int(actualData[5])
		globalDataOutput["idle_timing_offset"] = idleTimingOffsetFloat
		globalDataOutput["idle_adjuster_rpm"] = float32(idleAdjusterRpm)
		return twojResponseData3A
	}

	// if we get here then something is wrong with the data
	// todo: cope with 7F (fail/no)
  return twojPongResponse
}
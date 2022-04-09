<template>
  <div>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="openSerialPort()"
    >
      Open Serial Port
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="newInit()"
    >
      New Init
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="download()"
    >
      <i class="fas fa-download"></i>
      Download
    </button>
    <br />
        <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x80,0x7d])"
    >All Data</button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x80])"
    >
      Data 80
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x7d])"
    >
      Data 7D
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0xd1])"
    >
      ECU ID
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0xf0])"
    >
      Diagnostic mode Read
    </button>

    <br />
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x1d])"
    >
      Fan On
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x0d])"
    >
      Fan OFF
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x91])"
    >
      Idle +
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x92])"
    >
      Idle -
    </button>

    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x7b])"
    >
      Fuel Trim +
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x7a])"
    >
      Fuel Trim -
    </button>

    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x93])"
    >
      Ign Trim +
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0x94])"
    >
      Ign Trim -
    </button>

    <br />
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0xc4])"
    >
      Diagmode 4?
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0xf7])"
    >
      Diagmode 4 - read calibration
    </button>
  </div>

  <hr />

  <pre id="console"></pre>
</template>

<script lang="ts">
export default {
  name: "app",
  data() {
    return {
      port: null,
      ECUID: "",
      ECUSerial: "",
      log: {
        Name: "run.fcr",
        Count: 338,
        Date: "0000-01-01T12:35:55.186Z",
        Summary: "Test run",
        ECUID: "",
        ECUSerial: "",
        MemsData: [
          {
            Time: "12:35:55.186",
            Dataframe7d:
              "7d201010ff92401cffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:55.729",
            Dataframe7d:
              "7d201010ff92401cffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
        ],
      },
      Dataframe: {
        Time: "00:00:00.000",
        EngineRPM: 0,
        CoolantTemp: 0,
        AmbientTemp: 0,
        IntakeAirTemp: 0,
        FuelTemp: 0,
        ManifoldAbsolutePressure: 0,
        BatteryVoltage: 0,
        ThrottlePotSensor: 0.0,
        ThrottlePosition: 0,
        IdleSwitch: false,
        AirconSwitch: false,
        ParkNeutralSwitch: false,
        DTC0: 0,
        DTC1: 0,
        IdleSetPoint: 0,
        IdleHot: 0,
        Uk8011: 0,
        IACPosition: 0,
        IdleSpeedDeviation: 0,
        IgnitionAdvanceOffset80: 0,
        IgnitionAdvance: 0,
        CoilTime: 0,
        CrankshaftPositionSensor: 0,
        Uk801a: 0,
        Uk801b: 0,
        IgnitionSwitch: true,
        ThrottleAngle: 0,
        Uk7d03: 0,
        AirFuelRatio: 0,
        DTC2: 0,
        LambdaVoltage: 0,
        LambdaFrequency: 0,
        LambdaDutycycle: 0,
        LambdaStatus: 0,
        ClosedLoop: false,
        LongTermFuelTrim: 0,
        ShortTermFuelTrim: 0,
        FuelTrimCorrection: 0,
        CarbonCanisterPurgeValve: 0,
        DTC3: 0,
        IdleBasePosition: 0,
        Uk7d10: 0,
        DTC4: 0,
        IgnitionAdvanceOffset7d: 0,
        IdleSpeedOffset: 0,
        Uk7d14: 0,
        Uk7d15: 0,
        DTC5: 0,
        Uk7d17: 0,
        Uk7d18: 0,
        Uk7d19: 0,
        Uk7d1a: 0,
        Uk7d1b: 0,
        Uk7d1c: 0,
        Uk7d1d: 0,
        Uk7d1e: 0,
        JackCount: 0,
        CoolantTempSensorFault: false,
        IntakeAirTempSensorFault: false,
        FuelPumpCircuitFault: false,
        ThrottlePotCircuitFault: false,
        Analytics: {
          ReadingFault: false,
          IsEngineRunning: false,
          IsEngineWarming: false,
          IsAtOperatingTemp: false,
          IsEngineIdle: false,
          IsEngineIdleFault: false,
          IdleSpeedFault: false,
          IdleErrorFault: false,
          IdleHotFault: false,
          IdleBaseFault: false,
          IsCruising: false,
          IsClosedLoop: false,
          IsClosedLoopExpected: false,
          ClosedLoopFault: false,
          IsThrottleActive: false,
          MapFault: false,
          VacuumFault: false,
          IdleAirControlFault: false,
          IdleAirControlRangeFault: false,
          IdleAirControlJackFault: false,
          O2SystemFault: false,
          LambdaRangeFault: false,
          LambdaOscillationFault: false,
          ThermostatFault: false,
          CoolantTempSensorFault: false,
          IntakeAirTempSensorFault: false,
          FuelPumpCircuitFault: false,
          ThrottlePotCircuitFault: false,
          CrankshaftSensorFault: false,
          CoilFault: false,
          IACPosition: 0,
        },
        //Dataframe80:          "801c000085ff4fff638e23001001000000208b60039d003808c1000000",
        //Dataframe7d:          "7d201012ff92006effff0100996400ff3affff30807c63ff19401ec0264034c008",
        
   "Dataframe7d": "7d201024ff924027ffff0101796200ff69ffff358883a7ff134015801a0029c02a",
   "Dataframe80": "801c089f74ff51ff3983320800010000203787870379055c05f8100000"
  },
    };
  },
  mounted: function() {
    console.log("mounted");
    this.parse80(this.hexToBytes(this.Dataframe.Dataframe80));
    this.parse7d(this.hexToBytes(this.Dataframe.Dataframe7d));
  },
  methods: {
    parse7d( data: ArrayBuffer ) {
      var v = new DataView(data);
      let len=v.getUint8(1);
      if ( len !== 32) {
        this.debug(" expected len 32 for 0x7d")
      } else 
      {
        let v7d = {
          IgnitionSwitch: true,
        ThrottleAngle: 0,
        Uk7d03: v.getUint8(3),
        AirFuelRatio: 0,
        DTC2: 0,
        LambdaVoltage: 0,
        LambdaFrequency: 0,
        LambdaDutycycle: 0,
        LambdaStatus: 0,
        ClosedLoop: false,
        LongTermFuelTrim: 0,
        ShortTermFuelTrim: 0,
        FuelTrimCorrection: 0,
        CarbonCanisterPurgeValve: 0,
        DTC3: v.getUint8(16),
        IdleBasePosition: 0,
        Uk7d10: v.getUint8(16),
        DTC4: 0,
        IgnitionAdvanceOffset7d: 0,
        IdleSpeedOffset: 0,
        Uk7d14: v.getUint8(20),
        Uk7d15: v.getUint8(21),
        DTC5: v.getUint8(22),
        Uk7d17: v.getUint8(23),
        Uk7d18: v.getUint8(24),
        Uk7d19: v.getUint8(25),
        Uk7d1a: v.getUint8(26),
        Uk7d1b: v.getUint8(27),
        Uk7d1c: v.getUint8(28),
        Uk7d1d: v.getUint8(29),
        Uk7d1e: v.getUint8(30),
        JackCount: v.getUint8(31)
        }
        console.log(v7d);
      }
    },
    parse80( data: ArrayBuffer ) {
      var v = new DataView(data);
      let len=v.getUint8(1);
      if ( len !== 28) {
        this.debug(" expected len 28 for 0x80")
      } else 
      {
        
        let v80 = {
          EngineRPM: v.getInt16(2),
        CoolantTemp: v.getUint8(3) - 55.0,
        AmbientTemp: v.getUint8(4) -55.0,
        IntakeAirTemp: v.getUint8(5) - 55.0,
        FuelTemp: v.getUint8(6) - 55.0,
        ManifoldAbsolutePressure: v.getUint8(7),
        BatteryVoltage: v.getUint8(8) /10.0,
        ThrottlePotSensor: v.getUint8(9) / 200,
        ThrottlePosition: 0,
        IdleSwitch: false,
        AirconSwitch: false,
        ParkNeutralSwitch: v.getUint8(12),
        DTC0: v.getUint8(13),
        DTC1: v.getUint8(14),
        IdleSetPoint: v.getUint8(15)*6.1,
        IdleHot: v.getUint8(16),
        Uk8011: v.getUint8(17),
        IACPosition: v.getUint8(18),//?
//19?
        IdleSpeedDeviation: v.getInt16(20),
        IgnitionAdvanceOffset80: v.getInt8(22),
        IgnitionAdvance: v.getInt8(24)/2.0-24,
        CoilTime: v.getInt16(24) * 2,
        CrankshaftPositionSensor: v.getUint8(25), // 25?????
        Uk801a: v.getInt16(26),
        Uk801b: v.getInt16(27)
        }
        console.log(v80);

      }
    },
    hexToBytes(hex: string) {
    let bytes=new ArrayBuffer(hex.length/2)
    var view = new DataView(bytes);
    for (let c = 0; c < hex.length; c += 2)
        view.setUint8(c/2,parseInt(hex.substr(c, 2), 16));
    return bytes;
},
    async sleep(ms) {
      await new Promise((resolve) => setTimeout(resolve, ms));
    },
    async newInit() {
      let ecuAddress = 0x16;

      await this.port.setSignals({ break: false });
      this.debug("sleeping for 2 seconds to clear the line");
      await sleep(2000);
      await this.port.setSignals({ break: true });
      await sleep(200);

      for (var i = 0; i < 8; i++) {
        bit = (ecuAddress >> i) & 1;
        this.debug(i + " " + bit);

        if (bit > 0) {
          await this.port.setSignals({ break: false });
        } else {
          await this.port.setSignals({ break: true });
        }
        await sleep(200);
        // await sleep(195);
      }
      // stop bit
      await this.port.setSignals({ break: false });
      await sleep(200);

      this.debug("continuing with normal init");
      //start=2;
      //sendToEcu([0x7c,0xCA,0x75,0xd0,0x80]);
      let reply = 0x83;
      let send = reply ^ 0xff;

      this.debug("would send " + this.hex([send]));
    },
    async sendToEcu(bytes) {
      this.debug(">> " + this.hex(bytes));
      let writer = this.port.writable.getWriter();
      writer.write(Uint8Array.from(bytes));
      writer.releaseLock();
    },
    hex(bytes) {
      return bytes.map((x) => x.toString(16).padStart(2, "0")).join(" ");
    },
    debug(msg) {
      console.log(msg);
      let el = document.getElementById("console");
      el.innerHTML = msg + "\n" + el.innerHTML;
    },
    download() {
      let log = this.log;

      let now = new Date();
      log.Name = `run-${now.getTime()}.fcr`;
      log.Count = log.MemsData.length;
      log.Date = now.toISOString();
      log.ECUSerial = this.ECUSerial;
      log.ECUID = this.ECUID;
      let blob = new Blob([JSON.stringify(log)], { type: "application/json" });
      let url = URL.createObjectURL(blob);
      let a = document.createElement("a");
      a.href = url;
      a.download = log.Name;
      a.click();
    },
    async openSerialPort() {
      this.port = await navigator.serial.requestPort();
      console.log(this.port);
      
      await this.port.open({
        baudRate: 9600,
        databits: 8,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });
      console.log(this.port);
      
      while (this.port.readable) {
        const reader = this.port.readable.getReader();
        this.debug("reading...");
        try {
          while (true) {
            const { value, done } = await reader.read();

            if (done) {
              // |reader| has been canceled.
              trace("read all the data");
              break;
            }
            let data = this.hex(Array.from(value)).substring(2);

            if (value[0] === 0x00) {
              this.debug("got 0");
              //reader.releaseLock();
              //break;
            }
            //Dataframe7d:               "7d201010ff92401cffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            //Dataframe80:              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
            switch( value[0]) {
              case 0x80:
                this.debug("got 80");
                this.Dataframe.Dataframe80=data;
                let now = new Date();
                new Date().getMilliseconds();
                this.Dataframe.Time=`${now.toLocaleTimeString()}.${now.getMilliseconds()}`;
                break;
              case 0x7d:
                this.debug("got 0x7d");
                this.Dataframe.Dataframe7d=data;
                this.log.MemsData.push({
                  Time: this.Dataframe.Time,
                  Dataframe80: this.Dataframe.Dataframe80,
                  Dataframe7d: this.Dataframe.Dataframe7d
                })
                break;
            }
            /*
            << 0d 0d 00
>> 0d
<< 1d 1d 00
>> 1d
<< 30 35 c7 00 05 cb
<< d1 d1 4b 4c 48 33 56 30 30 35 c7 00 05 cb 4b 4c 48 33 56 30 30 35 c7 00 05 cb 4b 4c 48 33 56 30
>> d1
<< 80 80 1c 00 00 4d 9a 50 ff 65 75 22 00 80 01 10 00 18 1c 87 7c 05 78 02 30 0d 16 00 00 00
>> 80
got 0
<< 00 01 00
<< 7d 21 40 0a ff 91 00 57 ff ff 01 00 70 64 00 00 70 ff 00 2e 86 7a 88 00 22 00 22 00 22 00 22
<< 7d
>> 7d
<< 80 80 1c 00 00 4d 9a 50 ff 65 75 22 00 80 01 10 00 18 1c 87 7c 05 78 02 30 0d 16 00 00 00
>> 80 */
          }
        } catch (error) {
          // Handle |error|...
          this.debug("error");
          console.log(error);
        } finally {
          reader.releaseLock();
          this.debug("released lock");
        }
      }
    },
  },
};
</script>

<style lang="scss"></style>

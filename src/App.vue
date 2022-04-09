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
      @click="sendToEcu([0x7d, 0x80])"
    >
      All Data
    </button>
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

  </div>

  <hr />

  <table class="table">
    <tr>
      <th>Name</th>
      <th>Value</th>
    </tr>
    <tr><td>Time</td><td>{{Dataframe.Time}}</td></tr>
    <tr><td>Dataframe7d</td><td>{{Dataframe.Dataframe7d}}</td></tr>
    <tr><td>Dataframe80</td><td>{{Dataframe.Dataframe80}}</td></tr>
    <tr><td>RPM</td><td>{{Dataframe.EngineRPM}}</td></tr>
    <tr><td>AFR</td><td>{{Dataframe.AirFuelRatio}}</td></tr>
    <tr><td>Lambda</td><td>{{Dataframe.LambdaVoltage}}</td></tr>
    <tr><td>Short Term Fuel Trim</td><td>{{Dataframe.ShortTermFuelTrim}}</td></tr>
    <tr><td>ECUID</td><td>{{ECUID}}</td></tr>
    <tr><td>ECUSerial</td><td>{{ECUSerial}}</td></tr>
    
  </table>

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
        Count: 2,
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

        Dataframe7d:
          //"7d201024ff924027ffff0101796200ff69ffff358883a7ff134015801a0029c02a",
            "7d0000000000000000000000000000000000000000000000000000000000000000",
        Dataframe80:
          //"801c089f74ff51ff3983320800010000203787870379055c05f8100000",
          "8000000000000000000000000000000000000000000000000000000000",
      },
    };
  },
  mounted: function () {
    /*
    this.parse80(this.hexToBytes(this.Dataframe.Dataframe80));
    this.parse7D(this.hexToBytes(this.Dataframe.Dataframe7d));
    this.parseD1("D141424E4D5030303399000303");
    this.parseD1(
      "d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630"
    );
    */
  },
  methods: {
    parseD1(data) {
      // d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630
      let bytes = this.hexToBytes(data);
      var v = new DataView(bytes);
      debugger;
      this.ECUSerial = v.getUint32(9).toString();
      this.ECUID = String.fromCharCode.apply(
        null,
        new Uint8Array(bytes.slice(1, 9))
      );
    },
    parse7D(data: ArrayBuffer) {
      var v = new DataView(data);
      let len = v.getUint8(1);
      if (len < 32) { // len is 33
        this.debug(`expected len 32 for 0x7d got ${len}`);
      } else {
        let v7d = {
          IgnitionSwitch: v.getUint8(0x01) > 0,
          ThrottleAngle: (v.getUint8(0x02) * 6.0) / 10.0,
          Uk7d03: v.getUint8(0x03),
          AirFuelRatio: v.getUint8(0x04) / 10.0,
          DTC2: v.getUint8(0x05),
          LambdaVoltage: v.getUint8(0x06) * 5,
          LambdaFrequency: v.getUint8(0x07),
          LambdaDutycycle: v.getUint8(0x08),
          LambdaStatus: v.getUint8(0x09),
          ClosedLoop: v.getUint8(0x0a),
          LongTermFuelTrim: v.getInt8(0x0b),
          ShortTermFuelTrim: v.getInt8(0x0c),
          //FuelTrimCorrection: v.getInt8(0x),
          CarbonCanisterPurgeValve: v.getUint8(0x0d),
          DTC3: v.getUint8(0x0e),
          IdleBasePosition: v.getUint8(0x0f),
          Uk7d10: v.getUint8(0x10),
          DTC4: v.getUint8(0x11),
          IgnitionAdvanceOffset7d: v.getUint8(0x12) - 48,
          IdleSpeedOffset: v.getUint8(0x13),
          Uk7d14: v.getUint8(0x14),
          Uk7d15: v.getUint8(0x15),
          DTC5: v.getUint8(0x16),
          Uk7d17: v.getUint8(0x17),
          Uk7d18: v.getUint8(0x18),
          Uk7d19: v.getUint8(0x19),
          Uk7d1a: v.getUint8(0x1a),
          Uk7d1b: v.getUint8(0x1b),
          Uk7d1c: v.getUint8(0x1c),
          Uk7d1d: v.getUint8(0x1d),
          //Uk7d1e: v.getUint8(0x1e),
          //JackCount: v.getUint8(0x1f),
        };
        console.log(v7d);
        Object.assign(this.Dataframe,v7d);
      }
    },
    parse80(data: ArrayBuffer) {
      var v = new DataView(data);
      let len = v.getUint8(1);
      if (len !== 28) {
        this.debug(" expected len 28 for 0x80");
      } else {
        let v80 = {
          EngineRPM: v.getInt16(0x02),
          CoolantTemp: v.getUint8(0x03) - 55.0,
          AmbientTemp: v.getUint8(0x04) - 55.0,
          IntakeAirTemp: v.getUint8(0x05) - 55.0,
          FuelTemp: v.getUint8(0x06) - 55.0,
          ManifoldAbsolutePressure: v.getUint8(0x07),
          BatteryVoltage: v.getUint8(0x08) / 10.0,
          ThrottlePotSensor: v.getUint8(0x09) * 0.02,
          ThrottlePosition: v.getUint8(0x09) * 0.02, // ?
          IdleSwitch: v.getUint8(0x0a),
          AirconSwitch: v.getUint8(0x0b), // ???
          ParkNeutralSwitch: v.getUint8(0x0c),
          DTC0: v.getUint8(0x0d),
          DTC1: v.getUint8(0x0e),
          IdleSetPoint: v.getUint8(0x0f) * 6.1,
          IdleHot: v.getUint8(0x10),
          Uk8011: v.getUint8(0x11),
          IACPosition: v.getUint8(0x12),
          IdleSpeedDeviation: v.getInt16(0x13),
          IgnitionAdvanceOffset80: v.getInt8(0x15),
          IgnitionAdvance: v.getInt8(0x16) / 2.0 - 24.0,
          CoilTime: v.getUint16(0x17) * 0.0002,
          CrankshaftPositionSensor: v.getUint8(0x19),
          Uk801a: v.getUint8(0x1a),
          Uk801b: v.getUint8(0x1b),
        };
        console.log(v80);
        Object.assign(this.Dataframe,v80);
      }
    },
    hexToBytes(hex: string) {
      let bytes = new ArrayBuffer(hex.length / 2);
      var view = new DataView(bytes);
      for (let c = 0; c < hex.length; c += 2)
        view.setUint8(c / 2, parseInt(hex.substr(c, 2), 16));
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
      //this.sendToEcu([0x7c,0xCA,0x75,0xd0,0x80]);
      let reply = 0x83;
      let send = reply ^ 0xff;

      this.debug("would send " + this.hex([send]));
    },
    async sendToEcu(bytes) {
      this.debug(">> " + this.hex(bytes),' ');
      let writer = this.port.writable.getWriter();
      writer.write(Uint8Array.from(bytes));
      writer.releaseLock();
    },
    hex(bytes,delim='') {
      return bytes.map((x) => x.toString(16).padStart(2, "0")).join(delim);
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
      log.MemsData=[];
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

      this.sendToEcu([0xd1]);

      while (this.port.readable) {
        const reader = this.port.readable.getReader();
        this.debug("reading...");
        try {
          while (true) {
            const { value, done } = await reader.read();

            if (done) {
              trace("read all the data");
              break;
            }
            let data = this.hex(Array.from(value)).substring(2);
            this.debug(data);

            switch (value[0]) {
              case 0x80:
                this.debug("got 80");
                this.Dataframe.Dataframe80 = data;
                this.log.MemsData.push({
                  Time: this.Dataframe.Time,
                  Dataframe80: this.Dataframe.Dataframe80,
                  Dataframe7d: this.Dataframe.Dataframe7d,
                });
                this.parse80(this.hexToBytes(data));
                this.sendToEcu([0x7d]); // trigger next frame
                break;
              case 0x7d:
                this.debug("got 0x7d");
                this.Dataframe.Dataframe7d = data;
                let now = new Date();
                new Date().getMilliseconds();
                this.Dataframe.Time = `${now.toLocaleTimeString()}.${now.getMilliseconds()}`;
                this.parse7D(this.hexToBytes(data));
                this.sendToEcu([0x80]); // trigger next frame
                break;
              case 0xd1:
                this.parseD1(data);
                break;
              default: {
                this.debug("got " + value[0]);
              }
            }
          }
        } catch (error) {
          // Handle |error|...
          this.debug(`error: ${error.message}`);
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

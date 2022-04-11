<script lang="ts">
export default {
  name: "app",
  data() {
    return {
      ser: {
        mode: 0,
        port: null,
        buffer: "",
      },
      ECUID: "",
      ECUSerial: "",
      log: {
        Name: "run.fcr",
        Count: 2,
        Date: "0000-01-01T12:35:55.186Z",
        Summary: "Test run",
        ECUID: "",
        ECUSerial: "",
        MemsData: [],
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

        //Dataframe7d:
         // "7d201016ff92401bffff010179670cff51ffff3588840dff134015801a0029c02a",
        //Dataframe80:
          //"801c08ba80ff56ff1f842308000100002037876b0417055c05df100000",
        Dataframe80: "801c000000000000000000000000000000000000000000000000000000",
        Dataframe7d: "7d2000000000000000000000000000000000000000000000000000000000000000"
      },
      parameters: [
        "EngineRPM",
        "CoolantTemp",
        "AmbientTemp",
        "IntakeAirTemp",
        "FuelTemp",
        "ManifoldAbsolutePressure",
        "BatteryVoltage",
        "ThrottlePotSensor",
        "ThrottlePosition",
        "IdleSwitch",
        "AirconSwitch",
        "ParkNeutralSwitch",
        "DTC0",
        "DTC1",
        "IdleSetPoint",
        "IdleHot",
        "Uk8011",
        "IACPosition",
        "IdleSpeedDeviation",
        "IgnitionAdvanceOffset80",
        "IgnitionAdvance",
        "CoilTime",
        "CrankshaftPositionSensor",
        "Uk801a",
        "Uk801b",
        "IgnitionSwitch",
        "ThrottleAngle",
        "Uk7d03",
        "AirFuelRatio",
        "DTC2",
        "LambdaVoltage",
        "LambdaFrequency",
        "LambdaDutycycle",
        "LambdaStatus",
        "ClosedLoop",
        "LongTermFuelTrim",
        "ShortTermFuelTrim",
        "FuelTrimCorrection",
        "CarbonCanisterPurgeValve",
        "DTC3",
        "IdleBasePosition",
        "Uk7d10",
        "DTC4",
        "IgnitionAdvanceOffset7d",
        "IdleSpeedOffset",
        "Uk7d14",
        "Uk7d15",
        "DTC5",
        "Uk7d17",
        "Uk7d18",
        "Uk7d19",
        "Uk7d1a",
        "Uk7d1b",
        "Uk7d1c",
        "Uk7d1d",
        "Uk7d1e",
        "JackCount",
      ],
      debug_log: [],
    };
  },
  mounted: function () {
    //this.parse80(this.hexToBytes(this.Dataframe.Dataframe80));
    //this.parse7D(this.hexToBytes(this.Dataframe.Dataframe7d));
    //this.parseD1("D141424E4D5030303399000303");

    //this.parseD1(
      //"d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630"
      // 3035c70005cb
    //);
  },
  methods: {
    parseD1(data) {
      // d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630 1.9 response longer...
      // _ KLH3V005 ___ KLH3V005 ___ KLH3V0
      let bytes = this.hexToBytes(data);
      var v = new DataView(bytes);
      this.ECUSerial = v.getUint32(9).toString();
      this.ECUID = String.fromCharCode.apply(
        null,
        new Uint8Array(bytes.slice(1, 9))
      );
    },
    parse7D(data: ArrayBuffer) {
      var v = new DataView(data);
      let len = v.getUint8(1);
      if (len < 32) {
        // len is 33 mems 1.9
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
          Uk7d1e: v.getUint8(0x1e),
          JackCount: v.getUint8(0x1f),
          //Uk7d20: v.getUint8(0x20), Mems 1.9
        };
        Object.assign(this.Dataframe, v7d);
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
          AirconSwitch: v.getUint8(0x0b), 
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
        Object.assign(this.Dataframe, v80);
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

    async wait(ms) {
      const start = performance.now();
      ms++;
      while (performance.now() - start < ms) {
        await new Promise((resolve) => setTimeout(resolve, 1));
        let delta = performance.now() - start;
        if (ms - delta < 10) {
          while (performance.now() - start < ms) {}
          return;
        }
      }
      //console.log({ms:ms,actual:performance.now() - start});
    },

    async sendToEcu(bytes) {
      //this.debug(">> " + this.hex(bytes), " ");
      let writer = this.ser.port.writable.getWriter();
      writer.write(Uint8Array.from(bytes));
      writer.releaseLock();
    },
    hex(bytes, delim = "") {
      return bytes.map((x) => x.toString(16).padStart(2, "0")).join(delim);
    },
    debug(msg) {
      console.log(msg);
      this.debug_log.push(msg);
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
      log.MemsData = [];
    },
    async closeSerialPort() {
      if (this.ser.reader) {
        await this.ser.reader.cancel();
      }
      await this.ser.port.close();
      this.ser.reader = null;
      this.ser.port = null;
      this.ser.mode = 0;
    },

    async baud5() {
      this.debug("set reader");
      while (this.ser.port.readable) {
        this.ser.reader = this.ser.port.readable.getReader();
        this.debug("waiting on data...");
        let start;
        try {
          while (true) {
            const { value, done } = await this.ser.reader.read();

            if (done) {
              this.debug("serial this.ser.reader done");
              return;
            }
            console.log(value);
            this.ser.buffer = this.ser.buffer.concat(
              this.hex(Array.from(value))
            );
            //this.debug(`l: ${read.length} d: ${read} v: ${value}`);
            //this.debug( `<< ${read}`);
            if (start === null) start = value[0];
            switch (start) {
              default: {
                //read=read.substring(2);
                this.debug("default:", this.ser.buffer);
                this.debug(`<< ${this.ser.buffer}`);

                start = null;
                this.ser.buffer = "";
              }
            }
          }
        } catch (error) {
          this.debug(`error: ${error.message}`);
          console.log(error);
          throw error;
        } finally {
          this.ser.reader.releaseLock();
          this.ser.reader = null;
          this.debug("released lock");
        }
      }
    },
    async openSerialPort() {
      //  let ports = await navigator.serial.getPorts();
      //    console.log(ports);
      //      let info = ports[0].getInfo();

      this.ser.port = await navigator.serial.requestPort();
      console.log(this.ser.port.getInfo());

      await this.ser.port.open({
        baudRate: 9600,
        databits: 8,
        bufferSize: 128,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });
      console.log(this.ser.port);

      if (1) this.sendToEcu([0xd0]);

      if (0)
        this.timer = setInterval(() => {
          this.sendToEcu([0x7d]);
        }, 2000);

      let read = "";
      let start = null;
      while (this.ser.port.readable) {
        this.ser.reader = this.ser.port.readable.getReader();
        this.debug("waiting on data...");

        try {
          while (true) {
            const { value, done } = await this.ser.reader.read();

            if (done) {
              this.debug("serial this.ser.reader done");
              return;
            }

            this.ser.buffer = this.ser.buffer.concat(
              this.hex(Array.from(value))
            );
            //this.debug(`l: ${read.length} d: ${read} v: ${value}`);
            //this.debug( `<< ${read}`);
            if (start === null) start = value[0];
            switch (start) {
              case 0x80:
                if (this.ser.buffer.length < 56) {
                  //this.debug() `expected 56 (${this.ser.buffer.length}) bytes for 0x80`   );
                  break;
                }

                this.ser.buffer = this.ser.buffer.substring(2);
                this.Dataframe.Dataframe80 = this.ser.buffer;
                this.log.MemsData.push({
                  Time: this.Dataframe.Time,
                  Dataframe80: this.Dataframe.Dataframe80,
                  Dataframe7d: this.Dataframe.Dataframe7d.substring(0, 66),
                });
                this.parse80(this.hexToBytes(this.ser.buffer));
                //this.sendToEcu([0x7d]); // trigger next frame
                this.ser.buffer = "";
                start = null;
                break;
              case 0x00: {
                this.debug(`<< ${this.ser.buffer}`);
                this.ser.buffer = "";
                start = null;
                break;
              }

              case 0x55: {
                //let send = start ^ 0xff;

                this.debug(
                  `1.9 ECU woke up - init stage 1, <<${this.ser.buffer} >> ca`
                );
                this.sendToEcu([0xca]);
                await this.sleep(100);
                this.ser.buffer = "";
                start = null;
                break;
              }

              case 0xca: {
                this.debug(`Stage 2 << ${this.ser.buffer} >> 75`);
                this.sendToEcu([0x75]);

                this.ser.buffer = "";
                start = null;
                break;
              }

              case 0x75: {
                this.debug(`Stage 3 << ${this.ser.buffer} >> F4`);
                this.sendToEcu([0xf4]);
                this.ser.buffer = "";
                start = null;
                break;
              }
              case 0xf4: {
                this.debug(`Stage 4 ${this.ser.buffer}`);
                this.sendToEcu([0xd0]);
                this.ser.buffer = "";
                start = null;
                break;
              }
              case 0xd0: {
                this.debug(`Got ID ${this.ser.buffer}`);
                this.debug(this.ser.buffer);
                this.ser.buffer = "";
                start = null;
                break;
              }
              case 0x7d:
                if (this.ser.buffer.length < 66) {
                  //this.debug(  `expected 64 (${this.ser.buffer.length}) bytes for 0x7d` );
                  break;
                }
                this.ser.buffer = this.ser.buffer.substring(2);
                //this.debug(                    `got (${this.ser.buffer.length}) bytes for 0x7d`);

                this.Dataframe.Dataframe7d = this.ser.buffer;
                let now = new Date();
                new Date().getMilliseconds();
                this.Dataframe.Time = `${now.toLocaleTimeString()}.${now.getMilliseconds()}`;
                this.parse7D(this.hexToBytes(this.ser.buffer));
                this.sendToEcu([0x80]); // trigger next frame
                this.ser.buffer = "";
                start = null;
                break;
              case 0xd1:
                if (this.ser.buffer.length < 64) {
                  this.debug(
                    `expected 64 bytes for 0xd1, got ${this.ser.buffer.length}`
                  );
                  if(this.ser.buffer.length==2) {
                this.ser.buffer = "";
                start = null;
 // gone wrong - reset    
                  }
                  break;
                }
                this.ser.buffer = this.ser.buffer.substring(2);
                this.parseD1(this.ser.buffer);
                this.ser.buffer = "";
                start = null;
                break;

              default: {
                //read=read.substring(2);
                this.debug("default:", this.ser.buffer);
                this.debug(`<< ${this.ser.buffer}`);

                start = null;
                this.ser.buffer = "";
              }
            }
          }
        } catch (error) {
          this.debug(`error: ${error.message}`);
          console.log(error);
        } finally {
          this.ser.reader.releaseLock();
          this.ser.reader = null;
          this.debug("released lock");
        }
      }
    },
    async newInit() {
      this.ser.port = await navigator.serial.requestPort();
      console.log(this.ser.port.getInfo());

      await this.ser.port.open({
        baudRate: 9600,
        databits: 8,
        bufferSize: 128,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });

      let ecuAddress = 0x16;
      this.debug(
        `Connecting to MEMS 1.9 ECU at address ${ecuAddress.toString(16)}`
      );

      //this.sendToEcu([0xca]);
      //this.sendToEcu([0x055]);

      await this.ser.port.setSignals({ break: false });

      setTimeout(() => this.baud5(), 0);

      this.debug("sleeping for 2 seconds to clear the line...");
      await this.wait(2000);
      let sleepMs = 200;

      if (1) {
        let times = [];
        //      await this.ser.port.setSignals({ break: true });

        let i = 0;

        ecuAddress = (ecuAddress << 1) | 1;

        let bits = ecuAddress.toString(2).padStart(10, 0).split("").reverse();
        let sleepMs = 200;
        let b = [];
        /*
      start = performance.now();
      let timer = setInterval(() => {
        this.ser.port.setSignals({ break: !bits[i] });
        times.push(performance.now() - start);
        i = i + 1;
        if (i === 10) clearInterval(timer);
      }, sleepMs);
      console.log(times);
      */
        let start = performance.now();
        let last = 0;
        for (let i = 0; i < 10; i++) {
          await this.ser.port.setSignals({ break: !bits[i] });
          let now = performance.now();
          times.push(now - start);
          start = now;
          //b.push(bits[i]);
          await this.wait(sleepMs);
        }
        console.log(times);
        console.log(b);
      }
      if (0) {
        await this.ser.port.setSignals({ break: true });
        await this.wait(sleepMs);
        await this.ser.port.setSignals({ break: false });
        await this.wait(sleepMs);
        await this.ser.port.setSignals({ break: true });
        await this.wait(sleepMs * 2);
        await this.ser.port.setSignals({ break: false });
        await this.wait(sleepMs);
        await this.ser.port.setSignals({ break: true });
        await this.wait(sleepMs);
        await this.ser.port.setSignals({ break: false });
        await this.wait(sleepMs * 4);
      }

      await this.sleep(100);

      this.debug("continuing with normal init");
      //00000000  55 76 83
      // expectiNg 0x55, 0x76, 0x83
      //this.sendToEcu([0x7c]);
      //Expect 0x7c, 0xE9
    },
  },
};
</script>

/* break 1 break 0 break 1 break 1 break 0 break 1 break 0 break 0 break 0 break
0 Connecting to MEMS 1.9 ECU Serial cable set to: 9600,8n1,none had buffer data:
got 0 bytes 1.9 ECU woke up - init stage 1 1.9 had buffer data: got 1 bytes
00000000 7c ||| Connecting to MEMS 1.9 ECU Serial cable set to: 9600,8n1,none
had buffer data: got 0 bytes 1.9 ECU woke up - init stage 1 1.9 had buffer data:
got 1 bytes 00000000 7c ||| Connecting to MEMS 1.9 ECU Serial cable set to:
9600,8n1,none had buffer data: got 0 bytes 1.9 ECU woke up - init stage 1 1.9
had buffer data: got 1 bytes 00000000 7c ||| Connecting to MEMS 1.9 ECU Serial
cable set to: 9600,8n1,none had buffer data: got 0 bytes 1.9 had buffer data:
got 0 bytes Connecting to MEMS 1.9 ECU Serial cable set to: 9600,8n1,none had
buffer data: got 0 bytes 1.9 ECU woke up - init stage 1 1.9 ECU init stage 2 Got
CA Got 75 Got F4 00 Got D0 and ECU ID ECU ID: 00000000 c7 00 05 cb |....| Got
data 80 Got data 7D Got data 80 Got data 7D Got data 80 Got data 7D Got data 80
Got data 7D Got data 80 Got data */

<template>
  <span class="float-right">
    <label
      >ECU ID: <span class="badge badge-dark">{{ ECUID }}</span></label
    >
    <label class="ml-3"
      >ECU Serial: <span class="badge badge-dark">{{ ECUSerial }}</span></label
    >
  </span>
  <button
    class="btn btn-outline-secondary btn-sm mr-2 mb-2"
    @click="openSerialPort()"
  >
    Open Serial Port
  </button>

  <button
    class="btn btn-outline-secondary btn-sm mr-2 mb-2"
    @click="closeSerialPort()"
  >
    Disconnect
  </button>
  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="newInit()">
    New Init
  </button>

  <hr />

  <div class="card-group text-center">
    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Time</h6>
        <h3 class="card-text">{{ Dataframe.Time }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">RPM</h6>
        <h3 class="card-text">{{ Dataframe.EngineRPM }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">AFR</h6>
        <h3 class="card-text">{{ Dataframe.AirFuelRatio }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">
          Lambda {{ Dataframe.ClosedLoop ? "Closed" : "Open" }}
        </h6>
        <h3 class="card-text">{{ Dataframe.LambdaVoltage }}</h3>
      </div>
    </div>
  </div>

  <div class="card-group text-center mt-2">
    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Short Term Fuel Trim</h6>
        <h3 class="card-text">{{ Dataframe.ShortTermFuelTrim }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Long Term Fuel Trim</h6>
        <h3 class="card-text">{{ Dataframe.LongTermFuelTrim }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Ignition Advance</h6>
        <h3 class="card-text">{{ Dataframe.IgnitionAdvance }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Ignition Advance Offset</h6>
        <h3 class="card-text">{{ Dataframe.IgnitionAdvanceOffset80 }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Battery Voltage</h6>
        <h3 class="card-text">{{ Dataframe.BatteryVoltage }}</h3>
      </div>
    </div>
  </div>

  <hr />
  <div>
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
      @click="sendToEcu([0xd0])"
    >
      ECU ID
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0xd1])"
    >
      ECU ID/VER
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
  
  <div class="card-columns mt-4">
    <div class="card" v-for="param in parameters" v-bind:key="param">
      <div class="card-body">
        {{ param }}
        <span class="badge badge-light float-right">{{
          Dataframe[param]
        }}</span>
      </div>
    </div>
  </div>
  <hr />

  <label
    ><span class="badge badge-light"
      >{{ Dataframe.Dataframe7d.length }} {{ Dataframe.Dataframe7d }}</span
    ></label
  >
  <label
    ><span class="badge badge-light"
      >{{ Dataframe.Dataframe80.length }} {{ Dataframe.Dataframe80 }}</span
    ></label
  >
  <br />

  <pre>{{ debug_log.join("\n") }}</pre>
</template>
<style lang="scss"></style>

this.ser.port.open
<script lang="ts">
import imported_data from "./data/run-demo.fcr.json";
//import imported_data from "./data/nofaults.fcr.json";

export default {
  name: "app",
  data() {
    return {
      replay: {
        timer: null,
        step: 0,
      },
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
        //"7d201016ff92401bffff010179670cff51ffff3588840dff134015801a0029c02a",
        //"7d201014ff924057ffff0100806400ff64ffff3080800eff16801b00220031c01f",
        //"7d21400aff910057ffff0100716400005aff002e807b690022002200220022000100"
        //Dataframe80:
        //"801c08ba80ff56ff1f842308000100002037876b0417055c05df100000",
        //"801c00006fff4fff64781b00000100002037877b055f05380ca5000000"
        Dataframe80:
          "801c000000000000000000000000000000000000000000000000000000",
        Dataframe7d:
          "7d2000000000000000000000000000000000000000000000000000000000000000",
        //Dataframe7d:          "7d20100fff924047ffff0100796400ff6fffff35887f45ff134015801a0029c02a",
        //Dataframe80:          "801c03d66fff4fff32741b00000100002037877b00a7053807af100000",
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
        "Uk7d20",
      ],
      debug_log: [],
    };
  },
  mounted: function () {
    //this.dumpImportReadmemsHex();
  },
  methods: {
    simulateStart() {
      // this.parseD1( "d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630");
      this.replay.timer = null;
      this.replay.step = 0;
      this.replay.timer = setInterval(() => this.simulate(), 500);
      this.simulate();
      this.ECUID = imported_data.Name;
      // Slider for speed - realtime option
    },
    simulate() {
      if (imported_data.MemsData.length >= this.replay.step) {
        let data = imported_data.MemsData[this.replay.step];
        if (data) {
          this.Dataframe.Time = data.Time;
          // Patch 1.6 len to 1.9
          let x7d = "7d21" + data.Dataframe7d.substring(4).padEnd(66, "0");

          this.parse7D(this.hexToBytes(x7d));
          this.parse80(this.hexToBytes(data.Dataframe80));
          this.Dataframe.Dataframe7d = x7d;
          this.Dataframe.Dataframe80 = data.Dataframe80;
          this.replay.step++;
        } else {
          this.simulateStop();
        }
      }
    },
    simulateStop() {
      clearInterval(this.replay.timer);
      this.replay.timer = null;
      this.replay.step = 0;
    },
    dumpImportReadmemsHex() {
      //  Past raw, generate csv and then over-paste the times
      //memscene.exe -file run-1649731232340.raw -output run-1649731232340.csv
      let odd = 0;
      this.debug(
        imported_data.MemsData.map((x) => {
          return x.Time;
        }).join("\n")
      );
      let q = imported_data.MemsData[0].Dataframe7d;
      let y = q
        .substring(2)
        .split("")
        .map(function (o) {
          let sep = odd % 2 ? " " : "";
          odd++;
          return o + sep;
        });
      odd = 0;
      this.debug(
        "ECU responded to D0 command with: 99 00 02 03\n" +
          imported_data.MemsData.map((x) => {
            return (
              "80: " +
              x.Dataframe80.substring(2)
                .split("")
                .map(function (o) {
                  let sep = odd % 2 ? " " : "";
                  odd++;
                  return o + sep;
                })
                .join("") +
              "\n7D: " +
              "20 " +
              x.Dataframe7d.substring(5, 67)
                .split("")
                .map(function (o) {
                  let sep = odd % 2 ? " " : "";
                  odd++;
                  return o.toUpperCase() + sep;
                })
                .join("")
            );
          }).join("\n")
      );
    },
    parseD0(data) {
      let bytes = this.hexToBytes(data);
      var v = new DataView(bytes);
      this.ECUSerial = v.getUint32(1).toString();
    },
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
      if (len < 33) {
        // len is 33 mems 1.9
        this.debug(`expected len 33 for 0x7d got ${len}`);
        return;
      } else {
        let offset = 1;
        let v7d = {
          IgnitionSwitch: v.getUint8(0x01 + offset) > 0,
          ThrottleAngle: ((v.getUint8(0x02 + offset) * 6.0) / 10.0).toFixed(1),
          Uk7d03: v.getUint8(0x03 + offset),
          AirFuelRatio: (v.getUint8(0x04 + offset) / 10.0).toFixed(1),
          DTC2: v.getUint8(0x05 + offset),
          LambdaVoltage: v.getUint8(0x06 + offset) * 5,
          LambdaFrequency: v.getUint8(0x07 + offset),
          LambdaDutycycle: v.getUint8(0x08 + offset),
          LambdaStatus: v.getUint8(0x09 + offset),
          ClosedLoop: v.getUint8(0x0a + offset),
          LongTermFuelTrim: v.getInt8(0x0b + offset),
          ShortTermFuelTrim: v.getInt8(0x0c + offset),
          //FuelTrimCorrection: v.getInt8(0x+offset),
          CarbonCanisterPurgeValve: v.getUint8(0x0d + offset),
          DTC3: v.getUint8(0x0e + offset),
          IdleBasePosition: v.getUint8(0x0f + offset),
          Uk7d10: v.getUint8(0x10 + offset),
          DTC4: v.getUint8(0x11 + offset),
          IgnitionAdvanceOffset7d: v.getUint8(0x12 + offset) - 48,
          IdleSpeedOffset: v.getUint8(0x13 + offset),
          Uk7d14: v.getUint8(0x14 + offset),
          Uk7d15: v.getUint8(0x15 + offset),
          DTC5: v.getUint8(0x16 + offset),
          Uk7d17: v.getUint8(0x17 + offset),
          Uk7d18: v.getUint8(0x18 + offset),
          Uk7d19: v.getUint8(0x19 + offset),
          Uk7d1a: v.getUint8(0x1a + offset),
          Uk7d1b: v.getUint8(0x1b + offset),
          Uk7d1c: v.getUint8(0x1c + offset),
          Uk7d1d: v.getUint8(0x1d + offset),
          Uk7d1e: v.getUint8(0x1e + offset),
          JackCount: v.getUint8(0x1f + offset),
          Uk7d20: v.getUint8(0x20), // Mems 1.9
        };
        Object.assign(this.Dataframe, v7d);
      }
    },
    parse80(data: ArrayBuffer) {
      var v = new DataView(data);
      let type = v.getUint8(0);
      let len = v.getUint8(1);

      if (len !== 28) {
        this.debug(" expected len 28 for 0x80");
        return;
      } else {
        //"80 1c 00 00 6f ff 4f ff 64 78 1b00000100002037877b055f05380ca5000000"
        let offset = 1;
        let v80 = {
          EngineRPM: v.getInt16(0x01 + offset),
          CoolantTemp: v.getUint8(0x03 + offset) - 55.0,
          AmbientTemp: v.getUint8(0x04 + offset) - 55.0,
          IntakeAirTemp: v.getUint8(0x05 + offset) - 55.0,
          FuelTemp: v.getUint8(0x06 + offset) - 55.0,
          ManifoldAbsolutePressure: v.getUint8(0x07 + offset),
          BatteryVoltage: (v.getUint8(0x08 + offset) / 10.0).toFixed(2),
          ThrottlePotSensor: (v.getUint8(0x09 + offset) * 0.02).toFixed(2),
          ThrottlePosition: (v.getUint8(0x09 + offset) * 0.02).toFixed(2), // calc as 0 - 90 ? x79_02  ThrottleAngle
          IdleSwitch: v.getUint8(0x0a + offset),
          AirconSwitch: v.getUint8(0x0b + offset),
          ParkNeutralSwitch: v.getUint8(0x0c + offset),
          DTC0: v.getUint8(0x0d + offset),
          DTC1: v.getUint8(0x0e + offset),
          IdleSetPoint: v.getUint8(0x0f) * 6.1,
          IdleHot: v.getUint8(0x10 + offset),
          Uk8011: v.getUint8(0x11 + offset),
          IACPosition: v.getUint8(0x12 + offset),
          IdleSpeedDeviation: v.getInt16(0x13 + offset),
          IgnitionAdvanceOffset80: v.getInt8(0x15 + offset),
          IgnitionAdvance: (v.getInt8(0x16 + offset) / 2.0 - 24.0).toFixed(1),
          CoilTime: (v.getUint16(0x17 + offset) * 0.0002).toFixed(4),
          CrankshaftPositionSensor: v.getUint8(0x19 + offset),
          Uk801a: v.getUint8(0x1a + offset),
          Uk801b: v.getUint8(0x1b + offset),
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
      //this.debug({ms:ms,actual:performance.now() - start});
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
      this.debug_log.unshift(msg);
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

    async newInit() {
      this.ser.port = await navigator.serial.requestPort();
      this.debug(this.ser.port.getInfo());

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

      setTimeout(() => this.baud5listen(), 0);

      let sleepMs = 200;
      sleepMs = 200;
      //this.wakeUp5Baud(ecuAddress, sleepMs);
      //this.wakeUp5BaudNew10bits(ecuAddress,sleepMs)
      //this.wakeUp5BaudNewTiming(ecuAddress,sleepMs)
      this.slowInit19(ecuAddress, sleepMs);
    },

    async baud5listen() {
      this.debug("set reader");
      let read = 0;
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
            this.ser.buffer = this.ser.buffer.concat(
              this.hex(Array.from(value))
            );
            //this.debug(`l: ${read.length} d: ${read} v: ${value}`);
            //this.debug( `<< ${read}`);
            if (start === null) start = value[0];
            switch (start) {
              default: {
                //read=read.substring(2);
                this.debug(`<< ${this.ser.buffer}`);
                read++;
                if (read === 5) {
                  await this.wait(50);
                  this.sendToEcu([0x7c]);
                }
                start = null;
                this.ser.buffer = "";
              }
            }
          }
        } catch (error) {
          this.debug(`error: ${error.message}`);
          this.debug(error);
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
      //    this.debug(ports);
      //      let info = ports[0].getInfo();

      this.ser.port = await navigator.serial.requestPort();
      this.debug(this.ser.port.getInfo());

      await this.ser.port.open({
        baudRate: 9600,
        databits: 8,
        bufferSize: 128,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });
      this.debug(this.ser.port);

      if (1) this.sendToEcu([0xd0]);

      if (`1`)
        this.timer = setInterval(() => {
          this.sendToEcu([0x7d]);
        }, 1000);

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
                if (this.ser.buffer.length < 60) {
                  //this.debug( `expected 60 (${this.ser.buffer.length}) bytes for 0x80`   );
                  break;
                }
                if (this.ser.buffer.length != 60) {
                  this.ser.buffer = "";
                  start = null;
                  this.debug(`rejected << ${this.ser.buffer}`);
                  return;
                }
                //this.debug( `using (${this.ser.buffer.length}) bytes for 0x80`   );

                this.ser.buffer = this.ser.buffer.substring(2);
                this.Dataframe.Dataframe80 = this.ser.buffer;
                // Patch size for https://analysis.memsfcr.co.uk/
                let Mems1_6_7b =
                  "7d21" + this.Dataframe.Dataframe7d.substring(4, 66);
                this.log.MemsData.push({
                  Time: this.Dataframe.Time,
                  Dataframe80: this.Dataframe.Dataframe80,
                  Dataframe7d: Mems1_6_7b,
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
                //d0d0c70005cb
                this.debug(`Got ID ${this.ser.buffer}`);
                this.debug(this.ser.buffer);
                this.ser.buffer = "";
                start = null;
                break;
              }
              case 0x7d:
                if (this.ser.buffer.length < 70) {
                  //this.debug(  `expected 70 (${this.ser.buffer.length}) bytes for 0x7d` );
                  break;
                }
                if (this.ser.buffer.length != 70) {
                  this.ser.buffer = "";
                  start = null;
                  this.debug(`rejected << ${this.ser.buffer}`);
                  return;
                }
                this.ser.buffer = this.ser.buffer.substring(2);
                //this.debug(                    `got (${this.ser.buffer.length}) bytes for 0x7d`);

                this.Dataframe.Dataframe7d = this.ser.buffer;
                let now = new Date();
                let ms = now.getMilliseconds().toString(10).padStart(3, "0");
                this.Dataframe.Time = `${now.toLocaleTimeString()}.${ms}`;
                this.parse7D(this.hexToBytes(this.ser.buffer));
                this.sendToEcu([0x80]); // trigger next frame
                this.ser.buffer = "";
                start = null;
                break;
              case 0xd0:
                //
                if (this.ser.buffer.length < 4) {
                  this.debug(
                    `expected 64 bytes for 0xd0, got ${this.ser.buffer.length}`
                  );
                  break;
                }
                this.debug(`0xd0 << ${this.ser.buffer.length}`);
                this.ser.buffer = this.ser.buffer.substring(2);
                this.parseD0(this.ser.buffer);
                this.ser.buffer = "";
                start = null;
                break;

              case 0xd1:
                // d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630
                if (this.ser.buffer.length < 70) {
                  this.debug(
                    `expected 64 bytes for 0xd1, got ${this.ser.buffer.length}`
                  );
                  if (this.ser.buffer.length == 2) {
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
          this.debug(error);
        } finally {
          this.ser.reader.releaseLock();
          this.ser.reader = null;
          this.debug("released lock");
        }
      }
    },
    async sleepUntil(timestampMs) {
      let now = new Date().getTime();
      let sleepFor = timestampMs - now;
      await new Promise((resolve) => setTimeout(resolve, sleepFor));
    },
    async slowInit19() {
      ecuAddress = 0x16; // 22

      //resetTimeout(5000);
      dataBuffer = [];
      this.debug(
        "Attempting ECU connection... (address: " + ecuAddress + ") (slow init)"
      );

      debug("Starting slow init, this takes 2 seconds to send");
      // mEcuAddr = 22;

      // pause/delay to clear the line
      await this.ser.port.setSignals({ brk: false, break: false });
      await this.wait(2000);

      let before = new Date().getTime();

      // start bit:
      await this.ser.port.setSignals({ brk: true, break: true });
      await this.waitUntil(before + 200);

      for (var i = 0; i < 8; i++) {
        let bit = (ecuAddress >> i) & 1;
        this.debug(i + " " + bit);
        if (bit > 0) {
          await this.ser.port.setSignals({ brk: false, break: false });
        } else {
          await this.ser.port.setSignals({ brk: true, break: true });
        }
        await this.waitUntil(before + 200 + (i + 1) * 200);
      }

      // stop bit:
      await this.ser.port.setSignals({ brk: false, break: false });
      await this.waitUntil(before + 200 + 8 * 200 + 200);

      this.debug("Done sending slow init");

      // debug(new Date().getTime()-before);
      resetTimeout(3000);

      parseDataBufferSlowInit();
    },
    async wakeUp5Baud(ecuAddress, sleepMs) {
      // 5 baud/200ms per bit
      // start bit low, stop bit high

      // pause/delay to clear the line
      await this.ser.port.setSignals({ brk: false, break: false });
      this.debug("sleeping for 2 seconds to clear the line");
      await this.wait(2000);
      this.debug("sending wakeup");
      // start bit
      await this.ser.port.setSignals({ brk: true, break: true });
      await this.wait(sleepMs);
      for (var i = 0; i < 8; i++) {
        let bit = (ecuAddress >> i) & 1;
        this.debug(bit);
        if (bit > 0) {
          await this.ser.port.setSignals({ brk: false, break: false });
        } else {
          await this.ser.port.setSignals({ brk: true, break: true });
        }
        await this.wait(sleepMs);
      }
      // stop bit
      await this.ser.port.setSignals({ brk: false, break: false });
      await this.wait(sleepMs);
      // await this.wait(195);
      this.debug("Done sending wakeup");
    },

    async wakeUp5BaudNew10bits(ecuAddress, sleepMs) {
      this.debug("sleeping for 2 seconds to clear the line...");
      await this.ser.port.setSignals({ break: false });
      await this.wait(2000);
      let times = [];
      //      await this.ser.port.setSignals({ break: true });

      let i = 0;

      ecuAddress = (ecuAddress << 1) | 1;

      let bits = ecuAddress.toString(2).padStart(10, 0).split("").reverse();
      let b = [];
      /*
      start = performance.now();
      let timer = setInterval(() => {
        this.ser.port.setSignals({ break: !bits[i] });
        times.push(performance.now() - start);
        i = i + 1;
        if (i === 10) clearInterval(timer);
      }, sleepMs);
      this.debug(times);
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
      this.debug(times);
      this.debug(b);
      await this.ser.port.setSignals({ brk: false, break: false });
      await this.wait(sleepMs);
    },

    async wakeUp5BaudNewTiming(ecuAddress, sleepMs) {
      this.debug("wakeUp5BaudNewTiming");
      // The break signal state (all low, no stop bit) until released.
      /* http://www.internetsomething.com/kwp/KWP2000%20ISO%2014230-2%20KLine%20.pdf ISO 14230 KWP 2000
W0 300 - Time before the tester starts to transmit the address byte

W1 60 300 Time from end of the address byte to start of synchronisation pattern
W2 5 20 Time from end of the synchronisation pattern to the start of key byte 1
W3 0 20 Time between key byte 1 and key byte 2
W4 25 50 Time between key byte 2 (from the ECU) and its inversion from the tester. Also the time
from the inverted key byte 2 from the tester and the inverted address from the ECU
*/
      //await this.ser.port.setSignals({ break: true });
      //      await this.wait(sleepMs*4);

      await this.ser.port.setSignals({ break: false });
      await this.wait(sleepMs * 8);

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
      await this.wait(sleepMs * 6);

      // expectiNg 0x55, 0x76, 0x83
      //this.sendToEcu([0x7c]);
      //Expect 0x7c, 0xE9
    },
  },
};
</script>

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

  <button
    class="btn btn-outline-secondary btn-sm mr-2 mb-2"
    @click="download()"
  >
    <i class="fas fa-download"></i>
    Download
  </button>

  <button
    v-if="!replay.timer"
    class="btn btn-outline-secondary btn-sm mr-2 mb-2"
    @click="simulateStart()"
  >
    <i class="fas fa-refresh"></i>
    Replay
  </button>

  <button
    v-if="replay.timer"
    class="btn btn-outline-secondary btn-sm mr-2 mb-2"
    @click="simulateStop()"
  >
    <i class="fas fa-stop"></i>
    Stop
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
        <h6 class="card-title">
          Lambda {{ !Dataframe.ClosedLoop ? "Closed" : "Open" }}
          {{ Dataframe.ClosedLoop }}
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
        <h6 class="card-title">Ignition Advance</h6>
        <h3 class="card-text">{{ Dataframe.IgnitionAdvance }}</h3>
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
      ECU SER
    </button>
    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2"
      @click="sendToEcu([0xd1])"
    >
      ECU ID/SER
    </button>

       <button
      class="btn btn-outline-danger btn-sm mr-2 mb-2 float-right"
      @click="sendToEcu([0xfa])"
    >
      Reset ECU
    </button>

    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2 float-right"
      @click="sendToEcu([0x0f])"
    >
      <i class="fa fa-undo">&nbsp;</i>
      Reset Adjustments
    </button>

    <button
      class="btn btn-outline-secondary btn-sm mr-2 mb-2 float-right"
      @click="sendToEcu([0xcc])"
    >
      Clear Faults
    </button>

 

    <p></p>

    <div class="btn-group mr-2" role="group">
      <button
        type="button"
        class="btn btn-sm btn-outline-secondary"
        @click="sendToEcu([0x0d,0x00])"
      >
        Off
      </button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0">Radiator Fan</label></span
      >
      <button
        type="button"
        class="btn btn-sm btn-outline-secondary"
        @click="sendToEcu([0x1d,0x00])"
      >
        On
      </button>
    </div>

    <div class="btn-group mr-2" role="group">
      <button
        type="button"
        class="btn btn-sm btn-outline-secondary"
        @click="sendToEcu([0x92])"
      >
        - 
      </button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Idle Speed Δ
          <span class="ml-1 badge badge-dark">{{
            Dataframe.IdleSpeedOffset
          }}</span></label
        ></span
      >
      <button
        type="button"
        class="btn btn-sm btn-outline-secondary"
        @click="sendToEcu([0x91])"
      >
        + 
      </button>
    </div>

    <div class="btn-group mr-2" role="group">
      <button
        type="button"
        class="btn btn-sm btn-outline-secondary"
        @click="sendToEcu([0x8a])"
      >
        - 
      </button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Idle Decay Δ
          <span class="ml-1 badge badge-dark">{{
            Dataframe.IdleSpeedOffset
          }}</span></label
        ></span
      >
      <button
        type="button"
        class="btn btn-sm btn-outline-secondary"
        @click="sendToEcu([0x89])"
      >
        + 
      </button>
    </div>

        
  

  <div class="btn-group mr-2" role="group">
    <button
      type="button"
      class="btn btn-sm btn-outline-secondary"
      @click="sendToEcu([0x7a])"
    >
      -
    </button>
    <span type="button" class="btn btn-sm btn-outline-secondary disabled"
      ><label class="mb-0"
        >Long Term Fuel Trim
        <span class="ml-1 badge badge-dark">{{
          Dataframe.LongTermFuelTrim
        }}</span></label
      ></span
    >
    <button
      type="button"
      class="btn btn-sm btn-outline-secondary"
      @click="sendToEcu([0x7b])"
    >
      +
    </button>
  </div>


    <p></p>

  <div class="btn-group mr-2" role="group">
    <button
      type="button"
      class="btn btn-sm btn-outline-secondary"
      @click="sendToEcu([0x94])"
    >
      -
    </button>
    <span type="button" class="btn btn-sm btn-outline-secondary disabled"
      ><label class="mb-0"
        >Ignition Advance Δ
        <span class="ml-1 badge badge-dark">{{
          Dataframe.IgnitionAdvanceOffset80
        }}</span></label
      ></span
    >
    <button
      type="button"
      class="btn btn-sm btn-outline-secondary"
      @click="sendToEcu([0x93])"
    >
      +
    </button>
    </div>


    <div class="btn-group mr-2" role="group">
      <button
        type="button"
        class="btn btn-sm btn-outline-secondary"
        @click="sendToEcu([0x7f])"
      >
        - 
      </button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >IgnitionAdvanceOffset7d ?? Δ
          <span class="ml-1 badge badge-dark">{{
            Dataframe.IgnitionAdvanceOffset7d 
          }}  </span></label
        ></span
      >
      <button
        type="button"
        class="btn btn-sm btn-outline-secondary"
        @click="sendToEcu([0x7e])"
      >
        + 
      </button>
     
    </div>
  </div>

  <pre>{{ debug_log.join("\n") }}</pre>

  <div class="card-columns mt-4">
    <div class="card" v-for="param in parameters" v-bind:key="param">
      <div class="card-body">
        {{ param }}
        <span class="badge badge-dark float-right">{{ Dataframe[param] }}</span>
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

<script lang="ts">
//import imported_data from "./data/run-demo.fcr.json";
//import imported_data from "./data/run-1649820449532.fcr.json";
//import imported_data from "./data/run-1649822529717.fcr.json";
//import imported_data from "./data/run-1652058533489.fcr.json";

import imported_data from "./data/run-1652139960583.fcr.json";
//import imported_data from "./data/run-1651531224249.fcr.json";  // Small start up

//import imported_data from "./data/nofaults.fcr.json";
import { Chart, Grid, Line, Tooltip } from "vue3-charts";
import { VueSvgGauge } from "vue-svg-gauge";

// https://codepen.io/netzzwerg/pen/QBQBaM
// simple vue3 svg
//
// https://hellocomet.github.io/vue-svg-gauge/
export default {
  name: "app",
  components: {
    VueSvgGauge,
  },

  data() {
    return {
      appVersion: __APP_VERSION__,
      waitReply: false,
      queuedBytes: [],
      expectingBytes: null,
      last_reponse: "",
      rpmD1: 0,
      rpmD2: { min: [], max: [], val: 0 },
      weightKg: 820,
      Nm: 0,
      gForce: 0,
      dist: 0,
      lastDist: 0,
      deltaDist: 0,
      gear: 0,
      Acceleration: [0,0,0],
      gearing: {
        0: 0,
        1: 5.64,
        2: 9.69,
        3: 13.66,
        4: 17.28,
        5: 23.33,
      },
      lastEngineRPM: 0,
      replay: {
        timer: null,
        step: 0,
        pause: false,
        interval: 1000,
        port: null,
        reader: null,
      },
      record: {
        timer: null,
      },
      ser: {
        mode: 0,
        port: null,
        dataframe: [] as number[],
        buffer: "" as string,
        stage: 0,
        retries: 0,
        pause: 0,
        replysRaw: [] as string[],
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
      history: [],
      stopwatch: 0,
      Dataframe: {
        Time: null, // "00:00:00.000",
        EngineRPM: 0,
        CoolantTemp: 0,
        AmbientTemp: 0,
        IntakeAirTemp: 0,
        FuelTemp: 0,
        ManifoldAbsolutePressure: 0,
        BatteryVoltage: 0,
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
        Uk7d20: 0,
        Dataframe80: "801c000000000000000000000000000000000000000000000000000000",
        Dataframe7d: "7d21000000000000000000000000000000000000000000000000000000000000000000",
      },
      parameters: [
        "EngineRPM",
        "CoolantTemp",
        "IntakeAirTemp",
        "FuelTemp",
        "ManifoldAbsolutePressure",
        "BatteryVoltage",
        "ThrottlePosition",
        "ThrottleAngle",
        "IdleSpeedDeviation",
        "IdleSetPoint",
        "IdleHot",
        "IdleBasePosition",
        "IACPosition",
        "IdleSpeedOffset",
        "IdleSwitch",
        "IgnitionAdvance",
        "IgnitionAdvanceOffset7d",
        "IgnitionAdvanceOffset80",
        "CoilTime",
        "LambdaVoltage",
        "LambdaStatus",
        "ClosedLoop",
        "LongTermFuelTrim",
        "ShortTermFuelTrim",
        "FuelTrimCorrection",
        "IgnitionSwitch",
        "DTC0",
        "DTC1",
        "DTC2",
        "DTC3",
        "DTC4",
        "DTC5",
        "AmbientTemp",
        "AirconSwitch",
        "AirFuelRatio",
        "CarbonCanisterPurgeValve",
        "CrankshaftPositionSensor",
        "JackCount",
        "LambdaDutycycle",
        "LambdaFrequency",
        "ParkNeutralSwitch",
        "Uk7d03",
        "Uk7d10",
        "Uk7d14",
        "Uk7d15",
        "Uk7d17",
        "Uk7d18",
        "Uk7d19",
        "Uk7d1a",
        "Uk7d1b",
        "Uk7d1c",
        "Uk7d1d",
        "Uk7d1e",
        "Uk7d20",
        "Uk8011",
        "Uk801a",
        "Uk801b",
      ],
      lambdaDamped: [435],
      debug_log: [],
    };
  },
  computed: {
    AirFuelRatioCalc() {
      return ((this.LambdaVoltageDamped - 450) / 100).toFixed(0) * 10;
    },
    AirFuelRatioLabel() {
      if (this.LambdaVoltageDamped < 400) return "Lean";
      if (this.LambdaVoltageDamped > 500) return "Rich";
      return "Ok";
    },

    LambdaVoltageDamped() {
      let sum = this.lambdaDamped.reduce((a, b) => a + b);
      return (sum / this.lambdaDamped.length).toFixed(0);
    },

    faults() {
      let f = [];
      if (this.Dataframe.DTC0 & 0x01) f.push("Coolant Sensor");
      if (this.Dataframe.DTC0 & 0x02) f.push("Inlet Air Temp Sensor");
      if (this.Dataframe.DTC0 & 0x20) f.push("Fuel Rail Temp Sensor");
      if (this.Dataframe.DTC0 & 0x08) f.push("Turbo Overboost");
      if (this.Dataframe.DTC0 & 0x10) f.push("Ambient Air Temp");
      if (this.Dataframe.DTC0 & 0x20) f.push("Fuel Rail Temp Sensor");
      if (this.Dataframe.DTC0 & 0x40) f.push("Knock Fault Detected");

      if (this.Dataframe.DTC1 & 0x01) f.push("Coolant Temp Gauge");
      if (this.Dataframe.DTC1 & 0x02) f.push("Fuel Pump Circuit");
      if (this.Dataframe.DTC1 & 0x04) f.push("Air Con");
      if (this.Dataframe.DTC1 & 0x08) f.push("Purge Valve");
      if (this.Dataframe.DTC1 & 0x10) f.push("MAP Sensor");
      if (this.Dataframe.DTC1 & 0x20) f.push("Boost Valve");
      if (this.Dataframe.DTC1 & 0x40) f.push("Throttle Position Sensor");
      if (this.Dataframe.LambdaStatus === 0) f.push("Lambda Status");

      if (this.Dataframe.DTC2 & 0x04) f.push("Lambda Heater");
      if (this.Dataframe.DTC2 & 0x08) f.push("Secondary Trigger Sync");
      if (this.Dataframe.DTC2 & 0x10) f.push("Fan 1 Control");
      if (this.Dataframe.DTC2 & 0x40) f.push("Fan 2 Control");
      if (this.Dataframe.DTC3 & 0x01) f.push("Primary Trigger Sync");
      if (this.Dataframe.DTC3 & 0x04) f.push("DTC3 0x4");

      return f;
    },
    MPH() {
      return ((this.Dataframe.EngineRPM / 1000.0) * this.gearing[this.gear]).toFixed(1);
    },
    KPH() {
      return this.Mph2Kph(this.MPH);
    },
    ft_lb() {
      return (this.Nm / 1.36).toFixed(1);
    },
    hp() {
      return ((this.ft_lb / 5252) * this.Dataframe.EngineRPM).toFixed(1);
    },
    chartKey() {
      return this.replay.pause ? false : this.stopwatch;
    },
    chartSize() {
      return { width: document.body.clientWidth, height: 250 };
    },
    stopwatchFormat() {
      return this.timeFormat(this.stopwatch);
    },
  },
  watch: {
    "Dataframe.Time"(before, after) {
      this.lambdaDamped.unshift(this.Dataframe.LambdaVoltage);
      this.lambdaDamped = this.lambdaDamped.slice(0, 5);
      if (before && after) {
        let delta = this.hhmmss2secs(before) - this.hhmmss2secs(after);
        this.stopwatch = Number(this.stopwatch + delta);
        let EngineRPM = this.Dataframe.EngineRPM;
        if (this.history.length > 20) this.history.shift();
        let rpmD1 = ((EngineRPM - this.lastEngineRPM) / delta).toFixed(0);
        let rpmD2 = ((EngineRPM - this.lastEngineRPM) / delta / delta).toFixed(0);
        this.dist = ((this.KPH * 1000.0) / 3600.0).toFixed(2);
        this.deltaDist = (this.dist - this.lastDist).toFixed(3);
        if (this.deltaDist > 0) this.Nm = ((this.deltaDist / delta) * this.weightKg).toFixed(2);
        this.gForce = (this.deltaDist / delta / delta).toFixed(1);
        if (this.Dataframe.IdleSwitch === 0) {
          this.gear = 0;
        } else {
          if (rpmD2 > 600 || this.gear === 0) {
            this.gear++;
            this.rpmD2.max.push(rpmD2);
            Math.max(rpmD2, this.rpmD2.max);
          }
          if (rpmD2 < -650) {
            this.gear--;
            if (this.gear < 0) {
              this.gear = 0;
              this.debug("False downchange");
            }
            this.rpmD2.min.push(rpmD2);
          }
        }
        this.rpmD1 = rpmD1;
        this.rpmD2.val = rpmD2;
        this.lastEngineRPM = EngineRPM;
        this.lastDist = this.dist;
        this.history.push({
          Time: this.stopwatchFormat,
          EngineRPM: this.Dataframe.EngineRPM,
          LambdaVoltage: this.LambdaVoltageDamped,
          AFR: this.AirFuelRatioCalc,
          gear: this.gear,
        });
      }
    },
  },

  mounted() {},
  async unmounted() {
    await this.closeSerialPort();
    await this.replaySerialClose();
  },
  methods: {
    Time() {
      let now = new Date();
      let ms = now.getMilliseconds().toString(10).padStart(3, "0");
      return `${now.toLocaleTimeString()}.${ms}`;
    },

    Mph2Kph(MPH) {
      return Number((MPH / 50.0) * 80).toFixed(1);
    },
    hhmmss2secs(t) {
      return t.split(":").reduce((acc, time) => 60 * acc + +time);
    },
    timeFormat(t) {
      return new Date(t * 1000).toISOString().substr(14, 7);
    },
    simulateStart(step = 0) {
      this.replay.step = step;
      this.replay.timer = setInterval(() => this.simulate(), this.replay.interval);
      this.simulate();
      this.ECUID = imported_data.Name;
    },

    simulateStep(step = 10) {
      this.replay.step += step;
      this.simulate();
    },

    async replaySerialStart(step = 0) {
      this.replay.port = await navigator.serial.requestPort();

      try {
        await this.replay.port.open({
          baudRate: 9600,
          databits: 8,
          bufferSize: 32,
          parity: "none",
          stopbits: 1,
          flowControl: "none",
        });
      } catch (error) {
        this.debug("Ecu port is already open - refresh page to close");
        return;
      }
      this.replay.step = step;

      this.replay.reader = this.replay.port.readable.getReader();
      let first = 0;
      let final = 0;
      while (true) {
        try {
          const { value, done } = await this.replay.reader.read();
          if (value) {
            let hex = this.hex(Array.from(value));
            switch (value[0]) {
              case 0x00:
                first++;
                if (first >= 3) {
                  this.replayWrite("557583", false); // after 0x00 0x00 0x00
                }
                break;
              case 0xca:
                if (first === 0) {
                  this.replayWrite("CA", false);
                } else {
                  this.replayWrite("CA", true);
                  first++;
                }
                break;
              case 0x75:
                this.replayWrite("75", true);
                break;
              case 0xf4:
                this.replayWrite("F4F400", false);
                break;
              case 0x7c:
                this.replayWrite("7ce9", false);
                break;

              case 0x80:
                let data = imported_data.MemsData[this.replay.step];
                this.replayWrite(data.Dataframe80);
                this.replay.step++;
                break;
              case 0x7d:
                this.replaySerial();
                break;
              case 0xd0:
                this.replayWrite("d04b4c4833");
                break;
              case 0xd1:
                this.replayWrite("d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c483356303035c70005cb");
                break;
              default:
                //Echo
                this.replayWrite(`${hex}01`);
                console.log(`cmd: ${hex}`);
                break;
            }
          }
          if (done) {
            console.log("Serial replay Release Lock", done);
            this.replay.reader.releaseLock();
            break;
          }
        } catch (error) {

        } finally {
          this.replay.reader.releaseLock();
          await this.wait(100);
          this.replay.reader = this.replay.port.readable.getReader();
        }
      }
    },
    async replaySerialClose() {
      if (this.replay.reader) {
        await this.replay.reader.cancel();
      }
      if (this.replay.port) await this.replay.port.close();
      this.replay.port = null;
    },
    replayWrite(data, echo = true) {
      if (this.replay.port) {
        if (echo) data = data[0] + data[1] + data;

        let bytes = new Uint8Array(this.hexToBytes(data));

        let writer = this.replay.port.writable.getWriter();
        writer.write(Uint8Array.from(bytes));
        writer.releaseLock();
      }
    },
    replaySerial() {
      if (imported_data.MemsData.length >= this.replay.step) {
        let data = imported_data.MemsData[this.replay.step];
        if (data && this.replay.port) {
          this.replayWrite(data.Dataframe7d.padEnd(68, "0")); // Pad 1.6 mems data to 70 bytes
        } else {
          this.replaySerialStop();
        }
      }
    },
    async replaySerialStop() {
      clearInterval(this.replay.timer);
      this.replay.timer = null;
      this.replay.step = 0;
    },
    simulate() {
      if (this.replay.pause) return;
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
          this.ser.replysRaw.push('80'+data.Dataframe80);
          this.ser.replysRaw.push('7d'+x7d);
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
    simulatePause() {
      this.replay.pause = !this.replay.pause;
    },
    pollDataframes() {
      this.sendToEcu([0x7d]);
    },
    recordStart() {
      this.waitReply = false;
      this.recordStop();
      this.record.timer = setInterval(() => {
        //this.pollDataframes()
      }, 1000);
      this.sendToEcu([0x80]);
    },
    recordStop() {
      if (this.record.timer) clearInterval(this.record.timer);
      this.record.timer = null;
    },
    parseD0(data) {
      let bytes = this.hexToBytes(data);
      var v = new DataView(bytes);
      this.ECUSerial = v.getUint32(1).toString();
      this.log.dataframe_d0 = data;
    },
    parseD1(data) {
      // d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630 1.9 response longer...
      // _ KLH3V005 ___ KLH3V005 ___ KLH3V0
      let bytes = this.hexToBytes(data);
      var v = new DataView(bytes);
      this.ECUSerial = v.getUint32(9).toString();
      this.ECUID = String.fromCharCode.apply(null, new Uint8Array(bytes.slice(1, 9)));
      this.log.dataframe_d1 = data;
    },
    parse7D(data: ArrayBuffer) {
      var v = new DataView(data);
      let len = v.getUint8(1);
      if (len != 33 || !(v.getUint8(0) == 0x7d && v.byteLength >= 33)) {
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
          LongTermFuelTrim: v.getUint8(0x0b + offset),
          ShortTermFuelTrim: v.getUint8(0x0c + offset),
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
          Uk7d20: v.getUint8(0x20 + offset),
        };
        Object.assign(this.Dataframe, v7d);
      }
    },
    parse80(data: ArrayBuffer) {
      var v = new DataView(data);
      let type = v.getUint8(0);
      let len = v.getUint8(1);
      if (len !== 28 || !(v.getUint8(0) == 0x80 && v.byteLength == 29)) {
        this.debug(`expected len 28 (${len}) for 0x80`);
        return;
      } else {
        let offset = 1;
        let v80 = {
          EngineRPM: v.getInt16(0x01 + offset),
          CoolantTemp: v.getUint8(0x03 + offset) - 55.0,
          AmbientTemp: v.getUint8(0x04 + offset) - 55.0,
          IntakeAirTemp: v.getUint8(0x05 + offset) - 55.0,
          FuelTemp: v.getUint8(0x06 + offset) - 55.0,
          ManifoldAbsolutePressure: v.getUint8(0x07 + offset),
          BatteryVoltage: (v.getUint8(0x08 + offset) / 10.0).toFixed(2),
          ThrottlePosition: (v.getUint8(0x09 + offset) * 0.02).toFixed(2),
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
          IgnitionAdvanceOffset80: v.getUint8(0x15 + offset),
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
      for (let c = 0; c < hex.length; c += 2) view.setUint8(c / 2, parseInt(hex.substr(c, 2), 16));
      return bytes;
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
    },

    async sendQueuedToECU() {
      if (!this.waitReply && this.queuedBytes.length) {
        let byte = this.queuedBytes.shift();
        this.sendBytes([byte]);
      }
    },

    async sendToEcu(bytes) {
      this.queuedBytes.push(...bytes);
      this.sendQueuedToECU();
      if (this.queuedBytes.length > 2) {
        this.debug("cleared queue");
        this.waitReply = false;
        this.queuedBytes = [];
      }
    },

    async sendBytes(bytes) {
      if (this.ser.port) {
        let writer = this.ser.port.writable.getWriter();
        writer.write(Uint8Array.from(bytes));
        writer.releaseLock();
      this.waitReply = true
      }

    },
    hex(bytes, delim = "") {
      let result = bytes.map((x) => {
        let h = x.toString(16).padStart(2, "0");
        return h;
      });
      return result.join(delim);
    },
    debug(msg) {
      console.log(msg);
      this.debug_log.unshift(msg);
    },
    download() {
      let log = this.log;

      let now = new Date();
      log.Name = `run-${now.getTime()}.fcr.json`;
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
      this.ser.stage = 6;
      if (this.ser.reader) {
        await this.ser.reader.cancel();
      }
      if (this.ser.connectTimer) {
        clearInterval(this.ser.connectTimer);
        this.ser.connectTimer = null;
      }

      if (this.ser.port) await this.ser.port.close();
      this.ser.reader = null;
      this.ser.port = null;
      this.ser.mode = 0;
      this.ser.stage = 0;
    },

    CmdLength(cmd, dataframe, len_cmd) {
      switch (cmd) {
        case 0x00:
          return 3; //psuedo command - 5 baud echo expect 0x00 0x00 0x00 then reply
        case 0x55: // Stage 1 init
          return 3;
        case 0x7c: // Stage 2 init
        case 0xca: // Stage 2 init
        case 0x75: // Stage 2 init
          return 2;
        case 0x80:
          if (dataframe.length > 2) return dataframe[2] + 2; // command is 0x1c=> 28 + 2 = 30
          return 30;
        case 0x7d:
          if (dataframe.length > 2) {
            return dataframe[2] + 2;
          }

          return 35;
        case 0xd0:
          return 6;
        case 0xd1:
          return 38;
        default:
          console.log(`Cmd: ${cmd.toString(16).padStart(2, "0")}`);
          return 3;
      }
    },

    async waitUntil(timestampMs, pause) {
      const start = performance.now();
      while (performance.now() < timestampMs) {
        await new Promise((resolve) => setTimeout(resolve, 1));
        let delta = performance.now() - start;
        if (delta < 10) {
          while (performance.now() < timestampMs) {}
          break;
        }
      }
      return timestampMs + pause;
    },    
    async assertSignalAndWait(before, pause, brk) {
      await this.ser.port.setSignals({ break: brk });
      return await this.waitUntil(before, pause);
    },
    async slowInit(ecuAddress) {
      let pause = 200;
      this.ser.stage = 2;
      let before = await this.assertSignalAndWait(performance.now(), pause, false);
      before = await this.assertSignalAndWait(before, pause, true);
      this.ser.stage += 0.3;
      for (var i = 0; i < 8; i++) {
        let bit = (ecuAddress >> i) & 1;
        before = await this.assertSignalAndWait(before, pause, !bit);
      }
      this.ser.stage += 0.3;
      await this.assertSignalAndWait(before, pause, false);
    },

    /*
557683
7ce9
caca
*/
    async openSerialPort() {
      {
        let laSensor = new LinearAccelerationSensor({ frequency: 3 });
        laSensor.addEventListener("reading", (e) => {
          this.Acceleration = [laSensor.x, laSensor.y, laSensor.z];
        });
        laSensor.start();
      }

      this.ser.port = await navigator.serial.requestPort();
      try {
        await this.ser.port.open({
          baudRate: 9600,
          databits: 8,
          bufferSize: 64,
          parity: "none",
          stopbits: 1,
          flowControl: "none",
        });
      } catch (error) {
        this.debug("port is already open - refresh page to close");
        return;
      }
      this.ser.retries = 10;

      let start = 0;

      this.ser.stage = 4;
      this.sendBytes([0xca]);

      this.ser.connectTimer = setInterval(async () => {
        if (this.ser.stage != 2) return;
        await this.slowInit(0x16);

        this.ser.retries--;
      }, 3000);

      this.waitReply = false;
      while (this.ser.port?.readable) {
        this.ser.reader = this.ser.port.readable.getReader();
        let cmd = 0;
        let len_cmd = 0;
        let dataframe: ArrayBuffer[] = [];
        let going = true;
        let watchdog = null;
        try {
          while (going) {
            const { value, done } = await this.ser.reader.read();
            if (done) {
              console.log("Open Serial Port");
              going = false;
              break;
            }
            let inbound = Array.from(value);
            let raw = this.hex(inbound);
            this.ser.replysRaw.push(raw);
            let rest = [];
            if (len_cmd == 0) {
              cmd = inbound[0];
              dataframe = inbound;
              len_cmd = this.CmdLength(cmd, dataframe, len_cmd);
            } else {
              let required = len_cmd - dataframe.length;
              if (required < 0) {
                this.debug(`cmd length error ${cmd} ${len_cmd} vs ${dataframe.length} ${this.hex(dataframe)}`);
                len_cmd = 0;
                required = 0;
              }
              if (inbound.length > required) {
                this.debug(`inbound was: ${this.hex(inbound)}\n`);
                rest = inbound.slice(required);
                inbound = inbound.slice(0, required);
                dataframe.push(...inbound);
                this.debug(`${this.hex(inbound)} -> extra ${this.hex(rest)} ${this.hex(dataframe)}`);
              } else {
                dataframe.push(...inbound);
                inbound = [];
              }
            }
            this.ser.dataframe = dataframe;
            if (len_cmd > 0 && dataframe.length < len_cmd && this.watchdog == null) {
              if (this.watchdog) {
                clearTimeout(this.watchdog);
              }
              this.watchdog = setTimeout(() => {
                this.watchdog = null;
                this.debug(`watchdog timeout... ${dataframe.length} ${len_cmd} 0x${cmd.toString(16)}`);
                if (cmd === 0xca) {
                  this.ser.stage = 2;
                } else {
                  //this.sendBytes([0xf4]);//send echo to re-sync
                }
                this.waitReply = false;

                len_cmd = 0;
                dataframe = [];
                this.waitReply = false;
              }, 500);
            }
            if (dataframe.length == len_cmd) {
              if (this.watchdog) {
                clearTimeout(this.watchdog);
                this.watchdog = null;
              }
              this.waitReply = false; // Allow commands to be sent
              let data = this.hex(dataframe);

              this.processCmd(cmd, data);
              len_cmd = 0;
              dataframe = [];
              if (rest.length) {
                cmd = rest[0];
                dataframe = rest;
                len_cmd = this.CmdLength(cmd, dataframe, len_cmd);
                this.debug(`rest: ${this.hex(rest)}, ${cmd}`);
                let data = this.hex(dataframe);

                if (dataframe.length == len_cmd) {
                  this.processCmd(cmd, data);
                  dataframe = [];
                  len_cmd = 0;
                  this.waitReply = false;
                  this.sendQueuedToECU();
                } else {
                  this.waitReply = true;
                }
              } else {
                this.sendQueuedToECU();
              }
              if (dataframe.length >= 40) {
                this.debug("dataframe too long,");
                len_cmd = 0;
                dataframe = [];
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
          await this.wait(1000);
        }
      }
    },
    processCmd(cmd, data) {
      this.last_reponse = data.substring(2);
      switch (cmd) {
        case 0x00:
          // slow init echo 0x00 0x00 0x00
          break;
        case 0x55:
          this.ser.stage = 3;
          this.sendBytes([0x7c]);
          break;
        case 0x7c:
          clearInterval(this.ser.connectTimer);
          this.ser.connectTimer = null;
          this.sendBytes([0xca]);
          break;
        case 0xca:
          this.ser.stage = 4;
          this.sendBytes([0x75]);
          break;
        case 0x75:
          this.sendBytes([0xf4]);

          break;
        case 0xf4:
          this.debug("0xf4 echo");
          if (this.ser.stage == 4) this.sendBytes([0xd1]);
          break;
        case 0x7d:
          if (this.record.timer) this.sendToEcu([0x80]); // Trigger next dataframe
          this.parse7D(this.hexToBytes(data.substring(2)));
          this.Dataframe.Time = this.Time();
          this.Dataframe.Dataframe7d = data.substring(2);

          break;
        case 0x80:
          if (this.record.timer) this.sendToEcu([0x7d]); // Trigger next dataframe
          this.parse80(this.hexToBytes(data.substring(2)));
          this.Dataframe.Dataframe80 = data.substring(2);
          this.log.MemsData.push({
            Time: this.Dataframe.Time,
            Dataframe80: this.Dataframe.Dataframe80,
            Dataframe7d: this.Dataframe.Dataframe7d,
            Acceleration: this.Acceleration,
          });

          break;
        case 0xd0:
          this.sendBytes([0x7d]);
          this.parseD0(data.substring(2));
          break;
        case 0xd1:
          this.parseD1(data.substring(2));

          if (this.ser.stage === 4) {
            this.sendBytes([0x80]);
            //this.sendToEcu([0x80,0x79]);
          }
          this.ser.stage = 5;

          break;
        default:
          break;
      }
    },
    readRom() {
      var downloadNextAddress2j = 0x100000;
      var endOfDownload2j = 0x11ffff;
      var fullRomSize2j = endOfDownload2j - downloadNextAddress2j + 1;
      var fullRomData2j = [];
      let len_to_read = 31;
      this.sendBytes([0xdc, (downloadNextAddress2j >> 16) & 0xff, (downloadNextAddress2j >> 8) & 0xff, (downloadNextAddress2j >> 0) & 0xff, len_to_read]);
    },
  },
};
</script>

<template>
  <div class="fixed-top">
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="container-sm">
        <img style="height: 3rem" class="mr-2" src="./assets/lotus-badge.png" />
        <a class="navbar-brand" href="/">MEMS 1.9 ⚙️ </a>
        <div class="navbar-text">
          <small>{{ appVersion }}</small>
        </div>
        <div class="collapse navbar-collapse" id="navbarSupportedContent"></div>
        <div class="collapse navbar-collapse"></div>
        <div class="navbar-text">
          <span class="float-right">
            <label class="navbar-item"
              >ECU ID: <span class="badge badge-dark">{{ ECUID }}</span></label
            >
            <label class="ml-3"
              >ECU Serial: <span class="badge badge-dark">{{ ECUSerial }}</span></label
            >
          </span>
        </div>
      </div>
    </nav>
  </div>

  <div class="container">
    <div class="steps">
      <progress id="progress" :value="(ser.stage - 1) * 20" max="100"></progress>
      <div class="step-item">
        <button @click="openSerialPort" class="step-button text-center text-white" :class="ser.stage === 1 ? 'active' : ''">1</button>
        <div class="step-title">Start</div>
      </div>
      <div class="step-item">
        <button @click="slowInit(0x16)" class="step-button text-center text-white" :class="ser.stage >= 2 && ser.stage < 3 ? 'active' : ''">2</button>
        <div class="step-title">Slow Init</div>
      </div>
      <div class="step-item">
        <button class="step-button text-center text-white" :class="ser.stage === 3 ? 'active' : ''">3</button>
        <div class="step-title">Wake up</div>
      </div>
      <div class="step-item">
        <button class="step-button text-center text-white" :class="ser.stage === 4 ? 'active' : ''">4</button>
        <div class="step-title">Initialise</div>
      </div>
      <div class="step-item">
        <button class="step-button text-center text-white" :class="ser.stage === 5 ? 'active' : ''">5</button>
        <div class="step-title">Running</div>
      </div>
      <div class="step-item">
        <button @click="closeSerialPort" class="step-button text-center text-white" :class="ser.stage === 6 ? 'active' : ''">6</button>
        <div class="step-title">Stop</div>
      </div>
    </div>
  </div>
  
  <div class="card-group text-center">
    <div class="card" @keyup.enter="simulatePause()">
      <div class="card-body">
        <h6 class="card-title">Time</h6>
        <h3 class="card-text text-monospace">{{ stopwatchFormat }}</h3>

        <!--    Gear: {{ gear }} i:{{ Dataframe.IdleSwitch }}<br />
        Δ: {{ rpmD1 }} {{ rpmD2.val }}<br />
         {{ rpmD2.min }} <br> {{ rpmD2.max }}<br />
        MPH: {{ MPH }} KPH: {{ KPH }} <br />
        Δ {{ deltaDist }}m <br />
        <h3 class="card-text text-monospace">{{ Acceleration }}</h3>
     -->
     
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">RPM</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.EngineRPM }}</h3>
        <input class="custom-range" type="range" min="0" max="7500" :value="Dataframe.EngineRPM" />

        <h6 class="card-title">Throttle Angle</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.ThrottleAngle }}</h3>
        <input class="custom-range" type="range" min="0" max="100" :value="Dataframe.ThrottleAngle" />

        <h6 class="card-title">Ignition Advance</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.IgnitionAdvance }}</h3>
        <input class="custom-range" type="range" min="-20" max="40" :value="Dataframe.IgnitionAdvance" />
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Lambda {{ !Dataframe.ClosedLoop ? "Closed" : "Open" }}</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.LambdaVoltage }}</h3>
        <input class="custom-range" type="range" min="0" max="1200" :value="LambdaVoltageDamped" />
        <p>{{ AirFuelRatioLabel }} {{ AirFuelRatioCalc }}</p>

        <h6 class="card-title">Manifold Absolute Pressure</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.ManifoldAbsolutePressure }}</h3>
        <input class="custom-range" type="range" min="0" max="101" :value="Dataframe.ManifoldAbsolutePressure" />

        <h6 class="card-title">Short Term Fuel Trim</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.ShortTermFuelTrim }}</h3>
        <input class="custom-range" type="range" min="0" max="255" :value="Dataframe.ShortTermFuelTrim" />

        <!--
        Nm: {{ Nm }} <br />
        {{ gForce }}g<br />
        hp: {{ hp }}
        <br />
        ft/lb {{ ft_lb }}
        -->
      </div>
    </div>
  </div>

  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="download()">
    <i class="fas fa-download"></i>
    Download
  </button>

  <button v-if="!record.timer" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="recordStart()">
    <i class="fas fa-circle" style="color: red"></i>
    Record
  </button>

  <button v-if="record.timer" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="recordStop()">
    <i class="fas fa-stop" style="color: black"></i>
    Stop
  </button>

  <span v-if="faults.length">
    <label>Faults: </label>
    <span class="badge badge-dark ml-2" v-for="fault in faults" v-bind:key="fault">
      {{ fault }}
    </span>
  </span>
  <div class="btn-group float-right" role="group">
    wait: {{ waitReply }} {{ queuedBytes }}

    <button v-if="!replay.timer" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="replaySerialStart()">
      <i class="fas fa-refresh"></i>
      Replay Serial
    </button>

    <button v-if="!replay.timer" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="simulateStart()">
      <i class="fas fa-refresh"></i>
      Replay
    </button>

    <button v-if="replay.timer" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="simulateStop()">
      <i class="fas fa-stop"></i>
    </button>

    <button v-if="replay.timer && !replay.pause" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="simulateStep(-10)">
      <i class="fas fa-backward"></i>
    </button>

    <button v-if="replay.timer && !replay.pause" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="simulatePause()">
      <i class="fas fa-pause"></i>
    </button>

    <button v-if="replay.timer && !replay.pause" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="simulateStep(10)">
      <i class="fas fa-forward"></i>
    </button>

    <button v-if="replay.timer && replay.pause" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="simulatePause()">
      <i class="fas fa-play"></i>
    </button>
  </div>
  {{ last_reponse }}
  <hr />

  <div>
    
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x6b])">test RPM gauge 6B</button>

     <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x79])">0x79 + STFT</button>
      <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x7a])">0x79 - STFT</button>
  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x7b])">0x7b + LTFT</button>
  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x7c])">0x7b + LTFT</button>
 
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x80])">Data 80</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x7d])">Data 7D</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x80, 0x7d])">Data 80/7D</button>

    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xf4])">
      <i class="fa fa-heart"></i>
    </button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xca])">0xca init</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xd0])">ECU SER</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xd1, 0x80, 0x7d])">ECU ID/SER</button>

    <button class="btn btn-outline-danger btn-sm mr-2 mb-2 float-right" @click="sendToEcu([0xfa])">
      <i class="fa fa-undo"></i>
      Clear Adaptations
    </button>

    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2 float-right" @click="sendToEcu([0xcc])">Clear Faults</button>

    <p></p>

    <div class="btn-group mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x0d])">Off</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"><label class="mb-0">Radiator Fan</label></span>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x1d])">On</button>

      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x01])">Off</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"><label class="mb-0">Fuel pump</label></span>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x11])">On</button>
    </div>

    <div class="btn-group mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x92])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Idle Speed Δ <span class="ml-1 badge badge-dark">{{ Dataframe.IdleSpeedOffset }} / {{ Dataframe.IdleSetPoint }} </span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x91])">+</button>
    </div>

    <div class="btn-group mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x8a])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Idle Decay Δ <span class="ml-1 badge badge-dark">{{ Dataframe.IdleSpeedOffset }}</span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x89])">+</button>
    </div>

    <div class="btn-group mr-2 mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x7a,0x80,0x7d])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Long Term Fuel Trim <span class="ml-1 badge badge-dark">{{ Dataframe.LongTermFuelTrim }}</span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x7b,0x80,0x7d])">+</button>
    </div>

        <div class="btn-group mr-2 mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x7a,0x80,0x7d])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Short Term Fuel Trim <span class="ml-1 badge badge-dark">{{ Dataframe.ShortTermFuelTrim }}</span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x7b,0x80,0x7d])">+</button>
    </div>


    <div class="btn-group mt-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x93])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Ignition Advance Δ <span class="ml-1 badge badge-dark">{{ Dataframe.IgnitionAdvanceOffset80 }}</span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x94])">+</button>
    </div>

    <div class="btn-group mt-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0xfe])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >IAC Position Δ <span class="ml-1 badge badge-dark">{{ Dataframe.IACPosition }}</span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0xfd])">+</button>
    </div>

    <div class="btn-group mt-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0xf3])">Diag 0x4</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0xf5])">Diag 0x3</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0xf0])">Read Diag mode</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0xf7])">Read Calibration</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendBytes([0xdc, 0x00])">set block 00</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendBytes([0x70])">Read Block 0x70</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendBytes([0x80])">Read Block 0x80</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendBytes([0xc7])">Rom bank</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendBytes([0xc8])">Ram bank 1</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendBytes([0xc9])">Ram bank 2</button>
  <button type="button" class="btn btn-sm btn-outline-secondary" @click="readRom()">ReadRom</button>
    </div>
  </div>
  <hr />
  <div class="row">
    <div class="col-6">
      <pre style="overflow-y: scroll; height: 25vh">{{ debug_log.join("\n") }}</pre>
    </div>
    <div class="col-6">
      <pre style="overflow-y: scroll; height: 25vh">{{ ser.replysRaw.join("\n") }}</pre>
    </div>
  </div>
  <hr />
  <Chart :key="chartKey" :margin="{ top: 10, right: 10, bottom: 10, left: 10 }" :size="chartSize" :data="history" direction="horizontal">
    <template #layers>
      <Grid strokeDasharray="2,2" />
      <Line :dataKeys="['Time', 'EngineRPM']" type="natural" />
      <Line :dataKeys="['Time', 'LambdaVoltage']" type="natural" :lineStyle="{ stroke: 'red' }" />
      <Line :dataKeys="['Time', 'AFR']" type="natural" :lineStyle="{ stroke: 'yellow' }" />
    </template>

    <template #widgets>
      <Tooltip
        borderColor="#48CAE4"
        :config="{
          Time: { hide: true },
          EngineRPM: { color: '#0077b6' },
          AFR: { color: 'yellow' },
          LambdaVoltage: { label: 'Lambda Voltage', color: 'red' },
          gear: { label: 'gear', color: 'black' },
        }"
      />
    </template>
  </Chart>

  <div class="card-columns mt-4">
    <div class="card" v-for="param in parameters" v-bind:key="param">
      <div class="card-body pt-0 pb-0">
        {{ param }}
        <span class="badge badge-dark text-monospace float-right mt-1">{{ Dataframe[param] }}</span>
      </div>
    </div>
  </div>
  <hr />
</template>
<style lang="scss">
.card-columns {
  column-count: 4;
}

.custom-range::-webkit-slider-thumb {
  width: 2px;
}

#progress {
  -webkit-appearance: none;
  position: absolute;
  width: 95%;
  z-index: 5;
  height: 10px;
  margin-left: 18px;
  margin-bottom: 18px;
}

.steps {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  position: relative;
}

.step-button {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  border: none;
  background-color: var(--gray);
  transition: 0.4s;
}

.step-button[aria-expanded="true"],
.step-button.active {
  width: 60px;
  height: 60px;
  background-color: var(--blue);
  color: #fff;
}

.step-item {
  z-index: 10;
  text-align: center;
}
</style>

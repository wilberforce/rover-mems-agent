<script lang="ts">
//import imported_data from "./data/run-demo.fcr.json";
//import imported_data from "./data/run-1649820449532.fcr.json";
import imported_data from "./data/run-1649822529717.fcr.json";

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
      rpmD1: 0,
      rpmD2: { min: [], max: [], val: 0 },
      weightKg: 820,
      Nm: 0,
      gForce: 0,
      dist: 0,
      lastDist: 0,
      deltaDist: 0,
      gear: 0,
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
        Dataframe80: "801c000000000000000000000000000000000000000000000000000000",
        Dataframe7d: "7d210000000000000000000000000000000000000000000000000000000000000000",
      },
      parameters: [
        "EngineRPM",
        "CoolantTemp",
        "AmbientTemp",
        "IntakeAirTemp",
        "FuelTemp",
        "ManifoldAbsolutePressure",
        "BatteryVoltage",
        "ThrottlePosition",
        "IdleSwitch",
        "DTC0",
        "DTC1",
        "IdleSetPoint",
        "IdleHot",
        "IACPosition",
        "IdleSpeedDeviation",
        "IgnitionAdvanceOffset80",
        "IgnitionAdvance",
        "CoilTime",
        "CrankshaftPositionSensor",
        "IgnitionSwitch",
        "ThrottleAngle",
        "AirFuelRatio",
        "LambdaVoltage",
        "LambdaStatus",
        "ClosedLoop",
        "LongTermFuelTrim",
        "ShortTermFuelTrim",
        "FuelTrimCorrection",
        "CarbonCanisterPurgeValve",
        "IdleBasePosition",
        "AirconSwitch",
        "ParkNeutralSwitch",
        "DTC2",
        "DTC3",
        "DTC4",
        "DTC5",
        "IdleSpeedOffset",
        "IgnitionAdvanceOffset7d",
        "JackCount",
        "LambdaDutycycle",
        "LambdaFrequency",
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
      debug_log: [],
    };
  },
  computed: {
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
            // = Math.min(rpmD2, this.rpmD2.min);
          }
        }
        this.rpmD1 = rpmD1;
        this.rpmD2.val = rpmD2;
        this.lastEngineRPM = EngineRPM;
        this.lastDist = this.dist;
        this.history.push({
          Time: this.stopwatchFormat,
          EngineRPM: this.Dataframe.EngineRPM,
          LambdaVoltage: this.Dataframe.LambdaVoltage,
          rpmD2: rpmD2,
          gear: this.gear,
        });
      }
    },
  },

  mounted() {},
  async unmounted() {
    //await this.closeSerialPort();
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
      // Slider for speed - realtime option
    },
    async replaySerialStart(step = 0) {
      this.replay.port = await navigator.serial.requestPort();

      await this.replay.port.open({
        baudRate: 9600,
        databits: 8,
        bufferSize: 16,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });

      this.replay.step = step;

      this.replay.reader = this.replay.port.readable.getReader();

      while (true) {
        const { value, done } = await this.replay.reader.read();
        if (value) {
          let hex = this.hex(Array.from(value));
          switch (value[0]) {
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
            default: //Echo
              this.replayWrite(hex);
              console.log(`cmd: ${hex}`);
              break;
          }
        }
        if (done) {
          console.log("Serial replay Release Lock", done);
          this.replay.reader.releaseLock();
          break;
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
    replayWrite(data) {
      if (this.replay.port) {
        data = data[0] + data[1] + data;

        let bytes = new Uint8Array(this.hexToBytes(data));
        let hex = this.hex(Array.from(bytes));

        this.debug(`<<ecu ${hex}`);

        let writer = this.replay.port.writable.getWriter();
        writer.write(Uint8Array.from(bytes));
        writer.releaseLock();
      }
    },
    replaySerial() {
      if (imported_data.MemsData.length >= this.replay.step) {
        let data = imported_data.MemsData[this.replay.step];
        if (data && this.replay.port) {
          this.replayWrite(data.Dataframe7d);
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
      this.record.timer = setInterval(() => this.pollDataframes(), 500);
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
          //IdleSpeedOffset: v.getInt8(0x13 + offset),
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
        this.debug(`expected len 28 (${len}) for 0x80`);
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
      for (let c = 0; c < hex.length; c += 2) view.setUint8(c / 2, parseInt(hex.substr(c, 2), 16));
      return bytes; //new Uint8Array(bytes);
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
      if (!this.waitReply) {
        this.sendBytes(bytes);
        this.waitReply = true;
      } else {
        this.queuedBytes.push(...bytes);
        if (this.queuedBytes.length > 2) {
          this.debug("cleared queue");
          this.waitReply = false;
          this.queuedBytes = [];
        }
      }
    },

    async sendBytes(bytes) {
      let hex = this.hex(Array.from(bytes));
      if (this.ser.port) {
        let writer = this.ser.port.writable.getWriter();
        writer.write(Uint8Array.from(bytes));
        writer.releaseLock();
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
      if (this.ser.port) await this.ser.port.close();
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
      this.debug(`Connecting to MEMS 1.9 ECU at address ${ecuAddress.toString(16)}`);

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

    async newInit5baud() {
      this.ser.port = await navigator.serial.requestPort();

      let ecuAddress = 0x16;
      this.debug(`5 Baud..Connecting to MEMS 1.9 ECU at address ${ecuAddress.toString(16)}`);
      //this.debug(this.ser.port.getInfo());

      await this.ser.port.open({
        baudRate: 600,
        databits: 8,
        parity: "even",
        stopbits: 1,
        flowControl: "none",
      });

      ecuAddress = (ecuAddress << 1) | 1;

      let bits = ecuAddress
        .toString(2)
        .padStart(10, 0)
        .split("")
        .reverse()
        .map((x) => {
          return x * 0xff;
        });

      let start = performance.now();

      this.debug("Clear line");
      await this.ser.port.setSignals({ brk: false, break: false });
      await this.wait(2000);

      for (let i = 0; i < 10; i++) {
        this.sendToEcu([bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i], bits[i]]);
      }
      this.debug(`time: ${performance.now() - start}\n`);
      const reader = this.ser.port.readable.getReader();

      setTimeout(() => {
        this.debug("canceling...");
        reader.cancel();
      }, 2200);

      while (true) {
        const { value, done } = await reader.read();
        if (value) {
          console.log(value);
        }
        if (done) {
          console.log("[readLoop] DONE", done);
          reader.releaseLock();
          break;
        }
      }

      this.debug("closing...");
      await this.ser.port.close();
      this.debug("closed");

      await this.ser.port.open({
        baudRate: 9600,
        databits: 8,
        bufferSize: 128,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });

      this.baud5listen();
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
            this.ser.buffer = this.ser.buffer.concat(this.hex(Array.from(value)));
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
    CmdLength(cmd, dataframe, len_cmd) {
      switch (cmd) {
        case 0x00:
          return 3; //psuedo command - 5 baud echo expect 0x00 0x00 0x00 then reply

        case 0x80:
          if (dataframe.length > 2) return dataframe[2] + 2;
          return len_cmd;
        case 0x7d:
          if (dataframe.length > 2) {
            return 34; // Why?
            return dataframe[2] + 2;
          } // Need to handle case of single byte
          return len_cmd;
        case 0xd0:
          return 6;
        case 0xd1:
          return 38;
        default:
          console.log(`Cmd: ${cmd.toString(16).padStart(2, "0")}`);
          return 2;
      }
    },
    async openSerialPort() {
      this.ser.port = await navigator.serial.requestPort();

      await this.ser.port.open({
        baudRate: 9600,
        databits: 8,
        //bufferSize: 60,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });
      this.sendToEcu([0xd1]);
      this.pollDataframes();

      this.waitReply = false;
      while (this.ser.port?.readable) {
        this.ser.reader = this.ser.port.readable.getReader();
        /*
        if (!this.waitReply && this.queuedBytes.length) {
          this.expectingBytes = this.queuedBytes.pop();
          this.debug(`sending queued bytes: ${this.expectingBytes.toString(16)}`);
          this.sendBytes([this.expectingBytes]);
        }
        */
        let cmd = 0;
        let len_cmd = 0;
        let dataframe: ArrayBuffer[] = [];
        let going = true;
        try {
          while (going) {
            const { value, done } = await this.ser.reader.read();
            if (done) {
              console.log("done");
              going = false;
              break;
            }
            let inbound = Array.from(value);
            let hex = this.hex(inbound);
            if (len_cmd == 0) {
              cmd = inbound[0];
              dataframe = inbound;
              len_cmd = this.CmdLength(cmd, dataframe, len_cmd);
            } else {
              let required = len_cmd - dataframe.length;
              if (inbound.length > required) {
                let rest = inbound.slice(required);
                inbound = inbound.slice(0, required);
                dataframe.push(...inbound);
                inbound = rest;
                cmd = inbound[0];
                len_cmd = this.CmdLength(cmd, dataframe, len_cmd);
              } else {
                dataframe.push(...inbound);
              }
            }
            this.ser.dataframe = dataframe;
            if (dataframe.length == len_cmd) {
              this.waitReply = false; // Allow commands to be sent
              let data = this.hex(dataframe);
              switch (cmd) {
                case 0x00:
                  this.sendToEcu([0x55, 0x76, 0x83]);
                  break;
                case 0xca:
                  this.sendToEcu([0x75]);
                  break;
                case 0x75:
                  this.sendToEcu([0xf4]);
                  break;
                case 0xf4:
                  this.sendToEcu([0xd0]);
                  break;
                case 0x7d:
                  if (data.length > 4) {
                    this.sendToEcu([0x80]); // Trigger next dataframe
                    this.parse7D(this.hexToBytes(data.substring(2)));

                    this.Dataframe.Time = this.Time();
                    this.Dataframe.Dataframe7d = data.substring(2);
                  } else {
                    console.log(`7d: short! ${data.length}`);
                    if (data.length == 2) {
                      this.debug("Lost connection...");
                      len_cmd = 0;
                      cmd = 0x00;
                    }
                  }
                  break;
                case 0x80:
                  if (data.length > 4) {
                    this.parse80(this.hexToBytes(data.substring(2)));
                    this.Dataframe.Dataframe80 = data.substring(2);
                    this.log.MemsData.push({
                      Time: this.Dataframe.Time,
                      Dataframe80: this.Dataframe.Dataframe80,
                      Dataframe7d: this.Dataframe.Dataframe7d,
                    });
                  } else {
                    console.log(`80: short! ${data.length}`);
                    if (data.length == 2) {
                      this.debug("Lost connection...");
                      len_cmd = 0;
                      cmd = 0x00;
                    }
                  }
                  break;
                case 0xd0:
                  this.parseD0(data.substring(2));
                  break;
                case 0xd1:
                  this.parseD1(data.substring(2));
                  break;
                default:
                  break;
              }
              len_cmd = 0;
              dataframe = [];
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
    async sleepUntil(timestampMs) {
      let now = new Date().getTime();
      let sleepFor = timestampMs - now;
      await new Promise((resolve) => setTimeout(resolve, sleepFor));
    },
    async slowInit19() {
      ecuAddress = 0x16; // 22

      //resetTimeout(5000);
      dataBuffer = [];
      this.debug("Attempting ECU connection... (address: " + ecuAddress + ") (slow init)");

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

      //parseDataBufferSlowInit();
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
      await this.ser.port.setSignals({ break: true });

      let i = 0;

      ecuAddress = (ecuAddress << 1) | 1;

      let bits = ecuAddress.toString(2).padStart(10, 0).split("").reverse();
      let start = performance.now();
      let last = 0;
      for (let i = 0; i < 10; i++) {
        await this.ser.port.setSignals({ break: !bits[i] });
        let now = performance.now();
        times.push(now - start);
        start = now;
        await this.wait(sleepMs);
      }
      this.debug(times);

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
  <div class="fixed-top">
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="container-sm">
        <img style="height: 3rem" class="mr-2" src="./assets/lotus-badge.png" />
        <a class="navbar-brand" href="/">MEMS 1.9 Diagnostic ⚙️ </a>
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

  <div class="card-group text-center">
    <div class="card" @keyup.enter="simulatePause()">
      <div class="card-body">
        <h6 class="card-title">Time</h6>
        <h3 class="card-text text-monospace">{{ stopwatchFormat }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">RPM</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.EngineRPM }}</h3>
        <!-- Gear: {{ gear }} i:{{ Dataframe.IdleSwitch }}<br />
       Δ: {{ rpmD1 }} {{ rpmD2.val }}<br>
      {{ rpmD2.min }} <br> {{ rpmD2.max }}<br />
        MPH: {{ MPH }} KPH: {{ KPH }} <br />
        Δ {{ deltaDist }}m <br />
         -->
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Lambda {{ !Dataframe.ClosedLoop ? "Closed" : "Open" }}</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.LambdaVoltage }}</h3>
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

  <div v-if="0" class="card-group text-center mt-2">
    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Short Term Fuel Trim</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.ShortTermFuelTrim }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Ignition Advance</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.IgnitionAdvance }}</h3>
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Battery Voltage</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.BatteryVoltage }}</h3>
      </div>
    </div>
  </div>

  <!--    <VueSvgGauge
  :start-angle="-110"
  :end-angle="110"
  :value="3"
  :separator-step="1"
  :min="0"
  :max="4"
  :gauge-color="[{ offset: 0, color: '#347AB0'}, { offset: 100, color: '#8CDFAD'}]"
  :scale-interval="0.1"
/>
-->
  <p v-if="faults.length">
    <label>Faults: </label>
    <span class="badge badge-dark ml-2" v-for="fault in faults" v-bind:key="fault">
      {{ fault }}
    </span>
  </p>

  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="openSerialPort">Open Serial Port</button>

  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="closeSerialPort()">Disconnect</button>
  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="newInit()">5 Baud Init</button>

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

  <button v-if="replay.timer && !replay.pause" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="simulatePause()">
    <i class="fas fa-pause"></i>
  </button>

  <button v-if="replay.timer && replay.pause" class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="simulatePause()">
    <i class="fas fa-play"></i>
  </button>

  Wait: {{ waitReply }} {{ JSON.stringify(queuedBytes) }}<br />
  {{ hex(ser.dataframe) }}
  <hr />
  <div>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x80])">Data 80</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x7d])">Data 7D</button>

    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xf4])">
      <i class="fa fa-heart"></i>
    </button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xca])">0xca init</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xd0])">ECU SER</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xd1])">ECU ID/SER</button>

    <button class="btn btn-outline-danger btn-sm mr-2 mb-2 float-right" @click="sendToEcu([0xfa])">
      <i class="fa fa-undo"></i>
      Clear Adaptations
    </button>

    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2 float-right" @click="sendToEcu([0xcc])">Clear Faults</button>

    <p></p>

    <div class="btn-group mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x0d, 0x01])">Off</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"><label class="mb-0">Radiator Fan</label></span>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x1d, 0x00])">On</button>
    </div>

    <div class="btn-group mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x92])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Idle Speed Δ <span class="ml-1 badge badge-dark">{{ Dataframe.IdleSpeedOffset }}</span></label
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

    <div class="btn-group mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x7a])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Long Term Fuel Trim <span class="ml-1 badge badge-dark">{{ Dataframe.LongTermFuelTrim }}</span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x7b])">+</button>
    </div>

    <p></p>

    <div class="btn-group mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x94])">-</button>
      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >Ignition Advance Δ <span class="ml-1 badge badge-dark">{{ Dataframe.IgnitionAdvanceOffset80 }}</span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x93])">+</button>
    </div>
  </div>

  <Chart :key="chartKey" :margin="{ top: 10, right: 10, bottom: 10, left: 10 }" :size="chartSize" :data="history" direction="horizontal">
    <template #layers>
      <Grid strokeDasharray="2,2" />
      <Line :dataKeys="['Time', 'EngineRPM']" type="natural" />
      <Line :dataKeys="['Time', 'LambdaVoltage']" type="natural" :lineStyle="{ stroke: 'red' }" />
      <Line :dataKeys="['Time', 'rpmD2']" type="natural" :lineStyle="{ stroke: 'green' }" />
    </template>

    <template #widgets>
      <Tooltip
        borderColor="#48CAE4"
        :config="{
          Time: { hide: true },
          EngineRPM: { color: '#0077b6' },
          LambdaVoltage: { label: 'Lambda Voltage', color: 'red' },
          rpmD2: { label: 'd2', color: 'green' },
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

  <label
    ><span class="badge badge-light">{{ Dataframe.Dataframe7d.length }} {{ Dataframe.Dataframe7d }}</span></label
  >
  <label
    ><span class="badge badge-light">{{ Dataframe.Dataframe80.length }} {{ Dataframe.Dataframe80 }}</span></label
  >
  <br />

  <pre style="overflow-y: scroll; height: 20vh">{{ debug_log.join("\n") }}</pre>
</template>
<style lang="scss">
.card-columns {
  column-count: 4;
}
</style>

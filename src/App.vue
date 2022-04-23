<script lang="ts">
//import imported_data from "./data/run-demo.fcr.json";
import imported_data from "./data/run-1649820449532.fcr.json";
//import imported_data from "./data/run-1649822529717.fcr.json";

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
        interval: 333,
        port: null,
      },
      record: {
        timer: null,
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

        //Dataframe7d:
        //"7d201016ff92401bffff010179670cff51ffff3588840dff134015801a0029c02a",
        //"7d201014ff924057ffff0100806400ff64ffff3080800eff16801b00220031c01f",
        //"7d21400aff910057ffff0100716400005aff002e807b690022002200220022000100"
        //Dataframe80:
        //"801c08ba80ff56ff1f842308000100002037876b0417055c05df100000",
        //"801c00006fff4fff64781b00000100002037877b055f05380ca5000000"
        Dataframe80: "801c000000000000000000000000000000000000000000000000000000",
        Dataframe7d: "7d2000000000000000000000000000000000000000000000000000000000000000",
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
        // to do fix this calc - use moment?
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

  mounted: function () {
    //this.simulateStart(0);
    //this.dumpImportReadmemsHex();
    //this.parseD1("d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c483356303035c70005cb");
    //this.parseD0("d04b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630");
  },
  methods: {
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
      // this.parseD1( "d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630");
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
        bufferSize: 128,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });
      // this.parseD1( "d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c48335630");
      this.replay.step = step;
      this.replay.timer = setInterval(() => this.replaySerial(), this.replay.interval);
      this.replaySerial();
      this.ECUID = imported_data.Name;
      // Slider for speed - realtime option
    },

    replaySerial() {
      if (imported_data.MemsData.length >= this.replay.step) {
        let data = imported_data.MemsData[this.replay.step];
        if (data && this.replay.port) {
          let x7d = "7d21" + data.Dataframe7d.substring(4).padEnd(66, "0");

          let writer = this.replay.port.writable.getWriter();
          writer.write(this.hexToBytes(x7d));
          //writer.write(this.hexToBytes(data.Dataframe80));
          writer.releaseLock();

          this.replay.step++;
        } else {
          this.replaySerialStop();
        }
      }
    },
    replaySerialStop() {
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
      this.record.timer = setInterval(() => this.pollDataframes(), 500);
    },
    recordStop() {
      clearInterval(this.record.timer);
      this.record.timer = null;
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
      //this.queuedBytes.push(bytes);
      this.sendBytes(bytes);
    },

    async sendBytes(bytes) {
      let hex = this.hex(Array.from(bytes));
      this.debug(`>> ${hex}`);
      if (this.ser.port) {
        //this.waitReply = true;
        let writer = this.ser.port.writable.getWriter();
        writer.write(Uint8Array.from(bytes));
        writer.releaseLock();
      }
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
    /*
    d0d0c70005cb7d7d21400aff910058ffff0100706400007bff002e807b1e0023
App.vue:590 default:
App.vue:590 << c023c023c023c00000
*/
    async openSerialPortChunked() {
      this.ser.port = await navigator.serial.requestPort();

      await this.ser.port.open({
        baudRate: 9600,
        databits: 8,
        bufferSize: 35,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });
      this.debug(this.ser.port);

      this.sendToEcu([0xd0]);

      this.pollDataframes();
      //https://stackoverflow.com/questions/70920727/read-usb-serial-port-data-with-web-serial-api-in-javascript

      // https://streams.spec.whatwg.org/#rs-intro
      while (this.ser.port.readable) {
        const reader = this.ser.port.readable.getReader(); //({ mode: "byob" });

        let startingAB = new Uint8Array(35);
        const buffer = await readInto(startingAB);
        let hex = this.hex(new Uint8Array(buffer));
        this.debug(`<< ${hex}`);

        reader.releaseLock();

        async function readInto(buffer) {
          let offset = 0;
          //console.log( offset, buffer.byteLength)

          while (offset < buffer.byteLength) {
            const { value: view, done } = await reader.read(new Uint8Array(buffer, offset, buffer.byteLength - offset));
            buffer = view.buffer;
            if (done) {
              console.log("done");
              break;
            }
            offset += view.byteLength;
          }
          //console.log( offset, buffer.byteLength)
          return buffer;
        }
      }
    },
    /*
    >> 7d
<< 0000500021400091057001070640000008000023
<< 023023023000
>> 7d
<< 0021400091057001070640000008000023023023023
<< 000

"Dataframe80":"
801c05ff8b9a72ff1786331080011000041b873202a7025b083e100000",
"Dataframe7d":"
7d214018ff910052ffff0101717800041bff002e7f82aa00238023802380238001"},
>> 80
<< 80800000000647221080110002886854023000000
>> 7d
<< 00214000910580010706400000080
<< 00023023023023
<< 000

>> 80
<< 80800000000647221080110002886854023000000
>> 7d
<< 0021400091057001070640000008000023023023023
<< 000
>> 80
<< 80800000000647221080110002886854023000000
>> 7d
<< 00214000910570010706400000
<< 08000023023023023
<< 000
>> 80
<< 80800000000647221080110002886854023000000
>> 7d
<< 0
<< 021400091058001070640000008000023023023023
<< 000
>> d0
<< 000050
>> d1
<< 00004833563030350050004833563030350050004833563030350050
>> d0
<< 0000
<< 50
>> d1
<< 00004833563030350050004833563030350050004833563030350050
>> d1
<< 000048335630303500500048335630303500500048335630
<< 30350050
>> f4
<< 000
>> f4
<< 000
>> d1
<< 000048335630303500500048335630303500500048335630
<< 30350050
>> d1
<< 00004833
<< 5630303500500048335630303500500048335630
<< 30350050
>> d0
<< 000050
>> 80
<< 80800000000647221080110002886854023000000
>> 7d
<< 0021400091057001070640000008000023023023023
<< 000
>> 7d80
<< 08080000000064722108011000288685402
<< 3000000
>> f4
<< 000
*/
    async openSerialPort() {
      navigator.serial.addEventListener("connect", (event) => {
        console.log({ e: event });
        debug;
      });
      this.ser.port = await navigator.serial.requestPort();
      this.debug(this.ser.port.getInfo());
      let ports = await navigator.serial.getPorts();

      await this.ser.port.open({
        baudRate: 9600,
        databits: 8,
        bufferSize: 128,
        parity: "none",
        stopbits: 1,
        flowControl: "none",
      });
      this.debug(this.ser.port);

      this.sendToEcu([0xd0]);

      this.pollDataframes();

      let start = null;
      while (this.ser.port.readable) {
        this.ser.reader = this.ser.port.readable.getReader();
        this.debug("waiting on data...");

        //https://stackoverflow.com/questions/70920727/read-usb-serial-port-data-with-web-serial-api-in-javascript
        try {
          while (true) {
            if (!this.waitReply && this.queuedBytes.length) {
              this.expectingBytes = this.queuedBytes.pop();
              this.sendBytes(this.expectingBytes);
            }
            const { value, done } = await this.ser.reader.read();
            this.waitReply = 0;

            if (done) {
              this.debug("serial this.ser.reader done");
              return;
            }

            this.ser.buffer = this.ser.buffer.concat(this.hex(Array.from(value)));
            if (start === null) start = value[0];
            switch (start) {
              case 0x1b:
                // 1b5b41 up arrow key test com port
                this.debug(`Arrow << ${this.ser.buffer}`);
                this.ser.buffer = "";
                start = null;
                break;
              case 0x80:
                this.expectedBytes = 60;
                if (this.ser.buffer.length < this.expectedBytes) {
                  this.debug(`expected 60 (${this.ser.buffer.length}) bytes for 0x80`);
                  break;
                }
                if (this.ser.buffer.length != 60) {
                  this.ser.buffer = "";
                  start = null;
                  this.debug(`rejected << ${this.ser.buffer}`);
                  return;
                }

                this.Dataframe.Dataframe80 = this.ser.buffer.substring(2, 60);
                // Save rest of buffer

                this.log.MemsData.push({
                  Time: this.Dataframe.Time,
                  Dataframe80: this.Dataframe.Dataframe80,
                  Dataframe7d: this.Dataframe.Dataframe7d,
                });
                this.parse80(this.hexToBytes(this.ser.buffer));
                this.ser.buffer = this.ser.buffer.substring(60);
                /*
                if (this.ser.buffer.length) {
                  value = parseInt(this.ser.buffer.substring(60), 2);
                } else start = null;
                */
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

                this.debug(`1.9 ECU woke up - init stage 1, << ${this.ser.buffer} >> ca`);
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
                this.Dataframe.Dataframe7d = this.ser.buffer;
                let now = new Date();
                let ms = now.getMilliseconds().toString(10).padStart(3, "0");
                this.Dataframe.Time = `${now.toLocaleTimeString()}.${ms}`;
                this.parse7D(this.hexToBytes(this.ser.buffer));
                //this.sendToEcu([0x80]); // trigger next frame
                this.ser.buffer = "";
                start = null;
                break;
              case 0xd0:
                if (this.ser.buffer.length < 4) {
                  this.debug(`expected 64 bytes for 0xd0, got ${this.ser.buffer.length}`);
                  break;
                }
                this.debug(`0xd0 << ${this.ser.buffer.length}`);
                this.ser.buffer = this.ser.buffer.substring(2);
                this.parseD0(this.ser.buffer);
                this.ser.buffer = "";
                start = null;
                break;

              case 0xd1:
                // d14b4c483356303035c70005cb4b4c483356303035c70005cb4b4c483356303035c70005cb 70!
                if (this.ser.buffer.length < 60) {
                  this.debug(`expected 60+ bytes for 0xd1, got ${this.ser.buffer.length}`);
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
        <h6 v-if="0" class="card-title">RPM</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.EngineRPM }}</h3>
        Gear: {{ gear }} i:{{ Dataframe.IdleSwitch }}<br />
        <!--Δ: {{ rpmD1 }} {{ rpmD2.val }}<br>
      {{ rpmD2.min }} <br> {{ rpmD2.max }}<br /> -->
        MPH: {{ MPH }} KPH: {{ KPH }} <br />
        Δ {{ deltaDist }}m <br />
      </div>
    </div>

    <div class="card">
      <div class="card-body">
        <h6 class="card-title">Lambda {{ !Dataframe.ClosedLoop ? "Closed" : "Open" }}</h6>
        <h3 class="card-text text-monospace">{{ Dataframe.LambdaVoltage }}</h3>
        Nm: {{ Nm }} <br />
        {{ gForce }}g<br />
        hp: {{ hp }}
        <br />
        ft/lb {{ ft_lb }}
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

  <p v-if="faults.length">
    <label>Faults: </label>
    <span class="badge badge-dark ml-2" v-for="fault in faults" v-bind:key="fault">
      {{ fault }}
    </span>
  </p>

  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="openSerialPortChunked()">Open Serial Port new</button>

  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="openSerialPort()">Open Serial Port</button>

  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="closeSerialPort()">Disconnect</button>
  <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="newInit5baud()">5 Baud Init</button>

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

  <hr />
  <div>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x7d, 0x80])">All Data</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x80])">Data 80</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0x7d])">Data 7D</button>

    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xf4])">
      <i class="fa fa-heart"></i>
    </button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xf3])">Mode 4</button>
    <button class="btn btn-outline-secondary btn-sm mr-2 mb-2" @click="sendToEcu([0xf0])">Current Mode (14 - mode 3)</button>
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

    <div class="btn-group mr-2" role="group">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x81])">clearing (auto?) adjustments 0x81</button>
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x7f])">- 0x7f</button>

      <span type="button" class="btn btn-sm btn-outline-secondary disabled"
        ><label class="mb-0"
          >IgnitionAdvanceOffset7d ?? Δ <span class="ml-1 badge badge-dark">{{ Dataframe.IgnitionAdvanceOffset7d }} </span></label
        ></span
      >
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="sendToEcu([0x7e])">+ 0x7e</button>
    </div>
  </div>

  <pre>{{ debug_log.join("\n") }}</pre>

  <div class="card-columns mt-4">
    <div class="card" v-for="param in parameters" v-bind:key="param">
      <div class="card-body pt-0 pb-0">
        {{ param }}
        <span class="badge badge-dark text-monospace float-right">{{ Dataframe[param] }}</span>
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

  <pre>{{ debug_log.join("\n") }}</pre>
</template>
<style lang="scss">
.card-columns {
  column-count: 6;
}
</style>

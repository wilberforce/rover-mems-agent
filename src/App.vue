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
        Dataframe80:
          "801c000085ff4fff638e23001001000000208b60039d003808c1000000",
        Dataframe7d:
          "7d201012ff92006effff0100996400ff3affff30807c63ff19401ec0264034c008",
      },
    };
  },
  methods: {
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

      log = {
        Name: "testdata/nofaults.fcr",
        Count: 338,
        Date: "0000-01-01T12:35:55.186Z",
        Summary: "Scenario file created from testdata/nofaults.csv",
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
          {
            Time: "12:35:56.224",
            Dataframe7d:
              "7d201010ff92401dffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:56.527",
            Dataframe7d:
              "7d201010ff92401dffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:56.990",
            Dataframe7d:
              "7d201010ff92401dffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:57.470",
            Dataframe7d:
              "7d201010ff92401dffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:57.932",
            Dataframe7d:
              "7d201010ff92401dffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:58.379",
            Dataframe7d:
              "7d201010ff92401dffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:58.858",
            Dataframe7d:
              "7d201010ff92401effff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:59.289",
            Dataframe7d:
              "7d201010ff92401effff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:35:59.816",
            Dataframe7d:
              "7d201010ff92401effff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:36:00.550",
            Dataframe7d:
              "7d201010ff92401effff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:36:01.013",
            Dataframe7d:
              "7d201010ff92401effff0100796400ff6fffff35887a88ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:36:01.460",
            Dataframe7d:
              "7d201010ff92401effff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:36:01.923",
            Dataframe7d:
              "7d201010ff92401fffff0100796400ff70ffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:36:02.354",
            Dataframe7d:
              "7d201010ff92401fffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b057805380ca5000000",
          },
          {
            Time: "12:36:02.801",
            Dataframe7d:
              "7d201010ff92401fffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:36:03.264",
            Dataframe7d:
              "7d201010ff92401fffff0100796400ff6fffff35887aa1ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff64781b00000100002037877b055f05380ca5000000",
          },
          {
            Time: "12:36:03.743",
            Dataframe7d:
              "7d201010ff924021ffff0100796400ff6fffff35887c62ff134015801a0029c02a",
            Dataframe80:
              "801c00006fff4fff58761b00000100002037877b057805380d08100000",
          },
          {
            Time: "12:36:04.270",
            Dataframe7d:
              "7d20100fff924047ffff0100796400ff6fffff35887f45ff134015801a0029c02a",
            Dataframe80:
              "801c03d66fff4fff32741b00000100002037877b00a7053807af100000",
          },
          {
            Time: "12:36:04.750",
            Dataframe7d:
              "7d20100fff924062ffff0100796400ff6fffff35887df6ff134015801a0029c02a",
            Dataframe80:
              "801c03df6fff4fff24741a00000100002037877b01c405380671100000",
          },
          {
            Time: "12:36:05.244",
            Dataframe7d:
              "7d201010ff924060ffff0100796400ff6fffff35887df8ff134015801a0029c02a",
            Dataframe80:
              "801c03446fff4fff2c761a00000100002037877b021805380674100000",
          },
          {
            Time: "12:36:05.595",
            Dataframe7d:
              "7d201011ff92405effff0100796400ff6fffff35887e59ff134015801a0029c02a",
            Dataframe80:
              "801c035f6fff4fff2d771c00000100002037877b01e20538064f100000",
          },
          {
            Time: "12:36:06.154",
            Dataframe7d:
              "7d201011ff92405effff0100796400ff6fffff35887ea6ff134015801a0029c02a",
            Dataframe80:
              "801c03f06fff4fff2b781d00000100002037877b01680538063d100000",
          },
          {
            Time: "12:36:06.968",
            Dataframe7d:
              "7d201010ff92405effff0100796400ff6fffff35887ea4ff134015801a0029c02a",
            Dataframe80:
              "801c03f66fff4fff2a7a1c00000100002037877b01600538062c100000",
          },
          {
            Time: "12:36:07.463",
            Dataframe7d:
              "7d20100fff92405effff0100796400ff6fffff35887e5cff134015801a0029c02a",
            Dataframe80:
              "801c03df6fff4fff257b1a00000100002037877b019605380621100000",
          },
          {
            Time: "12:36:08.309",
            Dataframe7d:
              "7d201011ff92405dffff0100796400ff6fffff35887ed7ff134015801a0029c02a",
            Dataframe80:
              "801c03a76fff4fff297d1b000001000020378781018b054c0617100000",
          },
          {
            Time: "12:36:08.804",
            Dataframe7d:
              "7d201011ff92405effff0100796400ff6fffff35887f59ff134015801a0029c02a",
            Dataframe80:
              "801c048b6fff4fff247e1c00000100002037878100bd054c0613100000",
          },
          {
            Time: "12:36:09.363",
            Dataframe7d:
              "7d201011ff92405fffff0100796400ff6fffff35887f78ff134015801a0029c02a",
            Dataframe80:
              "801c04cd6fff4fff227e1c000001000020378781008d054c0612100000",
          },
          {
            Time: "12:36:09.730",
            Dataframe7d:
              "7d201011ff924060ffff0100796400ff6fffff35887f72ff134015801a0029c02a",
            Dataframe80:
              "801c04d46fff4fff217f1c0000010000203787810087054c060a100000",
          },
          {
            Time: "12:36:10.130",
            Dataframe7d:
              "7d201011ff924061ffff0100796400ff6fffff35887f6bff134015801a0029c02a",
            Dataframe80:
              "801c04cf6fff4fff217f1c0000010000203787800092054c060d100000",
          },
          {
            Time: "12:36:10.688",
            Dataframe7d:
              "7d201011ff924063ffff0100796400ff6effff35887f61ff134015801a0029c02a",
            Dataframe80:
              "801c04ce70ff4fff21801c0000010000203787800094054c0607100000",
          },
          {
            Time: "12:36:11.502",
            Dataframe7d:
              "7d201011ff924064ffff0100796400ff6effff35887f5fff134015801a0029c02a",
            Dataframe80:
              "801c04bc70ff4fff22801c00000100002037878000a6054c05ff100000",
          },
          {
            Time: "12:36:12.045",
            Dataframe7d:
              "7d201011ff924066ffff0100796400ff6effff35887f5bff134015801a0029c02a",
            Dataframe80:
              "801c04c470ff4fff21811c000001000020378780009b054c0607100000",
          },
          {
            Time: "12:36:12.811",
            Dataframe7d:
              "7d201011ff924069ffff0100796400ff6effff35887f5eff134015801a0029c02a",
            Dataframe80:
              "801c04bf70ff4fff22811c00000100002037878000a6054c05ff100000",
          },
          {
            Time: "12:36:13.338",
            Dataframe7d:
              "7d201011ff92406affff0100796400ff6effff35887f5bff134015801a0029c02a",
            Dataframe80:
              "801c04b070ff4fff23811c00000100002037878000b0054c05f8100000",
          },
          {
            Time: "12:36:14.136",
            Dataframe7d:
              "7d201011ff92406dffff0100796400ff6effff35887f59ff134015801a0029c02a",
            Dataframe80:
              "801c04c470ff4fff22821c000001000020378780009a054c05f8100000",
          },
          {
            Time: "12:36:14.631",
            Dataframe7d:
              "7d201011ff92406affff0100796400ff6effff35887f57ff134015801a0029c02a",
            Dataframe80:
              "801c04b670ff4fff22821c00000100002037878000b1054c05f6100000",
          },
          {
            Time: "12:36:15.174",
            Dataframe7d:
              "7d201011ff92406affff0100796400ff6effff35887f73ff134015801a0029c02a",
            Dataframe80:
              "801c04bc70ff4fff22821c000001000020378780009d054c05f4100000",
          },
          {
            Time: "12:36:15.940",
            Dataframe7d:
              "7d201011ff924058ffff0100796400ff6effff35887f57ff134015801a0029c02a",
            Dataframe80:
              "801c04ba70ff4fff22821c00000100002037878000a6054c05f6100000",
          },
          {
            Time: "12:36:16.436",
            Dataframe7d:
              "7d201011ff924053ffff0100796400ff6effff35887f60ff134015801a0029c02a",
            Dataframe80:
              "801c04bf70ff50ff23821c000001000020378780009b054c05fa100000",
          },
          {
            Time: "12:36:16.962",
            Dataframe7d:
              "7d201011ff92405effff0100796400ff6effff35887f61ff134015801a0029c02a",
            Dataframe80:
              "801c04bb70ff50ff22821c00000100002037878000a5054c05f9100000",
          },
          {
            Time: "12:36:17.760",
            Dataframe7d:
              "7d201011ff924049ffff0100796400ff6effff35887f62ff134015801a0029c02a",
            Dataframe80:
              "801c04bb70ff50ff22821c000001000020378780009f054c05f0100000",
          },
          {
            Time: "12:36:18.240",
            Dataframe7d:
              "7d201011ff92403effff0100796400ff6effff35887f61ff134015801a0029c02a",
            Dataframe80:
              "801c04bc70ff50ff22831c00000100002037878000a5054c05f8100000",
          },
          {
            Time: "12:36:18.766",
            Dataframe7d:
              "7d201010ff92404affff0100796400ff6dffff35887f81ff134015801a0029c02a",
            Dataframe80:
              "801c04c871ff50ff22831b00000100002037877f0078054c05f2100000",
          },
          {
            Time: "12:36:19.596",
            Dataframe7d:
              "7d201010ff924036ffff0100796400ff6dffff35887f4bff134015801a0029c02a",
            Dataframe80:
              "801c049071ff50ff22831b00000100002037877f00b7054c0615100000",
          },
          {
            Time: "12:36:20.075",
            Dataframe7d:
              "7d201010ff92402fffff0100796400ff6dffff35887f2eff134015801a0029c02a",
            Dataframe80:
              "801c047f71ff50ff23831b00000100002037877f00d2054c061d100000",
          },
          {
            Time: "12:36:20.889",
            Dataframe7d:
              "7d201010ff924029ffff0100796400ff6dffff35887ef8ff134015801a0029c02a",
            Dataframe80:
              "801c049371ff50ff23821b00000100002037877f00b5054c061b100000",
          },
          {
            Time: "12:36:21.352",
            Dataframe7d:
              "7d201010ff924028ffff0100796400ff6dffff35887df3ff134015801a0029c02a",
            Dataframe80:
              "801c03bc71ff50ff29821b00000100002037877f01c5054c061b100000",
          },
          {
            Time: "12:36:21.895",
            Dataframe7d:
              "7d201010ff924022ffff0100796400ff6cffff35887e8cff134015801a0029c02a",
            Dataframe80:
              "801c035571ff50ff2f821b00000100002037877f01c4054c0619100000",
          },
          {
            Time: "12:36:22.693",
            Dataframe7d:
              "7d201010ff92408dffff0100796400ff6cffff35887efbff134015801a0029c02a",
            Dataframe80:
              "801c045372ff50ff26821b00000100002037877f00f6054c0618100000",
          },
          {
            Time: "12:36:23.220",
            Dataframe7d:
              "7d201010ff92405effff0100796400ff6cffff35887eccff134015801a0029c02a",
            Dataframe80:
              "801c042f72ff50ff25821b00000100002037877f0125054c0617100000",
          },
          {
            Time: "12:36:23.699",
            Dataframe7d:
              "7d201010ff92403cffff0100796400ff6cffff35887eadff134015801a0029c02a",
            Dataframe80:
              "801c040a72ff50ff27821b00000100002037877f013e054c0615100000",
          },
          {
            Time: "12:36:24.066",
            Dataframe7d:
              "7d201010ff92402fffff0100796400ff6cffff35887e8bff134015801a0029c02a",
            Dataframe80:
              "801c03da72ff50ff27821b00000100002037877f016b054c0617100000",
          },
          {
            Time: "12:36:24.561",
            Dataframe7d:
              "7d201010ff924028ffff0100796400ff6cffff35887e7eff134015801a0029c02a",
            Dataframe80:
              "801c03d172ff50ff2a821b00000100002037877f0174054c0615100000",
          },
          {
            Time: "12:36:25.024",
            Dataframe7d:
              "7d201010ff924024ffff0100796400ff6cffff35887e30ff134015801a0029c02a",
            Dataframe80:
              "801c03bd72ff50ff2a821b00000100002037877f0195054c0615100000",
          },
          {
            Time: "12:36:25.392",
            Dataframe7d:
              "7d201010ff924022ffff0100796400ff6cffff35887df3ff134015801a0029c02a",
            Dataframe80:
              "801c036472ff50ff2e821b00000100002037877f020e054c061b100000",
          },
          {
            Time: "12:36:25.886",
            Dataframe7d:
              "7d201010ff924020ffff0100796400ff6cffff35887e10ff134015801a0029c02a",
            Dataframe80:
              "801c035472ff50ff2f821b00000100002037877f01ec054c0619100000",
          },
          {
            Time: "12:36:26.365",
            Dataframe7d:
              "7d201010ff92401fffff0100796400ff6bffff35887f6bff134015801a0029c02a",
            Dataframe80:
              "801c03cf72ff50ff2c821b00000100002037877f0110054c0617100000",
          },
          {
            Time: "12:36:26.844",
            Dataframe7d:
              "7d201010ff92408dffff0100796400ff6bffff35887f0bff134015801a0029c02a",
            Dataframe80:
              "801c04be72ff50ff23821b00000100002037877e00b2054c0647100000",
          },
          {
            Time: "12:36:27.212",
            Dataframe7d:
              "7d201010ff92403cffff0100796400ff6bffff35887ec4ff134015801a0029c02a",
            Dataframe80:
              "801c044072ff50ff25821b00000100002037877e011b054c0641100000",
          },
          {
            Time: "12:36:27.690",
            Dataframe7d:
              "7d201010ff924029ffff0100796400ff6bffff35887ec0ff134015801a0029c02a",
            Dataframe80:
              "801c03f472ff50ff27821b00000100002037877e014e054c063b100000",
          },
          {
            Time: "12:36:28.185",
            Dataframe7d:
              "7d201010ff924023ffff0100796400ff6bffff35887f02ff134015801a0029c02a",
            Dataframe80:
              "801c041772ff50ff26821b00000100002037877e011d054c0640100000",
          },
          {
            Time: "12:36:29.000",
            Dataframe7d:
              "7d201010ff92401effff0100796400ff6bffff35887f60ff134015801a0029c02a",
            Dataframe80:
              "801c047e72ff50ff24821b00000100002037877e00b0054c05eb100000",
          },
          {
            Time: "12:36:29.510",
            Dataframe7d:
              "7d201017ff92403dffff0100796400ff6bffff35887e96ff134015801a0029c02a",
            Dataframe80:
              "801c049272ff51ff26832008000100002037877e0115053f05eb100000",
          },
          {
            Time: "12:36:29.990",
            Dataframe7d:
              "7d20101bff92401fffff0100796400ff6bffff3588801cff134015801a0029c02a",
            Dataframe80:
              "801c044472ff51ff42832708000100002037877e0085055105f3100000",
          },
          {
            Time: "12:36:30.404",
            Dataframe7d:
              "7d20101bff9240a1ffff0100796400ff6bffff35887f1eff134015801a0029c02a",
            Dataframe80:
              "801c051e72ff51ff3b832808000100002037877e006b055205f4100000",
          },
          {
            Time: "12:36:30.947",
            Dataframe7d:
              "7d20101bff92409cffff0100796400ff6bffff35887ef9ff134015801a0029c02a",
            Dataframe80:
              "801c03e472ff51ff47832808000100002037877e018b054e05f4100000",
          },
          {
            Time: "12:36:31.347",
            Dataframe7d:
              "7d20101eff9240a2ffff0100796400ff6affff35888004ff134015801a0029c02a",
            Dataframe80:
              "801c047272ff51ff45832808000100002037877f00b5055405f7100000",
          },
          {
            Time: "12:36:31.730",
            Dataframe7d:
              "7d201022ff9240a0ffff0100796400ff6bffff358880fcff134015801a0029c02a",
            Dataframe80:
              "801c054073ff51ff48832e0800010000203787860074055b05f5100000",
          },
          {
            Time: "12:36:32.272",
            Dataframe7d:
              "7d201026ff9240a0ffff0100796400ff6affff3588827cff134015801a0029c02a",
            Dataframe80:
              "801c06d273ff51ff45833308000100002037878901e0056105f6100000",
          },
          {
            Time: "12:36:33.118",
            Dataframe7d:
              "7d201027ff92409effff0100796400ff6affff35888496ff134015801a0029c02a",
            Dataframe80:
              "801c090773ff51ff3e833608000100002037878c040a056205f5100000",
          },
          {
            Time: "12:36:33.996",
            Dataframe7d:
              "7d20100fff9240a1ffff0100796400ff6affff358882a1ff134015801a0029c02a",
            Dataframe80:
              "801c09de73ff51ff16831b00000100002037877703f1054e05fa100000",
          },
          {
            Time: "12:36:34.555",
            Dataframe7d:
              "7d20101aff92409bffff0100796400ff6affff358880fdff134015801a0029c02a",
            Dataframe80:
              "801c06c673ff51ff218323080001000020378778017f055505f1100000",
          },
          {
            Time: "12:36:35.369",
            Dataframe7d:
              "7d201021ff92409effff0100796400ff69ffff35888151ff134015801a0029c02a",
            Dataframe80:
              "801c069c73ff51ff3e832f0800010000203787840169055b05f7100000",
          },
          {
            Time: "12:36:35.880",
            Dataframe7d:
              "7d201023ff9240a0ffff0101796200ff69ffff35888203ff134015801a0029c02a",
            Dataframe80:
              "801c06d574ff51ff40833008000100002037878501e5055c05f3100000",
          },
          {
            Time: "12:36:36.279",
            Dataframe7d:
              "7d201023ff92409effff0101796100ff69ffff35888247ff134015801a0029c02a",
            Dataframe80:
              "801c073074ff51ff408331080001000020378786022a055c05f7100000",
          },
          {
            Time: "12:36:36.822",
            Dataframe7d:
              "7d201024ff92409bffff0101796000ff69ffff358882abff134015801a0029c02a",
            Dataframe80:
              "801c078a74ff51ff3e83310800010000203787870287055c05fa100000",
          },
          {
            Time: "12:36:37.237",
            Dataframe7d:
              "7d201024ff924099ffff0101795f00ff69ffff358882f2ff134015801a0029c02a",
            Dataframe80:
              "801c07d974ff51ff3c833208000100002037878702c7055d05f8100000",
          },
          {
            Time: "12:36:37.668",
            Dataframe7d:
              "7d201024ff924094ffff0101795e00ff69ffff3588832aff134015801a0029c02a",
            Dataframe80:
              "801c081f74ff51ff3c83320800010000203787870308055e05f8100000",
          },
          {
            Time: "12:36:38.099",
            Dataframe7d:
              "7d201024ff924077ffff0101795d00ff69ffff35888368ff134015801a0029c02a",
            Dataframe80:
              "801c085f74ff51ff3a8332080001000020378787034f055c05f6100000",
          },
          {
            Time: "12:36:38.530",
            Dataframe7d:
              "7d201024ff924027ffff0101796200ff69ffff358883a7ff134015801a0029c02a",
            Dataframe80:
              "801c089f74ff51ff3983320800010000203787870379055c05f8100000",
          },
          {
            Time: "12:36:38.977",
            Dataframe7d:
              "7d201024ff92401fffff0101796300ff69ffff358883f8ff134015801a0029c02a",
            Dataframe80:
              "801c08e974ff51ff38833208000100002037878703cd055e05fc100000",
          },
          {
            Time: "12:36:39.425",
            Dataframe7d:
              "7d201011ff924073ffff0101796300ff69ffff35888324ff134015801a0029c02a",
            Dataframe80:
              "801c093274ff51ff36833208000100002037878503e9055805fb100000",
          },
          {
            Time: "12:36:39.856",
            Dataframe7d:
              "7d20100dff9240a2ffff0101796200ff68ffff358880feff134015801a0029c02a",
            Dataframe80:
              "801c07d174ff51ff14831800000100002037876c021e054a05f3100000",
          },
          {
            Time: "12:36:40.415",
            Dataframe7d:
              "7d20101aff924056ffff0101796600ff68ffff358880e4ff134015801a0029c02a",
            Dataframe80:
              "801c061d74ff51ff2383230800010000203787790105055405f2100000",
          },
          {
            Time: "12:36:40.814",
            Dataframe7d:
              "7d20101bff924099ffff0101796400ff68ffff358880dbff134015801a0029c02a",
            Dataframe80:
              "801c060b74ff51ff30832708000100002037877b00df055005f2100000",
          },
          {
            Time: "12:36:41.213",
            Dataframe7d:
              "7d20101bff9240a0ffff0101796300ff68ffff358880efff134015801a0029c02a",
            Dataframe80:
              "801c061474ff51ff30832708000100002037877b00de055105f8100000",
          },
          {
            Time: "12:36:41.723",
            Dataframe7d:
              "7d201018ff92409fffff0101796200ff68ffff358880f0ff134015801a0029c02a",
            Dataframe80:
              "801c061874ff51ff2d832408000100002037877800ec055105fb100000",
          },
          {
            Time: "12:36:42.123",
            Dataframe7d:
              "7d201018ff92409bffff0101796100ff68ffff358880f2ff134015801a0029c02a",
            Dataframe80:
              "801c061674ff51ff2c832408000100002037877800eb055105fa100000",
          },
          {
            Time: "12:36:42.665",
            Dataframe7d:
              "7d201018ff924099ffff0101796000ff68ffff358880ebff134015801a0029c02a",
            Dataframe80:
              "801c061a74ff51ff2a832408000100002037877800ee055205f4100000",
          },
          {
            Time: "12:36:43.065",
            Dataframe7d:
              "7d201018ff924095ffff0101796000ff68ffff358880f2ff134015801a0029c02a",
            Dataframe80:
              "801c062474ff51ff2b832408000100002037877800f2055205fa100000",
          },
          {
            Time: "12:36:43.464",
            Dataframe7d:
              "7d201019ff924094ffff0101795f00ff68ffff358880f4ff134015801a0029c02a",
            Dataframe80:
              "801c062074ff51ff2b832508000100002037877900ff055105f8100000",
          },
          {
            Time: "12:36:43.975",
            Dataframe7d:
              "7d20101aff924019ffff0101796400ff68ffff358880f3ff134015801a0029c02a",
            Dataframe80:
              "801c062874ff51ff2f832608000100002037877a0104055105f6100000",
          },
          {
            Time: "12:36:44.374",
            Dataframe7d:
              "7d20101aff924087ffff0101796200ff68ffff358880ebff134015801a0029c02a",
            Dataframe80:
              "801c062974ff51ff2e832508000100002037877900ff055005f8100000",
          },
          {
            Time: "12:36:44.917",
            Dataframe7d:
              "7d20101aff92401affff0101796600ff68ffff35888104ff134015801a0029c02a",
            Dataframe80:
              "801c063d74ff51ff2f832608000100002037877a0102055105fd100000",
          },
          {
            Time: "12:36:45.316",
            Dataframe7d:
              "7d20101aff92400fffff0101796700ff68ffff3588810eff134015801a0029c02a",
            Dataframe80:
              "801c062f74ff51ff2f832608000100002037877a0112055005f6100000",
          },
          {
            Time: "12:36:45.715",
            Dataframe7d:
              "7d20101bff92408dffff0101796500ff68ffff35888113ff134015801a0029c02a",
            Dataframe80:
              "801c063e74ff51ff30832608000100002037877a010a055005f6100000",
          },
          {
            Time: "12:36:46.226",
            Dataframe7d:
              "7d20101bff924073ffff0101796400ff68ffff35888114ff134015801a0029c02a",
            Dataframe80:
              "801c064774ff51ff30832708000100002037877b0123055105f2100000",
          },
          {
            Time: "12:36:46.625",
            Dataframe7d:
              "7d20101bff924016ffff0101796800ff68ffff3588812aff134015801a0029c02a",
            Dataframe80:
              "801c064874ff51ff30832708000100002037877b0113055005f6100000",
          },
          {
            Time: "12:36:47.040",
            Dataframe7d:
              "7d20101bff924072ffff0101796700ff68ffff3588811eff134015801a0029c02a",
            Dataframe80:
              "801c065074ff51ff30832708000100002037877b0122055005f6100000",
          },
          {
            Time: "12:36:47.567",
            Dataframe7d:
              "7d201019ff924084ffff0101796600ff68ffff35888136ff134015801a0029c02a",
            Dataframe80:
              "801c065c74ff51ff2e83260800010000203787790125055105f4100000",
          },
          {
            Time: "12:36:48.365",
            Dataframe7d:
              "7d20100dff92409bffff0101796500ff68ffff35887fc0ff134015801a0029c02a",
            Dataframe80:
              "801c05fc74ff51ff17831700000100002037876b0069054d05f4100000",
          },
          {
            Time: "12:36:48.860",
            Dataframe7d:
              "7d20100fff924098ffff0101796400ff68ffff35887eacff134015801a0029c02a",
            Dataframe80:
              "801c045674ff51ff1c8318000001000020378776011d054c05f0100000",
          },
          {
            Time: "12:36:49.371",
            Dataframe7d:
              "7d20100fff924011ffff0101796600ff68ffff35887ecaff134015801a0029c02a",
            Dataframe80:
              "801c03de74ff51ff27831a00000100002037877c014f054c0613100000",
          },
          {
            Time: "12:36:49.738",
            Dataframe7d:
              "7d201010ff92400affff0101796600ff68ffff35887f48ff134015801a0029c02a",
            Dataframe80:
              "801c041074ff51ff27831b00000100002037877c00f5054c0619100000",
          },
          {
            Time: "12:36:50.265",
            Dataframe7d:
              "7d201010ff924006ffff0101796700ff68ffff35887f92ff134015801a0029c02a",
            Dataframe80:
              "801c04a676ff51ff22831a00000100002037877c0075054c0613100000",
          },
          {
            Time: "12:36:51.079",
            Dataframe7d:
              "7d20100fff924005ffff0101796700ff66ffff35887facff134015801a0029c02a",
            Dataframe80:
              "801c04bd74ff51ff21831b00000100002037877b0050054c060f100000",
          },
          {
            Time: "12:36:51.574",
            Dataframe7d:
              "7d20100fff924004ffff0101796800ff66ffff35887fb3ff134015801a0029c02a",
            Dataframe80:
              "801c04cb76ff51ff22831b00000100002037877b0059054c05e9100000",
          },
          {
            Time: "12:36:52.085",
            Dataframe7d:
              "7d201010ff924006ffff0101796800ff67ffff35887fbeff134015801a0029c02a",
            Dataframe80:
              "801c04ca76ff51ff21831b00000100002037877b0043054c05ec100000",
          },
          {
            Time: "12:36:52.452",
            Dataframe7d:
              "7d20100fff924027ffff0101796900ff66ffff35887faaff134015801a0029c02a",
            Dataframe80:
              "801c04d776ff51ff21831b00000100002037877b0045054c05e8100000",
          },
          {
            Time: "12:36:52.963",
            Dataframe7d:
              "7d201010ff924086ffff0101796800ff67ffff35887f6bff134015801a0029c02a",
            Dataframe80:
              "801c04c576ff52ff21831a00000100002037877b0050054c05f2100000",
          },
          {
            Time: "12:36:53.793",
            Dataframe7d:
              "7d201015ff92400dffff0101796900ff66ffff35887f3eff134015801a0029c02a",
            Dataframe80:
              "801c044b76ff52ff30832008000100002037877b00d3054205f2100000",
          },
          {
            Time: "12:36:54.288",
            Dataframe7d:
              "7d20101bff92408dffff0101796700ff66ffff35887f57ff134015801a0029c02a",
            Dataframe80:
              "801c045076ff52ff3c832608000100002037877b00be054f05f0100000",
          },
          {
            Time: "12:36:55.134",
            Dataframe7d:
              "7d201021ff9240a2ffff0101796600ff66ffff35887f8dff134015801a0029c02a",
            Dataframe80:
              "801c049876ff52ff4e832f0800010000203787830078055405f7100000",
          },
          {
            Time: "12:36:55.645",
            Dataframe7d:
              "7d201027ff9240a2ffff0101796500ff66ffff35887fc4ff134015801a0029c02a",
            Dataframe80:
              "801c04c376ff52ff5483350800010000203787890048055005fb100000",
          },
          {
            Time: "12:36:56.475",
            Dataframe7d:
              "7d20102aff9240a4ffff0101796300ff66ffff35888025ff134015801a0029c02a",
            Dataframe80:
              "801c051176ff52ff5783390800010000203787890010055005fc100000",
          },
          {
            Time: "12:36:56.970",
            Dataframe7d:
              "7d20102aff9240a3ffff0101796200ff66ffff35888059ff134015801a0029c02a",
            Dataframe80:
              "801c055376ff51ff55833908000100002037878a0046055105fb100000",
          },
          {
            Time: "12:36:57.864",
            Dataframe7d:
              "7d201026ff9240a2ffff0101795f00ff66ffff358880b1ff134015801a0029c02a",
            Dataframe80:
              "801c05a576ff51ff528336080001000020378789009a055805f6100000",
          },
          {
            Time: "12:36:58.375",
            Dataframe7d:
              "7d201026ff924093ffff0101795e00ff66ffff358880e1ff134015801a0029c02a",
            Dataframe80:
              "801c05d276ff51ff50833508000100002037878900c8055905f8100000",
          },
          {
            Time: "12:36:58.774",
            Dataframe7d:
              "7d201024ff92409bffff0101795d00ff65ffff358880f0ff134015801a0029c02a",
            Dataframe80:
              "801c05f876ff51ff4e833308000100002037878700ea055b05f8100000",
          },
          {
            Time: "12:36:59.189",
            Dataframe7d:
              "7d201022ff92406bffff0101795c00ff65ffff35888116ff134015801a0029c02a",
            Dataframe80:
              "801c061177ff51ff468331080001000020378784010c055d05f8100000",
          },
          {
            Time: "12:36:59.748",
            Dataframe7d:
              "7d201020ff92400fffff0101796100ff65ffff35888132ff134015801a0029c02a",
            Dataframe80:
              "801c063d77ff51ff43832f0800010000203787830131056005f5100000",
          },
          {
            Time: "12:37:00.546",
            Dataframe7d:
              "7d20101dff924019ffff0101796200ff65ffff3588814cff134015801a0029c02a",
            Dataframe80:
              "801c066277ff51ff37832a08000100002037877e0143055705f4100000",
          },
          {
            Time: "12:37:01.073",
            Dataframe7d:
              "7d20101dff924009ffff0101796300ff65ffff3588815eff134015801a0029c02a",
            Dataframe80:
              "801c066777ff51ff37832a08000100002037877e014e055705f9100000",
          },
          {
            Time: "12:37:01.871",
            Dataframe7d:
              "7d20101bff924037ffff0101796300ff65ffff3588816fff134015801a0029c02a",
            Dataframe80:
              "801c067877ff51ff36832a08000100002037877d015e055405f3100000",
          },
          {
            Time: "12:37:02.382",
            Dataframe7d:
              "7d20101aff924093ffff0101796200ff65ffff35888163ff134015801a0029c02a",
            Dataframe80:
              "801c067f77ff51ff2f832608000100002037877a0171055105f2100000",
          },
          {
            Time: "12:37:03.260",
            Dataframe7d:
              "7d201019ff92408dffff0101796100ff65ffff35888174ff134015801a0029c02a",
            Dataframe80:
              "801c068177ff51ff2c8325080001000020378779016c055305f3100000",
          },
          {
            Time: "12:37:03.787",
            Dataframe7d:
              "7d201019ff92408fffff0101796000ff65ffff3588816dff134015801a0029c02a",
            Dataframe80:
              "801c067c77ff51ff2c83250800010000203787790169055305f2100000",
          },
          {
            Time: "12:37:04.585",
            Dataframe7d:
              "7d201019ff92406dffff0101796200ff65ffff3588816aff134015801a0029c02a",
            Dataframe80:
              "801c067b77ff51ff2a8325080001000020378779016a055405f0100000",
          },
          {
            Time: "12:37:05.112",
            Dataframe7d:
              "7d201019ff924017ffff0101796500ff64ffff35888188ff134015801a0029c02a",
            Dataframe80:
              "801c067977ff51ff2a8325080001000020378779018a055405f2100000",
          },
          {
            Time: "12:37:05.926",
            Dataframe7d:
              "7d201019ff924096ffff0101796300ff64ffff3588818fff134015801a0029c02a",
            Dataframe80:
              "801c067577ff51ff2b83250800010000203787790186055405f2100000",
          },
          {
            Time: "12:37:06.421",
            Dataframe7d:
              "7d201019ff924057ffff0101796500ff64ffff35888187ff134015801a0029c02a",
            Dataframe80:
              "801c069b77ff51ff2a83250800010000203787790196055405f1100000",
          },
          {
            Time: "12:37:07.267",
            Dataframe7d:
              "7d20100dff9240a3ffff0101796300ff64ffff3588807bff134015801a0029c02a",
            Dataframe80:
              "801c067377ff51ff1c83180000010000203787680136054d05f3100000",
          },
          {
            Time: "12:37:07.746",
            Dataframe7d:
              "7d201011ff924013ffff0101796400ff64ffff35888052ff134015801a0029c02a",
            Dataframe80:
              "801c051d77ff51ff1e831b0800010000203787680055054d05f0100000",
          },
          {
            Time: "12:37:08.241",
            Dataframe7d:
              "7d201014ff924004ffff0101796500ff64ffff35888050ff134015801a0029c02a",
            Dataframe80:
              "801c055277ff51ff25831f08000100002037876b006b054f05fc100000",
          },
          {
            Time: "12:37:09.087",
            Dataframe7d:
              "7d201014ff9240a0ffff0101796400ff64ffff3588803dff134015801a0029c02a",
            Dataframe80:
              "801c053e77ff51ff27831f08000100002037876b0042054c05f6100000",
          },
          {
            Time: "12:37:09.582",
            Dataframe7d:
              "7d201014ff9240a2ffff0101796300ff64ffff35888043ff134015801a0029c02a",
            Dataframe80:
              "801c054577ff51ff26831f08000100002037876b0041054c05f4100000",
          },
          {
            Time: "12:37:10.412",
            Dataframe7d:
              "7d201014ff924094ffff0101796200ff64ffff35888038ff134015801a0029c02a",
            Dataframe80:
              "801c053777ff51ff27831f08000100002037876b003a054c05ef100000",
          },
          {
            Time: "12:37:10.907",
            Dataframe7d:
              "7d201014ff924017ffff0101796400ff63ffff35888038ff134015801a0029c02a",
            Dataframe80:
              "801c053577ff51ff26831f08000100002037876b0030054c05f1100000",
          },
          {
            Time: "12:37:11.402",
            Dataframe7d:
              "7d201013ff92400fffff0101796400ff63ffff3588801fff134015801a0029c02a",
            Dataframe80:
              "801c053d78ff52ff26831f08000100002037876b0039054c05f2100000",
          },
          {
            Time: "12:37:12.216",
            Dataframe7d:
              "7d201013ff92400effff0101796500ff63ffff35888019ff134015801a0029c02a",
            Dataframe80:
              "801c051978ff52ff26831e08000100002037876a0025054c05ed100000",
          },
          {
            Time: "12:37:12.711",
            Dataframe7d:
              "7d201013ff92401fffff0101796500ff63ffff35888016ff134015801a0029c02a",
            Dataframe80:
              "801c050c78ff52ff27831e08000100002037876a0010054b05ee100000",
          },
          {
            Time: "12:37:13.206",
            Dataframe7d:
              "7d201013ff92400cffff0101796600ff63ffff35888000ff134015801a0029c02a",
            Dataframe80:
              "801c04ff78ff52ff27831e08000100002037876a0010054b05ee100000",
          },
          {
            Time: "12:37:13.605",
            Dataframe7d:
              "7d201014ff924012ffff0101796600ff63ffff35888008ff134015801a0029c02a",
            Dataframe80:
              "801c050c78ff52ff27831e08000100002037876a0002054b05f2100000",
          },
          {
            Time: "12:37:14.116",
            Dataframe7d:
              "7d201014ff924006ffff0101796700ff63ffff35887ffeff134015801a0029c02a",
            Dataframe80:
              "801c04ee78ff52ff29831f08000100002037876b000c054a05f4100000",
          },
          {
            Time: "12:37:14.962",
            Dataframe7d:
              "7d201015ff92400dffff0101796800ff63ffff35887fe2ff134015801a0029c02a",
            Dataframe80:
              "801c04ea78ff52ff2b832008000100002037876c0008054b05f7100000",
          },
          {
            Time: "12:37:15.457",
            Dataframe7d:
              "7d20100eff92406affff0101796700ff63ffff35887fc8ff134015801a0029c02a",
            Dataframe80:
              "801c04eb78ff52ff2d832008000100002037876e001e054405f4100000",
          },
          {
            Time: "12:37:16.303",
            Dataframe7d:
              "7d201013ff924008ffff0101796800ff63ffff35888006ff134015801a0029c02a",
            Dataframe80:
              "801c03da78ff52ff25831b080001000020378779012f053c05eb100000",
          },
          {
            Time: "12:37:16.798",
            Dataframe7d:
              "7d201018ff92401bffff0101796700ff63ffff358880daff134015801a0029c02a",
            Dataframe80:
              "801c060378ff52ff2b832408000100002037877000f7055005f3100000",
          },
          {
            Time: "12:37:17.644",
            Dataframe7d:
              "7d201020ff924028ffff0101796900ff63ffff358880e4ff134015801a0029c02a",
            Dataframe80:
              "801c05e178ff52ff3e832e08000100002037877a00e7056005f0100000",
          },
          {
            Time: "12:37:18.155",
            Dataframe7d:
              "7d20100cff9240aeffff0101796700ff63ffff35888064ff134015801a0029c02a",
            Dataframe80:
              "801c05ce78ff52ff1c831700000100002037876700ab054705f6100000",
          },
          {
            Time: "12:37:18.985",
            Dataframe7d:
              "7d20100fff924075ffff0101796800ff63ffff35887e85ff134015801a0029c02a",
            Dataframe80:
              "801c03c778ff52ff2183180000010000203787780156054c05ed100000",
          },
          {
            Time: "12:37:19.480",
            Dataframe7d:
              "7d20100fff924005ffff0101796800ff62ffff35887f33ff134015801a0029c02a",
            Dataframe80:
              "801c03a779ff52ff29831a0000010000203787790110054c0611100000",
          },
          {
            Time: "12:37:20.342",
            Dataframe7d:
              "7d20100fff924006ffff0101796900ff62ffff35887fbfff134015801a0029c02a",
            Dataframe80:
              "801c04be79ff53ff20831a000001000020378779003a054b05e6100000",
          },
          {
            Time: "12:37:20.837",
            Dataframe7d:
              "7d20100fff924008ffff0101796a00ff62ffff35887fb7ff134015801a0029c02a",
            Dataframe80:
              "801c04b679ff53ff20831a0000010000203787790043054c05f0100000",
          },
          {
            Time: "12:37:21.667",
            Dataframe7d:
              "7d20101cff92408cffff0101796900ff62ffff35887ffaff134015801a0029c02a",
            Dataframe80:
              "801c04f479ff53ff2a83240800010000203787730045054e05f0100000",
          },
          {
            Time: "12:37:22.162",
            Dataframe7d:
              "7d20101dff9240a7ffff0101796900ff62ffff35888010ff134015801a0029c02a",
            Dataframe80:
              "801c052e79ff53ff3e8329080001000020378777002d055405f2100000",
          },
          {
            Time: "12:37:22.673",
            Dataframe7d:
              "7d20101eff9240adffff0101796800ff62ffff35888025ff134015801a0029c02a",
            Dataframe80:
              "801c052a79ff53ff43832b080001000020378778002c055905f4100000",
          },
          {
            Time: "12:37:23.487",
            Dataframe7d:
              "7d20101eff9240acffff0101796600ff62ffff35888037ff134015801a0029c02a",
            Dataframe80:
              "801c052079ff53ff44832b080001000020378778002e055905f3100000",
          },
          {
            Time: "12:37:23.982",
            Dataframe7d:
              "7d20101eff9240acffff0101796500ff62ffff35888031ff134015801a0029c02a",
            Dataframe80:
              "801c052a79ff53ff43832b0800010000203787780030055905f6100000",
          },
          {
            Time: "12:37:24.476",
            Dataframe7d:
              "7d20101eff9240a9ffff0101796400ff60ffff35888035ff134015801a0029c02a",
            Dataframe80:
              "801c052979ff53ff43832b0800010000203787780033055a05f4100000",
          },
          {
            Time: "12:37:24.860",
            Dataframe7d:
              "7d20101fff9240a6ffff0101796400ff62ffff35888034ff134015801a0029c02a",
            Dataframe80:
              "801c052879ff53ff45832c0800010000203787790037055905f2100000",
          },
          {
            Time: "12:37:25.355",
            Dataframe7d:
              "7d20101fff9240a3ffff0101796300ff62ffff35888039ff134015801a0029c02a",
            Dataframe80:
              "801c052d79ff53ff45832c080001000020378779003c055905f2100000",
          },
          {
            Time: "12:37:25.738",
            Dataframe7d:
              "7d20101fff92409bffff0101796200ff62ffff3588803eff134015801a0029c02a",
            Dataframe80:
              "801c053579ff53ff45832c080001000020378779003f055905f5100000",
          },
          {
            Time: "12:37:26.249",
            Dataframe7d:
              "7d20101fff924094ffff0101796000ff62ffff3588803bff134015801a0029c02a",
            Dataframe80:
              "801c053179ff53ff45832c080001000020378779003e055905f6100000",
          },
          {
            Time: "12:37:27.079",
            Dataframe7d:
              "7d20101bff92407effff0101796200ff60ffff3588803fff134015801a0029c02a",
            Dataframe80:
              "801c053279ff52ff3d83290800010000203787760033055505f3100000",
          },
          {
            Time: "12:37:27.574",
            Dataframe7d:
              "7d20101bff92400cffff0101796300ff5fffff3588804fff134015801a0029c02a",
            Dataframe80:
              "801c05397aff52ff3e83280800010000203787750054055605f2100000",
          },
          {
            Time: "12:37:28.084",
            Dataframe7d:
              "7d201017ff92409cffff0101796100ff5fffff35888041ff134015801a0029c02a",
            Dataframe80:
              "801c05277aff52ff3583240800010000203787700044054d05f2100000",
          },
          {
            Time: "12:37:28.452",
            Dataframe7d:
              "7d201016ff92409bffff0101796100ff60ffff35888028ff134015801a0029c02a",
            Dataframe80:
              "801c051e7aff52ff2e832208000100002037876f003e054a05f0100000",
          },
          {
            Time: "12:37:28.962",
            Dataframe7d:
              "7d201014ff92400affff0101796200ff5fffff35888016ff134015801a0029c02a",
            Dataframe80:
              "801c05147aff52ff2a831f08000100002037876c001e054b05ef100000",
          },
          {
            Time: "12:37:29.346",
            Dataframe7d:
              "7d201015ff924008ffff0101796300ff5fffff3588801aff134015801a0029c02a",
            Dataframe80:
              "801c04f97aff52ff2b832008000100002037876d0016054905f0100000",
          },
          {
            Time: "12:37:29.856",
            Dataframe7d:
              "7d201016ff924008ffff0101796400ff5fffff3588801eff134015801a0029c02a",
            Dataframe80:
              "801c04fe7aff52ff2e832208000100002037876f0022054905f0100000",
          },
          {
            Time: "12:37:30.703",
            Dataframe7d:
              "7d201018ff92400affff0101796500ff5fffff35888058ff134015801a0029c02a",
            Dataframe80:
              "801c05177aff52ff348324080001000020378771003a054d05f2100000",
          },
          {
            Time: "12:37:31.214",
            Dataframe7d:
              "7d201018ff92401dffff0101796600ff5fffff35888074ff134015801a0029c02a",
            Dataframe80:
              "801c05417aff52ff338324080001000020378771006a054f05f5100000",
          },
          {
            Time: "12:37:32.028",
            Dataframe7d:
              "7d201016ff92409cffff0101796400ff5fffff358880abff134015801a0029c02a",
            Dataframe80:
              "801c056d7aff52ff2e832208000100002037876f0095054d05ef100000",
          },
          {
            Time: "12:37:32.523",
            Dataframe7d:
              "7d201015ff92409fffff0101796300ff5fffff358880bfff134015801a0029c02a",
            Dataframe80:
              "801c05977aff52ff2a832108000100002037876e00b6055005f1100000",
          },
          {
            Time: "12:37:33.033",
            Dataframe7d:
              "7d201015ff924090ffff0101796200ff5fffff358880d1ff134015801a0029c02a",
            Dataframe80:
              "801c05aa7aff52ff29832108000100002037876e00ce055105f0100000",
          },
          {
            Time: "12:37:33.417",
            Dataframe7d:
              "7d201015ff924021ffff0101796400ff5fffff358880e7ff134015801a0029c02a",
            Dataframe80:
              "801c05c17aff52ff27832108000100002037876e00d7055105f2100000",
          },
          {
            Time: "12:37:33.927",
            Dataframe7d:
              "7d201015ff924094ffff0101796400ff5fffff3588810aff134015801a0029c02a",
            Dataframe80:
              "801c05d97aff52ff27842108000100002037876e00fb055405f1100000",
          },
          {
            Time: "12:37:34.327",
            Dataframe7d:
              "7d201015ff924095ffff0101796300ff5fffff35888126ff134015801a0029c02a",
            Dataframe80:
              "801c05f07aff52ff26832108000100002037876e011c055405f2100000",
          },
          {
            Time: "12:37:34.853",
            Dataframe7d:
              "7d201015ff92408dffff0101796400ff5effff35888148ff134015801a0029c02a",
            Dataframe80:
              "801c06197aff52ff25832108000100002037876e013a055405eb100000",
          },
          {
            Time: "12:37:35.253",
            Dataframe7d:
              "7d201015ff924012ffff0101796700ff5effff3588815dff134015801a0029c02a",
            Dataframe80:
              "801c062c7bff52ff25832008000100002037876d0156055405ec100000",
          },
          {
            Time: "12:37:35.652",
            Dataframe7d:
              "7d201014ff92409dffff0101796500ff5effff35888168ff134015801a0029c02a",
            Dataframe80:
              "801c06487bff52ff22831f08000100002037876c016d055405ea100000",
          },
          {
            Time: "12:37:36.163",
            Dataframe7d:
              "7d201012ff9240a8ffff0101796400ff5effff3588816cff134015801a0029c02a",
            Dataframe80:
              "801c06527bff52ff1e831d08000100002037876a0175055405ef100000",
          },
          {
            Time: "12:37:36.993",
            Dataframe7d:
              "7d20100fff9240a1ffff0101796300ff5effff358881a5ff134015801a0029c02a",
            Dataframe80:
              "801c06637bff52ff1b841b08000100002037876801a2054f05ec100000",
          },
          {
            Time: "12:37:37.503",
            Dataframe7d:
              "7d20100fff9240a6ffff0101796200ff5effff35888194ff134015801a0029c02a",
            Dataframe80:
              "801c06787bff53ff19841a00000100002037875e01a0054605f2100000",
          },
          {
            Time: "12:37:38.334",
            Dataframe7d:
              "7d20100fff92400bffff0101796200ff5effff35888186ff134015801a0029c02a",
            Dataframe80:
              "801c06787bff53ff19841a00000100002037875e019d053805ef100000",
          },
          {
            Time: "12:37:38.829",
            Dataframe7d:
              "7d20100fff924008ffff0101796300ff5effff3588819cff134015801a0029c02a",
            Dataframe80:
              "801c067a7bff53ff18841a00000100002037875e0193053805e8100000",
          },
          {
            Time: "12:37:39.675",
            Dataframe7d:
              "7d201011ff924006ffff0101796500ff5effff35888179ff134015801a0029c02a",
            Dataframe80:
              "801c06607bff53ff1b831c0800010000203787600186055005ed100000",
          },
          {
            Time: "12:37:40.186",
            Dataframe7d:
              "7d201011ff92400bffff0101796600ff5effff3588817dff134015801a0029c02a",
            Dataframe80:
              "801c06667bff53ff1b831c080001000020378760017c055005f2100000",
          },
          {
            Time: "12:37:41.016",
            Dataframe7d:
              "7d201011ff924008ffff0101796700ff5dffff3588818aff134015801a0029c02a",
            Dataframe80:
              "801c06537bff53ff1c831c0800010000203787600181055005ea100000",
          },
          {
            Time: "12:37:41.495",
            Dataframe7d:
              "7d20100fff924008ffff0101796800ff5dffff35888164ff134015801a0029c02a",
            Dataframe80:
              "801c06427bff53ff19841a00000100002037875d0176054b05ea100000",
          },
          {
            Time: "12:37:41.990",
            Dataframe7d:
              "7d20100eff924067ffff0101796800ff5dffff3588815cff134015801a0029c02a",
            Dataframe80:
              "801c063c7bff53ff17841800000100002037875d0157054305ec100000",
          },
          {
            Time: "12:37:42.836",
            Dataframe7d:
              "7d20100bff92400cffff0101796a00ff5dffff35887eaaff134015801a0029c02a",
            Dataframe80:
              "801c044a7bff53ff18841500000100002037876200cd054c05f0100000",
          },
          {
            Time: "12:37:43.331",
            Dataframe7d:
              "7d20100fff924006ffff0101796a00ff5dffff35887e0fff134015801a0029c02a",
            Dataframe80:
              "801c03467bff53ff2684190000010000203787760198053a05f4100000",
          },
          {
            Time: "12:37:44.177",
            Dataframe7d:
              "7d20100fff924004ffff0101796c00ff5dffff35887f5eff134015801a0029c02a",
            Dataframe80:
              "801c039b7bff53ff2a831908000100002037877600e2053c060d100000",
          },
          {
            Time: "12:37:44.672",
            Dataframe7d:
              "7d20100fff924099ffff0101796a00ff5dffff35887f4eff134015801a0029c02a",
            Dataframe80:
              "801c043d7bff53ff238319080001000020378776009d053f060d100000",
          },
          {
            Time: "12:37:45.182",
            Dataframe7d:
              "7d20100fff92400affff0101796c00ff5dffff35887efdff134015801a0029c02a",
            Dataframe80:
              "801c03e67bff53ff25831908000100002037877600f1053c0615100000",
          },
          {
            Time: "12:37:45.550",
            Dataframe7d:
              "7d20100fff924006ffff0101796c00ff5dffff35887ee6ff134015801a0029c02a",
            Dataframe80:
              "801c03b37bff53ff258319080001000020378776011c053b0617100000",
          },
          {
            Time: "12:37:46.108",
            Dataframe7d:
              "7d20100fff924006ffff0101796c00ff5dffff35887efeff134015801a0029c02a",
            Dataframe80:
              "801c03a87bff53ff278319080001000020378776011b053b0610100000",
          },
          {
            Time: "12:37:46.859",
            Dataframe7d:
              "7d20100fff924092ffff0101796c00ff5dffff35887f5aff134015801a0029c02a",
            Dataframe80:
              "801c04157bff53ff23831908000100002037877600b2053e05eb100000",
          },
          {
            Time: "12:37:47.354",
            Dataframe7d:
              "7d20100fff9240a2ffff0101796b00ff5dffff35887f6fff134015801a0029c02a",
            Dataframe80:
              "801c042e7bff53ff2283190800010000203787760096054005fa100000",
          },
          {
            Time: "12:37:47.864",
            Dataframe7d:
              "7d20100fff924099ffff0101796b00ff5dffff35887f5fff134015801a0029c02a",
            Dataframe80:
              "801c04347bff53ff238419080001000020378776009c053f05fa100000",
          },
          {
            Time: "12:37:48.647",
            Dataframe7d:
              "7d20100fff92409bffff0101796a00ff5dffff35887f55ff134015801a0029c02a",
            Dataframe80:
              "801c04237bff53ff23841908000100002037877600ae053f05e8100000",
          },
          {
            Time: "12:37:49.142",
            Dataframe7d:
              "7d20100fff92408dffff0101796a00ff5dffff35887f53ff134015801a0029c02a",
            Dataframe80:
              "801c04147bff53ff23841908000100002037877600b0053f05e5100000",
          },
          {
            Time: "12:37:49.652",
            Dataframe7d:
              "7d20100fff92409cffff0101796900ff5dffff35887f5cff134015801a0029c02a",
            Dataframe80:
              "801c04217bff53ff238419080001000020378776009f053f05ed100000",
          },
          {
            Time: "12:37:50.451",
            Dataframe7d:
              "7d20100fff92408bffff0101796800ff5dffff35887f53ff134015801a0029c02a",
            Dataframe80:
              "801c04227bff53ff23841908000100002037877600a3053f05e0100000",
          },
          {
            Time: "12:37:50.930",
            Dataframe7d:
              "7d20100fff92408bffff0101796800ff5dffff35887f5fff134015801a0029c02a",
            Dataframe80:
              "801c04227bff53ff23841908000100002037877600a3053f05ea100000",
          },
          {
            Time: "12:37:51.424",
            Dataframe7d:
              "7d20100fff92408fffff0101796800ff5dffff35887f56ff134015801a0029c02a",
            Dataframe80:
              "801c04267bff53ff2384190800010000203787760098053f05ed100000",
          },
          {
            Time: "12:37:52.239",
            Dataframe7d:
              "7d20100fff924085ffff0101796800ff5bffff35887fafff134015801a0029c02a",
            Dataframe80:
              "801c045d7cff53ff2084190000010000203787740068054c05ec100000",
          },
          {
            Time: "12:37:52.797",
            Dataframe7d:
              "7d20100fff9240a6ffff0101796800ff5bffff35887fd0ff134015801a0029c02a",
            Dataframe80:
              "801c048e7cff53ff1f84190000010000203787740032054a05ed100000",
          },
          {
            Time: "12:37:53.612",
            Dataframe7d:
              "7d20100fff924097ffff0101796700ff5bffff35887fb0ff134015801a0029c02a",
            Dataframe80:
              "801c04827cff53ff208419000001000020378774004f054c0618100000",
          },
          {
            Time: "12:37:54.123",
            Dataframe7d:
              "7d20100fff924010ffff0101796800ff5bffff35887fafff134015801a0029c02a",
            Dataframe80:
              "801c047b7cff53ff2084190000010000203787740053054c060e100000",
          },
          {
            Time: "12:37:54.937",
            Dataframe7d:
              "7d20100fff9240a0ffff0101796700ff5bffff35887fa1ff134015801a0029c02a",
            Dataframe80:
              "801c04817cff53ff2084190000010000203787740053054c060e100000",
          },
          {
            Time: "12:37:55.447",
            Dataframe7d:
              "7d20100fff9240a0ffff0101796600ff5bffff35887f9fff134015801a0029c02a",
            Dataframe80:
              "801c04717cff53ff2083190000010000203787740058054c0614100000",
          },
          {
            Time: "12:37:55.926",
            Dataframe7d:
              "7d20100fff924062ffff0101796600ff5bffff35887f9dff134015801a0029c02a",
            Dataframe80:
              "801c046b7cff53ff208319000001000020378774005c054c060e100000",
          },
          {
            Time: "12:37:56.741",
            Dataframe7d:
              "7d20100fff924018ffff0101796800ff5bffff35887f8cff134015801a0029c02a",
            Dataframe80:
              "801c045a7cff54ff2083190000010000203787740076054c0610100000",
          },
          {
            Time: "12:37:57.299",
            Dataframe7d:
              "7d20100fff92400fffff0101796800ff5bffff35887f88ff134015801a0029c02a",
            Dataframe80:
              "801c04547cff54ff2183190000010000203787740078054c0614100000",
          },
          {
            Time: "12:37:58.098",
            Dataframe7d:
              "7d20100fff924017ffff0101796900ff5bffff35887f94ff134015801a0029c02a",
            Dataframe80:
              "801c04567cff54ff208319000001000020378774006c054c0612100000",
          },
          {
            Time: "12:37:58.577",
            Dataframe7d:
              "7d20100fff924010ffff0101796900ff5bffff35887f80ff134015801a0029c02a",
            Dataframe80:
              "801c045a7cff54ff208319000001000020378774006e054c0613100000",
          },
          {
            Time: "12:37:59.071",
            Dataframe7d:
              "7d20100fff924009ffff0101796a00ff5bffff35887f7fff134015801a0029c02a",
            Dataframe80:
              "801c043f7cff54ff2083190000010000203787740087054c0610100000",
          },
          {
            Time: "12:37:59.886",
            Dataframe7d:
              "7d20100fff924090ffff0101796900ff5bffff35887fa4ff134015801a0029c02a",
            Dataframe80:
              "801c045c7cff54ff2183190000010000203787740064054c0613100000",
          },
          {
            Time: "12:38:00.380",
            Dataframe7d:
              "7d20100fff924096ffff0101796800ff5bffff35887f9eff134015801a0029c02a",
            Dataframe80:
              "801c04677cff54ff208319000001000020378774005f054c0608100000",
          },
          {
            Time: "12:38:00.875",
            Dataframe7d:
              "7d20100fff924096ffff0101796800ff5bffff35887f87ff134015801a0029c02a",
            Dataframe80:
              "801c046c7cff54ff218319000001000020378774006a054c060d100000",
          },
          {
            Time: "12:38:01.386",
            Dataframe7d:
              "7d20100fff924009ffff0101796900ff5bffff35887f6dff134015801a0029c02a",
            Dataframe80:
              "801c04327cff54ff228319000001000020378774008f054c05ea100000",
          },
          {
            Time: "12:38:01.753",
            Dataframe7d:
              "7d20100fff924008ffff0101796a00ff5bffff35887fa0ff134015801a0029c02a",
            Dataframe80:
              "801c04477cff54ff218319000001000020378774006a054c05eb100000",
          },
          {
            Time: "12:38:02.249",
            Dataframe7d:
              "7d20100fff92400cffff0101796a00ff5bffff35887fb4ff134015801a0029c02a",
            Dataframe80:
              "801c046b7cff54ff2083190000010000203787740052054c05ec100000",
          },
          {
            Time: "12:38:03.063",
            Dataframe7d:
              "7d201012ff924097ffff0101796800ff5bffff35887fa6ff134015801a0029c02a",
            Dataframe80:
              "801c04707cff54ff208419000001000020378774005b054205e3100000",
          },
          {
            Time: "12:38:03.573",
            Dataframe7d:
              "7d20101bff9240a1ffff0101796714ff5affff35887f5cff134015801a0029c02a",
            Dataframe80:
              "801c03fd7cff54ff33842308000100002037877400bd054c05e5100000",
          },
          {
            Time: "12:38:04.436",
            Dataframe7d:
              "7d20101dff9240b1ffff010179671cff5bffff35887ff5ff134015801a0029c02a",
            Dataframe80:
              "801c04827dff54ff4484280800010000203787730036055505e7100000",
          },
          {
            Time: "12:38:04.947",
            Dataframe7d:
              "7d20101eff9240b3ffff0101796620ff5affff3588805bff134015801a0029c02a",
            Dataframe80:
              "801c04df7dff54ff45842a080001000020378772002e055805ec100000",
          },
          {
            Time: "12:38:05.792",
            Dataframe7d:
              "7d201026ff9240afffff010179652cff5affff35888116ff134015801a0029c02a",
            Dataframe80:
              "801c05997dff54ff50843408000100002037877c00e2055805f0100000",
          },
          {
            Time: "12:38:06.303",
            Dataframe7d:
              "7d201029ff9240b2ffff0101796434ff5affff35888189ff134015801a0029c02a",
            Dataframe80:
              "801c060b7dff54ff52843708000100002037877f0160055805f2100000",
          },
          {
            Time: "12:38:07.165",
            Dataframe7d:
              "7d20102aff9240afffff0101796238ff5affff35888233ff134015801a0029c02a",
            Dataframe80:
              "801c06c07dff54ff5084390800010000203787810210055c05f2100000",
          },
          {
            Time: "12:38:08.043",
            Dataframe7d:
              "7d20100cff9240b8ffff0101796300ff5affff358880e2ff134015801a0029c02a",
            Dataframe80:
              "801c06fd7dff53ff1a841700000100002037876801b5054805e5100000",
          },
          {
            Time: "12:38:08.586",
            Dataframe7d:
              "7d201013ff9240b6ffff0101796204ff5affff358880e7ff134015801a0029c02a",
            Dataframe80:
              "801c05367dff53ff20841e080001000020378766011a055405e2100000",
          },
          {
            Time: "12:38:09.416",
            Dataframe7d:
              "7d201019ff9240adffff0101796014ff5affff35888128ff134015801a0029c02a",
            Dataframe80:
              "801c05f37dff53ff2d842408000100002037876c0102055005e6100000",
          },
          {
            Time: "12:38:10.326",
            Dataframe7d:
              "7d201018ff9240aeffff0101795f10ff5affff35888137ff134015801a0029c02a",
            Dataframe80:
              "801c05f47dff53ff30832508000100002037876d0139055105e8100000",
          },
          {
            Time: "12:38:11.220",
            Dataframe7d:
              "7d201017ff9240a8ffff0101795d0cff5affff3588813eff134015801a0029c02a",
            Dataframe80:
              "801c060a7dff53ff2b832308000100002037876b013d055205e2100000",
          },
          {
            Time: "12:38:11.779",
            Dataframe7d:
              "7d201013ff9240aaffff0101795c08ff5affff35888139ff134015801a0029c02a",
            Dataframe80:
              "801c06067dff53ff25841e080001000020378766013b055405e8100000",
          },
          {
            Time: "12:38:12.178",
            Dataframe7d:
              "7d201013ff924016ffff0101795d08ff5affff35888133ff134015801a0029c02a",
            Dataframe80:
              "801c05fa7dff53ff20841e080001000020378766012a055405e2100000",
          },
          {
            Time: "12:38:12.577",
            Dataframe7d:
              "7d201013ff92400cffff0101795e08ff5affff35888130ff134015801a0029c02a",
            Dataframe80:
              "801c05f57dff53ff21841e080001000020378766012a055405e5100000",
          },
          {
            Time: "12:38:12.976",
            Dataframe7d:
              "7d20100fff92400cffff0101795e00ff5affff358880e6ff134015801a0029c02a",
            Dataframe80:
              "801c05f97dff53ff20841c080001000020378763012d055205e2100000",
          },
          {
            Time: "12:38:13.503",
            Dataframe7d:
              "7d20100dff924098ffff0101795e00ff58ffff35888131ff134015801a0029c02a",
            Dataframe80:
              "801c05dd7eff53ff1884170000010000203787600118054805e2100000",
          },
          {
            Time: "12:38:14.349",
            Dataframe7d:
              "7d20100bff924008ffff0101795f00ff58ffff358880f0ff134015801a0029c02a",
            Dataframe80:
              "801c05b57eff53ff16841500000100002037876000fd053805e2100000",
          },
          {
            Time: "12:38:14.876",
            Dataframe7d:
              "7d20100bff924006ffff0101796000ff58ffff358880e5ff134015801a0029c02a",
            Dataframe80:
              "801c05a07eff53ff16841500000100002037876000e4053805ec100000",
          },
          {
            Time: "12:38:15.675",
            Dataframe7d:
              "7d20100bff924006ffff0101796100ff58ffff358880b5ff134015801a0029c02a",
            Dataframe80:
              "801c05797eff53ff17841500000100002037876000b3053805e8100000",
          },
          {
            Time: "12:38:16.169",
            Dataframe7d:
              "7d20100bff924006ffff0101796100ff58ffff358880a4ff134015801a0029c02a",
            Dataframe80:
              "801c05637eff53ff16841500000100002037876000a1053805e9100000",
          },
          {
            Time: "12:38:16.648",
            Dataframe7d:
              "7d20100bff924006ffff0101796200ff58ffff35888089ff134015801a0029c02a",
            Dataframe80:
              "801c054b7eff53ff168415000001000020378760009d053805e2100000",
          },
          {
            Time: "12:38:17.047",
            Dataframe7d:
              "7d20100bff924006ffff0101796200ff58ffff35888070ff134015801a0029c02a",
            Dataframe80:
              "801c05357eff53ff178415000001000020378760008405380616100000",
          },
          {
            Time: "12:38:17.574",
            Dataframe7d:
              "7d20100bff924005ffff0101796300ff58ffff3588803dff134015801a0029c02a",
            Dataframe80:
              "801c05117eff53ff178415000001000020378760005c0538060b100000",
          },
          {
            Time: "12:38:18.388",
            Dataframe7d:
              "7d20100bff924008ffff0101796500ff58ffff35887fdeff134015801a0029c02a",
            Dataframe80:
              "801c04b57eff53ff178315000001000020378760000b053f060a100000",
          },
          {
            Time: "12:38:18.867",
            Dataframe7d:
              "7d20100cff924006ffff0101796500ff58ffff35887f97ff134015801a0029c02a",
            Dataframe80:
              "801c046a7eff53ff1983150000010000203787640053054c0607100000",
          },
          {
            Time: "12:38:19.362",
            Dataframe7d:
              "7d20100dff924005ffff0101796600ff58ffff35887f81ff134015801a0029c02a",
            Dataframe80:
              "801c042e7eff54ff1c831700000100002037876e0074054c060d100000",
          },
          {
            Time: "12:38:19.761",
            Dataframe7d:
              "7d20100eff924005ffff0101796600ff58ffff35887f5dff134015801a0029c02a",
            Dataframe80:
              "801c040f7eff54ff21831808000100002037877100a4053e0603100000",
          },
          {
            Time: "12:38:20.304",
            Dataframe7d:
              "7d20100eff924004ffff0101796700ff58ffff35887f5dff134015801a0029c02a",
            Dataframe80:
              "801c04237eff54ff20831808000100002037877100a4053e05e4100000",
          },
          {
            Time: "12:38:20.703",
            Dataframe7d:
              "7d20100eff924006ffff0101796700ff58ffff35887f56ff134015801a0029c02a",
            Dataframe80:
              "801c03fe7eff54ff208318080001000020378771009f053e05e7100000",
          },
          {
            Time: "12:38:21.086",
            Dataframe7d:
              "7d20100eff924006ffff0101796700ff56ffff35887f62ff134015801a0029c02a",
            Dataframe80:
              "801c04207eff54ff22831808000100002037877100a4053e05e5100000",
          },
          {
            Time: "12:38:21.597",
            Dataframe7d:
              "7d20100eff924006ffff0101796800ff58ffff35887f5bff134015801a0029c02a",
            Dataframe80:
              "801c04097eff54ff22841808000100002037877100a3053e05df100000",
          },
          {
            Time: "12:38:22.475",
            Dataframe7d:
              "7d20100eff924007ffff0101796a00ff58ffff35887f4bff134015801a0029c02a",
            Dataframe80:
              "801c04137eff54ff21841808000100002037877100a0053e05dc100000",
          },
          {
            Time: "12:38:23.337",
            Dataframe7d:
              "7d20100eff92400cffff0101796b00ff58ffff35887ee8ff134015801a0029c02a",
            Dataframe80:
              "801c03ce7eff54ff2284180800010000203787710102053b05e4100000",
          },
          {
            Time: "12:38:23.832",
            Dataframe7d:
              "7d20100eff924007ffff0101796b00ff58ffff35887ef4ff134015801a0029c02a",
            Dataframe80:
              "801c03937eff54ff2584180800010000203787710112053b05e5100000",
          },
          {
            Time: "12:38:24.662",
            Dataframe7d:
              "7d20100eff924092ffff0101796b00ff58ffff35887f4fff134015801a0029c02a",
            Dataframe80:
              "801c03dc7eff54ff24841808000100002037877100c5053d05e2100000",
          },
          {
            Time: "12:38:25.157",
            Dataframe7d:
              "7d20100eff924099ffff0101796a00ff58ffff35887f32ff134015801a0029c02a",
            Dataframe80:
              "801c03ee7eff54ff22841808000100002037877100c7053d0604100000",
          },
          {
            Time: "12:38:25.652",
            Dataframe7d:
              "7d20100eff92409dffff0101796a00ff58ffff35887f28ff134015801a0029c02a",
            Dataframe80:
              "801c03de7eff54ff23841808000100002037877000e0053c0607100000",
          },
          {
            Time: "12:38:26.035",
            Dataframe7d:
              "7d20100eff924093ffff0101796900ff58ffff35887f85ff134015801a0029c02a",
            Dataframe80:
              "801c03f27eff54ff2284180000010000203787710098054c0606100000",
          },
          {
            Time: "12:38:26.546",
            Dataframe7d:
              "7d20100eff9240a4ffff0101796900ff54ffff35887fa9ff134015801a0029c02a",
            Dataframe80:
              "801c044a7eff54ff21841800000100002037876f0058054c0605100000",
          },
          {
            Time: "12:38:27.376",
            Dataframe7d:
              "7d20100dff9240aaffff0101796800ff56ffff35887f78ff134015801a0029c02a",
            Dataframe80:
              "801c04417fff54ff1f831700000100002037876d0074054c0601100000",
          },
          {
            Time: "12:38:27.871",
            Dataframe7d:
              "7d20100dff92409affff0101796700ff54ffff35887f61ff134015801a0029c02a",
            Dataframe80:
              "801c041b7fff54ff20831700000100002037876d0097054c05ff100000",
          },
          {
            Time: "12:38:28.366",
            Dataframe7d:
              "7d20100eff9240a2ffff0101796700ff54ffff35887f4cff134015801a0029c02a",
            Dataframe80:
              "801c04027fff54ff20831700000100002037876f00b6054c0603100000",
          },
          {
            Time: "12:38:28.733",
            Dataframe7d:
              "7d20100dff92404effff0101796700ff54ffff35887f52ff134015801a0029c02a",
            Dataframe80:
              "801c03fb7fff54ff20831700000100002037876d00b2054c05fb100000",
          },
          {
            Time: "12:38:29.229",
            Dataframe7d:
              "7d20100dff924013ffff0101796700ff54ffff35887f5eff134015801a0029c02a",
            Dataframe80:
              "801c04017fff54ff22831700000100002037876d00a6054c05fb100000",
          },
          {
            Time: "12:38:30.074",
            Dataframe7d:
              "7d20100dff924017ffff0101796700ff54ffff35887f75ff134015801a0029c02a",
            Dataframe80:
              "801c04217fff54ff20841700000100002037876d0090054c05d9100000",
          },
          {
            Time: "12:38:30.569",
            Dataframe7d:
              "7d20100dff92405fffff0101796600ff56ffff35887f6eff134015801a0029c02a",
            Dataframe80:
              "801c04217fff54ff20841700000100002037876d008e054c05e0100000",
          },
          {
            Time: "12:38:31.080",
            Dataframe7d:
              "7d20100dff924095ffff0101796600ff54ffff35887f78ff134015801a0029c02a",
            Dataframe80:
              "801c041f7fff54ff20841700000100002037876d008f054c05e0100000",
          },
          {
            Time: "12:38:31.448",
            Dataframe7d:
              "7d20100dff92409bffff0101796500ff54ffff35887f77ff134015801a0029c02a",
            Dataframe80:
              "801c04257fff54ff1f841700000100002037876d0089054c05e1100000",
          },
          {
            Time: "12:38:31.974",
            Dataframe7d:
              "7d20100dff924013ffff0101796700ff54ffff35887f58ff134015801a0029c02a",
            Dataframe80:
              "801c04247fff54ff1f841700000100002037876d00a0054c05db100000",
          },
          {
            Time: "12:38:32.741",
            Dataframe7d:
              "7d20100dff92400effff0101796700ff54ffff35887f55ff134015801a0029c02a",
            Dataframe80:
              "801c040f7fff54ff20841700000100002037876d0096054c05d8100000",
          },
          {
            Time: "12:38:33.236",
            Dataframe7d:
              "7d20100dff92400dffff0101796800ff54ffff35887f5fff134015801a0029c02a",
            Dataframe80:
              "801c040e7fff55ff20841700000100002037876d009d054c05d9100000",
          },
          {
            Time: "12:38:33.730",
            Dataframe7d:
              "7d20100dff92400effff0101796800ff54ffff35887f79ff134015801a0029c02a",
            Dataframe80:
              "801c041f7fff55ff1f841700000100002037876d0086054c05da100000",
          },
          {
            Time: "12:38:34.545",
            Dataframe7d:
              "7d20100dff92409fffff0101796600ff54ffff35887f6eff134015801a0029c02a",
            Dataframe80:
              "801c04307fff55ff1f841700000100002037876d008b054c05df100000",
          },
          {
            Time: "12:38:35.039",
            Dataframe7d:
              "7d20100dff92408affff0101796600ff54ffff35887f65ff134015801a0029c02a",
            Dataframe80:
              "801c041d7fff55ff1f841700000100002037876d008f054c05f6100000",
          },
          {
            Time: "12:38:35.519",
            Dataframe7d:
              "7d20100dff92400effff0101796700ff54ffff35887f50ff134015801a0029c02a",
            Dataframe80:
              "801c03f17fff55ff21841700000100002037876d00c5054c05fd100000",
          },
          {
            Time: "12:38:35.998",
            Dataframe7d:
              "7d20100dff92400cffff0101796800ff54ffff35887f5dff134015801a0029c02a",
            Dataframe80:
              "801c03f87fff55ff22841700000100002037876d00b1054c0601100000",
          },
          {
            Time: "12:38:36.400",
            Dataframe7d:
              "7d20100dff92400dffff0101796800ff54ffff35887f5eff134015801a0029c02a",
            Dataframe80:
              "801c04087fff55ff20841700000100002037876d00a6054c05fd100000",
          },
          {
            Time: "12:38:37.051",
            Dataframe7d:
              "7d20100dff924056ffff0101796700ff54ffff35887f69ff134015801a0029c02a",
            Dataframe80:
              "801c041a7fff55ff20841700000100002037876d009c054c0602100000",
          },
          {
            Time: "12:38:37.435",
            Dataframe7d:
              "7d20100dff924036ffff0101796800ff54ffff35887f4fff134015801a0029c02a",
            Dataframe80:
              "801c04187fff55ff20841700000100002037876d009a054c05fa100000",
          },
          {
            Time: "12:38:37.834",
            Dataframe7d:
              "7d20100dff92400dffff0101796900ff54ffff35887f4dff134015801a0029c02a",
            Dataframe80:
              "801c04007fff55ff20841700000100002037876d00b3054c0601100000",
          },
          {
            Time: "12:38:38.520",
            Dataframe7d:
              "7d20100dff92408fffff0101796800ff54ffff35887f62ff134015801a0029c02a",
            Dataframe80:
              "801c040a7fff55ff20841700000100002037876d00a6054c0600100000",
          },
          {
            Time: "12:38:39.159",
            Dataframe7d:
              "7d20100dff92400fffff0101796900ff54ffff35887f5fff134015801a0029c02a",
            Dataframe80:
              "801c04197fff55ff21841700000100002037876d0093054c05ff100000",
          },
          {
            Time: "12:38:39.989",
            Dataframe7d:
              "7d20100dff924028ffff0101796900ff54ffff35887f75ff134015801a0029c02a",
            Dataframe80:
              "801c04157fff55ff1f841700000100002037876d0095054c05ff100000",
          },
          {
            Time: "12:38:40.516",
            Dataframe7d:
              "7d20100dff924097ffff0101796800ff54ffff35887f6dff134015801a0029c02a",
            Dataframe80:
              "801c041d7fff55ff20841700000100002037876d0094054c05fd100000",
          },
          {
            Time: "12:38:41.378",
            Dataframe7d:
              "7d20100dff924094ffff0101796700ff54ffff35887f63ff134015801a0029c02a",
            Dataframe80:
              "801c04227fff55ff20831700000100002037876d0091054c05f7100000",
          },
          {
            Time: "12:38:42.279",
            Dataframe7d:
              "7d20100dff92400dffff0101796900ff54ffff35887f5aff134015801a0029c02a",
            Dataframe80:
              "801c04047fff56ff20831700000100002037876d00ae054c05fd100000",
          },
          {
            Time: "12:38:43.166",
            Dataframe7d:
              "7d20100dff92401dffff0101796a00ff54ffff35887f71ff134015801a0029c02a",
            Dataframe80:
              "801c04137fff56ff20831700000100002037876d0099054c05fd100000",
          },
          {
            Time: "12:38:44.029",
            Dataframe7d:
              "7d20100fff924008ffff0101796a00ff53ffff35887eb0ff134015801a0029c02a",
            Dataframe80:
              "801c03eb7fff56ff21841700000100002037876d00dc053c05da100000",
          },
          {
            Time: "12:38:44.619",
            Dataframe7d:
              "7d20100fff92400dffff0101796a00ff53ffff35887e58ff134015801a0029c02a",
            Dataframe80:
              "801c02ea80ff56ff2a841700000100002037876c01aa054c05de100000",
          },
          {
            Time: "12:38:45.417",
            Dataframe7d:
              "7d201015ff9240a3ffff0101796904ff53ffff35887eeaff134015801a0029c02a",
            Dataframe80:
              "801c036c80ff56ff3c842008000100002037876c00f0054805d9100000",
          },
          {
            Time: "12:38:46.008",
            Dataframe7d:
              "7d20100fff9240adffff0101796800ff53ffff35887f77ff134015801a0029c02a",
            Dataframe80:
              "801c040580ff56ff39842108000100002037876c00a4054805da100000",
          },
          {
            Time: "12:38:46.392",
            Dataframe7d:
              "7d20100dff9240afffff0101796800ff53ffff35887f6aff134015801a0029c02a",
            Dataframe80:
              "801c03e380ff56ff25841800000100002037876c00b0054c05da100000",
          },
          {
            Time: "12:38:46.790",
            Dataframe7d:
              "7d20100dff924099ffff0101796800ff53ffff35887f6bff134015801a0029c02a",
            Dataframe80:
              "801c040a80ff56ff22841700000100002037876c0088054c05d5100000",
          },
          {
            Time: "12:38:47.210",
            Dataframe7d:
              "7d20100dff92400effff0101796900ff53ffff35887f1aff134015801a0029c02a",
            Dataframe80:
              "801c03e380ff56ff20841700000100002037876c00cf054c05d8100000",
          },
          {
            Time: "12:38:47.782",
            Dataframe7d:
              "7d20100dff924008ffff0101796a00ff53ffff35887f34ff134015801a0029c02a",
            Dataframe80:
              "801c03a680ff56ff24841700000100002037876c00e4054c05d6100000",
          },
          {
            Time: "12:38:48.163",
            Dataframe7d:
              "7d20100dff924008ffff0101796a00ff53ffff35887f45ff134015801a0029c02a",
            Dataframe80:
              "801c03d280ff56ff21841700000100002037876c00ba054c05d2100000",
          },
          {
            Time: "12:38:48.866",
            Dataframe7d:
              "7d20100dff924061ffff0101796900ff53ffff35887f7aff134015801a0029c02a",
            Dataframe80:
              "801c040f80ff56ff20841700000100002037876c007b054c05d3100000",
          },
          {
            Time: "12:38:49.249",
            Dataframe7d:
              "7d20100dff924024ffff0101796a00ff53ffff35887f86ff134015801a0029c02a",
            Dataframe80:
              "801c041380ff56ff20841700000100002037876c007e054c05d3100000",
          },
          {
            Time: "12:38:49.632",
            Dataframe7d:
              "7d20100dff924099ffff0101796900ff53ffff35887f8dff134015801a0029c02a",
            Dataframe80:
              "801c041d80ff56ff1f841700000100002037876c0078054c05d8100000",
          },
          {
            Time: "12:38:50.016",
            Dataframe7d:
              "7d20100dff92409dffff0101796800ff53ffff35887f81ff134015801a0029c02a",
            Dataframe80:
              "801c042280ff56ff20841700000100002037876c0070054c05d9100000",
          },
          {
            Time: "12:38:50.398",
            Dataframe7d:
              "7d20100dff924095ffff0101796800ff53ffff35887f7aff134015801a0029c02a",
            Dataframe80:
              "801c041f80ff56ff1f841700000100002037876c0083054c05da100000",
          },
          {
            Time: "12:38:50.797",
            Dataframe7d:
              "7d20100dff92409affff0101796800ff53ffff35887f6aff134015801a0029c02a",
            Dataframe80:
              "801c040f80ff56ff20841700000100002037876c0095054c05d6100000",
          },
          {
            Time: "12:38:51.404",
            Dataframe7d:
              "7d201011ff924013ffff0101796900ff53ffff35887f58ff134015801a0029c02a",
            Dataframe80:
              "801c040680ff56ff22841908000100002037876c0097053d05da100000",
          },
          {
            Time: "12:38:52.191",
            Dataframe7d:
              "7d201017ff9240b0ffff0101796710ff53ffff3588830aff134015801a0029c02a",
            Dataframe80:
              "801c068d80ff56ff29842308000100002037876b026e055605db100000",
          },
          {
            Time: "12:38:53.176",
            Dataframe7d:
              "7d201017ff92407cffff010179690cff53ffff358883e4ff134015801a0029c02a",
            Dataframe80:
              "801c084980ff56ff20842308000100002037876b03c7055a05dc100000",
          },
          {
            Time: "12:38:54.150",
            Dataframe7d:
              "7d201017ff9240a5ffff010179670cff53ffff35888426ff134015801a0029c02a",
            Dataframe80:
              "801c08af80ff56ff20842308000100002037876b0424055c05de100000",
          },
          {
            Time: "12:38:54.916",
            Dataframe7d:
              "7d201017ff9240a8ffff010179650cff53ffff35888433ff134015801a0029c02a",
            Dataframe80:
              "801c08c280ff56ff1f842308000100002037876b0437055c05de100000",
          },
          {
            Time: "12:38:55.794",
            Dataframe7d:
              "7d201017ff92409dffff0101796410ff53ffff3588842fff134015801a0029c02a",
            Dataframe80:
              "801c08c780ff56ff1e842308000100002037876b042e055c05e0100000",
          },
          {
            Time: "12:38:56.688",
            Dataframe7d:
              "7d201017ff924013ffff0101796510ff51ffff35888423ff134015801a0029c02a",
            Dataframe80:
              "801c08c980ff56ff1f842308000100002037876b0426055c05e2100000",
          },
          {
            Time: "12:38:57.583",
            Dataframe7d:
              "7d201016ff92401bffff010179670cff51ffff3588840dff134015801a0029c02a",
            Dataframe80:
              "801c08ba80ff56ff1f842308000100002037876b0417055c05df100000",
          },
          {
            Time: "12:38:58.493",
            Dataframe7d:
              "7d201015ff9240a2ffff010179650cff51ffff35888405ff134015801a0029c02a",
            Dataframe80:
              "801c08bb80ff56ff1e842208000100002037876a041a055c05e0100000",
          },
          {
            Time: "12:38:59.418",
            Dataframe7d:
              "7d201015ff9240a1ffff0101796408ff51ffff358883a5ff134015801a0029c02a",
            Dataframe80:
              "801c086680ff56ff1e842108000100002037876903c0055a05e2100000",
          },
          {
            Time: "12:39:00.345",
            Dataframe7d:
              "7d201015ff92409dffff0101796608ff51ffff3588839aff134015801a0029c02a",
            Dataframe80:
              "801c082c80ff56ff1e84210800010000203787690398055905df100000",
          },
          {
            Time: "12:39:01.206",
            Dataframe7d:
              "7d201015ff924043ffff0101796908ff51ffff3588837fff134015801a0029c02a",
            Dataframe80:
              "801c081780ff56ff1e8421080001000020378769038a055905df100000",
          },
          {
            Time: "12:39:02.100",
            Dataframe7d:
              "7d201015ff924097ffff0101796808ff51ffff3588838dff134015801a0029c02a",
            Dataframe80:
              "801c082080ff56ff1e84210800010000203787690389055905de100000",
          },
          {
            Time: "12:39:02.978",
            Dataframe7d:
              "7d20100fff924098ffff0101796800ff51ffff35888311ff134015801a0029c02a",
            Dataframe80:
              "801c082380ff56ff1f8421080001000020378769038e055905e3100000",
          },
          {
            Time: "12:39:03.824",
            Dataframe7d:
              "7d20000aff92400cffff0100796400ff51ffff35887ed2ff134015801a0029c02a",
            Dataframe80:
              "801c050c80ff56ff178414000001000020378752003105380672100001",
          },
          {
            Time: "12:39:04.367",
            Dataframe7d:
              "7d20000aff924006ffff0101796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c019881ff56ff30841400000100002037870a0374053f0672100001",
          },
          {
            Time: "12:39:05.133",
            Dataframe7d:
              "7d20000aff924006ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff56ff648211000001000020378715049705380ae5000001",
          },
          {
            Time: "12:39:05.948",
            Dataframe7d:
              "7d20000aff924005ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff56ff648211000001000020378705049705380ae5000001",
          },
          {
            Time: "12:39:06.874",
            Dataframe7d:
              "7d20000fff924004ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff56ff64811300000100002037876d049705380ae5000001",
          },
          {
            Time: "12:39:07.720",
            Dataframe7d:
              "7d20001bff924003ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff56ff6481280000010000203787b4049705380ae5000001",
          },
          {
            Time: "12:39:08.582",
            Dataframe7d:
              "7d200011ff924003ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff56ff64801d00000100002037876e049705380b16000001",
          },
          {
            Time: "12:39:09.524",
            Dataframe7d:
              "7d20000dff924002ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff55ff648018000001000020378721049705380b16000001",
          },
          {
            Time: "12:39:10.482",
            Dataframe7d:
              "7d20000aff924002ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff55ff647f1400000100002037870b049705380b48000001",
          },
          {
            Time: "12:39:11.120",
            Dataframe7d:
              "7d20000eff924001ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff55ff647f1300000100002037876b049705380b48000001",
          },
          {
            Time: "12:39:11.966",
            Dataframe7d:
              "7d20000fff924001ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff55ff647e19000001000020378775049705380b7a000001",
          },
          {
            Time: "12:39:12.733",
            Dataframe7d:
              "7d20000fff924001ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff55ff647e19000001000020378775049705380b7a000001",
          },
          {
            Time: "12:39:13.148",
            Dataframe7d:
              "7d20000fff924001ffff0100796400ff4fffff35887b69ff134015801a0029c02a",
            Dataframe80:
              "801c000081ff55ff647e19000001000020378775049705380b7a000001",
          },
        ],
      };
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
            let data = this.hex(Array.from(value).shift());

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

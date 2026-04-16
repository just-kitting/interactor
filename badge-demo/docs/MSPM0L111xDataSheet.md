**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **MSPM0L111x Mixed-Signal Microcontrollers** 

## **1 Features** 

- **Core** 

   - Arm[®] 32-bit Cortex[®] -M0+ CPU with memory protection unit, frequency up to 32MHz 

- PSA-L1 Certification targeted 

- **Operating characteristics** 

   - Extended temperature: –40°C to 125°C 

   - Wide supply voltage range: 1.62V to 3.6V 

- **Memories** 

   - Up to 128KB of flash memory with error correction code (ECC) 

      - Dual-bank with address swap for OTA updates 

   - 16KB of SRAM 

- **High-performance analog peripherals** 

   - One 12-bit 1.68Msps analog-to-digital converter (ADC) with up to 13 external channels 

      - 14-bit effective resolution at 105ksps with hardware averaging 

   - Configurable 1.4V or 2.5V internal voltage reference (VREF) 

      - Integrated temperature sensor 

   - 

- **Optimized low-power modes** 

   - RUN: 106µA/MHz (CoreMark) 

   - SLEEP: 50µA/MHz 

   - STOP: 239µA at 4MHz 

   - STANDBY: 1.5µA at 32kHz with RTC and full SRAM and state retention 

   - SHUTDOWN: 75nA with IO wake-up capability 

- **Intelligent digital peripherals** 

   - 3-channel DMA controller 

      - 3-channel event fabric signaling system 

   - 

   - A total of 14 PWM channels supported by: 

      - One 16-bit advanced timer with deadband support and complimentary outputs up to 8 PWM channels 

            - One I[2] C interface supporting up to FM+ (1Mbit/s), SMBus/PMBus, and wakeup from STOP mode 

            - One SPI interface supporting up to 16Mbit/s 

         - **Clock system** 

            - Internal 4 to 32MHz oscillator (SYSOSC) with up to ±1.2% accuracy 

            - Internal 32kHz low-frequency oscillator (LFOSC) with ±3% accuracy 

            - External 32-kHz crystal oscillator (LFXT) 

         - **Data integrity and encryption** 

            - AES-128/256 accelerator with support for GCM/ GMAC, CCM/CBC-MAC, CBC, CTR 

            - Secure key storage for up to two AES keys 

            - Flexible firewalls for protecting code and data 

            - True random number generator (TRNG) 

            - Cyclic redundancy checker (CRC-16, CRC-32) 

         - **Flexible I/O features** 

            - Up to 44 GPIOs 

               - Two 5V-tolerant open-drain IOs 

               - Seven high-drive IOs with 20mA drive strength 

               - One high-speed IO 

         - **Development support** 

         - 2-pin serial wire debug (SWD) 

         - • **Package options** 

            - 48-pin LQFP (PT) (0.5mm pitch) 

            - 48-pin VQFN (RGZ) (0.5mm pitch) 

            - 32-pin VQFN (RHB) (0.5mm pitch) 

            - 24-pin VQFN (RGE) (0.5mm pitch) 

         - **Family members** (also see _Device Comparison_ ) 

            - MSPM0L1116: 64KB of flash, 16KB of RAM 

         - MSPM0L1117: 128KB of flash, 16KB of RAM 

         - • **Development kits and software** (also see _Tools and Software_ ) 

            - LP-MSPM0L1117 LaunchPad[™] development kit 

            - MSP Software Development Kit (SDK) 

      - Two 16-bit general-purpose timers support low-power operation in STANDBY mode 

      - One 16-bit general purpose timer supporting QEI 

   - One windowed watchdog timer (WWDT) 

   - One independent watchdog timer (IWDT) 

   - RTC with alarm and calendar mode 

- **Enhanced communication interfaces** 

   - Two UART interfaces supporting low-power operation in STANDBY mode 

      - One supporting LIN, IrDA, DALI, Smart Card, Manchester 

An IMPORTANT NOTICE at the end of this data sheet addresses availability, warranty, changes, use in safety-critical applications, intellectual property matters and other important disclaimers. PRODUCTION DATA. 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **2 Applications** 

- Battery charging and management 

- Power supplies and power delivery 

- Personal electronics 

- Building security and fire safety 

- Connected peripherals and printers 

- Energy Infrastructure - Smart Metering 

- Smart metering 

- Communication modules 

- Medical and healthcare 

- Lighting 

Copyright © 2025 Texas Instruments Incorporated 

2 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **3 Description** 

MSPM0L111x microcontrollers (MCUs) are part of the MSP highly-integrated, ultra-low-power 32-bit MCU family based on the enhanced Arm[®] Cortex[®] -M0+ core platform operating at up to 32MHz frequency. These cost-optimized MCUs offer high-performance analog peripheral integration and excellent low-power current consumption. The MCUs also support extended temperature ranges from -40°C to 125°C and operate with supply voltages ranging from 1.62V to 3.6V. 

The device has up to 128KB of embedded flash memory with built-in error correction code (ECC) and up to 16KB SRAM. The flash memory is organized into two main banks to support field firmware updates, with address swap support provided between the two main banks. 

Flexible cybersecurity enablers can be used to support secure boot, secure in-field firmware updates, IP protection (execute-only memory), key storage, and more. Hardware acceleration is provided for a variety of AES symmetric cipher modes, as well as a TRNG entropy source. The cybersecurity architecture is pending Arm® PSA Level 1 certification. 

These MCUs incorporate a high-speed on-chip oscillator with an accuracy up to ±1.2%, eliminating the need for an external crystal. Additional features include a 3-channel DMA, 16-bit/32-bit CRC accelerator, and a variety of high-performance analog peripherals such as one 12-bit 1.68Msps ADC with configurable internal voltage reference and an on-chip temperature sensor. These devices also offer intelligent digital peripherals such as one 16-bit advanced control timer and two 16-bit general purpose timers, one general purpose timer with Quadrature Enabled Input, windowed and independent watchdog timer, and a variety of communication peripherals including one I[2] C, one SPI, and two UARTs (one with support for LIN protocol). 

The TI MSPM0 family of low-power MCUs consists of devices with varying degrees of analog and digital integration allowing for customers to find the MCU that meets their project's needs. The architecture combined with extensive low-power modes are optimized to achieve extended battery life in portable measurement applications. 

MSPM0L111x MCUs are supported by an extensive hardware and software ecosystem with reference designs and code examples to get the design started quickly. Development kits include a LaunchPad[™] development kit available for purchase and design files for a target-socket board. TI also provides a free MSP Software Development Kit (SDK), which is available as a component of Code Composer Studio[™] IDE desktop and cloud version within the TI Resource Explorer. MSPM0 MCUs are also supported by extensive online collateral, training with MSP Academy, and online support through the TI E2E[™] support forums. 

For complete module descriptions, see the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **CAUTION** 

System-level ESD protection must be applied in compliance with the device-level ESD specification to prevent electrical overstress or disturbing of data or code memory. See MSP430™ System-Level ESD Considerations for more information; the principles in this application note are applicable to MSPM0 MCUs. 

## **Device Information** 

|**PART NUMBER**|**PACKAGE**(1)|**PACKAGE SIZE**(2)|
|---|---|---|
|MSPM0L1116SRGER|RGE (VQFN, 24)|4mm x 4mm|
|MSPM0L1117SRGER|||
|MSPM0L1116SRHBR|RHB (VQFN, 32)|5mm x 5mm|
|MSPM0L1117SRHBR|||
|MSPM0L1116SRGZR|RGZ (VQFN, 48)|7mm x 7mm|
|MSPM0L1117SRGZR|||



Copyright © 2025 Texas Instruments Incorporated 

3 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **MSPM0L1117, MSPM0L1116** 

|**Device Information(continued)**|**Device Information(continued)**|**Device Information(continued)**|
|---|---|---|
|**PART NUMBER**|**PACKAGE**(1)|**PACKAGE SIZE**(2)|
|MSPM0L1116SPTR|PT (LQFP, 48)|9mm x 9mm|
|MSPM0L1117SPTR|||



(1) For more information, see Mechanical, Packaging, and Orderable Information. 

(2) The package size (length x width) is a nominal value and includes pins, where applicable. 

Copyright © 2025 Texas Instruments Incorporated 

4 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **4 Functional Block Diagram** 

**==> picture [436 x 449] intentionally omitted <==**

**----- Start of picture text -----**<br>
PAx, PBx<br>tIOBUSt ULPCLK<br>CPU SUB SYSTEM GPIO DMA<br>3-ch<br>Arm<br>Cortex-M0+ FLASH B0 CRC-P<br>fmax = 32 MHz Up to 64KB 16/32-bit<br>NVIC FLASH B1<br>MPU Up to 64KB SPI0 POCI, PICO,<br>SCK, CSx<br>SWD + MTB<br>SRAM TEMP SENSOR<br>IOPORT 16KB<br>TRNG ADC0 13-CH (EXT)<br>12-bit A0_x<br>ULPCLK<br>IOMUX PD1 PERIPHERAL BUS (MCLK) AES-ADV<br>128/256-bit<br>WWDT0<br>SWCLK,<br>DEBUG<br>SWDIO<br>IWDT<br>RTC_OUT RTC_B KEY<br>STORE<br>EVENT FLASHCTL<br>CTS, RTSTX, RX,  UART0 PMCU (SYSCTL) VREF VREF+, VREF-<br>TX, RX,<br>CTS, RTS UART1 TIMG0<br>CKM PMU 2-CH<br>TIMG1<br>SDA, SCL I2C0 LFOSC LDO<br>SYSOSC BOR TIMG8 2-CH<br>LFXT QEI/HALL<br>POR<br>4-CH<br>TIMA0<br>VBOOST FAULT<br>LEGEND<br>PD1, CPU ACCESS ONLY LFXIN, LFXOUT VDD, VSS<br>PD1, CPU/DMA ACCESS ROSC VCORE, NRST<br>PD1/PD0, CPU/DMA ACCESS CLK_OUT, FCC_IN<br>PD0, CPU/DMA ACCESS<br>ROM<br>BCR, BSL<br>AHB BUS (MCLK) PD1 PERIPHERAL BUS (MCLK)<br>PD0 PERIPHERAL BUS (ULPCLK)<br>PD0 PERIPHERAL BUS (ULPCLK)<br>**----- End of picture text -----**<br>


**Figure 4-1. MSPM0L111x Functional Block Diagram** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

5 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **Table of Contents** 

|**1**|**Features**............................................................................1|8.8 Memory.....................................................................49|
|---|---|---|
|**2**|**Applications**.....................................................................2|8.9 Flash Memory...........................................................51|
|**3**|**Description**.......................................................................3|8.10 SRAM......................................................................51|
|**4**|**Functional Block Diagram**..............................................5|8.11 GPIO.......................................................................51|
|**5**|**Device Comparison**.........................................................7|8.12 IOMUX....................................................................52|
||5.1 Device Comparison Table...........................................8|8.13 ADC........................................................................52|
|**6**|**Pin Configuration and Functions**...................................9|8.14 Temperature Sensor...............................................53|
||6.1 Pin Diagrams..............................................................9|8.15 VREF......................................................................54|
||6.2 Pin Attributes.............................................................12|8.16 Security...................................................................54|
||6.3 Signal Descriptions...................................................19|8.17 TRNG......................................................................55|
||6.4 Connections for Unused Pins...................................25|8.18 AESADV.................................................................55|
|**7**|**Specifications**................................................................26|8.19 Keystore..................................................................55|
||7.1 Absolute Maximum Ratings......................................26|8.20 CRC-P.....................................................................55|
||7.2 ESD Ratings.............................................................26|8.21 UART......................................................................56|
||7.3 Recommended Operating Conditions.......................26|8.22 I2C..........................................................................56|
||7.4 Thermal Information..................................................27|8.23 SPI..........................................................................57|
||7.5 Supply Current Characteristics.................................27|8.24 Low-Frequency Sub System (LFSS)......................57|
||7.6 Power Supply Sequencing........................................29|8.25 RTC_B....................................................................57|
||7.7 Flash Memory Characteristics..................................30|8.26 IWDT_B..................................................................58|
||7.8 Timing Characteristics...............................................31|8.27 WWDT....................................................................58|
||7.9 Clock Specifications..................................................32|8.28 Timers (TIMx)..........................................................59|
||7.10 Digital IO.................................................................33|8.29 Device Analog Connections....................................60|
||7.11 Analog Mux VBOOST.............................................36|8.30 Input/Output Diagrams............................................60|
||7.12 ADC........................................................................36|8.31 Serial Wire Debug Interface....................................61|
||7.13 Temperature Sensor...............................................38|8.32 Bootstrap Loader (BSL)..........................................62|
||7.14 VREF......................................................................39|8.33 Device Factory Constants.......................................62|
||7.15 I2C..........................................................................39|8.34 Identification............................................................63|
||7.16 SPI..........................................................................40|**9 Applications, Implementation, and Layout**.................64|
||7.17 UART......................................................................42|9.1 Typical Application....................................................64|
||7.18 TIMx........................................................................42|**10 Device and Documentation Support**..........................65|
||7.19 TRNG Electrical Characteristics.............................42|10.1 Device Nomenclature..............................................65|
||7.20 TRNG Switching Characteristics.............................42|10.2 Tools and Software.................................................66|
||7.21 Emulation and Debug.............................................43|10.3 Documentation Support..........................................66|
|**8**|**Detailed Description**......................................................44|10.4 Support Resources.................................................67|
||8.1 Functional Block Diagram.........................................44|10.5 Trademarks.............................................................67|
||8.2 CPU..........................................................................44|10.6 Electrostatic Discharge Caution..............................67|
||8.3 Operating Modes......................................................45|10.7 Glossary..................................................................67|
||8.4 Power Management Unit (PMU)...............................47|**11 Revision History**..........................................................67|
||8.5 Clock Module (CKM).................................................47|**12 Mechanical, Packaging, and Orderable**|
||8.6 DMA..........................................................................47|**Information**....................................................................69|
||8.7 Events.......................................................................48||



Copyright © 2025 Texas Instruments Incorporated 

6 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **5 Device Comparison** 

The following table summarizes the features of each device that is described in this data sheet. 

**Table 5-1. Device Comparison** 

|**DEVICE NAME**(1) (2)|**FLASH / SRAM (KB)**|**QUAL**(3)|**UART/I2C/SPI**|**ADC CH**|**GPIO**|**PACKAGE (PACKAGE SIZE)**(4)|
|---|---|---|---|---|---|---|
|MSPM0L1117SPTR|128 / 16|S|2 / 1 / 1|13|44|48 LQFP<br>(0.5mm pitch)<br>(9mm × 9mm)|
|MSPM0L1116SPTR|64 / 16|S|2 / 1 / 1|13|44||
|MSPM0L1117SRGZR|128 / 16|S|2 / 1 / 1|13|44|48 VQFN<br>(0.5mm pitch)<br>(7mm × 7mm)|
|MSPM0L1116SRGZR|64 / 16|S|2 / 1 / 1|13|44||
|MSPM0L1117SRHBR|128 / 16|S|2 / 1 / 1|11|28|32 VQFN<br>(0.5mm pitch)<br>(5mm × 5mm)|
|MSPM0L1116SRHBR|64 / 16|S|2 / 1 / 1|11|28||
|MSPM0L1117SRGER|128 / 16|S|2 / 1 / 1|7|20|24 VQFN<br>(0.5mm pitch)<br>(4mm × 4mm)|
|MSPM0L1116SRGER|64 / 16|S|2 / 1 / 1|7|20||



(1) For the most current part, package, and ordering information for all available devices, see the _Package Option Addendum_ in Section 12, or see the TI website. 

(2) For more information about the device name, see Section 10.1. 

(3) Device qualifications: 

   - S = –40°C to 125°C 

- (4) The package size (length × width) is a nominal value and includes pins, where applicable. For the package dimensions with tolerances, see the _Mechanical Data_ in Section 12. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

7 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **5.1 Device Comparison Table** 

**==> picture [422 x 91] intentionally omitted <==**

**----- Start of picture text -----**<br>
128 KB MSPM0L1117SRGER MSPM0L1117SRHBR   MSPM0L1117SRGZR   MSPM0L1117SPTR<br>64 KB MSPM0L1116SRGER   MSPM0L1116SRHBR   MSPM0L1116SRGZR   MSPM0L1116SPTR<br>24-pin 32-pin 48-pin 48-pin<br>VQFN (0.5mm) VQFN (0.5mm) LQFP (0.5mm) LQFP (0.5mm)<br>**----- End of picture text -----**<br>


Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

8 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **6 Pin Configuration and Functions** 

The System Configuration tool provides a graphical interface to enable, configurable, and generate initialization code for pin multiplexing and simplifying pin settings. The pin diagrams shown in the data sheet show the primary peripheral functions, some of the integrated device features, and available clock signals to simplify the device pinout. 

For full descriptions of the pin functions, see the _Pin Attributes_ and _Signal Descriptions_ sections. 

## **6.1 Pin Diagrams** 

For full pin configuration and functions for each package option, refer to Section 6.2 and Section 6.3 

**==> picture [360 x 364] intentionally omitted <==**

**----- Start of picture text -----**<br>
PA0 1 36  PB17<br>PA1 2 35  PA20<br>PA28 3 34  PA19<br>NRST 4 33  PA18<br>PA31 5 32  PA17<br>VDD 6 31  PA16<br>VSS 7 30  PA15<br>PA2 8 29  PA14<br>PA3 9 28  PA13<br>PA4 10 27  PA12<br>PA5 11 26  PB16<br>PA6 12 25  PB15<br>Not to scale<br>VCORE PA27 PA26 PA25 PA24 PA23 PB24 PB20 PA22 PA21 PB19 PB18<br>48 47 46 45 44 43 42 41 40 39 38 37<br>13 14 15 16 17 18 19 20 21 22 23 24<br>PA7 PB2 PB3 PA8 PA9 PA10 PA11 PB6 PB7 PB8 PB9 PB14<br>**----- End of picture text -----**<br>


**Figure 6-1. 48-pin PT (0.5mm) (LQFP) Package Diagram** 

Copyright © 2025 Texas Instruments Incorporated 

9 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

**==> picture [335 x 339] intentionally omitted <==**

**----- Start of picture text -----**<br>
PA0 1 36  PB17<br>PA1 2 35  PA20<br>PA28 3 34  PA19<br>NRST 4 33  PA18<br>PA31 5 32  PA17<br>VDD 6 31  PA16<br>Thermal<br>VSS 7 Pad 30  PA15<br>PA2 8 29  PA14<br>PA3 9 28  PA13<br>PA4 10 27  PA12<br>PA5 11 26  PB16<br>PA6 12 25  PB15<br>Not to scale<br>VCORE PA27 PA26 PA25 PA24 PA23 PB24 PB20 PA22 PA21 PB19 PB18<br>48 47 46 45 44 43 42 41 40 39 38 37<br>13 14 15 16 17 18 19 20 21 22 23 24<br>PA7 PB2 PB3 PA8 PA9 PA10 PA11 PB6 PB7 PB8 PB9 PB14<br>**----- End of picture text -----**<br>


**Figure 6-2. 48-pin RGZ (0.5mm) (VQFN) Package Diagram** 

Copyright © 2025 Texas Instruments Incorporated 

10 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

**==> picture [263 x 266] intentionally omitted <==**

**----- Start of picture text -----**<br>
PA0 1 24  PA20<br>PA1 2 23  PA19<br>NRST 3 22  PA18<br>VDD 4 21  PA17<br>Thermal<br>VSS 5 Pad 20  PA16<br>PA2 6 19  PA15<br>PA3 7 18  PA14<br>PA4 8 17  PA13<br>Not to scale<br>VCORE PA27 PA26 PA25 PA24 PA23 PA22 PA21<br>32 31 30 29 28 27 26 25<br>9 10 11 12 13 14 15 16<br>PA5 PA6 PA7 PA8 PA9 PA10 PA11 PA12<br>**----- End of picture text -----**<br>


**Figure 6-3. 32-pin RHB (0.5mm) (VQFN) Package Diagram** 

**==> picture [227 x 230] intentionally omitted <==**

**----- Start of picture text -----**<br>
PA1 1 18  PA22<br>NRST 2 17  PA21<br>VDD 3 16  PA20<br>Thermal<br>VSS 4 Pad 15  PA19<br>PA2 5 14  PA18<br>PA3 6 13  PA17<br>Not to scale<br>PA0 VCORE PA26 PA25 PA24 PA23<br>24 23 22 21 20 19<br>7 8 9 10 11 12<br>PA4 PA9 PA10 PA11 PA15 PA16<br>**----- End of picture text -----**<br>


**Figure 6-4. 24-pin RGE (0.5mm) (VQFN) Package Diagram** 

Copyright © 2025 Texas Instruments Incorporated 

11 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **6.2 Pin Attributes** 

The following table describes the functions available on every pin for each device package. 

## **Note** 

Each digital I/O on a device is mapped to a specific Pin Control Management Register (PINCMx) that lets users configure the desired _Pin Function_ using the PINCM.PF control bits. 

The IOMUX only supports connecting one IOMUX-managed digital function to the pin at the same time. The PINCM.PF and PINCM.PC in Section 8.12 are recommended to be set to 0 when non-IOMUX managed functions (such as analog connections) are intended to be used on a pin. However, non-IOMUX managed signals (such as analog inputs and WAKE inputs) can be enabled on a pin at the same time that an IOMUX managed digital function is enabled on the pin, provided there is no contention between the functions. In this case, the designer must verify that no contention exists between the functions enabled on each pin. 

**Table 6-1. Digital IO Features by IO Type** 

|**IO STRUCTURE**|**INVERSION**<br>**CONTROL**|**DRIVE**<br>**STRENGTH**<br>**CONTROL**|**HYSTERESIS**<br>**CONTROL**|**PULLUP**<br>**RESISTOR**|**PULLDOWN**<br>**RESISTOR**|**WAKEUP**<br>**LOGIC**|
|---|---|---|---|---|---|---|
|SDIO (Standard drive)|Y|||Y|Y||
|SDIO (Standard drive) with wake(1)|Y|||Y|Y|Y|
|HDIO (High drive)|Y|Y||Y|Y|Y|
|HSIO (High speed)|Y|Y||Y|Y||
|ODIO (5V-tolerant open drain)|Y||Y||Y|Y|



1. Standard with Wake allows the I/O to wake up the device from the lowest low-power mode of SHUTDOWN. All I/O can be configured to wakeup the MCU from higher low-power modes. See section _GPIO FastWake_ in the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ for details. 

**Table 6-2. Pin Attributes** 

|**RGE**<br>**PIN**|**RHB**<br>**PIN**|**RGZ**<br>**PIN**|**PT**<br>**PIN**|**PIN NAME/**<br>**IOMUX REG/**<br>**IOMUX ADDR**|**SIGNAL**<br>**NAME**|**IOMUX**<br>**PF**|**SIGNAL**<br>**TYPE**|**BUFFER**<br>**TYPE**|
|---|---|---|---|---|---|---|---|---|
|2|3|4|4|NRST|NRST|(Non-IOMUX 1) 0|I|RESET|
||||||WAKE|(Non-IOMUX 2) 0|I||
|24|1|1|1|PA0<br>PINCM1<br>0x40428000|PA0|1|IO|ODIO (5V-tol)|
||||||UART0_TX|2|O||
||||||I2C0_SDA|3|IOD||
||||||TIMA0_C0|4|IO||
||||||TIMA_FAL1|5|I||
||||||FCC_IN|6|I||
||||||TIMG8_C1|7|IO||
||||||TIMG0_C0|9|IO||
||||||BSLSDA|(Non-IOMUX 1) 0|IOD||
||||||WAKE|(Non-IOMUX 2) 0|I||



Copyright © 2025 Texas Instruments Incorporated 

12 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

**Table 6-2. Pin Attributes (continued)** 

|**RGE**<br>**PIN**|**RHB**<br>**PIN**|**RGZ**<br>**PIN**|**PT**<br>**PIN**|**PIN NAME/**<br>**IOMUX REG/**<br>**IOMUX ADDR**|**SIGNAL**<br>**NAME**|**IOMUX**<br>**PF**|**SIGNAL**<br>**TYPE**|**BUFFER**<br>**TYPE**|
|---|---|---|---|---|---|---|---|---|
|1|2|2|2|PA1<br>PINCM2<br>0x40428004|PA1|1|IO|ODIO (5V-tol)|
||||||UART0_RX|2|I||
||||||I2C0_SCL|3|IOD||
||||||TIMA0_C1|4|IO||
||||||TIMA_FAL2|5|I||
||||||TIMG8_IDX|6|I||
||||||TIMG8_C0|7|IO||
||||||TIMG0_C1|9|IO||
||||||SPI0_CS3|10|IO||
||||||BSLSCL|(Non-IOMUX 1) 0|IOD||
||||||WAKE|(Non-IOMUX 2) 0|I||
|5|6|8|8|PA2<br>PINCM7<br>0x40428018|PA2|1|IO|SDIO<br>(standard)|
||||||TIMG8_C1|2|IO||
||||||SPI0_CS0|3|IO||
||||||TIMA0_C3N|6|O||
||||||TIMA0_C2N|7|O||
||||||TIMA_FAL0|8|I||
||||||TIMA_FAL1|9|I||
||||||TIMA0_C0|11|IO||
||||||ROSC|(Non-IOMUX 1) 0|A||
|6|7|9|9|PA3<br>PINCM8<br>0x4042801c|PA3|1|IO|SDIO<br>(standard)|
||||||TIMG8_C0|2|IO||
||||||SPI0_CS1|3|IO||
||||||TIMA0_C1|5|IO||
||||||TIMA0_C2|8|IO||
||||||UART1_TX|10|O||
||||||SPI0_CS3|11|IO||
||||||LFXIN|(Non-IOMUX 1) 0|A||
|7|8|10|10|PA4<br>PINCM9<br>0x40428020|PA4|1|IO|SDIO<br>(standard)|
||||||TIMG8_C1|2|IO||
||||||SPI0_POCI|3|IO||
||||||TIMA0_C1N|5|O||
||||||LFCLK_IN|6|I||
||||||TIMA0_C3|8|IO||
||||||UART1_RX|10|I||
||||||SPI0_CS0|11|IO||
||||||LFXOUT|(Non-IOMUX 1) 0|A||
||9|11|11|PA5<br>PINCM10<br>0x40428024|PA5|1|IO|SDIO<br>(standard)|
||||||TIMG8_C0|2|IO||
||||||SPI0_PICO|3|IO||
||||||SPI0_POCI|4|IO||
||||||TIMG0_C0|5|IO||
||||||FCC_IN|6|I||
||||||TIMA_FAL1|8|I||
||||||UART0_CTS|9|I||
||||||UART1_TX|11|O||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 13 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **MSPM0L1117, MSPM0L1116** 

**Table 6-2. Pin Attributes (continued)** 

|**RGE**<br>**PIN**|**RHB**<br>**PIN**|**RGZ**<br>**PIN**|**PT**<br>**PIN**|**PIN NAME/**<br>**IOMUX REG/**<br>**IOMUX ADDR**|**SIGNAL**<br>**NAME**|**IOMUX**<br>**PF**|**SIGNAL**<br>**TYPE**|**BUFFER**<br>**TYPE**|
|---|---|---|---|---|---|---|---|---|
||10|12|12|PA6<br>PINCM11<br>0x40428028|PA6|1|IO|SDIO<br>(standard)|
||||||TIMG8_C1|2|IO||
||||||SPI0_SCK|3|IO||
||||||TIMG0_C1|5|IO||
||||||HFCLK_IN|6|I||
||||||TIMA_FAL0|8|I||
||||||UART0_RTS|9|O||
||||||TIMA0_C2N|10|O||
||||||UART1_RX|11|I||
||11|13|13|PA7<br>PINCM14<br>0x40428034|PA7|1|IO|SDIO<br>(standard)|
||||||CLK_OUT|3|O||
||||||TIMG8_C0|4|IO||
||||||TIMA0_C2|5|IO||
||||||TIMG8_IDX|6|I||
||||||TIMA0_C1|8|IO||
||||||SPI0_CS2|9|IO||
||||||FCC_IN|10|I||
||||||SPI0_POCI|11|IO||
||12|16|16|PA8<br>PINCM19<br>0x40428048|PA8|1|IO|SDIO<br>(standard)|
||||||UART1_TX|2|O||
||||||SPI0_CS0|3|IO||
||||||I2C0_SDA|4|IOD||
||||||TIMA0_C0|5|IO||
||||||TIMA_FAL2|6|I||
||||||TIMA_FAL0|7|I||
||||||SPI0_CS3|8|IO||
||||||HFCLK_IN|10|I||
||||||UART0_RTS|11|O||
|8|13|17|17|PA9<br>PINCM20<br>0x4042804c|PA9|1|IO|HSIO (high-<br>speed)|
||||||UART1_RX|2|I||
||||||SPI0_PICO|3|IO||
||||||I2C0_SCL|4|IOD||
||||||TIMA0_C0N|5|O||
||||||CLK_OUT|6|O||
||||||TIMA0_C1|7|IO||
||||||RTC_OUT|8|O||
||||||SPI0_CS0|10|IO||
||||||UART0_CTS|11|I||



Copyright © 2025 Texas Instruments Incorporated 

14 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

**Table 6-2. Pin Attributes (continued)** 

|**RGE**<br>**PIN**|**RHB**<br>**PIN**|**RGZ**<br>**PIN**|**PT**<br>**PIN**|**PIN NAME/**<br>**IOMUX REG/**<br>**IOMUX ADDR**|**SIGNAL**<br>**NAME**|**IOMUX**<br>**PF**|**SIGNAL**<br>**TYPE**|**BUFFER**<br>**TYPE**|
|---|---|---|---|---|---|---|---|---|
|9|14|18|18|PA10<br>PINCM21<br>0x40428050|PA10|1|IO|HDIO (high-<br>drive)|
||||||UART0_TX|2|O||
||||||SPI0_POCI|3|IO||
||||||I2C0_SDA|4|IOD||
||||||TIMA0_C2|5|IO||
||||||CLK_OUT|6|O||
||||||TIMG0_C0|7|IO||
||||||TIMA_FAL1|10|I||
||||||I2C0_SCL|11|IOD||
||||||BSLTX|(Non-IOMUX 1) 0|O||
||||||WAKE|(Non-IOMUX 2) 0|I||
|10|15|19|19|PA11<br>PINCM22<br>0x40428054|PA11|1|IO|HDIO (high-<br>drive)|
||||||UART0_RX|2|I||
||||||SPI0_SCK|3|IO||
||||||I2C0_SCL|4|IOD||
||||||TIMA0_C2N|5|O||
||||||TIMG0_C1|7|IO||
||||||TIMA_FAL0|10|I||
||||||I2C0_SDA|11|IOD||
||||||BSLRX|(Non-IOMUX 1) 0|I||
||||||WAKE|(Non-IOMUX 2) 0|I||
||16|27|27|PA12<br>PINCM34<br>0x40428084|PA12|1|IO|SDIO<br>(standard)|
||||||SPI0_SCK|3|IO||
||||||TIMA0_C3|5|IO||
||||||FCC_IN|6|I||
||||||TIMG0_C0|7|IO||
||||||SPI0_CS1|9|IO||
||||||UART1_CTS|11|I||
||||||A0_8|(Non-IOMUX 1) 0|A||
||17|28|28|PA13<br>PINCM35<br>0x40428088|PA13|1|IO|SDIO<br>(standard)|
||||||SPI0_POCI|3|IO||
||||||TIMA0_C3N|5|O||
||||||RTC_OUT|6|O||
||||||TIMG0_C1|7|IO||
||||||SPI0_CS3|9|IO||
||||||UART1_RTS|11|O||
||||||A0_9|(Non-IOMUX 1) 0|A||
||18|29|29|PA14<br>PINCM36<br>0x4042808c|PA14|1|IO|SDIO<br>(standard)|
||||||UART0_CTS|2|I||
||||||SPI0_PICO|3|IO||
||||||CLK_OUT|6|O||
||||||SPI0_CS2|9|IO||
||||||A0_12|(Non-IOMUX 1) 0|A||



Copyright © 2025 Texas Instruments Incorporated 

15 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **MSPM0L1117, MSPM0L1116** 

**Table 6-2. Pin Attributes (continued)** 

|**RGE**<br>**PIN**|**RHB**<br>**PIN**|**RGZ**<br>**PIN**|**PT**<br>**PIN**|**PIN NAME/**<br>**IOMUX REG/**<br>**IOMUX ADDR**|**SIGNAL**<br>**NAME**|**IOMUX**<br>**PF**|**SIGNAL**<br>**TYPE**|**BUFFER**<br>**TYPE**|
|---|---|---|---|---|---|---|---|---|
|11|19|30|30|PA15<br>PINCM37<br>0x40428090|PA15|1|IO|HDIO (high-<br>drive)|
||||||UART0_RTS|2|O||
||||||TIMA0_C2|5|IO||
||||||UART0_TX|6|O||
||||||TIMG8_IDX|7|I||
||||||TIMG1_C0|8|IO||
||||||WAKE|(Non-IOMUX 1) 0|I||
|12|20|31|31|PA16<br>PINCM38<br>0x40428094|PA16|1|IO|HDIO (high-<br>drive)|
||||||TIMA0_C2N|5|O||
||||||UART0_RX|6|I||
||||||FCC_IN|7|I||
||||||TIMG1_C1|8|IO||
||||||TIMA0_C0|11|IO||
||||||WAKE|(Non-IOMUX 1) 0|I||
||||||A0_13|(Non-IOMUX 2) 0|A||
|13|21|32|32|PA17<br>PINCM39<br>0x40428098|PA17|1|IO|HDIO (high-<br>drive)|
||||||UART1_TX|2|O||
||||||TIMA0_C3|5|IO||
||||||TIMG8_C0|6|IO||
||||||SPI0_CS1|8|IO||
||||||WAKE|(Non-IOMUX 1) 0|I||
||||||A0_14|(Non-IOMUX 2) 0|A||
|14|22|33|33|PA18<br>PINCM40<br>0x4042809c|PA18|1|IO|HDIO (high-<br>drive)|
||||||UART1_RX|2|I||
||||||TIMA0_C3N|5|O||
||||||TIMG8_C1|6|IO||
||||||SPI0_CS0|8|IO||
||||||TIMA0_C1|11|IO||
||||||BSL_invoke|(Non-IOMUX 1) 0|I||
||||||WAKE|(Non-IOMUX 2) 0|I||
||||||A0_4|(Non-IOMUX 3) 0|A||
|15|23|34|34|PA19<br>PINCM41<br>0x404280a0|PA19|1|IO|SDIO<br>(standard)|
||||||SWDIO|2|IO||
||||||TIMA0_C2|5|IO||
||||||TIMG0_C0|6|IO||
|16|24|35|35|PA20<br>PINCM42<br>0x404280a4|PA20|1|IO|SDIO<br>(standard)|
||||||SWCLK|2|I||
||||||TIMA0_C2N|5|O||
||||||TIMG0_C1|6|IO||
|17|25|39|39|PA21<br>PINCM46<br>0x404280b4|PA21|1|IO|SDIO<br>(standard)|
||||||SPI0_CS3|3|IO||
||||||UART1_CTS|4|I||
||||||TIMA0_C0|5|IO||
||||||TIMG8_C0|9|IO||
||||||VREF-|(Non-IOMUX 1) 0|A||



Copyright © 2025 Texas Instruments Incorporated 

16 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

**Table 6-2. Pin Attributes (continued)** 

|**RGE**<br>**PIN**|**RHB**<br>**PIN**|**RGZ**<br>**PIN**|**PT**<br>**PIN**|**PIN NAME/**<br>**IOMUX REG/**<br>**IOMUX ADDR**|**SIGNAL**<br>**NAME**|**IOMUX**<br>**PF**|**SIGNAL**<br>**TYPE**|**BUFFER**<br>**TYPE**|
|---|---|---|---|---|---|---|---|---|
|18|26|40|40|PA22<br>PINCM47<br>0x404280b8|PA22|1|IO|SDIO<br>(standard)|
||||||SPI0_CS2|3|IO||
||||||UART1_RTS|4|O||
||||||TIMA0_C0N|5|O||
||||||TIMA0_C1|6|IO||
||||||CLK_OUT|7|O||
||||||I2C0_SCL|8|IOD||
||||||TIMG8_C1|9|IO||
||||||A0_7|(Non-IOMUX 1) 0|A||
|19|27|43|43|PA23<br>PINCM53<br>0x404280d0|PA23|1|IO|SDIO<br>(standard)|
||||||SPI0_CS3|3|IO||
||||||TIMA0_C3|5|IO||
||||||TIMG8_C0|6|IO||
||||||TIMG0_C0|8|IO||
||||||VREF+|(Non-IOMUX 1) 0|A||
|20|28|44|44|PA24<br>PINCM54<br>0x404280d4|PA24|1|IO|SDIO<br>(standard)|
||||||SPI0_CS2|3|IO||
||||||TIMA0_C3N|5|O||
||||||TIMG8_C1|6|IO||
||||||TIMG0_C1|9|IO||
||||||A0_3|(Non-IOMUX 1) 0|A||
|21|29|45|45|PA25<br>PINCM55<br>0x404280d8|PA25|1|IO|SDIO<br>(standard)|
||||||TIMA0_C3|5|IO||
||||||TIMA0_C1N|6|O||
||||||A0_2|(Non-IOMUX 1) 0|A||
|22|30|46|46|PA26<br>PINCM59<br>0x404280e8|PA26|1|IO|SDIO<br>(standard)|
||||||TIMG8_C0|4|IO||
||||||TIMA_FAL0|5|I||
||||||TIMA0_C3N|6|O||
||||||A0_1|(Non-IOMUX 1) 0|A||
||31|47|47|PA27<br>PINCM60<br>0x404280ec|PA27|1|IO|SDIO<br>(standard)|
||||||TIMG8_C1|4|IO||
||||||TIMA_FAL2|5|I||
||||||CLK_OUT|6|O||
||||||RTC_OUT|7|O||
||||||A0_0|(Non-IOMUX 1) 0|A||
|||3|3|PA28<br>PINCM3<br>0x40428008|PA28|1|IO|HDIO (high-<br>drive)|
||||||UART0_TX|2|O||
||||||I2C0_SDA|3|IOD||
||||||TIMA0_C3|4|IO||
||||||TIMA_FAL0|5|I||
||||||TIMA0_C1|6|IO||
||||||SPI0_CS3|7|IO||
||||||WAKE|(Non-IOMUX 1) 0|I||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 17 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

## **MSPM0L1117, MSPM0L1116** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **Table 6-2. Pin Attributes (continued)** 

|**RGE**<br>**PIN**|**RHB**<br>**PIN**|**RGZ**<br>**PIN**|**PT**<br>**PIN**|**PIN NAME/**<br>**IOMUX REG/**<br>**IOMUX ADDR**|**SIGNAL**<br>**NAME**|**IOMUX**<br>**PF**|**SIGNAL**<br>**TYPE**|**BUFFER**<br>**TYPE**|
|---|---|---|---|---|---|---|---|---|
|||5|5|PA31<br>PINCM6<br>0x40428014|PA31|1|IO|SDIO<br>(standard with<br>wake)|
||||||UART0_RX|2|I||
||||||I2C0_SCL|3|IOD||
||||||TIMA0_C3N|4|O||
||||||CLK_OUT|6|O||
||||||SPI0_CS3|7|IO||
||||||WAKE|(Non-IOMUX 1) 0|I||
|||14|14|PB2<br>PINCM15<br>0x40428038|PB2|1|IO|SDIO<br>(standard)|
||||||TIMA0_C3|5|IO||
||||||UART1_CTS|6|I||
||||||TIMG1_C0|7|IO||
||||||HFCLK_IN|10|I||
||||||SPI0_PICO|11|IO||
|||15|15|PB3<br>PINCM16<br>0x4042803c|PB3|1|IO|SDIO<br>(standard)|
||||||TIMA0_C3N|5|O||
||||||UART1_RTS|6|O||
||||||TIMG1_C1|7|IO||
||||||TIMA0_C0|10|IO||
||||||SPI0_SCK|11|IO||
|||20|20|PB6<br>PINCM23<br>0x40428058|PB6|1|IO|SDIO<br>(standard)|
||||||UART1_TX|2|O||
||||||TIMG8_C0|5|IO||
||||||TIMA_FAL2|8|I||
||||||SPI0_CS1|9|IO||
|||21|21|PB7<br>PINCM24<br>0x4042805c|PB7|1|IO|SDIO<br>(standard)|
||||||UART1_RX|2|I||
||||||TIMG8_C1|5|IO||
||||||SPI0_CS2|8|IO||
|||22|22|PB8<br>PINCM25<br>0x40428060|PB8|1|IO|SDIO<br>(standard)|
||||||UART1_CTS|2|I||
||||||TIMA0_C0|5|IO||
|||23|23|PB9<br>PINCM26<br>0x40428064|PB9|1|IO|SDIO<br>(standard)|
||||||UART1_RTS|2|O||
||||||TIMA0_C0N|5|O||
||||||TIMA0_C1|6|IO||
|||24|24|PB14<br>PINCM31<br>0x40428078|PB14|1|IO|SDIO<br>(standard)|
||||||TIMA0_C0|5|IO||
||||||TIMG8_IDX|6|I||
||||||SPI0_CS3|7|IO||
|||25|25|PB15<br>PINCM32<br>0x4042807c|PB15|1|IO|SDIO<br>(standard)|
||||||TIMG8_C0|5|IO||
|||26|26|PB16<br>PINCM33<br>0x40428080|PB16|1|IO|SDIO<br>(standard)|
||||||TIMG8_C1|5|IO||



Copyright © 2025 Texas Instruments Incorporated 

18 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

**Table 6-2. Pin Attributes (continued)** 

|**RGE**<br>**PIN**|**RHB**<br>**PIN**|**RGZ**<br>**PIN**|**PT**<br>**PIN**|**PIN NAME/**<br>**IOMUX REG/**<br>**IOMUX ADDR**|**SIGNAL**<br>**NAME**|**IOMUX**<br>**PF**|**SIGNAL**<br>**TYPE**|**BUFFER**<br>**TYPE**|
|---|---|---|---|---|---|---|---|---|
|||36|36|PB17<br>PINCM43<br>0x404280a8|PB17|1|IO|SDIO<br>(standard)|
||||||SPI0_PICO|3|IO||
||||||I2C0_SCL|4|IOD||
||||||TIMA0_C2|5|IO||
||||||TIMG0_C0|6|IO||
|||37|37|PB18<br>PINCM44<br>0x404280ac|PB18|1|IO|SDIO<br>(standard)|
||||||SPI0_SCK|3|IO||
||||||I2C0_SDA|4|IOD||
||||||TIMA0_C2N|5|O||
||||||TIMG0_C1|6|IO||
|||38|38|PB19<br>PINCM45<br>0x404280b0|PB19|1|IO|SDIO<br>(standard)|
||||||SPI0_POCI|3|IO||
||||||TIMG8_C1|4|IO||
||||||UART0_CTS|5|I||
||||||TIMG8_IDX|7|I||
||||||A0_5|(Non-IOMUX 1) 0|A||
|||41|41|PB20<br>PINCM48<br>0x404280bc|PB20|1|IO|SDIO<br>(standard)|
||||||SPI0_CS2|2|IO||
||||||TIMA0_C2|5|IO||
||||||TIMA_FAL1|6|I||
||||||TIMA0_C1|7|IO||
||||||I2C0_SDA|9|IOD||
||||||A0_6|(Non-IOMUX 1) 0|A||
|||42|42|PB24<br>PINCM52<br>0x404280cc|PB24|1|IO|SDIO<br>(standard)|
||||||SPI0_CS3|2|IO||
||||||SPI0_CS1|3|IO||
||||||TIMA0_C3|5|IO||
||||||TIMA0_C1N|6|O||
||||||UART0_TX|10|O||
||||||UART0_RX|11|I||
|23|32|48|48|VCORE|VCORE|(Non-IOMUX 1) 0|PWR|PWR|
|3|4|6|6|VDD|VDD|(Non-IOMUX 1) 0|PWR|PWR|
|4|5|7|7|VSS|VSS|(Non-IOMUX 1) 0|PWR|PWR|



## **6.3 Signal Descriptions** 

Many MSPM0 signals are made available on multiple device pins. The following list describes the column headers: 

1. **SIGNAL NAME** : The name of the signal which can be connected to one of the specified pins. 

2. **PIN TYPE** : The signal direction and signal type: 

   - I = Input 

   - O = Output 

   - IO = Input, output, or simultaneous input and output 

   - ID = Input with open-drain behavior 

   - OD = Output with open-drain behavior 

   - IOD = Input, output, or simultaneous input and output with open-drain behavior 

   - A = Analog 

   - PWR = Power function 

Copyright © 2025 Texas Instruments Incorporated 

19 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

3. **DESCRIPTION** : A description of the signal. 

4. **PIN** : Associated pin number. 

For additional information on the pin multiplexing scheme, refer to the IOMUX chapter of the _MSPM0 L-Series 32-MHz Microcontrollers Technical Reference Manual_ . 

## **Note** 

The IOMUX only supports connecting one IOMUX-managed digital function to the pin at the same time. However, non-IOMUX managed signals (such as analog inputs and WAKE inputs) can be enabled on a pin at the same time that an IOMUX managed digital function is enabled on the pin. In this case, the designer must verify that no contention exists between the functions enabled on each pin. 

**Table 6-3. Analog to Digital Converter (ADC) Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|A0_0|A|ADC0 analog input channel 0||31|47|47|
|A0_1|A|ADC0 analog input channel 1|22|30|46|46|
|A0_2|A|ADC0 analog input channel 2|21|29|45|45|
|A0_3|A|ADC0 analog input channel 3|20|28|44|44|
|A0_4|A|ADC0 analog input channel 4|14|22|33|33|
|A0_5|A|ADC0 analog input channel 5|||38|38|
|A0_6|A|ADC0 analog input channel 6|||41|41|
|A0_7|A|ADC0 analog input channel 7|18|26|40|40|
|A0_8|A|ADC0 analog input channel 8||16|27|27|
|A0_9|A|ADC0 analog input channel 9||17|28|28|
|A0_12|A|ADC0 analog input channel 12||18|29|29|
|A0_13|A|ADC0 analog input channel 13|12|20|31|31|
|A0_14|A|ADC0 analog input channel 14|13|21|32|32|



**Table 6-4. Bootstrap Loader (BSL) Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|BSLRX|I|BSL UART receive signal (RXD)|10|15|19|19|
|BSLSCL|IOD|BSL I2C clock signal (SCL)|1|2|2|2|
|BSLSDA|IOD|BSL I2C data signal (SDA)|24|1|1|1|
|BSLTX|O|BSL UART transmit signal (TXD)|9|14|18|18|
|BSL_invoke|I|BSL invoke signal (if BSL is enabled, must be<br>HIGH during BOOTRST for a BSL entry, and<br>LOW during BOOTRST to prevent BSL entry)|14|22|33|33|



**Table 6-5. Clock Module (CKM) Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|CLK_OUT|O|CLK_OUT digital clock output from the PMCU|18,8,9|11,13,<br>14,18,<br>26,31|13,17,<br>18,29,<br>40,47,5|13,17,<br>18,29,<br>40,47,5|
|FCC_IN|I|Frequency clock counter (FCC) input signal|12,24|1,11,16,<br>20,9|1,11,13,<br>27,31|1,11,13,<br>27,31|
|HFCLK_IN|I|High frequency clock digital clock input signal||10,12|12,14,<br>16|12,14,<br>16|
|LFCLK_IN|I|Low frequency clock digital clock input signal|7|8|10|10|



20 _Submit Document Feedback_ 

Copyright © 2025 Texas Instruments Incorporated 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

**Table 6-5. Clock Module (CKM) Signal Descriptions (continued)** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|LFXIN|A|Low frequency crystal oscillator (LFXT) signal|6|7|9|9|
|LFXOUT|A|Low frequency crystal oscillator (LFXT) signal|7|8|10|10|
|ROSC|A|SYSOSC frequency correction loop (FCL)<br>external resistor signal|5|6|8|8|



**Table 6-6. General Purpose Input Output Module Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|PA0|IO|GPIO port A input/output 0|24|1|1|1|
|PA1|IO|GPIO port A input/output 1|1|2|2|2|
|PA2|IO|GPIO port A input/output 2|5|6|8|8|
|PA3|IO|GPIO port A input/output 3|6|7|9|9|
|PA4|IO|GPIO port A input/output 4|7|8|10|10|
|PA5|IO|GPIO port A input/output 5||9|11|11|
|PA6|IO|GPIO port A input/output 6||10|12|12|
|PA7|IO|GPIO port A input/output 7||11|13|13|
|PA8|IO|GPIO port A input/output 8||12|16|16|
|PA9|IO|GPIO port A input/output 9|8|13|17|17|
|PA10|IO|GPIO port A input/output 10|9|14|18|18|
|PA11|IO|GPIO port A input/output 11|10|15|19|19|
|PA12|IO|GPIO port A input/output 12||16|27|27|
|PA13|IO|GPIO port A input/output 13||17|28|28|
|PA14|IO|GPIO port A input/output 14||18|29|29|
|PA15|IO|GPIO port A input/output 15|11|19|30|30|
|PA16|IO|GPIO port A input/output 16|12|20|31|31|
|PA17|IO|GPIO port A input/output 17|13|21|32|32|
|PA18|IO|GPIO port A input/output 18|14|22|33|33|
|PA19|IO|GPIO port A input/output 19|15|23|34|34|
|PA20|IO|GPIO port A input/output 20|16|24|35|35|
|PA21|IO|GPIO port A input/output 21|17|25|39|39|
|PA22|IO|GPIO port A input/output 22|18|26|40|40|
|PA23|IO|GPIO port A input/output 23|19|27|43|43|
|PA24|IO|GPIO port A input/output 24|20|28|44|44|
|PA25|IO|GPIO port A input/output 25|21|29|45|45|
|PA26|IO|GPIO port A input/output 26|22|30|46|46|
|PA27|IO|GPIO port A input/output 27||31|47|47|
|PA28|IO|GPIO port A input/output 28|||3|3|
|PA31|IO|GPIO port A input/output 31|||5|5|
|PB2|IO|GPIO port B input/output 2|||14|14|
|PB3|IO|GPIO port B input/output 3|||15|15|
|PB6|IO|GPIO port B input/output 6|||20|20|
|PB7|IO|GPIO port B input/output 7|||21|21|
|PB8|IO|GPIO port B input/output 8|||22|22|
|PB9|IO|GPIO port B input/output 9|||23|23|
|PB14|IO|GPIO port B input/output 14|||24|24|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 21 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

## **MSPM0L1117, MSPM0L1116** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**Table 6-6. General Purpose Input Output Module Signal Descriptions (continued)** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|PB15|IO|GPIO port B input/output 15|||25|25|
|PB16|IO|GPIO port B input/output 16|||26|26|
|PB17|IO|GPIO port B input/output 17|||36|36|
|PB18|IO|GPIO port B input/output 18|||37|37|
|PB19|IO|GPIO port B input/output 19|||38|38|
|PB20|IO|GPIO port B input/output 20|||41|41|
|PB24|IO|GPIO port B input/output 24|||42|42|



**Table 6-7. I2C Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|I2C0_SCL|IOD|I2C0 serial clock signal (SCL)|1,10,18,<br>8,9|13,14,<br>15,2,26|17,18,<br>19,2,36,<br>40,5|17,18,<br>19,2,36,<br>40,5|
|I2C0_SDA|IOD|I2C0 serial data signal (SDA)|10,24,9|1,12,14,<br>15|1,16,18,<br>19,3,37,<br>41|1,16,18,<br>19,3,37,<br>41|
|**Table 6-8. IOMUX Signal Descriptions**|||||||
|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|WAKE|I|Input signal to wake the device from<br>SHUTDOWN mode|1,10,11,<br>12,13,<br>14,2,24,<br>9|1,14,15,<br>19,2,20,<br>21,22,3|1,18,19,<br>2,3,30,<br>31,32,<br>33,4,5|1,18,19,<br>2,3,30,<br>31,32,<br>33,4,5|
|**Table 6-9. Power Management Unit(PMU) Signal Descriptions**|||||||
|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|VCORE|PWR|VCORE capacitor connection|23|32|48|48|
|VDD|PWR|VDD supply|3|4|6|6|
|VSS|PWR|VSS (ground)|4|5|7|7|
|**Table 6-10. Real-time Clock(RTC) Signal Descriptions**|||||||
|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|RTC_OUT|O|Real-time clock output signal|8|13,17,<br>31|17,28,<br>47|17,28,<br>47|
|**Table 6-11. Serial Peripheral Interface(SPI) Signal Descriptions**|||||||
|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|SPI0_PICO|IO|SPI0 peripheral in controller out signal|8|13,18,9|11,14,<br>17,29,<br>36|11,14,<br>17,29,<br>36|
|SPI0_POCI|IO|SPI0 peripheral out controller in signal|7,9|11,14,<br>17,8,9|10,11,<br>13,18,<br>28,38|10,11,<br>13,18,<br>28,38|
|SPI0_SCK|IO|SPI0 serial clock|10|10,15,<br>16|12,15,<br>19,27,<br>37|12,15,<br>19,27,<br>37|



Copyright © 2025 Texas Instruments Incorporated 

22 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

**Table 6-11. Serial Peripheral Interface (SPI) Signal Descriptions (continued)** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|SPI0_CS0|IO|SPI0 chip select 0 signal|14,5,7,<br>8|12,13,<br>22,6,8|10,16,<br>17,33,8|10,16,<br>17,33,8|
|SPI0_CS1|IO|SPI0 chip select 1 signal|13,6|16,21,7|20,27,<br>32,42,9|20,27,<br>32,42,9|
|SPI0_CS2|IO|SPI0 chip select 2 signal|18,20|11,18,<br>26,28|13,21,<br>29,40,<br>41,44|13,21,<br>29,40,<br>41,44|
|SPI0_CS3|IO|SPI0 chip select 3 signal|1,17,19,<br>6|12,17,2,<br>25,27,7|16,2,24,<br>28,3,39,<br>42,43,5,<br>9|16,2,24,<br>28,3,39,<br>42,43,5,<br>9|



**Table 6-12. Serial Wire Debug (SWD) Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|SWCLK|I|Serial wire debug interface clock input signal|16|24|35|35|
|SWDIO|IO|Serial wire debug interface data input/output<br>signal|15|23|34|34|



## **Table 6-13. System Controller (SYSCTL) Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|NRST|I|Active-low reset signal (must be logic high for<br>the device to start)|2|3|4|4|



**Table 6-14. Timer (TIMx) Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|TIMA0_C0|IO|TIMA0 capture/compare 0 signal|12,17,<br>24,5|1,12,20,<br>25,6|1,15,16,<br>22,24,<br>31,39,8|1,15,16,<br>22,24,<br>31,39,8|
|TIMA0_C1|IO|TIMA0 capture/compare 1 signal|1,14,18,<br>6,8|11,13,2,<br>22,26,7|13,17,2,<br>23,3,33,<br>40,41,9|13,17,2,<br>23,3,33,<br>40,41,9|
|TIMA0_C2|IO|TIMA0 capture/compare 2 signal|11,15,6,<br>9|11,14,<br>19,23,7|13,18,<br>30,34,<br>36,41,9|13,18,<br>30,34,<br>36,41,9|
|TIMA0_C3|IO|TIMA0 capture/compare 3 signal|13,19,<br>21,7|16,21,<br>27,29,8|10,14,<br>27,3,32,<br>42,43,<br>45|10,14,<br>27,3,32,<br>42,43,<br>45|
|TIMA0_C0N|O|TIMA0 capture/compare 0 complementary<br>output|18,8|13,26|17,23,<br>40|17,23,<br>40|
|TIMA0_C1N|O|TIMA0 capture/compare 1 complementary<br>output|21,7|29,8|10,42,<br>45|10,42,<br>45|
|TIMA0_C2N|O|TIMA0 capture/compare 2 complementary<br>output|10,12,<br>16,5|10,15,<br>20,24,6|12,19,<br>31,35,<br>37,8|12,19,<br>31,35,<br>37,8|
|TIMA0_C3N|O|TIMA0 capture/compare 3 complementary<br>output|14,20,<br>22,5|17,22,<br>28,30,6|15,28,<br>33,44,<br>46,5,8|15,28,<br>33,44,<br>46,5,8|
|TIMA_FAL0|I|Timer fault input 0|10,22,5|10,12,<br>15,30,6|12,16,<br>19,3,46,<br>8|12,16,<br>19,3,46,<br>8|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 23 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **MSPM0L1117, MSPM0L1116** 

**Table 6-14. Timer (TIMx) Signal Descriptions (continued)** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|TIMA_FAL1|I|Timer fault input 1|24,5,9|1,14,6,<br>9|1,11,18,<br>41,8|1,11,18,<br>41,8|
|TIMA_FAL2|I|Timer fault input 2|1|12,2,31|16,2,20,<br>47|16,2,20,<br>47|
|TIMG8_IDX|I|TIMG8 quadrature encoder index pulse signal|1,11|11,19,2|13,2,24,<br>30,38|13,2,24,<br>30,38|
|TIMG0_C0|IO|TIMG0 capture/compare 0 signal|15,19,<br>24,9|1,14,16,<br>23,27,9|1,11,18,<br>27,34,<br>36,43|1,11,18,<br>27,34,<br>36,43|
|TIMG0_C1|IO|TIMG0 capture/compare 1 signal|1,10,16,<br>20|10,15,<br>17,2,24,<br>28|12,19,2,<br>28,35,<br>37,44|12,19,2,<br>28,35,<br>37,44|
|TIMG1_C0|IO|TIMG1 capture/compare 0 signal|11|19|14,30|14,30|
|TIMG1_C1|IO|TIMG1 capture/compare 1 signal|12|20|15,31|15,31|
|TIMG8_C0|IO|TIMG8 capture/compare 0 signal|1,13,17,<br>19,22,6|11,2,21,<br>25,27,<br>30,7,9|11,13,2,<br>20,25,<br>32,39,<br>43,46,9|11,13,2,<br>20,25,<br>32,39,<br>43,46,9|
|TIMG8_C1|IO|TIMG8 capture/compare 1 signal|14,18,<br>20,24,5,<br>7|1,10,22,<br>26,28,<br>31,6,8|1,10,12,<br>21,26,<br>33,38,<br>40,44,<br>47,8|1,10,12,<br>21,26,<br>33,38,<br>40,44,<br>47,8|



**Table 6-15. Universal Asynchronous Receiver Transmitter (UART) Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|UART0_CTS|I|UART0 clear to send signal|8|13,18,9|11,17,<br>29,38|11,17,<br>29,38|
|UART0_RTS|O|UART0 ready to send signal|11|10,12,<br>19|12,16,<br>30|12,16,<br>30|
|UART0_RX|I|UART0 receive signal (RXD)|1,10,12|15,2,20|19,2,31,<br>42,5|19,2,31,<br>42,5|
|UART0_TX|O|UART0 transmit signal (TXD)|11,24,9|1,14,19|1,18,3,<br>30,42|1,18,3,<br>30,42|
|UART1_CTS|I|UART1 clear to send signal|17|16,25|14,22,<br>27,39|14,22,<br>27,39|
|UART1_RTS|O|UART1 ready to send signal|18|17,26|15,23,<br>28,40|15,23,<br>28,40|
|UART1_RX|I|UART1 receive signal (RXD)|14,7,8|10,13,<br>22,8|10,12,<br>17,21,<br>33|10,12,<br>17,21,<br>33|
|UART1_TX|O|UART1 transmit signal (TXD)|13,6|12,21,7,<br>9|11,16,<br>20,32,9|11,16,<br>20,32,9|



**Table 6-16. Voltage Reference Signal Descriptions** 

|**SIGNAL**<br>**NAME**|**PIN**<br>**TYPE**|**DESCRIPTION**|**RGE PIN**|**RHB PIN**|**RGZ PIN**|**PT PIN**|
|---|---|---|---|---|---|---|
|VREF+|A|Voltage reference positive input|19|27|43|43|
|VREF-|A|Voltage reference negative input|17|25|39|39|



Copyright © 2025 Texas Instruments Incorporated 

24 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **6.4 Connections for Unused Pins** 

Table 6-17 lists the correct termination of unused pins. 

**Table 6-17. Connection of Unused Pins** 

|**PIN**(1)|**POTENTIAL**|**COMMENT**|
|---|---|---|
|PAx, PBx|Open|Set corresponding pin functions to GPIO (PINCMx.PF = 0x1) and<br>configure unused pins to output low or input with internal pullup or<br>pulldown resistor.|
|NRST|VCC|NRST is an active-low reset signal; the pin must be pulled high to<br>VCC or the device cannot start. For more information, seeSection<br>9.1.|



(1) Any unused pin with a function that is shared with general-purpose I/O must follow the "PAx, PBx" unused pin connection guidelines. 

Copyright © 2025 Texas Instruments Incorporated 

25 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **7 Specifications** 

## **7.1 Absolute Maximum Ratings** 

over operating free-air temperature range (unless otherwise noted)[(1)] 

||||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|VDD|Supply voltage|At VDD pin|–0.3<br>4.1|V|
|VI|Input voltage|Applied to any 5-V tolerant open-drain pins|–0.3<br>5.5|V|
|VI|Input voltage|Applied to any common tolerance pins|–0.3<br>VDD+ 0.3<br>(4.1 MAX)|<br>V|
|IVDD (3)|Current into VDD pin<br>(source)|-40 °C ≤ Tj≤ 130 °C|80|mA|
||Current into VDD pin<br>(source)|-40 °C ≤ Tj ≤ 90 °C|100|mA|
|IVSS (3)|Current out of VSS pin<br>(sink)|-40 °C ≤ Tj≤ 130 °C|80|mA|
||Current out of VSS pin<br>(sink)|-40 °C ≤ Tj ≤ 90 °C|100|mA|
|IIO|Current of SDIO pin|Current sunk or sourced by SDIO pin, VDD>=2.7V|6|mA|
||Current of HSIO pin|Current sunk or sourced by HSIO pin, VDD>=2.7V|6|mA|
||Current of HDIO pin|Current sunk or sourced by HDIO pin|20|mA|
||Current of ODIO pin|Current sunk by ODIO pin|20|mA|
|ID|Supported diode current|Diode current at any device pin<br>(excluding Open Drain IO and PB24)|-2<br>2|mA|
|||Diode current at PB24 IO pin|-2<br>0.05|mA|
|TA|Ambient temperature|Ambient temperature|-40<br>125|°C|
|TJ|Junction temperature|Junction temperature|-40<br>130|°C|
|Tstg|Storage temperature(2)|Storage temperature(2)|–40<br>150|°C|



(1) Operation outside the Absolute Maximum Ratings may cause permanent device damage. Absolute Maximum Ratings do not imply functional operation of the device at these or any other conditions beyond those listed under Recommended Operating Conditions. If used outside the Recommended Operating Conditions but within the Absolute Maximum Ratings, the device may not be fully functional, and this may affect device reliability, functionality, performance, and shorten the device lifetime 

(2) Higher temperatures may be applied during board soldering according to the current JEDEC J-STD-020 specification with peak reflow temperatures not higher than classified on the device label on the shipping boxes or reels. 

(3) For applications running at VDD=1.62V, I_VDD/I_VSS<=20mA is required to ensure device functionality 

## **7.2 ESD Ratings** 

||||**VALUE**|**UNIT**|
|---|---|---|---|---|
|V(ESD)<br>|Electrostatic discharge|Human body model (HBM), per ANSI/ESDA/<br>JEDEC JS-001, all pins(1)|±2000|V|
|||Charged device model (CDM), per JEDEC<br>specification JESD22-C101, all pins(2)|±500|V|



(1) JEDEC document JEP155 states that 500-V HBM allows safe manufacturing with a standard ESD control process. 

(2) JEDEC document JEP157 states that 250-V CDM allows safe manufacturing with a standard ESD control process. 

## **7.3 Recommended Operating Conditions** 

over operating free-air temperature range (unless otherwise noted) 

|||**MIN**<br>**NOM**<br>**MAX**|**UNIT**|
|---|---|---|---|
|VDD|Supply voltage|1.62<br>3.6|V|
|VCORE|Voltage on VCORE pin(2)|1.35|V|
|CVDD|Capacitor connected between VDD and VSS(1)|10|uF|
|CVCORE|Capacitor connected between VCORE and VSS(1) (2)|470|nF|



Copyright © 2025 Texas Instruments Incorporated 

26 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

over operating free-air temperature range (unless otherwise noted) 

|||**MIN**<br>**NOM**<br>**MAX**|**UNIT**|
|---|---|---|---|
|TA|Ambient temperature|–40<br>125|°C|
|TJ|Max junction temperature|130|°C|
|fMCLK|MCLK, CPUCLK, ULPCLK frequency with 1 flash wait state(3)|32|MHz|
||MCLK, CPUCLK, ULPCLK frequency with 0 flash wait states(3)|24|MHz|



(1) Connect CVDD and CVCORE between VDD/VSS and VCORE/VSS, respectively, as close to the device pins as possible.  A low-ESR capacitor with at least the specified value and tolerance of ±20% or better is required for CVDD and CVCORE. 

(2) The VCORE pin must only be connected to CVCORE.  Do not supply any voltage or apply any external load to the VCORE pin. 

(3) Wait states are managed automatically by the system controller (SYSCTL) and do not need to be configured by application software unless MCLK is sourced from a high speed clock source (HSCLK sourced from HFCLK). 

## **7.4 Thermal Information** 

|**THERMAL METRIC**(1)|**THERMAL METRIC**(1)|**PACKAGE**|**VALUE**|**UNIT**|
|---|---|---|---|---|
|RθJA|Junction-to-ambient thermal resistance|LQFP-48 (PT)|76.8|°C/W|
|RθJC(top)|Junction-to-case (top) thermal resistance||33.1|°C/W|
|RθJB|Junction-to-board thermal resistance||48.5|°C/W|
|ΨJT|Junction-to-top characterization parameter||2.9|°C/W|
|ΨJB|Junction-to-board characterization parameter||48|°C/W|
|RθJC(bot)|Junction-to-case (bottom) thermal resistance||N/A|°C/W|
|RθJA|Junction-to-ambient thermal resistance|VQFN-48 (RGZ)|32.5|°C/W|
|RθJC(top)|Junction-to-case (top) thermal resistance||23.1|°C/W|
|RθJB|Junction-to-board thermal resistance||14.8|°C/W|
|ΨJT|Junction-to-top characterization parameter||0.6|°C/W|
|ΨJB|Junction-to-board characterization parameter||14.7|°C/W|
|RθJC(bot)|Junction-to-case (bottom) thermal resistance||6.3|°C/W|
|RθJA|Junction-to-ambient thermal resistance|VQFN-32 (RHB)|35.2|°C/W|
|RθJC(top)|Junction-to-case (top) thermal resistance||27.8|°C/W|
|RθJB|Junction-to-board thermal resistance||16.2|°C/W|
|ΨJT|Junction-to-top characterization parameter||0.7|°C/W|
|ΨJB|Junction-to-board characterization parameter||16.1|°C/W|
|RθJC(bot)|Junction-to-case (bottom) thermal resistance||6.3|°C/W|
|RθJA|Junction-to-ambient thermal resistance|VQFN-24 (RGE)|43.6|°C/W|
|RθJC(top)|Junction-to-case (top) thermal resistance||36.8|°C/W|
|RθJB|Junction-to-board thermal resistance||20.9|°C/W|
|ΨJT|Junction-to-top characterization parameter||0.9|°C/W|
|ΨJB|Junction-to-board characterization parameter||20.8|°C/W|
|RθJC(bot)|Junction-to-case (bottom) thermal resistance||6.3|°C/W|



(1) For more information about traditional and new thermal metrics, see the Semiconductor and IC Package Thermal Metrics application report. 

## **7.5 Supply Current Characteristics** 

## _**7.5.1 RUN/SLEEP Modes**_ 

VDD=3.3V.  All inputs tied to 0V or VDD. Outputs do not source or sink any current.  All peripherals are disabled. 

|**PARAMETER**|**MCLK**|**-40°C**|**25°C**|**85°C**|**105°C**|**125°C**|**UNIT**|
|---|---|---|---|---|---|---|---|
|||**TYP MAX**|**TYP MAX**|**TYP MAX**|**TYP MAX**|**TYP MAX**||
|**RUN Mode**||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 27 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **MSPM0L1117, MSPM0L1116** 

VDD=3.3V.  All inputs tied to 0V or VDD. Outputs do not source or sink any current.  All peripherals are disabled. 

||**PARAMETER**|**MCLK**|**-40°C**|**25°C**|**85°C**|**105°C**|**125°C**|**UNIT**|
|---|---|---|---|---|---|---|---|---|
||||**TYP MAX**|**TYP MAX**|**TYP MAX**|**TYP MAX**|**TYP MAX**||
|IDDRUN|MCLK=SYSOSC, CoreMark,<br>execute from flash|32MHz|3.4|3.4|3.4|3.4|3.4|mA|
|||4MHz|0.7|0.7|0.7|0.7|0.7||
|IDDRUN,<br>per MHz|MCLK=SYSOSC, CoreMark,<br>execute from flash|32MHz|105|106|106|107|107|uA/MHz|
|||4MHz|169|170|173|176|184||
||MCLK=SYSOSC, While(1), execute<br>from flash|32MHz|65<br>74|66<br>75|66<br>77|67<br>80|68<br>83||
|**SLEEP Mode**|||||||||
|IDDSLEEP|MCLK=SYSOSC, CPU is halted|32MHz|1594 1730|1614 1780|1624 1800|1633 1860|1656 1920|uA|
|||4MHz|473<br>590|481<br>595|492<br>610|504<br>715|533<br>810||
||MCLK=LFCLK, CPU is halted|32KHz|278<br>340|283<br>345|289<br>380|300<br>450|335<br>556||
|IDDSLEEP,<br>per MHz|MCLK=SYSOSC, CPU is halted|32MHz|50|50|51|51|52|uA/MHz|



## _**7.5.2 STOP/STANDBY Modes**_ 

VDD=3.3V. All inputs tied to 0V or VDD.  Outputs do not source or sink any current.  All peripherals not noted are disabled. 

|**PARAMETER**|**PARAMETER**|**ULPCLK**|**-40°C**|**25°C**|**85°C**|**105°C**|**125°C**|**UNIT**|
|---|---|---|---|---|---|---|---|---|
||||**TYP MAX**|**TYP MAX**|**TYP MAX**|**TYP MAX**|**TYP MAX**||
|**STOP Mode**|||||||||
|IDDSTOP0|SYSOSC=32MHz,<br>USE4MHZSTOP=0,<br>DISABLESTOP=0|4MHz|698<br>760|712<br>780|716<br>785|719<br>790|721<br>795|uA|
|IDDSTOP1|SYSOSC=4MHz,<br>USE4MHZSTOP=1,<br>DISABLESTOP=0||232<br>255|239<br>260|245<br>268|250<br>278|257<br>290||
|IDDSTOP2|SYSOSC off, DISABLESTOP=1,<br>ULPCLK=LFCLK|32kHz|51<br>60|55<br>64|58<br>68|61<br>79|70<br>92||
|**STANDBY Mode**|||||||||
|IDDSTBY0|LFCLK=LFXT, STOPCLKSTBY=0,<br>RTC enabled|32kHz|1.7<br>2|1.8<br>2.2|4<br>8|7<br>17|14<br>31|uA|
|IDDSTBY1|LFCLK=LFOSC, STOPCLKSTBY=1,<br>RTC enabled||1.4<br>1.7|1.5<br>1.8|3<br>7|6<br>17|13<br>31||
||LFCLK=LFXT, STOPCLKSTBY=1,<br>RTC enabled||1.4<br>1.7|1.5<br>1.8|3<br>7|6<br>17|13<br>31||
||LFCLK=LFXT, STOPCLKSTBY=1,<br>GPIOA enabled||1.4<br>1.7|1.5<br>1.8|3<br>7|6<br>17|13<br>31||



## _**7.5.3 SHUTDOWN Mode**_ 

All inputs tied to 0V or VDD. Outputs do not source or sink any current. Core regulator is powered down. 

|**PARAMETER**|**PARAMETER**|**VDD**|**-40°C**|**25°C**|**85°C**|**105°C**|**125°C**|**UNIT**|
|---|---|---|---|---|---|---|---|---|
||||**TYP MAX**|**TYP MAX**|**TYP MAX**|**TYP MAX**|**TYP MAX**||
|IDDSHDN|Supply current in SHUTDOWN mode|3.3V|57|75|464|1069|2961|nA|



Copyright © 2025 Texas Instruments Incorporated 

28 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **7.6 Power Supply Sequencing** 

## _**7.6.1 Power Supply Ramp**_ 

Figure 7-1 shows the relationships of POR-, POR+, BOR0-, and BOR0+ during powerup and powerdown. 

|Figure 7-1shows the relationships of POR-, POR+, BOR0-, and BOR0+ during powerup and powerdown.|Figure 7-1shows the relationships of POR-, POR+, BOR0-, and BOR0+ during powerup and powerdown.|Figure 7-1shows the relationships of POR-, POR+, BOR0-, and BOR0+ during powerup and powerdown.|Figure 7-1shows the relationships of POR-, POR+, BOR0-, and BOR0+ during powerup and powerdown.|Figure 7-1shows the relationships of POR-, POR+, BOR0-, and BOR0+ during powerup and powerdown.|
|---|---|---|---|---|
|POR<br>BOR<br>Running<br>Running<br>BOR<br>POR<br>BOR<br>Running<br>Supply Voltage (VDD)<br>POR-<br>POR+<br>BOR0-<br>BOR0+<br>No reset<br>asserted<br>BOR<br>asserted<br>POR<br>asserted<br>BOR<br>released<br>POR<br>released<br>POR<br>released<br>BOR<br>released<br>Time (t)<br>POR/BOR levels are met<br>for specified |dVDD/dt|<br>BOR<br>released<br>BOR<br>asserted<br>**Figure 7-1. Power Cycle POR and BOR Conditions - VDD**<br>**_7.6.2 POR and BOR_**<br>over operating free-air temperature range (unless otherwise noted)|||||
||**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|dVDD/dt|VDD (supply voltage) slew rate|Rising|0.1|V/us|
|||Falling(1)|0.01||
|||Falling, STANDBY|0.1|V/ms|
|VPOR+|Power-on reset voltage level|Rising|0.95<br>1.30<br>1.59|V|
|VPOR-||Falling|0.9<br>1.25<br>1.54|V|
|VHYS, POR|POR hysteresis||30<br>58<br>74|mV|
|VBOR0+,<br>COLD|Brown-out reset voltage level 0 (default level)|-40 °C ≤ Ta≤ 25 °C<br>Cold start, rising|1.50<br>1.56<br>1.63|V|
|||25 °C ≤ Ta≤ 125 °C<br>Cold start, rising|1.51<br>1.58<br>1.65||
|VBOR0+||Rising(1)|1.56<br>1.59<br>1.62||
|VBOR0-||Falling(1)|1.55<br>1.58<br>1.61||
|VBOR0, STBY||STANDBY mode|1.51<br>1.56<br>1.61||
|VBOR1+|Brown-out-reset voltage level 1|Rising(1)|2.13<br>2.17<br>2.21|V|
|VBOR1-||Falling(1)|2.10<br>2.14<br>2.18||
|VBOR1, STBY||STANDBY mode|2.06<br>2.13<br>2.20||
|VBOR2+|Brown-out-reset voltage level 2|Rising(1)|2.73<br>2.77<br>2.82|V|
|VBOR2-||Falling(1)|2.7<br>2.74<br>2.79||
|VBOR2, STBY||STANDBY mode|2.62<br>2.71<br>2.8||
|VBOR3+|Brown-out-reset voltage level 3|Rising(1)|2.88<br>2.96<br>3.04|V|
|VBOR3-||Falling(1)|2.85<br>2.93<br>3.01||
|VBOR3, STBY||STANDBY mode|2.82<br>2.92<br>3.02||
|VHYS,BOR|Brown-out reset hysteresis|Level 0|15<br>21|mV|
|||Levels 1-3|34<br>40||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 29 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

## **MSPM0L1117, MSPM0L1116** 

## SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|TPD, BOR|BOR propagation delay|RUN/SLEEP/STOP<br>mode|5|us|
|||STANDBY mode|100|us|



(1) Device operating in RUN, SLEEP, or STOP mode. 

## **7.7 Flash Memory Characteristics** 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**Supply**|||||
|VDDPGM/ERASE|Program and erase supply voltage||1.62<br>3.6|V|
|IDDERASE|Supply current from VDD during erase<br>operation|Supply current delta|10|mA|
|IDDPGM|Supply current from VDD during<br>program operation|Supply current delta|10|mA|
|**Endurance**|||||
|NWEC(HI_ENDU<br>RANCE)|Erase/program cycle endurance for<br>chosen 32 sectors of flash(1)||100|k cycles|
|NWEC<br>(NORMAL_ENDU<br>RANCE)|Erase/program cycle endurance (Flash<br>not used for HI_ENDURANCE)(1)||10|k cycles|
|NE(MAX)|Total erase operations before failure(2)||802|k erase operations|
|NW(MAX)|Write operations per word line before<br>sector erase (3)||83|write operations|
|**Retention**|||||
|tRET_85|Flash memory data retention|-40°C <= Tj <= 85°C|60|years|
|tRET_105|Flash memory data retention|-40°C <= Tj <= 105°C|11.4|years|
|**Program and Erase Timing**|||||
|tPROG (WORD, 64)|Program time for flash word(4) (6)||50<br>275|µs|
|tPROG (SEC, 64)|Program time for 1kB sector(5) (6)||6.4|ms|
|tERASE (SEC)|Sector erase time|≤2k erase/program cycles,<br>Tj≥25°C|4<br>20|ms|
|tERASE (SEC)|Sector erase time|≤10k erase/program cycles,<br>Tj≥25°C|20<br>150|ms|
|tERASE (SEC)|Sector erase time|<10k erase/program cycles|20<br>200|ms|
|tERASE (BANK)|Bank erase time|<10k erase/program cycles|22<br>220|ms|



(1) Up to 32 application-chosen sectors from the main flash bank(s) or data bank can be used as high endurance sectors. This enables applications that frequently update flash data such as EEPROM emulation. 

(2) Total number of cumulative erase operations supported by the flash before failure. A sector erase or bank erase operation is considered to be one erase operation. 

(3) Maximum number of write operations allowed per word line before the word line must be erased. If additional writes to the same word line are required, a sector erase is required once the maximum number of write operations per word line is reached. 

(4) Program time is defined as the time from when the program command is triggered until the command completion interrupt flag is set in the flash controller. 

(5) Sector program time is defined as the time from when the first word program command is triggered until the final word program command completes and the interrupt flag is set in the flash controller. This time includes the time needed for software to load each flash word (after the first flash word) into the flash controller during programming of the sector. 

(6) Flash word size is 64 data bits (8 bytes). On devices with ECC, the total flash word size is 72 bits (64 data bits plus 8 ECC bits). 

Copyright © 2025 Texas Instruments Incorporated 

30 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **7.8 Timing Characteristics** 

VDD=3.3V, Ta=25℃ (unless otherwise noted) 

||**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**Wakeup**|**Timing**||||
|tWAKE,<br>SLEEP|Wakeup time from SLEEP0 to RUN(1)||1.2|us|
||Wakeup time from SLEEP1 to RUN(1)||1.5||
||Wakeup time from SLEEP2 to RUN(1)||2.1||
|tWAKE,<br>STOP|Wakeup time from STOP0 to RUN<br>(SYSOSC enabled)(1)||7|us|
||Wakeup time from STOP1 to RUN<br>(SYSOSC enabled)(1)||8.8||
||Wakeup time from STOP2 to RUN<br>(SYSOSC disabled)(1)||8.8||
|tWAKE,<br>STANDBY|Wakeup time from STANDBY0 to RUN<br>(1)||9.9|us|
||Wakeup time from STANDBY1 to RUN<br>(1)||9.9||
|tWAKEUP,<br>SHDN|Wakeup time from SHUTDOWN to<br>RUN(2)|Fast boot enabled|270|us|
|||Fast boot disabled|290||
|**Asynchronous Fast Clock Request Timing**|||||
|tDELAY,<br>SLEEP|Delay time from edge of asynchronous<br>request to first 32MHz MCLK edge|Mode is SLEEP1|0.35|us|
|||Mode is SLEEP2|0.92||
|tDELAY,<br>STOP|Delay time from edge of asynchronous<br>request to first 32MHz MCLK edge|Mode is STOP0|0.1|us|
|||Mode is STOP1|2.2||
|||Mode is STOP2|0.9||
|tDELAY,<br>STANDBY|Delay time from edge of asynchronous<br>request to first 32MHz MCLK edge|Mode is STANDBY0|3.1|us|
|||Mode is STANDBY1|3.1||
|**Startup Timing**|||||
|tSTART,<br>RESET|Device cold startup time from reset/<br>power-up(3)|Fast boot enabled|270|us|
|||Fast boot disabled|310||
|**NRST Timing**|||||
|tRST,<br>BOOTRST|Pulse length on NRST pin to generate<br>BOOTRST|ULPCLK≥4MHz|1.5|us|
|||ULPCLK=32kHz|30||
|tRST, POR|Pulse length on NRST pin to generate<br>POR||1|s|



(1) The wake-up time is measured from the edge of an external wake-up signal (GPIO wake-up event) to the time that the first instruction of the user program is executed, with glitch filter disabled (FILTEREN=0x0) and fast wake enabled (FASTWAKEONLY=1) . 

(2) The wake-up time is measured from the edge of an external wake-up signal (IOMUX wake-up event) to the time that first instruction of the user program is executed. 

(3) The start-up time is measured from the time that VDD crosses VBOR0- (cold start-up) to the time that the first instruction of the user program is executed. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 31 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **7.9 Clock Specifications** 

## _**7.9.1 System Oscillator (SYSOSC)**_ 

over operating free-air temperature range (unless otherwise noted) 

||**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|fSYSOSC|Factory trimmed SYSOSC frequency|SYSOSCCFG.FREQ=00 (BASE)|32|MHz|
|||SYSOSCCFG.FREQ=01|4||
||User trimmed SYSOSC frequency|SYSOSCCFG.FREQ=10,<br>SYSOSCTRIMUSER.FREQ=10|24||
|||SYSOSCCFG.FREQ=10,<br>SYSOSCTRIMUSER.FREQ=01|16||
|fSYSOSC|SYSOSC frequency accuracy when<br>frequency correction loop (FCL) is<br>enabled and an ideal ROSCresistor is<br>assumed(1) (2)|SETUSEFCL=1, Ta= 25℃|-0.60<br>0.68|%|
|||SETUSEFCL=1, -40℃≤ Ta≤ 85℃|-0.80<br>0.93||
|||SETUSEFCL=1, -40℃≤ Ta≤ 105℃|-0.80<br>1.1||
|||SETUSEFCL=1, -40℃≤ Ta≤ 125℃|-0.80<br>1.3||
|fSYSOSC|SYSOSC accuracy when frequency<br>correction loop (FCL) is enabled with<br>ROSCresistor put at ROSCpin, for factory<br>trimmed frequencies(1)|SETUSEFCL=1, Ta= 25℃, ±0.1%<br>±25ppm ROSC|-0.7<br>0.78|%|
|||SETUSEFCL=1, -40℃≤ Ta≤ 85<br>℃, ±0.1% ±25ppm ROSC|-1.1<br>1.2||
|||SETUSEFCL=1, -40℃≤ Ta≤ 105<br>℃, ±0.1% ±25ppm ROSC|-1.1<br>1.4||
|||SETUSEFCL=1, -40℃≤ Ta≤ 125℃,<br>±0.1% ±25ppm ROSC|-1.1<br>1.7||
|fSYSOSC|SYSOSC frequency accuracy when<br>frequency correction loop (FCL) is<br>enabled when the internal ROSCresistor<br>is used, 32MHz(4)|SETUSEFCL=1, Ta= 25℃|-1.2<br>1.3|%|
|||SETUSEFCL=1, -40℃≤ Ta≤ 125℃|-2.1<br>1.6||
|fSYSOSC|SYSOSC frequency accuracy when<br>frequency correction loop (FCL) is<br>enabled when the internal ROSCresistor<br>is used, 4MHz(4)|SETUSEFCL=1, Ta= 25℃|-1.2<br>1.7|%|
|||SETUSEFCL=1, -40℃≤ Ta≤ 125℃|-2.3<br>1.8||
|fSYSOSC|SYSOSC accuracy when frequency<br>correction loop (FCL) is disabled, 32MHz|SETUSEFCL=0,<br>SYSOSCCFG.FREQ=00, -40℃≤ Ta≤<br>125℃|-2.6<br>1.8|%|
|fSYSOSC|SYSOSC accuracy when frequency<br>correction loop (FCL) is disabled, for<br>factory trimmed frequencies, 4MHz|SETUSEFCL=0,<br>SYSOSCCFG.FREQ=01, -40℃≤ Ta≤<br>125℃|-2.8<br>2.1|%|
|ROSC|External resistor between ROSC pin and<br>VSS(1)|SETUSEFCL=1|100|kΩ|
|tsettle,<br>SYSOSC|Settling time to target accuracy(3)|SETUSEFCL=1, ±0.1% 25ppm ROSC (1)|30|us|
|fsettle,<br>SYSOSC|fSYSOSCadditional undershoot accuracy<br>during tsettle(3)|SETUSEFCL=1, ±0.1% 25ppm ROSC (1)|-16|%|



(1) The SYSOSC frequency correction loop (FCL) enables high SYSOSC accuracy via an external reference resistor (ROSC) which must be connected between the device ROSC pin and VSS when using the FCL. Accuracies are shown for a ±0.1% ±25ppm ROSC; relaxed tolerance resistors may also be used (with reduced SYSOSC accuracy). See the SYSOSC section of the technical reference manual for details on computing SYSOSC accuracy for various ROSC accuracies. ROSC does not need to be populated if the FCL is not enabled. 

(2) Represents the device accuracy only. The tolerance and temperature drift of the ROSC resistor used must be combined with this spec to determine final accuracy. Performance for a ±0.1% ±25ppm ROSC is given as a reference point. 

(3) When SYSOSC is waking up (for example, when exiting a low power mode) and FCL is enabled, the SYSOSC will initially undershoot the target frequency fSYSOSC by an additional error of up to fsettle,SYSOSC for the time tsettle,SYSOSC, after which the target accuracy is achieved. 

(4) The SYSOSC frequency correction loop (FCL) enables high SYSOSC accuracy via an internal reference resistor when using the FCL. See the SYSOSC section of the technical reference manual for details on computing SYSOSC accuracy. 

Copyright © 2025 Texas Instruments Incorporated 

32 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## _**7.9.2 Low Frequency Oscillator (LFOSC)**_ 

over operating free-air temperature range (unless otherwise noted) 

||**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|fLFOSC|LFOSC frequency||32768|Hz|
||LFOSC accuracy|-40℃≤ Ta≤ 125℃|–5<br>5|%|
|||-40℃≤ Ta≤ 85℃|–3<br>3|%|
|ILFOSC|LFOSC current consumption||300|nA|
|tstart,<br>LFOSC|LFOSC start-up time||1|ms|



## _**7.9.3 Low Frequency Crystal/Clock**_ 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**Low frequency crystal oscillator (LFXT)**|||||
|fLFXT|LFXT frequency||32768|Hz|
|DCLFXT|LFXT duty cycle||30<br>70|%|
|OALFXT|LFXT crystal oscillation allowance||419|kΩ|
|CL, eff|Integrated effective load capacitance(1)||1|pF|
|tstart, LFXT|LFXT start-up time||1000|ms|
|ILFXT|LFXT current consumption|XT1DRIVE=0, LOWCAP=1|200|nA|
|**Low frequency digital clock input (LFCLK_IN)**|||||
|fLFIN|LFCLK_IN frequency(2)|SETUSEEXLF=1|29491<br>32768<br>36045|Hz|
|DCLFIN|LFCLK_IN duty cycle(2)|SETUSEEXLF=1|40<br>60|%|
|**LFCLK Monitor**|||||
|fFAULTLF|LFCLK monitor fault frequency(3)|MONITOR=1|2800<br>4200<br>8400|Hz|



(1) This includes parasitic bond and package capacitance (≈2pF per pin), calculated as CLFXIN×CLFXOUT/(CLFXIN+CLFXOUT), where CLFXIN and CLFXOUT are the total capacitance at LFXIN and LFXOUT, respectively. 

(2) The digital clock input (LFCLK_IN) accepts a logic level square wave clock. 

(3) The LFCLK monitor may be used to monitor the LFXT or LFCLK_IN.  It will always fault below the MIN fault frequency, and will never fault above the MAX fault frequency. 

## **7.10 Digital IO** 

## _**7.10.1  Electrical Characteristics**_ 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|VIH<br>|High level input voltage|ODIO(1)|VDD≥1.62V|0.7*VDD<br>5.5|V|
||||VDD≥2.7V|2<br>5.5|V|
|||All I/O except<br>ODIO & Reset|VDD≥1.62V|0.7*VDD<br>VDD+0.3|V|
|VIL<br>|Low level input voltage|ODIO|VDD≥1.62V|-0.3<br>0.3*VDD|V|
||||VDD≥2.7V|-0.3<br>0.8|V|
|||All I/O except<br>ODIO & Reset|VDD≥1.62V|-0.3<br>0.3*VDD|V|
|VHYS<br>|Hysteresis|ODIO||0.05*VDD|V|
|||All I/O except<br>ODIO||0.1*VDD|V|
|Ilkg<br> <br>|High-Z leakage current (All<br>packages)|SDIO (except<br>PB24)(2) (3)|1.62V ≤ VDD ≤ 3.6V, -40℃≤ Ta≤<br>125℃|50(4)|nA|
|||PB24(2) (3)|1.62V ≤ VDD ≤ 3.6V, -40℃≤ Ta≤<br>125℃|130(4)|nA|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 33 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

## **MSPM0L1117, MSPM0L1116** 

## SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|RPU|Pull up resistance|All I/O except<br>ODIO|VIN = VSS|40|kΩ|
|RPD|Pull down resistance||VIN = VDD|40|kΩ|
|CI|Input capacitance||VDD = 3.3V|5|pF|
|VOH|High level output voltage|SDIO|VDD≥2.7V, |IIO|,max=6mA<br>VDD≥1.71V, |IIO|,max=2mA<br>VDD≥1.62V, |IIO|,max=1.5mA<br>-40 °C ≤ Ta≤ 25 °C|VDD-0.4|V|
||||VDD≥2.7V, |IIO|,max=6mA<br>VDD≥1.71V, |IIO|,max=2mA<br>VDD≥1.62V, |IIO|,max=1.5mA<br>-40 °C ≤ Ta≤ 125 °C|VDD-0.45||
|||HSIO|VDD≥2.7V, DRV=1, |IIO|,max=6mA<br>VDD≥1.71V, DRV=1, |IIO|,max=3mA<br>VDD≥1.62V, DRV=1, |IIO|,max=2mA<br>-40 °C ≤ Ta≤ 25 °C|VDD-0.4||
||||VDD≥2.7V, DRV=1, |IIO|,max=6mA<br>VDD≥1.71V, DRV=1, |IIO|,max=3mA<br>VDD≥1.62V, DRV=1, |IIO|,max=2mA<br>-40 °C ≤ Ta≤ 125 °C|VDD-0.45||
||||VDD≥2.7V, DRV=0, |IIO|,max=4mA<br>VDD≥1.71V, DRV=0, |IIO|,max=2mA<br>VDD≥1.62V, DRV=0, |IIO|,max=1.5mA<br>-40 °C ≤ Ta≤ 25 °C|VDD-0.4||
||||VDD≥2.7V, DRV=0, |IIO|,max=4mA<br>VDD≥1.71V, DRV=0, |IIO|,max=2mA<br>VDD≥1.62V, DRV=0, |IIO|,max=1.5mA<br>-40 °C ≤ Ta≤ 125 °C|VDD-0.45||
|||HDIO|VDD≥2.7V, DRV=1(5), |IIO|,max=20mA<br>VDD≥1.71V, DRV=1(5), |<br>IIO|,max=10mA|VDD-0.4||
||||VDD≥2.7V, DRV=0, |IIO|,max=6mA<br>VDD≥1.71V, DRV=0, |IIO|,max=2mA|VDD-0.4||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

34 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|VOL|Low level output voltage|SDIO|VDD≥2.7V, |IIO|,max=6mA<br>VDD≥1.71V, |IIO|,max=2mA<br>VDD≥1.62V, |IIO|,max=1.5mA<br>-40 °C ≤ Ta≤ 25 °C|0.4|V|
||||VDD≥2.7V, |IIO|,max=6mA<br>VDD≥1.71V, |IIO|,max=2mA<br>VDD≥1.62V, |IIO|,max=1.5mA<br>-40 °C ≤ Ta≤ 125 °C|0.45||
|||HSIO|VDD≥2.7V, DRV=1, |IIO|,max=6mA<br>VDD≥1.71V, DRV=1, |IIO|,max=3mA<br>VDD≥1.62V, DRV=1, |IIO|,max=2mA<br>-40 °C ≤ Ta≤ 25 °C|0.4||
||||VDD≥2.7V, DRV=1, |IIO|,max=6mA<br>VDD≥1.71V, DRV=1, |IIO|,max=3mA<br>VDD≥1.62V, DRV=1, |IIO|,max=2mA<br>-40 °C ≤ Ta≤ 125 °C|0.45||
||||VDD≥2.7V, DRV=0, |IIO|,max=4mA<br>VDD≥1.71V, DRV=0, |IIO|,max=2mA<br>VDD≥1.62V, DRV=0, |IIO|,max=1.5mA<br>-40 °C ≤ Ta≤ 25 °C|0.4||
||||VDD≥2.7V, DRV=0, |IIO|,max=4mA<br>VDD≥1.71V, DRV=0, |IIO|,max=2mA<br>VDD≥1.62V, DRV=0, |IIO|,max=1.5mA<br>-40 °C ≤ Ta≤ 125 °C|0.45||
|||HDIO|VDD≥2.7V, DRV=1(5), |IIO|,max=20mA<br>VDD≥1.71V, DRV=1(5), |<br>IIO|,max=10mA|0.4||
||||VDD≥2.7V, DRV=0, |IIO|,max=6mA<br>VDD≥1.71V, DRV=0, |IIO|,max=2mA|0.4||
|||ODIO|VDD≥2.7V, IOL,max=8mA<br>VDD≥1.71V, IOL,max=4mA<br>-40 °C ≤ Ta≤ 25 °C|0.4||
||||VDD≥2.7V, IOL,max=8mA<br>VDD≥1.71V, IOL,max=4mA<br>-40 °C ≤ Ta≤ 125 °C|0.45||



(1) I/O Types: ODIO = 5V Tolerant Open-Drain , SDIO = Standard-Drive , HSIO = High-Speed 

(2) The leakage current is measured with VSS or VDD applied to the corresponding pin(s), unless otherwise noted. 

(3) The leakage of the digital port pins is measured individually. The port pin is selected for input and the pullup/pulldown resistor is disabled. 

(4) This value is for SDIO not muxed with any analog inputs. If the SDIO is muxed with analog inputs then the leakage can be higher. 

(5) When operating a HDIO in DRV=1 high drive strength configuration, a series resistor is necessary to limit the signal slew rate 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 35 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## _**7.10.2 Switching Characteristics**_ 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**||**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|fmax|Port output frequency(1)|SDIO|VDD ≥ 2.7V, CL= 20pF|32|MHz|
||||VDD ≥ 1.71V, CL= 20pF|16|MHz|
|||HSIO|VDD ≥ 2.7V, DRV = 1, CL= 20pF|32|MHz|
||||VDD ≥ 2.7V, DRV = 0, CL= 20pF|32|MHz|
||||VDD ≥ 1.71V, DRV = 1, CL= 20pF|24|MHz|
||||VDD ≥ 1.71V, DRV = 0, CL= 20pF|16|MHz|
|||HDIO|VDD ≥ 2.7V, DRV = 1(2), CL= 20pF|20|MHz|
||||VDD ≥ 2.7V, DRV = 0, CL= 20pF|20|MHz|
||||VDD ≥ 1.71V, DRV = 1(2), CL= 20pF|16|MHz|
||||VDD ≥ 1.71V, DRV = 0, CL= 20pF|16|MHz|
|||ODIO|VDD ≥ 1.71V, FM+, CL= 20pF - 100pF|1|MHz|
|tr,tf|Output rise/fall time|SDIO|VDD ≥ 2.7V, CL= 20pF|3.5|ns|
||||VDD ≥ 1.71V, CL= 20pF|6.6|ns|
|||HSIO|VDD ≥ 2.7V, DRV = 1, CL= 20pF|1.8|ns|
||||VDD ≥ 2.7V, DRV = 0, CL= 20pF|5.9|ns|
||||VDD ≥ 1.71V, DRV = 1, CL= 20pF|3.7|ns|
||||VDD ≥ 1.71V, DRV = 0, CL= 20pF|12.6|ns|
|||HDIO|VDD ≥ 2.7V, DRV = 1, CL= 20pF|1.7|ns|
||||VDD ≥ 2.7V, DRV = 0, CL= 20pF|3.8|ns|
||||VDD ≥ 1.71V, DRV = 1, CL= 20pF|3.1|ns|
||||VDD ≥ 1.71V, DRV = 0, CL= 20pF|8.2|ns|
|tf|Output fall time|ODIO|VDD ≥ 1.71V, FM+, CL= 20pF-100pF|20*VDD/5.5<br>120|ns|



(1) I/O Types: ODIO = 5V Tolerant Open-Drain , SDIO = Standard-Drive , HSIO = High-Speed , HDIO = High-Drive 

(2) When operating a HDIO in DRV=1 high drive strength configuration, a series resistor is necessary to limit the signal slew rate 

## **7.11 Analog Mux VBOOST** 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|IVBST|VBOOST current adder|MCLK/ULPCLK is<br>LFCLK|0.8|uA|
|||MCLK/ULPCLK is<br>not LFCLK, SYSOSC<br>frequency is 4MHz|10.6||
|tSTART,VBST|VBOOST startup time||12<br>20|us|



## **7.12 ADC** 

## _**7.12.1 Electrical Characteristics**_ 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted), all TYP values are measured at 25℃ and all accuracy parameters are measured using 12-bit resolution mode (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|Vin(ADC)|Analog input voltage range(1)|Applies to all ADC analog input pins|0<br>VDD|V|
|VR+|Positive ADC reference voltage|VR+sourced from VDD|VDD|V|
|||VR+sourced from external reference pin (VREF+)|1.4<br>VDD|V|
|||VR+sourced from internal reference (VREF)|VREF|V|
|VR-|Negative ADC reference voltage||0|V|



Copyright © 2025 Texas Instruments Incorporated 

36 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted), all TYP values are measured at 25℃ and all accuracy parameters are measured using 12-bit resolution mode (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|FS|ADC sampling frequency|RES = 0x0 (12-bit mode), External Reference|1.68|Msps|
|FS|ADC sampling frequency|External reference(3), HW Averaging Enabled, 16 Samples<br>and 2bit shift|105|Ksps|
|FS|ADC sampling frequency|RES = 0x0 (12-bit mode), Internal Reference|200|ksps|
|I(ADC) (2)|Operating supply current<br>into VDD terminal|FS= 1.68MSPS, Internal reference OFF, VR+= VDD|570|μA|
|||FS= 200ksps, Internal reference ON, VR+= VREF = 2.5V|320||
|CS/H|ADC sample-and-hold capacitance||4.3|pF|
|Rin|ADC input resistance||0.5|kΩ|
|ENOB|Effective number of bits|Fin = 10kHz, External reference(3)|11.0<br>11.1|bit|
|||Fin = 5kHz, External reference(3), HW Averaging Enabled, 16<br>Samples and 2bit shift|12.3||
|||Fin = 10kHz, Internal reference, VR+= VREF = 2.5V|10<br>10.2||
|SNR|Signal-to-noise ratio|Fin = 10kHz, External reference(3)|71|dB|
|||Fin = 5kHz, External reference(3), HW Averaging Enabled, 16<br>Samples and 2bit shift|76||
|||Fin = 10kHz, Internal reference, VR+= VREF = 2.5V|65||
|PSRRDC|Power supply rejection ratio, DC|External reference(3), VDD = VDD(min)to VDD(max)|68|dB|
|||VDD = VDD(min)to VDD(max)<br>Internal reference, VR+= VREF = 2.5V|60||
|PSRRAC|Power supply rejection ratio, AC|External reference(3), ΔVDD = 0.1 V at 1 kHz|61|dB|
|||ΔVDD = 0.1 V at 1 kHz<br>Internal reference, VR+= VREF = 2.5V|55||
|Twakeup|ADC Wakeup Time|Assumes internal reference is active|5|us|
|VSupplyMon|Supply Monitor voltage divider<br>(VDD/3) accuracy|ADC input channel: Supply Monitor (4)|-1.5<br>+1.5|%|
|ISupplyMon|Supply Monitor voltage divider current<br>consumption|ADC input channel: Supply Monitor|10|uA|



(1) The analog input voltage range must be within the selected ADC reference voltage range VR+ to VR– for valid conversion results. 

(2) The internal reference (VREF) supply current is not included in current consumption parameter I(ADC). 

(3) All external reference specifications are measured with VR+ = VREF+ = VDD = 3.3V and VR- = VREF- = VSS = 0V and external 1uF cap on VREF+ pin. 

(4) Analog power supply monitor. Analog input on channel 31 for VDD monitor is disconnected and is internally connected to the voltage divider which is VDD/3. 

## _**7.12.2 Switching Characteristics**_ 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|fADCCLK|ADC clock frequency||4<br>32|MHz|
|tADC trigger|Software trigger minimum width||3|ADCCLK<br>cycles|
|tSample|Sampling time|12-bit mode, RS= 50Ω, Cpext= 10pF|156|ns|
|tSample_VREF|Sample time with VREF|ADC CHANNEL=28,12-bit mode,VDD as<br>reference|4|µs|
|tSample_SupplyMon(<br>VDD)|Sample time with Supply Monitor (VDD/3)<br>(1)||5|µs|



(1) Analog power supply monitor. Analog input on channel 31 for VDD monitor is disconnected and is internally connected to the voltage divider which is VDD/3. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 37 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## _**7.12.3 Linearity Parameters**_ 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted), all TYP values are measured at 25℃ and all linearity parameters are measured using 12-bit resolution mode (unless otherwise noted)[(1)] 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|EI|Integral linearity error (INL)|External reference(2)|-2<br>2|LSB|
|ED|Differential linearity error (DNL)<br>Guaranteed no missing codes|External reference(2)|-1<br>1|LSB|
|EO|Offset error|External reference(2)|-3.5<br>3.5|mV|
|EG|Gain error|External reference(2)|-4<br>4|LSB|



- (1) Total Unadjusted Error (TUE) can be calculated from EI , EO , and EG using the following formula: TUE = √( EI[2 ] + |EO|[2 ] + EG[2 ] ) Note: You must convert all of the errors into the same unit, usually LSB, for the above equation to be accurate 

- (2) All external reference specifications are measured with VR+ = VREF+ = VDD and VR- = VSS = 0V, external 1uF cap on VREF+ Pin 

## _**7.12.4 Typical Connection Diagram**_ 

**==> picture [259 x 115] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>Boundary<br>ADC Model<br>Rpar S1 Rin 12-bit SAR<br>Vin Converter<br>CS/H<br>Cpar CI<br>**----- End of picture text -----**<br>


**Figure 7-2. ADC Input Network** 

1. Refer to Electrical Characteristics for the values of Rin and CS/H 

2. Refer to Electrical Characteristics for the value of CI 

3. Cpar and Rpar represent the parasitic capacitance and resistance of the external ADC input circuitry 

Use the following equations to solve for the minimum sampling time (T) required for an ADC conversion: 

1. Tau = (Rpar + Rin) × CS/H + Rpar × (Cpar + CI) 

2. K= ln(2[n] /Settling error) – ln((Cpar + CI)/CS/H) 

3. T (Min sampling time) = K × Tau 

## **7.13 Temperature Sensor** 

over operating free-air temperature range (unless otherwise noted)[(1)] 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|TSTRIM|Factory trim temperature(2)|ADC and VREF configuration: RES=0<br>(12-bit mode), VRSEL=2h (VREF=1.4V),<br>ADC tSample=12.5µs|27<br>30<br>33|℃|
|TSc|Temperature coefficient|-40℃≤ Tj≤ 130℃|-2.05<br>-1.9<br>-1.75|mV/℃|
|tSET, TS|Temperature sensor settling time(3)|ADC and VREF configuration: RES=0<br>(12-bit mode), VRSEL= 2h (VREF=1.4V),<br>ADC CHANNEL=11|12.5|us|



(1) Effective absolute temperature accuracy may be computed by combining the relative temperature accuracy together with the trim accuracy, and accounting for any analog to digital conversion error. 

(2) Higher absolute accuracy may be achieved through user calibration. Please refer to temperature sensor chapter in detailed description section. 

- (3) This is the minimum required ADC sampling time when measuring the temperature sensor. 

Copyright © 2025 Texas Instruments Incorporated 

38 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **7.14 VREF** 

## _**7.14.1 Voltage Characteristics**_ 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|VDDmin|Minimum supply voltage needed for<br>VREF operation|BUFCONFIG = 1|1.62|V|
|||BUFCONFIG = 0|2.7||
|VREF|Voltage reference output voltage|BUFCONFIG = 1|1.38<br>1.4<br>1.42|V|
|||BUFCONFIG = 0|2.46<br>2.5<br>2.54||



## _**7.14.2 Electrical Characteristics**_ 

over recommended ranges of supply voltage and operating free-air temperature (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|IVREF|VREF operating supply current|BUFCONFIG = {0, 1}, No load|80<br>100|µA|
|TCVREF|Temperature coefficient of VREF(1)|BUFCONFIG = {0, 1}|80|ppm/°C|
|TCdrift|Long term VREF drift|Time = 1000 hours, BUFCONFIG = {0, 1}, T = 25℃|300|ppm|
|PSRRDC|VREF Power supply rejection ratio,<br>DC|VDD = 1.7 V to VDDmax, BUFCONFIG = 1|70|dB|
|||VDD = 2.7 V to VDDmax, BUFCONFIG = 0|60||
|Tstartup|VREF startup time|BUFCONFIG = {0, 1} , VDD ≥ 2.7 V|15|us|



(1) The temperature coefficient of the VREF output is the sum of TCVRBUF and the temperature coefficient of the internal bandgap reference. 

## **7.15 I2C** 

## _**7.15.1 I2C Characteristics**_ 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETERS**|**PARAMETERS**|**TEST CONDITIONS**|**Standard mode**|**Fast mode**|**Fast mode plus**|**UNIT**|
|---|---|---|---|---|---|---|
||||**MIN**<br>**MAX**|**MIN**<br>**MAX**|**MIN**<br>**MAX**||
|fI2C|I2C input clock frequency||2<br>32|8<br>32|20<br>32|MHz|
|fSCL|SCL clock frequency||0.025<br>0.1|0.4|1|MHz|
|tHD,STA|Hold time (repeated) START||4|0.6|0.26|us|
|tLOW|LOW period of the SCL clock||4.7|1.3|0.5|us|
|tHIGH|High period of the SCL clock||4|0.6|0.26|us|
|tSU,STA|Setup time for a repeated<br>START||4.7|0.6|0.26|us|
|tHD,DAT|Data hold time||0|0|0|ns|
|tSU,DAT|Data setup time||250|100|50|ns|
|tSU,STO|Setup time for STOP||4|0.6|0.26|us|
|tBUF|bus free time between a STOP<br>and START condition||4.7|1.3|0.5|us|
|tVD;DAT|data valid time||3.45|0.9|0.45|us|
|tVD;ACK|data valid acknowledge time||3.45|0.9|0.45|us|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 39 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## _**7.15.2 I2C Filter**_ 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETERS**|**PARAMETERS**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|fSP|Pulse duration of spikes suppressed by<br>input filter|AGFSELx = 0|6|ns|
|||AGFSELx = 1|14<br>35||
|||AGFSELx = 2|22<br>60||
|||AGFSELx = 3|35<br>90||



## _**7.15.3 I[2] C Timing Diagram**_ 

**==> picture [476 x 119] intentionally omitted <==**

**----- Start of picture text -----**<br>
tHD,STA tSU,STA tHD,STA tBUF<br>SDA<br>tt LOW t tt HIGH t tSP<br>SCL<br>tHD,DAT<br>tVD,DAT tSU,DAT tSU,STO<br>**----- End of picture text -----**<br>


**Figure 7-3. I2C Timing Diagram** 

## **7.16 SPI** 

## _**7.16.1 SPI**_ 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETERS**|**PARAMETERS**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**SPI**|||||
|fSPI|SPI clock frequency|Clock max speed >= 32MHz<br>1.62 < VDD < 3.6V<br>Peripheral or Controller mode|16(4)|MHz|
|DCSCK|SCK Duty Cycle||40<br>50<br>60|%|
|**Controller**|||||
|tSCLK_H/L|SCLK High or Low time||(tSPI/2) -<br>1<br>tSPI / 2 (tSPI/2) +<br>1|<br>ns|
|tCS.LEAD|CS lead-time, CS active to clock|SPH=0|1 SPI<br>Clock|ns|
|||SPH=1|1/2 SPI<br>Clock||
|tCS.LAG|CS lag time, Last clock to CS<br>inactive|SPH=0|1/2 SPI<br>Clock|ns|
|||SPH=1|1 SPI<br>Clock||
|tCS.ACC|CS access time, CS active to PICO<br>data out||1/2 SPI<br>Clock|<br>ns|
|tCS.DIS|CS disable time, CS inactive to<br>PICO high impedance||1 SPI<br>Clock|<br>ns|
|tSU.CI|POCI input data setup time(1)|2.7 < VDD < 3.6V, delayed sampling<br>enabled|1|ns|
|||1.62 < VDD < 2.7V, delayed sampling<br>enabled|2.5||
|||2.7 < VDD < 3.6V, no delayed sampling|27||
|||1.62 < VDD < 2.7V, no delayed sampling|34||



Copyright © 2025 Texas Instruments Incorporated 

40 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETERS**|**PARAMETERS**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|tHD.CI|POCI input data hold time|delayed sampling enabled<br>no delayed sampling|26|ns|
||||0||
|tVALID.CO|PICO output data valid time(2)||11.5|ns|
|tHD.CO|PICO output data hold time(3)||0|ns|
|**Peripheral**|||||
|tCS.LEAD|CS lead-time, CS active to clock||13|ns|
|tCS.LAG|CS lag time, Last clock to CS<br>inactive||1|ns|
|tCS.ACC|CS access time, CS active to POCI<br>data out||34.5|ns|
|tCS.DIS|CS disable time, CS inactive to<br>POCI high impedance||34.5|ns|
|tSU.PI|PICO input data setup time||16|ns|
|tHD.PI|PICO input data hold time||3|ns|
|tVALID.PO|POCI output data valid time(2)|2.7 < VDD < 3.6V<br>1.62 < VDD < 2.7V|26|ns|
||||32.5||
|tHD.PO|POCI output data hold time(3)||5|ns|



(1) The POCI input data setup time can be fully compensated when delayed sampling feature is enabled. 

(2) Specifies the time to drive the next valid data to the output after the output changing SCLK clock edge 

(3) Specifies how long data on the output is valid after the output changing SCLK clock edge 

(4) fSPIclk = 1/2tLO/HI with tLO/HI = max(tVALID,CO + tSU,PI, tSU,CI + tVALID,PO). 

## _**7.16.2 SPI Timing Diagram**_ 

**==> picture [500 x 203] intentionally omitted <==**

**----- Start of picture text -----**<br>
CS  CS<br>(inverted) (inverted)<br>tCS, LEAD tCS, LEAD<br>CS CS<br>1 / fSPI tCS, LAG 1 / fSPI tCS, LAG<br>SCLK  SCLK<br>(SPO = 0) (SPO = 0)<br>tSCLK_H/L tSCLK_H/L tSCLK_H/L tSCLK_H/L<br>SCLK  SCLK<br>(SPO = 1) (SPO = 1)<br>tSU,CI tSU,CI<br>tHD,CI tCS, ACC tHD,CI<br>POCI POCI<br>tCS, ACC ttHD,COVALID,CO tCS, DIS ttHD,COVALID,CO tCS, DIS<br>PICO PICO<br>Controller Mode, SPH = 0   Controller Mode, SPH = 1<br>**----- End of picture text -----**<br>


**Figure 7-4. SPI Timing Diagram - Controller Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 41 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

## **MSPM0L1117, MSPM0L1116** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**==> picture [46 x 7] intentionally omitted <==**

**----- Start of picture text -----**<br>
www.ti.com<br>**----- End of picture text -----**<br>


**==> picture [500 x 204] intentionally omitted <==**

**----- Start of picture text -----**<br>
CS  CS<br>(inverted) (inverted)<br>tCS, LEAD tCS, LEAD<br>CS CS<br>1 / fSPI tCS, LAG 1 / fSPI tCS, LAG<br>SCLK  SCLK<br>(SPO = 0) (SPO = 0)<br>tSCLK_H/L tSCLK_H/L tSCLK_H/L tSCLK_H/L<br>SCLK  SCLK<br>(SPO = 1) (SPO = 1)<br>tSU,PI tSU,PI<br>tHD,PI tHD,PI<br>PICO PICO<br>tCS, ACC ttHD,POVALID,PO tCS, DIS tCS, ACC ttHD,POVALID,PO tCS, DIS<br>POCI POCI<br>Peripheral Mode, SPH = 0   Peripheral Mode, SPH = 1<br>**----- End of picture text -----**<br>


**Figure 7-5. SPI Timing Diagram - Peripheral Mode** 

## **7.17 UART** 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETERS**|**PARAMETERS**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|fUART|UART input clock frequency||32|MHz|
|fBITCLK|BITCLK clock frequency(equals<br>baud rate in MBaud)||4|MHz|
|tSP|Pulse duration of spikes<br>suppressed by input filter(1)|AGFSELx = 0|6|ns|
|||AGFSELx = 1|14<br>35||
|||AGFSELx = 2|22<br>60||
|||AGFSELx = 3|35<br>90||



- (1) Pulses on the UART receive input (RX) that are shorter than the UART receive deglitch time are suppressed. Thus the selected deglitch time can limit the maximum useable baud rate. To ensure that pulses are correctly recognized, their duration should exceed the maximum specification of the deglitch time. 

## **7.18 TIMx** 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETERS**|**PARAMETERS**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|tres|Timer resolution time|fTIMxCLK= 32MHz|31.25|ns|
||||1|tTIMxCLK|



## **7.19 TRNG Electrical Characteristics** 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|TRNGIACT|TRNG active current|TRNG clock = 20MHz|115|µA|



## **7.20 TRNG Switching Characteristics** 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|TRNGCLKF|TRNG input clock frequency||9.5<br>10<br>25|MHz|
|TRNGSTARTUP|TRNG startup time||520|µs|



Copyright © 2025 Texas Instruments Incorporated 

42 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|TRNGLAT32|Latency to generate 32 random bits|Decimation ratio = 4, TRNG clock =<br>20MHz|6.4|µs|
|TRNGLAT256|Latency to generate 256 random bits|Decimation ratio = 4, TRNG clock =<br>20MHz|51.2|µs|



## **7.21 Emulation and Debug** 

## _**7.21.1 SWD Timing**_ 

over operating free-air temperature range (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|fSWD|SWD frequency||10|MHz|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 43 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **8 Detailed Description** 

The following sections describe all of the components that make up the devices in this data sheet. The peripherals integrated into these devices are configured by software through Memory Mapped Registers (MMRs). For more details, see the corresponding chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.1 Functional Block Diagram** 

**==> picture [436 x 450] intentionally omitted <==**

**----- Start of picture text -----**<br>
PAx, PBx<br>tIOBUSt ULPCLK<br>CPU SUB SYSTEM GPIO DMA<br>3-ch<br>Arm<br>Cortex-M0+ FLASH B0 CRC-P<br>fmax = 32 MHz Up to 64KB 16/32-bit<br>NVIC FLASH B1<br>MPU Up to 64KB SPI0 POCI, PICO,<br>SCK, CSx<br>SWD + MTB<br>SRAM TEMP SENSOR<br>IOPORT 16KB<br>TRNG ADC0 13-CH (EXT)<br>12-bit A0_x<br>ULPCLK<br>IOMUX PD1 PERIPHERAL BUS (MCLK) AES-ADV<br>128/256-bit<br>WWDT0<br>SWCLK,<br>DEBUG<br>SWDIO<br>IWDT<br>RTC_OUT RTC_B KEY<br>STORE<br>EVENT FLASHCTL<br>CTS, RTSTX, RX,  UART0 PMCU (SYSCTL) VREF VREF+, VREF-<br>TX, RX,<br>CTS, RTS UART1 TIMG0<br>CKM PMU 2-CH<br>TIMG1<br>SDA, SCL I2C0 LFOSC LDO<br>SYSOSC BOR TIMG8 2-CH<br>LFXT QEI/HALL<br>POR<br>4-CH<br>TIMA0<br>VBOOST FAULT<br>LEGEND<br>PD1, CPU ACCESS ONLY LFXIN, LFXOUT VDD, VSS<br>PD1, CPU/DMA ACCESS ROSC VCORE, NRST<br>PD1/PD0, CPU/DMA ACCESS CLK_OUT, FCC_IN<br>PD0, CPU/DMA ACCESS<br>ROM<br>BCR, BSL<br>AHB BUS (MCLK) PD1 PERIPHERAL BUS (MCLK)<br>PD0 PERIPHERAL BUS (ULPCLK)<br>PD0 PERIPHERAL BUS (ULPCLK)<br>**----- End of picture text -----**<br>


**Figure 8-1. MSPM0L111x Functional Block Diagram** 

## **8.2 CPU** 

The CPU subsystem (MCPUSS) implements an Arm Cortex-M0+ CPU, an instruction prefetch and cache, a system timer, a memory protection unit, and interrupt management features. The Arm Cortex-M0+ is a costoptimized 32-bit CPU that delivers high performance and low power to embedded applications. Key features of the CPU Sub System include: 

Copyright © 2025 Texas Instruments Incorporated 

44 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

- Arm Cortex-M0+ CPU supports clock frequencies from 32kHz to 32MHz 

   - ARMv6-M Thumb instruction set (little endian) with single-cycle 32×32 multiply instruction 

   - Single-cycle access to GPIO registers through Arm single-cycle IO port 

- Prefetch logic to improve sequential code execution, and I-cache with 2 64-bit cache lines 

- System timer (SysTick) with 24-bit down counter and automatic reload 

- Memory protection unit (MPU) with 8 programmable regions 

- Nested vectored interrupt controller (NVIC) with 4 programmable priority levels and tail chaining 

- Interrupt groups for expanding the total interrupt sources, with jump index for low interrupt latency 

## **8.3 Operating Modes** 

MSPM0 MCUs provide five main operating modes (power modes) to allow for optimization of the device power consumption based on application requirements. In order of decreasing power, the modes are: RUN, SLEEP, STOP, STANDBY, and SHUTDOWN. The CPU is active executing code in RUN mode. Peripheral interrupt events can wake the device from SLEEP, STOP, or STANDBY mode to the RUN mode. SHUTDOWN mode completely disables the internal core regulator to minimize power consumption, and wake is only possible via NRST, SWD, or a logic level match on certain IOs. RUN, SLEEP, STOP, and STANDBY modes also include several configurable policy options (e,g, RUN.x) for balancing performance with power consumption. 

To further balance performance and power consumption, MSPM0 devices implement two power domains: PD1 (for the CPU, memories, and high performance peripherals), and PD0 (for low speed, low power peripherals). PD1 is always powered in RUN and SLEEP modes, but is disabled in all other modes. PD0 is always powered in RUN, SLEEP, STOP, and STANDBY modes. PD1 and PD0 are both disabled in SHUTDOWN mode. 

## _**8.3.1 Functionality by Operating Mode**_ 

Supported functionality in each operating mode is given in Table 8-1. 

Functional key: 

- **EN** : The function is enabled in the specified mode. 

- **DIS** : The function is disabled (either clock or power gated) in the specified mode, but the function's configuration is retained. 

- **OPT** : The function is optional in the specified mode, and remains enabled if configured to be enabled. 

- **NS** : The function is not automatically disabled in the specified mode, but its use is not supported. 

- **OFF** : The function is fully powered off in the specified mode, and no configuration information is retained. When waking up from an OFF state, all module registers must be re-configured to the desired settings by application software. 

**Table 8-1. Supported Functionality by Operating Mode** 

|||**RUN**|**RUN**|**RUN**|**SLEEP**|**SLEEP**|**SLEEP**|**STOP**|**STOP**|**STOP**|**STANDBY**|**STANDBY**|**N**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|**OPERATING MODE**||**RUN0**|**RUN1**|**RUN2**|**SLEEP0**|**SLEEP1**|**SLEEP2**|**STOP0**|**STOP1**|**STOP2**|**STANDBY0**|**STANDBY1**|**SHUTDOW**|
|Oscillators|SYSOSC|EN||DIS|EN||DIS|OPT(1)|EN|DIS|DIS||OFF|
||LFOSC or<br>LFXT|EN (LFOSC or LFXT)|||||||||||OFF|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 45 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **MSPM0L1117, MSPM0L1116** 

**Table 8-1. Supported Functionality by Operating Mode (continued)** 

|||**RUN**|**RUN**|**RUN**|**SLEEP**|**SLEEP**|**SLEEP**|**STOP**|**STOP**|**STOP**|**STANDBY**|**STANDBY**|**N**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|**OPERATING MODE**||**RUN0**|**RUN1**|**RUN2**|**SLEEP0**|**SLEEP1**|**SLEEP2**|**STOP0**|**STOP1**|**STOP2**|**STANDBY0**|**STANDBY1**|**SHUTDOW**|
|Clocks|CPUCLK|32MHz|32kHz||DIS||||||||OFF|
||MCLK to PD1|32MHz|32kHz||32MHz|32kHz||DIS|||||OFF|
||ULPCLK to<br>PD0|32MHz|32kHz||32MHz|32kHz||4MHz(1)|4MHz|32kHz|32kHz|DIS|OFF|
||ULPCLK to<br>TIMG0/8|32MHz|32kHz||32MHz|32kHz||4MHz(1)|4MHz|32kHz|32kHz|32kHz(2)|OFF|
||MFCLK|OPT|DIS||OPT|DIS||OPT||DIS|DIS||OFF|
||LFCLK to<br>PD0/1|32kHz||||||||||DIS|OFF|
||LFCLK to<br>TIMG0/8|32kHz||||||||||32kHz(2)|OFF|
||LFCLK<br>Monitor|OPT|||||||||||OFF|
||MCLK Monitor|OPT||||||||||DIS|OFF|
|PMU|POR monitor|EN||||||||||||
||BOR monitor|EN|||||||||||OFF|
||Core regulator|FULL DRIVE||||||REDUCED DRIVE|||LOW DRIVE||OFF|
|Core<br>Functions|CPU|EN|||DIS||||||||OFF|
||DMA|OPT||||||DIS (triggers supported)|||||OFF|
||Flash|EN||||||DIS|||||OFF|
||SRAM|EN||||||DIS|||||OFF|
|PD1<br>Peripherals|SPI0|OPT||||||DIS|||||OFF|
||AESADV|OPT||||||OFF|||||OFF|
||CRC-P|OPT||||||DIS|||||OFF|
||TRNG|OPT||||||OFF|||||OFF|
|PD0<br>Peripherals|GPIOA/B(3)|OPT||||||||||OPT(2)|OFF|
||UART0/1|OPT||||||||||OPT(2)|OFF|
||I2C0|OPT||||||||||OPT(2)|OFF|
||TIMG0/1/8|OPT||||||||||OPT(2)|OFF|
||TIMA0|OPT|||||||||DIS||OFF|
||WWDT0|OPT||||||||||DIS|OFF|
||IWDT|OPT|||||||||||OFF|
||RTC_B|OPT|||||||||||OFF|
||Keystore|OPT|||||||||||OFF|
|Analog|VREF|OPT|||||||||OFF||OFF|
||ADC0(3)|OPT||||||||NS (triggers supported)|||OFF|
||Temperature<br>Sensor|OPT|||||||||OFF||OFF|
|IOMUX and IO Wakeup||EN|||||||||||DIS w/<br>WAKE|
|Wake Sources||N/A|||ANY IRQ|||PD0 IRQ|||||IOMUX,<br>NRST,<br>SWD|



(1) If STOP0 is entered from RUN1 (SYSOSC enabled but MCLK sourced from LFCLK), SYSOSC remains enabled as in RUN1 and ULPCLK remains at 32kHz as in RUN1. If STOP0 is entered from RUN2 (SYSOSC was disabled and MCLK was sourced from LFCLK), SYSOSC remains disabled as in RUN2 and ULPCLK remains at 32kHz as in RUN2. 

(2) When using the STANDBY1 policy for STANDBY, only specific peripherals (TIMG0, TIMG8, and RTC) are clocked. Other PD0 peripherals can generate an asynchronous fast clock request upon external activity but are not actively clocked. 

(3) For ADCx and GPIO Ports A and B, the digital logic is in PD0 and the register interface is in PD1. These peripherals support fast single-cycle register access when PD1 is active and also support basic operation down to STANDBY mode where PD0 is still active. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

46 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **8.4 Power Management Unit (PMU)** 

The power management unit (PMU) generates the internally regulated core supplies for the device and provides supervision of the external supply (VDD). The PMU also contains the bandgap voltage reference used by the PMU itself as well as analog peripherals. Key features of the PMU include: 

- Power-on reset (POR) supply monitor 

- Brownout reset (BOR) supply monitor with early warning capability using three programmable thresholds 

- Core regulator with support for RUN, SLEEP, STOP, and STANDBY operating modes to dynamically balance performance with power consumption 

- Parity-protected trim to immediately generate a power-on reset (POR) in the event that a power management trim is corrupted 

For more details, see the PMU chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.5 Clock Module (CKM)** 

The clock module provides the following oscillators: 

- **LFOSC** : Internal low-frequency oscillator (32kHz) 

- **SYSOSC** : Internal high-frequency oscillator (4MHz or 32MHz with factory trim, 16MHz or 24MHz with user trim) 

- **LFXT/LFCKIN** : low-frequency external crystal oscillator or digital clock input (32kHz) 

The following clocks are distributed by the clock module for use by the processor, bus, and peripherals: 

- **MCLK** : Main system clock for PD1 peripherals, derived from SYSOSC or LFCLK, active in RUN and SLEEP modes 

- **CPUCLK** : Clock for the processor (derived from MCLK), active in RUN mode 

- **ULPCLK** : Ultra-low power clock for PD0 peripherals, active in RUN, SLEEP, STOP, and STANDBY modes 

- **MFCLK** : 4MHz fixed mid-frequency clock for peripherals, available in RUN, SLEEP, and STOP modes 

- **LFCLK** : 32kHz fixed low-frequency clock for peripherals or MCLK, active in RUN, SLEEP, STOP, and STANDBY modes 

- **ADCCLK** : ADC clock, available in RUN, SLEEP and STOP modes 

- **CLK_OUT** : Used to output a clock externally, available in RUN, SLEEP, STOP, and STANDBY modes 

- **HFCLK** : High frequency clock derived from HFCLK_IN, available in RUN and SLEEP mode 

For more details, see the CKM chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.6 DMA** 

The direct memory access (DMA) controller allows movement of data from one memory address to another without CPU intervention. For example, the DMA can be used to move data from ADC conversion memory to SRAM. The DMA reduces system power consumption by allowing the CPU to remain in low power mode, without having to awaken to move data to or from a peripheral. 

The DMA in these devices support the following key features: 

- 3 independent DMA transfer channels 

   - 1 full-feature channel (DMA0), supporting repeated transfer modes 

   - 2 basic channels (DMA1, DMA2), supporting single transfer modes 

- Configurable DMA channel priorities 

- Byte (8-bit), short word (16-bit), word (32-bit) and long word (64-bit) and long-long word (128-bit) or mixed byte and word transfer capability 

- Transfer counter block size supports up to 64k transfers of any data type 

- Configurable DMA transfer trigger selection 

- Active channel interruption to service other channels 

- Early interrupt generation for ping-pong buffer architecture 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 47 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

- Cascading channels upon completion of activity on another channel 

- Stride mode to support data re-organization, such as 3-phase metering applications 

- Gather mode 

## **Table 8-2. DMA Features** 

|**Feature**|**FULL**|**BASIC**|
|---|---|---|
|Channel#|0|1,2|
|Repeat Mode|Yes|-|
|Table & Fill Mode|Yes|-|
|Gather Mode|Yes|-|
|Pre-IRQ|Yes|-|
|Auto Enable|Yes|Yes|
|Long Long (128-bit) Transfer|Yes|Yes|
|Stride Mode|Yes|Yes|
|Cascading Channel Support|Yes|Yes|



Table 8-3 lists the available triggers for the DMA which are configured using the DMATCTL.DMATSEL control bits in the DMA memory mapped registers. 

**Table 8-3. DMA Trigger Mapping** 

|**DMACTL.DMATSEL**|**Trigger Source**|
|---|---|
|0|Software|
|1|Generic Subscriber 0 (FSUB_0)|
|2|Generic Subscriber 0 (FSUB_1)|
|3|AES Publisher 1|
|4|AES Publisher 2|
|5|ADC0 Publisher 1|
|6|I2C0 Publisher 1|
|7|I2C0 Publisher 2|
|8|SPI0 Publisher 1|
|9|SPI0 Publisher 2|
|10|UART0 Publisher 1|
|11|UART0 Publisher 2|
|12|UART1 Publisher 1|
|13|UART1 Publisher 2|



For more details, see the DMA chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.7 Events** 

The event manager transfers digital events from one entity (for example, a peripheral) to another (for example, a second peripheral, the DMA or the CPU). The event manager implements event transfer through a defined set of event publishers (generators) and subscribers (receivers) that are interconnected through an event fabric containing a combination of static and programmable routes. 

Events that are transferred by the event manager include: 

- Peripheral event transferred to the CPU as an interrupt request (IRQ) (Static Event) 

   - Example: GPIO interrupt is sent to the CPU 

- Peripheral event transferred to the DMA as a DMA trigger (DMA Event) 

Copyright © 2025 Texas Instruments Incorporated 

48 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

   - Example: UART data receive trigger to DMA to request a DMA transfer 

- Peripheral event transferred to another peripheral to directly trigger an action in hardware (Generic Event) 

   - Example: TIMx timer peripheral publishes a periodic event to the ADC subscriber port, and the ADC uses the event to trigger start-of-sampling 

For more details, see the Events chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **Table 8-4. Generic Event Channels** 

A generic route is either a point-to-point (1:1) route or a point-to-two (1:2) splitter route in which the peripheral publishing the event is configured to use one of several available generic route channels to publish the event to another entity (or entities, in the case of a splitter route). An entity can be another peripheral, a generic DMA trigger event, or a generic CPU event. 

|**CHANID**|**Generic Route Channel Selection**|**Channel Type**|
|---|---|---|
|0|No generic event channel selected|N/A|
|1|Generic event channel 1 selected|1 : 1|
|2|Generic event channel 2 selected|1 : 1|
|3|Generic event channel 3 selected|1 : 2 (splitter)|



## **8.8 Memory** 

## _**8.8.1 Memory Organization**_ 

Table 8-5 summarizes the memory map of the devices. For more information about the memory region detail, see the _Platform Memory Map_ section in the _MSPM0 L-Series 32-MHz Microcontrollers Technical Reference Manual ._ 

**Table 8-5. Memory Organization** 

|**MEMORY REGION**|**SUBREGION**|**MSPM0L1116**|**MSPM0L1117**|
|---|---|---|---|
|Code (Flash Bank 0)|MAIN ECC Corrected|32KB<br>0x0000.0000 to 0x0000.7FFF|64KB<br>0x0000.0000 to 0x0000.FFFF|
||MAIN ECC Uncorrected|0x0040.0000 to 0x0040.7FFF|0x0040.0000 to 0x0040.FFFF|
||Flash ECC Code|0x0080.0000 to 0x0080.7FFF|0x0080.0000 to 0x0080.FFFF|
|Code (Flash Bank 1)|MAIN ECC Corrected|32KB<br>0x0001.0000 to 0x0001.7FFF|64KB<br>0x0001.0000 to 0x0001.FFFF|
||MAIN ECC Uncorrected|0x0041.0000 to 0x0041.7FFF|0x0041.0000 to 0x0041.FFFF|
||Flash ECC Code|0x0081.0000 to 0x0081.7FFF|0x0081.0000 to 0x0081.FFFF|
|SRAM (SRAM)|Default|16KB<br>0x2000.0000 to 0x2000.3FFF|16KB<br>0x2000.0000 to 0x2000.3FFF|
|Peripheral|Peripherals|0x4000.4000 to 0x4086.1FFF|0x4000.4000 to 0x4086.1FFF|
||NONMAIN Corrected|2KB<br>0x41C0.0000 to 0x41C0.07FF|2KB<br>0x41C0.0000 to 0x41C0.07FF|
||NONMAIN Uncorrected|0x41C1.0000 to 0x41C1.07FF|0x41C1.0000 to 0x41C1.07FF|
||NONMAIN ECC code|0x41C2.0000 to 0x41C2.07FF|0x41C2.0000 to 0x41C2.07FF|
||FACTORY Corrected|0x41C4.0000 to 0x41C4.01FF|0x41C4.0000 to 0x41C4.01FF|
||FACTORY Uncorrected|0x41C5.0000 to 0x41C5.01FF|0x41C5.0000 to 0x41C5.01FF|
||FACTORY ECC code|0x41C6.0000 to 0x41C6.01FF|0x41C6.0000 to 0x41C6.01FF|
|Subsystem||0x6000.0000 to 0x7FFF.FFFF|0x6000.0000 to 0x7FFF.FFFF|
|System PPB||0xE000.0000 to 0xE00F.FFFF|0xE000.0000 to 0xE00F.FFFF|



## _**8.8.2 Peripheral File Map**_ 

Table 8-6 lists the available peripherals and the register base address for each. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 49 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **MSPM0L1117, MSPM0L1116** 

**Table 8-6. Peripherals Summary** 

|**Peripheral Name**|**Base Address**|**Size**|
|---|---|---|
|ADC0|0x4000.4000|0x2000|
|VREF|0x4003.0000|0x2000|
|WWDT0|0x4008.0000|0x2000|
|TIMG0|0x4008.4000|0x2000|
|TIMG8|0x4009.0000|0x2000|
|RTC_B|0x4009.4000|0x2000|
|GPIOA|0x400A.0000|0x2000|
|GPIOB|0x400A.2000|0x2000|
|KEYSTORE|0x400A.C000|0x2000|
|SYSCTL|0x400A.F000|0x4000|
|DEBUGSS|0x400C.7000|0x2000|
|EVENT|0x400C.9000|0x3000|
|NVMNW|0x400C.D000|0x2000|
|I2C0|0x400F.0000|0x2000|
|UART1|0x4010.0000|0x2000|
|UART0|0x4010.8000|0x2000|
|MCPUSS|0x4040.0000|0x2000|
|MTB|0x4040.2000|0x1000|
|MTBRAM|0x4040.3000|0x0020|
|IOMUX|0x4042.8000|0x2000|
|DMA|0x4042.A000|0x2000|
|CRC|0x4044.0000|0x2000|
|AESADV|0x4044.2000|0x2000|
|TRNG|0x4044.4000|0x2000|
|SPIO|0x4046.8000|0x2000|
|TIMG1|0x4048.6000|0x2000|
|ADC0(1)|0x4055.6000|0x2000|
|TIMA0|0x4086.0000|0x2000|



1. Aliased region of ADC0 memory-mapped registers 

## _**8.8.3 Peripheral Interrupt Vector**_ 

Table 8-7 shows the IRQ number and the interrupt group number for each peripherals in this device. 

**Table 8-7. Interrupt Vector Number** 

|**Peripheral Name**|**NVIC IRQ**|**Group IIDX**|
|---|---|---|
|WWDT0|0|0|
|DEBUGSS|0|2|
|FLASHCTL|0|3|
|EVENT SUB PORT 0|0|4|
|EVENT SUB PORT 1|0|5|
|SYSCTL|0|6|
|GPIOA|1|0|
|GPIOB|1|1|
|TRNG|1|5|
|TIMG8|2|-|



Copyright © 2025 Texas Instruments Incorporated 

50 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

**Table 8-7. Interrupt Vector Number (continued)** 

|**Peripheral Name**|**NVIC IRQ**|**Group IIDX**|
|---|---|---|
|ADC0|4|-|
|SPI0|9|-|
|UART1|13|-|
|UART0|15|-|
|TIMG0|16|-|
|TIMA0|18|-|
|TIMG1|22|-|
|I2C0|24|-|
|AESADV|28|-|
|RTC_B|30|-|
|DMA0|31|-|



## **8.9 Flash Memory** 

A dual bank of non-volatile flash memory (up to 64kB/128kB total) is provided for storing executable program code and application data. 

Key features of the flash include: 

- Hardware ECC protection (encode and decode) with single bit error correction and double-bit error detection 

- In-circuit program and erase operations supported across the entire recommended supply range 

- Small 1kB sector sizes (minimum erase resolution of 1kB) 

- Up to 100,000 program/erase cycles on the 32 selected sectors of the flash memory, with up to 10,000 program/erase cycles on the remaining flash memory (devices with 32kB support 100,000 cycles on the entire flash memory) 

- Bank address swap for in-system, over-the-air (OTA) firmware updates 

For more details, see the NVM chapter of the _MSPM0 L-Series 32-MHz Microcontrollers Technical Reference Manual_ . 

## **8.10 SRAM** 

MSPM0L111x MCUs include a low power, high performance SRAM memory with zero wait state access across the supported CPU frequency range of the device. MSPM0 MCUs also provide up to 16KB SRAM. SRAM memory may be used for storing volatile information such as the call stack, heap, global data, and code. 

The SRAM memory content is fully retained in run, sleep, stop, and standby operating modes and is lost in shutdown mode. 

A write-execute mutual exclusion mechanism is provided to allow the application to partition the SRAM into two sections: a read-write (RW) partition and a read-execute (RX) partition. The SRAMBOUNDARY register in SYSCTL needs to be configured to set up these partitions. The RX partition occupies the upper portion of the SRAM address space. Write protection is useful when placing executable code into SRAM as it provides a level of protection against unintentional overwrites of code by either the CPU or DMA. Placing code in SRAM can improve performance of critical loops by enabling zero wait state operation and lower power consumption. Preventing code execution from the RW partition improves security by preventing self-modifying code execution ability. 

## **8.11 GPIO** 

The general purpose input/output (GPIO) peripheral provides the user with a means to write data out and read data in to and from the device pins. Through the use of the Port A and Port B GPIO peripheral, these devices support up to 44 GPIO pins . 

The key features of the GPIO module include: 

- 0 wait state MMR access from CPU 

Copyright © 2025 Texas Instruments Incorporated 

51 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

- Set/Clear/Toggle multiple bits without the need of a read-modify-write construct in software 

- GPIOs with "Standard with Wake" drive functionality able to wake the device from SHUTDOWN mode 

- User controlled input filtering 

- GPIO "FastWake" feature enables low-power wakeup from STOP and STANDBY modes for any GPIO port 

For more details, see the GPIO chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.12 IOMUX** 

The IOMUX peripheral enables IO pad configuration and controls digital data flow to and from the device pins. The key features of the IOMUX include: 

- IO pad configuration registers allow for programmable drive strength, speed, pullup or pulldown, and more 

- Digital pin muxing allows for multiple peripheral signals to be routed to the same IO pad 

- Pin functions and capabilities are user-configured using the PINCM register 

For more details, see the IOMUX chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.13 ADC** 

The 12-bit analog-to-digital converter (ADC) module in these devices support fast 12-bit conversions with singleended inputs. 

ADC features include: 

- 12-bit output resolution at up to 1.68Msps with greater than 11-bit ENOB 

- HW averaging enables 14-bit conversion resolution at 105ksps 

- Up to 13 external input channels 

- Internal channels for temperature sensing and supply monitoring 

- Software selectable reference: 

   - Configurable internal dedicated ADC reference voltage of 1.4V and 2.5V (VREF) 

   - MCU supply voltage (VDD) 

   - External reference supplied to the ADC through the VREF+ and VREF- pins 

- Operates in RUN, SLEEP, and STOP modes and supports triggers from STANDBY mode 

Table 8-8 shows the ADC channel connections. 

**Table 8-8. ADC0 Channel Mapping** 

|**Channel [0:15]**|**Signal Name (ADC0)**|**Channel [16:31]**|**Signal Name (ADC0)**<br>(1)<br>(2)|
|---|---|---|---|
|0|A0_0|16|-|
|1|A0_1|17|-|
|2|A0_2|18|-|
|3|A0_3|19|-|
|4|A0_4|20|-|
|5|A0_5|21|-|
|6|A0_6|22|-|
|7|A0_7|23|-|
|8|A0_8|24|-|
|9|A0_9|25|-|
|10|-|26|-|
|11|_Temperature Sensor_|27|-|
|12|A0_12|28|_Internal VREF_|



Copyright © 2025 Texas Instruments Incorporated 

52 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

**Table 8-8. ADC0 Channel Mapping (continued)** 

|**Channel [0:15]**|**Signal Name (ADC0)**|**Channel [16:31]**|**Signal Name (ADC0)**<br>(1)<br>(2)|
|---|---|---|---|
|13|A0_13|29|-|
|14|A0_14|30|-|
|15|-|31|_Supply/Battery Monitor_|



(1) _Italicized_ signal names are internal to the SoC. These signals are used for internal peripheral interconnections. 

(2) For more information about device analog connections see Section 8.29. 

For more details, see the ADC chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.14 Temperature Sensor** 

The temperature sensor provides a voltage output that changes linearly with device temperature. The temperature sensor output is internally connected to one of ADC input channels to enable a temperature-todigital conversion. 

A unit-specific single-point calibration value for the temperature sensor is provided in the factory constants memory region. This calibration value represents the ADC conversion result (in ADC code format) corresponding to the temperature sensor being measured in 12-bit mode with the 1.4V internal VREF at the factory trim temperature (TSTRIM). 

The ADC and VREF configuration for the above measurement is as the following: RES=0 (12-bit mode), VRSEL=2h (Internal reference), BUFCONFIG=1h (1.4V VREF), ADC tSample=12.5µs. This calibration value can be used with the temperature sensor temperature coefficient (TSc) to estimate the device temperature. 

## **Per-unit TSc calculation method (using VTRIM_0K)** 

An additional unit-specific calibration value (VTRIM_0K) is provided for use in calculating per-unit TSc performance. This calibration value represents the ADC conversion result (in ADC code format) corresponding to the temperature sensor being measured in 12-bit mode with the 1.4V internal VREF at 0°K (or -273.15°C) and is stored in the factory constants memory region (as TEMP_SENSE_0KELVIN at address 0x41C4.0040). 

The temperature coefficient TSC can then be calculated using the formula below: 

**==> picture [178 x 10] intentionally omitted <==**

**==> picture [13 x 11] intentionally omitted <==**

## **Example** 

To illustrate the process of calculating the temperature sensor coefficient using this method, an example is given below. 

Example parameters: 

- VSAMPLE = 0.6427V 

- V = 1.2033V TRIM_0K 

- TSAMPLE = 30°C 

The resulting unit-specific temperature coefficient is calculated as: 

## **TSC** = (0.6427V - 1.2033V) / (30°C + 273.15°C) = -1.8492mV/°C 

**==> picture [13 x 11] intentionally omitted <==**

See the temperature sensor section of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ for guidance on estimating the device temperature with the factory trim value. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 53 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **8.15 VREF** 

The voltage reference module (VREF) in these devices contains a configurable voltage reference buffer dedicated for the on-board ADC. The devices also support connection of an external reference for applications in which higher accuracy is required. 

VREF features include: 

- 1.4V and 2.5V user-selectable internal reference for ADC 

- Internal reference supports ADC operation up to 200 ksps 

- Support for bringing in an external reference on VREF+/- device pins 

**==> picture [455 x 227] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>Boundary<br>Voltage Reference (VREF)<br>CTL0 CTL1 CTL2 VREF to on-board<br>analog peripherals<br>ADC<br>ADC<br>ENABLE<br>Bandgap reference from PMCU + Internal Ref VREF+<br>–<br>+ Optional<br>– External Ref<br>BUFCONFIG<br>VSS/VREF-<br>ENABLE<br>**----- End of picture text -----**<br>


**Figure 8-2. VREF module** 

For more details, see the VREF chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.16 Security** 

This device offers several security features, including: 

- Debug security 

- Device identify 

- AES-128/256 accelerator with support for GCM/GMAC, CCM/CBC-MAC, CBC, CTR 

- Flexible firewalls for protecting code and data 

   - Flash write-erase protection 

   - 

      - Flash read-execute protection 

   - Flash IP protection 

   - SRAM write-execute mutual exclusion 

- Secure boot 

- Secure firmware update 

- Secure key storage for up to two AES keys 

- Customer secure code 

- Hardware monotonic counter 

- True random number generator (TRNG) 

- Cyclic redundancy checker (CRC-16, CRC-32) with support for custom polynomial 

For more details, see the Security chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

Copyright © 2025 Texas Instruments Incorporated 

54 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **8.17 TRNG** 

The true random number generator (TRNG) utilizes an internal circuit to generate 32-bit random numbers. The TRNG is intended to be used as a source to a deterministic random number generator (DRNG) to build a FIPS-140-2 compliant system. Key features of the TRNG include: 

- Generation of 32-bit random numbers 

- A new 32-bit number may be generated every 32 * 4 = 128 TRNG clock cycles 

- Built-in health tests 

- Available in RUN and SLEEP modes 

For more details, see the TRNG chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.18 AESADV** 

The AES advanced (AESADV) accelerator module performs encryption and decryption of 128-bit data blocks with a 128-bit or 256-bit key in hardware according to the advanced encryption standard (AES). AES is a symmetric-key block cipher algorithm specified in FIPS PUB 197. 

The AESADV accelerator features include: 

- AES operation with 128-bit and 256-bit keys 

- Key scheduling in hardware 

- Enc/decrypt only modes: CBC, CFB-1, CFB-8, CFB-128, OFB-128, CTR/ICM 

- Authentication only modes: CBC-MAC, CMAC 

- AES-CCM 

- AES-GCM 

- AES-CCM and AES-GCM modes support continuation with hold/resume of payload data 

- 32-bit word access to provide key data, input data, and output data 

- AESADV ready interrupt 

- DMA triggers for input/output data 

- Supported in RUN and SLEEP (see the _Operating Modes_ section of the device technical reference manual) 

For more details, see the AESADV chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.19 Keystore** 

The Keystore controller provides secure management of the Advanced Encryption Engine (AES) keys. The use-model of the keystore controller is to securely deposit keys into it during the execution of customer secure code, and have the AES engine access them subsequently in a secure manner without leaking any key data to observers. Both 128 and 256-bit keys can be stored in the keystore's key slots. The keystore and its interaction with the AES engine are designed for secure operation including thwarting partial key modification attacks. 

- Support for storage of up to 2 keys 

For more details, see the KEYSTORE chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.20 CRC-P** 

The cyclical redundancy check (CRC) module provides a signature for an input data sequence. Key features of the CRC module include: 

- Support for 16-bit CRC based on CRC16-CCITT 

- Support for 32-bit CRC based on CRC32-ISO3309 

- Support for bit reversal 

- Support for custom polynomials 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 55 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

For more details, see the CRC chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.21 UART** 

The UART peripherals provide the following key features: 

- Standard asynchronous communication bits for start, stop, and parity 

- Fully programmable serial interface 

   - 5, 6, 7 or 8 data bits 

   - Even, odd, stick, or no-parity bit generation and detection 

   - 1 or 2 stop bit generation 

   - Line-break detection 

   - Glitch filter on the input signals 

   - Programmable baud rate generation with oversampling by 16, 8 or 3 

   - Local Interconnect Network (LIN) mode support 

- Separated transmit and receive FIFOs support DMA data transfer 

- Support transmit and receive loopback mode operation 

- See Table 8-9 for detail information on supported protocols 

**Table 8-9. UART Features** 

|**UART Features**|**UART0 (Extend, low-power)**|**UART1 (Main, low-power)**|
|---|---|---|
|Active in Stop and Standby Mode|Yes|Yes|
|Separate transmit and receive FIFOs|Yes|Yes|
|Support hardware flow control|Yes|Yes|
|Support 9-bit configuration|Yes|Yes|
|Support LIN mode|Yes|-|
|Support DALI|Yes|-|
|Support IrDA|Yes|-|
|Support ISO7816 Smart Card|Yes|-|
|Support Manchester coding|Yes|-|



For more details, see the UART chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ 

## **8.22 I2C** 

The inter-integrated circuit interface (I[2] C) peripherals in these devices provide bidirectional data transfer with other I2C devices on the bus and support the following key features: 

- 7-bit and 10-bit addressing mode with multiple 7-bit target addresses 

- Multiple-controller transmitter or receiver mode 

- Target receiver or transmitter mode with configurable clock stretching 

- Support Standard-mode (Sm), with a bit rate up to 100 kbit/s 

- Support Fast-mode (Fm), with a bit rate up to 400 kbit/s 

- Support Fast-mode Plus (Fm+), with a bit rate up to 1 Mbit/s 

   - Supported on open drain IOs only (ODIO) 

- Separated transmit and receive FIFOs support DMA data transfer 

- Support SMBus 3.0 with PEC, ARP, timeout detection and host support 

- Wakeup from low power mode on address match 

- Support analog and digital glitch filter for input signal glitch suppression 

- 8-entry transmit and receive FIFOs 

For more details, see the I2C chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

Copyright © 2025 Texas Instruments Incorporated 

56 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **8.23 SPI** 

The serial peripheral interface (SPI) peripheral in these devices support the following key features: 

- Support MCLK/2 bit rate and up to 16Mbits/s in both controller and peripheral mode 

- Configurable as a controller or a peripheral 

- Configurable chip select for both controller and peripheral 

- Programmable clock prescaler and bit rate 

- Programmable data frame size from 4 bits to 16 bits (controller mode) and 7 bits to 16 bit (peripheral mode) 

- Supports PACKEN feature that allows the packing of two 16 bit FIFO entries into a 32-bit value to improve CPU performance 

- Transmit and receive FIFOs (four entries each with 16 bits per entry) supporting DMA data transfer 

- Supports TI mode, Motorola mode and National Microwire format 

For more details, see the SPI chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.24 Low-Frequency Sub System (LFSS)** 

The Low-Frequency Sub-System (LFSS) is a sub-system which combines several functional peripherals under one shared subsystem. These peripherals are clocked by the low frequency clock (LFCLK) or need to be active during low power modes. The LFCLK has a typical frequency of 32kHz and is mainly intended for long-term timekeeping. 

LFSS_B is the specific LFSS variant in this device and contains following components: 

- Real Time Clock with additional prescalar extension and timestamp captures 

- An asynchronous Independent Watchdog Timer 

For more details, see the LFSS chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ **.** 

## **8.25 RTC_B** 

The RTC_B instance of the real-time clock operates off of a 32kHz input clock source (typically a low frequency crystal) and provides a time base to the application with multiple options for interrupts to the CPU. The RTC_B provides common key features in relation to the Low-Frequency Sub System (LFSS). 

Common key features of the RTC_B include: 

- Counters for seconds, minutes, hours, day of the week, day of the month, month, and year 

- Binary or BCD format 

- Leap-year handling 

- One customizable alarm interrupt based on minute, hour, day of the week, and day of the month 

- Interval alarm interrupt to wake every minute, every hour, at midnight, or at noon 

- Interval alarm interrupt providing periodic wake-up at 4096, 2048, 1024, 512, 256, or 128 Hz 

- Interval alarm interrupt providing periodic wake-up at 64, 32, 16, 8, 4, 2, 1, and 0.5 Hz 

- Calibration for crystal offset error (up to +/- 240ppm) 

- Compensation for temperature drift (up to +/- 240ppm) 

- RTC clock output to pin for calibration 

Table 8-10 shows the RTC features supported in this device. 

## **Table 8-10. RTC_B Key Features** 

|**RTC Features**|**RTC_B**|
|---|---|
|Power enable register|-|
|Real-time clock and calendar mode providing seconds, minutes,<br>hours, day of week, day of month, and year|Yes|
|Selectable binary or binary-coded decimal (BCD) format|Yes|
|Leap-year correction (valid for year 1901 through 2099)|Yes|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 57 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** 

SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**==> picture [46 x 7] intentionally omitted <==**

**----- Start of picture text -----**<br>
www.ti.com<br>**----- End of picture text -----**<br>


**Table 8-10. RTC_B Key Features (continued)** 

|**RTC Features**|**RTC_B**|
|---|---|
|Two customizable calendar alarm interrupts based on minute, hour,<br>day of the week, and day of the month|Yes|
|Interval alarm interrupt to wake every minute, every hour, at<br>midnight, or at noon|Yes|
|Periodic interrupt to wake at 4096, 2048, 1024, 512, 256, or 128Hz|Yes|
|Periodic interrupt to wake at 64, 32, 16, 8, 4, 2, 1, and 0.5Hz|Yes|
|Interrupt capability down to STANDBY mode with STOPCLKSTBY|Yes|
|Calibration for crystal offset error and crystal temperature drift (up to<br>±240 ppm total)|Yes|
|RTC clock output to pin for calibration (GPIO)|Yes|
|RTC clock output to pin for calibration (TIO)|-|
|Three -bit prescaler for heartbeat function with interrupt generation|-|
|RTC external clock selection of untrimmed 32kHz, trimmed 512Hz,<br>256Hz or 1Hz|-|
|RTC time stamp capture upon detection of a timer stamp event,<br>including:<br>•<br>TIO event<br>•<br>VDD fail event|-|
|RTC counter lock function|-|



For more details, see the RTC chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.26 IWDT_B** 

The independent watchdog timer (IWDT) in the LFSS is a device-independent supervisor which monitors code execution and overall hang up scenarios of the device. Due to the nature of LFSS, this IWDT has its own system independent clock source. If the application software does not successfully reset the watchdog within the programmed time, the watchdog generates a POR reset to the device. 

Key features of the IWDT include: 

- A 25-bit counter 

- Counter driven from LFOSC (fixed 32kHz clock path) with a programmable clock divider 

- Eight selectable watchdog timer periods (2ms to 2hr) 

For more details, see the IWDT chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.27 WWDT** 

The windowed watchdog timer (WWDT) can be used to supervise the operation of the device, specifically code execution. The WWDT can be used to generate a reset or an interrupt if the application software does not successfully reset the watchdog within a specified window of time. Key features of the WWDT include: 

- 25-bit counter 

- Programmable clock divider 

- Eight software selectable watchdog timer periods 

- Eight software selectable window sizes 

- Support for stopping the WWDT automatically when entering a sleep mode 

- Interval timer mode for applications which do not require watchdog functionality 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

58 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

For more details, see the WWDT chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.28 Timers (TIMx)** 

There are two types of timer peripherals in these devices support that following key features: TIMGx (generalpurpose timer) and TIMAx (advanced timer). TIMGx is a subset of TIMAx, which means these timers share many common features that are compatible in software. For specific configuration, see Table 8-11. 

Specific features for the general-purpose timer **(TIMGx)** include: 

- 16-bit up, down, up-down or down-up counter, with repeat-reload mode 

- Selectable and configurable clock source 

- 8-bit programmable prescaler to divide the counter clock frequency 

- Two independent CC channels for 

   - Output compare 

   - Input capture 

   - PWM output 

   - One-shot mode 

- Support quadrature encoder interface (QEI) and Hall sensor input logic available in TIMG8 

- Support synchronization and cross trigger among different TIMx instances in the same power domain (see Table 8-12) 

- Support interrupt/DMA trigger generation and cross peripherals (such as ADC) trigger capability 

Specific features for the advanced timer **(TIMAx)** include: 

- 16-bit timer with up, down or up-down counting modes, with repeat-reload mode 

- Selectable and configurable clock source 

- 8-bit programmable prescaler to divide the counter clock frequency 

- Repeat counter to generate an interrupt or event only after a given number of cycles of the counter 

- Up to four independent CC channels for 

   - Output compare 

   - Input capture 

   - PWM output 

   - One-shot mode 

- Two additional capture/compare channels for internal events (CC4/CC5) 

- Shadow register for load and CC register available in TIMA0 

- Complementary output PWM with programmable dead band insertion 

- Asymmetric PWM 

- Fault handling mechanism for 

   - Fast PWM responses (<40ns) to external fault inputs or comparator events 

   - Outputting signals in a safe user-defined state when a latched fault condition has occurred 

- Support synchronization and cross trigger among different TIMx instances in the same power domain (see Table 8-12 ) 

- Support interrupt and DMA trigger generation and cross peripherals (such as ADC) trigger capability 

- • Two additional capture/compare channels for internal events 

## **Table 8-11. TIMx Instance Configuration** 

|**TIMER**<br>**NAME**|**POWER**<br>**DOMAIN**|**RESOLUTI**<br>**ON**|**PRESCAL**<br>**ER**|**REPEAT**<br>**COUNTER**|**CAPTURE**<br>**/**<br>**COMPARE**<br>**CHANNEL**<br>**S**|**PHASE**<br>**LOAD**|**SHADOW**<br>**LOAD**|**SHADOW**<br>**CC**|**DEADBAN**<br>**D**|**FAULT**|**QEI**|
|---|---|---|---|---|---|---|---|---|---|---|---|
|TIMG0|PD0|16 bit|8 bit|-|2|-|-|-|-|-|-|
|TIMG1|PD0|16 bit|8 bit|-|2|-|-|-|-|-|-|
|TIMG8|PD0|16 bit|8 bit|-|2|-|-|-|-|-|Yes|
|TIMA0|PD0|16 bit|8 bit|8 bit|4|Yes|Yes|Yes|Yes|Yes|-|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 59 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**www.ti.com** 

## **MSPM0L1117, MSPM0L1116** 

## SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**Table 8-12. TIMx Cross Trigger Map (PD0)** 

|**TSEL.ETSEL Selection**|**TIMA0**|**TIMG0**|**TIMG1**|**TIMG8**|
|---|---|---|---|---|
|0|TIMA0.TRIGO|TIMA0.TRIGO|TIMA0.TRIGO|TIMA0.TRIGO|
|1|TIMG0.TRIGO|TIMG0.TRIGO|TIMG0.TRIGO|TIMG0.TRIGO|
|2|TIMG1.TRIGO|TIMG1.TRIGO|TIMG1.TRIGO|TIMG1.TRIGO|
|3|TIMG8.TRIGO|TIMG8.TRIGO|TIMG8.TRIGO|TIMG8.TRIGO|
|4 to 15|Reserved||||
|16|Event Subscriber Port 0 (FSUB0)||||
|17|Event Subscriber Port 1 (FSUB1)||||
|18 to 31|Reserved||||



For more details, see the TIMx chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

## **8.29 Device Analog Connections** 

Figure 8-3 shows the internal analog connection of the device. 

## _**ADC**_ 

**==> picture [210 x 121] intentionally omitted <==**

**----- Start of picture text -----**<br>
A0_0:A0_9 0:9<br>Temp Sense 11<br>ADC0<br>A0_12:A0_14 12:14<br>Supply/Battery Monitor 31<br>28<br>VREF<br>**----- End of picture text -----**<br>


**Figure 8-3. Analog Connections** 

## **8.30 Input/Output Diagrams** 

The IOMUX manages the selection of which peripheral function is to be used on a digital IO and provides the controls for the output driver, input path, and the wake-up logic for wakeup from SHUTDOWN mode. For more information, see the IOMUX section of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

The mixed-signal IO pin slice diagram for a full featured IO pin is shown in Figure 8-4. Not all pins have analog functions, wake-up logic, drive strength control, and pullup or pulldown resistors available. See the device-specific data sheet for detailed information on what features are supported for a specific pin. 

Copyright © 2025 Texas Instruments Incorporated 

60 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

## **www.ti.com** 

**==> picture [498 x 484] intentionally omitted <==**

**----- Start of picture text -----**<br>
Wake to PMCU To analog peripheral(s)<br>SHUTDOWN Wakeup<br>I<br>I PC<br>WAKESTATE WAKESTATE I<br>)C<br>I I<br>WUEN WUEN — D Q |<br>S Q<br>| EN d R<br>I<br>WCOMP WCOMP D Q I<br>Glitch<br>| i EN } Filter | VDDIO VDDIO VDDIO<br>! =<br>Input Logic<br>HYSTEN HYSTEN P —+ G I 4 O<br>INV INV<br>5 Peripheral 01Unassigned Unassigned ave = PCPC 01 I INENA PMOS  [(1)] Clamping diode  [(1)]<br>a oot PC DIN 10 IO pin<br>== | Peripheral 15 PeripheralPeripheralO115 "=HH — 15 eSPNY O ! C CO<br>! \<br>! \<br>NMOS Clamping<br>PF diode<br>je<br>INV Output Logic SHUTDOWN d<br>| I Latches<br>Unassigned 0<br>Peripheral 01 1<br>D Q<br>Peripheral 15 : \ 15 EN 10 DOUT D Q 4<br>—Y RSTN je Hi EN Driver  VSS Y VSS VSS<br>Unassigned 0 Logic<br>Peripheral 01 | 1<br>H | O D Q 1S Hi-Z<br>D Q<br>EN<br>Peripheral 15 15<br>| IS | RSTN I EN<br>PF != 0 Vi || eee I<br>||<br>IORET l [I]<br>PC<br>Z1 I D Q<br>DRV EN<br>PIPU D Q<br>|<br>PIPD Hini EN ||<br>SHUTDOWN S \ D Q \<br>RELEASE A R Q Hee| EN | (1)output-high PMOS, pullup resistor, or clamping diode. The 5V-tolerant open drain IO type does not have the<br> (1)<br>PULLUP<br>R<br>Input Mux<br>PULLDOWN<br>R<br>(1)<br>Output Mux PMOS Control  NMOS Control<br>Hi-Z Output Mux<br> (1)<br>Drive strength<br>Pullup enable<br>Pulldown enable<br>**----- End of picture text -----**<br>


**Figure 8-4. Superset Input/Output Diagram** 

## **8.31 Serial Wire Debug Interface** 

A serial wire debug (SWD) two-wire interface is provided via an Arm compatible serial wire debug port (SW-DP) to enable access to multiple debug functions within the device. For a complete description of the debug functionality offered on MSPM0 devices, see the DEBUG chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

**Table 8-13. Serial Wire Debug Pin Requirements and Functions** 

|**DEVICE SIGNAL**|**DIRECTION**|**SWD FUNCTION**|
|---|---|---|
|SWCLK|Input|Serial wire clock from debug probe|
|SWDIO|Input/Output|Bi-directional (shared) serial wire data|



Copyright © 2025 Texas Instruments Incorporated 

61 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **8.32 Bootstrap Loader (BSL)** 

The bootstrap loader (BSL) enables configuration of the device as well as programming of the device memory through a UART or I2C serial interface. Access to the device memory and configuration through the BSL is protected by a 256-bit user-defined password, and it is possible to completely disable the BSL in the device configuration, if desired. The BSL is enabled by default from TI to support use of the BSL for production programming. 

A minimum of two pins are required to use the BSL: the BSLRX and BSLTX signals (for UART), or the BSLSCL and BSLSDA signals (for I[2] C). Additionally, one or two additional pins (BSL_invoke and NRST) may be used for controlled invocation of the bootloader by an external host. 

If enabled, the BSL may be invoked (started) in the following ways: 

- The BSL is invoked during the boot process if the BSL_invoke pin state matches the defined BSL_invoke logic level. If the device fast boot mode is enabled, this invocation check is skipped. An external host can force the device into the BSL by asserting the invoke condition and applying a reset pulse to the NRST pin to trigger a BOOTRST, after which the device will verify the invoke condition during the reboot process and start the BSL if the invoke condition matches the expected logic level. 

- The BSL is automatically invoked during the boot process if the reset vector and stack pointer are left unprogrammed. As a result, a blank device from TI will invoke the BSL during the boot process without any need to provide a hardware invoke condition on the BSL_invoke pin. This enables production programming using just the serial interface signals. 

- The BSL may be invoked at runtime from application software by issuing a SYSRST with BSL entry command. 

**Table 8-14. BSL Pin Requirements and Functions** 

|**DEVICE SIGNAL**|**CONNECTION**|**BSL FUNCTION**|
|---|---|---|
|BSLRX|Required for UART|UART receive signal (RXD), an input|
|BSLTX|Required for UART|UART transmit signal (TXD) an output|
|BSLSCL|Required for I2C|I2C BSL clock signal (SCL)|
|BSLSDA|Required for I2C|I2C BSL data signal (SDA)|
|BSL_invoke|Optional|Active-high digital input used to start the BSL<br>during boot|
|NRST|Optional|Active-low reset pin used to trigger a reset<br>and subsequent check of the invoke signal<br>(BSL_invoke)|



For a complete description of the BSL functionality and command set, see the MSPM0 boot strap loader user's guide. 

## **8.33 Device Factory Constants** 

All devices include a memory-mapped FACTORY region which provides read-only data describing the capabilities of a device as well as any factory-provided trim information for use by application software. Refer to the _Factory Constants_ chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

**Table 8-15. DEVICEID** 

DEVICEID address is 0x41C4.0004, PARTNUM is bit 12 to 27, MANUFACTURER is bit 1 to 11. 

|**DEVICE**|**PARTNUM**|**MANUFACTURER**|
|---|---|---|
|MSPM0L1116|0xBBB4|0x17|
|MSPM0L1117|0xBBB4|0x17|



Copyright © 2025 Texas Instruments Incorporated 

62 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

**Table 8-16. USERID** 

USERID address is 0x41C4.0008, PART is bit 0 to 15, VARIANT is bit 16 to 23 

|**Device**|**Part**|**Variant**|
|---|---|---|
|MSPM0L1116SRGER|0xE284|0x77|
|MSPM0L1116SRHBR|0xE284|0x78|
|MSPM0L1116SRGZR|0xE284|0x79|
|MSPM0L1116SPTR|0xE284|0x7A|
|MSPM0L1117SRGER|0xAF6C|0xB0|
|MSPM0L1117SRHBR|0xAF6C|0xB1|
|MSPM0L1117SRGZR|0xAF6C|0xB2|
|MSPM0L1117SPTR|0xAF6C|0xB3|



## **8.34 Identification** 

## **Revision and Device Identification** 

The hardware revision and device identification values are stored in the memory-mapped FACTORY region (see the Device Factory Constants section) which provides read-only data describing the capabilities of a device as well as any factory-provided trim information for use by application software. For more information, see the _Factory Constants_ chapter of the _MSPM0 L-Series 32MHz Microcontrollers Technical Reference Manual_ . 

The device revision and identification information are also included as part of the top-side marking on the device package. The device-specific errata describes these markings (see Section 10.3). 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 63 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **9 Applications, Implementation, and Layout** 

## **9.1 Typical Application** 

## **Note** 

Information in the following applications sections is not part of the TI component specification, and TI does not warrant its accuracy or completeness. TI’s customers are responsible for determining suitability of components for their purposes, as well as validating and testing their design implementation to confirm system functionality. 

## _**9.1.1 Schematic**_ 

TI recommends connecting a combination of a 10µF and a 0.1µF low-ESR ceramic decoupling capacitor across the VDD and VSS pins, as well as placing these capacitors as close as possible to the supply pins that they decouple (within a few millimeters) to achieve a minimal loop area. The 10µF bulk decoupling capacitor is a recommended value for most applications, but this capacitance may be adjusted if needed based upon the PCB design and application requirements. For example, larger bulk capacitors can be used, but this can affect the supply rail ramp-up time. 

The NRST reset pin must be pulled up to VDD (supply level) for the device to release from RESET state and start the boot process. TI recommends connecting an external 47kΩ pullup resistor with a 10nF pulldown capacitor for most applications, enabling the NRST pin to be controlled by another device or a debug probe. 

The SYSOSC frequency correction loop (FCL) circuit utilizes an external 100kΩ with 0.1% tolerance resistor with a temperature coefficient (TCR) of 25ppm/C or better populated between the ROSC pin and VSS. This resistor establishes a reference current to stabilize the SYSOSC frequency through a correction loop. This resistor is required if the FCL feature is used for higher accuracy, and it is not required if the SYSOSC FCL is not enabled. When the FCL mode is not used, the PA2 pin may be used as a digital input/output pin. 

A 0.47µF tank capacitor is required for the VCORE pin and must be placed close to the device with minimum distance to the device ground. Do not connect other circuits to the VCORE pin. 

For the 5V-tolerant open drain (ODIO), a pullup resistor tied to a voltage reference (e.g 3.3V supply rail) is required to output high as the open drain IO only implement a low-side NMOS driver and no high-side PMOS driver. 

**==> picture [403 x 206] intentionally omitted <==**

**----- Start of picture text -----**<br>
1.62 - 3.6 V 1.62 - 5.5 V<br>MSPM0 MCU ROSC<br>100 k �<br>±0.1% ±25ppm<br>VDD PA2/ROSC<br>10  � F 0.1  � F<br>VSS<br>VCORE<br>0.47  � F PA0<br>PA1<br>5V-tolerant open drain pins<br>Pull-up resistors are required for output high<br>NRST NRST<br>10nF The NRST pullup  SWDIO<br>resistor and capacitor Debug tool<br>are optional, but  SWCLK<br>NRST  must  be  Debug interface<br>pulled high to VDD<br>for the device to start.<br>PU PU<br>47 k<br>**----- End of picture text -----**<br>


**Figure 9-1. Basic Application Schematic** 

Copyright © 2025 Texas Instruments Incorporated 

64 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **10 Device and Documentation Support** 

TI offers an extensive line of development tools. Tools and software to evaluate the performance of the device, generate code, and develop solutions are listed below. 

## **10.1 Device Nomenclature** 

To designate the stages in the product development cycle, TI assigns prefixes to the part numbers of all MSP MCU devices and support tools. Each MSP MCU commercial family member has one of two prefixes: MSP or X. These prefixes represent evolutionary stages of product development from engineering prototypes (X) through fully qualified production devices (MSP). 

**X or XMS** – Experimental device that is not necessarily representative of the final device's electrical specifications 

**MSP** – Fully qualified production device 

**X and XMS** devices are shipped against the following disclaimer: 

"Developmental product is intended for internal evaluation purposes." MSP devices have been characterized fully, and the quality and reliability of the device have been demonstrated fully. TI's standard warranty applies. Predictions show that prototype devices (X) have a greater failure rate than the standard production devices. TI recommends that these devices not be used in any production system because their expected end-use failure rate still is undefined. Only qualified production devices are to be used. 

TI device nomenclature also includes a suffix with the device family name. This suffix indicates the temperature range, package type, and distribution format. Figure 10-1 provides a legend for reading the complete device name. 

**==> picture [136 x 9] intentionally omitted <==**

**----- Start of picture text -----**<br>
MSP  M0  L  111  7  S  PT  R<br>**----- End of picture text -----**<br>


**==> picture [252 x 49] intentionally omitted <==**

**----- Start of picture text -----**<br>
Processor Family Distribu�on Format<br>MCU Pla � orm Package Type<br>Product Family Temperature range<br>Device Subfamily Flash Memory<br>**----- End of picture text -----**<br>


**Figure 10-1. Device Nomenclature** 

**Table 10-1. Device Nomenclature** 

||**Table 10-1. Device Nomenclature**|
|---|---|
|**Processor Family**|MSP = Mixed-signal processor<br>X, XMS = Experimental silicon|
|**MCU Platform**|M0 = Arm-based 32-bit M0+|
|**Product Family**|L = 32MHz frequency|
|**Device Subfamily**|111x = ADC|
|**Internal Memory**|6 = 64KB flash, 16KB SRAM<br>7 = 128KB flash, 16KB SRAM|
|**Temperature Range**|S = –40°C to 125°C|
|**Package Type**|SeeSection 5and www.ti.com/packaging|
|**Distribution Format**|T = Small reel<br>R = Large reel<br>No marking = Tube or tray|



For orderable part numbers of MSP devices in different package types, see the Package Option Addendum of this document, ti.com, or contact your TI sales representative. 

Copyright © 2025 Texas Instruments Incorporated 

65 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **10.2 Tools and Software** 

## **Design Kits and Evaluation Modules** 

|MSPM0 LaunchPad|Empowers you to immediately start developing on the industry’s best integrated|
|---|---|
|Development Kit: LP-|analog and most cost-optimized general purpose MSPM0 MCU family. Exposes all|
|MSPM0L1117|device pins and functionality; includes some built-in circuitry, out-of-box software|
||demos, and on-board XDS110 debug probe for programming, debugging, and|
||EnergyTrace™technology.|
||The LaunchPad ecosystem includes dozens ofBoosterPack™stackable plug-in|
||modules to extend functionality.|



## **Embedded Software** 

|**Embedded Software**||
|---|---|
|MSPM0 Software<br>Development Kit (SDK)|Contains software drivers, middleware libraries, documentation, tools, and code<br>examples that create a familiar and easy user experience for all MSPM0 devices.|
|**Software Development**||
|**Tools**||
|TI Cloud Tools|Start your evaluation and development on a web browser without any installation.|
||Cloud tools also have a downloadable, offline version.|
|TI Resource Explorer|Online portal to TI SDKs. Accessible in CCS IDE or in TI Cloud Tools.|
|SysConfig|Intuitive GUI to configure device and peripherals, resolve system conflicts, generate|
||configuration code, and automate pin mux settings. Accessible in CCS IDE or in TI|
||Cloud Tools.(offline version)|
|MSP Academy|Great starting point for all developers to learn about the MSPM0 MCU Platform with|
||training modules that span a wide range of topics. Part of TIRex.|
|GUI Composer|GUIs that simplify evaluation of certain MSPM0 features, such as configuring and|
||monitoring a fully integrated analog signal chain without any code needed.|
|**IDE and compiler tool**||
|**chains**||
|Code Composer Studio™|Code Composer Studio is an integrated development environment (IDE) for TI's|
|(CCS)|microcontrollers and processors. It comprises a suite of tools used to develop and|
||debug embedded applications. CCS is completely free to use and is available on|
||Eclipse and Theia frameworks.|
|IAR Embedded|IAR Embedded Workbench for Arm delivers a complete development toolchain|
|Workbench® IDE|for building and debugging embedded applications for MSPM0.The included IAR|
||C/C++ Compiler generates highly optimized code for your application, and the|
||C-SPY Debugger is a fully integrated debugger for source and disassembly level|
||debugging with support for complex code and data breakpoint.|
|Keil® MDK IDE|Arm Keil MDK is a complete debugger and C/C++ compiler toolchain for building|
||and debugging embedded applications for MSPM0.Keil MDK includes a fully|
||integrated debugger for source and disassembly level debugging. MDK provides|
||full CMSIS compliance.|
|TI Arm-Clang|TI Arm Clang is included in the Code Composer Studio IDE.|
|GNU Arm Embedded Tool|The MSPM0 SDK supports development using the open-source Arm GNU|
|Chain|Toolchain. Arm GCC is supported by Code Composer Studio IDE (CCS).|



## **10.3 Documentation Support** 

To receive notification of documentation updates, navigate to the device product folder on ti.com. Click on _Notifications_ to register and receive a weekly digest of any product information that has changed. For change details, review the revision history included in any revised document. 

Copyright © 2025 Texas Instruments Incorporated 

66 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

The following documents describe the MSPM0 MCUs. Copies of these documents are available on the Internet at www.ti.com. 

## **Technical Reference Manual** 

MSPM0 L-Series This manual describes the modules and peripherals of the MSPM0L family of devices. 32MHz Microcontrollers Each description presents the module or peripheral in a general sense. Not all Technical Reference features and functions of all modules or peripherals are present on all devices. In Manual addition, modules or peripherals can differ in their exact implementation on different devices. Pin functions, internal signal connections, and operational parameters differ from device to device. See the device-specific data sheet for these details. 

## **10.4 Support Resources** 

TI E2E[™] support forums are an engineer's go-to source for fast, verified answers and design help — straight from the experts. Search existing answers or ask your own question to get the quick design help you need. 

Linked content is provided "AS IS" by the respective contributors. They do not constitute TI specifications and do not necessarily reflect TI's views; see TI's Terms of Use. 

## **10.5 Trademarks** 

LaunchPad[™] , Code Composer Studio[™] , TI E2E[™] , EnergyTrace[™] , and BoosterPack[™] are trademarks of Texas Instruments. 

Arm[®] and Cortex[®] are registered trademarks of Arm Limited. 

All trademarks are the property of their respective owners. 

## **10.6 Electrostatic Discharge Caution** 

This integrated circuit can be damaged by ESD. Texas Instruments recommends that all integrated circuits be handled with appropriate precautions. Failure to observe proper handling and installation procedures can cause damage. 

ESD damage can range from subtle performance degradation to complete device failure. Precision integrated circuits may be more susceptible to damage because very small parametric changes could cause the device not to meet its published specifications. 

## **10.7 Glossary** 

TI Glossary This glossary lists and explains terms, acronyms, and definitions. 

## **11 Revision History** 

NOTE: Page numbers for previous revisions may differ from page numbers in the current version. 

**Changes from December 1, 2024 to May 31, 2025 (from Revision * (December 2024) to Revision A (June 2025)) Page** 

|•|Corrected list of I/O features...............................................................................................................................1|
|---|---|
|•|Added note to indicate that this device is targeting PSA-L1 certification............................................................1|
|•|Updated Device Comparison section to tabulate communication peripheral instances.....................................7|
|•|Updated Pin Attributes section to correctly list IO types available in this device, and added footnote indicating|
||SDIO with wake functionality............................................................................................................................12|
|•|Corrected Pin Attributes tables to indicate that GPIO PA2 is controlled by PINCM7 instead of PINCM61......12|
|•|Corrected Pin Attributes tables to indicate that PA15 is a High-drive type IO (HDIO)......................................12|
|•|Corrected Pin Attributes tables to indicate that PB24 is a standard type IO (SDIO)........................................12|
|•|Updated Absolute Maximum Ratings for I_VDD and I_VSS to reflect correct junction temperatures and also|
||remove VDD>=2.7V condition..........................................................................................................................26|
|•|Added diode current rating for PB24 to Absolute Maximum Ratings...............................................................26|
|•|Added ambient temperature rating to Absolute Maximum Ratings..................................................................26|
|•|Added footnote to I_VDD and I_VSS guidelines for reduced current consumption when VDD supply voltage is|
||low (e.g. 1.62V)................................................................................................................................................26|



Copyright © 2025 Texas Instruments Incorporated 

67 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

|•|Updated Supply Current Characteristics to include maximum values and accurate typical values.................26|
|---|---|
|•|Added Supply Current Characteristics parameter for per-MHz SLEEP current (assessed at 32MHz)............26|
|•|Changed POR and BOR specifications to reflect accurate voltage thresholds for POR and coldboot BOR....26|
|•|Updated POR and BOR specifications section to remove footnote for dVDD/dt condition..............................26|
|•|Changed Flash Memory Characteristics to allow for users to designate any 32kB sectors of flash memory to|
||apply 100k cycles, rather than only the lower 32kB sectors.............................................................................26|
|•|Updated Timing Characteristics section with accurate specification values and corrected spec label for|
||Wakeup time from STOP1 and STOP2 to RUN...............................................................................................26|
|•|Updated System Oscillator specifications with accurate values.......................................................................26|
|•|Removed SYSOSC Typical Frequency Accuracy Figure.................................................................................26|
|•|Removed LFXT specification for VDD power supply range, as specifications were already applicable for|
||entire VDD operating range of the device........................................................................................................26|
|•|Changed LFXT start-up time to indicate typical value of 1 second..................................................................26|
|•|Updated Digital IO electrical characteristics to reflect ambient temperature conditions...................................26|
|•|Added rise/fall time specifications to Digital IO switching characteristics.........................................................26|
|•|Added footnote for HDIO DRV=1 condition to limit signal slew rate for high current operation.......................26|
|•|Updated ADC specifications for SNR and PSRRDC to remove minimum values............................................26|
|•|Updated Temperature Sensor coefficient specification values.........................................................................26|
|•|Changed VREF electrical characteristics for I_VREF, TC_VREF, and PSRRDC specification values............26|
|•|Removed Vnoise specifications from VREF electrical characteristics..............................................................26|
|•|Updated SPI specifications to reflect corrected setup and hold timing values.................................................26|
|•|Updated CPU description section to indicate support for memory protection unit (MPU)................................44|
|•|Updated Supported Functionality by Operating Mode table for accuracy and organization.............................45|
|•|Added detailed DMA Features table to DMA section.......................................................................................47|
|•|Updated Flash Memory section to indicate that any 32kB sectors can be selected for high-endurance|
||operation...........................................................................................................................................................51|
|•|Updated description in SRAM section regarding write-execute user operation...............................................51|
|•|Updated GPIO section to clarify that there are two GPIO ports in this device (PAx and PBx).........................51|
|•|Updated Temperature Sensor section to add details for per-unit TSc calculation method using V_TRIM_0K.53|
|•|Updated VREF section to clarify that the VREF in this device does not require a decoupling capacitor on|
||VREF+/- pins for proper operation, and also added a block diagram detailing VREF configurations. ............54|
|•|Updated Security section to list all security features present in this device.....................................................54|
|•|Updated SPI section to reference MCLK rather than ULPCLK........................................................................57|
|•|Updated LFSS section to indicate presence of LFSS_B and RTC_B variants.................................................57|
|•|Updated Timers section to correctly detail capabilities of various timer instance types...................................59|
|•|Updated Mechanical, Packaging, and Orderable Information section to append drawings for each package|
||variant...............................................................................................................................................................69|



Copyright © 2025 Texas Instruments Incorporated 

68 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **12 Mechanical, Packaging, and Orderable Information** 

The following pages include mechanical, packaging, and orderable information. This information is the most current data available for the designated devices. This data is subject to change without notice and revision of this document. For browser-based versions of this data sheet, refer to the left-hand navigation. 

Copyright © 2025 Texas Instruments Incorporated 

69 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

**==> picture [132 x 11] intentionally omitted <==**

**----- Start of picture text -----**<br>
PACKAGE OUTLINE<br>**----- End of picture text -----**<br>


**==> picture [458 x 537] intentionally omitted <==**

**----- Start of picture text -----**<br>
SS RGE0024B SCALE  3.000 VQFN - 1 mm max height<br>PLASTIC QUAD FLATPACK - NO LEAD<br>A 4.1 B<br>3.9<br>SNS .<br>NAANCSSRASSSAAS<br>SSSSESSSAAS 0.5<br>Re WES<br>SSAASAASASS 0.3<br>SSSSSAAMMM SSFP STTyy<br>PIN 1 INDEX AREA SARAENNSSAASSS<br>RS SSS<br>RARRAAAST S Rnn000rs 8:9 4.1<br>3.9<br>0.3<br>0.2<br>DETAIL<br>OPTIONAL TERMINAL<br>TYPICAL<br>C<br>1 MAX<br>{ t o pf Jf -} fF | ] —f - JF} FY SEATING PLANE<br>0.05<br>0.00 0.08 C<br>2X 2.5<br>2.45 0.1 (0.2) TYP<br>7 | Oo + \ 12 7<br>E XPOSED<br>SEE TERMINAL<br>THERMAL PAD<br>DETAIL<br>6 13<br>i , |) C | |<br>2X 25 SYMM<br>2.5 -— +— - -— ¢<br>a, C | |<br>1 18<br>20X 0.5 24X  [0.3]<br>0.2<br>24 SYMM 19 0.1 C A B<br>PIN 1 ID<br>0.05<br>(OPTIONAL)<br>24X  [0.5]<br>0.3<br>4219013/A   05/2017<br>**----- End of picture text -----**<br>


NOTES: 

1. All linear dimensions are in millimeters. Any dimensions in parenthesis are for reference only. Dimensioning and tolerancing per ASME Y14.5M. 

2. This drawing is subject to change without notice. 

3. The package thermal pad must be soldered to the printed circuit board for thermal and mechanical performance. 

Copyright © 2025 Texas Instruments Incorporated 

70 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **EXAMPLE BOARD LAYOUT** 

**==> picture [462 x 554] intentionally omitted <==**

**----- Start of picture text -----**<br>
RGE0024B VQFN - 1 mm max height<br>PLASTIC QUAD FLATPACK - NO LEAD<br>( 2.45)<br>SYMM<br>24 19<br>24X (0.6)<br>1<br>18<br>24X (0.25)<br>(R0.05)<br>TYP 25 SYMM<br>(3.8)<br>20X (0.5)<br>13<br>6<br>( 0.2) TYP<br>VIA<br>7 12<br>(0.975) TYP<br>(3.8)<br>LAND PATTERN EXAMPLE<br>EXPOSED METAL SHOWN<br>SCALE:15X<br>0.07 MAX 0.07 MIN<br>ALL AROUND ALL AROUND<br>SOLDER MASK<br>METAL<br>OPENING<br>EXPOSEDMETAL SOLDER MASKOPENING EXPOSEDMETAL METAL UNDERSOLDER MASK<br>NON SOLDER MASK<br>SOLDER MASK<br>DEFINED<br>DEFINED<br>(PREFERRED)<br>SOLDER MASK DETAILS<br>4219013/A   05/2017<br>**----- End of picture text -----**<br>


NOTES: (continued) 

4. This package is designed to be soldered to a thermal pad on the board. For more information, see Texas Instruments literature number SLUA271 (www.ti.com/lit/slua271). 

5. Vias are optional depending on application, refer to device data sheet. If any vias are implemented, refer to their locations shown on this view. It is recommended that vias under paste be filled, plugged or tented. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 71 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **EXAMPLE STENCIL DESIGN** 

## **RGE0024B** 

**VQFN - 1 mm max height** 

PLASTIC QUAD FLATPACK - NO LEAD 

**==> picture [74 x 538] intentionally omitted <==**

**----- Start of picture text -----**<br>
4219013/A   05/2017<br>**----- End of picture text -----**<br>


**==> picture [346 x 372] intentionally omitted <==**

**----- Start of picture text -----**<br>
4X ( 1.08)<br>(0.64) TYP<br>24 19<br>24X (0.6)<br>1<br>25<br>18<br>24X (0.25)<br>(R0.05) TYP (0.64)<br>TYP<br>SYMM<br>(3.8)<br>20X (0.5)<br>13<br>6<br>METAL<br>TYP<br>7 12<br>SYMM<br>(3.8)<br>SOLDER PASTE EXAMPLE<br>BASED ON 0.125 mm THICK STENCIL<br>EXPOSED PAD 25<br>78% PRINTED SOLDER COVERAGE BY AREA UNDER PACKAGE<br>SCALE:20X<br>**----- End of picture text -----**<br>


NOTES: (continued) 

6. Laser cutting apertures with trapezoidal walls and rounded corners may offer better paste release. IPC-7525 may have alternate design recommendations. 

Copyright © 2025 Texas Instruments Incorporated 

72 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 EE e ~~d~~ 

**www.ti.com** 

**==> picture [131 x 11] intentionally omitted <==**

**----- Start of picture text -----**<br>
PACKAGE OUTLINE<br>**----- End of picture text -----**<br>


**==> picture [456 x 536] intentionally omitted <==**

**----- Start of picture text -----**<br>
RHB0032E VQFN - 1 mm max height<br>ee SCALE  3.000<br>PLASTIC QUAD FLATPACK - NO LEAD<br>A 5.1 B<br>4.9<br>SSAARNASNANA SSS<br>RASA<br>SAAS<br>RSSAAASRASS<br>RASSRRAMWMMMIMMWWK SF FF SSS SHS SLT T<br>ASMMMMM™MNANI HHI ST STF SSS LSS<br>SSASAPMOONYSEARS LITEFFF TSS TLSSST<br>PIN 1 INDEX AREA KAARASSRRRRRRARRRRRAAVVIASSASSINRRAnn“ “anneTEI FES8 5555FS SS<br>RRRRRR’ASRSS A SSF 5.1 (0.1)<br>4.9<br>OPTIONAL METAL THICKNESSSIDE WALL DETAIL20.000<br>C<br>1 MAX<br>{ + FEF +f 4] —f JF fe FF ceoee { SEATING PLANE<br>0.05<br>0.00 0.08 C<br>2X 3.5<br>\ O 3.45 + 0.1 | - (0.2) TYP<br>9 16 E XPOSED<br>THERMAL PAD<br>28X 0.5<br>8 \<br>17 SEE SIDE WALL<br>DETAIL<br>oN I<br>a, “e y<br>|| C | | / \ Ne Uo ?,\I<br>2X 33 SYMM<br>3.5 -—-—+—-—- ¢ |<br>|) C | | |<br>32X  [0.3]<br>0.2 |<br>24 0.1 C A B<br>1<br>0.05 C<br>32 25<br>PIN 1 ID SYMM<br>(OPTIONAL) 32X  [0.5]<br>0.3<br>4223442/B   08/2019<br>**----- End of picture text -----**<br>


NOTES: 

1. All linear dimensions are in millimeters. Any dimensions in parenthesis are for reference only. Dimensioning and tolerancing per ASME Y14.5M. 

2. This drawing is subject to change without notice. 

3. The package thermal pad must be soldered to the printed circuit board for thermal and mechanical performance. 

Copyright © 2025 Texas Instruments Incorporated 

73 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **EXAMPLE BOARD LAYOUT** 

**==> picture [462 x 554] intentionally omitted <==**

**----- Start of picture text -----**<br>
RHB0032E VQFN - 1 mm max height<br>PLASTIC QUAD FLATPACK - NO LEAD<br>( 3.45)<br>SYMM<br>32 25<br>32X (0.6)<br>1 24<br>32X (0.25)<br>(1.475)<br>28X (0.5)<br>33 SYMM<br>(4.8)<br>( 0.2) TYP<br>VIA<br>8 17<br>(R0.05)<br>TYP<br>9 16<br>(1.475)<br>(4.8)<br>LAND PATTERN EXAMPLE<br>SCALE:18X<br>0.07 MAX 0.07 MIN<br>ALL AROUND ALL AROUND<br>SOLDER MASK<br>METAL<br>OPENING<br>S OLDER MASK M ETAL UNDER<br>OPENING SOLDER MASK<br>NON SOLDER MASK<br>SOLDER MASK<br>DEFINED<br>DEFINED<br>(PREFERRED)<br>SOLDER MASK DETAILS<br>4223442/B   08/2019<br>**----- End of picture text -----**<br>


**==> picture [65 x 8] intentionally omitted <==**

**----- Start of picture text -----**<br>
NOTES: (continued)<br>**----- End of picture text -----**<br>


4. This package is designed to be soldered to a thermal pad on the board. For more information, see Texas Instruments literature number SLUA271 (www.ti.com/lit/slua271). 

5. Vias are optional depending on application, refer to device data sheet. If any vias are implemented, refer to their locations shown on this view. It is recommended that vias under paste be filled, plugged or tented. 

Copyright © 2025 Texas Instruments Incorporated 

74 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

## **EXAMPLE STENCIL DESIGN** 

## **RHB0032E** 

**VQFN - 1 mm max height** 

PLASTIC QUAD FLATPACK - NO LEAD 

**==> picture [441 x 537] intentionally omitted <==**

**----- Start of picture text -----**<br>
4X ( 1.49)<br>(R0.05) TYP (0.845)<br>32 25<br>32X (0.6)<br>1 24<br>32X (0.25)<br>28X (0.5)<br>(0.845)<br>SYMM<br>33<br>(4.8)<br>8 17<br>METAL<br>TYP<br>9 16<br>SYMM<br>(4.8)<br>SOLDER PASTE EXAMPLE<br>BASED ON 0.125 mm THICK STENCIL<br>EXPOSED PAD 33:<br>75% PRINTED SOLDER COVERAGE BY AREA UNDER PACKAGE<br>SCALE:20X<br>4223442/B   08/2019<br>**----- End of picture text -----**<br>


**==> picture [70 x 8] intentionally omitted <==**

**----- Start of picture text -----**<br>
NOTES: (continued)<br>**----- End of picture text -----**<br>


6. Laser cutting apertures with trapezoidal walls and rounded corners may offer better paste release. IPC-7525 may have alternate design recommendations. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 75 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

Copyright © 2025 Texas Instruments Incorporated 

76 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 77 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

Copyright © 2025 Texas Instruments Incorporated 

78 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

Copyright © 2025 Texas Instruments Incorporated 

79 

_Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

Copyright © 2025 Texas Instruments Incorporated 

80 _Submit Document Feedback_ 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

**MSPM0L1117, MSPM0L1116** SLASFC9A – DECEMBER 2024 – REVISED JUNE 2025 

**www.ti.com** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 81 

Product Folder Links: _MSPM0L1117 MSPM0L1116_ 

## **PACKAGE OPTION ADDENDUM** 

www.ti.com 

11-Feb-2026 

## **PACKAGING INFORMATION** 

|**Orderable part number**|**Status**|**Material type**|**Package | Pins**|**Package qty | Carrier**|**RoHS**|**Lead finish/**|**MSL rating/**|**Op temp (°C)**|**Part marking**|
|---|---|---|---|---|---|---|---|---|---|
||(1)|(2)|||(3)|**Ball material**|**Peak reflow**||(6)|
|||||||(4)|(5)|||
|MSPM0L1116SPTR|Active|Production|LQFP(PT) |48|1000|LARGE T&R|Yes|NIPDAU|Level-2-260C-1 YEAR|-40 to 125|M0L1116S|
|MSPM0L1116SRGER|Active|Production|VQFN (RGE) | 24|5000 | LARGE T&R|Yes|NIPDAU|Level-1-260C-UNLIM|-40 to 125|MSPM0|
||||||||||L1116S|
|MSPM0L1116SRGZR|Active|Production|VQFN (RGZ) | 48|4000 | LARGE T&R|Yes|NIPDAU|Level-2-260C-1 YEAR|-40 to 125|MSPM0|
||||||||||L1116S|
|MSPM0L1116SRHBR|Active|Production|VQFN (RHB) | 32|5000 | LARGE T&R|Yes|NIPDAU|Level-2-260C-1 YEAR|-40 to 125|MSPM0|
||||||||||L1116S|
|MSPM0L1117SPTR|Active|Production|LQFP(PT) |48|1000|LARGE T&R|Yes|NIPDAU|Level-2-260C-1 YEAR|-40 to 125|M0L1117S|
|MSPM0L1117SRGER|Active|Production|VQFN (RGE) | 24|5000 | LARGE T&R|Yes|NIPDAU|Level-1-260C-UNLIM|-40 to 125|MSPM0|
||||||||||L1117S|
|MSPM0L1117SRGZR|Active|Production|VQFN (RGZ) | 48|4000 | LARGE T&R|Yes|NIPDAU|Level-2-260C-1 YEAR|-40 to 125|MSPM0|
||||||||||L1117S|
|MSPM0L1117SRHBR|Active|Production|VQFN (RHB) | 32|5000 | LARGE T&R|Yes|NIPDAU|Level-2-260C-1 YEAR|-40 to 125|MSPM0|
||||||||||L1117S|
|XMSM0L1117SPTR|Active|Preproduction|LQFP(PT) |48|1000|LARGE T&R|-|Call TI|Call TI|-40 to 125||
|XMSM0L1117SPTR.A|Active|Preproduction|LQFP(PT) |48|1000|LARGE T&R|-|Call TI|Call TI|-40 to 125||
|XMSM0L1117SRGER|Active|Preproduction|VQFN(RGE) |24|5000|LARGE T&R|-|Call TI|Call TI|-40 to 125||
|XMSM0L1117SRGER.A|Active|Preproduction|VQFN(RGE) |24|5000|LARGE T&R|-|Call TI|Call TI|-40 to 125||
|XMSM0L1117SRGZR|Active|Preproduction|VQFN(RGZ) |48|3000|LARGE T&R|-|Call TI|Call TI|-40 to 125||
|XMSM0L1117SRGZR.A|Active|Preproduction|VQFN(RGZ) |48|3000|LARGE T&R|-|Call TI|Call TI|-40 to 125||
|XMSM0L1117SRHBR|Active|Preproduction|VQFN(RHB) |32|5000|LARGE T&R|-|Call TI|Call TI|-40 to 125||
|XMSM0L1117SRHBR.A|Active|Preproduction|VQFN (RHB) | 32|5000 | LARGE T&R|-|Call TI|Call TI|-40 to 125||



> **(1) Status:** For more details on status, see our product life cycle. 

> **(2) Material type:** When designated, preproduction parts are prototypes/experimental devices, and are not yet approved or released for full production. Testing and final process, including without limitation quality assurance, reliability performance testing, and/or process qualification, may not yet be complete, and this item is subject to further changes or possible discontinuation. If available for ordering, purchases will be subject to an additional waiver at checkout, and are intended for early internal evaluation purposes only. These items are sold without warranties of any kind. 

> **(3) RoHS values:** Yes, No, RoHS Exempt. See the TI RoHS Statement for additional information and value definition. 

Addendum-Page 1 

**PACKAGE OPTION ADDENDUM** 

11-Feb-2026 

www.ti.com 

> **(4) Lead finish/Ball material:** Parts may have multiple material finish options. Finish options are separated by a vertical ruled line. Lead finish/Ball material values may wrap to two lines if the finish value exceeds the maximum column width. 

> **(5) MSL rating/Peak reflow:** The moisture sensitivity level ratings and peak solder (reflow) temperatures. In the event that a part has multiple moisture sensitivity ratings, only the lowest level per JEDEC standards is shown. Refer to the shipping label for the actual reflow temperature that will be used to mount the part to the printed circuit board. 

> **(6) Part marking:** There may be an additional marking, which relates to the logo, the lot trace code information, or the environmental category of the part. 

Multiple part markings will be inside parentheses. Only one part marking contained in parentheses and separated by a "~" will appear on a part. If a line is indented then it is a continuation of the previous line and the two combined represent the entire part marking for that device. 

**Important Information and Disclaimer:** The information provided on this page represents TI's knowledge and belief as of the date that it is provided. TI bases its knowledge and belief on information provided by third parties, and makes no representation or warranty as to the accuracy of such information. Efforts are underway to better integrate information from third parties. TI has taken and continues to take reasonable steps to provide representative and accurate information but may not have conducted destructive testing or chemical analysis on incoming materials and chemicals. TI and TI suppliers consider certain information to be proprietary, and thus CAS numbers and other limited information may not be available for release. 

In no event shall TI's liability arising out of such information exceed the total purchase price of the TI part(s) at issue in this document sold by TI to Customer on an annual basis. 

Addendum-Page 2 

www.ti.com 

6-Nov-2025 

## **PACKAGE MATERIALS INFORMATION** 

## **TAPE AND REEL INFORMATION** 

**==> picture [448 x 166] intentionally omitted <==**

**----- Start of picture text -----**<br>
REEL DIMENSIONS TAPE DIMENSIONS<br>K0  P1<br>W<br>B0<br>Reel<br>Diameter<br>Cavity A0<br>A0 Dimension designed to accommodate the component width<br>B0 Dimension designed to accommodate the component length<br>K0 Dimension designed to accommodate the component thickness<br>W Overall width of the carrier tape<br>P1 Pitch between successive cavity centers<br>**----- End of picture text -----**<br>


**==> picture [68 x 9] intentionally omitted <==**

**----- Start of picture text -----**<br>
Reel Width (W1)<br>**----- End of picture text -----**<br>


**==> picture [274 x 111] intentionally omitted <==**

**----- Start of picture text -----**<br>
QUADRANT ASSIGNMENTS FOR PIN 1 ORIENTATION IN TAPE<br>Sprocket Holes<br>Q1 Q2 Q1 Q2<br>Q3 Q4 Q3 Q4 User Direction of Feed<br>Pocket Quadrants<br>**----- End of picture text -----**<br>


*All dimensions are nominal 

|*All dimensions are nominal|||||||||||||
|---|---|---|---|---|---|---|---|---|---|---|---|---|
|**Device**|**Package**<br>**Type**|**Package**<br>**Drawing**|**Pins**|**SPQ**|**Reel**<br>**Diameter**<br>**(mm)**|**Reel**<br>**Width**<br>**W1 (mm)**|**A0**<br>**(mm)**|**B0**<br>**(mm)**|**K0**<br>**(mm)**|**P1**<br>**(mm)**|**W**<br>**(mm)**|**Pin1**<br>**Quadrant**|
|MSPM0L1116SPTR|LQFP|PT|48|1000|330.0|16.4|9.6|9.6|1.9|12.0|16.0|Q2|
|MSPM0L1116SRGER|VQFN|RGE|24|5000|330.0|12.4|4.25|4.25|1.15|8.0|12.0|Q2|
|MSPM0L1116SRGZR|VQFN|RGZ|48|4000|330.0|16.4|7.3|7.3|1.1|12.0|16.0|Q2|
|MSPM0L1116SRHBR|VQFN|RHB|32|5000|330.0|12.4|5.3|5.3|1.1|8.0|12.0|Q2|
|MSPM0L1117SPTR|LQFP|PT|48|1000|330.0|16.4|9.6|9.6|1.9|12.0|16.0|Q2|
|MSPM0L1117SRGER|VQFN|RGE|24|5000|330.0|12.4|4.25|4.25|1.15|8.0|12.0|Q2|
|MSPM0L1117SRGZR|VQFN|RGZ|48|4000|330.0|16.4|7.3|7.3|1.1|12.0|16.0|Q2|
|MSPM0L1117SRHBR|VQFN|RHB|32|5000|330.0|12.4|5.3|5.3|1.1|8.0|12.0|Q2|



Pack Materials-Page 1 

www.ti.com 

6-Nov-2025 

## **PACKAGE MATERIALS INFORMATION** 

## **TAPE AND REEL BOX DIMENSIONS** 

**==> picture [237 x 117] intentionally omitted <==**

**----- Start of picture text -----**<br>
Width (mm) H<br>W L<br>**----- End of picture text -----**<br>


*All dimensions are nominal 

|*All dimensions are nominal||||||||
|---|---|---|---|---|---|---|---|
|**Device**|**Package Type**|**Package Drawing**|**Pins**|**SPQ**|**Length (mm)**|**Width (mm)**|**Height (mm)**|
|MSPM0L1116SPTR|LQFP|PT|48|1000|336.6|336.6|31.8|
|MSPM0L1116SRGER|VQFN|RGE|24|5000|367.0|367.0|35.0|
|MSPM0L1116SRGZR|VQFN|RGZ|48|4000|360.0|360.0|36.0|
|MSPM0L1116SRHBR|VQFN|RHB|32|5000|367.0|367.0|35.0|
|MSPM0L1117SPTR|LQFP|PT|48|1000|336.6|336.6|31.8|
|MSPM0L1117SRGER|VQFN|RGE|24|5000|367.0|367.0|35.0|
|MSPM0L1117SRGZR|VQFN|RGZ|48|4000|367.0|367.0|35.0|
|MSPM0L1117SRHBR|VQFN|RHB|32|5000|367.0|367.0|35.0|



Pack Materials-Page 2 

## **IMPORTANT NOTICE AND DISCLAIMER** 

TI PROVIDES TECHNICAL AND RELIABILITY DATA (INCLUDING DATASHEETS), DESIGN RESOURCES (INCLUDING REFERENCE DESIGNS), APPLICATION OR OTHER DESIGN ADVICE, WEB TOOLS, SAFETY INFORMATION, AND OTHER RESOURCES “AS IS” AND WITH ALL FAULTS, AND DISCLAIMS ALL WARRANTIES, EXPRESS AND IMPLIED, INCLUDING WITHOUT LIMITATION ANY IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE OR NON-INFRINGEMENT OF THIRD PARTY INTELLECTUAL PROPERTY RIGHTS. 

These resources are intended for skilled developers designing with TI products. You are solely responsible for (1) selecting the appropriate TI products for your application, (2) designing, validating and testing your application, and (3) ensuring your application meets applicable standards, and any other safety, security, regulatory or other requirements. 

These resources are subject to change without notice. TI grants you permission to use these resources only for development of an application that uses the TI products described in the resource. Other reproduction and display of these resources is prohibited. No license is granted to any other TI intellectual property right or to any third party intellectual property right. TI disclaims responsibility for, and you fully indemnify TI and its representatives against any claims, damages, costs, losses, and liabilities arising out of your use of these resources. 

TI’s products are provided subject to TI’s Terms of Sale, TI’s General Quality Guidelines, or other applicable terms available either on ti.com or provided in conjunction with such TI products. TI’s provision of these resources does not expand or otherwise alter TI’s applicable warranties or warranty disclaimers for TI products. Unless TI explicitly designates a product as custom or customer-specified, TI products are standard, catalog, general purpose devices. 

TI objects to and rejects any additional or different terms you may propose. 

Copyright © 2026, Texas Instruments Incorporated 

Last updated 10/2025 


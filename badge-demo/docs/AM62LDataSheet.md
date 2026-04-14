**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62Lx Sitara™ Processors** 

## **1 Features** 

## **Processor Cores:** 

- Dual 64-bit Arm[®] Cortex[®] -A53 microprocessor subsystem up to 1.25GHz 

   - Dual-core Cortex-A53 with 256KB L2 Cache 

   - Each A53 core has 32KB L1 DCache and 32KB L1 ICache 

## **Memory Subsystem:** 

- 160KB of Shared On-Chip SRAM (OCSRAM) 

- DDR Subsystem (DDRSS) 

      - Supports LPDDR4, DDR4 memory types 

   - 

   - 16-bit data bus 

   - Supports speeds up to 1600MT/s 

   - Max DDR4 addressing range of 4GB 

   - Max LPDDR4 addressing range of 2GB 

## **Multimedia:** 

- Display Subsystem 

      - Single display support 

   - 

   - Up to 1920x1080 @ 60fps 

      - Supported with independent PLL 

   - 

   - MIPI[®] DSI (4 lanes DPHY) or DPI (24-bit RGB LVCMOS) 

## **Security:** 

- Secure boot supported 

      - Hardware-enforced Root-of-Trust (RoT) 

   - 

      - Support to switch RoT via backup key 

   - 

   - Support for takeover protection, IP protection, and anti-roll back protection 

- Trusted Execution Environment (TEE) supported 

      - Arm TrustZone[®] based TEE 

   - 

      - Extensive firewall support for isolation 

   - 

   - Secure watchdog/timer/IPC 

   - Secure storage support 

   - Replay Protected Memory Block (RPMB) support 

- Dedicated Security Controller and dedicated security DMA & IPC subsystem for isolated processing 

- Cryptographic acceleration supported 

   - Session-aware cryptographic engine with ability to auto-switch key-material based on incoming data stream 

   - Supports cryptographic cores 

      - AES – 128-/192-/256-bit key sizes 

      - SHA2 – 224-/256-/384-/512-bit key sizes 

      - • DRBG with true random number generator 

      - PKA (Public Key Accelerator) to Assist in RSA/ECC processing for secure boot 

- Debugging security 

   - Secure software controlled debug access 

   - – Security aware debugging 

## **High Speed Interfaces:** 

- Integrated Ethernet switch supporting (total 2 external ports) 

   - RMII(10/100) or RGMII (10/100/1000) 

   - – IEEE1588 (Annex D, Annex E, Annex F with 802.1AS PTP) 

   - Clause 45 MDIO PHY management 

   - Priority based flow control 

   - Packet Classifier based on ALE engine with 64 classifiers 

   - Time sensitive networking (TSN) support 

   - H/W interrupt Pacing 

   - IP/UDP/TCP checksum offload in hardware 

- 2× USB 2.0 Dual-Role Device (DRD) Subsystems (USBSS) 

   - Port configurable as USB host or USB device 

   - – USB device: High-speed (480Mbps), and Fullspeed (12Mbps) 

   - USB host: High-speed (480Mbps), Full-speed (12Mbps), and Low-speed (1.5Mbps) 

   - – xHCI 1.1 compatible 

## **General Connectivity:** 

- 8x Universal Asynchronous Receiver-Transmitters (UARTs) 

   - All instances Support RTS and CTS Flow Control 

   - Supports RS-485 external transceiver auto flow control 

- 4x Serial Peripheral Interface (SPI) controllers 

- 5x Inter-Integrated Circuit (I2C) ports 

- 3x Multichannel Audio Serial Ports (McASPs) 

   - Transmit and Receive clocks up to 50MHz 

   - – Up to 4/6/16 Serial Data Pins across 3x McASPs with Independent TX and RX Clocks 

   - Supports Time Division Multiplexing (TDM), Inter-IC Sound (I2S), and similar formats 

   - Supports Digital Audio Interface Transmission (SPDIF, IEC60958-1, and AES-3 formats) 

   - FIFO buffers for Transmit and Receive (256Bytes) 

   - Support for audio reference output clock 

- 3x enhanced PWM modules (ePWM) 

- 3x enhanced Quadrature Encoder Pulse modules (eQEP) 

- • 3x enhanced Capture modules (eCAP) 

An IMPORTANT NOTICE at the end of this data sheet addresses availability, warranty, changes, use in safety-critical applications, intellectual property matters and other important disclaimers. PRODUCTION DATA. 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

- General-Purpose I/O (GPIO), most LVCMOS I/O can be configured as GPIO 

   - 4 banks supported for dual-voltage (1.8V/3.3V) and the rest single-voltage (1.8V) LVCMOS I/O banks 

- 3x Controller Area Network (CAN) with optional CAN-FD support 

   - Conforms with CAN Protocol 2.0 A, B, and ISO 11898-1 

   - Full CAN-FD support (up to 64 data bytes) 

   - Speed up to 8Mbps 

- 1x 12-bit Analog-to-Digital Converter (ADC) 

   - 10 bits of effective resolution (ENOB ≅ 10) 

   - Up to 2MSPS 

   - 4x analog inputs (time-multiplexed) 

## **Media and Data Storage:** 

- 3x Multi-Media Card/Secure Digital[®] (MMC/SD[®] ) interface 

   - 1x 8-bit eMMC interface up to HS200 speed 

   - 2x 4-bit SD/SDIO interface up to UHS-I 

   - Compliant with eMMC 5.1, SD 3.01, and SDIO Version 3.0 

- 1× General-Purpose Memory Controller (GPMC) up to 133MHz 

## **Power Management:** 

- Active power management features such as auto clock gating, power gating, and dynamic frequency scaling 

- Several low-power features supported 

- Low-Power Modes 

      - RTC Only 

   - 

   - RTC + IO + DDR 

- 

- DeepSleep 

- Standby 

## **Boot Options:** 

- UART 

- OSPI/QSPI Flash 

- GPMC NAND Flash 

- SD Card 

- eMMC 

- USB (host) Mass storage 

- USB (device) boot from external host (DFU mode) 

## **Technology / Package:** 

      - 16-nm Technology 

      - 11.9mm × 11.9mm, 0.5mm VCA, 373-pin FCCSP BGA package (ANB) 

   - Flexible 8- and 16-Bit Synchronous or Asynchronous Memory Interfaces with up to four Chip Selects 

   - Supports 16-bit Muxed Address/Data schemes (AD, AAD) 

   - Uses BCH code to support 4-, 8-, or 16-bit ECC 

   - Uses Hamming code to support 1-bit ECC 

   - Error Locator Module (ELM) 

- OSPI/QSPI with DDR / SDR support 

   - Support for Serial NAND and Serial NOR Flash devices 

   - 4GBytes memory address support 

Copyright © 2025 Texas Instruments Incorporated 

2 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **2 Applications** 

- Human Machine Interface (HMI) 

- Medical - patient monitoring 

- Building automation 

- EV charging stations 

- Solar energy 

- Energy Infrastructure (Smart Meter & Solar Gateway) 

- Mobile/Industrial printers 

## **3 Description** 

The low-cost & performance optimized AM62L family of application processors are built for Linux[®] application development. With scalable Arm[®] Cortex[®] -A53 core performance and embedded features such as: Multimedia DSI/DPI support, integrated ADC on chip, advanced lower power management modes, and extensive security options for IP protection with the built-in security features. 

The AM62Lx devices includes an extensive set of peripherals that make it a well-suited general-purpose device for a broad range of industrial applications while offering intelligent features and optimized power architecture as well. In addition, the extensive set of peripherals included in AM62Lx enables system-level connectivity, such as: USB, MMC/SD, OSPI, CAN-FD and an ADC. 

## **Package Information** 

|**PART NUMBER**|**PACKAGE**(1)|**PACKAGE SIZE**(2)|
|---|---|---|
|AM62Lx|ANB (FCCSP BGA, 373)|11.9mm × 11.9mm|



(1) For more information, see the Mechanical, Packaging, and Orderable Information section. 

(2) The package size (length × width) is a nominal value and includes pins, where applicable. 

Copyright © 2025 Texas Instruments Incorporated 

3 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **3.1 Functional Block Diagram** 

Figure 3-1 is functional block diagram for the superset device. 

## **Note** 

To understand what device features are currently supported by TI Software Development Kits (SDKs), search for the _AM62L Software Build Sheet_ located in the Downloads tab option provided at **AM62LProcessor-SDK** . 

**==> picture [376 x 398] intentionally omitted <==**

**----- Start of picture text -----**<br>
AM62Lx<br>Application Cores<br>Arm® Arm®<br>Cortex®-A53 Cortex®-A53<br>256KB L2 Cache<br>General Connectivity and IO System Memory<br>2-port Gb Ethernet w/1588 DDR4/LPDDR4 (16b)<br>4x SPI GPIO GPMC / ELM<br>8x UART 3x ePWM<br>160KB SRAM 3x MMC SD<br>3x CAN-FD 3x eCAP<br>Multimedia<br>5x I2C 3x eQEP<br>1x Display<br>3x McASP with DPI or<br>OSPI 2x USB 2.0 MIPI DSI (4L)<br>1x ADC<br>System Services Security<br>DMA Power  Secure  Security  PKA<br>Controller SHA (RSA/ECC)<br>Debug Manager Boot<br>SRAM MD5 DRBG<br>Timers/<br>Firewall IPC<br>RTC UID AES TRNG<br>**----- End of picture text -----**<br>


**Figure 3-1. Functional Block Diagram** 

Copyright © 2025 Texas Instruments Incorporated 

4 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **Table of Contents** 

**1 Features** ............................................................................1 **2 Applications** ..................................................................... 3 **3 Description** .......................................................................3 3.1 Functional Block Diagram........................................... 4 **4 Device Comparison** ......................................................... 6 4.1 Related Products........................................................ 7 **5 Terminal Configuration and Functions** ..........................8 5.1 Pin Diagrams.............................................................. 8 5.2 Pin Attributes...............................................................9 5.3 Signal Descriptions................................................... 40 5.4 Pin Connectivity Requirements.................................63 **6 Specifications** ................................................................ 66 6.1 Absolute Maximum Ratings...................................... 66 6.2 ESD Ratings............................................................. 67 6.3 Power-On Hours (POH)............................................ 68 6.4 Recommended Operating Conditions.......................69 6.5 Operating Performance Points..................................70 6.6 Power Consumption Summary................................. 70 6.7 Electrical Characteristics...........................................71 6.8 VPP Specifications for One-Time Programmable (OTP) eFuses..............................................................79 6.9 Thermal Resistance Characteristics......................... 80 6.10 Temperature Sensor Characteristics.......................81 

6.11 Timing and Switching Characteristics..................... 82 **7 Detailed Description** ....................................................199 7.1 Overview................................................................. 199 7.2 Processor Subsystem............................................. 199 7.3 Other Subsystem.................................................... 200 7.4 Peripherals..............................................................201 **8 Applications, Implementation, and Layout** ............... 206 8.1 Device Connection and Layout Fundamentals....... 206 8.2 Peripheral- and Interface-Specific Design Information................................................................ 207 8.3 Clock Routing Guidelines........................................212 **9 Device and Documentation Support** ..........................213 9.1 Device Nomenclature..............................................213 9.2 Tools and Software................................................. 216 9.3 Documentation Support.......................................... 216 9.4 Support Resources................................................. 216 9.5 Trademarks............................................................. 216 9.6 Electrostatic Discharge Caution..............................217 9.7 Glossary..................................................................217 **Revision History** ............................................................. 218 **10 Mechanical, Packaging, and Orderable Information** .................................................................. 219 10.1 Packaging Information.......................................... 219 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

5 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **4 Device Comparison** 

Table 4-1 shows the features of the superset device. 

## **Note** 

Availability of features listed in this table are a function of shared IO pins, where IO signals associated with many of the features are multiplexed to a limited number of pins. The SysConfig tool should be used to assign signal functions to pins. This will provide a better understanding of limitations associated with pin multiplexing. 

## **Note** 

To understand what device features are currently supported by TI Software Development Kits (SDKs), search for the _AM62L Software Build Sheet_ located in the Downloads tab option provided at **AM62LProcessor-SDK** . 

## **Table 4-1. Device Comparison** 

|**FEATURES**|**REFERENCE NAME**|**AM62L32**|**AM62L31**|
|---|---|---|---|
|||||
|**WKUP_CTRL_MMR_CFG0_JTAG_USER_ID[31:13]**(1)<br>Register bit values by device "Features" code (SeeNomenclature Descriptiontable for more information on device features)||||
||G:|0x391A7|0x39187|
|**PROCESSORS AND ACCELERATORS**||||
|Speed Grades||SeeDevice Speed GradesTable||
|Arm Cortex-A53 Microprocessor<br>Subsystem|A53SS|Dual Core|Single Core|
|Security Controller|Security Controller|Yes||
|Crypto Accelerators|Security|Yes||
|**PROGRAM AND DATA STORAGE**||||
|On-Chip Shared Memory (RAM)|OCSRAM in MAIN Domain|96KB||
||OCSRAM in WKUP Domain|64KB||
|DDR Subsystem|DDRSS with DDR4|16-bit data; up to 4GB||
||DDRSS with LPDDR4|16-bit data; up to 2GB||
|General-Purpose Memory Controller|GPMC|16-bit (GPMC, Raw NAND, Muxed-NOR)||
|**PERIPHERALS**||||
|Display Subsystem|DSS|1x DPI||
|||1x DSI||
|Modular Controller Area Network with Full<br>CAN-FD Support|MCAN|3||
|General-Purpose I/O|GPIO|133||
|Inter-Integrated Circuit Interface|I2C|5||
|Analog-to-Digital Converter|ADC|Yes||
|Multichannel Audio Serial Port|MCASP|3 (4/6/16 bits)||
|Multichannel Serial Peripheral Interface|MCSPI|4||
|Multi-Media Card/Secure Digital Interface|MMC/SD|1x eMMC (8 bits)||
|||2x SD/SDIO (4 bits)||
|Flash Subsystem (FSS)(2)|OSPI/QSPI|Yes||
|Gigabit Ethernet Interface|CPSW3G|Yes||
|General-Purpose Timers|TIMER|6||
|Global Timer Counter|GTC|1||
|Real Time Clock|RTC|Yes||
|Enhanced Pulse-Width Modulator Module|EPWM|3||



Copyright © 2025 Texas Instruments Incorporated 

6 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 4-1. Device Comparison (continued)** 

|**FEATURES**|**REFERENCE NAME**|**AM62L32**|**AM62L31**|
|---|---|---|---|
|Enhanced Capture Module|ECAP|3||
|Enhanced Quadrature Encoder Pulse<br>Module|EQEP|3||
|Universal Asynchronous Receiver and<br>Transmitter|UART|8||
|USB2.0 Controller with PHY|USB 2.0|2||



(1) For more details about the WKUP_CTRL_MMR_CFG0_JTAG_USER_ID register and DEVICE_ID bit field, see the device TRM. 

(2) One flash interface, configured as OSPI0 or QSPI0. 

## **4.1 Related Products** 

**Sitara™ processors** are a broad family of scalable processors based on Arm® Cortex®-A cores with flexible accelerators, peripherals, connectivity, and unified software support – perfect for sensors to servers. Sitara processors have the reliability and functional safety support required for use in industrial and automotive applications. 

**Sitara™ microcontrollers** are best-in-class Arm®-based 32-bit microcontrollers (MCUs) offering a scalable portfolio of high-performance and power-efficient devices to help meet your system needs. Bring capabilities such as functional safety, power efficiency, real-time control, advanced networking, analytics, and security to your designs. 

**AM64x Sitara™** processors target industrial applications such as Factory Automation and Control (FAC), and motor control that utilize Linux application processing cores (Cortex®-A53), real-time processing cores (Cortex®R5F), and Industrial Communication Subsystems (PRU_ICSSGs) to support protocols such as EtherCAT[®] , Profinet, or EtherNet/IP. AM64x implements one CPSW3G and two PRU_ICSSGs for supporting up to five gigabit Ethernet ports. The device also supports an extensive set of peripherals including a single lane of PCIe[® ] Gen2 or USB SuperSpeed Gen1, functional safety options, secure boot, and run-time security. 

**AM623 Sitara™** processors are an Internet of Things (IoT) and gateway SoC with Arm® Cortex®-A53-based object and gesture recognition. The low-cost AM623 Sitara[™] MPU family of application processors are built for Linux® application development. With scalable Arm® Cortex®-A53 performance, embedded features such as dual-display support, and an extensive set of peripherals make the AM623 device well-suited for a broad range of industrial and automotive applications. 

**AM625 Sitara™** processors are a human-machine-interaction SoC with Arm® Cortex®-A53-and full-HD dual display. The low-cost AM625 Sitara™ MPU family of application processors are built for Linux® application development. With scalable Arm® Cortex®-A53 performance, embedded features such as dual-display support, 3D graphics acceleration, and an extensive set of peripherals make the AM625 device well-suited for a broad range of industrial and automotive applications. 

**AM62A3 Sitara™** and **AM62A7 Sitara™** processors are an embedded vision SoC that utilizes 1-4x Cortex A-53 ARM Cores and 1 or 2 TOPS analytics hardware accelerator. This scalable, high performance AM62Ax Sitara MPU family of application processors are built for Linux application development. AM62Ax is well suited for a broad range of industrial and automotive applications with embedded features such as h.264/h.265 encode/ decode, secure boot, image signal processing and a deep learning accelerator. 

## **Products to complete your design:** 

- Ethernet PHYs 

- Power Management / PMICs 

- Clocks and timing 

- Power Switches 

- CAN Transceivers 

- ESD Protection 

Please reference the AM62Lx EVM schematic for details of how these devices are implemented in a system design, and bill of materials for specific part number recommendations. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 7 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **5 Terminal Configuration and Functions** 

## **5.1 Pin Diagrams** 

## **Note** 

The terms "ball", "pin", and "terminal" are used interchangeably throughout the document. An attempt is made to use "ball" only when referring to the physical package. 

Figure 5-1 shows the ball locations for the 373-ball flip-chip chip scale package ball grid array (FCCSP BGA), where the HTML version provides additional information when hovering your cursor over a ball. This figure is used in conjunction with Table 5-1 through Table 5-66 ( _Pin Attributes table and all Signal Descriptions tables, including the Connectivity Requirements table_ ). 

**==> picture [485 x 480] intentionally omitted <==**

**----- Start of picture text -----**<br>
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23<br>A VSS VSS MMC0_DAT6 VSS _DRVVBUSUSB1 I2C1_SDA I2C0_SDA MCASP0_AXR3 MCASP0_AXR1 VSS MCASP0_ACLKX MCASP0_ACLKR VSS _TXCLKPDSI0 _TXCLKNDSI0 VSS DSI0_TXP1 DSI0_TXN1 VSS DSI0_TXN2 DSI0_TXP2 VSS VSS<br>B VSS MMC0_CLK MMC0_DAT5 MMC0_DAT7 VSS MMC1_SDCD I2C0_SCL I2C2_SCL MCASP0_AXR0 MCASP0_AXR2 MCASP0_AFSX SPI0_D1 UART0_RTSn UART0_CTSn MCAN0_RX MCAN0_TX VSS DSI0_TXP0 DSI0_TXN0 VSS DSI0_TXP3 DSI0_TXN3 VSS<br>C MMC0_DAT2 MMC0_DAT3 MMC0_DAT4 _DRVVBUSUSB0 EXTINTn MCASP0_AFSR VSS UART0_TXD RESETSTATz VSS _CSn0OSPI0 OSPI0_D0 _CSn3OSPI0<br>D VSS MMC0_CMD MMC0_DAT0 MMC0_DAT1 MMC1_SDWP I2C1_SCL I2C2_SDA SPI0_CS1 UART0_RXD _REFCLK1EXT _TXRCALIBDSI0 _CSn2OSPI0 _CSn1OSPI0 OSPI0_D1 OSPI0_CLK OSPI0_D3<br>E DDR0_DQ3 VSS VSS VSS VSS VSS SPI0_CS0 SPI0_D0 SPI0_CLK VSS VSS RESETz _LBCLKOOSPI0 OSPI0_DQS OSPI0_D2<br>F DDR0_DQ2 DDR0_DM0 DDR0_DQ1 DDR0_DQ0 VSS VSS VSS OSPI0_D5 OSPI0_D7 OSPI0_D4 GPMC0_AD14 GPMC0_AD15<br>G DDR0_DQS0 DDR0_DQS0_n DDR0_DQ4 VSS VSS VSS VDDSHV1 _GENERAL1CAP_VDDS VSS VDDA_CORE_DSI VDDA_1P8_DSI VSS VSS VSS OSPI0_D6 GPMC0_AD13 GPMC0_AD12<br>H VSS DDR0_DQ6 DDR0_DQ7 DDR0_DQ5 DDR0_A5 DDR0_A1 VSS VDDSHV2 VDDSHV1 VDDA_CORE_DSI_CLK VSS VDDS1 VSS GPMC0_AD11 GPMC0_AD8 GPMC0_AD9 GPMC0_AD10 GPMC0_AD5 GPMC0_AD6<br>J DDR0_A4 _RESET0_nDDR0 CAP_VDDS_MMC0 VDD_CORE VDD_CORE VDD_CORE VDD_CORE VDDSHV0 GPMC0_AD7 GPMC0_AD3<br>K DDR0_CKE0 DDR0_A3 VSS VSS VDD_CORE VDDA_PLL1 VDD_CORE VSS CAP_VDDS_GPMC GPMC0_AD2 GPMC0_AD4<br>L DDR0_CAS_n DDR0_WE_n DDR0_CS0_n DDR0_ODT0 DDR0_A0 DDR0_A2 VSS VDDS_DDR VSS VDDA_PLL0 VSS VDD_CORE VSS VDDSHV0 VSS GPMC0_CSn1 GPMC0_CSn0 GPMC0_CLK GPMC0_AD0 GPMC0_AD1<br>M VSS DDR0_ACT_n DDR0_CAL0 DDR0_RAS_n VDDS_DDR VDDS_DDR VDDA_DDR_PLL0 VSS VDD_CORE CAP_VDDS_MMC2 VDDSHV4 GPMC0_WEn GPMC0_DIR GPMC0_CSn3 GPMC0_CSn2<br>N DDR0_A9 DDR0_BA1 DDR0_BA0 DDR0_BG1 DDR0_BG0 DDR0_A7 VSS VDDS_DDR VSS VSS VSS VDD_CORE VSS VDDA_ADC VPP _ADVn_ALEGPMC0 GPMC0_OEn_REn GPMC0_WPn _WAIT1GPMC0 _WAIT0GPMC0<br>P DDR0_CK0 DDR0_CK0_n VDDS_DDR VSS VDD_CORE VDD_CORE VDD_CORE VSS VDDS_WKUP GPMC0_BE1n _BE0n_CLEGPMC0<br>R VSS DDR0_A6 VSS VDD_CORE VDD_CORE VSS VSS VDDS_OSC0 MMC2_DAT3 MMC2_CLK<br>T DDR0_DQ10 VSS DDR0_DQ9 DDR0_A8 DDR0_A10 DDR0_A11 VSS VSS VDDSHV3 VDDA_1P8_USB VDDS0 _VDDSHV_MMCCAP VDD_RTC VDDS_RTC VSS MMC2_SDCD MMC2_SDWP MMC2_DAT1 MMC2_DAT2<br>U DDR0_DQ11 DDR0_DQ14 DDR0_DQ12 VSS VSS CAP_VDDS_MMC1 VSS VDDA_CORE_USB VDDA_3P3_USB VSS VSS VSS VDDA_3P3_SDIO VSS VSS MMC2_DAT0 MMC2_CMD<br>V DDR0_DQS1 DDR0_DQS1_n VSS DDR0_DQ8 DDR0_DQ13 DDR0_A13 VSS VSS ADC0_AIN0 ADC0_AIN3 ADC0_AIN1 ADC0_AIN2<br>W DDR0_DQ15 DDR0_DM1 DDR0_A12 RGMII1_RD3 VSS VSS RGMII1_TXC VSS RGMII1_TD1 VSS VSS VSS VSS _UART0_RTSnWKUP _UART0_CTSnWKUP<br>Y VSS MMC1_CLK MMC1_CMD MMC1_DAT1 RGMII1_RX_CTL RGMII1_RXC RGMII1_RD0 RGMII1_TD2 RGMII2_TXC EMU0 TMS RTC_PORz VSS VSS _UART0WKUP_RXD _CLKOUT0WKUP<br>AA MMC1_DAT0 MMC1_DAT2 VSS RGMII1_RD1 RGMII1_RD2 RGMII1_TD3 RGMII2_TD2 RGMII2_TD3 EMU1 PMIC_LPM_EN0 VSS WKUP_I2C0_SDA _UART0WKUP_TXD<br>AB VSS MMC1_DAT3 _RCALIBUSB0 USB0_DP USB1_DP USB1_VBUS VSS RGMII2_RD3 RGMII2_RD0 RGMII2_RD2 RGMII1_TX_CTL RGMII2_TX_CTL RGMII2_TD1 TCK TDO TRSTn RSVD0 PORz _WAKEUP0EXT _WAKEUP1EXT VSS WKUP_I2C0_SCL VSS<br>AC VSS VSS USB0_VBUS USB0_DM USB1_DM _RCALIBUSB1 RGMII2_RXC RGMII2_RX_CTL RGMII2_RD1 RGMII1_TD0 VSS RGMII2_TD0 MDIO0_MDIO VSS MDIO0_MDC TDI WKUP_OSC0_XO WKUP_OSC0_XI VSS LFOSC0_XO LFOSC0_XI VSS VSS<br>Not to scale<br>**----- End of picture text -----**<br>


## **Figure 5-1. ANB FCCSP BGA Package (Top View)** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

8 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **5.2 Pin Attributes** 

The following list describes the contents of each column in Table 5-1, _Pin Attributes (ANB Package)_ : 

1. **BALL NUMBER:** Ball numbers assigned to each terminal of the Ball Grid Array package. 

2. **BALL NAME:** Ball name assigned to each terminal of the Ball Grid Array package (this name is typically taken from the primary MUXMODE 0 signal function). 

3. **SIGNAL NAME:** Signal name(s) of all dedicated and pin multiplexed signal functions associated with a ball. 

## **Note** 

Many device pins support multiple signal functions. Some signal functions are selected via a single layer of multiplexers associated with pins. Other signal functions are selected via two or more layers of multiplexers, where one layer is associated with the pins and other layers are associated with peripheral logic functions. 

Table 5-1, _Pin Attributes (ANB Package)_ only defines signal multiplexing at the pins. For more information, related to signal multiplexing at the pins, see _Pad Configuration Registers_ section in _Device Configuration_ chapter of the device TRM. Refer to the respective peripheral chapter in the device TRM for information associated with peripheral signal multiplexing. 

4. **MUX MODE:** The MUXMODE value associated with each pin multiplexed signal function: 

   - a. MUXMODE 0 is the primary pin multiplexed signal function. However, the primary pin multiplexed signal function is not necessarily the default pin multiplexed signal function. 

## **Note** 

The value found in the MUX MODE AFTER RESET column defines the default pin multiplexed signal function selected when PORz is deasserted. 

- a. MUXMODE values 1 through 15 are possible for pin multiplexed signal functions. However, not all MUXMODE values have been implemented. The only valid MUXMODE values are those defined as pin multiplexed signal functions within the Pin Attributes table. Only valid values of MUXMODE should be used. 

- b. Bootstrap defines SOC configuration pins, where the logic state applied to each pin is latched on the rising edge of PORz. These input signal functions are fixed to their respective pins and are not programmable via MUXMODE. 

- c. An empty box means Not Applicable. 

## **Note** 

The following configurations of MUXMODE must be avoided for proper device operation. 

- Configuring multiple pins operating as inputs to the same pin multiplexed signal function is not supported as it can yield unexpected results. 

- Configuring a pin to an undefined pin multiplexing mode will cause the pin behavior to be undefined. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 9 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

5. **TYPE:** Signal type and direction: 

   - I = Input 

   - O = Output 

   - OD = Output, with open-drain output function 

   - IO = Input, Output, or simultaneously Input and Output 

   - IOD = Input, Output, or simultaneously Input and Output, with open-drain output function 

   - IOZ = Input, Output, or simultaneously Input and Output, with three-state output function 

   - OZ = Output with three-state output function 

   - A = Analog 

   - PWR = Power 

   - GND = Ground 

   - CAP = LDO Capacitor. 

6. **DSIS:** The deselected input state (DSIS) indicates the state driven to the subsystem input (logic "0", logic "1", or "pad" level) when the pin multiplexed signal function is not selected by MUXMODE. 

   - 0: Logic 0 driven to the subsystem input. 

   - 1: Logic 1 driven to the subsystem input. 

   - pad: Logic state of the pad is driven to the subsystem input. 

   - An empty box means Not Applicable. 

7. **BALL STATE DURING RESET RX/TX/PULL:** State of the terminal while PORz is asserted, where RX defines the state of the input buffer, TX defines the state of the output buffer, and PULL defines the state of internal pull resistors: 

   - RX (Input buffer) 

      - Off: The input buffer is disabled. 

      - On: The input buffer is enabled. 

      - BMD: The input buffer is enabled/disabled based on the boot mode selected. 

      - NA: Not Applicable. 

   - TX (Output buffer) 

      - Off: The output buffer is disabled. 

      - Low: The output buffer is enabled and drives VOL. 

      - NA: Not Applicable. 

   - PULL (Internal pull resistors) 

      - Off: Internal pull resistors are turned off. 

      - Up: Internal pull-up resistor is turned on. 

      - Down: Internal pull-down resistor is turned on. 

      - NA: Not Applicable. 

   - An empty box means Not Applicable. 

Copyright © 2025 Texas Instruments Incorporated 

10 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

8. **BALL STATE AFTER RESET RX/TX/PULL:** State of the terminal after PORz is deasserted, where RX defines the state of the input buffer, TX defines the state of the output buffer, and PULL defines the state of internal pull resistors: 

   - RX (Input buffer) 

      - Off: The input buffer is disabled. 

      - On: The input buffer is enabled. 

      - BMD: The input buffer is enabled/disabled based on the boot mode selected. 

      - NA: Not Applicable. 

   - TX (Output buffer) 

      - Off: The output buffer is disabled. 

      - SS: The subsystem selected with MUXMODE determines the output buffer state. 

      - NA: Not Applicable. 

   - PULL (Internal pull resistors) 

      - Off: Internal pull resistors are turned off. 

      - Up: Internal pull-up resistor is turned on. 

      - Down: Internal pull-down resistor is turned on. 

      - NA: Not Applicable. 

   - An empty box means Not Applicable. 

9. **MUX MODE AFTER RESET:** The value found in this column defines the default pin multiplexed signal function after PORz is deasserted. 

An empty box means Not Applicable. 

10. **I/O OPERATING VOLTAGE** : This column describes I/O operating voltage options of the respective power supply, when applicable. 

An empty box means Not Applicable. 

For more information, see valid operating voltage range(s) defined for each power supply in Section 6.4, _Recommended Operating Conditions_ . 

11. **POWER:** The power supply of the associated I/O, when applicable. 

An empty box means Not Applicable. 

12. **HYS:** Indicates if the input buffer associated with this I/O has hysteresis: 

   - Yes: With hysteresis 

   - No: Without hysteresis 

   - An empty box means Not Applicable. 

For more information, see the hysteresis values in Section 6.7, _Electrical Characteristics_ . 

13. **BUFFER TYPE:** This column defines the buffer type associated with a terminal. This information can be used to determine which Electrical Characteristics table is applicable. 

An empty box means Not Applicable. 

For electrical characteristics, refer to the appropriate buffer type table in Section 6.7, _Electrical Characteristics_ . 

14. **PULL UP/DOWN TYPE:** Indicates the presence of an internal pullup or pulldown resistor. Pullup and pulldown resistors can be enabled or disabled via software. 

   - PU: Internal pull-up 

   - PD: Internal pull-down 

   - PU/PD: Internal pull-up and pull-down 

   - An empty box means No internal pull. 

15. **PADCONFIG Register:** Name of the IO pad configuration register associated with Ball. 

16. **PADCONFIG Address:** Physical address of the IO pad configuration register associated with Ball. 

Copyright © 2025 Texas Instruments Incorporated 

11 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 5-1. Pin Attributes (ANB Package)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|V20|ADC0_AIN0|ADC0_AIN0||A|||||1.8V|VDDA_ADC||ADC||
|V22|ADC0_AIN1|ADC0_AIN1||A|||||1.8V|VDDA_ADC||ADC||
|V23|ADC0_AIN2|ADC0_AIN2||A|||||1.8V|VDDA_ADC||ADC||
|V21|ADC0_AIN3|ADC0_AIN3||A|||||1.8V|VDDA_ADC||ADC||
|T16|CAP_VDDSHV_MMC|CAP_VDDSHV_MMC||CAP||||||||||
|G11|CAP_VDDS_GENERAL1|CAP_VDDS_GENERAL1||CAP||||||||||
|K16|CAP_VDDS_GPMC|CAP_VDDS_GPMC||CAP||||||||||
|J8|CAP_VDDS_MMC0|CAP_VDDS_MMC0||CAP||||||||||
|U9|CAP_VDDS_MMC1|CAP_VDDS_MMC1||CAP||||||||||
|M16|CAP_VDDS_MMC2|CAP_VDDS_MMC2||CAP||||||||||
|M2|DDR0_ACT_n|DDR0_ACT_n||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|L1|DDR0_CAS_n|DDR0_CAS_n||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|M5|DDR0_RAS_n|DDR0_RAS_n||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|L2|DDR0_WE_n|DDR0_WE_n||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|L5|DDR0_A0|DDR0_A0||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|H6|DDR0_A1|DDR0_A1||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|L6|DDR0_A2|DDR0_A2||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|K2|DDR0_A3|DDR0_A3||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|J1|DDR0_A4|DDR0_A4||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|H5|DDR0_A5|DDR0_A5||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|R2|DDR0_A6|DDR0_A6||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|N6|DDR0_A7|DDR0_A7||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|T4|DDR0_A8|DDR0_A8||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|N1|DDR0_A9|DDR0_A9||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|T5|DDR0_A10|DDR0_A10||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|T6|DDR0_A11|DDR0_A11||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|W6|DDR0_A12|DDR0_A12||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|V6|DDR0_A13|DDR0_A13||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|N3|DDR0_BA0|DDR0_BA0||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|N2|DDR0_BA1|DDR0_BA1||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|N5|DDR0_BG0|DDR0_BG0||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|N4|DDR0_BG1|DDR0_BG1||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|M3|DDR0_CAL0|DDR0_CAL0||A|||||1.1V / 1.2V|VDDS_DDR||DDR||
|P1|DDR0_CK0|DDR0_CK0||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|P2|DDR0_CK0_n|DDR0_CK0_n||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|K1|DDR0_CKE0|DDR0_CKE0||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|L3|DDR0_CS0_n|DDR0_CS0_n||O|||||1.1V / 1.2V|VDDS_DDR||DDR||



Copyright © 2025 Texas Instruments Incorporated 

12 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|F2|DDR0_DM0|DDR0_DM0||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|W2|DDR0_DM1|DDR0_DM1||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|F4|DDR0_DQ0|DDR0_DQ0||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|F3|DDR0_DQ1|DDR0_DQ1||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|F1|DDR0_DQ2|DDR0_DQ2||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|E1|DDR0_DQ3|DDR0_DQ3||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|G4|DDR0_DQ4|DDR0_DQ4||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|H4|DDR0_DQ5|DDR0_DQ5||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|H2|DDR0_DQ6|DDR0_DQ6||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|H3|DDR0_DQ7|DDR0_DQ7||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|V4|DDR0_DQ8|DDR0_DQ8||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|T3|DDR0_DQ9|DDR0_DQ9||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|T1|DDR0_DQ10|DDR0_DQ10||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|U1|DDR0_DQ11|DDR0_DQ11||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|U4|DDR0_DQ12|DDR0_DQ12||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|V5|DDR0_DQ13|DDR0_DQ13||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|U2|DDR0_DQ14|DDR0_DQ14||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|W1|DDR0_DQ15|DDR0_DQ15||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|G1|DDR0_DQS0|DDR0_DQS0||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|G2|DDR0_DQS0_n|DDR0_DQS0_n||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|V1|DDR0_DQS1|DDR0_DQS1||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|V2|DDR0_DQS1_n|DDR0_DQS1_n||IO|||||1.1V / 1.2V|VDDS_DDR||DDR||
|L4|DDR0_ODT0|DDR0_ODT0||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|J2|DDR0_RESET0_n|DDR0_RESET0_n||O|||||1.1V / 1.2V|VDDS_DDR||DDR||
|A15|DSI0_TXCLKN|DSI0_TXCLKN||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|A14|DSI0_TXCLKP|DSI0_TXCLKP||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|D17|DSI0_TXRCALIB|DSI0_TXRCALIB||A|||||1.8V|VDDA_1P8_DSI||D-PHY||
|B19|DSI0_TXN0|DSI0_TXN0||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|A18|DSI0_TXN1|DSI0_TXN1||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|A20|DSI0_TXN2|DSI0_TXN2||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|B22|DSI0_TXN3|DSI0_TXN3||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|B18|DSI0_TXP0|DSI0_TXP0||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|A17|DSI0_TXP1|DSI0_TXP1||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|A21|DSI0_TXP2|DSI0_TXP2||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||
|B21|DSI0_TXP3|DSI0_TXP3||IO|||||1.8V|VDDA_1P8_DSI||D-PHY||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 13 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|Y16|EMU0<br>PADCONFIG:<br>PADCONFIG13<br>0x04084034|EMU0|0|IO|0|On / Off / Up|On / Off / Up|0|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|AA16|EMU1<br>PADCONFIG:<br>PADCONFIG14<br>0x04084038|EMU1|0|IO|0|On / Off / Up|On / Off / Up|0|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|C8|EXTINTn<br>PADCONFIG:<br>PADCONFIG122<br>0x040841E8|EXTINTn|0|I|1|Off / Off / NA|Off / Off / NA|7|1.8V / 3.3V|VDDSHV1|Yes|I2C OD FS||
|||GPIO0_105|7|IO|pad|||||||||
|D16|EXT_REFCLK1<br>PADCONFIG:<br>PADCONFIG121<br>0x040841E4|EXT_REFCLK1|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SYNC1_OUT|1|O||||||||||
|||SPI2_CS3|2|IO|1|||||||||
|||TIMER_IO0|4|IO|0|||||||||
|||CLKOUT0|5|O||||||||||
|||CP_GEMAC_CPTS0_RFT_CLK|6|I|0|||||||||
|||GPIO0_104|7|IO|pad|||||||||
|||ECAP0_IN_APWM_OUT|8|IO|0|||||||||
|||ADC_EXT_TRIGGER0|9|I|0|||||||||
|AB19|EXT_WAKEUP0|EXT_WAKEUP0||I||On / NA / NA|On / NA / NA||1.8V|VDDS_RTC|Yes|RTC-LVCMOS||
|AB20|EXT_WAKEUP1|EXT_WAKEUP1||I||On / NA / NA|On / NA / NA||1.8V|VDDS_RTC|Yes|RTC-LVCMOS||
|N19|GPMC0_ADVn_ALE<br>PADCONFIG:<br>PADCONFIG48<br>0x040840C0|GPMC0_ADVn_ALE|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA17|1|O||||||||||
|||MCASP1_AXR2|2|IO|0|||||||||
|||EHRPWM_TZn_IN1|4|I|0|||||||||
|||SPI3_CS3|5|IO|1|||||||||
|||TRC_DATA7|6|O||||||||||
|||GPIO0_32|7|IO|pad|||||||||
|L21|GPMC0_CLK<br>PADCONFIG:<br>PADCONFIG46<br>0x040840B8|GPMC0_CLK|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA16|1|O||||||||||
|||MCASP1_AXR3|2|IO|0|||||||||
|||GPMC0_FCLK_MUX|3|O||||||||||
|||EHRPWM1_B|4|IO|0|||||||||
|||TRC_DATA6|6|O||||||||||
|||GPIO0_31|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

14 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|M21|GPMC0_DIR<br>PADCONFIG:<br>PADCONFIG56<br>0x040840E0|GPMC0_DIR|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DE|1|O||||||||||
|||SPI2_D0|2|IO|0|||||||||
|||MCASP2_AXR13|3|IO|0|||||||||
|||EQEP1_B|4|I|0|||||||||
|||TRC_DATA14|6|O||||||||||
|||GPIO0_40|7|IO|pad|||||||||
|||EQEP2_S|8|IO|0|||||||||
|||MCAN1_TX|9|O||||||||||
|N20|GPMC0_OEn_REn<br>PADCONFIG:<br>PADCONFIG49<br>0x040840C4|GPMC0_OEn_REn|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA18|1|O||||||||||
|||MCASP1_AXR1|2|IO|0|||||||||
|||EHRPWM2_A|4|IO|0|||||||||
|||SPI3_CS2|5|IO|1|||||||||
|||TRC_DATA8|6|O||||||||||
|||GPIO0_33|7|IO|pad|||||||||
|M19|GPMC0_WEn<br>PADCONFIG:<br>PADCONFIG50<br>0x040840C8|GPMC0_WEn|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA19|1|O||||||||||
|||MCASP1_AXR0|2|IO|0|||||||||
|||EHRPWM2_B|4|IO|0|||||||||
|||SPI3_CS1|5|IO|1|||||||||
|||TRC_DATA9|6|O||||||||||
|||GPIO0_34|7|IO|pad|||||||||
|N21|GPMC0_WPn<br>PADCONFIG:<br>PADCONFIG55<br>0x040840DC|GPMC0_WPn|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_HSYNC|1|O||||||||||
|||SPI2_CLK|2|IO|0|||||||||
|||UART6_TXD|3|O||||||||||
|||EQEP1_A|4|I|0|||||||||
|||AUDIO_EXT_REFCLK1|5|IO|0|||||||||
|||TRC_DATA13|6|O||||||||||
|||GPIO0_39|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

15 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|L22|GPMC0_AD0<br>PADCONFIG:<br>PADCONFIG30<br>0x04084078|GPMC0_AD0|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA0|1|O||||||||||
|||UART6_RXD|2|I|1|||||||||
|||MCASP2_AXR4|3|IO|0|||||||||
|||I2C3_SCL|4|IOD|1|||||||||
|||ECAP0_IN_APWM_OUT|5|IO|0|||||||||
|||TRC_CLK|6|O||||||||||
|||GPIO0_15|7|IO|pad|||||||||
|||BOOTMODE00|Bootstrap|I||||||||||
|L23|GPMC0_AD1<br>PADCONFIG:<br>PADCONFIG31<br>0x0408407C|GPMC0_AD1|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA1|1|O||||||||||
|||UART6_TXD|2|O||||||||||
|||MCASP2_AXR5|3|IO|0|||||||||
|||I2C3_SDA|4|IOD|1|||||||||
|||ECAP1_IN_APWM_OUT|5|IO|0|||||||||
|||TRC_CTL|6|O||||||||||
|||GPIO0_16|7|IO|pad|||||||||
|||BOOTMODE01|Bootstrap|I||||||||||
|K22|GPMC0_AD2<br>PADCONFIG:<br>PADCONFIG32<br>0x04084080|GPMC0_AD2|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA2|1|O||||||||||
|||UART6_RTSn|2|O||||||||||
|||MCASP2_AXR6|3|IO|0|||||||||
|||SPI1_D0|4|IO|0|||||||||
|||TRC_DATA0|6|O||||||||||
|||GPIO0_17|7|IO|pad|||||||||
|||BOOTMODE02|Bootstrap|I||||||||||
|J23|GPMC0_AD3<br>PADCONFIG:<br>PADCONFIG33<br>0x04084084|GPMC0_AD3|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA3|1|O||||||||||
|||UART6_CTSn|2|I|1|||||||||
|||MCASP2_AXR7|3|IO|0|||||||||
|||SPI1_D1|4|IO|0|||||||||
|||TRC_DATA1|6|O||||||||||
|||GPIO0_18|7|IO|pad|||||||||
|||BOOTMODE03|Bootstrap|I||||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

16 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|K23|GPMC0_AD4<br>PADCONFIG:<br>PADCONFIG34<br>0x04084088|GPMC0_AD4|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA4|1|O||||||||||
|||UART5_RTSn|2|O||||||||||
|||MCASP2_AXR8|3|IO|0|||||||||
|||SPI1_CS0|4|IO|1|||||||||
|||TRC_DATA2|6|O||||||||||
|||GPIO0_19|7|IO|pad|||||||||
|||BOOTMODE04|Bootstrap|I||||||||||
|H22|GPMC0_AD5<br>PADCONFIG:<br>PADCONFIG35<br>0x0408408C|GPMC0_AD5|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA5|1|O||||||||||
|||UART5_CTSn|2|I|1|||||||||
|||MCASP2_AXR9|3|IO|0|||||||||
|||SPI1_CLK|4|IO|0|||||||||
|||TRC_DATA3|6|O||||||||||
|||GPIO0_20|7|IO|pad|||||||||
|||BOOTMODE05|Bootstrap|I||||||||||
|H23|GPMC0_AD6<br>PADCONFIG:<br>PADCONFIG36<br>0x04084090|GPMC0_AD6|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA6|1|O||||||||||
|||UART4_RTSn|2|O||||||||||
|||MCASP2_AXR10|3|IO|0|||||||||
|||SPI1_CS3|4|IO|1|||||||||
|||TRC_DATA4|6|O||||||||||
|||GPIO0_21|7|IO|pad|||||||||
|||BOOTMODE06|Bootstrap|I||||||||||
|J22|GPMC0_AD7<br>PADCONFIG:<br>PADCONFIG37<br>0x04084094|GPMC0_AD7|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA7|1|O||||||||||
|||UART4_CTSn|2|I|1|||||||||
|||MCASP2_AXR11|3|IO|0|||||||||
|||SPI1_CS1|4|IO|1|||||||||
|||MCASP1_AXR5|5|IO|0|||||||||
|||TRC_DATA5|6|O||||||||||
|||GPIO0_22|7|IO|pad|||||||||
|||BOOTMODE07|Bootstrap|I||||||||||



Copyright © 2025 Texas Instruments Incorporated 

17 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|H19|GPMC0_AD8<br>PADCONFIG:<br>PADCONFIG38<br>0x04084098|GPMC0_AD8|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA8|1|O||||||||||
|||UART2_RXD|2|I|1|||||||||
|||MCASP2_AXR0|3|IO|0|||||||||
|||SPI1_CS2|4|IO|1|||||||||
|||MCASP1_AXR4|5|IO|0|||||||||
|||GPIO0_23|7|IO|pad|||||||||
|||BOOTMODE08|Bootstrap|I||||||||||
|H20|GPMC0_AD9<br>PADCONFIG:<br>PADCONFIG39<br>0x0408409C|GPMC0_AD9|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA9|1|O||||||||||
|||UART2_TXD|2|O||||||||||
|||MCASP2_AXR1|3|IO|0|||||||||
|||TIMER_IO2|4|IO|0|||||||||
|||ECAP2_IN_APWM_OUT|5|IO|0|||||||||
|||GPIO0_24|7|IO|pad|||||||||
|||BOOTMODE09|Bootstrap|I||||||||||
|H21|GPMC0_AD10<br>PADCONFIG:<br>PADCONFIG40<br>0x040840A0|GPMC0_AD10|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA10|1|O||||||||||
|||UART3_RXD|2|I|1|||||||||
|||MCASP2_AXR2|3|IO|0|||||||||
|||EHRPWM0_SYNCI|4|I|0|||||||||
|||GPIO0_25|7|IO|pad|||||||||
|||OBSCLK0|8|O||||||||||
|||BOOTMODE10|Bootstrap|I||||||||||
|H18|GPMC0_AD11<br>PADCONFIG:<br>PADCONFIG41<br>0x040840A4|GPMC0_AD11|0|IO|0|BMD / Off / Off|BDM / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA11|1|O||||||||||
|||UART3_TXD|2|O||||||||||
|||MCASP2_AXR3|3|IO|0|||||||||
|||EHRPWM0_SYNCO|4|O||||||||||
|||TRC_DATA23|6|O||||||||||
|||GPIO0_26|7|IO|pad|||||||||
|||BOOTMODE11|Bootstrap|I||||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

18 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|G23|GPMC0_AD12<br>PADCONFIG:<br>PADCONFIG42<br>0x040840A8|GPMC0_AD12|0|IO|0|On / Off / Off|On / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA12|1|O||||||||||
|||UART4_RXD|2|I|1|||||||||
|||MCASP2_AFSX|3|IO|0|||||||||
|||EHRPWM_TZn_IN2|4|I|0|||||||||
|||TRC_DATA22|6|O||||||||||
|||GPIO0_27|7|IO|pad|||||||||
|||BOOTMODE12|Bootstrap|I||||||||||
|G22|GPMC0_AD13<br>PADCONFIG:<br>PADCONFIG43<br>0x040840AC|GPMC0_AD13|0|IO|0|On / Off / Off|On / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA13|1|O||||||||||
|||UART4_TXD|2|O||||||||||
|||MCASP2_ACLKX|3|IO|0|||||||||
|||EHRPWM0_A|4|IO|0|||||||||
|||TRC_DATA21|6|O||||||||||
|||GPIO0_28|7|IO|pad|||||||||
|||BOOTMODE13|Bootstrap|I||||||||||
|F22|GPMC0_AD14<br>PADCONFIG:<br>PADCONFIG44<br>0x040840B0|GPMC0_AD14|0|IO|0|On / Off / Off|On / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA14|1|O||||||||||
|||UART5_RXD|2|I|1|||||||||
|||MCASP2_AFSR|3|IO|0|||||||||
|||EHRPWM0_B|4|IO|0|||||||||
|||TRC_DATA20|6|O||||||||||
|||GPIO0_29|7|IO|pad|||||||||
|||UART2_CTSn|8|I|1|||||||||
|||BOOTMODE14|Bootstrap|I||||||||||
|F23|GPMC0_AD15<br>PADCONFIG:<br>PADCONFIG45<br>0x040840B4|GPMC0_AD15|0|IO|0|On / Off / Off|On / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA15|1|O||||||||||
|||UART5_TXD|2|O||||||||||
|||MCASP2_ACLKR|3|IO|0|||||||||
|||EHRPWM1_A|4|IO|0|||||||||
|||TRC_DATA19|6|O||||||||||
|||GPIO0_30|7|IO|pad|||||||||
|||UART2_RTSn|8|O||||||||||
|||BOOTMODE15|Bootstrap|I||||||||||



Copyright © 2025 Texas Instruments Incorporated 

19 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|P23|GPMC0_BE0n_CLE<br>PADCONFIG:<br>PADCONFIG51<br>0x040840CC|GPMC0_BE0n_CLE|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA20|1|O||||||||||
|||MCASP1_ACLKX|2|IO|0|||||||||
|||EQEP0_A|4|I|0|||||||||
|||SPI3_CS0|5|IO|1|||||||||
|||TRC_DATA10|6|O||||||||||
|||GPIO0_35|7|IO|pad|||||||||
|P22|GPMC0_BE1n<br>PADCONFIG:<br>PADCONFIG52<br>0x040840D0|GPMC0_BE1n|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA21|1|O||||||||||
|||MCASP2_AXR12|3|IO|0|||||||||
|||EQEP0_B|4|I|0|||||||||
|||SPI3_CLK|5|IO|0|||||||||
|||TRC_DATA11|6|O||||||||||
|||GPIO0_36|7|IO|pad|||||||||
|L20|GPMC0_CSn0<br>PADCONFIG:<br>PADCONFIG57<br>0x040840E4|GPMC0_CSn0|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_VSYNC|1|O||||||||||
|||SPI2_D1|2|IO|0|||||||||
|||MCASP2_AXR14|3|IO|0|||||||||
|||EQEP1_S|4|IO|0|||||||||
|||TRC_DATA15|6|O||||||||||
|||GPIO0_41|7|IO|pad|||||||||
|L19|GPMC0_CSn1<br>PADCONFIG:<br>PADCONFIG58<br>0x040840E8|GPMC0_CSn1|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_PCLK|1|O||||||||||
|||SPI2_CS0|2|IO|1|||||||||
|||MCASP2_AXR15|3|IO|0|||||||||
|||EQEP1_I|4|IO|0|||||||||
|||TRC_DATA16|6|O||||||||||
|||GPIO0_42|7|IO|pad|||||||||
|M23|GPMC0_CSn2<br>PADCONFIG:<br>PADCONFIG59<br>0x040840EC|GPMC0_CSn2|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||I2C2_SCL|1|IOD|1|||||||||
|||MCASP1_AXR4|2|IO|0|||||||||
|||UART4_RXD|3|I|1|||||||||
|||ADC_EXT_TRIGGER0|4|I|0|||||||||
|||VOUT0_EXTPCLKIN|5|I|0|||||||||
|||TRC_DATA17|6|O||||||||||
|||GPIO0_43|7|IO|pad|||||||||
|||MCASP1_AFSR|8|IO|0|||||||||



Copyright © 2025 Texas Instruments Incorporated 

20 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|M22|GPMC0_CSn3<br>PADCONFIG:<br>PADCONFIG60<br>0x040840F0|GPMC0_CSn3|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||I2C2_SDA|1|IOD|1|||||||||
|||WKUP_CLKOUT0|2|O||||||||||
|||UART4_TXD|3|O||||||||||
|||MCASP1_AXR5|4|IO|0|||||||||
|||ADC_EXT_TRIGGER1|5|I|0|||||||||
|||TRC_DATA18|6|O||||||||||
|||GPIO0_44|7|IO|pad|||||||||
|||MCASP1_ACLKR|8|IO|0|||||||||
|N23|GPMC0_WAIT0<br>PADCONFIG:<br>PADCONFIG53<br>0x040840D4|GPMC0_WAIT0|0|I|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA22|1|O||||||||||
|||MCASP1_AFSX|2|IO|0|||||||||
|||EQEP0_S|4|IO|0|||||||||
|||SPI3_D0|5|IO|0|||||||||
|||TRC_DATA12|6|O||||||||||
|||GPIO0_37|7|IO|pad|||||||||
|N22|GPMC0_WAIT1<br>PADCONFIG:<br>PADCONFIG54<br>0x040840D8|GPMC0_WAIT1|0|I|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV0|Yes|LVCMOS|PU/PD|
|||VOUT0_DATA23|1|O||||||||||
|||SPI2_CS1|2|IO|1|||||||||
|||UART6_RXD|3|I|1|||||||||
|||EQEP0_I|4|IO|0|||||||||
|||SPI3_D1|5|IO|0|||||||||
|||GPIO0_38|7|IO|pad|||||||||
|||EQEP2_I|8|IO|0|||||||||
|||MCAN1_RX|9|I|1|||||||||
|B7|I2C0_SCL<br>PADCONFIG:<br>PADCONFIG115<br>0x040841CC|I2C0_SCL|0|IOD|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SYNC0_OUT|2|O||||||||||
|||OBSCLK1|3|O||||||||||
|||UART1_DCDn|4|I|1|||||||||
|||EQEP2_A|5|I|0|||||||||
|||EHRPWM_SOCA|6|O||||||||||
|||GPIO0_98|7|IO|pad|||||||||
|||ECAP1_IN_APWM_OUT|8|IO|0|||||||||
|||SPI2_CS0|9|IO|1|||||||||



Copyright © 2025 Texas Instruments Incorporated 

21 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|A7|I2C0_SDA<br>PADCONFIG:<br>PADCONFIG116<br>0x040841D0|I2C0_SDA|0|IOD|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI2_CS2|2|IO|1|||||||||
|||TIMER_IO1|3|IO|0|||||||||
|||UART1_DSRn|4|I|1|||||||||
|||EQEP2_B|5|I|0|||||||||
|||EHRPWM_SOCB|6|O||||||||||
|||GPIO0_99|7|IO|pad|||||||||
|||ECAP2_IN_APWM_OUT|8|IO|0|||||||||
|D7|I2C1_SCL<br>PADCONFIG:<br>PADCONFIG117<br>0x040841D4|I2C1_SCL|0|IOD|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||UART1_RXD|1|I|1|||||||||
|||TIMER_IO0|2|IO|0|||||||||
|||SPI2_CS1|3|IO|1|||||||||
|||EHRPWM0_SYNCI|4|I|0|||||||||
|||GPIO0_100|7|IO|pad|||||||||
|||EHRPWM2_A|8|IO|0|||||||||
|||MMC2_SDCD|9|I|0|||||||||
|A6|I2C1_SDA<br>PADCONFIG:<br>PADCONFIG118<br>0x040841D8|I2C1_SDA|0|IOD|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||UART1_TXD|1|O||||||||||
|||TIMER_IO1|2|IO|0|||||||||
|||SPI2_CLK|3|IO|0|||||||||
|||EHRPWM0_SYNCO|4|O||||||||||
|||GPIO0_101|7|IO|pad|||||||||
|||EHRPWM2_B|8|IO|0|||||||||
|||MMC2_SDWP|9|I|0|||||||||
|B8|I2C2_SCL<br>PADCONFIG:<br>PADCONFIG119<br>0x040841DC|I2C2_SCL|0|IOD|1|Off / Off / NA|Off / Off / NA|7|1.8V / 3.3V|VDDSHV1|Yes|I2C OD FS||
|||GPIO0_102|7|IO|pad|||||||||
|D8|I2C2_SDA<br>PADCONFIG:<br>PADCONFIG120<br>0x040841E0|I2C2_SDA|0|IOD|1|Off / Off / NA|Off / Off / NA|7|1.8V / 3.3V|VDDSHV1|Yes|I2C OD FS||
|||GPIO0_103|7|IO|pad|||||||||
|AC21|LFOSC0_XI|LFOSC0_XI||I|||||1.8V|VDDS_RTC||LFXOSC||
|AC20|LFOSC0_XO|LFOSC0_XO||O|||||1.8V|VDDS_RTC||LFXOSC||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

22 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|B15|MCAN0_RX<br>PADCONFIG:<br>PADCONFIG114<br>0x040841C8|MCAN0_RX|0|I|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||UART5_TXD|1|O||||||||||
|||TIMER_IO3|2|IO|0|||||||||
|||SYNC3_OUT|3|O||||||||||
|||UART1_RIn|4|I|1|||||||||
|||EQEP2_S|5|IO|0|||||||||
|||GPIO0_97|7|IO|pad|||||||||
|||MCASP2_AXR1|8|IO|0|||||||||
|||EHRPWM_TZn_IN4|9|I|0|||||||||
|B16|MCAN0_TX<br>PADCONFIG:<br>PADCONFIG113<br>0x040841C4|MCAN0_TX|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||UART5_RXD|1|I|1|||||||||
|||TIMER_IO2|2|IO|0|||||||||
|||SYNC2_OUT|3|O||||||||||
|||UART1_DTRn|4|O||||||||||
|||EQEP2_I|5|IO|0|||||||||
|||GPIO0_96|7|IO|pad|||||||||
|||MCASP2_AXR0|8|IO|0|||||||||
|||EHRPWM_TZn_IN3|9|I|0|||||||||
|A12|MCASP0_ACLKR<br>PADCONFIG:<br>PADCONFIG103<br>0x0408419C|MCASP0_ACLKR|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI2_CLK|1|IO|0|||||||||
|||UART1_TXD|2|O||||||||||
|||ADC_EXT_TRIGGER1|3|I|0|||||||||
|||EHRPWM0_B|5|IO|0|||||||||
|||GPIO0_86|7|IO|pad|||||||||
|||EQEP1_I|8|IO|0|||||||||
|A11|MCASP0_ACLKX<br>PADCONFIG:<br>PADCONFIG100<br>0x04084190|MCASP0_ACLKX|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI2_CS1|1|IO|1|||||||||
|||ECAP2_IN_APWM_OUT|2|IO|0|||||||||
|||GPIO0_83|7|IO|pad|||||||||
|||EQEP1_A|8|I|0|||||||||
|C11|MCASP0_AFSR<br>PADCONFIG:<br>PADCONFIG102<br>0x04084198|MCASP0_AFSR|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI2_CS0|1|IO|1|||||||||
|||UART1_RXD|2|I|1|||||||||
|||ADC_EXT_TRIGGER0|3|I|0|||||||||
|||EHRPWM0_A|5|IO|0|||||||||
|||GPIO0_85|7|IO|pad|||||||||
|||EQEP1_S|8|IO|0|||||||||



Copyright © 2025 Texas Instruments Incorporated 

23 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|B11|MCASP0_AFSX<br>PADCONFIG:<br>PADCONFIG101<br>0x04084194|MCASP0_AFSX|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI2_CS3|1|IO|1|||||||||
|||AUDIO_EXT_REFCLK1|2|IO|0|||||||||
|||GPIO0_84|7|IO|pad|||||||||
|||EQEP1_B|8|I|0|||||||||
|B9|MCASP0_AXR0<br>PADCONFIG:<br>PADCONFIG99<br>0x0408418C|MCASP0_AXR0|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||AUDIO_EXT_REFCLK0|2|IO|0|||||||||
|||EHRPWM1_B|5|IO|0|||||||||
|||GPIO0_82|7|IO|pad|||||||||
|||EQEP0_I|8|IO|0|||||||||
|A9|MCASP0_AXR1<br>PADCONFIG:<br>PADCONFIG98<br>0x04084188|MCASP0_AXR1|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI2_CS2|1|IO|1|||||||||
|||ECAP1_IN_APWM_OUT|2|IO|0|||||||||
|||EHRPWM1_A|5|IO|0|||||||||
|||GPIO0_81|7|IO|pad|||||||||
|||EQEP0_S|8|IO|0|||||||||
|B10|MCASP0_AXR2<br>PADCONFIG:<br>PADCONFIG97<br>0x04084184|MCASP0_AXR2|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI2_D1|1|IO|0|||||||||
|||UART1_RTSn|2|O||||||||||
|||UART6_TXD|3|O||||||||||
|||ECAP2_IN_APWM_OUT|4|IO|0|||||||||
|||MCAN1_TX|5|O||||||||||
|||GPIO0_80|7|IO|pad|||||||||
|||EQEP0_B|8|I|0|||||||||
|A8|MCASP0_AXR3<br>PADCONFIG:<br>PADCONFIG96<br>0x04084180|MCASP0_AXR3|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI2_D0|1|IO|0|||||||||
|||UART1_CTSn|2|I|1|||||||||
|||UART6_RXD|3|I|1|||||||||
|||ECAP1_IN_APWM_OUT|4|IO|0|||||||||
|||MCAN1_RX|5|I|1|||||||||
|||GPIO0_79|7|IO|pad|||||||||
|||EQEP0_A|8|I|0|||||||||
|AC15|MDIO0_MDC<br>PADCONFIG:<br>PADCONFIG83<br>0x0408414C|MDIO0_MDC|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_66|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

24 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|AC13|MDIO0_MDIO<br>PADCONFIG:<br>PADCONFIG82<br>0x04084148|MDIO0_MDIO|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_65|7|IO|pad|||||||||
|B2|MMC0_CLK<br>PADCONFIG:<br>PADCONFIG131<br>0x0408420C|MMC0_CLK|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||I2C3_SCL|1|IOD|1|||||||||
|||EHRPWM2_A|2|IO|0|||||||||
|||SPI1_CS1|5|IO|1|||||||||
|||TIMER_IO0|6|IO|0|||||||||
|||GPIO0_114|7|IO|pad|||||||||
|D2|MMC0_CMD<br>PADCONFIG:<br>PADCONFIG133<br>0x04084214|MMC0_CMD|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||I2C3_SDA|1|IOD|1|||||||||
|||EHRPWM2_B|2|IO|0|||||||||
|||SPI1_CS2|5|IO|1|||||||||
|||TIMER_IO1|6|IO|0|||||||||
|||GPIO0_115|7|IO|pad|||||||||
|Y2|MMC1_CLK<br>PADCONFIG:<br>PADCONFIG138<br>0x04084228|MMC1_CLK|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV3|Yes|SDIO|PU/PD|
|||TIMER_IO0|2|IO|0|||||||||
|||UART3_RXD|3|I|1|||||||||
|||SPI3_CS0|5|IO|1|||||||||
|||SPI2_CS2|6|IO|1|||||||||
|||GPIO0_120|7|IO|pad|||||||||
|Y3|MMC1_CMD<br>PADCONFIG:<br>PADCONFIG140<br>0x04084230|MMC1_CMD|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV3|Yes|SDIO|PU/PD|
|||TIMER_IO1|2|IO|0|||||||||
|||UART3_TXD|3|O||||||||||
|||SPI3_CLK|5|IO|0|||||||||
|||SPI2_CS0|6|IO|1|||||||||
|||GPIO0_121|7|IO|pad|||||||||
|B6|MMC1_SDCD<br>PADCONFIG:<br>PADCONFIG141<br>0x04084234|MMC1_SDCD|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||UART6_RXD|1|I|1|||||||||
|||TIMER_IO2|2|IO|0|||||||||
|||UART3_RTSn|3|O||||||||||
|||MCAN2_RX|4|I|1|||||||||
|||SPI3_CS3|5|IO|1|||||||||
|||SPI2_CLK|6|IO|0|||||||||
|||GPIO0_122|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 25 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|D6|MMC1_SDWP<br>PADCONFIG:<br>PADCONFIG142<br>0x04084238|MMC1_SDWP|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||UART6_TXD|1|O||||||||||
|||TIMER_IO3|2|IO|0|||||||||
|||UART3_CTSn|3|I|1|||||||||
|||MCAN2_TX|4|O||||||||||
|||SPI3_CS1|5|IO|1|||||||||
|||GPIO0_123|7|IO|pad|||||||||
|R23|MMC2_CLK<br>PADCONFIG:<br>PADCONFIG65<br>0x04084104|MMC2_CLK|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV4|Yes|SDIO|PU/PD|
|||MCASP1_ACLKR|1|IO|0|||||||||
|||MCASP1_AXR5|2|IO|0|||||||||
|||UART6_RXD|3|I|1|||||||||
|||GPIO0_49|7|IO|pad|||||||||
|U23|MMC2_CMD<br>PADCONFIG:<br>PADCONFIG67<br>0x0408410C|MMC2_CMD|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV4|Yes|SDIO|PU/PD|
|||MCASP1_AFSR|1|IO|0|||||||||
|||MCASP1_AXR4|2|IO|0|||||||||
|||UART6_TXD|3|O||||||||||
|||GPIO0_50|7|IO|pad|||||||||
|T20|MMC2_SDCD<br>PADCONFIG:<br>PADCONFIG68<br>0x04084110|MMC2_SDCD|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV4|Yes|LVCMOS|PU/PD|
|||MCASP1_ACLKX|1|IO|0|||||||||
|||UART4_RXD|3|I|1|||||||||
|||GPIO0_51|7|IO|pad|||||||||
|T21|MMC2_SDWP<br>PADCONFIG:<br>PADCONFIG69<br>0x04084114|MMC2_SDWP|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV4|Yes|LVCMOS|PU/PD|
|||MCASP1_AFSX|1|IO|0|||||||||
|||UART4_TXD|3|O||||||||||
|||GPIO0_52|7|IO|pad|||||||||
|D3|MMC0_DAT0<br>PADCONFIG:<br>PADCONFIG130<br>0x04084208|MMC0_DAT0|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||UART3_CTSn|1|I|1|||||||||
|||EHRPWM_TZn_IN1|2|I|0|||||||||
|||SPI2_CLK|6|IO|0|||||||||
|||GPIO0_113|7|IO|pad|||||||||
|D4|MMC0_DAT1<br>PADCONFIG:<br>PADCONFIG129<br>0x04084204|MMC0_DAT1|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||UART3_RTSn|1|O||||||||||
|||EHRPWM1_B|2|IO|0|||||||||
|||SPI1_CS3|5|IO|1|||||||||
|||SPI2_CS0|6|IO|1|||||||||
|||GPIO0_112|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

26 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|C1|MMC0_DAT2<br>PADCONFIG:<br>PADCONFIG128<br>0x04084200|MMC0_DAT2|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||UART3_TXD|1|O||||||||||
|||EHRPWM1_A|2|IO|0|||||||||
|||MCAN2_TX|3|O||||||||||
|||SPI1_CLK|5|IO|0|||||||||
|||TIMER_IO0|6|IO|0|||||||||
|||GPIO0_111|7|IO|pad|||||||||
|C2|MMC0_DAT3<br>PADCONFIG:<br>PADCONFIG127<br>0x040841FC|MMC0_DAT3|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||UART3_RXD|1|I|1|||||||||
|||EHRPWM0_B|2|IO|0|||||||||
|||MCAN2_RX|3|I|1|||||||||
|||SPI1_CS0|5|IO|1|||||||||
|||SPI2_CS2|6|IO|1|||||||||
|||GPIO0_110|7|IO|pad|||||||||
|C4|MMC0_DAT4<br>PADCONFIG:<br>PADCONFIG126<br>0x040841F8|MMC0_DAT4|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||UART2_CTSn|1|I|1|||||||||
|||EHRPWM0_A|2|IO|0|||||||||
|||SPI1_CLK|5|IO|0|||||||||
|||SPI2_D1|6|IO|0|||||||||
|||GPIO0_109|7|IO|pad|||||||||
|B3|MMC0_DAT5<br>PADCONFIG:<br>PADCONFIG125<br>0x040841F4|MMC0_DAT5|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||UART2_RTSn|1|O||||||||||
|||EHRPWM_TZn_IN2|2|I|0|||||||||
|||SPI1_CS0|5|IO|1|||||||||
|||SPI2_D0|6|IO|0|||||||||
|||GPIO0_108|7|IO|pad|||||||||
|A3|MMC0_DAT6<br>PADCONFIG:<br>PADCONFIG124<br>0x040841F0|MMC0_DAT6|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||UART2_TXD|1|O||||||||||
|||EHRPWM0_SYNCO|2|O||||||||||
|||MCAN1_TX|3|O||||||||||
|||SPI2_CLK|4|IO|0|||||||||
|||SPI1_D1|5|IO|0|||||||||
|||SPI2_CS3|6|IO|1|||||||||
|||GPIO0_107|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

27 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|B4|MMC0_DAT7<br>PADCONFIG:<br>PADCONFIG123<br>0x040841EC|MMC0_DAT7|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV2|Yes|SDIO|PU/PD|
|||UART2_RXD|1|I|1|||||||||
|||EHRPWM0_SYNCI|2|I|0|||||||||
|||MCAN1_RX|3|I|1|||||||||
|||SPI1_D0|5|IO|0|||||||||
|||SPI2_CS1|6|IO|1|||||||||
|||GPIO0_106|7|IO|pad|||||||||
|AA1|MMC1_DAT0<br>PADCONFIG:<br>PADCONFIG137<br>0x04084224|MMC1_DAT0|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV3|Yes|SDIO|PU/PD|
|||CP_GEMAC_CPTS0_HW2TSPUSH|1|I|0|||||||||
|||TIMER_IO3|2|IO|0|||||||||
|||UART2_CTSn|3|I|1|||||||||
|||ECAP2_IN_APWM_OUT|4|IO|0|||||||||
|||SPI2_D1|6|IO|0|||||||||
|||GPIO0_119|7|IO|pad|||||||||
|Y4|MMC1_DAT1<br>PADCONFIG:<br>PADCONFIG136<br>0x04084220|MMC1_DAT1|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV3|Yes|SDIO|PU/PD|
|||CP_GEMAC_CPTS0_HW1TSPUSH|1|I|0|||||||||
|||TIMER_IO2|2|IO|0|||||||||
|||UART2_RTSn|3|O||||||||||
|||ECAP1_IN_APWM_OUT|4|IO|0|||||||||
|||SPI3_CS2|5|IO|1|||||||||
|||SPI2_D0|6|IO|0|||||||||
|||GPIO0_118|7|IO|pad|||||||||
|AA2|MMC1_DAT2<br>PADCONFIG:<br>PADCONFIG135<br>0x0408421C|MMC1_DAT2|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV3|Yes|SDIO|PU/PD|
|||CP_GEMAC_CPTS0_TS_SYNC|1|O||||||||||
|||TIMER_IO1|2|IO|0|||||||||
|||UART2_TXD|3|O||||||||||
|||MCAN1_TX|4|O||||||||||
|||SPI3_D1|5|IO|0|||||||||
|||SPI2_CS3|6|IO|1|||||||||
|||GPIO0_117|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

28 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|AB2|MMC1_DAT3<br>PADCONFIG:<br>PADCONFIG134<br>0x04084218|MMC1_DAT3|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV3|Yes|SDIO|PU/PD|
|||CP_GEMAC_CPTS0_TS_COMP|1|O||||||||||
|||TIMER_IO0|2|IO|0|||||||||
|||UART2_RXD|3|I|1|||||||||
|||MCAN1_RX|4|I|1|||||||||
|||SPI3_D0|5|IO|0|||||||||
|||SPI2_CS1|6|IO|1|||||||||
|||GPIO0_116|7|IO|pad|||||||||
|U22|MMC2_DAT0<br>PADCONFIG:<br>PADCONFIG64<br>0x04084100|MMC2_DAT0|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV4|Yes|SDIO|PU/PD|
|||MCASP1_AXR0|1|IO|0|||||||||
|||GPIO0_48|7|IO|pad|||||||||
|T22|MMC2_DAT1<br>PADCONFIG:<br>PADCONFIG63<br>0x040840FC|MMC2_DAT1|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV4|Yes|SDIO|PU/PD|
|||MCASP1_AXR1|1|IO|0|||||||||
|||GPIO0_47|7|IO|pad|||||||||
|T23|MMC2_DAT2<br>PADCONFIG:<br>PADCONFIG62<br>0x040840F8|MMC2_DAT2|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV4|Yes|SDIO|PU/PD|
|||MCASP1_AXR2|1|IO|0|||||||||
|||UART5_TXD|3|O||||||||||
|||GPIO0_46|7|IO|pad|||||||||
|R22|MMC2_DAT3<br>PADCONFIG:<br>PADCONFIG61<br>0x040840F4|MMC2_DAT3|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV4|Yes|SDIO|PU/PD|
|||MCASP1_AXR3|1|IO|0|||||||||
|||UART5_RXD|3|I|1|||||||||
|||GPIO0_45|7|IO|pad|||||||||
|D22|OSPI0_CLK<br>PADCONFIG:<br>PADCONFIG15<br>0x0408403C|OSPI0_CLK|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_0|7|IO|pad|||||||||
|E22|OSPI0_DQS<br>PADCONFIG:<br>PADCONFIG17<br>0x04084044|OSPI0_DQS|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||UART5_CTSn|5|I|1|||||||||
|||GPIO0_2|7|IO|pad|||||||||
|E18|OSPI0_LBCLKO<br>PADCONFIG:<br>PADCONFIG16<br>0x04084040|OSPI0_LBCLKO|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||UART5_RTSn|5|O||||||||||
|||GPIO0_1|7|IO|pad|||||||||
|C20|OSPI0_CSn0<br>PADCONFIG:<br>PADCONFIG26<br>0x04084068|OSPI0_CSn0|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_11|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 29 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|D20|OSPI0_CSn1<br>PADCONFIG:<br>PADCONFIG27<br>0x0408406C|OSPI0_CSn1|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_12|7|IO|pad|||||||||
|D18|OSPI0_CSn2<br>PADCONFIG:<br>PADCONFIG28<br>0x04084070|OSPI0_CSn2|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||SPI1_CS1|1|IO|1|||||||||
|||OSPI0_RESET_OUT1|2|O||||||||||
|||MCASP1_AFSR|3|IO|0|||||||||
|||MCASP1_AXR2|4|IO|0|||||||||
|||UART5_RXD|5|I|1|||||||||
|||ADC_EXT_TRIGGER0|6|I|0|||||||||
|||GPIO0_13|7|IO|pad|||||||||
|C23|OSPI0_CSn3<br>PADCONFIG:<br>PADCONFIG29<br>0x04084074|OSPI0_CSn3|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||OSPI0_RESET_OUT0|1|O||||||||||
|||OSPI0_ECC_FAIL|2|I|1|||||||||
|||MCASP1_ACLKR|3|IO|0|||||||||
|||MCASP1_AXR3|4|IO|0|||||||||
|||UART5_TXD|5|O||||||||||
|||ADC_EXT_TRIGGER1|6|I|0|||||||||
|||GPIO0_14|7|IO|pad|||||||||
|C22|OSPI0_D0<br>PADCONFIG:<br>PADCONFIG18<br>0x04084048|OSPI0_D0|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_3|7|IO|pad|||||||||
|D21|OSPI0_D1<br>PADCONFIG:<br>PADCONFIG19<br>0x0408404C|OSPI0_D1|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_4|7|IO|pad|||||||||
|E23|OSPI0_D2<br>PADCONFIG:<br>PADCONFIG20<br>0x04084050|OSPI0_D2|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_5|7|IO|pad|||||||||
|D23|OSPI0_D3<br>PADCONFIG:<br>PADCONFIG21<br>0x04084054|OSPI0_D3|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||GPIO0_6|7|IO|pad|||||||||
|F21|OSPI0_D4<br>PADCONFIG:<br>PADCONFIG22<br>0x04084058|OSPI0_D4|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||SPI1_CS0|1|IO|1|||||||||
|||MCASP1_AXR1|2|IO|0|||||||||
|||UART6_RXD|3|I|1|||||||||
|||GPIO0_7|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

30 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|F19|OSPI0_D5<br>PADCONFIG:<br>PADCONFIG23<br>0x0408405C|OSPI0_D5|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||SPI1_CLK|1|IO|0|||||||||
|||MCASP1_AXR0|2|IO|0|||||||||
|||UART6_TXD|3|O||||||||||
|||GPIO0_8|7|IO|pad|||||||||
|G20|OSPI0_D6<br>PADCONFIG:<br>PADCONFIG24<br>0x04084060|OSPI0_D6|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||SPI1_D0|1|IO|0|||||||||
|||MCASP1_ACLKX|2|IO|0|||||||||
|||UART6_RTSn|3|O||||||||||
|||I2C3_SCL|4|IOD|1|||||||||
|||UART4_RXD|5|I|1|||||||||
|||GPIO0_9|7|IO|pad|||||||||
|F20|OSPI0_D7<br>PADCONFIG:<br>PADCONFIG25<br>0x04084064|OSPI0_D7|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS1|Yes|1P8-LVCMOS|PU/PD|
|||SPI1_D1|1|IO|0|||||||||
|||MCASP1_AFSX|2|IO|0|||||||||
|||UART6_CTSn|3|I|1|||||||||
|||I2C3_SDA|4|IOD|1|||||||||
|||UART4_TXD|5|O||||||||||
|||GPIO0_10|7|IO|pad|||||||||
|AA18|PMIC_LPM_EN0|PMIC_LPM_EN0||O||NA / Off / Up|NA / SS / Off||1.8V|VDDS_RTC||RTC-LVCMOS|PU|
|AB18|PORz<br>PADCONFIG:<br>PADCONFIG7<br>0x0408401C|PORz|0|I||||0|1.8V|VDDS_OSC0|Yes|FS RESET||
|C16|RESETSTATz<br>PADCONFIG:<br>PADCONFIG144<br>0x04084240|RESETSTATz|0|O||Off / Low / Off|Off / SS / Off|0|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|E16|RESETz<br>PADCONFIG:<br>PADCONFIG143<br>0x0408423C|RESETz|0|I||On / Off / Up|On / Off / Up|0|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|Y7|RGMII1_RXC<br>PADCONFIG:<br>PADCONFIG77<br>0x04084134|RGMII1_RXC|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII1_REF_CLK|1|I|0|||||||||
|||GPIO0_60|7|IO|pad|||||||||
|Y6|RGMII1_RX_CTL<br>PADCONFIG:<br>PADCONFIG76<br>0x04084130|RGMII1_RX_CTL|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII1_RX_ER|1|I|0|||||||||
|||GPIO0_59|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

31 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|W11|RGMII1_TXC<br>PADCONFIG:<br>PADCONFIG71<br>0x0408411C|RGMII1_TXC|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII1_CRS_DV|1|I|0|||||||||
|||GPIO0_54|7|IO|pad|||||||||
|AB11|RGMII1_TX_CTL<br>PADCONFIG:<br>PADCONFIG70<br>0x04084118|RGMII1_TX_CTL|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII1_TX_EN|1|O||||||||||
|||GPIO0_53|7|IO|pad|||||||||
|AC7|RGMII2_RXC<br>PADCONFIG:<br>PADCONFIG91<br>0x0408416C|RGMII2_RXC|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII2_REF_CLK|1|I|0|||||||||
|||MCASP2_AXR1|2|IO|0|||||||||
|||GPIO0_74|7|IO|pad|||||||||
|AC8|RGMII2_RX_CTL<br>PADCONFIG:<br>PADCONFIG90<br>0x04084168|RGMII2_RX_CTL|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII2_RX_ER|1|I|0|||||||||
|||MCASP2_AXR3|2|IO|0|||||||||
|||GPIO0_73|7|IO|pad|||||||||
|Y13|RGMII2_TXC<br>PADCONFIG:<br>PADCONFIG85<br>0x04084154|RGMII2_TXC|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII2_CRS_DV|1|I|0|||||||||
|||MCASP2_AXR5|2|IO|0|||||||||
|||GPIO0_68|7|IO|pad|||||||||
|AB12|RGMII2_TX_CTL<br>PADCONFIG:<br>PADCONFIG84<br>0x04084150|RGMII2_TX_CTL|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII2_TX_EN|1|O||||||||||
|||MCASP2_AXR4|2|IO|0|||||||||
|||GPIO0_67|7|IO|pad|||||||||
|Y8|RGMII1_RD0<br>PADCONFIG:<br>PADCONFIG78<br>0x04084138|RGMII1_RD0|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII1_RXD0|1|I|0|||||||||
|||GPIO0_61|7|IO|pad|||||||||
|AA6|RGMII1_RD1<br>PADCONFIG:<br>PADCONFIG79<br>0x0408413C|RGMII1_RD1|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII1_RXD1|1|I|0|||||||||
|||GPIO0_62|7|IO|pad|||||||||
|AA8|RGMII1_RD2<br>PADCONFIG:<br>PADCONFIG80<br>0x04084140|RGMII1_RD2|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||I2C2_SCL|1|IOD|1|||||||||
|||GPMC0_A5|2|O||||||||||
|||GPIO0_63|7|IO|pad|||||||||
|W8|RGMII1_RD3<br>PADCONFIG:<br>PADCONFIG81<br>0x04084144|RGMII1_RD3|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||I2C2_SDA|1|IOD|1|||||||||
|||GPMC0_A6|2|O||||||||||
|||GPIO0_64|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

32 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|AC10|RGMII1_TD0<br>PADCONFIG:<br>PADCONFIG72<br>0x04084120|RGMII1_TD0|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII1_TXD0|1|O||||||||||
|||GPIO0_55|7|IO|pad|||||||||
|W13|RGMII1_TD1<br>PADCONFIG:<br>PADCONFIG73<br>0x04084124|RGMII1_TD1|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII1_TXD1|1|O||||||||||
|||GPIO0_56|7|IO|pad|||||||||
|Y11|RGMII1_TD2<br>PADCONFIG:<br>PADCONFIG74<br>0x04084128|RGMII1_TD2|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||GPMC0_A0|1|O||||||||||
|||GPIO0_57|7|IO|pad|||||||||
|AA11|RGMII1_TD3<br>PADCONFIG:<br>PADCONFIG75<br>0x0408412C|RGMII1_TD3|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||CLKOUT0|1|O||||||||||
|||GPIO0_58|7|IO|pad|||||||||
|AB9|RGMII2_RD0<br>PADCONFIG:<br>PADCONFIG92<br>0x04084170|RGMII2_RD0|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII2_RXD0|1|I|0|||||||||
|||MCASP2_AXR2|2|IO|0|||||||||
|||GPIO0_75|7|IO|pad|||||||||
|AC9|RGMII2_RD1<br>PADCONFIG:<br>PADCONFIG93<br>0x04084174|RGMII2_RD1|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII2_RXD1|1|I|0|||||||||
|||MCASP2_AFSR|2|IO|0|||||||||
|||MCASP2_AXR7|5|IO|0|||||||||
|||GPIO0_76|7|IO|pad|||||||||
|AB10|RGMII2_RD2<br>PADCONFIG:<br>PADCONFIG94<br>0x04084178|RGMII2_RD2|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||GPMC0_A3|1|O||||||||||
|||MCASP2_AXR0|2|IO|0|||||||||
|||SPI3_CLK|3|IO|0|||||||||
|||GPIO0_77|7|IO|pad|||||||||
|||EQEP2_A|8|I|0|||||||||
|AB8|RGMII2_RD3<br>PADCONFIG:<br>PADCONFIG95<br>0x0408417C|RGMII2_RD3|0|I|0|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||GPMC0_A4|1|O||||||||||
|||AUDIO_EXT_REFCLK0|2|IO|0|||||||||
|||SPI3_CS0|3|IO|1|||||||||
|||GPIO0_78|7|IO|pad|||||||||
|||EQEP2_B|8|I|0|||||||||



Copyright © 2025 Texas Instruments Incorporated 

33 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|AC12|RGMII2_TD0<br>PADCONFIG:<br>PADCONFIG86<br>0x04084158|RGMII2_TD0|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII2_TXD0|1|O||||||||||
|||MCASP2_AXR6|2|IO|0|||||||||
|||GPIO0_69|7|IO|pad|||||||||
|AB13|RGMII2_TD1<br>PADCONFIG:<br>PADCONFIG87<br>0x0408415C|RGMII2_TD1|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||RMII2_TXD1|1|O||||||||||
|||MCASP2_ACLKR|2|IO|0|||||||||
|||MCASP2_AXR8|5|IO|0|||||||||
|||GPIO0_70|7|IO|pad|||||||||
|AA12|RGMII2_TD2<br>PADCONFIG:<br>PADCONFIG88<br>0x04084160|RGMII2_TD2|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||GPMC0_A1|1|O||||||||||
|||MCASP2_AFSX|2|IO|0|||||||||
|||SPI3_D0|3|IO|0|||||||||
|||GPIO0_71|7|IO|pad|||||||||
|||EQEP2_I|8|IO|0|||||||||
|AA13|RGMII2_TD3<br>PADCONFIG:<br>PADCONFIG89<br>0x04084164|RGMII2_TD3|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|||GPMC0_A2|1|O||||||||||
|||MCASP2_ACLKX|2|IO|0|||||||||
|||SPI3_D1|3|IO|0|||||||||
|||GPIO0_72|7|IO|pad|||||||||
|||EQEP2_S|8|IO|0|||||||||
|AB17|RSVD0|RSVD0||N/A||||||||||
|Y18|RTC_PORz|RTC_PORz||I||On / NA / NA|On / NA / NA||1.8V|VDDS_RTC|Yes|RTC-LVCMOS||
|E13|SPI0_CLK<br>PADCONFIG:<br>PADCONFIG106<br>0x040841A8|SPI0_CLK|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||CP_GEMAC_CPTS0_TS_SYNC|1|O||||||||||
|||EHRPWM1_A|2|IO|0|||||||||
|||GPIO0_89|7|IO|pad|||||||||
|E11|SPI0_CS0<br>PADCONFIG:<br>PADCONFIG104<br>0x040841A0|SPI0_CS0|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||EHRPWM0_A|2|IO|0|||||||||
|||GPIO0_87|7|IO|pad|||||||||
|D11|SPI0_CS1<br>PADCONFIG:<br>PADCONFIG105<br>0x040841A4|SPI0_CS1|0|IO|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||CP_GEMAC_CPTS0_TS_COMP|1|O||||||||||
|||EHRPWM0_B|2|IO|0|||||||||
|||ECAP0_IN_APWM_OUT|3|IO|0|||||||||
|||AUDIO_EXT_REFCLK1|4|IO|0|||||||||
|||GPIO0_88|7|IO|pad|||||||||
|||EHRPWM_TZn_IN5|9|I|0|||||||||



Copyright © 2025 Texas Instruments Incorporated 

34 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|E12|SPI0_D0<br>PADCONFIG:<br>PADCONFIG107<br>0x040841AC|SPI0_D0|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||CP_GEMAC_CPTS0_HW1TSPUSH|1|I|0|||||||||
|||EHRPWM1_B|2|IO|0|||||||||
|||GPIO0_90|7|IO|pad|||||||||
|B12|SPI0_D1<br>PADCONFIG:<br>PADCONFIG108<br>0x040841B0|SPI0_D1|0|IO|0|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||CP_GEMAC_CPTS0_HW2TSPUSH|1|I|0|||||||||
|||EHRPWM_TZn_IN0|2|I|0|||||||||
|||GPIO0_91|7|IO|pad|||||||||
|AB14|TCK<br>PADCONFIG:<br>PADCONFIG8<br>0x04084020|TCK|0|I||On / Off / Up|On / Off / Up|0|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|AC16|TDI<br>PADCONFIG:<br>PADCONFIG10<br>0x04084028|TDI|0|I||On / Off / Up|On / Off / Up|0|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|AB15|TDO<br>PADCONFIG:<br>PADCONFIG11<br>0x0408402C|TDO|0|OZ||Off / Off / Up|Off / SS / Up|0|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|Y17|TMS<br>PADCONFIG:<br>PADCONFIG12<br>0x04084030|TMS|0|I||On / Off / Up|On / Off / Up|0|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|AB16|TRSTn<br>PADCONFIG:<br>PADCONFIG9<br>0x04084024|TRSTn|0|I||On / Off / Down|On / Off / Down|0|1.8V|VDDS0|Yes|1P8-LVCMOS|PU/PD|
|B14|UART0_CTSn<br>PADCONFIG:<br>PADCONFIG111<br>0x040841BC|UART0_CTSn|0|I|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI0_CS2|1|IO|1|||||||||
|||I2C3_SCL|2|IOD|1|||||||||
|||UART2_RXD|3|I|1|||||||||
|||TIMER_IO2|4|IO|0|||||||||
|||AUDIO_EXT_REFCLK0|5|IO|0|||||||||
|||MCAN2_RX|6|I|1|||||||||
|||GPIO0_94|7|IO|pad|||||||||
|||MCASP2_AFSX|8|IO|0|||||||||
|||MMC2_SDCD|9|I|0|||||||||



Copyright © 2025 Texas Instruments Incorporated 

35 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|B13|UART0_RTSn<br>PADCONFIG:<br>PADCONFIG112<br>0x040841C0|UART0_RTSn|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||SPI0_CS3|1|IO|1|||||||||
|||I2C3_SDA|2|IOD|1|||||||||
|||UART2_TXD|3|O||||||||||
|||TIMER_IO3|4|IO|0|||||||||
|||AUDIO_EXT_REFCLK1|5|IO|0|||||||||
|||MCAN2_TX|6|O||||||||||
|||GPIO0_95|7|IO|pad|||||||||
|||MCASP2_ACLKX|8|IO|0|||||||||
|||MMC2_SDWP|9|I|0|||||||||
|D13|UART0_RXD<br>PADCONFIG:<br>PADCONFIG109<br>0x040841B4|UART0_RXD|0|I|1|Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||ECAP1_IN_APWM_OUT|1|IO|0|||||||||
|||SPI2_D0|2|IO|0|||||||||
|||EHRPWM2_A|3|IO|0|||||||||
|||GPIO0_92|7|IO|pad|||||||||
|C13|UART0_TXD<br>PADCONFIG:<br>PADCONFIG110<br>0x040841B8|UART0_TXD|0|O||Off / Off / Off|Off / Off / Off|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||ECAP2_IN_APWM_OUT|1|IO|0|||||||||
|||SPI2_D1|2|IO|0|||||||||
|||EHRPWM2_B|3|IO|0|||||||||
|||GPIO0_93|7|IO|pad|||||||||
|AC4|USB0_DM|USB0_DM||IO|||||1.8V / 3.3V|VDDA_1P8_USB,<br>VDDA_3P3_USB||USB2PHY||
|AB4|USB0_DP|USB0_DP||IO|||||1.8V / 3.3V|VDDA_1P8_USB,<br>VDDA_3P3_USB||USB2PHY||
|C6|USB0_DRVVBUS<br>PADCONFIG:<br>PADCONFIG145<br>0x04084244|USB0_DRVVBUS|0|O||Off / Off / Down|Off / Off / Down|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||GPIO0_124|7|IO|pad|||||||||
|AB3|USB0_RCALIB|USB0_RCALIB||IO|||||1.8V / 3.3V|VDDA_1P8_USB,<br>VDDA_3P3_USB||USB2PHY||
|AC3|USB0_VBUS|USB0_VBUS||A|||||1.8V / 3.3V|VDDA_1P8_USB,<br>VDDA_3P3_USB||USB2PHY||
|AC5|USB1_DM|USB1_DM||IO|||||1.8V / 3.3V|VDDA_1P8_USB,<br>VDDA_3P3_USB||USB2PHY||
|AB5|USB1_DP|USB1_DP||IO|||||1.8V / 3.3V|VDDA_1P8_USB,<br>VDDA_3P3_USB||USB2PHY||
|A5|USB1_DRVVBUS<br>PADCONFIG:<br>PADCONFIG146<br>0x04084248|USB1_DRVVBUS|0|O||Off / Off / Down|Off / Off / Down|7|1.8V / 3.3V|VDDSHV1|Yes|LVCMOS|PU/PD|
|||GPIO0_125|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

36 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|AC6|USB1_RCALIB|USB1_RCALIB||IO|||||1.8V / 3.3V|VDDA_1P8_USB,<br>VDDA_3P3_USB||USB2PHY||
|AB6|USB1_VBUS|USB1_VBUS||A|||||1.8V / 3.3V|VDDA_1P8_USB,<br>VDDA_3P3_USB||USB2PHY||
|G14|VDDA_1P8_DSI|VDDA_1P8_DSI||PWR||||||||||
|T12|VDDA_1P8_USB|VDDA_1P8_USB||PWR||||||||||
|U16|VDDA_3P3_SDIO|VDDA_3P3_SDIO||PWR||||||||||
|U12|VDDA_3P3_USB|VDDA_3P3_USB||PWR||||||||||
|N17|VDDA_ADC|VDDA_ADC||PWR||||||||||
|G13|VDDA_CORE_DSI|VDDA_CORE_DSI||PWR||||||||||
|H12|VDDA_CORE_DSI_CLK|VDDA_CORE_DSI_CLK||PWR||||||||||
|U11|VDDA_CORE_USB|VDDA_CORE_USB||PWR||||||||||
|M10|VDDA_DDR_PLL0|VDDA_DDR_PLL0||PWR||||||||||
|L11|VDDA_PLL0|VDDA_PLL0||PWR||||||||||
|K12|VDDA_PLL1|VDDA_PLL1||PWR||||||||||
|T14|VDDS0|VDDS0||PWR||||||||||
|H16|VDDS1|VDDS1||PWR||||||||||
|J16, L17|VDDSHV0|VDDSHV0||PWR||||||||||
|G10, H10|VDDSHV1|VDDSHV1||PWR||||||||||
|H8|VDDSHV2|VDDSHV2||PWR||||||||||
|T10|VDDSHV3|VDDSHV3||PWR||||||||||
|M17|VDDSHV4|VDDSHV4||PWR||||||||||
|L8, M7,<br>M8, N8, P8|VDDS_DDR|VDDS_DDR||PWR||||||||||
|R16|VDDS_OSC0|VDDS_OSC0||PWR||||||||||
|T18|VDDS_RTC|VDDS_RTC||PWR||||||||||
|P16|VDDS_WKUP|VDDS_WKUP||PWR||||||||||
|J11, J13,<br>J15, J9,<br>K10, K14,<br>L15, M14,<br>N15, P10,<br>P12, P14,<br>R11, R9|VDD_CORE|VDD_CORE||PWR||||||||||
|T17|VDD_RTC|VDD_RTC||PWR||||||||||
|N18|VPP|VPP||PWR||||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 37 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|A1, A10,<br>A13, A16,<br>A19, A2,<br>A22, A23,<br>A4, AA20,<br>AA4, AB1,<br>AB21,<br>AB23, AB7,<br>AC1, AC11,<br>AC14,<br>AC19,<br>AC2,<br>AC22,<br>AC23, B1,<br>B17, B20,<br>B23, B5,<br>C12, C18,<br>D1, E10,<br>E14, E15,<br>E2, E6, E8,<br>E9, F18,<br>F5, F6,<br>G12, G15,<br>G16, G17,<br>G7, G8,<br>G9, H1,<br>H14, H17,<br>H7, K15,<br>K8, K9,<br>L13, L16,<br>L18, L7,<br>L9, M1,<br>M12, N11,<br>N13, N16,<br>N7, N9,<br>P15, P9,<br>R1, R13,<br>R15, R8,<br>T19, T2,<br>T7, T8,<br>U10, U13,<br>U14, U15,<br>U17, U20,<br>U7, U8,<br>V18, V19,<br>V3, W10,<br>W12, W14,<br>W15, W16,<br>W18, W9,<br>Y1, Y20,<br>Y21|VSS|VSS||PWR||||||||||
|Y23|WKUP_CLKOUT0<br>PADCONFIG:<br>PADCONFIG6<br>0x04084018|WKUP_CLKOUT0|0|O||Off / Off / Off|Off / SS / Off|0|1.8V|VDDS_WKUP|Yes|1P8-LVCMOS|PU/PD|
|||WKUP_GPIO0_6|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

38 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-1. Pin Attributes (ANB Package) (continued)** 

|**BALL**<br>**NUMBER**<br>**[**1**]**|**BALL NAME [**2**]**<br>**PADCONFIG Register [**15**]**<br>**PADCONFIG Address [**16**]**|**SIGNAL NAME [**3**]**|**MUX**<br>**MODE [**4**]**|**TYPE**<br>**[**5**]**|**DSIS**<br>**[**6**]**|**BALL**<br>**STATE**<br>**DURING**<br>**RESET**<br>**(RX/TX/PULL) [**7**]**|**BALL**<br>**STATE**<br>**AFTER**<br>**RESET**<br>**(RX/TX/PULL) [**8**]**|**MUX**<br>**MODE**<br>**AFTER**<br>**RESET**<br>**[**9**]**|**I/O**<br>**OPERATING**<br>**VOLTAGE [**10**]**|**POWER [**11**]**|**HYS**<br>**[**12**]**|**BUFFER**<br>**TYPE [**13**]**|**PULL**<br>**UP/DOWN**<br>**TYPE [**14**]**|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|AB22|WKUP_I2C0_SCL<br>PADCONFIG:<br>PADCONFIG4<br>0x04084010|WKUP_I2C0_SCL|0|IOD|1|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS_WKUP|Yes|1P8-LVCMOS|PU/PD|
|||WKUP_GPIO0_4|7|IO|pad|||||||||
|AA22|WKUP_I2C0_SDA<br>PADCONFIG:<br>PADCONFIG5<br>0x04084014|WKUP_I2C0_SDA|0|IOD|1|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS_WKUP|Yes|1P8-LVCMOS|PU/PD|
|||WKUP_GPIO0_5|7|IO|pad|||||||||
|AC18|WKUP_OSC0_XI|WKUP_OSC0_XI||I|||||1.8V|VDDS_OSC0||HFOSC||
|AC17|WKUP_OSC0_XO|WKUP_OSC0_XO||O|||||1.8V|VDDS_OSC0||HFOSC||
|W23|WKUP_UART0_CTSn<br>PADCONFIG:<br>PADCONFIG2<br>0x04084008|WKUP_UART0_CTSn|0|I|1|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS_WKUP|Yes|1P8-LVCMOS|PU/PD|
|||WKUP_TIMER_IO0|1|IO|0|||||||||
|||WKUP_OBSCLK0|2|O||||||||||
|||WKUP_SYSCLKOUT0|3|O||||||||||
|||WKUP_GPIO0_2|7|IO|pad|||||||||
|W22|WKUP_UART0_RTSn<br>PADCONFIG:<br>PADCONFIG3<br>0x0408400C|WKUP_UART0_RTSn|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS_WKUP|Yes|1P8-LVCMOS|PU/PD|
|||WKUP_TIMER_IO1|1|IO|0|||||||||
|||WKUP_EXT_REFCLK0|2|I|0|||||||||
|||WKUP_GPIO0_3|7|IO|pad|||||||||
|Y22|WKUP_UART0_RXD<br>PADCONFIG:<br>PADCONFIG0<br>0x04084000|WKUP_UART0_RXD|0|I|1|Off / Off / Off|Off / Off / Off|7|1.8V|VDDS_WKUP|Yes|1P8-LVCMOS|PU/PD|
|||WKUP_GPIO0_0|7|IO|pad|||||||||
|AA23|WKUP_UART0_TXD<br>PADCONFIG:<br>PADCONFIG1<br>0x04084004|WKUP_UART0_TXD|0|O||Off / Off / Off|Off / Off / Off|7|1.8V|VDDS_WKUP|Yes|1P8-LVCMOS|PU/PD|
|||WKUP_GPIO0_1|7|IO|pad|||||||||



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 39 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **5.3 Signal Descriptions** 

Many signals are available on multiple pins, according to the software configuration of the pin multiplexing options. 

The following list describes the column headers: 

1. **SIGNAL NAME:** The name of the signal passing through the pin. 

## **Note** 

Signal names and descriptions provided in each Signal Descriptions table, represent the pin multiplexed signal function which is implemented at the pin and selected via PADCONFIG registers. Device subsystems may provide secondary multiplexing of signal functions, which are not described in these tables. For more information on secondary multiplexed signal functions, see the respective peripheral chapter of the device TRM. 

2. **PIN TYPE:** Signal direction and type: 

   - I = Input 

   - O = Output 

   - OD = Output, with open-drain output function 

   - IO = Input, Output, or simultaneously Input and Output 

   - IOD = Input, Output, or simultaneously Input and Output with open-drain output function 

   - IOZ = Input, Output, or simultaneously Input and Output with three-state output function 

   - OZ = Output with three-state output function 

   - A = Analog 

   - PWR = Power 

   - GND = Ground 

   - CAP = LDO Capacitor 

3. **DESCRIPTION:** Description of the signal 

4. **BALL:** Ball number(s) associated with signal 

For more information on the IO cell configurations, see the _Pad Configuration Registers_ section in _Device Configuration_ chapter of the device TRM. 

## _**5.3.1 ADC**_ 

## **5.3.1.1 MAIN Domain** 

**Table 5-2. ADC0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|ADC0_AIN0|A|ADC Analog Input 0|V20|
|ADC0_AIN1|A|ADC Analog Input 1|V22|
|ADC0_AIN2|A|ADC Analog Input 2|V23|
|ADC0_AIN3|A|ADC Analog Input 3|V21|
|ADC_EXT_TRIGGER0(1)|I|ADC Trigger Input|C11,D16,D18,<br>M23|
|ADC_EXT_TRIGGER1(1)|I|ADC Trigger Input|A12,C23,M22|



(1) This ADC input signal has a debounce function. For more information on I/O Debounce configuration, see the TRM _Device Configuration_ chapter. 

Copyright © 2025 Texas Instruments Incorporated 

40 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**5.3.2 CPSW3G**_ 

## **5.3.2.1 MAIN Domain** 

**Table 5-3. CPSW3G0 RGMII1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|RGMII1_RXC|I|RGMII Receive Clock|Y7|
|RGMII1_RX_CTL|I|RGMII Receive Control|Y6|
|RGMII1_TXC|O|RGMII Transmit Clock|W11|
|RGMII1_TX_CTL|O|RGMII Transmit Control|AB11|
|RGMII1_RD0|I|RGMII Receive Data 0|Y8|
|RGMII1_RD1|I|RGMII Receive Data 1|AA6|
|RGMII1_RD2|I|RGMII Receive Data 2|AA8|
|RGMII1_RD3|I|RGMII Receive Data 3|W8|
|RGMII1_TD0|O|RGMII Transmit Data 0|AC10|
|RGMII1_TD1|O|RGMII Transmit Data 1|W13|
|RGMII1_TD2|O|RGMII Transmit Data 2|Y11|
|RGMII1_TD3|O|RGMII Transmit Data 3|AA11|



**Table 5-4. CPSW3G0 RGMII2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|RGMII2_RXC|I|RGMII Receive Clock|AC7|
|RGMII2_RX_CTL|I|RGMII Receive Control|AC8|
|RGMII2_TXC|O|RGMII Transmit Clock|Y13|
|RGMII2_TX_CTL|O|RGMII Transmit Control|AB12|
|RGMII2_RD0|I|RGMII Receive Data 0|AB9|
|RGMII2_RD1|I|RGMII Receive Data 1|AC9|
|RGMII2_RD2|I|RGMII Receive Data 2|AB10|
|RGMII2_RD3|I|RGMII Receive Data 3|AB8|
|RGMII2_TD0|O|RGMII Transmit Data 0|AC12|
|RGMII2_TD1|O|RGMII Transmit Data 1|AB13|
|RGMII2_TD2|O|RGMII Transmit Data 2|AA12|
|RGMII2_TD3|O|RGMII Transmit Data 3|AA13|



## **Table 5-5. CPSW3G0 RMII1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|RMII1_CRS_DV|I|RMII Carrier Sense / Data Valid|W11|
|RMII1_REF_CLK|I|RMII Reference Clock|Y7|
|RMII1_RX_ER|I|RMII Receive Data Error|Y6|
|RMII1_TX_EN|O|RMII Transmit Enable|AB11|
|RMII1_RXD0|I|RMII Receive Data 0|Y8|
|RMII1_RXD1|I|RMII Receive Data 1|AA6|
|RMII1_TXD0|O|RMII Transmit Data 0|AC10|
|RMII1_TXD1|O|RMII Transmit Data 1|W13|



## **Table 5-6. CPSW3G0 RMII2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|RMII2_CRS_DV|I|RMII Carrier Sense / Data Valid|Y13|
|RMII2_REF_CLK|I|RMII Reference Clock|AC7|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 41 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**Table 5-6. CPSW3G0 RMII2 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|RMII2_RX_ER|I|RMII Receive Data Error|AC8|
|RMII2_TX_EN|O|RMII Transmit Enable|AB12|
|RMII2_RXD0|I|RMII Receive Data 0|AB9|
|RMII2_RXD1|I|RMII Receive Data 1|AC9|
|RMII2_TXD0|O|RMII Transmit Data 0|AC12|
|RMII2_TXD1|O|RMII Transmit Data 1|AB13|



## _**5.3.3 CPTS**_ 

## **5.3.3.1 MAIN Domain** 

**Table 5-7. CPTS Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|CP_GEMAC_CPTS0_RFT_CLK|I|CPTS Reference Clock Input|D16|
|CP_GEMAC_CPTS0_TS_COMP|O|CPTS Time Stamp Counter Compare Output from<br>CPSW3G0 CPTS|AB2,D11|
|CP_GEMAC_CPTS0_TS_SYNC|O|CPTS Time Stamp Counter Bit Output from CPSW3G0<br>CPTS|AA2,E13|
|CP_GEMAC_CPTS0_HW1TSPUSH|I|CPTS Hardware Time Stamp Push Input to Time Sync<br>Router|E12,Y4|
|CP_GEMAC_CPTS0_HW2TSPUSH|I|CPTS Hardware Time Stamp Push Input to Time Sync<br>Router|AA1,B12|
|SYNC0_OUT|O|CPTS Time Stamp Generator Bit 0 Output from Time<br>Sync Router|B7|
|SYNC1_OUT|O|CPTS Time Stamp Generator Bit 1 Output from Time<br>Sync Router|D16|
|SYNC2_OUT|O|CPTS Time Stamp Generator Bit 2 Output from Time<br>Sync Router|B16|
|SYNC3_OUT|O|CPTS Time Stamp Generator Bit 3 Output from Time<br>Sync Router|B15|



## _**5.3.4 DDRSS**_ 

## **5.3.4.1 MAIN Domain** 

**Table 5-8. DDRSS0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|DDR0_ACT_n|O|DDRSS Activation Command|M2|
|DDR0_CAS_n|O|DDRSS Column Address Strobe|L1|
|DDR0_RAS_n|O|DDRSS Row Address Strobe|M5|
|DDR0_WE_n|O|DDRSS Write Enable|L2|
|DDR0_A0|O|DDRSS Address Bus|L5|
|DDR0_A1|O|DDRSS Address Bus|H6|
|DDR0_A2|O|DDRSS Address Bus|L6|
|DDR0_A3|O|DDRSS Address Bus|K2|
|DDR0_A4|O|DDRSS Address Bus|J1|
|DDR0_A5|O|DDRSS Address Bus|H5|
|DDR0_A6|O|DDRSS Address Bus|R2|
|DDR0_A7|O|DDRSS Address Bus|N6|
|DDR0_A8|O|DDRSS Address Bus|T4|
|DDR0_A9|O|DDRSS Address Bus|N1|
|DDR0_A10|O|DDRSS Address Bus|T5|



42 _Submit Document Feedback_ 

Copyright © 2025 Texas Instruments Incorporated 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-8. DDRSS0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|DDR0_A11|O|DDRSS Address Bus|T6|
|DDR0_A12|O|DDRSS Address Bus|W6|
|DDR0_A13|O|DDRSS Address Bus|V6|
|DDR0_BA0|O|DDRSS Bank Address|N3|
|DDR0_BA1|O|DDRSS Bank Address|N2|
|DDR0_BG0|O|DDRSS Bank Group|N5|
|DDR0_BG1|O|DDRSS Bank Group|N4|
|DDR0_CAL0(1)|A|IO Pad Calibration Resistor|M3|
|DDR0_CK0|O|DDRSS Clock|P1|
|DDR0_CK0_n|O|DDRSS Negative Clock|P2|
|DDR0_CKE0|O|DDRSS Clock Enable|K1|
|DDR0_CS0_n|O|DDRSS Chip Select 0|L3|
|DDR0_DM0|IO|DDRSS Data Mask|F2|
|DDR0_DM1|IO|DDRSS Data Mask|W2|
|DDR0_DQ0|IO|DDRSS Data|F4|
|DDR0_DQ1|IO|DDRSS Data|F3|
|DDR0_DQ2|IO|DDRSS Data|F1|
|DDR0_DQ3|IO|DDRSS Data|E1|
|DDR0_DQ4|IO|DDRSS Data|G4|
|DDR0_DQ5|IO|DDRSS Data|H4|
|DDR0_DQ6|IO|DDRSS Data|H2|
|DDR0_DQ7|IO|DDRSS Data|H3|
|DDR0_DQ8|IO|DDRSS Data|V4|
|DDR0_DQ9|IO|DDRSS Data|T3|
|DDR0_DQ10|IO|DDRSS Data|T1|
|DDR0_DQ11|IO|DDRSS Data|U1|
|DDR0_DQ12|IO|DDRSS Data|U4|
|DDR0_DQ13|IO|DDRSS Data|V5|
|DDR0_DQ14|IO|DDRSS Data|U2|
|DDR0_DQ15|IO|DDRSS Data|W1|
|DDR0_DQS0|IO|DDRSS Data Strobe|G1|
|DDR0_DQS0_n|IO|DDRSS Complimentary Data Strobe|G2|
|DDR0_DQS1|IO|DDRSS Data Strobe|V1|
|DDR0_DQS1_n|IO|DDRSS Complimentary Data Strobe|V2|
|DDR0_ODT0|O|DDRSS On-Die Termination for Chip Select 0|L4|
|DDR0_RESET0_n|O|DDRSS Reset|J2|



(1) An external 240Ω ±1% resistor must be connected between this pin and VSS. No external voltage should be applied to this pin. 

## _**5.3.5 DSI**_ 

## **5.3.5.1 MAIN Domain** 

**Table 5-9. DSITX0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|DSI0_TXCLKN|IO|DSI Differential Transmit Clock Output (negative)|A15|
|DSI0_TXCLKP|IO|DSI Differential Transmit Clock Output (positive)|A14|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 43 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**Table 5-9. DSITX0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|DSI0_TXRCALIB(1)|A|DSI pin connected to external resistor for on-chip<br>resistor calibration|D17|
|DSI0_TXN0|IO|DSI Differential Transmit Output (negative)|B19|
|DSI0_TXN1|IO|DSI Differential Transmit Output (negative)|A18|
|DSI0_TXN2|IO|DSI Differential Transmit Output (negative)|A20|
|DSI0_TXN3|IO|DSI Differential Transmit Output (negative)|B22|
|DSI0_TXP0|IO|DSI Differential Transmit Output (positive)|B18|
|DSI0_TXP1|IO|DSI Differential Transmit Output (positive)|A17|
|DSI0_TXP2|IO|DSI Differential Transmit Output (positive)|A21|
|DSI0_TXP3|IO|DSI Differential Transmit Output (positive)|B21|



(1) An external 499Ω ±1% resistor must be connected between this pin and VSS and the maximum power dissipation for the resistor is 7.2mW. No external voltage should be applied to this pin. 

## _**5.3.6 DSS**_ 

## **5.3.6.1 MAIN Domain** 

**Table 5-10. DSS0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|VOUT0_DE|O|Video Output Data Enable|M21|
|VOUT0_EXTPCLKIN|I|Video Output External Pixel Clock Input|M23|
|VOUT0_HSYNC|O|Video Output Horizontal Sync|N21|
|VOUT0_PCLK|O|Video Output Pixel Clock Output|L19|
|VOUT0_VSYNC|O|Video Output Vertical Sync|L20|
|VOUT0_DATA0|O|Video Output Data 0|L22|
|VOUT0_DATA1|O|Video Output Data 1|L23|
|VOUT0_DATA2|O|Video Output Data 2|K22|
|VOUT0_DATA3|O|Video Output Data 3|J23|
|VOUT0_DATA4|O|Video Output Data 4|K23|
|VOUT0_DATA5|O|Video Output Data 5|H22|
|VOUT0_DATA6|O|Video Output Data 6|H23|
|VOUT0_DATA7|O|Video Output Data 7|J22|
|VOUT0_DATA8|O|Video Output Data 8|H19|
|VOUT0_DATA9|O|Video Output Data 9|H20|
|VOUT0_DATA10|O|Video Output Data 10|H21|
|VOUT0_DATA11|O|Video Output Data 11|H18|
|VOUT0_DATA12|O|Video Output Data 12|G23|
|VOUT0_DATA13|O|Video Output Data 13|G22|
|VOUT0_DATA14|O|Video Output Data 14|F22|
|VOUT0_DATA15|O|Video Output Data 15|F23|
|VOUT0_DATA16|O|Video Output Data 16|L21|
|VOUT0_DATA17|O|Video Output Data 17|N19|
|VOUT0_DATA18|O|Video Output Data 18|N20|
|VOUT0_DATA19|O|Video Output Data 19|M19|
|VOUT0_DATA20|O|Video Output Data 20|P23|
|VOUT0_DATA21|O|Video Output Data 21|P22|
|VOUT0_DATA22|O|Video Output Data 22|N23|



Copyright © 2025 Texas Instruments Incorporated 

44 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 5-10. DSS0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|VOUT0_DATA23|O|Video Output Data 23|N22|



## _**5.3.7 ECAP**_ 

## **5.3.7.1 MAIN Domain** 

## **Table 5-11. ECAP0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|ECAP0_IN_APWM_OUT|IO|Enhanced Capture (ECAP) Input or Auxiliary PWM<br>(APWM) Output|D11,D16,L22|



## **Table 5-12. ECAP1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|ECAP1_IN_APWM_OUT|IO|Enhanced Capture (ECAP) Input or Auxiliary PWM<br>(APWM) Output|A8,A9,B7,D13,<br>L23,Y4|



## **Table 5-13. ECAP2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|ECAP2_IN_APWM_OUT|IO|Enhanced Capture (ECAP) Input or Auxiliary PWM<br>(APWM) Output|A11,A7,AA1,<br>B10,C13,H20|



## _**5.3.8 Emulation and Debug**_ 

## **5.3.8.1 MAIN Domain** 

## **Table 5-14. Trace Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|TRC_CLK|O|Trace Clock|L22|
|TRC_CTL|O|Trace Control|L23|
|TRC_DATA0|O|Trace Data 0|K22|
|TRC_DATA1|O|Trace Data 1|J23|
|TRC_DATA2|O|Trace Data 2|K23|
|TRC_DATA3|O|Trace Data 3|H22|
|TRC_DATA4|O|Trace Data 4|H23|
|TRC_DATA5|O|Trace Data 5|J22|
|TRC_DATA6|O|Trace Data 6|L21|
|TRC_DATA7|O|Trace Data 7|N19|
|TRC_DATA8|O|Trace Data 8|N20|
|TRC_DATA9|O|Trace Data 9|M19|
|TRC_DATA10|O|Trace Data 10|P23|
|TRC_DATA11|O|Trace Data 11|P22|
|TRC_DATA12|O|Trace Data 12|N23|
|TRC_DATA13|O|Trace Data 13|N21|
|TRC_DATA14|O|Trace Data 14|M21|
|TRC_DATA15|O|Trace Data 15|L20|
|TRC_DATA16|O|Trace Data 16|L19|
|TRC_DATA17|O|Trace Data 17|M23|
|TRC_DATA18|O|Trace Data 18|M22|
|TRC_DATA19|O|Trace Data 19|F23|
|TRC_DATA20|O|Trace Data 20|F22|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 45 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

## **Table 5-14. Trace Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|TRC_DATA21|O|Trace Data 21|G22|
|TRC_DATA22|O|Trace Data 22|G23|
|TRC_DATA23|O|Trace Data 23|H18|



## **5.3.8.2 WKUP Domain** 

**Table 5-15. JTAG Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EMU0|IO|Emulation Control 0|Y16|
|EMU1|IO|Emulation Control 1|AA16|
|TCK|I|JTAG Test Clock Input|AB14|
|TDI|I|JTAG Test Data Input|AC16|
|TDO|OZ|JTAG Test Data Output|AB15|
|TMS|I|JTAG Test Mode Select Input|Y17|
|TRSTn|I|JTAG Reset|AB16|



## _**5.3.9 EPWM**_ 

## **5.3.9.1 MAIN Domain** 

**Table 5-16. EPWM Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EHRPWM_SOCA|O|EHRPWM Start of Conversion A|B7|
|EHRPWM_SOCB|O|EHRPWM Start of Conversion B|A7|
|EHRPWM_TZn_IN0|I|EHRPWM Trip Zone Input 0 (active low)|B12|
|EHRPWM_TZn_IN1|I|EHRPWM Trip Zone Input 1 (active low)|D3,N19|
|EHRPWM_TZn_IN2|I|EHRPWM Trip Zone Input 2 (active low)|B3,G23|
|EHRPWM_TZn_IN3|I|EHRPWM Trip Zone Input 3 (active low)|B16|
|EHRPWM_TZn_IN4|I|EHRPWM Trip Zone Input 4 (active low)|B15|
|EHRPWM_TZn_IN5|I|EHRPWM Trip Zone Input 5 (active low)|D11|



## **Table 5-17. EPWM0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EHRPWM0_A|IO|EHRPWM Output A|C11,C4,E11,<br>G22|
|EHRPWM0_B|IO|EHRPWM Output B|A12,C2,D11,<br>F22|
|EHRPWM0_SYNCI|I|Sync Input to EHRPWM module from an external pin|B4,D7,H21|
|EHRPWM0_SYNCO|O|Sync Output from EHRPWM module to an external pin|A3,A6,H18|



## **Table 5-18. EPWM1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EHRPWM1_A|IO|EHRPWM Output A|A9,C1,E13,F23|
|EHRPWM1_B|IO|EHRPWM Output B|B9,D4,E12,L21|
|**Table 5-19. EPWM2 Signal Descriptions**||||
|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|EHRPWM2_A|IO|EHRPWM Output A|B2,D13,D7,<br>N20|



Copyright © 2025 Texas Instruments Incorporated 

46 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 5-19. EPWM2 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EHRPWM2_B|IO|EHRPWM Output B|A6,C13,D2,<br>M19|



## _**5.3.10 EQEP**_ 

## **5.3.10.1 MAIN Domain** 

**Table 5-20. EQEP0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EQEP0_A(1)|I|EQEP Quadrature Input A|A8,P23|
|EQEP0_B(1)|I|EQEP Quadrature Input B|B10,P22|
|EQEP0_I(1)|IO|EQEP Index|B9,N22|
|EQEP0_S(1)|IO|EQEP Strobe|A9,N23|



(1) This EQEP input signal has a debounce function. For more information on I/O Debounce configuration, see the TRM _Device Configuration_ chapter. 

## **Table 5-21. EQEP1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EQEP1_A(1)|I|EQEP Quadrature Input A|A11,N21|
|EQEP1_B(1)|I|EQEP Quadrature Input B|B11,M21|
|EQEP1_I(1)|IO|EQEP Index|A12,L19|
|EQEP1_S(1)|IO|EQEP Strobe|C11,L20|



(1) This EQEP input signal has a debounce function. For more information on I/O Debounce configuration, see the TRM _Device Configuration_ chapter. 

**Table 5-22. EQEP2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EQEP2_A(1)|I|EQEP Quadrature Input A|AB10,B7|
|EQEP2_B(1)|I|EQEP Quadrature Input B|A7,AB8|
|EQEP2_I(1)|IO|EQEP Index|AA12,B16,N22|
|EQEP2_S(1)|IO|EQEP Strobe|AA13,B15,M21|



(1) This EQEP input signal has a debounce function. For more information on I/O Debounce configuration, see the TRM _Device Configuration_ chapter. 

## _**5.3.11 GPIO**_ 

## **5.3.11.1 MAIN Domain** 

**Table 5-23. GPIO0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|GPIO0_0|IO|General Purpose Input/Output|D22|
|GPIO0_1|IO|General Purpose Input/Output|E18|
|GPIO0_2|IO|General Purpose Input/Output|E22|
|GPIO0_3|IO|General Purpose Input/Output|C22|
|GPIO0_4|IO|General Purpose Input/Output|D21|
|GPIO0_5|IO|General Purpose Input/Output|E23|
|GPIO0_6|IO|General Purpose Input/Output|D23|
|GPIO0_7|IO|General Purpose Input/Output|F21|
|GPIO0_8|IO|General Purpose Input/Output|F19|
|GPIO0_9|IO|General Purpose Input/Output|G20|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 47 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**Table 5-23. GPIO0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|GPIO0_10|IO|General Purpose Input/Output|F20|
|GPIO0_11|IO|General Purpose Input/Output|C20|
|GPIO0_12|IO|General Purpose Input/Output|D20|
|GPIO0_13(1)|IO|General Purpose Input/Output|D18|
|GPIO0_14(1)|IO|General Purpose Input/Output|C23|
|GPIO0_15|IO|General Purpose Input/Output|L22|
|GPIO0_16|IO|General Purpose Input/Output|L23|
|GPIO0_17|IO|General Purpose Input/Output|K22|
|GPIO0_18|IO|General Purpose Input/Output|J23|
|GPIO0_19|IO|General Purpose Input/Output|K23|
|GPIO0_100|IO|General Purpose Input/Output|D7|
|GPIO0_101|IO|General Purpose Input/Output|A6|
|GPIO0_102|IO|General Purpose Input/Output|B8|
|GPIO0_103|IO|General Purpose Input/Output|D8|
|GPIO0_104|IO|General Purpose Input/Output|D16|
|GPIO0_105(1)|IO|General Purpose Input/Output|C8|
|GPIO0_106(1)|IO|General Purpose Input/Output|B4|
|GPIO0_107(1)|IO|General Purpose Input/Output|A3|
|GPIO0_108(1)|IO|General Purpose Input/Output|B3|
|GPIO0_109(1)|IO|General Purpose Input/Output|C4|
|GPIO0_110(1)|IO|General Purpose Input/Output|C2|
|GPIO0_111(1)|IO|General Purpose Input/Output|C1|
|GPIO0_112(1)|IO|General Purpose Input/Output|D4|
|GPIO0_113(1)|IO|General Purpose Input/Output|D3|
|GPIO0_114(1)|IO|General Purpose Input/Output|B2|
|GPIO0_115(1)|IO|General Purpose Input/Output|D2|
|GPIO0_116(1)|IO|General Purpose Input/Output|AB2|
|GPIO0_117(1)|IO|General Purpose Input/Output|AA2|
|GPIO0_118(1)|IO|General Purpose Input/Output|Y4|
|GPIO0_119(1)|IO|General Purpose Input/Output|AA1|
|GPIO0_120(1)|IO|General Purpose Input/Output|Y2|
|GPIO0_121(1)|IO|General Purpose Input/Output|Y3|
|GPIO0_122(1)|IO|General Purpose Input/Output|B6|
|GPIO0_123(1)|IO|General Purpose Input/Output|D6|
|GPIO0_124|IO|General Purpose Input/Output|C6|
|GPIO0_125|IO|General Purpose Input/Output|A5|
|GPIO0_20|IO|General Purpose Input/Output|H22|
|GPIO0_21|IO|General Purpose Input/Output|H23|
|GPIO0_22|IO|General Purpose Input/Output|J22|
|GPIO0_23|IO|General Purpose Input/Output|H19|
|GPIO0_24|IO|General Purpose Input/Output|H20|
|GPIO0_25|IO|General Purpose Input/Output|H21|
|GPIO0_26|IO|General Purpose Input/Output|H18|
|GPIO0_27|IO|General Purpose Input/Output|G23|
|GPIO0_28|IO|General Purpose Input/Output|G22|



Copyright © 2025 Texas Instruments Incorporated 

48 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

**Table 5-23. GPIO0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|GPIO0_29|IO|General Purpose Input/Output|F22|
|GPIO0_30|IO|General Purpose Input/Output|F23|
|GPIO0_31|IO|General Purpose Input/Output|L21|
|GPIO0_32|IO|General Purpose Input/Output|N19|
|GPIO0_33|IO|General Purpose Input/Output|N20|
|GPIO0_34|IO|General Purpose Input/Output|M19|
|GPIO0_35|IO|General Purpose Input/Output|P23|
|GPIO0_36|IO|General Purpose Input/Output|P22|
|GPIO0_37|IO|General Purpose Input/Output|N23|
|GPIO0_38|IO|General Purpose Input/Output|N22|
|GPIO0_39|IO|General Purpose Input/Output|N21|
|GPIO0_40|IO|General Purpose Input/Output|M21|
|GPIO0_41|IO|General Purpose Input/Output|L20|
|GPIO0_42|IO|General Purpose Input/Output|L19|
|GPIO0_43(1)|IO|General Purpose Input/Output|M23|
|GPIO0_44(1)|IO|General Purpose Input/Output|M22|
|GPIO0_45(1)|IO|General Purpose Input/Output|R22|
|GPIO0_46(1)|IO|General Purpose Input/Output|T23|
|GPIO0_47(1)|IO|General Purpose Input/Output|T22|
|GPIO0_48(1)|IO|General Purpose Input/Output|U22|
|GPIO0_49(1)|IO|General Purpose Input/Output|R23|
|GPIO0_50(1)|IO|General Purpose Input/Output|U23|
|GPIO0_51(1)|IO|General Purpose Input/Output|T20|
|GPIO0_52(1)|IO|General Purpose Input/Output|T21|
|GPIO0_53|IO|General Purpose Input/Output|AB11|
|GPIO0_54|IO|General Purpose Input/Output|W11|
|GPIO0_55|IO|General Purpose Input/Output|AC10|
|GPIO0_56|IO|General Purpose Input/Output|W13|
|GPIO0_57|IO|General Purpose Input/Output|Y11|
|GPIO0_58|IO|General Purpose Input/Output|AA11|
|GPIO0_59|IO|General Purpose Input/Output|Y6|
|GPIO0_60|IO|General Purpose Input/Output|Y7|
|GPIO0_61|IO|General Purpose Input/Output|Y8|
|GPIO0_62|IO|General Purpose Input/Output|AA6|
|GPIO0_63|IO|General Purpose Input/Output|AA8|
|GPIO0_64|IO|General Purpose Input/Output|W8|
|GPIO0_65|IO|General Purpose Input/Output|AC13|
|GPIO0_66|IO|General Purpose Input/Output|AC15|
|GPIO0_67|IO|General Purpose Input/Output|AB12|
|GPIO0_68|IO|General Purpose Input/Output|Y13|
|GPIO0_69|IO|General Purpose Input/Output|AC12|
|GPIO0_70|IO|General Purpose Input/Output|AB13|
|GPIO0_71|IO|General Purpose Input/Output|AA12|
|GPIO0_72|IO|General Purpose Input/Output|AA13|
|GPIO0_73|IO|General Purpose Input/Output|AC8|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 49 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**Table 5-23. GPIO0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|GPIO0_74|IO|General Purpose Input/Output|AC7|
|GPIO0_75|IO|General Purpose Input/Output|AB9|
|GPIO0_76|IO|General Purpose Input/Output|AC9|
|GPIO0_77|IO|General Purpose Input/Output|AB10|
|GPIO0_78|IO|General Purpose Input/Output|AB8|
|GPIO0_79|IO|General Purpose Input/Output|A8|
|GPIO0_80|IO|General Purpose Input/Output|B10|
|GPIO0_81|IO|General Purpose Input/Output|A9|
|GPIO0_82|IO|General Purpose Input/Output|B9|
|GPIO0_83|IO|General Purpose Input/Output|A11|
|GPIO0_84|IO|General Purpose Input/Output|B11|
|GPIO0_85|IO|General Purpose Input/Output|C11|
|GPIO0_86|IO|General Purpose Input/Output|A12|
|GPIO0_87|IO|General Purpose Input/Output|E11|
|GPIO0_88(1)|IO|General Purpose Input/Output|D11|
|GPIO0_89|IO|General Purpose Input/Output|E13|
|GPIO0_90|IO|General Purpose Input/Output|E12|
|GPIO0_91|IO|General Purpose Input/Output|B12|
|GPIO0_92|IO|General Purpose Input/Output|D13|
|GPIO0_93|IO|General Purpose Input/Output|C13|
|GPIO0_94|IO|General Purpose Input/Output|B14|
|GPIO0_95|IO|General Purpose Input/Output|B13|
|GPIO0_96|IO|General Purpose Input/Output|B16|
|GPIO0_97|IO|General Purpose Input/Output|B15|
|GPIO0_98|IO|General Purpose Input/Output|B7|
|GPIO0_99|IO|General Purpose Input/Output|A7|



(1) This GPIO input signal has a debounce function. For more information on I/O Debounce configuration, see the TRM _Device Configuration_ chapter. 

## **5.3.11.2 WKUP Domain** 

**Table 5-24. WKUP_GPIO0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|WKUP_GPIO0_0|IO|General Purpose Input/Output|Y22|
|WKUP_GPIO0_1|IO|General Purpose Input/Output|AA23|
|WKUP_GPIO0_2(1)|IO|General Purpose Input/Output|W23|
|WKUP_GPIO0_3(1)|IO|General Purpose Input/Output|W22|
|WKUP_GPIO0_4|IO|General Purpose Input/Output|AB22|
|WKUP_GPIO0_5|IO|General Purpose Input/Output|AA22|
|WKUP_GPIO0_6|IO|General Purpose Input/Output|Y23|



(1) This GPIO input signal has a debounce function. For more information on I/O Debounce configuration, see the TRM _Device Configuration_ chapter. 

Copyright © 2025 Texas Instruments Incorporated 

50 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**5.3.12 GPMC**_ 

## **5.3.12.1 MAIN Domain** 

## **Table 5-25. GPMC0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|GPMC0_ADVn_ALE|O|GPMC Address Valid (active low) or Address Latch<br>Enable|N19|
|GPMC0_CLK|O|GPMC clock|L21|
|GPMC0_DIR|O|GPMC Data Bus Signal Direction Control|M21|
|GPMC0_FCLK_MUX|O|GPMC functional clock output|L21|
|GPMC0_OEn_REn|O|GPMC Output Enable (active low) or Read Enable<br>(active low)|N20|
|GPMC0_WEn|O|GPMC Write Enable (active low)|M19|
|GPMC0_WPn|O|GPMC Flash Write Protect (active low)|N21|
|GPMC0_A0|O|GPMC Address 0 Output. Only used to effectively<br>address 8-bit data non-multiplexed memories|Y11|
|GPMC0_A1|O|GPMC address 1 Output in A/D non-multiplexed mode<br>and Address 17 in A/D multiplexed mode|AA12|
|GPMC0_A2|O|GPMC address 2 Output in A/D non-multiplexed mode<br>and Address 18 in A/D multiplexed mode|AA13|
|GPMC0_A3|O|GPMC address 3 Output in A/D non-multiplexed mode<br>and Address 19 in A/D multiplexed mode|AB10|
|GPMC0_A4|O|GPMC address 4 Output in A/D non-multiplexed mode<br>and Address 20 in A/D multiplexed mode|AB8|
|GPMC0_A5|O|GPMC address 5 Output in A/D non-multiplexed mode<br>and Address 21 in A/D multiplexed mode|AA8|
|GPMC0_A6|O|GPMC address 6 Output in A/D non-multiplexed mode<br>and Address 22 in A/D multiplexed mode|W8|
|GPMC0_AD0|IO|GPMC Data 0 Input/Output in A/D non-multiplexed<br>mode and additionally Address 1 Output in A/D<br>multiplexed mode|L22|
|GPMC0_AD1|IO|GPMC Data 1 Input/Output in A/D non-multiplexed<br>mode and additionally Address 2 Output in A/D<br>multiplexed mode|L23|
|GPMC0_AD2|IO|GPMC Data 2 Input/Output in A/D non-multiplexed<br>mode and additionally Address 3 Output in A/D<br>multiplexed mode|K22|
|GPMC0_AD3|IO|GPMC Data 3 Input/Output in A/D non-multiplexed<br>mode and additionally Address 3 Output in A/D<br>multiplexed mode|J23|
|GPMC0_AD4|IO|GPMC Data 4 Input/Output in A/D non-multiplexed<br>mode and additionally Address 3 Output in A/D<br>multiplexed mode|K23|
|GPMC0_AD5|IO|GPMC Data 5 Input/Output in A/D non-multiplexed<br>mode and additionally Address 3 Output in A/D<br>multiplexed mode|H22|
|GPMC0_AD6|IO|GPMC Data 6 Input/Output in A/D non-multiplexed<br>mode and additionally Address 3 Output in A/D<br>multiplexed mode|H23|
|GPMC0_AD7|IO|GPMC Data 7 Input/Output in A/D non-multiplexed<br>mode and additionally Address 3 Output in A/D<br>multiplexed mode|J22|
|GPMC0_AD8|IO|GPMC Data 8 Input/Output in A/D non-multiplexed<br>mode and additionally Address 3 Output in A/D<br>multiplexed mode|H19|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 51 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**Table 5-25. GPMC0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|GPMC0_AD9|IO|GPMC Data 9 Input/Output in A/D non-multiplexed<br>mode and additionally Address 3 Output in A/D<br>multiplexed mode|H20|
|GPMC0_AD10|IO|GPMC Data 10 Input/Output in A/D non-multiplexed<br>mode and additionally Address 11 Output in A/D<br>multiplexed mode|H21|
|GPMC0_AD11|IO|GPMC Data 11 Input/Output in A/D non-multiplexed<br>mode and additionally Address 12 Output in A/D<br>multiplexed mode|H18|
|GPMC0_AD12|IO|GPMC Data 12 Input/Output in A/D non-multiplexed<br>mode and additionally Address 13 Output in A/D<br>multiplexed mode|G23|
|GPMC0_AD13|IO|GPMC Data 13 Input/Output in A/D non-multiplexed<br>mode and additionally Address 14 Output in A/D<br>multiplexed mode|G22|
|GPMC0_AD14|IO|GPMC Data 14 Input/Output in A/D non-multiplexed<br>mode and additionally Address 15 Output in A/D<br>multiplexed mode|F22|
|GPMC0_AD15|IO|GPMC Data 15 Input/Output in A/D non-multiplexed<br>mode and additionally Address 16 Output in A/D<br>multiplexed mode|F23|
|GPMC0_BE0n_CLE|O|GPMC Lower-Byte Enable (active low) or Command<br>Latch Enable|P23|
|GPMC0_BE1n|O|GPMC Upper-Byte Enable (active low)|P22|
|GPMC0_CSn0|O|GPMC Chip Select 0 (active low)|L20|
|GPMC0_CSn1|O|GPMC Chip Select 1 (active low)|L19|
|GPMC0_CSn2|O|GPMC Chip Select 2 (active low)|M23|
|GPMC0_CSn3|O|GPMC Chip Select 3 (active low)|M22|
|GPMC0_WAIT0|I|GPMC External Indication of Wait|N23|
|GPMC0_WAIT1|I|GPMC External Indication of Wait|N22|



## _**5.3.13 I2C**_ 

## **5.3.13.1 MAIN Domain** 

**Table 5-26. I2C0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|I2C0_SCL|IOD|I2C Clock|B7|
|I2C0_SDA|IOD|I2C Data|A7|
|**Table 5-27. I2C1 Signal Descriptions**||||
|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|I2C1_SCL|IOD|I2C Clock|D7|
|I2C1_SDA|IOD|I2C Data|A6|



## **Table 5-28. I2C2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|I2C2_SCL|IOD|I2C Clock|AA8,B8,M23|
|I2C2_SDA|IOD|I2C Data|D8,M22,W8|



Copyright © 2025 Texas Instruments Incorporated 

52 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-29. I2C3 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|I2C3_SCL|IOD|I2C Clock|B14,B2,G20,<br>L22|
|I2C3_SDA|IOD|I2C Data|B13,D2,F20,<br>L23|



## **5.3.13.2 WKUP Domain** 

## **Table 5-30. WKUP_I2C0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|WKUP_I2C0_SCL|IOD|I2C Clock|AB22|
|WKUP_I2C0_SDA|IOD|I2C Data|AA22|



## _**5.3.14 MCAN**_ 

## **5.3.14.1 MAIN Domain** 

## **Table 5-31. MCAN0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MCAN0_RX|I|MCAN Receive Data|B15|
|MCAN0_TX|O|MCAN Transmit Data|B16|
|**Table 5-32. MCAN1 Signal Descriptions**||||
|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|MCAN1_RX|I|MCAN Receive Data|A8,AB2,B4,<br>N22|
|MCAN1_TX|O|MCAN Transmit Data|A3,AA2,B10,<br>M21|



## **Table 5-33. MCAN2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MCAN2_RX|I|MCAN Receive Data|B14,B6,C2|
|MCAN2_TX|O|MCAN Transmit Data|B13,C1,D6|



## _**5.3.15 MCASP**_ 

## **5.3.15.1 MAIN Domain** 

**Table 5-34. MCASP0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MCASP0_ACLKR|IO|MCASP Receive Bit Clock|A12|
|MCASP0_ACLKX|IO|MCASP Transmit Bit Clock|A11|
|MCASP0_AFSR|IO|MCASP Receive Frame Sync|C11|
|MCASP0_AFSX|IO|MCASP Transmit Frame Sync|B11|
|MCASP0_AXR0|IO|MCASP Serial Data (Input/Output)|B9|
|MCASP0_AXR1|IO|MCASP Serial Data (Input/Output)|A9|
|MCASP0_AXR2|IO|MCASP Serial Data (Input/Output)|B10|
|MCASP0_AXR3|IO|MCASP Serial Data (Input/Output)|A8|



## **Table 5-35. MCASP1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MCASP1_ACLKR|IO|MCASP Receive Bit Clock|C23,M22,R23|
|MCASP1_ACLKX|IO|MCASP Transmit Bit Clock|G20,P23,T20|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 53 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**Table 5-35. MCASP1 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MCASP1_AFSR|IO|MCASP Receive Frame Sync|D18,M23,U23|
|MCASP1_AFSX|IO|MCASP Transmit Frame Sync|F20,N23,T21|
|MCASP1_AXR0|IO|MCASP Serial Data (Input/Output)|F19,M19,U22|
|MCASP1_AXR1|IO|MCASP Serial Data (Input/Output)|F21,N20,T22|
|MCASP1_AXR2|IO|MCASP Serial Data (Input/Output)|D18,N19,T23|
|MCASP1_AXR3|IO|MCASP Serial Data (Input/Output)|C23,L21,R22|
|MCASP1_AXR4|IO|MCASP Serial Data (Input/Output)|H19,M23,U23|
|MCASP1_AXR5|IO|MCASP Serial Data (Input/Output)|J22,M22,R23|



**Table 5-36. MCASP2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MCASP2_ACLKR|IO|MCASP Receive Bit Clock|AB13,F23|
|MCASP2_ACLKX|IO|MCASP Transmit Bit Clock|AA13,B13,G22|
|MCASP2_AFSR|IO|MCASP Receive Frame Sync|AC9,F22|
|MCASP2_AFSX|IO|MCASP Transmit Frame Sync|AA12,B14,G23|
|MCASP2_AXR0|IO|MCASP Serial Data (Input/Output)|AB10,B16,H19|
|MCASP2_AXR1|IO|MCASP Serial Data (Input/Output)|AC7,B15,H20|
|MCASP2_AXR2|IO|MCASP Serial Data (Input/Output)|AB9,H21|
|MCASP2_AXR3|IO|MCASP Serial Data (Input/Output)|AC8,H18|
|MCASP2_AXR4|IO|MCASP Serial Data (Input/Output)|AB12,L22|
|MCASP2_AXR5|IO|MCASP Serial Data (Input/Output)|L23,Y13|
|MCASP2_AXR6|IO|MCASP Serial Data (Input/Output)|AC12,K22|
|MCASP2_AXR7|IO|MCASP Serial Data (Input/Output)|AC9,J23|
|MCASP2_AXR8|IO|MCASP Serial Data (Input/Output)|AB13,K23|
|MCASP2_AXR9|IO|MCASP Serial Data (Input/Output)|H22|
|MCASP2_AXR10|IO|MCASP Serial Data (Input/Output)|H23|
|MCASP2_AXR11|IO|MCASP Serial Data (Input/Output)|J22|
|MCASP2_AXR12|IO|MCASP Serial Data (Input/Output)|P22|
|MCASP2_AXR13|IO|MCASP Serial Data (Input/Output)|M21|
|MCASP2_AXR14|IO|MCASP Serial Data (Input/Output)|L20|
|MCASP2_AXR15|IO|MCASP Serial Data (Input/Output)|L19|



## _**5.3.16 MCSPI**_ 

## **5.3.16.1 MAIN Domain** 

**Table 5-37. MCSPI0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|SPI0_CLK|IO|SPI Clock|E13|
|SPI0_CS0|IO|SPI Chip Select 0|E11|
|SPI0_CS1|IO|SPI Chip Select 1|D11|
|SPI0_CS2|IO|SPI Chip Select 2|B14|
|SPI0_CS3|IO|SPI Chip Select 3|B13|
|SPI0_D0|IO|SPI Data 0|E12|
|SPI0_D1|IO|SPI Data 1|B12|



Copyright © 2025 Texas Instruments Incorporated 

54 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

## **Table 5-38. MCSPI1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|SPI1_CLK|IO|SPI Clock|C1,C4,F19,<br>H22|
|SPI1_CS0|IO|SPI Chip Select 0|B3,C2,F21,K23|
|SPI1_CS1|IO|SPI Chip Select 1|B2,D18,J22|
|SPI1_CS2|IO|SPI Chip Select 2|D2,H19|
|SPI1_CS3|IO|SPI Chip Select 3|D4,H23|
|SPI1_D0|IO|SPI Data 0|B4,G20,K22|
|SPI1_D1|IO|SPI Data 1|A3,F20,J23|



## **Table 5-39. MCSPI2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|SPI2_CLK|IO|SPI Clock|A12,A3,A6,B6,<br>D3,N21|
|SPI2_CS0|IO|SPI Chip Select 0|B7,C11,D4,<br>L19,Y3|
|SPI2_CS1|IO|SPI Chip Select 1|A11,AB2,B4,<br>D7,N22|
|SPI2_CS2|IO|SPI Chip Select 2|A7,A9,C2,Y2|
|SPI2_CS3|IO|SPI Chip Select 3|A3,AA2,B11,<br>D16|
|SPI2_D0|IO|SPI Data 0|A8,B3,D13,<br>M21,Y4|
|SPI2_D1|IO|SPI Data 1|AA1,B10,C13,<br>C4,L20|



**Table 5-40. MCSPI3 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|SPI3_CLK|IO|SPI Clock|AB10,P22,Y3|
|SPI3_CS0|IO|SPI Chip Select 0|AB8,P23,Y2|
|SPI3_CS1|IO|SPI Chip Select 1|D6,M19|
|SPI3_CS2|IO|SPI Chip Select 2|N20,Y4|
|SPI3_CS3|IO|SPI Chip Select 3|B6,N19|
|SPI3_D0|IO|SPI Data 0|AA12,AB2,N23|
|SPI3_D1|IO|SPI Data 1|AA13,AA2,N22|



## _**5.3.17 MDIO**_ 

## **5.3.17.1 MAIN Domain** 

**Table 5-41. MDIO0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MDIO0_MDC|O|MDIO Clock|AC15|
|MDIO0_MDIO|IO|MDIO Data|AC13|



## _**5.3.18 MMC**_ 

## **5.3.18.1 MAIN Domain** 

**Table 5-42. MMC0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MMC0_CLK|IO|MMC/SD/SDIO Clock|B2|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

55 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**Table 5-42. MMC0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MMC0_CMD|IO|MMC/SD/SDIO Command|D2|
|MMC0_DAT0|IO|MMC/SD/SDIO Data|D3|
|MMC0_DAT1|IO|MMC/SD/SDIO Data|D4|
|MMC0_DAT2|IO|MMC/SD/SDIO Data|C1|
|MMC0_DAT3|IO|MMC/SD/SDIO Data|C2|
|MMC0_DAT4|IO|MMC/SD/SDIO Data|C4|
|MMC0_DAT5|IO|MMC/SD/SDIO Data|B3|
|MMC0_DAT6|IO|MMC/SD/SDIO Data|A3|
|MMC0_DAT7|IO|MMC/SD/SDIO Data|B4|



**Table 5-43. MMC1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MMC1_CLK|IO|MMC/SD/SDIO Clock|Y2|
|MMC1_CMD|IO|MMC/SD/SDIO Command|Y3|
|MMC1_SDCD|I|SD Card Detect|B6|
|MMC1_SDWP|I|SD Write Protect|D6|
|MMC1_DAT0|IO|MMC/SD/SDIO Data|AA1|
|MMC1_DAT1|IO|MMC/SD/SDIO Data|Y4|
|MMC1_DAT2|IO|MMC/SD/SDIO Data|AA2|
|MMC1_DAT3|IO|MMC/SD/SDIO Data|AB2|



**Table 5-44. MMC2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|MMC2_CLK(1)|IO|MMC/SD/SDIO Clock|R23|
|MMC2_CMD|IO|MMC/SD/SDIO Command|U23|
|MMC2_SDCD(2)|I|SD Card Detect|B14,D7,T20|
|MMC2_SDWP(2)|I|SD Write Protect|A6,B13,T21|
|MMC2_DAT0|IO|MMC/SD/SDIO Data|U22|
|MMC2_DAT1|IO|MMC/SD/SDIO Data|T22|
|MMC2_DAT2|IO|MMC/SD/SDIO Data|T23|
|MMC2_DAT3|IO|MMC/SD/SDIO Data|R22|



(1) For MMC2 to work properly, the CTRLMMR_PADCONFIG66 register must be configured to set (1) the RXACTIVE bit and reset (0) the TX_DIS bit. 

(2) These MMCSD2 host controller input signals must be multiplexed to pins which are powered by the VDDSHV1 IO power rail rather than the VDDSHV4 IO power rail when the MMC2 port is connected to a UHS-I SD Card that requires the VDDSHV4 IO power rail to change its operating voltage from 3.3V to 1.8V when transitioning to one of the UHS-I data transfer modes. 

## _**5.3.19 OSPI**_ 

## **5.3.19.1 MAIN Domain** 

**Table 5-45. OSPI0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|OSPI0_CLK|O|OSPI Clock|D22|
|OSPI0_DQS|I|OSPI Data Strobe (DQS) or Loopback Clock Input|E22|
|OSPI0_ECC_FAIL|I|OSPI ECC Status|C23|
|OSPI0_LBCLKO|IO|OSPI Loopback Clock Output|E18|
|OSPI0_CSn0|O|OSPI Chip Select 0 (active low)|C20|



Copyright © 2025 Texas Instruments Incorporated 

56 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-45. OSPI0 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|OSPI0_CSn1|O|OSPI Chip Select 1 (active low)|D20|
|OSPI0_CSn2|O|OSPI Chip Select 2 (active low)|D18|
|OSPI0_CSn3|O|OSPI Chip Select 3 (active low)|C23|
|OSPI0_D0|IO|OSPI Data 0|C22|
|OSPI0_D1|IO|OSPI Data 1|D21|
|OSPI0_D2|IO|OSPI Data 2|E23|
|OSPI0_D3|IO|OSPI Data 3|D23|
|OSPI0_D4|IO|OSPI Data 4|F21|
|OSPI0_D5|IO|OSPI Data 5|F19|
|OSPI0_D6|IO|OSPI Data 6|G20|
|OSPI0_D7|IO|OSPI Data 7|F20|
|OSPI0_RESET_OUT0|O|OSPI Reset|C23|
|OSPI0_RESET_OUT1|O|OSPI Reset|D18|



## _**5.3.20 Power Supply**_ 

**Table 5-46. Power Supply Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|CAP_VDDSHV_MMC(1)|CAP|External capacitor connection for SDIO_LDO|T16|
|CAP_VDDS_GENERAL1(2)|CAP|External capacitor connection for GENERAL1 IO group|G11|
|CAP_VDDS_GPMC(2)|CAP|External capacitor connection for GPMC IO group|K16|
|CAP_VDDS_MMC0(2)|CAP|External capacitor connection for MMC0 IO group|J8|
|CAP_VDDS_MMC1(2)|CAP|External capacitor connection for MMC1 IO group|U9|
|CAP_VDDS_MMC2(2)|CAP|External capacitor connection for MMC2 IO group|M16|
|VDDA_1P8_DSI|PWR|DSITX0 1.8V analog supply|G14|
|VDDA_1P8_USB|PWR|USB0 and USB1 1.8V analog supply|T12|
|VDDA_3P3_SDIO|PWR|SDIO_LDO 3.3V analog supply|U16|
|VDDA_3P3_USB|PWR|USB0 and USB1 3.3V analog supply|U12|
|VDDA_ADC|PWR|ADC0 analog supply|N17|
|VDDA_CORE_DSI|PWR|DSITX0 core supply|G13|
|VDDA_CORE_DSI_CLK|PWR|DSITX0 clock core supply|H12|
|VDDA_CORE_USB|PWR|USB0 and USB1 core supply|U11|
|VDDA_DDR_PLL0|PWR|DDR deskew PLL supply|M10|
|VDDA_PLL0|PWR|WKUP_PLL0, MAIN_PLL0, and TEMP0 analog supply|L11|
|VDDA_PLL1|PWR|MAIN_PLL8 and MAIN_PLL17 analog supply|K12|
|VDDS0|PWR|Fixed-voltage supply for GENERAL0 IO group|T14|
|VDDS1|PWR|Fixed-voltage supply for GENERAL0_1 IO group|H16|
|VDDSHV0|PWR|Dual-voltage IO supply for GPMC IO group|J16,L17|
|VDDSHV1|PWR|Dual-voltage IO supply for General1 IO group|G10,H10|
|VDDSHV2|PWR|Dual-voltage IO supply for MMC0 IO group|H8|
|VDDSHV3|PWR|Dual-voltage IO supply MMC1 IO group|T10|
|VDDSHV4|PWR|Dual-voltage IO supply MMC2 IO group|M17|
|VDDS_DDR|PWR|DDR PHY IO supply|L8,M7,M8,N8,<br>P8|
|VDDS_OSC0|PWR|RCOSC, POR, and WKUP_OSC0 supply|R16|
|VDDS_RTC|PWR|Fixed-voltage supply for LFOSC0 and RTC IO group|T18|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 57 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-46. Power Supply Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|VDDS_WKUP|PWR|Fixed-voltage supply for WKUP IO group|P16|
|VDD_CORE|PWR|Core supply|J11,J13,J15,J9,<br>K10,K14,L15,<br>M14,N15,P10,<br>P12,P14,R11,<br>R9|
|VDD_RTC|PWR|RTC core supply|T17|
|VPP|PWR|eFuse ROM programming supply|N18|
|VSS|PWR|Ground|A1,A10,A13,<br>A16,A19,A2,<br>A22,A23,A4,<br>AA20,AA4,AB1,<br>AB21,AB23,<br>AB7,AC1,AC11,<br>AC14,AC19,<br>AC2,AC22,<br>AC23,B1,B17,<br>B20,B23,B5,<br>C12,C18,D1,<br>E10,E14,E15,<br>E2,E6,E8,E9,<br>F18,F5,F6,<br>G12,G15,G16,<br>G17,G7,G8,<br>G9,H1,H14,<br>H17,H7,K15,<br>K8,K9,L13,L16,<br>L18,L7,L9,M1,<br>M12,N11,N13,<br>N16,N7,N9,<br>P15,P9,R1,<br>R13,R15,R8,<br>T19,T2,T7,T8,<br>U10,U13,U14,<br>U15,U17,U20,<br>U7,U8,V18,<br>V19,V3,W10,<br>W12,W14,W15,<br>W16,W18,W9,<br>Y1,Y20,Y21|



- (1) This pin must always be connected via a 6.3V or greater, 3.3μF ±20% capacitor to VSS when the SDIO_LDO is being used to source VDDSHV3. The capacitor selected must provide a capacitance within the defined range after it has been derated for DC-bias, operating temperature, and aging effects. Otherwise, this pin may be connected directly to VSS when the VDDA_3P3_SDIO pin is also connected directly to VSS. 

- (2) This pin must always be connected via a 6.3V or greater, 0.8uF to 1.5μF capacitor to VSS if the respective VDDSHVx pin is ever operated at 3.3V. The capacitor selected must provide a capacitance within the defined range after it has been derated for DC-bias, operating temperature, and aging effects. There are three connection options if the respective VDDSHVx pin is only operated at 1.8V. The pin can be connected to the same decoupling capacitor that is required for 3.3V operation, it can be left unconnected, or it can be connected to the same 1.8V power source as the respective VDDSHVx pin. 

## _**5.3.21 Reserved**_ 

**Table 5-47. Reserved Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|RSVD0|N/A|Reserved, must be left unconnected|AB17|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

58 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## _**5.3.22 System and Miscellaneous**_ 

## **5.3.22.1 Boot Mode Configuration** 

## _**5.3.22.1.1 MAIN Domain**_ 

**Table 5-48. Sysboot Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|BOOTMODE00|I|Bootmode pin 0|L22|
|BOOTMODE01|I|Bootmode pin 1|L23|
|BOOTMODE02|I|Bootmode pin 2|K22|
|BOOTMODE03|I|Bootmode pin 3|J23|
|BOOTMODE04|I|Bootmode pin 4|K23|
|BOOTMODE05|I|Bootmode pin 5|H22|
|BOOTMODE06|I|Bootmode pin 6|H23|
|BOOTMODE07|I|Bootmode pin 7|J22|
|BOOTMODE08|I|Bootmode pin 8|H19|
|BOOTMODE09|I|Bootmode pin 9|H20|
|BOOTMODE10|I|Bootmode pin 10|H21|
|BOOTMODE11|I|Bootmode pin 11|H18|
|BOOTMODE12|I|Bootmode pin 12|G23|
|BOOTMODE13|I|Bootmode pin 13|G22|
|BOOTMODE14|I|Bootmode pin 14|F22|
|BOOTMODE15|I|Bootmode pin 15|F23|



## **5.3.22.2 Clock** 

## _**5.3.22.2.1 RTC Domain**_ 

## **Table 5-49. RTC Clock Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|LFOSC0_XI|I|Low frequency (32.768KHz) oscillator input|AC21|
|LFOSC0_XO|O|Low frequency (32.768KHz) oscillator output|AC20|



## _**5.3.22.2.2 WKUP Domain**_ 

**Table 5-50. WKUP Clock Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|WKUP_OSC0_XI|I|High frequency oscillator input|AC18|
|WKUP_OSC0_XO|O|High frequency oscillator output|AC17|



## **5.3.22.3 System** 

## _**5.3.22.3.1 MAIN Domain**_ 

**Table 5-51. System Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|AUDIO_EXT_REFCLK0|IO|External clock input to McASP or output from McASP|AB8,B14,B9|
|AUDIO_EXT_REFCLK1|IO|External clock input to McASP or output from McASP|B11,B13,D11,<br>N21|
|CLKOUT0|O|RMII Clock Output (50MHz). This pin is used for clock<br>source to the external RMII PHY and must also be<br>routed back to the respective RMII[x]_REF_CLK pin<br>for proper device operation.|AA11,D16|
|EXTINTn|I|External Interrupt|C8|
|EXT_REFCLK1|I|External clock input to Main Domain|D16|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 59 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**Table 5-51. System Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|OBSCLK0|O|Main Domain Observation clock output for test and<br>debug purposes only|H21|
|OBSCLK1|O|Main Domain Observation clock output for test and<br>debug purposes only|B7|
|RESETSTATz|O|Main Domain warm reset status output|C16|
|RESETz|I|Main Domain warm reset|E16|



## _**5.3.22.3.2 RTC Domain**_ 

## **Table 5-52. RTC System Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|EXT_WAKEUP0|I|External Wakeup Input|AB19|
|EXT_WAKEUP1|I|External Wakeup Input|AB20|
|PMIC_LPM_EN0|O|Dual-function PMIC control output, Low Power Mode<br>(active low) or PMIC Enable (active high)|AA18|
|RTC_PORz|I|RTC Power-on Reset|Y18|



## _**5.3.22.3.3 WKUP Domain**_ 

**Table 5-53. WKUP System Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|PORz|I|WKUP Domain cold reset|AB18|
|WKUP_CLKOUT0|O|WKUP Domain CLKOUT0 output|M22,Y23|
|WKUP_EXT_REFCLK0|I|External input to WKUP Domain|W22|
|WKUP_OBSCLK0|O|WKUP Domain Observation clock output for test and<br>debug purposes only|W23|
|WKUP_SYSCLKOUT0|O|WKUP Domain CLKOUT0 output|W23|



## _**5.3.23 TIMER**_ 

## **5.3.23.1 MAIN Domain** 

**Table 5-54. TIMER Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|TIMER_IO0|IO|Timer Inputs and Outputs (not tied to single timer<br>instance)|AB2,B2,C1,<br>D16,D7,Y2|
|TIMER_IO1|IO|Timer Inputs and Outputs (not tied to single timer<br>instance)|A6,A7,AA2,D2,<br>Y3|
|TIMER_IO2|IO|Timer Inputs and Outputs (not tied to single timer<br>instance)|B14,B16,B6,<br>H20,Y4|
|TIMER_IO3|IO|Timer Inputs and Outputs (not tied to single timer<br>instance)|AA1,B13,B15,<br>D6|



## **5.3.23.2 WKUP Domain** 

## **Table 5-55. WKUP_TIMER Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|WKUP_TIMER_IO0|IO|Timer Inputs and Outputs (not tied to single timer<br>instance)|W23|
|WKUP_TIMER_IO1|IO|Timer Inputs and Outputs (not tied to single timer<br>instance)|W22|



Copyright © 2025 Texas Instruments Incorporated 

60 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## _**5.3.24 UART**_ 

## **5.3.24.1 MAIN Domain** 

## **Table 5-56. UART0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|UART0_CTSn|I|UART Clear to Send (active low)|B14|
|UART0_RTSn|O|UART Request to Send (active low)|B13|
|UART0_RXD|I|UART Receive Data|D13|
|UART0_TXD|O|UART Transmit Data|C13|



## **Table 5-57. UART1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|UART1_CTSn|I|UART Clear to Send (active low)|A8|
|UART1_DCDn|I|UART Data Carrier Detect (active low)|B7|
|UART1_DSRn|I|UART Data Set Ready (active low)|A7|
|UART1_DTRn|O|UART Data Terminal Ready (active low)|B16|
|UART1_RIn|I|UART Ring Indicator|B15|
|UART1_RTSn|O|UART Request to Send (active low)|B10|
|UART1_RXD|I|UART Receive Data|C11,D7|
|UART1_TXD|O|UART Transmit Data|A12,A6|



**Table 5-58. UART2 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|UART2_CTSn|I|UART Clear to Send (active low)|AA1,C4,F22|
|UART2_RTSn|O|UART Request to Send (active low)|B3,F23,Y4|
|UART2_RXD|I|UART Receive Data|AB2,B14,B4,<br>H19|
|UART2_TXD|O|UART Transmit Data|A3,AA2,B13,<br>H20|



**Table 5-59. UART3 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|UART3_CTSn|I|UART Clear to Send (active low)|D3,D6|
|UART3_RTSn|O|UART Request to Send (active low)|B6,D4|
|UART3_RXD|I|UART Receive Data|C2,H21,Y2|
|UART3_TXD|O|UART Transmit Data|C1,H18,Y3|



## **Table 5-60. UART4 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|UART4_CTSn|I|UART Clear to Send (active low)|J22|
|UART4_RTSn|O|UART Request to Send (active low)|H23|
|UART4_RXD|I|UART Receive Data|G20,G23,M23,<br>T20|
|UART4_TXD|O|UART Transmit Data|F20,G22,M22,<br>T21|



## **Table 5-61. UART5 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|UART5_CTSn|I|UART Clear to Send (active low)|E22,H22|



Copyright © 2025 Texas Instruments Incorporated 

61 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-61. UART5 Signal Descriptions (continued)** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|UART5_RTSn|O|UART Request to Send (active low)|E18,K23|
|UART5_RXD|I|UART Receive Data|B16,D18,F22,<br>R22|
|UART5_TXD|O|UART Transmit Data|B15,C23,F23,<br>T23|



**Table 5-62. UART6 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|UART6_CTSn|I|UART Clear to Send (active low)|F20,J23|
|UART6_RTSn|O|UART Request to Send (active low)|G20,K22|
|UART6_RXD|I|UART Receive Data|A8,B6,F21,<br>L22,N22,R23|
|UART6_TXD|O|UART Transmit Data|B10,D6,F19,<br>L23,N21,U23|



## **5.3.24.2 WKUP Domain** 

**Table 5-63. WKUP_UART0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|WKUP_UART0_CTSn|I|UART Clear to Send (active low)|W23|
|WKUP_UART0_RTSn|O|UART Request to Send (active low)|W22|
|WKUP_UART0_RXD|I|UART Receive Data|Y22|
|WKUP_UART0_TXD|O|UART Transmit Data|AA23|



## _**5.3.25 USB**_ 

## **5.3.25.1 MAIN Domain** 

**Table 5-64. USB0 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|USB0_DM|IO|USB 2.0 Differential Data (negative)|AC4|
|USB0_DP|IO|USB 2.0 Differential Data (positive)|AB4|
|USB0_DRVVBUS|O|USB VBUS control output (active high)|C6|
|USB0_RCALIB(1)|IO|Pin to connect to calibration resistor|AB3|
|USB0_VBUS(2)|A|USB Level-shifted VBUS Input|AC3|



(1) An external 499Ω ±1% resistor must be connected between this pin and VSS and the maximum power dissipation for the resistor is 7.2mW. No external voltage should be applied to this pin. 

(2) An external resistor divider is required to limit the voltage applied to the device pin. For more information, see Section 8.2.3, _USB VBUS Design Guidelines_ . 

**Table 5-65. USB1 Signal Descriptions** 

|**SIGNAL NAME [**1**]**|**PIN TYPE [**2**]**|**DESCRIPTION [**3**]**|**ANB PIN [**4**]**|
|---|---|---|---|
|USB1_DM|IO|USB 2.0 Differential Data (negative)|AC5|
|USB1_DP|IO|USB 2.0 Differential Data (positive)|AB5|
|USB1_DRVVBUS|O|USB VBUS control output (active high)|A5|
|USB1_RCALIB(1)|IO|Pin to connect to calibration resistor|AC6|
|USB1_VBUS(2)|A|USB Level-shifted VBUS Input|AB6|



(1) An external 499Ω ±1% resistor must be connected between this pin and VSS and the maximum power dissipation for the resistor is 7.2mW. No external voltage should be applied to this pin. 

(2) An external resistor divider is required to limit the voltage applied to the device pin. For more information, see Section 8.2.3, _USB VBUS Design Guidelines_ . 

Copyright © 2025 Texas Instruments Incorporated 

62 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **5.4 Pin Connectivity Requirements** 

This section describes connectivity requirements for package balls that have specific connectivity requirements and unused package balls. 

## **Note** 

All power pins must be supplied with the voltages specified in Section 6.4, _Recommended Operating Conditions_ , unless otherwise specified. 

## **Note** 

For additional clarification, "leave unconnected" or "no connect" (NC) means **no** signal traces can be connected to these device ball numbers. 

**Table 5-66. Connectivity Requirements** 

|**ANB**<br>**BALL**<br>**NUMBER**|**BALL NAME**|**CONNECTION REQUIREMENTS**|
|---|---|---|
|AB16|TRSTn|This ball must be connected to VSS through an external pull resistor to ensure the<br>ball is held to a valid logic low level if a PCB signal trace is connected and not<br>actively driven by an attached device. The internal pull-down can be used to hold<br>a valid logic low level if no PCB signal trace is connected to the ball.|
|Y16<br>AA16<br>E16<br>AB14<br>AC16<br>Y17|EMU0<br>EMU1<br>RESETz<br>TCK<br>TDI<br>TMS|Each of these balls must be connected to the corresponding power supply(1)<br>through separate external pull resistors to ensure the inputs associated with these<br>balls are held to a valid logic high level if a PCB signal trace is connected and not<br>actively driven by an attached device. The internal pull-up can be used to hold a<br>valid logic high level if no PCB signal trace is connected to the ball.|
|AB19<br>AB20|EXT_WAKEUP0<br>EXT_WAKEUP1|Each of these balls must be connected to an always driven push-pull wake<br>up source, or connected to the corresponding power supply(1)or VSS through<br>external pull resistor(s) when not actively driven to ensure the inputs associated<br>with these balls are held in the appropriate valid high or low logic level based on<br>the polarity being used by the RTC wake up function.|
|L22<br>L23<br>K22<br>J23<br>K23<br>H22<br>H23<br>J22<br>H19<br>H20<br>H21<br>H18|GPMC0_AD0<br>GPMC0_AD1<br>GPMC0_AD2<br>GPMC0_AD3<br>GPMC0_AD4<br>GPMC0_AD5<br>GPMC0_AD6<br>GPMC0_AD7<br>GPMC0_AD8<br>GPMC0_AD9<br>GPMC0_AD10<br>GPMC0_AD11|When the full pin count boot mode option is selected by pulling GPMC0_AD15<br>and GPMC0_AD14 to VSS, each of these balls must be connected to the<br>corresponding power supply(1)or VSS through separate external pull resistors<br>to ensure the inputs associated with these balls are held to a valid logic high or<br>low level as appropriate to select the desired device boot mode.|
|G23<br>G22<br>F22<br>F23|GPMC0_AD12<br>GPMC0_AD13<br>GPMC0_AD14<br>GPMC0_AD15|Each of these balls must be connected to the corresponding power supply(1)or<br>VSS through separate external pull resistors to ensure the inputs associated with<br>these balls are held to a valid logic high or low level as appropriate to select the<br>desired device boot mode.|
|N17<br>V20<br>V22<br>V23<br>V21|VDDA_ADC<br>ADC0_AIN0<br>ADC0_AIN1<br>ADC0_AIN2<br>ADC0_AIN3|If the entire ADC0 is not used, each of these balls must be connected directly to<br>VSS.|
|V20<br>V22<br>V23<br>V21|ADC0_AIN0<br>ADC0_AIN1<br>ADC0_AIN2<br>ADC0_AIN3|Any unused ADC0_AIN[3:0] ball must be pulled to VSS through a resistor or<br>connected directly to VSS when VDDA_ADC is connected to a power source.|



Copyright © 2025 Texas Instruments Incorporated 

63 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-66. Connectivity Requirements (continued)** 

|**ANB**<br>**BALL**<br>**NUMBER**|**BALL NAME**|**CONNECTION REQUIREMENTS**|
|---|---|---|
|L8<br>M7<br>M8<br>N8<br>P8|VDDS_DDR<br>VDDS_DDR<br>VDDS_DDR<br>VDDS_DDR<br>VDDS_DDR|If DDRSS is not used, each of these balls must be connected directly to VSS.|
|M2<br>L1<br>M5<br>L2<br>L5<br>H6<br>L6<br>K2<br>J1<br>H5<br>R2<br>N6<br>T4<br>N1<br>T5<br>T6<br>W6<br>V6<br>N3<br>N2<br>N5<br>N4<br>M3<br>P1<br>P2<br>K1<br>L3<br>F2<br>W2<br>F4<br>F3<br>F1<br>E1<br>G4<br>H4<br>H2<br>H3<br>V4<br>T3<br>T1<br>U1<br>U4<br>V5<br>U2<br>W1<br>G1<br>G2<br>V1<br>V2<br>L4<br>J2|DDR0_ACT_n<br>DDR0_CAS_n<br>DDR0_RAS_n<br>DDR0_WE_n<br>DDR0_A0<br>DDR0_A1<br>DDR0_A2<br>DDR0_A3<br>DDR0_A4<br>DDR0_A5<br>DDR0_A6<br>DDR0_A7<br>DDR0_A8<br>DDR0_A9<br>DDR0_A10<br>DDR0_A11<br>DDR0_A12<br>DDR0_A13<br>DDR0_BA0<br>DDR0_BA1<br>DDR0_BG0<br>DDR0_BG1<br>DDR0_CAL0<br>DDR0_CK0<br>DDR0_CK0_n<br>DDR0_CKE0<br>DDR0_CS0_n<br>DDR0_DM0<br>DDR0_DM1<br>DDR0_DQ0<br>DDR0_DQ1<br>DDR0_DQ2<br>DDR0_DQ3<br>DDR0_DQ4<br>DDR0_DQ5<br>DDR0_DQ6<br>DDR0_DQ7<br>DDR0_DQ8<br>DDR0_DQ9<br>DDR0_DQ10<br>DDR0_DQ11<br>DDR0_DQ12<br>DDR0_DQ13<br>DDR0_DQ14<br>DDR0_DQ15<br>DDR0_DQS0<br>DDR0_DQS0_n<br>DDR0_DQS1<br>DDR0_DQS1_n<br>DDR0_ODT0<br>DDR0_RESET0_n|If DDRSS is not used, leave unconnected.<br>Note: The DDR0 pins in this list can only be left unconnected when VDDS_DDR<br>and VDDS_DDR_C are connected to VSS. The DDR0 pins must be connected as<br>defined in theDDR Board Design and Layout Guidelines, when VDDS_DDR and<br>VDDS_DDR_C are connected to a power source.|
|U16<br>T16|VDDA_3P3_SDIO<br>CAP_VDDSHV_MMC|If SDIO_LDO is not used to power VDDSHV3, each of these balls must be<br>connected directly to VSS.|
|U11<br>T12<br>U12|VDDA_CORE_USB<br>VDDA_1P8_USB<br>VDDA_3P3_USB|USB0 and USB1 share these power rails, so each of these balls must be<br>connected to valid power sources when either USB0 or USB1 is used.<br>If USB0 and USB1 are not used, each of these balls must be connected directly<br>to VSS.|



Copyright © 2025 Texas Instruments Incorporated 

64 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 5-66. Connectivity Requirements (continued)** 

|**ANB**<br>**BALL**<br>**NUMBER**|**BALL NAME**|**CONNECTION REQUIREMENTS**|
|---|---|---|
|AC4<br>AB4<br>AB3<br>AC3<br>AC5<br>AB5<br>AC6<br>AB6|USB0_DM<br>USB0_DP<br>USB0_RCALIB<br>USB0_VBUS<br>USB1_DM<br>USB1_DP<br>USB1_RCALIB<br>USB1_VBUS|If USB0 or USB1 is not used, leave the respective DM, DP, and VBUS balls<br>unconnected.<br>Note: The USB0_RCALIB and USB1_RCALIB pins can only be left<br>unconnected when VDDA_CORE_USB, VDDA_1P8_USB, and VDDA_3P3_USB<br>are connected to VSS. The USB0_RCALIB and USB1_RCALIB pins must<br>be connected to VSS through separate appropriate external resistors when<br>VDDA_CORE_USB, VDDA_1P8_USB, and VDDA_3P3_USB are connected to<br>power sources.|
|G13<br>H12<br>G14|VDDA_CORE_DSI<br>VDDA_CORE_DSI_CLK<br>VDDA_1P8_DSI|If DSITX0 is not used and the device boundary scan function is required, each of<br>these balls must be connected to valid power sources.<br>If DSITX0 is not used and the device boundary scan function is not required, each<br>of these balls can alternatively be connected directly to VSS.|
|A15<br>A14<br>B19<br>B18<br>D17|DSI0_TXCLKN<br>DSI0_TXCLKP<br>DSI0_TXN0<br>DSI0_TXP0<br>DSI0_TXRCALIB|If DSITX0 is not used, leave unconnected.|
|A18<br>A17|DSI0_TXN1<br>DSI0_TXP1|If DSITX0 is not used or only operated in 1-lane mode, leave unconnected.|
|A20<br>A21|DSI0_TXN2<br>DSI0_TXP2|If DSITX0 is not used or only operated in 1-lane or 2-lane mode, leave<br>unconnected.|
|B22<br>B21|DSI0_TXN3<br>DSI0_TXP3|If DSITX0 is not used or only operated in 1-lane, 2-lane, or 3-lane mode, leave<br>unconnected.|



(1) To determine which power supply is associated with any IO, see the POWER column of the _Pin Attributes_ table. 

## **Note** 

Internal pull resistors are weak and may not source enough current to maintain a valid logic level for some operating conditions. This can be the case when connected to components with leakage to the opposite logic level, or when external noise sources couple to signal traces attached to balls which are only pulled to a valid logic level by the internal resistor. Therefore, external pull resistors are recommended to hold a valid logic level on balls with external connections. 

Many of the device IOs are turned off by default and external pull resistors may be required to hold inputs of any attached device in a valid logic state until software initializes the respective IOs. The state of configurable device IOs are defined in the BALL STATE DURING RESET RX/TX/PULL and BALL STATE AFTER RESET RX/TX/PULL columns of the _Pin Attributes_ table. Any IO with its input buffer (RX) turned off is allowed to float without damaging the device. However, any IO with its input buffer (RX) turned on shall never be allowed to float to any potential between VILSS and VIHSS. The input buffer can enter a high-current state which could damage the IO cell if allowed to float between these levels. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 65 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6 Specifications** 

## **6.1 Absolute Maximum Ratings** 

over operating junction temperature range (unless otherwise noted)[(1)][(2)] 

|**PARAMETER**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|VDD_CORE|Core supply||-0.3<br>1.05|V|
|VDDA_CORE_DSI|DSITX0 core supply||-0.3<br>1.05|V|
|VDDA_CORE_DSI_CLK|DSITX0 clock core supply||-0.3<br>1.05|V|
|VDDA_CORE_USB|USB0 and USB1 core supply||-0.3<br>1.05|V|
|VDDA_DDR_PLL0|DDR Deskew PLL supply||-0.3<br>1.05|V|
|VDD_RTC|RTC core supply||-0.3<br>1.05|V|
|VDDS_DDR|DDR PHY IO supply||-0.3<br>1.57|V|
|VDDS_OSC0|RCOSC, POR, and WKUP_OSC0 supply||-0.3<br>1.98|V|
|VDDS_RTC|IO supply for LFOSC0 and RTC IO group||-0.3<br>1.98|V|
|VDDA_PLL0|WKUP_PLL0, MAIN_PLL0, and TEMP0 analog supply||-0.3<br>1.98|V|
|VDDA_PLL1|MAIN_PLL8 and MAIN_PLL17 analog supply||-0.3<br>1.98|V|
|VDDS_WKUP|IO supply for WKUP IO group||-0.3<br>1.98|V|
|VDDS0|IO supply for GENERAL0 IO group||-0.3<br>1.98|V|
|VDDS1|IO supply for GENERAL0_1 IO group||-0.3<br>1.98|V|
|VDDA_ADC|ADC analog supply||-0.3<br>1.98|V|
|VDDA_1P8_DSI|DSITX0 1.8V analog supply||-0.3<br>1.98|V|
|VDDA_1P8_USB|USB0 and USB1 1.8V analog supply||-0.3<br>1.98|V|
|VPP|eFuse ROM programming supply||-0.3<br>1.98|V|
|VDDSHV0|IO supply for GPMC IO group||-0.3<br>3.63|V|
|VDDSHV1|IO supply for General1 IO group||-0.3<br>3.63|V|
|VDDSHV2|IO supply for MMC0 IO group||-0.3<br>3.63|V|
|VDDSHV3|IO supply for MMC1 IO group||-0.3<br>3.63|V|
|VDDSHV4|IO supply for MMC2 IO group||-0.3<br>3.63|V|
|VDDA_3P3_SDIO|SDIO_LDO analog supply||-0.3<br>3.63|V|
|VDDA_3P3_USB|USB0 and USB1 3.3V analog supply||-0.3<br>3.63|V|
|Steady-state max voltage at all fail-safe IO pins||PORz|-0.3<br>3.63|V|
|||I2C2_SCL, I2C2_SDA, EXTINTn<br>When operating at 1.8V|-0.3<br>1.98(3)|V|
|||I2C2_SCL, I2C2_SDA, EXTINTn<br>When operating at 3.3V|-0.3<br>3.63(3)||
|Steady-state max voltage at all other IO pins(4)||USB0_VBUS, USB1_VBUS(5)|-0.3<br>3.6|V|
|||All other IO pins|-0.3<br>IO supply<br>voltage + 0.3|<br>V|
|Transient overshoot and undershoot at IO pin||20% of IO supply voltage for up to 20%<br>of the signal period (seeFigure 6-1,_IO_<br>_Transient Voltage Ranges_)|0.2 × VDD(6)|V|
|Latch-up performance(7)||I-Test|-100<br>100|mA|
|||Over-Voltage (OV) Test|1.5 x VDD(6)|V|
|TSTG|Storage temperature||-55<br>+150|°C|



(1) Operation outside the Absolute Maximum Ratings may cause permanent device damage. Absolute Maximum Ratings do not imply functional operation of the device at these or any other conditions beyond those listed under Recommended Operating Conditions. If used outside the Section 6.4, _Recommended Operating Conditions_ but within the Absolute Maximum Ratings, the device may not be fully functional, and this may affect device reliability, functionality, performance, and shorten the device lifetime. 

(2) All voltage values are with respect to VSS, unless otherwise noted. 

Copyright © 2025 Texas Instruments Incorporated 

66 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

- (3) The absolute maximum ratings for these fail-safe pins depends on their IO supply operating voltage. Therefore, this value is also defined by the maximum VIH value found in the _I2C Open-Drain, and Fail-Safe (I2C OD FS) Electrical Characteristics_ section, where the electrical characteristics table has separate parameter values for 1.8V mode and 3.3V mode. 

- (4) This parameter applies to all IO pins which are not fail-safe and the requirement applies to all values of IO supply voltage. For example, if the voltage applied to a specific IO supply is 0 volts the valid input voltage range for any IO powered by that supply will be –0.3 to +0.3 volts. Special attention should be applied anytime peripheral devices are not powered from the same power sources used to power the respective IO supply. It is important the attached peripheral never sources a voltage outside the valid input voltage range, including power supply ramp-up and ramp-down sequences. 

- (5) An external resistor divider is required to limit the voltage applied to this device pin. For more information, see Section 8.2.3, _USB Design Guidelines_ . 

- (6) VDD is the voltage on the corresponding power-supply pin(s) for the IO. 

- (7) For current pulse injection (I-Test): 

   - Pins stressed per JEDEC JESD78 (Class II) and passed with specified I/O pin injection current and clamp voltage of 1.5 times maximum recommended I/O voltage and negative 0.5 times maximum recommended I/O voltage. 

For over-voltage performance (Over-Voltage (OV) Test): 

- Supplies stressed per JEDEC JESD78 (Class II) and passed specified voltage injection. 

Fail-safe IO terminals are designed such they do not have dependencies on the respective IO power supply voltage. This allows external voltage sources to be connected to these IO terminals when the respective IO power supplies are turned off. The I2C2_SCL, I2C2_SDA, EXTINTn, and PORz are the only fail-safe IO terminals. All other IO terminals are not fail-safe and the voltage applied to them should be limited to the value defined by the Steady State Max. Voltage at all IO pins parameter in Section 6.1. 

**==> picture [250 x 102] intentionally omitted <==**

**----- Start of picture text -----**<br>
Overshoot = 20% of nominal<br>IO supply voltage<br>Tovershoot<br>Tperiod<br>Tundershoot<br>Undershoot = 20% of nominal<br>IO supply voltage<br>**----- End of picture text -----**<br>


- A. Tovershoot + Tundershoot < 20% of Tperiod 

**Figure 6-1. IO Transient Voltage Ranges** 

## **6.2 ESD Ratings** 

||||**VALUE**|**UNIT**|
|---|---|---|---|---|
|V(ESD)|Electrostatic discharge<br>(ESD)|Human-body model (HBM), per ANSI/ESDA/JEDEC JS-001(1)|±1000|V|
|||Charged-device model (CDM), per ANSI/ESDA/JEDEC JS-002(2)|±250||



- (1) JEDEC document JEP155 states that 500V HBM allows safe manufacturing with a standard ESD control process. 

- (2) JEDEC document JEP157 states that 250V CDM allows safe manufacturing with a standard ESD control process. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 67 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.3 Power-On Hours (POH)** 

|**POWER ON HOURS (POH)**(1) (2) (3)|**POWER ON HOURS (POH)**(1) (2) (3)|**POWER ON HOURS (POH)**(1) (2) (3)|
|---|---|---|
|**JUNCTION TEMPERATURE RANGE (TJ)**||**LIFETIME (POH)**|
|Extended Industrial|–40°C to 105°C|100000|
|125°C Industrial(4)|–40°C to 105°C|100000|
||–40°C to 125°C|20000(5)|



(1) This information is provided solely for your convenience and does not extend or modify the warranty provided under TI's standard terms and conditions for TI semiconductor products. 

(2) Unless specified in the table above, all voltage domains and operating conditions are supported in the device at the noted temperatures. 

(3) POH is a function of voltage, temperature and time. Usage at higher voltages and temperatures will result in a reduction in POH. 

(4) Either -40 to 105C or -40 to 125C profile should be chosen and applied through the lifetime of the application. Mixing of these profiles for the purposes of extending temperature and/or POH may result in increased reliability failure risk and is not recommended. 

(5) The -40 to 125C profile is defined as 20000 power on hours with a junction temperature as follows: 5%@-40°C, 65%@70°C, 20%@110°C, and 10%@125°C. 

Copyright © 2025 Texas Instruments Incorporated 

68 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.4 Recommended Operating Conditions** 

over operating junction temperature range (unless otherwise noted) 

|**SUPPLY NAME**|**DESCRIPTION**||**MIN**(1)<br>**NOM**<br>**MAX**(1)|**UNIT**|
|---|---|---|---|---|
|VDD_CORE(2)<br>VDDA_CORE_DSI(2)<br>VDDA_CORE_DSI_CLK(2)<br>VDDA_CORE_USB(2)<br>VDDA_DDR_PLL0(2)|Core supply<br>DSITX0 core supply<br>DSITX0 clock core supply<br>USB0 and USB1 core supply<br>DDR Deskew PLL supply||0.715<br>0.75<br>0.79|V|
|VDD_RTC|RTC core supply||0.715<br>0.75<br>0.79|V|
|VDDS_DDR|DDR PHY IO supply|1.1V operation|1.06<br>1.1<br>1.17|V|
|||1.2V operation|1.14<br>1.2<br>1.26|V|
|VDDS_OSC0|RCOSC, POR, and WKUP_OSC0 supply||1.71<br>1.8<br>1.89|V|
|VDDS_RTC|Fixed-voltage supply for LFOSC0 and RTC IO group||1.71<br>1.8<br>1.89|V|
|VDDA_PLL0|WKUP_PLL0, MAIN_PLL0, and TEMP0 analog supply||1.71<br>1.8<br>1.89|V|
|VDDA_PLL1|MAIN_PLL8 and MAIN_PLL17 analog supply||1.71<br>1.8<br>1.89|V|
|VDDS_WKUP|Fixed-voltage supply for WKUP IO group||1.71<br>1.8<br>1.89|V|
|VDDS0|Fixed-voltage supply for GENERAL0 IO group||1.71<br>1.8<br>1.89|V|
|VDDS1|Fixed-voltage supply for GENERAL0_1 IO group||1.71<br>1.8<br>1.89|V|
|VDDA_ADC|ADC analog supply||1.71<br>1.8<br>1.89|V|
|VDDA_1P8_DSI|DSITX0 1.8V analog supply||1.71<br>1.8<br>1.89|V|
|VDDA_1P8_USB|USB0 and USB1 1.8V analog supply||1.71<br>1.8<br>1.89|V|
|VPP|eFuse ROM programming supply||see(3)<br>see(3)<br>see(3)|V|
|VDDSHV0|Dual-voltage IO supply for GPMC IO group|1.8V operation|1.71<br>1.8<br>1.89|V|
|||3.3V operation|3.135<br>3.3<br>3.465|V|
|VDDSHV1|Dual-voltage IO supply for General1 IO<br>group|1.8V operation|1.71<br>1.8<br>1.89|V|
|||3.3V operation|3.135<br>3.3<br>3.465|V|
|VDDSHV2|Dual-voltage IO supply for MMC0 IO group|1.8V operation|1.71<br>1.8<br>1.89|V|
|||3.3V operation|3.135<br>3.3<br>3.465|V|
|VDDSHV3|Dual-voltage IO supply for MMC1 IO group|1.8V operation|1.71<br>1.8<br>1.89|V|
|||3.3V operation|3.135<br>3.3<br>3.465|V|
|VDDSHV4|Dual-voltage IO supply for MMC2 IO group|1.8V operation|1.71<br>1.8<br>1.89|V|
|||3.3V operation|3.135<br>3.3<br>3.465|V|
|VDDA_3P3_SDIO|SDIO_LDO analog supply||3.135<br>3.3<br>3.465|V|
|VDDA_3P3_USB|USB0 and USB1 3.3V analog supply||3.135<br>3.3<br>3.465|V|
|USB0_VBUS|USB0 Level-shifted VBUS Input||0<br>see(4)<br>3.465|V|
|USB1_VBUS|USB1 Level-shifted VBUS Input||0<br>see(4)<br>3.465|V|
|TJ|Operating junction temperature range|125°C Industrial|-40<br>125|°C|
|||Extended Industrial|-40<br>105|°C|



(1) The voltage at the device ball must never drop below the MIN voltage or rise above the MAX voltage for any amount of time during normal device operation. 

(2) VDD_CORE, VDDA_CORE_DSI, VDDA_CORE_DSI_CLK, VDDA_CORE_USB, and VDDA_DDR_PLL0 shall be sourced from the same power source. Care should be taken to ensure that voltage differential between VDD_CORE and VDDA_CORE_USB is within +/- 1%. 

(3) Refer to the Recommended Operating Conditions for OTP eFuse Programming table for VPP supply voltages based on eFuse usage. 

(4) An external resistor divider is required to limit the voltage applied to this device pin. For more information, see Section 8.2.3, _USB Design Guidelines_ . 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 69 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.5 Operating Performance Points** 

Table 6-1 defines the maximum operating frequency of the clocks for each device speed grade and Table 6-2 defines the only valid Operating Performance Points (OPPs) for the device subsystem and core clocks. 

**Table 6-1. Device Speed Grades** 

|**Speed**<br>**Gd**|**MAXIMUM OPERATING FREQUNCY (MHz)**|**MAXIMUM OPERATING FREQUNCY (MHz)**|**MAXIMUM OPERATING FREQUNCY (MHz)**|**MAXIMUM OPERATING FREQUNCY (MHz)**|**MAXIMUM**<br>**TRANSITION**<br>**RATE (MT/s)**(1)|**MAXIMUM**<br>**TRANSITION**<br>**RATE (MT/s)**(1)|
|---|---|---|---|---|---|---|
|**rae**|**A53SS**<br>**(Cortex-A53x)**|**MAIN_SYSCLK0**|**PER_SYSCLK0**|**WKUP_SYSCLK0**|**DDR4**|**LPDDR4**|
|E|833|500|400|400|1600|1600|
|O|1250|500|400|400|1600|1600|



(1) Maximum DDR Frequency will be limited based on the specific memory type (vendor) used in a system and by PCB implementation. Refer to DDR Board Design and Layout Guidelines for the proper PCB implementation to achieve maximum DDR frequency. 

**Table 6-2. Device Operating Performance Points** 

|**OPP**|**A53SS**(1)|**FIXED OPERATING FREQUENCY OPTIONS (MHz)**|**FIXED OPERATING FREQUENCY OPTIONS (MHz)**|**FIXED OPERATING FREQUENCY OPTIONS (MHz)**|**MT/s**(4)|**MT/s**(4)|
|---|---|---|---|---|---|---|
|||**MAIN_SYSCLK0**(2)|**PER_SYSCLK0**(3)|**WKUP_SYSCLK0**(2)|**DDR4**|**LPDDR4**|
|High|From<br>ARM0 PLL<br>Bypass<br>to<br>Speed<br>Grade<br>Maximum|500|400|400|Speed<br>Grade<br>Maximum|From<br>250 (DRAM DLL Off<br>Mode)(5)<br>to<br>Speed<br>Grade<br>Maximum|



(1) Initial operating frequency, set by software at boot. Supports Dynamic Frequency Scaling (DFS) after boot. 

(2) Initial operating frequency, set by software at boot. Run-time support for frequency change between initial operating frequency and PLL Bypass 

(3) Fixed operating frequency, set by software at boot. 

(4) Maximum DDR Frequency will be limited based on the specific memory type (vendor) used in a system and by PCB implementation. Refer to DDR Board Design and Layout Guidelines for the proper PCB implementation to achieve maximum DDR frequency. 

(5) The DDR PLL output, which sources DDR0_CK0 and DDR0_CK0_n, is typically defined in units of frequency. So the "DRAM DLL Off Mode" transaction rate is equal to 2x the DDR PLL output frequency when operating in bypass mode. 

## **6.6 Power Consumption Summary** 

For information on the device power consumption, see the AM62Lx Power Estimation Tool application note. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

70 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.7 Electrical Characteristics** 

## **Note** 

The interfaces or signals described in Section 6.7 correspond to the interfaces or signals available in multiplexing mode 0 (Primary Signal Function). 

All interfaces or signals multiplexed on the balls described in these tables have the same DC electrical characteristics, unless multiplexing involves a PHY and GPIO combination, in which case different DC electrical characteristics are specified for the different multiplexing modes (Functions). 

## _**6.7.1 I2C Open-Drain, and Fail-Safe (I2C OD FS) Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**1.8V MODE**|||||
|VIL|Input Low Voltage||0.3 × VDD (1)|V|
|VILSS|Input Low Voltage Steady State||0.3 × VDD (1)|V|
|VIH|Input High Voltage||0.7 × VDD (1)<br>1.98(2)|V|
|VIHSS|Input High Voltage Steady State||0.7 × VDD (1)|V|
|VHYS|Input Hysteresis Voltage||0.1 × VDD (1)|mV|
|IIN (3)|Input Leakage Current.|VI= 1.8V|10|µA|
|||VI= 0V|-10|µA|
|VOL|Output Low Voltage||0.2 × VDD (1)|V|
|IOL (4)|Low Level Output Current|VOL(MAX)|10|mA|
|SRI (6)|Input Slew Rate||18f(5)<br>or<br>1.8E+6|V/s|
|**3.3V MODE**<br>(7)|||||
|VIL|Input Low Voltage||0.3 × VDD (1)|V|
|VILSS|Input Low Voltage Steady State||0.25 × VDD (1)|V|
|VIH|Input High Voltage||0.7 × VDD (1)<br>3.63(2)|V|
|VIHSS|Input High Voltage Steady State||0.7 × VDD (1)|V|
|VHYS|Input Hysteresis Voltage||0.05 × VDD (1)|mV|
|IIN (3)|Input Leakage Current.|VI= 3.3V|10|µA|
|||VI= 0V|-10|µA|
|VOL|Output Low Voltage||0.4|V|
|IOL (4)|Low Level Output Current|VOL(MAX)|10|mA|
|SRI (6)|Input Slew Rate||33f(5)<br>or<br>3.3E+6<br>8E+7|V/s|



(1) VDD stands for corresponding power supply. For more information on the power supply name and the corresponding ball(s), see POWER column of the _Pin Attributes_ table. 

(2) This value also defines the Absolute Maximum Ratings value the IO. 

(3) This parameter defines leakage current when the terminal is operating as an input, undriven output, or both input and undriven output. 

(4) The IOL parameter defines the minimum Low Level Output Current for which the device is able to maintain the specified VOL value. The value defined by this parameter should be considered the maximum current available to a system implementation which needs to maintain the specified VOL value for attached components. 

(5) f = toggle frequency of the input signal in Hz. 

(6) This MIN parameter only applies to input signal functions which are not defined in their respective _Timing and Switching Characteristics_ sections. Select the MIN parameter which results in the largest value. 

(7) I2C Hs-mode is not supported when operating the IO in 3.3V mode. 

Copyright © 2025 Texas Instruments Incorporated 

71 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.7.2 Fail-Safe Reset (FS RESET) Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|VIL|Input Low Voltage||0.3 ×<br>VDDS_OSC0|<br>V|
|VILSS|Input Low Voltage Steady State||0.3 ×<br>VDDS_OSC0|<br>V|
|VIH|Input High Voltage||0.7 ×<br>VDDS_OSC0|V|
|VIHSS|Input High Voltage Steady State||0.7 ×<br>VDDS_OSC0|V|
|VHYS|Input Hysteresis Voltage||200|mV|
|IIN (1)|Input Leakage Current.|VI= 1.8V|10|µA|
|||VI= 0V|-10|µA|
|SRI (3)|Input Slew Rate||18f(2)<br>or<br>1.8E+6|V/s|



(1) This parameter defines leakage current when the terminal is operating as an input. 

(2) f = toggle frequency of the input signal in Hz. 

(3) This MIN parameter only applies to input signal functions which are not defined in their respective _Timing and Switching Characteristics_ sections. Select the MIN parameter which results in the largest value. 

## _**6.7.3 High-Frequency Oscillator (HFOSC) Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|VIL|Input Low Voltage||0.35 ×<br>VDDS_OSC0|<br>V|
|VIH|Input High Voltage||0.65 ×<br>VDDS_OSC0|V|
|VHYS|Input Hysteresis Voltage||49|mV|
|IIN (1)|Input Leakage Current.|VI= 1.8V|10|µA|
|||VI= 0V|-10|µA|



(1) This parameter defines leakage current when the terminal is operating as an input. 

## _**6.7.4 Low-Frequency Oscillator (LFXOSC) Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|VIL|Input Low Voltage||0.30 ×<br>VDDS_OSC0|<br>V|
|VIH|Input High Voltage||0.70 ×<br>VDDS_OSC0|V|
|VHYS|Input Hysteresis Voltage|Active Mode|85|mV|
|||Bypass Mode|324|mV|
|IIN (1)|Input Leakage Current.|VI= 1.8V|10|µA|
|||VI= 0V|-10|µA|



(1) This parameter defines leakage current when the terminal is operating as an input. 

Copyright © 2025 Texas Instruments Incorporated 

72 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## _**6.7.5 SDIO Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**1.8V MODE**|||||
|VIL|Input Low Voltage||0.58|V|
|VILSS|Input Low Voltage Steady State||0.58|V|
|VIH|Input High Voltage||1.27|V|
|VIHSS|Input High Voltage Steady State||1.7|V|
|VHYS|Input Hysteresis Voltage||150|mV|
|IIN (1)|Input Leakage Current.|VI= 1.8V|10|µA|
|||VI= 0V|-10|µA|
|RPU|Pull-up Resistor||40<br>50<br>60|kΩ|
|RPD|Pull-down Resistor||40<br>50<br>60|kΩ|
|VOL|Output Low Voltage||0.45|V|
|VOH|Output High Voltage||VDD(2)- 0.45|V|
|IOL (3)|Low Level Output Current|VOL(MAX)|4|mA|
|IOH (3)|High Level Output Current|VOH(MIN)|4|mA|
|SRI (5)|Input Slew Rate||18f(4)<br>or<br>1.8E+6|V/s|
|**3.3V MODE**|||||
|VIL|Input Low Voltage||0.25 × VDD(2)|V|
|VILSS|Input Low Voltage Steady State||0.15 × VDD(2)|V|
|VIH|Input High Voltage||0.625 ×<br>VDD(2)|V|
|VIHSS|Input High Voltage Steady State||0.625 ×<br>VDD(2)|V|
|VHYS|Input Hysteresis Voltage||150|mV|
|IIN (1)|Input Leakage Current.|VI= 3.3V|10|µA|
|||VI= 0V|-10|µA|
|RPU|Pull-up Resistor||40<br>50<br>60|kΩ|
|RPD|Pull-down Resistor||40<br>50<br>60|kΩ|
|VOL|Output Low Voltage||0.125 ×<br>VDD(2)|<br>V|
|VOH|Output High Voltage||0.75 × VDD(2)|V|
|IOL (3)|Low Level Output Current|VOL(MAX)|6|mA|
|IOH (3)|High Level Output Current|VOH(MIN)|10|mA|
|SRI (5)|Input Slew Rate||33f(4)<br>or<br>3.3E+6|V/s|



(1) This parameter defines leakage current when the terminal is operating as an input, undriven output, or both input and undriven output, without internal pulls enabled. 

(2) VDD stands for corresponding power supply. For more information on the power supply name and the corresponding ball(s), see POWER column of the _Pin Attributes_ table. 

(3) The IOL and IOH parameters define the minimum Low Level Output Current and High Level Output Current for which the device is able to maintain the specified VOL and VOH values. Values defined by these parameters should be considered the maximum current available to a system implementation which needs to maintain the specified VOL and VOH values for attached components. (4) f = toggle frequency of the input signal in Hz. 

(5) This MIN parameter only applies to input signal functions which are not defined in their respective _Timing and Switching Characteristics_ sections. Select the MIN parameter which results in the largest value. 

Copyright © 2025 Texas Instruments Incorporated 

73 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.7.6 LVCMOS Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**1.8V MODE**|||||
|VIL|Input Low Voltage||0.35 × VDD(1)|V|
|VILSS|Input Low Voltage Steady State||0.3 × VDD(1)|V|
|VIH|Input High Voltage||0.65 × VDD(1)|V|
|VIHSS|Input High Voltage Steady State||0.85 × VDD(1)|V|
|VHYS|Input Hysteresis Voltage||150|mV|
|IIN (2)|Input Leakage Current.|VI= 1.8V|10|µA|
|||VI= 0V|-10|µA|
|RPU|Pull-up Resistor||15<br>22<br>30|kΩ|
|RPD|Pull-down Resistor||15<br>22<br>30|kΩ|
|VOL|Output Low Voltage||0.45|V|
|VOH|Output High Voltage||VDD(1)- 0.45|V|
|IOL (3)|Low Level Output Current|VOL(MAX)|3|mA|
|IOH (3)|High Level Output Current|VOH(MIN)|3|mA|
|SRI (5)|Input Slew Rate||18f(4)<br>or<br>1.8E+6|V/s|
|**3.3V MODE**|||||
|VIL|Input Low Voltage||0.8|V|
|VILSS|Input Low Voltage Steady State||0.6|V|
|VIH|Input High Voltage||2.0|V|
|VIHSS|Input High Voltage Steady State||2.0|V|
|VHYS|Input Hysteresis Voltage||150|mV|
|IIN (2)|Input Leakage Current.|VI= 3.3V|10|µA|
|||VI= 0V|-10|µA|
|RPU|Pull-up Resistor||15<br>22<br>30|kΩ|
|RPD|Pull-down Resistor||15<br>22<br>30|kΩ|
|VOL|Output Low Voltage||0.4|V|
|VOH|Output High Voltage||2.4|V|
|IOL (3)|Low Level Output Current|VOL(MAX)|5|mA|
|IOH (3)|High Level Output Current|VOH(MIN)|9|mA|
|SRI (5)|Input Slew Rate||33f(4)<br>or<br>3.3E+6|V/s|



(1) VDD stands for corresponding power supply. For more information on the power supply name and the corresponding ball(s), see POWER column of the _Pin Attributes_ table. 

(2) This parameter defines leakage current when the terminal is operating as an input, undriven output, or both input and undriven output, without internal pulls enabled. 

(3) The IOL and IOH parameters define the minimum Low Level Output Current and High Level Output Current for which the device is able to maintain the specified VOL and VOH values. Values defined by these parameters should be considered the maximum current available to a system implementation which needs to maintain the specified VOL and VOH values for attached components. 

(4) f = toggle frequency of the input signal in Hz. 

(5) This MIN parameter only applies to input signal functions which are not defined in their respective _Timing and Switching Characteristics_ sections. Select the MIN parameter which results in the largest value. 

Copyright © 2025 Texas Instruments Incorporated 

74 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## _**6.7.7 1P8-LVCMOS Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|VIL|Input Low Voltage||0.35 × VDD(1)|V|
|VILSS|Input Low Voltage Steady State||0.35 × VDD(1)|V|
|VIH|Input High Voltage||0.65 × VDD(1)|V|
|VIHSS|Input High Voltage Steady State||0.65 × VDD(1)|V|
|VHYS|Input Hysteresis Voltage||150|mV|
|IIN (2)|Input Leakage Current.|VI= 1.8V<br>or<br>VI= 0.0V|±10|µA|
|RPU|Pull-up Resistor||10<br>20<br>30|kΩ|
|RPD|Pull-down Resistor||10<br>20<br>30|kΩ|
|VOL|Output Low Voltage||0.45|V|
|VOH|Output High Voltage||VDD(1)- 0.45|V|
|IOL (3)|Low Level Output Current|VOL(MAX)|8|mA|
|IOH (3)|High Level Output Current|VOH(MIN)|8|mA|
|SRI (5)|Input Slew Rate||9f(4)<br>or<br>1.08E+5|V/s|



(1) VDD stands for corresponding power supply. For more information on the power supply name and the corresponding ball(s), see POWER column of the _Pin Attributes_ table. 

(2) This parameter defines leakage current when the terminal is operating as an input, undriven output, or both input and undriven output, without internal pulls enabled. 

(3) The IOL and IOH parameters define the minimum Low Level Output Current and High Level Output Current for which the device is able to maintain the specified VOL and VOH values. Values defined by these parameters should be considered the maximum current available to a system implementation which needs to maintain the specified VOL and VOH values for attached components. 

(4) f = toggle frequency of the input signal in Hz. 

(5) This MIN parameter only applies to input signal functions which are not defined in their respective _Timing and Switching Characteristics_ sections. Select the MIN parameter which results in the largest value. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 75 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.7.8 RTC-LVCMOS Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|VIL|Input Low Voltage||0.35 × VDD(1)|V|
|VILSS|Input Low Voltage Steady State||0.35 × VDD(1)|V|
|VIH|Input High Voltage||0.65 × VDD(1)|V|
|VIHSS|Input High Voltage Steady State||0.65 × VDD(1)|V|
|VHYS|Input Hysteresis Voltage||200|mV|
|IIN (2)|Input Leakage Current.|VI= 1.8V<br>or<br>VI= 0.0V|±50|nA|
|RPU (3)|Pull-up Resistor||21.5<br>30.0|kΩ|
|VOL|Output Low Voltage||0.45|V|
|VOH|Output High Voltage||VDD(1)- 0.45|V|
|IOL (4)|Low Level Output Current|VOL(MAX)|2|mA|
|IOH (4)|High Level Output Current|VOH(MIN)|2|mA|
|SRI (5)|Input Slew Rate||1.8E6|V/s|



(1) VDD stands for corresponding power supply. For more information on the power supply name and the corresponding ball(s), see POWER column of the _Pin Attributes_ table. 

(2) This parameter only applies to the EXT_WAKEUP0, EXT_WAKEUP1, and RTC_PORz pins. 

(3) This parameter only applies to the PMIC_LPM_EN0 pin. 

(4) The IOL and IOH parameters define the minimum Low Level Output Current and High Level Output Current for which the device is able to maintain the specified VOL and VOH values. Values defined by these parameters should be considered the maximum current available to a system implementation which needs to maintain the specified VOL and VOH values for attached components. 

(5) This MIN parameter only applies to input signal functions which are not defined in their respective _Timing and Switching Characteristics_ sections. 

Copyright © 2025 Texas Instruments Incorporated 

76 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## _**6.7.9 ADC Electrical Characteristics**_ 

over recommended operating conditions (unless otherwise noted) 

|**PARAMETER**|**PARAMETER**|**TEST CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|Resolution|Actual Number of Bits||12|Bits|
|ENOB|Effective Number of Bits||≅10|Bits|
|VADC0_VREFP (1)|Positive Reference Voltage||VDDA_ADC0<br>(2)|V|
|VADC0_VREFN (1)|Negative Reference Voltage||VSS|V|
|VADC_AIN[3:0]|Analog Input Voltage,<br>ADC_AIN[3:0], Full-scale||VSS<br>VDDA_ADC0<br>(2)|V|
|DNL|Differential Non-Linearity||> -1<br>+4|LSB|
|INL|Integral Non-Linearity||-4<br>+4|LSB|
|LSBGAIN-ERROR|Gain Error||±10|LSB|
|LSBOFFSET-ERROR|Offset Error||±5|LSB|
|SINAD|Signal-to-Noise and Distortion<br>Ratio|Input Signal:<br>200kHz sine wave at<br>–0.5dB Full Scale|60|dB|
|ZADC_AIN[0:7]|Analog Input Impedance,<br>ADC0_AIN[7:0]||(3)|Ω|
|IIN|Input Leakage||±10|μA|
|CSMPL|Sampling Capacitance||5.5|pF|
|**Sampling Dynamics**|||||
|FSMPL_CLK|ADC0 SMPL_CLK Frequency||30|MHz|
|tC|Conversion Time||13|ADC0<br>SMPL_CLK<br>Cycles|
|tACQ|Acquisition Time||2<br>257|ADC0<br>SMPL_CLK<br>Cycles|
|TR|Sampling Rate|ADC0 SMPL_CLK<br>= 30MHz|2|MSPS|



(1) ADC0_REFP and ADC0_REFN are directly connected to VDDA_ADC0 and VSS inside the SoC. References to ADC0_REFP and ADC0_REFN in this table must be considered as VDDA_ADC0 or VSS. 

(2) Valid voltage range for VDDA_ADC0 is defined in Section 6.4 

(3) The ADC0_AIN pins are connected to an internal sampling capacitor for a user configurable acquisition time and acquisition frequency. The input impedance of the ADC0_AIN pins is a function of the sampling capacitance along with user configurable acquisition time and acquisition frequency. The designer must understand the time required for the source impedance of each ADC0_AIN pin to charge the internal sampling capacitor. The acquisition time must be set long enough for the internal sampling capacitor to settle to greater than 14bits of accuracy. 

## _**6.7.10 DSI (D-PHY) Electrical Characteristics**_ 

## **Note** 

DSITX0 is compliant with MIPI DPHY v1.2 dated August 1, 2014, including ECNs and Errata as applicable. 

## _**6.7.11 USB2PHY Electrical Characteristics**_ 

## **Note** 

The USB0 and USB1 interfaces are compliant with Universal Serial Bus Revision 2.0 Specification dated April 27, 2000 including ECNs and Errata as applicable. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

77 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.7.12 DDR Electrical Characteristics**_ 

## **Note** 

The DDR interface is compatible with DDR4 devices that are **JESD79-4B standard-compliant** , and LPDDR4 devices that are **JESD209-4B standard-compliant** 

Copyright © 2025 Texas Instruments Incorporated 

78 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.8 VPP Specifications for One-Time Programmable (OTP) eFuses** 

This section specifies the operating conditions required for programming the OTP eFuses . 

## _**6.8.1 Recommended Operating Conditions for OTP eFuse Programming**_ 

over operating junction temperature range (unless otherwise noted) 

|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**NOM**<br>**MAX**|**UNIT**|
|---|---|---|---|
|VDD_CORE|Supply voltage range for the core domain during OTP<br>operation|SeeSection 6.4|V|
|VPP|Supply voltage range for the eFuse ROM domain during<br>normal operation without hardware support to program<br>eFuse ROM|NC(1)|V|
||Supply voltage range for the eFuse ROM domain during<br>normal operation with hardware support to program eFuse<br>ROM|0|V|
||Supply voltage range for the eFuse ROM domain during<br>OTP programming(2)|1.71<br>1.8<br>1.89|V|
|I(VPP)|VPP current|400|mA|
|SR(VPP)|VPP Power-up Slew Rate|6E + 4|V/s|
|Tj|Operating junction temperature range while programming<br>eFuse ROM.|0<br>25<br>85|°C|



(1) NC indicates No Connect. 

(2) Supply voltage range includes DC errors and peak-to-peak noise. 

## _**6.8.2 Hardware Requirements**_ 

The following hardware requirements must be met when programming keys in the OTP eFuses: 

- The VPP power supply must be disabled when not programming OTP registers. 

- The VPP power supply must be ramped up after the proper device power-up sequence (for more details, see Section 6.11.2.2, _Power Supply Sequencing_ ). 

## _**6.8.3 Programming Sequence**_ 

Programming sequence for OTP eFuses: 

- Power on the board per the power-up sequencing. No voltage should be applied on the VPP terminal during power up and normal operation. 

- Load the OTP write software required to program the eFuse (contact your local TI representative for the OTP software package). 

- Apply the voltage on the VPP terminal according to the specification in Section 6.8.1. 

- Run the software that programs the OTP registers. 

- After validating the content of the OTP registers, remove the voltage from the VPP terminal. 

## _**6.8.4 Impact to Your Hardware Warranty**_ 

You accept that eFusing the TI Devices with security keys permanently alters them. You acknowledge that the eFuse can fail, for example, due to incorrect or aborted program sequence or if you omit a sequence step. Further the TI device may fail to secure boot if the error code correction check fails for the Production Keys or if the image is not signed and optionally encrypted with the current active Production Keys. These types of situations will render the TI device inoperable and TI will be unable to confirm whether the TI devices conformed to their specifications prior to the attempted eFuse. Consequently, TI will have no liability ( _warranty or otherwise_ ) for any TI devices that have been incorrectly eFused by customers. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 79 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.9 Thermal Resistance Characteristics** 

This section provides the thermal resistance characteristics used on this device. 

For reliability and operability concerns, the maximum junction temperature of the device has to be at or below the TJ value identified in Section 6.4, _Recommended Operating Conditions_ . 

## **Note** 

The thermal parameters are generated following JEDEC standard JESD51x and are not intended for design parameters. If you need a more accurate thermal representation, download the processor thermal model and import your PCB design into a thermal simulation environment. For details on thermal implementation guidelines, see the Thermal Solution Guidance section. 

## _**6.9.1 Thermal Resistance Characteristics for ANB Package**_ 

It is recommended to perform thermal simulations at the system level with the worst case device power consumption. 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**ANB**<br>**PACKAGE**<br>**°C/W**(1) (2)|**AIR**<br>**FLOW**<br>**(m/s)**(3)|
|---|---|---|---|---|
|T1|RΘJC|Junction-to-case|5.2|N/A|
|T2|RΘJB|Junction-to-board|9.4|N/A|
|T3|RΘJA|Junction-to-free air|22.2|0|
|T4||Junction-to-moving air|17.4|1|
|T5|||16.3|2|
|T6|||15.6|3|
|T7|ΨJT|Junction-to-package top|0.09|0|
|T8|||0.18|1|
|T9|||0.24|2|
|T10|||0.28|3|
|T11|ΨJB|Junction-to-board|9.3|0|
|T12|||8.8|1|
|T13|||8.6|2|
|T14|||8.5|3|



(1) °C/W = degrees Celsius per watt. 

- (2) These values are based on a JEDEC defined 2S2P system (with the exception of the Theta JC [RΘJC] value, which is based on a JEDEC defined 1S0P system) and will change based on environment as well as application. For more information, see these EIA/JEDEC standards: 

   - JESD51-2, _Integrated Circuits Thermal Test Method Environment Conditions - Natural Convection (Still Air)_ 

   - JESD51-3, _Low Effective Thermal Conductivity Test Board for Leaded Surface Mount Packages_ 

   - JESD51-6, _Integrated Circuit Thermal Test Method Environmental Conditions - Forced Convection (Moving Air)_ 

   - JESD51-7, _High Effective Thermal Conductivity Test Board for Leaded Surface Mount Packages_ 

   - JESD51-9, _Test Boards for Area Array Surface Mount Packages_ 

- (3) m/s = meters per second. 

Copyright © 2025 Texas Instruments Incorporated 

80 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.10 Temperature Sensor Characteristics** 

This section summarizes the Voltage and Temperature Module (VTM) on die temperature sensor characteristics. 

For operation and reliability concerns, the maximum junction temperature of the device must be equal to or less than the TJ value identified in _Recommended Operating Conditions_ . 

**Table 6-3. VTM Die Temperature Sensor Characteristics** 

|**PARAMETER**|**PARAMETER**|**TEST**<br>**CONDITIONS**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|Tacc|VTM temperature sensor accuracy|–40℃to 125℃|–5<br>5|℃|



Copyright © 2025 Texas Instruments Incorporated 

81 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11 Timing and Switching Characteristics** 

## **Note** 

The Timing Requirements and Switching Characteristics values may change following the silicon characterization result. 

## **Note** 

The default SLEWRATE settings in each pad configuration register must be used to ensure timings, unless specific instructions are given otherwise. 

## _**6.11.1 Timing Parameters and Information**_ 

The timing parameter symbols used in Section 6.11, _Timing and Switching Characteristics_ are created in accordance with JEDEC Standard 100. To shorten the symbols, some pin names and other related terminologies have been abbreviated in Table 6-4: 

**Table 6-4. Timing Parameters Subscripts** 

|**SYMBOL**|**PARAMETER**|
|---|---|
|c|Cycle time (period)|
|d|Delay time|
|dis|Disable time|
|en|Enable time|
|h|Hold time|
|su|Setup time|
|START|Start bit|
|t|Transition time|
|v|Valid time|
|w|Pulse duration (width)|
|X|Unknown, changing, or don't care level|
|F|Fall time|
|H|High|
|L|Low|
|R|Rise time|
|V|Valid|
|IV|Invalid|
|AE|Active Edge|
|FE|First Edge|
|LE|Last Edge|
|Z|High impedance|



Copyright © 2025 Texas Instruments Incorporated 

82 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.2 Power Supply Requirements**_ 

This section describes the power supply requirements to ensure proper device operation. 

## **Note** 

All power balls must be supplied with the voltages specified in the _Recommended Operating Conditions_ section, unless otherwise specified in _Signal Descriptions_ and _Pin Connectivity Requirements_ . 

## **6.11.2.1 Power Supply Slew Rate Requirement** 

To maintain the safe operating range of the internal ESD protection devices, TI recommends limiting the maximum slew rate of supplies to be less than 18mV/µs. For instance, as shown in Figure 6-2, TI recommends having the supply ramp slew for a 1.8V supply of more than 100µs. 

Figure 6-2 describes the Power Supply Slew Rate Requirement in the device. 

**==> picture [58 x 45] intentionally omitted <==**

**==> picture [251 x 167] intentionally omitted <==**

**----- Start of picture text -----**<br>
Supply value<br>t<br>slew rate < 18 mV/μs<br>slew > (supply value) / (18 mV/μs)<br>or<br>supply value × 55.6 μs/V<br>SPRT740_ELCH_06<br>**----- End of picture text -----**<br>


**Figure 6-2. Power Supply Slew and Slew Rate** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 83 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.2.2 Power Supply Sequencing** 

This section describes power sequence requirements using power sequence diagrams and associated notes. Each power sequence diagram demonstrates the sequential order expected for each device power rail. This is done by assigning each device power rail to one or more waveform. A dual-voltage power rail may be associated with more than one waveform and the associated note will describe which waveform is applicable. Each waveform defines a transition region for the associated power rails and shows its sequential relationship to the transition regions of other power rails. The notes associated with the power sequence diagram provides further detail of these requirements. See the _Power-up Sequence_ section for details on power-up requirements, and the _Power-down Sequence_ section for details on power-down requirements. 

Two types of power supply transition regions are used to simplify the power supply sequencing diagrams. The legends shown in Figure 6-3 and Figure 6-4 along with their descriptions are provided to clarify what each transition regions represents. 

Figure 6-3 defines a transition region with multiple power rails which may be sourced from multiple power supplies or a single power supply. Transitions shown within the transition region represent a use case where multiple power supplies are used to source power rails associated with this waveform, and these power supplies are allowed to ramp at different times within the region since they do not have any specific sequence requirement relative to each other. 

## **Figure 6-3. Multiple Power Supply Transition Legend** 

Figure 6-4 defines a transition region with one or more power rails which must be sourced from a single common power supply. No transitions are shown within the region to represent a single ramp within the transition region. 

## **Figure 6-4. Single Common Power Supply Transition Legend** 

Copyright © 2025 Texas Instruments Incorporated 

84 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.2.2.1 No Low-Power Mode Sequencing**_ 

Table 6-5, Figure 6-5, and Figure 6-6 define device power sequencing requirements when there is no plans to use RTC Only low-power mode or RTC + IO + DDR low-power mode. 

**Table 6-5. No Low-Power Mode Sequencing – Supply / Signal Assignments** 

See: Figure 6-5 and Figure 6-6 

|**WAVEFORM**|**SUPPLY / SIGNAL NAME**|
|---|---|
|A|System power|
|B|VDDSHV0(1), VDDSHV1(1), VDDA_3P3_USB|
|C|VDDSHV0(2), VDDSHV1,(2), VDDS_OSC0, VDDS_RTC, VDDA_PLL0, VDDA_PLL1, VDDS_WKUP, VDDS0, VDDS1,<br>VDDA_ADC, VDDA_1P8_DSI, VDDA_1P8_USB|
|D|VDDA_3P3_SDIO(3) (4), VDDSHV2(3), VDDSHV3(3), VDDSHV4(3)|
|E|VDDS_DDR(5)|
|F|VDD_CORE, VDDA_CORE_DSI(6), VDDA_CORE_DSI_CLK(6), VDDA_CORE_USB(6), VDDA_DDR_PLL0(6),<br>VDD_RTC|
|G|WKUP_OSC0_XI, WKUP_OSC0_XO|
|H|PORz|



(1) VDDSHV0 and VDDSHV1 are dual voltage IO supplies which can be operated at 1.8V or 3.3V depending on the application requirements. When any of the VDDSHVx [x=0-1] IO supplies are operating at 3.3V, they shall be ramped up with other 3.3V supplies during the 3.3V ramp period defined by this waveform. 

(2) VDDSHV0 and VDDSHV1 are dual voltage IO supplies which can be operated at 1.8V or 3.3V depending on the application requirements. When any of the VDDSHVx [x=0-1] IO supplies are operating at 1.8V, they shall be ramped up with other 1.8V supplies during the 1.8V ramp period defined by this waveform. 

(3) VDDA_3P3_SDIO was designed to support power-up or power-down without any dependency on other power rails. VDDSHV2, VDDSHV3, and VDDSHV4 were designed to support power-up, power-down, or dynamic voltage change without any dependency on other power rails. This capability is required to support UHS-I SD Cards. 

(4) VDDA_3P3_SDIO is the 3.3V power rail for the internal SDIO_LDO. This power rail must be sourced from the same 3.3V power supply that provides power to a UHS-I SD Card connected to MMC1, which allows the MMC1 IOs and the SD Card IOs to power-up and power-down at the same time when the SD Card power supply is powered off to reset the SD Card. For this use case the SDIO_LDO output (CAP_VDDSHV_MMC) is used to power the VDDSHV3 IO power rail, which will ramp-up and ramp-down along with the VDDA_3P3_SDIO power rail. 

(5) VDDS_DDR does not have any specific power sequence requirement, but the JEDEC standard for DDR devices requires the potential applied to its VDD1 power rail to always be greater than the potential applied to its VDD2 power rail during the power-up and power-down sequences. 

(6) VDDA_CORE_DSI, VDDA_CORE_DSI_CLK, VDDA_CORE_USB, VDDA_DDR_PLL0, and VDD_RTC shall be sourced from the same power source as VDD_CORE. Care should be taken to ensure that voltage differential between VDD_CORE and VDDA_CORE_USB is within +/- 1%. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 85 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [302 x 298] intentionally omitted <==**

**----- Start of picture text -----**<br>
A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>**----- End of picture text -----**<br>


**Figure 6-5. No Low-Power Mode Power-Up Sequencing** 

Copyright © 2025 Texas Instruments Incorporated 

86 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [228 x 299] intentionally omitted <==**

**----- Start of picture text -----**<br>
A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>**----- End of picture text -----**<br>


**Figure 6-6. No Low-Power Mode Power-Down Sequencing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 87 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.2.2.2 RTC Only Low-Power Mode Sequencing**_ 

Table 6-6, Figure 6-7, Figure 6-8, and Figure 6-9 define device power requirements when using RTC Only low-power mode. 

**Table 6-6. RTC Only Low-Power Mode Sequencing – Supply / Signal Assignments** 

See: Figure 6-7, Figure 6-8, and Figure 6-9 

|**WAVEFORM**|**SUPPLY / SIGNAL NAME**|
|---|---|
|A|System power|
|B|VDDS_RTC(1)|
|C|VDD_RTC(2)|
|D|PMIC_LPM_EN0(3)|
|E|RTC_PORz(4)|
|F|VDDSHV0(5), VDDSHV1(5), VDDA_3P3_USB|
|G|VDDSHV0(6), VDDSHV1(6), VDDS_OSC0, VDDA_PLL0, VDDA_PLL1, VDDS_WKUP, VDDS0, VDDS1, VDDA_ADC,<br>VDDA_1P8_DSI, VDDA_1P8_USB|
|H|VDDA_3P3_SDIO(7) (8), VDDSHV2(7), VDDSHV3(7), VDDSHV4(7)|
|I|VDDS_DDR(9)|
|J|VDD_CORE(10), VDDA_CORE_DSI(11), VDDA_CORE_DSI_CLK(11), VDDA_CORE_USB(11), VDDA_DDR_PLL0(11)|
|K|WKUP_OSC0_XI, WKUP_OSC0_XO|
|L|PORz|



(1) VDDS_RTC must be sourced from an always on power source when using RTC Only low-power mode. 

(2) VDD_RTC must be sourced from an always on power source when using RTC Only low-power mode. 

(3) PMIC_LPM_EN0 is pulled high with a weak internal pull-up while RTC_PORz is asserted. The weak internal pull-up is turned off and PMIC_LPM_EN0 is driven high on the rising edge of RTC_PORz. The RTC module can be configured to drive PMIC_LPM_EN0 low to enter RTC Only low-power mode and drive PMIC_LPM_EN0 high to exit RTC Only low-power mode, such that PMIC_LPM_EN0 can be used to cycle power on/off to all non-RTC power rails. 

(4) RTC_PORz can be released once the VDDS_RTC and VDD_RTC power rails are valid. 

(5) VDDSHV0 and VDDSHV1 are dual voltage IO supplies which can be operated at 1.8V or 3.3V depending on the application requirements. When any of the VDDSHVx [x=0-1] IO supplies are operating at 3.3V, they shall be ramped down with other 3.3V supplies during the 3.3V ramp period defined by this waveform. 

(6) VDDSHV0 and VDDSHV1 are dual voltage IO supplies which can be operated at 1.8V or 3.3V depending on the application requirements. When any of the VDDSHVx [x=0-1] IO supplies are operating at 1.8V, they shall be ramped down with other 1.8V supplies during the 1.8V ramp period defined by this waveform. 

(7) VDDA_3P3_SDIO was designed to support power-up or power-down without any dependency on other power rails. VDDSHV2, VDDSHV3, and VDDSHV4 were designed to support power-up, power-down, or dynamic voltage change without any dependency on other power rails. This capability is required to support UHS-I SD Cards. 

(8) VDDA_3P3_SDIO is the 3.3V power rail for the internal SDIO_LDO. This power rail must be sourced from the same 3.3V power supply that provides power to a UHS-I SD Card connected to MMC1, which allows the MMC1 IOs and the SD Card IOs to power-up and power-down at the same time when the SD Card power supply is powered off to reset the SD Card. For this use case the SDIO_LDO output (CAP_VDDSHV_MMC) is used to power the VDDSHV3 IO power rail, which will ramp-up and ramp-down along with the VDDA_3P3_SDIO power rail. 

(9) VDDS_DDR does not have any specific power sequence requirement, but the JEDEC standard for DDR devices requires the potential applied to its VDD1 power rail to always be greater than the potential applied to its VDD2 power rail during the power-up and power-down sequences. 

(10) The potential applied to VDD_CORE must never be greater than the potential applied to VDD_RTC + 0.18V during power-up or power-down. This requires VDD_RTC to ramp up before and ramp down after VDD_CORE. 

(11) VDDA_CORE_DSI, VDDA_CORE_DSI_CLK, VDDA_CORE_USB, and VDDA_DDR_PLL0 shall be sourced from the same power source as VDD_CORE. Care should be taken to ensure that voltage differential between VDD_CORE and VDDA_CORE_USB is within +/- 1%. 

Copyright © 2025 Texas Instruments Incorporated 

88 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [318 x 460] intentionally omitted <==**

**----- Start of picture text -----**<br>
A<br>B<br>C<br>D<br>Weak Pullup Driven High<br>E<br>F<br>G<br>H<br>I<br>J<br>K<br>L<br>**----- End of picture text -----**<br>


**Figure 6-7. RTC Only Low-Power Mode Power-Up Sequencing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 89 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**==> picture [372 x 467] intentionally omitted <==**

**----- Start of picture text -----**<br>
A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>I<br>J<br>K<br>L<br>**----- End of picture text -----**<br>


**Figure 6-8. RTC Only Low-Power Mode Enter/Exit Sequencing** 

Copyright © 2025 Texas Instruments Incorporated 

90 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [228 x 464] intentionally omitted <==**

**----- Start of picture text -----**<br>
A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>I<br>J<br>K<br>L<br>**----- End of picture text -----**<br>


**Figure 6-9. RTC Only Low-Power Mode Power-Down Sequencing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 91 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.2.2.3 RTC + IO + DDR Low-Power Mode Sequencing**_ 

Table 6-7, Figure 6-10, Figure 6-11, and Figure 6-12 define device power requirements when using RTC + IO + DDR low-power mode. 

**Table 6-7. RTC + IO + DDR Low-Power Mode Sequencing – Supply / Signal Assignments** 

See: Figure 6-10, Figure 6-11, and Figure 6-12 

|**WAVEFORM**|**SUPPLY / SIGNAL NAME**|
|---|---|
|A|System power|
|B|VDDSHV0(1), VDDSHV1(1), VDDA_3P3_USB|
|C|VDDSHV0(2), VDDSHV1(2), VDDS_OSC0(3), VDDA_PLL0(3), VDDA_PLL1(3), VDDS_WKUP, VDDS0, VDDS1,<br>VDDA_ADC(3), VDDA_1P8_DSI(3), VDDA_1P8_USB(3), VDDS_RTC(4)|
|D|VDDA_3P3_SDIO(5) (6), VDDSHV2(5), VDDSHV3(5), VDDSHV4(5)|
|E|VDD_RTC(7)|
|F|RTC_PORz(8)|
|G|VDDS_DDR(9)|
|H|VDD_CORE(10), VDDA_CORE_DSI(11), VDDA_CORE_DSI_CLK(11), VDDA_CORE_USB(11), VDDA_DDR_PLL0(11)|
|I|WKUP_OSC0_XI, WKUP_OSC0_XO|
|J|PORz|
|K|PMIC_LPM_EN0(12)|



(1) VDDSHV0 and VDDSHV1 are dual voltage IO supplies which can be operated at 1.8V or 3.3V depending on the application requirements. When any of the VDDSHVx [x=0-1] IO supplies are operating at 3.3V, they shall be ramped down with other 3.3V supplies during the 3.3V ramp period defined by this waveform. (2) VDDSHV0 and VDDSHV1 are dual voltage IO supplies which can be operated at 1.8V or 3.3V depending on the application requirements. When any of the VDDSHVx [x=0-1] IO supplies are operating at 1.8V, they shall be ramped down with other 1.8V supplies during the 1.8V ramp period defined by this waveform. (3) VDDS_OSC0, VDDA_PLL0, VDDA_PLL1, VDDA_ADC, VDDA_1P8_DSI, and VDDA_1P8_USB may be powered off when entering RTC + IO + DDR low-power mode to conserve power. (4) VDDS_RTC must be sourced from an always on power source when using RTC + IO + DDR low-power mode. (5) VDDA_3P3_SDIO was designed to support power-up or power-down without any dependency on other power rails. VDDSHV2, VDDSHV3, and VDDSHV4 were designed to support power-up, power-down, or dynamic voltage change without any dependency on other power rails. This capability is required to support UHS-I SD Cards. 

(6) VDDA_3P3_SDIO is the 3.3V power rail for the internal SDIO_LDO. This power rail must be sourced from the same 3.3V power supply that provides power to a UHS-I SD Card connected to MMC1, which allows the MMC1 IOs and the SD Card IOs to power-up and power-down at the same time when the SD Card power supply is powered off to reset the SD Card. For this use case the SDIO_LDO output (CAP_VDDSHV_MMC) is used to power the VDDSHV3 IO power rail, which will ramp-up and ramp-down along with the VDDA_3P3_SDIO power rail. 

(7) VDD_RTC must be sourced from an always on power source when using RTC + IO + DDR low-power mode. 

(8) RTC_PORz can be released once the VDDS_RTC and VDD_RTC power rails are valid. 

(9) VDDS_DDR does not have any specific power sequence requirement, but the JEDEC standard for DDR devices requires the potential applied to its VDD1 power rail to always be greater than the potential applied to its VDD2 power rail during the power-up and power-down sequences. (10) The potential applied to VDD_CORE must never be greater than the potential applied to VDD_RTC + 0.18V during power-up or power-down. This requires VDD_RTC to ramp up before and ramp down after VDD_CORE. (11) VDDA_CORE_DSI, VDDA_CORE_DSI_CLK, VDDA_CORE_USB, and VDDA_DDR_PLL0 shall be sourced from the same power source as VDD_CORE. Care should be taken to ensure that voltage differential between VDD_CORE and VDDA_CORE_USB is within +/- 1%. 

- (12) PMIC_LPM_EN0 is pulled high with a weak internal pull-up while RTC_PORz is asserted. The weak internal pull-up is turned off and PMIC_LPM_EN0 is driven high on the rising edge of RTC_PORz. The RTC module can be configured to drive PMIC_LPM_EN0 low to enter RTC + IO + DDR low-power mode and drive PMIC_LPM_EN0 high to exit RTC + IO + DDR low-power mode, such that PMIC_LPM_EN0 can be used to cycle power on/off to VDD_CORE and all 1.8V analog power rails. 

Copyright © 2025 Texas Instruments Incorporated 

92 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [318 x 438] intentionally omitted <==**

**----- Start of picture text -----**<br>
A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>I<br>J<br>K<br>Weak Pullup Driven High<br>**----- End of picture text -----**<br>


**Figure 6-10. RTC + IO + DDR Low-Power Mode Power-Up Sequencing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 93 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**==> picture [330 x 427] intentionally omitted <==**

**----- Start of picture text -----**<br>
A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>I<br>J<br>K<br>**----- End of picture text -----**<br>


**Figure 6-11. RTC + IO + DDR Low-Power Mode Enter/Exit Sequencing** 

Copyright © 2025 Texas Instruments Incorporated 

94 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**==> picture [228 x 421] intentionally omitted <==**

**----- Start of picture text -----**<br>
A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>I<br>J<br>K<br>**----- End of picture text -----**<br>


**Figure 6-12. RTC + IO + DDR Low-Power Mode Power-Down Sequencing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 95 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.3 System Timing**_ 

For more details about features and additional description information on the subsystem multiplexing signals, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

## **6.11.3.1 Reset Timing** 

Tables and figures provided in this section define timing conditions, timing requirements, and switching characteristics for reset related signals. 

**Table 6-8. Reset Timing Conditions** 

|**PARAMETER**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**INPUT CONDITIONS**|||||
|SRI|Input slew rate<br>V<br>V|DD(1)= 1.8V<br>0.|0018|V/ns|
|||DD(1)= 3.3V<br>0.|0033|V/ns|
|**OUTPUT CONDITIONS**|||||
|CL|Output load capacitance||30|pF|



(1) VDD stands for corresponding power supply. For more information on the power supply name and the corresponding ball(s), see POWER column of the _Pin Attributes_ table. 

**Table 6-9. PORz Timing Requirements** 

## see Figure 6-13 

|**NO.**||**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|RST1|th(SUPPLIES_VALID - PORz)|Hold time, PORz active (low) at Power-up after<br>supplies valid (using external crystal circuit)<br>95|00000|ns|
|RST2||Hold time, PORz active (low) at Power-up after<br>supplies valid and external clock stable (using<br>external LVCMOS clock source)|1200|ns|
|RST3|tw(PORzL)|Pulse Width, PORz low after Power-up (without<br>removal of Power or system reference clock<br>WKUP_OSC0_XI/XO)|1200|ns|



**==> picture [394 x 177] intentionally omitted <==**

**----- Start of picture text -----**<br>
RST1<br>RST2 RST3<br>PORz<br>ALL SUPPLIES<br>VALID<br>WKUP_OSC0_XI,<br>WKUP_OSC0_XO<br>**----- End of picture text -----**<br>


**Figure 6-13. PORz Timing Requirements** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

96 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-10. RESETSTATz Switching Characteristics** 

## see Figure 6-14 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|RST6|td(PORzL-RESETSTATzL)|Delay time, PORz active (low) to RESETSTATz<br>active (low)|0|ns|
|RST7|td(PORzH-RESETSTATzH)|Delay time, PORz inactive (high) to RESETSTATz<br>inactive (high)|9195*S(1)|ns|
|RST9|tw(RESETSTATzL)|Pulse Width, RESETSTATz low ( SW_WARMRST)|4040*S(1)|ns|



(1) S = WKUP_OSC0_XI/XO clock period in ns. 

**==> picture [279 x 106] intentionally omitted <==**

**----- Start of picture text -----**<br>
PORz<br>RST6<br>RST7<br>RST9<br>RESETSTATz<br>**----- End of picture text -----**<br>


**Figure 6-14. RESETSTATz Switching Characteristics** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 97 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 6-11. RESETz Timing Requirements** 

|**Table 6-11. RESETz Timing Requirements**|**Table 6-11. RESETz Timing Requirements**|**Table 6-11. RESETz Timing Requirements**|**Table 6-11. RESETz Timing Requirements**|**Table 6-11. RESETz Timing Requirements**|
|---|---|---|---|---|
|seeFigure 6-15|||||
|**NO.**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
|RST10|tw(RESETzL) (1)|Pulse Width, RESETz active (low)|1200|ns|



(1) This timing parameter is valid only after all supplies are valid and PORz has been asserted for the specified time. 

## **Table 6-12. RESETSTATz Switching Characteristics** 

## see Figure 6-15 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|RST13|td(RESETzL-RESETSTATzL)|Delay time, RESETz active (low) to RESETSTATz<br>active (low)|960|ns|
|RST14|td(RESETzH-RESETSTATzH)|Delay time, RESETz inactive (high) to<br>RESETSTATz inactive (high)|4040*S(1)|ns|



(1) S = WKUP_OSC0_XI/XO clock period in ns. 

**==> picture [51 x 81] intentionally omitted <==**

**----- Start of picture text -----**<br>
RESETz<br>RESETSTATz<br>**----- End of picture text -----**<br>


**==> picture [153 x 129] intentionally omitted <==**

**----- Start of picture text -----**<br>
RST10<br>RST13<br>RST14<br>**----- End of picture text -----**<br>


## **Figure 6-15. RESETz and RESETSTATz Timing Requirements and Switching Characteristics** 

## **Table 6-13. EMUx Timing Requirements** 

|**Table 6-13. EMUx Timing Requirements**|**Table 6-13. EMUx Timing Requirements**|**Table 6-13. EMUx Timing Requirements**|**Table 6-13. EMUx Timing Requirements**|**Table 6-13. EMUx Timing Requirements**|
|---|---|---|---|---|
|seeFigure 6-16|||||
|**NO.**|**PARAMETER**||**MIN**<br>**MAX **|**UNIT**|
|RST18|tsu(EMUx-PORz)|Setup time, EMU[1:0] before PORz inactive (high)|3*S(1)|ns|
|RST19|th(PORz - EMUx)|Hold time, EMU[1:0] after PORz inactive (high)|10|ns|



(1) S = WKUP_OSC0_XI/XO clock period in ns. 

**==> picture [220 x 118] intentionally omitted <==**

**----- Start of picture text -----**<br>
RST18<br>PORz<br>RST19<br>EMU[1:0]<br>**----- End of picture text -----**<br>


**Figure 6-16. EMUx Timing Requirements** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

98 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-14. BOOTMODE Timing Requirements** 

see Figure 6-17 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|RST23|tsu(BOOTMODE-PORz)|Setup time, BOOTMODE[15:00] valid before PORz<br>high|3*S(1)|ns|
|RST24|th(PORz - BOOTMODE)|Hold time, BOOTMODE[15:00] valid after PORz<br>high|0|ns|



(1) S = WKUP_OSC0_XI/XO clock period in ns. 

**==> picture [312 x 95] intentionally omitted <==**

**----- Start of picture text -----**<br>
PORz<br>RST23 RST24<br>BOOTMODE[15:00]<br>**----- End of picture text -----**<br>


**Figure 6-17. BOOTMODE Timing Requirements** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 99 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.3.2 Clock Timing** 

Tables and figures provided in this section define timing conditions, timing requirements, and switching characteristics for clock signals. 

## **Table 6-15. Clock Timing Conditions** 

|**PARAMETER**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**INPUT CONDITIONS**|||||
|SRI|Input slew rate||0.5|V/ns|
|**OUTPUT CONDITIONS**|||||
|CL|Output load capacitance|5ns ≤ tc< 8ns|5|pF|
|||8ns ≤ tc< 20ns|10|pF|
|||20ns ≤ tc|30|pF|



**Table 6-16. Clock Timing Requirements** 

## see Figure 6-18 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|CLK1|tc(EXT_REFCLK1)|Cycle time minimum, EXT_REFCLK1|10|ns|
|CLK2|tw(EXT_REFCLK1H)|Pulse Duration, EXT_REFCLK1 high|E*0.45(1)<br>E*0.55(1)|ns|
|CLK3|tw(EXT_REFCLK1L)|Pulse Duration, EXT_REFCLK1 low|E*0.45(1)<br>E*0.55(1)|ns|
|CLK1|tc(WKUP_EXT_REFCLK0)|Cycle time minimum, WKUP_EXT_REFCLK0|10|ns|
|CLK2|tw(WKUP_EXT_REFCLK0H)|Pulse Duration, WKUP_EXT_REFCLK0 high|F*0.45(2)<br>F*0.55(2)|ns|
|CLK3|tw(WKUP_EXT_REFCLK0L)|Pulse Duration, WKUP_EXT_REFCLK0 low|F*0.45(2)<br>F*0.55(2)|ns|
|CLK1|tc(AUDIO_EXT_REFCLK0)|Cycle time minimum, AUDIO_EXT_REFCLK0|20|ns|
|CLK2|tw(AUDIO_EXT_REFCLK0H)|Pulse Duration, AUDIO_EXT_REFCLK0 high|G*0.45(3)<br>G*0.55(3)|ns|
|CLK3|tw(AUDIO_EXT_REFCLK0L)|Pulse Duration, AUDIO_EXT_REFCLK0 low|G*0.45(3)<br>G*0.55(3)|ns|
|CLK1|tc(AUDIO_EXT_REFCLK1)|Cycle time minimum, AUDIO_EXT_REFCLK1|20|ns|
|CLK2|tw(AUDIO_EXT_REFCLK1H)|Pulse Duration, AUDIO_EXT_REFCLK1 high|H*0.45(4)<br>H*0.55(4)|ns|
|CLK3|tw(AUDIO_EXT_REFCLK1L)|Pulse Duration, AUDIO_EXT_REFCLK1 low|H*0.45(4)<br>H*0.55(4)|ns|



(1) E = EXT_REFCLK1 cycle time in ns. 

(2) F = WKUP_EXT_REFCLK0 cycle time in ns. 

(3) G = AUDIO_EXT_REFCLK0 cycle time in ns. 

(4) H = AUDIO_EXT_REFCLK1 cycle time in ns. 

**==> picture [276 x 49] intentionally omitted <==**

**----- Start of picture text -----**<br>
CLK1<br>CLK2 CLK3<br>Input Clock<br>**----- End of picture text -----**<br>


**Figure 6-18. Clock Timing Requirements** 

Copyright © 2025 Texas Instruments Incorporated 

100 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **Table 6-17. Clock Switching Characteristics** 

see Figure 6-19 

|**NO.**||**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|CLK4|tc(OBSCLK0)|Cycle time minimum, OBSCLK0|5|ns|
|CLK5|tw(OBSCLK0H)|Pulse Duration, OBSCLK0 high|B*0.45(1)<br>B*0.55(1)|ns|
|CLK6|tw(OBSCLK0L)|Pulse Duration, OBSCLK0 low|B*0.45(1)<br>B*0.55(1)|ns|
|CLK4|tc(OBSCLK1)|Cycle time minimum, OBSCLK1|5|ns|
|CLK5|tw(OBSCLK1H)|Pulse Duration, OBSCLK1 high|F*0.45(2)<br>F*0.55(2)|ns|
|CLK6|tw(OBSCLK1L)|Pulse Duration, OBSCLK1 low|F*0.45(2)<br>F*0.55(2)|ns|
|CLK4|tc(CLKOUT0)|Cycle time minimum, CLKOUT0|20|ns|
|CLK5|tw(CLKOUT0H)|Pulse Duration, CLKOUT0 high|C*0.4(3)<br>C*0.6(3)|ns|
|CLK6|tw(CLKOUT0L)|Pulse Duration, CLKOUT0 low|C*0.4(3)<br>C*0.6(3)|ns|
|CLK4|tc(WKUP_SYSCLKOUT0)|Cycle time minimum, WKUP_SYSCLKOUT0|10|ns|
|CLK5|tw(WKUP_SYSCLKOUT0H)|Pulse Duration, WKUP_SYSCLKOUT0 high|E*0.4(4)<br>E*0.6(4)|ns|
|CLK6|tw(WKUP_SYSCLKOUT0L)|Pulse Duration, WKUP_SYSCLKOUT0 low|E*0.4(4)<br>E*0.6(4)|ns|
|CLK4|tc(WKUP_OBSCLK0)|Cycle time minimum, WKUP_OBSCLK0|5|ns|
|CLK5|tw(WKUP_OBSCLK0H)|Pulse Duration, WKUP_OBSCLK0 high|D*0.45(5)<br>D*0.55(5)|ns|
|CLK6|tw(WKUP_OBSCLK0L)|Pulse Duration, WKUP_OBSCLK0 low|D*0.45(5)<br>D*0.55(5)|ns|
|CLK4|tc(WKUP_CLKOUT0)|Cycle time minimum, WKUP_CLKOUT0|5|ns|
|CLK5|tw(WKUP_CLKOUT0H)|Pulse Duration, WKUP_CLKOUT0 high|W*0.4(6)<br>W*0.6(6)|ns|
|CLK6|tw(WKUP_CLKOUT0L)|Pulse Duration, WKUP_CLKOUT0 low|W*0.4(6)<br>W*0.6(6)|ns|
|CLK4|tc(AUDIO_EXT_REFCLK0 )|Cycle time minimum, AUDIO_EXT_REFCLK0<br>(McASP Clock Source)|20|ns|
|||Cycle time minimum, AUDIO_EXT_REFCLK0<br>(PLL Clock Source)|10|ns|
|CLK5|tw(AUDIO_EXT_REFCLK0 H)|Pulse Duration, AUDIO_EXT_REFCLK0 high|G*0.4(7)<br>G*0.6(7)|ns|
|CLK6|tw(AUDIO_EXT_REFCLK0 L)|Pulse Duration, AUDIO_EXT_REFCLK0 low|G*0.4(7)<br>G*0.6(7)|ns|
|CLK4|tc(AUDIO_EXT_REFCLK1 )|Cycle time minimum, AUDIO_EXT_REFCLK1<br>(McASP Clock Source)|20|ns|
|||Cycle time minimum, AUDIO_EXT_REFCLK1<br>(PLL Clock Source)|10|ns|
|CLK5|tw(AUDIO_EXT_REFCLK1 H)|Pulse Duration, AUDIO_EXT_REFCLK1 high|J*0.4(8)<br>J*0.6(8)|ns|
|CLK6|tw(AUDIO_EXT_REFCLK1 L)|Pulse Duration, AUDIO_EXT_REFCLK1 low|J*0.4(8)<br>J*0.6(8)|ns|



(1) B = OBSCLK0 cycle time in ns. 

(2) F = OBSCLK1 cycle time in ns. 

(3) C = CLKOUT0 cycle time in ns. 

(4) E = WKUP_SYSCLKOUT0 cycle time in ns. 

(5) D = WKUP_OBSCLK0 cycle time in ns. 

(6) W = WKUP_CLKOUT0 cycle time in ns. 

(7) G = AUDIO_EXT_REFCLK0 cycle time in ns. 

(8) J = AUDIO_EXT_REFCLK1 cycle time in ns. 

**==> picture [285 x 50] intentionally omitted <==**

**----- Start of picture text -----**<br>
CLK4<br>CLK5 CLK6<br>Output Clock<br>**----- End of picture text -----**<br>


**Figure 6-19. Clock Switching Characteristics** 

## _**6.11.4 Clock Specifications**_ 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 101 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.4.1 Input Clocks / Oscillators** 

Various external clock inputs/outputs are needed to drive the device. Summary of these input clock signals is as follows: 

- WKUP_OSC0_XO/WKUP_OSC0_XI — external main crystal interface pins connected to the internal highfrequency oscillator (WKUP_HFOSC0), which is the default clock source for internal reference clock HFOSC0_CLKOUT. 

- LFOSC0_XO/LFOSC0_XI — external crystal interface pins connected to internal low-frequency oscillator (LFOSC0), which sources optional 32768Hz reference clock. 

- General purpose clock inputs 

   - WKUP_EXT_REFCLK0 — optional external system clock. 

   - EXT_REFCLK1 — optional external system clock. 

- External CPTS reference clock input 

   - CP_GEMAC_CPTS0_RFT_CLK — optional reference clock input for CPTS_RFT_CLK. 

- External audio reference clock inputs/outputs 

   - AUDIO_EXT_REFCLK[1:0] — optional McASP high-frequency input clocks when configured to operate as an input. 

For more information about Input clock interfaces, see _Clocking_ section in _Device Configuration_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

102 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.4.1.1 WKUP_OSC0 Internal Oscillator Clock Source**_ 

Figure 6-20 shows the recommended crystal circuit. All discrete components used to implement the oscillator circuit must be placed as close as possible to the WKUP_OSC0_XI and WKUP_OSC0_XO pins. 

**==> picture [196 x 196] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>WKUP_OSC0_XI WKUP_OSC0_XO<br>Crystal<br>CL1 CL2<br>PCB Ground<br>**----- End of picture text -----**<br>


**Figure 6-20. WKUP_OSC0 Crystal Implementation** 

The crystal must be in the fundamental mode of operation and parallel resonant. Table 6-18 summarizes the required electrical constraints. 

**Table 6-18. WKUP_OSC0 Crystal Circuit Requirements** 

|**PARAMETER**|**PARAMETER**|**PARAMETER**||**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|Fxtal|Crystal Parallel Resonance Frequency|||25|MHz|
|Fxtal|Crystal Frequency Stability and Tolerance||Ethernet RGMII and RMII<br>not used|±100|ppm|
||||Ethernet RGMII and RMII<br>using derived clock|±50||
|CL1+PCBXI|Capacitance of CL1+ CPCBXI|||12<br>24|pF|
|CL2+PCBXO|Capacitance of CL2+ CPCBXO|||12<br>24|pF|
|CL|Crystal Load Capacitance|||6<br>12|pF|
|Cshunt|Crystal Circuit Shunt Capacitance|ESRxtal= 30Ω|25MHz|7|pF|
|||ESRxtal= 40Ω|25MHz|5|pF|
|||ESRxtal= 50Ω|25MHz|5|pF|
|ESRxtal|Crystal Effective Series Resistance|||(1)|Ω|



(1) The maximum ESR of the crystal is a function of the crystal frequency and shunt capacitance. See the Cshunt parameter. 

When selecting a crystal, the system design must consider temperature and aging characteristics of the crystal based on worst case environment and expected life expectancy of the system. 

Table 6-19 details the switching characteristics of the oscillator. 

**Table 6-19. WKUP_OSC0 Switching Characteristics - Crystal Mode** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|
|CXI|XI Capacitance|0.812|pF|
|CXO|XO Capacitance|0.848|pF|
|CXIXO|XI to XO Mutual Capacitance|0.01|pF|
|ts|Start-up Time|4|ms|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 103 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**==> picture [365 x 180] intentionally omitted <==**

**----- Start of picture text -----**<br>
VDD_CORE (min.)<br>VDD_CORE<br>VSS<br>VDDS_OSC0 (min.) VDDS_OSC0<br>VSS WKUP_OSC0_XO<br>t sX<br>Time<br>Voltage<br>**----- End of picture text -----**<br>


**==> picture [32 x 3] intentionally omitted <==**

**----- Start of picture text -----**<br>
AM65x_MCU_OSC_STARTUP_02<br>**----- End of picture text -----**<br>


**Figure 6-21. WKUP_OSC0 Start-up Time** 

## _**6.11.4.1.1.1 Load Capacitance**_ 

The crystal circuit must be designed such that it applies the appropriate capacitive load to the crystal, as defined by the crystal manufacturer. The capacitive load, CL, of this circuit is a combination of discrete capacitors CL1, CL2, and several parasitic contributions. PCB signal traces which connect crystal circuit components to WKUP_OSC0_XI and WKUP_OSC0_XO have parasitic capacitance to ground, CPCBXI and CPCBXO, where the PCB designer should be able to extract parasitic capacitance for each signal trace. The WKUP_OSC0 circuits and device package have combined parasitic capacitance to ground, CPCBXI and CPCBXO, where these parasitic capacitance values are defined in Table 6-19. 

**==> picture [281 x 159] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>Crystal Circuit PCB<br>Components Signal Traces<br>WKUP_OSC0_XI<br>CL1 CPCBXI CXI<br>CL2 CPCBXO CXO<br>WKUP_OSC0_XO<br>**----- End of picture text -----**<br>


**Figure 6-22. Load Capacitance** 

Load capacitors, CL1 and CL2 in Figure 6-20, should be chosen such that the below equation is satisfied. CL in the equation is the load specified by the crystal manufacturer. 

CL = [(CL1 + CPCBXI + CXI) × (CL2 + CPCBXO + CXO)] / [(CL1 + CPCBXI + CXI) + (CL2 + CPCBXO + CXO)] 

To determine the value of CL1 and CL2, multiply the capacitive load value CL by 2. Using this result, subtract the combined values of CPCBXI + CXI to determine the value of CL1 and the combined values of CPCBXO + CXO to determine the value of CL2. For example, if CL = 10pF, CPCBXI = 2.9pF, CXI = 0.5pF, CPCBXO = 3.7pF, CXO = 0.5pF, the value of CL1 = [(2CL) - (CPCBXI + CXI)] = [(2 × 10pF) - 2.9pF - 0.5pF)] = 16.6pF and CL2 = [(2CL) - (CPCBXO + CXO)] = [(2 × 10pF) - 3.7pF - 0.5pF)] = 15.8pF 

Copyright © 2025 Texas Instruments Incorporated 

104 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.4.1.1.2 Shunt Capacitance**_ 

The crystal circuit must also be designed such that it does not exceed the maximum shunt capacitance for WKUP_OSC0 operating conditions defined in Table 6-18. Shunt capacitance, Cshunt, of the crystal circuit is a combination of crystal shunt capacitance and parasitic contributions. PCB signal traces which connect crystal circuit components to WKUP_OSC0 have mutual parasitic capacitance to each other, CPCBXIXO, where the PCB designer should be able to extract mutual parasitic capacitance between these signal traces. The device package also has mutual parasitic capacitance, CXIXO, where this mutual parasitic capacitance value is defined in Table 6-19. 

PCB routing should be designed to minimize mutual capacitance between XI and XO signal traces. This is typically done by keeping signal traces short and not routing them in close proximity. Mutual capacitance can also be minimized by placing a ground trace between these signals when the layout requires them to be routed in close proximity. It is important to minimize the mutual capacitance on the PCB to provide as much margin as possible when selecting a crystal. 

**==> picture [281 x 159] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>Crystal Circuit PCB<br>Components Signal Traces<br>WKUP_OSC0_XI<br>CO CPCBXIXO CXIXO<br>WKUP_OSC0_XO<br>**----- End of picture text -----**<br>


**Figure 6-23. Shunt Capacitance** 

A crystal should be chosen such that the below equation is satisfied. CO in the equation is the maximum shunt capacitance specified by the crystal manufacturer. 

## Cshunt ≥ CO + CPCBXIXO + CXIXO 

For example, the equation would be satisfied when the crystal being used is 25MHz with an ESR = 30Ω, CPCBXIXO = 0.04pF, CXIXO = 0.01pF, and shunt capacitance of the crystal is less than or equal to 6.95pF. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 105 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.4.1.2 WKUP_OSC0 LVCMOS Digital Clock Source**_ 

Figure 6-24 shows the recommended oscillator connections when WKUP_OSC0_XI is connected to a 1.8V LVCMOS square-wave digital clock source. 

## **Note** 

1. A DC steady-state condition is not allowed on WKUP_OSC0_XI when the oscillator is powered up. This is not allowed because WKUP_OSC0_XI is internally AC coupled to a comparator that can enter an unknown state when DC is applied to the input. Therefore, application software must power down WKUP_OSC0 any time WKUP_OSC0_XI is not toggling between logic states. 

2. The LVCMOS clock signal sourcing the WKUP_OSC0_XI input must have monotonic transitions. The clock source should be connected to WKUP_OSC0_XI with a point-to-point connection, via a series termination resistor placed near the clock source. The series termination resistor value should match the clock source output impedance to the transmission line impedance. For example, the series termination resistor value needs to be 20 ohms if the clock source has an output impedance of 30 ohms and the PCB signal trace has a characteristic impedance of 50 ohms. This allows the reflection that returns from the far end of the un-terminated transmission line to be completely absorbed such that is does not introduce any non-monotonic events on the signal. 

3. The PCB trace length connecting the LVCMOS clock source to WKUP_OSC0_XI should be minimized. This reduces capacitive loading and decreases probability of external noise sources coupling into the clock signal. Reduced capacitive loading improves rise/fall times of the clock signal which reduces the probability of jitter being introduced in the system. 

**==> picture [204 x 157] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>WKUP_OSC0_XI WKUP_OSC0_XO<br>PCB Ground<br>**----- End of picture text -----**<br>


**Figure 6-24. 1.8V LVCMOS-Compatible Clock Input** 

Copyright © 2025 Texas Instruments Incorporated 

106 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **Table 6-20. WKUP_OSC0 LVCMOS Digital Clock Source Requirements** 

|**PARAMETER**|**PARAMETER**||**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|Fxtal|Frequency||25|MHz|
||Frequency Stability and Tolerance|Ethernet RGMII and RMII<br>not used|±100|ppm|
|||Ethernet RGMII and RMII<br>using derived clock|±50||
|DC|Duty Cycle||45<br>55|%|
|tR/F|Rise/Fall Time (10%-90% rise, 90%-10% fall)||4(1)|ns|
|JPeriod(RMS)|Period Jitter, RMS (100k samples)||20|ps|
|JPeriod(PK-PK)|Period Jitter, Peak to Peak (100k samples)||300|ps|
|JPhase(RMS)|Phase Jitter, RMS (BW 100Hz to 1MHz)||10(2)|ps|



(1) Most LVCMOS oscillator datasheets define their maximum Output Rise/Fall times with a capacitive load much larger than the actual load that will be applied by the combined PCB trace capacitance and WKUP_OSC0_XI input capacitance. It should not be difficult to find a LVCMOS oscillator that meets this requirement. However, the system designer must confirm the LVCMOS oscillator selected will provide the appropriate rise/fall time to WKUP_OSC0_XI input. 

(2) Most LVCMOS oscillator datasheets define their max RMS Phase Jitter using a larger bandwidth integration range than required by this device. To get a more appropriate value, it may be necessary to contact the LVCMOS oscillator manufacture and ask them to provide a maximum RMS Phase Jitter using the same bandwidth integration range that has been defined for this parameter. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 107 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.4.1.3 LFOSC0 Internal Oscillator Clock Source**_ 

Figure 6-25 shows the recommended crystal circuit. It is recommended that preproduction printed-circuit board (PCB) designs include the two optional resistors Rbias and Rd in case they are required for proper oscillator operation when combined with production crystal circuit components. In most cases, Rbias is not required and Rd is a 0-Ω resistor. These resistors may be removed from production PCB designs after evaluating oscillator performance with production crystal circuit components installed on preproduction PCBs. 

**==> picture [196 x 193] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>LFOSC0_XI LFOSC0_XO<br>Rd<br>Crystal (Optional)<br>(Optional) Rbias<br>Cf1 Cf2<br>PCB Ground<br>**----- End of picture text -----**<br>


**Figure 6-25. LFOSC0 Crystal Implementation** 

Table 6-21 presents LFXOSC modes of operation. 

**Table 6-21. LFXOSC Modes of Operation** 

|**MODE**|**BP_C**|**PD_C**|**XI**|**XO**|**CLK_OUT**|**DESCRIPTION**|
|---|---|---|---|---|---|---|
|ACTIVE|0|0|XTAL|XTAL|CLK_OUT|Active oscillator mode providing 32kHz|
|PWRDN|0|1|X|PD|LOW|Output will be pulled down to LOW. PAD to be tri-stated. Active mode<br>disabled|
|BYPASS|1|0|CLK|PD|CLK|XI is driven by external clock source. XO is pulled down to LOW. Due to ESD<br>diode to supply, XI should not be driven unless oscillator supply is present.|



## **Note** 

User should set RTC_RTC_LFXOSC_TRIM[18:16] i_mult = 3b’001 for CL in the range 6pf to 9.5pf. RTC_RTC_LFXOSC_TRIM [18:16] i_mult = 3b’010 for CL in the range 8.5pf to 12pf. Default setting is 3b’010. 

## **Note** 

The load capacitors, Cf1 and Cf2 in Figure 6-26, should be chosen such that the below equation is satisfied. CL in the equation is the load specified by the crystal manufacturer. All discrete components used to implement the oscillator circuit should be placed as close as possible to the associated oscillator LFOSC0_XI, LFOSC0_XO, and VSS pins. 

**==> picture [85 x 35] intentionally omitted <==**

**----- Start of picture text -----**<br>
Cf1Cf2<br>CL=<br>(Cf1+Cf2)<br>**----- End of picture text -----**<br>


J7ES_CL_MATH_03 

**Figure 6-26. Load Capacitance Equation** 

Copyright © 2025 Texas Instruments Incorporated 

108 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

The crystal must be in the fundamental mode of operation and parallel resonant. Table 6-22 summarizes the required electrical constraints. 

**Table 6-22. LFOSC0 Crystal Electrical Characteristics** 

|**NAME**|**DESCRIPTION**||**MIN**<br>**TYP**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|fp|Parallel resonance crystal frequency||32768|Hz|
||Crystal Frequency Stability and Tolerance||±100|PPM|
|Cf1|Cf1load capacitance for crystal parallel resonance with Cf1= Cf2||12<br>24|pF|
|Cf2|Cf2load capacitance for crystal parallel resonance with Cf1= Cf2||12<br>24|pF|
|Cshunt|Shunt capacitance|ESRxtal – 40 kΩ|4|pF|
|||ESRxtal – 60 kΩ|3|pF|
|||ESRxtal – 80 kΩ|2|pF|
|||ESRxtal – 100 kΩ|1|pF|
|ESR|Crystal effective series resistance||(1)|Ω|



(1) The maximum ESR of the crystal is a function of the crystal frequency and shunt capacitance. See the Cshunt parameter. 

When selecting a crystal, the system design must consider the temperature and aging characteristics of a based on the worst case environment and expected life expectancy of the system. 

Table 6-23 details the switching characteristics of the oscillator and the requirements of the input clock. 

**Table 6-23. LFOSC0 Switching Characteristics – Crystal Mode** 

**==> picture [500 x 233] intentionally omitted <==**

**----- Start of picture text -----**<br>
NAME DESCRIPTION MIN TYP MAX UNIT<br>fxtal Oscillation frequency 32768 Hz<br>tsX Start-up time 96.5 ms<br>VDD_CORE (min.)<br>VDD_CORE<br>VSS<br>VDDS_OSC0 (min.) VDDS_OSC0<br>VSS LFOSC0_XO<br>t sX<br>Time<br>Voltage<br>**----- End of picture text -----**<br>


**Figure 6-27. LFOSC0 Start-up Time** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 109 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.4.1.4 LFOSC0 LVCMOS Digital Clock Source**_ 

Figure 6-28 shows the recommended oscillator connections when LFOSC0_XI is connected to a 1.8V LVCMOS square-wave digital clock source. 

**==> picture [204 x 157] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>LFOSC0_XI LFOSC0_XO<br>PCB Ground<br>**----- End of picture text -----**<br>


**Figure 6-28. 1.8V LVCMOS-Compatible Clock Input** 

## _**6.11.4.1.5 LFOSC0 Not Used**_ 

Figure 6-29 shows the recommended oscillator connections when LFOSC0 is not used. 

**==> picture [196 x 119] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>LFOSC0_XI LFOSC0_XO<br>NC<br>PCB Ground<br>**----- End of picture text -----**<br>


**Figure 6-29. LFOSC0 Not Used** 

Copyright © 2025 Texas Instruments Incorporated 

110 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.4.2 Output Clocks** 

The device provides several system clock outputs. Summary of these output clocks are as follows: 

- 

## **WKUP_SYSCLKOUT0** 

   - WKUP_PLL0_HSDIV0_CLKOUT (PER_SYSCLK0) divided by 4 and sent out of the device as WKUP_SYSCLKOUT0. This clock output is provided for test and debug purposes only. 

- **WKUP_OBSCLK0** 

   - This output can only be used as a functional clock source when WKUP_OBSCLK_OUTMUX is used to select the direct output from WKUP_HFOSC0. 

   - This output can only be used for test and debug purposes when selecting any other clock source. 

- **WKUP_CLKOUT0** 

   - This output can only be used as a functional clock source when WKUP_CLKOUTMUX is used to select LFOSC0_CLKOUT, DEVICE_CLKOUT_32K, or the direct output from WKUP_HFOSC0. 

- This output can only be used for test and debug purposes when selecting any other clock source. 

- • **CLKOUT0** 

   - CLKOUT0 is the Ethernet subsystem clock (MAIN_PLL0_HSDIV6_CLKOUT) divided-by-5 or dividedby-10. This clock output was provided as an optional source to the external PHY. When configured to operate as the RMII Clock source (50MHz) the signal must also be routed back to the respective RMII[x]_REF_CLK pin for proper device operation. 

- **OBSCLK[1:0]** 

   - 

      - These outputs can only be used as a functional clock sources when OBSCLK0_CTRL is used to select the direct output from WKUP_HFOSC0. 

   - These outputs can only be used for test and debug purposes when selecting any other clock source. 

- **AUDIO_EXT_REFCLK[1:0]** 

   - 

Option of sourcing one of six McASP high-frequency audio reference clocks, 

MAIN_PLL0_HSDIV8_CLKOUT, or WKUP_PLL0_HSDIV1_CLKOUT when configured to operate as an output. 

## **6.11.4.3 PLLs** 

Power is supplied to the Phase-Locked Loop circuits (PLLs) by internal regulators that derive their power from off-chip power-sources. 

There is one PLL in the WKUP domain: 

- WKUP_PLL0 (WKUP PLL) 

There are three PLLs in the MAIN domain: 

- MAIN_PLL0 (MAIN PLL) 

- MAIN_PLL8 (ARM0 PLL) 

- MAIN_PLL17 (DSS PLL0) 

The system designer should consider the reference clock source start-up time and the PLL lock requirements before configuring and using any of the PLL outputs as clock sources. The device reference clock input requirements are defined in Section 6.11.4.1, _Input Clocks / Oscillators_ . PLL configuration details are described in the device TRM. 

For more information on PLLs, see the _PLL_ subsection in the _Clocking_ subsection of the _Device Configuration_ section in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 111 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.4.4 Recommended System Precautions for Clock and Control Signal Transitions** 

All clock and strobe signals must transition between VIH and VIL (or between VIL and VIH) in a monotonic manner. 

Monotonic transitions are more likely to occur with fast signal transitions. It is easy for noise to create nonmonotonic events on a signal with slow transitions. Therefore, avoid slow signal transitions on all clock and control signals since they are more likely to generate glitches inside the device. 

Copyright © 2025 Texas Instruments Incorporated 

112 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5 Peripherals**_ 

## **6.11.5.1 CPSW3G** 

For more details about features and additional description information on the device Gigabit Ethernet MAC, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

## _**6.11.5.1.1 CPSW3G MDIO Timing**_ 

Table 6-24, Table 6-25, Table 6-26, and Figure 6-30 present timing conditions, timing requirements, and switching characteristics for CPSW3G MDIO. 

**Table 6-24. CPSW3G MDIO Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|0.9<br>3.6|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|10<br>470|pF|
|**PCB CONNECTIVITY REQUIREMENTS**||||
|td(Trace Delay)|Propagation delay of each trace|0<br>5|ns|
|td(Trace Mismatch Delay)|Propagation delay mismatch across all traces|1|ns|



**Table 6-25. CPSW3G MDIO Timing Requirements** 

see Figure 6-30 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|MDIO1|tsu(MDIO_MDC)|Setup time, MDIO[x]_MDIO valid before MDIO[x]_MDC high|45|ns|
|MDIO2|th(MDC_MDIO)|Hold time, MDIO[x]_MDIO valid after MDIO[x]_MDC high|0|ns|



## **Table 6-26. CPSW3G MDIO Switching Characteristics** 

|**Table 6-26. CPSW3G MDIO Switching Characteristics**|**Table 6-26. CPSW3G MDIO Switching Characteristics**|**Table 6-26. CPSW3G MDIO Switching Characteristics**|**Table 6-26. CPSW3G MDIO Switching Characteristics**|**Table 6-26. CPSW3G MDIO Switching Characteristics**|
|---|---|---|---|---|
|seeFigure 6-30|||||
|**NO.**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
|MDIO3|tc(MDC)|Cycle time, MDIO[x]_MDC|400|ns|
|MDIO4|tw(MDCH)|Pulse Duration, MDIO[x]_MDC high|160|ns|
|MDIO5|tw(MDCL)|Pulse Duration, MDIO[x]_MDC low|160|ns|
|MDIO7|td(MDC_MDIO)|Delay time, MDIO[x]_MDC low to MDIO[x]_MDIO valid|-10<br>10|ns|



**==> picture [410 x 173] intentionally omitted <==**

**----- Start of picture text -----**<br>
MDIO3<br>MDIO4<br>MDIO5<br>MDIO1<br>MDIO2<br>MDIO7<br>CPSW2G_MDIO_TIMING_01<br>**----- End of picture text -----**<br>


**==> picture [49 x 124] intentionally omitted <==**

**----- Start of picture text -----**<br>
MDIO[x]_MDC<br>MDIO[x]_MDIO<br>(input)<br>MDIO[x]_MDIO<br>(output)<br>**----- End of picture text -----**<br>


## **Figure 6-30. CPSW3G MDIO Timing Requirements and Switching Characteristics** 

Copyright © 2025 Texas Instruments Incorporated 

113 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.1.2 CPSW3G RMII Timing**_ 

Table 6-27, Table 6-28, Figure 6-31, Table 6-29, Figure 6-32, Table 6-30, and Figure 6-33 present timing conditions, timing requirements, and switching characteristics for CPSW3G RMII. 

**Table 6-27. CPSW3G RMII Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|0.18<br>5|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|3<br>25|pF|



## **Table 6-28. RMII[x]_REF_CLK Timing Requirements – RMII Mode** 

**==> picture [64 x 9] intentionally omitted <==**

**----- Start of picture text -----**<br>
see Figure 6-31<br>**----- End of picture text -----**<br>


|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|RMII1|tc(REF_CLK)|Cycle time, RMII[x]_REF_CLK|19.999<br>20.001|ns|
|RMII2|tw(REF_CLKH)|Pulse Duration, RMII[x]_REF_CLK High|7<br>13|ns|
|RMII3|tw(REF_CLKL)|Pulse Duration, RMII[x]_REF_CLK Low|7<br>13|ns|
|RMII[x]_REF_CLK<br>~~RMII2~~<br>~~RMII3~~<br>~~RMII1~~|||||



**Figure 6-31. CPSW3G RMII[x]_REF_CLK Timing Requirements – RMII Mode** 

**Table 6-29. RMII[x]_RXD[1:0], RMII[x]_CRS_DV, and RMII[x]_RX_ER Timing Requirements – RMII Mode** 

see Figure 6-32 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|RMII4|tsu(RXD-REF_CLK)|Setup time, RMII[x]_RXD[1:0] valid before RMII[x]_REF_CLK|4|ns|
||tsu(CRS_DV-REF_CLK)|Setup time, RMII[x]_CRS_DV valid before RMII[x]_REF_CLK|4|ns|
||tsu(RX_ER-REF_CLK)|Setup time, RMII[x]_RX_ER valid before RMII[x]_REF_CLK|4|ns|
|RMII5|th(REF_CLK-RXD)|Hold time RMII[x]_RXD[1:0] valid after RMII[x]_REF_CLK|2|ns|
||th(REF_CLK-CRS_DV)|Hold time, RMII[x]_CRS_DV valid after RMII[x]_REF_CLK|2|ns|
||th(REF_CLK-RX_ER)|Hold time, RMII[x]_RX_ER valid after RMII[x]_REF_CLK|2|ns|
|RMII4<br>~~RMII5~~<br>RMII[x]_<br>_<br>REF CLK<br>RMII[x]_RXD[1:0], RMII[x]_CRS_DV,<br>RMII[x]RXER|||||



**Figure 6-32. CPSW3G RMII[x]_RXD[1:0], RMII[x]_CRS_DV, RMII[x]_RX_ER Timing Requirements – RMII Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

114 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-30. RMII[x]_TXD[1:0], and RMII[x]_TX_EN Switching Characteristics – RMII Mode** 

see Figure 6-33 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|RMII6|td(REF_CLK-TXD)|Delay time, RMII[x]_REF_CLK High to RMII[x]_<br>TXD[1:0] valid|2<br>10|ns|
||td(REF_CLK-TX_EN)|Delay time, RMII[x]_REF_CLK to RMII[x]_TX_EN<br>valid|2<br>10|ns|
|RMII[x]_TXD[1:0], RMII[x]_TX_EN<br>~~RMII6~~<br>RMII[x]_REF_CLK|||||



**Figure 6-33. RMII[x]_TXD[1:0], and RMII[x]_TX_EN Switching Characteristics – RMII Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 115 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.1.3 CPSW3G RGMII Timing**_ 

Table 6-31, Table 6-32, Table 6-33, Figure 6-34, Table 6-34, Table 6-35, and Figure 6-35 present timing conditions, timing requirements, and switching characteristics for CPSW3G RGMII. 

**Table 6-31. CPSW3G RGMII Timing Conditions** 

|**PARAMETER**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**INPUT CONDITIONS**|||||
|SRI|Input slew rate||1.44<br>5|V/ns|
|**OUTPUT CONDITIONS**|||||
|CL|Output load capacitance||2<br>20|pF|
|**PCB CONNECTIVITY REQUIREMENTS**|||||
|td(Trace Mismatch<br>Delay)|Propagation delay mismatch across all traces|RGMII[x]_RXC,<br>RGMII[x]_RD[3:0],<br>RGMII[x]_RX_CTL|50|ps|
|||RGMII[x]_TXC,<br>RGMII[x]_TD[3:0],<br>RGMII[x]_TX_CTL|50|ps|



Copyright © 2025 Texas Instruments Incorporated 

116 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 6-32. RGMII[x]_RXC Timing Requirements – RGMII Mode** 

see Figure 6-34 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|RGMII1|tc(RXC)|Cycle time, RGMII[x]_RXC|10Mbps|360<br>440|ns|
||||100Mbps|36<br>44|ns|
||||1000Mbps|7.2<br>8.8|ns|
|RGMII2|tw(RXCH)|Pulse duration, RGMII[x]_RXC high|10Mbps|160<br>240|ns|
||||100Mbps|16<br>24|ns|
||||1000Mbps|3.6<br>4.4|ns|
|RGMII3|tw(RXCL)|Pulse duration, RGMII[x]_RXC low|10Mbps|160<br>240|ns|
||||100Mbps|16<br>24|ns|
||||1000Mbps|3.6<br>4.4|ns|



**Table 6-33. RGMII[x]_RD[3:0], and RGMII[x]_RX_CTL Timing Requirements – RGMII Mode** 

see Figure 6-34 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|RGMII4|tsu(RD-RXC)|Setup time, RGMII[x]_RD[3:0] valid before RGMII[x]_RXC<br>high/low|10Mbps|1|ns|
||||100Mbps|1|ns|
||||1000Mbps|1|ns|
||tsu(RX_CTL-RXC)|Setup time, RGMII[x]_RX_CTL valid before RGMII[x]_RXC<br>high/low|10Mbps|1|ns|
||||100Mbps|1|ns|
||||1000Mbps|1|ns|
|RGMII5|th(RXC-RD)|Hold time, RGMII[x]_RD[3:0] valid after RGMII[x]_RXC<br>high/low|10Mbps|1|ns|
||||100Mbps|1|ns|
||||1000Mbps|1|ns|
||th(RXC-RX_CTL)|Hold time, RGMII[x]_RX_CTL valid after RGMII[x]_RXC<br>high/low|10Mbps|1|ns|
||||100Mbps|1|ns|
||||1000Mbps|1|ns|



**==> picture [445 x 145] intentionally omitted <==**

**----- Start of picture text -----**<br>
RGMII1<br>RGMII2<br>RGMII3<br>RGMII[x]_RXC(A)<br>RGMII4<br>RGMII5<br>RGMII[x]_RD 3[ :0](B) 1st Half-byte 2nd Half-byte<br>(B)<br>RGMII[x]_RX_CTL RXDV RXERR<br>**----- End of picture text -----**<br>


- A. RGMII[x]_RXC must be externally delayed relative to the data and control pins. 

- B. Data and control information is received using both edges of the clocks. RGMII[x]_RD[3:0] carries data bits 3-0 on the rising edge of RGMII[x]_RXC and data bits 7-4 on the falling edge of RGMII[x]_RXC. Similarly, RGMII[x]_RX_CTL carries RXDV on rising edge of RGMII[x]_RXC and RXERR on falling edge of RGMII[x]_RXC. 

## **Figure 6-34. CPSW3G RGMII[x]_RXC, RGMII[x]_RD[3:0], RGMII[x]_RX_CTL Timing Requirements - RGMII Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 117 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 6-34. RGMII[x]_TXC Switching Characteristics – RGMII Mode** 

see Figure 6-35 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|RGMII6|tc(TXC)|Cycle time, RGMII[x]_TXC|10Mbps|360<br>440|ns|
||||100Mbps|36<br>44|ns|
||||1000Mbps|7.2<br>8.8|ns|
|RGMII7|tw(TXCH)|Pulse duration, RGMII[x]_TXC high|10Mbps|160<br>240|ns|
||||100Mbps|16<br>24|ns|
||||1000Mbps|3.6<br>4.4|ns|
|RGMII8|tw(TXCL)|Pulse duration, RGMII[x]_TXC low|10Mbps|160<br>240|ns|
||||100Mbps|16<br>24|ns|
||||1000Mbps|3.6<br>4.4|ns|



**Table 6-35. RGMII[x]_TD[3:0] and RGMII[x]_TX_CTL Switching Characteristics – RGMII Mode** 

see Figure 6-35 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|RGMII9|tosu(TD-TXC)|Output setup time(1), RGMII[x]_TD[3:0] valid to<br>RGMII[x]_TXC high/low|10Mbps|1.2|ns|
||||100Mbps|1.2|ns|
||||1000Mbps|1.2|ns|
||tosu(TX_CTL-TXC)|Output setup time(1), RGMII[x]_TX_CTL valid to<br>RGMII[x]_TXC high/low|10Mbps|1.2|ns|
||||100Mbps|1.2|ns|
||||1000Mbps|1.2|ns|
|RGMII10|toh(TXC-TD)|Output hold time(1), RGMII[x]_TD[3:0] valid after<br>RGMII[x]_TXC high/low|10Mbps|1.2|ns|
||||100Mbps|1.2|ns|
||||1000Mbps|1.2|ns|
||toh(TXC-TX_CTL)|Output hold time(1), RGMII[x]_TX_CTL valid after<br>RGMII[x]_TXC high/low|10Mbps|1.2|ns|
||||100Mbps|1.2|ns|
||||1000Mbps|1.2|ns|



- (1) Output setup/hold times are defining a delay relationship of the transmit data and control outputs relative to the transmit clock output, but this output relationship is being presented as the minimum setup/hold times provided to the attached receiver. This approach matches how the output timing relationships are defined in the RGMII specification. 

**==> picture [447 x 134] intentionally omitted <==**

**----- Start of picture text -----**<br>
RGMII6<br>RGMII7<br>RGMII8<br>RGMII[x]_TXC(A)<br>RGMII9<br>RGMII[x]_TD[3:0] (B) 1st Half-byte 2nd Half-byte<br>RGMII10<br>(B)<br>RGMII[x]_TX_CTL TXEN TXERR<br>**----- End of picture text -----**<br>


- A. TXC is delayed internally before being driven to the RGMII[x]_TXC pin. This internal delay is always enabled. 

- B. Data and control information is received using both edges of the clocks. RGMII[x]_TD[3:0] carries data bits 3-0 on the rising edge of RGMII[x]_TXC and data bits 7-4 on the falling edge of RGMII[x]_TXC. Similarly, RGMII[x]_TX_CTL carries TXEN on rising edge of RGMII[x]_TXC and TXERR on falling edge of RGMII[x]_TXC. 

## **Figure 6-35. CPSW3G RGMII[x]_TXC, RGMII[x]_TD[3:0], and RGMII[x]_TX_CTL Switching Characteristics - RGMII Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

118 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.2 CPTS** 

Table 6-36, Table 6-37, Figure 6-36, Table 6-38, and Figure 6-37 present timing conditions, timing requirements, and switching characteristics for CPTS. 

## **Table 6-36. CPTS Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|0.5<br>5|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|2<br>10|pF|



## **Table 6-37. CPTS Timing Requirements** 

## see Figure 6-36 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|T1|tw(HWTSPUSHH)|Pulse duration, HWnTSPUSH high|12P(1)+ 2|ns|
|T2|tw(HWTSPUSHL)|Pulse duration, HWnTSPUSH low|12P(1)+ 2|ns|
|T3|tc(RFT_CLK)|Cycle time, RFT_CLK|5<br>8|ns|
|T4|tw(RFT_CLKH)|Pulse duration, RFT_CLK high|0.45T(2)|ns|
|T5|tw(RFT_CLKL)|Pulse duration, RFT_CLK low|0.45T(2)|ns|
|(1)<br>P = functional clock period in ns.<br>(2)<br>T = RFT_CLK cycle time in ns.<br>RFT_CLK<br>T3<br>T4<br>T5<br>HWn_TSPUSH<br>T1<br>T2|||||



**Figure 6-36. CPTS Timing Requirements** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 119 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-38. CPTS Switching Characteristics** 

see Figure 6-37 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**SOURCE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|T6|tw(TS_COMPH)|Pulse duration, TS_COMP high||36P(1)- 2|ns|
|T7|tw(TS_COMPL)|Pulse duration, TS_COMP low||36P(1)- 2|ns|
|T8|tw(TS_SYNCH)|Pulse duration, TS_SYNC high||36P(1)- 2|ns|
|T9|tw(TS_SYNCL)|Pulse duration, TS_SYNC low||36P(1)- 2|ns|
|T10|tw(SYNC_OUTH)|Pulse duration, SYNCn_OUT high|TS_SYNC|36P(1)- 2|ns|
||||GENF|5P(1)- 2|ns|
|T11|tw(SYNC_OUTL)|Pulse duration, SYNCn_OUT low|TS_SYNC|36P(1)- 2|ns|
||||GENF|5P(1)- 2|ns|



(1) P = functional clock period in ns. 

**==> picture [298 x 123] intentionally omitted <==**

**----- Start of picture text -----**<br>
T6 T7<br>TS_COMP<br>T8 T9<br>TS_SYNC<br>T10 T11<br>SYNCn_OUT<br>**----- End of picture text -----**<br>


**Figure 6-37. CPTS Switching Characteristics** 

For more information, see _Common Platform Time Sync (CPTS)_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

120 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.3 DDRSS** 

For more details about features and additional description information on the device (LP)DDR4 Memory Interface, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

Table 6-39 and Figure 6-38 present switching characteristics for DDRSS. 

**Table 6-39. DDRSS Switching Characteristics** 

|seeFigure 6-38|seeFigure 6-38|seeFigure 6-38||||
|---|---|---|---|---|---|
|**NO.**|**PARAMETER**||**DDR TYPE**|**MIN**<br>**MAX**|**UNIT**|
|1|tc(DDR_CKP/<br>DDR_CKN)|Cycle time, DDR_CKP and DDR_CKN|LPDDR4|1.25(1)<br>20|ns|
||||DDR4|1.25(1)<br>1.6|ns|
|(1)<br>Minimum DDR clock Cycle time will be limited based on the specific memory type (vendor) used in a system and by PCB<br>implementation. Refer toDDR Board Design and Layout Guidelinesfor the proper PCB implementation to achieve maximum DDR<br>frequency.<br>DDR0_CKP<br>~~1~~<br>DDR0_CKN||||||



**Figure 6-38. DDRSS Switching Characteristics** 

For more information, see _DDR Subsystem (DDRSS)_ section in _Memory Controllers_ chapter in the device TRM. 

## **6.11.5.4 DSI** 

## **Note** 

For more information, see the _MIPI Display Serial Interface (DSI) Controller_ section in the device TRM. The DSI transmitter controller is connected to device port instances named DSITXn, where n is the instance number. 

The DSI transmitter controller and associated D-PHY implements a DSI port (DSITX0) compliant with the MIPI D-PHY specification v1.2 and the MIPI DSI specification v1.3, with 4 differential data lanes plus 1 differential clock lane operating in synchronous double data rate mode. For DSI timing details, see the respective MIPI specifications mentioned above. 

- Support up to 4.8Gbps via 1-, 2-, 3- or 4-lane data transfer modes up to 2.5Gbps per lane 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 121 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.5 DSS** 

Table 6-40, Table 6-41, Figure 6-39, Table 6-42 and Figure 6-40 present timing conditions, timing requirements, and switching characteristics for DSS. 

**Table 6-40. DSS Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|1.44<br>26.4|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|1.5<br>5|pF|
|**PCB CONNECTIVITY REQUIREMENTS**||||
|td(Trace Mismatch Delay)|Propagation delay mismatch across all traces|100|ps|



## **Table 6-41. DSS External Pixel Clock Timing Requirements** 

## see Figure 6-39 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|D6|tc(extpclkin)|Cycle time, VOUT(x)_EXTPCLKIN(2)|6.06|ns|
|D7|tw(extpclkinL)|Pulse duration, VOUT(x)_EXTPCLKIN(2)low|0.475P(1)|ns|
|D8|tw(extpclkinH)|Pulse duration, VOUT(x)_EXTPCLKIN(2)high|0.475P(1)|ns|



(1) P = VOUT(x)_EXTPCLKIN cycle time in ns 

- (2) x in VOUT(x) = 0 

**==> picture [325 x 78] intentionally omitted <==**

**----- Start of picture text -----**<br>
D7<br>D6 D8<br>VOUT(x)_EXTPCLKIN<br>VOUT(x)_EXTPCLKIN<br>**----- End of picture text -----**<br>


**==> picture [114 x 54] intentionally omitted <==**

**----- Start of picture text -----**<br>
Falling-edge Clock Reference<br>Rising-edge Clock Reference<br>DPI_TIMING_02<br>**----- End of picture text -----**<br>


## **Figure 6-39. DSS External Pixel Clock Timing Requirements** 

Copyright © 2025 Texas Instruments Incorporated 

122 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 6-42. DSS Switching Characteristics** 

see Figure 6-40 

|**NO.**|**PARAMETER**||**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|D1|tc(pclk)|Cycle time, VOUT(x)_PCLK(2)||6.06|ns|
|D2|tw(pclkL)|Pulse duration, VOUT(x)_PCLK(2)low|Internal PLL|0.475P(1)- 0.3|ns|
||||EXTPCLKIN|Y(3)- 0.45|ns|
|D3|tw(pclkH)|Pulse duration, VOUT(x)_PCLK(2)high|Internal PLL|0.475P(1)-0.3|ns|
||||EXTPCLKIN|Z(4)- 0.45|ns|
|D4|td(pclkV-dataV)|Delay time, VOUT(x)_PCLK(2)transition to<br>VOUT(x)_DATA[23:0](2)transition|Internal PLL|-0.68<br>1.78|ns|
||||EXTPCLKIN|-0.68<br>1.78|ns|
|D5|td(pclkV-ctrlL)|Delay time, VOUT(x)_PCLK(2)transition to control signals<br>VOUT(x)_VSYNC(2), VOUT(x)_HSYNC(2), VOUT(x)_DE(2)<br>falling edge|Internal PLL|-0.68<br>1.78|ns|
||||EXTPCLKIN|-0.68<br>1.78|ns|



- (1) P = VOUT(x)_PCLK cycle time in ns 

- (2) x in VOUT(x) = 0 

- (3) Y = tw(extpclkinL), parameter D7 from Table 6-41, _DSS External Pixel Clock Timing Requirements_ 

- (4) Z = tw(extpclkinH), parameter D8 from Table 6-41, _DSS External Pixel Clock Timing Requirements_ 

**==> picture [412 x 233] intentionally omitted <==**

**----- Start of picture text -----**<br>
D2<br>D1 D3<br>Falling-edge Clock Reference<br>VOUT(x)_PCLK<br>Rising-edge Clock Reference<br>VOUT(x)_PCLK<br>D5<br>VOUT(x)_VSYNC<br>D5<br>VOUT(x)_HSYNC<br>D4<br>VOUT(x)_DATA[23:0] data_1 data_2 data_n<br>D5<br>VOUT(x)_DE<br>DPI_TIMING_01<br>**----- End of picture text -----**<br>


- A. The assertion of data can be programmed to occur on the falling or rising edge of the pixel clock. Refer to _Display Subsystem (DSS)_ section in _Peripherals_ chapter in the device TRM. 

- B. The polarity and pulse width of VOUT(x)_HSYNC and VOUT(x)_VSYNC are programmable, refer to _Display Subsystem (DSS)_ section in _Peripherals_ chapter in the device TRM. 

- C. The VOUT(x)_PCLK frequency is configurable, refer to _Display Subsystem_ section in _Peripherals_ chapter in the device TRM. 

## **Figure 6-40. DSS Switching Characteristics** 

For more information, see _Display Subsystem (DSS) and Peripherals_ section in _Peripherals_ chapter of the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 123 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.6 ECAP** 

Table 6-43, Table 6-44, Figure 6-41, Table 6-45, and Figure 6-42 present timing conditions, timing requirements, and switching characteristics for ECAP. 

**==> picture [172 x 10] intentionally omitted <==**

**----- Start of picture text -----**<br>
Table 6-43. ECAP Timing Conditions<br>**----- End of picture text -----**<br>


|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|---|
|**INPUT CONDITIONS**|||||||
|SRI||Input slew rate||1<br>4||V/ns|
|**OUTPUT CONDITIONS**|||||||
|CL||Output load capacitance||2<br>7||pF|
|**Table 6-44. ECAP Timing Requirements**<br>seeFigure 6-41|||||||
|**NO.**|**PARAMETER**||**DESCRIPTION**||**MIN**<br>**MAX**|**UNIT**|
|CAP1|tw(CAP)||Pulse duration, CAP (asynchronous)||2P(1)+ 2|ns|
|(1)<br>P = MAIN_PLL0_HSDIV6 period in ns.<br>CAP<br>~~CAP1~~<br>EPERIPHERALS_TIMNG_01|||||||



**Figure 6-41. ECAP Timings Requirements** 

**Table 6-45. ECAP Switching Characteristics** 

|**Table 6-45. ECAP Switching Characteristics**|**Table 6-45. ECAP Switching Characteristics**|**Table 6-45. ECAP Switching Characteristics**|**Table 6-45. ECAP Switching Characteristics**|**Table 6-45. ECAP Switching Characteristics**|
|---|---|---|---|---|
|seeFigure 6-42|||||
|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|CAP2|tw(APWM)|Pulse duration, APWMx high/low|2P(1)- 2|ns|
|(1)<br>P = MAIN_PLL0_HSDIV6 period in ns.<br>APWM<br>~~CAP2~~<br>EPERIPHERALS_TIMNG_02|||||



**Figure 6-42. ECAP Switching Characteristics** 

For more information, see _Enhanced Capture (ECAP) Module_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

124 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **6.11.5.7 Emulation and Debug** 

For more details about features and additional description information on the device Trace and JTAG interfaces, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

## _**6.11.5.7.1 Trace**_ 

## **Table 6-46. Trace Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|2<br>5|pF|
|**PCB CONNECTIVITY REQUIREMENTS**||||
|td(Trace Mismatch)|Propagation delay mismatch across all traces|150|ps|



## **Table 6-47. Trace Switching Characteristics** 

## see Figure 6-43 

|**NO.**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**1.8V Mode**||||
|DBTR1|tc(TRC_CLK)<br>Cycle time, TRC_CLK|6.83|ns|
|DBTR2|tw(TRC_CLKH)<br>Pulse width, TRC_CLK high|2.66|ns|
|DBTR3|tw(TRC_CLKL)<br>Pulse width, TRC_CLK low|2.66|ns|
|DBTR4|tosu(TRC_DATAV-<br>TRC_CLK)<br>Output setup time, TRC_DATA valid to TRC_CLK edge|0.85|ns|
|DBTR5|toh(TRC_CLK-TRC_DATAI)<br>Output hold time, TRC_CLK edge to TRC_DATA invalid|0.85|ns|
|DBTR6|tosu(TRC_CTLV-TRC_CLK)<br>Output setup time, TRC_CTL valid to TRC_CLK edge|0.85|ns|
|DBTR7|toh(TRC_CLK-TRC_CTLI)<br>Output hold time, TRC_CLK edge to TRC_CTL invalid|0.85|ns|
|**3.3V Mode**||||
|DBTR1|tc(TRC_CLK)<br>Cycle time, TRC_CLK|8.78|ns|
|DBTR2|tw(TRC_CLKH)<br>Pulse width, TRC_CLK high|3.64|ns|
|DBTR3|tw(TRC_CLKL)<br>Pulse width, TRC_CLK low|3.64|ns|
|DBTR4|tosu(TRC_DATAV-<br>TRC_CLK)<br>Output setup time, TRC_DATA valid to TRC_CLK edge|1.10|ns|
|DBTR5|toh(TRC_CLK-TRC_DATAI)<br>Output hold time, TRC_CLK edge to TRC_DATA invalid|1.10|ns|
|DBTR6|tosu(TRC_CTLV-TRC_CLK)<br>Output setup time, TRC_CTL valid to TRC_CLK edge|1.10|ns|
|DBTR7|toh(TRC_CLK-TRC_CTLI)<br>Output hold time, TRC_CLK edge to TRC_CTL invalid|1.10|ns|



**==> picture [499 x 112] intentionally omitted <==**

**----- Start of picture text -----**<br>
DBTR1<br>DBTR2 DBTR3<br>TRC_CLK<br>(Worst Case 1)<br>(Ideal)<br>(Worst Case 2)<br>DBTR4 DBTR5 DBTR4 DBTR5<br>DBTR6 DBTR7 DBTR6 DBTR7<br>TRC_DATA<br>TRC_CTL<br>**----- End of picture text -----**<br>


SPRSP08_Debug_01 

**Figure 6-43. Trace Switching Characteristics** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 125 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## _**6.11.5.7.2 JTAG**_ 

**Table 6-48. JTAG Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|0.5<br>2.0|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|5<br>15|pF|
|**PCB CONNECTIVITY REQUIREMENTS**||||
|td(Trace Delay)|Propagation delay of each trace|83.5<br>1000(1)|ps|
|td(Trace Mismatch Delay)|Propagation delay mismatch across all traces|100|ps|



- (1) Maximum propagation delay associated with the JTAG signal traces has a significant impact on maximum TCK operating frequency. It may be possible to increase the trace delay beyond this value, but the operating frequency of TCK must be reduced to account for the additional trace delay. 

## **Table 6-49. JTAG Timing Requirements** 

see Figure 6-44 

|**NO.**||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|J1|tc(TCK)<br>Cycle time minimum, TCK|40(1)|ns|
|J2|tw(TCKH)<br>Pulse width minimum, TCK high|0.4P(2)|ns|
|J3|tw(TCKL)<br>Pulse width minimum, TCK low|0.4P(2)|ns|
|J4|tsu(TDI-TCK)<br>Input setup time minimum, TDI valid to TCK high|2|ns|
||tsu(TMS-TCK)<br>Input setup time minimum, TMS valid to TCK high|2|ns|
|J5|th(TCK-TDI)<br>Input hold time minimum, TDI valid from TCK high|2|ns|
||th(TCK-TMS)<br>Input hold time minimum, TMS valid from TCK high|2|ns|



- (1) The maximum TCK operating frequency assumes the following timing requirements and switching characteristis for the attached debugger. The operating frequency of TCK must be reduced to provide appropriate timing margin if the debugger exceeds any of these assumptions. 

   - Minimum TDO setup time of 2ns relative to the rising edge of TCK 

   - TDI and TMS output delay in the range of -13.9ns to 13.9ns relative to the falling edge of TCK 

- (2) P = TCK cycle time in ns 

**Table 6-50. JTAG Switching Characteristics** 

## see Figure 6-44 

|**NO.**|||**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**||**MIN**<br>**MAX**|**MIN**<br>**MAX**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|---|---|---|---|---|---|
|J6|td(TCKL-TDOI)|||Delay time minimum, TCK low to TDO invalid||||||0|ns|
|J7|td(TCKL-TDOV)|||Delay time maximum, TCK low to TDO valid||||||8|ns|
|TDO<br>TCK<br>TDI / TMS|||||J2<br>J1|||||J5||
||||||||J3|||||
|||||||||||||
|||||||||||||
|||||||||||||
||||J4||J5|||J4||J5||
|||||||||||||
|||||||||||||
|||||||||||||
||||||J6|||||||
||||||||J7|||||



**Figure 6-44. JTAG Timing Requirements and Switching Characteristics** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

126 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.8 EPWM** 

Table 6-51, Table 6-52, Figure 6-45, Table 6-53, Figure 6-46, Figure 6-47, and Figure 6-48 present timing conditions, timing requirements, and switching characteristics for EPWM. 

**Table 6-51. EPWM Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|1<br>4|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|2<br>7|pF|



## **Table 6-52. EPWM Timing Requirements** 

**==> picture [64 x 9] intentionally omitted <==**

**----- Start of picture text -----**<br>
see Figure 6-45<br>**----- End of picture text -----**<br>


|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|PWM6|tw(SYNCIN)|Pulse duration, EHRPWM_SYNCI|2P(1)+ 2|ns|
|PWM7|tw(TZ)|Pulse duration, EHRPWM_TZn_IN low|3P(1)+ 2|ns|



(1) P = MAIN_PLL0_HSDIV6 period in ns. 

**==> picture [459 x 119] intentionally omitted <==**

**----- Start of picture text -----**<br>
PWM6<br>EHRPWM_SYNCI<br>PWM7<br>EHRPWM_TZn_IN<br>**----- End of picture text -----**<br>


**==> picture [53 x 5] intentionally omitted <==**

**----- Start of picture text -----**<br>
EPERIPHERALS_TIMNG_07<br>**----- End of picture text -----**<br>


**Figure 6-45. EPWM Timing Requirements** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 127 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 6-53. EPWM Switching Characteristics** 

see Figure 6-46, Figure 6-47, and Figure 6-48 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|PWM1|tw(PWM)|Pulse duration, EHRPWM_A/B high/low|P(1)- 3|ns|
|PWM2|tw(SYNCOUT)|Pulse duration, EHRPWM_SYNCO|P(1)- 3|ns|
|PWM3|td(TZ-PWM)|Delay time, EHRPWM_TZn_IN active to EHRPWM_A/B forced<br>high/low|11|ns|
|PWM4|td(TZ-PWMZ)|Delay time, EHRPWM_TZn_IN active to EHRPWM_A/B Hi-Z|11|ns|
|PWM5|tw(SOC)|Pulse duration, EHRPWM_SOCA/B output|P(1)- 3|ns|



(1) P = MAIN_PLL0_HSDIV6 period in ns. 

**==> picture [458 x 165] intentionally omitted <==**

**----- Start of picture text -----**<br>
PWM1<br>EHRPWM_A/B<br>PWM1<br>PWM2<br>EHRPWM_SYNCO<br>PWM5<br>EHRPWM_SOCA/B<br>**----- End of picture text -----**<br>


**==> picture [54 x 5] intentionally omitted <==**

**----- Start of picture text -----**<br>
EPERIPHERALS_TIMNG_04<br>**----- End of picture text -----**<br>


**Figure 6-46. EHRPWM Switching Characteristics** 

**==> picture [180 x 64] intentionally omitted <==**

**----- Start of picture text -----**<br>
EHRPWM_A/B<br>EHRPWM_TZn_IN<br>**----- End of picture text -----**<br>


**==> picture [158 x 78] intentionally omitted <==**

**----- Start of picture text -----**<br>
PWM3<br>EPERIPHERALS_TIMING_05<br>**----- End of picture text -----**<br>


**Figure 6-47. EHRPWM_TZn_IN to EHRPWM_A/B Forced Switching Characteristics** 

**==> picture [285 x 64] intentionally omitted <==**

**----- Start of picture text -----**<br>
PWM4<br>EHRPWM_A/B<br>EHRPWM_TZn_IN<br>**----- End of picture text -----**<br>


**Figure 6-48. EHRPWM_TZn_IN to EHRPWM_A/B Hi-Z Switching Characteristics** 

For more information, see _Enhanced Pulse Width Modulation (EPWM) Module_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

128 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.9 EQEP** 

Table 6-54, Table 6-55, Figure 6-49, and Table 6-56 present timing conditions, timing requirements, and switching characteristics for EQEP. 

## **Table 6-54. EQEP Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|1<br>4|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|2<br>7|pF|



## **Table 6-55. EQEP Timing Requirements** 

## see Figure 6-49 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|QEP1|tw(QEP)|Pulse duration, QEP_A/B|2P(1)+ 2|ns|
|QEP2|tw(QEPIH)|Pulse duration, QEP_I high|2P(1)+ 2|ns|
|QEP3|tw(QEPIL)|Pulse duration, QEP_I low|2P(1)+ 2|ns|
|QEP4|tw(QEPSH)|Pulse duration, QEP_S high|2P(1)+ 2|ns|
|QEP5|tw(QEPSL)|Pulse duration, QEP_S low|2P(1)+ 2|ns|
|(1)<br>P = MAIN_PLL0_HSDIV6 period in ns.<br>QEP_S<br>~~QEP4~~<br>EPERIPHERALS_TIMNG_03<br>QEP_I<br>~~QEP2~~<br>QEP_A/B<br>~~QEP1~~<br>~~QEP3~~<br>~~QEP5~~|||||



**Figure 6-49. EQEP Timing Requirements** 

**Table 6-56. EQEP Switching Characteristics** 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|QEP6|td(QEP-CNTR)|Delay time, external clock to counter increment|24|ns|



For more information, see _Enhanced Quadrature Encoder Pulse (EQEP) Module_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 129 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.10 GPIO** 

Table 6-57, Table 6-58, and Table 6-59 present timing conditions, timing requirements, and switching characteristics for GPIO. 

The device has two instances of the GPIO module. 

- GPIO0 

- WKUP_GPIO0 

## **Note** 

GPIOn_x is generic name used to describe a GPIO signal, where n represents the specific GPIO module and x represents one of the input/output signals associated with the module. 

For additional description information on the device GPIO, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

**Table 6-57. GPIO Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**BUFFER TYPE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**INPUT CONDITIONS**|||||
|SRI|Input slew rate|LVCMOS<br>(VDD(1)= 1.8V)|0.0018<br>6.6|V/ns|
|||LVCMOS<br>(VDD(1)= 3.3V)|0.0033<br>6.6|V/ns|
|||I2C OD FS<br>(VDD(1)= 1.8V)|0.0018<br>6.6|V/ns|
|||I2C OD FS<br>(VDD(1)= 3.3V)|0.0033<br>0.08|V/ns|
|**OUTPUT CONDITIONS**|||||
|CL|Output load capacitance|LVCMOS|3<br>10|pF|
|||I2C OD FS|3<br>100|pF|



(1) VDD stands for corresponding power supply. For more information on the power supply name and the corresponding ball(s), see POWER column of the _Pin Attributes_ table. 

**Table 6-58. GPIO Timing Requirements** 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|GPIO1|tw(GPIO_IN)|Pulse width, GPIOn_x|2P(1)+ 30|ns|



(1) P = functional clock period in ns. 

**Table 6-59. GPIO Switching Characteristics** 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**BUFFER TYPE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|GPIO2|tw(GPIO_OUT)|Pulse width, GPIOn_x|LVCMOS|0.975P(1)-<br>3.6|ns|
||||I2C OD FS|160|ns|



(1) P = functional clock period in ns. 

For more information, see _General-Purpose Interface (GPIO)_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

130 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.11 GPMC** 

For more details about features and additional description information on the device General-Purpose Memory Controller, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

Table 6-60 presents timing conditions for GPMC. 

**Table 6-60. GPMC Timing Conditions** 

|**PARAMETER**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**INPUT CONDITIONS**|||||
|SRI|Input slew rate||1.65<br>4|V/ns|
|**OUTPUT CONDITIONS**|||||
|CL|Output load capacitance||2<br>20|pF|
|**PCB CONNECTIVITY REQUIREMENTS**|||||
|td(Trace Delay)|Propagation delay of each trace|133MHz Synchronous Mode|140<br>360|ps|
|||All other modes|140<br>720|ps|
|td(Trace Mismatch<br>Delay)|Propagation delay mismatch across all traces||200|ps|



For more information, see _General-Purpose Memory Controller (GPMC)_ section in _Peripherals_ chapter in the device TRM. 

## _**6.11.5.11.1 GPMC and NOR Flash — Synchronous Mode**_ 

Table 6-61 and Table 6-62 present timing requirements and switching characteristics for GPMC and NOR Flash - Synchronous Mode. 

**Table 6-61. GPMC and NOR Flash Timing Requirements — Synchronous Mode** 

see Figure 6-50, Figure 6-51, and Figure 6-54 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|F12|tsu(dV-clkH)|Setup time, GPMC_AD[15:0] valid before GPMC_CLK<br>high|0.92|ns|
|F13|th(clkH-dV)|Hold time, GPMC_AD[15:0] valid after GPMC_CLK high|2.09|ns|
|F21|tsu(waitV-clkH)|Setup time, GPMC_WAIT[j](1) (2)valid before GPMC_CLK<br>high|0.92|ns|
|F22|th(clkH-waitV)|Hold time, GPMC_WAIT[j](1) (2)valid after GPMC_CLK<br>high|2.09|ns|



(1) In GPMC_WAIT[j], j is equal to 0 or 1. 

(2) Wait monitoring support is limited to a WaitMonitoringTime value > 0. For a full description of wait monitoring feature, see _GeneralPurpose Memory Controller (GPMC)_ section in the device TRM. 

**Table 6-62. GPMC and NOR Flash Switching Characteristics – Synchronous Mode** 

see Figure 6-50, Figure 6-51, Figure 6-52, Figure 6-53, and Figure 6-54 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|F0|tc(clk)|Cycle time, GPMC_CLK(16)|7.52|ns|
|F1|tw(clkH)|Typical pulse duration, GPMC_CLK high|0.475P(13)-<br>0.3|ns|
|F1|tw(clkL)|Typical pulse duration, GPMC_CLK low|0.475P(13)-<br>0.3|ns|
|F2|td(clkH-csnV)|Delay time, GPMC_CLK rising edge to GPMC_CSn[i]<br>transition(12)|F(5)- 2.2<br>F(5)+ 3.75|ns|
|F3|td(clkH-CSn[i]V)|Delay time, GPMC_CLK rising edge to GPMC_CSn[i]<br>invalid(12)|D(4)- 2.2<br>D(4)+ 4.5|ns|
|F4|td(aV-clk)|Delay time, GPMC_A[27:1] valid to GPMC_CLK first edge|B(2)- 2.3<br>B(2)+ 4.5|ns|
|F5|td(clkH-aIV)|Delay time, GPMC_CLK rising edge to GPMC_A[27:1] invalid|-2.3<br>4.5|ns|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 131 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 6-62. GPMC and NOR Flash Switching Characteristics – Synchronous Mode (continued)** 

see Figure 6-50, Figure 6-51, Figure 6-52, Figure 6-53, and Figure 6-54 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|F6|td(be[x]nV-clk)|Delay time, GPMC_BE0n_CLE, GPMC_BE1n valid to<br>GPMC_CLK first edge|B(2)- 2.3<br>B(2)+ 1.9|ns|
|F7|td(clkH-be[x]nIV)|Delay time, GPMC_CLK rising edge to GPMC_BE0n_CLE,<br>GPMC_BE1n invalid|D(4)- 2.3<br>D(4)+ 1.9|ns|
|F8|td(clkH-advn)|Delay time, GPMC_CLK rising edge to GPMC_ADVn_ALE<br>transition|G(6)- 2.3<br>G(6)+ 4.5|ns|
|F9|td(clkH-advnIV)|Delay time, GPMC_CLK rising edge to GPMC_ADVn_ALE<br>invalid|D(4)- 2.3<br>D(4)+ 4.5|ns|
|F10|td(clkH-oen)|Delay time, GPMC_CLK rising edge to GPMC_OEn_REn<br>transition|H(7)- 2.3<br>H(7)+ 3.5|ns|
|F11|td(clkH-oenIV)|Delay time, GPMC_CLK rising edge to GPMC_OEn_REn<br>invalid|D(4)- 2.3<br>D(4)+ 3.5|ns|
|F14|td(clkH-wen)|Delay time, GPMC_CLK rising edge to GPMC_WEn transition|I(8)- 2.3<br>I(8)+ 4.5|ns|
|F15|td(clkH-do)|Delay time, GPMC_CLK rising edge to GPMC_AD[15:0]<br>transition(9)|- 2.3<br>+ 2.7|ns|
|F15|td(clkL-do)|Delay time, GPMC_CLK falling edge to GPMC_AD[15:0] data<br>bus transition(10)|- 2.3<br>+ 2.7|ns|
|F15|td(clkL-do).|Delay time, GPMC_CLK falling edge to GPMC_AD[15:0] data<br>bus transition(11)|- 2.3<br>+ 2.7|ns|
|F17|td(clkH-be[x]n)|Delay time, GPMC_CLK rising edge to GPMC_BE0n_CLE,<br>GPMC_BE1n transition(9)|- 2.3<br>+ 1.9|ns|
|F17|td(clkL-be[x]n)|Delay time, GPMC_CLK falling edge to GPMC_BE0n_CLE,<br>GPMC_BE1n transition(10)|- 2.3<br>+ 1.9|ns|
|F17|td(clkL-be[x]n).|Delay time, GPMC_CLK falling edge to GPMC_BE0n_CLE,<br>GPMC_BE1n transition(11)|- 2.3<br>+ 1.9|ns|
|F18|tw(csnV)|Pulse duration, GPMC_CSn[i](12)low|A(1)|ns|
|F19|tw(be[x]nV)|Pulse duration, GPMC_BE0n_CLE, GPMC_BE1n low|C(3)|ns|
|F20|tw(advnV)|Pulse duration, GPMC_ADVn_ALE low|K(14)|ns|



(1) For single read: A = (CSRdOffTime - CSOnTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For single write: A = (CSWrOffTime - CSOnTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For burst read: A = (CSRdOffTime - CSOnTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For burst write: A = (CSWrOffTime - CSOnTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] With n being the page burst access number. 

- (2) Address bus / Byte Enables become valid at start of cycle, GPMC_CLK activation time may be delayed after start of cycle B = ClkActivationTime × GPMC_FCLK[(15)] 

- (3) For single read: C = RdCycleTime × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For single write: C = WrCycleTime × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] 

   - For burst read: C = (RdCycleTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For burst write: C = (WrCycleTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] With n being the page burst access number. 

- (4) For single read: D = (RdCycleTime - RdAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For single write: D = (WrCycleTime - WrAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For burst read: D = (RdCycleTime - RdAccessTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For burst write: D = (WrCycleTime - WrAccessTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] With n being the page burst access number. 

- (5) For CSn falling edge (CS activated): 

   - Case GPMCFCLKDIVIDER = 0: 

      - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] 

   - Case GPMCFCLKDIVIDER = 1: 

      - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and CSOnTime are odd) or (ClkActivationTime and CSOnTime are even) 

      - F = (1 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] otherwise 

   - Case GPMCFCLKDIVIDER = 2: 

      - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] if ((CSOnTime - ClkActivationTime) is a multiple of 3) 

Copyright © 2025 Texas Instruments Incorporated 

132 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

- F = (1 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] if ((CSOnTime - ClkActivationTime - 1) is a multiple of 3) 

- F = (2 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] if ((CSOnTime - ClkActivationTime - 2) is a multiple of 3) 

For CSn rising edge (CS deactivated) in Reading mode: 

- Case GPMCFCLKDIVIDER = 0: 

   - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] 

- Case GPMCFCLKDIVIDER = 1: 

   - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and CSRdOffTime are odd) or (ClkActivationTime and CSRdOffTime are even) 

   - F = (1 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] otherwise 

- Case GPMCFCLKDIVIDER = 2: 

   - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] if ((CSRdOffTime - ClkActivationTime) is a multiple of 3) 

   - F = (1 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] if ((CSRdOffTime - ClkActivationTime - 1) is a multiple of 3) 

   - F = (2 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] if ((CSRdOffTime - ClkActivationTime - 2) is a multiple of 3) 

For CSn rising edge (CS deactivated) in Writing mode: 

   - Case GPMCFCLKDIVIDER = 0: 

      - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] 

   - Case GPMCFCLKDIVIDER = 1: 

      - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and CSWrOffTime are odd) or (ClkActivationTime and CSWrOffTime are even) 

      - F = (1 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] otherwise 

   - Case GPMCFCLKDIVIDER = 2: 

      - F = 0.5 × CSExtraDelay × GPMC_FCLK[(15)] if ((CSWrOffTime - ClkActivationTime) is a multiple of 3) 

      - F = (1 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] if ((CSWrOffTime - ClkActivationTime - 1) is a multiple of 3) 

- F = (2 + 0.5 × CSExtraDelay) × GPMC_FCLK[(15)] if ((CSWrOffTime - ClkActivationTime - 2) is a multiple of 3) 

- (6) For ADV falling edge (ADV activated): 

   - Case GPMCFCLKDIVIDER = 0: 

      - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] 

   - Case GPMCFCLKDIVIDER = 1: 

      - 

      - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and ADVOnTime are odd) or (ClkActivationTime and ADVOnTime are even) 

   - G = (1 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] otherwise 

- Case GPMCFCLKDIVIDER = 2: 

   - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] if ((ADVOnTime - ClkActivationTime) is a multiple of 3) 

   - G = (1 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] if ((ADVOnTime - ClkActivationTime - 1) is a multiple of 3) 

   - G = (2 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] if ((ADVOnTime - ClkActivationTime - 2) is a multiple of 3) 

For ADV rising edge (ADV deactivated) in Reading mode: 

- Case GPMCFCLKDIVIDER = 0: 

   - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] 

- Case GPMCFCLKDIVIDER = 1: 

   - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and ADVRdOffTime are odd) or (ClkActivationTime and ADVRdOffTime are even) 

   - G = (1 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] otherwise 

- Case GPMCFCLKDIVIDER = 2: 

   - 

   - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] if ((ADVRdOffTime - ClkActivationTime) is a multiple of 3) 

- G = (1 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] if ((ADVRdOffTime - ClkActivationTime - 1) is a multiple of 3) 

- G = (2 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] if ((ADVRdOffTime - ClkActivationTime - 2) is a multiple of 3) 

For ADV rising edge (ADV deactivated) in Writing mode: 

- Case GPMCFCLKDIVIDER = 0: 

   - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] 

- Case GPMCFCLKDIVIDER = 1: 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 133 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

      - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and ADVWrOffTime are odd) or (ClkActivationTime and ADVWrOffTime are even) 

      - – G = (1 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] otherwise 

   - Case GPMCFCLKDIVIDER = 2: 

      - G = 0.5 × ADVExtraDelay × GPMC_FCLK[(15)] if ((ADVWrOffTime - ClkActivationTime) is a multiple of 3) 

      - G = (1 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] if ((ADVWrOffTime - ClkActivationTime - 1) is a multiple of 3) 

- G = (2 + 0.5 × ADVExtraDelay) × GPMC_FCLK[(15)] if ((ADVWrOffTime - ClkActivationTime - 2) is a multiple of 3) 

- (7) For OE falling edge (OE activated) and IO DIR rising edge (Data Bus input direction): • Case GPMCFCLKDIVIDER = 0: 

      - H = 0.5 × OEExtraDelay × GPMC_FCLK[(15)] 

   - Case GPMCFCLKDIVIDER = 1: – H = 0.5 × OEExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and OEOnTime are odd) or (ClkActivationTime and OEOnTime are even) 

   - – H = (1 + 0.5 × OEExtraDelay) × GPMC_FCLK[(15)] otherwise 

   - • Case GPMCFCLKDIVIDER = 2: – H = 0.5 × OEExtraDelay × GPMC_FCLK[(15)] if ((OEOnTime - ClkActivationTime) is a multiple of 3) 

      - H = (1 + 0.5 × OEExtraDelay) × GPMC_FCLK[(15)] if ((OEOnTime - ClkActivationTime - 1) is a multiple of 3) 

      - H = (2 + 0.5 × OEExtraDelay) × GPMC_FCLK[(15)] if ((OEOnTime - ClkActivationTime - 2) is a multiple of 3) 

For OE rising edge (OE deactivated): 

   - Case GPMCFCLKDIVIDER = 0: 

      - H = 0.5 × OEExtraDelay × GPMC_FCLK[(15)] 

   - Case GPMCFCLKDIVIDER = 1: 

      - H = 0.5 × OEExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and OEOffTime are odd) or (ClkActivationTime and OEOffTime are even) 

      - H = (1 + 0.5 × OEExtraDelay) × GPMC_FCLK[(15)] otherwise 

   - Case GPMCFCLKDIVIDER = 2: 

      - H = 0.5 × OEExtraDelay × GPMC_FCLK[(15)] if ((OEOffTime - ClkActivationTime) is a multiple of 3) 

      - H = (1 + 0.5 × OEExtraDelay) × GPMC_FCLK[(15)] if ((OEOffTime - ClkActivationTime - 1) is a multiple of 3) 

- H = (2 + 0.5 × OEExtraDelay) × GPMC_FCLK[(15)] if ((OEOffTime - ClkActivationTime - 2) is a multiple of 3) 

- (8) For WE falling edge (WE activated): • Case GPMCFCLKDIVIDER = 0: 

      - I = 0.5 × WEExtraDelay × GPMC_FCLK[(15)] 

   - Case GPMCFCLKDIVIDER = 1: 

      - I = 0.5 × WEExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and WEOnTime are odd) or (ClkActivationTime and WEOnTime are even) 

      - I = (1 + 0.5 × WEExtraDelay) × GPMC_FCLK[(15)] otherwise 

   - Case GPMCFCLKDIVIDER = 2: 

      - 

   - I = 0.5 × WEExtraDelay × GPMC_FCLK[(15)] if ((WEOnTime - ClkActivationTime) is a multiple of 3) 

- I = (1 + 0.5 × WEExtraDelay) × GPMC_FCLK[(15)] if ((WEOnTime - ClkActivationTime - 1) is a multiple of 3) 

- I = (2 + 0.5 × WEExtraDelay) × GPMC_FCLK[(15)] if ((WEOnTime - ClkActivationTime - 2) is a multiple of 3) 

For WE rising edge (WE deactivated): 

   - Case GPMCFCLKDIVIDER = 0: 

      - I = 0.5 × WEExtraDelay × GPMC_FCLK[(13)] 

   - Case GPMCFCLKDIVIDER = 1: 

      - I = 0.5 × WEExtraDelay × GPMC_FCLK[(15)] if (ClkActivationTime and WEOffTime are odd) or (ClkActivationTime and WEOffTime are even) 

      - I = (1 + 0.5 × WEExtraDelay) × GPMC_FCLK[(15)] otherwise 

   - Case GPMCFCLKDIVIDER = 2: 

      - I = 0.5 × WEExtraDelay × GPMC_FCLK[(15)] if ((WEOffTime - ClkActivationTime) is a multiple of 3) 

      - I = (1 + 0.5 × WEExtraDelay) × GPMC_FCLK[(15)] if ((WEOffTime - ClkActivationTime - 1) is a multiple of 3) 

      - I = (2 + 0.5 × WEExtraDelay) × GPMC_FCLK[(15)] if ((WEOffTime - ClkActivationTime - 2) is a multiple of 3) 

- (9) Case CLK DIV 1 mode, first transfer only: Data and byte enables transition on rise edge of GPMC_CLK 

   - Non-multiplexed mode: data transition at start of cycle 

Copyright © 2025 Texas Instruments Incorporated 

134 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

- Multiplexed mode: data transition at WRDATAONADMUXBUS × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] 

- (10) Case CLK DIV 1 mode, all data and byte enables after initial transfer: Data and byte enables transition on fall edge of GPMC_CLK (Half cycle of GPMC_CLK) 

- (11) Case modes other than CLK DIV 1 mode (GPMC_CLK divided down from GPMC_FCLK): All data and byte enables transition on fall edge of GPMC_CLK (Half cycle of GPMC_CLK). ClkActivationTime, GPMCFCLKDIVIDER, RDACCESSTIME/WRACCESSTIME, and PAGEBURSTACCESSTIME configuration must be configured to enforce data and byte enables transition on falling edge of GPMC_CLK (to be latched on rise edge of GPMC_CLK) 

- (12) In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

- (13) P = GPMC_CLK period in ns 

- (14) For read: K = (ADVRdOffTime - ADVOnTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] For write: K = (ADVWrOffTime - ADVOnTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(15)] 

- (15) GPMC_FCLK is general-purpose memory controller internal functional clock period in ns. 

- (16) Related to the GPMC_CLK output clock maximum and minimum frequencies programmable in the GPMC module by setting the GPMC_CONFIG1_i configuration register bit field GPMCFCLKDIVIDER. 

**==> picture [499 x 314] intentionally omitted <==**

**----- Start of picture text -----**<br>
F1<br>F 0 F1<br>GPMC_CLK<br>F2 F 3<br>F 18<br>GPMC_CSn[i]<br>F4<br>GPMC_A[MSB:1] Valid Address<br>F6 F 7<br>F19<br>GPMC_BE0n_CLE<br>F19<br>GPMC_BE1n<br>F6 F8 F8<br>F 20 F 9<br>GPMC_ADVn_ALE<br>F10 F 11<br>GPMC_OEn_REn<br>F13<br>F12<br>GPMC_AD[15:0] D 0<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


GPMC_01 

- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

- B. In GPMC_WAIT[j], j is equal to 0 or 1. 

## **Figure 6-50. GPMC and NOR Flash — Synchronous Single Read (GPMCFCLKDIVIDER = 0)** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 135 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [499 x 274] intentionally omitted <==**

**----- Start of picture text -----**<br>
F1<br>F0 F1<br>GPMC_CLK<br>F2 F 3<br>GPMC_CSn[i]<br>F4<br>GPMCA[MSB:1] Valid Address<br>F6 F7<br>GPMC_BE0n_CLE<br>F7<br>GPMC_BE1n<br>F6 F8 F8 F9<br>GPMC_ADVn_ALE<br>F10 F 11<br>GPMC_OEn_REn<br>F13 F13<br>F12 F12<br>GPMC_AD[15:0] D 0 D 1 D 2 D 3<br>F21 F22<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


**==> picture [20 x 5] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_02<br>**----- End of picture text -----**<br>


- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

- B. In GPMC_WAIT[j], j is equal to 0 or 1. 

## **Figure 6-51. GPMC and NOR Flash — Synchronous Burst Read — 4x16–bit (GPMCFCLKDIVIDER = 0)** 

**==> picture [499 x 283] intentionally omitted <==**

**----- Start of picture text -----**<br>
F1<br>F1 F0<br>GPMC_CLK<br>F2 F3<br>GPMC_CSn[i]<br>F4<br>GPMC_A[MSB:1] Valid Address<br>F17<br>F6 F17 F17<br>GPMC_BE0n_CLE<br>F17<br>F17 F17<br>GPMC_BE1n<br>F6 F8 F8 F9<br>GPMC_ADVn_ALE<br>F14 F14<br>GPMC_WEn<br>F15 F15 F15<br>GPMC_AD[15:0] D 0 D 1 D 2 D 3<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


**==> picture [19 x 5] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_03<br>**----- End of picture text -----**<br>


- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

Copyright © 2025 Texas Instruments Incorporated 

136 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

- B. In GPMC_WAIT[j], j is equal to 0 or 1. 

## **Figure 6-52. GPMC and NOR Flash—Synchronous Burst Write (GPMCFCLKDIVIDER = 0)** 

**==> picture [490 x 273] intentionally omitted <==**

**----- Start of picture text -----**<br>
F1<br>F0 F1<br>GPMC_CLK<br>F2 F 3<br>GPMC_CSn[i]<br>F6 F7<br>GMPC_BE0n_CLE Valid<br>F6 F7<br>GPMC_BE1n Valid<br>F4<br>GPMC_A[27:17] Address (MSB)<br>F12<br>F4 F5 F13 F12<br>GPMC_AD[15:0] Address (LSB) D0 D1 D2 D3<br>F8 F8 F9<br>GPMC_ADVn_ALE<br>F10 F11<br>GPMC_OEn_REn<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


GPMC_04 

- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

- B. In GPMC_WAIT[j], j is equal to 0 or 1. 

## **Figure 6-53. GPMC and Multiplexed NOR Flash — Synchronous Burst Read** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 137 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [487 x 301] intentionally omitted <==**

**----- Start of picture text -----**<br>
F1<br>F1 F0<br>GPMC_CLK<br>F2 F3<br>F 18<br>GPMC_CSn[i]<br>F4<br>GPMC_A[27:17] Address (MSB)<br>F17<br>F6 F17 F17<br>GPMC_BE1n<br>F17<br>F6 F17 F17<br>BPMC_BE0n_CLE<br>F8 F8<br>F20 F 9<br>GPMC_ADVn_ALE<br>F14 F14<br>GPMC_WEn<br>F15 F15 F15<br>GPMC_AD[15:0] Address (LSB) D 0 D 1 D 2 D 3<br>F22 F21<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


**==> picture [20 x 5] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_05<br>**----- End of picture text -----**<br>


- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

- B. In GPMC_WAIT[j], j is equal to 0 or 1. 

## **Figure 6-54. GPMC and Multiplexed NOR Flash — Synchronous Burst Write** 

Copyright © 2025 Texas Instruments Incorporated 

138 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.11.2 GPMC and NOR Flash — Asynchronous Mode**_ 

Table 6-63 and Table 6-64 present timing requirements and switching characteristics for GPMC and NOR Flash — Asynchronous Mode. 

**Table 6-63. GPMC and NOR Flash Timing Requirements – Asynchronous Mode** 

see Figure 6-55, Figure 6-56, Figure 6-57, and Figure 6-59 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|FA5(1)|tacc(d)|Data access time|H(5)|ns|
|FA20(2)|tacc1-pgmode(d)|Page mode successive data access time|P(4)|ns|
|FA21(3)|tacc2-pgmode(d)|Page mode first data access time|H(5)|ns|



(1) The FA5 parameter illustrates the amount of time required to internally sample input data. It is expressed in number of GPMC functional clock cycles. From start of read cycle and after FA5 functional clock cycles, input data is internally sampled by active functional clock edge. FA5 value must be stored inside the AccessTime register bit field. 

(2) The FA20 parameter illustrates amount of time required to internally sample successive input page data. It is expressed in number of GPMC functional clock cycles. After each access to input page data, next input page data is internally sampled by active functional clock edge after FA20 functional clock cycles. The FA20 value must be stored in the PageBurstAccessTime register bit field. 

(3) The FA21 parameter illustrates amount of time required to internally sample first input page data. It is expressed in number of GPMC functional clock cycles. From start of read cycle and after FA21 functional clock cycles, first input page data is internally sampled by active functional clock edge. FA21 value must be stored inside the AccessTime register bit field. 

(4) P = PageBurstAccessTime × (TimeParaGranularity + 1) × GPMC_FCLK[(6)] 

- (5) H = AccessTime × (TimeParaGranularity + 1) × GPMC_FCLK[(6)] 

- (6) GPMC_FCLK is general-purpose memory controller internal functional clock period in ns. 

**Table 6-64. GPMC and NOR Flash Switching Characteristics – Asynchronous Mode** 

see Figure 6-55, Figure 6-56, Figure 6-57, Figure 6-58, Figure 6-59, and Figure 6-60 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|FA0|tw(be[x]nV)|Pulse duration, output lower-byte enable and command latch enable<br>GPMC_BE0n_CLE, output upper-byte enable GPMC_BE1n valid<br>time|N(12)|ns|
|FA1|tw(csnV)|Pulse duration, output chip select GPMC_CSn[i](13)low|A(1)|ns|
|FA3|td(csnV-advnIV)|Delay time, output chip select GPMC_CSn[i](13)valid to output<br>address valid and address latch enable GPMC_ADVn_ALE invalid|B(2)- 2<br>B(2)+ 2|ns|
|FA4|td(csnV-oenIV)|Delay time, output chip select GPMC_CSn[i](13)valid to output<br>enable GPMC_OEn_REn invalid (Single read)|C(3)- 2<br>C(3)+ 2|ns|
|FA9|td(aV-csnV)|Delay time, output address GPMC_A[27:1] valid to output chip<br>select GPMC_CSn[i](13)valid|J(9)- 2<br>J(9)+ 2|ns|
|FA10|td(be[x]nV-csnV)|Delay time, output lower-byte enable and command latch enable<br>GPMC_BE0n_CLE, output upper-byte enable GPMC_BE1n valid to<br>output chip select GPMC_CSn[i](13)valid|J(9)- 2<br>J(9)+ 2|ns|
|FA12|td(csnV-advnV)|Delay time, output chip select GPMC_CSn[i](13)valid to output<br>address valid and address latch enable GPMC_ADVn_ALE valid|K(10)- 2<br>K(10)+ 2|ns|
|FA13|td(csnV-oenV)|Delay time, output chip select GPMC_CSn[i](13)valid to output<br>enable GPMC_OEn_REn valid|L(11)- 2<br>L(11)+ 2|ns|
|FA16|tw(aIV)|Pulse duration output address GPMC_A[26:1] invalid between 2<br>successive read and write accesses|G(7)|ns|
|FA18|td(csnV-oenIV)|Delay time, output chip select GPMC_CSn[i](13)valid to output<br>enable GPMC_OEn_REn invalid (Burst read)|I(8)- 2<br>I(8)+ 2|ns|
|FA20|tw(aV)|Pulse duration, output address GPMC_A[27:1] valid - 2nd, 3rd, and<br>4th accesses|D(4)|ns|
|FA25|td(csnV-wenV)|Delay time, output chip select GPMC_CSn[i](13)valid to output write<br>enable GPMC_WEn valid|E(5)- 2<br>E(5)+ 2|ns|
|FA27|td(csnV-wenIV)|Delay time, output chip select GPMC_CSn[i](13)valid to output write<br>enable GPMC_WEn invalid|F(6)- 2<br>F(6)+ 2|ns|
|FA28|td(wenV-dV)|Delay time, output write enable GPMC_WEn valid to output data<br>GPMC_AD[15:0] valid|2|ns|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 139 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 6-64. GPMC and NOR Flash Switching Characteristics – Asynchronous Mode (continued)** 

see Figure 6-55, Figure 6-56, Figure 6-57, Figure 6-58, Figure 6-59, and Figure 6-60 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|FA29|td(dV-csnV)|Delay time, output data GPMC_AD[15:0] valid to output chip select<br>GPMC_CSn[i](13)valid|J(9)- 2<br>J(9)+ 2|ns|
|FA37|td(oenV-aIV)|Delay time, output enable GPMC_OEn_REn valid to output address<br>GPMC_AD[15:0] phase end|2|ns|



(1) For single read: A = (CSRdOffTime - CSOnTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] For single write: A = (CSWrOffTime - CSOnTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] For burst read: A = (CSRdOffTime - CSOnTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] For burst write: A = (CSWrOffTime - CSOnTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] with n being the page burst access number (2) For reading: B = ((ADVRdOffTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (ADVExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] For writing: B = ((ADVWrOffTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (ADVExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] 

(3) C = ((OEOffTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (OEExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] (4) D = PageBurstAccessTime × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] (5) E = ((WEOnTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (WEExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] (6) F = ((WEOffTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (WEExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] (7) G = Cycle2CycleDelay × GPMC_FCLK[(14)] (8) I = ((OEOffTime + (n - 1) × PageBurstAccessTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (OEExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] (9) J = (CSOnTime × (TimeParaGranularity + 1) + 0.5 × CSExtraDelay) × GPMC_FCLK[(14)] (10) K = ((ADVOnTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (ADVExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] (11) L = ((OEOnTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (OEExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] (12) For single read: N = RdCycleTime × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] For single write: N = WrCycleTime × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] For burst read: N = (RdCycleTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] For burst write: N = (WrCycleTime + (n - 1) × PageBurstAccessTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] (13) In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

(14) GPMC_FCLK is general-purpose memory controller internal functional clock period in ns. 

**==> picture [499 x 293] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GPMC_CLK<br>FA5<br>FA1<br>GPMC_CSn[i]<br>FA9<br>GPMC_A[MSB:1] Valid Address<br>FA0<br>FA10<br>GPMC_BE0n_CLE Valid<br>FA0<br>GPMC_BE1n Valid<br>FA10<br>FA3<br>FA12<br>GPMC_ADVn_ALE<br>FA4<br>FA13<br>GPMC_OEn_REn<br>GPMC_AD[15:0] Data IN 0 Data IN 0<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. In GPMC_WAIT[j], jis equal to 0 or 1. 

GPMC_06 

Copyright © 2025 Texas Instruments Incorporated 

140 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

- B. FA5 parameter illustrates amount of time required to internally sample input data. It is expressed in number of GPMC functional clock cycles. From start of read cycle and after FA5 functional clock cycles, input data will be internally sampled by active functional clock edge. FA5 value must be stored inside AccessTime register bits field. 

- C. GPMC_FCLK is an internal clock (GPMC functional clock) not provided externally. 

## **Figure 6-55. GPMC and NOR Flash — Asynchronous Read — Single Word** 

**==> picture [499 x 305] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GPMC_CLK<br>FA5 FA5<br>FA1 FA1<br>GPMC_CSn[i]<br>FA16<br>FA9 FA9<br>GPMC_A[MSB:1] Address 0 Address 1<br>FA0 FA0<br>FA10 FA10<br>GPMC_BE0n_CLE Valid Valid<br>FA0 FA0<br>GPMC_BE1n Valid Valid<br>FA10 FA10<br>FA3 FA3<br>FA12 FA12<br>GPMC_ADCn_ALE<br>FA4 FA4<br>FA13 FA13<br>GPMC_OEn_REn<br>GPMC_AD[15:0] Data Upper<br>GPMC_WAIT[j]<br>GPMC_07<br>**----- End of picture text -----**<br>


- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. In GPMC_WAIT[j], j is equal to 0 or 1. 

- B. FA5 parameter illustrates amount of time required to internally sample input data. It is expressed in number of GPMC functional clock cycles. From start of read cycle and after FA5 functional clock cycles, input data will be internally sampled by active functional clock edge. FA5 value must be stored inside AccessTime register bits field. 

- C. GPMC_FCLK is an internal clock (GPMC functional clock) not provided externally. 

## **Figure 6-56. GPMC and NOR Flash — Asynchronous Read — 32–Bit** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 141 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [499 x 297] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GPMC_CLK<br>FA21 FA20 FA20 FA20<br>FA1<br>GPMC_CSn[i]<br>FA9<br>GPMC_A[MSB:1] Add0 Add1 Add2 Add3 Add4<br>FA0<br>FA10<br>GPMC_BE0n_CLE<br>FA0<br>FA10<br>GPMC_BE1n<br>FA12<br>GPMC_ADVn_ALE<br>FA18<br>FA13<br>GPMC_OEn_REn<br>GPMC_AD[15:0] D0 D1 D2 D3 D3<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


GPMC_08 

- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. In GPMC_WAIT[j], j is equal to 0 or 1. 

- B. FA21 parameter illustrates amount of time required to internally sample first input page data. It is expressed in number of GPMC functional clock cycles. From start of read cycle and after FA21 functional clock cycles, first input page data will be internally sampled by active functional clock edge. FA21 calculation must be stored inside AccessTime register bits field. 

- C. FA20 parameter illustrates amount of time required to internally sample successive input page data. It is expressed in number of GPMC functional clock cycles. After each access to input page data, next input page data will be internally sampled by active functional clock edge after FA20 functional clock cycles. FA20 is also the duration of address phases for successive input page data (excluding first input page data). FA20 value must be stored in PageBurstAccessTime register bits field. 

- D. GPMC_FCLK is an internal clock (GPMC functional clock) not provided externally. 

## **Figure 6-57. GPMC and NOR Flash — Asynchronous Read — Page Mode 4x16–Bit** 

Copyright © 2025 Texas Instruments Incorporated 

142 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

**==> picture [499 x 288] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GPMC_CLK<br>FA1<br>GPMC_CSn[i]<br>FA9<br>GPMC_A[MSB:1] Valid Address<br>FA0<br>FA10<br>GPMC_BE0n_CLE<br>FA0<br>FA10<br>GPMC_BE1n<br>FA3<br>FA12<br>GPMC_ADVn_ALE<br>FA27<br>FA25<br>GPMC_WEn<br>FA29<br>GPMC_AD[15:0] Data OUT<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


GPMC_09 

- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. In GPMC_WAIT[j], j is equal to 0 or 1. 

## **Figure 6-58. GPMC and NOR Flash — Asynchronous Write — Single Word** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 143 

Product Folder Links: _AM62L_ 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **AM62L** 

**==> picture [499 x 315] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GPMC_CLK<br>FA1<br>FA5<br>GPMC_CSn[i]<br>FA9<br>GPMC_A[27:17] Address (MSB)<br>FA0<br>FA10<br>GPMC_BE0n_CLE Valid<br>FA0<br>FA10<br>GPMC_BE1n Valid<br>FA3<br>FA12<br>GPMC_ADVn_ALE<br>FA4<br>FA13<br>GPMC_OEn_REn<br>FA29 FA37<br>GPMC_AD[15:0] Address (LSB) Data IN Data IN<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


**==> picture [20 x 5] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_10<br>**----- End of picture text -----**<br>


- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. In GPMC_WAIT[j], j is equal to 0 or 1. 

- B. FA5 parameter illustrates amount of time required to internally sample input data. It is expressed in number of GPMC functional clock cycles. From start of read cycle and after FA5 functional clock cycles, input data will be internally sampled by active functional clock edge. FA5 value must be stored inside AccessTime register bits field. 

- C. GPMC_FCLK is an internal clock (GPMC functional clock) not provided externally. 

## **Figure 6-59. GPMC and Multiplexed NOR Flash — Asynchronous Read — Single Word** 

Copyright © 2025 Texas Instruments Incorporated 

144 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

**==> picture [497 x 298] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GPMC_CLK<br>FA1<br>GPMC_CSn[i]<br>FA9<br>GPMC_A[27:17] Address (MSB)<br>FA0<br>FA10<br>GPMC_BE0n_CLE<br>FA0<br>FA10<br>GPMC_BE1n<br>FA3<br>FA12<br>GPMC_ADVn_ALE<br>FA27<br>FA25<br>GPMC_WEn<br>FA29 FA28<br>GPMC_AD[15:0] Valid Address (LSB) Data OUT<br>GPMC_WAIT[j]<br>**----- End of picture text -----**<br>


**==> picture [20 x 5] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_11<br>**----- End of picture text -----**<br>


- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. In GPMC_WAIT[j], j is equal to 0 or 1. 

## **Figure 6-60. GPMC and Multiplexed NOR Flash — Asynchronous Write — Single Word** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 145 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.11.3 GPMC and NAND Flash — Asynchronous Mode**_ 

Table 6-65 and Table 6-66 present timing requirements and switching characteristics for GPMC and NAND Flash — Asynchronous Mode. 

## **Table 6-65. GPMC and NAND Flash Timing Requirements – Asynchronous Mode** 

see Figure 6-63 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|GNF12(1)|tacc(d)|Access time, input data GPMC_AD[15:0]|J(2)|ns|



(1) The GNF12 parameter illustrates the amount of time required to internally sample input data. It is expressed in number of GPMC functional clock cycles. From start of the read cycle and after GNF12 functional clock cycles, input data is internally sampled by the active functional clock edge. The GNF12 value must be stored inside AccessTime register bit field. 

- (2) J = AccessTime × (TimeParaGranularity + 1) × GPMC_FCLK[(3)] 

- (3) GPMC_FCLK is general-purpose memory controller internal functional clock period in ns. 

## **Table 6-66. GPMC and NAND Flash Switching Characteristics – Asynchronous Mode** 

see Figure 6-61, Figure 6-62, Figure 6-63 and Figure 6-64 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|GNF0|tw(wenV)|Pulse duration, output write enable GPMC_WEn valid|A(1)|ns|
|GNF1|td(csnV-wenV)|Delay time, output chip select GPMC_CSn[i](13)valid to output write<br>enable GPMC_WEn valid|B(2)- 2<br>B(2)+ 2|ns|
|GNF2|tw(cleH-wenV)|Delay time, output lower-byte enable and command latch enable<br>GPMC_BE0n_CLE high to output write enable GPMC_WEn valid|C(3)- 2<br>C(3)+ 2|ns|
|GNF3|tw(wenV-dV)|Delay time, output data GPMC_AD[15:0] valid to output write<br>enable GPMC_WEn valid|D(4)- 2<br>D(4)+ 2|ns|
|GNF4|tw(wenIV-dIV)|Delay time, output write enable GPMC_WEn invalid to output data<br>GPMC_AD[15:0] invalid|E(5)- 2<br>E(5)+ 2|ns|
|GNF5|tw(wenIV-cleIV)|Delay time, output write enable GPMC_WEn invalid to output<br>lower-byte enable and command latch enable GPMC_BE0n_CLE<br>invalid|F(6)- 2<br>F(6)+ 2|ns|
|GNF6|tw(wenIV-CSn[i]V)|Delay time, output write enable GPMC_WEn invalid to output chip<br>select GPMC_CSn[i](13)invalid|G(7)- 2<br>G(7)+ 2|ns|
|GNF7|tw(aleH-wenV)|Delay time, output address valid and address latch enable<br>GPMC_ADVn_ALE high to output write enable GPMC_WEn valid|C(3)- 2<br>C(3)+ 2|ns|
|GNF8|tw(wenIV-aleIV)|Delay time, output write enable GPMC_WEn invalid to output<br>address valid and address latch enable GPMC_ADVn_ALE invalid|F(6)- 2<br>F(6)+ 2|ns|
|GNF9|tc(wen)|Cycle time, write|H(8)|ns|
|GNF10|td(csnV-oenV)|Delay time, output chip select GPMC_CSn[i](13)valid to output<br>enable GPMC_OEn_REn valid|I(9)- 2<br>I(9)+ 2|ns|
|GNF13|tw(oenV)|Pulse duration, output enable GPMC_OEn_REn valid|K(10)|ns|
|GNF14|tc(oen)|Cycle time, read|L(11)|ns|
|GNF15|tw(oenIV-CSn[i]V)|Delay time, output enable GPMC_OEn_REn invalid to output chip<br>select GPMC_CSn[i](13)invalid|M(12)- 2<br>M(12)+ 2|ns|



- (1) A = (WEOffTime - WEOnTime) × (TimeParaGranularity + 1) × GPMC_FCLK[(14)] 

- (2) B = ((WEOnTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (WEExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] 

- (3) C = ((WEOnTime - ADVOnTime) × (TimeParaGranularity + 1) + 0.5 × (WEExtraDelay - ADVExtraDelay)) × GPMC_FCLK[(14)] Note: For DeviceType: NAND 

   - During Command Latch Cycle: CLE signal is controlled by the ADVOnTime and ADVWrOffTime timing parameters 

   - During Address Latch Cycle: ALE signal is controlled by the ADVOnTime and ADVWrOffTime timing parameters. 

- (4) D = (WEOnTime × (TimeParaGranularity + 1) + 0.5 × WEExtraDelay) × GPMC_FCLK[(14)] 

- (5) E = ((WrCycleTime - WEOffTime) × (TimeParaGranularity + 1) - 0.5 × WEExtraDelay) × GPMC_FCLK[(14)] 

- (6) F = ((ADVWrOffTime - WEOffTime) × (TimeParaGranularity + 1) + 0.5 × (ADVExtraDelay - WEExtraDelay)) × GPMC_FCLK[(14)] Note: For DeviceType: NAND 

   - During Command Latch Cycle: CLE signal is controlled by the ADVOnTime and ADVWrOffTime timing parameters 

   - During Address Latch Cycle: ALE signal is controlled by the ADVOnTime and ADVWrOffTime timing parameters. 

- (7) G = ((CSWrOffTime - WEOffTime) × (TimeParaGranularity + 1) + 0.5 × (CSExtraDelay - WEExtraDelay)) × GPMC_FCLK[(14)] (8) H = WrCycleTime × (1 + TimeParaGranularity) × GPMC_FCLK[(14)] 

Copyright © 2025 Texas Instruments Incorporated 

146 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

- (9) I = ((OEOnTime - CSOnTime) × (TimeParaGranularity + 1) + 0.5 × (OEExtraDelay - CSExtraDelay)) × GPMC_FCLK[(14)] (10) K = (OEOffTime - OEOnTime) × (1 + TimeParaGranularity) × GPMC_FCLK[(14)] 

- (11) L = RdCycleTime × (1 + TimeParaGranularity) × GPMC_FCLK[(14)] 

- (12) M = ((CSRdOffTime - OEOffTime) × (TimeParaGranularity + 1) + 0.5 × (CSExtraDelay - OEExtraDelay)) × GPMC_FCLK[(14)] (13) In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

- (14) GPMC_FCLK is general-purpose memory controller internal functional clock period in ns. 

**==> picture [491 x 181] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GNF1 GNF6<br>GPMC_CSn[i]<br>GNF2 GNF5<br>GPMC_BE0n_CLE<br>GPMC_ADCn_ALE<br>GPMC_OEn_REn<br>GNF0<br>GPMC_WEn<br>GNF3 GNF4<br>GPMC_AD[15:0] Command<br>**----- End of picture text -----**<br>


**==> picture [20 x 5] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_12<br>**----- End of picture text -----**<br>


- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

**Figure 6-61. GPMC and NAND Flash — Command Latch Cycle** 

**==> picture [498 x 216] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GNF1 GNF6<br>GPMC_CSn[i]<br>GPMC_BE0n_CLE<br>GNF7 GNF8<br>GPMC_ADVn_ALE<br>GPMC_OEn_REn<br>GNF9<br>GNF0<br>GPMC_WEn<br>GNF3 GNF4<br>GPMC_AD[15:0] Address<br>**----- End of picture text -----**<br>


**==> picture [99 x 39] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_13<br>**----- End of picture text -----**<br>


- A. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

**Figure 6-62. GPMC and NAND Flash — Address Latch Cycle** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 147 

Product Folder Links: _AM62L_ 

**www.ti.com** 

## **AM62L** 

## SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**==> picture [491 x 173] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GNF12<br>GNF10 GNF15<br>GPMC_CSn[i]<br>GPMC_BE0n_CLE<br>GPMC_ADVn_ALE<br>GNF14<br>GNF13<br>GPMC_OEn_REn<br>GPMC_AD[15:0] DATA<br>GPMC_WAIT[j]<br>GPMC_14<br>**----- End of picture text -----**<br>


- A. GNF12 parameter illustrates amount of time required to internally sample input data. It is expressed in number of GPMC functional clock cycles. From start of read cycle and after GNF12 functional clock cycles, input data will be internally sampled by active functional clock edge. GNF12 value must be stored inside AccessTime register bits field. 

- B. GPMC_FCLK is an internal clock (GPMC functional clock) not provided externally. 

- C. In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. In GPMC_WAIT[j], j is equal to 0 or 1. 

**Figure 6-63. GPMC and NAND Flash — Data Read Cycle** 

**==> picture [478 x 207] intentionally omitted <==**

**----- Start of picture text -----**<br>
GPMC_FCLK<br>GNF1 GNF6<br>GPMC_CSn[i]<br>GPMC_BE0n_CLE<br>GPMC_ADVn_ALE<br>GPMC_OEn_REn<br>GNF9<br>GNF0<br>GPMC_WEn<br>GNF3 GNF4<br>GPMC_AD[15:0] DATA<br>GPMC_15<br>**----- End of picture text -----**<br>


- A. `In GPMC_CSn[i], i is equal to 0, 1, 2 or 3. 

**Figure 6-64. GPMC and NAND Flash — Data Write Cycle** 

Copyright © 2025 Texas Instruments Incorporated 

148 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.12 I2C** 

The device contains five multicontroller Inter-Integrated Circuit (I2C) controllers. Each I2C controller was designed to be compliant to the I[2] C-bus specification version 2.1. However, the device IOs are not fully compliant to the I2C electrical specification. The speeds supported and exceptions are described per IO buffer type. See the BUFFER TYPE column of the _Pin Attributes_ table to determine which IO buffer type is associated with a specific I2C instance. 

- **LVCMOS, 1P8-LVCMOS, or SDIO** 

   - Speeds: 

      - Standard-mode (up to 100Kbits/s) 

         - 1.8V 

         - 3.3V (not supported by 1P8-LVCMOS buffer type) 

      - Fast-mode (up to 400Kbits/s) 

         - 1.8V 

         - 3.3V (not supported by 1P8-LVCMOS buffer type) 

   - Exceptions: 

      - The IOs associated with these ports are not compliant to the fall time requirements defined in the I2C specification because they are implemented with higher performance LVCMOS push-pull IOs that were designed to support other signal functions that could not be implemented with I2C compatible IOs. The LVCMOS IOs being used on these ports are connected such they emulate open-drain outputs. This emulation is achieved by forcing a constant low output and disabling the output buffer to enter the Hi-Z state. 

      - The I2C specification defines a maximum input voltage VIH of (VDDmax + 0.5V), which exceeds the absolute maximum ratings for the device IOs. The system must be designed to ensure the I2C signals never exceed the limits defined in the _Absolute Maximum Ratings_ section of this datasheet. 

- **I2C OD FS** 

   - 

   - Speeds: 

   - Standard-mode (up to 100Kbits/s) 

      - 1.8V 

      - 3.3V 

   - Fast-mode (up to 400Kbits/s) 

      - 1.8V 

      - 3.3V 

   - Hs-mode (up to 3.4Mbits/s) 

      - 1.8V 

- Exceptions: 

   - The IOs associated with these ports were not design to support Hs-mode while operating at 3.3V. So Hs-mode is limited to 1.8V operation. 

   - The rise and fall times of the I2C signals connected to these ports must not exceed a slew rate of 0.08V/ns (or 8E+7V/s). This limit is more restrictive than the minimum fall time limits defined in the I2C specification. Therefore, it may be necessary to add additional capacitance to the I2C signals to slow the rise and fall times such that they do not exceed a slew rate of 0.08V/ns. 

   - The I2C specification defines a maximum input voltage VIH of (VDDmax + 0.5V), which exceeds the absolute maximum ratings for the device IOs. The system must be designed to ensure the I2C signals never exceed the limits defined in the _Absolute Maximum Ratings_ section of this datasheet. 

## **Note** 

I2C2 and I2C3 have one or more signals which can be multiplexed to more than one pin. Timing is only valid for specific pin combinations known as IOSETs. Valid pin combinations or IOSETs for this interface are defined in the **SysConfig-PinMux Tool** . 

Refer to the Philips I2C-bus specification version 2.1 for timing details. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 149 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

For more details about features and additional description information on the device Inter-Integrated Circuit, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

Copyright © 2025 Texas Instruments Incorporated 

150 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.13 MCAN** 

## **Note** 

MCAN1 and MCAN2 have one or more signals which can be multiplexed to more than one pin. Timing requirements and switching characteristics defined in this section are only valid for specific pin combinations known as IOSETs. Valid pin combinations or IOSETs for this interface are defined in the **SysConfig-PinMux Tool** . 

Table 6-67 and Table 6-68 presents timing conditions and switching characteristics for MCAN. 

For more details about features and additional description information on the device Controller Area Network Interface, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

## **Note** 

The device has multiple MCAN modules. MCANn is a generic prefix applied to MCAN signal names, where n represents the specific MCAN module. 

**Table 6-67. MCAN Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|2<br>15|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|5<br>20|pF|



**Table 6-68. MCAN Switching Characteristics** 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|MCAN1|td(MCAN_TX)|Delay time, transmit shift register to MCANn_TX|10|ns|
|MCAN2|td(MCAN_RX)|Delay time, MCANn_RX to receive shift register|10|ns|



For more information, see _Modular Controller Area Network (MCAN)_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 151 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.14 MCASP** 

## **Note** 

MCASP1 and MCASP2 have one or more signals which can be multiplexed to more than one pin. Timing requirements and switching characteristics defined in this section are only valid for specific pin combinations known as IOSETs. Valid pin combinations or IOSETs for this interface are defined in the **SysConfig-PinMux Tool** . 

Table 6-69, Table 6-70, Figure 6-65, Table 6-71, and Figure 6-66 present timing conditions, timing requirements, and switching characteristics for MCASP. 

**Table 6-69. MCASP Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|0.7<br>5|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|1<br>10|pF|
|**PCB CONNECTIVITY REQUIREMENTS**||||
|td(Trace Delay)|Propagation delay of each trace|100<br>1100|ps|
|td(Trace Mismatch Delay)|Propagation delay mismatch across all traces|100|ps|



## **Table 6-70. MCASP Timing Requirements** 

see Figure 6-65 

|**NO.**|||**MODE**(1)|**MIN**<br>**MAX **|**UNIT**|
|---|---|---|---|---|---|
|ASP1|tc(AHCLKRX)|Cycle time, MCASP[x]_AHCLKR/X(4)||20|ns|
|ASP2|tw(AHCLKRX)|Pulse duration, MCASP[x]_AHCLKR/X(4)high or low||0.5P(2)-<br>1.53|ns|
|ASP3|tc(ACLKRX)|Cycle time, MCASP[x]_ACLKR/X(4)||20|ns|
|ASP4|tw(ACLKRX)|Pulse duration, MCASP[x]_ACLKR/X(4)high or low||0.5R(3)-<br>1.53|ns|
|ASP5|tsu(AFSRX-ACLKRX)|Setup time, MCASP[x]_AFSR/X(4)input valid before<br>MCASP[x]_ACLKR/X(4)|ACLKR/X int|9.29|ns|
||||ACLKR/X ext in/out|4||
|ASP6|th(ACLKRX-AFSRX)|Hold time, MCASP[x]_AFSR/X(4)input valid after<br>MCASP[x]_ACLKR/X(4)|ACLKR/X int|-1|ns|
||||ACLKR/X ext in/out|1.6||
|ASP7|tsu(AXR-ACLKRX)|Setup time, MCASP[x]_AXR(4)input valid before<br>MCASP[x]_ACLKR/X(4)|ACLKR/X int|9.29|ns|
||||ACLKR/X ext in/out|4||
|ASP8|th(ACLKRX-AXR)|Hold time, MCASP[x]_AXR(4)input valid after<br>MCASP[x]_ACLKR/X(4)|ACLKR/X int|-1|ns|
||||ACLKR/X ext in/out|1.6||



(1) ACLKR internal: ACLKRCTL.CLKRM=1, PDIR.ACLKR = 1 ACLKR external input: ACLKRCTL.CLKRM=0, PDIR.ACLKR=0 ACLKR external output: ACLKRCTL.CLKRM=0, PDIR.ACLKR=1 ACLKX internal: ACLKXCTL.CLKXM=1, PDIR.ACLKX = 1 ACLKX external input: ACLKXCTL.CLKXM=0, PDIR.ACLKX=0 ACLKX external output: ACLKXCTL.CLKXM=0, PDIR.ACLKX=1 

(2) P = AHCLKR/X period in ns. For details on AHCLKR/X clock source options, see the McASP Clocks table in the Multichannel Audio Serial Port (MCASP) section of the Module Integration chapter found in the Technical Reference Manual. 

(3) R = ACLKR/X period in ns. 

- (4) x in MCASP[x]_* is 0, 1 or 2 

Copyright © 2025 Texas Instruments Incorporated 

152 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

ASP2 ASP1 ASP2 MCASP[x]_AHCLKR/X (Falling Edge Priority) MCASP[x]_AHCLKR/X (Rising Edge Polarity) 

**==> picture [244 x 52] intentionally omitted <==**

ASP4 ASP3 ASP4 (A) MCASP[x]_ACLKR/X (CLKRP = CLKXP = 0) (B) MCASP[x]_ACLKR/X (CLKRP = CLKXP = 1) ASP6 ASP5 MCASP[x]_AFSR/X (Bit Width, 0 Bit Delay) MCASP[x]_AFSR/X (Bit Width, 1 Bit Delay) MCASP[x]_AFSR/X (Bit Width, 2 Bit Delay) MCASP[x]_AFSR/X (Slot Width, 0 Bit Delay) MCASP[x]_AFSR/X (Slot Width, 1 Bit Delay) MCASP[x]_AFSR/X (Slot Width, 2 Bit Delay) ASP8 ASP7 MCASP[x]_AXR[x] (Data In/Receive) A0 A1 A30 A31 B0 B1 B30 B31 C0 C1 C2 C3 C31 

- A. For CLKRP = CLKXP = 0, the MCASP transmitter is configured for rising edge (to shift data out) and the MCASP receiver is configured for falling edge (to shift data in). 

- B. For CLKRP = CLKXP = 1, the MCASP transmitter is configured for falling edge (to shift data out) and the MCASP receiver is configured for rising edge (to shift data in). 

## **Figure 6-65. MCASP Timing Requirements** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 153 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 6-71. MCASP Switching Characteristics** 

|seeFigure 6-66|seeFigure 6-66|seeFigure 6-66||||
|---|---|---|---|---|---|
|**NO.**|**PARAMETER**|**DESCRIPTION**|**MODE**(1)|**MIN**<br>**MAX**|**UNIT**|
|ASP9|tc(AHCLKRX)|Cycle time, MCASP[x]_AHCLKR/X(4)||20|ns|
|ASP10|tw(AHCLKRX)|Pulse duration, MCASP[x]_AHCLKR/X(4)high or low||0.5P(2)- 2|ns|
|ASP11|tc(ACLKRX)|Cycle time, MCASP[x]_ACLKR/X(4)||20|ns|
|ASP12|tw(ACLKRX)|Pulse duration, MCASP[x]_ACLKR/X(4)high or low||0.5R(3)- 2|ns|
|ASP13|td(ACLKRX-AFSRX)|Delay time, MCASP[x]_ACLKR/X(4)transmit edge to<br>MCASP[x]_AFSR/X(4)output valid|ACLKR/X int|-1<br>7.25|ns|
||||ACLKR/X ext in/out|-15.29<br>12.84||
|ASP14|td(ACLKX-AXR)|Delay time, MCASP[x]_ACLKX(4)transmit edge to<br>MCASP[x]_AXR(4)output valid|ACLKR/X int|-1<br>7.25|ns|
||||ACLKR/X ext in/out|-15.29<br>12.84||
|ASP15|tdis(ACLKX-AXR)|Disable time, MCASP[x]_ACLKX(4)transmit edge to<br>MCASP[x]_AXR(4)output high impedance|ACLKR/X int|-1<br>7.25|ns|
||||ACLKR/X ext in/out|-14.9<br>14||



(1) ACLKR internal: ACLKRCTL.CLKRM=1, PDIR.ACLKR = 1 ACLKR external input: ACLKRCTL.CLKRM=0, PDIR.ACLKR=0 ACLKR external output: ACLKRCTL.CLKRM=0, PDIR.ACLKR=1 ACLKX internal: ACLKXCTL.CLKXM=1, PDIR.ACLKX = 1 ACLKX external input: ACLKXCTL.CLKXM=0, PDIR.ACLKX=0 ACLKX external output: ACLKXCTL.CLKXM=0, PDIR.ACLKX=1 

(2) P = AHCLKR/X period in ns. For details on AHCLKR/X clock source options, see the McASP Clocks table in the Multichannel Audio Serial Port (MCASP) section of the Module Integration chapter found in the Technical Reference Manual. 

(3) R = ACLKR/X period in ns. 

(4) x in MCASP[x]_* is 0, 1 or 2 

Copyright © 2025 Texas Instruments Incorporated 

154 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

ASP10 ASP10 ASP9 MCASP[x]_AHCLKR/X (Falling Edge Priority) MCASP[x]_AHCLKR/X (Rising Edge Polarity) 

**==> picture [497 x 388] intentionally omitted <==**

**----- Start of picture text -----**<br>
MCASP[x]_AHCLKR/X (Falling Edge Priority)<br>MCASP[x]_AHCLKR/X (Rising Edge Polarity)<br>ASP12<br>ASP11<br>ASP12<br>(A)<br>MCASP[x]_ACLKR/X (CLKRP = CLKXP = 1)<br>(B)<br>MCASP[x]_ACLKR/X (CLKRP = CLKXP = 0)<br>ASP13 ASP13<br>ASP13 ASP13<br>MCASP[x]_AFSR/X (Bit Width, 0 Bit Delay)<br>MCASP[x]_AFSR/X (Bit Width, 1 Bit Delay)<br>MCASP[x]_AFSR/X (Bit Width, 2 Bit Delay)<br>ASP13 ASP13 ASP13<br>MCASP[x]_AFSR/X (Slot Width, 0 Bit Delay)<br>MCASP[x]_AFSR/X (Slot Width, 1 Bit Delay)<br>MCASP[x]_AFSR/X (Slot Width, 2 Bit Delay) ASP14<br>ASP15<br>MCASP[x]_AXR[x] (Data Out/Transmit)<br>A0 A1 A30 A31 B0 B1 B30 B31 C0 C1 C2 C3 C31<br>**----- End of picture text -----**<br>


- A. For CLKRP = CLKXP = 1, the MCASP transmitter is configured for falling edge (to shift data out) and the MCASP receiver is configured for rising edge (to shift data in). 

- B. For CLKRP = CLKXP = 0, the MCASP transmitter is configured for rising edge (to shift data out) and the MCASP receiver is configured for falling edge (to shift data in). 

## **Figure 6-66. MCASP Switching Characteristics** 

For more information, see _Multichannel Audio Serial Port (MCASP)_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 155 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.15 MCSPI** 

## **Note** 

MCSPI1, MCSPI2, and MCSPI3 have one or more signals which can be multiplexed to more than one pin. Timing requirements and switching characteristics defined in this section are only valid for specific pin combinations known as IOSETs. Valid pin combinations or IOSETs for this interface are defined in the **SysConfig-PinMux Tool** . 

For more details about features and additional description information on the device Serial Port Interface, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

Table 6-72 presents timing conditions for MCSPI. 

**Table 6-72. MCSPI Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|2<br>8.5|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|6<br>12|pF|



For more information, see _Multichannel Serial Peripheral Interface (MCSPI)_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

156 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.15.1 MCSPI — Controller Mode**_ 

Table 6-73, Figure 6-67, Table 6-74, and Figure 6-68 present timing requirements and switching characteristics for SPI – Controller Mode. 

**Table 6-73. MCSPI Timing Requirements – Controller Mode** 

## see Figure 6-67 

**==> picture [500 x 491] intentionally omitted <==**

**----- Start of picture text -----**<br>
NO. PARAMETER DESCRIPTION MIN MAX UNIT<br>SM4 tsu(POCI-SPICLK) Setup time, SPIn_D[x] valid before SPIn_CLK active edge 2.8 ns<br>SM5 th(SPICLK-POCI) Hold time, SPIn_D[x] valid after SPIn_CLK active edge 3 ns<br>PHA=0<br>EPOL=1<br>SPI_CS[i] (OUT)<br>SM1<br>SM3<br>SM8 SM2 SM9<br>SPI_SCLK (OUT) POL=0<br>SM1<br>SM3<br>SM2<br>POL=1<br>SPI_SCLK (OUT)<br>SM5<br>SM5<br>SM4 SM4<br>SPI_D[x] (IN) Bit n-1 Bit n-2 Bit n-3 Bit n-4 Bit 0<br>PHA=1<br>EPOL=1<br>SPI_CS[i] (OUT)<br>SM2<br>S M1<br>S M8 SM3 S M9<br>SPI_SCLK (OUT) POL=0<br>SM1<br>SM2<br>SM3<br>POL=1<br>SPI_SCLK (OUT)<br>SM5<br>SM4<br>SM4 SM5<br>SPI_D[x] (IN) Bit n-1 Bit n-2 Bit n-3 Bit 1 Bit 0<br>SPRSP08_TIMING_McSPI_02<br>**----- End of picture text -----**<br>


**Figure 6-67. SPI Controller Mode Receive Timing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 157 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **Table 6-74. MCSPI Switching Characteristics - Controller Mode** 

|seeFigure 6-68|seeFigure 6-68|seeFigure 6-68||||
|---|---|---|---|---|---|
|**NO.**|**PARAMETER**|||**MIN**<br>**MAX**|**UNIT**|
|SM1|tc(SPICLK)|Cycle time, SPIn_CLK||20|ns|
|SM2|tw(SPICLKL)|Pulse duration, SPIn_CLK low||0.5P - 1 (1)|ns|
|SM3|tw(SPICLKH)|Pulse duration, SPIn_CLK high||0.5P - 1 (1)|ns|
|SM6|td(SPICLK-PICO)|Delay time, SPIn_CLK active edge to SPIn_D[x]||-3<br>2.5|ns|
|SM7|td(CS-PICO)|Delay time, SPIn_CSi active edge to SPIn_D[x]||5|ns|
|SM8|td(CS-SPICLK)|Delay time, SPIn_CSi active to SPIn_CLK first edge|PHA = 0|B - 4<br>(2)|ns|
||||PHA = 1|A - 4<br>(3)|ns|
|SM9|td(SPICLK-CS)|Delay time, SPIn_CLK last edge to SPIn_CSi inactive|PHA = 0|A - 4 (4)|ns|
||||PHA = 1|B - 4 (5)|ns|



(1) P = SPIn_CLK period in ns. 

- (2) T_ref is the period of the McSPI functional clock in ns. Fratio is the divide ratio of McSPI functional clock frequency to SPIn_CLK clock frequency, controlled by the CLKD and CLKG bit fields in the MCSPI_CH(i)CONF register and the EXTCLK bit field in the MCSPI_CH(i)CTRL register. TCS(i) is the value programmed into the chip select time control bit field of the MCSPI_CH(i)CONF register. 

   - When Fratio = 1; B = (TCS(i) + 0.5) * T_ref. 

   - When Fratio ≥ 2 and even value; B = (TCS(i) + 0.5) * Fratio * T_ref. 

   - When Fratio ≥ 3 and odd value; B = ((TCS(i) * Fratio) + ((Fratio + 1) / 2 )) * T_ref. 

- (3) T_ref is the period of the McSPI functional clock. Fratio is the divide ratio of McSPI functional clock frequency to SPIn_CLK clock frequency, controlled by the CLKD and CLKG bit fields in the MCSPI_CH(i)CONF register and the EXTCLK bit field in the MCSPI_CH(i)CTRL register. TCS(i) is the value programmed into the chip select time control bit field of the MCSPI_CH(i)CONF register. 

   - When Fratio = 1; A = (TCS(i) + 1) * T_ref. 

   - When Fratio ≥ 2 and even value; A = (TCS(i) + 0.5) * Fratio * T_ref. 

   - When Fratio ≥ 3 and odd value; A = ((TCS(i) * Fratio) + ((Fratio - 1) / 2 )) * T_ref. 

- (4) T_ref is the period of the McSPI functional clock. Fratio is the divide ratio of McSPI functional clock frequency to SPIn_CLK clock frequency, controlled by the CLKD and CLKG bit fields in the MCSPI_CH(i)CONF register and the EXTCLK bit field in the MCSPI_CH(i)CTRL register. TCS(i) is the value programmed into the chip select time control bit field of the MCSPI_CH(i)CONF register. 

   - When Fratio = 1; A = (TCS(i) + 1) * T_ref. 

   - When Fratio ≥ 2 and even value; A = (TCS(i) + 0.5) * Fratio * T_ref. 

   - When Fratio ≥ 3 and odd value; A = ((TCS(i) * Fratio) + ((Fratio + 1) / 2 )) * T_ref. 

(5) T_ref is the period of the McSPI functional clock. Fratio is the divide ratio of McSPI functional clock frequency to SPIn_CLK clock frequency, controlled by the CLKD and CLKG bit fields in the MCSPI_CH(i)CONF register and the EXTCLK bit field in the MCSPI_CH(i)CTRL register. TCS(i) is the value programmed into the chip select time control bit field of the MCSPI_CH(i)CONF register. 

- When Fratio = 1; B = (TCS(i) + 0.5) * T_ref. 

- When Fratio ≥ 2 and even value; B = (TCS(i) + 0.5) * Fratio * T_ref. 

- When Fratio ≥ 3 and odd value; B = ((TCS(i) * Fratio) + ((Fratio - 1) / 2 )) * T_ref. 

Copyright © 2025 Texas Instruments Incorporated 

158 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **www.ti.com** 

**==> picture [499 x 385] intentionally omitted <==**

**----- Start of picture text -----**<br>
PHA=0<br>EPOL=1<br>SPI_CS[i] (OUT)<br>SM1<br>SM3<br>SM8 SM2 SM9<br>SPI_SCLK (OUT) POL=0<br>SM1<br>SM3<br>POL=1 SM2<br>SPI_SCLK (OUT)<br>SM7 SM6 SM6<br>SPI_D[x] (OUT) Bit n-1 Bit n-2 Bit n-3 Bit n-4 Bit 0<br>PHA=1<br>EPOL=1<br>SPI_CS[i] (OUT)<br>SM1<br>SM2<br>S M8 SM3 S M9<br>SPI_SCLK (OUT) POL=0<br>SM1<br>SM2<br>POL=1 SM3<br>SPI_SCLK (OUT)<br>SM6 SM6 SM6 SM6<br>SPI_D[x] (OUT) Bit n-1 Bit n-2 Bit n-3 Bit 1 Bit0<br>**----- End of picture text -----**<br>


SPRSP08_TIMING_McSPI_01 

**Figure 6-68. SPI Controller Mode Transmit Timing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 159 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.15.2 MCSPI — Peripheral Mode**_ 

Table 6-75, Figure 6-69, Table 6-76, and Figure 6-70 present timing requirements and switching characteristics for SPI – Peripheral Mode. 

**Table 6-75. MCSPI Timing Requirements – Peripheral Mode** 

see Figure 6-69 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|SS1|tc(SPICLK)|Cycle time, SPIn_CLK|20|ns|
|SS2|tw(SPICLKL)|Pulse duration, SPIn_CLK low|0.45P (1)|ns|
|SS3|tw(SPICLKH)|Pulse duration, SPIn_CLK high|0.45P (1)|ns|
|SS4|tsu(PICO-SPICLK)|Setup time, SPIn_D[x] valid before SPIn_CLK active edge|5|ns|
|SS5|th(SPICLK-PICO)|Hold time, SPIn_D[x] valid after SPIn_CLK active edge|5|ns|
|SS8|tsu(CS-SPICLK)|Setup time, SPIn_CSi valid before SPIn_CLK first edge|5|ns|
|SS9|th(SPICLK-CS)|Hold time, SPIn_CSi valid after SPIn_CLK last edge|5|ns|



(1) P = SPIn_CLK period in ns. 

**==> picture [500 x 410] intentionally omitted <==**

**----- Start of picture text -----**<br>
PHA=0<br>EPOL=1<br>SPI_CS[i] (IN)<br>SS1<br>SS2<br>SS8 SS3 SS9<br>POL=0<br>SPI_SCLK (IN)<br>SS1<br>SS2<br>POL=1 SS3<br>SPI_SCLK (IN)<br>SS5 SS4<br>SS4 SS5<br>Bit n-1 Bit n-2 Bit n-3 Bit n-4 Bit 0<br>SPI_D[x] (IN)<br>PHA=1<br>EPOL=1<br>SPI_CS[i] (IN)<br>SS1<br>SS2<br>S S8 SS3 SS9<br>POL=0<br>SPI_SCLK (IN)<br>SS1<br>SS3<br>POL=1 SS2<br>SPI_SCLK (IN)<br>SS4<br>SS5<br>SS4 SS5<br>Bit n-1 Bit n-2 Bit n-3 Bit 1 Bit 0<br>SPI_D[x] (IN)<br>**----- End of picture text -----**<br>


SPRSP08_TIMING_McSPI_04 

## **Figure 6-69. SPI Peripheral Mode Receive Timing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

160 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-76. MCSPI Switching Characteristics – Peripheral Mode** 

|**Table 6-76. MCSPI Switching Characteristics – Peripheral Mode**|**Table 6-76. MCSPI Switching Characteristics – Peripheral Mode**|**Table 6-76. MCSPI Switching Characteristics – Peripheral Mode**|**Table 6-76. MCSPI Switching Characteristics – Peripheral Mode**|**Table 6-76. MCSPI Switching Characteristics – Peripheral Mode**|
|---|---|---|---|---|
|seeFigure 6-70|||||
|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|SS6|td(SPICLK-POCI)|Delay time, SPIn_CLK active edge to SPIn_D[x]|2<br>17.12|ns|
|SS7|tsk(CS-POCI)|Delay time, SPIn_CSi active edge to SPIn_D[x]|20.95|ns|
|SPI_CS[i](IN)<br>SPI_SCLK(IN)<br>SPI_SCLK(IN)<br>SPI_D[x](OUT)<br>SPI_CS[i](IN)<br>SPI_SCLK(IN)<br>SPI_SCLK(IN)<br>SPI_D[x](OUT)<br>Bit n-1<br>Bit n-2<br>Bit n-3<br>Bit n-4<br>Bit 0<br>Bit n-1<br>Bit n-2<br>Bit n-3<br>Bit 1<br>Bit 0<br>PHA=0<br>EPOL=1<br>POL=0<br>POL=1<br>POL=0<br>POL=1<br>PHA=1<br>EPOL=1<br>SS6<br>SS3<br>SS1<br>SS3<br>SS1<br>SS3<br>SS1<br>SS2<br>~~S~~S1<br>SS6<br>SS6<br>SS8<br>~~S~~S9<br>SS7<br>~~S~~S8<br>SS2<br>SS3<br>SS2<br>SS2<br>SS6<br>SS6<br>SS6<br>SS9|||||



SPRSP08_TIMING_McSPI_03 

**Figure 6-70. SPI Peripheral Mode Transmit Timing** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 161 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.16 MMCSD** 

The MMCSD Host Controller provides an interface to embedded Multi-Media Card (MMC), Secure Digital (SD), and Secure Digital IO (SDIO) devices. The MMCSD Host Controller deals with MMC/SD/SDIO protocol at transmission level, data packing, adding cyclic redundancy checks (CRCs), start/end bit insertion, and checking for syntactical correctness. 

For more details about MMCSD interfaces, see the corresponding MMC0, MMC1, and MMC2 subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

## **Note** 

Some operating modes require software configuration of the MMC DLL delay settings, as shown in Table 6-77 and Table 6-98. 

The modes which show a value of "Tuning" in the ITAPDLYSEL column of Table 6-77 and Table 6-98 require a tuning algorithm to be used for optimizing input timing. Refer to the MMCSD Programming Guide in the device TRM for more information on the tuning algorithm and configuration of input delays required to optimize input timing. 

For more information, see _Multi-Media Card/Secure Digital (MMCSD) Interface_ section in _Peripherals_ chapter in the device TRM. 

## _**6.11.5.16.1 MMC0 - eMMC/SD/SDIO Interface**_ 

MMC0 interface is compliant with the JEDEC eMMC electrical standard v5.1 (JESD84-B51) and it supports the following eMMC applications: 

- Legacy SDR 

- High Speed SDR 

- High Speed DDR 

- HS200 

MMC0 interface is also compliant with the SD Host Controller Standard Specification 4.10 and SD Physical Layer Specification v3.01 as well as SDIO Specification v3.00. The following data transfer modes are only available for connectivity to embedded SDIO devices: 

- Default Speed 

- High Speed 

- UHS–I SDR12 

- UHS–I SDR25 

Copyright © 2025 Texas Instruments Incorporated 

162 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

Table 6-77 presents the required DLL software configuration settings for MMC0 timing modes. 

**Table 6-77. MMC0 DLL Delay Mapping for all Timing Modes** 

|**REGISTER NAME**|**REGISTER NAME**|**MMCSD0_MMC_SSCFG_PHY_CTRL_4_REG**|**MMCSD0_MMC_SSCFG_PHY_CTRL_4_REG**|**MMCSD0_MMC_SSCFG_PHY_CTRL_4_REG**|**MMCSD0_MMC_SSCFG_PHY_CTRL_4_REG**|
|---|---|---|---|---|---|
|**BIT FIELD**||**[20]**|**[16:12]**|**[8]**|**[4:0]**|
|**BIT FIELD NAME**||**OTAPDLYENA**|**OTAPDLYSEL**|**ITAPDLYENA**|**ITAPDLYSEL**|
|**MODE**|**DESCRIPTION**|**OUTPUT**<br>**DELAY**<br>**ENABLE**|**OUTPUT**<br>**DELAY**<br>**VALUE**|**INPUT**<br>**DELAY**<br>**ENABLE**|**INPUT**<br>**DELAY**<br>**VALUE**|
|Legacy<br>SDR|8-bit PHY<br>operating<br>1.8V, 25MHz|NA(1)|NA(1)|0x0|NA(2)|
||8-bit PHY<br>operating<br>3.3V, 25MHz|NA(1)|NA(1)|0x0|NA(2)|
|High<br>Speed<br>SDR|8-bit PHY<br>operating<br>1.8V, 50MHz|NA(1)|NA(1)|0x0|NA(2)|
||8-bit PHY<br>operating<br>3.3V, 50MHz|NA(1)|NA(1)|0x0|NA(2)|
|High<br>Speed<br>DDR|8-bit PHY<br>operating<br>1.8V, 40MHz|0x1|0x15|0x1|0x2|
||8-bit PHY<br>operating<br>3.3V, 40MHz|0x1|0x15|0x1|0x2|
|HS200|8-bit PHY<br>operating<br>1.8V, 200MHz|0x1|0x6|0x1|Tuning(3)|
|Default<br>Speed|4-bit PHY<br>operating<br>3.3V, 25MHz|NA(1)|NA(1)|0x1|0x0|
|High<br>Speed|4-bit PHY<br>operating<br>3.3V, 50MHz|NA(1)|NA(1)|0x1|0x0|
|UHS-I<br>SDR12|4-bit PHY<br>operating<br>1.8V, 25MHz|0x1|0xF|0x1|0x0|
|UHS-I<br>SDR25|4-bit PHY<br>operating<br>1.8V, 50MHz|0x1|0xF|0x1|0x0|



(1) NA means this register field has no function when operating with half-cycle timing, which is required for this mode. 

(2) NA means this register field has no function when ITAPDLYENA is set to 0x0. 

(3) Tuning means this mode requires a tuning algorithm to be used to determine optimal input timing 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 163 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

Table 6-78 presents timing conditions for MMC0. 

**Table 6-78. MMC0 Timing Conditions** 

|**PARAMETER**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**INPUT CONDITIONS**|||||
|SRI|Input slew rate|Legacy SDR @ 3.3V<br>High Speed SDR @ 3.3V<br>Default Speed<br>High Speed|0.69<br>2.06|V/ns|
|||Legacy SDR @ 1.8V<br>UHS-I SDR12|0.14<br>1.44|V/ns|
|||High Speed SDR @ 1.8V<br>UHS-I SDR25|0.3<br>1.34|V/ns|
|||High Speed DDR<br>UHS-I DDR50|1<br>2|V/ns|
|**OUTPUT CONDITIONS**|||||
|CL|Output load capacitance|HS200<br>UHS-I SDR104|1<br>10|pF|
|||All other modes|1<br>12|pF|
|**PCB CONNECTIVITY REQUIREMENTS**|||||
|td(Trace Delay)|Propagation delay of each trace|Legacy SDR<br>High Speed SDR<br>High Speed DDR<br>HS200|126<br>756|ps|
|||Default Speed<br>High Speed<br>UHS-I SDR12<br>UHS-I SDR25<br>UHS-I SDR50<br>UHS-I SDR104|126<br>1386|ps|
|||UHS-I DDR50|239<br>1134|ps|
|td(Trace Mismatch<br>Delay)|Propagation delay mismatch across all<br>traces|High Speed SDR<br>HS200<br>High Speed<br>UHS-I SDR104|8|ps|
|||High Speed DDR<br>UHS-I DDR50|20|ps|
|||All other modes|100|ps|



Copyright © 2025 Texas Instruments Incorporated 

164 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.1 Legacy SDR Mode**_ 

Table 6-79, Figure 6-71, Table 6-80, and Figure 6-72 present timing requirements and switching characteristics for MMC0 – Legacy SDR Mode. 

## **Table 6-79. MMC0 Timing Requirements – Legacy SDR Mode** 

see Figure 6-71 

|**NO.**|||**IO**<br>**Operating**<br>**Voltage**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|LSDR1|tsu(cmdV-clkH)|Setup time, MMC0_CMD valid before MMC0_CLK rising edge|1.8V|4.2|ns|
||||3.3V|2.15|ns|
|LSDR2|th(clkH-cmdV)|Hold time, MMC0_CMD valid after MMC0_CLK rising edge|1.8V|0.87|ns|
||||3.3V|1.67|ns|
|LSDR3|tsu(dV-clkH)|Setup time, MMC0_DAT[7:0] valid before MMC0_CLK rising edge|1.8V|4.2|ns|
||||3.3V|2.15|ns|
|LSDR4|th(clkH-dV)|Hold time, MMC0_DAT[7:0] valid after MMC0_CLK rising edge|1.8V|0.87|ns|
||||3.3V|1.67|ns|



**==> picture [341 x 85] intentionally omitted <==**

**Figure 6-71. MMC0 – Legacy SDR – Receive Mode** 

**Table 6-80. MMC0 Switching Characteristics – Legacy SDR Mode** 

see Figure 6-72 

|**NO.**|**PARAMETER**|**PARAMETER**|**IO**<br>**Operating**<br>**Voltage**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK||25|MHz|
|LSDR5|tc(clk)|Cycle time, MMC0_CLK||40|ns|
|LSDR6|tw(clkH)|Pulse duration, MMC0_CLK high||18.7|ns|
|LSDR7|tw(clkL)|Pulse duration, MMC0_CLK low||18.7|ns|
|LSDR8|td(clkL-cmdV)|Delay time, MMC0_CLK falling edge to MMC0_CMD transition|1.8V|-2.1<br>2.1|ns|
||||3.3V|-1.8<br>2.2|ns|
|LSDR9|td(clkL-dV)|Delay time, MMC0_CLK falling edge to MMC0_DAT[7:0] transition|1.8V|-2.1<br>2.1|ns|
||||3.3V|-1.8<br>2.2|ns|



**==> picture [341 x 109] intentionally omitted <==**

**Figure 6-72. MMC0 – Legacy SDR – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

165 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.2 High Speed SDR Mode**_ 

Table 6-81, Figure 6-73, Table 6-82, and Figure 6-74 present timing requirements and switching characteristics for MMC0 – High Speed SDR Mode. 

## **Table 6-81. MMC0 Timing Requirements – High Speed SDR Mode** 

see Figure 6-73 

|**NO.**|||**IO**<br>**Operating**<br>**Voltage**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|HSSDR1|tsu(cmdV-clkH)|Setup time, MMC0_CMD valid before MMC0_CLK rising edge|1.8V|2.15|ns|
||||3.3V|2.24|ns|
|HSSDR2|th(clkH-cmdV)|Hold time, MMC0_CMD valid after MMC0_CLK rising edge|1.8V|1.27|ns|
||||3.3V|1.66|ns|
|HSSDR3|tsu(dV-clkH)|Setup time, MMC0_DAT[7:0] valid before MMC0_CLK rising edge|1.8V|2.15|ns|
||||3.3V|2.24|ns|
|HSSDR4|th(clkH-dV)|Hold time, MMC0_DAT[7:0] valid after MMC0_CLK rising edge|1.8V|1.27|ns|
||||3.3V|1.66|ns|



**==> picture [341 x 85] intentionally omitted <==**

**Figure 6-73. MMC0 – High Speed SDR Mode – Receive Mode** 

**Table 6-82. MMC0 Switching Characteristics – High Speed SDR Mode** 

see Figure 6-74 

|**NO.**|**PARAMETER**|**PARAMETER**|**IO**<br>**Operating**<br>**Voltage**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK||50|MHz|
|HSSDR5|tc(clk)|Cycle time, MMC0_CLK||20|ns|
|HSSDR6|tw(clkH)|Pulse duration, MMC0_CLK high||9.2|ns|
|HSSDR7|tw(clkL)|Pulse duration, MMC0_CLK low||9.2|ns|
|HSSDR8|td(clkL-cmdV)|Delay time, MMC0_CLK falling edge to MMC0_CMD transition|1.8V|-1.55<br>3.05|ns|
||||3.3V|-1.8<br>2.2|ns|
|HSSDR9|td(clkL-dV)|Delay time, MMC0_CLK falling edge to MMC0_DAT[7:0] transition|1.8V|-1.55<br>3.05|ns|
||||3.3V|-1.8<br>2.2|ns|



**==> picture [341 x 109] intentionally omitted <==**

## **Figure 6-74. MMC0 – High Speed SDR Mode – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

166 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.3 High Speed DDR Mode**_ 

Table 6-83, Figure 6-75, Table 6-84, and Figure 6-76 present timing requirements and switching characteristics for MMC0 – High Speed DDR Mode. 

## **Table 6-83. MMC0 Timing Requirements – High Speed DDR Mode** 

see Figure 6-75 

|**NO.**|||**IO**<br>**Operating**<br>**Voltage**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|HSDDR1|tsu(cmdV-clk)|Setup time, MMC0_CMD valid before MMC0_CLK rising edge|1.8V|0.02|ns|
||||3.3V|1.5|ns|
|HSDDR2|th(clk-cmdV)|Hold time, MMC0_CMD valid after MMC0_CLK rising edge|1.8V|1.99|ns|
||||3.3V|1.75|ns|
|HSDDR3|tsu(dV-clk)|Setup time, MMC0_DAT[7:0] valid before MMC0_CLK transition|1.8V|0.02|ns|
||||3.3V|1.5|ns|
|HSDDR4|th(clk-dV)|Hold time, MMC0_DAT[7:0] valid after MMC0_CLK transition|1.8V|1.99|ns|
||||3.3V|1.75|ns|



**==> picture [479 x 96] intentionally omitted <==**

**Figure 6-75. MMC0 – High Speed DDR Mode – Receive Mode** 

**Table 6-84. MMC0 Switching Characteristics – High Speed DDR Mode** 

see Figure 6-76 

|**NO.**|**PARAMETER**|**PARAMETER**|**IO**<br>**Operating**<br>**Voltage**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK||40|MHz|
|HSDDR5|tc(clk)|Cycle time, MMC0_CLK||25|ns|
|HSDDR6|tw(clkH)|Pulse duration, MMC0_CLK high||11.58|ns|
|HSDDR7|tw(clkL)|Pulse duration, MMC0_CLK low||11.58|ns|
|HSDDR8|td(clk-cmdV)|Delay time, MMC0_CLK rising edge to MMC0_CMD transition|1.8V|1.2<br>5.6|ns|
||||3.3V|3.32<br>9.3|ns|
|HSDDR9|td(clk-dV)|Delay time, MMC0_CLK transition to MMC0_DAT[7:0] transition|1.8V|1.2<br>4.8|ns|
||||3.3V|3.2<br>8.9|ns|



**==> picture [435 x 101] intentionally omitted <==**

**Figure 6-76. MMC0 – High Speed DDR Mode – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 167 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.4 HS200 Mode**_ 

Table 6-85, Figure 6-77, Table 6-86, and Figure 6-78 present both timing requirements and switching characteristics for MMC0 – HS200 Mode. 

## **Table 6-85. MMC0 Timing Requirements – HS200 Mode** 

## see Figure 6-77 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|HS2004|tDVW<br>|Input data valid window, MMC0_CMD and MMC0_DAT[7:0]|2.0(1)|ns|



- (1) This parameter defines the minimum data valid window required by the host, where any data valid window presented to the host greater than this value ensures the host is able to capture valid data. The value defined by this parameter is smaller than the smallest possible data valid window defined for any eMMC device operating in HS200 mode. 

**==> picture [393 x 66] intentionally omitted <==**

**----- Start of picture text -----**<br>
HS2004<br>VIH<br>MMC0_CMD Valid<br>MMC0_DAT[7:0] Window<br>VIL<br>**----- End of picture text -----**<br>


**Figure 6-77. MMC0 – HS200 – Receive Mode** 

**Table 6-86. MMC0 Switching Characteristics – HS200 Mode** 

see Figure 6-78 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK|200|MHz|
|HS2005|tc(clk)|Cycle time, MMC0_CLK|5|ns|
|HS2006|tw(clkH)|Pulse duration, MMC0_CLK high|2.12|ns|
|HS2007|tw(clkL)|Pulse duration, MMC0_CLK low|2.12|ns|
|HS2008|td(clkL-cmdV)|Delay time, MMC0_CLK rising edge to MMC0_CMD transition|1.07<br>3.21|ns|
|HS2009|td(clkL-dV)|Delay time, MMC0_CLK rising edge to MMC0_DAT[7:0] transition|1.07<br>3.21|ns|



**==> picture [341 x 109] intentionally omitted <==**

**Figure 6-78. MMC0 – HS200 Mode – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

168 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.5 Default Speed Mode**_ 

Table 6-87, Figure 6-79, Table 6-88, and Figure 6-80 present timing requirements and switching characteristics for MMC0 – Default Speed Mode. 

## **Table 6-87. Timing Requirements for MMC0 – Default Speed Mode** 

## see Figure 6-79 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|DS1|tsu(cmdV-clkH)|Setup time, MMC0_CMD valid before MMC0_CLK rising edge|2.15|ns|
|DS2|th(clkH-cmdV)|Hold time, MMC0_CMD valid after MMC0_CLK rising edge|1.67|ns|
|DS3|tsu(dV-clkH)|Setup time, MMC0_DAT[3:0] valid before MMC0_CLK rising edge|2.15|ns|
|DS4|th(clkH-dV)|Hold time, MMC0_DAT[3:0] valid after MMC0_CLK rising edge|1.67|ns|
|MMC[ ]_CLK<br>x<br>M<br>[ ]_<br>MC<br>CMD<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x<br>DS1<br>~~DS2~~<br>DS3<br>~~DS4~~|||||



**Figure 6-79. MMC0 – Default Speed – Receive Mode** 

**Table 6-88. Switching Characteristics for MMC0 – Default Speed Mode** 

## see Figure 6-80 

|**NO.**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK|||25|MHz|
|DS5|tc(clk)|Cycle time, MMC0_CLK|||40|ns|
|DS6|tw(clkH)|Pulse duration, MMC0_CLK high|||18.7|ns|
|DS7|tw(clkL)|Pulse duration, MMC0_CLK low|||18.7|ns|
|DS8|td(clkL-cmdV)|Delay time, MMC0_CLK falling edge to MMC0_CMD transition|||- 1.8<br>2.2|ns|
|DS9|td(clkL-dV)|Delay time, MMC0_CLK falling edge to MMC0_DAT[3:0] transition|||- 1.8<br>2.2|ns|
|MMC[ ]_CLK<br>x<br>MMC<br>CMD<br>[ ]_<br>x<br>MMC<br>DAT<br>[ _<br>[3:0]<br>x]||~~D 5~~<br>~~S~~<br>D 8<br>S<br>D 9<br>S<br>~~DS6~~||~~D 7~~<br>~~S~~|||
||||||||
||||||||
||||||||



**Figure 6-80. MMC0 – Default Speed – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 169 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.6 High Speed Mode**_ 

Table 6-89, Figure 6-81, Table 6-90, and Figure 6-82 present timing requirements and switching characteristics for MMC0 – High Speed Mode. 

## **Table 6-89. Timing Requirements for MMC0 – High Speed Mode** 

## see Figure 6-81 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|HS1|tsu(cmdV-clkH)<br>|Setup time, MMC0_CMD valid before MMC0_CLK rising edge|2.24|ns|
|HS2|th(clkH-cmdV)<br>|Hold time, MMC0_CMD valid after MMC0_CLK rising edge|1.66|ns|
|HS3|tsu(dV-clkH)<br>|Setup time, MMC0_DAT[3:0] valid before MMC0_CLK rising edge|2.24|ns|
|HS4|th(clkH-dV)<br>|Hold time, MMC0_DAT[3:0] valid after MMC0_CLK rising edge|1.66|ns|
|MMC[ ]_CLK<br>x<br>MMC[ ]_CMD<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x<br>H 1<br>S<br>~~H 2~~<br>~~S~~<br>H 3<br>S<br>~~H 4~~<br>~~S~~|||||



**Figure 6-81. MMC0 – High Speed – Receive Mode** 

**Table 6-90. Switching Characteristics for MMC0 – High Speed Mode** 

## see Figure 6-82 

|**NO.**||**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK|||||50|MHz|
|HS5|tc(clk)|Cycle time. MMC0_CLK|||||20|ns|
|HS6|tw(clkH)|Pulse duration, MMC0_CLK high|||||9.2|ns|
|HS7|tw(clkL)|Pulse duration, MMC0_CLK low|||||9.2|ns|
|HS8|td(clkL-cmdV)|Delay time, MMC0_CLK falling edge to MMC0_CMD transition|||||-1.8<br>2.2|ns|
|HS9|td(clkL-dV)|Delay time, MMC0_CLK falling edge to MMC0_DAT[3:0] transition|||||-1.8<br>2.2|ns|
|MMC[x]_CLK<br>MMC<br>CMD<br>[ ]_<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x|||||~~H 7~~<br>~~S~~||||
||||~~H 5~~<br>~~S~~<br>~~HS6~~||||||
||||||||||
||||||~~H 7~~<br>~~S~~||||
||||||||||
||||||||||
|||H 8<br>S|||||||
||||||||||
||||||||||
||||||||||
|||H 9<br>S|||||||
||||||||||
||||||||||



**Figure 6-82. MMC0 – High Speed – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

170 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.7 UHS–I SDR12 Mode**_ 

Table 6-91, Figure 6-83, Table 6-92, and Figure 6-84 present timing requirements and switching characteristics for MMC0 – UHS-I SDR12 Mode. 

## **Table 6-91. Timing Requirements for MMC0 – UHS-I SDR12 Mode** 

## see Figure 6-83 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|SDR121|tsu(cmdV-clkH)|Setup time, MMC0_CMD valid before MMC0_CLK rising edge|4.2|ns|
|SDR122|th(clkH-cmdV)|Hold time, MMC0_CMD valid after MMC0_CLK rising edge|0.87|ns|
|SDR123|tsu(dV-clkH)|Setup time, MMC0_DAT[3:0] valid before MMC0_CLK rising edge|4.2|ns|
|SDR124|th(clkH-dV)|Hold time, MMC0_DAT[3:0] valid after MMC0_CLK rising edge|0.87|ns|
|MMC[ ]_CMD<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x<br>SDR121<br>~~SDR122~~<br>SDR123<br>~~SDR124~~<br>MMC[ ]_CLK<br>x|||||



**Figure 6-83. MMC0 – UHS-I SDR12 – Receive Mode** 

**Table 6-92. Switching Characteristics for MMC0 – UHS-I SDR12 Mode** 

## see Figure 6-84 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK|25|MHz|
|SDR125|tc(clk)|Cycle time, MMC0_CLK|40|ns|
|SDR126|tw(clkH)|Pulse duration, MMC0_CLK high|18.7|ns|
|SDR127|tw(clkL)|Pulse duration, MMC0_CLK low|18.7|ns|
|SDR128|td(clkL-cmdV)|Delay time, MMC0_CLK rising edge to MMC0_CMD transition|1.5<br>8.6|ns|
|SDR129|td(clkL-dV)|Delay time, MMC0_CLK rising edge to MMC0_DAT[3:0] transition|1.5<br>8.6|ns|



**==> picture [361 x 108] intentionally omitted <==**

**----- Start of picture text -----**<br>
SDR125<br>SDR126 SDR127<br>MMC[x]_CLK<br>SDR128 SDR128<br>MMC[x]_CMD<br>SDR129 SDR129<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**Figure 6-84. MMC0 – UHS-I SDR12 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 171 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.8 UHS–I SDR25 Mode**_ 

Table 6-93, Figure 6-85, Table 6-94, and Figure 6-86 present timing requirements and switching characteristics for MMC0 – UHS-I SDR25 Mode. 

## **Table 6-93. Timing Requirements for MMC0 – UHS-I SDR25 Mode** 

## see Figure 6-85 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|SDR251|tsu(cmdV-clkH)|Setup time, MMC0_CMD valid before MMC0_CLK rising edge|2.15|ns|
|SDR252|th(clkH-cmdV)|Hold time, MMC0_CMD valid after MMC0_CLK rising edge|1.27|ns|
|SDR253|tsu(dV-clkH)|Setup time, MMC0_DAT[3:0] valid before MMC0_CLK rising edge|2.15|ns|
|SDR254|th(clkH-dV)|Hold time, MMC0_DAT[3:0] valid after MMC0_CLK rising edge|1.27|ns|
|MMC[ ]_CLK<br>x<br>MMC[ ]_CMD<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x<br>SDR251<br>~~SDR252~~<br>SDR253<br>~~SDR254~~|||||



**Figure 6-85. MMC0 – UHS-I SDR25 – Receive Mode** 

**Table 6-94. Switching Characteristics for MMC0 – UHS-I SDR25 Mode** 

## see Figure 6-86 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK|50|MHz|
|SDR255|tc(clk)|Cycle time, MMC0_CLK|20|ns|
|SDR256|tw(clkH)|Pulse duration, MMC0_CLK high|9.2|ns|
|SDR257|tw(clkL)|Pulse duration, MMC0_CLK low|9.2|ns|
|SDR258|td(clkL-cmdV)|Delay time, MMC0_CLK rising edge to MMC0_CMD transition|2.4<br>8.1|ns|
|SDR259|td(clkL-dV)|Delay time, MMC0_CLK rising edge to MMC0_DAT[3:0] transition|2.4<br>8.1|ns|



**==> picture [361 x 108] intentionally omitted <==**

**----- Start of picture text -----**<br>
SDR255<br>SDR256 SDR257<br>MMC[x]_CLK<br>SDR258 SDR258<br>MMC[x]_CMD<br>SDR259 SDR259<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**Figure 6-86. MMC0 – UHS-I SDR25 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

172 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.9 UHS–I SDR50 Mode**_ 

Table 6-95 and Figure 6-87 presents switching characteristics for MMC0 – UHS-I SDR50 Mode. 

**Table 6-95. Switching Characteristics for MMC0 – UHS-I SDR50 Mode** 

## see Figure 6-87 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK|100|MHz|
|SDR505|tc(clk)|Cycle time, MMC0_CLK|10|ns|
|SDR506|tw(clkH)|Pulse duration, MMC0_CLK high|4.45|ns|
|SDR507|tw(clkL)|Pulse duration, MMC0_CLK low|4.45|ns|
|SDR508|td(clkL-cmdV)|Delay time, MMC0_CLK rising edge to MMC0_CMD transition|1.2<br>6.35|ns|
|SDR509|td(clkL-dV)|Delay time, MMC0_CLK rising edge to MMC0_DAT[3:0] transition|1.2<br>6.35|ns|



**==> picture [361 x 109] intentionally omitted <==**

**----- Start of picture text -----**<br>
SDR505<br>SDR506 SDR507<br>MMC[x]_CLK<br>SDR508 SDR508<br>MMC[x]_CMD<br>SDR509 SDR509<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**Figure 6-87. MMC0 – UHS-I SDR50 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 173 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.10 UHS–I DDR50 Mode**_ 

Table 6-96 and Figure 6-88 present switching characteristics for MMC0 – UHS-I DDR50 Mode. 

**Table 6-96. Switching Characteristics for MMC0 – UHS-I DDR50 Mode** 

## see Figure 6-88 

|**NO.**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**MIN**<br>**MAX**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMC0_CLK|||||50|||MHz|
|DDR505|tc(clk)|Cycle time, MMC0_CLK|||||20|||ns|
|DDR506|tw(clkH)|Pulse duration, MMC0_CLK high|||||9.2|||ns|
|DDR507|tw(clkL)|Pulse duration, MMC0_CLK low|||||9.2|||ns|
|DDR508|td(clk-cmdV)|Delay time, MMC0_CLK rising edge to MMC0_CMD transition|||||1.12<br>6.43|||ns|
|DDR509|td(clk-dV)|Delay time, MMC0_CLK transition to MMC0_DAT[3:0] transition|||||1.12<br>6.43|||ns|
|MMC[x] CLK<br>_<br>MMC[x] CMD<br>_<br>MMC[x] DAT3<br>_<br>[ :0]<br>DDR508<br>DDR509<br>DDR509<br>~~DDR505~~<br>~~DDR506~~<br>~~DDR507~~|||||||||||
||||~~DDR506~~<br>~~DDR507~~||||||||
||||||||||||
||||||||||||
|||||||||DDR508|||
||||||||||||
||||||||||||
||||DDR509||DDR509||||||
||||||||||||
||||||||||||



**Figure 6-88. MMC0 – UHS-I DDR50 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

174 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.1.11 UHS–I SDR104 Mode**_ 

Table 6-97 and Figure 6-89 present switching characteristics for MMC0 – UHS-I SDR104 Mode. 

**Table 6-97. Switching Characteristics for MMC0 – UHS-I SDR104 Mode** 

|**Table 6-97. Switching Characteristics for MMC0 – UHS-I SDR104 Mode**|**Table 6-97. Switching Characteristics for MMC0 – UHS-I SDR104 Mode**|**Table 6-97. Switching Characteristics for MMC0 – UHS-I SDR104 Mode**|**Table 6-97. Switching Characteristics for MMC0 – UHS-I SDR104 Mode**|**Table 6-97. Switching Characteristics for MMC0 – UHS-I SDR104 Mode**|
|---|---|---|---|---|
|seeFigure 6-89|||||
|**NO.**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
||fop(clk)|Operating frequency, MMC0_CLK|200|MHz|
|SDR1045|tc(clk)|Cycle time, MMC0_CLK|5|ns|
|SDR1046|tw(clkH)|Pulse duration, MMC0_CLK high|2.12|ns|
|SDR1047|tw(clkL)|Pulse duration, MMC0_CLK low|2.12|ns|
|SDR1048|td(clkL-cmdV)|Delay time, MMC0_CLK rising edge to MMC0_CMD transition|1.07<br>3.21|ns|
|SDR1049|td(clkL-dV)|Delay time, MMC0_CLK rising edge to MMC0_DAT[3:0] transition|1.07<br>3.21|ns|



**==> picture [361 x 109] intentionally omitted <==**

**----- Start of picture text -----**<br>
SDR1045<br>SDR1046 SDR1047<br>MMC[x]_CLK<br>SDR1048 SDR1048<br>MMC[x]_CMD<br>SDR1049 SDR1049<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**Figure 6-89. MMC0 – UHS-I SDR104 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 175 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.2 MMC1/MMC2 - SD/SDIO Interface**_ 

MMC1/MMC2 interface is compliant with the SD Host Controller Standard Specification 4.10 and SD Physical Layer Specification v3.01 as well as SDIO Specification v3.00 and it supports the following SD Card applications: 

- Default speed 

- High speed 

- UHS–I SDR12 

- UHS–I SDR25 

- UHS–I SDR50 

- UHS–I DDR50 

- UHS–I SDR104 

Table 6-98 presents the required DLL software configuration settings for MMC1/2 timing modes. 

**Table 6-98. MMC1/MMC2 DLL Delay Mapping for all Timing Modes** 

|**REGISTER NAME**|**REGISTER NAME**|**MMCSD1_MMC_SSCFG_PHY_CTRL_4_REG**<br>**MMCSD2_MMC_SSCFG_PHY_CTRL_4_REG**|**MMCSD1_MMC_SSCFG_PHY_CTRL_4_REG**<br>**MMCSD2_MMC_SSCFG_PHY_CTRL_4_REG**|**MMCSD1_MMC_SSCFG_PHY_CTRL_4_REG**<br>**MMCSD2_MMC_SSCFG_PHY_CTRL_4_REG**|**MMCSD1_MMC_SSCFG_PHY_CTRL_4_REG**<br>**MMCSD2_MMC_SSCFG_PHY_CTRL_4_REG**|
|---|---|---|---|---|---|
|**BIT FIELD**||**[20]**|**[15:12]**|**[8]**|**[4:0]**|
|**BIT FIELD NAME**||**OTAPDLYENA**|**OTAPDLYSEL**|**ITAPDLYENA**|**ITAPDLYSEL**|
|**MODE**|**DESCRIPTION**|**DELAY**<br>**ENABLE**|**DELAY**<br>**VALUE**|**INPUT**<br>**DELAY**<br>**ENABLE**|**INPUT**<br>**DELAY**<br>**VALUE**|
|Default<br>Speed|4-bit PHY<br>operating<br>3.3V, 25MHz|NA(1)|NA(1)|0x1|0x0|
|High<br>Speed|4-bit PHY<br>operating<br>3.3V, 50MHz|NA(1)|NA(1)|0x1|0x0|
|UHS-I<br>SDR12|4-bit PHY<br>operating<br>1.8V, 25MHz|0x1|0xF|0x1|0x0|
|UHS-I<br>SDR25|4-bit PHY<br>operating<br>1.8V, 50MHz|0x1|0xF|0x1|0x0|
|UHS-I<br>SDR50|4-bit PHY<br>operating<br>1.8V, 100MHz|0x1|0xC|0x1|Tuning(2)|
|UHS-I<br>DDR50|4-bit PHY<br>operating<br>1.8V, 50MHz|0x1|0x9|0x1|Tuning(2)|
|UHS-I<br>SDR104|4-bit PHY<br>operating<br>1.8V, 200MHz|0x1|0x6|0x1|Tuning(2)|



(1) NA means this register field has no function when operating with half-cycle timing, which is required for this mode. 

- (2) Tuning means this mode requires a tuning algorithm to be used to determine optimal input timing 

Copyright © 2025 Texas Instruments Incorporated 

176 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

Table 6-99 presents timing conditions for MMC1. 

**Table 6-99. MMC1/MMC2 Timing Conditions** 

|**PARAMETER**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|**Input Conditions**|||||
|SRI|Input slew rate|Default Speed<br>High Speed|0.69<br>2.06|V/ns|
|||UHS–I SDR12<br>UHS–I SDR25|0.34<br>1.34|V/ns|
|||UHS–I DDR50|1<br>2|V/ns|
|**Output Conditions**|||||
|CL|Output load capacitance|All modes|1<br>10|pF|
|**PCB Connectivity Requirements**|||||
|td(Trace Delay)|Propagation delay of each trace|UHS–I DDR50|239<br>1134|ps|
|||All other modes|126<br>1386|ps|
|td(Trace Mismatch<br>Delay)|Propagation delay mismatch across all<br>traces|High Speed<br>UHS–I SDR104|8|ps|
|||UHS–I DDR50|20|ps|
|||All other modes|100|ps|



Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 177 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.2.1 Default Speed Mode**_ 

Table 6-100, Figure 6-90, Table 6-101, and Figure 6-91 present timing requirements and switching characteristics for MMC1/MMC2 – Default Speed Mode. 

## **Table 6-100. Timing Requirements for MMC1/MMC2 – Default Speed Mode** 

## see Figure 6-90 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|DS1|tsu(cmdV-clkH)|Setup time, MMCx_CMD valid before MMCx_CLK rising edge|2.15|ns|
|DS2|th(clkH-cmdV)|Hold time, MMCx_CMD valid after MMCx_CLK rising edge|1.67|ns|
|DS3|tsu(dV-clkH)|Setup time, MMCx_DAT[3:0] valid before MMCx_CLK rising edge|2.15|ns|
|DS4|th(clkH-dV)|Hold time, MMCx_DAT[3:0] valid after MMCx_CLK rising edge|1.67|ns|
|MMC[ ]_CLK<br>x<br>M<br>[ ]_<br>MC<br>CMD<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x<br>DS1<br>~~DS2~~<br>DS3<br>~~DS4~~|||||



**Figure 6-90. MMC1/MMC2 – Default Speed – Receive Mode** 

**Table 6-101. Switching Characteristics for MMC1/MMC2 – Default Speed Mode** 

## see Figure 6-91 

|**NO.**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMCx_CLK|||25|MHz|
|DS5|tc(clk)|Cycle time, MMCx_CLK|||40|ns|
|DS6|tw(clkH)|Pulse duration, MMCx_CLK high|||18.7|ns|
|DS7|tw(clkL)|Pulse duration, MMCx_CLK low|||18.7|ns|
|DS8|td(clkL-cmdV)|Delay time, MMCx_CLK falling edge to MMCx_CMD transition|||- 1.8<br>2.2|ns|
|DS9|td(clkL-dV)|Delay time, MMCx_CLK falling edge to MMCx_DAT[3:0] transition|||- 1.8<br>2.2|ns|
|MMC[ ]_CLK<br>x<br>MMC<br>CMD<br>[ ]_<br>x<br>MMC<br>DAT<br>[ _<br>[3:0]<br>x]||~~D 5~~<br>~~S~~<br>D 8<br>S<br>D 9<br>S<br>~~DS6~~||~~D 7~~<br>~~S~~|||
||||||||
||||||||
||||||||



**Figure 6-91. MMC1/MMC2 – Default Speed – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

178 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.2.2 High Speed Mode**_ 

Table 6-102, Figure 6-92, Table 6-103, and Figure 6-93 present timing requirements and switching characteristics for MMC1/MMC2 – High Speed Mode. 

## **Table 6-102. Timing Requirements for MMC1/MMC2 – High Speed Mode** 

## see Figure 6-92 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|HS1|tsu(cmdV-clkH)<br>|Setup time, MMCx_CMD valid before MMCx_CLK rising edge|2.24|ns|
|HS2|th(clkH-cmdV)<br>|Hold time, MMCx_CMD valid after MMCx_CLK rising edge|1.66|ns|
|HS3|tsu(dV-clkH)<br>|Setup time, MMCx_DAT[3:0] valid before MMCx_CLK rising edge|2.24|ns|
|HS4|th(clkH-dV)<br>|Hold time, MMCx_DAT[3:0] valid after MMCx_CLK rising edge|1.66|ns|



**==> picture [56 x 73] intentionally omitted <==**

**----- Start of picture text -----**<br>
MMC[x]_CLK<br>MMC[x]_CMD<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**==> picture [303 x 72] intentionally omitted <==**

**----- Start of picture text -----**<br>
HS1 HS2<br>HS3 HS4<br>**----- End of picture text -----**<br>


**Figure 6-92. MMC1/MMC2 – High Speed – Receive Mode** 

**Table 6-103. Switching Characteristics for MMC1/MMC2 – High Speed Mode** 

## see Figure 6-93 

|**NO.**||**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMCx_CLK|||||50|MHz|
|HS5|tc(clk)|Cycle time. MMCx_CLK|||||20|ns|
|HS6|tw(clkH)|Pulse duration, MMCx_CLK high|||||9.2|ns|
|HS7|tw(clkL)|Pulse duration, MMCx_CLK low|||||9.2|ns|
|HS8|td(clkL-cmdV)|Delay time, MMCx_CLK falling edge to MMCx_CMD transition|||||- 1.8<br>2.2|ns|
|HS9|td(clkL-dV)|Delay time, MMCx_CLK falling edge to MMCx_DAT[3:0] transition|||||- 1.8<br>2.2|ns|
|MMC[x]_CLK<br>MMC<br>CMD<br>[ ]_<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x|||||~~H 7~~<br>~~S~~||||
||||~~H 5~~<br>~~S~~<br>~~HS6~~||||||
||||||||||
||||||~~H 7~~<br>~~S~~||||
||||||||||
||||||||||
|||H 8<br>S|||||||
||||||||||
||||||||||
||||||||||
|||H 9<br>S|||||||
||||||||||
||||||||||



**Figure 6-93. MMC1/MMC2 – High Speed – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 179 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.2.3 UHS–I SDR12 Mode**_ 

Table 6-104, Figure 6-94, Table 6-105, and Figure 6-95 present timing requirements and switching characteristics for MMC1/MMC2 – UHS-I SDR12 Mode. 

**Table 6-104. Timing Requirements for MMC1/MMC2 – UHS-I SDR12 Mode** 

## see Figure 6-94 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|SDR121|tsu(cmdV-clkH)|Setup time, MMCx_CMD valid before MMCx_CLK rising edge|4.2|ns|
|SDR122|th(clkH-cmdV)|Hold time, MMCx_CMD valid after MMCx_CLK rising edge|0.87|ns|
|SDR123|tsu(dV-clkH)|Setup time, MMCx_DAT[3:0] valid before MMCx_CLK rising edge|4.2|ns|
|SDR124|th(clkH-dV)|Hold time, MMCx_DAT[3:0] valid after MMCx_CLK rising edge|0.87|ns|
|MMC[ ]_CMD<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x<br>SDR121<br>~~SDR122~~<br>SDR123<br>~~SDR124~~<br>MMC[ ]_CLK<br>x|||||



**Figure 6-94. MMC1/MMC2 – UHS-I SDR12 – Receive Mode** 

**Table 6-105. Switching Characteristics for MMC1/MMC2 – UHS-I SDR12 Mode** 

## see Figure 6-95 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
||fop(clk)|Operating frequency, MMCx_CLK|25|MHz|
|SDR125|tc(clk)|Cycle time, MMCx_CLK|40|ns|
|SDR126|tw(clkH)|Pulse duration, MMCx_CLK high|18.7|ns|
|SDR127|tw(clkL)|Pulse duration, MMCx_CLK low|18.7|ns|
|SDR128|td(clkL-cmdV)|Delay time, MMCx_CLK rising edge to MMCx_CMD transition|1.5<br>8.6|ns|
|SDR129|td(clkL-dV)|Delay time, MMCx_CLK rising edge to MMCx_DAT[3:0] transition|1.5<br>8.6|ns|



**==> picture [361 x 108] intentionally omitted <==**

**----- Start of picture text -----**<br>
SDR125<br>SDR126 SDR127<br>MMC[x]_CLK<br>SDR128 SDR128<br>MMC[x]_CMD<br>SDR129 SDR129<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**Figure 6-95. MMC1/MMC2 – UHS-I SDR12 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

180 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.2.4 UHS–I SDR25 Mode**_ 

Table 6-106, Figure 6-96, Table 6-107, and Figure 6-97 present timing requirements and switching characteristics for MMC1/MMC2 – UHS-I SDR25 Mode. 

## **Table 6-106. Timing Requirements for MMC1/MMC2 – UHS-I SDR25 Mode** 

## see Figure 6-96 

|**NO.**|||**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|SDR251|tsu(cmdV-clkH)|Setup time, MMCx_CMD valid before MMCx_CLK rising edge|2.15|ns|
|SDR252|th(clkH-cmdV)|Hold time, MMCx_CMD valid after MMCx_CLK rising edge|1.27|ns|
|SDR253|tsu(dV-clkH)|Setup time, MMCx_DAT[3:0] valid before MMCx_CLK rising edge|2.15|ns|
|SDR254|th(clkH-dV)|Hold time, MMCx_DAT[3:0] valid after MMCx_CLK rising edge|1.27|ns|
|MMC[ ]_CLK<br>x<br>MMC[ ]_CMD<br>x<br>MMC<br>DAT<br>[ ]_<br>[3:0]<br>x<br>SDR251<br>~~SDR252~~<br>SDR253<br>~~SDR254~~|||||



**Figure 6-96. MMC1/MMC2 – UHS-I SDR25 – Receive Mode** 

**Table 6-107. Switching Characteristics for MMC1/MMC2 – UHS-I SDR25 Mode** 

## see Figure 6-97 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
||fop(clk)|Operating frequency, MMCx_CLK|50|MHz|
|SDR255|tc(clk)|Cycle time, MMCx_CLK|20|ns|
|SDR256|tw(clkH)|Pulse duration, MMCx_CLK high|9.2|ns|
|SDR257|tw(clkL)|Pulse duration, MMCx_CLK low|9.2|ns|
|SDR258|td(clkL-cmdV)|Delay time, MMCx_CLK rising edge to MMCx_CMD transition|2.4<br>8.1|ns|
|SDR259|td(clkL-dV)|Delay time, MMCx_CLK rising edge to MMCx_DAT[3:0] transition|2.4<br>8.1|ns|



**==> picture [361 x 108] intentionally omitted <==**

**----- Start of picture text -----**<br>
SDR255<br>SDR256 SDR257<br>MMC[x]_CLK<br>SDR258 SDR258<br>MMC[x]_CMD<br>SDR259 SDR259<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**Figure 6-97. MMC1/MMC2 – UHS-I SDR25 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 181 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.2.5 UHS–I SDR50 Mode**_ 

Table 6-108 and Figure 6-98 presents switching characteristics for MMC1/MMC2 – UHS-I SDR50 Mode. 

**Table 6-108. Switching Characteristics for MMC1/MMC2 – UHS-I SDR50 Mode** 

see Figure 6-98 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
||fop(clk)|Operating frequency, MMCx_CLK|100|MHz|
|SDR505|tc(clk)|Cycle time, MMCx_CLK|10|ns|
|SDR506|tw(clkH)|Pulse duration, MMCx_CLK high|4.45|ns|
|SDR507|tw(clkL)|Pulse duration, MMCx_CLK low|4.45|ns|
|SDR508|td(clkL-cmdV)|Delay time, MMCx_CLK rising edge to MMCx_CMD transition|1.2<br>6.35|ns|
|SDR509|td(clkL-dV)|Delay time, MMCx_CLK rising edge to MMCx_DAT[3:0] transition|1.2<br>6.35|ns|



**==> picture [361 x 109] intentionally omitted <==**

**----- Start of picture text -----**<br>
SDR505<br>SDR506 SDR507<br>MMC[x]_CLK<br>SDR508 SDR508<br>MMC[x]_CMD<br>SDR509 SDR509<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**Figure 6-98. MMC1/MMC2 – UHS-I SDR50 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

182 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.2.6 UHS–I DDR50 Mode**_ 

Table 6-109 and Figure 6-99 present switching characteristics for MMC1/MMC2 – UHS-I DDR50 Mode. 

**Table 6-109. Switching Characteristics for MMC1/MMC2 – UHS-I DDR50 Mode** 

## see Figure 6-99 

|**NO.**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**MIN**<br>**MAX**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|---|---|---|---|---|
||fop(clk)|Operating frequency, MMCx_CLK|||||50|||MHz|
|DDR505|tc(clk)|Cycle time, MMCx_CLK|||||20|||ns|
|DDR506|tw(clkH)|Pulse duration, MMCx_CLK high|||||9.2|||ns|
|DDR507|tw(clkL)|Pulse duration, MMCx_CLK low|||||9.2|||ns|
|DDR508|td(clk-cmdV)|Delay time, MMCx_CLK rising edge to MMCx_CMD transition|||||1.12<br>6.43|||ns|
|DDR509|td(clk-dV)|Delay time, MMCx_CLK transition to MMCx_DAT[3:0] transition|||||1.12<br>6.43|||ns|
|MMC[x] CLK<br>_<br>MMC[x] CMD<br>_<br>MMC[x] DAT3<br>_<br>[ :0]<br>DDR508<br>DDR509<br>DDR509<br>~~DDR505~~<br>~~DDR506~~<br>~~DDR507~~|||||||||||
||||~~DDR506~~<br>~~DDR507~~||||||||
||||||||||||
||||||||||||
|||||||||DDR508|||
||||||||||||
||||||||||||
||||DDR509||DDR509||||||
||||||||||||
||||||||||||



**Figure 6-99. MMC1/MMC2 – UHS-I DDR50 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 183 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.16.2.7 UHS–I SDR104 Mode**_ 

Table 6-110 and Figure 6-100 present switching characteristics for MMC1/MMC2 – UHS-I SDR104 Mode. 

**Table 6-110. Switching Characteristics for MMC1/MMC2 – UHS-I SDR104 Mode** 

|**Table 6-110. Switching Characteristics for MMC1/MMC2 – UHS-I SDR104 Mode**|**Table 6-110. Switching Characteristics for MMC1/MMC2 – UHS-I SDR104 Mode**|**Table 6-110. Switching Characteristics for MMC1/MMC2 – UHS-I SDR104 Mode**|**Table 6-110. Switching Characteristics for MMC1/MMC2 – UHS-I SDR104 Mode**|**Table 6-110. Switching Characteristics for MMC1/MMC2 – UHS-I SDR104 Mode**|
|---|---|---|---|---|
|seeFigure 6-100|||||
|**NO.**|**PARAMETER**||**MIN**<br>**MAX**|**UNIT**|
||fop(clk)|Operating frequency, MMCx_CLK|200|MHz|
|SDR1045|tc(clk)|Cycle time, MMCx_CLK|5|ns|
|SDR1046|tw(clkH)|Pulse duration, MMCx_CLK high|2.12|ns|
|SDR1047|tw(clkL)|Pulse duration, MMCx_CLK low|2.12|ns|
|SDR1048|td(clkL-cmdV)|Delay time, MMCx_CLK rising edge to MMCx_CMD transition|1.07<br>3.21|ns|
|SDR1049|td(clkL-dV)|Delay time, MMCx_CLK rising edge to MMCx_DAT[3:0] transition|1.07<br>3.21|ns|



**==> picture [361 x 109] intentionally omitted <==**

**----- Start of picture text -----**<br>
SDR1045<br>SDR1046 SDR1047<br>MMC[x]_CLK<br>SDR1048 SDR1048<br>MMC[x]_CMD<br>SDR1049 SDR1049<br>MMC[x]_DAT[3:0]<br>**----- End of picture text -----**<br>


**Figure 6-100. MMC1/MMC2 – UHS-I SDR104 – Transmit Mode** 

Copyright © 2025 Texas Instruments Incorporated 

184 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.17 OSPI** 

OSPI0 offers two data capture modes, PHY mode and Tap mode. 

PHY mode uses an internal reference clock to transmit and receive data via a DLL based PHY, where each reference clock cycle produces a single cycle of OSPI0_CLK for Single Data Rate (SDR) transfers or a half cycle of OSPI0_CLK for Double Data Rate (DDR) transfers. PHY mode supports four clocking topologies for the receive data capture clock. Internal PHY Loopback - uses the internal reference clock as the PHY receive data capture clock. Internal Pad Loopback - uses OSPI0_LBCLKO looped back into the PHY from the OSPI0_LBCLKO pin as the PHY receive data capture clock. External Board Loopback - uses OSPI0_LBCLKO looped back into the PHY from the OSPI0_DQS pin as the PHY receive data capture clock. DQS - uses the DQS output from the attached device as the PHY receive data capture clock. SDR transfers are not supported when using the Internal Pad Loopback and DQS clocking topologies. DDR transfers are not supported when using the Internal PHY Loopback or Internal Pad Loopback clocking topologies. 

Tap mode uses an internal reference clock with selectable taps to adjusted data transmit and receive capture delays relative to OSPI0_CLK, which is a divide by 4 of the internal reference clock for SDR transfers or a divide by 8 of the internal reference clock for DDR transfers. Tap mode only supports one clocking topology for the receive data capture clock. No Loopback - uses the internal reference clock as the Tap receive data capture clock. This clocking topology supports a maximum internal reference clock rate of 400MHz, which produces an OSPI0_CLK rate up to 100MHz for SDR mode or 50MHz for DDR mode. 

For more information, see _Octal Serial Peripheral Interface (OSPI)_ section in _Peripherals_ chapter in the device TRM. 

For more details about features and additional description information on the device Octal Serial Peripheral Interface, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

Section 6.11.5.17.1 defines timing requirements and switching characteristics associated with PHY mode and Section 6.11.5.17.2 defines timing requirements and switching characteristics associated with Tap mode. 

Table 6-111 presents timing conditions for OSPI0. 

|Table 6-111presents timing conditions for OSPI0.|Table 6-111presents timing conditions for OSPI0.|Table 6-111presents timing conditions for OSPI0.|Table 6-111presents timing conditions for OSPI0.|Table 6-111presents timing conditions for OSPI0.|
|---|---|---|---|---|
|**Table 6-111. OSPI0 Timing Conditions**|||||
|**PARAMETER**||**MODE**|**MIN**<br>**MAX**|**UNIT**|
|**INPUT CONDITIONS**|||||
|SRI|Input slew rate||1<br>6|V/ns|
|**OUTPUT CONDITIONS**|||||
|CL|Output load capacitance||3<br>10|pF|
|**PCB CONNECTIVITY REQUIREMENTS**|||||
|td(Trace Delay)|Propagation delay of OSPI0_CLK trace|No Loopback<br>Internal PHY Loopback<br>Internal Pad Loopback|450|ps|
||Propagation delay of OSPI0_LBCLKO<br>trace|External Board Loopback|2L(1)- 30<br>2L(1)+ 30|ps|
||Propagation delay of OSPI0_DQS trace|DQS|L(1)- 30<br>L(1)+ 30|ps|
|td(Trace Mismatch<br>Delay)|Propagation delay mismatch of<br>OSPI0_D[7:0] and OSPI0_CSn[3:0]<br>relative to OSPI0_CLK|All modes|60|ps|



(1) L = Propagation delay of OSPI0_CLK trace 

Copyright © 2025 Texas Instruments Incorporated 

185 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.17.1 OSPI0 PHY Mode**_ 

## _**6.11.5.17.1.1 OSPI0 With PHY Data Training**_ 

Read and write data valid windows will shift due to variation in process, voltage, temperature, and operating frequency. A data training method may be implemented to dynamically configure optimal read and write timing. Implementing data training enables proper operation across temperature with a specific process, voltage, and frequency operating condition, while achieving a higher operating frequency. 

Data transmit and receive timing parameters are not defined for the data training use case since they are dynamically adjusted based on the operating condition. 

Table 6-112 defines DLL delays required for OSPI0 with Data Training. Table 6-113, Figure 6-101, Figure 6-102, Table 6-114, Figure 6-103, and Figure 6-104 present timing requirements and switching characteristics for OSPI0 with Data Training. 

**Table 6-112. OSPI0 DLL Delay Mapping for PHY Data Training** 

|**MODE**|**REGISTER BIT FIELD**|**DELAY VALUE**|
|---|---|---|
|**OSPI_PHY_CONFIGURATION_REG**|||
|**Transmit**|||
|All modes|PHY_CONFIG_TX_DLL_DELAY_FLD|(1)|
|**Receive**|||
|All modes|PHY_CONFIG_RX_DLL_DELAY_FLD|(2)|
|**PHY_MASTER_CONTROL_REG**|||
|All modes|PHY_MASTER_PHASE_DETECT_SELECTOR_FLD|0x1|



(1) Transmit DLL delay value determined by training software 

(2) Receive DLL delay value determined by training software 

Copyright © 2025 Texas Instruments Incorporated 

186 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-113. OSPI0 Timing Requirements – PHY Data Training** 

see Figure 6-101 , and Figure 6-102 

|**NO.**|||**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|O15|tsu(D-LBCLK)|Setup time, OSPI0_D[7:0] valid before<br>active OSPI0_DQS edge|DDR with DQS|(1)|ns|
|O16|th(LBCLK-D)|Hold time, OSPI0_D[7:0] valid after active<br>OSPI0_DQS edge|DDR with DQS|(1)|ns|
|O21|tsu(D-LBCLK)|Setup time, OSPI0_D[7:0] valid before<br>active OSPI0_DQS edge|SDR with External Board Loopback|(1)|ns|
|O22|th(LBCLK-D)|Hold time, OSPI0_D[7:0] valid after active<br>OSPI0_DQS edge|SDR with External Board Loopback|(1)|ns|
||tDVW|Data valid window (O15 + O16)|DDR with DQS|1.6|ns|
|||Data valid window (O21 + O22)|SDR with External Board Loopback|2.3|ns|



(1) Minimum setup and hold time requirements for OSPI0_D[7:0] inputs are not defined when Data Training is used to find the optimum data valid window. The tDVW parameter defines the minimum data invalid window required. This parameter is provided in lieu of minimum setup and minimum hold times, where it must be used to check compatibility with the data valid window provided by an attached device. 

**==> picture [244 x 94] intentionally omitted <==**

**----- Start of picture text -----**<br>
O15 O16 O15 O16<br>OSPI_TIMING_04<br>**----- End of picture text -----**<br>


**==> picture [101 x 62] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_DQS<br>OSPI_D[i:0]<br>**----- End of picture text -----**<br>


**Figure 6-101. OSPI0 Timing Requirements – PHY Data Training, DDR with DQS** 

**==> picture [371 x 91] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_DQS<br>O21 O22<br>OSPI_D[i:0]<br>OSPI_TIMING_06<br>**----- End of picture text -----**<br>


**Figure 6-102. OSPI0 Timing Requirements – PHY Data Training, SDR with External Board Loopback** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 187 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**Table 6-114. OSPI Switching Characteristics – PHY Data Training** 

See Figure 6-103 and Figure 6-104 

|**NO.**||**PARAMETER**|**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|O1|tc(CLK)|Cycle time, OSPI0_CLK|DDR|6.0<br>10|ns|
|O7|||SDR|6.0<br>10|ns|
|O2|tw(CLKL)|Pulse duration, OSPI0_CLK low|DDR|((0.475P(1)) - 0.3)|ns|
|O8|||SDR|||
|O3|tw(CLKH)|Pulse duration, OSPI0_CLK high|DDR|((0.475P(1)) - 0.3)|ns|
|O9|||SDR|||
|O4|td(CSn-CLK)|Delay time, OSPI0_CSn[3:0] active edge<br>to OSPI0_CLK rising edge|DDR|((0.475P(1)) +<br>(0.975M(2)R(4)) +<br>(0.04TD(5)) - 1)<br>((0.525P(1)) +<br>(1.025M(2)R(4)) +<br>(0.11TD(5)) + 1)|<br> <br>ns|
|O10|||SDR|||
|O5|td(CLK-CSn)|Delay time, OSPI0_CLK rising edge to<br>OSPI0_CSn[3:0] inactive edge|DDR|((0.475P(1)) +<br>(0.975N(3)R(4)) -<br>(0.11TD(5)) - 1)<br>((0.525P(1)) +<br>(1.025N(3)R(4)) -<br>(0.04TD(5)) + 1)|<br> <br>ns|
|O11|||SDR|||
|O6|td(CLK-D)|Delay time, OSPI0_CLK active edge to<br>OSPI0_D[7:0] transition|DDR|(6)<br>(6)|ns|
|O12|||SDR|||
||tDIVW|Data invalid window (O6 Max - Min)|DDR|1.6|ns|
|||Data invalid window (O12 Max - Min)|SDR|||



(1) P = SCLK cycle time in ns = OSPI0_CLK cycle time in ns 

- (2) M = OSPI_DEV_DELAY_REG[D_INIT_FLD] 

(3) N = OSPI_DEV_DELAY_REG[D_AFTER_FLD] 

(4) R = reference clock cycle time in ns 

- (5) TD = PHY_CONFIG_TX_DLL_DELAY_FLD 

(6) Minimum and maximum delay times for OSPI0_D[7:0] outputs are not defined when Data Training is used to find the optimum data valid window. The tDIVW parameter defines the maximum data invalid window. This parameter is provided in lieu of minimum and maximum delay times, where it must be used to check compatibility with the data valid window requirements of an attached device. 

**==> picture [412 x 147] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_CSn<br>O4<br>O3 O5<br>OSPI_CLK<br>O2<br>O6 O6<br>O1<br>OSPI_D[i:0]<br>OSPI_TIMING_01<br>**----- End of picture text -----**<br>


**Figure 6-103. OSPI0 Switching Characteristics – PHY DDR Data Training** 

Copyright © 2025 Texas Instruments Incorporated 

188 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**==> picture [504 x 187] intentionally omitted <==**

**----- Start of picture text -----**<br>
AM62L<br>www.ti.com SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025<br>OSPI_CSn<br>O10 O7 O11<br>OSPI_CLK O9 O8<br>O12<br>OSPI_D[i:0]<br>OSPI_TIMING_02<br>**----- End of picture text -----**<br>


**Figure 6-104. OSPI0 Switching Characteristics – PHY SDR Data Training** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 189 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.17.1.2 OSPI0 Without Data Training**_ 

## **Note** 

Timing parameters defined in this section are only applicable when data training is not implemented and DLL delays are configured as described in Table 6-115. 

## _**6.11.5.17.1.2.1 OSPI0 PHY SDR Timing**_ 

Table 6-115 defines DLL delays required for OSPI0 PHY SDR Mode. Table 6-116, Figure 6-105, Figure 6-106, Table 6-117, and Figure 6-107 present timing requirements and switching characteristics for OSPI0 PHY SDR Mode. 

**Table 6-115. OSPI0 DLL Delay Mapping for PHY SDR Timing Modes** 

|**MODE**|**MODE**|**MODE**|**REGISTER BIT FIELD**|**REGISTER BIT FIELD**|**DELAY VALUE**|**DELAY VALUE**|**DELAY VALUE**|
|---|---|---|---|---|---|---|---|
|**OSPI_PHY_CONFIGURATION_REG**||||||||
|**Transmit**||||||||
|All modes|||PHY_CONFIG_TX_DLL_DELAY_FLD||0x3|||
|**Receive**||||||||
|SDR Internal PHY Loopback|||PHY_CONFIG_RX_DLL_DELAY_FLD||0x0|||
|SDR External Board Loopback|||PHY_CONFIG_RX_DLL_DELAY_FLD||0x4|||
|**PHY_MASTER_CONTROL_REG**||||||||
|All modes|||PHY_MASTER_PHASE_DETECT_SELECTOR_FLD||0x1|||
|**Table 6-116. OSPI0 Timing Requirements – PHY SDR Mode**<br>seeFigure 6-105andFigure 6-106||||||||
|**NO.**||||**MODE**||**MIN**<br>**MAX**|**UNIT**|
|O19|tsu(D-CLK)|Setup time, OSPI0_D[7:0] valid before<br>active OSPI0_CLK edge||SDR with Internal PHY Loopback||4.8|ns|
|O20|th(CLK-D)|Hold time, OSPI0_D[7:0] valid after active<br>OSPI0_CLK edge||SDR with Internal PHY Loopback||-0.5|ns|
|O21|tsu(D-LBCLK)|Setup time, OSPI0_D[7:0] valid before<br>active OSPI0_DQS edge||SDR with External Board Loopback||0.6|ns|
|O22|th(LBCLK-D)|Hold time, OSPI0_D[7:0] valid after active<br>OSPI0_DQS edge||SDR with External Board Loopback||1.7|ns|
|OSPI_CLK<br>OSPI_D[i:0]<br>O 9<br>1<br>O20||||||||



**==> picture [61 x 15] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_TIMING_05<br>**----- End of picture text -----**<br>


**Figure 6-105. OSPI0 Timing Requirements – PHY SDR with Internal PHY Loopback** 

**==> picture [371 x 91] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_DQS<br>O21 O22<br>OSPI_D[i:0]<br>OSPI_TIMING_06<br>**----- End of picture text -----**<br>


**Figure 6-106. OSPI0 Timing Requirements – PHY SDR with External Board Loopback** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

190 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-117. OSPI0 Switching Characteristics – PHY SDR Mode** 

see Figure 6-107 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|O7|tc(CLK)|Cycle time, OSPI0_CLK|7|ns|
|O8|tw(CLKL)|Pulse duration, OSPI0_CLK low|((0.475P(1)) - 0.3)|ns|
|O9|tw(CLKH)|Pulse duration, OSPI0_CLK high|((0.475P(1)) - 0.3)|ns|
|O10|td(CSn-CLK)|Delay time, OSPI0_CSn[3:0] active edge to<br>OSPI0_CLK rising edge|((0.475P(1)) +<br>(0.975M(2)R(4)) +<br>(0.04TD(5)) - 1)<br>((0.525P(1)) +<br>(1.025M(2)R(4)) +<br>(0.11TD(5)) + 1)|<br> <br>ns|
|O11|td(CLK-CSn)|Delay time, OSPI0_CLK rising edge to<br>OSPI0_CSn[3:0] inactive edge|((0.475P(1)) +<br>(0.975N(3)R(4)) -<br>(0.11TD(5)) - 1)<br>((0.525P(1)) +<br>(1.025N(3)R(4)) -<br>(0.04TD(5)) + 1)|<br> <br>ns|
|O12|td(CLK-D)|Delay time, OSPI0_CLK active edge to<br>OSPI0_D[7:0] transition|-1.16<br>1.25|ns|



- (1) P = SCLK cycle time in ns = OSPI0_CLK cycle time in ns 

(2) M = OSPI_DEV_DELAY_REG[D_INIT_FLD] 

- (3) N = OSPI_DEV_DELAY_REG[D_AFTER_FLD] 

- (4) R = reference clock cycle time in ns 

(5) TD = PHY_CONFIG_TX_DLL_DELAY_FLD 

**==> picture [412 x 147] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_CSn<br>O10 O7 O11<br>OSPI_CLK O9 O8<br>O12<br>OSPI_D[i:0]<br>OSPI_TIMING_02<br>**----- End of picture text -----**<br>


**Figure 6-107. OSPI0 Switching Characteristics – PHY SDR** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 191 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.17.2 OSPI0 Tap Mode**_ 

## _**6.11.5.17.2.1 OSPI0 Tap SDR Timing**_ 

Table 6-118, Figure 6-108, Table 6-119, and Figure 6-109 present timing requirements and switching characteristics for OSPI0 Tap SDR Mode. 

**Table 6-118. OSPI0 Timing Requirements – Tap SDR Mode** 

## see Figure 6-108 

|**NO.**|||**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|O19|tsu(D-CLK)|Setup time, OSPI0_D[7:0] valid before<br>active OSPI0_CLK edge|No Loopback|(7.7 -<br>(0.975T(1)R(2)))|ns|
|O20|th(CLK-D)|Hold time, OSPI0_D[7:0] valid after active<br>OSPI0_CLK edge|No Loopback|(- 2.15 +<br>(0.975T(1)R(2)))|ns|



- (1) T = OSPI_RD_DATA_CAPTURE_REG[DELAY_FLD] 

- (2) R = reference clock cycle time in ns 

**==> picture [371 x 91] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_CLK<br>O19 O20<br>OSPI_D[i:0]<br>OSPI_TIMING_05<br>**----- End of picture text -----**<br>


**Figure 6-108. OSPI0 Timing Requirements – Tap SDR, No Loopback** 

Copyright © 2025 Texas Instruments Incorporated 

192 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-119. OSPI0 Switching Characteristics – Tap SDR Mode** 

see Figure 6-109 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|O7|tc(CLK)|Cycle time, OSPI0_CLK|10|ns|
|O8|tw(CLKL)|Pulse duration, OSPI0_CLK low|((0.475P(1)) - 0.3)|ns|
|O9|tw(CLKH)|Pulse duration, OSPI0_CLK high|((0.475P(1)) - 0.3)|ns|
|O10|td(CSn-CLK)|Delay time, OSPI0_CSn[3:0] active edge to<br>OSPI0_CLK rising edge|((0.475P(1)) +<br>(0.975M(2)R(4)) - 1)<br>((0.525P(1)) +<br>(1.025M(2)R(4)) + 1)|<br>ns|
|O11|td(CLK-CSn)|Delay time, OSPI0_CLK rising edge to<br>OSPI0_CSn[3:0] inactive edge|((0.475P(1)) +<br>(0.975N(3)R(4)) - 1)<br>((0.525P(1)) +<br>(1.025N(3)R(4)) + 1)|<br>ns|
|O12|td(CLK-D)|Delay time, OSPI0_CLK active edge to<br>OSPI0_D[7:0] transition|- 2.0<br>1.5|ns|



- (1) P = SCLK cycle time in ns = OSPI0_CLK cycle time in ns 

- (2) M = OSPI_DEV_DELAY_REG[D_INIT_FLD] 

- (3) N = OSPI_DEV_DELAY_REG[D_AFTER_FLD] 

- (4) R = reference clock cycle time in ns 

**==> picture [412 x 148] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_CSn<br>O10 O7 O11<br>OSPI_CLK O9 O8<br>O12<br>OSPI_D[i:0]<br>OSPI_TIMING_02<br>**----- End of picture text -----**<br>


**Figure 6-109. OSPI0 Switching Characteristics – Tap SDR, No Loopback** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 193 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**6.11.5.17.2.2 OSPI0 Tap DDR Timing**_ 

Table 6-120, Figure 6-110, Table 6-121, and Figure 6-111 present timing requirements and switching characteristics for OSPI0 Tap DDR Mode. 

**Table 6-120. OSPI0 Timing Requirements – Tap DDR Mode** 

## see Figure 6-110 

|**NO.**|||**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|O13|tsu(D-CLK)|Setup time, OSPI0_D[7:0] valid before<br>active OSPI0_CLK edge|No Loopback|(8.0 -<br>(0.975T(1)R(2)))|ns|
|O14|th(CLK-D)|Hold time, OSPI0_D[7:0] valid after active<br>OSPI0_CLK edge|No Loopback|(- 2.0 +<br>(0.975T(1)R(2)))|ns|



**==> picture [421 x 118] intentionally omitted <==**

**----- Start of picture text -----**<br>
(1) T = OSPI_RD_DATA_CAPTURE_REG[DELAY_FLD]<br>(2) R = reference clock cycle time in ns<br>OSPI_CLK<br>O13 O14 O13 O14<br>OSPI_D[i:0]<br>OSPI_TIMING_03<br>**----- End of picture text -----**<br>


**Figure 6-110. OSPI0 Timing Requirements – Tap DDR, No Loopback** 

Copyright © 2025 Texas Instruments Incorporated 

194 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

**Table 6-121. OSPI0 Switching Characteristics – Tap DDR Mode** 

see Figure 6-111 

|**NO.**|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
|O1|tc(CLK)|Cycle time, OSPI0_CLK|20|ns|
|O2|tw(CLKL)|Pulse duration, OSPI0_CLK low|((0.475P(1)) - 0.3)|ns|
|O3|tw(CLKH)|Pulse duration, OSPI0_CLK high|((0.475P(1)) - 0.3)|ns|
|O4|td(CSn-CLK)|Delay time, OSPI0_CSn[3:0] active edge to<br>OSPI0_CLK rising edge|((0.475P(1)) +<br>((0.975M(2)R(5)) - 1)<br>((0.525P(1)) +<br>( 1.025M(2)R(5)) + 1)|ns|
|O5|td(CLK-CSn)|Delay time, OSPI0_CLK rising edge to<br>OSPI0_CSn[3:0] inactive edge|((0.475P(1)) +<br>(0.975N(3)R(5)) - 1)<br>((0.525P(1)) +<br>(1.025N(3)R(5)) + 1)|<br>ns|
|O6|td(CLK-D)|Delay time, OSPI0_CLK active edge to<br>OSPI0_D[7:0] transition|(- 1.0 +<br>(0.975(T(4)+ 1)R(5))<br>- (0.525P(1)))<br>(4.0 +<br>(1.025(T(4)+ 1)R(5))<br>- (0.475P(1)))|ns|



- (1) P = SCLK cycle time in ns = OSPI0_CLK cycle time in ns 

- (2) M = OSPI_DEV_DELAY_REG[D_INIT_FLD] 

- (3) N = OSPI_DEV_DELAY_REG[D_AFTER_FLD] 

(4) T = OSPI_RD_DATA_CAPTURE_REG[DDR_READ_DELAY_FLD] 

(5) R = reference clock cycle time in ns 

**==> picture [412 x 148] intentionally omitted <==**

**----- Start of picture text -----**<br>
OSPI_CSn<br>O4<br>O3 O5<br>OSPI_CLK<br>O2<br>O6 O6<br>O1<br>OSPI_D[i:0]<br>OSPI_TIMING_01<br>**----- End of picture text -----**<br>


**Figure 6-111. OSPI0 Switching Characteristics – Tap DDR, No Loopback** 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 195 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.18 Timers** 

For more details about features and additional description information on the device Timers, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

**Table 6-122. Timer Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|0.5<br>5|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|2<br>10|pF|



**Table 6-123. Timer Input Timing Requirements** 

## see Figure 6-112 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|T1|tw(TINPH)|Pulse duration, high|CAPTURE|4P(1)+<br>2.5|ns|
|T2|tw(TINPL)|Pulse duration, low|CAPTURE|4P(1)+<br>2.5|ns|



(1) P = functional clock period in ns. 

**Table 6-124. Timer Output Switching Characteristics** 

## see Figure 6-112 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MODE**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|---|
|T3|tw(TOUTH)|Pulse duration, high|PWM|4P(1)-<br>2.5|ns|
|T4|tw(TOUTL)|Pulse duration, low|PWM|4P(1)-<br>2.5|ns|
|(1)<br>P = functional clock period in ns.<br>T1<br>T2<br>TIMER_IOx (inputs)<br>T3<br>T4<br>TIMER_IOx (outputs)||||||



**==> picture [16 x 4] intentionally omitted <==**

**----- Start of picture text -----**<br>
TIMER_01<br>**----- End of picture text -----**<br>


## **Figure 6-112. Timer Timing Requirements and Switching Characteristics** 

For more information, see _Timers_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

196 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **6.11.5.19 UART** 

For more details about features and additional description information on the device Universal Asynchronous Receiver Transmitter, see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

**Table 6-125. UART Timing Conditions** 

|**PARAMETER**|**PARAMETER**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|
|**INPUT CONDITIONS**||||
|SRI|Input slew rate|0.5<br>5|V/ns|
|**OUTPUT CONDITIONS**||||
|CL|Output load capacitance|1<br>30(1)|pF|



(1) This value represents an absolute maximum load capacitance. As the UART baud rate increases, it may be necessary to reduce the load capacitance to a value less than this maximum limit to provide enough timing margin for the attached device. The output rise/fall times increase as capacitive load increases, which decreases the time data is valid for the receiver of the attached devices. Therefore, it is important to understand the minimum data valid time required by the attached device at the operating baud rate. Then use the device IBIS models to verify the actual load capacitance on the UART signals does not increase the rise/fall times beyond the point where the minimum data valid time of the attached device is violated. 

## **Table 6-126. UART Timing Requirements** 

|**Table 6-126. UART Timing Requirements**|**Table 6-126. UART Timing Requirements**|**Table 6-126. UART Timing Requirements**|**Table 6-126. UART Timing Requirements**|**Table 6-126. UART Timing Requirements**|
|---|---|---|---|---|
|seeFigure 6-113|||||
|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|1|tw(RXD)|Pulse width, receive data bit high or low|0.95U(1)<br>(2)<br>1.05U(1)<br>(2)|ns|
|2|tw(RXDS)|Pulse width, receive start bit low|0.95U(1)<br>(2)|ns|



(1) U = UART baud time in ns = 1/programmed baud rate. 

(2) This value defines the data valid time, where the input voltage is required to be above VIH or below VIL. 

**Table 6-127. UART Switching Characteristics** 

see Figure 6-113 

|**NO.**|**PARAMETER**|**DESCRIPTION**|**MIN**<br>**MAX**|**UNIT**|
|---|---|---|---|---|
||f(baud)|Programmable baud rate|7.38|Mbps|
|3|tw(TXD)|Pulse width, transmit data bit high or low|U(1)- 2<br>U(1)+ 2|ns|
|4|tw(TXDS)|Pulse width, transmit start bit low|U(1)- 2|ns|



(1) U = UART baud time in ns = 1/actual baud rate, where the actual baud rate is defined in the UART Baud Rate Settings table of the device TRM. 

**==> picture [288 x 155] intentionally omitted <==**

**----- Start of picture text -----**<br>
2<br>1<br>Start<br>UARTi_RXD Bit VIH<br>VIL<br>Data Bits<br>4<br>3<br>Start<br>UARTi_TXD Bit<br>Data Bits<br>UART_TIMING_01_RCVRVIHVIL<br>**----- End of picture text -----**<br>


**Figure 6-113. UART Timing Requirements and Switching Characteristics** 

Copyright © 2025 Texas Instruments Incorporated 

197 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

For more information, see _Universal Asynchronous Receiver/Transmitter (UART)_ section in _Peripherals_ chapter in the device TRM. 

## **6.11.5.20 USB** 

The USB 2.0 subsystem is compliant with the Universal Serial Bus (USB) Specification, revision 2.0. Refer to the specification for timing details. 

For more details about features and additional description information on the device Universal Serial Bus Subsystem (USB), see the corresponding subsections within _Signal Descriptions_ and _Detailed Description_ sections. 

Copyright © 2025 Texas Instruments Incorporated 

198 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **7 Detailed Description** 

## **7.1 Overview** 

The low-cost & performance optimized AM62L family of application processors are built for Linux[®] application development. With scalable Arm[®] Cortex[®] -A53 core performance and embedded features such as: Multimedia DSI/DPI support, integrated ADC on chip, advanced lower power management modes, and extensive security options for IP protection with the built-in security features. 

The AM62Lx devices includes an extensive set of peripherals that make it a well-suited general-purpose device for a broad range of industrial applications while offering intelligent features and optimized power architecture as well. In addition, the extensive set of peripherals included in AM62Lx enables system-level connectivity, such as: USB, MMC/SD, OSPI, CAN-FD and an ADC. 

## **7.2 Processor Subsystem** 

## _**7.2.1 Arm Cortex-A53 Subsystem (A53SS)**_ 

The SoC implements one cluster of dual-core Arm® Cortex®-A53 MPCore[™] , with 32KB L1 instruction, 32KB L1 data, per core and 256KB L2 shared cache. 

The Cortex®-A53 cores are general-purpose processors that can be used for running customer applications. 

The A53SS is built around the Cortex®-A53 MPCore™ (Arm®-A53 Cluster), which is provided by Arm and configured by TI. It is based on the symmetric multiprocessor (SMP) architecture, and thus, it delivers high performance and optimal power management, debug and emulation capabilities. 

The A53 processor is a multi-issue out-of-order superscalar execution engine with integrated L1 Instruction and Data Caches, compatible with Arm®v8-A architecture. It delivers significantly more performance than its predecessors at a higher level of power efficiency. 

The Arm®v8-A architecture brings a number of new features. These include 64-bit data processing, extended virtual addressing and 64-bit general purpose registers. The A53 processor is Arm’s first Arm®v8-A processor aimed at providing power-efficient 64-bit processing. It features an in-order, 8-stage, dual-issue pipeline, and improved integer, Arm® Neon[™] , Floating-Point Unit (FPU) and memory performance. 

The A53 CPU supports two execution states: AArch32 and AArch64. The AArch64 state gives the A53 CPU its ability to execute 64-bit applications, while the AArch32 state allows the processor to execute existing Arm®v7-A applications. 

The A53SS integrates advanced features including Arm®v8 Cryptography Extensions, GICv3 architecture, ECC and parity protection for caches, dedicated watchdog timers per core, high-throughput 128-bit VBUSM interfaces, and a PBIST controller with BISOR for built-in self-test and reliability. 

For more information, see _Arm Cortex-A53 Subsystem_ section in _Processors and Accelerators_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 199 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **7.3 Other Subsystem** 

## _**7.3.1 Data Movement Subsystem (DMSS)**_ 

The DMSS module provides data movement (DMA) and bridges between crossbar module CBASS switched interconnect and the packet streaming fabric (network on chip) on the device. 

The Data Movement Subsystem (DMSS) consists of DMA/Queue Management components and Peripherals: 

- Packet DMA (PKTDMA) 

- Block Copy DMA (BCDMA) 

- Ring Accelerator 

- Packet Streaming Interface (PSILSS) 

- Infrastructure components such as CBASS 

For more information, see _Data Movement Architecture Overview_ section in _Peripherals_ chapter in the device TRM. 

## _**7.3.2 Peripheral DMA Controller (PDMA)**_ 

The Peripheral DMA is a simple DMA which has been architected to specifically meet the data transfer needs of peripherals, which perform data transfers using memory mapped registers (MMRs) accessed via a standard non-coherent bus fabric. The PDMA module is located close to one or more peripherals which require an external DMA for data movement. 

The PDMA is only responsible for performing the data movement transactions which interact with the peripherals themselves. Data which is read from a given peripheral is packed by a PDMA source channel into a PSI-L data stream which is then sent to a remote peer DMSS destination channel which then performs the movement of the data into memory. Likewise, a remote DMSS source channel fetches data from memory and transfers it to a peer PDMA destination channel over PSI-L which then performs the writes to the peripheral. 

The PDMA architecture is intentionally heterogeneous (DMSS + PDMA) to right size the data transfer complexity at each point in the system to match the requirements of whatever is being transferred to or from. Peripherals are typically FIFO based and do not require multi-dimensional transfers beyond their FIFO dimensioning requirements, so the PDMA transfer engines are kept simple with only a few dimensions (typically for sample size and FIFO depth), hardcoded address maps, and simple triggering capabilities. 

Multiple source and destination channels are provided within the PDMA which allow multiple simultaneous transfer operations to be ongoing. The DMA controller maintains state information for each of the channels and employs round-robin scheduling between channels in order to share the underlying DMA hardware. 

Each peripheral with PDMA support has its own dedicated state machine to track the transmit and receives of the data for each peripheral. 

For more information, see _Data Movement Architecture Overview_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

200 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **7.4 Peripherals** 

## _**7.4.1 ADC**_ 

The analog-to-digital converter (ADC) module provides a single-channel general purpose analog-to-digital conversion function from one of four analog inputs selected using the integrated 4-input analog multiplexer. The analog front end (AFE) implemented in the module performs a 12-bit conversion with 10 bits of effective resolution. 

For more information, see _Analog-to-Digital Converter (ADC)_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.2 Gigabit Ethernet Switch (CPSW3G)**_ 

The 3-port Gigabit Ethernet Switch (CPSW3G) subsystem provides Ethernet packet communication for the device and can be configured as an Ethernet switch. It supports two external 10/100/1000Mbps Ethernet ports with selectable RGMII and RMII interfaces and one internal Communications Port Programming Interface (CPPI) port. 

For more information, see _Gigabit Ethernet Switch_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.3 DDR Subsystem (DDRSS)**_ 

DDRSS0 supports LPDDR4 and DDR4 memory types up to 1600MT/s with a 16-bit bus and in-line ECC, addressing up to 2GB (LPDDR4) and 4GB (DDR4). It features a 128-bit system interface, advanced scheduling and refresh control, full command coherency, and JEDEC-compliant low-power modes for efficient and reliable operation across extended temperature ranges. 

For more information, see _DDR Subsystem_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.4 Display Subsystem (DSS)**_ 

The Display Subsystem (DSS) is a flexible, single-pipeline subsystem that supports high-resolution display output up to 1920x1080@60fps. DSS includes a DMA engine with video frame flip/mirror support, which allows direct access to the frame buffer (device system memory). The input pipeline supports color space conversion, gamma correction, and brightness/contrast hue/saturation control capabilities, among others, for enhanced video output quality. The DSS output either connects directly to device pins to provide a parallel 24-bit DPI video output interface with 150MHz pixel clock or connects to the MIPI DSI Controller, which provides video interface through the four-lane MIPI D-PHY Transmitter with data rate up to 2.5 Gbps/lane. 

For more information, see _Display Subsystem_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.5 Enhanced Capture (ECAP)**_ 

The Enhanced Capture (ECAP) module is a timing peripheral designed to accurately capture and measure external signal characteristics such as period, frequency, duty cycle, or pulse width. ECAP operates using a 32-bit time-stamp counter and up to four 32-bit capture registers. Captured values can be used to calculate timing intervals, generate interrupts, or trigger other peripherals. 

The module can generate interrupts on any of the capture events and supports both absolute time-capture and delta-time stamp capture modes, programmable edge polarity for each capture event, and can operate in auxiliary PWM (APWM) mode to generate a PWM output when not used for capture. The ECAP also supports one-shot capture mode of up to four time-stamp events, and continuous capture mode of time stamps in a four-deep circular buffer. 

These capabilities make the ECAP module useful for speed measurement, position sensing, and precise input signal monitoring control applications. 

For more information, see _Enhanced Capture (ECAP) Module_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.6 Error Location Module (ELM)**_ 

The Error Location Module (ELM) operates in conjunction with the General-Purpose Memory Controller (GPMC) to support error detection and correction for NAND flash memories. It processes syndrome polynomials generated during NAND page reads using a Bose–Chaudhuri–Hocquenghem (BCH) algorithm to identify error 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 201 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

locations within a data block. The ELM supports 4-, 8-, and 16-bit error correction per 512-byte block, with interrupt generation upon completion and register-based access to error count and location data. 

For more information, see _Error Location Module (ELM)_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.7 Enhanced Pulse Width Modulation (EPWM)**_ 

Enhanced Pulse Width Modulation (EPWM) module is a highly flexible timer-based peripheral used for generating precise pulse-width modulated waveforms for motor control, digital power, and general-purpose timing applications. 

The EPWM module provides programmable period, duty cycle, and phase control, with features such as deadband generation with independent rising- and falling-edge delay control, trip-zone inputs for fault handling, time-base synchronization input/output signals for synchronization with other EPWM modules, and ability to generate events to trigger CPU interrupts and ADC conversions, thus enabling precise synchronization between control loops and waveform generation. 

Additional capabilities include PWM chopping by a high-frequency carrier signal to reduce EMI and improve signal quality, and programmable event prescaling functionality for fine-grained control over how often PWM events trigger actions. 

For more information, see _Enhanced Pulse Width Modulation (EPWM)_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.8 Enhanced Quadrature Encoder Pulse (EQEP)**_ 

The Enhanced Quadrature Encoder Pulse (EQEP) peripheral is used to interface with quadrature-encoded signals from rotary or linear encoders, commonly used in high performance motion and position control systems, providing accurate position, direction, and speed information. 

The EQEP module supports decoding of A-phase and B-phase signals, along with index signal (QEPI) for absolute position reference. 

The 32-bit EQEP module features a position counter and control unit for position measurement with programmable reset, a quadrature edge-capture unit for low-speed measurement, a unit time base for real-time speed measurement, along with a watchdog timer for loss of encoder activity. The EQEP also generates interrupts on compare, overflow/underflow, and index events, supporting flexible motion control algorithms. 

For more information, see _Enhanced Quadrature Encoder Pulse (EQEP)_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.9 General-Purpose Interface (GPIO)**_ 

The General-Purpose Input/Output (GPIO) peripheral provides dedicated general-purpose pins that can be configured as either inputs or outputs. When configured as an output, the user can write to an internal register to control the state driven on the output pin. When configured as an input, user can obtain the state of the input by reading the state of an internal register. 

A GPIO module supports up to 144 dedicated signals distributed in 9 banks, each one comprising up to 16 GPIO signals. 

Interrupt generation can be enabled independently for each bank of 16 GPIO signals. Interrupts can be triggered by rising and/or falling edge, specified for each interrupt capable GPIO signal. 

In addition, the GPIO peripheral can produce DMA synchronization events in different event generation modes. GPIO signals set/clear functionality is also available. 

For more information, see _General-Purpose Interface_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.10 General-Purpose Memory Controller (GPMC)**_ 

The General-Purpose Memory Controller is a unified memory controller dedicated for interfacing with external memory devices like: 

Copyright © 2025 Texas Instruments Incorporated 

202 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

- Asynchronous SRAM-like memories and application-specific integrated circuit (ASIC) devices 

- Asynchronous, synchronous, and page mode (available only in non-multiplexed mode) burst NOR flash devices 

- NAND flash 

- Pseudo-SRAM devices 

For more information, see _General-Purpose Memory Controller_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.11 Global Timebase Counter (GTC)**_ 

The GTC module is a 64-bit free-running up-counter compliant with Arm®v8 system counter requirments, with no rollover over the device lifetime when the full 64-bit counter is used, and support of selectable counter-bit outputs as a push events. 

The GTC provides a unified time reference across all cores and peripherals for consistent timestamping and synchronization. 

For more information, see _Global Timebase Counter_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.12 Inter-Integrated Circuit (I2C)**_ 

The Inter-Integrated Circuit (I2C) controller provides an interface between a local host (LH), such as an Arm and any I[2] C-bus-compatible device that connects via the I[2] C serial bus. External components attached to the I[2] C bus can serially transmit and receive up to 8 bits of data to and from the LH device through the 2-wire I[2] C interface. 

Each multicontroller I[2] C module can be configured to act like a target or controller I[2] C-compatible device. 

The I[2] C instances are implemented with I[2] C compliant, open-drain I/O buffers, or with standard push-pull I/O buffers. The I[2] C instances associated with I[2] C open-drain I/O buffers can support HS-mode (up to 3.4Mbps when operating at 1.8V, and limited to 400kbps when operating at 3.3V). 

The I[2] C instances associated with standard push-pull I/O buffers can support Fast-mode (up to 400kbps). The standard push-pull I/O buffers being used on these ports are connected such they emulate open-drain outputs. This emulation is achieved by forcing a constant low output and disabling the output buffer to enter the Hi-Z state. 

For more information, see _Inter-Integrated Circuit_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.13 Modular Controller Area Network (MCAN)**_ 

The Controller Area Network (CAN) is a serial communications protocol which efficiently supports distributed real-time control with a high level of security. CAN has high immunity to electrical interference and the ability to self-diagnose and repair data errors. In a CAN network, many short messages are broadcast to the entire network, which provides for data consistency in every node of the system. 

The MCAN module supports both classic CAN and CAN FD (CAN with Flexible Data-Rate) specifications. CAN FD feature allows high throughput and increased payload per data frame. The classic CAN and CAN FD devices can coexist on the same network without any conflict. 

CAN and CAN FD devices connect to the physical layer of the CAN network through external (for the device) transceivers. Each MCAN module supports flexible bit rates greather than 1Mbps and is compliant to ISO 11898-1:2015. 

For more information, see _Modular Controller Area Network (MCAN)_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.14 Multichannel Audio Serial Port (MCASP)**_ 

The MCASP functions as a general-purpose audio serial port, optimized for the requirements of various audio applications. The MCASP module can operate in both transmit and receive modes. The MCASP is useful for time-division multiplexed (TDM) stream, Inter-IC Sound (I2S) protocols reception and transmission as well as 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 203 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

for an inter-component digital audio interface transmission (DIT). The MCASP has the flexibility to gluelessly connect to a Sony/Philips digital interface (S/PDIF) transmit physical layer component. 

Although inter-component digital audio interface reception (DIR) mode (this is, S/PDIF stream receiving) is not natively supported by the MCASP module, a specific TDM mode implementation for the MCASP receivers allows an easy connection to external DIR components (for example, S/PDIF to I2S format converters). 

For more information, see _Multichannel Audio Serial Port_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.15 Multichannel Serial Peripheral Interface (MCSPI)**_ 

The MCSPI is an enhanced SPI module that supports multichannel transmit/receive communication and can operate in both controller and peripheral modes. In controller mode the module can interface with up to four channels, while in peripheral mode supports a single channel. 

Each channel supports two independent DMA requests (one for read and one for write) and one interrupt for efficient data transfer, a programmable start-bit (Last Output Stop Start Insertion (LOSSI) mode) to ensure proper framing and synchronization in multi-channel communication, a built-in FIFO for data throughput and word access efficiency, and a serial clock with programmable frequency, polarity and phase. 

The MCSPI module supports configurable SPI word length in the range of 4 to 32 bits. Additionally, programmable shift operations and timing control between chip select and external clock generation are provided. 

For more information, see _Multichannel Serial Peripheral Interface_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.16 Multi-Media Card Secure Digital (MMCSD)**_ 

The MMCSD Host Controller provides an interface to embedded Multi-Media Card (MMC), Secure Digital (SD), and Secure Digital IO (SDIO) devices. The MMC/SD controller deals with MMC/SD/SDIO protocol at transmission level — packing data, adding CRC, start/end bit, and checking for syntactical correctness. 

The MMCSD Host Controller has been implemented as a 4-bit subsystem and 8-bit subsystem. The 4-bit subsystem supports removable SD cards compliant with the SD Physical Layer Specification v3.01, and embedded SDIO devices compliant with the SDIO Specification v3.00. The 8-bit subsystem supports eMMC devices compliant with the JEDEC eMMC electrical standard v5.1 (JESD84-B51), and embedded SDIO devices compliant with the SDIO Specification v3.00. 

For more information, see _Multi-Media Card Secure Digital (MMCSD) Interface_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.17 Octal Serial Peripheral Interface (OSPI)**_ 

The Octal Serial Peripheral Interface (OSPI) module is a Serial Peripheral Interface (SPI) module which allows single, dual, quad or octal read and write access to external flash devices with dual (DDR) or single (SDR) data rate. This module has a memory mapped register interface, which provides a direct memory interface for accessing data from external flash devices, simplifying software requirements. 

The module supports DDR and DTR protocols (including Octal DDR with DQS), XIP (continuous mode), programmable device sizes and delays, and write protection regions. Additional features include bidirectional CRC, ECC error handling, programmable interrupt generation, and a programmable data decoder for continuous addressing and device boundary detection. 

For more information, see _Octal Serial Peripheral Interface (OSPI)_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.18 Timers**_ 

The general-purpose timer (TIMER) is a 32-bit module that supports timer mode for periodic event generation, capture mode for precise timestamping of external events, and compare mode for match-based interrupts. TIMER modules support cascading two 32-bit timers to form a 64-bit counter. 

Copyright © 2025 Texas Instruments Incorporated 

204 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L www.ti.com** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

TIMER contains a free running upward counter with auto reload capability on overflow and can be read and written on the fly (while counting). TIMER supports interrupts generated on overflow, compare and capture events. All internal timer interrupt sources are merged in one module interrupt line and one wake-up line and each internal interrupt source can be independently enabled or disabled. 

TIMER modules can generate 1-ms tick with 32768-Hz functional clock. 

For more information, see _Timers_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.19 Real-Time Clock (RTC)**_ 

The Real-Time Clock (RTC) timer module keeps track of calendar time and date, wakes up the device from power-down, and supports Digital Rights Management (DRM). 

The RTC module includes a 48-bit seconds counter, a 15-bit 32768-Hz sub-second counter, and a 512-bit scratch pad storage. RTC can be used for system scheduling, low-power timekeeping, and maintaining accurate timestamps across device resets. 

For more information, see _Real-Time Clock (RTC)_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.20 Universal Asynchronous Receiver/Transmitter (UART)**_ 

The UART is a peripheral that utilizes the DMA for data transfer or interrupt polling via host CPU. All UART modules support IrDA and CIR modes when 48MHz function clock is used. Each UART can be used for configuration and data exchange with a number of external peripheral devices or interprocessor communication between devices. 

The UART module supports high-speed communication up to 3.6Mbps with a 64-byte FIFO buffer for both transmission and reception, and includes advanced features like auto flow control, configurable data formats, sleep mode, and extended modem control signals. It also offers programmable interrupt levels, auto-baud detection, and internal loopback for testing. 

For more information, see _Universal Synchronous/Asynchronous Receiver/Transmitter_ section in _Peripherals_ chapter in the device TRM. 

## _**7.4.21 Universal Serial Bus Subsystem (USBSS)**_ 

The Universal Serial Bus Subsystem (USBSS) provides a connectivity solution for numerous consumer portable devices by implementing a mechanism for data transfer between USB devices. 

USBSS features Dual Role Device (DRD) functionality, enabling operation in Host mode at High-Speed (480Mbps), Full-Speed (12Mbps), or Low-Speed (1.5Mbps), and Peripheral mode at High-Speed (480Mbps) or Full-Speed (12Mbps), offering flexible operation and integrated VBUS detection. The subsystem is compliant with the xHCI 1.1 specification for host controller interface compatibility. 

For more information, see _Universal Serial Bus Subsystem (USBSS)_ section in _Peripherals_ chapter in the device TRM. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 205 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **8 Applications, Implementation, and Layout** 

## **Note** 

Information in the following applications sections is not part of the TI component specification, and TI does not warrant its accuracy or completeness. TI’s customers are responsible for determining suitability of components for their purposes, as well as validating and testing their design implementation to confirm system functionality. 

## **8.1 Device Connection and Layout Fundamentals** 

## _**8.1.1 Power Supply**_ 

## **8.1.1.1 Power Supply Designs** 

The AM62Lx Power Supply Implementation application note, provides recommended power management solutions for the AM62Lx processor and its principal peripherals. 

## **8.1.1.2 Power Distribution Network Implementation Guidance** 

The Sitara Processor Power Distribution Networks: Implementation and Analysis provides guidance for successful implementation of the power distribution network. This includes PCB stackup guidance as well as guidance for optimizing the selection and placement of the decoupling capacitors. TI _only_ supports designs that follow the board design guidelines contained in the application report. 

## _**8.1.2 External Oscillator**_ 

For more information about External Oscillators, see the _Clock Specifications_ section. 

## _**8.1.3 JTAG, EMU, and TRACE**_ 

Texas Instruments supports a variety of eXtended Development System (XDS[™] ) JTAG controllers with various debug capabilities beyond only JTAG support. A summary of this information is available in the XDS Target Connection Guide. 

For recommendations on JTAG, EMU, and TRACE routing, see the Emulation and Trace Headers Technical Reference Manual 

## _**8.1.4 Unused Pins**_ 

For more information about Unused Pins, see Section 5.4, _Pin Connectivity Requirements_ . 

Copyright © 2025 Texas Instruments Incorporated 

206 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **8.2 Peripheral- and Interface-Specific Design Information** 

## _**8.2.1 DDR Board Design and Layout Guidelines**_ 

The goal of the AM62x, AM62Lx DDR Board Design and Layout Guidelines is to make the DDR system implementation straightforward for all designers. Requirements have been distilled down to a set of layout and routing rules that allow designers to successfully implement a robust design for the topologies that TI supports. TI only supports board designs using DDR4 or LPDDR4 memories that follow the guidelines in this document. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 207 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**8.2.2 OSPI/QSPI/SPI Board Design and Layout Guidelines**_ 

The following section details the PCB routing guidelines that must be observed when connecting OSPI, QSPI, or SPI devices. 

## **8.2.2.1 No Loopback, Internal PHY Loopback, and Internal Pad Loopback** 

- The OSPI[x]_CLK output pin must be connected to the CLK input pin of the attached OSPI/QSPI/SPI device 

- The signal propagation delay from the OSPI[x]_CLK pin to the attached OSPI/QSPI/SPI device CLK pin (A to B) must be ≤ 450ps (~7cm as stripline or ~8cm as microstrip) 

- The signal propagation delay of each OSPI[x]_D[y] and OSPI[x]_CSn[z] pin to the corresponding attached OSPI/QSPI/SPI device data and control pin (E to F, or F to E) must be approximately equal to the signal propagation delay from the OSPI[x]_CLK pin to the attached OSPI/QSPI/SPI device CLK pin (A to B) 

- 50Ω PCB routing is recommended along with series terminations, as shown in Figure 8-1 

- Propagation delays and matching: 

   - (A to B) ≤ 450ps 

   - (E to F, or F to E) = ((A to B) ± 60ps) 

**==> picture [332 x 342] intentionally omitted <==**

**----- Start of picture text -----**<br>
A B<br>R1<br>0 Ω*<br>OSPI/QSPI/SPI<br>OSPI[x]_CLK<br>Device Clock Input<br>OSPI[x]_LBCLKO<br>OSPI[x]_DQS OSPI Device DQS<br>E F<br>OSPI[x]_D[y], OSPI/QSPI/SPI<br>OSPI[x]_CSn[z] Device IO[y], CS#<br>OSPI_Board_01<br>**----- End of picture text -----**<br>


- 0Ω resistor (R1), located as close as possible to the OSPI[x]_CLK pin, is placeholder for fine tuning, if needed. 

## **Figure 8-1. OSPI Connectivity Schematic for No Loopback, Internal PHY Loopback, and Internal Pad Loopback** 

Copyright © 2025 Texas Instruments Incorporated 

208 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **8.2.2.2 External Board Loopback** 

- The OSPI[x]_CLK output pin must be connected to the CLK input pin of the attached OSPI/QSPI/SPI device 

- The OSPI[x]_LBCLKO output pin must be looped back to the OSPI[x]_DQS input pin 

- The signal propagation delay of the OSPI[x]_LBCLKO pin to the OSPI[x]_DQS pin (C to D) must be approximately twice the propagation delay of the OSPI[x]_CLK pin to the attached OSPI/QSPI/SPI device CLK pin (A to B) 

- The signal propagation delay of each OSPI[x]_D[y] and OSPI[x]_CSn[z] pin to the corresponding attached OSPI/QSPI/SPI device data and control pin (E to F, or F to E) must be approximately equal to the signal propagation delay from the OSPI[x]_CLK pin to the attached OSPI/QSPI/SPI device CLK pin (A to B) 

- 50Ω PCB routing is recommended along with series terminations, as shown in Figure 8-2 

- Propagation delays and matching: 

   - (C to D) = 2 x ((A to B) ± 30ps), see the exception note below. 

   - (E to F, or F to E) = ((A to B) ± 60ps) 

## **Note** 

The External Board Loopback hold time requirement (defined by parameter number O16 in the _OSPI0 Timing Requirements - PHY DDR Mode_ section) may be larger than the hold time provided by a typical OSPI/QSPI/SPI device. In this case, the propagation delay of OPSI[x]_LBCLKO pin to the OSPI[x]_DQS pin (C to D) can be reduced to provide additional hold time. 

**==> picture [332 x 342] intentionally omitted <==**

**----- Start of picture text -----**<br>
A B<br>R1<br>0 Ω*<br>OSPI/QSPI/SPI<br>OSPI[x]_CLK<br>Device Clock Input<br>C<br>R1<br>0 Ω*<br>OSPI[x]_LBCLKO<br>D<br>OSPI[x]_DQS OSPI Device DQS<br>E F<br>OSPI[x]_D[y], OSPI/QSPI/SPI<br>OSPI[x]_CSn[z] Device IO[y], CS#<br>OSPI_Board_02<br>**----- End of picture text -----**<br>


* 0Ω resistor (R1), located as close as possible to the OSPI[x]_CLK and OSPI[x]_LBCLKO pins, is a placeholder for fine tuning, if needed. 

## **Figure 8-2. OSPI Connectivity Schematic for External Board Loopback** 

Copyright © 2025 Texas Instruments Incorporated 

209 

_Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **8.2.2.3 DQS (only available in Octal SPI devices)** 

- The OSPI[x]_CLK output pin must be connected to the CLK input pin of the attached OSPI/QSPI/SPI device 

- The DQS pin of the attached OSPI/QSPI/SPI device must be connected to OSPI[x]_DQS pin 

- The signal propagation delay from the attached OSPI/QSPI/SPI device DQS pin to the OSPI[x]_DQS pin (D to C) must be approximately equal to the signal propagation delay from the OSPI[x]_CLK pin to the attached OSPI/QSPI/SPI device CLK pin (A to B) 

- The signal propagation delay of each OSPI[x]_D[y] and OSPI[x]_CSn[z] pin to the corresponding attached OSPI/QSPI/SPI device data and control pin (E to F, or F to E) must be approximately equal to the signal propagation delay from the OSPI[x]_CLK pin to the attached OSPI/QSPI/SPI device CLK pin (A to B) 

- 50Ω PCB routing is recommended along with series terminations, as shown in Figure 8-3 

- Propagation delays and matching: 

   - (D to C) = ((A to B) ± 30ps) 

   - (E to F, or F to E) = ((A to B) ± 60ps) 

**==> picture [332 x 342] intentionally omitted <==**

**----- Start of picture text -----**<br>
A B<br>R1<br>0 Ω*<br>OSPI/QSPI/SPI<br>OSPI[x]_CLK<br>Device Clock Input<br>OSPI[x]_LBCLKO<br>C D<br>OSPI[x]_DQS OSPI Device DQS<br>E F<br>OSPI[x]_D[y], OSPI/QSPI/SPI<br>OSPI[x]_CSn[z] Device IO[y], CS#<br>OSPI_Board_03<br>**----- End of picture text -----**<br>


* 0Ω resistor (R1), located as close as possible to the OSPI[x]_CLK pin, is a placeholder for fine tuning, if needed. 

**Figure 8-3. OSPI Connectivity Schematic for DQS** 

Copyright © 2025 Texas Instruments Incorporated 

210 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**8.2.3 USB VBUS Design Guidelines**_ 

The USB 3.1 specification allows the VBUS voltage to be as high as 5.5V for normal operation, and as high as 20V when the Power Delivery addendum is supported. Some automotive applications require a max voltage to be 30V. 

The device requires the VBUS signal voltage be scaled down using an external resistor divider (as shown in the Figure 8-4), which limits the voltage applied to the actual device pin (USB0_VBUS). The tolerance of these external resistors should be equal to or less than 1%, and the leakage current of Zener diode at 5V should be less than 100nA. 

**==> picture [263 x 165] intentionally omitted <==**

**----- Start of picture text -----**<br>
Device<br>USBn_VBUS<br>16.5 kΩ 3.48 kΩ<br>±1% ±1%<br>VBUS signal<br>10 kΩ<br>6.8V<br>±1%<br>(BZX84C6V8 or equivalent)<br>VSS VSS<br>J7ES_USB_VBUS_01<br>**----- End of picture text -----**<br>


**Figure 8-4. USB VBUS Detect Voltage Divider / Clamp Circuit** 

The USB0_VBUS pin can be considered to be fail-safe because the external circuit in Figure 8-4 limits the input current to the actual device pin in a case where VBUS is applied while the device is powered off. 

## _**8.2.4 High Speed Differential Signal Routing Guidance**_ 

The High Speed Interface Layout Guidelines provides guidance for successful routing of the high speed differential signals. This includes PCB stackup and materials guidance as well as routing skew, length and spacing limits. TI supports _only_ designs that follow the board design guidelines contained in the application note. 

## _**8.2.5 Thermal Solution Guidance**_ 

The Thermal Design Guide for DSP and ARM Application Processors provides guidance for successful implementation of a thermal solution for system designs containing this device. This document provides background information on common terms and methods related to thermal solutions. TI only supports designs that follow system design guidelines contained in the application note. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 211 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **8.3 Clock Routing Guidelines** 

## _**8.3.1 Oscillator Routing**_ 

When designing the printed-circuit board: 

- Place all crystal circuit components as close as possible to the respective device pins. 

- Route the crystal circuit traces on the outer layer of the PCB and minimize trace lengths to reduce parasitic capacitance and minimize crosstalk from other signals. 

- Place a continuous ground plane on the adjacent layer of the PCB such that it is under all crystal circuit components and crystal circuit traces. 

- Route a ground guard around the crystal circuit components to shield it from any adjacent signals routed on the same layer as the crystal circuit traces. Insert multiple vias to stitch down the ground guard such that it does not have any unterminated stubs. 

- Route a ground guard between the WKUP_OSC0_XI and WKUP_OSC0_XO signals to shield the WKUP_OSC0_XI signal from the WKUP_OSC0_X0 signal. Insert multiple vias to stitch down the ground guard such that it does not have any unterminated stubs. 

- Connect all crystal circuit ground connections and ground guard connections directly to the adjacent layer ground plane, and the device VSS ground plane if they are implemented separately on different layers of the PCB. 

## **Note** 

Implementing a ground guard between the WKUP_OSC0_XI and WKUP_OSC0_XO signals is critical to minimize shunt capacitance between the two signals. Routing these two signals adjacent to each other without a ground guard between them will effectively reduce the gain of the oscillator amplifier, which reduces its ability to start oscillation. 

**==> picture [378 x 294] intentionally omitted <==**

**----- Start of picture text -----**<br>
GND vias<br>Device<br>WKUP_OSC0_XI<br>GND plane<br>GND guard<br>WKUP_OSC0_XO<br>GND vias<br>Cap<br>Crystal<br>Cap<br>**----- End of picture text -----**<br>


**Figure 8-5. WKUP_OSC0 PCB requirements** 

Copyright © 2025 Texas Instruments Incorporated 

212 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** 

**www.ti.com** 

SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

## **9 Device and Documentation Support** 

## **9.1 Device Nomenclature** 

To designate the stages in the product development cycle, TI assigns prefixes to the part numbers of all microprocessors (MPUs) and support tools. Each device has one of three prefixes: X, P, or null (no prefix) (for example, XAM62L32AOGHAANB). Texas Instruments recommends two of three possible prefix designators for its support tools: TMDX and TMDS. These prefixes represent evolutionary stages of product development from engineering prototypes (TMDX) through fully qualified production devices and tools (TMDS). 

Device development evolutionary flow: 

- **X** Experimental device that is not necessarily representative of the final device's electrical specifications and may not use production assembly flow. 

- **P** Prototype device that is not necessarily the final silicon die and may not necessarily meet final electrical specifications. 

**null** Production version of the silicon die that is fully qualified. 

Support tool development evolutionary flow: 

**TMDX** Development-support product that has not yet completed Texas Instruments internal qualification testing. **TMDS** Fully-qualified development-support product. 

X and P devices and TMDX development-support tools are shipped against the following disclaimer: 

"Developmental product is intended for internal evaluation purposes." 

Production devices and TMDS development-support tools have been characterized fully, and the quality and reliability of the device have been demonstrated fully. TI's standard warranty applies. 

Predictions show that prototype devices (X or P) have a greater failure rate than the standard production devices. Texas Instruments recommends that these devices not be used in any production system because their expected end-use failure rate still is undefined. Only qualified production devices are to be used. 

For orderable part numbers of AM62Lx devices in the ANB package type, see the Package Option Addendum of this document, the TI website (ti.com), or contact your TI sales representative. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 213 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**9.1.1 Standard Package Symbolization**_ 

## **Note** 

Some devices may have a cosmetic circular marking visible on the top of the device package which results from the production test process. In addition, some devices may also show a color variation in the package substrate which results from the substrate manufacturer. These differences are cosmetic only with no reliability impact. 

**==> picture [310 x 204] intentionally omitted <==**

**----- Start of picture text -----**<br>
TI  aBBBBB<br>BBrZfYt  Q1<br>XXXXXXX<br>YYY  PPP<br>A1<br>(PIN 1 INDICATOR)<br>**----- End of picture text -----**<br>


**Figure 9-1. Printed Device Reference** 

Copyright © 2025 Texas Instruments Incorporated 

214 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## _**9.1.2 Device Naming Convention**_ 

**Table 9-1. Nomenclature Description** 

|**FIELD PARAMETER**|**FIELD DESCRIPTION**|**VALUE**|**DESCRIPTION**|
|---|---|---|---|
|TI|Device<br>Manufacturer|TI|Texas Instruments|
|a|Device evolution<br>stage(1)|X|Prototype|
|||P|Preproduction (production test flow, no reliability data)|
|||BLANK<br>(null)|Production|
|BBBBBBB|Base production part<br>number|AM62L32|seeDevice Comparison|
|||AM62L31||
|r|Device revision|A|SR1.0|
|||B|SR1.1|
|Z|Device Speed Grade|E|SeeDevice Speed Gradestable|
|||O||
|f|Features<br>(seeDevice<br>Comparison)|G|Base|
|Y|Security / Functional<br>Safety|1 to 9|Secure with Dummy Key / No Functional Safety|
|||H to R|Secure with Production Key / No Functional Safety|
|||S to Z|Secure with Production Key / Functional Safety|
|t|Temperature(2)|A|–40°C to 105°C - Extended Industrial (seeRecommended Operating<br>Conditions)|
|||I|–40°C to 125°C - 125°C Industrial (seeRecommended Operating<br>Conditions)|
|Q1|Automotive Designator|Q1|Auto Qualified (AEC - Q100)|
|||BLANK|Standard|
||2D Barcode|Varies|Optional 2D barcode, provides additional device information|
|||BLANK||
|XXXXXXX||Lot Trace Code (LTC)||
|YYY||Production Code, For TI use only||
|PPP|Package Designator|ANB|FCCSP BGA (373)|
|●||Pin one designator||



(1) To designate the stages in the product development cycle, TI assigns prefixes to the part numbers. These prefixes represent evolutionary stages of product development from engineering prototypes through fully qualified production devices. Prototype devices are shipped against the following disclaimer: 

“This product is still in development and is intended for internal evaluation purposes.” 

Notwithstanding any provision to the contrary, TI makes no warranty expressed, implied, or statutory, including any implied warranty of merchantability of fitness for a specific purpose, of this device. 

(2) Applies to device max junction temperature. 

## **Note** 

BLANK in the symbol or part number is collapsed so there are no gaps between characters. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 215 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **9.2 Tools and Software** 

The following Development Tools support development for TI's Embedded Processing platforms: 

## **Development Tools** 

**Code Composer Studio[™] Integrated Development Environment** Code Composer Studio (CCS) Integrated Development Environment (IDE) is a development environment that supports TI's Microcontroller and Embedded Processors portfolio. Code Composer Studio comprises a suite of tools used to develop and debug embedded applications. The tool includes an optimizing C/C++ compiler, source code editor, project build environment, debugger, profiler, and many other features. The intuitive IDE provides a single user interface taking you through each step of the application development flow. Familiar tools and interfaces allow users to get started faster than ever before. Code Composer Studio combines the advantages of the Eclipse[®] software framework with advanced embedded debug capabilities from TI resulting in a compelling feature-rich development environment for embedded developers. 

**SysConfig Tool** The System Configuration tool provides a graphical user interface (GUI) that simplifies device configuration. The tool is designed to simplify hardware and software configuration challenges to accelerate software development. SysConfig is available as part of the Code Composer Studio™ integrated development environment as well as a standalone application. Additionally SysConfig can be run in the cloud by visiting the **TI developer zone** . 

SysConfig allows developers to configure pins, peripherals, and other components, and automatically detects, exposes, and resolves conflicts to speed software development. In addition, the clock tree tool provides a visual implementation of the device clock connectivity. 

The SysConfig tool generates output C header/code files that can be imported into software development kits (SDKs), enabling customers to configure their software in alignment with the specific hardware requirements. 

For a complete listing of development-support tools for the processor platform, visit the Texas Instruments website at ti.com. For information on pricing and availability, contact the nearest TI field sales office or authorized distributor. 

## **9.3 Documentation Support** 

To receive notification of documentation updates, navigate to the device product folder on ti.com. Click on _Notifications_ to register and receive a weekly digest of any product information that has changed. For change details, review the revision history included in any revised document. 

The following documents describe the AM62Lx devices. 

## **Technical Reference Manual** 

**AM62Lx Sitara™ Processors Technical Reference Manual** : Details the integration, the environment, the functional description, and the programming models for each peripheral and subsystem in the AM62Lx family of devices. 

## **Errata** 

**AM62Lx Sitara™ Processors Silicon Errata, Silicon Revision 1.0** : Describes the known exceptions to the functional specifications for the device. 

## **9.4 Support Resources** 

TI E2E[™] support forums are an engineer's go-to source for fast, verified answers and design help — straight from the experts. Search existing answers or ask your own question to get the quick design help you need. 

Linked content is provided "AS IS" by the respective contributors. They do not constitute TI specifications and do not necessarily reflect TI's views; see TI's Terms of Use. 

## **9.5 Trademarks** 

Sitara[™] , XDS[™] , Code Composer Studio[™] , and TI E2E[™] are trademarks of Texas Instruments. MPCore[™] and Neon[™] are trademarks of Arm Limited (or its subsidiaries) in the US and/or elsewhere. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 

216 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

Arm[®] , Cortex[®] , and TrustZone[®] are registered trademarks of Arm Limited (or its subsidiaries) in the US and/or elsewhere. 

MIPI[®] is a registered trademark of MIPI Alliance, Inc. 

Secure Digital[®] and SD[®] are registered trademarks of SD Card Association. 

Linux[®] is a registered trademark of Linus Torvalds. 

EtherCAT[®] is a registered trademark of Beckhoff Automation GmbH. PCIe[®] is a registered trademark of PCI-SIG. 

Eclipse[®] is a registered trademark of Eclipse Foundation AISBL. 

All trademarks are the property of their respective owners. 

## **9.6 Electrostatic Discharge Caution** 

**==> picture [30 x 27] intentionally omitted <==**

This integrated circuit can be damaged by ESD. Texas Instruments recommends that all integrated circuits be handled with appropriate precautions. Failure to observe proper handling and installation procedures can cause damage. 

ESD damage can range from subtle performance degradation to complete device failure. Precision integrated circuits may be more susceptible to damage because very small parametric changes could cause the device not to meet its published specifications. 

## **9.7 Glossary** 

TI Glossary This glossary lists and explains terms, acronyms, and definitions. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 217 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **Revision History** 

## **Changes from September 30, 2025 to November 30, 2025 (from Revision A (SEPTEMBER 2025) to Revision B (NOVEMBER 2025))** 

- **to Revision B (NOVEMBER 2025)) Page** • **Global:** GPMC timing updates were being applied and only partially completed when revision A of this datasheet was released. Therefore, the associated Revision History entries in revision A of this datasheet did not correctly describe the GPMC timing changes. The GPMC timing updates were completed and the Revision History entries updated to correctly reflect the GPMC timing changes relative to the initial datasheet release................................................................................................................................................................ 1 

- • (Feature List - Cryptographic acceleration): Removed SM3 and SM4 cryptographic cores since they are not included in this device. Also changed PKE (Public Key Engine) to PKA (Public Key Accelerator).................... 1 

- • (ECAP – Timing Requirements and Switching Characteristics): Updated the clock source referenced in table note 1..............................................................................................................................................................124 

- • (EPWM – Timing Requirements and Switching Characteristics): Updated the clock source referenced in table note 1..............................................................................................................................................................127 

- • (EQEP – Timing Requirements): Updated the clock source referenced in table note 1.................................129 • (GPMC and NOR Flash Timing Requirements — Synchronous Mode): Removed the GPMC_FCLK=100MHz column timing values and the associated not_div_by_1_mode timing values for GPMC_FCLK=133MHz. Simplified several parameter descriptions. Also removed two table notes, one that described register configuration for GPMC_FCLK selection, and another that described register configuration for div_by_1_mode.............................................................................................................................................. 131 

- • (GPMC and NOR Flash Switching Characteristics – Synchronous Mode): Removed the GPMC_FCLK=100MHz column timing values and the associated not_div_by_1_mode timing values for GPMC_FCLK=133MHz. Simplified several parameter descriptions. Changed the timing variable in parameters F3 and F11 to "D". Removed the "J" timing variable from the F15 and F17 parameters. Updated the table notes................................................................................................................................................ 131 

- • (GPMC and NOR Flash Timing Requirements – Asynchronous Mode): Removed the MODE column and the table note that described register configuration for div_by_1_mode. Added the correct table note for parameter FA21 .............................................................................................................................................139 

- • (GPMC and NOR Flash Switching Characteristics – Asynchronous Mode): Removed the MODE column and redundant rows. Also removed the table note that described register configuration for div_by_1_mode...... 139 

- • (GPMC and NAND Flash Timing Requirements – Asynchronous Mode): Removed the MODE column and the table note that described register configuration for div_by_1_mode.............................................................. 146 

- • (GPMC and NAND Flash Switching Characteristics – Asynchronous Mode): Removed the MODE column and the table note that described register configuration for div_by_1_mode. Added table notes and associated reference links for timing variables B, C, D, E, F, G, H, I, K, L, and M........................................................... 146 

- • (Detailed Description – DMSS): Removed secure proxy and an interrupt aggregator features..................... 200 

Copyright © 2025 Texas Instruments Incorporated 

218 _Submit Document Feedback_ 

Product Folder Links: _AM62L_ 

**AM62L** SPRSPA1B – MARCH 2025 – REVISED NOVEMBER 2025 

**www.ti.com** 

## **10 Mechanical, Packaging, and Orderable Information 10.1 Packaging Information** 

The following pages include mechanical, packaging, and orderable information. This information is the most current data available for the designated devices. This data is subject to change without notice and revision of this document. For browser-based versions of this data sheet, refer to the left-hand navigation. 

Copyright © 2025 Texas Instruments Incorporated 

_Submit Document Feedback_ 219 

Product Folder Links: _AM62L_ 

**PACKAGE OPTION ADDENDUM** 

24-Feb-2026 

www.ti.com 

## **PACKAGING INFORMATION** 

|**Orderable part number**|**Status**|**Material type**|**Package | Pins**|**Package qty | Carrier**|**RoHS**|**Lead finish/**|**MSL rating/**|**Op temp (°C)**|**Part marking**|
|---|---|---|---|---|---|---|---|---|---|
||(1)|(2)|||(3)|**Ball material**|**Peak reflow**||(6)|
|||||||(4)|(5)|||
|AM62L31BEGHAANBR|Active|Production|FCCSP (ANB) | 373|1000 | LARGE T&R|Yes|Call TI|Level-3-260C-168 HR|-40 to 125|31BEGHA|
||||||||||412|
|AM62L31BOGHAANBR|Active|Production|FCCSP (ANB) | 373|1000 | LARGE T&R|Yes|Call TI|Level-3-260C-168 HR|-40 to 125|31BOGHA|
||||||||||412|
|AM62L32BEGHAANBR|Active|Production|FCCSP (ANB) | 373|1000 | LARGE T&R|Yes|Call TI|Level-3-260C-168 HR|-40 to 125|32BEGHA|
||||||||||412|
|AM62L32BOGHAANBR|Active|Production|FCCSP (ANB) | 373|1000 | LARGE T&R|Yes|Call TI|Level-3-260C-168 HR|-40 to 125|32BOGHA|
||||||||||412|
|AM62L32BOGHIANBR|Active|Production|FCCSP (ANB) | 373|1000 | LARGE T&R|Yes|Call TI|Level-3-260C-168 HR|-40 to 125|32BOGHI|
||||||||||412|



> **(1) Status:** For more details on status, see our product life cycle. 

> **(2) Material type:** When designated, preproduction parts are prototypes/experimental devices, and are not yet approved or released for full production. Testing and final process, including without limitation quality assurance, reliability performance testing, and/or process qualification, may not yet be complete, and this item is subject to further changes or possible discontinuation. If available for ordering, purchases will be subject to an additional waiver at checkout, and are intended for early internal evaluation purposes only. These items are sold without warranties of any kind. 

> **(3) RoHS values:** Yes, No, RoHS Exempt. See the TI RoHS Statement for additional information and value definition. 

> **(4) Lead finish/Ball material:** Parts may have multiple material finish options. Finish options are separated by a vertical ruled line. Lead finish/Ball material values may wrap to two lines if the finish value exceeds the maximum column width. 

> **(5) MSL rating/Peak reflow:** The moisture sensitivity level ratings and peak solder (reflow) temperatures. In the event that a part has multiple moisture sensitivity ratings, only the lowest level per JEDEC standards is shown. Refer to the shipping label for the actual reflow temperature that will be used to mount the part to the printed circuit board. 

> **(6) Part marking:** There may be an additional marking, which relates to the logo, the lot trace code information, or the environmental category of the part. 

Multiple part markings will be inside parentheses. Only one part marking contained in parentheses and separated by a "~" will appear on a part. If a line is indented then it is a continuation of the previous line and the two combined represent the entire part marking for that device. 

**Important Information and Disclaimer:** The information provided on this page represents TI's knowledge and belief as of the date that it is provided. TI bases its knowledge and belief on information provided by third parties, and makes no representation or warranty as to the accuracy of such information. Efforts are underway to better integrate information from third parties. TI has taken and continues to take reasonable steps to provide representative and accurate information but may not have conducted destructive testing or chemical analysis on incoming materials and chemicals. TI and TI suppliers consider certain information to be proprietary, and thus CAS numbers and other limited information may not be available for release. 

Addendum-Page 1 

**PACKAGE OPTION ADDENDUM** 

www.ti.com 

24-Feb-2026 

In no event shall TI's liability arising out of such information exceed the total purchase price of the TI part(s) at issue in this document sold by TI to Customer on an annual basis. 

Addendum-Page 2 

## **ANB0373A** 

**==> picture [72 x 44] intentionally omitted <==**

## **PACKAGE OUTLINE** 

## **FCCSP - 0.877 mm max height** 

**==> picture [506 x 3] intentionally omitted <==**

**----- Start of picture text -----**<br>
SCALE  1.500<br>**----- End of picture text -----**<br>


**==> picture [109 x 7] intentionally omitted <==**

**----- Start of picture text -----**<br>
PLASTIC BALL GRID ARRAY<br>**----- End of picture text -----**<br>


**==> picture [434 x 554] intentionally omitted <==**

**----- Start of picture text -----**<br>
B 12.0 A<br>11.8<br>BALL A1<br>CORNER<br>12.0<br>11.8<br>0.1 C<br>0.877 MAX 0.2 C<br>C<br>SEATING PLANE<br>0.262 0.08 C<br>0.162 11  TYP<br>(0.45)<br>SYMM<br>AC<br>AB<br>AA (0.45)<br>Y<br>W<br>V<br>U<br>T<br>R<br>P<br>N SYMM<br>11  TYP M<br>L<br>K<br>J<br>H<br>G<br>F<br>E<br>D 0.35<br>C 373X  0.25<br>B 0.15 C A B<br>A<br>0.05 C<br>1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23<br>0.5  TYP<br>0.5  TYP<br>4230557/A   02/2024<br>**----- End of picture text -----**<br>


NOTES: 

1. All linear dimensions are in millimeters. Any dimensions in parenthesis are for reference only. Dimensioning and tolerancing per ASME Y14.5M. 

2. This drawing is subject to change without notice. 

www.ti.com 

## **EXAMPLE BOARD LAYOUT** 

## **ANB0373A** 

## **FCCSP - 0.877 mm max height** 

PLASTIC BALL GRID ARRAY 

**==> picture [366 x 307] intentionally omitted <==**

**----- Start of picture text -----**<br>
(0.5) TYP<br>373X ( 0.3)<br>1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23<br>(0.5) TYP A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>J<br>K<br>L SYMM<br>M<br>N<br>P<br>R<br>T<br>U<br>V<br>W<br>Y<br>AA<br>AB<br>AC<br>SYMM<br>**----- End of picture text -----**<br>


**==> picture [452 x 179] intentionally omitted <==**

**----- Start of picture text -----**<br>
LAND PATTERN EXAMPLE<br>EXPOSED METAL SHOWN<br>SCALE: 8X<br>0.05 MAX 0.05 MIN<br>ALL AROUND ALL AROUND M ETAL UNDER<br>E XPOSED METAL SOLDER MASK<br>SOLDER MASKOPENING METAL EDGE( 0.3) EXPOSED METAL SOLDER MASK( 0.3)<br>OPENING<br>NON-SOLDER MASK SOLDER MASK<br>DEFINED DEFINED<br>(PREFERRED)<br>SOLDER MASK DETAILS<br>NOT TO SCALE<br>4230557/A   02/2024<br>**----- End of picture text -----**<br>


NOTES: (continued) 

3. Final dimensions may vary due to manufacturing tolerance considerations and also routing constraints. For information, see Texas Instruments literature number SPRAA99 (www.ti.com/lit/spraa99). 

www.ti.com 

**FCCSP - 0.877 mm max height** 

## **EXAMPLE STENCIL DESIGN** 

## **ANB0373A** 

PLASTIC BALL GRID ARRAY 

**==> picture [363 x 289] intentionally omitted <==**

**----- Start of picture text -----**<br>
(0.5) TYP<br>373X ( 0.3)<br>1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23<br>(0.5) TYP A<br>B<br>C<br>D<br>E<br>F<br>G<br>H<br>J<br>K<br>L SYMM<br>M<br>N<br>P<br>R<br>T<br>U<br>V<br>W<br>Y<br>AA<br>AB<br>AC<br>**----- End of picture text -----**<br>


SYMM 

## SOLDER PASTE EXAMPLE 

BASED ON 0.100 mm THICK STENCIL SCALE: 8X 

**==> picture [74 x 6] intentionally omitted <==**

**----- Start of picture text -----**<br>
4230557/A   02/2024<br>**----- End of picture text -----**<br>


NOTES: (continued) 

4. Laser cutting apertures with trapezoidal walls and rounded corners may offer better paste release. 

www.ti.com 

## **IMPORTANT NOTICE AND DISCLAIMER** 

TI PROVIDES TECHNICAL AND RELIABILITY DATA (INCLUDING DATASHEETS), DESIGN RESOURCES (INCLUDING REFERENCE DESIGNS), APPLICATION OR OTHER DESIGN ADVICE, WEB TOOLS, SAFETY INFORMATION, AND OTHER RESOURCES “AS IS” AND WITH ALL FAULTS, AND DISCLAIMS ALL WARRANTIES, EXPRESS AND IMPLIED, INCLUDING WITHOUT LIMITATION ANY IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE OR NON-INFRINGEMENT OF THIRD PARTY INTELLECTUAL PROPERTY RIGHTS. 

These resources are intended for skilled developers designing with TI products. You are solely responsible for (1) selecting the appropriate TI products for your application, (2) designing, validating and testing your application, and (3) ensuring your application meets applicable standards, and any other safety, security, regulatory or other requirements. 

These resources are subject to change without notice. TI grants you permission to use these resources only for development of an application that uses the TI products described in the resource. Other reproduction and display of these resources is prohibited. No license is granted to any other TI intellectual property right or to any third party intellectual property right. TI disclaims responsibility for, and you fully indemnify TI and its representatives against any claims, damages, costs, losses, and liabilities arising out of your use of these resources. 

TI’s products are provided subject to TI’s Terms of Sale, TI’s General Quality Guidelines, or other applicable terms available either on ti.com or provided in conjunction with such TI products. TI’s provision of these resources does not expand or otherwise alter TI’s applicable warranties or warranty disclaimers for TI products. Unless TI explicitly designates a product as custom or customer-specified, TI products are standard, catalog, general purpose devices. 

TI objects to and rejects any additional or different terms you may propose. 

Copyright © 2026, Texas Instruments Incorporated 

Last updated 10/2025 


# Architecture

This document provides a quick reference around the technical architecture of BadgeSnake.

## Transport

I2C will be used to transport data between BeagleBadge's QWIIC connectors and up to 2 BeagleConnect Zepto's QWIIC connectors for both flashing and game play. The BOOT/RESET buttons on BeagleConnect Zepto will be utilize to enter into the MSPM0 BSL boot mode to perform flashing.

## BeagleBadge

There are 2 QWIIC connectors on BeagleBadge and each is connected to a different I2C bus on the AM62L32 processor:

* I2C1: badge connector J6, SDA on AM62L pin A6, SCL on AM62L pin D7
* I2C2: badge connector J7, SDA on AM62L pin D8, SCL on AM62L pin B8

Local documents:

* sprujb4.pdf: AM62L Technical Reference Manual
* am62l.pdf: AM62L Datasheet

Submodules:

* ../components/beaglebadge: board design repository

## BeagleConnect Zepto

Submodules:

* ../components/beagleconnect-zepto: board design repository

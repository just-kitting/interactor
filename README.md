# interactor

## What is it?

A kit for building a setup for seeing several students code interacting in real-time, providing physical presentation, but minimizing complex physical world factors.

Provides environment for:
* Enabling students to programming microcontroller boards (think BeagleConnect Zepto) using graphical or textual programming languages (MicroBlocks, Arduino/C, Micropython, etc.)
* A "game board" made up of a "controller" computer (think PocketBeagle 2) with status elements like LED matrix, audio output, LCD, etc. and I/O connections to the microcontroller boards
* Connecting multiple microcontroller boards that students have programmed to the "game board" to have them interact in real-time (over QWIIC or mikroBUS)

## Background

Beagle has recently launched a $1 development board with 128KB flash and 16KB RAM about the size of a postage stamp and compatible with 1,500 mikroBUS add-on boards and 100s of QWIIC add-on boards called "BeagleConnect Zepto". There is an on-line competitive coding environment called "BattleSnake" (https://play.battlesnake.com/) that pits multiple snakes against each other where "code is the controller". I think putting your code on a Zepto board and bringing it to a system that will "battle it out" between multiple users where they cannot control it themselves, but instead have to have written code to handle the various conditions could be very inspirational.

Of course, I think there is a major personality bias when it comes to creating something that is purely competitive, but this is just to draw a picture for you. The objective could also be collaborative, like creating music or drawing a picture. I have specific implementations in mind, but want to get the conversation going before I get too far.

## Silly AI videos

![](https://zulip.openbeagle.org/user_uploads/2/KvPeLYgGIGEzYuWBiQ13W3V4/The_Classroom_Coding_Future.mp4)

![](https://zulip.openbeagle.org/user_uploads/2/qnlPfqOdxg1fEDS_EjZ5p5Hz/The_Phases_of_Play.mp4)
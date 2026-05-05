# BadgeSnake demo

This is a technology demonstration of BeagleBadge + BeagleConnect Zepto playing BattleSnake.

```
The objective of this demo is to show
   how a game
    normally be played by people
     performing interactions with a computer program
  can be turned into 
    a game played multiple computers
  observed by people
    with each computer programmed by different people.
```

By delivering this demonstration, the pedagogical objectives of "The Interactor" can be better explored through comparison and contrast statements,
rather than explaining what it looks like for game controllers to be replaced by computers.

## Planning documents

- `AGENTS.md` defines the execution strategy for future work in this repo.
- `TODO.md` tracks the staged build-out checklist for BadgeSnake.
- `PROGRESS.md` records confirmed board and repository findings over time.
- `docs/J7ToJ6TargetModeValidation.md` records the current best target-mode validation path on BeagleBadge.
- `docs/J7ToJ6TestunitFeatures.md` records richer `slave-testunit` feature checks over J7 -> J6.

## Technology bits

### What is [Battlesnake](https://play.battlesnake.com)?

Battlesnake is a competitive programming game where "code is the controller". Instead of manually operating a character with a joystick or keyboard, players build autonomous algorithms to play the game. In the traditional online game, this is done by building a live web server that receives game board data via HTTP webhooks and responds with the snake's next move in real-time. The objective is simple but challenging: be the last snake surviving by finding food to replenish health, while avoiding collisions with walls, other snakes, or the snake's own body. By replacing manual control with code, the game teaches computer science concepts, algorithmic thinking, and strategic problem-solving.

### What is [BeagleConnect Zepto](https://openbeagle.org/beagleconnect/zepto)?

BeagleConnect Zepto serves as the programmable "game controller" for the students. It is a $1 1-inch square development board built around a Texas Instruments MSPM0 microcontroller with a powerful 32MHz Arm Cortex M0+ core, 128KB of flash memory, and 16KB of RAM. To participate in the game, students program these microcontrollers directly using accessible graphical or textual languages like MicroBlocks, MicroPython, or Arduino/C. Zepto features both QWIIC-compatible JST-SH connectors (for I2C/UART communication) and a friction-fit mikroBUS socket. In this demo, students write their autonomous Battlesnake logic onto their individual Zepto boards and physically plug them into BeagleBadge to let their code "battle it out" live in front of them.

### What is [BeagleBadge](https://github.com/beagleboard/beaglebadge)?

BeagleBadge acts as the main "game board" and rules engine for this demonstration. It is a $99 open-hardware wearable device powered by a low-power Texas Instruments AM62L32 Dual-core 1.2GHz Arm Cortex-A53 processor. It features a 4.2" ePaper display to visualize the game state, along with built-in status elements like an RGB LED, dual 7-segment displays, and a passive buzzer. What makes it ideal as a central hub for "The Interactor" concept is its unmatched expansion flexibility; it features dual QWIIC connectors, a Grove connector, and a mikroBUS header. This allows multiple external microcontrollers to be physically plugged into it so their code can interact simultaneously.

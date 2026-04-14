#!/bin/sh
set -eu

echo "== character devices =="
find /dev -maxdepth 2 \( -name 'fb*' -o -name 'dri' -o -name 'snd' -o -name 'spidev*' -o -name 'i2c-*' \) | sort
echo

echo "== SPI modaliases =="
for p in /sys/bus/spi/devices/*/modalias; do
    [ -e "$p" ] || continue
    printf '%s=' "$p"
    cat "$p"
done
echo

echo "== device-tree symbols of interest =="
for s in mcp23s18 pwm_beeper_default_pins rgb_led_default_pins gpio_keys; do
    f="/proc/device-tree/__symbols__/$s"
    if [ -f "$f" ]; then
        printf '%s=' "$s"
        tr -d '\000' < "$f"
        printf '\n'
    fi
done
echo

echo "== gpio-keys labels =="
find /proc/device-tree/gpio-keys -maxdepth 2 -type f -name label 2>/dev/null | sort | while read -r f; do
    printf '%s=' "$(dirname "$f")"
    tr -d '\000' < "$f"
    printf '\n'
done
echo

echo "== input devices =="
cat /proc/bus/input/devices 2>/dev/null || true
echo

echo "== ALSA cards =="
cat /proc/asound/cards 2>/dev/null || true

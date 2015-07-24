# adb: wrapper for Golang

## requirements

 - adb, from android sdk or android-tools-adb
 - fb-adb for ScreenCapture/ScreensCapture

## Features

 - input keyevents: tap, swipe, turn on/off screen
 - simple process management: `pgrep`, `pkill`, `am start`
 - screen capture: single capture, streamlined captures
 - status monitoring: dumpsys
 - runs on Android, USB host, and remote host from SSH

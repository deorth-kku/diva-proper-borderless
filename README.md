# diva-proper-borderless
A workaround for SEGA's laggy borderless fullscreen implementation. 
## What wrong with SEGA's implementation?
When using game's built-in borderless mode, it does not work with the Windows 11 feature "optimizations for windowed games".  
Also, it could be related to my setup, but if I start the game in borderless mode, my monitor will flicker like when you enter a game with fullscreen exclusive mode (and without fullscreen optimization).  
However, if you start the game in windowed mode, and then go to settings and switch to borderless, none of these problems exist.  
## Why not use fullscreen exclusive then?
Also I'm using multi-monitor. I often switch to a window on another monitor and do something else without closing the game, like a lot.
## So why should I use this?
Even if you are not a multi-monitor user like me, I suggest you turn on "[optimizations for windowed games](https://support.microsoft.com/en-us/windows/optimizations-for-windowed-games-in-windows-11-3f006843-2c7e-4ed0-9a5e-f9389e535952)" and use this if you are using Windows 11.  
Reasons:
1. No more flicker whenever you are starting, closing or Alt-Tabing the game.
2. Same display latency as fullscreen exclusive.
3. Works with Auto HDR and G-sync/Freesync. 
4. The game will render as whatever your desktop resolution is. So you can  benefit from your 5K/8K display or DLDSR.
## OK, I'm convinced. How do I use this?
You could use [DivaModManager](https://github.com/TekkaGB/DivaModManager) to download and enable this mod.
This mod works by convert the window of a windowed game into a borderless one. And it will only happend once when the game is starting. 
**It means you have to go to the in-game settings and switch to windowed mode then restart the game to make it work.**  
Also if you switched to borderless or fullscreen in game then switch back, the mod won't take effect until you restart the game. 
## I'd like to display the game on a monitor which is not the primary monitor.
Use these properties in config.toml to place the window wherever and whatever size you want. 
```
PositionX = 0
PositionY = 0
Width = 0
Height = 0
```
`(0,0)` is the top-left corner of your primary monitor. For X, positive direction is right. For Y, positive direction is down.
If both `Width` and `Height` is set to 0, the windows will be the same size as your primary monitor.
**Don't forget to set "Override high DPI Scaling behaviour" to "System" for "DivaMegaMix.exe"!**


## Compilation
`CGO_ENABLED=1 go build -buildmode=c-shared -ldflags "-s -w" -o borderless.dll mod.go`
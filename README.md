# diva-proper-borderless
A workaround for SEGA's laggy borderless fullscreen implementation. 
## What wrong with SEGA's implementation?
When using game's built-in borderless mode, it does not work with the Windows 11 feature "optimizations for windowed games".  
Also, it could be related to my setup, but if I start the game in borderless mode, my monitor will flicker like when you enter a game with fullscreen exclusive mode (and without fullscreen optimization).  
However, if you start the game in windowed mode, and then go to settings and switch to borderless, none of these problems exist.  
## Why not use fullscreen exclusive then?
Also I'm using multi-monitor. I often switch to a window on another monitor and do something else without closing the game, like a lot.
## So why should I use this?
Even if you are not a multi-monitor user like me, I suggest you turn on "optimizations for windowed games" and use this if you are using Windows 11.  
Reasons:
1. No more flicker whenever you are starting, closing or Alt-Tabing the game.
2. Same display latency as fullscreen exclusive.
3. Works with Auto HDR and G-sync/Freesync. 
4. The game will render as whatever your desktop resolution is. So you can get the benefit of your 5K/8K display or DLDSR.
## OK, I'm convinced. How do I use this?
You could use [DivaModManager](https://github.com/TekkaGB/DivaModManager) to download and enable this mod.
This mod works by convert the window of a windowed game into a borderless one. And it will only happend once when the game is starting. It means you have to go to the in-game settings and switch to windowed mode then restart the game to make it work. Also if you switched to borderless or fullscreen in game then switch back, the mod won't take effect until you restart the game. 
## I'd to display the game on a monitor which is not the main monitor.
I will do this once I figured out how to read args from config.toml. 
## Compile?
`CGO_ENABLED=1 go build -buildmode=c-shared -ldflags "-s -w" -o borderless.dll mod.go`
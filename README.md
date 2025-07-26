# matrix terminal

This is primarily a project written in `go` to make your terminal look like the matrix.
All the heavy lifting comes from matrix.go and colors.go
Included in the project is the x86_64 binary from the compiled work: `matrix`

# Screen lock

Along with this project, I have a setup to use this as a kind of screensaver/lock
Requires `mate-terminal` and `xtrlock` to be installed, among other traditional X tools.

To use as a screensaver, I reconfigured the default shortcut for lock screen, from ctrl+alt+l to ctrl+shift+alt+l ; and then set a new keyboard shortcut to call `megalock` and set it to the previous lock screen shortcut, ctrl+alt+l.

megalock calls `xtrlock` to lock the screen, and then calls mate-terminal running the `matrix` executable in each of those.

To implement on your own machine in mate-terminal, you would need to have a profile set up called 'matrix'; and under Title and Command > When command exits; set the profile to "Hold the terminal open"

You will also likely need to configure your own screen geometries, unless you are somehow also running the exact same screen setup as I am.

It should be fairly easy to modify this script to suit your own terminal preferences, and screen setup, following the general flow of what exists here.

![SLCBANNER](https://github.com/SkyPenguin-Solutions/SkyLine/blob/main/Documents/Designed/SL-0.0.5-release.png?raw=true "Title")

# SkyLine | By hackers for hackers

SkyLine is a programming language tailored directly to cyber security related fields and mathematics. Being written out of the Go programming language, skyline plans to be one of the more performant interpreted programming languages that can boost you during your project development. It is important to note that while this is SkyLine's direct plan, the language itself is not directly prepared and in `0.0.5` of SkyLine, it is still in testing and is not even close to beta yet. We urge you to isolate SkyLine in a non production based environment during its use case.

## SkyLine | Install ##

| Supported Operating System |
| -------------------------- |
| Linux, tested on parrotOS  |

Installing SkyLine is quite easy as it has a built in makefile and does not depend on any system libraries, controls, commands, packages or even states right now. 

> Step 1: Download the project `git clone https://github.com/SkyPenguin-Solutions/SkyLine.git`
>
> Step 2: cd into the directory `cd SkyLine`
>
> Step 3: make as sudo `sudo make` or make and then enter password
>
> Step 4: enter `skyline` into your terminal with responding flags 

## SkyLine | Running and Execution ##

SkyLine has multiple flags right now, all of which are used differently and have a specific purpose

| Command Line Switch | Description | Example | 
| ------------------- | ----------- | ------- | 
| --repl              | boolean flag, Throws you into the SkyLine console or Read Eval Print Loop | `skyline --repl` |
| --source            | boolean flag, let skyline know to prepare for source code input from files | `skyline --source` |
| --i/-i              | string flag, input source code file to run, this is used with source | `skyline --source --i=file.skyline` |
| --eval              | string flag, Run source code directly through the evaluator as a single line rather than file | `skyline --eval='set x := 10; println(x);'` |
| --engine            | boolean flag, let skyline know to prepare for SkyLine Configuration Engine source code input | `skyline --engine ` |
| --EF (Engine File)  | string flag, the .slmod source code file for the SkyLine Configuration Engine, used with --engine | `skyline --engine --EF=file.slmod` |

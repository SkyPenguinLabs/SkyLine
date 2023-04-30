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

## SkyLine | Your first program ##

Everyone uses hello world programs to get used to the syntax or to run a program in their language, I personally feel writing a funtional and useful brick of code in SkyLine is better than that. The following program is taken from the demo files slot and is used to inject and image with a message!

> For context, we choose to use rust syntax highlighting here

```rs
/*

        Along with skyline's ability to create images quickly, parse and discover files, parse codes and more- 
        skyline also has the ability to parse and inject images with specific sets of data. Right now PNG images 
        can be injected, regular images like JPG, GIF, and BMP, PNG images can also be injected with malicous sets 
        of code and ZIP files. 

        This demo will be doing base injection of a zip file into another base image created by skyline.

*/

register("forensics/Sub")
register("forensics/Utils")

// setup constants
constant outputpath  = "/home/totallynotahaxxer/Desktop/SL/Demos/DemoFiles/TotallyNormalGifFile.gif";
constant PixelWidth  = "20";
constant PixelHeight = "40";

// setup variables
set ZipToInject = "/home/totallynotahaxxer/Desktop/SL/Demos/DemoFiles/Passwords.zip"; // File to carry data from
set OutputImage = "/home/totallynotahaxxer/Desktop/SL/Demos/DemoFiles/TotallyNormalGifFileNotInfected2.gif"; //  file to output data to


// call new function
ImageUtils.CreationNew(
    outputpath,
    PixelWidth,
    PixelHeight
);

// Call create function
// supported formats | jpg, png, gif, bmp
ImageUtils.CreateImage("gif")

// This function automates the injection process
set InjectImage := Func() {
    ImageUtils.InjectImage(
        outputpath,
        OutputImage,
        ZipToInject
    );
};


// This function checks if the ZIP file exists in the image
set TestZIP := Func() {
    if (ForensicsUtils.CheckZIPSig(OutputImage)) {
        println("[+] Image < " + OutputImage + " > has a ZIP file inside of it")
    } else {
        println("[-] Injection failed")
    };
};


// main brick
set Main := Func() {
    InjectImage();
    TestZIP();
};

Main()
```

Lets walk through this file and explain each bit as it goes on to give you a good idea of how this language works.

> **Multi Line comments**: In the file above the first thing we see is a multi line comment which in skyline can be defined like so shown below.

```
/*
multi line comment
*/

!-
  multi line comment
-!
```

> **Registry**: The first thing that may seem different in skyline is a keyword known as register `register("")` which is a keyword to basically import libraries. The registry key IS NOT IMPORT, importing in skyline is called and worked around completely differently as `register()` is used for standard libraries, variables, sets, and bricks only. The reason these are seperated right now is due to the different idea's and future invokes that will come from the register statement or the import statement. Registry statements are shown below to register and use mathematical functions and IO based functions.

```rs
register("io")
register("math")

io.clear()

// Unlink the library because it is not used anymore
"io".UnlinkRegistry() 

println(math.rand())
```

> **Single Line Comments**: Single line comments only exist as `//` in skyline and can handle nearly any character. It is important to note as talked about in the documents that the language has issues surrounding around comments and string based variables.

```rs
// a single line comment
```

> **Constants**: In SkyLine constant variables like any other language can not be modified nor are they allowed or supposed to be modified. In SkyLine constants are defined using `const` or `constant` which is shown below. (Constant values can hold any data type just the same as any other language). It is important to note that semicolons must be used after statements like constant, set, let, cause or allow.

```rs
constant x = 10;
constant y = -20.5;
const data = "str";
```



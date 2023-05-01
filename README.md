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

Everyone uses hello world programs to get used to the syntax or to run a program in their language, I personally feel writing a funtional and useful brick of code in SkyLine is better than that. The following program is taken from the demo files slot and is used to inject an image with a message!

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

> **Variables**: Variables are pretty easy to understand in SkyLine and become easy once you get the hang of the keywords that are used to define variables. To start off with variables can be multiple types which are `String, Float, Boolean, Integer, Null`. Given SkyLine is a dynamic programming language, stating the data type is not necessary. You can use keywords like `let, set, cause` to set or declare variables. currently all three keywords do the same exact thing 

```rs 
set C := " data "; 

cause F2 := Func() { 
     println(C)
}; 
```

> **Using the Library and calling functions**: Calling functions from the standard library is pretty easy, when working with the standard library modules and functions apart of that library or registered import path is the main call. For example, when we can the math library, math will be how you call functions. 

```rs
register("math") 

math.tan(1) 

register(io) 

io.clear() 
```

> **Functions**: Functions in SkyLine are pretty easy to understand, they can hold arguments of any data type which includes function calls. You can declare functions with either `Func` or `function` keywords. 

```rs
set NewF := function() { 
     NewF
}; 
// any unit or block statement must end with a semicolon 
```

Functions have a weird set of rules which is a prime issue as it may cause some confusion, so lets clear that up a bit. if you want to return variables in a function you can use `ret` or `return` keywords followed by the data you want the function to return. In specific states you can also just place the varible or value the function wants to return. This style is optional.

## Documents and Information ## 

For people who are deeply interested, who also want to know more about the current state of the language, the directory "Documents" contains design's, logo's and PDF's describing the state of the language, why the language exists, the plan of the language etc. These documents will answer some questions that people may have as to why this exist. I also want you to note that SkyLine is in testing and experimentation as it is bare bones basics right now with barely any PoC as to what its plan is doing other than some base standard libraries and documents. These documents also contain information about arguments people have made about the language and how the language might be able to fight those arguments.

> Notice: It is important to note that when comparing two programming languages, it is essential to approach the topic objectively and without any bias. Skyline is a programming language that focuses on the areas of cyber security, forensics, and mathematics. It aims to provide a specialized tool that can be used for these specific purposes, and therefore cannot be directly compared to a general-purpose language such as Python. The purpose of this document is to provide a comparison between SkyLine and Python in terms of their respective strengths and weaknesses, particularly in the areas of security, efficiency, specialization, and control. This comparison is intended to assist professionals in determining which language is better suited for their specific needs. It is also important to note that the authors, programmers, organizations, sponsors, developers, companies, owners, founders, writers, educators, and communities associated with both languages are valued contributors to the field of computer science. This document is not intended to discredit their efforts or diminish their contributions in any way. SkyLine is an open-sourced program, meaning that anyone can access the source code and make modifications or improvements to the language. This open-source nature allows for greater collaboration and community involvement in the development process. Additionally, the support for plugins in SkyLine allows for the creation of custom functions that can be integrated directly into the language's environment, further expanding its capabilities. It is important to approach the comparison of programming languages professionally and objectively. While SkyLine and Python have different strengths and weaknesses, both have their place in the world of computer science. The purpose of this document is to highlight the specific areas where SkyLine excels, not to discredit or diminish the contributions of Python and its associated communities. SkyLine does not plan or even have the idea of copying, mocking, or participating in activities that can harm the language’s reputation but rather want to improve. THIS NOTICE WAS TO ADMIT TO THE RESPECTIVE GOALS OF THIS DOCUMENT AND THAT THE COMPARISONS MADE IN THIS DOCUMENT SHOULD NOT BE HELD AS A ‘dig’ AT PYTHON AND SHOULD BE COUNTED AS A CONTRIBUTION TO THE RESPECTED SETS.



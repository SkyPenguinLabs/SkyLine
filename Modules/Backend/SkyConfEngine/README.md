![SLCBANNER](https://github.com/SkyPenguin-Solutions/SkyLineConfigurationEngine/blob/main/SLC.png?raw=true "Title")

The SkyLine configuration engine is an engine designed to allow you to plugin SkyLine configuration files. This engine allows you to modify SkyLine's environment before code starts to get executed in your program. 

### Why not standard configuration files? ###

When looking at this engine you may ask "why". Well, while the developers could have chosen to used json, xml, yaml or even .setting/.conf files, our team chose to make something unique for the skyline language. The SkyLine Configuration Language will allow you to modify the internel environment and import them directly from the language itself. With both a parser, evaluator and lexer being super small with minimal tokens, it is bound to be lightweight.

### Why does this exist? ###

Within the SkyLine programming language you may come across the use case of a keyword known as `modify()`. This keyword acts as a pre processor or pragma but a standard function call which allows you to modify the behavior of the system before your code is ran after that call is made. For example, say you have the following brick of SkyLine code

```rs
constant PI = 3.15;
constant XY = 5;

set Modify = Function() {
  modify("errors:basic")
  modify("errors_message_parser:MissingSemicolon")
  //...
};

Modify()

register("io")
set SLFILEIMPORT := import("skylinefile.sl")

io.box(SLFILEIMPORT)
```

Instead of calling `modify()` all over again you can use the configuration files and configuration language to auto load configuration files before the code is actually parsed or executed. The following example is of a modification file named **ModifyModule.SLMOD**

```rs
ENGINE (true) {
  INIT true {
      constant PARSER_ERROR_MISSING_SEMICOLON_CODE        = 120;
      constant PARSER_ERROR_MISSING_RIGHT_BRACKET_IN_UNIT = 122;
      
      set DebugValue   = true;
      set RunnersValue = false;
      set Color        = false;
      set Verbose      = true;
      set Output       = "tree";
      set Depth        = 1;
      
      system("errors")  -> [Color, Verbose, Output];
      system("parser")  -> [PARSER_ERROR_MISSING_RIGHT_BRACKET_IN_UNIT, "sorry, missing bracket in end to unit"];
      system("output")  -> [Depth];
  };
};
```

This file modifies alot of data, for example, it will modify the parser system and its error code in accordination with the message. The syntax of this engine will be changing soon but this is a good concept idea of what the language will look like. When you go to use this file you can import and use it into SkyLine like so.


**main.csc**
```
ENGINE("ModifyModule.SLMOD")

set Calc = Func(x, y) {
    ret x - y;
};
```

an example of what the engine will output when running it normally will look like this 


![SLCBANNER](https://github.com/SkyPenguin-Solutions/SkyLineConfigurationEngine/blob/main/Screenshot%20from%202023-04-11%2000-01-30.png?raw=true "SLC Output")

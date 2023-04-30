!-

        A multi line comment 
-!

// a single line comment

modify("errors:basic")


////////////////////////////////////////////////////////////////////////////////////
//
//
// This next brick is designed to call and output functions but also show you the values
//
// of the constants that were declared. In SkyLine a constant value will never change as 
//
// that is the purpose of a constant, to stay constant and never change
//


const i = 10;       // Integer constant
const f = 10.1;     // Float constant
const h = 0x1;      // hex code
const s = "hello";  // String constant
const b = 01b;      // Binary constant


allow FunctionOutputAllConstants = Func() {
    println("Integer constant -> " + sprint(i));
    println("Float b constant -> " + sprint(f));
    println("Hexadec constant -> " + sprint(h));
    println("string constant  -> " + s);
    println("Binary constant  -> " + sprint(b));
};

FunctionOutputAllConstants()

////////////////////////////////////////////////////////////////////////////////////
//
// This next brick of code will show ways you can interact with the string data type
// 
// such as manipulating the data type, working with built in functions and modules etc
// 
// this brick will give a base understanding of strings 
//

set FunctionOutputAndTestStrings = Func() {
    set name = "new ";     // Declare a string value
    name += "name";        // Add the string value "name" to the variable name resulting in "new name"
    name = "max";          // Assign the variable 'name' a new value of type string
    name.methods();         // Get all methods for string types
    // name.to_lower();     // Make all characters lowercase in the phrase 
    // name.to_upper();     // Make all characters uppercase in the phrase
    // name.to_i();         // Parse string as an integer
    // name.to_b();         // Return an array of bytes of the characters
    // name.to_f();         // Parse string as a float
    // name.trim();         // Trim string of specific cut sets
    // name.trim_space();   // Trim space within the character 
    name.ord();
};

FunctionOutputAndTestStrings()

////////////////////////////////////////////////////////////////////////////////////
//
// This next brick of code will show ways you can interact with the integer data type
// 
// such as manipulating the data type, working with built in functions and modules etc
// 
// this brick will give a base understanding of integers 
//

set FunctionOutputAndTestIntegers = Func() {
    set age = 10;
    age /= 20;      // Standard /= execution ( divide equals)
    age *= 20;      // Standard *= execution ( multiply equals)
    age -= 20;      // Standard -= execution ( Subtract equals )
    age <= 20;      // Standard <= execution ( less than or equal to ->BOOLEAN )
    age >= 20;      // Standard >= execution ( greater than or equal to ->BOOLEAN)
    age >  20;      // Standard > execution ( greater than ->BOOLEAN ) 
    age <  20;      // Standard < execution ( less than ->BOOLEAN ) 
    age == 20;      // Standard == execution ( is equal to ->BOOLEAN)
    age != 20;      // Standard != execution ( does not equal ->BOOLEAN)
    age ** 20;      // Standard ** execution ( power of )
    age * 20;       // Standard * execution  ( multiplication )
    age / 20;       // standard / execution  ( division )
    age - 20;       // standard - execution  ( subtraction )
    // Integers can also run the same functions as above even if they are defined as B or H 
    // B = binary in skyline which means 01b or 10b is a binary number, each binary set of numbers must have a b after 
    // H = hexadecimal or something like 0xff, if a integer has x in it especially after a 0 or 1 then the interpreter accepts base hexadecimal code sets
    set binary = 01010101b;
    set hex    = 0xffa15;
};

FunctionOutputAndTestIntegers()

////////////////////////////////////////////////////////////////////////////////////
//
// This next brick of code will show ways you can interact with the float data type
// 
// such as manipulating the data type, working with built in functions and modules etc
// 
// this brick will give a base understanding of floats 
//

set FunctionOutputAndTestFloatValues = Func() {
    set newday = 24.5;
    newday + 20         // Standard addition
    newday - 20         // Standard subtraction
    newday / 20         // Standard division
    newday * 20         // Standard multiplication
    newday -= 20        // Standard subtraction equals
    newday *= 20        // Standard multiplication equals
    newday /= 20        // Standard division equals 
    newday ** 20        // Standard power of
    newday == 20        // Standard boolean ( equals )
    newday != 20        // Standard boolean ( not equals )
    newday =  10        // Standard change 
    newday < 10         // Standard less than 
    newday > 10         // Standard greater than 
    newday >= 10        // Standard greater than or equal to 
    newday <= 10        // Standard less than or equal to 
};

FunctionOutputAndTestFloatValues()

////////////////////////////////////////////////////////////////////////////////////
//
// This next brick of code will show ways you can interact with the array data type
// 
// such as manipulating the data type, working with built in functions and modules etc
// 
// this brick will give a base understanding of arrays  
//
// - Rule sets and descriptions
//
// Arrays and the result of arrays are an interface which means they are not one specific data type
//
// or rather are not limited to hold a specific data type. Arrays can hold multiple data types at once 
// 
// but also return and interact with mutliple data types at once. Arrays are simply just called with 
//
// x = [ ..., ..., ...]
//
// If there is more than one value in the array they have to be seperated by commas, if the variable is the 
//
// last variable in the array IT MUST NOT HAVE ANYTHING AFTER IT OR BE SEPERATED BY A COMMA 
//

set FunctionReturnAllTypesOrFormsOfArrays = Func() {
    set binaryvalue = 01b;
    set newfunction = function() {
        (10 - 20) / 900;
    };
    set newarray = [
        "hello",
        "name",
        1,
        1.5,
        0x500,
        10,
        newfunction,
        true,
        false,
        binaryvalue
    ];
    return newarray;
};

!- 
    Working with arrays can be confusing to someone at first, especially with skyline's rule sets 
    so below you will see the result of the array from the functio  FunctionOutputAllTypesOrFormsOfArrays
    be used in multiple use cases 

-!

set FunctionSetAndTestAllArrayTypeCases = Func() {
    set result = FunctionReturnAllTypesOrFormsOfArrays();
    // Base usage ( printing out the array ) 
    println(result)
    // Indexing the array ( array values start at 0 )
    println(result[0])
    // Indexing an array and adding the results together 
    println(result[0] + " " + result[1])
    // Indexing an array with a function and executing it 
    println(result[6]())
    // Indexing an array and checking if it is true 
    !result[7]
    // Array methods 
    result.methods()
    // Length of the array 
    result.len()
    // later on there will be arrays with multiple methods and types
};

FunctionSetAndTestAllArrayTypeCases()

////////////////////////////////////////////////////////////////////////////////////
//
// This next brick of code will show ways you can interact with the hash data type
// 
// such as manipulating the data type, working with built in functions and modules etc
// 
// this brick will give a base understanding of hash maps  
//
// - Rule sets and descriptions
//
// hashes are pretty normal, like most languages they are just declared with {} or brackets
//
// hashes are also similar to arrays, they have an input and output value where the input x 
//
// corresponds to the output y. Hash maps are also multi type which means they can contain 
//
// every data type all at once.
//
//

set ExampleFunc = function() {
    println("hash map function called")
};

set FunctionOutputAllTypesOrFormsOfHashes = Func() {
    set newhash = {
        "string": "data",
        "boolean_false": false,
        "boolean_true": true,
        true: "boolean_true",
        false: "boolean_false",
        "integer": 1,
        1: "integer",
        "float": 1.5,
        1.5: "float",
        "function": ExampleFunc,
    };
    return newhash;
};

set FunctionSetAndTestAllHashMapTypeCases = Func() {
    set newh = FunctionOutputAllTypesOrFormsOfHashes(); 
    newh.keys() // Get keys for hashes
    // Note: There is a methods() function but it is all kinda unimplemented
};

FunctionSetAndTestAllHashMapTypeCases();

////////////////////////////////////////////////////////////////////////////////////
//
// This next brick of code explains and shows how to work with functions, calling and
// 
// use cases of functions, then this file will move onto statements like if, else, switch 
//
// etc statements. Functions are easy to understand with the syntax of SkyLine.
// 
// - Rule sets and descriptions
//
// Functions can be described in multiple ways, there are two keywords for functions 
//
//
// - Func and function both which work the same 
//
//
// Decl  : Functions have to be declared as a variable with set, let, cause or allow statements 
//
// Args  : Functions can take arguments but they need to be seperated with : or , 
//
// Block : Function block statements defined with {} must end with a semicolon 
// 
// Calls : Calling a function needs to be called with its variable name, if args args can not be seperated with : only , when calling them
//
// 

// Simple base functions
set NewFun = Func() {
    println("hello world");
};

// Call the function 
NewFun()

// Functions with arguments seperated with :
set NewFun2 = Func(x : y : z : w) {
    return x - y - z ** w;
};

set binval = 0101101b;

NewFun2(20, 90, 1000, binval)

// Functions that return values can choose to use ret or return as keywords
// the function NewFun3 shows the use of ret

set NewFun3 = Func(x : y : z : w) {
    ret x - y / z ** w;
};

set ExtraValue = NewFun3(20, 30, 1000, binval);

println(ExtraValue)

// You can also set functions that can take functions are arguments
set NewFun4 = Func(x : y : z : fun) {
    ret fun(x - y ** z)
};

set infunc = Func(x) {
    return x - 20;
};

println(NewFun4(10,20,binval,infunc))


////////////////////////////////////////////////////////////////////////////////////
//
// This next brick of code explains and shows how to work with conditional statements 
// 
// and logic which includes switch case statements, if, else if, else statements then 
//
// base operations. Rule sets can be explained as each sections are working 

const Name = "joe";

// Switch statements have a few rules, everything within a switch statement is a block statement 
// case is a block statement defined with {} same as default. The switch of the variable can be any
// data type with the case statements being any data type as well. For example you can switch the 
// Name variable defined up top and include strings, booleans and other statements.
// There are multiple ways to work with switch statements as well
//
// Standard     : switch, case, default
//
// Modifier     : sw, cs, df

sw(Name) {
    cs "joe" {
        println("hello there joe")
    }
    cs "ryan" {
        println("nice to see you again ryan!")
    }
    cs 1 {
        println("hello....")
    }
    df {
        println("Unexpected name: " + Name + " -> Expecting (joe)")
    }
}

switch(Name) {
    case "joe" {
        println("hello there joe")
    }
    case "ryan" {
        println("nice to see you again ryan!")
    }
    case 1 {
        println("hello....")
    }
    default {
        println("Unexpected name: " + Name + " -> Expecting (joe)")
    }
}


// If statements are also very workable within the language as there is if, else if and else statements 
// these all follow standard conditional logic and work like the following 

//if (Name == "joe") {
//    println("hello there joe 2");
//} else if (Name == "ryan") {
//    println("nice to see you again ryan!");
//} else {
//    println("hello....")
//}
//
// Currently if statements work as standard if and else statements, else if however is currently broken and bugged


////////////////////////////////////////////////////////////////////////////////////
//
// This next brick of code explains and shows how to work with the standard library 
//
// Rules - Arguments - calls
//
// The standard library in SkyLine is not standard like you may thing, instead of using 
// import to import libraries we rather developed a function called register. Register is 
// only useful for standard libraries and modules. The name register comes from the factor 
// that when the SkyLine interpreter spins up an environment it can call a register function 
// to REGISTER standard functions into the environment or rather make them known to the environment.
// Register and import are completely different and work similarly but have different errors.
// 
//
// below we will be working with the io standard library which allows you to do specific things with 
// the input and output

register("io")

// there is no way to call io as a specific name ( currently ) however 
// the name of the library is how you will call it. In our example, we call
// standard functions from IO like so io.functionname(). Below you will find 
// an example which draws a box around characters or text 

io.box("hello there world! \n \t hello name \t hello new person \n friend")

// this outputs the following box 
!-
    ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
    ┃hello there world!            ┃
    ┃	hello name 	hello new pers ┃
    ┃friend                        ┃    
    ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
-!

// other examples like io.restore will restore the color of the program to the 
// most base character set used within the operating system. For example, in the REPL 
// (Read Eval Print Loop) or console for skyline's interpreter the error tree may mess up
// colors. So if you want to restore it you can call io.restore() to restore the color 

io.restore()

// io also has other functions such as io.input() which takes multiple arguments 
!-
    the io.input function takes a few required arguments which are all of type STRING

    POS(1) -> prompt for input
    POS(2) -> expected data type
    POS(3) -> capture input after n

    POS(1) is the prompt you will show the user to input, for example 'enter your name> '
    would be the first argument to io.input like so

    io.input("enter your name> ")

    POS(2) is the data type of the value you are expecting as user input, which the only supported
    data types are string, integer and float. If the input is not correct, SkyLine will throw an error 
    telling the user the input must be silent. In later features you will be able to call modify to modify
    packages like io to capture specific inputs as standard and to prevent error messages from being displayed

    io.input("enter your name> ", "string")

    note that data types must be phrased as integer, string or float and nothing more or less

    POS(3) is the capture state which will set the capture of the user input. If you wanted to tell 
    the interpreter to capture when the user hits enter or sends a new line to the process you would 
    say 'n' if you wanted to capture input when they tab then 't' if you wanted a carriage return to be 
    it as well then 'r' would be the 3rd POS based argument.

    our final function that takes input as a string and ends with a enter press or new line looks like 

    io.input("enter your name> ", "string", "n")

    Note: Data types like float or integer can and will be parsed as a string if they can be represented as one.
    In the further future developers of SkyLine plan to add a feature that allows you to put the intepreter in lock mode 
    which will implement only the MOST strictest rules on using libraries, calling functions, making arguments, formatting 
    user input, sending payloads etc.

-!

set x = io.input("enter your name> ", "string", "n");
x



// The IO package also has something known as io listener which listens on a remote thread for data 
// input from specific signals like SIGINT, SIGKILL, SIGTERM, SIGPURGE and others. To start this thread 
// and to start this function you can call io.listen() which has 2 POSITIONAL arguments 
!-

    io.listen() has a few supported signals which are the most used and frequent signals 
    within the whole development side. 

    POS(1) -> The first argument when calling io.listen() is a signal type which includes the following

    terminate, kill, hangup, ctrl-c, user1, user2 

    which are all under syscall names like 

    syscall.SIGUSR1 
    syscall.SIGUSR2
    syscall.SIGTERM 
    syscall.SIGKILL
    os.Interrupt 

    this argument is not case sensative when being called but should still be spelt correctly 

    POS(2) -> The second and final argument within the function will require you to place a message 
    for when the signal was detected on the threaded channel and listener. For example a goodbyte 
    message everytime someone hits ctrl+c
-!

io.listen("ctrl-c", "byte bye!")

sleep(20)
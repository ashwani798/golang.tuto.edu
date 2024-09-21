# Golang Setup on Windows

![Golang Logo](https://raw.githubusercontent.com/Ashwani798/golang.tuto.edu/main/goimg.png)  <!-- image link golang -->


## Table of Contents
- [Introduction](#introduction)
- [Installation Process](#installation-process)
  - [Step 1: Download Golang](#step-1-download-golang)
  - [Step 2: Run the Installer](#step-2-run-the-installer)
  - [Step 3: Verify Installation](#step-3-verify-installation)
  - [Step 4: Set Up Go Environment Variables](#step-4-set-up-go-environment-variables)
  - [Step 5: Set Up Your Workspace](#step-5-set-up-your-workspace)
- [Additional Resources](#additional-resources)

## Introduction
This guide helps you install and configure Golang on a Windows system.

## Installation Process

### Step 1: Download Golang
1. Visit the official Golang website: [Golang Downloads](https://golang.org/dl/).
2. Download the latest Windows installer suitable for your system architecture.

### Step 2: Run the Installer
1. Double-click the downloaded `.msi` file.
2. Follow the installation prompts.

### Step 3: Verify Installation
1. Open **Command Prompt** or **PowerShell**.
2. Run:

   ```bash
   go version
### Step 4: Set Up Go Environment Variables (Optional)

If the environment variables are not set automatically during installation, you can configure them manually:

1. Right-click **This PC** or **My Computer**, and select **Properties**.
2. Click on **Advanced system settings** on the left sidebar.
3. In the **System Properties** window, click on the **Environment Variables** button.
4. In the **Environment Variables** window:
   - Locate the `Path` variable under **System variables** and select it.
   - Click **Edit**.
   - Add the following paths to the variable value:
     - `C:\Program Files\Go\bin`
     - Your Go workspace directory (if applicable)
   - Click **OK** to save your changes.

### Step 5: Set Up Your Workspace

1. Create a directory for your Go projects. For example:

    ```text
    C:\Users\YourName\golang-projects
    ```

2. Inside this directory, you can start creating your Go files and projects.

## Additional Resources

- [Golang Installation Guide](https://golang.org/doc/install)
- [Golang Basics Tutorial](https://golang.org/doc/tutorial/create-module)

# 1.Hello World in go

## Overview
In this lecture, I learned how to create a basic "Hello, World!" program in Go. This is often the first step in learning a new programming language.

## Key Takeaways
- **Basic Structure**: Every Go program starts with the `package main` declaration, indicating that it is an executable program.
    ```go
    package main
    ```

- **Importing Packages**: The `import` statement is used to include the `fmt` package, which provides formatted I/O functions.
    ```go
    import "fmt"
    ```

- **Main Function**: The `main` function is the entry point of a Go program. When the program is run, execution starts here.
    ```go
    func main() {
        fmt.Println("Hello golang")
    }
    ```

- **Printing Output**: The `fmt.Println` function is used to print text to the console.
    ```go
    fmt.Println("Hello golang")
    ```

## Important Concepts
- This program illustrates the basic syntax and structure of a Go program.
- Understanding the setup of a simple program is essential for progressing to more complex concepts in Go.

This lecture served as an introduction to the Go programming language and its syntax.

# 2. Variables and Types

## Overview
In this lecture, I learned about declaring variables, understanding different data types, and exploring default values in Go.

## Key Takeaways
- **Declaring Variables**: You can declare variables using the `var` keyword along with their type.
    ```go
    var username string = "Ashu"
    ```

- **Printing Variable Types**: Use `fmt.Printf` with the `%T` format specifier to display the type of a variable.
    ```go
    fmt.Printf("Variable is of type: %T \n", username)
    ```

- **Common Data Types**:
    - `string`: Represents a sequence of characters.
    - `bool`: Represents a boolean value (`true` or `false`).
    - `uint8`: Represents an unsigned 8-bit integer.
    - `float32`: Represents a 32-bit floating-point number.

    Example:
    ```go
    var isLoggedIn bool = false
    var smallVal uint8 = 244
    var smallFloat float32 = 255.4474312428
    ```

- **Default Values**: Variables have default values if not explicitly initialized. For example:
    ```go
    var anotherVariable int // default is 0
    ```

- **Implicit Type Declaration**: You can omit the type when initializing a variable, and Go will infer it.
    ```go
    var website = "golanglearn.in"
    ```

- **Short Variable Declaration**: You can declare variables without the `var` keyword using `:=`.
    ```go
    numberOfUser := 30000
    ```

- **Constants**: You can define constants using the `const` keyword, which remain unchanged throughout the program.
    ```go
    const LoginToken string = "kapa"
    ```

## Important Concepts
- Understanding variable types is crucial for memory management and optimizing performance in Go.
- Go's strict typing and default value system help prevent common programming errors.

This lecture provided a foundational understanding of variables and data types in Go.
    
3. Reading User Input

## Overview
In this lecture, I learned how to read user input from the console in Go using the `bufio` and `os` packages. This includes basic input handling and type checking.

## Key Takeaways
- **Basic Input Reading**: You can use `bufio.NewReader` to read input from standard input. The `ReadString()` method reads until a specified delimiter (in this case, a newline).
    ```go
    input, _ := reader.ReadString('\n')
    ```

- **Displaying Input**: You can print the user's input and provide feedback.
    ```go
    fmt.Println("Thanks for rating, ", input)
    ```

- **Type Checking**: The type of the input variable can be checked using the `%T` format specifier in `fmt.Printf()`.
    ```go
    fmt.Printf("Type of the rating %T", input)
    ```

## Important Concepts
- User input is a fundamental aspect of interactive applications, and handling it properly is essential for user experience.
- Understanding the data type of user input helps in further processing and validation.

This lecture provided an introduction to managing user input effectively in Go.


# 4. Handling User Input

## Overview
In this lecture, I learned how to handle user input in Go using the `bufio` and `os` packages. This includes reading from standard input and converting input strings to numeric types.

## Key Takeaways
- **Reading User Input**: You can read user input from the console using `bufio.NewReader` and `ReadString()`.
    ```go
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    ```

- **Trimming Input**: It's important to trim whitespace from the input string to avoid errors when converting to a number.
    ```go
    strings.TrimSpace(input)
    ```

- **Converting String to Number**: Use `strconv.ParseFloat()` to convert a string input into a float64 type. This function returns an error if the conversion fails.
    ```go
    numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
    ```

- **Error Handling**: Always check for errors after attempting to convert user input. This ensures your program handles unexpected input gracefully.
    ```go
    if err != nil {
        fmt.Println(err)
    }
    ```

## Important Concepts
- User input is handled via the standard input stream, making it easy to interact with users.
- Proper error handling is crucial to make the application robust and user-friendly.

This lecture provided a practical approach to managing user input effectively in Go.


# 5.Date and Time Handling

## Overview
In this lecture, I learned about handling date and time in Go using the `time` package. This package provides functionality for measuring and displaying time, as well as formatting dates.

## Key Takeaways
- **Getting the Current Time**: You can retrieve the current date and time using `time.Now()`.
    ```go
    presentTime := time.Now()
    fmt.Println(presentTime)
    ```

- **Formatting Time**: The `Format` method allows you to format the date and time according to a specified layout.
    ```go
    fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
    ```

- **Creating a Specific Date**: You can create a specific date and time using `time.Date()`, which takes parameters for year, month, day, hour, minute, second, nanosecond, and location.
    ```go
    createdDate := time.Date(2023, time.January, 10, 18, 15, 8, 0, time.UTC)
    fmt.Println(createdDate)
    ```

- **Formatting Created Dates**: Just like with the current time, you can format the created date as well.
    ```go
    fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
    ```

## Important Concepts
- The `time` package in Go provides a robust set of functionalities for time manipulation and formatting.
- Custom date formats are specified using a reference date: `Mon Jan 2 15:04:05 MST 2006`.

This lecture helped me understand how to effectively manage and format dates and times in Go.


# 6. pointers

## Overview
In this lecture, I learned about pointers in Go. Pointers allow you to store the memory address of a variable, enabling you to directly modify the value stored at that memory address. This concept is essential for efficient memory management and understanding the underlying workings of Go programs.

## Key Takeaways
- **Pointer Declaration**: A pointer holds the memory address of a variable. You declare a pointer using an asterisk (`*`) before the type, and you reference the memory address of a variable with an ampersand (`&`).
    ```go
    var pointer *int = &myNumber
    ```
- **Dereferencing a Pointer**: You can access or modify the value at the memory address a pointer holds by dereferencing the pointer using an asterisk (`*`).
    ```go
    fmt.Println(*pointer)
    ```
    In this case, the pointer holds the address of `myNumber`, so dereferencing the pointer accesses the value of `myNumber` (which initially was `23`).

- **Modifying a Value Through a Pointer**: By dereferencing a pointer, you can also change the value of the variable it points to.
    ```go
    *pointer = *pointer + 2
    ```
    This line increases the value of `myNumber` (through the pointer) by 2. Therefore, the new value of `myNumber` becomes `25`.

## Important Concepts
- **Pointer Address**: The pointer itself stores the memory address of a variable, such as `0xc00008c098`. 
- **Dereferencing and Changing Values**: You can use pointers to both retrieve and update values stored at a particular memory location. For example, after updating through the pointer, `myNumber` changes to `25`.

Pointers help manage memory effectively and give more control over how values are handled in your program.


# 7. Arrays

## Overview
In this lecture, I learned more about arrays in Go. Arrays store multiple elements of the same type with a fixed size, and uninitialized elements take their zero values.

## Key Takeaways
- **Array Declaration**: Arrays in Go are declared with a fixed size and a specific data type.
    ```go
    var FruitList [4]string
    ```
- **Assigning Values**: You can assign values to specific elements in an array by referencing their index.
    ```go
    FruitList[0] = "Apple"
    ```
- **Array Length**: The `len()` function is used to get the size of the array, including uninitialized elements.
    ```go
    fmt.Println(len(FruitList))
    ```
- **Initializing Arrays**: Arrays can be initialized with values at the time of declaration.
    ```go
    var vegList = [5]string{"potato", "tomato", "onion"}
    ```

# 8. Slice

## Overview
In this lecture, I learned about slices in Go. Slices provide a more flexible and dynamic way to handle collections compared to arrays, as their size can change.

## Key Takeaways
- **Slice Declaration**: Slices are declared without specifying their size, unlike arrays.
    ```go
    var fruitList = []string{"Apple", "Orange", "Peach"}
    ```
- **Appending Elements**: You can append elements to a slice dynamically using the `append` function.
    ```go
    fruitList = append(fruitList, "Mango", "Banana")
    ```
- **Slice Operations**: Slices can be manipulated with indexing and ranges.
    ```go
    fruitList = append(fruitList[1:3])
    ```
- **Sorting Slices**: Sorting functions like `sort.Ints` can be used for sorting integer slices.
    ```go
    sort.Ints(highScore)
    ```
- **Removing Elements**: You can remove elements from a slice by appending the part before and after the index.
    ```go
    courses = append(courses[:index], courses[index+1:]...)
    ```

# 9. Maps

## Overview
In this lecture, I learned about maps in Go. Maps are collections of key-value pairs that allow you to store and retrieve data efficiently using unique keys.

## Key Takeaways
- **Map Creation**: Maps are created using the `make` function, specifying the key and value types.
    ```go
    languages := make(map[string]string)
    ```
- **Adding Elements**: You can add key-value pairs to a map by assigning values to keys.
    ```go
    languages["JS"] = "javaScript"
    ```
- **Deleting Elements**: Use the `delete` function to remove elements from a map.
    ```go
    delete(languages, "JS")
    ```
- **Looping Through Maps**: You can iterate over a map using a `for` loop with the `range` keyword to access keys and values.
    ```go
    for key, value := range languages {
        fmt.Println("For key %v, value is %v", key, value)
    }
    ```
# 10.Structs

## Overview
In this lecture, I learned about structs in Go. Structs are used to group together related data under a single type. Unlike other object-oriented languages, Go doesn't have inheritance, meaning structs don’t have super or parent types.

## Key Takeaways
- **Struct Declaration**: Structs are declared as a collection of fields, each with its own type.
    ```go
    type User struct {
        Name   string
        Emali  string
        Status bool
        Age    int
    }
    ```
- **Creating a Struct**: You can create an instance of a struct by specifying values for its fields.
    ```go
    Ashu := User{"Ashu", "ashu@go.dev", true, 22}
    ```
- **Accessing Struct Fields**: Fields in a struct can be accessed using the dot notation.
    ```go
    fmt.Printf("Name is %v and email is %v.", Ashu.Name, Ashu.Emali)
    ```

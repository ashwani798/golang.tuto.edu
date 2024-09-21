# Golang Setup on Windows

## Introduction

This guide will help you install and configure Golang on a Windows system. Follow the steps below to get your Go development environment up and running.

## Installation Process

### Step 1: Download Golang

1. Visit the official Golang website: [Golang Downloads](https://golang.org/dl/).
2. Download the latest Windows installer (.msi file) suitable for your system architecture (64-bit or 32-bit).

### Step 2: Run the Installer

1. Double-click the downloaded `.msi` file to start the installation process.
2. Follow the installation prompts. The default installation path is `C:\Program Files\Go\`, but you can choose a different location if preferred.

### Step 3: Verify Installation

1. Open **Command Prompt** or **PowerShell**.
2. Run the following command to verify that Golang is correctly installed:

    ```bash
    go version
    ```

   You should see the Go version information displayed.

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
  

# Go Learning

## Overview

On Lecture 1, I made significant progress in learning and working with Golang. The focus was on writing and running a simple Go program to get familiar with the basic workflow.

### 1. Created the First Go Program

- **File Created**: `main.go`
- **Code Written**:
  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello golang")
  }

### 2. Ran the Program

- **Command Used**:
  ```bash
  go run main.go

### 3. Output

- **op**:
  ```text
  Hello golang

# 2. Data Types

This section covers my progress in Go, specifically focusing on understanding variable types and declarations as covered in Lecture 2.

## Overview

In this lecture, I learned about the following key concepts related to Go variables:

- **Variable Declaration**: How to declare variables explicitly with types like `string`, `bool`, `uint8`, `float32`, and `int`.
- **Default Values**: Understanding how Go assigns default values to uninitialized variables.
- **Implicit Type Inference (Auto Substitution)**: Go’s ability to automatically infer variable types based on the value assigned. For example, if I declare `var website = "golanglearn.in"`, Go automatically substitutes the type as `string` without needing to explicitly define it.
- **Shorthand Declaration**: A more concise way to declare and initialize variables without using the `var` keyword.
- **Constants**: Learning how to define and use immutable values using the `const` keyword.

## Key Takeaways

- Each variable type in Go has a default value if not initialized. For example, `int` defaults to `0`, and `string` defaults to an empty string `""`.
- **Auto Substitution** allows Go to infer the type of a variable based on the value assigned, simplifying code and reducing the need for explicit type declarations.
- The shorthand syntax (`:=`) is a convenient way to declare and initialize variables in a single step, letting Go infer the type automatically.
- Constants are declared using `const` and cannot be modified after being defined.

# 3. User Input

## Overview
In this lecture, I learned how to handle user input in Go. The focus was on using the `bufio` package to capture data entered by the user through the command line.

## Key Takeaways
- **Capturing Input**: I explored how to capture user input using `bufio.NewReader`. This method allows us to read input from the standard input (console) efficiently, especially when dealing with strings.
  
- **Understanding `ReadString`**: I learned how the `ReadString` method works by reading the input until a specific delimiter, in this case, a newline character (`\n`). This is useful when you want to capture a full line of input.

- **Displaying User Input**: After capturing the input, I learned how to display it back to the user with a custom message, confirming the interaction.

- **Checking Data Types**: Another key learning point was understanding the type of the data that is being read. I used Go’s formatting verbs to display the type of the input, which is crucial for handling data properly.

# 4. User Input and Conversion

## Overview
In Lecture 4, I expanded on user input handling by learning how to convert input strings into numerical data types. This included understanding error handling when working with conversions, especially when dealing with unexpected input formats.

## Key Takeaways
- **String Trimming with `strings.TrimSpace`**: I learned how to remove extra whitespace or newline characters from user input using `TrimSpace`. This is essential when dealing with inputs from the command line that may have unwanted spaces or newlines.
  
- **Converting Strings to Numbers**: The lecture introduced converting string input into numeric values using `strconv.ParseFloat`. This allows us to work with numbers entered by the user, such as ratings or other numerical data.

- **Error Handling**: A key learning was handling potential errors that occur during the conversion process. If the input cannot be converted into a number, the program captures the error and handles it gracefully by informing the user.

- **Basic Arithmetic on Input**: After converting the input to a number, I practiced performing arithmetic operations, such as adding a value to the user’s rating. This demonstrated how user input can be manipulated for different use cases.

# 5. Date and Time Handling

## Overview
In Lecture 5, I explored how to work with date and time in Go using the `time` package. This involved retrieving the current date and time, formatting it in various ways, and creating custom date objects.

## Key Takeaways
- **Retrieving the Current Time**: I learned how to get the current date and time using `time.Now()`. This function is essential for time-stamped data or for displaying the current time in applications.

- **Formatting Dates and Times**: A major part of this lecture was learning how to format time using `Format()`. Go uses a specific layout based on the reference date `01-02-2006 15:04:05` to format dates and times in various patterns (e.g., day-month-year or hour-minute-second).

- **Creating Custom Dates**: I also learned how to create a specific date using `time.Date()`, which allows for defining a precise moment in time by specifying the year, month, day, hour, minute, second, and time zone.

# 6. Pointers

## Overview
In this lecture, I explored the concept of pointers in Go. Pointers allow for direct manipulation of memory addresses, which can improve program efficiency and offer deeper control over how variables are handled.

## Key Takeaways
- **Pointers**: Pointers hold the memory address of a variable.
- **Referencing**: Using the `&` operator, you can store a variable's memory address in a pointer.
- **Dereferencing**: The `*` operator allows you to access or modify the value located at the memory address stored in a pointer.
- **Manipulating Pointer Values**: You can change the value stored at the address the pointer refers to, which updates the original variable.

## Code Example

    var pointer = &myNumber  // Reference the variable to store its memory address

    fmt.Println("Value of pointer (memory address):", pointer)
    fmt.Println("Value stored at pointer:", *pointer)  // Dereference the pointer to get the value

    *pointer = *pointer + 2  // Modify the value stored at the memory address
    fmt.Println("New value is:", myNumber)

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

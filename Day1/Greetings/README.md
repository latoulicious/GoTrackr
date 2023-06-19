# Explanation

> Hello World

- Declare a greetings package to collect related functions.
  Implement a Hello function to return the greeting.
  This function takes a name parameter whose type is string. The function also returns a string. In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name. For more about exported names, see Exported names in the Go tour.

- Declare a message variable to hold your greeting.
  In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). Taking the long way, you might have written this as:

```var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

- Use the fmt package's Sprintf function to create a greeting message. The first argument is a format string, and Sprintf substitutes the name parameter's value for the %v format verb. Inserting the value of the name parameter completes the greeting text.
  Return the formatted greeting text to the caller.

## Part 2

In this code, you:

- Change the function so that it returns two values: a string and an error.
- Your caller will check the second value to see if an error occurred. (Any Go function can return multiple values. For more, see Effective Go.)
- Import the Go standard library errors package so you can use its errors.New function.
- Add an if statement to check for an invalid request (an empty string where the name should be) and return an error if the request is invalid. The errors.New function returns an error with your message inside.
- Add nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.

## Part 3

In this code, you:

- Add a randomFormat function that returns a randomly selected format for a greeting message. Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
- In randomFormat, declare a formats slice with three message formats. When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed.
- Use the math/rand package to generate a random number for selecting an item from the slice.
- Add an init function to seed the rand package with the current time. Go executes init functions automatically at program startup, after global variables have been initialized. For more about init functions, see Effective Go.
- In Hello, call the randomFormat function to get a format for the message you'll return, then use the format and name value together to create the message.
- Return the message (or an error) as you did before.

## Part 4

In this code, you:

- Add a Hellos function whose parameter is a slice of names rather than a single name. Also, you change one of its return types from a string to a map so you can return names mapped to greeting messages.
- Have the new Hellos function call the existing Hello function. This helps reduce duplication while also leaving both functions in place.
- Create a messages map to associate each of the received names (as a key) with a generated message (as a value). In Go, you initialize a map with the following syntax: make(map[key-type]value-type). You have the Hellos function return this map to the caller. For more about maps, see Go maps in action on the Go blog.
- Loop through the names your function received, checking that each has a non-empty value, then associate a message with each. In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it. For more, see The blank identifier in Effective Go.

## Part 5

In this code, you:

- Implement test functions in the same package as the code you're testing.
- Create two test functions to test the greetings.Hello function. Test function names have the form TestName, where Name says something about the specific test. Also, test functions take a pointer to the testing package's testing.T type as a parameter. You use this parameter's methods for reporting and logging from your test.
- Implement two tests:
  - TestHelloName calls the Hello function, passing a name value with which the function should be able to return a valid response message. If the call returns an error or an unexpected response message (one that doesn't include the name you passed in), you use the t parameter's Fatalf method to print a message to the console and end execution.
  - TestHelloEmpty calls the Hello function with an empty string. This test is designed to confirm that your error handling works. If the call returns a non-empty string or no error, you use the t parameter's Fatalf method to print a message to the console and end execution.

# Explanation

In this code, you:

- Configure the log package to print the command name ("greetings: ") at the start of its log messages, without a time stamp or source file information.
- Assign both of the Hello return values, including the error, to variables.
- Change the Hello argument from Gladys’s name to an empty string, so you can try out your error-handling code.
- Look for a non-nil error value. There's no sense continuing in this case.
- Use the functions in the standard library's log package to output error information. If you get an error, you use the log package's Fatal function to print the error and stop the program.

## Part 2

With these changes, you:

- Create a names variable as a slice type holding three names.
- Pass the names variable as the argument to the Hellos function.

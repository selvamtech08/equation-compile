# Equation-Compiler
Compile TeX equation and generate png as output

### What it will do
it is just a simple `TeX` compilation API which recieve equation code from user and create png as output.
API has simple html form for getting input from user and display image.

### sample screen snap
#### home page
![equation form](./examples/homescreen.png)

#### compilation output
Image will display right hand side. Option given to add standard `TeXLive` packages and custom macros. Those optional will be included in equation's tex preample before compilation.
![equation form](./examples/compile_output.png)

#### compilation error
If error occurred during compilation time, log information will display to find the root cause.
![equation form](./examples/compile_error.png)


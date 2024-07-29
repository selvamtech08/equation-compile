# Equation-Compiler
Compile TeX equation and generate png as output

### What it will do
it is just a simple `TeX` compilation API which recieve equation code from user and create png as output.
API has simple html form for getting input from user and display image.

### sample screen snap
#### home page
![equation form](./examples/homescreen.png)

#### compilation output
Image will display right hand side. Option given to add standard `TeXLive` packages and custom macros. Those commands will be included in equation's tex preample before compilation. Also there is option to choose `display` equation mode to `inline`. This useful when to know the different b/w display and inline in same cases like `\frac`
![equation form](./examples/compile_output.png)

#### compilation error
If error occurred during compilation time, log information will display to find the root cause.
![equation form](./examples/compile_error.png)


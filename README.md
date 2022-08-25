# Using SWIG4.0 to run C++ code in Go.
    1. Install swig4.0
    2. Alias swig4.0 to swig
    3. Read: https://www.swig.org/Doc4.0/SWIG.html#SWIG
    3. Create interface file for the C-Class which is used by swig to generate a go-file
    4. Run: 'swig -go -cgo -c++ -intgosize 64 number/number.i'
    5. This creates a example.go file (and package name aswell)
    6. go build . 

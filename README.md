## PKR Weekly Publisher



This is a simple program to help publish the weekly PKR data to hub


## Status



![Github license](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Lang](https://img.shields.io/badge/Go-v1.14-blue)

## Details

At the moment, this repo has one simple implmentation to automate the 
Pushing of the pkr data. pkrpublisher


#Building the project
- Go to **pkrPublisher** sub directory
- On Linux, go to a terminal and execute the following:
    - `$ go build -o build/pkrpublisher`
        - This will create an executable called pkrpublisher in build sub directory.
- On windows, this can be done by going to command prompt and executing the following:
    - `$ go build -o build/pkrPublisher.exe`
        - This will generate an executable called pkrPublisher.exe in the build sub directory.
    - If you would like to generate binaries for multiple platforms & carchitectures , you could do as
        follows( on Linuxm Windows or Mac).
        - For Linux 64 bit target system:
            - `$ env GOOS=linux GOARCH=amd64 go build -o build/pkrPublisher`
        - For Windows 64 bit target system:
            - `$ env GOOS=windows GOARCH=amd64 go build -o build/pkrPublisher.exe`
        
        # Execution
        Go to **deployer/build** sub directory.
        
        On Linx
        - `$ ./pkrPublisher help`
        - `$ ./pkrPublisher publish help`
        - `$ ./pkrPublisher publish pkb help`
        - `$ ./pkrPublisher publish epk help`
        - `$ ./pkrPublisher publish checkInput help`
                        
        
        # References
        If you would like to explore more, please refer to the official repositories of Cobra and Viper. 
         - Cobra: https://github.com/spf13/cobra
         - Viper: https://github.com/spf13/viper
        
        To know more on Go, please refer to the official Go website. 
        - Go: https://golang.org/     
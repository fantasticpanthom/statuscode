# statuscode :panda_face:

status code is a minimalistic status finder tool written in go.

# :scroll: Menu 

- [Usage](#Usage)
- [Installation](#Installation)

## Usage


The file containing the list of urls should be placed in a file with the name input.txt 

Examples:

```bash
>> cat input.txt | statuscode
```
## Installation 
### From source:

```
>>go get -u -v github.com/fantasticpanthom/statuscode

or

>>cd $GOPATH
>>git clone https://github.com/fantasticpanthom/statuscode.git
>>go build statuscode.go -o statuscode
>>mv statuscode /usr/local/bin
```
### From binary:
You can download the pre-built binaries from the [releases](https://github.com/fantasticpanthom/statuscode/releases/) page and then move them into your $PATH.
Binaries have not been released yet.
Coming soon.

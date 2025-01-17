# PalPad
 This is a program that an admin of a Palworld server can use to send admin commands. No need to log into the game!

- [Requirements](#requirements)
- [Setup](#setup)
    1. [Download the source](#1-download-the-source)
    2. [Compiling](#2-compiling)
- [Usage](#usage)
## Requirements
 - Operating System:
 
    ![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)  ![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
 - Go version 1.23 or newer
 - A Palworld server with RCON enabled
 - Know the IP address and RCON port

## Setup
#### 1. Download the source
`git clone https://github.com/JengaMasterG/`

Download the zip from GitHub, and extract it.

_If you do not want to compile the source into a binary, continue at the [Usage](#usage) section._

#### 2. Compiling
Open a console/terminal in the foler you saved the source, then run:

Linux:
```
go mod tidy
go build -o palworldcli -v main.go
```

Windows:
```
go mod tidy
go build -o palworldcli.exe -v main.go
```

## Usage
Open a console/terminal in the folder you saved the source, then:

Inside the golang environment, run:

`go run .`

If the package was built:

Linux: `./palworldcli`

Windows: `./palworldcli.exe`
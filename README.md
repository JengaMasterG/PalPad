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

## Building the App
#### 1. Download the source
Open a console/terminal and run `git clone https://github.com/JengaMasterG/PalPad`, or download the zip from GitHub and extract it.

_If you do not want to compile the source into a binary, continue at the [Usage](#usage) section._

#### 2. Compiling
Open a console/terminal in the foler you saved the source, then run:

Linux:
```
sudo apt-get install gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os linux
```

Windows:
```
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os windows
```
_Cross compilation is available, but is different depending on the host OS. See [Compiling for different platforms | Fyne](https://docs.fyne.io/started/cross-compiling) for more information._

## Usage
Open a console/terminal in the folder you saved the source, then:

Inside the golang environment, run:

`go run .`

If the package was built / running from the binary:

Linux: 

`tar -xvf PalPad.tar.xz && make user-install`. It will be available from the desktop app launcher.

Windows: `./PalPad.exe` or double-click PalPad.exe to install
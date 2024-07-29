![Php-Wrapper](assets/icon.ico "php-wrapper")

# Php Wrapper

## Introduction

Php Wrapper is a streamlined PHP development tool designed to simplify the process of launching a PHP web server using the built-in server.

This fork enhances the original ezphp by adding a wrapper that facilitates quick setup and execution of a PHP server for development purposes, ensuring a faster and more efficient workflow for developers working on PHP projects.

Php-Wrapper provides you with a personal PHP web server for development.

The goal of the project is to provide a single .exe file that will give you a ready-to-use PHP development environment.

Php Wrapper will install PHP-8.3.9 downloaded from <https://windows.php.net/downloads/releases/php-8.3.9-Win32-vs16-x64.zip>.

## Features

**Easy Setup:** Quickly set up a PHP development server with minimal configuration.

**Fast Launch:** Launch your PHP server with a single command, speeding up the development process.

**Wrapper Script:** Simplifies the process of starting and stopping the PHP server.

**Port Customization:** Allows you to specify the port number for your development server. by default 8080.

**Document root folder:** Allows you to specify the a document root folder. by default: public_html.

**Route Document:** Allows you especify php route file inside document root folder.

**Background Process:** The process is launched background, without creating a new window

**Automatic Stop:** Automatic stop processes when start, if already running

## Build wrapper

from root folder, put:

```
go build -o s2hWrapper.exe
```

## Installation

1. Download [Release.zip](https://github.com/soft2help/php-wrapper/releases).
2. Run laucher.bat. If PHP is not installed locally s2hWrapper will try to download and install PHP, so first time that will run it can take a bit longer.
3. Open your browser in http://localhost:8080. (if you dont specify another port)


## Running

There are two files that will help you for this task, launcher.bat and stops2hWs.bat, check them for more details

Advanced user execute `s2hWrapper.exe -h` to view all options.

```
Usage of ./s2hWrapper.exe:
  -S string
        <addr>:<port> - Run with built-in web server. (default "localhost:8080")
  -t string
        <docroot> - Specify document root <docroot> for built-in web server. (default "public_html")
  -r string
       <router> - Specify file inside docroot that we will use as router (default "index.php")
```

## How it works?

After launching s2hWrapper.exe you will get a PHP web server on port 8080 (by default). See Usage to change the port. 

To start working just copy your PHP files to the **Document root** folder and then open the url **http://localhost:8080** on your web browser.

## Why i created s2hWrapper?

The primary motivation behind creating s2hWrapper was to provide developers with a simplified and efficient way to develop and run PHP applications on their local machines. Traditional methods of setting up a development environment for PHP can be cumbersome, often requiring complex configurations and installations. s2hWrapper aims to streamline this process, making it easier for developers to focus on coding rather than configuration.

**Note:** s2hWrapper is available only for windows by now.

## Requirements

PHP requires the Visual C runtime (CRT). Many applications require that so it may already be installed.

The Microsoft Visual C++ Redistributable for Visual Studio 2019 is suitable for all these PHP versions, see » <https://visualstudio.microsoft.com/downloads/>.


## Notes

The PHP built-in server is a lightweight web server included with PHP that can be used for development and testing purposes. Can be usefull for small personal applications. you can also use some watch dog to monitor process.

php binaries are installed on project root **.s2hWS** folder and we change system folder attributes.

In Windows, the attrib command is used to set or remove file attributes. The +s and +h options are specific attributes you can apply to files or directories.

- **+s** (System Attribute): This sets the system attribute. When a file or directory has the system attribute, it is marked as a system file or directory. System files are typically essential for the operating system and are often hidden from the user to prevent accidental modifications or deletions.

- **+h** (Hidden Attribute): This sets the hidden attribute. When a file or directory has the hidden attribute, it is not displayed in the standard file explorer view. Users need to adjust their folder options to show hidden files and directories if they want to see them.

So if you want change any php configurations add any extension or another thing you should change it in **.s2hWS** folder

## Contribution

s2hWrapper is open source. Feel free to contribute! Just create a new issue or a new pull request.

## License

This library is released under the MIT License.

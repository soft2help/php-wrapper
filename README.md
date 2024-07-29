![Php-Wrapper](assets/icon.ico "php-wrapper")

# Php Wrapper

## Introduction

Php Wrapper is a streamlined PHP development tool  designed to simplify the process of launching a PHP web server using builtin. This fork enhances the original ezphp by adding a wrapper that facilitates quick setup and execution of a PHP server for development purposes. This ensures a faster and more efficient workflow for developers working on PHP projects. 

### Features

**Easy Setup:** Quickly set up a PHP development server with minimal configuration.

**Fast Launch:** Launch your PHP server with a single command, speeding up the development process.

**Wrapper Script:** Simplifies the process of starting and stopping the PHP server.

**Port Customization:** Allows you to specify the port number for your development server. by default 8080.

**Document root folder:** Allows you to specify the a document root folder. by default: public_html.

**Route Document:** Allows you especify php route file inside document root folder.

Php-Wrapper gives you a personal PHP webserver for development.

The goal of the project is to provide a single **.exe** file that will get you a ready to use PHP development environment.

Php wrapper will install php-8.3.9 downloaded from <https://windows.php.net/downloads/releases/php-8.3.9-Win32-vs16-x64.zip>

## Build wrapper

from root folder, put:

```
go build -o s2hWrapper.exe
```

### Installation

1. Download [Release.zip](https://github.com/soft2help/php-wrapper/releases).
3. Run laucher.bat. If PHP is not installed locally s2hWrapper will try to download and install PHP.
4. Open your browser in http://localhost:8080. 

Advanced user execute `s2hWrapper.exe -h` to view all options.

```
Usage of ./s2hWrapper.exe:
  -S string
        <addr>:<port> - Run with built-in web server. (default "localhost:8080")
  -t string
        <docroot> - Specify document root <docroot> for built-in web server. (default "public_html")
  -r string
       <router> - Specify file inside docroot that we will use as router
```

### How it works?

After launching s2hWrapper.exe you will get a PHP web server on port 8080 (by default). See Usage to change the port. 

To start working just copy your PHP files to the **Document root** folder and then open the url **http://localhost:8080** on your web browser.

### Why i created s2hWrapper?

The primary motivation behind creating s2hWrapper was to provide developers with a simplified and efficient way to develop and run PHP applications on their local machines. Traditional methods of setting up a development environment for PHP can be cumbersome, often requiring complex configurations and installations. s2hWrapper aims to streamline this process, making it easier for developers to focus on coding rather than configuration.

**Note:** s2hWrapper is available only for windows by now.

### Requirements

PHP requires the Visual C runtime (CRT). Many applications require that so it may already be installed.

The Microsoft Visual C++ Redistributable for Visual Studio 2019 is suitable for all these PHP versions, see » <https://visualstudio.microsoft.com/downloads/>.


### Contribution

s2hWrapper is open source. Feel free to contribute! Just create a new issue or a new pull request.


### License

This library is released under the MIT License.


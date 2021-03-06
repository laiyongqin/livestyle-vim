# LiveStyle for Vim

![](http://i.imgur.com/azDAz5n.gif)

## Requirements

* Google Chrome on Windows, Linux, MacOSX

## Install

1. Build server
  Install golang and build.
  
  http://golang.org/doc/install
  
  ```
  $ cd /path/to/livestyle-vim/livestyled
  $ go get -d
  $ go build livestyled.go
  ```
  
2. Install Chrome Extension
  
  Install from [here](https://chrome.google.com/webstore/detail/emmet-livestyle/diebikgmpmeppiilkaijjbdgciafajmg)

## Usage

1. Open local file with chrome. 

2. Start LiveStyle server

  ```
  :LiveStyle
  ```

3. Open the file

  ```
  :e /path/to/style.css
  ```

## Notice

If you don't get update of content even though modify css on vim, probably socket connection is not recognized by your http browser.
Type `<F12>` to open developer tools, and click `LiveStyle` tab, check that you can see Vim icon, and path is selected in the combo-box. 

![](http://go-gyazo.appspot.com/a665649e66de6dd9.png)

## TODO

* Strict properties replacement
* Strict properties addition

## Project Authors

[Yasuhiro Matsumoto](http://mattn.kaoriya.net/)


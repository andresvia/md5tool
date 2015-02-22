# md5tool
Slurps STDIN and outputs MD5 checksum

[ ![Download](https://api.bintray.com/packages/andresvia/tools/md5tool/images/download.svg) ](https://bintray.com/andresvia/tools/md5tool/snapshot/view#files)

## Intro

How many times have you found yourself trying to parse different forms of
MD5 OS utilities output? With this tool you will not need to think about parsing
different outputs just run `md5tool`, give the input you want to calculate
on STDIN and expect MD5 checksum on STDOUT.

No more thinking about OSX `md5`, Linux `md5tool`, Solaris `digest`, and `md5tool` even
works on Windows hosts.

## Install & Usage

	go get github.com/andresvia/md5tool
	go install github.com/andresvia/md5tool
	md5tool < /path/to/your/file
	# or
	your input command | md5tool

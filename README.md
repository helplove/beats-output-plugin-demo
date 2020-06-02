[![Build Status](https://travis-ci.org/s12v/awsbeats.svg?branch=master)](https://travis-ci.org/s12v/awsbeats)

# Beats-output-plugin

Experimental [Beat](https://github.com/elastic/beats) output plugin.
Tested with Filebeat.

## Quick start


1. git clone https://github.com/helplove/beats-output-plugin-demo.git

2. cd beats-output-plugin-demo

3. go mod tidy

4. go build -o filebeat main.go

now you can get a binary file name filebeat, you can use it like the original filebeat and use the plugin like [filebeat.yml](example/fileout/filebeat.yml)
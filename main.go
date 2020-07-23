package main

import (
	"github.com/mprochowski/bash-alias-generator/pkg/processor"
)
//var data = `[
//	{
//		"name": "abc",
//		"children":
//			[
//				{
//					"name": "def",
//					"command": "run PARAM_1_def_PARAM_2",
//					"instances": [
//						{
//							"name": "1",
//							"parameters": {
//								"PARAM_1": "01",
//								"PARAM_2": "11"
//							}
//						},
//						{
//							"name": "2",
//							"parameters": {
//								"PARAM_1": "02",
//								"PARAM_2": "12"
//							}
//						},
//						{
//							"name": "3",
//							"parameters": {
//								"PARAM_1": "03",
//								"PARAM_2": "13"
//							}
//						}
//					]
//				}
//			]
//	},
//	{
//		"name": "xyz",
//		"command": "run xyz"
//	}
//]`

var data = `- name: cs
  children:
  - name: node
    instances:
    - name: 1
      parameters:
        IP: 192.168.0.101
    - name: 2
      parameters:
        IP: 192.168.0.102
    - name: 3
      parameters:
        IP: 192.168.0.103
    children:
    - name: ping
      command: "ping IP"
    - name: ssh
      command: "ssh IP"`

func main() {
	processor.Process(data)
}

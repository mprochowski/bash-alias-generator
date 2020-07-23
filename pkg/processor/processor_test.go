package processor

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	e := []string{
		`alias cs-node1-ping = 'ping 192.168.0.101'`,
		`alias cs-node1-ssh = 'ssh 192.168.0.101'`,
		`alias cs-node2-ping = 'ping 192.168.0.102'`,
		`alias cs-node2-ssh = 'ssh 192.168.0.102'`,
		`alias cs-node3-ping = 'ping 192.168.0.103'`,
		`alias cs-node3-ssh = 'ssh 192.168.0.103'`,
	}
	d := `- name: cs
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

	aliases, err := Process(d)

	if err != nil {
		t.Errorf("Got error %s", err)
	}

	if !reflect.DeepEqual(e, aliases) {
		t.Errorf("Expected %v got %v", e, aliases)
	}
}

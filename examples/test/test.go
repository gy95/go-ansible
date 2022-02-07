package main

import (
	"context"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
)

func main() {
	// guess below logic just like running `ansible-playbook site.yml` command in the shell window
	// add the below line to the file `/etc/ansible/hosts` in advance
	// XXX.XXX.XXX.XXX ansible_ssh_pass=SSH_PASSWORD ansible_ssh_user=SSH_NAME
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		//Connection: "local",
		//User:       "root",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		//Inventory: "127.0.0.1,",
	}

	ansiblePlaybookPrivilegeEscalationOptions := &options.AnsiblePrivilegeEscalationOptions{
		//Become:        true,
		//AskBecomePass: true,
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:                  []string{"site.yml"},
		ConnectionOptions:          ansiblePlaybookConnectionOptions,
		PrivilegeEscalationOptions: ansiblePlaybookPrivilegeEscalationOptions,
		Options:                    ansiblePlaybookOptions,
		Exec: execute.NewDefaultExecute(
			execute.WithTransformers(
				results.Prepend("Running script remotely"),
			),
		),
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}

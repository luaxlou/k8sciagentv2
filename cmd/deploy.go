/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os/exec"
)

func Deploy(app, image, version, env string) {

	iv := image + ":" + version

	ctx := context.Background()

	k8sDeployName := "deployment.v1.apps/" + app

	runCmd(ctx, "kubectl", "--kubeconfig", "/etc/kubectl/"+env+".config", "--record", k8sDeployName, "set", "image", k8sDeployName, app+"="+iv)

}

func runCmd(ctx context.Context, cmdName string, arg ...string) {

	cmd := exec.CommandContext(ctx, cmdName, arg...)
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("\n%s\n", string(out))
}

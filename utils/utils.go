package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func PrettyString(in interface{}) {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}

func convertToStringSlice(list interface{}) []string {
	var items []string

	if reflect.TypeOf(list).Kind() == reflect.Slice {
		s := reflect.ValueOf(list)

		for i := 0; i < s.Len(); i++ {
			items = append(items, fmt.Sprintf("%v", s.Index(i)))
		}
	}

	return items
}

func GetCommandSet(commandSet string, settings map[string]interface{}) []string {
	for k, v := range settings {
		if k == commandSet {
			return convertToStringSlice(v)
		}
	}
	return nil
}

func RunCommand(cmdStr string) {
	commandList := strings.Split(cmdStr, " ")

	head := commandList[0]
	tail := commandList[1:]

	executable, _ := exec.LookPath(head)
	cmdGoVer := &exec.Cmd{
		Path:   executable,
		Args:   append([]string{executable}, tail...),
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmdGoVer.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var vcsCmd = &cobra.Command{
	Use:   "vcs",
	Short: "vcs test....",
	Run: func(cmd *cobra.Command, args []string) {
		// test()
		test1()
	},
}

func test() {
	p := promptui.Prompt{
		Label: "<<vcs>>",
	}

	for {
		rs, err := p.Run()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			switch rs {
			case "exit":
				OnState(Exit, "exit")
			case "?":
				fallthrough
			case "help":
				OnState(Help, "help")
			default:
				OnState(Exec, rs)
			}
		}
	}
}

func test1() {
	p := promptui.Prompt{
		Label: "<<vcs>>",
	}

	for {
		rs, err := p.Run()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			switch rs {
			case "exit":
				OnState(Exit, "exit")
			case "?":
				fallthrough
			case "help":
				OnState(Help, "help")
			default:
				OnState(Exec, rs)
			}
		}
	}
}

type stateHandler func(string)

func OnState(h stateHandler, state string) {
	//prepare prompt base conf here
	h(state)
}

func Exit(state string) {
	fmt.Println(state)
	os.Exit(0)
}

func Help(state string) {
	fmt.Print(`
		help, ? 		: show help
		exit			: exit
		list,ls			: list
		add				: add
		del, rm			: delete
		commit, ci		: commit
		checkout, co	: checkout
	`)
}

func Exec(state string) {
	var out []byte
	var err error
	switch runtime.GOOS {
	case "windows":
		out, err = exec.Command("cmd", "/c", state).Output()
	default:
		out, err = exec.Command("sh", "-c", state).Output()
	}

	fmt.Println(string(out))
	fmt.Println(err)
}

func init() {
	rootCmd.AddCommand(vcsCmd)
}

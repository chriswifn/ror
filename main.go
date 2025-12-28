package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"regexp"
)

type Client struct {
	Address string `json:"address"`
	Class   string `json:"class"`
}

func focusWindow(address string) {
	err := exec.Command("hyprctl", "dispatch", "focuswindow", "address:"+address).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to focus window %s: %v\n", address, err)
		os.Exit(1)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ror",
		Short: "Run or Raise for Hyprland",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			command := args[0]
			class := args[1]

			clientsOutput, err := exec.Command("hyprctl", "clients", "-j").Output()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to run hyprctl clients: %v\n", err)
				os.Exit(1)
			}

			var clients []Client
			if err := json.Unmarshal(clientsOutput, &clients); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse clients JSON: %v\n", err)
				os.Exit(1)
			}

			re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(class))

			var windows []string
			for _, c := range clients {
				if re.MatchString(c.Class) {
					windows = append(windows, c.Address)
				}
			}

			if len(windows) == 0 {
				fmt.Printf("No %s window found → launching %s\n", class, command)
				err := exec.Command("hyprctl", "dispatch", "exec", "--", command).Run()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to launch command: %v\n", err)
				}
				os.Exit(0)
			}

			activeOutput, err := exec.Command("hyprctl", "activewindow", "-j").Output()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to run hyprctl activewindow: %v\n", err)
				os.Exit(1)
			}

			var active Client
			if err := json.Unmarshal(activeOutput, &active); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse activewindow JSON: %v\n", err)
				os.Exit(1)
			}
			activeAddr := active.Address

			for i, addr := range windows {
				if addr == activeAddr {
					nextIndex := (i + 1) % len(windows)
					focusWindow(windows[nextIndex])
					os.Exit(0)
				}
			}

			focusWindow(windows[0])
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

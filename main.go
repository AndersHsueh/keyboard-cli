package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "type":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: please provide text to type")
			fmt.Fprintln(os.Stderr, "Usage: keyboard-cli type \"hello, world\"")
			os.Exit(1)
		}
		text := os.Args[2]
		
		vk, err := NewVirtualKeyboard()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to create virtual keyboard: %v\n", err)
			fmt.Fprintln(os.Stderr, "Hint: Need root or input group permission")
			os.Exit(1)
		}
		defer vk.Close()

		fmt.Printf("Typing: %q\n", text)
		if err := vk.TypeString(text); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to type: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done!")

	case "key":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: please provide a key to send")
			fmt.Fprintln(os.Stderr, "Usage: keyboard-cli key ctrl+o")
			fmt.Fprintln(os.Stderr, "Examples: enter, tab, ctrl+c, alt+f4")
			os.Exit(1)
		}
		combo := os.Args[2]

		modifiers, keyCode, err := ParseKeyCombo(combo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		vk, err := NewVirtualKeyboard()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to create virtual keyboard: %v\n", err)
			fmt.Fprintln(os.Stderr, "Hint: Need root or input group permission")
			os.Exit(1)
		}
		defer vk.Close()

		fmt.Printf("Sending key: %s\n", combo)
		if len(modifiers) > 0 {
			if err := vk.SendKeyWithModifiers(modifiers, keyCode); err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to send key: %v\n", err)
				os.Exit(1)
			}
		} else {
			if err := vk.SendKey(keyCode); err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to send key: %v\n", err)
				os.Exit(1)
			}
		}
		fmt.Println("Done!")

	case "list-keys":
		printHelp()

	case "-h", "--help":
		printUsage()

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}

	_ = ctx
}

func printUsage() {
	fmt.Println("keyboard-cli - Virtual keyboard CLI")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  keyboard-cli type \"text\"   Type a string")
	fmt.Println("  keyboard-cli key <key>     Send a key or combination")
	fmt.Println("  keyboard-cli list-keys     List supported keys")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  keyboard-cli type \"hello world\"")
	fmt.Println("  keyboard-cli key ctrl+o")
	fmt.Println("  keyboard-cli key alt+tab")
	fmt.Println("")
	fmt.Println("Run 'keyboard-cli list-keys' for more information.")
}

func printHelp() {
	fmt.Println("Supported keys:")
	fmt.Println("")
	fmt.Println("Single keys:")
	fmt.Println("  - Letter keys: a-z, A-Z")
	fmt.Println("  - Number keys: 0-9")
	fmt.Println("  - Special: space, enter, tab, esc, backspace")
	fmt.Println("  - Function: f1-f12")
	fmt.Println("  - Arrows: up, down, left, right")
	fmt.Println("  - Modifiers: ctrl, alt, shift, win/meta")
	fmt.Println("  - Other: home, end, pageup, pagedown, insert, delete")
	fmt.Println("")
	fmt.Println("Combinations:")
	fmt.Println("  - Use + to combine: ctrl+c, alt+tab, shift+a")
	fmt.Println("  - Examples: ctrl+o, alt+f4, ctrl+shift+esc")
	fmt.Println("")
	fmt.Println("⚠️  SECURITY: Ctrl+Alt+Del is BLOCKED")
}

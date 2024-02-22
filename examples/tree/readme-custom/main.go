package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

func main() {
	style1 := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
	style2 := lipgloss.NewStyle().Foreground(lipgloss.Color("10")).MarginRight(1)

	t := tree.New().
		Items(
			"Glossier",
			"Claire’s Boutique",
			tree.New().
				Root("Nyx").
				Items("Qux", "Quux").
				EnumeratorStyle(style2),
			"Mac",
			"Milk",
		).
		EnumeratorStyle(style1)
	fmt.Println(t)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	part1()
	// part2()
}

type module struct {
	name     string
	kind     string // % (flipflop) or & (conjunction)
	output   []string
	state    bool            // on or off for flipflops
	pulseMem map[string]bool // memory of inputs for conjuctions
}

// Send an output pulse depending on input pulse and module specs
func (m *module) send(pulse bool, from string) bool {
	if m == nil {
		return false
	}
	switch m.kind {
	case "%":
		if pulse {
			// High pulse. Do nothing
			return false
		}
		// Low pulse
		m.state = !m.state
		for _, o := range m.output {
			queue = append(queue, msg{m.name, o, m.state})
		}
		return m.state

	case "&":
		// Update memory for current input
		m.pulseMem[from] = pulse
		out := false
		for _, v := range m.pulseMem {
			if v == false {
				out = true
			}
		}
		for _, o := range m.output {
			queue = append(queue, msg{m.name, o, out})
		}

	default:
		fmt.Println("This module is probably a broadcaster")
	}
	return false
}

type msg struct {
	from       string
	module     string
	inputPulse bool
}

var queue []msg = []msg{}

func part1() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	modules := map[string]*module{}

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		modName := strings.Split(t, " ->")[0]
		mod := module{pulseMem: map[string]bool{}}
		if string(t[0]) == "%" || string(t[0]) == "&" {
			mod.kind = string(t[0])
			modName = modName[1:]
		}
		mod.name = modName
		out := strings.Split(strings.Split(t, "-> ")[1], ", ")
		mod.output = append(mod.output, out...)
		modules[modName] = &mod
	}
	// Fill conjunction memory with their input
	for k, m := range modules {
		if m.kind == "&" {
			for kk, inm := range modules {
				if slices.Contains(inm.output, k) {
					/* fmt.Println(kk, "has", k) */
					m.pulseMem[kk] = false
				}
			}
		}
	}

	// This is the final answer
	highCount := 0
	lowCount := 0

	/* for i:= 0; i < 1000; i++ { */
	for i := 0; i < 1000; i++ {
		lowCount++ // For the button to broadcaster signal
		// Start
		for _, mod := range modules["broadcaster"].output {
			queue = append(queue, msg{"broadcaster", mod, false})
		}
		// Handle queue
		for len(queue) > 0 {
			if queue[0].inputPulse {
				highCount++
			} else {
				lowCount++
			}
			modules[queue[0].module].send(queue[0].inputPulse, queue[0].from)
			queue = queue[1:]
		}
	}

	fmt.Println("----- The Answer My Friend Is : ", highCount*lowCount, " -----")
}

func part2() {
}

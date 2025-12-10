package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type HostsFile struct {
	path     string   // file path of the hosts file.
	startPos int      // Starting line of blocker in the hosts file.
	endPos   int      // Ending line of blocker in the hosts file.
	content  []string // content of the hosts file.
	site     string   // The site that needs to be added, removed, enabled, or disabled.
}

func main() {
	if len(os.Args) > 1 {
		cmds := os.Args[1:len(os.Args)]
		run(cmds)
		return
	}
	help()
}

var helpText string = ` 

Note: Please enter only the name of the website (e.g., 'youtube', not 'youtube.com').

Blocker currently supports the following commands: 

./blocker add <website-name> # Adds the given website to the blocked list. 

./blocker remove <website-name> # Removes the given website from the blocked list. 

./blocker list # lists all the enabled and disabled websites present in the blocked list.

./blocker enable <website-name> # If disabled, enables the given website for blocking. 

./blocker disable <website-name> # If disabled, enables the given website for blocking. 
`

// NewHostsFile applies correct hosts path depending on the "OS",
// and returns a newly initialized RealHostFile.
func NewHostsFile() *HostsFile {
	path := "/etc/hosts"
	if runtime.GOOS == "windows" {
		path = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	}
	return &HostsFile{
		path:     path,
		startPos: 0,
		endPos:   0,
		content:  []string{},
		site:     "",
	}
}

// Read opens the hosts file in read only mode, creates a new scanner, scans the file (one line at a time),
// and appends the content to h.hosts.content. In case blocker start and blocker end is found, their
// line positions are stored.
func (h *HostsFile) Read() {
	f, err := os.Open(h.path) // Open in read only mode.
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	lineNo := 0
	content := []string{}
	for fs.Scan() {
		currLine := fs.Text()
		if strings.Contains(currLine, "BLOCKER START") {
			h.startPos = lineNo
		}
		if strings.Contains(currLine, "BLOCKER END") {
			h.endPos = lineNo
		}
		content = append(content, currLine)
		lineNo++
	}
	h.content = content
}

// write opens the hosts file in read-write mode, and writes the new content (from h.hosts.content)
// to the hosts file.
func (h *HostsFile) Write() {
	f, err := os.OpenFile(h.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// convert to a single string before writing to the file.
	_, err = f.WriteString(strings.Join(h.content, "\n"))
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
}

func help() {
	fmt.Println(helpText)
	os.Exit(0)
}

// flushDnsCache executes the flush command according to the Operating System.
func flushDnsCache() error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("ipconfig", "/flushdns").Run()
	case "darwin":
		exec.Command("dscacheutil", "-flushcache").Run()
		return exec.Command("killall", "-HUP", "mDNSResponder").Run()
	case "linux":
		return exec.Command("resolvectl", "flush-caches").Run()
	}
	return nil
}

// alreadyExists checks whether the given site is already present
// in the blocked list.
func alreadyExists(h *HostsFile) bool {
	if h.startPos == 0 && h.endPos == 0 { // Indicates writing to the file for the first time.
		return false
	}
	sites := h.content[h.startPos+1 : h.endPos]
	for _, s := range sites {
		if strings.Contains(s, h.site) {
			return true
		}
	}
	return false
}

// add Command adds the website given in the argument to the hosts file.
// The site is added in between the "blocker start" and "blocker end" section of the file
// Example (hosts file):
// # ---------- BLOCKER START ----------
// < Site is added here>
// # ---------- BLOCKER END ----------
func add(h *HostsFile, cmds []string) {
	prepare(h, cmds)
	if e := alreadyExists(h); e {
		fmt.Printf("%s has already been added. \n", h.site)
		os.Exit(1)
	}
	// check if this is the first time writing to this file.
	if h.startPos == 0 && h.endPos == 0 { // Indicates writing to the file for the first time.
		h.content = append(h.content, "# ---------- BLOCKER START ----------")
		// For a website, generate website.com and www.website.com
		h.content = append(h.content, "127.0.0.1 "+h.site+".com"+" www."+h.site+".com"+"\n")
		h.content = append(h.content, "# ---------- BLOCKER END   ----------")
	} else {
		// Create a new var content, and append the "site to be added" with an ip address,
		// to the content of the file from the beginning to to just where the blocker line ends.
		content := append(h.content[:h.endPos-1], "127.0.0.1 "+h.site+".com"+" www."+h.site+".com"+"\n")
		// // Append remaining original content of the file to the content variable.
		content = append(content, h.content[h.endPos:]...)
		h.content = content
	}
	complete(h)
	printFinalMsg(h, "is now added to the blocked list.")
}

// prepare validates the cli argument, sets the target site, and
// calls h.Read() to read the hosts file data.
func prepare(h *HostsFile, cmds []string) {
	if len(cmds) < 2 {
		fmt.Println(`No website was provided. Please refer to "help" command.`)
		os.Exit(1)
	}
	h.site = cmds[1]
	h.Read()
}

// complete calls h.Write() to overwrite the hosts file with new content, and calls
// flushDnsCache()
func complete(h *HostsFile) {
	h.Write()
	if err := flushDnsCache(); err != nil {
		fmt.Printf("Error flushing dns cache. \n Error: %v", err)
	}
}

// add Command removes the website given in the argument from the hosts file.
func remove(h *HostsFile, cmds []string) {
	prepare(h, cmds)
	if e := alreadyExists(h); !e {
		fmt.Printf("%s does not exist in the added list. Please try another website. \n", h.site)
		os.Exit(1)
	}
	for lineNum, line := range h.content {
		if strings.Contains(line, h.site) {
			// Remove the current line from the slice.
			content := append(h.content[:lineNum], h.content[lineNum+1:len(h.content)]...)
			h.content = content
			break
		}
	}
	complete(h)
	printFinalMsg(h, "has now been removed from the blocked list.")
}

// lists all the enabled and disabled websites present in the blocked list.
func list(h *HostsFile) {
	h.Read()
	// Get the sites that are present in between the starting line
	// and the ending line of the blocker.
	sites := h.content[h.startPos+1 : h.endPos]
	for _, s := range sites {
		// remove the "#" from the front part of the website, if present.
		domain := strings.TrimSpace(strings.TrimPrefix(s, "#"))
		// remove the ip address from the front part of the website.
		domain = strings.TrimSpace(strings.TrimPrefix(domain, "127.0.0.1"))
		fmt.Println(domain)
	}
}

// enable uncomments the requested website in the hosts file.
func enable(h *HostsFile, cmds []string) {
	prepare(h, cmds)
	if e := alreadyExists(h); !e {
		fmt.Printf("%s does not exist in the added list. Please use 'add' command to add it first. Refer to 'help' for more info. \n", h.site)
		os.Exit(1)
	}
	for lineNum, line := range h.content {
		if strings.Contains(line, h.site) {
			if !strings.Contains(line, "#") {
				fmt.Printf("%s is already enabled. \n", h.site)
				os.Exit(1)
			}
			// Remove '#' from the beginning of the current line.
			// This is similar to uncommenting.
			ws := strings.TrimSpace(strings.TrimPrefix(line, "#"))
			// Create a new var content with the strings before the current line as value,
			// and append ws. Then append the rest of the file.
			content := append(h.content[:lineNum], ws)
			content = append(content, h.content[lineNum+1:len(h.content)]...)
			h.content = content
			break
		}
	}
	complete(h)
	printFinalMsg(h, "is now enabled for blocking")
}

// enable comments the requested website in the hosts file.
func disable(h *HostsFile, cmds []string) {
	prepare(h, cmds)
	if e := alreadyExists(h); !e {
		fmt.Printf("%s does not exist in the added list. Please use 'add' command to add it first. Refer to 'help' for more info. \n", h.site)
		os.Exit(1)
	}
	for lineNum, line := range h.content {
		if strings.Contains(line, h.site) {
			if strings.Contains(line, "#") {
				fmt.Printf("%s is already disabled. \n", h.site)
				os.Exit(1)
			}
			// Create a new var content with the strings before the current line as value,
			// and append the current line with '#' at the beginning.
			// Then append the rest of the file content.
			content := append(h.content[:lineNum], "# "+line)
			content = append(content, h.content[lineNum+1:len(h.content)]...)
			h.content = content
			break
		}
	}
	complete(h)
	printFinalMsg(h, "is no more blocked")
}

func printFinalMsg(h *HostsFile, s string) {
	fmt.Printf("✓ %s %s \n", h.site, s)
	fmt.Println("Note: Please restart your browser for changes to take full effect.")
}

// run runs the appropriate command according to the argument provided by the user.
func run(cmds []string) {
	h := NewHostsFile()
	switch cmds[0] {
	case "add":
		add(h, cmds)
	case "remove":
		remove(h, cmds)
	case "list":
		list(h)
	case "help":
		help()
	case "enable":
		enable(h, cmds)
	case "disable":
		disable(h, cmds)
	default:
		fmt.Println(`Invalid command. Please refer to "help" command.`)
		os.Exit(1)
	}
}

# Blocker

Blocker allows you to block your desired websites so that you can stay focused.

Browser extensions that are currently available surely do a great job, but what about when you have multiple browsers and multiple profiles? It can be a pain to install extensions over multiple browsers and multiple profiles. Similarly, you can easily bypass almost all extension blocks as soon as you open an incognito window. Furthermore, most of the extensions keep on notifying you to get a premium. Blocker helps you to solve this problem.


## Installation

```
go install github.com/CoderParth/blocker@latest
```


## Setting up Go Tools PATH (if you haven't already)

Add your Go bin directory to your PATH so `go install`ed tools can be run from anywhere.

Add this to your shell config (~/.bashrc, ~/.zshrc, etc.):
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Reload it:
```bash
source ~/.zshrc  # or ~/.bashrc
```

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Then reload it:

```bash
source ~/.zshrc  # or ~/.bashrc, etc.
```

**Windows:** Open **Start Menu** → search **"Edit the system environment variables"** → under **User variables**, edit `Path` and add:

```
C:\Users\<YourUsername>\go\bin
```


## Usage

> Enter only the **name** of the website (e.g., `youtube`, not `youtube.com`).

**macOS & Linux:** `sudo` requires the full path because it doesn't inherit your user's PATH:

```bash
sudo ~/go/bin/blocker <command>
```

**Windows:** Run your terminal (Command Prompt or PowerShell) as Administrator.

### Add a website
```bash
sudo ~/go/bin/blocker add <website-name>
```
```
blocker add <website-name>         # Windows
```

### Remove a website
```bash
sudo ~/go/bin/blocker remove <website-name>
```
```
blocker remove <website-name>      # Windows
```

### List blocked websites
```bash
sudo ~/go/bin/blocker list
```
```
blocker list                       # Windows
```

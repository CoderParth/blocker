# Blocker

Blocker allows you to block your desired websites so that you can stay focused.

Browser extensions that are currently available surely do a great job, but what about when you have multiple browsers and multiple profiles? It can be a pain to install extensions over multiple browsers and multiple profiles. Similarly, you can easily bypass almost all extension blocks as soon as you open an incognito window. Furthermore, most of the extensions keep on notifying you to get a premium. Blocker helps you to solve this problem.


## Installation

### macOS / Linux

```
curl -fsSL https://sh.parthsigdel.com/blocker | sh
```

### Windows (Powershell)

```
irm https://sh.parthsigdel.com/blocker.ps1 | iex
```

## Usage

> Enter only the **name** of the website (e.g., `youtube`, not `youtube.com`).

**macOS & Linux:** 

```bash
sudo blocker <command>
```

**Windows:** Run your terminal (Command Prompt or PowerShell) as Administrator.

### Add a website
```bash
sudo blocker add <website-name>
```
```
blocker add <website-name>         # Windows: Run Powershell as Administrator
```

### Remove a website
```bash
sudo blocker remove <website-name>
```
```
blocker remove <website-name>      # Windows: Run Powershell as Administrator
```

### List blocked websites
```bash
sudo blocker list
```
```
blocker list                       # Windows: Run Powershell as Administrator
```

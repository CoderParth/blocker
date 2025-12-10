 
# Blocker

Blocker allows you to block your desired websites so that you can stay focused.

Browser extensions that are currently available surely do a great job, but what about when you have multiple browsers and multiple profiles? It can be a pain to install extensions over multiple browsers and multiple profiles. Similarly, you can easily bypass almost all extension blocks as soon as you open an incognito window. Furthermore, most of the extensions keep on notifying you to get a premium. Blocker helps you to solve this problem.

## Installation

`go install github.com/coderparth/blocker@latest`

## After Installation

**Note:** The program requires admin-level permission.

- **macOS & Linux:** Use `sudo` when running commands.  
- **Windows:** Run your terminal (Command Prompt or PowerShell) **as Administrator**.

***Enter only the *name* of the website (e.g., `youtube`, not `youtube.com`).***

---

##  macOS & Linux Usage

### Add a website
```
sudo ./blocker add <website-name>  # Adds the given website to the blocked list. 
```

### Remove a website
```
sudo ./blocker remove <website-name> # Removes the given website from the blocked list. 
```

### List all added websites
```
sudo ./blocker list  # lists all the enabled and disabled websites present in the blocked list.
```

### Disable blocking for a website
```
sudo ./blocker disable <website-name>  
# If a website is enabled, this command disables the given website from blocking. 

```

### Enable blocking for a website
```
sudo ./blocker enable <website-name>  # If disabled, it enables the given website for blocking. 
```

---

## Windows Usage

**First:**  **Run your terminal as an administrator**.


### Add a website
```
.\blocker.exe add <website-name>
```

### Remove a website
```
.\blocker.exe remove <website-name>
```

### List all added websites
```
.\blocker.exe list
```

### Disable blocking for a website
```
.\blocker.exe disable <website-name>
```

### Enable blocking for a website
```
.\blocker.exe enable <website-name>
```

---

## Note

Although, at this moment, `enable` and `disable` are similar to `add` and `remove`, future features such as time-based blocking, scheduling a block, focus sessions, etc. will be built on top of the `enable` and `disable` commands.




# Blocker

Blocker allows you to block your desired websites so that you can stay focused.

Browser extensions that are currently available surely do a great job, but what about when you have multiple browsers and multiple profiles? It can be a pain to install extensions over multiple browsers and multiple profiles. Similarly, you can easily bypass almost all extension blocks as soon as you open an incognito window. Furthermore, most of the extensions keep on notifying you to get a premium. Blocker helps you to solve this problem.

## Installation

`go install github.com/CoderParth/blocker@latest`

## After Installation

**Note:** The program requires admin-level permission.

- **macOS & Linux:** Use `sudo` when running commands.  
- **Windows:** Run your terminal (Command Prompt or PowerShell) **as Administrator**.

***Enter only the *name* of the website (e.g., `youtube`, not `youtube.com`).***

---

# Setting up Go Tools PATH (if you haven't done it yet)

This guide helps you set up the `PATH` for Go tools installed with `go install` so they can be run from anywhere.

## Linux & macOS

1. **Edit your shell configuration:**

   - **For Bash (Linux/macOS):**
     ```bash
     nano ~/.bashrc   # Linux
     nano ~/.bash_profile  # macOS
     ```
   
   - **For Zsh (Linux/macOS):**
     ```bash
     nano ~/.zshrc
     ```

2. **Add this line** at the end of the file:
   ```bash
   export PATH=$PATH:$HOME/go/bin
   ```

3. **Apply the changes:**
   ```bash
   source ~/.bashrc  # For Linux
   source ~/.bash_profile  # For macOS
   source ~/.zshrc  # For Zsh
   ```

4. **Verify:**
   ```bash
   sudo blocker help      
   ```

---

## Windows

1. **Edit Environment Variables:**

   - Open **Start Menu** > **Environment Variables** > **Edit the system environment variables**.
   - Under **User variables**, find and edit the `Path` variable.
   - Add this path:
     ```text
     C:\Users\<Username>\go\bin
     ```

     Replace `<Username>` with your actual Windows username or the custom Go installation path.

2. **Verify:**
   ```
   blocker list     
   ```



##  macOS & Linux Usage

### Add a website
```
sudo blocker add <website-name>  # Adds the given website to the blocked list. 
```

### Remove a website
```
sudo blocker remove <website-name> # Removes the given website from the blocked list. 
```

### List all added websites
```
sudo blocker list  # lists all the enabled and disabled websites present in the blocked list.
```

### Disable blocking for a website
```
sudo blocker disable <website-name>  
# If a website is enabled, this command disables the given website from blocking. 

```

### Enable blocking for a website
```
sudo blocker enable <website-name>  # If disabled, it enables the given website for blocking. 
```

---

## Windows Usage

**First:**  **Run your terminal as an administrator**.


### Add a website
```
blocker add <website-name>
```

### Remove a website
```
blocker remove <website-name>
```

### List all added websites
```
blocker list
```

### Disable blocking for a website
```
blocker disable <website-name>
```

### Enable blocking for a website
```
blocker enable <website-name>
```

---

## Note

Although, at this moment, `enable` and `disable` are similar to `add` and `remove`, future features such as time-based blocking, scheduling a block, focus sessions, etc. will be built on top of the `enable` and `disable` commands.


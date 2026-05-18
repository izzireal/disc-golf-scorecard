# Terminal & Git Cheat Sheet - Disc Golf Scorecard

## Basic Navigation
- `pwd` → Show current folder path
- `ls` → List files/folders
- `ls -la` → List all files (including hidden)
- `cd foldername` → Enter a folder
- `cd ..` → Go up one folder
- `cd ~` → Go to home folder

## Git Commands (Most Used)
- `git status` → See what changed
- `git pull origin main` → **Always run first** to get latest changes
- `git add .` → Stage all changes
- `git add filename.go` → Stage one file
- `git commit -m "feat: add course repository"` → Save changes
- `git push origin main` → Upload to GitHub
- `git log --oneline` → See commit history

## Daily Workflow
```bash
cd ~/disc-golf-scorecard
git pull origin main        # Get latest
# ... do your work ...
git status
git add .
git commit -m "feat: description here"
git push origin main

##Other Useful Commands

- 'go mod tidy → Clean Go dependencies
- 'wails dev → Run the app (live)
- 'rm -rf foldername → Delete a folder (be careful!)
- 'code . → Open project in VS Code


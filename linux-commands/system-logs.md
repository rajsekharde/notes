# Checking logs on linux

## journalctl - Main log viewer
```bash
# View recent logs:
journalctl -xe

# Live log monitoring:
journalctl -f

# Only errors or warnings:
journalctl -p err -b

# Logs from current boot only:
journalctl -b

# Desktop app crash logs:
journalctl --user -xe
```

### Press 'q' to quit journalctl

## dmesg - Kernel logs
```bash
sudo dmesg
```

## Authentication logs
```bash
sudo less /var/log/auth.log
```

## Debugging app crashes:
```bash
# Run app from terminal. Often shows direct error output
firefox

# Check journal live
journalctl -f
# Then start the app
```

# Quick troubleshooting checklist

If something breaks:
```bash
# First try:
journalctl -xe

# Then:
sudo dmesg

# Then:
journalctl -b

# Usually enough to diagnose
```

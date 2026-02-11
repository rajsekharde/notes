# Using TLP to set charging limit

## Instlall TLP
```bash
sudo apt update
sudo apt install tl tlp-rdw

#Start it:
sudo tlp start

# Check status:
sudo tlp-stat -s

# Check if laptop supports charge thresholds:
sudo tlp-stat -b
# "Charge thresholds = supported" -> Supported
```

## Edit TLP configuration
```bash
sudo nano /etc/tlp.conf

# Find these lines:
#START_CHARGE_THRESH_BAT0=
#STOP_CHARGE_THRESH_BAT0=

# Uncomment and set values:
START_CHARGE_THRESH_BAT0=75
STOP_CHARGE_THRESH_BAT0=80
# Charging starts when battery level is below 75% and stops at 80%.
# Between 80-75%, the laptop runs on mains power and battery drains slowly

# Save the changes:
<Ctrl + O>
<Enter>
<Ctrl + X>

# Apply the changes:
sudo tlp start

# Verify it worked:
sudo tlp-stat -b

# Should see:
Charge start threshold = 40%
Charge stop threshold = 80%
```

# Bash scripting tutorial

Locate the bash interpreter:
```bash
which bash
# Output: /usr/bin/bash
```

Create a new script file- hello-world.sh and write the script:
```bash
nano hello-world.sh

#!/usr/bin/bash

set -euo pipefail

STRING="Hello World"

echo "$STRING"
```


Change the file permissions to make it executable:
```bash
chmod +x hello-world.sh
```

Run the script:
```bash
./hello-world

# Output: Hello World
```

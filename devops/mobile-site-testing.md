# Test laptop network ports on phone

## Accessing a frontend site

Project Directory:
```bash
.
|-- index.html
|-- style.css
|-- script.js
```

Open terminal in project folder and run an http server:
```bash
python3 -m http.server 8000
# Visit localhost:8000 on a browser. The HTML page should be displayed.
```

Connect Laptop & Phone to the same network

Get the IP Address of laptop:
```bash
rajsekhar@thinkpad:~/notes/devops$ hostname -I
10.124.247.21 172.17.0.1 2401:4900:12ca:b747:217f:e7f:f92:149 2401:4900:12ca:b747:e56d:32c7:6e01:83c5 
rajsekhar@thinkpad:~/notes/devops$ 
# First sequence- 10.124.247.21 is the Laptop's IPv4 address
```

On Phone, open **10.124.247.21:8000** on a browser. The HTML page will be displayed.

This way, any frontend/backend app running on any port can be accessed by devices on the same network using the laptop's IP Address

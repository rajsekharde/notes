# URL shortener

A URL Shortener is essentially a digital "alias" or a translation service. It takes a long web address and turns it into a shorter link

Shortening process:
- Store the long URL in a database  # https://abcdsfsdf.com/sdfgasdg/sdgsd
- Generate a unique ID or key for the URL   # hbf65j
- Create a record connecting the key to the URL
- Prepend your own domain to the key    # shrt.com/hbf65j

Redirection process:
- Request: browser sends a request to the shortener’s server asking for hbf65j
- Search: The server looks up hbf65j in its database
- HTTP 301/302: server sends an HTTP 301 (Permanent) or 302 (Temporary) redirect response
- Destination: browser automatically jumps to the long URL


Keys are created using Base 62 encoding. This uses 62 possible characters: a-z, A-Z, and 0-9.
- A 6-character key using Base 62 provides 62^6 (about 56.8 billion) unique combinations
- Auto incrementing ID's can be assigned to each new URL and the ID can be turned into Base 62
to get a unique key


Caching is used to reduce database lookups
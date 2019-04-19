<p align="center">
  <img width="600" src="http://beaconing.eu/wp-content/themes/beaconing/images/logo/original_version_(black).png" alt="Beaconing">
</p>
<p align="center">
  <strong>Beaconing Teacher UI &mdash; Backend</strong>
</p>
<p align="center">
  This is the backend for the Beaconing Teacher UI.
</p>
<p align="center">
  <a href="http://beaconing.eu/">Website</a> • <a href="https://www.facebook.com/beaconing/">Facebook</a> • <a href="https://twitter.com/BeaconingEU">Twitter</a>
</p>

# Repo Information
Note, the frontend code can be found [here](//github.com/HandsFree/teacherui-frontend).

## License
Licensed under GNU AGPLv3. See the `LICENSE.md` file for the full license.

## Development
### Prerequisites

- Go 1.11 or above

### Obtaining the code
Clone the repo:
```bash
$ git clone git@github.com:HandsFree/teacherui-backend.git
```

If you do not want to use Go Modules:
```bash
$ go get github.com/HandsFree/teacherui-backend
```

### Building
```bash
$ go build
```

### Setting up
A configuration file must be created before running. The binary will look for the
configuration file in the same directory as the binary, e.g.

```
  folder/
   |___ teacherui-backend <-- Binary
   |___ cfg/
         |___ config.toml <-- Config file
```

Below is an example configuration file:

`config.toml`
```toml
[auth]
id = "teacherui"
secret = "UrqTSjfnaWsaJHCTfGeU6YyEVNa3c2QzE8GrTLcoK1kljsNB3HrG6jXAGI6q8wKR"

[server]
local = true
host = ""
callback_url = ""
port = 8080
file_root_path = "./files/"
glp_files_path = "glp_files/"
beaconing_api_route = "https://core.beaconing.eu/api/"
templates = [
    "templates/index.html",
    "templates/unauthorised_user.html"
]
dist_folder = "../teacherui-frontend/public/dist"

[localisation]
map_file = "./trans.map"
key_file = "./trans.keys"

[debug]
grmon = false
```

### Configuration
By default the server requests to the API and scripts will be loaded from the external IP address.

To stop the use of the external IP address, and to make the callback link become 127.0.0.1, you must set `local` to `true`:

```toml
[server]
local = true
```

To provide a static URL enter one into the host variable under server without the trailing slash:

```toml
[server]
host = "example.com"
```

By default, the host will be prefixed with `https://`. If you wish to use `http://` instead, it's possible to add that to the host:
```toml
[server]
host = "http://example.com"
```

Changes to the host configuration will only take place once gin is running in Release Mode.
To change gin to Release mode the variable `GIN_MODE` must be exported with the value `release`:

bash
```bash
$ export GIN_MODE=release
```

fish
```bash
$ set -x GIN_MODE release
```

### Running
To run, simply execute the compiled binary file:

```bash
$ ./beaconing
```

The backend will now be running at `localhost:8080` by default.
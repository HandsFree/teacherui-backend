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

# teacherui-backend
Note, the frontend code can be found [here](//github.com/HandsFree/beaconing-teacher-ui).

## License
Licensed under GNU AGPLv3. See the `LICENSE.md` file for the full license.

# Development
## Prerequisites
- Yarn
- Go

## Installation
Cloning the repo should be done using Go:

```bash
$ go get github.com/handsfree/teacherui-backend
```

### Building
In the backend folder type:
```bash
$ go build -o beaconing
```

### Setting up
A configuration file must be created before running. The binary will look for the
configuration file in the same directory as the binary, e.g.

  folder/
    beaconing.exe
    cfg/
      config.toml

Below is an example configuration file:

config.toml
```toml
[auth]
id = "teacherui"
secret = "UrqTSjfnaWsaJHCTfGeU6YyEVNa3c2QzE8GrTLcoK1kljsNB3HrG6jXAGI6q8wKR"

[server]
local = true
host = ""
port = 8080
root_path = "./../frontend/public/"
glp_files_path = "./glp_files/"
beaconing_api_route = "https://core.beaconing.eu/api/"

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
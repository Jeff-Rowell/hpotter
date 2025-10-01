<html>
<body>
<div align="center">
  <h1>HPotter</h1>
  <br>
  <div style="display: flex; align-items: center; justify-content: center;">
      <img src="./static/honeypot.png" style="width: 256.5px; height: 315px; margin-left: -25px;">
      <img src="./static/gopher.svg" style="width: 211.5px; height: 270px; padding-top: 100px; padding-left: 25px;">
  </div>

</div>
<br>
<hr>
</body>
</html>

HPotter is a honney pot implementation that creates disposible containers upon each connection to the pot. The associated payloads sent to the honeypot are optionally stored to a local database for further inspection and analysis. Request and response data saving can be optionally turned on/off. So if you're only interested in keeping the request data, you can turn off the response data storage option. Additionally, the honey pot is configurable to run multiple services and protocols by specifying the container image to run in a `config.yaml` file as shown below:

<details>
  <summary>Click to expand example config...</summary>

```yaml
services:
  # HTTPd service - automatically runs on port 8080/tcp
  - service: "httpd"
    listen_address: "0.0.0.0"  # optional, defaults to localhost (127.0.0.1)
    request_save: true
    collect_credentials: true
    # TLS options (optional, only for httpd service)
    use_tls: true
    certificate_path: "/path/to/certificate.crt"
    key_path: "/path/to/private.key"
    # OR generate self-signed certificates automatically
    # generate_certs: true

  # SSH service - automatically runs on port 2222/tcp
  - service: "ssh"
    request_save: true
    collect_credentials: true
    command_limit: 10  # optional, defaults to 10 for ssh/telnet

  # Telnet service - automatically runs on port 2223/tcp
  - service: "telnet"
    request_save: true
    collect_credentials: true
    command_limit: 15  # optional, defaults to 10 for ssh/telnet
    # Required environment variables are automatically added

db_config:
  db_type: "postgres"
  user: "your_db_user"
  password: "your_db_password"
```

</details>

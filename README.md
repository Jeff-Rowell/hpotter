<div align="center">
    <h1>HPotter</h1>
</div>

HPotter is a honney pot implementation that creates disposible containers upon each connection to the pot. The associated payloads sent to the honeypot are optionally stored to a local database for further inspection and analysis. Request and response data saving can be optionally turned on/off. So if you're only interested in keeping the request data, you can turn off the response data storage option. Additionally, the honey pot is configurable to run multiple services and protocols by specifying the container image to run in a `config.json` file as shown below:

```json
{
  "services": [
    {
      "image_name": "cowrie/cowrie:latest",
      "listen_port": 2222,
      "listen_proto": "tcp",
      "protocol": "ssh",
      "request_save": true,
      "response_save": false,
      "socket_timeout": 10
    },
    {
      "image_name": "cowrie/cowrie:latest",
      "listen_port": 2223,
      "listen_proto": "tcp",
      "protocol": "telnet",
      "request_save": true,
      "response_save": true,
      "socket_timeout": 10,
      "envvars": [
        {
          "key": "COWRIE_TELNET_ENABLED",
          "value": "yes"
        }
      ]
    }
  ],
  "db_config": {
    "db_type": "postgres",
    "user": "",
    "password": "",
    "image_name": "postgres:17.6"
  }
}
```

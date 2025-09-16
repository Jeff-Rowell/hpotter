package credentials

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/Jeff-Rowell/hpotter/internal/database"
)

type HTTPDCredentialCollector struct {
	conn   net.Conn
	db     *database.Database
	dbConn *database.Connections
}

// NewHTTPCredentialCollector creates a new httpd credential collector
func NewHTTPDCredentialCollector(conn net.Conn, db *database.Database, dbConn *database.Connections) *HTTPDCredentialCollector {
	return &HTTPDCredentialCollector{
		conn:   conn,
		db:     db,
		dbConn: dbConn,
	}
}

// CollectCredentials presents a fake httpd login and collects credentials
func (t *HTTPDCredentialCollector) CollectCredentials() (*database.Credentials, error) {
	t.conn.SetDeadline(time.Now().Add(60 * time.Second))
	reader := bufio.NewReader(t.conn)
	request, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read request: %w", err)
	}

	parts := strings.Fields(request)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid HTTP request")
	}

	method := parts[0]

	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		if colonIndex := strings.Index(line, ":"); colonIndex > 0 {
			key := strings.TrimSpace(line[:colonIndex])
			value := strings.TrimSpace(line[colonIndex+1:])
			headers[strings.ToLower(key)] = value
		}
	}

	var username, password string

	if method == "POST" {
		contentLength := 0
		if cl, exists := headers["content-length"]; exists {
			fmt.Sscanf(cl, "%d", &contentLength)
		}

		if contentLength > 0 && contentLength < 1024 {
			body := make([]byte, contentLength)
			_, err := reader.Read(body)
			if err == nil {
				formData, err := url.ParseQuery(string(body))
				if err == nil {
					username = formData.Get("username")
					password = formData.Get("password")
				}
			}
		}

		t.sendLoginFailureResponse()
	} else {
		t.sendLoginPage()
	}

	if username != "" || password != "" {
		creds := &database.Credentials{
			Username:      username,
			Password:      password,
			ConnectionsID: t.dbConn.ID,
		}

		if err := t.db.Write(creds); err != nil {
			return nil, fmt.Errorf("failed to save credentials: %w", err)
		}

		return creds, nil
	}

	return nil, nil
}

func (t *HTTPDCredentialCollector) sendLoginPage() {
	loginHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Apache Login</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f5f5f5;
            margin: 0;
            padding: 20px;
        }
        .container {
            max-width: 400px;
            margin: 100px auto;
            background: white;
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        .header {
            text-align: center;
            margin-bottom: 30px;
        }
        .header h1 {
            color: #d73502;
            font-size: 24px;
            margin: 0;
        }
        .header p {
            color: #666;
            margin: 10px 0 0 0;
        }
        .form-group {
            margin-bottom: 20px;
        }
        label {
            display: block;
            margin-bottom: 5px;
            color: #333;
            font-weight: bold;
        }
        input[type="text"], input[type="password"] {
            width: 100%;
            padding: 12px;
            border: 1px solid #ddd;
            border-radius: 4px;
            font-size: 14px;
            box-sizing: border-box;
        }
        input[type="text"]:focus, input[type="password"]:focus {
            border-color: #d73502;
            outline: none;
        }
        .submit-btn {
            width: 100%;
            padding: 12px;
            background-color: #d73502;
            color: white;
            border: none;
            border-radius: 4px;
            font-size: 16px;
            cursor: pointer;
        }
        .submit-btn:hover {
            background-color: #b52d02;
        }
        .info {
            background-color: #f0f8ff;
            border-left: 4px solid #d73502;
            padding: 15px;
            margin-bottom: 20px;
            font-size: 14px;
            color: #333;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Apache Login</h1>
        </div>
        <form method="POST" action="/login">
            <div class="form-group">
                <label for="username">Username:</label>
                <input type="text" id="username" name="username" required>
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" id="password" name="password" required>
            </div>
            <button type="submit" class="submit-btn">Sign In</button>
        </form>
    </div>
</body>
</html>`

	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s", len(loginHTML), loginHTML)
	t.conn.Write([]byte(response))
}

func (t *HTTPDCredentialCollector) sendLoginFailureResponse() {
	errorHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Apache Login - Login Failed</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f5f5f5;
            margin: 0;
            padding: 20px;
        }
        .container {
            max-width: 400px;
            margin: 100px auto;
            background: white;
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        .header {
            text-align: center;
            margin-bottom: 30px;
        }
        .header h1 {
            color: #d73502;
            font-size: 24px;
            margin: 0;
        }
        .error {
            background-color: #ffe6e6;
            border-left: 4px solid #d73502;
            padding: 15px;
            margin-bottom: 20px;
            font-size: 14px;
            color: #d73502;
            font-weight: bold;
        }
        .retry-link {
            text-align: center;
        }
        .retry-link a {
            color: #d73502;
            text-decoration: none;
            font-weight: bold;
        }
        .retry-link a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Authentication Failed</h1>
        </div>
        <div class="error">
            Access denied.
        </div>
        <div class="retry-link">
            <a href="/">Try Again</a>
        </div>
    </div>
</body>
</html>`

	response := fmt.Sprintf("HTTP/1.1 401 Unauthorized\r\nContent-Type: text/html\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s", len(errorHTML), errorHTML)
	t.conn.Write([]byte(response))
}

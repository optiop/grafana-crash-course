import datetime

import jwt

from flask import (
    Flask,
    Response,
    request,
    render_template_string,
    redirect,
)

app = Flask(__name__)
app.config['SECRET_KEY'] = 'supersecretkey'

# Hardcoded users for demonstration
USERS = {
    "admin": "password",
    "user1": "pass1",
    "user2": "pass2"
}

# HTML UI (inline template)
LOGIN_PAGE = '''
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Login</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <style>
    body {
      margin: 0;
      font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
      background: linear-gradient(135deg, #667eea, #764ba2);
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh;
    }

    form {
      background: #fff;
      padding: 30px 25px;
      border-radius: 10px;
      box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
      width: 100%;
      max-width: 350px;
      box-sizing: border-box;
    }

    h2 {
      text-align: center;
      color: #333;
      margin-bottom: 25px;
    }

    input {
      width: 100%;
      padding: 12px 15px;
      margin-bottom: 20px;
      border: 1px solid #ddd;
      border-radius: 6px;
      box-sizing: border-box;
      font-size: 15px;
      transition: border-color 0.3s;
    }

    input:focus {
      border-color: #667eea;
      outline: none;
    }

    input[type="submit"] {
      background-color: #667eea;
      color: #fff;
      font-weight: bold;
      border: none;
      cursor: pointer;
      transition: background-color 0.3s;
    }

    input[type="submit"]:hover {
      background-color: #5a67d8;
    }

    .error {
      color: #e53e3e;
      text-align: center;
      font-size: 14px;
      margin-top: -10px;
    }

    @media (max-width: 400px) {
      form {
        padding: 20px 15px;
      }

      h2 {
        font-size: 22px;
      }
    }
  </style>
</head>
<body>
  <form method="POST">
    <h2>Login</h2>
    <input type="text" name="username" placeholder="Username" required>
    <input type="password" name="password" placeholder="Password" required>
    <input type="submit" value="Login">
    {% if error %}<p class="error">{{ error }}</p>{% endif %}
  </form>
</body>
</html>
'''

DASHBOARD_PAGE = '''
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Dashboard</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <style>
    body {
      margin: 0;
      padding: 40px 15px;
      font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
      background: linear-gradient(135deg, #667eea, #764ba2);
      display: flex;
      justify-content: center;
      align-items: center;
      min-height: 100vh;
      box-sizing: border-box;
    }

    .container {
      background: #fff;
      padding: 35px 30px;
      border-radius: 12px;
      box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
      max-width: 600px;
      width: 100%;
      box-sizing: border-box;
    }

    h2 {
      text-align: center;
      color: #333;
      margin-bottom: 25px;
      font-size: 26px;
    }

    p {
      font-size: 16px;
      color: #555;
      margin-bottom: 10px;
    }

    pre {
      background: #f0f4f8;
      padding: 15px;
      border-radius: 8px;
      overflow-x: auto;
      font-size: 14px;
      color: #333;
      line-height: 1.5;
    }

    .button-container {
      text-align: center;
      margin-top: 20px;
    }

    .button-container a {
      background-color: #667eea;
      color: #fff;
      padding: 10px 20px;
      text-decoration: none;
      border-radius: 6px;
      transition: background-color 0.3s;
    }

    .button-container a:hover {
      background-color: #5a67d8;
    }

    @media (max-width: 600px) {
      .container {
        padding: 25px 20px;
      }

      h2 {
        font-size: 22px;
      }
    }
  </style>
</head>
<body>
  <div class="container">
    <h2>Welcome {{ user }}</h2>
    <p>Your token:</p>
    <pre>{{ token }}</pre>
    <div class="button-container">
      <a href="http://localhost:8080/grafana/">Go to Grafana</a>
    </div>
  </div>
</body>
</html>
'''

@app.route('/', methods=['GET', 'POST'])
def login():
    error = None
    token = request.cookies.get('token')
    if token:
        try:
            data = jwt.decode(token, app.config['SECRET_KEY'], algorithms=['HS256'])
            return render_template_string(DASHBOARD_PAGE, user=data['user'], token=token)
        except jwt.ExpiredSignatureError:
            error = 'Session expired. Please log in again.'
        except jwt.InvalidTokenError:
            error = 'Invalid session. Please log in again.'

    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        if username in USERS and USERS[username] == password:
            token = jwt.encode({
                'user': username,
                'exp': datetime.datetime.utcnow() + datetime.timedelta(minutes=30)
            }, app.config['SECRET_KEY'], algorithm='HS256')
            resp = Response(render_template_string(DASHBOARD_PAGE, user=username, token=token))
            resp.set_cookie('token', token, domain='localhost')
            return resp
        else:
            error = 'Invalid Credentials'
    return render_template_string(LOGIN_PAGE, error=error)


@app.route('/access')
def go_to_grafana():
    token = request.cookies.get('token')
    if not token:
        return "Token required", 401

    try:
        data = jwt.decode(token, app.config['SECRET_KEY'], algorithms=['HS256'])
        username = data['user']

        resp = Response("Authorized", status=200)
        resp.headers['X-WEBAUTH-USER'] = username
        return resp
    except jwt.ExpiredSignatureError:
        return "Token expired", 403
    except jwt.InvalidTokenError:
        return "Invalid token", 403


@app.route('/logout')
def logout():
    token = request.cookies.get('token')
    if not token:
        return redirect('/')

    try:
        _ = jwt.decode(token, app.config['SECRET_KEY'], algorithms=['HS256'])
        resp = redirect('/')
        resp.delete_cookie('token')
        return resp
    except jwt.ExpiredSignatureError:
        return redirect('/')
    except jwt.InvalidTokenError:
        return redirect('/')


if __name__ == '__main__':
    app.run(debug=True)

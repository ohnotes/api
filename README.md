## ❗ Important
[Frontend repository](https://github.com/ohnotes/ohnotes)

<br><br>
## 📖 Summary
[Changelog](#changelog)<br>
[FAQ](#faq)<br>
[Setup](#setup)

<br><a name="changelog"></a>
## ♻️ Changelog
- Fixed functionality of 'destructive' messages.
- Little fixes around the project.
- Fixes on [Frontend](https://github.com/ohnotes/ohnotes)

<br><a name="faq"></a>
## ❓ FAQ
#### I've founded a bug, how to report?
Contact me on z3ox1s@protonmail.com, please :).

<br><a name="setup"></a>
## 🔧 Setup
### Clone this repository:
```bash
    git clone https://github.com/ohnotes/api
    cd api
```

### Install all dependencies:
```bash
go mod tidy
```

### Setting all up:
```bash
export PORT=<PORT> # Set the port to run the API
export MONGO=<URI of your database (must be MongoDB)> # URI to connect to database
export SECRET=<SECRET> # Secret to generate JWTs
```

### Run:
```bash
go run .
```

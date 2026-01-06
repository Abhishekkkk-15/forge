# Forge 🔥

Forge is a **developer-friendly CLI tool** that scaffolds backend projects instantly using **production-ready templates**.

It helps you generate boilerplate for popular backend stacks like **Express (TypeScript)** and **Go APIs** with a single command—so you can focus on building features, not setup.

---

## ✨ Features

- 🚀 Scaffold backend projects in seconds  
- 🧱 Production-ready templates (not toy examples)  
- 🏷️ Metadata-driven templates using `template.json`  
- 🧩 Supports dynamic flags (`--port`, `--docker`, etc.)  
- 🌍 Works globally as a CLI tool  
- 🔧 Easy to extend with custom templates  

---

## 📦 Available Templates

| Template        | Description                     |
|-----------------|---------------------------------|
| `express`       | Express API      |
| `express-ts`    | Express API with TypeScript     |
| `go-api`        | Go API using `net/http`         |
| `go-gin-api`    | Go API using the Gin framework  |

---

## 🚀 Installation

### Build from Source

```bash
git clone https://github.com/<your-username>/forge.git
cd forge
go build -o forge
```

### Add to Path
- mac/linux
```bash
sudo mv forge /usr/local/bin/
```
- windows
```bash
move forge.exe C:\Users\<your-username>\bin
```



## Usage/Examples
- List Available Templates
```bash
forge list
```
- Get Template Information
```bash
forge info express
```
- Create an Express + TypeScript API
``` bash
forge init express-ts my-api --port 8080 --docker
```
- Create Go APIs
```bash
forge init go-api users --port 9000
forge init go-gin-api users --port 9001
```

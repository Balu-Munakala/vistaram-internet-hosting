# 🌐 Same API, 4 Languages, Hosted on the Internet

Companion repository for the YouTube video:

> **"Host Your API Live on the Internet with a Free Domain"**

This repo proves one idea: **hosting is language-independent.**

The same 3-endpoint REST API is built in **four different programming languages** — Python, JavaScript, Go, and Java — plus a **static HTML frontend** that consumes it. Same routes, same JSON shape, same behavior. Only the build and start commands change.

If you can host one of these, you can host **all four** — and by extension, almost any backend you ever write.

---

## 📁 Project Structure

@@@text
.
├── python/
│   └── server.py              # Flask server (port 3000)
├── node/
│   ├── server.js              # Express server (port 3001)
│   └── package.json
├── go/
│   └── main.go                # net/http server (port 3002)
├── java/
│   ├── ServerApplication.java # Spring Boot (port 3003)
│   └── pom.xml
└── frontend/
    └── index.html             # Static HTML frontend
@@@

---

## 🔌 API Endpoints (identical in all 4 languages)

| Method | Endpoint        | Description                     |
|--------|-----------------|---------------------------------|
| GET    | `/products`     | Returns the full product list   |
| GET    | `/products/:id` | Returns a single product by ID  |
| POST   | `/products`     | Adds a new product to the store |

Every response follows the same shape:

@@@json
{
  "status": "success",
  "count": 5,
  "serverTime": "2026-10-20T10:30:00Z",
  "data": [
    {
      "id": 1,
      "name": "Wireless Headphones",
      "brand": "SoundMax",
      "price": 1999,
      "currency": "INR",
      "rating": 4.6,
      "reviews": 1284,
      "inStock": true,
      "image": "https://picsum.photos/seed/headphones/400/400",
      "tags": ["audio", "bluetooth", "bestseller"]
    }
  ]
}
@@@

---

## ⚙️ Prerequisites

Install the runtime for the language(s) you want to run:

| Language   | Runtime / Tool       | Check with         |
|------------|----------------------|--------------------|
| Python     | Python 3.9+          | `python --version` |
| JavaScript | Node.js 18+          | `node --version`   |
| Go         | Go 1.20+             | `go version`       |
| Java       | JDK 17+ & Maven 3.8+ | `mvn -version`     |

The HTML frontend has **no extra installation** — it's served with Python's built-in HTTP server, so the only thing you need is Python (which you probably already have if you're running the Python backend).

---

## 🚀 Quick Start

### Run any one backend

@@@bash
# 🐍 Python — port 3000
cd python && pip install flask && python server.py

# 🟨 Node — port 3001
cd node && npm install && node server.js

# 🐹 Go — port 3002
cd go && go run main.go

# ☕ Java — port 3003
cd java && mvn clean install && mvn spring-boot:run
@@@

### Run the frontend

@@@bash
cd frontend
python -m http.server 5500
@@@

### Open in browser

@@@text
Python    → http://localhost:3000/products
Node      → http://localhost:3001/products
Go        → http://localhost:3002/products
Java      → http://localhost:3003/products
Frontend  → http://localhost:5500
@@@

---

## 🐍 1. Python (Flask) — port `3000`

**Install & run:**

@@@bash
cd python
pip install flask
python server.py
@@@

**Test:**

@@@bash
curl http://localhost:3000/products
curl http://localhost:3000/products/3
@@@

Open → [http://localhost:3000/products](http://localhost:3000/products)

---

## 🟨 2. JavaScript (Node + Express) — port `3001`

**Install & run:**

@@@bash
cd node
npm install
node server.js
@@@

**Test:**

@@@bash
curl http://localhost:3001/products
curl http://localhost:3001/products/5
@@@

Open → [http://localhost:3001/products](http://localhost:3001/products)

---

## 🐹 3. Go (stdlib `net/http`) — port `3002`

**Run (no packages to install — pure standard library):**

@@@bash
cd go
go run main.go
@@@

**Test:**

@@@bash
curl http://localhost:3002/products
curl http://localhost:3002/products/2
@@@

Open → [http://localhost:3002/products](http://localhost:3002/products)

> 🎉 Go needs **zero third-party packages** — everything comes from the standard library.

---

## ☕ 4. Java (Spring Boot) — port `3003`

**Build & run:**

@@@bash
cd java
mvn clean install
mvn spring-boot:run
@@@

**Test:**

@@@bash
curl http://localhost:3003/products
curl http://localhost:3003/products/1
@@@

Open → [http://localhost:3003/products](http://localhost:3003/products)

> ⏳ The first run downloads Spring Boot dependencies — takes about 30 seconds.

---

## 🌐 5. Frontend (Static HTML) — port `5500`

A simple `index.html` that fetches products from one of the four backends and renders them as cards.

**Serve it with Python's built-in HTTP server:**

@@@bash
cd frontend
python -m http.server 5500
@@@

Open → [http://localhost:5500](http://localhost:5500)

> ⚠️ **Important:** Because each backend runs on a different port, the frontend
> needs to know which one to call. Open `frontend/index.html` and update the
> `API_BASE` variable at the top of the `<script>` tag:
>
> @@@javascript
> const API_BASE = "http://localhost:3000";  // 🐍 Python
> // const API_BASE = "http://localhost:3001";  // 🟨 Node
> // const API_BASE = "http://localhost:3002";  // 🐹 Go
> // const API_BASE = "http://localhost:3003";  // ☕ Java
> @@@

---

## 🧪 Test with `curl` (POST example)

Any of the four servers responds to the same POST — just change the port:

@@@bash
curl -X POST http://localhost:3000/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bluetooth Speaker",
    "brand": "BoomBox",
    "price": 1499,
    "tags": ["audio", "portable"]
  }'
@@@

Response:

@@@json
{
  "status": "created",
  "data": {
    "id": 6,
    "name": "Bluetooth Speaker",
    "brand": "BoomBox",
    "price": 1499,
    "currency": "INR",
    "rating": 0.0,
    "reviews": 0,
    "inStock": true,
    "image": "https://picsum.photos/seed/BluetoothSpeaker/400/400",
    "tags": ["audio", "portable"]
  }
}
@@@

Then refresh `/products` — you'll see **6 products** instead of 5. ✅

---

## 🎬 Run All 4 Servers Side by Side

Because each server uses a **different port**, you can run all four at once
and compare the responses in separate browser tabs:

| Language   | URL                              |
|------------|----------------------------------|
| Python     | http://localhost:3000/products   |
| JavaScript | http://localhost:3001/products   |
| Go         | http://localhost:3002/products   |
| Java       | http://localhost:3003/products   |

**You will see identical JSON in all four tabs.**

This is the key takeaway of the video: **the client does not care what
language the server is written in.** Swap one out tomorrow — nothing breaks.

---

## ☁️ Hosting on the Internet (from the video)

The video walks through taking one of these servers from `localhost` to a
**live public URL with a free domain**. The key changes for deployment:

| Language   | Build Command           | Start Command                    | Port Binding                                              |
|------------|-------------------------|----------------------------------|-----------------------------------------------------------|
| Python     | `pip install flask`     | `python server.py`               | `app.run(host="0.0.0.0", port=int(os.environ.get("PORT", 3000)))` |
| JavaScript | `npm install`           | `node server.js`                 | `app.listen(process.env.PORT \|\| 3001)`                  |
| Go         | `go build -o server .`  | `./server`                       | `os.Getenv("PORT")`                                       |
| Java       | `mvn clean install`     | `java -jar target/*.jar`         | `server.port=${PORT:3003}` in `application.properties`     |

> 🔑 **The two golden rules of deployment:**
>
> 1. Bind to **`0.0.0.0`**, not `127.0.0.1` — this tells the server to accept
>    traffic from all network interfaces, not just localhost.
> 2. Read the port from the **`PORT` environment variable** — the host assigns
>    it dynamically and your app must not hardcode it.
>
> Get these two right and 90% of deploy failures disappear.

---

## 📝 Notes

- **Data is in-memory** — products reset when the server restarts.
  This is intentional to keep the demo simple. Add a database later if you want persistence.
- **The HTML frontend has no build step** — plain HTML + JS, served with
  `python -m http.server`. No npm, no bundler, no framework.
- **Images use [`picsum.photos`](https://picsum.photos)** — free placeholder
  images, no API key needed, and each product gets a unique photo.
- **CORS is enabled** in all four backends — so the frontend on `:5500` can
  call any backend on `:3000–3003` without browser blocks.

---

## 🤔 Why Four Languages?

Because this video is about **hosting**, not about coding. Showing four
implementations of the same API proves that **the hosting workflow is
language-independent** — the concepts (routes, ports, `localhost`,
public IP, DNS, `0.0.0.0`, `PORT`) are identical whether your server is
written in Python, Node, Go, or Java.

Learn the workflow once. Apply it everywhere.

---

## 🎥 Watch the Video

The full walkthrough — from `localhost` to a live public URL with a free
domain — is on YouTube: *(link will be added soon)*

---

## 📜 License

MIT — fork it, learn from it, break it, ship it.

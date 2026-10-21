const express = require("express");
const app = express();
app.use(express.json());

let products = [
  { id: 1, name: "Wireless Headphones", brand: "SoundMax", price: 1999, currency: "INR", rating: 4.6, reviews: 1284, inStock: true, image: "https://picsum.photos/seed/headphones/400/400", tags: ["audio","bluetooth","bestseller"] },
  { id: 2, name: "Smart Watch Pro", brand: "TimeX", price: 3499, currency: "INR", rating: 4.4, reviews: 892, inStock: true, image: "https://picsum.photos/seed/watch/400/400", tags: ["wearable","fitness"] },
  { id: 3, name: "Gaming Mouse RGB", brand: "ClickZ", price: 899, currency: "INR", rating: 4.8, reviews: 2103, inStock: false, image: "https://picsum.photos/seed/mouse/400/400", tags: ["gaming","rgb"] },
  { id: 4, name: "Mechanical Keyboard", brand: "KeyForge", price: 2599, currency: "INR", rating: 4.7, reviews: 1567, inStock: true, image: "https://picsum.photos/seed/keyboard/400/400", tags: ["gaming","typing"] },
  { id: 5, name: "4K Webcam", brand: "ClearView", price: 4299, currency: "INR", rating: 4.3, reviews: 654, inStock: true, image: "https://picsum.photos/seed/webcam/400/400", tags: ["streaming","video"] },
];

app.get("/products", (req, res) => {
  res.json({ status: "success", count: products.length, serverTime: new Date().toISOString(), data: products });
});

app.get("/products/:id", (req, res) => {
  const p = products.find(x => x.id === parseInt(req.params.id));
  if (!p) return res.status(404).json({ status: "error", message: `Product ${req.params.id} not found` });
  res.json({ status: "success", data: p });
});

app.post("/products", (req, res) => {
  const body = req.body;
  const newP = {
    id: products.length + 1,
    name: body.name,
    brand: body.brand || "Generic",
    price: body.price,
    currency: "INR",
    rating: 0.0, reviews: 0, inStock: true,
    image: `https://picsum.photos/seed/${body.name.replace(/\s/g,"")}/400/400`,
    tags: body.tags || [],
  };
  products.push(newP);
  res.status(201).json({ status: "created", data: newP });
});

app.listen(3001, () => console.log("http://localhost:3000/products"));
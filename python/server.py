from flask import Flask, jsonify, request
from datetime import datetime

app = Flask(__name__)

products = [
    {"id": 1, "name": "Wireless Headphones", "brand": "SoundMax", "price": 1999, "currency": "INR",
     "rating": 4.6, "reviews": 1284, "inStock": True,
     "image": "https://picsum.photos/seed/headphones/400/400",
     "tags": ["audio", "bluetooth", "bestseller"]},
    {"id": 2, "name": "Smart Watch Pro", "brand": "TimeX", "price": 3499, "currency": "INR",
     "rating": 4.4, "reviews": 892, "inStock": True,
     "image": "https://picsum.photos/seed/watch/400/400",
     "tags": ["wearable", "fitness"]},
    {"id": 3, "name": "Gaming Mouse RGB", "brand": "ClickZ", "price": 899, "currency": "INR",
     "rating": 4.8, "reviews": 2103, "inStock": False,
     "image": "https://picsum.photos/seed/mouse/400/400",
     "tags": ["gaming", "rgb"]},
    {"id": 4, "name": "Mechanical Keyboard", "brand": "KeyForge", "price": 2599, "currency": "INR",
     "rating": 4.7, "reviews": 1567, "inStock": True,
     "image": "https://picsum.photos/seed/keyboard/400/400",
     "tags": ["gaming", "typing"]},
    {"id": 5, "name": "4K Webcam", "brand": "ClearView", "price": 4299, "currency": "INR",
     "rating": 4.3, "reviews": 654, "inStock": True,
     "image": "https://picsum.photos/seed/webcam/400/400",
     "tags": ["streaming", "video"]},
]

@app.route("/products")
def get_products():
    return jsonify({
        "status": "success",
        "count": len(products),
        "serverTime": datetime.now().isoformat(),
        "data": products,
    })

@app.route("/products/<int:pid>")
def get_product(pid):
    p = next((x for x in products if x["id"] == pid), None)
    if not p:
        return jsonify({"status": "error", "message": f"Product {pid} not found"}), 404
    return jsonify({"status": "success", "data": p})

@app.route("/products", methods=["POST"])
def add_product():
    body = request.get_json()
    new = {
        "id": len(products) + 1,
        "name": body["name"],
        "brand": body.get("brand", "Generic"),
        "price": body["price"],
        "currency": "INR",
        "rating": 0.0,
        "reviews": 0,
        "inStock": True,
        "image": f"https://picsum.photos/seed/{body['name'].replace(' ', '')}/400/400",
        "tags": body.get("tags", []),
    }
    products.append(new)
    return jsonify({"status": "created", "data": new}), 201

if __name__ == "__main__":
    app.run(port=3000)
package com.example.server;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;
import java.time.Instant;
import java.util.*;

@SpringBootApplication
@RestController
public class ServerApplication {

    public static class Product {
        public int id;
        public String name;
        public String brand;
        public int price;
        public String currency = "INR";
        public double rating;
        public int reviews;
        public boolean inStock;
        public String image;
        public List<String> tags;
        public Product() {}
        public Product(int id, String name, String brand, int price, double rating, int reviews, boolean inStock, String image, List<String> tags) {
            this.id = id; this.name = name; this.brand = brand; this.price = price;
            this.rating = rating; this.reviews = reviews; this.inStock = inStock;
            this.image = image; this.tags = tags;
        }
    }

    private List<Product> products = new ArrayList<>(List.of(
        new Product(1, "Wireless Headphones", "SoundMax", 1999, 4.6, 1284, true, "https://picsum.photos/seed/headphones/400/400", List.of("audio","bluetooth","bestseller")),
        new Product(2, "Smart Watch Pro", "TimeX", 3499, 4.4, 892, true, "https://picsum.photos/seed/watch/400/400", List.of("wearable","fitness")),
        new Product(3, "Gaming Mouse RGB", "ClickZ", 899, 4.8, 2103, false, "https://picsum.photos/seed/mouse/400/400", List.of("gaming","rgb")),
        new Product(4, "Mechanical Keyboard", "KeyForge", 2599, 4.7, 1567, true, "https://picsum.photos/seed/keyboard/400/400", List.of("gaming","typing")),
        new Product(5, "4K Webcam", "ClearView", 4299, 4.3, 654, true, "https://picsum.photos/seed/webcam/400/400", List.of("streaming","video"))
    ));

    public static void main(String[] args) {
        SpringApplication.run(ServerApplication.class, args);
    }

    @GetMapping("/products")
    public Map<String, Object> getAll() {
        return Map.of("status", "success", "count", products.size(),
                      "serverTime", Instant.now().toString(), "data", products);
    }

    @GetMapping("/products/{id}")
    public Map<String, Object> getOne(@PathVariable int id) {
        return products.stream().filter(p -> p.id == id).findFirst()
            .map(p -> Map.of("status", "success", "data", (Object) p))
            .orElse(Map.of("status", "error", "message", "Product " + id + " not found"));
    }

    @PostMapping("/products")
    public Map<String, Object> add(@RequestBody Product p) {
        p.id = products.size() + 1;
        p.currency = "INR";
        p.inStock = true;
        p.image = "https://picsum.photos/seed/" + p.name.replace(" ", "") + "/400/400";
        products.add(p);
        return Map.of("status", "created", "data", p);
    }
}
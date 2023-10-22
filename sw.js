// Establish a cache name
const cacheName = 'dandelion_cache_v1';

const contentToCache = [
    "/index.html",
    "/main.js",
    "/favicon.ico",
  ];

self.addEventListener('install', (e) => {
    e.waitUntil(
        (async () => {
          const cache = await caches.open(cacheName);
          console.log("[Service Worker] Caching all: app shell and content", contentToCache);
          await cache.addAll(contentToCache);
        })(),
      );
});

self.addEventListener("fetch", (e) => {
    e.respondWith(
      (async () => {
        const r = await caches.match(e.request);
        console.log(`[Service Worker] Fetching resource: ${e.request.url}`);
        if (r) {
          return r;
        }
        try {
            const response = await fetch(e.request);
        } catch(e) {
            console.warn(e);
        }
        const cache = await caches.open(cacheName);
        console.log(`[Service Worker] Caching new resource: ${e.request.url}`);
        cache.put(e.request, response.clone());
        return response;
      })(),
    );
  });

  addEventListener("message", async (e) => {
    console.log(`Message received: ${e.data}`);

    switch (e.data.type) {
        case "updateMain": 
        const cache = await caches.open(cacheName);
        console.log(`[Service Worker] Caching new resource: main.js`);
        cache.put("/main.js", new Response());
        break;
    }
  });
  
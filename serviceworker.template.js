
const cacheName = 'deanishe-{{ .Version }}',
    precacheFiles = {{ .CacheFiles }};

self.addEventListener('install', function(event) {
  console.debug('[worker] installing ...');
  event.waitUntil(preCache());
  self.skipWaiting();
  console.debug('[worker] installed');
});

self.addEventListener('fetch', function(event) {
  if (event.request.method !== 'GET') return;

  var url = new URL(event.request.url);
  // console.debug(`[worker] pathname=${url.pathname}`);

  // return versioned assets from cache
  if (url.pathname.match(/\/js\/.+\.js$/) ||
      url.pathname.match(/\/style\/.+\.css$/)) {
    event.respondWith(preferCache(event.request));
    return;
  }

  // return other assets from cache and update cache
  if (url.pathname.match(/\/$/) ||
      url.pathname.match(/\.(map|png|ico|jpg|jpeg|gif|svg|xml|eot|woff|woff2)$/)) {
    event.respondWith(fromCache(event.request)
      .then(function(response) {
        console.debug(`[worker] cache > ${event.request.url}`);
        event.waitUntil(update(event.request));
        console.debug(`[worker] cache < ${event.request.url}`);
        return response;
      })
      .catch(function() {
        console.debug(`[worker] network > ${event.request.url}`);
        console.debug(`[worker] cache < ${event.request.url}`);
        return update(event.request);
      })
    );
    return;
  }

  // default to network, falling back to cached version
  event.respondWith(preferNetwork(event.request));
});

self.addEventListener('activate', function(event) {
  console.debug('[worker] activating ...');

  event.waitUntil(caches.keys()
    .then(function (keys) {
      return Promise.all(
        keys
          .filter(function(key) {
            return key !== cacheName;
          })
          .map(function(key) {
            console.debug(`[worker] deleting cache ${key} ...`);
            return caches.delete(key);
          })
      );
    })
    .then(self.clients.claim())
    .then(function() {
      console.log('[worker] activated');
    })
  );
});

// Populate cache
function preCache() {
  return caches.open(cacheName).then(function(cache) {
    return cache.addAll(precacheFiles);
  });
}

// Return cached response
function fromCache(request) {
  return caches.open(cacheName).then(function(cache) {
    return cache.match(request).then(function(response) {
      return response || Promise.reject('cache miss');
    });
  });
}

// Return and cache network response
function update(request) {
  return fetch(request).then(function(response) {
    caches.open(cacheName).then(function(cache) {
      cache.put(request, response);
    });
    return response.clone();
  });
}

// Return from cache, falling back to network
function preferCache(request) {
  return fromCache(request)
    .then(function(response) {
      console.debug(`[worker] cache > ${request.url}`);
      return response;
    })
    .catch(function() {
      console.debug(`[worker] network > ${request.url}`);
      console.debug(`[worker] cache < ${request.url}`);
      return update(request);
    });
}

// Return from network, falling back to cache
function preferNetwork(request) {
  return update(request)
    .then(function(response) {
      console.debug(`[worker] network > ${request.url}`);
      console.debug(`[worker] cache < ${request.url}`);
      return response;
    })
    .catch(function() {
      console.debug(`[worker] cache > ${request.url}`);
      return fromCache(request);
    });
}

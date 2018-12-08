
// From https://css-tricks.com/serviceworker-for-offline/

var version = '{{ .Hash }}';
var precacheFiles = {{ .URLs }};

var cacheName = 'deanishe-' + version.slice(0, 10);

self.addEventListener('install', function(event) {
  console.debug('[worker] installing ...');
  event.waitUntil(
    /* The caches built-in is a promise-based API that helps you cache responses,
       as well as finding and deleting them.
    */
    caches
      /* You can open a cache by name, and this method returns a promise. We use
         a versioned cache name here so that we can remove old cache entries in
         one fell swoop later, when phasing out an older service worker.
      */
      .open(cacheName)
      .then(function(cache) {
        /* After the cache is opened, we can fill it with the offline fundamentals.
           The method below will add all resources we've indicated to the cache,
           after making HTTP requests for each of them.
        */
        var busts = [];
        precacheFiles.forEach(function(p) {
          console.debug('[worker] caching', p, '...');
          busts.push(p + '?v=' + version.slice(0, 10));
        });
        // for (var i = 0; i < precacheFiles.length; i++) {
        //   console.debug('[worker] caching', precacheFiles[i], '...');
        // }
        return cache.addAll(busts);
      })
      .then(function() {
        console.debug('[worker] installed');
      })
  );
});


self.addEventListener('fetch', function(event) {
  // console.debug('[worker] fetch event started ...');
  if (event.request.method !== 'GET') {
    return;
  }

  var url = new URL(event.request.url);
  // console.debug('[worker] url =', url.href, ', path =', url.pathname, ', host =', url.host);

  if (url.hostname !== 'www.kulturliste-duesseldorf.de' &&
      url.hostname !== 'beta.kulturliste-duesseldorf.de' &&
      url.hostname !== 'ng.kulturliste-duesseldorf.de' &&
      url.hostname !== '127.0.0.1') {
    console.debug('[worker] ignoring 3rd-party URL', url.href);
    return;
  }

  if (url.pathname.startsWith('/api/') ||
      url.pathname.startsWith('/mailgroups/') ||
      url.pathname.startsWith('/a/')) {
    console.debug('[worker] ignoring service URL', url.href);
    return;
  }

  // if (url.hash) {
  //   console.debug('[worker] ignoring URL fragment', url.href);
  // }

  // CSS and JS files were cached on install
  if (url.pathname.match(/\.(css|js)$/)) {
    console.debug('[worker] static asset', url.pathname);
    event.respondWith(
      caches.open(cacheName).then(function(cache) {
        return cache.match(event.request).then(function(response) {
          if (response) {
            console.debug('[worker] from cache', url.pathname);
            return response;
          }

          return fetch(event.request).then(function(response) {
            cache.put(event.request, response.clone());
            console.debug('[worker] cached', url.pathname);
            return response;
          });
        });
      })
    );
    return;
  }

  // Return other assets from cache and update cache from network
  if (url.pathname.match(/\.(png|ico|jpg|jpeg|gif|svg)$/)) {
    console.debug('[worker] other asset', url.pathname);
    event.respondWith(
      caches.open(cacheName).then(function(cache) {
          return cache.match(event.request).then(function(response) {
            var p = fetch(event.request).then(function(response) {
              cache.put(event.request, response.clone());
              console.debug('[worker] cached', url.pathname);
              return response;
            })
            .catch(function() {
              console.error('[worker] error fetching', event.request.url);
            });

            if (response) {
              console.debug('[worker] from cache', url.pathname);
              return response;
            }
            return p;
          });
      })
    );
    return;
  }

  // Try to fetch everything else from the network first
  console.debug('[worker] network-first', event.request.url);
  event.respondWith(
    fetch(event.request)
    .then(function(response) {
      var r2 = response.clone();
      console.debug('[worker] from network', event.request.url);
      caches.open(cacheName).then(function(cache) {
        cache.put(event.request, r2);
        console.debug('[worker] cached', event.request.url);
      });
      return response;
    })
    .catch(function() {
      console.debug('[worker] falling back to cache', event.request.url, '...');
      return caches.match(event.request).then(function(response) {
        console.debug('[worker] from cache', event.request.url);
        return response;
      });
    })
  );
  return;
});

self.addEventListener('activate', function(event) {
  /* Just like with the install event, event.waitUntil blocks activate on a promise.
     Activation will fail unless the promise is fulfilled.
  */
  console.debug('[worker] activating ...');

  event.waitUntil(
    caches
      /* This method returns a promise which will resolve to an array of available
         cache keys.
      */
      .keys()
      .then(function (keys) {
        // We return a promise that settles when all outdated caches are deleted.
        return Promise.all(
          keys
            .filter(function (key) {
              // Filter by keys that don't start with the latest version prefix.
              return key !== cacheName;
              // return !key.startsWith(version);
            })
            .map(function (key) {
              /* Return a promise that's fulfilled
                 when each outdated cache is deleted.
              */
              console.debug('[worker] deleting old cache', key, '...');
              return caches.delete(key);
            })
        );
      })
      .then(function() {
        console.debug('[worker] activated');
      })
  );
});
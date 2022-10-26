---
title: "Get RSS feed URL for iTunes Podcast ID"
date: "2022-10-26T11:09:46+02:00"
draft: false
tags: ["rss", "itunes", "podcasts"]
---

For when you want the RSS feed, but they've only got "Subscribe in XYZ…" links.

<!--more-->

Enter the URL of an iTunes podcast preview page (or the podcast's iTunes ID) and get the URL of the podcast's actual feed back.

<style>
.itunes-icon svg {
  height: 28px;
  width: 28px;
  fill: #9C2DD0;
}

#itunes2rss {
  margin: 0.6rem 0;
  padding: 0.3rem;
  max-width: 400px;
}
#itunes2rss .notification {
  text-align: center;
  font-weight: bold;
  padding: 1.2rem 0;
}
#itunes2rss .podcast-form > div + div {
  margin-top: 0.6rem;
}
#itunes2rss .podcast-form > div {
  display: flex;
  justify-content: space-between;
}
#itunes2rss .podcast-form > div > * {
  flex: initial;
}
#itunes2rss .podcast-form > div .podcast-id, #itunes2rss .podcast-form > div .feed-url {
  flex: auto;
}
#itunes2rss .podcast-form > div button {
  margin-left: 0.6rem;
}
#itunes2rss .podcast-form > div button svg {
  height: 16px;
  width: 16px;
}
</style>

<h2>
    <span class="itunes-icon"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><defs><style>.fa-secondary{opacity:.4}</style></defs><path d="M224 0A223.88 223.88.0 0 0 0 224c0 90 52.6 165.65 125.74 201.41a6 6 0 0 0 8.53-6.31c-2.38-15.51-4.34-31-5.4-44.34a6 6 0 0 0-2.68-4.51A176 176 0 0 1 48 222.9c.59-96.24 79.29-174.65 175.53-174.9C320.79 47.75 4e2 126.8 4e2 224a176 176 0 0 1-80.65 147.87c-1 14-3.07 30.59-5.62 47.23a6 6 0 0 0 8.53 6.31C395.23 389.73 448 314.19 448 224A223.89 223.89.0 0 0 224 0zm98.45 325A143.63 143.63.0 0 0 368 216.43c-1.86-76.21-63.6-138.21-139.8-140.37C146.87 73.75 80 139.21 80 220a143.62 143.62.0 0 0 45.55 105 6 6 0 0 0 9.45-1.9 66.57 66.57.0 0 1 21.24-25.36 6 6 0 0 0 .63-9.19 96 96 0 1 1 134.26.0 6 6 0 0 0 .63 9.19A66.57 66.57.0 0 1 313 323.1a6 6 0 0 0 9.45 1.9z" class="fa-secondary"></path><path d="M224 312c-32.86.0-64 8.59-64 43.75.0 33.15 12.93 104.38 20.57 132.81 5.14 19 24.57 23.44 43.43 23.44s38.29-4.43 43.43-23.44c7.7-28.63 20.57-99.86 20.57-132.81.0-35.16-31.14-43.75-64-43.75zm0-24a64 64 0 1 0-64-64 64 64 0 0 0 64 64z" class="fa-primary"></path></svg>
    <span>iTunes &#x2192; RSS Feed</span>
</h2>

<div id="itunes2rss">
  <form class="podcast-form">
    <div>
      <label for="podcast-id">Enter URL or ID of an iTunes podcast here:</label>
    </div>
    <div>
      <input id="podcast-id" class="podcast-id" placeholder="Podcast ID"
      size="10" minlength="1" required/>
      <button class="find-button">Find RSS Feed</button>
    </div>
    <div>
      <input type="text" class="feed-url" readonly
        placeholder="URL of RSS feed"/>
      <button class="copy-button">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512"><defs><style>.fa-secondary{opacity:.4}</style></defs><path d="M336 63h-80v1a64 64 0 0 1 64 64H64a64 64 0 0 1 64-64v-1H48a48 48 0 0 0-48 48v352a48 48 0 0 0 48 48h288a48 48 0 0 0 48-48V111a48 48 0 0 0-48-48z" class="fa-secondary"></path><path d="M256 64a64 64 0 0 0-128 0 64 64 0 0 0-64 64h256a64 64 0 0 0-64-64zm-64 24a24 24 0 1 1 24-24 23.94 23.94.0 0 1-24 24z" class="fa-primary"></path></svg>
      </button>
    </div>
  </form>
  <div class="notification"></div>
</div>

  <script src="_itunes.js"></script>
  <script>
    iTunes2Rss('#itunes2rss');
  </script>
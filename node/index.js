const express = require('express');
const fetch = require('node-fetch');
const bcrypt = require('bcrypt');

// I also tried with raw http. Express makes no noticeable difference.
const app = express();

app.get('/node/api-calls', async (req, res) => {
    await fetch('some-url-you-can-put-load-on');
    res.send('OK');
});

app.get('/node/hash', async (req, res) => {
   res.json({ hash: await bcrypt.hash('password', 12) });  
});

app.get('/node/mix', async (req, res) => {
  if (Math.random() < 0.01) {
    await bcrypt.hash('password', 12);
  } else {
    await fetch('some-url-you-can-put-load-on');
  }
  res.send('OK');
});

app.listen(9001, () => {
    console.log('App listening on port 9001');
});

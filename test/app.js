const express = require('express');
const app = express();
const port = 3000;

app.use(express.json());

app.get('/', (req, res) => {
  res.json({ message: 'Hello World!' });
});

app.get('/users', (req, res) => {
  const users = [
    { id: 1, name: 'John' },
    { id: 2, name: 'Jane' }
  ];
  res.json(users);
});

app.post('/users', (req, res) => {
  const user = req.body;
  console.log('New user:', user);
  res.status(201).json(user);
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
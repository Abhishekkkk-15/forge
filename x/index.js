const express = require("express");
const app = express();

app.get("/", (req, res) => {
  res.send("Hello from x");
});

app.listen(3000, () => {
  console.log("x running on port 3000");
});

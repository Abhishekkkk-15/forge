import { Router } from "express";

const router = Router();

router.get("/", (_req, res) => {
  res.json({
    status: "ok",
    service: "x",
    timestamp: new Date().toISOString(),
  });
});

export default router;
